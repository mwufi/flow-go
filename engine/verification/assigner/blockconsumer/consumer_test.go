package blockconsumer

import (
	"sync"
	"testing"
	"time"

	"github.com/dgraph-io/badger/v2"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/engine/testutil"
	"github.com/onflow/flow-go/engine/verification/utils"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module"
	"github.com/onflow/flow-go/module/metrics"
	"github.com/onflow/flow-go/module/trace"
	bstorage "github.com/onflow/flow-go/storage/badger"
	"github.com/onflow/flow-go/utils/unittest"
)

// TestBlockToJob evaluates that a block can be converted to a job,
// and its corresponding job can be converted back to the same block.
func TestBlockToJob(t *testing.T) {
	block := unittest.BlockFixture()
	actual, err := toBlock(toJob(&block))
	require.NoError(t, err)
	require.Equal(t, &block, actual)
}

func TestProduceConsume(t *testing.T) {
	// pushing 10 finalized blocks sequentially to block reader, with 3 workers on consumer and the block processor
	// blocking on the blocks, results in processor only receiving the first three finalized blocks:
	// 10 blocks sequentially --> block reader --> consumer can read and push 3 blocks at a time to processor --> blocking processor.
	t.Run("pushing 10 blocks, blocking, receives 3", func(t *testing.T) {
		received := make([]*flow.Block, 0)
		lock := &sync.Mutex{}
		neverFinish := func(notifier module.ProcessingNotifier, block *flow.Block) {
			lock.Lock()
			defer lock.Unlock()
			received = append(received, block)

			// this processor never notifies consumer that it is done with the block.
			// hence from consumer perspective, it is blocking on each received block.
		}

		withConsumer(t, 10, 3, neverFinish, func(consumer *BlockConsumer, blocks []*flow.Block) {
			<-consumer.Ready()

			for i := 0; i < len(blocks); i++ {
				// consumer is only required to be "notified" that a new finalized block available.
				// It keeps track of the last finalized block it has read, and read the next height upon
				// getting notified as follows:
				consumer.OnFinalizedBlock(&model.Block{})
			}

			<-consumer.Done()

			// expects the processor receive only the first 3 blocks (since it is blocked on those, hence no
			// new block is fetched to process).
			unittest.RequireBlockListsMatchElements(t, received, blocks[:3])
		})
	})

	// pushing 10 finalized blocks sequentially to block reader, with 3 workers on consumer and the processor finishes processing
	// each received block immediately, results in processor receiving all 10 blocks:
	// 10 blocks sequentially --> block reader --> consumer can read and push 3 blocks at a time to processor --> processor finishes blocks
	// immediately.
	t.Run("pushing 10 blocks, non-blocking, receives 10", func(t *testing.T) {
		received := make([]*flow.Block, 0)
		lock := &sync.Mutex{}
		var processAll sync.WaitGroup
		alwaysFinish := func(notifier module.ProcessingNotifier, block *flow.Block) {
			lock.Lock()
			defer lock.Unlock()

			received = append(received, block)

			go func() {
				notifier.Notify(block.ID())
				processAll.Done()
			}()
		}

		withConsumer(t, 10, 3, alwaysFinish, func(consumer *BlockConsumer, blocks []*flow.Block) {
			<-consumer.Ready()
			processAll.Add(len(blocks))

			for i := 0; i < len(blocks); i++ {
				// consumer is only required to be "notified" that a new finalized block available.
				// It keeps track of the last finalized block it has read, and read the next height upon
				// getting notified as follows:
				consumer.OnFinalizedBlock(&model.Block{})
			}

			// waits until all blocks finish processing
			unittest.RequireReturnsBefore(t, processAll.Wait, 1*time.Second, "could not process all blocks on time")
			<-consumer.Done()

			// expects the mock engine receive all 10 blocks.
			unittest.RequireBlockListsMatchElements(t, received, blocks)
		})
	})

	// pushing 100 finalized blocks concurrently to block reader, with 3 workers on consumer and the processor finishes processing
	// all blocks immediately, results in the processor receiving all 100 blocks:
	// 100 blocks concurrently --> block reader --> consumer can read and push 3 blocks at a time to processor --> processor finishes blocks
	// immediately.
	t.Run("pushing 100 blocks concurrently, non-blocking, receives 100", func(t *testing.T) {
		received := make([]*flow.Block, 0)
		lock := &sync.Mutex{}
		var processAll sync.WaitGroup

		alwaysFinish := func(notifier module.ProcessingNotifier, block *flow.Block) {
			lock.Lock()
			defer lock.Unlock()

			received = append(received, block)

			go func() {
				notifier.Notify(block.ID())
				processAll.Done()
			}()
		}

		withConsumer(t, 100, 3, alwaysFinish, func(consumer *BlockConsumer, blocks []*flow.Block) {
			<-consumer.Ready()
			processAll.Add(len(blocks))

			for i := 0; i < len(blocks); i++ {
				go func() {
					// consumer is only required to be "notified" that a new finalized block available.
					// It keeps track of the last finalized block it has read, and read the next height upon
					// getting notified as follows:
					consumer.OnFinalizedBlock(&model.Block{})
				}()
			}

			// waits until all finished
			unittest.RequireReturnsBefore(t, processAll.Wait, 1*time.Second, "could not process all blocks on time")
			<-consumer.Done()

			// expects the mock engine receive all 100 blocks.
			unittest.RequireBlockListsMatchElements(t, received, blocks)
		})
	})

}

// withConsumer is a test helper that sets up a block consumer with specified number of workers.
// The block consumer that operates on a block reader with a chain of specified number of finalized blocks
// ready to read and consumer.
func withConsumer(
	t *testing.T,
	blockCount int,
	workerCount int,
	process func(notifier module.ProcessingNotifier, block *flow.Block),
	withBlockConsumer func(*BlockConsumer, []*flow.Block),
) {
	unittest.RunWithBadgerDB(t, func(db *badger.DB) {
		maxProcessing := int64(workerCount)

		processedHeight := bstorage.NewConsumerProgress(db, module.ConsumeProgressVerificationBlockHeight)
		collector := &metrics.NoopCollector{}
		tracer := &trace.NoopTracer{}
		participants := unittest.IdentityListFixture(5, unittest.WithAllRoles())
		s := testutil.CompleteStateFixture(t, collector, tracer, participants)

		engine := &mockFinalizedBlockProcessor{
			process: process,
		}

		consumer, _, err := NewBlockConsumer(unittest.Logger(),
			processedHeight,
			s.Storage.Blocks,
			s.State,
			engine,
			maxProcessing)
		require.NoError(t, err)

		// generates a chain of blocks in the form of root <- R1 <- C1 <- R2 <- C2 <- ... where Rs are distinct reference
		// blocks (i.e., containing guarantees), and Cs are container blocks for their preceding reference block,
		// Container blocks only contain receipts of their preceding reference blocks. But they do not
		// hold any guarantees.
		root, err := s.State.Params().Root()
		require.NoError(t, err)
		results := utils.CompleteExecutionResultChainFixture(t, root, blockCount/2)
		blocks := extendStateWithFinalizedBlocks(t, results, s.State)
		// makes sure that we generated a block chain of requested length.
		require.Len(t, blocks, blockCount)

		withBlockConsumer(consumer, blocks)
	})
}

// mockFinalizedBlockProcessor provides a FinalizedBlockProcessor with a plug-and-play process method.
type mockFinalizedBlockProcessor struct {
	notifier module.ProcessingNotifier
	process  func(module.ProcessingNotifier, *flow.Block)
}

func (e *mockFinalizedBlockProcessor) ProcessFinalizedBlock(block *flow.Block) {
	e.process(e.notifier, block)
}

func (e *mockFinalizedBlockProcessor) WithBlockConsumerNotifier(notifier module.ProcessingNotifier) {
	e.notifier = notifier
}
