import { ethers } from 'ethers'

/**
 * Core required deployment configuration.
 */
interface RequiredDeployConfig {
  /**
   * Number of confirmations to wait when deploying contracts.
   */
  numDeployConfirmations?: number

  /**
   * Address that will own the entire system on L1 during the deployment process. This address will
   * not own the system after the deployment is complete, ownership will be transferred to the
   * final system owner.
   */
  controller?: string

  /**
   * The L2 genesis script uses this to fill the storage of the L1Block info predeploy. The rollup
   * config script uses this to fill the L1 genesis info for the rollup. The Output oracle deploy
   * script may use it if the L2 starting timestamp is undefined, assuming the L2 genesis is set up
   * with this.
   */
  l1StartingBlockTag: string

  /**
   * Chain ID for the L1 network.
   */
  l1ChainID: number

  /**
   * Chain ID for the L2 network.
   */
  l2ChainID: number

  /**
   * Number of seconds in between each L2 block.
   */
  l2BlockTime: number

  /**
   * Sequencer batches may not be more than maxSequencerDrift seconds after the L1 timestamp of the
   * end of the sequencing window end.
   */
  maxSequencerDrift: number

  /**
   * Number of L1 blocks per sequencing window.
   */
  sequencerWindowSize: number

  /**
   * Number of L1 blocks that a frame stays valid when included in L1.
   */
  channelTimeout: number

  /**
   * Address of the key the sequencer uses to sign blocks on the P2P layer.
   */
  p2pSequencerAddress: string

  /**
   * L1 address that batches are sent to.
   */
  batchInboxAddress: string

  /**
   * Acceptable batch-sender address, to filter transactions going into the batchInboxAddress on L1 for data.
   * Warning: this address is hardcoded now, but is intended to become governed via L1.
   */
  batchSenderAddress: string

  /**
   * Address of the trusted validator.
   */
  validatorPoolTrustedValidator: string

  /**
   * Amount of the required bond in hex value.
   */
  validatorPoolRequiredBondAmount: string

  /**
   * Max number of unbonds when trying unbond.
   */
  validatorPoolMaxUnbond: number

  /**
   * The duration of one submission round in seconds.
   */
  validatorPoolRoundDuration: number

  /**
   * The output index where ValidatorPool is terminated after in hex value.
   */
  validatorPoolTerminateOutputIndex: string

  /**
   * Address of the trusted validator.
   */
  validatorManagerTrustedValidator: string

  /**
   * Amount of the minimum register amount in hex value.
   */
  validatorManagerMinRegisterAmount: string

  /**
   * Amount of the minimum activation amount in hex value.
   */
  validatorManagerMinActivateAmount: string

  /**
   * The first output index after MPT transition.
   */
  validatorManagerMptFirstOutputIndex: string

  /**
   * The delay to finalize the commission rate change in seconds.
   */
  validatorManagerCommissionChangeDelaySeconds: number

  /**
   * The duration of one submission round in seconds.
   */
  validatorManagerRoundDurationSeconds: number

  /**
   * The duration of jail period in seconds in output non-submissions penalty.
   */
  validatorManagerSoftJailPeriodSeconds: number

  /**
   * The duration of jail period in seconds in slashing penalty.
   */
  validatorManagerHardJailPeriodSeconds: number

  /**
   * The threshold of output non-submission to be jailed.
   */
  validatorManagerJailThreshold: number

  /**
   * Max number of output finalizations when distributing reward.
   */
  validatorManagerMaxFinalizations: number

  /**
   * Amount of the base reward in hex value.
   */
  validatorManagerBaseReward: string

  /**
   * Address of the KGH NFT contract.
   */
  assetManagerKgh: string

  /**
   * Address of the validator reward vault.
   */
  assetManagerVault: string

  /**
   * The duration of minimum delegation period in seconds.
   */
  assetManagerMinDelegationPeriod: number

  /**
   * Amount of the bond amount in hex value.
   */
  assetManagerBondAmount: string

  /**
   * Output Oracle submission interval in L2 blocks.
   */
  l2OutputOracleSubmissionInterval: number

  /**
   * Starting block number for the output oracle.
   * Must be greater than or equal to the L2 genesis block, and the first L2 output will correspond to this value.
   */
  l2OutputOracleStartingBlockNumber?: number

  /**
   * Starting timestamp for the output oracle.
   * MUST be the same as the timestamp of the L2OO start block.
   */
  l2OutputOracleStartingTimestamp?: number

  /**
   * Output finalization period in seconds.
   */
  finalizationPeriodSeconds: number

  /**
   * The period seconds for which challenges can be created per each output.
   */
  colosseumCreationPeriodSeconds: number

  /**
   * Dummy hash to be used to compute zkEVM proof as a padding if
   * the number of transaction is less than maximum number of transactions.
   */
  colosseumDummyHash: string

  /**
   * Maximum the number of transaction are allowed in a block when computing zkEVM proof.
   */
  colosseumMaxTxs: number

  /**
   * List of segments length that must be submitted for each turn of the challenge.
   * A value represented by a comma-separated string like `9,6,5,6`
   */
  colosseumSegmentsLengths: string

  /**
   * Owner of the ProxyAdmin contract.
   */
  proxyAdminOwner: string

  /**
   * L1 recipient of fees accumulated in the ProtocolVault.
   */
  protocolVaultRecipient: string

  /**
   * L1 recipient of fees accumulated in the L1FeeVault.
   */
  l1FeeVaultRecipient: string

  /**
   * Timeout seconds of bisection in the Colosseum.
   */
  colosseumBisectionTimeout: number

  /**
   * Timeout seconds of proving in the Colosseum.
   */
  colosseumProvingTimeout: number

  /**
   * The value used by line 459 of the ZK verifier contract.
   */
  zkVerifierHashScalar: string

  /**
   * The value used by line 1173 of the ZK verifier contract.
   */
  zkVerifierM56Px: string

  /**
   * The value used by line 1173 of the ZK verifier contract.
   */
  zkVerifierM56Py: string

  /**
   * Address of the SP1VerifierGateway contract.
   */
  zkProofVerifierSP1Verifier: string

  /**
   * The verification key for the zkVM program.
   */
  zkProofVerifierVKey: string

  /**
   * Governor voting delay in block.
   */
  governorVotingDelayBlocks: number

  /**
   * Governor voting period in block.
   */
  governorVotingPeriodBlocks: number

  /**
   * Governor proposal threshold.
   */
  governorProposalThreshold: number

  /**
   * Quorum as a fraction of the token's total supply.
   */
  governorVotesQuorumFractionPercent: number

  /**
   * Initial minimum delay for operations.
   */
  timeLockMinDelaySeconds: number

  /**
   * L2 : Initial minimum delay for operations.
   */
  l2TimeLockMinDelaySeconds: number

  /**
   * L2 : Governor voting period in block.
   */
  l2GovernorVotingPeriodBlocks: number

  /**
   * The salt to determine GovernanceTokenProxy address on L1 and L2.
   */
  governanceTokenProxySalt: string

  /**
   * The owner of MintManager on L1 and L2.
   */
  mintManagerOwner: string
}

/**
 * Optional deployment configuration when spinning up an L1 network as part of the deployment.
 */
interface OptionalL1DeployConfig {
  cliqueSignerAddress: string
  l1BlockTime: number
  l1GenesisBlockNonce: string
  l1GenesisBlockGasLimit: string
  l1GenesisBlockDifficulty: string
  l1GenesisBlockMixHash: string
  l1GenesisBlockCoinbase: string
  l1GenesisBlockNumber: string
  l1GenesisBlockGasUsed: string
  l1GenesisBlockParentHash: string
  l1GenesisBlockBaseFeePerGas: string
  governanceTokenNotUseCreate2: boolean
  l1MintManagerRecipients: string[]
  l1MintManagerShares: string[]
}

/**
 * Optional deployment configuration when spinning up an L2 network as part of the deployment.
 */
interface OptionalL2DeployConfig {
  l2GenesisBlockNonce: string
  l2GenesisBlockGasLimit: string
  l2GenesisBlockDifficulty: string
  l2GenesisBlockMixHash: string
  l2GenesisBlockNumber: string
  l2GenesisBlockGasUsed: string
  l2GenesisBlockParentHash: string
  l2GenesisBlockBaseFeePerGas: string
  l2GenesisBlockCoinbase: string
  eip1559Denominator: number
  eip1559Elasticity: number
  gasPriceOracleOverhead: number
  gasPriceOracleScalar: number
  validatorRewardScalar: number
  l2MintManagerRecipients: string[]
  l2MintManagerShares: string[]
}

/**
 * Full deployment configuration.
 */
export type DeployConfig = RequiredDeployConfig &
  Partial<OptionalL1DeployConfig> &
  Partial<OptionalL2DeployConfig>

/**
 * Deployment configuration specification for the hardhat plugin.
 */
export const deployConfigSpec: {
  [K in keyof DeployConfig]: {
    type: 'string' | 'number' | 'boolean' | 'address'
    default?: any
  }
} = {
  numDeployConfirmations: {
    type: 'number',
    default: 1,
  },
  l1StartingBlockTag: {
    type: 'string',
  },
  l1ChainID: {
    type: 'number',
  },
  l2ChainID: {
    type: 'number',
  },
  l2BlockTime: {
    type: 'number',
  },
  maxSequencerDrift: {
    type: 'number',
  },
  sequencerWindowSize: {
    type: 'number',
  },
  channelTimeout: {
    type: 'number',
  },
  p2pSequencerAddress: {
    type: 'address',
  },
  batchInboxAddress: {
    type: 'address',
  },
  batchSenderAddress: {
    type: 'address',
  },
  validatorPoolTrustedValidator: {
    type: 'address',
  },
  validatorPoolRequiredBondAmount: {
    type: 'string', // uint256
  },
  validatorPoolMaxUnbond: {
    type: 'number',
  },
  validatorPoolRoundDuration: {
    type: 'number',
  },
  validatorPoolTerminateOutputIndex: {
    type: 'string', // uint256
  },
  validatorManagerTrustedValidator: {
    type: 'address',
  },
  validatorManagerMinRegisterAmount: {
    type: 'string', // uint128
  },
  validatorManagerMinActivateAmount: {
    type: 'string', // uint128
  },
  validatorManagerMptFirstOutputIndex: {
    type: 'string', // uint256
  },
  validatorManagerCommissionChangeDelaySeconds: {
    type: 'number',
  },
  validatorManagerRoundDurationSeconds: {
    type: 'number',
  },
  validatorManagerSoftJailPeriodSeconds: {
    type: 'number',
  },
  validatorManagerHardJailPeriodSeconds: {
    type: 'number',
  },
  validatorManagerJailThreshold: {
    type: 'number',
  },
  validatorManagerMaxFinalizations: {
    type: 'number',
  },
  validatorManagerBaseReward: {
    type: 'string', // uint128
  },
  assetManagerKgh: {
    type: 'address',
  },
  assetManagerVault: {
    type: 'address',
  },
  assetManagerMinDelegationPeriod: {
    type: 'number',
  },
  assetManagerBondAmount: {
    type: 'string', // uint128
  },
  l2OutputOracleSubmissionInterval: {
    type: 'number',
  },
  l2OutputOracleStartingBlockNumber: {
    type: 'number',
    default: 0,
  },
  l2OutputOracleStartingTimestamp: {
    type: 'number',
  },
  finalizationPeriodSeconds: {
    type: 'number',
    default: 2,
  },
  proxyAdminOwner: {
    type: 'address',
  },
  protocolVaultRecipient: {
    type: 'address',
  },
  l1FeeVaultRecipient: {
    type: 'address',
  },
  cliqueSignerAddress: {
    type: 'address',
    default: ethers.constants.AddressZero,
  },
  l1BlockTime: {
    type: 'number',
    default: 15,
  },
  l1GenesisBlockNonce: {
    type: 'string', // uint64
    default: '0x0',
  },
  l1GenesisBlockGasLimit: {
    type: 'string',
    default: ethers.BigNumber.from(15_000_000).toHexString(),
  },
  l1GenesisBlockDifficulty: {
    type: 'string', // uint256
    default: '0x1',
  },
  l1GenesisBlockMixHash: {
    type: 'string', // bytes32
    default: ethers.constants.HashZero,
  },
  l1GenesisBlockCoinbase: {
    type: 'address',
    default: ethers.constants.AddressZero,
  },
  l1GenesisBlockNumber: {
    type: 'string', // uint64
    default: '0x0',
  },
  l1GenesisBlockGasUsed: {
    type: 'string', // uint64
    default: '0x0',
  },
  l1GenesisBlockParentHash: {
    type: 'string', // bytes32
    default: ethers.constants.HashZero,
  },
  l1GenesisBlockBaseFeePerGas: {
    type: 'string', // uint256
    default: ethers.BigNumber.from(1000_000_000).toHexString(), // 1 gwei
  },
  l2GenesisBlockNonce: {
    type: 'string', // uint64
    default: '0x0',
  },
  l2GenesisBlockGasLimit: {
    type: 'string',
    default: ethers.BigNumber.from(15_000_000).toHexString(),
  },
  l2GenesisBlockDifficulty: {
    type: 'string', // uint256
    default: '0x1',
  },
  l2GenesisBlockMixHash: {
    type: 'string', // bytes32
    default: ethers.constants.HashZero,
  },
  l2GenesisBlockNumber: {
    type: 'string', // uint64
    default: '0x0',
  },
  l2GenesisBlockGasUsed: {
    type: 'string', // uint64
    default: '0x0',
  },
  l2GenesisBlockParentHash: {
    type: 'string', // bytes32
    default: ethers.constants.HashZero,
  },
  l2GenesisBlockBaseFeePerGas: {
    type: 'string', // uint256
    default: ethers.BigNumber.from(1000_000_000).toHexString(), // 1 gwei
  },
  gasPriceOracleOverhead: {
    type: 'number',
    default: 2100,
  },
  gasPriceOracleScalar: {
    type: 'number',
    default: 1_000_000,
  },
  validatorRewardScalar: {
    type: 'number',
    default: 0,
  },
  colosseumCreationPeriodSeconds: {
    type: 'number',
  },
  colosseumBisectionTimeout: {
    type: 'number',
    default: 3600,
  },
  colosseumProvingTimeout: {
    type: 'number',
    default: 3600,
  },
  colosseumDummyHash: {
    type: 'string', // bytes32
    default: ethers.constants.HashZero,
  },
  colosseumMaxTxs: {
    type: 'number',
    default: 0,
  },
  colosseumSegmentsLengths: {
    type: 'string', // comma-separated segments lengths
  },
  zkVerifierHashScalar: {
    type: 'string', // uint256
  },
  zkVerifierM56Px: {
    type: 'string', // uint256
  },
  zkVerifierM56Py: {
    type: 'string', // uint256
  },
  zkProofVerifierSP1Verifier: {
    type: 'address',
    default: ethers.constants.AddressZero,
  },
  zkProofVerifierVKey: {
    type: 'string', // bytes32
    default: ethers.constants.HashZero,
  },
  governorVotingDelayBlocks: {
    type: 'number',
  },
  governorVotingPeriodBlocks: {
    type: 'number',
  },
  governorProposalThreshold: {
    type: 'number',
  },
  governorVotesQuorumFractionPercent: {
    type: 'number',
  },
  timeLockMinDelaySeconds: {
    type: 'number',
  },
  l2GovernorVotingPeriodBlocks: {
    type: 'number',
  },
  l2TimeLockMinDelaySeconds: {
    type: 'number',
  },
  governanceTokenNotUseCreate2: {
    type: 'boolean',
    default: false,
  },
  governanceTokenProxySalt: {
    type: 'string', // bytes32
    default: ethers.constants.HashZero,
  },
  mintManagerOwner: {
    type: 'address',
  },
}
