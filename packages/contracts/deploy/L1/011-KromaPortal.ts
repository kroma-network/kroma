import '@kroma/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getContractFromArtifact,
  getDeploymentAddress,
} from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l2OutputOracleProxyAddress = await getDeploymentAddress(
    hre,
    'L2OutputOracleProxy'
  )

  const validatorPoolProxyAddress = await getDeploymentAddress(
    hre,
    'ValidatorPoolProxy'
  )

  const Artifact__SystemConfigProxy = await hre.deployments.get(
    'SystemConfigProxy'
  )

  const securityCouncilProxyAddress = await getDeploymentAddress(
    hre,
    'SecurityCouncilProxy'
  )

  const ZKMerkleTrie = await getContractFromArtifact(hre, 'ZKMerkleTrie')

  await deploy(hre, 'KromaPortal', {
    args: [
      l2OutputOracleProxyAddress,
      validatorPoolProxyAddress,
      securityCouncilProxyAddress,
      false,
      Artifact__SystemConfigProxy.address,
      ZKMerkleTrie.address,
    ],
    isProxyImpl: true,
    initArgs: [false],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'L2_ORACLE',
        l2OutputOracleProxyAddress
      )
      await assertContractVariable(
        contract,
        'VALIDATOR_POOL',
        validatorPoolProxyAddress
      )
      await assertContractVariable(
        contract,
        'GUARDIAN',
        securityCouncilProxyAddress
      )
      await assertContractVariable(
        contract,
        'SYSTEM_CONFIG',
        Artifact__SystemConfigProxy.address
      )
      await assertContractVariable(
        contract,
        'ZK_MERKLE_TRIE',
        ZKMerkleTrie.address
      )
    },
  })
}

deployFn.tags = ['KromaPortal', 'setup', 'l1', 'mpt']

export default deployFn
