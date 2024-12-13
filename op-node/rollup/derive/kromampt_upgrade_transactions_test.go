package derive

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/params"
)

func TestMPTSourcesMatchSpec(t *testing.T) {
	for _, test := range []struct {
		source       UpgradeDepositSource
		expectedHash string
	}{
		{
			source:       deployBaseFeeVaultSource,
			expectedHash: "0xdd882729c9833d55eb4ad39f08f578c0b92e171234e1b4ddc8f91d7878cd8686",
		},
		{
			source:       updateBaseFeeVaultProxySource,
			expectedHash: "0x747e2dc0ea5ed790a8746601b59c7588bfdce93e5d4dc0f6fd92a4d32c07596d",
		},
		{
			source:       deployL1FeeVaultSource,
			expectedHash: "0xb88c42fa962a1aa3d495761f85b93fd496a9607b57ffda85a1e3840b0d854794",
		},
		{
			source:       updateL1FeeVaultProxySource,
			expectedHash: "0xd4b39da88babd8b9c3d96bfc88b2c4be2a977c3b2629d0f2498fc290e1c523be",
		},
		{
			source:       deploySequencerFeeVaultSource,
			expectedHash: "0x4618df4b1692fcaee1da6b0532641b361f5dbfc3bb5714ae22802f989dc59c6f",
		},
		{
			source:       updateSequencerFeeVaultProxySource,
			expectedHash: "0xaf7e0d5c4b92ad5918fe238414031d8a82b87841ed70f1ba247d50121f9c04eb",
		},
		{
			source:       deployL1BlockMPTSource,
			expectedHash: "0x3c459db0fc6553f1b48a3b7f1e20de2e80d7f28976085817070699d1502dbe74",
		},
		{
			source:       updateL1BlockMPTProxySource,
			expectedHash: "0xb3a45c33c31a2322c4c4db99788dcae68606b98e4c90ea9fb3d1bc93ddfbf3bf",
		},
	} {
		require.Equal(t, common.HexToHash(test.expectedHash), test.source.SourceHash())
	}
}

func TestMPTNetworkTransactions(t *testing.T) {
	// test chainId-specific configurations
	tests := []struct {
		name    string
		chainId *big.Int
	}{
		{
			name:    "Local devnet configuration",
			chainId: big.NewInt(KromaLocalDevnetChainID),
		},
		{
			name:    "Kroma configuration",
			chainId: big.NewInt(params.KromaMainnetChainID),
		},
		{
			name:    "Kroma Sepolia configuration",
			chainId: big.NewInt(params.KromaSepoliaChainID),
		},
		{
			name:    "Kroma Holesky configuration",
			chainId: big.NewInt(params.KromaDevnetChainID),
		},
		{
			name:    "Unsupported chain. Defaults to local devnet",
			chainId: big.NewInt(9999),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			upgradeTxns, err := KromaMPTNetworkUpgradeTransactions(tt.chainId)
			require.NoError(t, err)
			deploymentData, err := getDeploymentData(tt.chainId)
			require.NoError(t, err)
			require.Len(t, upgradeTxns, 8)

			deployBaseFeeVaultSender, deployBaseFeeVault := toDepositTxn(t, upgradeTxns[1])
			require.Equal(t, deployBaseFeeVaultSender, BaseFeeVaultDeployerAddress)
			require.Equal(t, deployBaseFeeVaultSource.SourceHash(), deployBaseFeeVault.SourceHash())
			require.Nil(t, deployBaseFeeVault.To())
			require.Equal(t, uint64(1_000_000), deployBaseFeeVault.Gas())
			require.Equal(t, fmt.Sprintf("0x%x", deploymentData.baseFeeVaultDeploymentBytecodeWithArg), hexutil.Bytes(deployBaseFeeVault.Data()).String())

			deployL1FeeVaultSender, deployL1FeeVault := toDepositTxn(t, upgradeTxns[2])
			require.Equal(t, deployL1FeeVaultSender, L1FeeVaultDeployerAddress)
			require.Equal(t, deployL1FeeVaultSource.SourceHash(), deployL1FeeVault.SourceHash())
			require.Nil(t, deployL1FeeVault.To())
			require.Equal(t, uint64(1_000_000), deployL1FeeVault.Gas())
			require.Equal(t, fmt.Sprintf("0x%x", deploymentData.l1FeeVaultDeploymentBytecodeWithArg), hexutil.Bytes(deployL1FeeVault.Data()).String())

			deploySequencerFeeVaultSender, deploySequencerFeeVault := toDepositTxn(t, upgradeTxns[3])
			require.Equal(t, deploySequencerFeeVaultSender, SequencerFeeVaultDeployerAddress)
			require.Equal(t, deploySequencerFeeVaultSource.SourceHash(), deploySequencerFeeVault.SourceHash())
			require.Nil(t, deploySequencerFeeVault.To())
			require.Equal(t, uint64(1_000_000), deploySequencerFeeVault.Gas())
			require.Equal(t, fmt.Sprintf("0x%x", deploymentData.sequencerFeeVaultDeploymentBytecodeWithArg), hexutil.Bytes(deploySequencerFeeVault.Data()).String())
		})
	}

	// test chainId-independent configurations
	t.Run("ChainId-Independent Configurations", func(t *testing.T) {
		chainID := big.NewInt(KromaLocalDevnetChainID)
		upgradeTxns, err := KromaMPTNetworkUpgradeTransactions(chainID) // Use default configuration
		require.NoError(t, err)
		require.Len(t, upgradeTxns, 8)

		deployL1BlockSender, deployL1Block := toDepositTxn(t, upgradeTxns[0])
		require.Equal(t, deployL1BlockSender, L1BlockMPTDeployerAddress)
		require.Equal(t, deployL1BlockMPTSource.SourceHash(), deployL1Block.SourceHash())
		require.Nil(t, deployL1Block.To())
		require.Equal(t, uint64(500_000), deployL1Block.Gas())
		require.Equal(t, hexutil.Bytes(l1BlockMPTDeploymentBytecode).String(), hexutil.Bytes(deployL1Block.Data()).String())

		updateL1BlockProxySender, updateL1BlockProxy := toDepositTxn(t, upgradeTxns[4])
		require.Equal(t, updateL1BlockProxySender, common.Address{})
		require.Equal(t, updateL1BlockMPTProxySource.SourceHash(), updateL1BlockProxy.SourceHash())
		require.NotNil(t, updateL1BlockProxy.To())
		require.Equal(t, *updateL1BlockProxy.To(), predeploys.L1BlockAddr)
		require.Equal(t, uint64(50_000), updateL1BlockProxy.Gas())
		require.Equal(t, common.FromHex("0x3659cfe6000000000000000000000000cd7467a8926d13f8b41ea035ff3761326c822bad"), updateL1BlockProxy.Data())

		updateBaseFeeVaultProxySender, updateBaseFeeVaultProxy := toDepositTxn(t, upgradeTxns[5])
		require.Equal(t, updateBaseFeeVaultProxySender, common.Address{})
		require.Equal(t, updateBaseFeeVaultProxySource.SourceHash(), updateBaseFeeVaultProxy.SourceHash())
		require.NotNil(t, updateBaseFeeVaultProxy.To())
		require.Equal(t, *updateBaseFeeVaultProxy.To(), predeploys.BaseFeeVaultAddr)
		require.Equal(t, uint64(50_000), updateBaseFeeVaultProxy.Gas())
		require.Equal(t, common.FromHex("3659cfe6000000000000000000000000db78bab44d9632e68348659dd47b4806ba276d89"), updateBaseFeeVaultProxy.Data())
	})
}

func TestCreateDeploymentBytecode(t *testing.T) {
	// example deployment bytecode, constructor argument
	bytecode := fmt.Sprintf("0x%x", "example bytecode")
	constructorArg := common.HexToAddress("0x1234")

	// expected ABI-encoded argument
	expectedEncodedArg := "0000000000000000000000000000000000000000000000000000000000001234"
	encodedArg, err := encodeConstructorArg(constructorArg)
	require.NoError(t, err)
	require.Equal(t, expectedEncodedArg, hex.EncodeToString(encodedArg))

	// create deployment bytecode
	deploymentBytecode, err := createDeploymentBytecode(common.FromHex(bytecode), constructorArg)
	require.NoError(t, err)

	// validate deployment bytecode
	expectedDeploymentBytecodeHex := bytecode + expectedEncodedArg
	actualDeploymentBytecodeHex := fmt.Sprintf("0x%s", hex.EncodeToString(deploymentBytecode))
	require.Equal(t, expectedDeploymentBytecodeHex, actualDeploymentBytecodeHex, "deployment bytecode mismatch")

	// validate the split between base bytecode and constructor argument
	actualBytecode := actualDeploymentBytecodeHex[:len(bytecode)]
	actualArg := actualDeploymentBytecodeHex[len(bytecode):]
	require.Equal(t, bytecode, actualBytecode, "base bytecode invalid")
	require.Equal(t, expectedEncodedArg, actualArg, "constructor argument invalid")
}
