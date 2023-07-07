package actions

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/components/node/eth"
)

func (v *L2Validator) ActValidateL2Output(t Testing, outputRoot eth.Bytes32, l2BlockNumber uint64) bool {
	isValid, err := v.guardian.ValidateL2Output(t.Ctx(), outputRoot, l2BlockNumber)
	require.NoError(t, err, "unable to validate l2Output")
	return isValid
}

func (v *L2Validator) ActConfirmTransaction(t Testing, transactionId *big.Int) common.Hash {
	tx, err := v.guardian.ConfirmTransaction(t.Ctx(), transactionId)
	require.NoError(t, err, "unable to confirm transaction")

	err = v.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	return tx.Hash()
}
