package genesis

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum-optimism/optimism/op-bindings/hardhat"
	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-chain-ops/immutables"
	"github.com/ethereum-optimism/optimism/op-chain-ops/state"
	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
)

var (
	ErrInvalidDeployConfig     = errors.New("invalid deploy config")
	ErrInvalidImmutablesConfig = errors.New("invalid immutables config")
)

// DeployConfig represents the deployment configuration for Optimism
type DeployConfig struct {
	L1StartingBlockTag *MarshalableRPCBlockNumberOrHash `json:"l1StartingBlockTag"`
	L1ChainID          uint64                           `json:"l1ChainID"`
	L2ChainID          uint64                           `json:"l2ChainID"`
	L2BlockTime        uint64                           `json:"l2BlockTime"`

	FinalizationPeriodSeconds uint64         `json:"finalizationPeriodSeconds"`
	MaxSequencerDrift         uint64         `json:"maxSequencerDrift"`
	SequencerWindowSize       uint64         `json:"sequencerWindowSize"`
	ChannelTimeout            uint64         `json:"channelTimeout"`
	P2PSequencerAddress       common.Address `json:"p2pSequencerAddress"`
	BatchInboxAddress         common.Address `json:"batchInboxAddress"`
	BatchSenderAddress        common.Address `json:"batchSenderAddress"`

	ValidatorPoolTrustedValidator   common.Address `json:"validatorPoolTrustedValidator"`
	ValidatorPoolRequiredBondAmount *hexutil.Big   `json:"validatorPoolRequiredBondAmount"`
	ValidatorPoolMaxUnbond          uint64         `json:"validatorPoolMaxUnbond"`
	ValidatorPoolRoundDuration      uint64         `json:"validatorPoolRoundDuration"`

	L2OutputOracleSubmissionInterval uint64 `json:"l2OutputOracleSubmissionInterval"`
	L2OutputOracleStartingTimestamp  int    `json:"l2OutputOracleStartingTimestamp"`

	L1BlockTime                 uint64         `json:"l1BlockTime"`
	L1GenesisBlockTimestamp     hexutil.Uint64 `json:"l1GenesisBlockTimestamp"`
	L1GenesisBlockNonce         hexutil.Uint64 `json:"l1GenesisBlockNonce"`
	CliqueSignerAddress         common.Address `json:"cliqueSignerAddress"` // proof of stake genesis if left zeroed.
	L1GenesisBlockGasLimit      hexutil.Uint64 `json:"l1GenesisBlockGasLimit"`
	L1GenesisBlockDifficulty    *hexutil.Big   `json:"l1GenesisBlockDifficulty"`
	L1GenesisBlockMixHash       common.Hash    `json:"l1GenesisBlockMixHash"`
	L1GenesisBlockCoinbase      common.Address `json:"l1GenesisBlockCoinbase"`
	L1GenesisBlockNumber        hexutil.Uint64 `json:"l1GenesisBlockNumber"`
	L1GenesisBlockGasUsed       hexutil.Uint64 `json:"l1GenesisBlockGasUsed"`
	L1GenesisBlockParentHash    common.Hash    `json:"l1GenesisBlockParentHash"`
	L1GenesisBlockBaseFeePerGas *hexutil.Big   `json:"l1GenesisBlockBaseFeePerGas"`

	L2GenesisBlockNonce         hexutil.Uint64 `json:"l2GenesisBlockNonce"`
	L2GenesisBlockGasLimit      hexutil.Uint64 `json:"l2GenesisBlockGasLimit"`
	L2GenesisBlockDifficulty    *hexutil.Big   `json:"l2GenesisBlockDifficulty"`
	L2GenesisBlockMixHash       common.Hash    `json:"l2GenesisBlockMixHash"`
	L2GenesisBlockNumber        hexutil.Uint64 `json:"l2GenesisBlockNumber"`
	L2GenesisBlockGasUsed       hexutil.Uint64 `json:"l2GenesisBlockGasUsed"`
	L2GenesisBlockParentHash    common.Hash    `json:"l2GenesisBlockParentHash"`
	L2GenesisBlockBaseFeePerGas *hexutil.Big   `json:"l2GenesisBlockBaseFeePerGas"`

	ColosseumCreationPeriodSeconds uint64      `json:"colosseumCreationPeriodSeconds"`
	ColosseumBisectionTimeout      uint64      `json:"colosseumBisectionTimeout"`
	ColosseumProvingTimeout        uint64      `json:"colosseumProvingTimeout"`
	ColosseumSegmentsLengths       string      `json:"colosseumSegmentsLengths"`
	ColosseumDummyHash             common.Hash `json:"colosseumDummyHash"`
	ColosseumMaxTxs                uint64      `json:"colosseumMaxTxs"`

	SecurityCouncilOwners []common.Address `json:"securityCouncilOwners"`

	// The initial value of the voting delay(unit:block)
	GovernorVotingDelayBlocks uint64 `json:"governorVotingDelayBlocks"`
	// The initial value of the voting period(unit:block)
	GovernorVotingPeriodBlocks uint64 `json:"governorVotingPeriodBlocks"`
	// The initial value of the proposal threshold(unit:token)
	GovernorProposalThreshold uint64 `json:"governorProposalThreshold"`
	// The initial value of the votes quorum fraction(unit:percent)
	GovernorVotesQuorumFractionPercent uint64 `json:"governorVotesQuorumFractionPercent"`
	// The latency value of the proposal executing(unit:second)
	TimeLockMinDelaySeconds uint64 `json:"timeLockMinDelaySeconds"`
	// The initial value of the L2 voting period(unit:block)
	L2GovernorVotingPeriodBlocks uint64 `json:"l2GovernorVotingPeriodBlocks"`
	// The latency value of the L2 proposal executing(unit:second)
	L2TimeLockMinDelaySeconds uint64 `json:"l2TimeLockMinDelaySeconds"`

	ZKVerifierHashScalar *hexutil.Big `json:"zkVerifierHashScalar"`
	ZKVerifierM56Px      *hexutil.Big `json:"zkVerifierM56Px"`
	ZKVerifierM56Py      *hexutil.Big `json:"zkVerifierM56Py"`

	// Owner of the ProxyAdmin predeploy, but in test it means super admin of the system
	ProxyAdminOwner common.Address `json:"proxyAdminOwner"`
	// L1 recipient of fees accumulated in the ProtocolVault
	ProtocolVaultRecipient common.Address `json:"protocolVaultRecipient"`
	// L1 recipient of fees accumulated in the L1FeeVaultRecipient
	L1FeeVaultRecipient common.Address `json:"l1FeeVaultRecipient"`
	// L1StandardBridge proxy address on L1
	L1StandardBridgeProxy common.Address `json:"l1StandardBridgeProxy"`
	// L1CrossDomainMessenger proxy address on L1
	L1CrossDomainMessengerProxy common.Address `json:"l1CrossDomainMessengerProxy"`
	// L1ERC721Bridge proxy address on L1
	L1ERC721BridgeProxy common.Address `json:"l1ERC721BridgeProxy"`
	// SystemConfig proxy address on L1
	SystemConfigProxy common.Address `json:"systemConfigProxy"`
	// KromaPortal proxy address on L1
	KromaPortalProxy common.Address `json:"kromaPortalProxy"`
	// ValidatorPool proxy address on L1
	ValidatorPoolProxy common.Address `json:"validatorPoolProxy"`
	// The initial value of the gas overhead
	GasPriceOracleOverhead uint64 `json:"gasPriceOracleOverhead"`
	// The initial value of the gas scalar
	GasPriceOracleScalar uint64 `json:"gasPriceOracleScalar"`
	// The initial value of the validator reward scalar
	ValidatorRewardScalar uint64 `json:"validatorRewardScalar"`

	DeploymentWaitConfirmations int `json:"deploymentWaitConfirmations"`

	EIP1559Elasticity  uint64 `json:"eip1559Elasticity"`
	EIP1559Denominator uint64 `json:"eip1559Denominator"`

	FundDevAccounts bool `json:"fundDevAccounts"`
}

// Check will ensure that the config is sane and return an error when it is not
func (d *DeployConfig) Check() error {
	if d.L1StartingBlockTag == nil {
		return fmt.Errorf("%w: L2StartingBlockTag cannot be nil", ErrInvalidDeployConfig)
	}
	if d.L1ChainID == 0 {
		return fmt.Errorf("%w: L1ChainID cannot be 0", ErrInvalidDeployConfig)
	}
	if d.L2ChainID == 0 {
		return fmt.Errorf("%w: L2ChainID cannot be 0", ErrInvalidDeployConfig)
	}
	if d.L2BlockTime == 0 {
		return fmt.Errorf("%w: L2BlockTime cannot be 0", ErrInvalidDeployConfig)
	}
	if d.FinalizationPeriodSeconds == 0 {
		return fmt.Errorf("%w: FinalizationPeriodSeconds cannot be 0", ErrInvalidDeployConfig)
	}
	if d.MaxSequencerDrift == 0 {
		return fmt.Errorf("%w: MaxSequencerDrift cannot be 0", ErrInvalidDeployConfig)
	}
	if d.SequencerWindowSize == 0 {
		return fmt.Errorf("%w: SequencerWindowSize cannot be 0", ErrInvalidDeployConfig)
	}
	if d.ChannelTimeout == 0 {
		return fmt.Errorf("%w: ChannelTimeout cannot be 0", ErrInvalidDeployConfig)
	}
	if d.P2PSequencerAddress == (common.Address{}) {
		return fmt.Errorf("%w: P2PSequencerAddress cannot be address(0)", ErrInvalidDeployConfig)
	}
	if d.BatchInboxAddress == (common.Address{}) {
		return fmt.Errorf("%w: BatchInboxAddress cannot be address(0)", ErrInvalidDeployConfig)
	}
	if d.BatchSenderAddress == (common.Address{}) {
		return fmt.Errorf("%w: BatchSenderAddress cannot be address(0)", ErrInvalidDeployConfig)
	}
	if d.ValidatorPoolTrustedValidator == (common.Address{}) {
		return fmt.Errorf("%w: ValidatorPoolTrustedValidator cannot be address(0)", ErrInvalidDeployConfig)
	}
	if d.ValidatorPoolRequiredBondAmount == nil {
		return fmt.Errorf("%w: ValidatorPoolRequiredBondAmount cannot be nil", ErrInvalidDeployConfig)
	}
	if d.ValidatorPoolMaxUnbond == 0 {
		return fmt.Errorf("%w: ValidatorPoolMaxUnbond cannot be 0", ErrInvalidDeployConfig)
	}
	if d.ValidatorPoolRoundDuration == 0 {
		return fmt.Errorf("%w: ValidatorPoolRoundDuration cannot be 0", ErrInvalidDeployConfig)
	}
	if d.L2OutputOracleSubmissionInterval == 0 {
		return fmt.Errorf("%w: L2OutputOracleSubmissionInterval cannot be 0", ErrInvalidDeployConfig)
	}
	if d.L2OutputOracleSubmissionInterval*d.L2BlockTime != d.ValidatorPoolRoundDuration*2 {
		return fmt.Errorf("%w: double of ValidatorPoolRoundDuration must equal to L2OutputOracleSubmissionInterval", ErrInvalidDeployConfig)
	}
	if d.L2OutputOracleStartingTimestamp == 0 {
		log.Warn("L2OutputOracleStartingTimestamp is 0")
	}
	if d.ProxyAdminOwner == (common.Address{}) {
		return fmt.Errorf("%w: ProxyAdminOwner cannot be address(0)", ErrInvalidDeployConfig)
	}
	if d.ProtocolVaultRecipient == (common.Address{}) {
		return fmt.Errorf("%w: ProtocolVaultRecipient cannot be address(0)", ErrInvalidDeployConfig)
	}
	if d.L1FeeVaultRecipient == (common.Address{}) {
		return fmt.Errorf("%w: L1FeeVaultRecipient cannot be address(0)", ErrInvalidDeployConfig)
	}
	if d.GasPriceOracleOverhead == 0 {
		log.Warn("GasPriceOracleOverhead is 0")
	}
	if d.GasPriceOracleScalar == 0 {
		return fmt.Errorf("%w: GasPriceOracleScalar cannot be 0", ErrInvalidDeployConfig)
	}
	if d.ValidatorRewardScalar == 0 {
		log.Warn("ValidatorRewardScalar is 0")
	}
	if d.L1StandardBridgeProxy == (common.Address{}) {
		return fmt.Errorf("%w: L1StandardBridgeProxy cannot be address(0)", ErrInvalidDeployConfig)
	}
	if d.L1CrossDomainMessengerProxy == (common.Address{}) {
		return fmt.Errorf("%w: L1CrossDomainMessengerProxy cannot be address(0)", ErrInvalidDeployConfig)
	}
	if d.L1ERC721BridgeProxy == (common.Address{}) {
		return fmt.Errorf("%w: L1ERC721BridgeProxy cannot be address(0)", ErrInvalidDeployConfig)
	}
	if d.SystemConfigProxy == (common.Address{}) {
		return fmt.Errorf("%w: SystemConfigProxy cannot be address(0)", ErrInvalidDeployConfig)
	}
	if d.KromaPortalProxy == (common.Address{}) {
		return fmt.Errorf("%w: KromaPortalProxy cannot be address(0)", ErrInvalidDeployConfig)
	}
	if d.EIP1559Denominator == 0 {
		return fmt.Errorf("%w: EIP1559Denominator cannot be 0", ErrInvalidDeployConfig)
	}
	if d.EIP1559Elasticity == 0 {
		return fmt.Errorf("%w: EIP1559Elasticity cannot be 0", ErrInvalidDeployConfig)
	}
	if d.L2GenesisBlockGasLimit == 0 {
		return fmt.Errorf("%w: L2 genesis block gas limit cannot be 0", ErrInvalidDeployConfig)
	}
	// When the initial resource config is made to be configurable by the DeployConfig, ensure
	// that this check is updated to use the values from the DeployConfig instead of the defaults.
	if uint64(d.L2GenesisBlockGasLimit) < uint64(defaultResourceConfig.MaxResourceLimit+defaultResourceConfig.SystemTxMaxGas) {
		return fmt.Errorf("%w: L2 genesis block gas limit is too small", ErrInvalidDeployConfig)
	}
	if d.L2GenesisBlockBaseFeePerGas == nil {
		return fmt.Errorf("%w: L2 genesis block base fee per gas cannot be nil", ErrInvalidDeployConfig)
	}
	if d.ColosseumCreationPeriodSeconds == 0 {
		return fmt.Errorf("%w: ColosseumCreationPeriodSeconds cannot be 0", ErrInvalidDeployConfig)
	}
	if d.ColosseumBisectionTimeout == 0 {
		return fmt.Errorf("%w: ColosseumBisectionTimeout cannot be 0", ErrInvalidDeployConfig)
	}
	if d.ColosseumProvingTimeout == 0 {
		return fmt.Errorf("%w: ColosseumProvingTimeout cannot be 0", ErrInvalidDeployConfig)
	}
	if d.ColosseumDummyHash == (common.Hash{}) {
		return fmt.Errorf("%w: ColosseumDummyHash cannot be 0", ErrInvalidDeployConfig)
	}
	if d.ColosseumMaxTxs == 0 {
		return fmt.Errorf("%w: ColosseumMaxTxs cannot be 0", ErrInvalidDeployConfig)
	}
	if len(strings.Split(d.ColosseumSegmentsLengths, ","))%2 > 0 {
		return fmt.Errorf("%w: ColosseumSegmentsLengths length cannot be an odd number", ErrInvalidDeployConfig)
	}
	if d.GovernorVotingPeriodBlocks == 0 {
		return fmt.Errorf("%w: GovernorVotingPeriodBlocks cannot be 0", ErrInvalidDeployConfig)
	}
	if d.GovernorProposalThreshold == 0 {
		return fmt.Errorf("%w: GovernorProposalThreshold cannot be 0", ErrInvalidDeployConfig)
	}
	if d.GovernorVotesQuorumFractionPercent > 100 {
		return fmt.Errorf("%w: GovernorVotesQuorumFractionPercent cannot be greater than 100", ErrInvalidDeployConfig)
	}
	if d.ZKVerifierHashScalar == nil {
		return fmt.Errorf("%w: ZKVerifierHashScalar cannot be nil", ErrInvalidDeployConfig)
	}
	if d.ZKVerifierM56Px == nil {
		return fmt.Errorf("%w: ZKVerifierM56Px cannot be nil", ErrInvalidDeployConfig)
	}
	if d.ZKVerifierM56Py == nil {
		return fmt.Errorf("%w: ZKVerifierM56Py cannot be nil", ErrInvalidDeployConfig)
	}
	return nil
}

// GetDeployedAddresses will get the deployed addresses of deployed L1 contracts
// required for the L2 genesis creation.
func (d *DeployConfig) GetDeployedAddresses(hh *hardhat.Hardhat) error {
	if d.L1StandardBridgeProxy == (common.Address{}) {
		l1StandardBridgeProxyDeployment, err := hh.GetDeployment("L1StandardBridgeProxy")
		if err != nil {
			return err
		}
		d.L1StandardBridgeProxy = l1StandardBridgeProxyDeployment.Address
	}

	if d.L1CrossDomainMessengerProxy == (common.Address{}) {
		l1CrossDomainMessengerProxyDeployment, err := hh.GetDeployment("L1CrossDomainMessengerProxy")
		if err != nil {
			return err
		}
		d.L1CrossDomainMessengerProxy = l1CrossDomainMessengerProxyDeployment.Address
	}

	if d.L1ERC721BridgeProxy == (common.Address{}) {
		// There is no legacy deployment of this contract
		l1ERC721BridgeProxyDeployment, err := hh.GetDeployment("L1ERC721BridgeProxy")
		if err != nil {
			return err
		}
		d.L1ERC721BridgeProxy = l1ERC721BridgeProxyDeployment.Address
	}

	if d.SystemConfigProxy == (common.Address{}) {
		systemConfigProxyDeployment, err := hh.GetDeployment("SystemConfigProxy")
		if err != nil {
			return err
		}
		d.SystemConfigProxy = systemConfigProxyDeployment.Address
	}

	if d.KromaPortalProxy == (common.Address{}) {
		kromaPortalProxyDeployment, err := hh.GetDeployment("KromaPortalProxy")
		if err != nil {
			return err
		}
		d.KromaPortalProxy = kromaPortalProxyDeployment.Address
	}

	if d.ValidatorPoolProxy == (common.Address{}) {
		validatorPoolProxyDeployment, err := hh.GetDeployment("ValidatorPoolProxy")
		if err != nil {
			return err
		}
		d.ValidatorPoolProxy = validatorPoolProxyDeployment.Address
	}

	return nil
}

// InitDeveloperDeployedAddresses will set the dev addresses on the DeployConfig
func (d *DeployConfig) InitDeveloperDeployedAddresses() error {
	d.L1StandardBridgeProxy = predeploys.DevL1StandardBridgeAddr
	d.L1CrossDomainMessengerProxy = predeploys.DevL1CrossDomainMessengerAddr
	d.L1ERC721BridgeProxy = predeploys.DevL1ERC721BridgeAddr
	d.KromaPortalProxy = predeploys.DevKromaPortalAddr
	d.SystemConfigProxy = predeploys.DevSystemConfigAddr
	d.ValidatorPoolProxy = predeploys.DevValidatorPoolAddr
	return nil
}

// RollupConfig converts a DeployConfig to a rollup.Config
func (d *DeployConfig) RollupConfig(l1StartBlock *types.Block, l2GenesisBlockHash common.Hash, l2GenesisBlockNumber uint64) (*rollup.Config, error) {
	if d.KromaPortalProxy == (common.Address{}) {
		return nil, errors.New("KromaPortalProxy cannot be address(0)")
	}
	if d.SystemConfigProxy == (common.Address{}) {
		return nil, errors.New("SystemConfigProxy cannot be address(0)")
	}

	return &rollup.Config{
		Genesis: rollup.Genesis{
			L1: eth.BlockID{
				Hash:   l1StartBlock.Hash(),
				Number: l1StartBlock.NumberU64(),
			},
			L2: eth.BlockID{
				Hash:   l2GenesisBlockHash,
				Number: l2GenesisBlockNumber,
			},
			L2Time: l1StartBlock.Time(),
			SystemConfig: eth.SystemConfig{
				BatcherAddr:           d.BatchSenderAddress,
				Overhead:              eth.Bytes32(common.BigToHash(new(big.Int).SetUint64(d.GasPriceOracleOverhead))),
				Scalar:                eth.Bytes32(common.BigToHash(new(big.Int).SetUint64(d.GasPriceOracleScalar))),
				GasLimit:              uint64(d.L2GenesisBlockGasLimit),
				ValidatorRewardScalar: eth.Bytes32(common.BigToHash(new(big.Int).SetUint64(d.ValidatorRewardScalar))),
			},
		},
		BlockTime:              d.L2BlockTime,
		MaxSequencerDrift:      d.MaxSequencerDrift,
		SeqWindowSize:          d.SequencerWindowSize,
		ChannelTimeout:         d.ChannelTimeout,
		L1ChainID:              new(big.Int).SetUint64(d.L1ChainID),
		L2ChainID:              new(big.Int).SetUint64(d.L2ChainID),
		BatchInboxAddress:      d.BatchInboxAddress,
		DepositContractAddress: d.KromaPortalProxy,
		L1SystemConfigAddress:  d.SystemConfigProxy,
	}, nil
}

// NewDeployConfig reads a config file given a path on the filesystem.
func NewDeployConfig(path string) (*DeployConfig, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("deploy config at %s not found: %w", path, err)
	}

	var config DeployConfig
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf("cannot unmarshal deploy config: %w", err)
	}

	return &config, nil
}

// NewDeployConfigWithNetwork takes a path to a deploy config directory
// and the network name. The config file in the deploy config directory
// must match the network name and be a JSON file.
func NewDeployConfigWithNetwork(network, path string) (*DeployConfig, error) {
	deployConfig := filepath.Join(path, network+".json")
	return NewDeployConfig(deployConfig)
}

// NewL2ImmutableConfig will create an ImmutableConfig given an instance of a
// DeployConfig and a block.
func NewL2ImmutableConfig(config *DeployConfig, block *types.Block) (immutables.ImmutableConfig, error) {
	immutable := make(immutables.ImmutableConfig)

	if config.L1StandardBridgeProxy == (common.Address{}) {
		return immutable, fmt.Errorf("L1StandardBridgeProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1CrossDomainMessengerProxy == (common.Address{}) {
		return immutable, fmt.Errorf("L1CrossDomainMessengerProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1ERC721BridgeProxy == (common.Address{}) {
		return immutable, fmt.Errorf("L1ERC721BridgeProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.ProtocolVaultRecipient == (common.Address{}) {
		return immutable, fmt.Errorf("ProtocolVaultRecipient cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1FeeVaultRecipient == (common.Address{}) {
		return immutable, fmt.Errorf("L1FeeVaultRecipient cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}

	immutable["L2StandardBridge"] = immutables.ImmutableValues{
		"otherBridge": config.L1StandardBridgeProxy,
	}
	immutable["L2CrossDomainMessenger"] = immutables.ImmutableValues{
		"otherMessenger": config.L1CrossDomainMessengerProxy,
	}
	immutable["L2ERC721Bridge"] = immutables.ImmutableValues{
		"messenger":   predeploys.L2CrossDomainMessengerAddr,
		"otherBridge": config.L1ERC721BridgeProxy,
	}
	immutable["KromaMintableERC721Factory"] = immutables.ImmutableValues{
		"bridge":        predeploys.L2ERC721BridgeAddr,
		"remoteChainId": new(big.Int).SetUint64(config.L1ChainID),
	}
	rewardDivider := config.FinalizationPeriodSeconds / (config.L2OutputOracleSubmissionInterval * config.L2BlockTime)
	immutable["ValidatorRewardVault"] = immutables.ImmutableValues{
		"validatorPoolAddress": config.ValidatorPoolProxy,
		"rewardDivider":        new(big.Int).SetUint64(rewardDivider),
	}
	immutable["L1FeeVault"] = immutables.ImmutableValues{
		"recipient": config.L1FeeVaultRecipient,
	}
	immutable["ProtocolVault"] = immutables.ImmutableValues{
		"recipient": config.ProtocolVaultRecipient,
	}

	return immutable, nil
}

// NewL2StorageConfig will create a StorageConfig given an instance of a
// Hardhat and a DeployConfig.
func NewL2StorageConfig(config *DeployConfig, block *types.Block) (state.StorageConfig, error) {
	storage := make(state.StorageConfig)

	if block.Number() == nil {
		return storage, errors.New("block number not set")
	}
	if block.BaseFee() == nil {
		return storage, errors.New("block base fee not set")
	}

	storage["L2ToL1MessagePasser"] = state.StorageValues{
		"msgNonce": 0,
	}
	storage["L2CrossDomainMessenger"] = state.StorageValues{
		"_initialized":     1,
		"_initializing":    false,
		"xDomainMsgSender": "0x000000000000000000000000000000000000dEaD",
		"msgNonce":         0,
	}
	storage["L1Block"] = state.StorageValues{
		"number":                block.Number(),
		"timestamp":             block.Time(),
		"basefee":               block.BaseFee(),
		"hash":                  block.Hash(),
		"sequenceNumber":        0,
		"batcherHash":           config.BatchSenderAddress.Hash(),
		"l1FeeOverhead":         config.GasPriceOracleOverhead,
		"l1FeeScalar":           config.GasPriceOracleScalar,
		"validatorRewardScalar": config.ValidatorRewardScalar,
	}
	storage["WETH9"] = state.StorageValues{
		"name":     "Wrapped Ether",
		"symbol":   "WETH",
		"decimals": 18,
	}
	storage["ProxyAdmin"] = state.StorageValues{
		"_owner": config.ProxyAdminOwner,
	}
	return storage, nil
}

type MarshalableRPCBlockNumberOrHash rpc.BlockNumberOrHash

func (m *MarshalableRPCBlockNumberOrHash) MarshalJSON() ([]byte, error) {
	r := rpc.BlockNumberOrHash(*m)
	if hash, ok := r.Hash(); ok {
		return json.Marshal(hash)
	}
	if num, ok := r.Number(); ok {
		// never errors
		text, _ := num.MarshalText()
		return json.Marshal(string(text))
	}
	return json.Marshal(nil)
}

func (m *MarshalableRPCBlockNumberOrHash) UnmarshalJSON(b []byte) error {
	var r rpc.BlockNumberOrHash
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	asMarshalable := MarshalableRPCBlockNumberOrHash(r)
	*m = asMarshalable
	return nil
}
