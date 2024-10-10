import fs from 'fs'
import path from 'path'

import hardhatConfig from '../hardhat.config'

const DEPLOY_SCRIPTS_PATH = path.resolve(hardhatConfig.paths.deploy as string)
const DEPLOY_SCRIPTS_PATH_L1 = path.join(DEPLOY_SCRIPTS_PATH, 'L1')
const DEPLOY_SCRIPTS_PATH_L2 = path.join(DEPLOY_SCRIPTS_PATH, 'L2')
const SCRIPT_EXT = '.ts'
const L1_ORDERED_NAMES = [
  'ProxyAdmin',
  'Proxies',
  'ZKMerkleTrie',
  'L1CrossDomainMessenger',
  'L1StandardBridge',
  'L1ERC721Bridge',
  'KromaMintableERC20Factory',
  'ValidatorPool',
  'L2OutputOracle',
  'ZKVerifier',
  'Colosseum',
  'SecurityCouncil',
  'KromaPortal',
  'SystemConfig',
  'SecurityCouncilToken',
  'TimeLock',
  'UpgradeGovernor',
  'L1GovernanceTokenProxy',
  'L1MintManager',
  'L1GovernanceToken',
  'AssetManager',
  'ValidatorManager',
  'ZKProofVerifier',
]
const L2_ORDERED_NAMES = [
  'L1Block',
  'L2CrossDomainMessenger',
  'L2StandardBridge',
  'L2ToL1MessagePasser',
  'L2ERC721Bridge',
  'GasPriceOracle',
  'ValidatorRewardVault',
  'ProtocolVault',
  'L1FeeVault',
  'KromaMintableERC20Factory',
  'KromaMintableERC721Factory',
  'GovernanceTokenProxy',
  'MintManager',
  'GovernanceToken',
]

const main = async () => {
  await sort(DEPLOY_SCRIPTS_PATH_L1, L1_ORDERED_NAMES)
  await sort(DEPLOY_SCRIPTS_PATH_L2, L2_ORDERED_NAMES)
}

const sort = async (deployPath: string, filenames: string[]) => {
  const existsFiles = await fs.promises.readdir(deployPath)

  for (const filename of existsFiles) {
    if (filename.indexOf('-') !== 3) {
      continue
    }

    const pureName = filename.slice(4, filename.indexOf(SCRIPT_EXT))

    if (!filenames.includes(pureName)) {
      throw new Error(
        `found unexpected file: ${filename}, expected: ${pureName}`
      )
    }

    await fs.promises.rename(
      path.join(deployPath, filename),
      path.join(deployPath, pureName + SCRIPT_EXT)
    )
  }

  for (let i = 0; i < filenames.length; i++) {
    const filename = filenames[i] + SCRIPT_EXT
    await fs.promises.rename(
      path.join(deployPath, filename),
      path.join(deployPath, `${i.toString().padStart(3, '0')}-${filename}`)
    )
  }
}

main()
  .then(() => {
    console.log('Done')
  })
  .catch(console.error)
