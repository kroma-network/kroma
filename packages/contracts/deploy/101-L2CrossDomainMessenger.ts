import '@kroma-network/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import { ethers } from 'ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { assertContractVariable, deploy } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const Artifact__L1CrossDomainMessenger = await hre.companionNetworks[
    'l1'
  ].deployments.get('L1CrossDomainMessengerProxy')

  await deploy(hre, 'L2CrossDomainMessenger', {
    args: [Artifact__L1CrossDomainMessenger.address],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'OTHER_MESSENGER',
        ethers.utils.getAddress(Artifact__L1CrossDomainMessenger.address)
      )
    },
  })
}

deployFn.tags = ['L2CrossDomainMessenger', 'l2']

export default deployFn
