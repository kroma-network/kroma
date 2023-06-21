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

  await deploy(hre, 'SecurityCouncil', {
    args: [ColosseumProxyAddress],
    isProxyImpl: true,
    initArgs: [
      hre.deployConfig.securityCouncilOwners,
      hre.deployConfig.securityCouncilNumConfirmationRequired,
    ],
    postDeployAction: async (contract) => {
      await assertContractVariable(contract, 'COLOSSEUM', ColosseumProxyAddress)
    },
  })
}

deployFn.tags = ['SecurityCouncil', 'setup', 'l1']

export default deployFn
