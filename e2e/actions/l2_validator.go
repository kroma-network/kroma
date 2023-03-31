package actions

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/wemixkanvas/kanvas/components/node/sources"
	validator "github.com/wemixkanvas/kanvas/components/validator"
	kcrypto "github.com/wemixkanvas/kanvas/utils/service/crypto"
	"github.com/wemixkanvas/kanvas/utils/service/txmgr"
)

type ValidatorCfg struct {
	OutputOracleAddr  common.Address
	ColosseumAddr     common.Address
	ValidatorKey      *ecdsa.PrivateKey
	AllowNonFinalized bool
}

type L2Validator struct {
	log     log.Logger
	l1      *ethclient.Client
	l2os    *validator.L2OutputSubmitter
	address common.Address
	lastTx  common.Hash
}

func NewL2Validator(t Testing, log log.Logger, cfg *ValidatorCfg, l1 *ethclient.Client, rollupCl *sources.RollupClient) *L2Validator {
	signer := func(chainID *big.Int) kcrypto.SignerFn {
		s := kcrypto.PrivateKeySignerFn(cfg.ValidatorKey, chainID)
		return func(_ context.Context, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return s(addr, tx)
		}
	}
	from := crypto.PubkeyToAddress(cfg.ValidatorKey.PublicKey)

	chainID, err := l1.ChainID(t.Ctx())
	require.NoError(t, err)

	validatorCfg := validator.Config{
		L2OutputOracleAddr: cfg.OutputOracleAddr,
		PollInterval:       time.Second,
		TxManagerConfig: txmgr.Config{
			ResubmissionTimeout:       5 * time.Second,
			ReceiptQueryInterval:      time.Second,
			NumConfirmations:          1,
			SafeAbortNonceTooLowCount: 4,
			From:                      from,
			Signer:                    signer(chainID),
		},
		L1Client:          l1,
		RollupClient:      rollupCl,
		AllowNonFinalized: cfg.AllowNonFinalized,
		From:              from,
		SignerFn:          signer(chainID),
	}

	l2os, err := validator.NewL2OutputSubmitter(t.Ctx(), validatorCfg, log)
	require.NoError(t, err)

	return &L2Validator{
		log:     log,
		l1:      l1,
		l2os:    l2os,
		address: crypto.PubkeyToAddress(cfg.ValidatorKey.PublicKey),
	}
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

	tx, err := v.l2os.CreateSubmitL2OutputTx(t.Ctx(), output)
	require.NoError(t, err)

	err = v.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	v.lastTx = tx.Hash()
}

func (v *L2Validator) LastSubmitL2OutputTx() common.Hash {
	return v.lastTx
}
