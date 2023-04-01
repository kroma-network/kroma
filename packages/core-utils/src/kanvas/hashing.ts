import { defaultAbiCoder } from '@ethersproject/abi'
import { BigNumberish, BigNumber } from '@ethersproject/bignumber'
import { keccak256 } from '@ethersproject/keccak256'

import { decodeVersionedNonce, encodeCrossDomainMessageV0 } from './encoding'

/**
 * Output oracle data.
 */
export interface OutputData {
  outputRoot: string
  l1Timestamp: number
  l2BlockNumber: number
  l2OutputIndex: number
}

/**
 * State commitment
 */
export interface OutputRootProof {
  version: string
  stateRoot: string
  messagePasserStorageRoot: string
  latestBlockhash: string
}

/**
 * Proof data required to finalize an L2 to L1 message.
 */
export interface CrossChainMessageProof {
  l2OutputIndex: number
  outputRootProof: OutputRootProof
  withdrawalProof: string[]
}

/**
 * Parameters that govern the L2OutputOracle.
 */
export type L2OutputOracleParameters = {
  submissionInterval: number
  startingBlockNumber: number
  l2BlockTime: number
}

/**
 * Hashes a cross domain message.
 *
 * @param nonce     The cross domain message nonce
 * @param sender    The sender of the cross domain message
 * @param target    The target of the cross domain message
 * @param value     The value being sent with the cross domain message
 * @param gasLimit  The gas limit of the cross domain execution
 * @param data      The data passed along with the cross domain message
 */
export const hashCrossDomainMessage = (
  nonce: BigNumber,
  sender: string,
  target: string,
  value: BigNumber,
  gasLimit: BigNumber,
  data: string
) => {
  const { version } = decodeVersionedNonce(nonce)
  if (version.eq(0)) {
    return hashCrossDomainMessageV0(
      nonce,
      sender,
      target,
      value,
      gasLimit,
      data
    )
  }
  throw new Error(`unknown version ${version.toString()}`)
}

/**
 * Hashes a V0 cross domain message
 *
 * @param nonce     The cross domain message nonce
 * @param sender    The sender of the cross domain message
 * @param target    The target of the cross domain message
 * @param value     The value being sent with the cross domain message
 * @param gasLimit  The gas limit of the cross domain execution
 * @param data      The data passed along with the cross domain message
 */
export const hashCrossDomainMessageV0 = (
  nonce: BigNumber,
  sender: string,
  target: string,
  value: BigNumberish,
  gasLimit: BigNumberish,
  data: string
) => {
  return keccak256(
    encodeCrossDomainMessageV0(nonce, sender, target, value, gasLimit, data)
  )
}

/**
 * Hashes a withdrawal
 *
 * @param nonce     The cross domain message nonce
 * @param sender    The sender of the cross domain message
 * @param target    The target of the cross domain message
 * @param value     The value being sent with the cross domain message
 * @param gasLimit  The gas limit of the cross domain execution
 * @param data      The data passed along with the cross domain message
 */
export const hashWithdrawal = (
  nonce: BigNumber,
  sender: string,
  target: string,
  value: BigNumber,
  gasLimit: BigNumber,
  data: string
): string => {
  const types = ['uint256', 'address', 'address', 'uint256', 'uint256', 'bytes']
  const encoded = defaultAbiCoder.encode(types, [
    nonce,
    sender,
    target,
    value,
    gasLimit,
    data,
  ])
  return keccak256(encoded)
}

/**
 * Hashes an output root proof
 *
 * @param proof OutputRootProof
 */
export const hashOutputRootProof = (proof: OutputRootProof): string => {
  return keccak256(
    defaultAbiCoder.encode(
      ['bytes32', 'bytes32', 'bytes32', 'bytes32'],
      [
        proof.version,
        proof.stateRoot,
        proof.messagePasserStorageRoot,
        proof.latestBlockhash,
      ]
    )
  )
}
