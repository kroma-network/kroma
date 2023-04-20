package actions

import (
	"context"
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
	chal "github.com/wemixkanvas/kanvas/components/validator/challenge"
	"github.com/wemixkanvas/kanvas/e2e/e2eutils"
	kcrypto "github.com/wemixkanvas/kanvas/utils/service/crypto"
	"github.com/wemixkanvas/kanvas/utils/service/txmgr"
)

type L2Challenger struct {
	log        log.Logger
	l1         *ethclient.Client
	challenger *validator.Challenger
	address    common.Address
}

func NewL2Challenger(t Testing, log log.Logger, cfg *ValidatorCfg, l1 *ethclient.Client, rollupCl *sources.RollupClient) *L2Challenger {
	signer := func(chainID *big.Int) kcrypto.SignerFn {
		s := kcrypto.PrivateKeySignerFn(cfg.ValidatorKey, chainID)
		return func(_ context.Context, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return s(addr, tx)
		}
	}
	from := crypto.PubkeyToAddress(cfg.ValidatorKey.PublicKey)

	chainID, err := l1.ChainID(t.Ctx())
	require.NoError(t, err)

	rollupConfig, err := rollupCl.RollupConfig(t.Ctx())
	require.NoError(t, err)

	validatorCfg := validator.Config{
		L2OutputOracleAddr: cfg.OutputOracleAddr,
		ColosseumAddr:      cfg.ColosseumAddr,
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
		RollupConfig:      rollupConfig,
		AllowNonFinalized: cfg.AllowNonFinalized,
		From:              from,
		SignerFn:          signer(chainID),
		ProofFetcher:      e2eutils.NewFetcher(log),
	}

	challenger, err := validator.NewChallenger(t.Ctx(), validatorCfg, log)
	require.NoError(t, err)

	return &L2Challenger{
		log:        log,
		l1:         l1,
		challenger: challenger,
		address:    crypto.PubkeyToAddress(cfg.ValidatorKey.PublicKey),
	}
}

func (c *L2Challenger) ActCreateChallenge(t Testing) common.Hash {
	isInProgress, err := c.challenger.IsChallengeInProgress()
	require.NoError(t, err)
	require.False(t, isInProgress, "another challenge is in progress")

	outputRange, err := c.challenger.GetInvalidOutputRange()
	require.NoError(t, err)
	require.NotNil(t, outputRange)
	tx, err := c.challenger.CreateChallenge(outputRange)
	require.NoError(t, err, "unable to create createChallenge tx")

	err = c.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	return tx.Hash()
}

func (c *L2Challenger) ActBisect(t Testing) common.Hash {
	status, err := c.challenger.GetStatusInProgress()
	require.NoError(t, err)

	var tx *types.Transaction

	if status == chal.StatusChallengerTurn || status == chal.StatusAsserterTurn {
		tx, err = c.challenger.Bisect()
		require.NoError(t, err, "unable to create bisect tx")
	} else {
		require.Fail(t, "invalid challenge status")
	}

	err = c.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	return tx.Hash()
}

func (c *L2Challenger) ActTimeout(t Testing) common.Hash {
	status, err := c.challenger.GetStatusInProgress()
	require.NoError(t, err)

	var tx *types.Transaction

	if status == chal.StatusAsserterTimeout {
		tx, err = c.challenger.AsserterTimeout()
	} else if status == chal.StatusChallengerTimeout {
		challengeId, err := c.challenger.LatestChallengeId()
		require.NoError(t, err)
		tx, err = c.challenger.ChallengerTimeout(challengeId)
		require.NoError(t, err)
	} else {
		require.Fail(t, "invalid challenge status")
	}

	require.NoError(t, err, "unable to create tx")

	err = c.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	return tx.Hash()
}

func (c *L2Challenger) ActProveFault(t Testing) common.Hash {
	status, err := c.challenger.GetStatusInProgress()
	require.NoError(t, err)
	require.Equal(t, status, chal.StatusProveReady)

	tx, err := c.challenger.ProveFault()
	require.NoError(t, err, "unable to create proveFault tx")

	err = c.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err)

	return tx.Hash()
}
