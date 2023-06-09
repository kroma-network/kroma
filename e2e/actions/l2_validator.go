package actions

import (
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/components/node/sources"
	validator "github.com/kroma-network/kroma/components/validator"
	"github.com/kroma-network/kroma/utils/service/txmgr"
)

type ValidatorCfg struct {
	OutputOracleAddr  common.Address
	ColosseumAddr     common.Address
	ValidatorKey      *ecdsa.PrivateKey
	AllowNonFinalized bool
}

type L2Validator struct {
	log          log.Logger
	l1           *ethclient.Client
	l2os         *validator.L2OutputSubmitter
	address      common.Address
	privKey      *ecdsa.PrivateKey
	contractAddr common.Address
	lastTx       common.Hash
}

func NewL2Validator(t Testing, log log.Logger, cfg *ValidatorCfg, l1 *ethclient.Client, rollupCl *sources.RollupClient) *L2Validator {
	from := crypto.PubkeyToAddress(cfg.ValidatorKey.PublicKey)

	rollupConfig, err := rollupCl.RollupConfig(t.Ctx())
	require.NoError(t, err)

	validatorCfg := validator.Config{
		L2OutputOracleAddr:     cfg.OutputOracleAddr,
		ChallengerPollInterval: time.Second,
		NetworkTimeout:         time.Second,
		L1Client:               l1,
		RollupClient:           rollupCl,
		RollupConfig:           rollupConfig,
		AllowNonFinalized:      cfg.AllowNonFinalized,
		// We use custom signing here instead of using the transaction manager.
		TxManager: &txmgr.SimpleTxManager{
			Config: txmgr.Config{From: from},
		},
	}

	l2os, err := validator.NewL2OutputSubmitter(validatorCfg, log)
	require.NoError(t, err)

	return &L2Validator{
		log:          log,
		l1:           l1,
		l2os:         l2os,
		address:      from,
		privKey:      cfg.ValidatorKey,
		contractAddr: cfg.OutputOracleAddr,
	}
}

// sendTx reimplements creating & sending transactions because we need to do the final send as async in
// the action tests while we do it synchronously in the real system.
func (v *L2Validator) sendTx(t Testing, data []byte) {
	gasTipCap := big.NewInt(2 * params.GWei)
	pendingHeader, err := v.l1.HeaderByNumber(t.Ctx(), big.NewInt(-1))
	require.NoError(t, err, "need l1 pending header for gas price estimation")
	gasFeeCap := new(big.Int).Add(gasTipCap, new(big.Int).Mul(pendingHeader.BaseFee, big.NewInt(2)))
	chainID, err := v.l1.ChainID(t.Ctx())
	require.NoError(t, err)
	nonce, err := v.l1.NonceAt(t.Ctx(), v.address, nil)
	require.NoError(t, err)

	gasLimit, err := v.l1.EstimateGas(t.Ctx(), ethereum.CallMsg{
		From:      v.address,
		To:        &v.contractAddr,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Data:      data,
	})
	require.NoError(t, err)

	rawTx := &types.DynamicFeeTx{
		Nonce:     nonce,
		To:        &v.contractAddr,
		Data:      data,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Gas:       gasLimit,
		ChainID:   chainID,
	}

	tx, err := types.SignNewTx(v.privKey, types.LatestSignerForChainID(chainID), rawTx)
	require.NoError(t, err, "need to sign tx")

	err = v.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err, "need to send tx")

	v.lastTx = tx.Hash()
}

func (v *L2Validator) CanSubmit(t Testing) bool {
	_, shouldSubmit, err := v.l2os.FetchNextOutputInfo(t.Ctx())
	require.NoError(t, err)
	return shouldSubmit
}

func (v *L2Validator) ActSubmitL2Output(t Testing) {
	output, shouldSubmit, err := v.l2os.FetchNextOutputInfo(t.Ctx())
	if !shouldSubmit {
		return
	}
	require.NoError(t, err)

	txData, err := v.l2os.SubmitL2OutputTxData(output)
	require.NoError(t, err)

	// Note: Use L1 instead of the output submitter's transaction manager because
	// this is non-blocking while the txmgr is blocking & deadlocks the tests
	v.sendTx(t, txData)
}

func (v *L2Validator) LastSubmitL2OutputTx() common.Hash {
	return v.lastTx
}
