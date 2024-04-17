package validator

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

	valpoolAddr, err := opservice.ParseAddress(ctx.String(flags.ValPoolAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorPool address: %w", err)
	}

	if err = sendTransaction(ctx, valpoolAddr, txData, amount); err != nil {
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

	if err = sendTransaction(ctx, valpoolAddr, txData, "0"); err != nil {
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

	if err = sendTransaction(ctx, valpoolAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func Approve(ctx *cli.Context) error {
	amount := ctx.String("amount")

	approveAmount, success := new(big.Int).SetString(amount, 10)
	if !success {
		return fmt.Errorf("failed to parse approve amount: %s", amount)
	}

	erc20Abi, err := bindings.ERC20MetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ERC20 ABI: %w", err)
	}

	assetManagerAddr, err := opservice.ParseAddress(ctx.String(flags.AssetManagerAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse AssetManager address: %w", err)
	}

	txData, err := erc20Abi.Pack("approve", assetManagerAddr, approveAmount)
	if err != nil {
		return fmt.Errorf("failed to create approve transaction data: %w", err)
	}

	govTokenAddr, err := getGovTokenAddress(ctx)
	if err != nil {
		return fmt.Errorf("failed to get gov token address: %w", err)
	}

	if err = sendTransaction(ctx, govTokenAddr, txData, "0"); err != nil {
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

	validatorAddr, err := getValidatorAddress(ctx)
	if err != nil {
		return fmt.Errorf("failed to get validator address: %w", err)
	}

	assetManagerAbi, err := bindings.AssetManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get AssetManager ABI: %w", err)
	}

	txData, err := assetManagerAbi.Pack("delegate", validatorAddr, delegateAmount)
	if err != nil {
		return fmt.Errorf("failed to create delegate transaction data: %w", err)
	}

	assetManagerAddr, err := opservice.ParseAddress(ctx.String(flags.AssetManagerAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse AssetManager address: %w", err)
	}

	if err = sendTransaction(ctx, assetManagerAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func InitUndelegate(ctx *cli.Context) error {
	amount := ctx.String("amount")

	undelegateAmount, success := new(big.Int).SetString(amount, 10)
	if !success {
		return fmt.Errorf("failed to parse delegate amount: %s", amount)
	}

	assetManagerAbi, err := bindings.AssetManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get AssetManager ABI: %w", err)
	}

	validatorAddr, err := getValidatorAddress(ctx)
	if err != nil {
		return fmt.Errorf("failed to get validator address: %w", err)
	}

	// Convert the given amount to shares by calling previewDelegate function.
	callData, err := assetManagerAbi.Pack("previewDelegate", validatorAddr, undelegateAmount)
	if err != nil {
		return fmt.Errorf("failed to create convert to share transaction data: %w", err)
	}

	assetManagerAddr, err := opservice.ParseAddress(ctx.String(flags.AssetManagerAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse AssetManager address: %w", err)
	}

	returnData, err := callContract(ctx, callData, assetManagerAddr)
	if err != nil {
		return fmt.Errorf("failed to call previewUndelegate function: %w", err)
	}

	shares := new(big.Int).SetBytes(returnData)

	txData, err := assetManagerAbi.Pack("initUndelegate", validatorAddr, shares)
	if err != nil {
		return fmt.Errorf("failed to create undelegate transaction data: %w", err)
	}

	if err = sendTransaction(ctx, assetManagerAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func FinalizeUndelegate(ctx *cli.Context) error {
	assetManagerAbi, err := bindings.AssetManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get AssetManager ABI: %w", err)
	}

	validatorAddr, err := getValidatorAddress(ctx)
	if err != nil {
		return fmt.Errorf("failed to get validator address: %w", err)
	}

	txData, err := assetManagerAbi.Pack("finalizeUndelegate", validatorAddr)
	if err != nil {
		return fmt.Errorf("failed to create finalize undelegate transaction data: %w", err)
	}

	assetManagerAddr, err := opservice.ParseAddress(ctx.String(flags.AssetManagerAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse AssetManager address: %w", err)
	}

	if err = sendTransaction(ctx, assetManagerAddr, txData, "0"); err != nil {
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

	if err = sendTransaction(ctx, assetManagerAddr, txData, "0"); err != nil {
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

	if err = sendTransaction(ctx, assetManagerAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func RegisterValidator(ctx *cli.Context) error {
	amount := ctx.String("amount")

	assets, success := new(big.Int).SetString(amount, 10)
	if !success {
		return fmt.Errorf("failed to parse amount: %s", amount)
	}

	commissionRate := uint8(ctx.Uint64("commissionRate"))
	commissionMaxChangeRate := uint8(ctx.Uint64("commissionMaxChangeRate"))

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

	if err = sendTransaction(ctx, valManagerAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func Unjail(ctx *cli.Context) error {
	valManagerAbi, err := bindings.ValidatorManagerMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ValidatorManager ABI: %w", err)
	}

	validatorAddr, err := getValidatorAddress(ctx)
	if err != nil {
		return fmt.Errorf("failed to get validator address: %w", err)
	}

	txData, err := valManagerAbi.Pack("tryUnjail", validatorAddr, false)
	if err != nil {
		return fmt.Errorf("failed to create try unjail transaction data: %w", err)
	}

	valManagerAddr, err := opservice.ParseAddress(ctx.String(flags.ValManagerAddressFlag.Name))
	if err != nil {
		return fmt.Errorf("failed to parse ValidatorManager address: %w", err)
	}

	if err = sendTransaction(ctx, valManagerAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func ChangeCommissionRate(ctx *cli.Context) error {
	commissionRate := uint8(ctx.Uint64("commissionRate"))

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

	if err = sendTransaction(ctx, valManagerAddr, txData, "0"); err != nil {
		return err
	}

	return nil
}

func sendTransaction(ctx *cli.Context, txTo common.Address, txData []byte, txValue string) error {
	txMgrConfig := txmgr.ReadCLIConfig(ctx)
	txManager, err := txmgr.NewSimpleTxManager("validator-balance", log.New(), &metrics.NoopTxMetrics{}, txMgrConfig)
	if err != nil {
		return fmt.Errorf("failed to create tx manager: %w", err)
	}

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

	_, err = txManager.Send(context.Background(), txCandidate)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}

	return nil
}

func getValidatorAddress(ctx *cli.Context) (common.Address, error) {
	txMgrConfig := txmgr.ReadCLIConfig(ctx)
	txManager, err := txmgr.NewSimpleTxManager("validator-balance", log.New(), &metrics.NoopTxMetrics{}, txMgrConfig)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to create tx manager: %w", err)
	}

	return txManager.Config.From, nil
}

func getGovTokenAddress(ctx *cli.Context) (common.Address, error) {
	assetManagerAbi, err := bindings.AssetManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get AssetManager ABI: %w", err)
	}

	callData, err := assetManagerAbi.Pack("ASSET_TOKEN")
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to create convert to share transaction data: %w", err)
	}

	assetManagerAddr, err := opservice.ParseAddress(ctx.String(flags.AssetManagerAddressFlag.Name))
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to parse AssetManager address: %w", err)
	}

	returnData, err := callContract(ctx, callData, assetManagerAddr)

	return common.BytesToAddress(returnData), nil
}

func callContract(ctx *cli.Context, callData []byte, contractAddr common.Address) ([]byte, error) {
	txMgrConfig := txmgr.ReadCLIConfig(ctx)
	txManager, err := txmgr.NewSimpleTxManager("validator-balance", log.New(), &metrics.NoopTxMetrics{}, txMgrConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create tx manager: %w", err)
	}

	return txManager.Backend.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddr,
		Data: callData,
	}, nil)
}
