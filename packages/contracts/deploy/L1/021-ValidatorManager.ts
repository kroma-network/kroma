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
        _commissionRateMinChangeSeconds:
          hre.deployConfig.validatorManagerCommissionMinChangeSeconds,
        _roundDurationSeconds:
          hre.deployConfig.validatorManagerRoundDurationSeconds,
        _jailPeriodSeconds: hre.deployConfig.validatorManagerJailPeriodSeconds,
        _jailThreshold: hre.deployConfig.validatorManagerJailThreshold,
        _maxOutputFinalizations:
          hre.deployConfig.validatorManagerMaxFinalizations,
        _baseReward: hre.deployConfig.validatorManagerBaseReward,
        _minRegisterAmount: hre.deployConfig.validatorManagerMinRegisterAmount,
        _minActivateAmount: hre.deployConfig.validatorManagerMinActivateAmount,
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
        'COMMISSION_RATE_MIN_CHANGE_SECONDS',
        hre.deployConfig.validatorManagerCommissionMinChangeSeconds
      )
      await assertContractVariable(
        contract,
        'ROUND_DURATION_SECONDS',
        hre.deployConfig.validatorManagerRoundDurationSeconds
      )
      await assertContractVariable(
        contract,
        'JAIL_PERIOD_SECONDS',
        hre.deployConfig.validatorManagerJailPeriodSeconds
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
    },
  })
}

deployFn.tags = ['ValidatorManager', 'setup', 'l1', 'validatorSystemUpgrade']

export default deployFn
