package actions

import (
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
	"github.com/kroma-network/kroma/op-e2e/e2eutils/validator"
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
	rt := defaultRuntime(gt, setupSequencerTest, deltaTimeOffset)
	rt.setupHonestValidator(false)

	// bind contracts
	rt.bindContracts()

	// submit outputs
	rt.setupOutputSubmitted(validator.ValidatorV1)

	checkRightOutputSubmitted(rt.t, rt.outputOracleContract, rt.sequencer, rt.seqEngine)
}

func RunValidatorManagerTest(gt *testing.T, deltaTimeOffset *hexutil.Uint64) {
	rt := defaultRuntime(gt, setupSequencerTest, deltaTimeOffset)

	// Redeploy and upgrade ValidatorPool to set the termination index to a smaller value for ValidatorManager test
	rt.assertRedeployValPoolToTerminate(defaultValPoolTerminationIndex)

	rt.setupHonestValidator(false)

	// bind contracts
	rt.bindContracts()

	rt.proceedWithBlocks(6)

	rt.depositToValPool(rt.validator)

	// Submit outputs to ValidatorPool until newTerminationIndex
	for i := uint64(0); new(big.Int).SetUint64(i).Cmp(defaultValPoolTerminationIndex) <= 0; i++ {
		rt.submitL2Output()
	}

	// assert if the ValidatorPool is terminated
	isValPoolTerminated := rt.validator.isValPoolTerminated(rt.t)
	require.True(rt.t, isValPoolTerminated, "ValPool should be terminated")

	rt.registerToValMgr(rt.validator)

	// create l2 output submission transactions until there is nothing left to submit
	submitAfterTransition := false
	for {
		waitTime := rt.validator.CalculateWaitTime(rt.t)
		if waitTime > 0 {
			break
		}
		rt.submitL2Output()
		submitAfterTransition = true
	}

	// Assert validator submitted at least one output after transition
	require.True(rt.t, submitAfterTransition)

	checkRightOutputSubmitted(rt.t, rt.outputOracleContract, rt.sequencer, rt.seqEngine)
}

// checkRightOutputSubmitted checks that L1 stored the expected output root
func checkRightOutputSubmitted(t StatefulTesting, outputOracleContract *bindings.L2OutputOracle, sequencer *L2Sequencer, seqEngine *L2Engine) {
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
