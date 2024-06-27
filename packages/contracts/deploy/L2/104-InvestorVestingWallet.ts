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
  const investorVestingWalletProxy = await deployProxy(
    hre,
    'InvestorVestingWalletProxy',
    deployer
  )

  // Deploy impl contract and upgrade proxy
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)

  await deployAndUpgradeByDeployer(hre, 'InvestorVestingWallet', {
    contract: 'KromaVestingWallet',
    args: [
      deployConfig.investorVestingWalletCliffDivider,
      deployConfig.investorVestingWalletCycleSeconds,
    ],
    isProxyImpl: true,
    initArgs: [
      deployConfig.investorVestingWalletBeneficiary,
      deployConfig.investorVestingWalletStartTimestamp,
      deployConfig.investorVestingWalletDurationSeconds,
    ],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'CLIFF_DIVIDER',
        deployConfig.investorVestingWalletCliffDivider
      )
      await assertContractVariable(
        contract,
        'VESTING_CYCLE',
        deployConfig.investorVestingWalletCycleSeconds
      )
    },
  })

  // Ensure variables are set correctly after initialization
  const upgradedProxy = await hre.ethers.getContractAt(
    'KromaVestingWallet',
    investorVestingWalletProxy.address
  )
  assertContractVariable(
    upgradedProxy,
    'beneficiary',
    deployConfig.investorVestingWalletBeneficiary
  )
  assertContractVariable(
    upgradedProxy,
    'start',
    deployConfig.investorVestingWalletStartTimestamp
  )
  assertContractVariable(
    upgradedProxy,
    'duration',
    deployConfig.investorVestingWalletDurationSeconds
  )

  // Change admin of InvestorVestingWalletProxy to mintManagerOwner
  const signer = hre.ethers.provider.getSigner(deployer)
  const contractWithSigner = investorVestingWalletProxy.connect(signer)
  const tx = await contractWithSigner.changeAdmin(deployConfig.mintManagerOwner)
  await tx.wait()
  console.log(
    `changed admin of "InvestorVestingWalletProxy" to ${deployConfig.mintManagerOwner}`
  )
}

deployFn.tags = ['InvestorVestingWallet', 'l2', 'tge']

export default deployFn
