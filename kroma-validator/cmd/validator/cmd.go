package validator

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	opservice "github.com/ethereum-optimism/optimism/op-service"
	"github.com/ethereum-optimism/optimism/op-service/optsutils"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum-optimism/optimism/op-service/txmgr/metrics"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
	"github.com/kroma-network/kroma/kroma-validator/flags"
)

func Register(ctx *cli.Context) error {
	amount := ctx.String("amount")

	assets, success := new(big.Int).SetString(amount, 10)
	if !success {
		return fmt.Errorf("failed to parse amount: %s", amount)
	}

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}

	assetManagerAddr, err := opservice.ParseAddress(ctx.String(flags.AssetManagerAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse AssetManager address: %w", err)
	}

	if err = approve(ctx, assets, txManager, assetManagerAddr); err != nil {
		return fmt.Errorf("failed to approve assets: %w", err)
	}

	commissionRate := uint8(ctx.Uint64("commission-rate"))
	withdrawAccount, err := opservice.ParseAddress(ctx.String(WithdrawAccountFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse withdraw address: %w", err)
	}

	valMgrAbi, err := bindings.ValidatorManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ValidatorManager ABI: %w", err)
	}

	txData, err := valMgrAbi.Pack("registerValidator", assets, commissionRate, withdrawAccount)
	if err != nil {
		return fmt.Errorf("failed to create register validator transaction data: %w", err)
	}

	valMgrAddr, err := opservice.ParseAddress(ctx.String(flags.ValMgrAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorManager address: %w", err)
	}

	if err = sendTransaction(txManager, valMgrAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func Activate(ctx *cli.Context) error {
	valMgrAbi, err := bindings.ValidatorManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ValidatorManager ABI: %w", err)
	}

	txData, err := valMgrAbi.Pack("activateValidator")
	if err != nil {
		return fmt.Errorf("failed to create activate transaction data: %w", err)
	}

	valMgrAddr, err := opservice.ParseAddress(ctx.String(flags.ValMgrAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorManager address: %w", err)
	}

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}

	if err = sendTransaction(txManager, valMgrAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func Unjail(ctx *cli.Context) error {
	valMgrAbi, err := bindings.ValidatorManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ValidatorManager ABI: %w", err)
	}

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}

	txData, err := valMgrAbi.Pack("tryUnjail")
	if err != nil {
		return fmt.Errorf("failed to create try unjail transaction data: %w", err)
	}

	valMgrAddr, err := opservice.ParseAddress(ctx.String(flags.ValMgrAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorManager address: %w", err)
	}

	if err = sendTransaction(txManager, valMgrAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func InitCommissionChange(ctx *cli.Context) error {
	commissionRate := uint8(ctx.Uint64("commission-rate"))

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}

	valMgrAbi, err := bindings.ValidatorManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ValidatorManager ABI: %w", err)
	}

	txData, err := valMgrAbi.Pack("initCommissionChange", commissionRate)
	if err != nil {
		return fmt.Errorf("failed to create init commission change transaction data: %w", err)
	}

	valMgrAddr, err := opservice.ParseAddress(ctx.String(flags.ValMgrAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorManager address: %w", err)
	}

	if err = sendTransaction(txManager, valMgrAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func FinalizeCommissionChange(ctx *cli.Context) error {
	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}

	valMgrAbi, err := bindings.ValidatorManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ValidatorManager ABI: %w", err)
	}

	txData, err := valMgrAbi.Pack("finalizeCommissionChange")
	if err != nil {
		return fmt.Errorf("failed to create finalize commission change transaction data: %w", err)
	}

	valMgrAddr, err := opservice.ParseAddress(ctx.String(flags.ValMgrAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorManager address: %w", err)
	}

	if err = sendTransaction(txManager, valMgrAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func DepositKro(ctx *cli.Context) error {
	amount := ctx.String("amount")

	depositAmount, success := new(big.Int).SetString(amount, 10)
	if !success {
		return fmt.Errorf("failed to parse deposit amount: %s", amount)
	}

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}

	assetManagerAddr, err := opservice.ParseAddress(ctx.String(flags.AssetManagerAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse AssetManager address: %w", err)
	}

	if err = approve(ctx, depositAmount, txManager, assetManagerAddr); err != nil {
		return fmt.Errorf("failed to approve assets: %w", err)
	}

	assetManagerAbi, err := bindings.AssetManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get AssetManager ABI: %w", err)
	}

	txData, err := assetManagerAbi.Pack("deposit", depositAmount)
	if err != nil {
		return fmt.Errorf("failed to create deposit transaction data: %w", err)
	}

	if err = sendTransaction(txManager, assetManagerAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

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

	valpoolAddr, err := opservice.ParseAddress(ctx.String(flags.ValPoolAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorPool address: %w", err)
	}

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}

	if err = sendTransaction(txManager, valpoolAddr, txData, amount); err != nil {
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

	valpoolAddr, err := opservice.ParseAddress(ctx.String(flags.ValPoolAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorPool address: %w", err)
	}

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}

	if err = sendTransaction(txManager, valpoolAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func WithdrawTo(ctx *cli.Context) error {
	address := ctx.String("address")
	toAddr := common.HexToAddress(address)

	amount := ctx.String("amount")
	withdrawAmount, success := new(big.Int).SetString(amount, 10)
	if !success {
		return errors.New("failed to parse withdraw amount")
	}

	valpoolABI, err := bindings.ValidatorPoolMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ValidatorPool ABI: %w", err)
	}

	txData, err := valpoolABI.Pack("withdrawTo", toAddr, withdrawAmount)
	if err != nil {
		return fmt.Errorf("failed to create withdrawTo transaction data: %w", err)
	}

	valpoolAddr, err := opservice.ParseAddress(ctx.String(flags.ValPoolAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorPool address: %w", err)
	}

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}

	if err = sendTransaction(txManager, valpoolAddr, txData, "0"); err != nil {
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

	valpoolAddr, err := opservice.ParseAddress(ctx.String(flags.ValPoolAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorPool address: %w", err)
	}

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}

	if err = sendTransaction(txManager, valpoolAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func approve(ctx *cli.Context, amount *big.Int, txManager *txmgr.SimpleTxManager, assetManagerAddr common.Address) error {
	erc20Abi, err := bindings.ERC20MetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ERC20 ABI: %w", err)
	}

	txData, err := erc20Abi.Pack("approve", assetManagerAddr, amount)
	if err != nil {
		return fmt.Errorf("failed to create approve transaction data: %w", err)
	}

	assetManagerContract, err := bindings.NewAssetManagerCaller(assetManagerAddr, txManager.Backend.(*ethclient.Client))
	if err != nil {
		return fmt.Errorf("failed to fetch AssetManager contract: %w", err)
	}

	assetTokenAddr, err := assetManagerContract.ASSETTOKEN(optsutils.NewSimpleCallOpts(ctx.Context))
	if err != nil {
		return fmt.Errorf("failed to fetch asset token address: %w", err)
	}

	if err = sendTransaction(txManager, assetTokenAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func sendTransaction(txManager *txmgr.SimpleTxManager, txTo common.Address, txData []byte, txValue string) error {
	value, success := new(big.Int).SetString(txValue, 10)
	if !success {
		return errors.New("failed to parse tx value")
	}

	txCandidate := txmgr.TxCandidate{
		TxData:   txData,
		To:       &txTo,
		GasLimit: 0,
		Value:    value,
	}

	_, err := txManager.Send(context.Background(), txCandidate)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}

	return nil
}

func newTxManager(ctx *cli.Context) (*txmgr.SimpleTxManager, error) {
	txMgrConfig := txmgr.ReadCLIConfig(ctx)
	txManager, err := txmgr.NewSimpleTxManager("validator-cmd", log.New(), &metrics.NoopTxMetrics{}, txMgrConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create tx manager: %w", err)
	}
	return txManager, nil
}
