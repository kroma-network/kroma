package balance

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	opservice "github.com/ethereum-optimism/optimism/op-service"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum-optimism/optimism/op-service/txmgr/metrics"
	"github.com/kroma-network/kroma/kroma-bindings/bindings"
	"github.com/kroma-network/kroma/kroma-validator/flags"
)

func Deposit(ctx *cli.Context) error {
	amount := ctx.String("amount")

	valpoolABI, err := bindings.ValidatorPoolMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ValidatorPool ABI: %w", err)
	}

	txData, err := valpoolABI.Pack("deposit")
	if err != nil {
		return fmt.Errorf("failed to create deposit transaction data: %w", err)
	}

	if err = sendTransaction(ctx, txData, amount); err != nil {
		return err
	}

	return nil
}

func Withdraw(ctx *cli.Context) error {
	amount := ctx.String("amount")

	withdrawAmount, success := new(big.Int).SetString(amount, 10)
	if !success {
		return errors.New("failed to parse withdraw amount")
	}

	valpoolABI, err := bindings.ValidatorPoolMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ValidatorPool ABI: %w", err)
	}

	txData, err := valpoolABI.Pack("withdraw", withdrawAmount)
	if err != nil {
		return fmt.Errorf("failed to create withdraw transaction data: %w", err)
	}

	if err = sendTransaction(ctx, txData, "0"); err != nil {
		return err
	}

	return nil
}

func Unbond(ctx *cli.Context) error {
	valpoolABI, err := bindings.ValidatorPoolMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ValidatorPool ABI: %w", err)
	}

	txData, err := valpoolABI.Pack("unbond")
	if err != nil {
		return fmt.Errorf("failed to create unbond transaction data: %w", err)
	}

	if err = sendTransaction(ctx, txData, "0"); err != nil {
		return err
	}

	return nil
}

func sendTransaction(ctx *cli.Context, txData []byte, txValue string) error {
	txMgrConfig := txmgr.ReadCLIConfig(ctx)
	txManager, err := txmgr.NewSimpleTxManager("validator-balance", log.New(), &metrics.NoopTxMetrics{}, txMgrConfig)
	if err != nil {
		return fmt.Errorf("failed to create tx manager: %w", err)
	}

	valpoolAddr, err := opservice.ParseAddress(ctx.String(flags.ValPoolAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorPool address: %w", err)
	}

	value, success := new(big.Int).SetString(txValue, 10)
	if !success {
		return errors.New("failed to parse tx value")
	}

	txCandidate := txmgr.TxCandidate{
		TxData:   txData,
		To:       &valpoolAddr,
		GasLimit: 0,
		Value:    value,
	}

	_, err = txManager.Send(context.Background(), txCandidate)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}

	return nil
}
