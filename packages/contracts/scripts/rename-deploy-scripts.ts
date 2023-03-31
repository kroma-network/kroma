import path from 'path'
import fs from 'fs'

// @ts-ignore
import hardhatConfig from '../hardhat.config'

const DEPLOY_SCRIPTS_PATH = path.resolve(hardhatConfig.paths.deploy as string)
const SCRIPT_EXT = '.ts'
const ORDERED_FILE_NAMES = [
  'ProxyAdmin',
  'Proxies',
  'ZKMerkleTrie',
  'L1CrossDomainMessenger',
  'L1StandardBridge',
  'L1ERC721Bridge',
  'KanvasMintableERC20Factory',
  'L2OutputOracle',
  'Colosseum',
  'KanvasPortal',
  'SystemConfig',
]

const main = async () => {
  const existsFiles = await fs.promises.readdir(DEPLOY_SCRIPTS_PATH)

  for (const filename of existsFiles) {
    if (filename.indexOf('-') !== 3) {
      continue
    }

    const pureName = filename.slice(4, filename.indexOf(SCRIPT_EXT))

    if (!ORDERED_FILE_NAMES.includes(pureName)) {
      throw new Error(
        `found unexpected file: ${filename}, expected: ${pureName}`
      )
    }

    await fs.promises.rename(
      path.join(DEPLOY_SCRIPTS_PATH, filename),
      path.join(DEPLOY_SCRIPTS_PATH, pureName + SCRIPT_EXT)
    )
  }

  for (let i = 0; i < ORDERED_FILE_NAMES.length; i++) {
    const filename = ORDERED_FILE_NAMES[i] + SCRIPT_EXT
    await fs.promises.rename(
      path.join(DEPLOY_SCRIPTS_PATH, filename),
      path.join(
        DEPLOY_SCRIPTS_PATH,
        `${i.toString().padStart(3, '0')}-${filename}`
      )
    )
  }
}

main()
  .then(() => {
    console.log('Done')
  })
  .catch(console.error)
