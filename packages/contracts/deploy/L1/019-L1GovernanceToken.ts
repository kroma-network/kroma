import '@kroma/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deployAndUpgradeByDeployer,
  getDeploymentAddress,
} from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1StandardBridgeProxyAddress = await getDeploymentAddress(
    hre,
    'L1StandardBridgeProxy'
  )
  const governanceTokenProxyAddress = await getDeploymentAddress(
    hre,
    'L1GovernanceTokenProxy'
  )
  const l1MintManagerAddress = await getDeploymentAddress(hre, 'L1MintManager')

  await deployAndUpgradeByDeployer(hre, 'L1GovernanceToken', {
    contract: 'GovernanceToken',
    args: [l1StandardBridgeProxyAddress, governanceTokenProxyAddress],
    initArgs: [l1MintManagerAddress],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'BRIDGE',
        l1StandardBridgeProxyAddress
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
  assertContractVariable(upgradedProxy, 'pendingOwner', l1MintManagerAddress)

  // Change admin of L1GovernanceTokenProxy to mintManagerOwner
  const { deployer } = await hre.getNamedAccounts()
  const signer = hre.ethers.provider.getSigner(deployer)
  const l1GovernanceTokenProxy = await hre.ethers.getContractAt(
    'Proxy',
    governanceTokenProxyAddress
  )
  const contractWithSigner = l1GovernanceTokenProxy.connect(signer)
  const tx = await contractWithSigner.changeAdmin(
    hre.deployConfig.mintManagerOwner
  )
  await tx.wait()
  console.log(
    `changed admin of "L1GovernanceTokenProxy" to ${hre.deployConfig.mintManagerOwner}`
  )
}

deployFn.tags = ['L1GovernanceToken', 'setup', 'l1', 'tge']

export default deployFn
