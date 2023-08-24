package actions

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	val "github.com/kroma-network/kroma/components/validator"
	chal "github.com/kroma-network/kroma/components/validator/challenge"
)

func (v *L2Validator) ActCreateChallenge(t Testing, outputIndex *big.Int) common.Hash {
	inChallengeCreationPeriod, err := v.challenger.IsInChallengeCreationPeriod(t.Ctx(), outputIndex)
	require.NoError(t, err, "unable to check challenge creation period")
	require.True(t, inChallengeCreationPeriod, "challenge creation period is past")

	outputs, err := v.challenger.OutputsAtIndex(t.Ctx(), outputIndex)
	require.NoError(t, err, "unable to fetch outputs")

	outputRange := v.challenger.ValidateOutput(outputIndex, outputs)
	require.NotNil(t, outputRange, "output is valid")

	outputDeleted := val.IsOutputDeleted(outputs.RemoteOutput.OutputRoot)
	require.False(t, outputDeleted, "output is already deleted")

	status, err := v.challenger.GetChallengeStatus(t.Ctx(), outputIndex, v.address)
	require.NoError(t, err, "unable to get challenge status")
	require.Condition(t, func() bool {
		return status == chal.StatusNone || status == chal.StatusChallengerTimeout
	}, "challenge is already in progress")

	hasEnoughDeposit, err := v.challenger.HasEnoughDeposit(t.Ctx())
	require.NoError(t, err, "unable to check challenger deposit")
	require.True(t, hasEnoughDeposit, "challenger not enough deposit to create challenge")

	tx, err := v.challenger.CreateChallenge(t.Ctx(), outputRange)
	require.NoError(t, err, "unable to create create challenge tx")

	err = v.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	return tx.Hash()
}

func (v *L2Validator) ActBisect(t Testing, outputIndex *big.Int, challenger common.Address, isAsserter bool) common.Hash {
	outputFinalized, err := v.challenger.IsOutputFinalized(t.Ctx(), outputIndex)
	require.NoError(t, err, "unable to get if output is finalized")
	require.False(t, outputFinalized, "output is already finalized")

	if isAsserter {
		outputs, err := v.challenger.OutputsAtIndex(t.Ctx(), outputIndex)
		require.NoError(t, err, "unable to fetch outputs")

		outputDeleted := val.IsOutputDeleted(outputs.RemoteOutput.OutputRoot)
		require.False(t, outputDeleted, "output is already deleted")
	}

	tx, err := v.challenger.Bisect(t.Ctx(), outputIndex, challenger)
	require.NoError(t, err, "unable to create bisect tx")

	err = v.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	return tx.Hash()
}

func (v *L2Validator) ActCancelChallenge(t Testing, outputIndex *big.Int) common.Hash {
	tx, err := v.challenger.CancelChallenge(t.Ctx(), outputIndex)
	require.NoError(t, err, "unable to create cancel challenge tx")

	err = v.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	return tx.Hash()
}

func (v *L2Validator) ActChallengerTimeout(t Testing, outputIndex *big.Int, challenger common.Address) common.Hash {
	outputFinalized, err := v.challenger.IsOutputFinalized(t.Ctx(), outputIndex)
	require.NoError(t, err, "unable to get if output is finalized")
	require.False(t, outputFinalized, "output is already finalized")

	outputs, err := v.challenger.OutputsAtIndex(t.Ctx(), outputIndex)
	require.NoError(t, err, "unable to fetch outputs")

	outputDeleted := val.IsOutputDeleted(outputs.RemoteOutput.OutputRoot)
	require.False(t, outputDeleted, "output is already deleted")

	tx, err := v.challenger.ChallengerTimeout(t.Ctx(), outputIndex, challenger)
	require.NoError(t, err, "unable to create challenger timeout tx")

	err = v.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	return tx.Hash()
}

func (v *L2Validator) ActProveFault(t Testing, outputIndex *big.Int, skipSelectPosition bool) common.Hash {
	outputFinalized, err := v.challenger.IsOutputFinalized(t.Ctx(), outputIndex)
	require.NoError(t, err, "unable to get if output is finalized")
	require.False(t, outputFinalized, "output is already finalized")

	tx, err := v.challenger.ProveFault(t.Ctx(), outputIndex, v.address, skipSelectPosition)
	require.NoError(t, err, "unable to create prove fault tx")

	err = v.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	return tx.Hash()
}
