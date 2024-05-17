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
      hre.deployConfig.assetManagerKghManager,
      scProxyAddress,
      valMgrProxyAddress,
      hre.deployConfig.assetManagerUndelegationPeriod,
      hre.deployConfig.assetManagerSlashingRate,
      hre.deployConfig.assetManagerMinSlashingAmount,
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
      await assertContractVariable(
        contract,
        'KGH_MANAGER',
        hre.deployConfig.assetManagerKghManager
      )
      await assertContractVariable(contract, 'SECURITY_COUNCIL', scProxyAddress)
      await assertContractVariable(
        contract,
        'VALIDATOR_MANAGER',
        valMgrProxyAddress
      )
      await assertContractVariable(
        contract,
        'UNDELEGATION_PERIOD',
        hre.deployConfig.assetManagerUndelegationPeriod
      )
      await assertContractVariable(
        contract,
        'SLASHING_RATE',
        hre.deployConfig.assetManagerSlashingRate
      )
      await assertContractVariable(
        contract,
        'MIN_SLASHING_AMOUNT',
        hre.deployConfig.assetManagerMinSlashingAmount
      )
    },
  })
}

deployFn.tags = ['AssetManager', 'setup', 'l1', 'validatorSystemUpgrade']

export default deployFn
