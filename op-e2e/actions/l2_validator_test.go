package actions

import (
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/testlog"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
)

func TestValidator(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)
	miner, seqEngine, sequencer := setupSequencerTest(t, sd, log)

	rollupSeqCl := sequencer.RollupClient()
	batcher := NewL2Batcher(log, sd.RollupCfg, &BatcherCfg{
		MinL1TxSize: 0,
		MaxL1TxSize: 128_000,
		BatcherKey:  dp.Secrets.Batcher,
	}, rollupSeqCl, miner.EthClient(), seqEngine.EthClient())

	validator := NewL2Validator(t, log, &ValidatorCfg{
		OutputOracleAddr:    sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorPoolAddr:   sd.DeploymentsL1.ValidatorPoolProxy,
		ColosseumAddr:       sd.DeploymentsL1.ColosseumProxy,
		SecurityCouncilAddr: sd.DeploymentsL1.SecurityCouncilProxy,
		ValidatorKey:        dp.Secrets.TrustedValidator,
		AllowNonFinalized:   false,
	}, miner.EthClient(), seqEngine.EthClient(), sequencer.RollupClient())

	// NOTE(chokobole): It is necessary to wait for one finalized (or safe if AllowNonFinalized
	// config is set) block to pass after each submission interval before submitting the output
	// root. For example, if the submission interval is set to 1800 blocks, the output root can
	// only be submitted at 1801 finalized blocks. In fact, the following code is designed to
	// create one or more finalized L2 blocks in order to pass the test. If Proto Dank Sharding
	// is introduced, the below code fix may no longer be necessary.
	for i := 0; i < 2; i++ {
		// L1 block
		miner.ActEmptyBlock(t)
		// L2 block
		sequencer.ActL1HeadSignal(t)
		sequencer.ActL2PipelineFull(t)
		sequencer.ActBuildToL1Head(t)
		// submit and include in L1
		batcher.ActSubmitAll(t)
		miner.includeL1Block(t, dp.Addresses.Batcher)
		// finalize the first and second L1 blocks, including the batch
		miner.ActL1SafeNext(t)
		miner.ActL1SafeNext(t)
		miner.ActL1FinalizeNext(t)
		miner.ActL1FinalizeNext(t)
		// derive and see the L2 chain fully finalize
		sequencer.ActL2PipelineFull(t)
		sequencer.ActL1SafeSignal(t)
		sequencer.ActL1FinalizedSignal(t)
	}

	// deposit bond for validator
	validator.ActDeposit(t, 1_000)
	miner.includeL1Block(t, validator.address)

	require.Equal(t, sequencer.SyncStatus().UnsafeL2, sequencer.SyncStatus().FinalizedL2)
	// create l2 output submission transactions until there is nothing left to submit
	for {
		waitTime := validator.CalculateWaitTime(t)
		if waitTime > 0 {
			break
		}
		// and submit it to L1
		validator.ActSubmitL2Output(t)
		// include output on L1
		miner.includeL1Block(t, validator.address)
		miner.ActEmptyBlock(t)
		// Check submission was successful
		receipt, err := miner.EthClient().TransactionReceipt(t.Ctx(), validator.LastSubmitL2OutputTx())
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "submission failed")
	}

	// check that L1 stored the expected output root
	outputOracleContract, err := bindings.NewL2OutputOracle(sd.DeploymentsL1.L2OutputOracleProxy, miner.EthClient())
	require.NoError(t, err)
	// NOTE(chokobole): Comment these 2 lines because of the reason above.
	// If Proto Dank Sharding is introduced, the below code fix may be restored.
	// block := sequencer.SyncStatus().FinalizedL2
	// outputOnL1, err := outputOracleContract.GetL2OutputAfter(nil, new(big.Int).SetUint64(block.Number))
	blockNum, err := outputOracleContract.LatestBlockNumber(nil)
	require.NoError(t, err)
	outputOnL1, err := outputOracleContract.GetL2OutputAfter(nil, blockNum)
	require.NoError(t, err)
	block, err := seqEngine.EthClient().BlockByNumber(t.Ctx(), blockNum)
	require.NoError(t, err)
	require.Less(t, block.Time(), outputOnL1.Timestamp.Uint64(), "output is registered with L1 timestamp of L2 tx output submission, past L2 block")
	outputComputed, err := sequencer.RollupClient().OutputAtBlock(t.Ctx(), blockNum.Uint64())
	require.NoError(t, err)
	require.Equal(t, eth.Bytes32(outputOnL1.OutputRoot), outputComputed.OutputRoot, "output roots must match")
}
