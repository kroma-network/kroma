package chaincfg

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
)

// TestGetRollupConfig tests that the configs sourced from the superchain-registry match
// the configs that were embedded in the op-node manually before the superchain-registry was utilized.
//
// The superchain-registry repository is a work in progress.
// At a later date, it will be proposed to, and must be approved by, Optimism Governance.
// Until that time, the configuration described in the superchain-registry is subject to change.
//
// This test ensures no op-node config-loading behavior changes before
// the superchain-registry is no longer deemed experimental.
func TestGetRollupConfig(t *testing.T) {
	configsByName := map[string]rollup.Config{
		/* [Kroma: START]
		"goerli":  goerliCfg,
		[Kroma: END] */
		"mainnet": mainnetCfg,
		"sepolia": sepoliaCfg,
	}

	for name, expectedCfg := range configsByName {
		gotCfg, err := GetRollupConfig(name)
		require.NoError(t, err)

		require.Equalf(t, expectedCfg, *gotCfg, "rollup-configs from superchain-registry must match for %v", name)
	}
}

var mainnetCfg = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0xe459c500b760ed52a1ad799bf578b257af2c76f6ebe061a4c62627e9c605bced"),
			Number: 18067255,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0xeab1dbcbd854942126643609f6b457e391b169c819b7e5d5042389ccf6012cbf"),
			Number: 0,
		},
		L2Time: 1693880387,
		SystemConfig: eth.SystemConfig{
			BatcherAddr:           common.HexToAddress("0x41b8cd6791de4d8f9e0eaf7861ac506822adce12"),
			Overhead:              eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000bc")),
			Scalar:                eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000a6fe0")),
			GasLimit:              30_000_000,
			ValidatorRewardScalar: eth.Bytes32(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000002710")),
		},
	},
	BlockTime:              2,
	MaxSequencerDrift:      600,
	SeqWindowSize:          3600,
	ChannelTimeout:         300,
	L1ChainID:              big.NewInt(1),
	L2ChainID:              big.NewInt(255),
	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000000000255"),
	DepositContractAddress: common.HexToAddress("0x31f648572b67e60ec6eb8e197e1848cc5f5558de"),
	L1SystemConfigAddress:  common.HexToAddress("0x3971eb866aa9b2b8afea8a7c816f3b7e8b195a35"),
	RegolithTime:           u64Ptr(0),
	CanyonTime:             u64Ptr(1708502400),
	DeltaTime:              u64Ptr(1709107200),
	EcotoneTime:            u64Ptr(1714032001),
	KromaMPTTime:           nil,
	FjordTime:              nil,
	InteropTime:            nil,
	/* [Kroma: START]
	ProtocolVersionsAddress: common.HexToAddress("0x8062AbC286f5e7D9428a0Ccb9AbD71e50d93b935"),
	[Kroma: END] */
}

var sepoliaCfg = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0x936e490e33e6e136ecd9095090e30ed7def3903ef2bae3e05966b376e493ad76"),
			Number: 3841490,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0x52ef8f66bb31c16326eb2072dd9b2fa734068728b845d5428f3a256a50bf252e"),
			Number: 0,
		},
		L2Time: 1688709132,
		SystemConfig: eth.SystemConfig{
			BatcherAddr:           common.HexToAddress("0xf15dc770221b99c98d4aaed568f2ab04b9d16e42"),
			Overhead:              eth.Bytes32(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000834")),
			Scalar:                eth.Bytes32(common.HexToHash("0x000000000000000000000000000000000000000000000000000000000016e360")),
			GasLimit:              30_000_000,
			ValidatorRewardScalar: eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000007d0")),
		},
	},
	BlockTime:              2,
	MaxSequencerDrift:      1200,
	SeqWindowSize:          3600,
	ChannelTimeout:         120,
	L1ChainID:              big.NewInt(11155111),
	L2ChainID:              big.NewInt(2358),
	BatchInboxAddress:      common.HexToAddress("0xfa79000000000000000000000000000000000001"),
	DepositContractAddress: common.HexToAddress("0x31ab8ed993a3be9aa2757c7d368dc87101a868a4"),
	L1SystemConfigAddress:  common.HexToAddress("0x398c8ea789968893095d86cba168378a4f452e33"),
	RegolithTime:           u64Ptr(0),
	CanyonTime:             u64Ptr(1707897600),
	DeltaTime:              u64Ptr(1708416000),
	EcotoneTime:            u64Ptr(1713340800),
	KromaMPTTime:           nil,
	FjordTime:              nil,
	InteropTime:            nil,
	/* [Kroma: START]
	ProtocolVersionsAddress: common.HexToAddress("0x79ADD5713B383DAa0a138d3C4780C7A1804a8090"),
	[Kroma: END] */
}
