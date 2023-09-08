import '@kroma/hardhat-deploy-config'
import assert from 'assert'

import { ethers } from 'ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { deploy, getDeploymentAddress } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const upgradeGovernorProxyAddress = await getDeploymentAddress(
    hre,
    'UpgradeGovernorProxy'
  )
  const timeLockProxyAddress = await getDeploymentAddress(hre, 'TimeLockProxy')
  const { deployer } = await hre.getNamedAccounts()

  await deploy(hre, 'TimeLock', {
    isProxyImpl: true,
    initializer: 'initialize(uint256,address[],address[],address)',
    initArgs: [
      hre.deployConfig.timeLockMinDelaySeconds,
      [upgradeGovernorProxyAddress],
      [upgradeGovernorProxyAddress],
      upgradeGovernorProxyAddress,
    ],
  })

  const artifact = await hre.deployments.get('TimeLock')
  const timeLock = new ethers.Contract(
    timeLockProxyAddress,
    artifact.abi,
    hre.ethers.provider.getSigner(deployer)
  )

  // Check variable
  assert(
    (await timeLock.getMinDelay()).toNumber() ===
      hre.deployConfig.timeLockMinDelaySeconds
  )
}

deployFn.tags = ['TimeLock', 'setup', 'l1']

export default deployFn
