import '@wemixkanvas/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getContractFromArtifact,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const L2OutputOracleProxy = await getContractFromArtifact(
    hre,
    'L2OutputOracleProxy'
  )

  const ZKMerkleTrie = await getContractFromArtifact(hre, 'ZKMerkleTrie')

  await deploy(hre, 'KanvasPortal', {
    args: [
      L2OutputOracleProxy.address,
      hre.deployConfig.portalGuardian,
      false,
      ZKMerkleTrie.address,
    ],
    isProxyImpl: true,
    initArgs: [false],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'L2_ORACLE',
        L2OutputOracleProxy.address
      )
      await assertContractVariable(
        contract,
        'GUARDIAN',
        hre.deployConfig.portalGuardian
      )
      await assertContractVariable(
        contract,
        'ZK_MERKLE_TRIE',
        ZKMerkleTrie.address
      )
    },
  })
}

deployFn.tags = ['KanvasPortal', 'setup']

export default deployFn
