import { Interface } from '@ethersproject/abi'
import { BigNumberish, BigNumber } from '@ethersproject/bignumber'

const iface = new Interface([
  'function relayMessage(uint256,address,address,uint256,uint256,bytes)',
])

const nonceMask = BigNumber.from(
  '0x0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff'
)

/**
 * Encodes the version into the nonce.
 *
 * @param nonce
 * @param version
 */
export const encodeVersionedNonce = (
  nonce: BigNumber,
  version: BigNumber
): BigNumber => {
  return version.or(nonce.shl(240))
}

/**
 * Decodes the version from the nonce and returns the unversioned nonce as well
 * as the version. The version is encoded in the first byte of
 * the nonce. Note that this nonce is the nonce held in the
 * CrossDomainMessenger.
 *
 * @param nonce
 */
export const decodeVersionedNonce = (
  nonce: BigNumber
): {
  version: BigNumber
  nonce: BigNumber
} => {
  return {
    version: nonce.shr(240),
    nonce: nonce.and(nonceMask),
  }
}

/**
 * Encodes a V0 cross domain message.
 *
 * @param nonce     The cross domain message nonce
 * @param sender    The sender of the cross domain message
 * @param target    The target of the cross domain message
 * @param value     The value being sent with the cross domain message
 * @param gasLimit  The gas limit of the cross domain execution
 * @param data      The data passed along with the cross domain message
 */
export const encodeCrossDomainMessageV0 = (
  nonce: BigNumber,
  sender: string,
  target: string,
  value: BigNumberish,
  gasLimit: BigNumberish,
  data: string
) => {
  return iface.encodeFunctionData(
    'relayMessage(uint256,address,address,uint256,uint256,bytes)',
    [nonce, sender, target, value, gasLimit, data]
  )
}

/**
 * Encodes a cross domain message. The version byte in the nonce determines
 * the serialization format that is used.
 *
 * @param nonce     The cross domain message nonce
 * @param sender    The sender of the cross domain message
 * @param target    The target of the cross domain message
 * @param value     The value being sent with the cross domain message
 * @param gasLimit  The gas limit of the cross domain execution
 * @param data      The data passed along with the cross domain message
 */
export const encodeCrossDomainMessage = (
  nonce: BigNumber,
  sender: string,
  target: string,
  value: BigNumber,
  gasLimit: BigNumber,
  data: string
) => {
  const { version } = decodeVersionedNonce(nonce)
  if (version.eq(0)) {
    return encodeCrossDomainMessageV0(
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
