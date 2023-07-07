import { DeployFunction } from 'hardhat-deploy/dist/types'

import { assertContractVariable, deploy } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const { deployer } = await hre.getNamedAccounts()

  await deploy(hre, 'ProxyAdmin', {
    args: [deployer],
    postDeployAction: async (contract) => {
      await assertContractVariable(contract, 'owner', deployer)
    },
  })
}

deployFn.tags = ['ProxyAdmin', 'setup', 'l1']

export default deployFn
