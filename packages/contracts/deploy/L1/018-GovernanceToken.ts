import '@kroma/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { predeploys } from '../../src'
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

  const l1MintManagerAddress = await getDeploymentAddress(hre, 'L1MintManager')

  await deploy(hre, 'L1GovernanceToken', {
    contract: 'GovernanceToken',
    args: [l1StandardBridgeProxyAddress, predeploys.GovernanceToken],
    isProxyImpl: true,
    initArgs: [l1MintManagerAddress],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'BRIDGE',
        l1StandardBridgeProxyAddress
      )
      await assertContractVariable(
        contract,
        'REMOTE_TOKEN',
        predeploys.GovernanceToken
      )
    },
  })
}

deployFn.tags = ['L1GovernanceToken', 'setup', 'l1', 'tge']

export default deployFn
