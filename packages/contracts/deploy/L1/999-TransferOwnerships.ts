import assert from 'assert'

import '@kroma-network/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'
import { HardhatRuntimeEnvironment } from 'hardhat/types'

import { getDeploymentAddress } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const timelockProxyAddress = await getDeploymentAddress(hre, 'TimeLockProxy')

  await transferProxyAdminOwnership(hre, timelockProxyAddress)
  await transferSystemConfigOwnership(hre, timelockProxyAddress)
}

const transferProxyAdminOwnership = async (
  hre: HardhatRuntimeEnvironment,
  newOwner: string
) => {
  const proxyAdminAddress = await getDeploymentAddress(hre, 'ProxyAdmin')
  let proxyAdmin = await hre.ethers.getContractAt(
    'ProxyAdmin',
    proxyAdminAddress
  )
  const currentProxyAdminOwner = await proxyAdmin.owner()
  if (currentProxyAdminOwner === newOwner) {
    console.log(
      'skip the ProxyAdmin owner transfer process because the owner has already been transferred.'
    )
    return
  }

  proxyAdmin = proxyAdmin.connect(
    hre.ethers.provider.getSigner(currentProxyAdminOwner)
  )

  const tx = await proxyAdmin.transferOwnership(newOwner, {
    from: currentProxyAdminOwner,
  })
  await tx.wait()

  assert((await proxyAdmin.owner()) === newOwner)
}

const transferSystemConfigOwnership = async (
  hre: HardhatRuntimeEnvironment,
  newOwner: string
) => {
  const systemConfigProxyAddress = await getDeploymentAddress(
    hre,
    'SystemConfigProxy'
  )
  let systemConfig = await hre.ethers.getContractAt(
    'SystemConfig',
    systemConfigProxyAddress
  )
  const currentSystemConfigOwner = await systemConfig.owner()
  if (currentSystemConfigOwner === newOwner) {
    console.log(
      'skip the SystemConfig owner transfer process because the owner has already been transferred.'
    )
  }

  systemConfig = systemConfig.connect(
    hre.ethers.provider.getSigner(currentSystemConfigOwner)
  )

  const tx = await systemConfig.transferOwnership(newOwner, {
    from: currentSystemConfigOwner,
  })
  await tx.wait()

  assert((await systemConfig.owner()) === newOwner)
}

deployFn.runAtTheEnd = true
deployFn.tags = ['setup', 'l1']

export default deployFn
