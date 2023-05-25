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

  await deploy(hre, 'ValidatorPool', {
    args: [
      l2OutputOracleProxyAddress,
      hre.deployConfig.validatorPoolTrustedValidator,
      hre.deployConfig.validatorPoolMinBondAmount,
    ],
    isProxyImpl: true,
    initArgs: [],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'L2_ORACLE',
        l2OutputOracleProxyAddress
      )
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
    },
  })
}

deployFn.tags = ['ValidatorPool', 'setup', 'l1']

export default deployFn
