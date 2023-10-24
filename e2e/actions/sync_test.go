package actions

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/rollup/derive"
	"github.com/kroma-network/kroma/components/node/rollup/sync"
	"github.com/kroma-network/kroma/components/node/sources"
	"github.com/kroma-network/kroma/components/node/testlog"
	"github.com/kroma-network/kroma/e2e/e2eutils"
)

func TestDerivationWithFlakyL1RPC(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlError) // mute all the temporary derivation errors that we forcefully create
	_, _, miner, proposer, _, syncer, _, batcher := setupReorgTestActors(t, dp, sd, log)

	rng := rand.New(rand.NewSource(1234))
	proposer.ActL2PipelineFull(t)
	syncer.ActL2PipelineFull(t)

	// build a L1 chain with 20 blocks and matching L2 chain and batches to test some derivation work
	miner.ActEmptyBlock(t)
	for i := 0; i < 20; i++ {
		proposer.ActL1HeadSignal(t)
		proposer.ActL2PipelineFull(t)
		proposer.ActBuildToL1Head(t)
		batcher.ActSubmitAll(t)
		miner.ActL1StartBlock(12)(t)
		miner.ActL1IncludeTx(batcher.batcherAddr)(t)
		miner.ActL1EndBlock(t)
	}
	// Make syncer aware of head
	syncer.ActL1HeadSignal(t)

	// Now make the L1 RPC very flaky: requests will randomly fail with 50% chance
	miner.MockL1RPCErrors(func() error {
		if rng.Intn(2) == 0 {
			return errors.New("mock rpc error")
		}
		return nil
	})

	// And sync the syncer
	syncer.ActL2PipelineFull(t)
	// syncer should be synced, even though it hit lots of temporary L1 RPC errors
	require.Equal(t, proposer.L2Unsafe(), syncer.L2Safe(), "syncer is synced")
}

func TestFinalizeWhileSyncing(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlError) // mute all the temporary derivation errors that we forcefully create
	_, _, miner, proposer, _, syncer, _, batcher := setupReorgTestActors(t, dp, sd, log)

	proposer.ActL2PipelineFull(t)
	syncer.ActL2PipelineFull(t)

	syncerStartStatus := syncer.SyncStatus()

	// Build an L1 chain with (FinalityDelay + 1) blocks, containing batches of L2 chain.
	// Enough to go past the FinalityDelay of the engine queue,
	// to make the syncer finalize while it syncs.
	miner.ActEmptyBlock(t)
	for i := 0; i < derive.FinalityDelay+1; i++ {
		proposer.ActL1HeadSignal(t)
		proposer.ActL2PipelineFull(t)
		proposer.ActBuildToL1Head(t)
		batcher.ActSubmitAll(t)
		miner.ActL1StartBlock(12)(t)
		miner.ActL1IncludeTx(batcher.batcherAddr)(t)
		miner.ActL1EndBlock(t)
	}
	l1Head := miner.l1Chain.CurrentHeader()
	// finalize all of L1
	miner.ActL1Safe(t, l1Head.Number.Uint64())
	miner.ActL1Finalize(t, l1Head.Number.Uint64())

	// Now signal L1 finality to the syncer, while the syncer is not synced.
	syncer.ActL1HeadSignal(t)
	syncer.ActL1SafeSignal(t)
	syncer.ActL1FinalizedSignal(t)

	// Now sync the syncer, without repeating the signal.
	// While it's syncing, it should finalize on interval now, based on the future L1 finalized block it remembered.
	syncer.ActL2PipelineFull(t)

	// Verify the syncer finalized something new
	require.Less(t, syncerStartStatus.FinalizedL2.Number, syncer.SyncStatus().FinalizedL2.Number, "syncer finalized L2 blocks during sync")
}

func TestUnsafeSync(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlInfo)

	sd, _, _, proposer, propEng, syncer, _, _ := setupReorgTestActors(t, dp, sd, log)
	propEngCl, err := sources.NewEngineClient(propEng.RPCClient(), log, nil, sources.EngineClientDefaultConfig(sd.RollupCfg))
	require.NoError(t, err)

	proposer.ActL2PipelineFull(t)
	syncer.ActL2PipelineFull(t)

	for i := 0; i < 10; i++ {
		// Build a L2 block
		proposer.ActL2StartBlock(t)
		proposer.ActL2EndBlock(t)
		// Notify new L2 block to syncer by unsafe gossip
		propHead, err := propEngCl.PayloadByLabel(t.Ctx(), eth.Unsafe)
		require.NoError(t, err)
		syncer.ActL2UnsafeGossipReceive(propHead)(t)
		// Handle unsafe payload
		syncer.ActL2PipelineFull(t)
		// Syncer must advance its unsafe head and engine sync target.
		require.Equal(t, proposer.L2Unsafe().Hash, syncer.L2Unsafe().Hash)
		// Check engine sync target updated.
		require.Equal(t, proposer.L2Unsafe().Hash, proposer.EngineSyncTarget().Hash)
		require.Equal(t, syncer.L2Unsafe().Hash, syncer.EngineSyncTarget().Hash)
	}
}

func TestEngineP2PSync(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlInfo)

	miner, propEng, proposer := setupProposerTest(t, sd, log)
	// Enable engine P2P sync
	_, syncer := setupSyncer(t, sd, log, miner.L1Client(t, sd.RollupCfg), &sync.Config{EngineSync: true})

	propEngCl, err := sources.NewEngineClient(propEng.RPCClient(), log, nil, sources.EngineClientDefaultConfig(sd.RollupCfg))
	require.NoError(t, err)

	proposer.ActL2PipelineFull(t)
	syncer.ActL2PipelineFull(t)

	syncerUnsafeHead := syncer.L2Unsafe()

	// Build a L2 block. This block will not be gossiped to syncer, so syncer can not advance chain by itself.
	proposer.ActL2StartBlock(t)
	proposer.ActL2EndBlock(t)

	for i := 0; i < 10; i++ {
		// Build a L2 block
		proposer.ActL2StartBlock(t)
		proposer.ActL2EndBlock(t)
		// Notify new L2 block to syncer by unsafe gossip
		propHead, err := propEngCl.PayloadByLabel(t.Ctx(), eth.Unsafe)
		require.NoError(t, err)
		syncer.ActL2UnsafeGossipReceive(propHead)(t)
		// Handle unsafe payload
		syncer.ActL2PipelineFull(t)
		// Syncer must advance only engine sync target.
		require.NotEqual(t, proposer.L2Unsafe().Hash, syncer.L2Unsafe().Hash)
		require.NotEqual(t, syncer.L2Unsafe().Hash, syncer.EngineSyncTarget().Hash)
		require.Equal(t, syncer.L2Unsafe().Hash, syncerUnsafeHead.Hash)
		require.Equal(t, proposer.L2Unsafe().Hash, syncer.EngineSyncTarget().Hash)
	}
}
