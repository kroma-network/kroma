package actions

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	val "github.com/kroma-network/kroma/kroma-validator"
	chal "github.com/kroma-network/kroma/kroma-validator/challenge"
	valhelper "github.com/kroma-network/kroma/op-e2e/e2eutils/validator"
	"github.com/kroma-network/kroma/op-e2e/testdata"
)

var challengerTests = []struct {
	name string
	f    func(ft *testing.T, deltaTimeOffset *hexutil.Uint64, version uint8)
}{
	{"ChallengeBasic", ChallengeBasic},
	{"ChallengeAsserterBisectTimeout", ChallengeAsserterBisectTimeout},
	{"ChallengeChallengerBisectTimeout", ChallengeChallengerBisectTimeout},
	{"ChallengeChallengerProvingTimeout", ChallengeChallengerProvingTimeout},
	{"ChallengeInvalidProofFail", ChallengeInvalidProofFail},
	{"ChallengeForceDeleteOutputBySecurityCouncil", ChallengeForceDeleteOutputBySecurityCouncil},
	{"MultipleChallenges", MultipleChallenges},
}

// TestChallengerBatchType run each challenger-related test case in singular batch mode and span batch mode.
func TestChallengerBatchType(t *testing.T) {
	for _, test := range challengerTests {
		test := test
		t.Run(test.name+"_SingularBatch", func(t *testing.T) {
			test.f(t, nil, valhelper.ValidatorV1)
		})
	}

	deltaTimeOffset := hexutil.Uint64(0)
	for _, test := range challengerTests {
		test := test
		t.Run(test.name+"_SpanBatch", func(t *testing.T) {
			test.f(t, &deltaTimeOffset, valhelper.ValidatorV1)
		})
	}
}

// TestValidatorSystemVersion run each challenge test case in ValidatorPool version and ValidatorManager version.
func TestValidatorSystemVersion(t *testing.T) {
	for _, test := range challengerTests {
		test := test
		t.Run(test.name+"_ValidatorPool", func(t *testing.T) {
			test.f(t, nil, valhelper.ValidatorV1)
		})
	}
	for _, test := range challengerTests {
		test := test
		t.Run(test.name+"_ValidatorManager", func(t *testing.T) {
			test.f(t, nil, valhelper.ValidatorV2)
		})
	}
}

func ChallengeBasic(t *testing.T, deltaTimeOffset *hexutil.Uint64, version uint8) {
	rt := defaultRuntime(t, setupSequencerTest, deltaTimeOffset)

	if version == valhelper.ValidatorV2 {
		rt.assertRedeployValPoolToTerminate(defaultValPoolTerminationIndex)
	}

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger1()
	rt.setupHonestGuardian()

	// bind contracts
	rt.bindContracts()

	// submit outputs
	rt.setupOutputSubmitted(version)

	// create challenge
	rt.setupChallenge(rt.challenger1, version)

	var beforeAsset, beforeAssetBonded *big.Int
	if version == valhelper.ValidatorV2 {
		_, _, beforeAsset, beforeAssetBonded, _ = rt.fetchValidatorStatus(rt.validator)
	}

interaction:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			rt.txHash = rt.challenger1.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, false)
			rt.includeL1BlockBySender(rt.challenger1.address)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			rt.txHash = rt.validator.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, true)
			rt.includeL1BlockBySender(rt.validator.address)
		case chal.StatusReadyToProve:
			rt.txHash = rt.challenger1.ActProveFault(rt.t, rt.outputIndex, false)
			rt.includeL1BlockBySender(rt.challenger1.address)
		case chal.StatusNone:
			// guardian validates deleted output by challenger is invalid after challenge is proven
			outputBlockNum := rt.outputOnL1.L2BlockNumber.Uint64()
			isEqual := rt.guardian.ActValidateL2Output(rt.t, rt.outputOnL1.OutputRoot, outputBlockNum)
			require.False(rt.t, isEqual, "deleted output is expected not equal but actually equal")
			break interaction
		default:
			break interaction
		}

		// check whether the submission was successful
		rt.receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
		require.NoError(rt.t, err)
		require.Equal(rt.t, types.ReceiptStatusSuccessful, rt.receipt.Status, "failed to progress interactive fault proof")
	}

	// check output is deleted
	remoteOutput, err := rt.outputOracleContract.GetL2Output(nil, rt.outputIndex)
	require.NoError(rt.t, err, "unable to get l2 output")
	outputDeleted := val.IsOutputDeleted(remoteOutput.OutputRoot)
	require.True(rt.t, outputDeleted, "invalid output is not deleted")

	// check output submitter is changed to challenger
	require.Equal(rt.t, remoteOutput.Submitter, rt.challenger1.address)

	// check the status of challenge is StatusNone(0)
	status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
	require.NoError(rt.t, err)
	require.Equal(rt.t, chal.StatusNone, status)

	if version == valhelper.ValidatorV1 {
		// check bond amount doubled after challenge proven
		bond, err := rt.valPoolContract.GetBond(nil, rt.outputIndex)
		require.NoError(rt.t, err)
		require.Equal(rt.t, big.NewInt(2*rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt().Int64()), bond.Amount)
	} else if version == valhelper.ValidatorV2 {
		// check asserter has been slashed
		valStatus, inJail, afterAsset, afterAssetBonded, slashingAmount := rt.fetchValidatorStatus(rt.validator)
		require.Equal(rt.t, val.StatusReady, valStatus)
		require.True(rt.t, inJail)
		require.Equal(rt.t, beforeAsset.Uint64()-slashingAmount.Uint64(), afterAsset.Uint64())
		require.Equal(rt.t, beforeAssetBonded.Uint64()-slashingAmount.Uint64(), afterAssetBonded.Uint64())
	}
}

func ChallengeAsserterBisectTimeout(t *testing.T, deltaTimeOffset *hexutil.Uint64, version uint8) {
	rt := defaultRuntime(t, setupSequencerTest, deltaTimeOffset)

	if version == valhelper.ValidatorV2 {
		rt.assertRedeployValPoolToTerminate(defaultValPoolTerminationIndex)
	}

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger1()
	rt.setupHonestGuardian()

	// bind contracts
	rt.bindContracts()

	// submit outputs
	rt.setupOutputSubmitted(version)

	// create challenge
	rt.setupChallenge(rt.challenger1, version)

	var beforeAsset, beforeAssetBonded *big.Int
	if version == valhelper.ValidatorV2 {
		_, _, beforeAsset, beforeAssetBonded, _ = rt.fetchValidatorStatus(rt.validator)
	}

interaction:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusAsserterTurn:
			// do nothing to trigger asserter timeout
			rt.miner.ActEmptyBlock(rt.t)
		case chal.StatusAsserterTimeout:
			rt.txHash = rt.challenger1.ActProveFault(rt.t, rt.outputIndex, true)
			rt.includeL1BlockBySender(rt.challenger1.address)
		case chal.StatusNone:
			// guardian validates deleted output by challenger is invalid after challenge is proven
			outputBlockNum := rt.outputOnL1.L2BlockNumber.Uint64()
			isEqual := rt.guardian.ActValidateL2Output(rt.t, rt.outputOnL1.OutputRoot, outputBlockNum)
			require.False(rt.t, isEqual, "deleted output is expected not equal but actually equal")
			break interaction
		default:
			break interaction
		}

		// check whether the submission was successful
		rt.receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
		require.NoError(rt.t, err)
		require.Equal(rt.t, types.ReceiptStatusSuccessful, rt.receipt.Status, "failed to progress interactive fault proof")
	}

	// check output is deleted
	remoteOutput, err := rt.outputOracleContract.GetL2Output(nil, rt.outputIndex)
	require.NoError(rt.t, err, "unable to get l2 output")
	outputDeleted := val.IsOutputDeleted(remoteOutput.OutputRoot)
	require.True(rt.t, outputDeleted, "invalid output is not deleted")

	// check output submitter is changed to challenger
	require.Equal(rt.t, remoteOutput.Submitter, rt.challenger1.address)

	// check the status of challenge is StatusNone(0)
	status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
	require.NoError(rt.t, err)
	require.Equal(rt.t, chal.StatusNone, status)

	if version == valhelper.ValidatorV1 {
		// check bond amount doubled after challenge proven
		bond, err := rt.valPoolContract.GetBond(nil, rt.outputIndex)
		require.NoError(rt.t, err)
		require.Equal(rt.t, big.NewInt(2*rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt().Int64()), bond.Amount)
	} else if version == valhelper.ValidatorV2 {
		// check asserter has been slashed
		valStatus, inJail, afterAsset, afterAssetBonded, slashingAmount := rt.fetchValidatorStatus(rt.validator)
		require.Equal(rt.t, val.StatusReady, valStatus)
		require.True(rt.t, inJail)
		require.Equal(rt.t, beforeAsset.Uint64()-slashingAmount.Uint64(), afterAsset.Uint64())
		require.Equal(rt.t, beforeAssetBonded.Uint64()-slashingAmount.Uint64(), afterAssetBonded.Uint64())
	}
}

func ChallengeChallengerBisectTimeout(t *testing.T, deltaTimeOffset *hexutil.Uint64, version uint8) {
	rt := defaultRuntime(t, setupSequencerTest, deltaTimeOffset)

	if version == valhelper.ValidatorV2 {
		rt.assertRedeployValPoolToTerminate(defaultValPoolTerminationIndex)
	}

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger1()

	// bind contracts
	rt.bindContracts()

	// submit outputs
	rt.setupOutputSubmitted(version)

	// create challenge
	rt.setupChallenge(rt.challenger1, version)

	var beforeAsset, beforeAssetBonded *big.Int
	if version == valhelper.ValidatorV2 {
		_, _, beforeAsset, beforeAssetBonded, _ = rt.fetchValidatorStatus(rt.challenger1)
	}

interaction:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// do nothing to trigger challenger timeout
			rt.miner.ActEmptyBlock(rt.t)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			rt.txHash = rt.validator.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, true)
			rt.includeL1BlockBySender(rt.validator.address)
		case chal.StatusChallengerTimeout:
			// call challenger timeout by validator
			rt.txHash = rt.validator.ActChallengerTimeout(rt.t, rt.outputIndex, rt.challenger1.address)
			rt.includeL1BlockBySender(rt.validator.address)
		default:
			break interaction
		}

		// check whether the submission was successful
		rt.receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
		require.NoError(rt.t, err)
		require.Equal(rt.t, types.ReceiptStatusSuccessful, rt.receipt.Status, "failed to progress interactive fault proof")
	}

	// check challenge is not proven, output is remained
	remoteOutput, err := rt.outputOracleContract.GetL2Output(nil, rt.outputIndex)
	require.NoError(rt.t, err, "unable to get l2 output")
	outputDeleted := val.IsOutputDeleted(remoteOutput.OutputRoot)
	require.False(rt.t, outputDeleted, "output is deleted when not proven")

	// check output submitter is not changed
	require.Equal(rt.t, remoteOutput.Submitter, rt.validator.address)

	// check the status of challenge is StatusNone(0)
	status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
	require.NoError(rt.t, err)
	require.Equal(rt.t, chal.StatusNone, status)

	if version == valhelper.ValidatorV1 {
		// check bond amount doubled after challenger timed out
		bond, err := rt.valPoolContract.GetBond(nil, rt.outputIndex)
		require.NoError(rt.t, err)
		require.Equal(rt.t, big.NewInt(2*rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt().Int64()), bond.Amount)
	} else if version == valhelper.ValidatorV2 {
		// check challenger has been slashed
		valStatus, inJail, afterAsset, afterAssetBonded, slashingAmount := rt.fetchValidatorStatus(rt.challenger1)
		require.Equal(rt.t, val.StatusReady, valStatus)
		require.True(rt.t, inJail)
		require.Equal(rt.t, beforeAsset.Uint64()-slashingAmount.Uint64(), afterAsset.Uint64())
		require.Equal(rt.t, beforeAssetBonded.Uint64()-slashingAmount.Uint64(), afterAssetBonded.Uint64())
	}
}

func ChallengeChallengerProvingTimeout(t *testing.T, deltaTimeOffset *hexutil.Uint64, version uint8) {
	rt := defaultRuntime(t, setupSequencerTest, deltaTimeOffset)

	if version == valhelper.ValidatorV2 {
		rt.assertRedeployValPoolToTerminate(defaultValPoolTerminationIndex)
	}

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger1()

	// bind contracts
	rt.bindContracts()

	// submit outputs
	rt.setupOutputSubmitted(version)

	// create challenge
	rt.setupChallenge(rt.challenger1, version)

	var beforeAsset, beforeAssetBonded *big.Int
	if version == valhelper.ValidatorV2 {
		_, _, beforeAsset, beforeAssetBonded, _ = rt.fetchValidatorStatus(rt.challenger1)
	}

interaction:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			rt.txHash = rt.challenger1.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, false)
			rt.includeL1BlockBySender(rt.challenger1.address)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			rt.txHash = rt.validator.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, true)
			rt.includeL1BlockBySender(rt.validator.address)
		case chal.StatusReadyToProve:
			// do nothing to trigger challenger proving timeout
			rt.miner.ActEmptyBlock(rt.t)
		case chal.StatusChallengerTimeout:
			// call challenger timeout by validator
			rt.txHash = rt.validator.ActChallengerTimeout(rt.t, rt.outputIndex, rt.challenger1.address)
			rt.includeL1BlockBySender(rt.validator.address)
		default:
			break interaction
		}

		// check whether the submission was successful
		rt.receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
		require.NoError(rt.t, err)
		require.Equal(rt.t, types.ReceiptStatusSuccessful, rt.receipt.Status, "failed to progress interactive fault proof")
	}

	// check challenge is not proven, output is remained
	remoteOutput, err := rt.outputOracleContract.GetL2Output(nil, rt.outputIndex)
	require.NoError(rt.t, err, "unable to get l2 output")
	outputDeleted := val.IsOutputDeleted(remoteOutput.OutputRoot)
	require.False(rt.t, outputDeleted, "output is deleted when not proven")

	// check output submitter is not changed
	require.Equal(rt.t, remoteOutput.Submitter, rt.validator.address)

	// check the status of challenge is StatusNone(0)
	status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
	require.NoError(rt.t, err)
	require.Equal(rt.t, chal.StatusNone, status)

	if version == valhelper.ValidatorV1 {
		// check bond amount doubled after challenger timed out
		bond, err := rt.valPoolContract.GetBond(nil, rt.outputIndex)
		require.NoError(rt.t, err)
		require.Equal(rt.t, big.NewInt(2*rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt().Int64()), bond.Amount)
	} else if version == valhelper.ValidatorV2 {
		// check challenger has been slashed
		valStatus, inJail, afterAsset, afterAssetBonded, slashingAmount := rt.fetchValidatorStatus(rt.challenger1)
		require.Equal(rt.t, val.StatusReady, valStatus)
		require.True(rt.t, inJail)
		require.Equal(rt.t, beforeAsset.Uint64()-slashingAmount.Uint64(), afterAsset.Uint64())
		require.Equal(rt.t, beforeAssetBonded.Uint64()-slashingAmount.Uint64(), afterAssetBonded.Uint64())
	}
}

func ChallengeInvalidProofFail(t *testing.T, deltaTimeOffset *hexutil.Uint64, version uint8) {
	rt := defaultRuntime(t, setupSequencerTest, deltaTimeOffset)

	if version == valhelper.ValidatorV2 {
		rt.assertRedeployValPoolToTerminate(defaultValPoolTerminationIndex)
	}

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger1()
	rt.setupMaliciousGuardian()

	// bind contracts
	rt.bindContracts()

	// submit outputs
	rt.setupOutputSubmitted(version)

	// create challenge
	rt.setupChallenge(rt.challenger1, version)

	var beforeAssetChal, beforeAssetVal, beforeAssetBondedChal, beforeAssetBondedVal *big.Int
	if version == valhelper.ValidatorV2 {
		_, _, beforeAssetChal, beforeAssetBondedChal, _ = rt.fetchValidatorStatus(rt.challenger1)
		_, _, beforeAssetVal, beforeAssetBondedVal, _ = rt.fetchValidatorStatus(rt.validator)
	}

interaction:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			rt.txHash = rt.challenger1.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, false)
			rt.includeL1BlockBySender(rt.challenger1.address)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			rt.txHash = rt.validator.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, true)
			rt.includeL1BlockBySender(rt.validator.address)
		case chal.StatusReadyToProve:
			rt.txHash = rt.challenger1.ActProveFault(rt.t, rt.outputIndex, false)
			rt.includeL1BlockBySender(rt.challenger1.address)
		case chal.StatusNone:
			// get txId from receipt
			var transactionId *big.Int
			for _, log := range rt.receipt.Logs {
				pLog, _ := rt.securityCouncilContract.SecurityCouncilFilterer.ParseValidationRequested(*log)
				if pLog != nil {
					transactionId = pLog.TransactionId
				}
			}
			require.NotNil(rt.t, transactionId, "unable to get transactionId")

			// check after challenge is proven, output is deleted
			remoteOutput, err := rt.outputOracleContract.GetL2Output(nil, rt.outputIndex)
			require.NoError(rt.t, err, "unable to get l2 output")
			outputDeleted := val.IsOutputDeleted(remoteOutput.OutputRoot)
			require.True(rt.t, outputDeleted, "output is not deleted")

			// guardian validates deleted output by challenger is valid, so confirm the transaction to roll back the challenge
			needConfirm := rt.guardian.ActCheckConfirmCondition(rt.t, rt.outputIndex, transactionId)
			require.True(rt.t, needConfirm, "confirmation condition is not met")
			outputBlockNum := rt.outputOnL1.L2BlockNumber.Uint64()
			isEqual := rt.guardian.ActValidateL2Output(rt.t, rt.outputOnL1.OutputRoot, outputBlockNum)
			require.True(rt.t, isEqual, "deleted output is expected equal but actually not equal")
			rt.txHash = rt.guardian.ActConfirmTransaction(rt.t, transactionId)
			rt.includeL1BlockBySender(rt.guardian.address)
			break interaction
		default:
			break interaction
		}

		// check whether the submission was successful
		rt.receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
		require.NoError(rt.t, err)
		require.Equal(rt.t, types.ReceiptStatusSuccessful, rt.receipt.Status, "failed to progress interactive fault proof")
	}

	// check challenge is dismissed, output is rolled back
	remoteOutput, err := rt.outputOracleContract.GetL2Output(nil, rt.outputIndex)
	require.NoError(rt.t, err, "unable to get l2 output")
	outputDeleted := val.IsOutputDeleted(remoteOutput.OutputRoot)
	require.False(rt.t, outputDeleted, "output is not rolled back")

	// check output submitter is rolled back to asserter
	require.Equal(rt.t, remoteOutput.Submitter, rt.validator.address)

	// check the status of challenge is StatusNone(0)
	status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
	require.NoError(rt.t, err)
	require.Equal(rt.t, chal.StatusNone, status)

	if version == valhelper.ValidatorV1 {
		// check bond amount doubled after challenge is proven incorrectly anyway
		bond, err := rt.valPoolContract.GetBond(nil, rt.outputIndex)
		require.NoError(rt.t, err)
		require.Equal(rt.t, big.NewInt(2*rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt().Int64()), bond.Amount)
	} else if version == valhelper.ValidatorV2 {
		// check challenger has been slashed
		valStatus, inJail, afterAsset, afterAssetBonded, slashingAmount := rt.fetchValidatorStatus(rt.challenger1)
		require.Equal(rt.t, val.StatusReady, valStatus)
		require.True(rt.t, inJail)
		require.Equal(rt.t, beforeAssetChal.Uint64()-slashingAmount.Uint64(), afterAsset.Uint64())
		require.Equal(rt.t, beforeAssetBondedChal.Uint64()-slashingAmount.Uint64(), afterAssetBonded.Uint64())

		// check asserter has been reverted slash by guardian
		valStatus, inJail, afterAsset, afterAssetBonded, slashingAmount = rt.fetchValidatorStatus(rt.validator)
		require.Equal(rt.t, val.StatusActive, valStatus)
		require.False(rt.t, inJail)
		require.Equal(rt.t, beforeAssetVal.Uint64(), afterAsset.Uint64())
		require.Equal(rt.t, beforeAssetBondedVal.Uint64(), afterAssetBonded.Uint64())
	}
}

func ChallengeForceDeleteOutputBySecurityCouncil(t *testing.T, deltaTimeOffset *hexutil.Uint64, version uint8) {
	rt := defaultRuntime(t, setupSequencerTest, deltaTimeOffset)

	if version == valhelper.ValidatorV2 {
		rt.assertRedeployValPoolToTerminate(defaultValPoolTerminationIndex)
	}

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger1()
	rt.setupHonestGuardian()

	// bind contracts
	rt.bindContracts()

	// submit outputs
	rt.setupOutputSubmitted(version)

	// create challenge
	rt.setupChallenge(rt.challenger1, version)

	var beforeAsset, beforeAssetBonded *big.Int
	if version == valhelper.ValidatorV2 {
		_, _, beforeAsset, beforeAssetBonded, _ = rt.fetchValidatorStatus(rt.validator)
	}

interaction:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			rt.txHash = rt.challenger1.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, false)
			rt.includeL1BlockBySender(rt.challenger1.address)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			rt.txHash = rt.validator.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, true)
			rt.includeL1BlockBySender(rt.validator.address)
		case chal.StatusChallengerTimeout:
			rt.txHash = rt.validator.ActChallengerTimeout(rt.t, rt.outputIndex, rt.challenger1.address)
			rt.includeL1BlockBySender(rt.validator.address)
		case chal.StatusReadyToProve:
			// do nothing
			rt.miner.ActEmptyBlock(rt.t)
		case chal.StatusNone:
			outputBlockNum := rt.outputOnL1.L2BlockNumber.Uint64()
			isEqual := rt.guardian.ActValidateL2Output(rt.t, rt.outputOnL1.OutputRoot, outputBlockNum)
			require.False(t, isEqual)

			rt.txHash = rt.guardian.ActForceDeleteOutput(rt.t, rt.outputIndex)
			rt.includeL1BlockBySender(rt.guardian.address)
			break interaction
		default:
			break interaction
		}

		// check whether the submission was successful
		rt.receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
		require.NoError(rt.t, err)
		require.Equal(rt.t, types.ReceiptStatusSuccessful, rt.receipt.Status, "failed to progress interactive fault proof")
	}

	confirmReceipt, err := rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
	require.NoError(rt.t, err)
	require.Equal(rt.t, types.ReceiptStatusSuccessful, confirmReceipt.Status, "failed to confirm")

	// check output is deleted
	remoteOutput, err := rt.outputOracleContract.GetL2Output(nil, rt.outputIndex)
	require.NoError(rt.t, err, "unable to get l2 output")
	outputDeleted := val.IsOutputDeleted(remoteOutput.OutputRoot)
	require.True(rt.t, outputDeleted, "invalid output is not deleted")

	// check output submitter is changed to security council
	securityCouncilAddr, err := rt.colosseumContract.SECURITYCOUNCIL(nil)
	require.NoError(rt.t, err)
	require.Equal(rt.t, remoteOutput.Submitter, securityCouncilAddr)

	if version == valhelper.ValidatorV1 {
		// check bond amount doubled after output is deleted forcefully
		bond, err := rt.valPoolContract.GetBond(nil, rt.outputIndex)
		require.NoError(rt.t, err)
		require.Equal(rt.t, big.NewInt(2*rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt().Int64()), bond.Amount)
	} else if version == valhelper.ValidatorV2 {
		// check asserter has been slashed
		valStatus, inJail, afterAsset, afterAssetBonded, slashingAmount := rt.fetchValidatorStatus(rt.validator)
		require.Equal(rt.t, val.StatusReady, valStatus)
		require.True(rt.t, inJail)
		require.Equal(rt.t, beforeAsset.Uint64()-slashingAmount.Uint64(), afterAsset.Uint64())
		require.Equal(rt.t, beforeAssetBonded.Uint64()-slashingAmount.Uint64(), afterAssetBonded.Uint64())
	}
}

func MultipleChallenges(t *testing.T, deltaTimeOffset *hexutil.Uint64, version uint8) {
	rt := defaultRuntime(t, setupSequencerTest, deltaTimeOffset)

	if version == valhelper.ValidatorV2 {
		rt.assertRedeployValPoolToTerminate(defaultValPoolTerminationIndex)
	}

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger1()
	rt.setupHonestChallenger2()
	rt.setupHonestGuardian()

	// bind contracts
	rt.bindContracts()

	// submit outputs
	rt.setupOutputSubmitted(version)

	// create challenges
	rt.setupChallenge(rt.challenger1, version)
	rt.setupChallenge(rt.challenger2, version)

	// progress challenge by challenger 1
interaction1:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			rt.txHash = rt.challenger1.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, false)
			rt.includeL1BlockBySender(rt.challenger1.address)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			rt.txHash = rt.validator.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, true)
			rt.includeL1BlockBySender(rt.validator.address)
		case chal.StatusReadyToProve:
			rt.txHash = rt.challenger1.ActProveFault(rt.t, rt.outputIndex, false)
			rt.includeL1BlockBySender(rt.challenger1.address)
		case chal.StatusNone:
			// guardian validates deleted output by challenger is invalid after challenge is proven
			outputBlockNum := rt.outputOnL1.L2BlockNumber.Uint64()
			isEqual := rt.guardian.ActValidateL2Output(rt.t, rt.outputOnL1.OutputRoot, outputBlockNum)
			require.False(rt.t, isEqual, "deleted output is expected not equal but actually equal")
			break interaction1
		default:
			break interaction1
		}

		// check whether the submission was successful
		rt.receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
		require.NoError(rt.t, err)
		require.Equal(rt.t, types.ReceiptStatusSuccessful, rt.receipt.Status, "failed to progress interactive fault proof")
	}

	// check output is deleted
	remoteOutput, err := rt.outputOracleContract.GetL2Output(nil, rt.outputIndex)
	require.NoError(rt.t, err, "unable to get l2 output")
	outputDeleted := val.IsOutputDeleted(remoteOutput.OutputRoot)
	require.True(rt.t, outputDeleted, "invalid output is not deleted")

	// check output submitter is changed to challenger
	require.Equal(rt.t, remoteOutput.Submitter, rt.challenger1.address)

	if version == valhelper.ValidatorV1 {
		// check pending bond amount before challenge is canceled
		balance, err := rt.valPoolContract.BalanceOf(nil, rt.challenger2.address)
		require.NoError(rt.t, err)
		require.Equal(rt.t, balance.Int64(), defaultDepositAmount-rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt().Int64())
	} else if version == valhelper.ValidatorV2 {
		// check bond amount before challenge is canceled
		_, _, _, bond, slashingAmount := rt.fetchValidatorStatus(rt.challenger2)
		require.Equal(t, slashingAmount.Uint64(), bond.Uint64())
	}

	// progress challenge by challenger 2
interaction2:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger2.address)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusAsserterTurn:
			// do nothing because output is already deleted
			rt.miner.ActEmptyBlock(rt.t)
		case chal.StatusAsserterTimeout:
			// call cancel challenge by challenger because output is already deleted
			rt.txHash = rt.challenger2.ActCancelChallenge(rt.t, rt.outputIndex)
			rt.includeL1BlockBySender(rt.challenger2.address)
		default:
			break interaction2
		}

		// check whether the submission was successful
		rt.receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
		require.NoError(rt.t, err)
		require.Equal(rt.t, types.ReceiptStatusSuccessful, rt.receipt.Status, "failed to progress interactive fault proof")
	}

	// check the status of challenge is StatusNone(0)
	status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger2.address)
	require.NoError(rt.t, err)
	require.Equal(rt.t, chal.StatusNone, status)

	if version == valhelper.ValidatorV1 {
		// check pending bond amount refunded after challenge canceled
		balance, err := rt.valPoolContract.BalanceOf(nil, rt.challenger2.address)
		require.NoError(rt.t, err)
		require.Equal(rt.t, balance.Int64(), int64(defaultDepositAmount))
	} else if version == valhelper.ValidatorV2 {
		// check bond amount released after challenge canceled
		_, _, _, bond, _ := rt.fetchValidatorStatus(rt.challenger2)
		require.Equal(t, uint64(0), bond.Uint64())
	}
}
