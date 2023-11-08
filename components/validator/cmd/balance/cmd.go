package balance

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum-optimism/optimism/op-service/txmgr/metrics"
	"github.com/kroma-network/kroma/components/validator/flags"
	"github.com/kroma-network/kroma/utils"
)

func Deposit(ctx *cli.Context) error {
	depositAmount := ctx.Uint64("amount")

	valpoolABI, err := bindings.ValidatorPoolMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ValidatorPool ABI: %w", err)
	}

	txData, err := valpoolABI.Pack("deposit")
	if err != nil {
		return fmt.Errorf("failed to create deposit transaction data: %w", err)
	}

	if err = sendTransaction(ctx, txData, depositAmount); err != nil {
		return err
	}

	return nil
}

func Withdraw(ctx *cli.Context) error {
	withdrawAmount := ctx.Uint64("amount")

	valpoolABI, err := bindings.ValidatorPoolMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ValidatorPool ABI: %w", err)
	}

	txData, err := valpoolABI.Pack("withdraw", new(big.Int).SetUint64(withdrawAmount))
	if err != nil {
		return fmt.Errorf("failed to create withdraw transaction data: %w", err)
	}

	if err = sendTransaction(ctx, txData, 0); err != nil {
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

	if err = sendTransaction(ctx, txData, 0); err != nil {
		return err
	}

	return nil
}

func sendTransaction(ctx *cli.Context, txData []byte, txValue uint64) error {
	txMgrConfig := txmgr.ReadCLIConfig(ctx)
	txManager, err := txmgr.NewSimpleTxManager("validator-balance", log.New(), &metrics.NoopTxMetrics{}, txMgrConfig)
	if err != nil {
		return fmt.Errorf("failed to create tx manager: %w", err)
	}

	valpoolAddr, err := utils.ParseAddress(ctx.GlobalString(flags.ValPoolAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorPool address: %w", err)
	}

	txCandidate := txmgr.TxCandidate{
		TxData:   txData,
		To:       &valpoolAddr,
		GasLimit: 0,
		Value:    new(big.Int).SetUint64(txValue),
	}

	_, err = txManager.Send(context.Background(), txCandidate)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}

	return nil
}
