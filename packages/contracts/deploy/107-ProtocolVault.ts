import '@kroma-network/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import { ethers } from 'ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { assertContractVariable, deploy } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)

  const protocolVaultRecipient = deployConfig.protocolVaultRecipient
  if (protocolVaultRecipient === ethers.constants.AddressZero) {
    throw new Error('ProtocolVault RECIPIENT zero address')
  }

  await deploy(hre, 'ProtocolVault', {
    args: [protocolVaultRecipient],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'RECIPIENT',
        ethers.utils.getAddress(protocolVaultRecipient)
      )
    },
  })
}

deployFn.tags = ['ProtocolVault', 'l2']

export default deployFn
