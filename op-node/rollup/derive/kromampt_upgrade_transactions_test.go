package derive

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	// Test chainId-specific configurations
	tests := []struct {
		name    string
		chainId *big.Int
	}{
		{
			name:    "Devnet Configuration",
			chainId: big.NewInt(901),
		},
		{
			name:    "Kroma Configuration",
			chainId: big.NewInt(255),
		},
		{
			name:    "Kroma Sepolia Configuration",
			chainId: big.NewInt(2358),
		},
		{
			name:    "Kroma Holesky Configuration",
			chainId: big.NewInt(7791),
		},
		{
			name:    "Unsupported Chain Defaults to Devnet",
			chainId: big.NewInt(9999),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			upgradeTxns, err := KromaMPTNetworkUpgradeTransactions(tt.chainId)
			deploymentData := getDeploymentData(tt.chainId)
			require.NoError(t, err)
			require.Len(t, upgradeTxns, 8)

			deployBaseFeeVaultSender, deployBaseFeeVault := toDepositTxn(t, upgradeTxns[1])
			require.Equal(t, deployBaseFeeVaultSender, BaseFeeVaultDeployerAddress)
			require.Equal(t, deployBaseFeeVaultSource.SourceHash(), deployBaseFeeVault.SourceHash())
			require.Nil(t, deployBaseFeeVault.To())
			require.Equal(t, uint64(1_000_000), deployBaseFeeVault.Gas())
			require.Equal(t, "0x"+hex.EncodeToString(deploymentData.BaseFeeVaultDeploymentBytecodeWithArg), hexutil.Bytes(deployBaseFeeVault.Data()).String())

			deployL1FeeVaultSender, deployL1FeeVault := toDepositTxn(t, upgradeTxns[2])
			require.Equal(t, deployL1FeeVaultSender, L1FeeVaultDeployerAddress)
			require.Equal(t, deployL1FeeVaultSource.SourceHash(), deployL1FeeVault.SourceHash())
			require.Nil(t, deployL1FeeVault.To())
			require.Equal(t, uint64(1_000_000), deployL1FeeVault.Gas())
			require.Equal(t, "0x"+hex.EncodeToString(deploymentData.L1FeeVaultDeploymentBytecodeWithArg), hexutil.Bytes(deployL1FeeVault.Data()).String())

			deploySequencerFeeVaultSender, deploySequencerFeeVault := toDepositTxn(t, upgradeTxns[3])
			require.Equal(t, deploySequencerFeeVaultSender, SequencerFeeVaultDeployerAddress)
			require.Equal(t, deploySequencerFeeVaultSource.SourceHash(), deploySequencerFeeVault.SourceHash())
			require.Nil(t, deploySequencerFeeVault.To())
			require.Equal(t, uint64(1_000_000), deploySequencerFeeVault.Gas())
			require.Equal(t, "0x"+hex.EncodeToString(deploymentData.SequencerFeeVaultDeploymentBytecodeWithArg), hexutil.Bytes(deploySequencerFeeVault.Data()).String())
		})
	}

	// Test chainId-independent configurations
	t.Run("ChainId-Independent Configurations", func(t *testing.T) {
		chainID := big.NewInt(1)
		deploymentData := getDeploymentData(chainID)
		upgradeTxns, err := KromaMPTNetworkUpgradeTransactions(chainID) // Use default configuration
		require.NoError(t, err)
		require.Len(t, upgradeTxns, 8)

		deployL1BlockSender, deployL1Block := toDepositTxn(t, upgradeTxns[0])
		require.Equal(t, deployL1BlockSender, L1BlockMPTDeployerAddress)
		require.Equal(t, deployL1BlockMPTSource.SourceHash(), deployL1Block.SourceHash())
		require.Nil(t, deployL1Block.To())
		require.Equal(t, uint64(500_000), deployL1Block.Gas())
		require.Equal(t, hexutil.Bytes(deploymentData.l1BlockMPTDeploymentBytecode).String(), hexutil.Bytes(deployL1Block.Data()).String())

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

func TestKromaMPTNetworkUpgradeTransactions(t *testing.T) {
	tests := []struct {
		name         string
		chainId      *big.Int
		expectedCode ChainConfig
	}{
		{
			name:         "Mainnet Configuration",
			chainId:      big.NewInt(1),
			expectedCode: ChainConfigs["1"],
		},
		{
			name:         "Goerli Configuration",
			chainId:      big.NewInt(5),
			expectedCode: ChainConfigs["5"],
		},
		{
			name:         "Unsupported Chain Defaults to Devnet",
			chainId:      big.NewInt(9999),
			expectedCode: ChainConfigs["1"],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			upgradeTxns, err := KromaMPTNetworkUpgradeTransactions(tt.chainId)
			require.NoError(t, err, "expected no error for chain ID %d", tt.chainId)
			require.NotNil(t, upgradeTxns, "expected transactions for chain ID %d", tt.chainId)

			// Validate the bytecode in the transactions
			require.Equal(t, tt.expectedCode.BaseFeeVaultDeploymentBytecode, ChainConfigs[tt.chainId.String()].BaseFeeVaultDeploymentBytecode, "BaseFeeVaultBytecode mismatch for chain ID %d", tt.chainId)
			require.Equal(t, tt.expectedCode.L1FeeVaultDeploymentBytecode, ChainConfigs[tt.chainId.String()].L1FeeVaultDeploymentBytecode, "L1FeeVaultBytecode mismatch for chain ID %d", tt.chainId)
			require.Equal(t, tt.expectedCode.SequencerFeeVaultDeploymentBytecode, ChainConfigs[tt.chainId.String()].SequencerFeeVaultDeploymentBytecode, "SequencerFeeVaultBytecode mismatch for chain ID %d", tt.chainId)
		})
	}
}

func TestCreateDeploymentBytecode(t *testing.T) {
	// example deployment bytecode, constructor argument
	bytecode := "0x61012060405234801561001157600080fd5b50604051610b0f380380610b0f8339810160408190526100309161005b565b600060808190526001600160a01b0390911660a052600160c081905260e0919091526101005261008b565b60006020828403121561006d57600080fd5b81516001600160a01b038116811461008457600080fd5b9392505050565b60805160a05160c05160e05161010051610a176100f860003960006103620152600061033901526000610310015260008181608701528181610191015281816102a7015281816103c4015281816104840152610623015260008181610157015261052b0152610a176000f3fe6080604052600436106100695760003560e01c80636ed39f62116100435780636ed39f621461010c57806384411d6514610121578063d3e5792b1461014557600080fd5b80630d9019e1146100755780633ccfd60b146100d357806354fd4d50146100ea57600080fd5b3661007057005b600080fd5b34801561008157600080fd5b506100a97f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156100df57600080fd5b506100e8610179565b005b3480156100f657600080fd5b506100ff610309565b6040516100ca91906108c8565b34801561011857600080fd5b506100e86103ac565b34801561012d57600080fd5b5061013760005481565b6040519081526020016100ca565b34801561015157600080fd5b506101377f000000000000000000000000000000000000000000000000000000000000000081565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610243576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f4665655661756c743a20746865206f6e6c7920726563697069656e742063616e60448201527f2063616c6c00000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b600061024d610527565b604080516020810182526000815290517fe11013dd0000000000000000000000000000000000000000000000000000000081529192507342000000000000000000000000000000000000099163e11013dd9184916102d4917f0000000000000000000000000000000000000000000000000000000000000000916188b891906004016108e2565b6000604051808303818588803b1580156102ed57600080fd5b505af1158015610301573d6000803e3d6000fd5b505050505050565b60606103347f0000000000000000000000000000000000000000000000000000000000000000610693565b61035d7f0000000000000000000000000000000000000000000000000000000000000000610693565b6103867f0000000000000000000000000000000000000000000000000000000000000000610693565b60405160200161039893929190610926565b604051602081830303815290604052905090565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610471576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f4665655661756c743a20746865206f6e6c7920726563697069656e742063616e60448201527f2063616c6c000000000000000000000000000000000000000000000000000000606482015260840161023a565b600061047b610527565b905060006104ba7f00000000000000000000000000000000000000000000000000000000000000005a8460405180602001604052806000815250610751565b905080610523576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4665655661756c743a20455448207472616e73666572206661696c6564000000604482015260640161023a565b5050565b60007f00000000000000000000000000000000000000000000000000000000000000004710156105ff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f4665655661756c743a207769746864726177616c20616d6f756e74206d75737460448201527f2062652067726561746572207468616e206d696e696d756d207769746864726160648201527f77616c20616d6f756e7400000000000000000000000000000000000000000000608482015260a40161023a565b600047905080600080828254610615919061099c565b9091555050604080518281527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166020820152338183015290517fc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba9181900360600190a1919050565b606060006106a08361076b565b600101905060008167ffffffffffffffff8111156106c0576106c06109db565b6040519080825280601f01601f1916602001820160405280156106ea576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846106f457509392505050565b600080600080845160208601878a8af19695505050505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106107b4577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef810000000083106107e0576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106107fe57662386f26fc10000830492506010015b6305f5e1008310610816576305f5e100830492506008015b612710831061082a57612710830492506004015b6064831061083c576064830492506002015b600a8310610848576001015b92915050565b60005b83811015610869578181015183820152602001610851565b83811115610878576000848401525b50505050565b6000815180845261089681602086016020860161084e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006108db602083018461087e565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff8316602082015260606040820152600061091d606083018461087e565b95945050505050565b6000845161093881846020890161084e565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610974816001850160208a0161084e565b6001920191820152835161098f81600284016020880161084e565b0160020195945050505050565b600082198211156109d6577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c634300080f000a"
	constructorArg := "0xA03c13C6597a0716D1525b7fDaD2fD95ECb49081"

	// expected ABI-encoded argument
	expectedEncodedArg := "000000000000000000000000a03c13c6597a0716d1525b7fdad2fd95ecb49081"
	encodedArg, err := encodeConstructorArg(constructorArg)
	require.NoError(t, err)
	require.Equal(t, expectedEncodedArg, encodedArg)

	// create deployment bytecode
	deploymentBytecode := createDeploymentBytecode(common.FromHex(bytecode), common.HexToAddress(constructorArg))

	// Validate deployment bytecode
	expectedDeploymentBytecodeHex := bytecode + expectedEncodedArg
	actualDeploymentBytecodeHex := "0x" + hex.EncodeToString(deploymentBytecode)
	require.Equal(t, expectedDeploymentBytecodeHex, actualDeploymentBytecodeHex, "deployment bytecode mismatch")

	// Validate the split between base bytecode and constructor argument
	actualBytecode := actualDeploymentBytecodeHex[:len(bytecode)]
	actualArg := actualDeploymentBytecodeHex[len(bytecode):]
	require.Equal(t, bytecode, actualBytecode, "base bytecode invalid")
	require.Equal(t, expectedEncodedArg, actualArg, "constructor argument invalid")
}

// Function to encode constructor arguments
func encodeConstructorArg(arg string) (string, error) {
	// Define ABI argument as an address type
	addressType, _ := abi.NewType("address", "", nil)
	constructorArgs := abi.Arguments{
		{
			Type: addressType,
		},
	}

	// Pack the argument
	encodedArgs, err := constructorArgs.Pack(common.HexToAddress(arg))
	if err != nil {
		return "", fmt.Errorf("failed to encode constructor argument: %w", err)
	}

	// Return encoded arguments as a hex string
	return hex.EncodeToString(encodedArgs), nil
}
