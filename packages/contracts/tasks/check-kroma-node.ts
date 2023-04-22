import { KromaNodeProvider } from '@kroma-network/core-utils'
import { task, types } from 'hardhat/config'

// TODO(tynes): add in config validation
task('check-kroma-node', 'Validate the config of the kroma-node')
  .addParam(
    'kromaNodeUrl',
    'URL of the Kroma Node.',
    'http://localhost:7545',
    types.string
  )
  .setAction(async (args) => {
    const provider = new KromaNodeProvider(args.kromaNodeUrl)

    const syncStatus = await provider.syncStatus()
    console.log(JSON.stringify(syncStatus, null, 2))

    const config = await provider.rollupConfig()
    console.log(JSON.stringify(config, null, 2))
  })
