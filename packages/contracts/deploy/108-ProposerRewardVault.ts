import '@kroma-network/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import { ethers } from 'ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { assertContractVariable, deploy } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)

  const ProposerRewardVaultRecipient = deployConfig.ProposerRewardVaultRecipient
  if (ProposerRewardVaultRecipient === ethers.constants.AddressZero) {
    throw new Error('ProposerRewardVault RECIPIENT zero address')
  }

  await deploy(hre, 'ProposerRewardVault', {
    args: [ProposerRewardVaultRecipient],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'RECIPIENT',
        ethers.utils.getAddress(ProposerRewardVaultRecipient)
      )
    },
  })
}

deployFn.tags = ['ProposerRewardVault', 'l2']

export default deployFn
