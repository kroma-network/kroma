import '@kroma/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getDeploymentAddress,
} from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)

  const GovernanceTokenProxyAddress = await getDeploymentAddress(
    hre,
    'GovernanceTokenProxy'
  )

  await deploy(hre, 'MintManager', {
    args: [
      GovernanceTokenProxyAddress,
      deployConfig.mintManagerOwner,
      deployConfig.l2MintManagerRecipients,
      deployConfig.l2MintManagerShares,
    ],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'GOVERNANCE_TOKEN',
        GovernanceTokenProxyAddress
      )
      await assertContractVariable(
        contract,
        'pendingOwner',
        deployConfig.mintManagerOwner
      )
    },
  })
}

deployFn.tags = ['MintManager', 'l2', 'tge']

export default deployFn
