import '@nomiclabs/hardhat-ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getDeploymentAddress,
} from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const zkVerifierProxyAddress = await getDeploymentAddress(
    hre,
    'ZKVerifierProxy'
  )
  const zkMerkleTrieAddress = await getDeploymentAddress(hre, 'ZKMerkleTrie')

  await deploy(hre, 'ZKProofVerifier', {
    args: [
      zkVerifierProxyAddress,
      hre.deployConfig.colosseumDummyHash,
      hre.deployConfig.colosseumMaxTxs,
      zkMerkleTrieAddress,
      hre.deployConfig.zkProofVerifierSP1Verifier,
      hre.deployConfig.zkProofVerifierVKey,
    ],
    isProxyImpl: true,
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'zkVerifier',
        zkVerifierProxyAddress
      )
      await assertContractVariable(
        contract,
        'dummyHash',
        hre.deployConfig.colosseumDummyHash
      )
      await assertContractVariable(
        contract,
        'maxTxs',
        hre.deployConfig.colosseumMaxTxs
      )
      await assertContractVariable(
        contract,
        'zkMerkleTrie',
        zkMerkleTrieAddress
      )
      await assertContractVariable(
        contract,
        'sp1Verifier',
        hre.deployConfig.zkProofVerifierSP1Verifier
      )
      await assertContractVariable(
        contract,
        'zkVmProgramVKey',
        hre.deployConfig.zkProofVerifierVKey
      )
    },
  })
}

deployFn.tags = ['ZKProofVerifier', 'setup', 'l1', 'mpt']

export default deployFn
