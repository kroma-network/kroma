import '@kroma/hardhat-deploy-config'

import assert from 'assert'

import { ethers } from 'ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { deploy, getDeploymentAddress } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)
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
      deployConfig.governorVotingDelayBlocks,
      deployConfig.l2GovernorVotingPeriodBlocks,
      deployConfig.governorProposalThreshold,
      deployConfig.governorVotesQuorumFractionPercent,
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
  assert(
    (await governor.votingDelay()).toNumber() ===
      deployConfig.governorVotingDelayBlocks
  )
  assert(
    (await governor.votingPeriod()).toNumber() ===
      deployConfig.l2GovernorVotingPeriodBlocks
  )
  assert(
    (await governor.proposalThreshold()).toNumber() ===
      deployConfig.governorProposalThreshold
  )
}

deployFn.tags = ['UpgradeGovernor', 'setup', 'l2']

export default deployFn
