import { DeployFunction } from 'hardhat-deploy/dist/types'

import { deployDeterministicProxy } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  // Deploy proxy contract with deployer as admin
  const { deployer } = await hre.getNamedAccounts()

  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)

  await deployDeterministicProxy(
    hre,
    'GovernanceTokenProxy',
    deployer,
    deployConfig.governanceTokenProxySalt
  )
}

deployFn.tags = ['GovernanceTokenProxy', 'l2', 'tge']

export default deployFn
