import '@kroma/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getDeploymentAddress,
} from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1StandardBridgeProxyAddress = await getDeploymentAddress(
    hre,
    'L1StandardBridgeProxy'
  )

  const governanceTokenProxyAddress = await getDeploymentAddress(
    hre,
    'L1GovernanceTokenProxy'
  )

  await deploy(hre, 'L1GovernanceToken', {
    contract: 'GovernanceToken',
    args: [l1StandardBridgeProxyAddress, governanceTokenProxyAddress],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'BRIDGE',
        l1StandardBridgeProxyAddress
      )
      await assertContractVariable(
        contract,
        'REMOTE_TOKEN',
        governanceTokenProxyAddress
      )
    },
  })
}

deployFn.tags = ['L1GovernanceToken', 'l1', 'tge']

export default deployFn
