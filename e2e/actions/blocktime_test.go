package actions

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/components/node/testlog"
	"github.com/kroma-network/kroma/e2e/e2eutils"
)

// TestBatchInLastPossibleBlocks tests that the derivation pipeline
// accepts a batch that is included in the last possible L1 block
// where there are also no other batches included in the proposer
// window.
// This is a regression test against the bug fixed in PR #4566
func TestBatchInLastPossibleBlocks(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	dp.DeployConfig.ProposerWindowSize = 4
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)

	sd, _, miner, proposer, proposerEngine, _, _, batcher := setupReorgTestActors(t, dp, sd, log)

	signer := types.LatestSigner(sd.L2Cfg.Config)
	cl := proposerEngine.EthClient()
	aliceTx := func() {
		n, err := cl.PendingNonceAt(t.Ctx(), dp.Addresses.Alice)
		require.NoError(t, err)
		tx := types.MustSignNewTx(dp.Secrets.Alice, signer, &types.DynamicFeeTx{
			ChainID:   sd.L2Cfg.Config.ChainID,
			Nonce:     n,
			GasTipCap: big.NewInt(2 * params.GWei),
			GasFeeCap: new(big.Int).Add(miner.l1Chain.CurrentBlock().BaseFee, big.NewInt(2*params.GWei)),
			Gas:       params.TxGas,
			To:        &dp.Addresses.Bob,
			Value:     e2eutils.Ether(2),
		})
		require.NoError(gt, cl.SendTransaction(t.Ctx(), tx))
	}
	makeL2BlockWithAliceTx := func() {
		aliceTx()
		proposer.ActL2StartBlock(t)
		proposerEngine.ActL2IncludeTx(dp.Addresses.Alice)(t) // include a test tx from alice
		proposer.ActL2EndBlock(t)
	}
	verifyChainStateOnProposer := func(l1Number, unsafeHead, unsafeHeadOrigin, safeHead, safeHeadOrigin uint64) {
		require.Equal(t, l1Number, miner.l1Chain.CurrentHeader().Number.Uint64())
		require.Equal(t, unsafeHead, proposer.L2Unsafe().Number)
		require.Equal(t, unsafeHeadOrigin, proposer.L2Unsafe().L1Origin.Number)
		require.Equal(t, safeHead, proposer.L2Safe().Number)
		require.Equal(t, safeHeadOrigin, proposer.L2Safe().L1Origin.Number)
	}

	// Make 8 L1 blocks & 17 L2 blocks.
	miner.ActL1StartBlock(4)(t)
	miner.ActL1EndBlock(t)
	proposer.ActL1HeadSignal(t)
	proposer.ActL2PipelineFull(t)
	makeL2BlockWithAliceTx()
	makeL2BlockWithAliceTx()
	makeL2BlockWithAliceTx()

	for i := 0; i < 7; i++ {
		batcher.ActSubmitAll(t)
		miner.ActL1StartBlock(4)(t)
		miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
		miner.ActL1EndBlock(t)
		proposer.ActL1HeadSignal(t)
		proposer.ActL2PipelineFull(t)
		makeL2BlockWithAliceTx()
		makeL2BlockWithAliceTx()
	}

	// 8 L1 blocks with 17 L2 blocks is the unsafe state.
	// Because we consistently batch submitted we are one epoch behind the unsafe head with the safe head
	verifyChainStateOnProposer(8, 17, 8, 15, 7)

	// Create the batch for L2 blocks 16 & 17
	batcher.ActSubmitAll(t)

	// L1 Block 8 contains the batch for L2 blocks 14 & 15
	// Then we create L1 blocks 9, 10, 11
	// The L1 origin of L2 block 16 is L1 block 8
	// At a proposer window of 4, should be possible to include the batch for L2 block 16 & 17 at L1 block 12

	// Make 3 more L1 + 6 L2 blocks
	for i := 0; i < 3; i++ {
		miner.ActL1StartBlock(4)(t)
		miner.ActL1EndBlock(t)
		proposer.ActL1HeadSignal(t)
		proposer.ActL2PipelineFull(t)
		makeL2BlockWithAliceTx()
		makeL2BlockWithAliceTx()
	}

	// At this point verify that we have not started auto generating blocks
	// by checking that L1 & the unsafe head have advanced as expected, but the safe head is the same.
	verifyChainStateOnProposer(11, 23, 11, 15, 7)

	// Check that the batch can go in on the last block of the proposer window
	miner.ActL1StartBlock(4)(t)
	miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
	miner.ActL1EndBlock(t)
	proposer.ActL1HeadSignal(t)
	proposer.ActL2PipelineFull(t)

	// We have one more L1 block, no more unsafe blocks, but advance one
	// epoch on the safe head with the submitted batches
	verifyChainStateOnProposer(12, 23, 11, 17, 8)
}

// TestLargeL1Gaps tests the case that there is a gap between two L1 blocks which
// is larger than the proposer drift.
// This test has the following parameters:
// L1 Block time: 4s. L2 Block time: 2s. Proposer Drift: 32s
//
// It generates 8 L1 blocks & 16 L2 blocks.
// Then generates an L1 block that has a time delta of 48s.
// It then generates the 24 L2 blocks.
// Then it generates 3 more L1 blocks.
// At this point it can verify that the batches where properly generated.
// Note: It batches submits when possible.
func TestLargeL1Gaps(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	dp.DeployConfig.L1BlockTime = 4
	dp.DeployConfig.L2BlockTime = 2
	dp.DeployConfig.ProposerWindowSize = 4
	dp.DeployConfig.MaxProposerDrift = 32
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)

	sd, _, miner, proposer, proposerEngine, syncer, _, batcher := setupReorgTestActors(t, dp, sd, log)

	signer := types.LatestSigner(sd.L2Cfg.Config)
	cl := proposerEngine.EthClient()
	aliceTx := func() {
		n, err := cl.PendingNonceAt(t.Ctx(), dp.Addresses.Alice)
		require.NoError(t, err)
		tx := types.MustSignNewTx(dp.Secrets.Alice, signer, &types.DynamicFeeTx{
			ChainID:   sd.L2Cfg.Config.ChainID,
			Nonce:     n,
			GasTipCap: big.NewInt(2 * params.GWei),
			GasFeeCap: new(big.Int).Add(miner.l1Chain.CurrentBlock().BaseFee, big.NewInt(2*params.GWei)),
			Gas:       params.TxGas,
			To:        &dp.Addresses.Bob,
			Value:     e2eutils.Ether(2),
		})
		require.NoError(gt, cl.SendTransaction(t.Ctx(), tx))
	}
	makeL2BlockWithAliceTx := func() {
		aliceTx()
		proposer.ActL2StartBlock(t)
		proposerEngine.ActL2IncludeTx(dp.Addresses.Alice)(t) // include a test tx from alice
		proposer.ActL2EndBlock(t)
	}

	verifyChainStateOnProposer := func(l1Number, unsafeHead, unsafeHeadOrigin, safeHead, safeHeadOrigin uint64) {
		require.Equal(t, l1Number, miner.l1Chain.CurrentHeader().Number.Uint64())
		require.Equal(t, unsafeHead, proposer.L2Unsafe().Number)
		require.Equal(t, unsafeHeadOrigin, proposer.L2Unsafe().L1Origin.Number)
		require.Equal(t, safeHead, proposer.L2Safe().Number)
		require.Equal(t, safeHeadOrigin, proposer.L2Safe().L1Origin.Number)
	}

	verifyChainStateOnSyncer := func(l1Number, unsafeHead, unsafeHeadOrigin, safeHead, safeHeadOrigin uint64) {
		require.Equal(t, l1Number, miner.l1Chain.CurrentHeader().Number.Uint64())
		require.Equal(t, unsafeHead, syncer.L2Unsafe().Number)
		require.Equal(t, unsafeHeadOrigin, syncer.L2Unsafe().L1Origin.Number)
		require.Equal(t, safeHead, syncer.L2Safe().Number)
		require.Equal(t, safeHeadOrigin, syncer.L2Safe().L1Origin.Number)
	}

	// Make 8 L1 blocks & 16 L2 blocks.
	miner.ActL1StartBlock(4)(t)
	miner.ActL1EndBlock(t)
	proposer.ActL1HeadSignal(t)
	proposer.ActL2PipelineFull(t)
	makeL2BlockWithAliceTx()
	makeL2BlockWithAliceTx()

	for i := 0; i < 7; i++ {
		batcher.ActSubmitAll(t)
		miner.ActL1StartBlock(4)(t)
		miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
		miner.ActL1EndBlock(t)
		proposer.ActL1HeadSignal(t)
		proposer.ActL2PipelineFull(t)
		makeL2BlockWithAliceTx()
		makeL2BlockWithAliceTx()
	}

	n, err := cl.NonceAt(t.Ctx(), dp.Addresses.Alice, nil)
	require.NoError(t, err)
	require.Equal(t, uint64(16), n) // 16 valid blocks with txns.

	verifyChainStateOnProposer(8, 16, 8, 14, 7)

	// Make the really long L1 block. Do include previous batches
	batcher.ActSubmitAll(t)
	miner.ActL1StartBlock(48)(t)
	miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
	miner.ActL1EndBlock(t)
	proposer.ActL1HeadSignal(t)
	proposer.ActL2PipelineFull(t)

	verifyChainStateOnProposer(9, 16, 8, 16, 8)

	// Make the L2 blocks corresponding to the long L1 block
	for i := 0; i < 24; i++ {
		makeL2BlockWithAliceTx()
	}
	verifyChainStateOnProposer(9, 40, 9, 16, 8)

	// Check how many transactions from alice got included on L2
	// We created one transaction for every L2 block. So we should have created 40 transactions.
	// The first 16 L2 block where included without issue.
	// Then over the long block, 32s proposer drift / 2s block time => 16 blocks with transactions
	// Then at the last L2 block we reached the next origin, and accept txs again => 17 blocks with transactions
	// That leaves 7 L2 blocks without transactions. So we should have 16+17 = 33 transactions on chain.
	n, err = cl.PendingNonceAt(t.Ctx(), dp.Addresses.Alice)
	require.NoError(t, err)
	require.Equal(t, uint64(40), n)

	n, err = cl.NonceAt(t.Ctx(), dp.Addresses.Alice, nil)
	require.NoError(t, err)
	require.Equal(t, uint64(33), n)

	// Make more L1 blocks to get past the proposer window for the large range.
	// Do batch submit the previous L2 blocks.
	batcher.ActSubmitAll(t)
	miner.ActL1StartBlock(4)(t)
	miner.ActL1IncludeTx(sd.RollupCfg.Genesis.SystemConfig.BatcherAddr)(t)
	miner.ActL1EndBlock(t)

	// We are not able to do eager batch derivation for these L2 blocks because
	// we reject batches with a greater timestamp than the drift.
	verifyChainStateOnProposer(10, 40, 9, 16, 8)

	for i := 0; i < 2; i++ {
		miner.ActL1StartBlock(4)(t)
		miner.ActL1EndBlock(t)
	}

	// Run the pipeline against the batches + to be auto-generated batches.
	proposer.ActL1HeadSignal(t)
	proposer.ActL2PipelineFull(t)
	verifyChainStateOnProposer(12, 40, 9, 40, 9)

	// Recheck nonce. Will fail if no batches where submitted
	n, err = cl.NonceAt(t.Ctx(), dp.Addresses.Alice, nil)
	require.NoError(t, err)
	require.Equal(t, uint64(33), n) // 16 valid blocks with txns. Get proposer drift non-empty (32/2 => 16) & 7 forced empty

	// Check that the syncer got the same result
	syncer.ActL1HeadSignal(t)
	syncer.ActL2PipelineFull(t)
	verifyChainStateOnSyncer(12, 40, 9, 40, 9)
	require.Equal(t, syncer.L2Safe(), proposer.L2Safe())
}
