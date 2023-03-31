package actions

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/wemixkanvas/kanvas/components/node/rollup/derive"
	"github.com/wemixkanvas/kanvas/components/node/testlog"
	"github.com/wemixkanvas/kanvas/e2e/e2eutils"
)

func setupSyncer(t Testing, sd *e2eutils.SetupData, log log.Logger, l1F derive.L1Fetcher) (*L2Engine, *L2Syncer) {
	jwtPath := e2eutils.WriteDefaultJWT(t)
	engine := NewL2Engine(t, log, sd.L2Cfg, sd.RollupCfg.Genesis.L1, jwtPath)
	engCl := engine.EngineClient(t, sd.RollupCfg)
	syncer := NewL2Syncer(t, log, l1F, engCl, sd.RollupCfg)
	return engine, syncer
}

func setupSyncerOnlyTest(t Testing, sd *e2eutils.SetupData, log log.Logger) (*L1Miner, *L2Engine, *L2Syncer) {
	miner := NewL1Miner(t, log, sd.L1Cfg)
	l1Cl := miner.L1Client(t, sd.RollupCfg)
	engine, syncer := setupSyncer(t, sd, log, l1Cl)
	return miner, engine, syncer
}

func TestL2Syncer_ProposerWindow(gt *testing.T) {
	t := NewDefaultTesting(gt)
	p := &e2eutils.TestParams{
		MaxProposerDrift:   10,
		ProposerWindowSize: 24,
		ChannelTimeout:     10,
	}
	dp := e2eutils.MakeDeployParams(t, p)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)
	miner, engine, syncer := setupSyncerOnlyTest(t, sd, log)
	miner.ActL1SetFeeRecipient(common.Address{'A'})

	// Make two proposer windows worth of empty L1 blocks. After we pass the first proposer window, the L2 chain should get blocks
	for miner.l1Chain.CurrentBlock().NumberU64() < sd.RollupCfg.ProposerWindowSize*2 {
		miner.ActL1StartBlock(10)(t)
		miner.ActL1EndBlock(t)

		syncer.ActL2PipelineFull(t)

		l1Head := miner.l1Chain.CurrentBlock().NumberU64()
		expectedL1Origin := uint64(0)
		// as soon as we complete the proposer window, we force-adopt the L1 origin
		if l1Head >= sd.RollupCfg.ProposerWindowSize {
			expectedL1Origin = l1Head - sd.RollupCfg.ProposerWindowSize
		}
		require.Equal(t, expectedL1Origin, syncer.SyncStatus().SafeL2.L1Origin.Number, "L1 origin is forced in, given enough L1 blocks pass by")
		require.LessOrEqual(t, miner.l1Chain.GetBlockByNumber(expectedL1Origin).Time(), engine.l2Chain.CurrentBlock().Time(), "L2 time higher than L1 origin time")
	}
	tip2N := syncer.SyncStatus()

	// Do a deep L1 reorg as deep as a proposer window, this should affect the safe L2 chain
	miner.ActL1RewindDepth(sd.RollupCfg.ProposerWindowSize)(t)

	// Without new L1 block, the L1 appears to not be synced, and the node shouldn't reorg
	syncer.ActL2PipelineFull(t)
	require.Equal(t, tip2N.SafeL2, syncer.SyncStatus().SafeL2, "still the same after syncer work")

	// Make a new empty L1 block with different data than there was before.
	miner.ActL1SetFeeRecipient(common.Address{'B'})
	miner.ActL1StartBlock(10)(t)
	miner.ActL1EndBlock(t)
	reorgL1Block := miner.l1Chain.CurrentBlock()

	// Still no reorg, we need more L1 blocks first, before the reorged L1 block is forced in by proposer window
	syncer.ActL2PipelineFull(t)
	require.Equal(t, tip2N.SafeL2, syncer.SyncStatus().SafeL2)

	for miner.l1Chain.CurrentBlock().NumberU64() < sd.RollupCfg.ProposerWindowSize*2 {
		miner.ActL1StartBlock(10)(t)
		miner.ActL1EndBlock(t)
	}

	// workaround: in L1Traversal we only recognize the reorg once we see origin N+1, we don't reorg to shorter L1 chains
	miner.ActL1StartBlock(10)(t)
	miner.ActL1EndBlock(t)

	// Now it will reorg
	syncer.ActL2PipelineFull(t)

	got := miner.l1Chain.GetBlockByHash(miner.l1Chain.GetBlockByHash(syncer.SyncStatus().SafeL2.L1Origin.Hash).Hash())
	require.Equal(t, reorgL1Block.Hash(), got.Hash(), "must have reorged L2 chain to the new L1 chain")
}
