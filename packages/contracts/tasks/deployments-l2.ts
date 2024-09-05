import fs from 'fs'
import path from 'path'

import { task } from 'hardhat/config'

import { predeploys } from '../src'
import { getImplementation } from '../src/deploy-utils'

task(
  'deployments-l2',
  'Migrates deployments on L2 to hardhat-deploy deployments.'
).setAction(async (args, hre) => {
  const networkName = hre.network.name
  const deploymentDir = path.join(hre.config.paths.deployments, networkName)

  try {
    await fs.promises.mkdir(deploymentDir, {})
  } catch (e: any) {
    if (e.code !== 'EEXIST') {
      throw e
    }
  }

  const chainIdPath = path.join(deploymentDir, '.chainId')
  try {
    await fs.promises.access(chainIdPath)
  } catch (e: any) {
    if (e.code === 'ENOENT') {
      const chainId = await hre.getChainId()
      await fs.promises.writeFile(chainIdPath, chainId)
    } else {
      throw e
    }
  }

  const proxy = await hre.ethers.getContractFactory('Proxy')
  const proxyAbi = JSON.parse(proxy.interface.format('json') as string)

  for (const [name, proxyAddr] of Object.entries(predeploys)) {
    if (name === 'Create2Deployer') {
      continue
    }

    const proxyName = name + 'Proxy'

    const proxyDepExists = await hre.deployments.getOrNull(proxyName)
    if (!proxyDepExists) {
      const deployedBytecode = await hre.ethers.provider.getCode(proxyAddr)
      const proxyDeployment = {
        address: proxyAddr,
        abi: proxyAbi,
        bytecode: proxy.bytecode,
        deployedBytecode,
      }

      await hre.deployments.save(proxyName, proxyDeployment)
    }

    const implDepExists = await hre.deployments.getOrNull(name)
    if (!implDepExists) {
      const impl = await hre.ethers.getContractFactory(name)
      const implAddr = await getImplementation(hre, proxyAddr)
      const implAbi = JSON.parse(impl.interface.format('json') as string)
      const deployedBytecode = await hre.ethers.provider.getCode(implAddr)

      const implDeployment = {
        address: implAddr,
        abi: implAbi,
        bytecode: impl.bytecode,
        deployedBytecode,
      }

      await hre.deployments.save(name, implDeployment)
    }
  }
})
