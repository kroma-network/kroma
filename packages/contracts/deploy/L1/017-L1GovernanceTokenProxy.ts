import { DeployFunction } from 'hardhat-deploy/dist/types'

import { deployDeterministicProxy } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  await deployDeterministicProxy(
    hre,
    'L1GovernanceTokenProxy',
    hre.deployConfig.mintManagerOwner,
    hre.deployConfig.governanceTokenProxySalt
  )
}

deployFn.tags = ['L1GovernanceTokenProxy', 'l1', 'tge']

export default deployFn
