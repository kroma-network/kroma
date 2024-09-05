import '@kroma/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { predeploys } from '../../src'
import {
  assertContractVariable,
  deployAndUpgradeByDeployer,
  getDeploymentAddress,
} from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)

  const governanceTokenProxyAddress = await getDeploymentAddress(
    hre,
    'GovernanceTokenProxy'
  )
  const mintManagerAddress = await getDeploymentAddress(hre, 'MintManager')

  await deployAndUpgradeByDeployer(hre, 'GovernanceToken', {
    args: [predeploys.L2StandardBridge, governanceTokenProxyAddress],
    initArgs: [mintManagerAddress],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'BRIDGE',
        predeploys.L2StandardBridge
      )
      await assertContractVariable(
        contract,
        'REMOTE_TOKEN',
        governanceTokenProxyAddress
      )
    },
  })

  // Ensure variables are set correctly after initialization
  const upgradedProxy = await hre.ethers.getContractAt(
    'GovernanceToken',
    governanceTokenProxyAddress
  )
  assertContractVariable(upgradedProxy, 'pendingOwner', mintManagerAddress)

  // Change admin of GovernanceTokenProxy to mintManagerOwner
  const { deployer } = await hre.getNamedAccounts()
  const signer = hre.ethers.provider.getSigner(deployer)
  const governanceTokenProxy = await hre.ethers.getContractAt(
    'Proxy',
    governanceTokenProxyAddress
  )
  const contractWithSigner = governanceTokenProxy.connect(signer)
  const tx = await contractWithSigner.changeAdmin(deployConfig.mintManagerOwner)
  await tx.wait()
  console.log(
    `changed admin of "GovernanceTokenProxy" to ${deployConfig.mintManagerOwner}`
  )
}

deployFn.tags = ['GovernanceToken', 'l2', 'tge']

export default deployFn
