import '@kroma-network/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import { ethers } from 'ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { predeploys } from '../src/constants'
import { assertContractVariable, deploy } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const KromaMintableERC721Factory = await hre.ethers.getContractAt(
    'KromaMintableERC721Factory',
    predeploys.KromaMintableERC721Factory
  )
  const remoteChainId = await KromaMintableERC721Factory.REMOTE_CHAIN_ID()

  await deploy(hre, 'KromaMintableERC721Factory', {
    args: [predeploys.L2StandardBridge, remoteChainId],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'BRIDGE',
        ethers.utils.getAddress(predeploys.L2StandardBridge)
      )
      await assertContractVariable(contract, 'REMOTE_CHAIN_ID', remoteChainId)
    },
  })
}

deployFn.tags = ['KromaMintableERC721Factory', 'l2']

export default deployFn
