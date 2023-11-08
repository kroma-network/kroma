package actions

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-node/eth"
)

func (v *L2Validator) ActValidateL2Output(t Testing, outputRoot eth.Bytes32, l2BlockNumber uint64) bool {
	isEqual, err := v.guardian.ValidateL2Output(t.Ctx(), outputRoot, l2BlockNumber)
	require.NoError(t, err, "unable to validate l2Output")
	return isEqual
}

func (v *L2Validator) ActConfirmTransaction(t Testing, outputIndex *big.Int, transactionId *big.Int) common.Hash {
	outputFinalized, err := v.guardian.IsOutputFinalized(t.Ctx(), outputIndex)
	require.NoError(t, err, "unable to get if output is finalized")
	require.False(t, outputFinalized, "output is already finalized")

	tx, err := v.guardian.ConfirmTransaction(t.Ctx(), transactionId)
	require.NoError(t, err, "unable to confirm transaction")

	err = v.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	return tx.Hash()
}

func (v *L2Validator) ActForceDeleteOutput(t Testing, outputIndex *big.Int) common.Hash {
	tx, err := v.guardian.RequestDeletion(t.Ctx(), outputIndex)
	require.NoError(t, err, "failed to create force delete tx")

	err = v.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	return tx.Hash()
}
