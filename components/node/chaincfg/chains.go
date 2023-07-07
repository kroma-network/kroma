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
			Hash:   common.HexToHash("0xd15978faa2adb2fe8678ab7bba43d288b355240b64389ca65f67d41df2dc40ac"),
			Number: 3835030,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0x4d5e05fc1480f767a0a6a33a69c3b66ac3ec382337731db22b6714c632046faa"),
			Number: 0,
		},
		L2Time: 1688629152,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0xf15dc770221b99c98d4aaed568f2ab04b9d16e42"),
			Overhead:    eth.Bytes32(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000834")),
			Scalar:      eth.Bytes32(common.HexToHash("0x000000000000000000000000000000000000000000000000000000000016e360")),
			GasLimit:    30_000_000,
		},
	},
	BlockTime:              2,
	MaxProposerDrift:       1200,
	ProposerWindowSize:     3600,
	ChannelTimeout:         120,
	L1ChainID:              big.NewInt(11155111),
	L2ChainID:              big.NewInt(2358),
	BatchInboxAddress:      common.HexToAddress("0xfa79000000000000000000000000000000000000"),
	DepositContractAddress: common.HexToAddress("0x16ceb19a3abf1a8b56f53db50eb22695b6ef7bcc"),
	L1SystemConfigAddress:  common.HexToAddress("0x29eeb4681bde3b2d0e739c5d82c5908dd0769b61"),
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
