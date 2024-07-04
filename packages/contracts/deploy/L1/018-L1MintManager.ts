import '@kroma/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getDeploymentAddress,
} from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1GovernanceTokenProxyAddress = await getDeploymentAddress(
    hre,
    'L1GovernanceTokenProxy'
  )

  await deploy(hre, 'L1MintManager', {
    contract: 'MintManager',
    args: [
      l1GovernanceTokenProxyAddress,
      hre.deployConfig.mintManagerOwner,
      hre.deployConfig.l1MintManagerRecipients,
      hre.deployConfig.l1MintManagerShares,
    ],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'GOVERNANCE_TOKEN',
        l1GovernanceTokenProxyAddress
      )
      await assertContractVariable(
        contract,
        'pendingOwner',
        hre.deployConfig.mintManagerOwner
      )
    },
  })
}

deployFn.tags = ['L1MintManager', 'setup', 'l1', 'tge']

export default deployFn
