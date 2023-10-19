import { predeploys } from '@kroma/contracts'

import {
  ContractsLike,
  L1ChainID,
  L2ChainID,
  L2ContractsLike,
} from '../interfaces'

export const DEPOSIT_CONFIRMATION_BLOCKS: {
  [ChainID in L2ChainID]: number
} = {
  [L2ChainID.KROMA_MAINNET]: 4 as const, // 4 slot
  [L2ChainID.KROMA_SEPOLIA]: 4 as const, // 4 slot
  [L2ChainID.KROMA_LOCAL_DEVNET]: 2 as const,
}

export const CHAIN_BLOCK_TIMES: {
  [ChainID in L1ChainID]: number
} = {
  [L1ChainID.MAINNET]: 12 as const,
  [L1ChainID.SEPOLIA]: 12 as const,
  [L1ChainID.LOCAL_DEVNET]: 3 as const,
}

/**
 * Full list of default L2 contract addresses.
 */
export const DEFAULT_L2_CONTRACT_ADDRESSES: L2ContractsLike = {
  L2CrossDomainMessenger: predeploys.L2CrossDomainMessenger,
  L2StandardBridge: predeploys.L2StandardBridge,
  L2ToL1MessagePasser: predeploys.L2ToL1MessagePasser,
  WETH9: predeploys.WETH9,
}

/**
 * Mapping of L2 chain IDs to the appropriate contract addresses for the deployments to the
 * given network. Simplifies the process of getting the correct contract addresses for a given
 * contract name.
 */
export const CONTRACT_ADDRESSES: {
  [ChainID in L2ChainID]: ContractsLike
} = {
  [L2ChainID.KROMA_MAINNET]: {
    l1: {
      KromaPortal: '0x31F648572b67e60Ec6eb8E197E1848CC5F5558de' as const,
      L1CrossDomainMessenger:
        '0x46B8bB4C5dd27bB42807Db477af4d1a7C8A5B746' as const,
      L1StandardBridge: '0x827962404D7104202C5aaa6b929115C8211d9596' as const,
      L2OutputOracle: '0x180c77aE51a9c505a43A2C7D81f8CE70cacb93A6' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.KROMA_SEPOLIA]: {
    l1: {
      KromaPortal: '0x31ab8eD993A3BE9Aa2757C7D368Dc87101A868a4' as const,
      L1CrossDomainMessenger:
        '0x69786A10c1A153191BF5A50B61e70F6934fcc0A2' as const,
      L1StandardBridge: '0x38C9a0a694AA0f92c05238484C3a9bdE1e85ddE4' as const,
      L2OutputOracle: '0x7291913342063fd10d31651735BAF3877D2F9645' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.KROMA_LOCAL_DEVNET]: {
    l1: {
      KromaPortal: '0x6900000000000000000000000000000000000003' as const,
      L1CrossDomainMessenger:
        '0x6900000000000000000000000000000000000006' as const,
      L1StandardBridge: '0x6900000000000000000000000000000000000007' as const,
      L2OutputOracle: '0x6900000000000000000000000000000000000004' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
}
