import '@kroma/hardhat-deploy-config'
import assert from 'assert'

import { ethers } from 'ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { deploy, getDeploymentAddress } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)
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
      deployConfig.l2TimeLockMinDelaySeconds,
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
      deployConfig.l2TimeLockMinDelaySeconds
  )
}

deployFn.tags = ['TimeLock', 'setup', 'l2']

export default deployFn
