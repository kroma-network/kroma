import '@kroma-network/hardhat-deploy-config'

import assert from 'assert'

import { ethers } from 'ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { deploy, getDeploymentAddress } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const upgradeGovernorProxyAddress = await getDeploymentAddress(
    hre,
    'UpgradeGovernorProxy'
  )
  const securityCouncilTokenProxyAddress = await getDeploymentAddress(
    hre,
    'SecurityCouncilTokenProxy'
  )
  const timeLockProxyAddress = await getDeploymentAddress(hre, 'TimeLockProxy')
  const { deployer } = await hre.getNamedAccounts()

  await deploy(hre, 'UpgradeGovernor', {
    isProxyImpl: true,
    initializer: 'initialize(address,address,uint256,uint256,uint256,uint256)',
    initArgs: [
      securityCouncilTokenProxyAddress,
      timeLockProxyAddress,
      hre.deployConfig.votingDelayBlocks,
      hre.deployConfig.votingPeriodBlocks,
      hre.deployConfig.proposalThresholdTokens,
      hre.deployConfig.votesQuorumFractionPercent,
    ],
  })

  const artifact = await hre.deployments.get('UpgradeGovernor')
  const governor = new ethers.Contract(
    upgradeGovernorProxyAddress,
    artifact.abi,
    hre.ethers.provider.getSigner(deployer)
  )

  // Check variable
  assert((await governor.timelock()) === timeLockProxyAddress)
  assert((await governor.votingDelay()) === hre.deployConfig.votingDelayBlocks)
  assert(
    (await governor.votingPeriod()) === hre.deployConfig.votingPeriodBlocks
  )
  assert(
    (await governor.proposalThreshold()) ===
      hre.deployConfig.proposalThresholdTokens
  )
}

deployFn.tags = ['UpgradeGovernor', 'setup', 'l1']

export default deployFn
