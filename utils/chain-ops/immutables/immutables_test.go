package immutables_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/utils/chain-ops/immutables"
)

func TestBuildKroma(t *testing.T) {
	results, err := immutables.BuildKroma(immutables.ImmutableConfig{
		"L2StandardBridge": {
			"otherBridge": common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"L2CrossDomainMessenger": {
			"otherMessenger": common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"L2ERC721Bridge": {
			"otherBridge": common.HexToAddress("0x1234567890123456789012345678901234567890"),
			"messenger":   common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"KromaMintableERC721Factory": {
			"remoteChainId": big.NewInt(1),
			"bridge":        common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"ValidatorRewardVault": {
			"recipient": common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"ProposerRewardVault": {
			"recipient": common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
		"ProtocolVault": {
			"recipient": common.HexToAddress("0x1234567890123456789012345678901234567890"),
		},
	}, false)
	require.Nil(t, err)
	require.NotNil(t, results)

	contracts := map[string]bool{
		"GasPriceOracle":             true,
		"L1Block":                    true,
		"L2CrossDomainMessenger":     true,
		"L2StandardBridge":           true,
		"L2ToL1MessagePasser":        true,
		"ValidatorRewardVault":       true,
		"ProtocolVault":              true,
		"ProposerRewardVault":        true,
		"KromaMintableERC20Factory":  true,
		"L2ERC721Bridge":             true,
		"KromaMintableERC721Factory": true,
	}

	// Only the exact contracts that we care about are being
	// modified
	require.Equal(t, len(results), len(contracts))

	for name, bytecode := range results {
		// There is bytecode there
		require.Greater(t, len(bytecode), 0)
		// It is in the set of contracts that we care about
		require.True(t, contracts[name])
	}
}
