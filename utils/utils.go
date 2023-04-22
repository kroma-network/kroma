package utils

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/kroma-network/kroma/components/node/client"
	"github.com/kroma-network/kroma/components/node/sources"
	"github.com/kroma-network/kroma/utils/service/crypto"
	"github.com/kroma-network/kroma/utils/service/txmgr"
)

const (
	// DefaultDialTimeout is default duration the service will wait on
	// startup to make a connection to either the L1 or L2 backends.
	DefaultDialTimeout = 5 * time.Second
)

// DialEthClientWithTimeout attempts to dial the L1 provider using the provided
// URL. If the dial doesn't complete within defaultDialTimeout seconds, this
// method will return an error.
func DialEthClientWithTimeout(ctx context.Context, url string) (*ethclient.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, DefaultDialTimeout)
	defer cancel()

	return ethclient.DialContext(ctx, url)
}

// DialRollupClientWithTimeout attempts to dial the RPC provider using the provided
// URL. If the dial doesn't complete within defaultDialTimeout seconds, this
// method will return an error.
func DialRollupClientWithTimeout(ctx context.Context, url string) (*sources.RollupClient, error) {
	ctx, cancel := context.WithTimeout(ctx, DefaultDialTimeout)
	defer cancel()

	rpcCl, err := rpc.DialContext(ctx, url)
	if err != nil {
		return nil, err
	}

	return sources.NewRollupClient(client.NewBaseRPCClient(rpcCl)), nil
}

// ParseAddress parses an ETH address from a hex string. This method will fail if
// the address is not a valid hexadecimal address.
func ParseAddress(address string) (common.Address, error) {
	if common.IsHexAddress(address) {
		return common.HexToAddress(address), nil
	}

	return common.Address{}, fmt.Errorf("invalid address: %v", address)
}

func CalcGasTipAndFeeCap(ctx context.Context, client *ethclient.Client) (*big.Int, *big.Int, error) {
	gasTipCap, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, nil, err
	}

	head, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	gasFeeCap := new(big.Int).Add(gasTipCap, new(big.Int).Mul(head.BaseFee, big.NewInt(2)))

	return gasTipCap, gasFeeCap, nil
}

func UpdateGasPrice(client *ethclient.Client, tx *types.Transaction, addr common.Address, signerFn crypto.SignerFn) txmgr.UpdateGasPriceFunc {
	return func(ctx context.Context) (*types.Transaction, error) {
		var newTx *types.Transaction

		nonce, err := client.PendingNonceAt(ctx, addr)
		if err != nil {
			nonce = tx.Nonce()
		}

		if tx.ChainId() != nil {
			gasTipCap, gasFeeCap, err := CalcGasTipAndFeeCap(ctx, client)
			if err != nil {
				return nil, err
			}
			newTx = types.NewTx(&types.DynamicFeeTx{
				ChainID:    tx.ChainId(),
				Nonce:      nonce,
				GasTipCap:  gasTipCap,
				GasFeeCap:  gasFeeCap,
				Gas:        tx.Gas(),
				To:         tx.To(),
				Value:      tx.Value(),
				Data:       tx.Data(),
				AccessList: tx.AccessList(),
			})
		} else if tx.GasPrice() != nil {
			gasPrice, err := client.SuggestGasPrice(ctx)
			if err != nil {
				return nil, err
			}
			newTx = types.NewTx(&types.LegacyTx{
				Nonce:    nonce,
				GasPrice: gasPrice,
				Gas:      tx.Gas(),
				To:       tx.To(),
				Value:    tx.Value(),
				Data:     tx.Data(),
			})
		}

		signedTx, err := signerFn(ctx, addr, newTx)
		if err != nil {
			return nil, err
		}

		return signedTx, nil
	}
}

func NewSimpleCallOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{Context: ctx}
}

func NewCallOptsWithSender(ctx context.Context, sender common.Address) *bind.CallOpts {
	return &bind.CallOpts{Context: ctx, From: sender}
}

func NewSimpleTxOpts(ctx context.Context, from common.Address, signerFn crypto.SignerFn) *bind.TransactOpts {
	return &bind.TransactOpts{
		From: from,
		Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return signerFn(ctx, addr, tx)
		},
		Context: ctx,
		NoSend:  true,
	}
}
