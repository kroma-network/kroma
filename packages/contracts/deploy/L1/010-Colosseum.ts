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
      securityCouncilProxyAddress,
      hre.deployConfig.colosseumGuardianPeriodSeconds,
      hre.deployConfig.colosseumMaxClockDurationSeconds,
      hre.deployConfig.colosseumChallengeGracePeriodSeconds,
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
        'SECURITY_COUNCIL',
        securityCouncilProxyAddress
      )
      await assertContractVariable(
        contract,
        'GUARDIAN_PERIOD',
        hre.deployConfig.colosseumGuardianPeriodSeconds
      )
      await assertContractVariable(
        contract,
        'MAX_CLOCK_DURATION_SECONDS',
        hre.deployConfig.colosseumMaxClockDurationSeconds
      )
      await assertContractVariable(
        contract,
        'CHALLENGE_GRACE_PERIOD',
        hre.deployConfig.colosseumChallengeGracePeriodSeconds
      )
    },
  })
}

deployFn.tags = ['Colosseum', 'setup', 'l1', 'validatorSystemUpgrade', 'mpt']

export default deployFn
