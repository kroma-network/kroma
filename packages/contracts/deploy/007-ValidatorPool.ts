import '@kroma-network/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getDeploymentAddress,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l2OutputOracleProxyAddress = await getDeploymentAddress(
    hre,
    'L2OutputOracleProxy'
  )
  const portalProxyAddress = await getDeploymentAddress(hre, 'KromaPortalProxy')

  await deploy(hre, 'ValidatorPool', {
    args: [
      l2OutputOracleProxyAddress,
      portalProxyAddress,
      hre.deployConfig.validatorPoolTrustedValidator,
      hre.deployConfig.validatorPoolMinBondAmount,
      hre.deployConfig.validatorPoolMaxUnbond,
      hre.deployConfig.validatorPoolRoundDuration,
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
      await assertContractVariable(
        contract,
        'TRUSTED_VALIDATOR',
        hre.deployConfig.validatorPoolTrustedValidator
      )
      await assertContractVariable(
        contract,
        'MIN_BOND_AMOUNT',
        hre.deployConfig.validatorPoolMinBondAmount
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
    },
  })
}

deployFn.tags = ['ValidatorPool', 'setup', 'l1']

export default deployFn
