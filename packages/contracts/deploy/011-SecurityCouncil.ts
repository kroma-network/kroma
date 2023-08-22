import '@kroma-network/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getDeploymentAddress,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const ColosseumProxyAddress = await getDeploymentAddress(
    hre,
    'ColosseumProxy'
  )

  const TimeLockProxyAddress = await getDeploymentAddress(hre, 'TimeLockProxy')

  await deploy(hre, 'SecurityCouncil', {
    args: [ColosseumProxyAddress, TimeLockProxyAddress],
    isProxyImpl: true,
    initializer: 'initialize(bool,address[],uint256)',
    initArgs: [
      false,
      hre.deployConfig.securityCouncilOwners,
      hre.deployConfig.securityCouncilNumConfirmationRequired,
    ],
    postDeployAction: async (contract) => {
      await assertContractVariable(contract, 'COLOSSEUM', ColosseumProxyAddress)
      await assertContractVariable(contract, 'GOVERNOR', TimeLockProxyAddress)
    },
  })
}

deployFn.tags = ['SecurityCouncil', 'setup', 'l1']

export default deployFn
