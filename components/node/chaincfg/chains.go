package chaincfg

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/rollup"
)

var Sepolia = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0x6b586b77f3fc109ae0820c917eee034c373386f9a182d6e636257947852c2216"),
			Number: 3197600,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0xa0428513c59752e86b882020a097e738bd7bd4b50a64f28ea605641b61a41f91"),
			Number: 0,
		},
		L2Time: 1680199224,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0xf1d8505e40e3f3dc57c104df7ad4e19b8f9d4165"),
			Overhead:    eth.Bytes32(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000834")),
			Scalar:      eth.Bytes32(common.HexToHash("0x000000000000000000000000000000000000000000000000000000000016e360")),
			GasLimit:    25_000_000,
		},
	},
	BlockTime:              2,
	MaxProposerDrift:       1200,
	ProposerWindowSize:     3600,
	ChannelTimeout:         120,
	L1ChainID:              big.NewInt(11155111),
	L2ChainID:              big.NewInt(2357),
	BatchInboxAddress:      common.HexToAddress("0xbac0000000000000000000000000000000000003"),
	DepositContractAddress: common.HexToAddress("0x9c818e93c0884f75f48d93a9bdb2e994f8d77b86"),
	L1SystemConfigAddress:  common.HexToAddress("0x472f1b9ea60e3ec09bc84b45b381d502a2ab51f6"),
	BlueTime:               u64Ptr(1683693240), // GMT: Wednesday, May 10, 2023 4:34:00 AM
}

var NetworksByName = map[string]rollup.Config{
	"sepolia": Sepolia,
}

var L2ChainIDToNetworkName = func() map[string]string {
	out := make(map[string]string)
	for name, netCfg := range NetworksByName {
		out[netCfg.L2ChainID.String()] = name
	}
	return out
}()

func AvailableNetworks() []string {
	var networks []string
	for name := range NetworksByName {
		networks = append(networks, name)
	}
	return networks
}

func GetRollupConfig(name string) (rollup.Config, error) {
	network, ok := NetworksByName[name]
	if !ok {
		return rollup.Config{}, fmt.Errorf("invalid network %s", name)
	}

	return network, nil
}

func u64Ptr(v uint64) *uint64 {
	return &v
}
