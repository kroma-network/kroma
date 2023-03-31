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
      hre.deployConfig.colosseumChallengeTimeout,
      hre.deployConfig.colosseumSegmentsLengths.split(','),
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
        'CHALLENGE_TIMEOUT',
        hre.deployConfig.colosseumChallengeTimeout
      )
    },
  })
}

deployFn.tags = ['Colosseum', 'ZKVerifier', 'setup']

export default deployFn
