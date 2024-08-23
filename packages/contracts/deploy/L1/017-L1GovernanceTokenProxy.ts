import { DeployFunction } from 'hardhat-deploy/dist/types'

import { deployDeterministicProxy, deployProxy } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  // Deploy proxy contract with deployer as admin
  const { deployer } = await hre.getNamedAccounts()
  const name = 'L1GovernanceTokenProxy'

  if (hre.deployConfig.governanceTokenNotUseCreate2) {
    await deployProxy(hre, name, deployer)
  } else {
    await deployDeterministicProxy(
      hre,
      name,
      deployer,
      hre.deployConfig.governanceTokenProxySalt
    )
  }
}

deployFn.tags = ['L1GovernanceTokenProxy', 'setup', 'l1', 'tge']

export default deployFn
