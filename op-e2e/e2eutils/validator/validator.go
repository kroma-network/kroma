package validator

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-e2e/config"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
)

type Helper struct {
	t                  *testing.T
	l1Client           *ethclient.Client
	l1ChainID          *big.Int
	l2ChainID          *big.Int
	l1BlockTime        uint64
	valPoolContract    *bindings.ValidatorPool
	ValMgrContract     *bindings.ValidatorManager
	AssetMgrContract   *bindings.AssetManager
	AssetTokenContract *bindings.GovernanceToken
}

func NewHelper(t *testing.T, l1Client *ethclient.Client, l1ChainID *big.Int, l2ChainID *big.Int, l1BlockTime uint64) *Helper {
	valPoolContract, err := bindings.NewValidatorPool(config.L1Deployments.ValidatorPoolProxy, l1Client)
	require.NoError(t, err)

	valMgrContract, err := bindings.NewValidatorManager(config.L1Deployments.ValidatorManagerProxy, l1Client)
	require.NoError(t, err)

	assetMgrContract, err := bindings.NewAssetManager(config.L1Deployments.AssetManagerProxy, l1Client)
	require.NoError(t, err)

	assetTokenContract, err := bindings.NewGovernanceToken(config.L1Deployments.L1GovernanceTokenProxy, l1Client)
	require.NoError(t, err)

	return &Helper{
		t:                  t,
		l1Client:           l1Client,
		l1ChainID:          l1ChainID,
		l2ChainID:          l2ChainID,
		l1BlockTime:        l1BlockTime,
		valPoolContract:    valPoolContract,
		ValMgrContract:     valMgrContract,
		AssetMgrContract:   assetMgrContract,
		AssetTokenContract: assetTokenContract,
	}
}

func (h *Helper) DepositToValPool(priv *ecdsa.PrivateKey, value *big.Int) {
	transactOpts, err := bind.NewKeyedTransactorWithChainID(priv, h.l1ChainID)
	require.NoError(h.t, err)
	transactOpts.Value = value

	tx, err := h.valPoolContract.Deposit(transactOpts)
	require.NoError(h.t, err)

	_, err = wait.ForReceiptOK(context.Background(), h.l1Client, tx.Hash())
	require.NoError(h.t, err)
}

func (h *Helper) UnbondValPool(priv *ecdsa.PrivateKey) bool {
	transactOpts, err := bind.NewKeyedTransactorWithChainID(priv, h.l1ChainID)
	require.NoError(h.t, err)

	tx, err := h.valPoolContract.Unbond(transactOpts)
	require.NoError(h.t, err)

	receipt, err := geth.WaitForTransaction(tx.Hash(), h.l1Client, time.Duration(3*h.l1BlockTime)*time.Second)
	require.NoError(h.t, err)

	return receipt.Status == types.ReceiptStatusSuccessful
}

func (h *Helper) RegisterToValMgr(priv *ecdsa.PrivateKey, amount *big.Int) {
	transactOpts, err := bind.NewKeyedTransactorWithChainID(priv, h.l1ChainID)
	require.NoError(h.t, err)

	tx, err := h.AssetTokenContract.Approve(transactOpts, config.L1Deployments.AssetManagerProxy, amount)
	require.NoError(h.t, err)

	_, err = wait.ForReceiptOK(context.Background(), h.l1Client, tx.Hash())
	require.NoError(h.t, err)

	tx, err = h.ValMgrContract.RegisterValidator(transactOpts, amount, uint8(10), uint8(2))
	require.NoError(h.t, err)

	_, err = wait.ForReceiptOK(context.Background(), h.l1Client, tx.Hash())
	require.NoError(h.t, err)
}

func (h *Helper) Delegate(priv *ecdsa.PrivateKey, validator common.Address, amount *big.Int) {
	transactOpts, err := bind.NewKeyedTransactorWithChainID(priv, h.l1ChainID)
	require.NoError(h.t, err)

	tx, err := h.AssetTokenContract.Approve(transactOpts, config.L1Deployments.AssetManagerProxy, amount)
	require.NoError(h.t, err)

	_, err = wait.ForReceiptOK(context.Background(), h.l1Client, tx.Hash())
	require.NoError(h.t, err)

	tx, err = h.AssetMgrContract.Delegate(transactOpts, validator, amount)
	require.NoError(h.t, err)

	_, err = wait.ForReceiptOK(context.Background(), h.l1Client, tx.Hash())
	require.NoError(h.t, err)
}

func (h *Helper) SendTransferTx(l2Seq *ethclient.Client, l2Sync *ethclient.Client, priv *ecdsa.PrivateKey) *types.Receipt {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	nonce, err := l2Seq.PendingNonceAt(ctx, crypto.PubkeyToAddress(priv.PublicKey))
	cancel()
	require.NoError(h.t, err)

	tx := types.MustSignNewTx(priv, types.LatestSignerForChainID(h.l2ChainID), &types.DynamicFeeTx{
		ChainID:   h.l2ChainID,
		Nonce:     nonce,
		To:        &common.Address{0xff, 0xff},
		Value:     common.Big1,
		GasTipCap: big.NewInt(10),
		GasFeeCap: big.NewInt(200),
		Gas:       21000,
	})

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Duration(h.l1BlockTime)*time.Second)
	err = l2Seq.SendTransaction(ctx, tx)
	cancel()
	require.NoError(h.t, err)

	_, err = geth.WaitForL2Transaction(tx.Hash(), l2Seq, 4*time.Duration(h.l1BlockTime)*time.Second)
	require.NoError(h.t, err)

	receipt, err := geth.WaitForL2Transaction(tx.Hash(), l2Sync, 4*time.Duration(h.l1BlockTime)*time.Second)
	require.NoError(h.t, err)

	return receipt
}
