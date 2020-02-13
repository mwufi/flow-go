package hotstuff

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/dapperlabs/flow-go/engine/consensus/hotstuff/types"
)

// Voter produces votes for the given block
type Voter struct {
	signer    Signer
	viewState *ViewState
	forks     Forks
	// Need to keep track of the last view we voted for so we don't double vote accidentally
	lastVotedView uint64
	log           zerolog.Logger
}

// NewVoter creates a new Voter instance
func (v *Voter) NewVoter(signer Signer, viewState *ViewState, forks Forks, log zerolog.Logger) *Voter {
	return &Voter{
		signer:        signer,
		viewState:     viewState,
		forks:         forks,
		lastVotedView: 0,
		log:           log,
	}
}

// ProduceVoteIfVotable will make a decision on whether it will vote for the given proposal, the returned
// boolean indicates whether to vote or not.
// In order to ensure that only a safe node will be voted, Voter will ask Forks whether a vote is a safe node or not.
// The curView is taken as input to ensure Voter will only vote for proposals at current view and prevent double voting.
// This method will only ever _once_ return a `non-nil, true` vote: the very first time it encounters a safe block of the
//  current view to vote for. Subsequently, voter does _not_ vote for any other block with the same (or lower) view.
// (including repeated calls with the initial block we voted for also return `nil, false`).
func (v *Voter) ProduceVoteIfVotable(bp *types.BlockProposal, curView uint64) (*types.Vote, bool) {
	if v.forks.IsSafeBlock(bp) {
		log.Info().Msg("received block is not a safe node, don't vote")
		return nil, false
	}

	if curView != bp.Block.View {
		log.Info().Uint64("view", bp.Block.View).Uint64("curView", curView).
			Msg("received block's view is not our current view, don't vote")
		return nil, false
	}

	if curView <= v.lastVotedView {
		log.Info().Uint64("lastVotedView", v.lastVotedView).Uint64("curView", curView).
			Msg("received block's view is <= lastVotedView, don't vote")
		return nil, false
	}

	vote, err := v.produceVote(bp)
	if err != nil {
		return nil, false
	}
	return vote, true
}

func (v *Voter) produceVote(bp *types.BlockProposal) (*types.Vote, error) {
	unsignedVote := types.NewUnsignedVote(bp.Block.View, bp.BlockID())
	vote, err := v.signer.SignVote(unsignedVote)
	if err != nil {
		return nil, fmt.Errorf("failed to sign the vote (blockID: %v, view: %v): %w", unsignedVote.BlockID, unsignedVote.View, err)
	}
	return vote, nil
}
