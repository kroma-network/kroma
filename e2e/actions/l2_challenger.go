package actions

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/components/node/eth"
	chal "github.com/kroma-network/kroma/components/validator/challenge"
)

func (v *L2Validator) ActCreateChallenge(t Testing, outputIndex *big.Int) common.Hash {
	isInProgress, err := v.challenger.IsChallengeInProgress(t.Ctx(), outputIndex)
	require.NoError(t, err)
	require.False(t, isInProgress, "another challenge is in progress")

	outputRange, err := v.challenger.ValidateOutput(t.Ctx(), outputIndex)
	require.NoError(t, err)
	require.NotNil(t, outputRange)
	tx, err := v.challenger.CreateChallenge(t.Ctx(), outputRange)
	require.NoError(t, err, "unable to create createChallenge tx")

	err = v.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	return tx.Hash()
}

func (v *L2Validator) ActBisect(t Testing, outputIndex *big.Int) common.Hash {
	tx, err := v.challenger.Bisect(t.Ctx(), outputIndex)
	require.NoError(t, err, "unable to create bisect tx")

	err = v.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	return tx.Hash()
}

func (v *L2Validator) ActTimeout(t Testing, outputIndex *big.Int) common.Hash {
	status, err := v.challenger.GetChallengeStatus(t.Ctx(), outputIndex)
	require.NoError(t, err)

	var tx *types.Transaction

	if status == chal.StatusChallengerTimeout {
		tx, err = v.challenger.ChallengerTimeout(t.Ctx(), outputIndex)
		require.NoError(t, err)
	} else {
		require.Fail(t, "invalid challenge status")
	}

	require.NoError(t, err, "unable to create tx")

	err = v.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	return tx.Hash()
}

func (v *L2Validator) ActProveFault(t Testing, outputIndex *big.Int) common.Hash {
	tx, err := v.challenger.ProveFault(t.Ctx(), outputIndex)
	require.NoError(t, err, "unable to create proveFault tx")

	err = v.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	return tx.Hash()
}

func (v *L2Validator) ActOutputAtBlockSafe(t Testing, blockNumber uint64) *eth.OutputResponse {
	output, err := v.challenger.OutputAtBlockSafe(t.Ctx(), blockNumber)
	require.NoError(t, err, "unable get output at block safe")

	return output
}
