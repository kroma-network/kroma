import { DeployFunction } from 'hardhat-deploy/dist/types'

import { deployDeterministicProxy } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  // Deploy proxy contract with deployer as admin
  const { deployer } = await hre.getNamedAccounts()

  await deployDeterministicProxy(
    hre,
    'L1GovernanceTokenProxy',
    deployer,
    hre.deployConfig.governanceTokenProxySalt
  )
}

deployFn.tags = ['L1GovernanceTokenProxy', 'l1', 'tge']

export default deployFn
