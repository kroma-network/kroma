import { DeployFunction } from 'hardhat-deploy/dist/types'

import { deployDeterministicProxy } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)

  await deployDeterministicProxy(
    hre,
    'GovernanceTokenProxy',
    deployConfig.mintManagerOwner,
    deployConfig.governanceTokenProxySalt
  )
}

deployFn.tags = ['GovernanceTokenProxy', 'l2', 'tge']

export default deployFn
