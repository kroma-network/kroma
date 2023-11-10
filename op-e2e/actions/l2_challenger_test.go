package actions

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-e2e/testdata"
	val "github.com/kroma-network/kroma/kroma-validator"
	chal "github.com/kroma-network/kroma/kroma-validator/challenge"
)

func TestChallenge(t *testing.T) {
	rt := defaultRuntime(t, setupSequencerTest)

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger1()
	rt.setupHonestGuardian()

	// bind contracts
	rt.bindChallengeContracts()

	// submit outputs
	rt.setupOutputSubmitted()

	// create challenge
	rt.setupChallenge(rt.challenger1)

interaction:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			rt.txHash = rt.challenger1.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, false)
			rt.IncludeL1Block(rt.challenger1.address)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			rt.txHash = rt.validator.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, true)
			rt.IncludeL1Block(rt.validator.address)
		case chal.StatusReadyToProve:
			rt.txHash = rt.challenger1.ActProveFault(rt.t, rt.outputIndex, false)
			rt.IncludeL1Block(rt.challenger1.address)
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

	// check bond amount doubled after challenge proven
	bond, err := rt.valPoolContract.GetBond(nil, rt.outputIndex)
	require.NoError(rt.t, err)
	require.Equal(rt.t, big.NewInt(2*rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt().Int64()), bond.Amount)
}

func TestChallengeAsserterBisectTimeout(t *testing.T) {
	rt := defaultRuntime(t, setupSequencerTest)

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger1()
	rt.setupHonestGuardian()

	// bind contracts
	rt.bindChallengeContracts()

	// submit outputs
	rt.setupOutputSubmitted()

	// create challenge
	rt.setupChallenge(rt.challenger1)

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
			rt.IncludeL1Block(rt.challenger1.address)
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

	// check bond amount doubled after challenge proven
	bond, err := rt.valPoolContract.GetBond(nil, rt.outputIndex)
	require.NoError(rt.t, err)
	require.Equal(rt.t, big.NewInt(2*rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt().Int64()), bond.Amount)
}

func TestChallengeChallengerBisectTimeout(t *testing.T) {
	rt := defaultRuntime(t, setupSequencerTest)

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger1()

	// bind contracts
	rt.bindChallengeContracts()

	// submit outputs
	rt.setupOutputSubmitted()

	// create challenge
	rt.setupChallenge(rt.challenger1)

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
			rt.IncludeL1Block(rt.validator.address)
		case chal.StatusChallengerTimeout:
			// call challenger timeout by validator
			rt.txHash = rt.validator.ActChallengerTimeout(rt.t, rt.outputIndex, rt.challenger1.address)
			rt.IncludeL1Block(rt.validator.address)
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

	// check bond amount doubled after challenger timed out
	bond, err := rt.valPoolContract.GetBond(nil, rt.outputIndex)
	require.NoError(rt.t, err)
	require.Equal(rt.t, big.NewInt(2*rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt().Int64()), bond.Amount)
}

func TestChallengeChallengerProvingTimeout(t *testing.T) {
	rt := defaultRuntime(t, setupSequencerTest)

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger1()

	// bind contracts
	rt.bindChallengeContracts()

	// submit outputs
	rt.setupOutputSubmitted()

	// create challenge
	rt.setupChallenge(rt.challenger1)

interaction:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			rt.txHash = rt.challenger1.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, false)
			rt.IncludeL1Block(rt.challenger1.address)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			rt.txHash = rt.validator.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, true)
			rt.IncludeL1Block(rt.validator.address)
		case chal.StatusReadyToProve:
			// do nothing to trigger challenger proving timeout
			rt.miner.ActEmptyBlock(rt.t)
		case chal.StatusChallengerTimeout:
			// call challenger timeout by validator
			rt.txHash = rt.validator.ActChallengerTimeout(rt.t, rt.outputIndex, rt.challenger1.address)
			rt.IncludeL1Block(rt.validator.address)
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

	// check bond amount doubled after challenger timed out
	bond, err := rt.valPoolContract.GetBond(nil, rt.outputIndex)
	require.NoError(rt.t, err)
	require.Equal(rt.t, big.NewInt(2*rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt().Int64()), bond.Amount)
}

func TestChallengeInvalidProofFail(t *testing.T) {
	rt := defaultRuntime(t, setupSequencerTest)

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger1()
	rt.setupMaliciousGuardian()

	// bind contracts
	rt.bindChallengeContracts()

	// submit outputs
	rt.setupOutputSubmitted()

	// create challenge
	rt.setupChallenge(rt.challenger1)

interaction:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			rt.txHash = rt.challenger1.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, false)
			rt.IncludeL1Block(rt.challenger1.address)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			rt.txHash = rt.validator.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, true)
			rt.IncludeL1Block(rt.validator.address)
		case chal.StatusReadyToProve:
			rt.txHash = rt.challenger1.ActProveFault(rt.t, rt.outputIndex, false)
			rt.IncludeL1Block(rt.challenger1.address)
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

			// guardian validates deleted output by challenger is invalid after challenge is proven
			outputBlockNum := rt.outputOnL1.L2BlockNumber.Uint64()
			isEqual := rt.guardian.ActValidateL2Output(rt.t, rt.outputOnL1.OutputRoot, outputBlockNum)
			require.True(rt.t, isEqual, "deleted output is expected equal but actually not equal")
			rt.txHash = rt.guardian.ActConfirmTransaction(rt.t, rt.outputIndex, transactionId)
			rt.IncludeL1Block(rt.guardian.address)
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

	// check bond amount doubled after challenge is proven incorrectly anyway
	bond, err := rt.valPoolContract.GetBond(nil, rt.outputIndex)
	require.NoError(rt.t, err)
	require.Equal(rt.t, big.NewInt(2*rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt().Int64()), bond.Amount)
}

func TestMultipleChallenges(t *testing.T) {
	rt := defaultRuntime(t, setupSequencerTest)

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger1()
	rt.setupHonestChallenger2()
	rt.setupHonestGuardian()

	// bind contracts
	rt.bindChallengeContracts()

	// submit outputs
	rt.setupOutputSubmitted()

	// create challenges
	rt.setupChallenge(rt.challenger1)
	rt.setupChallenge(rt.challenger2)

	// progress challenge by challenger 1
interaction1:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			rt.txHash = rt.challenger1.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, false)
			rt.IncludeL1Block(rt.challenger1.address)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			rt.txHash = rt.validator.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, true)
			rt.IncludeL1Block(rt.validator.address)
		case chal.StatusReadyToProve:
			rt.txHash = rt.challenger1.ActProveFault(rt.t, rt.outputIndex, false)
			rt.IncludeL1Block(rt.challenger1.address)
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

	// check pending bond amount before challenge is canceled
	balance, err := rt.valPoolContract.BalanceOf(nil, rt.challenger2.address)
	require.NoError(rt.t, err)
	require.Equal(rt.t, balance.Int64(), defaultDepositAmount-rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt().Int64())

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
			// call bisect by challenger
			rt.txHash = rt.challenger2.ActProveFault(rt.t, rt.outputIndex, true)
			rt.IncludeL1Block(rt.challenger2.address)
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

	// check pending bond amount refunded after challenge canceled
	balance, err = rt.valPoolContract.BalanceOf(nil, rt.challenger2.address)
	require.NoError(rt.t, err)
	require.Equal(rt.t, balance.Int64(), int64(defaultDepositAmount))
}

func TestChallengeForceDeleteOutputBySecurityCouncil(t *testing.T) {
	rt := defaultRuntime(t, setupSequencerTest)
	rt.SetCreationPeriod(9)

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger1()
	rt.setupHonestGuardian()

	// bind contracts
	rt.bindChallengeContracts()

	// submit outputs
	rt.setupOutputSubmitted()

	// create challenge
	rt.setupChallenge(rt.challenger1)

interaction:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex, rt.challenger1.address)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			rt.txHash = rt.challenger1.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, false)
			rt.IncludeL1Block(rt.challenger1.address)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			rt.txHash = rt.validator.ActBisect(rt.t, rt.outputIndex, rt.challenger1.address, true)
			rt.IncludeL1Block(rt.validator.address)
		case chal.StatusAsserterTimeout:
			rt.txHash = rt.challenger1.ActProveFault(rt.t, rt.outputIndex, true)
			rt.IncludeL1Block(rt.challenger1.address)
		case chal.StatusChallengerTimeout:
			rt.txHash = rt.validator.ActChallengerTimeout(rt.t, rt.outputIndex, rt.challenger1.address)
			rt.IncludeL1Block(rt.validator.address)
		case chal.StatusReadyToProve:
			// do nothing
			rt.miner.ActEmptyBlock(rt.t)
		case chal.StatusNone:
			if rt.IsCreationEnded() {
				outputBlockNum := rt.outputOnL1.L2BlockNumber.Uint64()
				isEqual := rt.guardian.ActValidateL2Output(rt.t, rt.outputOnL1.OutputRoot, outputBlockNum)
				require.False(t, isEqual)

				votes, _ := rt.securityCouncilContract.GetVotes(nil, rt.guardian.address)
				fmt.Printf("asdfasdfasdf %d\n", votes)

				rt.txHash = rt.guardian.ActForceDeleteOutput(rt.t, rt.outputIndex)
				rt.IncludeL1Block(rt.challenger1.address)
				break interaction
			}
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

	// check output submitter is changed to challenger
	securityCouncilAddr, err := rt.colosseumContract.SECURITYCOUNCIL(nil)
	require.NoError(rt.t, err)
	require.Equal(rt.t, remoteOutput.Submitter, securityCouncilAddr)

	// check bond amount doubled after challenge proven
	bond, err := rt.valPoolContract.GetBond(nil, rt.outputIndex)
	require.NoError(rt.t, err)
	require.Equal(rt.t, big.NewInt(2*rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt().Int64()), bond.Amount)
}
