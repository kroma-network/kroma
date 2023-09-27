import '@kroma/hardhat-deploy-config'
import assert from 'assert'

import { ethers } from 'ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { deploy, getDeploymentAddress } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)
  const securityCouncilTokenProxyAddress = await getDeploymentAddress(
    hre,
    'SecurityCouncilTokenProxy'
  )
  const { deployer } = await hre.getNamedAccounts()
  await deploy(hre, 'SecurityCouncilToken', {
    isProxyImpl: true,
    initializer: 'initialize(address)',
    initArgs: [deployer],
  })

  const artifact = await hre.deployments.get('SecurityCouncilToken')
  const token = new ethers.Contract(
    securityCouncilTokenProxyAddress,
    artifact.abi,
    hre.ethers.provider.getSigner(deployer)
  )

  // Check variable
  assert((await token.name()) === 'KromaSecurityCouncil')
  assert((await token.symbol()) === 'KSC')
  assert((await token.owner()) === deployer)

  // Minting to guardians
  for (const [
    index,
    guardian,
  ] of deployConfig.securityCouncilOwners.entries()) {
    const balance = await token.balanceOf(guardian)
    if (balance.toNumber() === 0) {
      const res = `${index + 1}.png`
      await token.safeMint(guardian, res)
    }
  }
}

deployFn.tags = ['SecurityCouncilToken', 'setup', 'l2']

export default deployFn
