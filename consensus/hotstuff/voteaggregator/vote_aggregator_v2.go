package voteaggregator

import (
	"fmt"

	"github.com/rs/zerolog"

	"github.com/onflow/flow-go/consensus/hotstuff"
	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/engine"
	"github.com/onflow/flow-go/engine/common/fifoqueue"
	"github.com/onflow/flow-go/engine/consensus/sealing/counters"
	"github.com/onflow/flow-go/module/lifecycle"
	"github.com/onflow/flow-go/module/mempool"
)

// defaultVoteAggregatorWorkers number of workers to dispatch events for vote aggregators
const defaultVoteAggregatorWorkers = 8

// defaultVoteQueueCapacity maximum capacity of block votes queue
const defaultVoteQueueCapacity = 1000

// VoteAggregator stores the votes and aggregates them into a QC when enough votes have been collected
// VoteAggregator is designed in a way that it can aggregate votes for collection & consensus clusters
// that is why implementation relies on dependency injection.
type VoteAggregatorV2 struct {
	unit                *engine.Unit
	lm                  *lifecycle.LifecycleManager
	log                 zerolog.Logger
	notifier            hotstuff.Consumer
	highestPrunedView   counters.StrictMonotonousCounter
	collectors          hotstuff.VoteCollectors
	queuedVotesNotifier engine.Notifier
	queuedVotes         *fifoqueue.FifoQueue
}

var _ hotstuff.VoteAggregatorV2 = &VoteAggregatorV2{}

// NewVoteAggregatorV2 creates an instance of vote aggregator
// Note: verifyingProcessorFactory is injected. Thereby, the code is agnostic to the
// different voting formats of main Consensus vs Collector consensus.
func NewVoteAggregatorV2(
	log zerolog.Logger,
	notifier hotstuff.Consumer,
	highestPrunedView uint64,
	collectors hotstuff.VoteCollectors,
) (*VoteAggregatorV2, error) {

	queuedVotes, err := fifoqueue.NewFifoQueue(
		fifoqueue.WithCapacity(defaultVoteQueueCapacity))
	if err != nil {
		return nil, fmt.Errorf("could not initialize votes queue")
	}

	aggregator := &VoteAggregatorV2{
		unit:                engine.NewUnit(),
		lm:                  lifecycle.NewLifecycleManager(),
		log:                 log,
		notifier:            notifier,
		highestPrunedView:   counters.NewMonotonousCounter(highestPrunedView),
		collectors:          collectors,
		queuedVotes:         queuedVotes,
		queuedVotesNotifier: engine.NewNotifier(),
	}

	return aggregator, nil
}

// Ready returns a ready channel that is closed once the engine has fully
// started. For the propagation engine, we consider the engine up and running
// upon initialization.
func (va *VoteAggregatorV2) Ready() <-chan struct{} {
	va.lm.OnStart(func() {
		// launch as many workers as we need
		for i := 0; i < defaultVoteAggregatorWorkers; i++ {
			va.unit.Launch(va.queuedVotesProcessingLoop)
		}

		<-va.unit.Ready()
	})
	return va.lm.Started()
}

func (va *VoteAggregatorV2) Done() <-chan struct{} {
	va.lm.OnStop(func() {
		<-va.unit.Done()
	})
	return va.lm.Stopped()
}

func (va *VoteAggregatorV2) queuedVotesProcessingLoop() {
	notifier := va.queuedVotesNotifier.Channel()
	for {
		select {
		case <-va.unit.Quit():
			return
		case <-notifier:
			err := va.processQueuedVoteEvents()
			if err != nil {
				va.log.Fatal().Err(err).Msg("internal error processing block incorporated queued message")
			}
		}
	}
}

func (va *VoteAggregatorV2) processQueuedVoteEvents() error {
	for {
		select {
		case <-va.unit.Quit():
			return nil
		default:
		}

		msg, ok := va.queuedVotes.Pop()
		if ok {
			err := va.processQueuedVote(msg.(*model.Vote))
			if err != nil {
				return fmt.Errorf("could not process pending vote: %w", err)
			}
			continue
		}

		// when there is no more messages in the queue, back to the loop to wait
		// for the next incoming message to arrive.
		return nil
	}
}

// processQueuedVote performs actual processing of queued votes, this method is called from multiple
// concurrent goroutines.
func (va *VoteAggregatorV2) processQueuedVote(vote *model.Vote) error {
	// TODO: log created
	collector, _, err := va.collectors.GetOrCreateCollector(vote.View)
	if err != nil {
		// ignore if our routine is outdated and some other one has pruned collectors
		if mempool.IsDecreasingPruningHeightError(err) {
			return nil
		}
		return fmt.Errorf("could not get collector for view %d: %w",
			vote.View, err)
	}
	err = collector.AddVote(vote)
	if err != nil {
		if model.IsDoubleVoteError(err) {
			doubleVoteErr := err.(model.DoubleVoteError)
			va.notifier.OnDoubleVotingDetected(doubleVoteErr.FirstVote, doubleVoteErr.ConflictingVote)
			return nil
		}

		return fmt.Errorf("could not process vote for view %d, blockID %v: %w",
			vote.View, vote.BlockID, err)
	}

	return nil
}

// AddVote checks if vote is stale and appends vote into processing queue
// actual vote processing will be called in other dispatching goroutine.
func (va *VoteAggregatorV2) AddVote(vote *model.Vote) error {
	// drop stale votes
	if vote.View <= va.highestPrunedView.Value() {
		return nil
	}

	// It's ok to silently drop votes in case our processing pipeline is full.
	// It means that we are probably catching up.
	if ok := va.queuedVotes.Push(vote); ok {
		va.queuedVotesNotifier.Notify()
	}

	return nil
}

// AddBlock notifies the VoteAggregator about a known block so that it can start processing
// pending votes whose block was unknown.
// It also verifies the proposer vote of a block, and return whether the proposer signature is valid.
// Expected error returns during normal operations:
// * model.InvalidBlockError if the block is invalid
func (va *VoteAggregatorV2) AddBlock(block *model.Proposal) error {
	// check if the block is for a view that has already been pruned (and is thus stale)
	if block.Block.View <= va.highestPrunedView.Value() {
		return nil
	}

	collector, _, err := va.collectors.GetOrCreateCollector(block.Block.View)
	if err != nil {
		// ignore if our routine is outdated and some other one has pruned collectors
		if mempool.IsDecreasingPruningHeightError(err) {
			return nil
		}
		return fmt.Errorf("could not get or create collector for block %v: %w", block.Block.BlockID, err)
	}

	err = collector.ProcessBlock(block)
	if err != nil {
		return fmt.Errorf("could not process block: %v, %w", block.Block.BlockID, err)
	}

	return nil
}

// InvalidBlock notifies the VoteAggregator about an invalid proposal, so that it
// can process votes for the invalid block and slash the voters. Expected error
// returns during normal operations:
// * mempool.DecreasingPruningHeightError if proposal's view has already been pruned
func (va *VoteAggregatorV2) InvalidBlock(proposal *model.Proposal) error {
	slashingVoteConsumer := func(vote *model.Vote) {
		if proposal.Block.BlockID == vote.BlockID {
			va.notifier.OnVoteForInvalidBlockDetected(vote, proposal)
		}
	}

	block := proposal.Block
	collector, _, err := va.collectors.GetOrCreateCollector(block.View)
	if err != nil {
		// ignore if our routine is outdated and some other one has pruned collectors
		if mempool.IsDecreasingPruningHeightError(err) {
			return nil
		}
		return fmt.Errorf("could not retrieve vote collector for view %d: %w", block.View, err)
	}
	// registering vote consumer will deliver all previously cached votes in strict order
	// and will keep delivering votes if more are collected
	collector.RegisterVoteConsumer(slashingVoteConsumer)
	return nil
}

// PruneUpToView will delete all votes equal or below to the given view, as well as related indexes.
func (va *VoteAggregatorV2) PruneUpToView(view uint64) {
	// if someone else has updated view in parallel don't bother doing extra work for cleaning, whoever
	// is able to advance counter will perform the cleanup
	if va.highestPrunedView.Set(view) {
		va.collectors.PruneUpToView(view)
	}
}