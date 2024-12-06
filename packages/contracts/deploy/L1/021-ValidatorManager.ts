import '@kroma/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import assert from 'assert'

import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getDeploymentAddress,
} from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l2OutputOracleProxyAddress = await getDeploymentAddress(
    hre,
    'L2OutputOracleProxy'
  )
  const assetManagerProxyAddress = await getDeploymentAddress(
    hre,
    'AssetManagerProxy'
  )

  assert(
    (hre.deployConfig.l2OutputOracleSubmissionInterval *
      hre.deployConfig.l2BlockTime) /
      2 ===
      hre.deployConfig.validatorManagerRoundDurationSeconds
  )

  await deploy(hre, 'ValidatorManager', {
    args: [
      {
        _l2Oracle: l2OutputOracleProxyAddress,
        _assetManager: assetManagerProxyAddress,
        _trustedValidator: hre.deployConfig.validatorManagerTrustedValidator,
        _commissionChangeDelaySeconds:
          hre.deployConfig.validatorManagerCommissionChangeDelaySeconds,
        _roundDurationSeconds:
          hre.deployConfig.validatorManagerRoundDurationSeconds,
        _softJailPeriodSeconds:
          hre.deployConfig.validatorManagerSoftJailPeriodSeconds,
        _hardJailPeriodSeconds:
          hre.deployConfig.validatorManagerHardJailPeriodSeconds,
        _jailThreshold: hre.deployConfig.validatorManagerJailThreshold,
        _maxOutputFinalizations:
          hre.deployConfig.validatorManagerMaxFinalizations,
        _baseReward: hre.deployConfig.validatorManagerBaseReward,
        _minRegisterAmount: hre.deployConfig.validatorManagerMinRegisterAmount,
        _minActivateAmount: hre.deployConfig.validatorManagerMinActivateAmount,
        _mptFirstOutputIndex:
          hre.deployConfig.validatorManagerMptFirstOutputIndex,
      },
    ],
    isProxyImpl: true,
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'L2_ORACLE',
        l2OutputOracleProxyAddress
      )
      await assertContractVariable(
        contract,
        'ASSET_MANAGER',
        assetManagerProxyAddress
      )
      await assertContractVariable(
        contract,
        'TRUSTED_VALIDATOR',
        hre.deployConfig.validatorManagerTrustedValidator
      )
      await assertContractVariable(
        contract,
        'MIN_REGISTER_AMOUNT',
        hre.deployConfig.validatorManagerMinRegisterAmount
      )
      await assertContractVariable(
        contract,
        'MIN_ACTIVATE_AMOUNT',
        hre.deployConfig.validatorManagerMinActivateAmount
      )
      await assertContractVariable(
        contract,
        'COMMISSION_CHANGE_DELAY_SECONDS',
        hre.deployConfig.validatorManagerCommissionChangeDelaySeconds
      )
      await assertContractVariable(
        contract,
        'ROUND_DURATION_SECONDS',
        hre.deployConfig.validatorManagerRoundDurationSeconds
      )
      await assertContractVariable(
        contract,
        'SOFT_JAIL_PERIOD_SECONDS',
        hre.deployConfig.validatorManagerSoftJailPeriodSeconds
      )
      await assertContractVariable(
        contract,
        'HARD_JAIL_PERIOD_SECONDS',
        hre.deployConfig.validatorManagerHardJailPeriodSeconds
      )
      await assertContractVariable(
        contract,
        'JAIL_THRESHOLD',
        hre.deployConfig.validatorManagerJailThreshold
      )
      await assertContractVariable(
        contract,
        'MAX_OUTPUT_FINALIZATIONS',
        hre.deployConfig.validatorManagerMaxFinalizations
      )
      await assertContractVariable(
        contract,
        'BASE_REWARD',
        hre.deployConfig.validatorManagerBaseReward
      )
      await assertContractVariable(
        contract,
        'MPT_FIRST_OUTPUT_INDEX',
        hre.deployConfig.validatorManagerMptFirstOutputIndex
      )
    },
  })
}

deployFn.tags = ['ValidatorManager', 'setup', 'l1', 'validatorSystemUpgrade']

export default deployFn
