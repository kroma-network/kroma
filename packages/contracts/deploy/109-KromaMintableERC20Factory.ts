import '@kroma-network/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import { ethers } from 'ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { predeploys } from '../src/constants'
import { assertContractVariable, deploy } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  await deploy(hre, 'KromaMintableERC20Factory', {
    args: [predeploys.L2StandardBridge],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'BRIDGE',
        ethers.utils.getAddress(predeploys.L2StandardBridge)
      )
    },
  })
}

deployFn.tags = ['KromaMintableERC20Factory', 'l2']

export default deployFn
