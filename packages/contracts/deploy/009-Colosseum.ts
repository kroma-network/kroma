import '@nomiclabs/hardhat-ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getDeploymentAddress,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  await deploy(hre, 'ZKVerifier')

  const zkVerifierAddress = await getDeploymentAddress(hre, 'ZKVerifier')
  const l2OutputOracleProxyAddress = await getDeploymentAddress(
    hre,
    'L2OutputOracleProxy'
  )

  await deploy(hre, 'Colosseum', {
    args: [
      l2OutputOracleProxyAddress,
      zkVerifierAddress,
      hre.deployConfig.l2OutputOracleSubmissionInterval,
      hre.deployConfig.colosseumBisectionTimeout,
      hre.deployConfig.colosseumProvingTimeout,
      hre.deployConfig.l2ChainID,
      hre.deployConfig.colosseumDummyHash,
      hre.deployConfig.colosseumMaxTxs,
      hre.deployConfig.colosseumSegmentsLengths.split(','),
      hre.deployConfig.portalGuardian,
    ],
    isProxyImpl: true,
    initArgs: [hre.deployConfig.colosseumSegmentsLengths.split(',')],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'L2_ORACLE',
        l2OutputOracleProxyAddress
      )
      await assertContractVariable(contract, 'ZK_VERIFIER', zkVerifierAddress)
      await assertContractVariable(
        contract,
        'L2_ORACLE_SUBMISSION_INTERVAL',
        hre.deployConfig.l2OutputOracleSubmissionInterval
      )
      await assertContractVariable(
        contract,
        'BISECTION_TIMEOUT',
        hre.deployConfig.colosseumBisectionTimeout
      )
      await assertContractVariable(
        contract,
        'PROVING_TIMEOUT',
        hre.deployConfig.colosseumProvingTimeout
      )
      await assertContractVariable(
        contract,
        'CHAIN_ID',
        hre.deployConfig.l2ChainID
      )
      await assertContractVariable(
        contract,
        'DUMMY_HASH',
        hre.deployConfig.colosseumDummyHash
      )
      await assertContractVariable(
        contract,
        'MAX_TXS',
        hre.deployConfig.colosseumMaxTxs
      )
      await assertContractVariable(
        contract,
        'GUARDIAN',
        hre.deployConfig.portalGuardian
      )
    },
  })
}

deployFn.tags = ['Colosseum', 'ZKVerifier', 'setup', 'l1']

export default deployFn
