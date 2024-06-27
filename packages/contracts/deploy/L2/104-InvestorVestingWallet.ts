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
  await deployProxy(hre, 'InvestorVestingWalletProxy', proxyAdminProxy)

  // Deploy impl contract and upgrade proxy
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)

  const investorVestingWallet = await deploy(hre, 'InvestorVestingWallet', {
    contract: 'KromaVestingWallet',
    isProxyImpl: true,
    initArgs: [
      deployConfig.investorVestingWalletBeneficiary,
      deployConfig.investorVestingWalletStartTimestamp,
      deployConfig.investorVestingWalletDurationSeconds,
      deployConfig.investorVestingWalletCliffDivider,
      deployConfig.investorVestingWalletCycleSeconds,
      deployConfig.investorVestingWalletOwner,
    ],
  })

  // Ensure variables are set correctly after initialization
  assertContractVariable(
    investorVestingWallet,
    'beneficiary',
    deployConfig.investorVestingWalletBeneficiary
  )
  assertContractVariable(
    investorVestingWallet,
    'start',
    deployConfig.investorVestingWalletStartTimestamp
  )
  assertContractVariable(
    investorVestingWallet,
    'duration',
    deployConfig.investorVestingWalletDurationSeconds
  )
  assertContractVariable(
    investorVestingWallet,
    'cliffDivider',
    deployConfig.investorVestingWalletCliffDivider
  )
  assertContractVariable(
    investorVestingWallet,
    'vestingCycle',
    deployConfig.investorVestingWalletCycleSeconds
  )
  assertContractVariable(
    investorVestingWallet,
    'owner',
    deployConfig.investorVestingWalletOwner
  )
}

deployFn.tags = ['InvestorVestingWallet', 'l2', 'tge']

export default deployFn
