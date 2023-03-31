import { KanvasNodeProvider } from '@wemixkanvas/core-utils'
import { task, types } from 'hardhat/config'

// TODO(tynes): add in config validation
task('check-kanvas-node', 'Validate the config of the kanvas-node')
  .addParam(
    'kanvasNodeUrl',
    'URL of the Kanvas Node.',
    'http://localhost:7545',
    types.string
  )
  .setAction(async (args) => {
    const provider = new KanvasNodeProvider(args.kanvasNodeUrl)

    const syncStatus = await provider.syncStatus()
    console.log(JSON.stringify(syncStatus, null, 2))

    const config = await provider.rollupConfig()
    console.log(JSON.stringify(config, null, 2))
  })
