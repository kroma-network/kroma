import '@kroma-network/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getDeploymentAddress,
} from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const colosseumProxyAddress = await getDeploymentAddress(
    hre,
    'ColosseumProxy'
  )

  await deploy(hre, 'SecurityCouncil', {
    args: [colosseumProxyAddress, hre.deployConfig.securityCouncilTokenOwner],
    isProxyImpl: true,
    initializer: 'initialize(bool,address[],uint256)',
    initArgs: [
      false,
      hre.deployConfig.securityCouncilOwners,
      hre.deployConfig.securityCouncilNumConfirmationRequired,
    ],
    postDeployAction: async (contract) => {
      await assertContractVariable(contract, 'COLOSSEUM', colosseumProxyAddress)
      await assertContractVariable(
        contract,
        'GOVERNOR',
        hre.deployConfig.securityCouncilTokenOwner
      )
    },
  })
}

deployFn.tags = ['SecurityCouncil', 'setup', 'l1']

export default deployFn
