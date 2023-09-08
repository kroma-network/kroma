import '@kroma/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { assertContractVariable, deploy } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)

  const validatorPoolProxy = await hre.companionNetworks['l1'].deployments.get(
    'ValidatorPoolProxy'
  )

  const submissionInterval = deployConfig.l2OutputOracleSubmissionInterval
  const l2BlockTime = deployConfig.l2BlockTime
  const finalizationPeriodSeconds = deployConfig.finalizationPeriodSeconds

  const rewardDivider =
    finalizationPeriodSeconds / (submissionInterval * l2BlockTime)
  if (rewardDivider < 1) {
    throw new Error('invalid reward divider value')
  }

  await deploy(hre, 'ValidatorRewardVault', {
    args: [validatorPoolProxy.address, rewardDivider],
    isProxyImpl: true,
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'VALIDATOR_POOL',
        validatorPoolProxy.address
      )
      await assertContractVariable(contract, 'REWARD_DIVIDER', rewardDivider)
    },
  })
}

deployFn.tags = ['ValidatorRewardVault', 'l2']

export default deployFn
