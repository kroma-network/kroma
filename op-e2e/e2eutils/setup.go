package e2eutils

import (
	"math/big"
	"os"
	"path"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	genesis2 "github.com/ethereum-optimism/optimism/op-chain-ops/genesis"
	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
)

var testingJWTSecret = [32]byte{123}

// WriteDefaultJWT writes a testing JWT to the temporary directory of the test and returns the path to the JWT file.
func WriteDefaultJWT(t TestingBase) string {
	// Sadly the geth node config cannot load JWT secret from memory, it has to be a file
	jwtPath := path.Join(t.TempDir(), "jwt_secret")
	if err := os.WriteFile(jwtPath, []byte(hexutil.Encode(testingJWTSecret[:])), 0o600); err != nil {
		t.Fatalf("failed to prepare jwt file for geth: %v", err)
	}
	return jwtPath
}

func uint64ToBig(in uint64) *hexutil.Big {
	return (*hexutil.Big)(new(big.Int).SetUint64(in))
}

// DeployParams bundles the deployment parameters to generate further testing inputs with.
type DeployParams struct {
	DeployConfig   *genesis2.DeployConfig
	MnemonicConfig *MnemonicConfig
	Secrets        *Secrets
	Addresses      *Addresses
}

// TestParams parametrizes the most essential rollup configuration parameters
type TestParams struct {
	MaxSequencerDrift   uint64
	SequencerWindowSize uint64
	ChannelTimeout      uint64
	L1BlockTime         uint64
}

func MakeDeployParams(t require.TestingT, tp *TestParams) *DeployParams {
	mnemonicCfg := DefaultMnemonicConfig
	secrets, err := mnemonicCfg.Secrets()
	require.NoError(t, err)
	addresses := secrets.Addresses()
	deployConfig := &genesis2.DeployConfig{
		L1ChainID:   900,
		L2ChainID:   901,
		L2BlockTime: 2,

		MaxSequencerDrift:   tp.MaxSequencerDrift,
		SequencerWindowSize: tp.SequencerWindowSize,
		ChannelTimeout:      tp.ChannelTimeout,
		P2PSequencerAddress: addresses.SequencerP2P,
		BatchInboxAddress:   common.Address{0: 0x42, 19: 0xff}, // tbd
		BatchSenderAddress:  addresses.Batcher,

		ValidatorPoolTrustedValidator:   addresses.TrustedValidator,
		ValidatorPoolRequiredBondAmount: uint64ToBig(1),
		ValidatorPoolMaxUnbond:          10,
		ValidatorPoolRoundDuration:      4,

		L2OutputOracleSubmissionInterval: 4,
		L2OutputOracleStartingTimestamp:  -1,

		L1BlockTime:                 tp.L1BlockTime,
		L1GenesisBlockNonce:         0,
		CliqueSignerAddress:         common.Address{}, // proof of stake, no clique
		L1GenesisBlockTimestamp:     hexutil.Uint64(time.Now().Unix()),
		L1GenesisBlockGasLimit:      30_000_000,
		L1GenesisBlockDifficulty:    uint64ToBig(1),
		L1GenesisBlockMixHash:       common.Hash{},
		L1GenesisBlockCoinbase:      common.Address{},
		L1GenesisBlockNumber:        0,
		L1GenesisBlockGasUsed:       0,
		L1GenesisBlockParentHash:    common.Hash{},
		L1GenesisBlockBaseFeePerGas: uint64ToBig(1000_000_000), // 1 gwei
		FinalizationPeriodSeconds:   10,

		L2GenesisBlockNonce:         0,
		L2GenesisBlockGasLimit:      30_000_000,
		L2GenesisBlockDifficulty:    uint64ToBig(0),
		L2GenesisBlockMixHash:       common.Hash{},
		L2GenesisBlockNumber:        0,
		L2GenesisBlockGasUsed:       0,
		L2GenesisBlockParentHash:    common.Hash{},
		L2GenesisBlockBaseFeePerGas: uint64ToBig(1000_000_000),

		ColosseumCreationPeriodSeconds: 10,
		ColosseumBisectionTimeout:      120,
		ColosseumProvingTimeout:        480,
		ColosseumDummyHash:             common.HexToHash("0xa1235b834d6f1f78f78bc4db856fbc49302cce2c519921347600693021e087f7"),
		ColosseumMaxTxs:                100,
		ColosseumSegmentsLengths:       "2,2,3,3",

		SecurityCouncilOwners: []common.Address{
			addresses.Challenger1,
			addresses.Challenger2,
			addresses.TrustedValidator,
		},

		GasPriceOracleOverhead:      2100,
		GasPriceOracleScalar:        1000_000,
		DeploymentWaitConfirmations: 1,
		ValidatorRewardScalar:       5000,

		ProxyAdminOwner:        addresses.ProxyAdminOwner,
		ProtocolVaultRecipient: common.Address{19: 2},
		L1FeeVaultRecipient:    common.Address{19: 3},

		EIP1559Elasticity:  10,
		EIP1559Denominator: 50,

		FundDevAccounts: false,

		GovernorVotingDelayBlocks:          0,
		GovernorVotingPeriodBlocks:         25,
		GovernorProposalThreshold:          1,
		GovernorVotesQuorumFractionPercent: 51,
		TimeLockMinDelaySeconds:            1,

		ZKVerifierHashScalar: (*hexutil.Big)(hexutil.MustDecodeBig("0x1545b1bf82c58ee35648bd877da9c5010193e82b036b16bf382acf31bc2ab576")),
		ZKVerifierM56Px:      (*hexutil.Big)(hexutil.MustDecodeBig("0x15ae1a8e3b993dd9aadc8f9086d1ea239d4cd5c09cfa445f337e1b60d7b3eb87")),
		ZKVerifierM56Py:      (*hexutil.Big)(hexutil.MustDecodeBig("0x2c702ede24f9db8c8c9a439975facd3872a888c5f84f58b3b5f5a5623bac945a")),
	}

	// Configure the DeployConfig with the expected developer L1
	// addresses.
	if err := deployConfig.InitDeveloperDeployedAddresses(); err != nil {
		panic(err)
	}

	return &DeployParams{
		DeployConfig:   deployConfig,
		MnemonicConfig: mnemonicCfg,
		Secrets:        secrets,
		Addresses:      addresses,
	}
}

// DeploymentsL1 captures the L1 addresses used in the deployment,
// commonly just the developer predeploys during testing,
// but later deployed contracts may be used in some tests too.
type DeploymentsL1 struct {
	L1CrossDomainMessengerProxy common.Address
	L1StandardBridgeProxy       common.Address
	ValidatorPoolProxy          common.Address
	L2OutputOracleProxy         common.Address
	ColosseumProxy              common.Address
	SecurityCouncilProxy        common.Address
	KromaPortalProxy            common.Address
	SystemConfigProxy           common.Address
	SecurityCouncilTokenProxy   common.Address
	TimeLockProxy               common.Address
	UpgradeGovernorProxy        common.Address
}

// SetupData bundles the L1, L2, rollup and deployment configuration data: everything for a full test setup.
type SetupData struct {
	L1Cfg         *core.Genesis
	L2Cfg         *core.Genesis
	RollupCfg     *rollup.Config
	DeploymentsL1 DeploymentsL1
}

// AllocParams defines genesis allocations to apply on top of the genesis generated by deploy parameters.
// These allocations override existing allocations per account,
// i.e. the allocations are merged with AllocParams having priority.
type AllocParams struct {
	L1Alloc          core.GenesisAlloc
	L2Alloc          core.GenesisAlloc
	PrefundTestUsers bool
}

var etherScalar = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

// Ether converts a uint64 Ether amount into a *big.Int amount in wei units, for allocating test balances.
func Ether(v uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(v), etherScalar)
}

// Setup computes the testing setup configurations from deployment configuration and optional allocation parameters.
func Setup(t require.TestingT, deployParams *DeployParams, alloc *AllocParams) *SetupData {
	deployConf := deployParams.DeployConfig
	l1Genesis, err := genesis2.BuildL1DeveloperGenesis(deployConf)
	require.NoError(t, err, "failed to create l1 genesis")
	if alloc.PrefundTestUsers {
		for _, addr := range deployParams.Addresses.All() {
			l1Genesis.Alloc[addr] = core.GenesisAccount{
				Balance: Ether(1e12),
			}
		}
	}
	for addr, val := range alloc.L1Alloc {
		l1Genesis.Alloc[addr] = val
	}

	l1Block := l1Genesis.ToBlock()

	l2Genesis, err := genesis2.BuildL2DeveloperGenesis(deployConf, l1Block, true)
	require.NoError(t, err, "failed to create l2 genesis")
	if alloc.PrefundTestUsers {
		for _, addr := range deployParams.Addresses.All() {
			l2Genesis.Alloc[addr] = core.GenesisAccount{
				Balance: Ether(1e12),
			}
		}
	}
	for addr, val := range alloc.L2Alloc {
		l2Genesis.Alloc[addr] = val
	}

	rollupCfg := &rollup.Config{
		Genesis: rollup.Genesis{
			L1: eth.BlockID{
				Hash:   l1Block.Hash(),
				Number: 0,
			},
			L2: eth.BlockID{
				Hash:   l2Genesis.ToBlock().Hash(),
				Number: 0,
			},
			L2Time:       uint64(deployConf.L1GenesisBlockTimestamp),
			SystemConfig: SystemConfigFromDeployConfig(deployConf),
		},
		BlockTime:              deployConf.L2BlockTime,
		MaxSequencerDrift:      deployConf.MaxSequencerDrift,
		SeqWindowSize:          deployConf.SequencerWindowSize,
		ChannelTimeout:         deployConf.ChannelTimeout,
		L1ChainID:              new(big.Int).SetUint64(deployConf.L1ChainID),
		L2ChainID:              new(big.Int).SetUint64(deployConf.L2ChainID),
		BatchInboxAddress:      deployConf.BatchInboxAddress,
		DepositContractAddress: predeploys.DevKromaPortalAddr,
		L1SystemConfigAddress:  predeploys.DevSystemConfigAddr,
	}

	deploymentsL1 := DeploymentsL1{
		L1CrossDomainMessengerProxy: predeploys.DevL1CrossDomainMessengerAddr,
		L1StandardBridgeProxy:       predeploys.DevL1StandardBridgeAddr,
		ValidatorPoolProxy:          predeploys.DevValidatorPoolAddr,
		L2OutputOracleProxy:         predeploys.DevL2OutputOracleAddr,
		ColosseumProxy:              predeploys.DevColosseumAddr,
		SecurityCouncilProxy:        predeploys.DevSecurityCouncilAddr,
		KromaPortalProxy:            predeploys.DevKromaPortalAddr,
		SystemConfigProxy:           predeploys.DevSystemConfigAddr,
		SecurityCouncilTokenProxy:   predeploys.DevSecurityCouncilTokenAddr,
		TimeLockProxy:               predeploys.DevTimeLockAddr,
		UpgradeGovernorProxy:        predeploys.DevUpgradeGovernorAddr,
	}

	return &SetupData{
		L1Cfg:         l1Genesis,
		L2Cfg:         l2Genesis,
		RollupCfg:     rollupCfg,
		DeploymentsL1: deploymentsL1,
	}
}

func SystemConfigFromDeployConfig(deployConfig *genesis2.DeployConfig) eth.SystemConfig {
	return eth.SystemConfig{
		BatcherAddr:           deployConfig.BatchSenderAddress,
		Overhead:              eth.Bytes32(common.BigToHash(new(big.Int).SetUint64(deployConfig.GasPriceOracleOverhead))),
		Scalar:                eth.Bytes32(common.BigToHash(new(big.Int).SetUint64(deployConfig.GasPriceOracleScalar))),
		GasLimit:              uint64(deployConfig.L2GenesisBlockGasLimit),
		ValidatorRewardScalar: eth.Bytes32(common.BigToHash(new(big.Int).SetUint64(deployConfig.ValidatorRewardScalar))),
	}
}
