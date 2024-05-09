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
	commissionMaxChangeRate := uint8(ctx.Uint64("commission-max-change-rate"))

	valManagerAbi, err := bindings.ValidatorManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ValidatorPool ABI: %w", err)
	}

	txData, err := valManagerAbi.Pack("registerValidator", assets, commissionRate, commissionMaxChangeRate)
	if err != nil {
		return fmt.Errorf("failed to create register validator transaction data: %w", err)
	}

	valManagerAddr, err := opservice.ParseAddress(ctx.String(flags.ValManagerAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorManager address: %w", err)
	}

	if err = sendTransaction(txManager, valManagerAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func Activate(ctx *cli.Context) error {
	valManagerAbi, err := bindings.ValidatorManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ValidatorManager ABI: %w", err)
	}

	txData, err := valManagerAbi.Pack("activateValidator")
	if err != nil {
		return fmt.Errorf("failed to create activate transaction data: %w", err)
	}

	valManagerAddr, err := opservice.ParseAddress(ctx.String(flags.ValManagerAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorManager address: %w", err)
	}

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}

	if err = sendTransaction(txManager, valManagerAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func Unjail(ctx *cli.Context) error {
	valManagerAbi, err := bindings.ValidatorManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ValidatorManager ABI: %w", err)
	}

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}
	validatorAddr := txManager.Config.From

	txData, err := valManagerAbi.Pack("tryUnjail", validatorAddr, false)
	if err != nil {
		return fmt.Errorf("failed to create try unjail transaction data: %w", err)
	}

	valManagerAddr, err := opservice.ParseAddress(ctx.String(flags.ValManagerAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorManager address: %w", err)
	}

	if err = sendTransaction(txManager, valManagerAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func ChangeCommissionRate(ctx *cli.Context) error {
	commissionRate := uint8(ctx.Uint64("commission-rate"))

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}

	valManagerAbi, err := bindings.ValidatorManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ValidatorManager ABI: %w", err)
	}

	txData, err := valManagerAbi.Pack("changeCommissionRate", commissionRate)
	if err != nil {
		return fmt.Errorf("failed to create change commission rate transaction data: %w", err)
	}

	valManagerAddr, err := opservice.ParseAddress(ctx.String(flags.ValManagerAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorManager address: %w", err)
	}

	if err = sendTransaction(txManager, valManagerAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func Delegate(ctx *cli.Context) error {
	amount := ctx.String("amount")

	delegateAmount, success := new(big.Int).SetString(amount, 10)
	if !success {
		return fmt.Errorf("failed to parse delegate amount: %s", amount)
	}

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}
	validatorAddr := txManager.Config.From

	assetManagerAddr, err := opservice.ParseAddress(ctx.String(flags.AssetManagerAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse AssetManager address: %w", err)
	}

	if err = approve(ctx, delegateAmount, txManager, assetManagerAddr); err != nil {
		return fmt.Errorf("failed to approve assets: %w", err)
	}

	assetManagerAbi, err := bindings.AssetManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get AssetManager ABI: %w", err)
	}

	txData, err := assetManagerAbi.Pack("delegate", validatorAddr, delegateAmount)
	if err != nil {
		return fmt.Errorf("failed to create delegate transaction data: %w", err)
	}

	if err = sendTransaction(txManager, assetManagerAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func InitUndelegate(ctx *cli.Context) error {
	amount := ctx.String("amount")

	undelegateAmount, success := new(big.Int).SetString(amount, 10)
	if !success {
		return fmt.Errorf("failed to parse undelegate amount: %s", amount)
	}

	assetManagerAbi, err := bindings.AssetManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get AssetManager ABI: %w", err)
	}

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}
	validatorAddr := txManager.Config.From

	assetManagerAddr, err := opservice.ParseAddress(ctx.String(flags.AssetManagerAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse AssetManager address: %w", err)
	}

	assetManagerContract, err := bindings.NewAssetManagerCaller(assetManagerAddr, txManager.Backend.(*ethclient.Client))
	if err != nil {
		return fmt.Errorf("failed to fetch AssetManager contract: %w", err)
	}

	shares, err := assetManagerContract.PreviewDelegate(optsutils.NewSimpleCallOpts(ctx.Context), validatorAddr, undelegateAmount)
	if err != nil {
		return fmt.Errorf("failed to preview delegate: %w", err)
	}

	txData, err := assetManagerAbi.Pack("initUndelegate", validatorAddr, shares)
	if err != nil {
		return fmt.Errorf("failed to create init undelegate transaction data: %w", err)
	}

	if err = sendTransaction(txManager, assetManagerAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func FinalizeUndelegate(ctx *cli.Context) error {
	assetManagerAbi, err := bindings.AssetManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get AssetManager ABI: %w", err)
	}

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}
	validatorAddr := txManager.Config.From

	txData, err := assetManagerAbi.Pack("finalizeUndelegate", validatorAddr)
	if err != nil {
		return fmt.Errorf("failed to create finalize undelegate transaction data: %w", err)
	}

	assetManagerAddr, err := opservice.ParseAddress(ctx.String(flags.AssetManagerAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse AssetManager address: %w", err)
	}

	if err = sendTransaction(txManager, assetManagerAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func InitClaimValidatorReward(ctx *cli.Context) error {
	amount := ctx.String("amount")

	claimAmount, success := new(big.Int).SetString(amount, 10)
	if !success {
		return fmt.Errorf("failed to parse claim amount: %s", amount)
	}

	assetManagerAbi, err := bindings.AssetManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get AssetManager ABI: %w", err)
	}

	txData, err := assetManagerAbi.Pack("initClaimValidatorReward", claimAmount)
	if err != nil {
		return fmt.Errorf("failed to create claim validator rewards transaction data: %w", err)
	}

	assetManagerAddr, err := opservice.ParseAddress(ctx.String(flags.AssetManagerAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse AssetManager address: %w", err)
	}

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}

	if err = sendTransaction(txManager, assetManagerAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func FinalizeClaimValidatorReward(ctx *cli.Context) error {
	assetManagerAbi, err := bindings.AssetManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get AssetManager ABI: %w", err)
	}

	txData, err := assetManagerAbi.Pack("finalizeClaimValidatorReward")
	if err != nil {
		return fmt.Errorf("failed to create finalize claim validator rewards transaction data: %w", err)
	}

	assetManagerAddr, err := opservice.ParseAddress(ctx.String(flags.AssetManagerAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse AssetManager address: %w", err)
	}

	txManager, err := newTxManager(ctx)
	if err != nil {
		return err
	}

	if err = sendTransaction(txManager, assetManagerAddr, txData, "0"); err != nil {
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
