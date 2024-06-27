import '@kroma/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  deployProxy,
  getDeploymentAddress,
} from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  // Deploy proxy contract
  const proxyAdminProxy = await getDeploymentAddress(hre, 'ProxyAdminProxy')
  await deployProxy(hre, 'TeamVestingWalletProxy', proxyAdminProxy)

  // Deploy impl contract and upgrade proxy
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)

  const teamVestingWallet = await deploy(hre, 'TeamVestingWallet', {
    contract: 'KromaVestingWallet',
    isProxyImpl: true,
    initArgs: [
      deployConfig.teamVestingWalletBeneficiary,
      deployConfig.teamVestingWalletStartTimestamp,
      deployConfig.teamVestingWalletDurationSeconds,
      deployConfig.teamVestingWalletCliffDivider,
      deployConfig.teamVestingWalletCycleSeconds,
      deployConfig.teamVestingWalletOwner,
    ],
  })

  // Ensure variables are set correctly after initialization
  assertContractVariable(
    teamVestingWallet,
    'beneficiary',
    deployConfig.teamVestingWalletBeneficiary
  )
  assertContractVariable(
    teamVestingWallet,
    'start',
    deployConfig.teamVestingWalletStartTimestamp
  )
  assertContractVariable(
    teamVestingWallet,
    'duration',
    deployConfig.teamVestingWalletDurationSeconds
  )
  assertContractVariable(
    teamVestingWallet,
    'cliffDivider',
    deployConfig.teamVestingWalletCliffDivider
  )
  assertContractVariable(
    teamVestingWallet,
    'vestingCycle',
    deployConfig.teamVestingWalletCycleSeconds
  )
  assertContractVariable(
    teamVestingWallet,
    'owner',
    deployConfig.teamVestingWalletOwner
  )
}

deployFn.tags = ['TeamVestingWallet', 'l2', 'tge']

export default deployFn
