package actions

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	kcrypto "github.com/ethereum-optimism/optimism/op-service/crypto"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
	"github.com/kroma-network/kroma/kroma-validator"
	validatormetrics "github.com/kroma-network/kroma/kroma-validator/metrics"
)

type ValidatorCfg struct {
	OutputOracleAddr     common.Address
	ColosseumAddr        common.Address
	SecurityCouncilAddr  common.Address
	ValidatorPoolAddr    common.Address
	ValidatorManagerAddr common.Address
	AssetManagerAddr     common.Address
	ValidatorKey         *ecdsa.PrivateKey
	AllowNonFinalized    bool
}

type L2Validator struct {
	log                      log.Logger
	l1                       *ethclient.Client
	l2os                     *validator.L2OutputSubmitter
	challenger               *validator.Challenger
	guardian                 *validator.Guardian
	address                  common.Address
	privKey                  *ecdsa.PrivateKey
	l2ooContractAddr         common.Address
	valPoolContractAddr      common.Address
	valManagerContractAddr   common.Address
	assetManagerContractAddr common.Address
	lastTx                   common.Hash
	cfg                      *validator.Config
}

func NewL2Validator(t Testing, log log.Logger, cfg *ValidatorCfg, l1 *ethclient.Client, l2 *ethclient.Client, rollupCl *sources.RollupClient) *L2Validator {
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
		L2OutputOracleAddr:              cfg.OutputOracleAddr,
		ValidatorPoolAddr:               cfg.ValidatorPoolAddr,
		ValidatorManagerAddr:            cfg.ValidatorManagerAddr,
		AssetManagerAddr:                cfg.AssetManagerAddr,
		ColosseumAddr:                   cfg.ColosseumAddr,
		SecurityCouncilAddr:             cfg.SecurityCouncilAddr,
		ChallengerPollInterval:          time.Second,
		OutputSubmitterRetryInterval:    time.Second,
		OutputSubmitterRoundBuffer:      30,
		OutputSubmitterAllowPublicRound: true,
		NetworkTimeout:                  time.Second,
		L1Client:                        l1,
		L2Client:                        l2,
		RollupClient:                    rollupCl,
		RollupConfig:                    rollupConfig,
		AllowNonFinalized:               cfg.AllowNonFinalized,
		ProofFetcher:                    e2eutils.NewFetcher(log, "../testdata/proof"),
		// We use custom signing here instead of using the transaction manager.
		TxManager: &txmgr.BufferedTxManager{
			SimpleTxManager: txmgr.SimpleTxManager{
				Config: txmgr.Config{
					From:   from,
					Signer: signer(chainID),
				},
			},
		},
	}

	l2os, err := validator.NewL2OutputSubmitter(validatorCfg, log, validatormetrics.NoopMetrics)
	require.NoError(t, err)
	err = l2os.InitConfig(t.Ctx())
	require.NoError(t, err)

	challenger, err := validator.NewChallenger(validatorCfg, log, validatormetrics.NoopMetrics)
	require.NoError(t, err)
	err = challenger.InitConfig(t.Ctx())
	require.NoError(t, err)

	guardian, err := validator.NewGuardian(validatorCfg, log)
	require.NoError(t, err)
	err = guardian.InitConfig(t.Ctx())
	require.NoError(t, err)

	return &L2Validator{
		log:                      log,
		l1:                       l1,
		l2os:                     l2os,
		challenger:               challenger,
		guardian:                 guardian,
		address:                  from,
		privKey:                  cfg.ValidatorKey,
		l2ooContractAddr:         cfg.OutputOracleAddr,
		valPoolContractAddr:      cfg.ValidatorPoolAddr,
		valManagerContractAddr:   cfg.ValidatorManagerAddr,
		assetManagerContractAddr: cfg.AssetManagerAddr,
		cfg:                      &validatorCfg,
	}
}

// sendTx reimplements creating & sending transactions because we need to do the final send as async in
// the action tests while we do it synchronously in the real system.
func (v *L2Validator) sendTx(t Testing, toAddr *common.Address, txValue *big.Int, data []byte, gasLimitMultiplier float64) {
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
		To:        toAddr,
		Value:     txValue,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Data:      data,
	})
	require.NoError(t, err)

	rawTx := &types.DynamicFeeTx{
		Nonce:     nonce,
		To:        toAddr,
		Value:     txValue,
		Data:      data,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Gas:       uint64(float64(gasLimit) * gasLimitMultiplier),
		ChainID:   chainID,
	}

	tx, err := types.SignNewTx(v.privKey, types.LatestSignerForChainID(chainID), rawTx)
	require.NoError(t, err, "need to sign tx")

	err = v.l1.SendTransaction(t.Ctx(), tx)
	require.NoError(t, err, "need to send tx")

	v.lastTx = tx.Hash()
}

func (v *L2Validator) CalculateWaitTime(t Testing) time.Duration {
	nextBlockNumber, err := v.l2os.FetchNextBlockNumber(t.Ctx())
	require.NoError(t, err)
	calculatedWaitTime := v.l2os.CalculateWaitTime(t.Ctx(), nextBlockNumber)
	return calculatedWaitTime
}

func (v *L2Validator) ActSubmitL2Output(t Testing) {
	nextBlockNumber, err := v.l2os.FetchNextBlockNumber(t.Ctx())
	require.NoError(t, err)

	output, err := v.l2os.FetchOutput(t.Ctx(), nextBlockNumber)
	require.NoError(t, err)

	txData, err := validator.SubmitL2OutputTxData(v.l2os.L2ooAbi(), output)
	require.NoError(t, err)

	// Note: Use L1 instead of the output submitter's transaction manager because
	// this is non-blocking while the txmgr is blocking & deadlocks the tests.
	// Also set gasLimitMultiplier above 1 because finalization process sets state variables from 0 to value.
	v.sendTx(t, &v.l2ooContractAddr, common.Big0, txData, 2)
}

func (v *L2Validator) LastSubmitL2OutputTx() common.Hash {
	return v.lastTx
}

func (v *L2Validator) ActDeposit(t Testing, depositAmount uint64) {
	valPoolABI, err := bindings.ValidatorPoolMetaData.GetAbi()
	require.NoError(t, err)

	txData, err := valPoolABI.Pack("deposit")
	require.NoError(t, err)

	v.sendTx(t, &v.valPoolContractAddr, new(big.Int).SetUint64(depositAmount), txData, 1)
}

func (v *L2Validator) ActRegisterValidator(t Testing, assets *big.Int) {
	valManagerABI, err := bindings.ValidatorManagerMetaData.GetAbi()
	require.NoError(t, err)

	txData, err := valManagerABI.Pack(
		"registerValidator",
		assets,
		uint8(10),
		uint8(2),
	)
	require.NoError(t, err)

	v.sendTx(t, &v.valManagerContractAddr, common.Big0, txData, 2)
}

func (v *L2Validator) ActApprove(t Testing, amount uint64) {
	tokenAddr := v.getGovTokenAddr(t)
	governanceTokenABI, err := bindings.GovernanceTokenMetaData.GetAbi()
	require.NoError(t, err)

	txData, err := governanceTokenABI.Pack("approve", &v.assetManagerContractAddr, new(big.Int).SetUint64(amount))
	require.NoError(t, err)

	v.sendTx(t, &tokenAddr, common.Big0, txData, 1)
}

func (v *L2Validator) fetchOutput(t Testing, blockNumber *big.Int) *eth.OutputResponse {
	output, err := v.l2os.FetchOutput(t.Ctx(), blockNumber)
	require.NoError(t, err)

	return output
}

func (v *L2Validator) getGovTokenAddr(t Testing) common.Address {
	assetManagerABI, err := bindings.AssetManagerMetaData.GetAbi()
	require.NoError(t, err)

	callData, err := assetManagerABI.Pack("ASSET_TOKEN")
	require.NoError(t, err)

	returnData, err := v.l1.CallContract(t.Ctx(), ethereum.CallMsg{
		To:   &v.assetManagerContractAddr,
		Data: callData,
	}, nil)
	require.NoError(t, err)

	tokenAddr := common.BytesToAddress(returnData)
	return tokenAddr
}
