import '@kroma/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { predeploys } from '../../src'
import { assertContractVariable, deploy, getDeploymentAddress } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const Artifact__L1GovernanceTokenProxy = await hre.companionNetworks[
    'l1'
    ].deployments.get('L1GovernanceTokenProxy')
  const l2StandardBridgeProxyAddress = await getDeploymentAddress(
    hre,
    'L2StandardBridgeProxy'
  )

  await deploy(hre, 'GovernanceToken', {
    args: [
      l2StandardBridgeProxyAddress,
      Artifact__L1GovernanceTokenProxy.address,
      predeploys.MintManager,
    ],
    isProxyImpl: true,
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'BRIDGE',
        l2StandardBridgeProxyAddress
      )
      await assertContractVariable(
        contract,
        'REMOTE_TOKEN',
        Artifact__L1GovernanceTokenProxy.address,
      )
      await assertContractVariable(contract, 'MINT_MANAGER', predeploys.MintManager)
    },
  })
}

deployFn.tags = ['GovernanceToken', 'setup', 'l2', 'tge']

export default deployFn