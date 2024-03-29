import '@kroma/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import { ethers } from 'ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { assertContractVariable, deploy } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)

  const l1FeeVaultRecipient = deployConfig.l1FeeVaultRecipient
  if (l1FeeVaultRecipient === ethers.constants.AddressZero) {
    throw new Error('L1FeeVault RECIPIENT zero address')
  }

  await deploy(hre, 'L1FeeVault', {
    args: [l1FeeVaultRecipient],
    isProxyImpl: true,
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'RECIPIENT',
        ethers.utils.getAddress(l1FeeVaultRecipient)
      )
    },
  })
}

deployFn.tags = ['L1FeeVault', 'l2']

export default deployFn
