import '@kroma/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getDeploymentAddress,
} from '../../src/deploy-utils'
import { ethers } from 'ethers'
import { predeploys } from '../../src'

const deployFn: DeployFunction = async (hre) => {
  const zeroAddress = ethers.constants.AddressZero
  const l1StandardBridgeProxyAddress = await getDeploymentAddress(
    hre,
    'L1StandardBridgeProxy'
  )

  await deploy(hre, 'L1GovernanceToken', {
    contract: 'GovernanceToken',
    args: [
      l1StandardBridgeProxyAddress,
      predeploys.GovernanceToken,
      zeroAddress,
    ],
    isProxyImpl: true,
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
      await assertContractVariable(contract, 'MINT_MANAGER', zeroAddress)
    },
  })
}

deployFn.tags = ['SecurityCouncil', 'setup', 'l1']

export default deployFn
