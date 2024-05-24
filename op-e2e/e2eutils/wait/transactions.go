package wait

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ForTransferTxOnL2(l2ChainID *big.Int, l2Seq, l2Sync *ethclient.Client,
	from *ecdsa.PrivateKey, to common.Address, value *big.Int) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	nonce, err := l2Seq.PendingNonceAt(ctx, crypto.PubkeyToAddress(from.PublicKey))
	cancel()
	if err != nil {
		return nil, err
	}

	tx := types.MustSignNewTx(from, types.LatestSignerForChainID(l2ChainID), &types.DynamicFeeTx{
		ChainID:   l2ChainID,
		Nonce:     nonce,
		To:        &to,
		Value:     value,
		GasTipCap: big.NewInt(10),
		GasFeeCap: big.NewInt(200),
		Gas:       21000,
	})

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	err = l2Seq.SendTransaction(ctx, tx)
	cancel()
	if err != nil {
		return nil, err
	}

	_, err = ForReceiptOK(context.Background(), l2Seq, tx.Hash())
	if err != nil {
		return nil, err
	}

	receipt, err := ForReceiptOK(context.Background(), l2Sync, tx.Hash())
	if err != nil {
		return nil, err
	}

	return receipt, nil
}
