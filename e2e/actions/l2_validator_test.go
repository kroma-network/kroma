package actions

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/bindings/bindings"
	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/testlog"
	"github.com/kroma-network/kroma/e2e/e2eutils"
)

func includeL1Block(t StatefulTesting, miner *L1Miner, sender common.Address) {
	miner.ActL1StartBlock(12)(t)
	miner.ActL1IncludeTx(sender)(t)
	miner.ActL1EndBlock(t)
}

func TestValidator(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)
	miner, propEngine, proposer := setupProposerTest(t, sd, log)

	rollupPropCl := proposer.RollupClient()
	batcher := NewL2Batcher(log, sd.RollupCfg, &BatcherCfg{
		MinL1TxSize: 0,
		MaxL1TxSize: 128_000,
		BatcherKey:  dp.Secrets.Batcher,
	}, rollupPropCl, miner.EthClient(), propEngine.EthClient())

	validator := NewL2Validator(t, log, &ValidatorCfg{
		OutputOracleAddr:  sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorKey:      dp.Secrets.Validator,
		AllowNonFinalized: false,
	}, miner.EthClient(), proposer.RollupClient())

	// NOTE(chokobole): After the Blue hard fork, it is necessary to wait for one finalized
	// (or safe if AllowNonFinalized config is set) block to pass after each submission interval
	// before submitting the output root.
	// For example, if the submission interval is set to 1800 blocks, before the Blue hard fork,
	// the output root could be submitted at 1800 finalized blocks. However, after the update,
	// the output root can only be submitted at 1801 finalized blocks.
	// In fact, the following code is designed to create one or more finalized L2 blocks
	// in order to pass the test after the Blue hard fork.
	// If Proto Dank Sharding is introduced, the below code fix may no longer be necessary.
	for i := 0; i < 2; i++ {
		// L1 block
		miner.ActEmptyBlock(t)
		// L2 block
		proposer.ActL1HeadSignal(t)
		proposer.ActL2PipelineFull(t)
		proposer.ActBuildToL1Head(t)
		// submit and include in L1
		batcher.ActSubmitAll(t)
		includeL1Block(t, miner, dp.Addresses.Batcher)
		// finalize the first and second L1 blocks, including the batch
		miner.ActL1SafeNext(t)
		miner.ActL1SafeNext(t)
		miner.ActL1FinalizeNext(t)
		miner.ActL1FinalizeNext(t)
		// derive and see the L2 chain fully finalize
		proposer.ActL2PipelineFull(t)
		proposer.ActL1SafeSignal(t)
		proposer.ActL1FinalizedSignal(t)
	}

	require.Equal(t, proposer.SyncStatus().UnsafeL2, proposer.SyncStatus().FinalizedL2)
	// create l2 output submission transactions until there is nothing left to submit
	for validator.CanSubmit(t) {
		// and submit it to L1
		validator.ActSubmitL2Output(t)
		// include output on L1
		includeL1Block(t, miner, dp.Addresses.Validator)
		// Check submission was successful
		receipt, err := miner.EthClient().TransactionReceipt(t.Ctx(), validator.LastSubmitL2OutputTx())
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "submission failed")
	}

	// check that L1 stored the expected output root
	outputOracleContract, err := bindings.NewL2OutputOracle(sd.DeploymentsL1.L2OutputOracleProxy, miner.EthClient())
	require.NoError(t, err)
	// NOTE(chokobole): Comment these 2 lines because of the reason above about the Blue hard fork.
	// If Proto Dank Sharding is introduced, the below code fix may be restored.
	// block := proposer.SyncStatus().FinalizedL2
	// outputOnL1, err := outputOracleContract.GetL2OutputAfter(nil, new(big.Int).SetUint64(block.Number))
	blockNum, err := outputOracleContract.LatestBlockNumber(nil)
	require.NoError(t, err)
	outputOnL1, err := outputOracleContract.GetL2OutputAfter(nil, blockNum)
	require.NoError(t, err)
	block, err := propEngine.EthClient().BlockByNumber(t.Ctx(), blockNum)
	require.NoError(t, err)
	require.Less(t, block.Time(), outputOnL1.Timestamp.Uint64(), "output is registered with L1 timestamp of L2 tx output submission, past L2 block")
	outputComputed, err := proposer.RollupClient().OutputAtBlock(t.Ctx(), blockNum.Uint64(), false)
	require.NoError(t, err)
	require.Equal(t, eth.Bytes32(outputOnL1.OutputRoot), outputComputed.OutputRoot, "output roots must match")
}
