import '@kroma/hardhat-deploy-config'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { predeploys } from '../../src'
import { assertContractVariable, deploy } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)
  await deploy(hre, 'MintManager', {
    args: [
      deployConfig.mintManagerMintActivatedBlock,
      deployConfig.mintManagerInitMintPerBlock,
      deployConfig.mintManagerSlidingWindowBlocks,
      deployConfig.mintManagerDecayingFactor,
    ],
    isProxyImpl: true,
    initArgs: [
      deployConfig.mintManagerRecipients,
      deployConfig.mintManagerShares,
    ],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'GOVERNANCE_TOKEN',
        predeploys.GovernanceToken
      )
      await assertContractVariable(
        contract,
        'MINT_ACTIVATED_BLOCK',
        deployConfig.mintManagerMintActivatedBlock
      )
      await assertContractVariable(
        contract,
        'INIT_MINT_PER_BLOCK',
        deployConfig.mintManagerInitMintPerBlock
      )
      await assertContractVariable(
        contract,
        'SLIDING_WINDOW_BLOCKS',
        deployConfig.mintManagerSlidingWindowBlocks
      )
      await assertContractVariable(
        contract,
        'DECAYING_FACTOR',
        deployConfig.mintManagerDecayingFactor
      )
    },
  })
}

deployFn.tags = ['MintManager', 'setup', 'l2', 'tge']

export default deployFn
