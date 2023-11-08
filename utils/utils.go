package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum-optimism/optimism/op-node/client"
	"github.com/ethereum-optimism/optimism/op-node/sources"
	"github.com/ethereum-optimism/optimism/op-service/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
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

func NewSimpleWatchOpts(ctx context.Context) *bind.WatchOpts {
	return &bind.WatchOpts{Context: ctx}
}
