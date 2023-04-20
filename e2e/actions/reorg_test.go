package actions

import (
	"math/big"
	"math/rand"
	"path"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/components/node/client"
	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/sources"
	"github.com/kroma-network/kroma/components/node/testlog"
	"github.com/kroma-network/kroma/e2e/e2eutils"
)

func setupReorgTest(t Testing, config *e2eutils.TestParams) (*e2eutils.SetupData, *e2eutils.DeployParams, *L1Miner, *L2Proposer, *L2Engine, *L2Syncer, *L2Engine, *L2Batcher) {
	dp := e2eutils.MakeDeployParams(t, config)

	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)

	return setupReorgTestActors(t, dp, sd, log)
}

func setupReorgTestActors(t Testing, dp *e2eutils.DeployParams, sd *e2eutils.SetupData, log log.Logger) (*e2eutils.SetupData, *e2eutils.DeployParams, *L1Miner, *L2Proposer, *L2Engine, *L2Syncer, *L2Engine, *L2Batcher) {
	miner, propEngine, proposer := setupProposerTest(t, sd, log)
	miner.ActL1SetFeeRecipient(common.Address{'A'})
	proposer.ActL2PipelineFull(t)
	syncEngine, syncer := setupSyncer(t, sd, log, miner.L1Client(t, sd.RollupCfg))
	rollupPropCl := proposer.RollupClient()
	batcher := NewL2Batcher(log, sd.RollupCfg, &BatcherCfg{
		MinL1TxSize: 0,
		MaxL1TxSize: 128_000,
		BatcherKey:  dp.Secrets.Batcher,
	}, rollupPropCl, miner.EthClient(), propEngine.EthClient())
	return sd, dp, miner, proposer, propEngine, syncer, syncEngine, batcher
}

func TestReorgOrphanBlock(gt *testing.T) {
	t := NewDefaultTesting(gt)
	sd, _, miner, proposer, _, syncer, syncerEngine, batcher := setupReorgTest(t, defaultRollupTestParams)
	syncEngClient := syncerEngine.EngineClient(t, sd.RollupCfg)

	proposer.ActL2PipelineFull(t)
	syncer.ActL2PipelineFull(t)

	// build empty L1 block
	miner.ActEmptyBlock(t)

	// Create L2 blocks, and reference the L1 head as origin
	proposer.ActL1HeadSignal(t)
	proposer.ActBuildToL1Head(t)

	// submit all new L2 blocks
	batcher.ActSubmitAll(t)

	// new L1 block with L2 batch
	miner.ActL1StartBlock(12)(t)
	miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
	batchTx := miner.l1Transactions[0]
	miner.ActL1EndBlock(t)

	// syncer picks up the L2 chain that was submitted
	syncer.ActL1HeadSignal(t)
	syncer.ActL2PipelineFull(t)
	require.Equal(t, syncer.L2Safe(), proposer.L2Unsafe(), "syncer syncs from proposer via L1")
	require.NotEqual(t, proposer.L2Safe(), proposer.L2Unsafe(), "proposer has not processed L1 yet")

	// orphan the L1 block that included the batch tx, and build a new different L1 block
	miner.ActL1RewindToParent(t)
	miner.ActL1SetFeeRecipient(common.Address{'B'})
	miner.ActEmptyBlock(t)
	miner.ActEmptyBlock(t) // needs to be a longer chain for reorg to be applied. TODO: maybe more aggressively react to reorgs to shorter chains?

	// sync syncer again. The L1 reorg excluded the batch, so now the previous L2 chain should be unsafe again.
	// However, the L2 chain can still be canonical later, since it did not reference the reorged L1 block
	syncer.ActL1HeadSignal(t)
	syncer.ActL2PipelineFull(t)
	require.Equal(t, syncer.L2Safe(), proposer.L2Safe(), "syncer rewinds safe when L1 reorgs out batch")
	ref, err := syncEngClient.L2BlockRefByLabel(t.Ctx(), eth.Safe)
	require.NoError(t, err)
	require.Equal(t, syncer.L2Safe(), ref, "syncer engine matches rollup client")

	// Now replay the batch tx in a new L1 block
	miner.ActL1StartBlock(12)(t)
	miner.ActL1SetFeeRecipient(common.Address{'C'})
	// note: the geth tx pool reorgLoop is too slow (responds to chain head events, but async),
	// and there's no way to manually trigger runReorg, so we re-insert it ourselves.
	require.NoError(t, miner.eth.TxPool().AddLocal(batchTx))
	// need to re-insert previously included tx into the block
	miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
	miner.ActL1EndBlock(t)

	// sync the syncer again: now it should be safe again
	syncer.ActL1HeadSignal(t)
	syncer.ActL2PipelineFull(t)
	require.Equal(t, syncer.L2Safe(), proposer.L2Unsafe(), "syncer syncs from proposer via replayed batch on L1")
	ref, err = syncEngClient.L2BlockRefByLabel(t.Ctx(), eth.Safe)
	require.NoError(t, err)
	require.Equal(t, syncer.L2Safe(), ref, "syncer engine matches rollup client")

	proposer.ActL1HeadSignal(t)
	proposer.ActL2PipelineFull(t)
	require.Equal(t, syncer.L2Safe(), proposer.L2Safe(), "syncer and proposer see same safe L2 block, while only syncer dealt with the orphan and replay")
}

func TestReorgFlipFlop(gt *testing.T) {
	t := NewDefaultTesting(gt)
	sd, _, miner, proposer, _, syncer, syncerEng, batcher := setupReorgTest(t, defaultRollupTestParams)
	minerCl := miner.L1Client(t, sd.RollupCfg)
	syncEngClient := syncerEng.EngineClient(t, sd.RollupCfg)
	checkSyncEngine := func() {
		// TODO: geth preserves L2 chain with origin A1 after flip-flopping to B?
		//ref, err := syncEngClient.L2BlockRefByLabel(t.Ctx(), eth.Unsafe)
		//require.NoError(t, err)
		//t.Logf("l2 unsafe head %s with origin %s", ref, ref.L1Origin)
		//require.NotEqual(t, syncer.L2Unsafe().Hash, ref.ParentHash, "TODO off by one, engine syncs A0 after reorging back from B, while rollup node only inserts up to A0 (excl.)")
		//require.Equal(t, syncer.L2Unsafe(), ref, "syncer safe head of engine matches rollup client")

		ref, err := syncEngClient.L2BlockRefByLabel(t.Ctx(), eth.Safe)
		require.NoError(t, err)
		require.Equal(t, syncer.L2Safe(), ref, "syncer safe head of engine matches rollup client")
	}

	proposer.ActL2PipelineFull(t)
	syncer.ActL2PipelineFull(t)

	// Start building chain A
	miner.ActL1SetFeeRecipient(common.Address{'A', 0})
	miner.ActEmptyBlock(t)
	blockA0, err := minerCl.L1BlockRefByLabel(t.Ctx(), eth.Unsafe)
	require.NoError(t, err)

	// Create L2 blocks, and reference the L1 head A0 as origin
	proposer.ActL1HeadSignal(t)
	proposer.ActBuildToL1Head(t)

	// submit all new L2 blocks
	batcher.ActSubmitAll(t)

	// new L1 block A1 with L2 batch
	miner.ActL1SetFeeRecipient(common.Address{'A', 1})
	miner.ActL1StartBlock(12)(t)
	miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
	batchTxA := miner.l1Transactions[0]
	miner.ActL1EndBlock(t)

	// syncer picks up the L2 chain that was submitted
	syncer.ActL1HeadSignal(t)
	syncer.ActL2PipelineFull(t)
	require.Equal(t, syncer.L2Safe().L1Origin, blockA0.ID(), "syncer syncs L2 chain with L1 A0 origin")
	checkSyncEngine()

	// Flip to chain B!
	miner.ActL1RewindToParent(t) // undo A1
	miner.ActL1RewindToParent(t) // undo A0
	// build B0
	miner.ActL1SetFeeRecipient(common.Address{'B', 0})
	miner.ActEmptyBlock(t)
	blockB0, err := minerCl.L1BlockRefByLabel(t.Ctx(), eth.Unsafe)
	require.NoError(t, err)
	require.Equal(t, blockA0.Number, blockB0.Number, "same height")
	require.NotEqual(t, blockA0.Hash, blockB0.Hash, "different content")

	// re-include the batch tx that submitted L2 chain data that pointed to A0, in the new block B1
	miner.ActL1SetFeeRecipient(common.Address{'B', 1})
	miner.ActL1StartBlock(12)(t)
	require.NoError(t, miner.eth.TxPool().AddLocal(batchTxA))
	miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
	miner.ActL1EndBlock(t)

	// make B2, the reorg is picked up when we have a new longer chain
	miner.ActL1SetFeeRecipient(common.Address{'B', 2})
	miner.ActEmptyBlock(t)
	blockB2, err := minerCl.L1BlockRefByLabel(t.Ctx(), eth.Unsafe)
	require.NoError(t, err)

	// now sync the syncer: some of the batches should be ignored:
	//	The safe head should have a genesis L1 origin, but past genesis, as some L2 blocks were built to get to A0 time
	syncer.ActL1HeadSignal(t)
	syncer.ActL2PipelineFull(t)
	require.Equal(t, sd.RollupCfg.Genesis.L1, syncer.L2Safe().L1Origin, "expected to be back at genesis origin after losing A0 and A1")

	require.NotZero(t, syncer.L2Safe().Number, "still preserving old L2 blocks that did not reference reorged L1 chain (assuming more than one L2 block per L1 block)")
	require.Equal(t, syncer.L2Safe(), syncer.L2Unsafe(), "head is at safe block after L1 reorg")
	checkSyncEngine()

	// and sync the proposer, then build some new L2 blocks, up to and including with L1 origin B2
	proposer.ActL1HeadSignal(t)
	proposer.ActL2PipelineFull(t)
	proposer.ActBuildToL1Head(t)
	require.Equal(t, proposer.L2Unsafe().L1Origin, blockB2.ID(), "B2 is the unsafe L1 origin of proposer now")

	// submit all new L2 blocks for chain B, and include in new block B3
	batcher.ActSubmitAll(t)
	miner.ActL1SetFeeRecipient(common.Address{'B', 3})
	miner.ActL1StartBlock(12)(t)
	miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
	miner.ActL1EndBlock(t)

	// sync the syncer to the L2 chain with origin B2
	syncer.ActL1HeadSignal(t)
	syncer.ActL2PipelineFull(t)
	require.Equal(t, syncer.L2Safe().L1Origin, blockB2.ID(), "B2 is the L1 origin of syncer now")
	checkSyncEngine()

	// Flop back to chain A!
	miner.ActL1RewindDepth(4)(t) // B3, B2, B1, B0
	pivotBlock, err := minerCl.L1BlockRefByLabel(t.Ctx(), eth.Unsafe)
	require.NoError(t, err)
	require.Equal(t, sd.RollupCfg.Genesis.L1, pivotBlock.ID(), "back at L1 genesis")
	miner.ActL1SetFeeRecipient(common.Address{'A', 0})
	miner.ActEmptyBlock(t)
	blockA0Again, err := minerCl.L1BlockRefByLabel(t.Ctx(), eth.Unsafe)
	require.NoError(t, err)
	require.Equal(t, blockA0, blockA0Again, "block A0 is back again after flip-flop")

	// Continue to build on A, and replay the old A-chain batch transaction (we can't replay B, since the nonce values would conflict)
	miner.ActL1SetFeeRecipient(common.Address{'A', 1})
	miner.ActEmptyBlock(t)

	miner.ActL1SetFeeRecipient(common.Address{'A', 2})
	miner.ActL1StartBlock(12)(t)
	require.NoError(t, miner.eth.TxPool().AddLocal(batchTxA)) // replay chain A batches, but now in A2 instead of A1
	miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
	miner.ActL1EndBlock(t)

	// build more L1 blocks, so the chain is long enough for reorg to be picked up
	miner.ActL1SetFeeRecipient(common.Address{'A', 3})
	miner.ActEmptyBlock(t)
	miner.ActL1SetFeeRecipient(common.Address{'A', 4})
	miner.ActEmptyBlock(t)
	blockA4, err := minerCl.L1BlockRefByLabel(t.Ctx(), eth.Unsafe)
	require.NoError(t, err)

	// sync syncer, and see if A0 is safe, using the old replayed batch for chain A
	syncer.ActL1HeadSignal(t)
	syncer.ActL2PipelineFull(t)
	require.Equal(t, syncer.L2Safe().L1Origin, blockA0.ID(), "B2 is the L1 origin of syncer now")
	checkSyncEngine()

	// sync proposer to the replayed L1 chain A
	proposer.ActL1HeadSignal(t)
	proposer.ActL2PipelineFull(t)
	require.Equal(t, syncer.L2Safe(), proposer.L2Safe(), "proposer reorgs to match syncer again")

	// and adopt the rest of L1 chain A into L2
	proposer.ActBuildToL1Head(t)

	// submit the new unsafe A blocks
	batcher.ActSubmitAll(t)
	miner.ActL1SetFeeRecipient(common.Address{'A', 5})
	miner.ActL1StartBlock(12)(t)
	miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
	miner.ActL1EndBlock(t)

	// sync syncer to what ths proposer submitted
	syncer.ActL1HeadSignal(t)
	syncer.ActL2PipelineFull(t)
	require.Equal(t, syncer.L2Safe(), proposer.L2Unsafe(), "syncer syncs from proposer")
	require.Equal(t, syncer.L2Safe().L1Origin, blockA4.ID(), "L2 chain origin is A4")
	checkSyncEngine()
}

// Deep Reorg Test
//
// Steps:
//  1. Create an L1 actor
//  2. Ask the L1 actor to build three proposer windows of empty blocks
//     2.a alice submits a transaction on l2 with an l1 origin of block #35
//     2.b in block #50, include the batch that contains the l2 block with alice's transaction as well
//     as all other blocks before it.
//  3. Ask the L2 proposer to build a chain that references these L1 blocks
//  4. Ask the batch submitter to submit remaining unsafe L2 blocks to L1
//  5. Ask the L1 to include this data
//  6. Rewind chain A 21 blocks
//  7. Ask the L1 actor to build one proposer window + 1 empty blocks on chain B
//  8. Ask the L1 actor to build an empty block in place of the batch submission block on chain A
//  9. Ask the L1 actor to create another empty block so that chain B is longer than chain A
//  10. Ask the L2 proposer to send a head signal and run one iteration of the derivation pipeline.
//  11. Ask the L2 proposer build a chain that references chain B's blocks
//  12. Sync the syncer and assert that the L2 safe head L1 origin has caught up with chain B
//  13. Ensure that the parent L2 block of the block that contains Alice's transaction still exists
//     after the L2 has re-derived from chain B.
//  14. Ensure that the L2 block that contained Alice's transaction before the reorg no longer exists.
//
// Chain A
// - 61 blocks total
//   - 60 empty blocks
//   - Alice submits her L2 transaction with an L1 Origin of block #35
//   - In block 50, submit the batch containing the L2 block with Alice's transaction.
//   - Block 61 includes batch with blocks [1, 60]
//
// Syncer
// - Prior to second batch submission, safe head origin is block A50
// - After batch, safe head origin is block A60
// - Unsafe head origin is A61
//
// Reorg L1 (start: block #61, depth: 22 blocks)
// - Rewind depth: Batch submission block + ProposerWindowSize+1 blocks
// - Wind back to block #39
//
// Before building L2 to L1 head / syncing syncer & proposer:
// Syncer
// - Unsafe head L1 origin is block #60
// - Safe head L1 origin is at genesis block #60
//
// Build Chain B
// - 62 blocks total
//   - 39 empty blocks left over from chain A
//   - 21 empty blocks
//   - empty block (61)
//   - empty block (62) <- Makes chain B longer than chain A, the re-org will be picked up
//
// After building L2 to L1 head:
// Syncer
// - Unsafe head is 62
// - Safe head is 42
func TestDeepReorg(gt *testing.T) {
	t := NewDefaultTesting(gt)

	// Create actor and synchronization engine client
	sd, dp, miner, proposer, propEngine, syncer, syncerEng, batcher := setupReorgTest(t, &e2eutils.TestParams{
		MaxProposerDrift:   40,
		ProposerWindowSize: 20,
		ChannelTimeout:     120,
		L1BlockTime:        4,
	})
	minerCl := miner.L1Client(t, sd.RollupCfg)
	l2Client := propEngine.EthClient()
	syncEngClient := syncerEng.EngineClient(t, sd.RollupCfg)
	checkSyncEngine := func() {
		ref, err := syncEngClient.L2BlockRefByLabel(t.Ctx(), eth.Safe)
		require.NoError(t, err)
		require.Equal(t, syncer.L2Safe(), ref, "syncer safe head of engine matches rollup client")
	}

	// Set up alice
	log := testlog.Logger(t, log.LvlDebug)
	addresses := e2eutils.CollectAddresses(sd, dp)
	l2UserEnv := &BasicUserEnv[*L2Bindings]{
		EthCl:          l2Client,
		Signer:         types.LatestSigner(sd.L2Cfg.Config),
		AddressCorpora: addresses,
		Bindings:       NewL2Bindings(t, l2Client, propEngine.GethClient()),
	}
	alice := NewCrossLayerUser(log, dp.Secrets.Alice, rand.New(rand.NewSource(0xa57b)), sd.RollupCfg)
	alice.L2.SetUserEnv(l2UserEnv)

	// Run one iteration of the L2 derivation pipeline
	proposer.ActL1HeadSignal(t)
	proposer.ActL2PipelineFull(t)
	syncer.ActL2PipelineFull(t)

	// Start building chain A
	miner.ActL1SetFeeRecipient(common.Address{0x0A, 0x00})

	// Create a var to store the ref for the second to last block of the second proposer window
	var blockA39 eth.L1BlockRef

	var aliceL2TxBlock types.Block
	// Mine enough empty blocks on L1 to reach two proposer windows.
	for i := uint64(0); i < sd.RollupCfg.ProposerWindowSize*3; i++ {
		// At block #50, send a batch to L1 containing all L2 blocks built up to this point.
		// This batch contains alice's transaction, and will be reorg'd out of the L1 chain
		// later in the test.
		if i == 50 {
			batcher.ActSubmitAll(t)

			miner.ActL1StartBlock(12)(t)
			miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
			miner.ActL1EndBlock(t)
		} else {
			miner.ActEmptyBlock(t)
		}

		// Get the second to last block of the first proposer window
		// This is used later to verify the head of chain B after rewinding
		// chain A 1 proposer window + 1 block + Block A1 (batch submission with two
		// proposer windows worth of transactions)
		if i == sd.RollupCfg.ProposerWindowSize*2-2 {
			var err error
			blockA39, err = minerCl.L1BlockRefByLabel(t.Ctx(), eth.Unsafe)
			require.NoError(t, err)
			require.Equal(t, uint64(39), blockA39.Number)
		}

		// Submit a dummy tx on L2 as alice with an L1 origin that will remain in the
		// canonical chain after the reorg. The batch that contains this transaction
		// will be submitted in block #50, which *will* be reorg'd out of the L1 chain, so the
		// L2 block that contains this transaction should no longer exist after L2 has
		// been re-derived from chain B later on in the test.
		if i == 35 {
			// Include alice's transaction on L2
			proposer.ActL2StartBlock(t)

			// Submit a dummy tx
			alice.L2.ActResetTxOpts(t)
			alice.L2.ActSetTxToAddr(&dp.Addresses.Bob)(t)
			alice.L2.ActMakeTx(t)

			// Include the tx in the block we're making
			propEngine.ActL2IncludeTx(alice.Address())(t)

			// Finalize the L2 block containing alice's transaction
			proposer.ActL2EndBlock(t)

			// Store the ref to the L2 block that the transaction was included in for later.
			b0, err := l2Client.BlockByNumber(t.Ctx(), big.NewInt(int64(proposer.L2Unsafe().Number)))
			require.NoError(t, err, "failed to fetch unsafe head of L2 after submitting alice's transaction")

			aliceL2TxBlock = *b0
		}

		// Ask proposer to handle new L1 head and build L2 blocks up to the L1 head
		proposer.ActL1HeadSignal(t)
		proposer.ActL2PipelineFull(t)
		proposer.ActBuildToL1Head(t)
	}

	// Get the last empty block built in the loop above.
	// This will be the last block in the third proposer window.
	blockA60, err := minerCl.L1BlockRefByLabel(t.Ctx(), eth.Unsafe)
	require.NoError(t, err)

	// Check that the safe head's L1 origin is block A50 before batch submission
	require.Equal(t, uint64(50), proposer.L2Safe().L1Origin.Number)
	// Check that the unsafe head's L1 origin is block A60
	require.Equal(t, blockA60.ID(), proposer.L2Unsafe().L1Origin)

	// Batch and submit all new L2 blocks that were built above to L1
	batcher.ActSubmitAll(t)

	// Build a new block on L1 that includes the L2 batch containing all blocks
	// between [51, 60]
	miner.ActL1SetFeeRecipient(common.Address{0x0A, 0x01})
	miner.ActL1StartBlock(12)(t)
	miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
	miner.ActL1EndBlock(t)

	// Handle the new head block on both the syncer and the proposer
	syncer.ActL1HeadSignal(t)
	proposer.ActL1HeadSignal(t)

	// Run one iteration of the L2 derivation pipeline on both the syncer and proposer
	syncer.ActL2PipelineFull(t)
	proposer.ActL2PipelineFull(t)

	// Ensure that the syncer picks up that the L2 blocks were submitted to L1
	// and marks them as safe.
	// We check that the L2 safe L1 origin is block A240, or the last block
	// within the second proposer window. This is the block directly before
	// the block that included the batch on chain A.
	require.Equal(t, blockA60.ID(), syncer.L2Safe().L1Origin)
	checkSyncEngine()

	// Perform a deep reorg the size of one proposer window + 2 blocks.
	// This will affect the safe L2 chain.
	miner.ActL1RewindToParent(t)                                   // Rewind the batch submission
	miner.ActL1RewindDepth(sd.RollupCfg.ProposerWindowSize + 1)(t) // Rewind one proposer window + 1 block

	// Ensure that the block we rewinded to on L1 is the second to last block of the first
	// proposer window.
	headAfterReorg, err := minerCl.L1BlockRefByLabel(t.Ctx(), eth.Unsafe)
	require.NoError(t, err)

	// Ensure that we landed on the intended L1 block after the reorg
	require.Equal(t, blockA39.ID(), headAfterReorg.ID())

	// Ensure that the safe L2 head has not been altered yet- we have not issued
	// a head signal to the proposer or syncer post reorg.
	require.Equal(t, blockA60.ID(), syncer.L2Safe().L1Origin)
	require.Equal(t, blockA60.ID(), proposer.L2Safe().L1Origin)
	// Ensure that the L2 unsafe head has not been altered yet- we have not issued
	// a head signal to the proposer or syncer post reorg.
	require.Equal(t, blockA60.ID(), syncer.L2Unsafe().L1Origin)
	require.Equal(t, blockA60.ID(), proposer.L2Unsafe().L1Origin)
	checkSyncEngine()

	// --------- [ CHAIN B ] ---------

	// Start building chain B
	miner.ActL1SetFeeRecipient(common.Address{0x0B, 0x00})
	// Mine enough empty blocks on L1 to reach three proposer windows or 60 blocks.
	// We already have 39 empty blocks on the rewinded L1 that are left over from chain A.
	for i := uint64(0); i < sd.RollupCfg.ProposerWindowSize+1; i++ {
		miner.ActEmptyBlock(t)

		// Ask proposer to handle new L1 head and build L2 blocks up to the L1 head
		proposer.ActL1HeadSignal(t)
		proposer.ActL2PipelineFull(t)
		proposer.ActBuildToL1Head(t)
	}

	// Get the last unsafe block on chain B after creating ProposerWindowSize+1 empty blocks
	blockB60, err := minerCl.L1BlockRefByLabel(t.Ctx(), eth.Unsafe)
	require.NoError(t, err)
	// Ensure blockB60 is #60 on chain B
	require.Equal(t, uint64(60), blockB60.Number)

	// Mine an empty block in place of the block that included the final batch on chain A
	miner.ActL1SetFeeRecipient(common.Address{0x0B, 0x01})
	miner.ActEmptyBlock(t)

	// Make block B62. the reorg is picked up when we have a new, longer chain.
	miner.ActL1SetFeeRecipient(common.Address{0x0B, 0x02})
	miner.ActEmptyBlock(t)
	blockB62, err := minerCl.L1BlockRefByLabel(t.Ctx(), eth.Unsafe)
	require.NoError(t, err)

	// Now sync the syncer. The batch from chain A is invalid, so it should have been ignored.
	// The safe head should have an origin at block B42
	syncer.ActL1HeadSignal(t)
	syncer.ActL2PipelineFull(t)

	// Ensure that the L2 Safe block is B42
	require.Equal(t, uint64(42), syncer.L2Safe().L1Origin.Number, "expected to be at block #42 after losing A40-61")
	require.NotZero(t, syncer.L2Safe().Number, "still preserving old L2 blocks that did not reference reorged L1 chain (assuming more than one L2 block per L1 block)")
	require.Equal(t, syncer.L2Safe(), syncer.L2Unsafe(), "L2 safe and unsafe head should be equal")
	checkSyncEngine()

	// Sync the proposer, then build some new L2 blocks, up to and including with L1 origin B62
	proposer.ActL1HeadSignal(t)
	proposer.ActL2PipelineFull(t)
	proposer.ActBuildToL1Head(t)
	require.Equal(t, proposer.L2Unsafe().L1Origin, blockB62.ID())

	// Sync the syncer to the L2 chain with origin B62
	// Run an iteration of the derivation pipeline and ensure that the L2 safe L1 origin is block B62
	syncer.ActL1HeadSignal(t)
	syncer.ActL2PipelineFull(t)
	require.Equal(t, uint64(42), syncer.L2Safe().L1Origin.Number, "expected to be at block #42 after losing A40-61 and building 23 blocks on reorged chain")
	require.Equal(t, syncer.L2Safe(), syncer.L2Unsafe(), "L2 safe and unsafe head should be equal")
	checkSyncEngine()

	// Ensure that the parent of the L2 block containing Alice's transaction still exists
	b0, err := l2Client.BlockByHash(t.Ctx(), aliceL2TxBlock.ParentHash())
	require.NoError(t, err, "Parent of the L2 block containing Alice's transaction should still exist on L2")
	require.Equal(t, b0.Hash(), aliceL2TxBlock.ParentHash())

	// Ensure that the L2 block containing Alice's transaction no longer exists.
	b1, err := l2Client.BlockByNumber(t.Ctx(), aliceL2TxBlock.Number())
	require.NoError(t, err, "A block that has the same number as the block that contained Alice's transaction should still exist on L2")
	require.Equal(t, b1.Number(), aliceL2TxBlock.Number())
	require.NotEqual(t, b1.Hash(), aliceL2TxBlock.Hash(), "L2 block containing Alice's transaction should no longer exist on L2")
}

type rpcWrapper struct {
	client.RPC
}

// TestRestartKromaGeth tests that the proposer can restart its execution engine without rollup-node restart,
// including recovering the finalized/safe state of L2 chain without reorging.
func TestRestartKromaGeth(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dbPath := path.Join(t.TempDir(), "testdb")
	dbOption := func(_ *ethconfig.Config, nodeCfg *node.Config) error {
		nodeCfg.DataDir = dbPath
		return nil
	}
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)
	jwtPath := e2eutils.WriteDefaultJWT(t)
	// L1
	miner := NewL1Miner(t, log, sd.L1Cfg)
	l1F, err := sources.NewL1Client(miner.RPCClient(), log, nil, sources.L1ClientDefaultConfig(sd.RollupCfg, false, sources.RPCKindBasic))
	require.NoError(t, err)
	// Proposer
	propEng := NewL2Engine(t, log, sd.L2Cfg, sd.RollupCfg.Genesis.L1, jwtPath, dbOption)
	engRpc := &rpcWrapper{propEng.RPCClient()}
	l2Cl, err := sources.NewEngineClient(engRpc, log, nil, sources.EngineClientDefaultConfig(sd.RollupCfg))
	require.NoError(t, err)
	proposer := NewL2Proposer(t, log, l1F, l2Cl, sd.RollupCfg, 0)

	batcher := NewL2Batcher(log, sd.RollupCfg, &BatcherCfg{
		MinL1TxSize: 0,
		MaxL1TxSize: 128_000,
		BatcherKey:  dp.Secrets.Batcher,
	}, proposer.RollupClient(), miner.EthClient(), propEng.EthClient())

	// start
	proposer.ActL2PipelineFull(t)

	miner.ActEmptyBlock(t)

	buildAndSubmit := func() {
		// build some blocks
		proposer.ActL1HeadSignal(t)
		proposer.ActBuildToL1Head(t)
		// submit the blocks, confirm on L1
		batcher.ActSubmitAll(t)
		miner.ActL1StartBlock(12)(t)
		miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
		miner.ActL1EndBlock(t)
		proposer.ActL2PipelineFull(t)
	}
	buildAndSubmit()

	// finalize the L1 data (first block, and the new block with batch)
	miner.ActL1SafeNext(t)
	miner.ActL1SafeNext(t)
	miner.ActL1FinalizeNext(t)
	miner.ActL1FinalizeNext(t)
	proposer.ActL1FinalizedSignal(t)
	proposer.ActL1SafeSignal(t)

	// build and submit more
	buildAndSubmit()
	// but only mark the L1 block with this batch as safe
	miner.ActL1SafeNext(t)
	proposer.ActL1SafeSignal(t)

	// build some more, these stay unsafe
	miner.ActEmptyBlock(t)
	proposer.ActL1HeadSignal(t)
	proposer.ActBuildToL1Head(t)

	statusBeforeRestart := proposer.SyncStatus()
	// before restart scenario: we have a distinct finalized, safe, and unsafe part of the L2 chain
	require.NotZero(t, statusBeforeRestart.FinalizedL2.L1Origin.Number)
	require.Less(t, statusBeforeRestart.FinalizedL2.L1Origin.Number, statusBeforeRestart.SafeL2.L1Origin.Number)
	require.Less(t, statusBeforeRestart.SafeL2.L1Origin.Number, statusBeforeRestart.UnsafeL2.L1Origin.Number)

	// close the proposer engine
	require.NoError(t, propEng.Close())
	// and start a new one with same db path
	propEngNew := NewL2Engine(t, log, sd.L2Cfg, sd.RollupCfg.Genesis.L1, jwtPath, dbOption)
	// swap in the new rpc. This is as close as we can get to reconnecting to a new in-memory rpc connection
	engRpc.RPC = propEngNew.RPCClient()

	// note: geth does not persist the safe block label, only the finalized block label
	safe, err := l2Cl.L2BlockRefByLabel(t.Ctx(), eth.Safe)
	require.NoError(t, err)
	finalized, err := l2Cl.L2BlockRefByLabel(t.Ctx(), eth.Finalized)
	require.NoError(t, err)
	require.Equal(t, statusBeforeRestart.FinalizedL2, safe, "expecting to revert safe head to finalized head upon restart")
	require.Equal(t, statusBeforeRestart.FinalizedL2, finalized, "expecting to keep same finalized head upon restart")

	// proposer runs pipeline, but now attached to the restarted geth node
	proposer.ActL2PipelineFull(t)
	require.Equal(t, statusBeforeRestart.UnsafeL2, proposer.L2Unsafe(), "expecting to keep same unsafe head upon restart")
	require.Equal(t, statusBeforeRestart.SafeL2, proposer.L2Safe(), "expecting the safe block to catch up to what it was before shutdown after syncing from L1, and not be stuck at the finalized block")
}

// TestConflictingL2Blocks tests that a second copy of the proposer stack cannot introduce an alternative
// L2 block (compared to something already secured by the first proposer):
// the alt block is not synced by the syncer, in unsafe and safe sync modes.
func TestConflictingL2Blocks(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)

	sd, _, miner, proposer, propEng, syncer, _, batcher := setupReorgTestActors(t, dp, sd, log)

	// Extra setup: a full alternative proposer, proposer engine, and batcher
	jwtPath := e2eutils.WriteDefaultJWT(t)
	altPropEng := NewL2Engine(t, log, sd.L2Cfg, sd.RollupCfg.Genesis.L1, jwtPath)
	altPropEngCl, err := sources.NewEngineClient(altPropEng.RPCClient(), log, nil, sources.EngineClientDefaultConfig(sd.RollupCfg))
	require.NoError(t, err)
	l1F, err := sources.NewL1Client(miner.RPCClient(), log, nil, sources.L1ClientDefaultConfig(sd.RollupCfg, false, sources.RPCKindBasic))
	require.NoError(t, err)
	altProposer := NewL2Proposer(t, log, l1F, altPropEngCl, sd.RollupCfg, 0)
	altBatcher := NewL2Batcher(log, sd.RollupCfg, &BatcherCfg{
		MinL1TxSize: 0,
		MaxL1TxSize: 128_000,
		BatcherKey:  dp.Secrets.Batcher,
	}, altProposer.RollupClient(), miner.EthClient(), altPropEng.EthClient())

	// And set up user Alice, using the alternative proposer endpoint
	l2Cl := altPropEng.EthClient()
	addresses := e2eutils.CollectAddresses(sd, dp)
	l2UserEnv := &BasicUserEnv[*L2Bindings]{
		EthCl:          l2Cl,
		Signer:         types.LatestSigner(sd.L2Cfg.Config),
		AddressCorpora: addresses,
		Bindings:       NewL2Bindings(t, l2Cl, altPropEng.GethClient()),
	}
	alice := NewCrossLayerUser(log, dp.Secrets.Alice, rand.New(rand.NewSource(1234)), sd.RollupCfg)
	alice.L2.SetUserEnv(l2UserEnv)

	proposer.ActL2PipelineFull(t)
	syncer.ActL2PipelineFull(t)
	altProposer.ActL2PipelineFull(t)

	// build empty L1 block
	miner.ActEmptyBlock(t)

	// Create L2 blocks, and reference the L1 head as origin
	proposer.ActL1HeadSignal(t)
	proposer.ActBuildToL1Head(t)

	// submit all new L2 blocks
	batcher.ActSubmitAll(t)

	// new L1 block with L2 batch
	miner.ActL1StartBlock(12)(t)
	miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
	miner.ActL1EndBlock(t)

	// syncer picks up the L2 chain that was submitted
	syncer.ActL1HeadSignal(t)
	syncer.ActL2PipelineFull(t)
	syncerHead := syncer.L2Unsafe()
	require.Equal(t, syncer.L2Safe(), proposer.L2Unsafe(), "syncer syncs from proposer via L1")
	require.Equal(t, syncer.L2Safe(), syncerHead, "syncer head is the same as that what was derived from L1")
	require.NotEqual(t, proposer.L2Safe(), proposer.L2Unsafe(), "proposer has not processed L1 yet")

	require.Less(t, altProposer.L2Unsafe().L1Origin.Number, proposer.L2Unsafe().L1Origin.Number, "alt-proposer is behind")

	// produce a conflicting L2 block with the alt proposer:
	// a new unsafe block that should not replace the existing safe block at the same height
	altProposer.ActL2StartBlock(t)
	// include tx to force the L2 block to really be different than the previous empty block
	alice.L2.ActResetTxOpts(t)
	alice.L2.ActSetTxToAddr(&dp.Addresses.Bob)(t)
	alice.L2.ActMakeTx(t)
	altPropEng.ActL2IncludeTx(alice.Address())(t)
	altProposer.ActL2EndBlock(t)

	conflictBlock := propEng.l2Chain.GetBlockByNumber(altProposer.L2Unsafe().Number)
	require.NotEqual(t, conflictBlock.Hash(), altProposer.L2Unsafe().Hash, "alt proposer has built a conflicting block")

	// give the unsafe block to the syncer, and see if it reorgs because of any unsafe inputs
	head, err := altPropEngCl.PayloadByLabel(t.Ctx(), eth.Unsafe)
	require.NoError(t, err)
	syncer.ActL2UnsafeGossipReceive(head)

	// make sure syncer has processed everything
	syncer.ActL2PipelineFull(t)

	// check if syncer is still following safe chain
	require.Equal(t, syncer.L2Unsafe(), syncerHead, "syncer must not accept the unsafe payload that orphans the safe payload")

	// now submit it to L1, and see if the syncer respects the inclusion order and preserves the original block
	altBatcher.ActSubmitAll(t)
	// include it in L1
	miner.ActL1StartBlock(12)(t)
	miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
	miner.ActL1EndBlock(t)
	l1Number := miner.l1Chain.CurrentHeader().Number.Uint64()

	// show latest L1 block with new batch data to syncer, and make it sync.
	syncer.ActL1HeadSignal(t)
	syncer.ActL2PipelineFull(t)
	require.Equal(t, syncer.SyncStatus().CurrentL1.Number, l1Number, "syncer has synced all new L1 blocks")
	require.Equal(t, syncer.L2Unsafe(), syncerHead, "syncer sticks to first included L2 block")

	// Now make the alt proposer aware of the L1 chain and derive the L2 chain like the syncer;
	// it should reorg out its conflicting blocks to get back in harmony with the syncer.
	altProposer.ActL1HeadSignal(t)
	altProposer.ActL2PipelineFull(t)
	require.Equal(t, syncer.L2Unsafe(), altProposer.L2Unsafe(), "alt-proposer gets back in harmony with syncer by reorging out its conflicting data")
	require.Equal(t, proposer.L2Unsafe(), altProposer.L2Unsafe(), "and gets back in harmony with original proposer")
}
