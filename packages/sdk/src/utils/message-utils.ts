import { hashWithdrawal } from '@kroma/core-utils'
import { ethers } from 'ethers'

import { LowLevelMessage } from '../interfaces'

/**
 * Utility for hashing a LowLevelMessage object.
 *
 * @param message LowLevelMessage object to hash.
 * @returns Hash of the given LowLevelMessage.
 */
export const hashLowLevelMessage = (message: LowLevelMessage): string => {
  return hashWithdrawal(
    message.messageNonce,
    message.sender,
    message.target,
    message.value,
    message.minGasLimit,
    message.message
  )
}

/**
 * Utility for hashing a message hash. This computes the storage slot
 * where the message hash will be stored in state. HashZero is used
 * because the first mapping in the contract is used.
 *
 * @param messageHash Message hash to hash.
 * @returns Hash of the given message hash.
 */
export const hashMessageHash = (messageHash: string): string => {
  const data = ethers.utils.defaultAbiCoder.encode(
    ['bytes32', 'uint256'],
    [messageHash, ethers.constants.HashZero]
  )
  return ethers.utils.keccak256(data)
}
