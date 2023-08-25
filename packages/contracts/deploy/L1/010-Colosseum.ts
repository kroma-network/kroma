import '@nomiclabs/hardhat-ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getDeploymentAddress,
} from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const zkVerifierProxyAddress = await getDeploymentAddress(hre, 'ZKVerifier')
  const l2OutputOracleProxyAddress = await getDeploymentAddress(
    hre,
    'L2OutputOracleProxy'
  )
  const securityCouncilProxyAddress = await getDeploymentAddress(
    hre,
    'SecurityCouncilProxy'
  )
  const zkMerkleTrieAddress = await getDeploymentAddress(hre, 'ZKMerkleTrie')

  await deploy(hre, 'Colosseum', {
    args: [
      l2OutputOracleProxyAddress,
      zkVerifierProxyAddress,
      hre.deployConfig.l2OutputOracleSubmissionInterval,
      hre.deployConfig.colosseumCreationPeriodSeconds,
      hre.deployConfig.colosseumBisectionTimeout,
      hre.deployConfig.colosseumProvingTimeout,
      hre.deployConfig.colosseumDummyHash,
      hre.deployConfig.colosseumMaxTxs,
      hre.deployConfig.colosseumSegmentsLengths.split(','),
      securityCouncilProxyAddress,
      zkMerkleTrieAddress,
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
        'ZK_VERIFIER',
        zkVerifierProxyAddress
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
        'SECURITY_COUNCIL',
        securityCouncilProxyAddress
      )
      await assertContractVariable(
        contract,
        'ZK_MERKLE_TRIE',
        zkMerkleTrieAddress
      )
    },
  })
}

deployFn.tags = ['Colosseum', 'setup', 'l1']

export default deployFn
