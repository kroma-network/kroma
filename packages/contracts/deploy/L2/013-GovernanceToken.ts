import '@kroma/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { predeploys } from '../../src'
import { assertContractVariable, deploy } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const Artifact__L1GovernanceTokenProxy = await hre.companionNetworks[
    'l1'
  ].deployments.get('L1GovernanceTokenProxy')

  await deploy(hre, 'GovernanceToken', {
    args: [
      predeploys.L2StandardBridge,
      Artifact__L1GovernanceTokenProxy.address,
    ],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'BRIDGE',
        predeploys.L2StandardBridge
      )
      await assertContractVariable(
        contract,
        'REMOTE_TOKEN',
        Artifact__L1GovernanceTokenProxy.address
      )
    },
  })
}

deployFn.tags = ['GovernanceToken', 'l2', 'tge']

export default deployFn
