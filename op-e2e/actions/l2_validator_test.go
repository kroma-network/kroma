package actions

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/kroma-network/kroma/kroma-bindings/bindings"
)

// TestValidatorBatchType run each validator-related test case in singular batch mode and span batch mode.
func TestValidatorBatchType(t *testing.T) {
	tests := []struct {
		name string
		f    func(gt *testing.T, deltaTimeOffset *hexutil.Uint64)
	}{
		{"RunValidatorPoolTest", RunValidatorPoolTest},
		{"RunValidatorManagerTest", RunValidatorManagerTest},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name+"_SingularBatch", func(t *testing.T) {
			test.f(t, nil)
		})
	}

	deltaTimeOffset := hexutil.Uint64(0)
	for _, test := range tests {
		test := test
		t.Run(test.name+"_SpanBatch", func(t *testing.T) {
			test.f(t, &deltaTimeOffset)
		})
	}
}

func RunValidatorPoolTest(gt *testing.T, deltaTimeOffset *hexutil.Uint64) {
	t := NewDefaultTesting(gt)

	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	dp.DeployConfig.L2GenesisDeltaTimeOffset = deltaTimeOffset
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LevelDebug)
	miner, seqEngine, sequencer := setupSequencerTest(t, sd, log)

	rollupSeqCl := sequencer.RollupClient()
	batcher := NewL2Batcher(log, sd.RollupCfg, DefaultBatcherCfg(dp),
		rollupSeqCl, miner.EthClient(), seqEngine.EthClient(), seqEngine.EngineClient(t, sd.RollupCfg))

	validator := NewL2Validator(t, log, &ValidatorCfg{
		OutputOracleAddr:     sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorPoolAddr:    sd.DeploymentsL1.ValidatorPoolProxy,
		ValidatorManagerAddr: sd.DeploymentsL1.ValidatorManagerProxy,
		AssetManagerAddr:     sd.DeploymentsL1.AssetManagerProxy,
		ColosseumAddr:        sd.DeploymentsL1.ColosseumProxy,
		SecurityCouncilAddr:  sd.DeploymentsL1.SecurityCouncilProxy,
		ValidatorKey:         dp.Secrets.TrustedValidator,
		AllowNonFinalized:    false,
	}, miner.EthClient(), seqEngine.EthClient(), sequencer.RollupClient())

	proceedWithBlocks(t, miner, sequencer, batcher, dp.Addresses.Batcher, 1)

	// deposit bond for validator
	validator.ActDeposit(t, 1_000)
	miner.includeL1Block(t, validator.address, 12)

	require.Equal(t, sequencer.SyncStatus().UnsafeL2, sequencer.SyncStatus().FinalizedL2)
	// create l2 output submission transactions until there is nothing left to submit
	for {
		waitTime := validator.CalculateWaitTime(t)
		if waitTime > 0 {
			break
		}
		submitOutput(t, validator, miner)
	}

	checkRightOutputSubmitted(t, sd.DeploymentsL1.L2OutputOracleProxy, miner, sequencer, seqEngine)
}

func RunValidatorManagerTest(gt *testing.T, deltaTimeOffset *hexutil.Uint64) {
	t := NewDefaultTesting(gt)

	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	dp.DeployConfig.L2GenesisDeltaTimeOffset = deltaTimeOffset
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LevelDebug)
	miner, seqEngine, sequencer := setupSequencerTest(t, sd, log)

	rollupSeqCl := sequencer.RollupClient()
	batcher := NewL2Batcher(log, sd.RollupCfg, DefaultBatcherCfg(dp),
		rollupSeqCl, miner.EthClient(), seqEngine.EthClient(), seqEngine.EngineClient(t, sd.RollupCfg))

	validator := NewL2Validator(t, log, &ValidatorCfg{
		OutputOracleAddr:     sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorPoolAddr:    sd.DeploymentsL1.ValidatorPoolProxy,
		ValidatorManagerAddr: sd.DeploymentsL1.ValidatorManagerProxy,
		AssetManagerAddr:     sd.DeploymentsL1.AssetManagerProxy,
		ColosseumAddr:        sd.DeploymentsL1.ColosseumProxy,
		SecurityCouncilAddr:  sd.DeploymentsL1.SecurityCouncilProxy,
		ValidatorKey:         dp.Secrets.TrustedValidator,
		AllowNonFinalized:    false,
	}, miner.EthClient(), seqEngine.EthClient(), sequencer.RollupClient())

	proceedWithBlocks(t, miner, sequencer, batcher, dp.Addresses.Batcher, 6)

	// deposit bond for validator
	validator.ActDeposit(t, 1_000)
	miner.includeL1Block(t, validator.address, 12)
	// Submit 16 outputs to ValidatorPool
	for i := 0; i < 16; i++ {
		submitOutput(t, validator, miner)
	}

	// assert if the ValidatorPool is terminated
	isValPoolTerminated := validator.isValPoolTerminated(t)
	require.True(t, isValPoolTerminated, "ValPool should be terminated")

	// approve governance token
	assets := new(big.Int).SetUint64(1_000)
	validator.ActApprove(t, assets)
	miner.includeL1Block(t, validator.address, 12)

	// register validator
	validator.ActRegisterValidator(t, assets)
	miner.includeL1Block(t, validator.address, 12)

	require.Equal(t, sequencer.SyncStatus().UnsafeL2, sequencer.SyncStatus().FinalizedL2)
	// create l2 output submission transactions until there is nothing left to submit
	for {
		waitTime := validator.CalculateWaitTime(t)
		if waitTime > 0 {
			break
		}
		submitOutput(t, validator, miner)
	}

	checkRightOutputSubmitted(t, sd.DeploymentsL1.L2OutputOracleProxy, miner, sequencer, seqEngine)
}

func proceedWithBlocks(t StatefulTesting, miner *L1Miner, sequencer *L2Sequencer, batcher *L2Batcher, batcherAddr common.Address, n int) {
	for i := 0; i < n; i++ {
		// L1 block
		miner.ActEmptyBlock(t)
		// L2 block
		sequencer.ActL1HeadSignal(t)
		sequencer.ActL2PipelineFull(t)
		sequencer.ActBuildToL1Head(t)
		// submit and include in L1
		batcher.ActSubmitAll(t)
		miner.includeL1Block(t, batcherAddr, 12)
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
}

func submitOutput(t StatefulTesting, validator *L2Validator, miner *L1Miner) {
	// submit to L1
	validator.ActSubmitL2Output(t)
	// include output on L1
	miner.includeL1Block(t, validator.address, 12)
	// Check submission was successful
	receipt, err := miner.EthClient().TransactionReceipt(t.Ctx(), validator.LastSubmitL2OutputTx())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "submission failed")
}

func checkRightOutputSubmitted(t StatefulTesting, l2OOAddr common.Address, miner *L1Miner, sequencer *L2Sequencer, seqEngine *L2Engine) {
	// check that L1 stored the expected output root
	outputOracleContract, err := bindings.NewL2OutputOracle(l2OOAddr, miner.EthClient())
	require.NoError(t, err)
	// NOTE: If Proto Dank Sharding is introduced, the below code fix may be restored.
	// block := sequencer.SyncStatus().FinalizedL2
	// outputOnL1, err := outputOracleContract.GetL2OutputAfter(nil, new(big.Int).SetUint64(block.Number))
	blockNum, err := outputOracleContract.LatestBlockNumber(&bind.CallOpts{})
	require.NoError(t, err)
	outputOnL1, err := outputOracleContract.GetL2OutputAfter(&bind.CallOpts{}, blockNum)
	require.NoError(t, err)
	block, err := seqEngine.EthClient().BlockByNumber(t.Ctx(), blockNum)
	require.NoError(t, err)
	require.Less(t, block.Time(), outputOnL1.Timestamp.Uint64(), "output is registered with L1 timestamp of L2 tx output submission, past L2 block")
	outputComputed, err := sequencer.RollupClient().OutputAtBlock(t.Ctx(), blockNum.Uint64())
	require.NoError(t, err)
	require.Equal(t, eth.Bytes32(outputOnL1.OutputRoot), outputComputed.OutputRoot, "output roots must match")
}
