import '@kroma/hardhat-deploy-config'
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

  const upgradeGovernorProxyAddress = await getDeploymentAddress(
    hre,
    'UpgradeGovernorProxy'
  )

  await deploy(hre, 'SecurityCouncil', {
    args: [colosseumProxyAddress, upgradeGovernorProxyAddress],
    isProxyImpl: true,
    postDeployAction: async (contract) => {
      await assertContractVariable(contract, 'COLOSSEUM', colosseumProxyAddress)
      await assertContractVariable(
        contract,
        'GOVERNOR',
        upgradeGovernorProxyAddress
      )
    },
  })
}

deployFn.tags = ['SecurityCouncil', 'setup', 'l1']

export default deployFn
