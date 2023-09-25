import assert from 'assert'

import '@kroma/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'
import { HardhatRuntimeEnvironment } from 'hardhat/types'

import { getDeploymentAddress } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)
  const timeLockProxyAddress = await getDeploymentAddress(hre, 'TimeLockProxy')
  await transferProxyAdminOwnership(hre, timeLockProxyAddress)

  const scTokenOwnerAddress = deployConfig.securityCouncilTokenOwner
  await transferSecurityCouncilTokenOwnership(hre, scTokenOwnerAddress)
}

const transferProxyAdminOwnership = async (
  hre: HardhatRuntimeEnvironment,
  newOwner: string
) => {
  console.log('transfer ProxyAdmin ownership to TimeLock')
  const proxyAdminProxyAddress = await getDeploymentAddress(hre, 'ProxyAdminProxy')
  let proxyAdmin = await hre.ethers.getContractAt(
    'ProxyAdmin',
    proxyAdminProxyAddress
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
  console.log('successfully transferred ownership of ProxyAdmin')
}

const transferSecurityCouncilTokenOwnership = async (
  hre: HardhatRuntimeEnvironment,
  newOwner: string
) => {
  console.log(
    'transfer SecurityCouncilToken ownership to SecurityCouncilTokenOwner'
  )
  const scTokenProxyAddress = await getDeploymentAddress(
    hre,
    'SecurityCouncilTokenProxy'
  )
  let scToken = await hre.ethers.getContractAt(
    'SecurityCouncilToken',
    scTokenProxyAddress
  )
  const currentScTokenOwner = await scToken.owner()
  if (
    hre.ethers.utils.getAddress(currentScTokenOwner) ===
    hre.ethers.utils.getAddress(newOwner)
  ) {
    console.log(
      'skip the SecurityCouncilToken owner transfer process because the owner has already been transferred.'
    )
  }

  scToken = scToken.connect(hre.ethers.provider.getSigner(currentScTokenOwner))

  const tx = await scToken.transferOwnership(newOwner, {
    from: currentScTokenOwner,
  })
  await tx.wait()

  assert(
    hre.ethers.utils.getAddress(await scToken.owner()) ===
      hre.ethers.utils.getAddress(newOwner)
  )
  console.log('successfully transferred ownership of SecurityCouncilToken')
}

deployFn.runAtTheEnd = true
deployFn.tags = ['setup', 'l2']

export default deployFn
