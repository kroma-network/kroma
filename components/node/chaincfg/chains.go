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
			Hash:   common.HexToHash("0x936e490e33e6e136ecd9095090e30ed7def3903ef2bae3e05966b376e493ad76"),
			Number: 3841490,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0x52ef8f66bb31c16326eb2072dd9b2fa734068728b845d5428f3a256a50bf252e"),
			Number: 0,
		},
		L2Time: 1688709132,
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
	BatchInboxAddress:      common.HexToAddress("0xfa79000000000000000000000000000000000001"),
	DepositContractAddress: common.HexToAddress("0x31ab8ed993a3be9aa2757c7d368dc87101a868a4"),
	L1SystemConfigAddress:  common.HexToAddress("0x398c8ea789968893095d86cba168378a4f452e33"),
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
