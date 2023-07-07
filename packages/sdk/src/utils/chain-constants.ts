import { predeploys } from '@kroma-network/contracts'

import {
  ContractsLike,
  L1ChainID,
  L2ChainID,
  L2ContractsLike,
} from '../interfaces'

export const DEPOSIT_CONFIRMATION_BLOCKS: {
  [ChainID in L2ChainID]: number
} = {
  [L2ChainID.KROMA_SEPOLIA]: 4 as const, // 2 epoch, 1 epoch = 32 slot
  [L2ChainID.KROMA_LOCAL_DEVNET]: 2 as const,
}

export const CHAIN_BLOCK_TIMES: {
  [ChainID in L1ChainID]: number
} = {
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
  [L2ChainID.KROMA_SEPOLIA]: {
    l1: {
      KromaPortal: '0x16cEb19A3ABF1A8B56f53dB50eb22695b6eF7BcC' as const,
      L1CrossDomainMessenger:
        '0xCfE879a845b7bdb1fC51B84F7607fb41044f4004' as const,
      L1StandardBridge: '0xB6a9294251FF3a920D7C0204A45B1F7FfE4D2983' as const,
      L2OutputOracle: '0xB70D7dBa8ac50842820E703C63022Ef52220410B' as const,
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
