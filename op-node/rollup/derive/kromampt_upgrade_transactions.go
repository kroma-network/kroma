package derive

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	oppredeploys "github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/kroma-network/kroma/kroma-bindings/bindings"
)

var (
	L1BlockMPTDeployerAddress        = common.HexToAddress("0x2551000000000000000000000000000000000000")
	BaseFeeVaultDeployerAddress      = common.HexToAddress("0x2551000000000000000000000000000000000001")
	L1FeeVaultDeployerAddress        = common.HexToAddress("0x2551000000000000000000000000000000000002")
	SequencerFeeVaultDeployerAddress = common.HexToAddress("0x2551000000000000000000000000000000000003")

	newL1BlockMPTAddress        = crypto.CreateAddress(L1BlockMPTDeployerAddress, 0)
	newBaseFeeVaultAddress      = crypto.CreateAddress(BaseFeeVaultDeployerAddress, 0)
	newL1FeeVaultAddress        = crypto.CreateAddress(L1FeeVaultDeployerAddress, 0)
	newSequencerFeeVaultAddress = crypto.CreateAddress(SequencerFeeVaultDeployerAddress, 0)

	deployBaseFeeVaultSource           = UpgradeDepositSource{Intent: "KromaMPT: BaseFee Vault Deployment"}
	updateBaseFeeVaultProxySource      = UpgradeDepositSource{Intent: "KromaMPT: BaseFee Vault Proxy Update"}
	deployL1FeeVaultSource             = UpgradeDepositSource{Intent: "KromaMPT: L1 Fee Vault Deployment"}
	updateL1FeeVaultProxySource        = UpgradeDepositSource{Intent: "KromaMPT: L1 Fee Vault Proxy Update"}
	deploySequencerFeeVaultSource      = UpgradeDepositSource{Intent: "KromaMPT: Sequencer Fee Vault Deployment"}
	updateSequencerFeeVaultProxySource = UpgradeDepositSource{Intent: "KromaMPT: Sequencer Fee Vault Proxy Update"}
	deployL1BlockMPTSource             = UpgradeDepositSource{Intent: "KromaMPT: L1 Block Deployment"}
	updateL1BlockMPTProxySource        = UpgradeDepositSource{Intent: "KromaMPT: L1 Block Proxy Update"}

	l1BlockMPTDeploymentBytecode        = common.FromHex(bindings.L1BlockMetaData.Bin)
	sequencerFeeVaultDeploymentBytecode = common.FromHex(bindings.ProtocolVaultMetaData.Bin)
	baseFeeVaultDeploymentBytecode      = common.FromHex(bindings.ProtocolVaultMetaData.Bin)
	l1FeeVaultDeploymentBytecode        = common.FromHex(bindings.L1FeeVaultMetaData.Bin)
)

func KromaMPTNetworkUpgradeTransactions() ([]hexutil.Bytes, error) {
	upgradeTxns := make([]hexutil.Bytes, 0, 8)

	deployL1BlockTransaction, err := types.NewTx(&types.DepositTx{
		SourceHash:          deployL1BlockMPTSource.SourceHash(),
		From:                L1BlockMPTDeployerAddress,
		To:                  nil,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 375_000,
		IsSystemTransaction: false,
		Data:                l1BlockMPTDeploymentBytecode,
	}).MarshalBinary()
	if err != nil {
		return nil, err
	}

	upgradeTxns = append(upgradeTxns, deployL1BlockTransaction)

	deployBaseFeeVault, err := types.NewTx(&types.DepositTx{
		SourceHash:          deployBaseFeeVaultSource.SourceHash(),
		From:                BaseFeeVaultDeployerAddress,
		To:                  nil,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 1_000_000,
		IsSystemTransaction: false,
		Data:                baseFeeVaultDeploymentBytecode,
	}).MarshalBinary()
	if err != nil {
		return nil, err
	}

	upgradeTxns = append(upgradeTxns, deployBaseFeeVault)

	deployL1FeeVault, err := types.NewTx(&types.DepositTx{
		SourceHash:          deployL1FeeVaultSource.SourceHash(),
		From:                L1FeeVaultDeployerAddress,
		To:                  nil,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 1_000_000,
		IsSystemTransaction: false,
		Data:                l1FeeVaultDeploymentBytecode,
	}).MarshalBinary()
	if err != nil {
		return nil, err
	}

	upgradeTxns = append(upgradeTxns, deployL1FeeVault)

	deploySequencerFeeVault, err := types.NewTx(&types.DepositTx{
		SourceHash:          deploySequencerFeeVaultSource.SourceHash(),
		From:                SequencerFeeVaultDeployerAddress,
		To:                  nil,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 1_000_000,
		IsSystemTransaction: false,
		Data:                sequencerFeeVaultDeploymentBytecode,
	}).MarshalBinary()
	if err != nil {
		return nil, err
	}

	upgradeTxns = append(upgradeTxns, deploySequencerFeeVault)

	updateL1BlockProxy, err := types.NewTx(&types.DepositTx{
		SourceHash:          updateL1BlockMPTProxySource.SourceHash(),
		From:                common.Address{},
		To:                  &oppredeploys.L1BlockAddr,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 50_000,
		IsSystemTransaction: false,
		Data:                upgradeToCalldata(newL1BlockMPTAddress),
	}).MarshalBinary()
	if err != nil {
		return nil, err
	}

	upgradeTxns = append(upgradeTxns, updateL1BlockProxy)

	updateBaseFeeVaultProxy, err := types.NewTx(&types.DepositTx{
		SourceHash:          updateBaseFeeVaultProxySource.SourceHash(),
		From:                common.Address{},
		To:                  &oppredeploys.BaseFeeVaultAddr,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 50_000,
		IsSystemTransaction: false,
		Data:                upgradeToCalldata(newBaseFeeVaultAddress),
	}).MarshalBinary()
	if err != nil {
		return nil, err
	}

	upgradeTxns = append(upgradeTxns, updateBaseFeeVaultProxy)

	updateL1FeeVaultProxy, err := types.NewTx(&types.DepositTx{
		SourceHash:          updateL1FeeVaultProxySource.SourceHash(),
		From:                common.Address{},
		To:                  &oppredeploys.L1FeeVaultAddr,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 50_000,
		IsSystemTransaction: false,
		Data:                upgradeToCalldata(newL1FeeVaultAddress),
	}).MarshalBinary()
	if err != nil {
		return nil, err
	}

	upgradeTxns = append(upgradeTxns, updateL1FeeVaultProxy)

	updateSequencerFeeVaultProxy, err := types.NewTx(&types.DepositTx{
		SourceHash:          updateSequencerFeeVaultProxySource.SourceHash(),
		From:                common.Address{},
		To:                  &oppredeploys.SequencerFeeVaultAddr,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 50_000,
		IsSystemTransaction: false,
		Data:                upgradeToCalldata(newSequencerFeeVaultAddress),
	}).MarshalBinary()
	if err != nil {
		return nil, err
	}

	upgradeTxns = append(upgradeTxns, updateSequencerFeeVaultProxy)

	return upgradeTxns, nil
}
