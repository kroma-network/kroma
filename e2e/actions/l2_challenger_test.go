package actions

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	chal "github.com/kroma-network/kroma/components/validator/challenge"
	"github.com/kroma-network/kroma/e2e/testdata"
)

func TestChallenge(t *testing.T) {
	rt := defaultRuntime(t)

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger()
	rt.setupHonestGuardian()

	// bind contracts
	rt.bindChallengeContracts()

	// create challenge
	rt.setupChallenge()

interaction:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			rt.txHash = rt.challenger.ActBisect(rt.t, rt.outputIndex)
			rt.miner.includeL1Block(rt.t, rt.challenger.address)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			rt.txHash = rt.validator.ActBisect(rt.t, rt.outputIndex)
			rt.miner.includeL1Block(rt.t, rt.validator.address)
		case chal.StatusReadyToProve:
			rt.txHash = rt.challenger.ActProveFault(rt.t, rt.outputIndex, false)
			rt.miner.includeL1Block(rt.t, rt.challenger.address)
		case chal.StatusProven:
			// validate l2 output submitted by challenger
			outputBlockNum := rt.outputOnL1.L2BlockNumber.Uint64()
			output := rt.challenger.ActOutputAtBlockSafe(rt.t, outputBlockNum)
			isValid := rt.guardian.ActValidateL2Output(rt.t, output.OutputRoot, outputBlockNum)
			require.True(rt.t, isValid)
			rt.txHash = rt.guardian.ActConfirmTransaction(rt.t, big.NewInt(0))
			rt.miner.includeL1Block(rt.t, rt.guardian.address)
		default:
			break interaction
		}

		// Check whether the submission was successful
		rt.receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
		require.NoError(rt.t, err)
		require.Equal(rt.t, types.ReceiptStatusSuccessful, rt.receipt.Status, "failed to progress interactive fault proof")
	}

	// Check the status of challenge is StatusApproved(7)
	status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex)
	require.NoError(rt.t, err)
	require.Equal(rt.t, chal.StatusApproved, status)
}

func TestChallengeAsserterBisectTimeout(t *testing.T) {
	rt := defaultRuntime(t)

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger()
	rt.setupHonestGuardian()

	// bind contracts
	rt.bindChallengeContracts()

	// create challenge
	rt.setupChallenge()

interaction:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusAsserterTurn:
			// do nothing to trigger asserter timeout
			rt.miner.ActEmptyBlock(rt.t)
		case chal.StatusAsserterTimeout:
			rt.txHash = rt.challenger.ActProveFault(rt.t, rt.outputIndex, true)
			rt.miner.includeL1Block(rt.t, rt.challenger.address)
		case chal.StatusProven:
			// validate l2 output submitted by challenger
			outputBlockNum := rt.outputOnL1.L2BlockNumber.Uint64()
			output := rt.challenger.ActOutputAtBlockSafe(rt.t, outputBlockNum)
			isValid := rt.guardian.ActValidateL2Output(rt.t, output.OutputRoot, outputBlockNum)
			require.True(rt.t, isValid)
			rt.txHash = rt.guardian.ActConfirmTransaction(rt.t, big.NewInt(0))
			rt.miner.includeL1Block(rt.t, rt.guardian.address)
		default:
			break interaction
		}

		// Check whether the submission was successful
		rt.receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
		require.NoError(rt.t, err)
		require.Equal(rt.t, types.ReceiptStatusSuccessful, rt.receipt.Status, "failed to progress interactive fault proof")
	}

	// Check the status of challenge is StatusApproved(7)
	status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex)
	require.NoError(rt.t, err)
	require.Equal(rt.t, chal.StatusApproved, status)
}

func TestChallengeChallengerBisectTimeout(t *testing.T) {
	rt := defaultRuntime(t)

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger()

	// bind contracts
	rt.bindChallengeContracts()

	// create challenge
	rt.setupChallenge()

interaction:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// do nothing to trigger challenger timeout
			rt.miner.ActEmptyBlock(rt.t)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			rt.txHash = rt.validator.ActBisect(rt.t, rt.outputIndex)
			rt.miner.includeL1Block(rt.t, rt.validator.address)
		default:
			break interaction
		}

		// Check whether the submission was successful
		rt.receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
		require.NoError(rt.t, err)
		require.Equal(rt.t, types.ReceiptStatusSuccessful, rt.receipt.Status, "failed to progress interactive fault proof")
	}

	// Check the status of challenge is StatusChallengerTimeout(3)
	status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex)
	require.NoError(rt.t, err)
	require.Equal(rt.t, chal.StatusChallengerTimeout, status)
}

func TestChallengeChallengerProvingTimeout(t *testing.T) {
	rt := defaultRuntime(t)

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger()

	// bind contracts
	rt.bindChallengeContracts()

	// create challenge
	rt.setupChallenge()

interaction:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			rt.txHash = rt.challenger.ActBisect(rt.t, rt.outputIndex)
			rt.miner.includeL1Block(rt.t, rt.challenger.address)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			rt.txHash = rt.validator.ActBisect(rt.t, rt.outputIndex)
			rt.miner.includeL1Block(rt.t, rt.validator.address)
		case chal.StatusReadyToProve:
			// do nothing to trigger challenger proving timeout
			rt.miner.ActEmptyBlock(rt.t)
		default:
			break interaction
		}

		// Check whether the submission was successful
		rt.receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
		require.NoError(rt.t, err)
		require.Equal(rt.t, types.ReceiptStatusSuccessful, rt.receipt.Status, "failed to progress interactive fault proof")
	}

	// Check the status of challenge is StatusChallengerTimeout(3)
	status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex)
	require.NoError(rt.t, err)
	require.Equal(rt.t, chal.StatusChallengerTimeout, status)
}

func TestChallengeInvalidProofFail(t *testing.T) {
	rt := defaultRuntime(t)

	rt.setTargetInvalidBlockNumber(testdata.TargetBlockNumber)
	rt.setupMaliciousValidator()
	rt.setupHonestChallenger()
	rt.setupMaliciousGuardian()

	// bind contracts
	rt.bindChallengeContracts()

	// create challenge
	rt.setupChallenge()

interaction:
	for {
		status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex)
		require.NoError(rt.t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			rt.txHash = rt.challenger.ActBisect(rt.t, rt.outputIndex)
			rt.miner.includeL1Block(rt.t, rt.challenger.address)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			rt.txHash = rt.validator.ActBisect(rt.t, rt.outputIndex)
			rt.miner.includeL1Block(rt.t, rt.validator.address)
		case chal.StatusReadyToProve:
			rt.txHash = rt.challenger.ActProveFault(rt.t, rt.outputIndex, false)
			rt.miner.includeL1Block(rt.t, rt.challenger.address)
		case chal.StatusProven:
			// validate l2 output submitted by challenger
			outputBlockNum := rt.outputOnL1.L2BlockNumber.Uint64()
			output := rt.challenger.ActOutputAtBlockSafe(rt.t, outputBlockNum)
			isValid := rt.guardian.ActValidateL2Output(rt.t, output.OutputRoot, outputBlockNum)
			require.False(rt.t, isValid)
			break interaction
		default:
			break interaction
		}

		// Check whether the submission was successful
		rt.receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
		require.NoError(rt.t, err)
		require.Equal(rt.t, types.ReceiptStatusSuccessful, rt.receipt.Status, "failed to progress interactive fault proof")
	}

	// Check the status of challenge is StatusProven(6)
	status, err := rt.colosseumContract.GetStatus(nil, rt.outputIndex)
	require.NoError(rt.t, err)
	require.Equal(rt.t, chal.StatusProven, status)
}
