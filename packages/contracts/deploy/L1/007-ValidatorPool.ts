import '@kroma/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
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
  const portalProxyAddress = await getDeploymentAddress(hre, 'KromaPortalProxy')
  const scProxyAddress = await getDeploymentAddress(hre, 'SecurityCouncilProxy')
  await deploy(hre, 'ValidatorPool', {
    args: [
      l2OutputOracleProxyAddress,
      portalProxyAddress,
      scProxyAddress,
      hre.deployConfig.validatorPoolTrustedValidator,
      hre.deployConfig.validatorPoolRequiredBondAmount,
      hre.deployConfig.validatorPoolMaxUnbond,
      hre.deployConfig.validatorPoolRoundDuration,
      hre.deployConfig.validatorPoolTerminateOutputIndex,
    ],
    isProxyImpl: true,
    initArgs: [],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'L2_ORACLE',
        l2OutputOracleProxyAddress
      )
      await assertContractVariable(contract, 'PORTAL', portalProxyAddress)
      await assertContractVariable(contract, 'SECURITY_COUNCIL', scProxyAddress)
      await assertContractVariable(
        contract,
        'TRUSTED_VALIDATOR',
        hre.deployConfig.validatorPoolTrustedValidator
      )
      await assertContractVariable(
        contract,
        'REQUIRED_BOND_AMOUNT',
        hre.deployConfig.validatorPoolRequiredBondAmount
      )
      await assertContractVariable(
        contract,
        'MAX_UNBOND',
        hre.deployConfig.validatorPoolMaxUnbond
      )
      await assertContractVariable(
        contract,
        'ROUND_DURATION',
        hre.deployConfig.validatorPoolRoundDuration
      )
      await assertContractVariable(
        contract,
        'TERMINATE_OUTPUT_INDEX',
        hre.deployConfig.validatorPoolTerminateOutputIndex
      )
    },
  })
}

deployFn.tags = ['ValidatorPool', 'setup', 'l1', 'validatorSystemUpgrade']

export default deployFn
