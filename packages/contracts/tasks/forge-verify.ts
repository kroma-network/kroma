import { spawn as spawn } from 'child_process'

import * as foundryup from '@foundry-rs/easy-foundryup'
import { ethers } from 'ethers'
import { task, types } from 'hardhat/config'
import 'hardhat-deploy'

interface ForgeVerifyArgs {
  chainId: string
  compilerVersion: string
  constructorArgs: string
  optimizerRuns: number
  contractAddress: string
  contractName: string
  etherscanApiKey: string
  verifier: string
  verifierUrl: string
  root: string
}

const verifyArgs = (opts: ForgeVerifyArgs): string[] => {
  const allArgs: string[] = []

  if (!opts.etherscanApiKey) {
    throw new Error('No Etherscan API key provided')
  }
  allArgs.push('--etherscan-api-key', opts.etherscanApiKey)
  if (!opts.chainId) {
    throw new Error(`No chain-id provided`)
  }
  allArgs.push(`--chain`, opts.chainId)

  if (opts.compilerVersion) {
    allArgs.push('--compiler-version', opts.compilerVersion)
  }
  if (opts.constructorArgs) {
    allArgs.push('--constructor-args', opts.constructorArgs)
  }
  if (typeof opts.optimizerRuns === 'number') {
    allArgs.push('--num-of-optimizations', opts.optimizerRuns.toString())
  }
  if (opts.verifier) {
    allArgs.push('--verifier', opts.verifier)
  }
  if (opts.verifierUrl) {
    allArgs.push('--verifier-url', opts.verifierUrl)
  }
  if (opts.root) {
    allArgs.push('--root', opts.root)
  }
  allArgs.push('--watch')

  if (!opts.contractAddress) {
    throw new Error('No contract address provided')
  }
  allArgs.push(opts.contractAddress)
  if (!opts.contractName) {
    throw new Error('No contract name provided')
  }
  allArgs.push(opts.contractName)
  return allArgs
}

const spawnVerify = async (opts: ForgeVerifyArgs): Promise<boolean> => {
  const args = ['verify-contract', ...verifyArgs(opts)]
  const forgeCmd = await foundryup.getForgeCommand()
  return new Promise((resolve) => {
    const process = spawn(forgeCmd, args, {
      stdio: 'inherit',
    })
    process.on('exit', (code) => {
      resolve(code === 0)
    })
  })
}

task('forge-contract-verify', 'Verify contracts using forge')
  .addOptionalParam(
    'contract',
    'Name of the contract to verify',
    '',
    types.string
  )
  .addOptionalParam(
    'etherscanApiKey',
    'Etherscan API key',
    process.env.ETHERSCAN_API_KEY,
    types.string
  )
  .addOptionalParam(
    'verifier',
    'Verification provider to use (etherscan, sourcify, blockscout)',
    'etherscan',
    types.string
  )
  .addOptionalParam('verifierUrl', 'Verifier URL', '', types.string)
  .addOptionalParam('root', "The project's root path", '', types.string)
  .setAction(async (args, hre) => {
    const deployments = await hre.deployments.all()
    if (args.contract !== '') {
      if (!deployments[args.contract]) {
        throw new Error(
          `Contract ${args.contract} not found in ${hre.network} deployments`
        )
      }
    }

    if (args.root === '') {
      args.root = hre.config.paths.root
    }

    for (const [contract, deployment] of Object.entries(deployments)) {
      if (args.contract !== '' && args.contract !== contract) {
        continue
      }

      const chainId = await hre.getChainId()
      const contractAddress = deployment.address
      const etherscanApiKey = args.etherscanApiKey
      const verifier = args.verifier
      const verifierUrl = args.verifierUrl
      const root = args.root

      let metadata = deployment.metadata as any
      // Handle double nested JSON stringify
      while (typeof metadata === 'string') {
        metadata = JSON.parse(metadata) as any
      }

      // Skip if metadata is undefined or null
      if (!metadata) {
        continue
      }

      const contractName = Object.entries(
        metadata.settings.compilationTarget
      )[0].join(':')
      const compilerVersion = metadata.compiler.version

      const iface = new ethers.utils.Interface(deployment.abi)
      const constructorArgs = iface.encodeDeploy(deployment.args)
      const optimizerRuns = metadata.settings.optimizer

      const success = await spawnVerify({
        chainId,
        compilerVersion,
        constructorArgs,
        optimizerRuns,
        contractAddress,
        contractName,
        etherscanApiKey,
        verifier,
        verifierUrl,
        root,
      })

      if (success) {
        console.log(`Contract verification successful for ${contractName}`)
      } else {
        console.log(`Contract verification unsuccessful for ${contractName}`)
      }
    }
  })
