import '@kroma-network/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { assertContractVariable, deploy } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  await deploy(hre, 'L2ToL1MessagePasser', {
    args: [],
    postDeployAction: async (contract) => {
      await assertContractVariable(contract, 'MESSAGE_VERSION', 0)
    },
  })
}

deployFn.tags = ['L2ToL1MessagePasser', 'l2']

export default deployFn
