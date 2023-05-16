import '@kroma-network/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { assertContractVariable, deploy } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  await deploy(hre, 'GasPriceOracle', {
    args: [],
    postDeployAction: async (contract) => {
      await assertContractVariable(contract, 'DECIMALS', 6)
    },
  })
}

deployFn.tags = ['GasPriceOracle', 'l2']

export default deployFn
