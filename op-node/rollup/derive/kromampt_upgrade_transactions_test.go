package derive

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	oppredeploys "github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/kroma-bindings/predeploys"
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
		{
			source:       deployGasPriceOracleMPTSource,
			expectedHash: "0x68219c2277a00907673390a65684be3c56297f65eb168ffadcd3d14cce3204c6",
		},
		{
			source:       updateGasPriceOracleMPTProxySource,
			expectedHash: "0x30f9d5d060ae85936d70ce7e51c48f3d609304321fb412b181e8c2e06b68ba15",
		},
		{
			source:       enableKromaMPTSource,
			expectedHash: "0xe0b5bb4bc6b821c057531a3ca4e07aa657c7d593610939bd7eb9d679344ff562",
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
			name:    "LocalDevnet",
			chainId: big.NewInt(KromaLocalDevnetChainID),
		},
		{
			name:    "KromaMainnet",
			chainId: big.NewInt(params.KromaMainnetChainID),
		},
		{
			name:    "KromaSepolia",
			chainId: big.NewInt(params.KromaSepoliaChainID),
		},
		{
			name:    "KromaHolesky",
			chainId: big.NewInt(params.KromaDevnetChainID),
		},
		{
			name:    "UnsupportedChain",
			chainId: big.NewInt(9999),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			upgradeTxns, err := KromaMPTNetworkUpgradeTransactions(test.chainId)
			require.NoError(t, err)
			deploymentData, err := getDeploymentData(test.chainId)
			require.NoError(t, err)
			require.Len(t, upgradeTxns, KromaMPTUpgradeTxCount)

			deployBaseFeeVaultSender, deployBaseFeeVault := toDepositTxn(t, upgradeTxns[1])
			require.Equal(t, deployBaseFeeVaultSender, BaseFeeVaultDeployerAddress)
			require.Equal(t, deployBaseFeeVaultSource.SourceHash(), deployBaseFeeVault.SourceHash())
			require.Nil(t, deployBaseFeeVault.To())
			require.Equal(t, uint64(1_000_000), deployBaseFeeVault.Gas())
			require.Equal(t, hexutil.Bytes(deploymentData.baseFeeVaultDeploymentBytecodeWithArg).String(), hexutil.Bytes(deployBaseFeeVault.Data()).String())

			deployL1FeeVaultSender, deployL1FeeVault := toDepositTxn(t, upgradeTxns[2])
			require.Equal(t, deployL1FeeVaultSender, L1FeeVaultDeployerAddress)
			require.Equal(t, deployL1FeeVaultSource.SourceHash(), deployL1FeeVault.SourceHash())
			require.Nil(t, deployL1FeeVault.To())
			require.Equal(t, uint64(1_000_000), deployL1FeeVault.Gas())
			require.Equal(t, hexutil.Bytes(deploymentData.l1FeeVaultDeploymentBytecodeWithArg).String(), hexutil.Bytes(deployL1FeeVault.Data()).String())

			deploySequencerFeeVaultSender, deploySequencerFeeVault := toDepositTxn(t, upgradeTxns[3])
			require.Equal(t, deploySequencerFeeVaultSender, SequencerFeeVaultDeployerAddress)
			require.Equal(t, deploySequencerFeeVaultSource.SourceHash(), deploySequencerFeeVault.SourceHash())
			require.Nil(t, deploySequencerFeeVault.To())
			require.Equal(t, uint64(1_000_000), deploySequencerFeeVault.Gas())
			require.Equal(t, hexutil.Bytes(deploymentData.sequencerFeeVaultDeploymentBytecodeWithArg).String(), hexutil.Bytes(deploySequencerFeeVault.Data()).String())
		})
	}

	// test chainId-independent configurations
	t.Run("ChainIDIndependent", func(t *testing.T) {
		chainID := big.NewInt(KromaLocalDevnetChainID)
		upgradeTxns, err := KromaMPTNetworkUpgradeTransactions(chainID) // Use default configuration
		require.NoError(t, err)
		require.Len(t, upgradeTxns, KromaMPTUpgradeTxCount)

		deployL1BlockSender, deployL1Block := toDepositTxn(t, upgradeTxns[0])
		require.Equal(t, deployL1BlockSender, L1BlockMPTDeployerAddress)
		require.Equal(t, deployL1BlockMPTSource.SourceHash(), deployL1Block.SourceHash())
		require.Nil(t, deployL1Block.To())
		require.Equal(t, uint64(500_000), deployL1Block.Gas())
		require.Equal(t, hexutil.Bytes(l1BlockMPTDeploymentBytecode).String(), hexutil.Bytes(deployL1Block.Data()).String())

		deployGasPriceOracleSender, deployGasPriceOracle := toDepositTxn(t, upgradeTxns[4])
		require.Equal(t, deployGasPriceOracleSender, GasPriceOracleMPTDeployerAddress)
		require.Equal(t, deployGasPriceOracleMPTSource.SourceHash(), deployGasPriceOracle.SourceHash())
		require.Nil(t, deployGasPriceOracle.To())
		require.Equal(t, uint64(1_500_000), deployGasPriceOracle.Gas())
		require.Equal(t, hexutil.Bytes(gasPriceOracleMPTDeploymentBytecode).String(), hexutil.Bytes(deployGasPriceOracle.Data()).String())

		updateL1BlockProxySender, updateL1BlockProxy := toDepositTxn(t, upgradeTxns[5])
		require.Equal(t, updateL1BlockProxySender, common.Address{})
		require.Equal(t, updateL1BlockMPTProxySource.SourceHash(), updateL1BlockProxy.SourceHash())
		require.NotNil(t, updateL1BlockProxy.To())
		require.Equal(t, *updateL1BlockProxy.To(), oppredeploys.L1BlockAddr)
		require.Equal(t, uint64(50_000), updateL1BlockProxy.Gas())
		require.Equal(t, common.FromHex("0x3659cfe6000000000000000000000000cd7467a8926d13f8b41ea035ff3761326c822bad"), updateL1BlockProxy.Data())

		updateBaseFeeVaultProxySender, updateBaseFeeVaultProxy := toDepositTxn(t, upgradeTxns[6])
		require.Equal(t, updateBaseFeeVaultProxySender, common.Address{})
		require.Equal(t, updateBaseFeeVaultProxySource.SourceHash(), updateBaseFeeVaultProxy.SourceHash())
		require.NotNil(t, updateBaseFeeVaultProxy.To())
		require.Equal(t, *updateBaseFeeVaultProxy.To(), oppredeploys.BaseFeeVaultAddr)
		require.Equal(t, uint64(50_000), updateBaseFeeVaultProxy.Gas())
		require.Equal(t, common.FromHex("3659cfe6000000000000000000000000db78bab44d9632e68348659dd47b4806ba276d89"), updateBaseFeeVaultProxy.Data())

		updateL1FeeVaultProxySender, updateL1FeeVaultProxy := toDepositTxn(t, upgradeTxns[7])
		require.Equal(t, updateL1FeeVaultProxySender, common.Address{})
		require.Equal(t, updateL1FeeVaultProxySource.SourceHash(), updateL1FeeVaultProxy.SourceHash())
		require.NotNil(t, updateL1FeeVaultProxy.To())
		require.Equal(t, *updateL1FeeVaultProxy.To(), oppredeploys.L1FeeVaultAddr)
		require.Equal(t, uint64(50_000), updateL1FeeVaultProxy.Gas())
		require.Equal(t, common.FromHex("3659cfe6000000000000000000000000d86e1a7c380f398bda3d598ee65c891ce5c3c8f0"), updateL1FeeVaultProxy.Data())

		updateSequencerFeeVaultProxySender, updateSequencerFeeVaultProxy := toDepositTxn(t, upgradeTxns[8])
		require.Equal(t, updateSequencerFeeVaultProxySender, common.Address{})
		require.Equal(t, updateSequencerFeeVaultProxySource.SourceHash(), updateSequencerFeeVaultProxy.SourceHash())
		require.NotNil(t, updateSequencerFeeVaultProxy.To())
		require.Equal(t, *updateSequencerFeeVaultProxy.To(), oppredeploys.SequencerFeeVaultAddr)
		require.Equal(t, uint64(50_000), updateSequencerFeeVaultProxy.Gas())
		require.Equal(t, common.FromHex("3659cfe600000000000000000000000013963c74d9c62d31ce3fcec0d46f6430a74ea79e"), updateSequencerFeeVaultProxy.Data())

		updateGasPriceOracleProxySender, updateGasPriceOracleProxy := toDepositTxn(t, upgradeTxns[9])
		require.Equal(t, updateGasPriceOracleProxySender, common.Address{})
		require.Equal(t, updateGasPriceOracleMPTProxySource.SourceHash(), updateGasPriceOracleProxy.SourceHash())
		require.NotNil(t, updateGasPriceOracleProxy.To())
		require.Equal(t, *updateGasPriceOracleProxy.To(), predeploys.GasPriceOracleAddr)
		require.Equal(t, uint64(50_000), updateGasPriceOracleProxy.Gas())
		require.Equal(t, common.FromHex("3659cfe60000000000000000000000008aa737409ba5a40950ad531a8f22aedf8b9a00ac"), updateGasPriceOracleProxy.Data())

		gpoSetKromaMPTSender, gpoSetKromaMPT := toDepositTxn(t, upgradeTxns[10])
		require.Equal(t, gpoSetKromaMPTSender, L1InfoDepositerAddress)
		require.Equal(t, enableKromaMPTSource.SourceHash(), gpoSetKromaMPT.SourceHash())
		require.NotNil(t, gpoSetKromaMPT.To())
		require.Equal(t, *gpoSetKromaMPT.To(), predeploys.GasPriceOracleAddr)
		require.Equal(t, uint64(80_000), gpoSetKromaMPT.Gas())
		require.Equal(t, common.FromHex("0x8cca6762"), gpoSetKromaMPT.Data())
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
