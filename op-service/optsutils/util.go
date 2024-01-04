package optsutils

import (
	"context"
	"github.com/ethereum-optimism/optimism/op-service/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

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
