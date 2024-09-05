import '@kroma/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getDeploymentAddress,
} from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1GovernanceTokenProxyAddress = await getDeploymentAddress(
    hre,
    'L1GovernanceTokenProxy'
  )
  const scProxyAddress = await getDeploymentAddress(hre, 'SecurityCouncilProxy')
  const valMgrProxyAddress = await getDeploymentAddress(
    hre,
    'ValidatorManagerProxy'
  )

  await deploy(hre, 'AssetManager', {
    args: [
      l1GovernanceTokenProxyAddress,
      hre.deployConfig.assetManagerKgh,
      scProxyAddress,
      hre.deployConfig.assetManagerVault,
      valMgrProxyAddress,
      hre.deployConfig.assetManagerMinDelegationPeriod,
      hre.deployConfig.assetManagerBondAmount,
    ],
    isProxyImpl: true,
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'ASSET_TOKEN',
        l1GovernanceTokenProxyAddress
      )
      await assertContractVariable(
        contract,
        'KGH',
        hre.deployConfig.assetManagerKgh
      )
      await assertContractVariable(contract, 'SECURITY_COUNCIL', scProxyAddress)
      await assertContractVariable(
        contract,
        'VALIDATOR_REWARD_VAULT',
        hre.deployConfig.assetManagerVault
      )
      await assertContractVariable(
        contract,
        'VALIDATOR_MANAGER',
        valMgrProxyAddress
      )
      await assertContractVariable(
        contract,
        'MIN_DELEGATION_PERIOD',
        hre.deployConfig.assetManagerMinDelegationPeriod
      )
      await assertContractVariable(
        contract,
        'BOND_AMOUNT',
        hre.deployConfig.assetManagerBondAmount
      )
    },
  })
}

deployFn.tags = ['AssetManager', 'setup', 'l1', 'validatorSystemUpgrade']

export default deployFn
