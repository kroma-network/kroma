import '@kroma/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { predeploys } from '../../src'
import { assertContractVariable, deploy } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)

  await deploy(hre, 'MintManager', {
    args: [
      predeploys.GovernanceToken,
      deployConfig.mintManagerOwner,
      deployConfig.l2MintManagerRecipients,
      deployConfig.l2MintManagerShares,
    ],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'GOVERNANCE_TOKEN',
        predeploys.GovernanceToken
      )
      await assertContractVariable(
        contract,
        'owner',
        deployConfig.mintManagerOwner
      )
    },
  })
}

deployFn.tags = ['MintManager', 'setup', 'l2', 'tge']

export default deployFn
