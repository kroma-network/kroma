/* Imports: External */
import { toRpcHexString } from '@wemixkanvas/core-utils'
import { BigNumber, ethers } from 'ethers'

/**
 * Generates a Merkle-Patricia trie proof for a given account and storage slot.
 *
 * @param provider RPC provider attached to an EVM-compatible chain.
 * @param blockNumber Block number to generate the proof at.
 * @param address Address to generate the proof for.
 * @param slot Storage slot to generate the proof for.
 * @returns Account proof and storage proof.
 */
export const makeStateTrieProof = async (
  provider: ethers.providers.JsonRpcProvider,
  blockNumber: number,
  address: string,
  slot: string
): Promise<{
  accountProof: string[]
  storageProof: string[]
  storageValue: BigNumber
  storageRoot: string
}> => {
  const proof = await provider.send('eth_getProof', [
    address,
    [slot],
    toRpcHexString(blockNumber),
  ])

  return {
    accountProof: proof.accountProof,
    storageProof: proof.storageProof[0].proof,
    storageValue: BigNumber.from(proof.storageProof[0].value),
    storageRoot: proof.storageHash,
  }
}
