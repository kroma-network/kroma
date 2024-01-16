import fs from 'fs'
import path from 'path'

import { task } from 'hardhat/config'
import { HardhatRuntimeEnvironment } from 'hardhat/types'

// This task exports contract addresses by name as JSON object,
// which is the same as the output of the OP's forge script execution.

task(
  'export-addresses',
  'Export deployed contract address by name as JSON object'
).setAction(async (_, hre: HardhatRuntimeEnvironment) => {
  const deploymentDir = path.join(
    hre.config.paths.deployments,
    hre.network.name
  )
  const files = await fs.promises.readdir(deploymentDir)

  const exportData = {}
  for (const file of files) {
    if (file.endsWith('.json')) {
      const data = await fs.promises.readFile(
        path.join(deploymentDir, file),
        'utf8'
      )
      const deployment = JSON.parse(data)
      const name = file.slice(0, -5)
      exportData[name] = deployment.address
    }
  }

  await fs.promises.writeFile(
    path.join(deploymentDir, '.deploy'),
    JSON.stringify(exportData, null, 4)
  )
})
