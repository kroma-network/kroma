import '@nomiclabs/hardhat-ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getDeploymentAddress,
} from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const zkProofVerifierProxyAddress = await getDeploymentAddress(
    hre,
    'ZKProofVerifierProxy'
  )
  const l2OutputOracleProxyAddress = await getDeploymentAddress(
    hre,
    'L2OutputOracleProxy'
  )
  const securityCouncilProxyAddress = await getDeploymentAddress(
    hre,
    'SecurityCouncilProxy'
  )

  await deploy(hre, 'Colosseum', {
    args: [
      l2OutputOracleProxyAddress,
      zkProofVerifierProxyAddress,
      hre.deployConfig.l2OutputOracleSubmissionInterval,
      hre.deployConfig.colosseumCreationPeriodSeconds,
      hre.deployConfig.colosseumBisectionTimeout,
      hre.deployConfig.colosseumProvingTimeout,
      hre.deployConfig.colosseumSegmentsLengths.split(','),
      securityCouncilProxyAddress,
    ],
    isProxyImpl: true,
    initArgs: [hre.deployConfig.colosseumSegmentsLengths.split(',')],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'L2_ORACLE',
        l2OutputOracleProxyAddress
      )
      await assertContractVariable(
        contract,
        'ZK_PROOF_VERIFIER',
        zkProofVerifierProxyAddress
      )
      await assertContractVariable(
        contract,
        'L2_ORACLE_SUBMISSION_INTERVAL',
        hre.deployConfig.l2OutputOracleSubmissionInterval
      )
      await assertContractVariable(
        contract,
        'CREATION_PERIOD_SECONDS',
        hre.deployConfig.colosseumCreationPeriodSeconds
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
        'SECURITY_COUNCIL',
        securityCouncilProxyAddress
      )
    },
  })
}

deployFn.tags = ['Colosseum', 'setup', 'l1', 'validatorSystemUpgrade', 'mpt']

export default deployFn
