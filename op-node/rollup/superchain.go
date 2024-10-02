package rollup

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/params"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/superchain-registry/superchain"
)

var OPStackSupport = params.ProtocolVersionV0{Build: [8]byte{}, Major: 6, Minor: 0, Patch: 0, PreRelease: 0}.Encode()

const (
	opMainnet = 10
	opGoerli  = 420
	opSepolia = 11155420

	labsGoerliDevnet   = 997
	labsGoerliChaosnet = 888
	labsSepoliaDevnet0 = 11155421

	baseGoerli  = 84531
	baseMainnet = 8453

	pgnMainnet = 424
	pgnSepolia = 58008
)

// LoadOPStackRollupConfig loads the rollup configuration of the requested chain ID from the superchain-registry.
// Some chains may require a SystemConfigProvider to retrieve any values not part of the registry.
func LoadOPStackRollupConfig(chainID uint64) (*Config, error) {
	chConfig, ok := superchain.OPChains[chainID]
	if !ok {
		return nil, fmt.Errorf("unknown chain ID: %d", chainID)
	}

	superChain, ok := superchain.Superchains[chConfig.Superchain]
	if !ok {
		return nil, fmt.Errorf("chain %d specifies unknown superchain: %q", chainID, chConfig.Superchain)
	}

	var genesisSysConfig eth.SystemConfig
	if sysCfg, ok := superchain.GenesisSystemConfigs[chainID]; ok {
		genesisSysConfig = eth.SystemConfig{
			BatcherAddr: common.Address(sysCfg.BatcherAddr),
			Overhead:    eth.Bytes32(sysCfg.Overhead),
			Scalar:      eth.Bytes32(sysCfg.Scalar),
			GasLimit:    sysCfg.GasLimit,
		}
	} else {
		return nil, fmt.Errorf("unable to retrieve genesis SystemConfig of chain %d", chainID)
	}

	addrs, ok := superchain.Addresses[chainID]
	if !ok {
		return nil, fmt.Errorf("unable to retrieve deposit contract address")
	}

	regolithTime := uint64(0)
	// three goerli testnets test-ran Bedrock and later upgraded to Regolith.
	// All other OP-Stack chains have Regolith enabled from the start.
	switch chainID {
	case baseGoerli:
		regolithTime = 1683219600
	case opGoerli:
		regolithTime = 1679079600
	case labsGoerliDevnet:
		regolithTime = 1677984480
	case labsGoerliChaosnet:
		regolithTime = 1692156862
	}

	cfg := &Config{
		Genesis: Genesis{
			L1: eth.BlockID{
				Hash:   common.Hash(chConfig.Genesis.L1.Hash),
				Number: chConfig.Genesis.L1.Number,
			},
			L2: eth.BlockID{
				Hash:   common.Hash(chConfig.Genesis.L2.Hash),
				Number: chConfig.Genesis.L2.Number,
			},
			L2Time:       chConfig.Genesis.L2Time,
			SystemConfig: genesisSysConfig,
		},
		// The below chain parameters can be different per OP-Stack chain,
		// but since none of the superchain chains differ, it's not represented in the superchain-registry yet.
		// This restriction on superchain-chains may change in the future.
		// Test/Alt configurations can still load custom rollup-configs when necessary.
		BlockTime:         2,
		MaxSequencerDrift: 600,
		SeqWindowSize:     3600,
		ChannelTimeout:    300,
		L1ChainID:         new(big.Int).SetUint64(superChain.Config.L1.ChainID),
		L2ChainID:         new(big.Int).SetUint64(chConfig.ChainID),
		RegolithTime:      &regolithTime,
		CanyonTime:        chConfig.CanyonTime,
		DeltaTime:         chConfig.DeltaTime,
		EcotoneTime:       chConfig.EcotoneTime,
		// TODO(seolaoh): uncomment this when geth updated
		// KromaMPTTime:           chConfig.KromaMPTTime,
		FjordTime:              chConfig.FjordTime,
		BatchInboxAddress:      common.Address(chConfig.BatchInboxAddr),
		DepositContractAddress: common.Address(addrs.OptimismPortalProxy),
		L1SystemConfigAddress:  common.Address(addrs.SystemConfigProxy),
	}
	if superChain.Config.ProtocolVersionsAddr != nil { // Set optional protocol versions address
		/* [Kroma: START]
		cfg.ProtocolVersionsAddress = common.Address(*superChain.Config.ProtocolVersionsAddr)
		[Kroma: END] */
	}
	if chainID == labsGoerliDevnet || chainID == labsGoerliChaosnet {
		cfg.ChannelTimeout = 120
		cfg.MaxSequencerDrift = 1200
	}
	if chainID == pgnSepolia {
		cfg.MaxSequencerDrift = 1000
		cfg.SeqWindowSize = 7200
	}
	return cfg, nil
}
