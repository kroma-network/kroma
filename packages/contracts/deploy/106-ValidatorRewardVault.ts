import '@kroma-network/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getDeploymentAddress,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)

  const validatorPoolProxyAddress = await getDeploymentAddress(
    hre,
    'ValidatorPoolProxy'
  )

  const submissionInterval = deployConfig.submissionInterval
  const l2BlockTime = deployConfig.l2BlockTime
  const finalizationPeriodSeconds = deployConfig.finalizationPeriodSeconds

  const rewardDivider =
    finalizationPeriodSeconds / (submissionInterval * l2BlockTime)
  if (rewardDivider < 1) {
    throw new Error('invalid reward divider value')
  }

  await deploy(hre, 'ValidatorRewardVault', {
    args: [validatorPoolProxyAddress, rewardDivider],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'VALIDATOR_POOL',
        validatorPoolProxyAddress
      )
      await assertContractVariable(contract, 'REWARD_DIVIDER', rewardDivider)
    },
  })
}

deployFn.tags = ['ValidatorRewardVault', 'l2']

export default deployFn
