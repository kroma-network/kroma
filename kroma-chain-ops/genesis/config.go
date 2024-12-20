package genesis

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/ethereum-optimism/optimism/op-chain-ops/state"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gstate "github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/kroma-network/kroma/kroma-bindings/predeploys"
	"github.com/kroma-network/kroma/kroma-chain-ops/immutables"
)

var (
	ErrInvalidDeployConfig     = errors.New("invalid deploy config")
	ErrInvalidImmutablesConfig = errors.New("invalid immutables config")
)

// DeployConfig represents the deployment configuration for an OP Stack chain.
// It is used to deploy the L1 contracts as well as create the L2 genesis state.
type DeployConfig struct {
	// L1StartingBlockTag is used to fill in the storage of the L1Block info predeploy. The rollup
	// config script uses this to fill the L1 genesis info for the rollup. The Output oracle deploy
	// script may use it if the L2 starting timestamp is nil, assuming the L2 genesis is set up
	// with this.
	L1StartingBlockTag *MarshalableRPCBlockNumberOrHash `json:"l1StartingBlockTag"`
	// L1ChainID is the chain ID of the L1 chain.
	L1ChainID uint64 `json:"l1ChainID"`
	// L2ChainID is the chain ID of the L2 chain.
	L2ChainID uint64 `json:"l2ChainID"`
	// L2BlockTime is the number of seconds between each L2 block.
	L2BlockTime uint64 `json:"l2BlockTime"`
	// FinalizationPeriodSeconds represents the number of seconds before an output is considered
	// finalized. This impacts the amount of time that withdrawals take to finalize and is
	// generally set to 1 week.
	FinalizationPeriodSeconds uint64 `json:"finalizationPeriodSeconds"`
	// MaxSequencerDrift is the number of seconds after the L1 timestamp of the end of the
	// sequencing window that batches must be included, otherwise L2 blocks including
	// deposits are force included.
	MaxSequencerDrift uint64 `json:"maxSequencerDrift"`
	// SequencerWindowSize is the number of L1 blocks per sequencing window.
	SequencerWindowSize uint64 `json:"sequencerWindowSize"`
	// ChannelTimeout is the number of L1 blocks that a frame stays valid when included in L1.
	ChannelTimeout uint64 `json:"channelTimeout"`
	// P2PSequencerAddress is the address of the key the sequencer uses to sign blocks on the P2P layer.
	P2PSequencerAddress common.Address `json:"p2pSequencerAddress"`
	// BatchInboxAddress is the L1 account that batches are sent to.
	BatchInboxAddress common.Address `json:"batchInboxAddress"`
	// BatchSenderAddress represents the initial sequencer account that authorizes batches.
	// Transactions sent from this account to the batch inbox address are considered valid.
	BatchSenderAddress common.Address `json:"batchSenderAddress"`
	// L2OutputOracleSubmissionInterval is the number of L2 blocks between outputs that are submitted
	// to the L2OutputOracle contract located on L1.
	L2OutputOracleSubmissionInterval uint64 `json:"l2OutputOracleSubmissionInterval"`
	// L2OutputOracleStartingTimestamp is the starting timestamp for the L2OutputOracle.
	// MUST be the same as the timestamp of the L2OO start block.
	L2OutputOracleStartingTimestamp int `json:"l2OutputOracleStartingTimestamp"`

	/* [Kroma: START]
	// L2OutputOracleStartingBlockNumber is the starting block number for the L2OutputOracle.
	// Must be greater than or equal to the first Bedrock block. The first L2 output will correspond
	// to this value plus the submission interval.
	L2OutputOracleStartingBlockNumber uint64 `json:"l2OutputOracleStartingBlockNumber"`
	// L2OutputOracleProposer is the address of the account that proposes L2 outputs.
	L2OutputOracleProposer common.Address `json:"l2OutputOracleProposer"`
	// L2OutputOracleChallenger is the address of the account that challenges L2 outputs.
	L2OutputOracleChallenger common.Address `json:"l2OutputOracleChallenger"`
	[Kroma: END] */

	// CliqueSignerAddress represents the signer address for the clique consensus engine.
	// It is used in the multi-process devnet to sign blocks.
	CliqueSignerAddress common.Address `json:"cliqueSignerAddress"`
	// L1UseClique represents whether or not to use the clique consensus engine.
	L1UseClique bool `json:"l1UseClique"`

	L1BlockTime                 uint64         `json:"l1BlockTime"`
	L1GenesisBlockTimestamp     hexutil.Uint64 `json:"l1GenesisBlockTimestamp"`
	L1GenesisBlockNonce         hexutil.Uint64 `json:"l1GenesisBlockNonce"`
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

	// L2GenesisRegolithTimeOffset is the number of seconds after genesis block that Regolith hard fork activates.
	// Set it to 0 to activate at genesis. Nil to disable Regolith.
	L2GenesisRegolithTimeOffset *hexutil.Uint64 `json:"l2GenesisRegolithTimeOffset,omitempty"`
	// L2GenesisCanyonTimeOffset is the number of seconds after genesis block that Canyon hard fork activates.
	// Set it to 0 to activate at genesis. Nil to disable Canyon.
	L2GenesisCanyonTimeOffset *hexutil.Uint64 `json:"l2GenesisCanyonTimeOffset,omitempty"`
	// L2GenesisDeltaTimeOffset is the number of seconds after genesis block that Delta hard fork activates.
	// Set it to 0 to activate at genesis. Nil to disable Delta.
	L2GenesisDeltaTimeOffset *hexutil.Uint64 `json:"l2GenesisDeltaTimeOffset,omitempty"`
	// L2GenesisEcotoneTimeOffset is the number of seconds after genesis block that Ecotone hard fork activates.
	// Set it to 0 to activate at genesis. Nil to disable Ecotone.
	L2GenesisEcotoneTimeOffset *hexutil.Uint64 `json:"l2GenesisEcotoneTimeOffset,omitempty"`
	// L2GenesisKromaMPTTimeOffset is the number of seconds after genesis block that Kroma MPT hard fork activates.
	// Set it to 0 to activate at genesis. Nil to disable Kroma MPT.
	L2GenesisKromaMPTTimeOffset *hexutil.Uint64 `json:"l2GenesisKromaMPTTimeOffset,omitempty"`
	// L2GenesisFjordTimeOffset is the number of seconds after genesis block that Fjord hard fork activates.
	// Set it to 0 to activate at genesis. Nil to disable Fjord.
	L2GenesisFjordTimeOffset *hexutil.Uint64 `json:"l2GenesisFjordTimeOffset,omitempty"`
	// L2GenesisInteropTimeOffset is the number of seconds after genesis block that the Interop hard fork activates.
	// Set it to 0 to activate at genesis. Nil to disable Interop.
	L2GenesisInteropTimeOffset *hexutil.Uint64 `json:"l2GenesisInteropTimeOffset,omitempty"`
	// L2GenesisBlockExtraData is configurable extradata. Will default to []byte("BEDROCK") if left unspecified.
	L2GenesisBlockExtraData []byte `json:"l2GenesisBlockExtraData"`
	// ProxyAdminOwner represents the owner of the ProxyAdmin predeploy on L2.
	ProxyAdminOwner common.Address `json:"proxyAdminOwner"`
	// L1 recipient of fees accumulated in the ProtocolVault
	ProtocolVaultRecipient common.Address `json:"protocolVaultRecipient"`
	// L1 recipient of fees accumulated in the L1FeeVaultRecipient
	L1FeeVaultRecipient common.Address `json:"l1FeeVaultRecipient"`
	// L1StandardBridgeProxy represents the address of the L1StandardBridgeProxy on L1 and is used
	// as part of building the L2 genesis state.
	L1StandardBridgeProxy common.Address `json:"l1StandardBridgeProxy"`
	// L1CrossDomainMessengerProxy represents the address of the L1CrossDomainMessengerProxy on L1 and is used
	// as part of building the L2 genesis state.
	L1CrossDomainMessengerProxy common.Address `json:"l1CrossDomainMessengerProxy"`
	// L1ERC721BridgeProxy represents the address of the L1ERC721Bridge on L1 and is used
	// as part of building the L2 genesis state.
	L1ERC721BridgeProxy common.Address `json:"l1ERC721BridgeProxy"`
	// SystemConfigProxy represents the address of the SystemConfigProxy on L1 and is used
	// as part of the derivation pipeline.
	SystemConfigProxy common.Address `json:"systemConfigProxy"`
	// KromaPortalProxy represents the address of the KromaPortalProxy on L1 and is used
	// as part of the derivation pipeline.
	KromaPortalProxy common.Address `json:"kromaPortalProxy"`
	// GasPriceOracleOverhead represents the initial value of the gas overhead in the GasPriceOracle predeploy.
	GasPriceOracleOverhead uint64 `json:"gasPriceOracleOverhead"`
	// GasPriceOracleScalar represents the initial value of the gas scalar in the GasPriceOracle predeploy.
	GasPriceOracleScalar uint64 `json:"gasPriceOracleScalar"`
	// EnableGovernance configures whether or not include governance token predeploy.
	EnableGovernance bool `json:"enableGovernance"`
	/* [Kroma: START]
	// GovernanceTokenSymbol represents the  ERC20 symbol of the GovernanceToken.
	GovernanceTokenSymbol string `json:"governanceTokenSymbol"`
	// GovernanceTokenName represents the ERC20 name of the GovernanceToken
	GovernanceTokenName string `json:"governanceTokenName"`
	// GovernanceTokenOwner represents the owner of the GovernanceToken. Has the ability
	// to mint and burn tokens.
	GovernanceTokenOwner common.Address `json:"governanceTokenOwner"`
	[Kroma: END] */
	// [Kroma: START]
	// GovernanceTokenNotUseCreate2 is used to determine whether not to use CREATE2 to deploy GovernanceTokenProxy.
	GovernanceTokenNotUseCreate2 bool `json:"governanceTokenNotUseCreate2,omitempty"`
	// GovernanceTokenProxySalt is used to determine GovernanceTokenProxy address on L1 and L2.
	GovernanceTokenProxySalt common.Hash `json:"governanceTokenProxySalt"`
	// MintManagerOwner represents the owner of the MintManager on L1 and L2. Has the ability to mint initially.
	MintManagerOwner common.Address `json:"mintManagerOwner"`
	// L1MintManagerRecipients is an array of recipient addresses to receive the minted governance tokens on L1.
	L1MintManagerRecipients []common.Address `json:"l1MintManagerRecipients"`
	// L1MintManagerShares is an array of each recipient's share of total minted tokens on L1.
	L1MintManagerShares []uint64 `json:"l1MintManagerShares"`
	// L2MintManagerRecipients is an array of recipient addresses to receive the minted governance tokens on L2.
	L2MintManagerRecipients []common.Address `json:"l2MintManagerRecipients"`
	// L2MintManagerShares is an array of each recipient's share of total minted tokens on L2.
	L2MintManagerShares []uint64 `json:"l2MintManagerShares"`
	// [Kroma: END]
	// DeploymentWaitConfirmations is the number of confirmations to wait during
	// deployment. This is DEPRECATED and should be removed in a future PR.
	DeploymentWaitConfirmations int `json:"deploymentWaitConfirmations"`
	// EIP1559Elasticity is the elasticity of the EIP1559 fee market.
	EIP1559Elasticity uint64 `json:"eip1559Elasticity"`
	// EIP1559Denominator is the denominator of EIP1559 base fee market.
	EIP1559Denominator uint64 `json:"eip1559Denominator"`
	// EIP1559DenominatorCanyon is the denominator of EIP1559 base fee market when Canyon is active.
	EIP1559DenominatorCanyon uint64 `json:"eip1559DenominatorCanyon"`
	// SystemConfigStartBlock represents the block at which the op-node should start syncing
	// from. It is an override to set this value on legacy networks where it is not set by
	// default. It can be removed once all networks have this value set in their storage.
	SystemConfigStartBlock uint64 `json:"systemConfigStartBlock"`

	/* [Kroma: START]
	// FaultGameAbsolutePrestate is the absolute prestate of Cannon. This is computed
	// by generating a proof from the 0th -> 1st instruction and grabbing the prestate from
	// the output JSON. All honest challengers should agree on the setup state of the program.
	FaultGameAbsolutePrestate common.Hash `json:"faultGameAbsolutePrestate"`
	// FaultGameMaxDepth is the maximum depth of the position tree within the fault dispute game.
	// `2^{FaultGameMaxDepth}` is how many instructions the execution trace bisection game
	// supports. Ideally, this should be conservatively set so that there is always enough
	// room for a full Cannon trace.
	FaultGameMaxDepth uint64 `json:"faultGameMaxDepth"`
	// FaultGameMaxDuration is the maximum amount of time (in seconds) that the fault dispute
	// game can run for before it is ready to be resolved. Each side receives half of this value
	// on their chess clock at the inception of the dispute.
	FaultGameMaxDuration uint64 `json:"faultGameMaxDuration"`
	// FaultGameGenesisBlock is the block number for genesis.
	FaultGameGenesisBlock uint64 `json:"faultGameGenesisBlock"`
	// FaultGameGenesisOutputRoot is the output root for the genesis block.
	FaultGameGenesisOutputRoot common.Hash `json:"faultGameGenesisOutputRoot"`
	// FaultGameSplitDepth is the depth at which the fault dispute game splits from output roots to execution trace claims.
	FaultGameSplitDepth uint64 `json:"faultGameSplitDepth"`
	// FaultGameWithdrawalDelay is the number of seconds that users must wait before withdrawing ETH from a fault game.
	FaultGameWithdrawalDelay uint64 `json:"faultGameWithdrawalDelay"`
	// PreimageOracleMinProposalSize is the minimum number of bytes that a large preimage oracle proposal can be.
	PreimageOracleMinProposalSize uint64 `json:"preimageOracleMinProposalSize"`
	// PreimageOracleChallengePeriod is the number of seconds that challengers have to challenge a large preimage proposal.
	PreimageOracleChallengePeriod uint64 `json:"preimageOracleChallengePeriod"`
	[Kroma: END] */

	// FundDevAccounts configures whether or not to fund the dev accounts. Should only be used
	// during devnet deployments.
	FundDevAccounts bool `json:"fundDevAccounts"`

	/* [Kroma: START]
	// RequiredProtocolVersion indicates the protocol version that
	// nodes are required to adopt, to stay in sync with the network.
	RequiredProtocolVersion params.ProtocolVersion `json:"requiredProtocolVersion"`
	// RequiredProtocolVersion indicates the protocol version that
	// nodes are recommended to adopt, to stay in sync with the network.
	RecommendedProtocolVersion params.ProtocolVersion `json:"recommendedProtocolVersion"`
	// ProofMaturityDelaySeconds is the number of seconds that a proof must be
	// mature before it can be used to finalize a withdrawal.
	ProofMaturityDelaySeconds uint64 `json:"proofMaturityDelaySeconds"`
	// DisputeGameFinalityDelaySeconds is an additional number of seconds a
	// dispute game must wait before it can be used to finalize a withdrawal.
	DisputeGameFinalityDelaySeconds uint64 `json:"disputeGameFinalityDelaySeconds"`
	// RespectedGameType is the dispute game type that the OptimismPortal
	// contract will respect for finalizing withdrawals.
	RespectedGameType uint32 `json:"respectedGameType"`
	// UseFaultProofs is a flag that indicates if the system is using fault
	// proofs instead of the older output oracle mechanism.
	UseFaultProofs bool `json:"useFaultProofs"`
	[Kroma: END] */

	// UsePlasma is a flag that indicates if the system is using op-plasma
	UsePlasma bool `json:"usePlasma,omitempty"`
	// DAChallengeWindow represents the block interval during which the availability of a data commitment can be challenged.
	DAChallengeWindow uint64 `json:"daChallengeWindow,omitempty"`
	// DAResolveWindow represents the block interval during which a data availability challenge can be resolved.
	DAResolveWindow uint64 `json:"daResolveWindow,omitempty"`
	// DABondSize represents the required bond size to initiate a data availability challenge.
	DABondSize uint64 `json:"daBondSize,omitempty"`
	// DAResolverRefundPercentage represents the percentage of the resolving cost to be refunded to the resolver
	// such as 100 means 100% refund.
	DAResolverRefundPercentage uint64 `json:"daResolverRefundPercentage,omitempty"`
	// DAChallengeProxy represents the L1 address of the DataAvailabilityChallenge contract.
	DAChallengeProxy common.Address `json:"daChallengeProxy,omitempty"`
	// When Cancun activates. Relative to L1 genesis.
	L1CancunTimeOffset *hexutil.Uint64 `json:"l1CancunTimeOffset,omitempty"`

	// [Kroma: START]
	// ValidatorPool proxy address on L1
	ValidatorPoolProxy common.Address `json:"validatorPoolProxy"`
	// The initial value of the validator reward scalar
	ValidatorRewardScalar uint64 `json:"validatorRewardScalar"`

	ValidatorPoolTrustedValidator   common.Address `json:"validatorPoolTrustedValidator"`
	ValidatorPoolRequiredBondAmount *hexutil.Big   `json:"validatorPoolRequiredBondAmount"`
	ValidatorPoolMaxUnbond          uint64         `json:"validatorPoolMaxUnbond"`
	ValidatorPoolRoundDuration      uint64         `json:"validatorPoolRoundDuration"`
	// ValidatorPoolTerminateOutputIndex is the output index where ValidatorPool is terminated after
	// in hex value.
	ValidatorPoolTerminateOutputIndex *hexutil.Big `json:"validatorPoolTerminateOutputIndex"`

	// ValidatorManagerTrustedValidator represents the address of the trusted validator.
	ValidatorManagerTrustedValidator common.Address `json:"validatorManagerTrustedValidator"`
	// ValidatorManagerMinRegisterAmount is the amount of the minimum register amount.
	ValidatorManagerMinRegisterAmount *hexutil.Big `json:"validatorManagerMinRegisterAmount"`
	// ValidatorManagerMinActivateAmount is the amount of the minimum activation amount.
	ValidatorManagerMinActivateAmount *hexutil.Big `json:"validatorManagerMinActivateAmount"`
	// ValidatorManagerMptFirstOutputIndex is the first output index after the MPT transition.
	// Only TrustedValidator is allowed to submit output. Challenges for this outputIndex are also restricted.
	ValidatorManagerMptFirstOutputIndex *hexutil.Big `json:"validatorManagerMptFirstOutputIndex"`
	// ValidatorManagerCommissionChangeDelaySeconds is the delay to finalize the commission rate change in seconds.
	ValidatorManagerCommissionChangeDelaySeconds uint64 `json:"validatorManagerCommissionChangeDelaySeconds"`
	// ValidatorManagerRoundDurationSeconds is the duration of one submission round in seconds.
	ValidatorManagerRoundDurationSeconds uint64 `json:"validatorManagerRoundDurationSeconds"`
	// ValidatorManagerSoftJailPeriodSeconds is the duration of jail period in seconds in output non-submissions penalty.
	ValidatorManagerSoftJailPeriodSeconds uint64 `json:"validatorManagerSoftJailPeriodSeconds"`
	// ValidatorManagerHardJailPeriodSeconds is the duration of jail period in seconds in slashing penalty.
	ValidatorManagerHardJailPeriodSeconds uint64 `json:"validatorManagerHardJailPeriodSeconds"`
	// ValidatorManagerJailThreshold is the threshold of output non-submission to be jailed.
	ValidatorManagerJailThreshold uint64 `json:"validatorManagerJailThreshold"`
	// ValidatorManagerMaxFinalizations is the max number of output finalizations when distributing
	// reward.
	ValidatorManagerMaxFinalizations uint64 `json:"validatorManagerMaxFinalizations"`
	// ValidatorManagerBaseReward is the amount of the base reward in hex value.
	ValidatorManagerBaseReward *hexutil.Big `json:"validatorManagerBaseReward"`

	// AssetManagerKgh represents the address of the KGH NFT contract.
	AssetManagerKgh common.Address `json:"assetManagerKgh"`
	// AssetManagerVault represents the address of the validator reward vault.
	AssetManagerVault common.Address `json:"assetManagerVault"`
	// AssetManagerMinDelegationPeriod is the duration of minimum delegation period in seconds.
	AssetManagerMinDelegationPeriod uint64 `json:"assetManagerMinDelegationPeriod"`
	// AssetManagerBondAmount is the bond amount.
	AssetManagerBondAmount *hexutil.Big `json:"assetManagerBondAmount"`

	ColosseumCreationPeriodSeconds uint64      `json:"colosseumCreationPeriodSeconds"`
	ColosseumBisectionTimeout      uint64      `json:"colosseumBisectionTimeout"`
	ColosseumProvingTimeout        uint64      `json:"colosseumProvingTimeout"`
	ColosseumSegmentsLengths       string      `json:"colosseumSegmentsLengths"`
	ColosseumDummyHash             common.Hash `json:"colosseumDummyHash"`
	ColosseumMaxTxs                uint64      `json:"colosseumMaxTxs"`

	// Owner of the SecurityCouncil
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

	// ZKProofVerifierSP1Verifier is the address of the SP1VerifierGateway contract.
	ZKProofVerifierSP1Verifier common.Address `json:"zkProofVerifierSP1Verifier"`
	// ZKProofVerifierVKey is the verification key for the zkVM program.
	ZKProofVerifierVKey common.Hash `json:"zkProofVerifierVKey"`
	// [Kroma: END]
}

// Copy will deeply copy the DeployConfig. This does a JSON roundtrip to copy
// which makes it easier to maintain, we do not need efficiency in this case.
func (d *DeployConfig) Copy() *DeployConfig {
	raw, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	cpy := DeployConfig{}
	if err = json.Unmarshal(raw, &cpy); err != nil {
		panic(err)
	}
	return &cpy
}

// Check will ensure that the config is sane and return an error when it is not
func (d *DeployConfig) Check() error {
	if d.L1StartingBlockTag == nil {
		return fmt.Errorf("%w: L1StartingBlockTag cannot be nil", ErrInvalidDeployConfig)
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
	if d.L2OutputOracleSubmissionInterval == 0 {
		return fmt.Errorf("%w: L2OutputOracleSubmissionInterval cannot be 0", ErrInvalidDeployConfig)
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
	if d.EIP1559Denominator == 0 {
		return fmt.Errorf("%w: EIP1559Denominator cannot be 0", ErrInvalidDeployConfig)
	}
	if d.L2GenesisCanyonTimeOffset != nil && d.EIP1559DenominatorCanyon == 0 {
		return fmt.Errorf("%w: EIP1559DenominatorCanyon cannot be 0 if Canyon is activated", ErrInvalidDeployConfig)
	}
	if d.EIP1559Elasticity == 0 {
		return fmt.Errorf("%w: EIP1559Elasticity cannot be 0", ErrInvalidDeployConfig)
	}
	if d.L2GenesisBlockGasLimit == 0 {
		return fmt.Errorf("%w: L2 genesis block gas limit cannot be 0", ErrInvalidDeployConfig)
	}
	// When the initial resource config is made to be configurable by the DeployConfig, ensure
	// that this check is updated to use the values from the DeployConfig instead of the defaults.
	if uint64(d.L2GenesisBlockGasLimit) < uint64(DefaultResourceConfig.MaxResourceLimit+DefaultResourceConfig.SystemTxMaxGas) {
		return fmt.Errorf("%w: L2 genesis block gas limit is too small", ErrInvalidDeployConfig)
	}
	if d.L2GenesisBlockBaseFeePerGas == nil {
		return fmt.Errorf("%w: L2 genesis block base fee per gas cannot be nil", ErrInvalidDeployConfig)
	}
	if d.EnableGovernance {
		/* [Kroma: START]
		if d.GovernanceTokenName == "" {
			return fmt.Errorf("%w: GovernanceToken.name cannot be empty", ErrInvalidDeployConfig)
		}
		if d.GovernanceTokenSymbol == "" {
			return fmt.Errorf("%w: GovernanceToken.symbol cannot be empty", ErrInvalidDeployConfig)
		}
		if d.GovernanceTokenOwner == (common.Address{}) {
			return fmt.Errorf("%w: GovernanceToken owner cannot be address(0)", ErrInvalidDeployConfig)
		}
		[Kroma: END] */
		// [Kroma: START]
		if d.GovernanceTokenProxySalt == (common.Hash{}) {
			return fmt.Errorf("%w: GovernanceTokenProxySalt cannot be empty hash", ErrInvalidDeployConfig)
		}
		if d.MintManagerOwner == (common.Address{}) {
			return fmt.Errorf("%w: MintManagerOwner cannot be address(0)", ErrInvalidDeployConfig)
		}
		if len(d.L1MintManagerRecipients) == 0 {
			return fmt.Errorf("%w: L1MintManagerRecipients array cannot be empty", ErrInvalidDeployConfig)
		}
		if len(d.L1MintManagerRecipients) != len(d.L1MintManagerShares) {
			return fmt.Errorf("%w: L1MintManagerRecipients and L1MintManagerShares must be the same length", ErrInvalidDeployConfig)
		}
		if len(d.L2MintManagerRecipients) == 0 {
			return fmt.Errorf("%w: L2MintManagerRecipients array cannot be empty", ErrInvalidDeployConfig)
		}
		if len(d.L2MintManagerRecipients) != len(d.L2MintManagerShares) {
			return fmt.Errorf("%w: L2MintManagerRecipients and L2MintManagerShares must be the same length", ErrInvalidDeployConfig)
		}
		// [Kroma: END]
	}
	// L2 block time must always be smaller than L1 block time
	if d.L1BlockTime < d.L2BlockTime {
		return fmt.Errorf("L2 block time (%d) is larger than L1 block time (%d)", d.L2BlockTime, d.L1BlockTime)
	}
	/* [Kroma: START]
	if d.RequiredProtocolVersion == (params.ProtocolVersion{}) {
		log.Warn("RequiredProtocolVersion is empty")
	}
	if d.RecommendedProtocolVersion == (params.ProtocolVersion{}) {
		log.Warn("RecommendedProtocolVersion is empty")
	}
	if d.ProofMaturityDelaySeconds == 0 {
		log.Warn("ProofMaturityDelaySeconds is 0")
	}
	if d.DisputeGameFinalityDelaySeconds == 0 {
		log.Warn("DisputeGameFinalityDelaySeconds is 0")
	}
	[Kroma: END] */

	if d.UsePlasma {
		if d.DAChallengeWindow == 0 {
			return fmt.Errorf("%w: DAChallengeWindow cannot be 0 when using plasma mode", ErrInvalidDeployConfig)
		}
		if d.DAResolveWindow == 0 {
			return fmt.Errorf("%w: DAResolveWindow cannot be 0 when using plasma mode", ErrInvalidDeployConfig)
		}
		if d.DAChallengeProxy == (common.Address{}) {
			return fmt.Errorf("%w: DAChallengeContract cannot be empty when using plasma mode", ErrInvalidDeployConfig)
		}
	}
	// checkFork checks that fork A is before or at the same time as fork B
	checkFork := func(a, b *hexutil.Uint64, aName, bName string) error {
		if a == nil && b == nil {
			return nil
		}
		if a == nil && b != nil {
			return fmt.Errorf("fork %s set (to %d), but prior fork %s missing", bName, *b, aName)
		}
		if a != nil && b == nil {
			return nil
		}
		if *a > *b {
			return fmt.Errorf("fork %s set to %d, but prior fork %s has higher offset %d", bName, *b, aName, *a)
		}
		return nil
	}
	if err := checkFork(d.L2GenesisRegolithTimeOffset, d.L2GenesisCanyonTimeOffset, "regolith", "canyon"); err != nil {
		return err
	}
	if err := checkFork(d.L2GenesisCanyonTimeOffset, d.L2GenesisDeltaTimeOffset, "canyon", "delta"); err != nil {
		return err
	}
	if err := checkFork(d.L2GenesisDeltaTimeOffset, d.L2GenesisEcotoneTimeOffset, "delta", "ecotone"); err != nil {
		return err
	}
	// [Kroma: START]
	if err := checkFork(d.L2GenesisEcotoneTimeOffset, d.L2GenesisKromaMPTTimeOffset, "ecotone", "mpt"); err != nil {
		return err
	}
	if err := checkFork(d.L2GenesisKromaMPTTimeOffset, d.L2GenesisFjordTimeOffset, "mpt", "fjord"); err != nil {
		return err
	}

	if d.ValidatorRewardScalar == 0 {
		log.Warn("ValidatorRewardScalar is 0")
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
	if d.ValidatorManagerTrustedValidator == (common.Address{}) {
		return fmt.Errorf("%w: ValidatorManagerTrustedValidator cannot be address(0)", ErrInvalidDeployConfig)
	}
	if d.ValidatorManagerMinRegisterAmount == nil {
		return fmt.Errorf("%w: ValidatorManagerMinRegisterAmount cannot be nil", ErrInvalidDeployConfig)
	}
	if d.ValidatorManagerMinActivateAmount == nil {
		return fmt.Errorf("%w: ValidatorManagerMinActivateAmount cannot be nil", ErrInvalidDeployConfig)
	}
	if d.ValidatorManagerMinActivateAmount.ToInt().Cmp(d.ValidatorManagerMinRegisterAmount.ToInt()) < 0 {
		return fmt.Errorf("%w: ValidatorManagerMinActivateAmount must equal or more than ValidatorManagerMinRegisterAmount", ErrInvalidDeployConfig)
	}
	if d.ValidatorManagerMptFirstOutputIndex == nil {
		return fmt.Errorf("%w: ValidatorManagerMptFirstOutputIndex cannot be nil", ErrInvalidDeployConfig)
	}
	if d.ValidatorManagerCommissionChangeDelaySeconds == 0 {
		return fmt.Errorf("%w: ValidatorManagerCommissionChangeDelaySeconds cannot be 0", ErrInvalidDeployConfig)
	}
	if d.ValidatorManagerRoundDurationSeconds == 0 {
		return fmt.Errorf("%w: ValidatorManagerRoundDurationSeconds cannot be 0", ErrInvalidDeployConfig)
	}
	if d.ValidatorManagerSoftJailPeriodSeconds == 0 {
		return fmt.Errorf("%w: ValidatorManagerSoftJailPeriodSeconds cannot be 0", ErrInvalidDeployConfig)
	}
	if d.ValidatorManagerHardJailPeriodSeconds == 0 {
		return fmt.Errorf("%w: ValidatorManagerHardJailPeriodSeconds cannot be 0", ErrInvalidDeployConfig)
	}
	if d.ValidatorManagerJailThreshold == 0 {
		return fmt.Errorf("%w: ValidatorManagerJailThreshold cannot be 0", ErrInvalidDeployConfig)
	}
	if d.ValidatorManagerMaxFinalizations == 0 {
		return fmt.Errorf("%w: ValidatorManagerMaxFinalizations cannot be 0", ErrInvalidDeployConfig)
	}
	if d.ValidatorManagerBaseReward == nil {
		return fmt.Errorf("%w: ValidatorManagerBaseReward cannot be nil", ErrInvalidDeployConfig)
	}
	if d.AssetManagerKgh == (common.Address{}) {
		return fmt.Errorf("%w: AssetManagerKgh cannot be address(0)", ErrInvalidDeployConfig)
	}
	if d.AssetManagerVault == (common.Address{}) {
		return fmt.Errorf("%w: AssetManagerVault cannot be address(0)", ErrInvalidDeployConfig)
	}
	if d.AssetManagerMinDelegationPeriod == 0 {
		return fmt.Errorf("%w: AssetManagerMinDelegationPeriod cannot be 0", ErrInvalidDeployConfig)
	}
	if d.AssetManagerBondAmount == nil {
		return fmt.Errorf("%w: AssetManagerBondAmount cannot be nil", ErrInvalidDeployConfig)
	}
	if d.L2OutputOracleSubmissionInterval*d.L2BlockTime != d.ValidatorManagerRoundDurationSeconds*2 {
		return fmt.Errorf("%w: double of ValidatorManagerRoundDurationSeconds must equal to L2OutputOracleSubmissionInterval", ErrInvalidDeployConfig)
	}
	if d.L2OutputOracleSubmissionInterval*d.L2BlockTime != d.ValidatorPoolRoundDuration*2 {
		return fmt.Errorf("%w: double of ValidatorPoolRoundDuration must equal to L2OutputOracleSubmissionInterval", ErrInvalidDeployConfig)
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
	if d.ZKProofVerifierSP1Verifier == (common.Address{}) {
		return fmt.Errorf("%w: ZKProofVerifierSP1Verifier cannot be address(0)", ErrInvalidDeployConfig)
	}
	if d.ZKProofVerifierVKey == (common.Hash{}) {
		return fmt.Errorf("%w: ZKProofVerifierVKey cannot be 0", ErrInvalidDeployConfig)
	}
	// [Kroma: END]

	return nil
}

// CheckAddresses will return an error if the addresses are not set.
// These values are required to create the L2 genesis state and are present in the deploy config
// even though the deploy config is required to deploy the contracts on L1. This creates a
// circular dependency that should be resolved in the future.
func (d *DeployConfig) CheckAddresses() error {
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
	return nil
}

// SetDeployments will merge a Deployments into a DeployConfig.
func (d *DeployConfig) SetDeployments(deployments *L1Deployments) {
	d.L1StandardBridgeProxy = deployments.L1StandardBridgeProxy
	d.L1CrossDomainMessengerProxy = deployments.L1CrossDomainMessengerProxy
	d.L1ERC721BridgeProxy = deployments.L1ERC721BridgeProxy
	d.SystemConfigProxy = deployments.SystemConfigProxy
	d.KromaPortalProxy = deployments.KromaPortalProxy

	// [Kroma: START]
	d.ValidatorPoolProxy = deployments.ValidatorPoolProxy
	// [Kroma: END]
}

func (d *DeployConfig) GovernanceEnabled() bool {
	return d.EnableGovernance
}

func (d *DeployConfig) RegolithTime(genesisTime uint64) *uint64 {
	if d.L2GenesisRegolithTimeOffset == nil {
		return nil
	}
	v := uint64(0)
	if offset := *d.L2GenesisRegolithTimeOffset; offset > 0 {
		v = genesisTime + uint64(offset)
	}
	return &v
}

func (d *DeployConfig) CanyonTime(genesisTime uint64) *uint64 {
	if d.L2GenesisCanyonTimeOffset == nil {
		return nil
	}
	v := uint64(0)
	if offset := *d.L2GenesisCanyonTimeOffset; offset > 0 {
		v = genesisTime + uint64(offset)
	}
	return &v
}

func (d *DeployConfig) DeltaTime(genesisTime uint64) *uint64 {
	if d.L2GenesisDeltaTimeOffset == nil {
		return nil
	}
	v := uint64(0)
	if offset := *d.L2GenesisDeltaTimeOffset; offset > 0 {
		v = genesisTime + uint64(offset)
	}
	return &v
}

func (d *DeployConfig) EcotoneTime(genesisTime uint64) *uint64 {
	if d.L2GenesisEcotoneTimeOffset == nil {
		return nil
	}
	v := uint64(0)
	if offset := *d.L2GenesisEcotoneTimeOffset; offset > 0 {
		v = genesisTime + uint64(offset)
	}
	return &v
}

func (d *DeployConfig) KromaMPTTime(genesisTime uint64) *uint64 {
	if d.L2GenesisKromaMPTTimeOffset == nil {
		return nil
	}
	v := uint64(0)
	if offset := *d.L2GenesisKromaMPTTimeOffset; offset > 0 {
		v = genesisTime + uint64(offset)
	}
	return &v
}

func (d *DeployConfig) FjordTime(genesisTime uint64) *uint64 {
	if d.L2GenesisFjordTimeOffset == nil {
		return nil
	}
	v := uint64(0)
	if offset := *d.L2GenesisFjordTimeOffset; offset > 0 {
		v = genesisTime + uint64(offset)
	}
	return &v
}

func (d *DeployConfig) InteropTime(genesisTime uint64) *uint64 {
	if d.L2GenesisInteropTimeOffset == nil {
		return nil
	}
	v := uint64(0)
	if offset := *d.L2GenesisInteropTimeOffset; offset > 0 {
		v = genesisTime + uint64(offset)
	}
	return &v
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
		RegolithTime:           d.RegolithTime(l1StartBlock.Time()),
		CanyonTime:             d.CanyonTime(l1StartBlock.Time()),
		DeltaTime:              d.DeltaTime(l1StartBlock.Time()),
		EcotoneTime:            d.EcotoneTime(l1StartBlock.Time()),
		KromaMPTTime:           d.KromaMPTTime(l1StartBlock.Time()),
		FjordTime:              d.FjordTime(l1StartBlock.Time()),
		InteropTime:            d.InteropTime(l1StartBlock.Time()),
		UsePlasma:              d.UsePlasma,
		DAChallengeAddress:     d.DAChallengeProxy,
		DAChallengeWindow:      d.DAChallengeWindow,
		DAResolveWindow:        d.DAResolveWindow,
	}, nil
}

// NewDeployConfig reads a config file given a path on the filesystem.
func NewDeployConfig(path string) (*DeployConfig, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("deploy config at %s not found: %w", path, err)
	}

	dec := json.NewDecoder(bytes.NewReader(file))
	dec.DisallowUnknownFields()

	var config DeployConfig
	if err := dec.Decode(&config); err != nil {
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

// L1Deployments represents a set of L1 contracts that are deployed.
type L1Deployments struct {
	/* [Kroma: START]
	AddressManager common.Address `json:"AddressManager"`
	BlockOracle                       common.Address `json:"BlockOracle"`
	DisputeGameFactory                common.Address `json:"DisputeGameFactory"`
	DisputeGameFactoryProxy           common.Address `json:"DisputeGameFactoryProxy"`
	[Kroma: END] */
	L1CrossDomainMessenger         common.Address `json:"L1CrossDomainMessenger"`
	L1CrossDomainMessengerProxy    common.Address `json:"L1CrossDomainMessengerProxy"`
	L1ERC721Bridge                 common.Address `json:"L1ERC721Bridge"`
	L1ERC721BridgeProxy            common.Address `json:"L1ERC721BridgeProxy"`
	L1StandardBridge               common.Address `json:"L1StandardBridge"`
	L1StandardBridgeProxy          common.Address `json:"L1StandardBridgeProxy"`
	L2OutputOracle                 common.Address `json:"L2OutputOracle"`
	L2OutputOracleProxy            common.Address `json:"L2OutputOracleProxy"`
	KromaMintableERC20Factory      common.Address `json:"KromaMintableERC20Factory"`
	KromaMintableERC20FactoryProxy common.Address `json:"KromaMintableERC20FactoryProxy"`
	KromaPortal                    common.Address `json:"KromaPortal"`
	KromaPortalProxy               common.Address `json:"KromaPortalProxy"`
	ProxyAdmin                     common.Address `json:"ProxyAdmin"`
	SystemConfig                   common.Address `json:"SystemConfig"`
	SystemConfigProxy              common.Address `json:"SystemConfigProxy"`
	/* [Kroma: START]
	ProtocolVersions                  common.Address `json:"ProtocolVersions"`
	ProtocolVersionsProxy             common.Address `json:"ProtocolVersionsProxy"`
	[Kroma: END] */
	DataAvailabilityChallenge      common.Address `json:"DataAvailabilityChallenge"`
	DataAvailabilityChallengeProxy common.Address `json:"DataAvailabilityChallengeProxy"`

	// [Kroma: START]
	Colosseum                 common.Address `json:"Colosseum"`
	ColosseumProxy            common.Address `json:"ColosseumProxy"`
	L1GovernanceToken         common.Address `json:"L1GovernanceToken"`
	L1GovernanceTokenProxy    common.Address `json:"L1GovernanceTokenProxy"`
	L1MintManager             common.Address `json:"L1MintManager"`
	Poseidon2                 common.Address `json:"Poseidon2"`
	SecurityCouncil           common.Address `json:"SecurityCouncil"`
	SecurityCouncilProxy      common.Address `json:"SecurityCouncilProxy"`
	SecurityCouncilToken      common.Address `json:"SecurityCouncilToken"`
	SecurityCouncilTokenProxy common.Address `json:"SecurityCouncilTokenProxy"`
	TimeLock                  common.Address `json:"TimeLock"`
	TimeLockProxy             common.Address `json:"TimeLockProxy"`
	UpgradeGovernor           common.Address `json:"UpgradeGovernor"`
	UpgradeGovernorProxy      common.Address `json:"UpgradeGovernorProxy"`
	ValidatorPool             common.Address `json:"ValidatorPool"`
	ValidatorPoolProxy        common.Address `json:"ValidatorPoolProxy"`
	AssetManager              common.Address `json:"AssetManager"`
	AssetManagerProxy         common.Address `json:"AssetManagerProxy"`
	ValidatorManager          common.Address `json:"ValidatorManager"`
	ValidatorManagerProxy     common.Address `json:"ValidatorManagerProxy"`
	ZKMerkleTrie              common.Address `json:"ZKMerkleTrie"`
	ZKVerifier                common.Address `json:"ZKVerifier"`
	ZKVerifierProxy           common.Address `json:"ZKVerifierProxy"`
	ZKProofVerifier           common.Address `json:"ZKProofVerifier"`
	ZKProofVerifierProxy      common.Address `json:"ZKProofVerifierProxy"`
	// [Kroma: END]
}

// GetName will return the name of the contract given an address.
func (d *L1Deployments) GetName(addr common.Address) string {
	val := reflect.ValueOf(d)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		if addr == val.Field(i).Interface().(common.Address) {
			return val.Type().Field(i).Name
		}
	}
	return ""
}

// Check will ensure that the L1Deployments are sane
func (d *L1Deployments) Check(deployConfig *DeployConfig) error {
	val := reflect.ValueOf(d)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		name := val.Type().Field(i).Name
		// Skip the non production ready contracts
		if name == "DisputeGameFactory" ||
			name == "DisputeGameFactoryProxy" ||
			name == "BlockOracle" {
			continue
		}
		if !deployConfig.UsePlasma &&
			(name == "DataAvailabilityChallenge" ||
				name == "DataAvailabilityChallengeProxy") {
			continue
		}
		if val.Field(i).Interface().(common.Address) == (common.Address{}) {
			return fmt.Errorf("%s is not set", name)
		}
	}
	return nil
}

// ForEach will iterate over each contract in the L1Deployments
func (d *L1Deployments) ForEach(cb func(name string, addr common.Address)) {
	val := reflect.ValueOf(d)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		name := val.Type().Field(i).Name
		cb(name, val.Field(i).Interface().(common.Address))
	}
}

// Copy will copy the L1Deployments struct
func (d *L1Deployments) Copy() *L1Deployments {
	cpy := L1Deployments{}
	data, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &cpy); err != nil {
		panic(err)
	}
	return &cpy
}

// NewL1Deployments will create a new L1Deployments from a JSON file on disk
// at the given path.
func NewL1Deployments(path string) (*L1Deployments, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("L1 deployments at %s not found: %w", path, err)
	}

	var deployments L1Deployments
	if err := json.Unmarshal(file, &deployments); err != nil {
		return nil, fmt.Errorf("cannot unmarshal L1 deployments: %w", err)
	}

	return &deployments, nil
}

// NewStateDump will read a Dump JSON file from disk
func NewStateDump(path string) (*gstate.Dump, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("dump at %s not found: %w", path, err)
	}

	var fdump ForgeDump
	if err := json.Unmarshal(file, &fdump); err != nil {
		return nil, fmt.Errorf("cannot unmarshal dump: %w", err)
	}
	dump := (gstate.Dump)(fdump)
	return &dump, nil
}

// ForgeDump is a simple alias for state.Dump that can read "nonce" as a hex string.
// It appears as if updates to foundry have changed the serialization of the state dump.
type ForgeDump gstate.Dump

func (d *ForgeDump) UnmarshalJSON(b []byte) error {
	type forgeDumpAccount struct {
		Balance     string                 `json:"balance"`
		Nonce       uint64                 `json:"nonce"`
		Root        hexutil.Bytes          `json:"root"`
		CodeHash    hexutil.Bytes          `json:"codeHash"`
		Code        hexutil.Bytes          `json:"code,omitempty"`
		Storage     map[common.Hash]string `json:"storage,omitempty"`
		Address     *common.Address        `json:"address,omitempty"`
		AddressHash hexutil.Bytes          `json:"key,omitempty"`
	}
	type forgeDump struct {
		Root     string                              `json:"root"`
		Accounts map[common.Address]forgeDumpAccount `json:"accounts"`
	}
	var dump forgeDump
	if err := json.Unmarshal(b, &dump); err != nil {
		return err
	}

	d.Root = dump.Root
	d.Accounts = make(map[string]gstate.DumpAccount)
	for addr, acc := range dump.Accounts {
		d.Accounts[addr.String()] = gstate.DumpAccount{
			Balance:     acc.Balance,
			Nonce:       (uint64)(acc.Nonce),
			Root:        acc.Root,
			CodeHash:    acc.CodeHash,
			Code:        acc.Code,
			Storage:     acc.Storage,
			Address:     acc.Address,
			AddressHash: acc.AddressHash,
		}
	}
	return nil
}

// NewL2ImmutableConfig will create an ImmutableConfig given an instance of a
// DeployConfig and a block.
func NewL2ImmutableConfig(config *DeployConfig, block *types.Block) (*immutables.PredeploysImmutableConfig, error) {
	if config.L1StandardBridgeProxy == (common.Address{}) {
		return nil, fmt.Errorf("L1StandardBridgeProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1CrossDomainMessengerProxy == (common.Address{}) {
		return nil, fmt.Errorf("L1CrossDomainMessengerProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1ERC721BridgeProxy == (common.Address{}) {
		return nil, fmt.Errorf("L1ERC721BridgeProxy cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.ProtocolVaultRecipient == (common.Address{}) {
		return nil, fmt.Errorf("ProtocolVaultRecipient cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}
	if config.L1FeeVaultRecipient == (common.Address{}) {
		return nil, fmt.Errorf("L1FeeVaultRecipient cannot be address(0): %w", ErrInvalidImmutablesConfig)
	}

	rewardDivider := config.FinalizationPeriodSeconds / (config.L2OutputOracleSubmissionInterval * config.L2BlockTime)

	cfg := immutables.PredeploysImmutableConfig{
		L2ToL1MessagePasser: struct{}{},
		/* [Kroma: START]
		DeployerWhitelist:   struct{}{},
		[Kroma: END] */
		WETH9: struct{}{},
		// [Kroma: START]
		L2CrossDomainMessenger: struct{ OtherMessenger common.Address }{
			OtherMessenger: config.L1CrossDomainMessengerProxy,
		},
		L2StandardBridge: struct {
			OtherBridge common.Address
		}{
			OtherBridge: config.L1StandardBridgeProxy,
		},
		// [Kroma: END]
		/* [Kroma: START]
		SequencerFeeVault: struct {
			Recipient           common.Address
			MinWithdrawalAmount *big.Int
			WithdrawalNetwork   uint8
		}{
			Recipient:           config.SequencerFeeVaultRecipient,
			MinWithdrawalAmount: (*big.Int)(config.SequencerFeeVaultMinimumWithdrawalAmount),
			WithdrawalNetwork:   config.SequencerFeeVaultWithdrawalNetwork.ToUint8(),
		},
		[Kroma: END] */
		L1BlockNumber:  struct{}{},
		GasPriceOracle: struct{}{},
		KromaL1Block:   struct{}{},
		/* [Kroma: START]
		GovernanceToken: struct{}{},
		LegacyMessagePasser: struct{}{},
		[Kroma: END]*/
		L2ERC721Bridge: struct {
			OtherBridge common.Address
			Messenger   common.Address
		}{
			OtherBridge: config.L1ERC721BridgeProxy,
			Messenger:   predeploys.L2CrossDomainMessengerAddr,
		},
		KromaMintableERC721Factory: struct {
			Bridge        common.Address
			RemoteChainId *big.Int
		}{
			Bridge:        predeploys.L2ERC721BridgeAddr,
			RemoteChainId: new(big.Int).SetUint64(config.L1ChainID),
		},
		KromaMintableERC20Factory: struct {
			Bridge common.Address
		}{
			Bridge: predeploys.L2StandardBridgeAddr,
		},
		ProxyAdmin: struct{}{},
		// [Kroma: START]
		ProtocolVault: struct {
			Recipient common.Address
		}{
			Recipient: config.ProtocolVaultRecipient,
		},
		L1FeeVault: struct {
			Recipient common.Address
		}{
			Recipient: config.L1FeeVaultRecipient,
		},
		// [Kroma: END]
		/* [Kroma: START]
		SchemaRegistry: struct{}{},
		EAS: struct {
			Name string
		}{
			Name: "EAS",
		},
		[Kroma: END] */
		Create2Deployer: struct{}{},
		// [Kroma: START]
		ValidatorRewardVault: struct {
			ValidatorPoolAddress common.Address
			RewardDivider        *big.Int
		}{
			ValidatorPoolAddress: config.ValidatorPoolProxy,
			RewardDivider:        new(big.Int).SetUint64(rewardDivider),
		},
		// [Kroma: END]
	}

	if err := cfg.Check(); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// NewL2StorageConfig will create a StorageConfig given an instance of a DeployConfig and genesis block.
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
	/* [Kroma: START]
	storage["L2StandardBridge"] = state.StorageValues{
		"_initialized":  1,
		"_initializing": false,
	}
	storage["L2ERC721Bridge"] = state.StorageValues{
		"_initialized":  1,
		"_initializing": false,
	}
	[Kroma: END] */
	storage["KromaL1Block"] = state.StorageValues{
		"number":                block.Number(),
		"timestamp":             block.Time(),
		"basefee":               block.BaseFee(),
		"hash":                  block.Hash(),
		"sequenceNumber":        0,
		"batcherHash":           eth.AddressAsLeftPaddedHash(config.BatchSenderAddress),
		"l1FeeOverhead":         config.GasPriceOracleOverhead,
		"l1FeeScalar":           config.GasPriceOracleScalar,
		"validatorRewardScalar": config.ValidatorRewardScalar,
	}
	/* [Kroma: START]
	storage["LegacyERC20ETH"] = state.StorageValues{
		"_name":   "Ether",
		"_symbol": "ETH",
	}
	[Kroma: END] */
	storage["WETH9"] = state.StorageValues{
		"name":     "Wrapped Ether",
		"symbol":   "WETH",
		"decimals": 18,
	}
	/* [Kroma: START]
	if config.EnableGovernance {
		storage["GovernanceToken"] = state.StorageValues{
			"_name":   config.GovernanceTokenName,
			"_symbol": config.GovernanceTokenSymbol,
			"_owner":  config.GovernanceTokenOwner,
		}
	}
	[Kroma: END] */
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

// Number wraps the rpc.BlockNumberOrHash Number method.
func (m *MarshalableRPCBlockNumberOrHash) Number() (rpc.BlockNumber, bool) {
	return (*rpc.BlockNumberOrHash)(m).Number()
}

// Hash wraps the rpc.BlockNumberOrHash Hash method.
func (m *MarshalableRPCBlockNumberOrHash) Hash() (common.Hash, bool) {
	return (*rpc.BlockNumberOrHash)(m).Hash()
}

// String wraps the rpc.BlockNumberOrHash String method.
func (m *MarshalableRPCBlockNumberOrHash) String() string {
	return (*rpc.BlockNumberOrHash)(m).String()
}
