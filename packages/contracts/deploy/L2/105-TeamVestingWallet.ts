import '@kroma/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deployAndUpgradeByDeployer,
  deployProxy,
} from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  // Deploy proxy contract with deployer as admin
  const { deployer } = await hre.getNamedAccounts()
  const teamVestingWalletProxy = await deployProxy(
    hre,
    'TeamVestingWalletProxy',
    deployer
  )

  // Deploy impl contract and upgrade proxy
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)

  await deployAndUpgradeByDeployer(hre, 'TeamVestingWallet', {
    contract: 'KromaVestingWallet',
    args: [
      deployConfig.teamVestingWalletCliffDivider,
      deployConfig.teamVestingWalletCycleSeconds,
    ],
    initArgs: [
      deployConfig.teamVestingWalletBeneficiary,
      deployConfig.teamVestingWalletStartTimestamp,
      deployConfig.teamVestingWalletDurationSeconds,
    ],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'CLIFF_DIVIDER',
        deployConfig.teamVestingWalletCliffDivider
      )
      await assertContractVariable(
        contract,
        'VESTING_CYCLE',
        deployConfig.teamVestingWalletCycleSeconds
      )
    },
  })

  // Ensure variables are set correctly after initialization
  const upgradedProxy = await hre.ethers.getContractAt(
    'KromaVestingWallet',
    teamVestingWalletProxy.address
  )
  assertContractVariable(
    upgradedProxy,
    'beneficiary',
    deployConfig.teamVestingWalletBeneficiary
  )
  assertContractVariable(
    upgradedProxy,
    'start',
    deployConfig.teamVestingWalletStartTimestamp
  )
  assertContractVariable(
    upgradedProxy,
    'duration',
    deployConfig.teamVestingWalletDurationSeconds
  )

  // Change admin of TeamVestingWalletProxy to mintManagerOwner
  const signer = hre.ethers.provider.getSigner(deployer)
  const contractWithSigner = teamVestingWalletProxy.connect(signer)
  const tx = await contractWithSigner.changeAdmin(deployConfig.mintManagerOwner)
  await tx.wait()
  console.log(
    `changed admin of "TeamVestingWalletProxy" to ${deployConfig.mintManagerOwner}`
  )
}

deployFn.tags = ['TeamVestingWallet', 'l2', 'tge']

export default deployFn
