package e2e

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/components/node/eth"
)

// TestMissingGasLimit tests that kroma-geth cannot build a block without gas limit while kroma is active in the chain config.
func TestMissingGasLimit(t *testing.T) {
	cfg := DefaultSystemConfig(t)
	cfg.DeployConfig.FundDevAccounts = false
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	kromaGeth, err := NewKromaGeth(t, ctx, &cfg)
	require.NoError(t, err)
	defer kromaGeth.Close()

	attrs, err := kromaGeth.CreatePayloadAttributes()
	require.NoError(t, err)
	// Remove the GasLimit from the otherwise valid attributes
	attrs.GasLimit = nil

	res, err := kromaGeth.StartBlockBuilding(ctx, attrs)
	require.ErrorIs(t, err, eth.InputError{})
	require.Equal(t, eth.InvalidPayloadAttributes, err.(eth.InputError).Code)
	require.Nil(t, res)
}

// TestInvalidDepositInFCU runs an invalid deposit through a FCU/GetPayload/NewPayload/FCU set of calls.
// This tests that deposits must always allow the block to be built even if they are invalid.
func TestInvalidDepositInFCU(t *testing.T) {
	cfg := DefaultSystemConfig(t)
	cfg.DeployConfig.FundDevAccounts = false
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	kromaGeth, err := NewKromaGeth(t, ctx, &cfg)
	require.NoError(t, err)
	defer kromaGeth.Close()

	// Create a deposit from alice that will always fail (not enough funds)
	fromAddr := cfg.Secrets.Addresses().Alice
	balance, err := kromaGeth.L2Client.BalanceAt(ctx, fromAddr, nil)
	require.Nil(t, err)
	require.Equal(t, 0, balance.Cmp(common.Big0))

	badDepositTx := types.NewTx(&types.DepositTx{
		From:  fromAddr,
		To:    &fromAddr, // send it to ourselves
		Value: big.NewInt(params.Ether),
		Gas:   25000,
	})

	// We are inserting a block with an invalid deposit.
	// The invalid deposit should still remain in the block.
	_, err = kromaGeth.AddL2Block(ctx, badDepositTx)
	require.NoError(t, err)

	// Deposit tx was included, but Alice still shouldn't have any ETH
	balance, err = kromaGeth.L2Client.BalanceAt(ctx, fromAddr, nil)
	require.Nil(t, err)
	require.Equal(t, 0, balance.Cmp(common.Big0))
}
