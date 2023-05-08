package actions

import (
	"testing"

	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/components/node/testlog"
	"github.com/kroma-network/kroma/e2e/e2eutils"
)

func TestShapellaL1Fork(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)

	sd := e2eutils.Setup(t, dp, defaultAlloc)
	activation := sd.L1Cfg.Timestamp + 24
	sd.L1Cfg.Config.ShanghaiTime = &activation
	log := testlog.Logger(t, log.LvlDebug)

	_, _, miner, proposer, _, syncer, _, batcher := setupReorgTestActors(t, dp, sd, log)

	require.False(t, sd.L1Cfg.Config.IsShanghai(miner.l1Chain.CurrentBlock().Time), "not active yet")

	// start nodes
	proposer.ActL2PipelineFull(t)
	syncer.ActL2PipelineFull(t)

	// build empty L1 blocks, crossing the fork boundary
	miner.ActEmptyBlock(t)
	miner.ActEmptyBlock(t)
	miner.ActEmptyBlock(t)

	// verify Shanghai is active
	l1Block := miner.l1Chain.CurrentBlock()
	require.True(t, sd.L1Cfg.Config.IsShanghai(l1Block.Time))

	// build L2 chain up to and including L2 blocks referencing shanghai L1 blocks
	proposer.ActL1HeadSignal(t)
	proposer.ActBuildToL1Head(t)
	miner.ActL1StartBlock(12)(t)
	batcher.ActSubmitAll(t)
	miner.ActL1IncludeTx(batcher.batcherAddr)(t)
	miner.ActL1EndBlock(t)

	// sync syncer
	syncer.ActL1HeadSignal(t)
	syncer.ActL2PipelineFull(t)
	// verify syncer accepted shanghai L1 inputs
	require.Equal(t, l1Block.Hash(), syncer.SyncStatus().SafeL2.L1Origin.Hash, "syncer synced L1 chain that includes shanghai headers")
	require.Equal(t, proposer.SyncStatus().UnsafeL2, syncer.SyncStatus().UnsafeL2, "syncer and sequencer agree")
}
