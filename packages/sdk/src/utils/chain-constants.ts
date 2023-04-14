import { predeploys } from '@wemixkanvas/contracts'

import {
  ContractsLike,
  L1ChainID,
  L2ChainID,
  L2ContractsLike,
} from '../interfaces'

export const DEPOSIT_CONFIRMATION_BLOCKS: {
  [ChainID in L2ChainID]: number
} = {
  [L2ChainID.KANVAS_AQUA]: 64 as const, // 2 epoch, 1 epoch = 32 slot
  [L2ChainID.KANVAS_SAIL]: 2 as const,
  [L2ChainID.KANVAS_LOCAL_DEVNET]: 2 as const,
}

export const CHAIN_BLOCK_TIMES: {
  [ChainID in L1ChainID]: number
} = {
  [L1ChainID.SEPOLIA]: 12 as const,
  [L1ChainID.KANVAS_EASEL]: 5 as const,
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
  [L2ChainID.KANVAS_AQUA]: {
    l1: {
      KanvasPortal: '0x9C818e93C0884f75f48d93a9BDB2E994f8d77b86' as const,
      L1CrossDomainMessenger:
        '0x6B6985865e71F0D6B7F8FA6915511b6AE72F778B' as const,
      L1StandardBridge: '0x972C765Ed4C7301d17828D1999BF24d17dAd9230' as const,
      L2OutputOracle: '0x29674FCFc8F24E96dE1c0caBf6366Be9E8A00FA1' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.KANVAS_SAIL]: {
    l1: {
      KanvasPortal: '0xA199e7ab96BF9DF52C52eb7BAb5572789a726d33' as const,
      L1CrossDomainMessenger:
        '0x8233369E29653b70E50E93d1276a50B8f2122a01' as const,
      L1StandardBridge: '0x6B99600daD0a1998337357696827381D122825F3' as const,
      L2OutputOracle: '0xF978b011bcf604b201996FEb3E53eD3D52F0A90F' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
  [L2ChainID.KANVAS_LOCAL_DEVNET]: {
    l1: {
      KanvasPortal: '0x6900000000000000000000000000000000000003' as const,
      L1CrossDomainMessenger:
        '0x6900000000000000000000000000000000000006' as const,
      L1StandardBridge: '0x6900000000000000000000000000000000000007' as const,
      L2OutputOracle: '0x6900000000000000000000000000000000000004' as const,
    },
    l2: DEFAULT_L2_CONTRACT_ADDRESSES,
  },
}
