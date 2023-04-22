import assert from 'assert'

import { Provider } from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'
import { sleep } from '@kroma-network/core-utils'
import '@kroma-network/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import { Contract, ethers } from 'ethers'
import { HardhatRuntimeEnvironment } from 'hardhat/types'
import 'hardhat-deploy'
import { ArtifactData } from 'hardhat-deploy/dist/types'

const IMPLEMENTATION_SLOT =
  '0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc'

interface DeployOptions {
  contract?: string | ArtifactData
  args?: any[]
  postDeployAction?: (contract: Contract) => Promise<void>
  isProxyImpl?: boolean
  initArgs?: any[]
}

/**
 * Wrapper around hardhat-deploy with some extra features.
 *
 * @param hre HardhatRuntimeEnvironment.
 * @param name Name to use for the deployment file.
 * @param opts Parameters for the deployment.
 * @param opts.contract Name of the contract to deploy.
 * @param opts.args Arguments to pass to the contract constructor.
 * @param opts.isProxyImpl Whether to update the implementation of the proxy.
 * @param opts.initArgs Arguments to pass to the proxy initializer.
 * @param opts.postDeployAction Action to perform after the contract is deployed.
 * @returns A deployed contract object.
 */
export const deploy = async (
  hre: HardhatRuntimeEnvironment,
  name: string,
  opts?: DeployOptions
): Promise<Contract | null> => {
  if (!opts) {
    opts = {}
  }

  const { deployer } = await hre.getNamedAccounts()
  const { differences, address } = await hre.deployments.fetchIfDifferent(
    name,
    {
      from: deployer,
      args: opts.args,
      contract: opts.contract,
    }
  )
  if (!differences) {
    console.log(`skipping ${name}, using existing deployment at ${address}`)
    return null
  }

  const result = await hre.deployments.deploy(name, {
    contract: opts.contract,
    from: deployer,
    args: opts.args,
    log: true,
    waitConfirmations: hre.deployConfig.numDeployConfirmations,
  })

  console.log(`deployed ${name} at ${result.address}`)
  // Always wait for the transaction to be mined, just in case.
  await hre.ethers.provider.waitForTransaction(result.transactionHash)

  // Check to make sure there is code
  const code = await hre.ethers.provider.getCode(result.address)
  if (code === '0x') {
    throw new Error(`no code for ${result.address}`)
  }

  // Create the contract object to return.
  const created = asAdvancedContract({
    confirmations: hre.deployConfig.numDeployConfirmations,
    contract: new Contract(
      result.address,
      result.abi,
      hre.ethers.provider.getSigner(deployer)
    ),
  })

  if (result.newlyDeployed && typeof opts.postDeployAction === 'function') {
    await opts.postDeployAction(created)
  }

  if (opts.isProxyImpl) {
    const proxyAdmin = await getContractFromArtifact(hre, 'ProxyAdmin', {
      signerOrProvider: hre.ethers.provider.getSigner(deployer),
    })
    const proxyName = name + 'Proxy'
    const proxy = await getContractFromArtifact(hre, proxyName)
    const hasImpl = await hasImplementation(hre, proxy.address)

    if (!opts.initArgs || hasImpl) {
      console.log(`upgrading "${proxyName}" to ${created.address}`)
      const tx = await proxyAdmin.upgrade(proxy.address, created.address)
      await hre.ethers.provider.waitForTransaction(tx.hash)
    } else {
      console.log(
        `upgrading "${proxyName}" to ${created.address} and initializing`
      )
      const tx = await proxyAdmin.upgradeAndCall(
        proxy.address,
        created.address,
        created.interface.encodeFunctionData('initialize', opts.initArgs)
      )
      await hre.ethers.provider.waitForTransaction(tx.hash)
    }
  }

  return created
}

/**
 * Deploys proxy contract using hardhat-deploy wrapper.
 * Proxy name must end with "Proxy"
 *
 * @param hre HardhatRuntimeEnvironment.
 * @param name Name to use for the proxy.
 * @param admin Address of ProxyAdmin contract.
 * @returns A deployed contract object.
 */
export const deployProxy = async (
  hre: HardhatRuntimeEnvironment,
  name: string,
  admin: string
): Promise<Contract | null> => {
  if (!name.endsWith('Proxy')) {
    throw new Error('proxy contract name must end with "Proxy"')
  }

  return deploy(hre, name, {
    contract: 'Proxy',
    args: [admin],
    postDeployAction: async (contract) => {
      await assertContractVariable(contract, 'admin', admin)
    },
  })
}

/**
 * Returns a version of the contract object which modifies all of the input contract's methods to
 * automatically await transaction receipts and confirmations. Will also throw if we timeout while
 * waiting for a transaction to be included in a block.
 *
 * @param opts Options for the contract.
 * @param opts.hre HardhatRuntimeEnvironment.
 * @param opts.contract Contract to wrap.
 * @returns Wrapped contract object.
 */
export const asAdvancedContract = (opts: {
  contract: Contract
  confirmations?: number
  gasPrice?: number
}): Contract => {
  // Temporarily override Object.defineProperty to bypass ether's object protection.
  const def = Object.defineProperty
  Object.defineProperty = (obj, propName, prop) => {
    prop.writable = true
    return def(obj, propName, prop)
  }

  const contract = new Contract(
    opts.contract.address,
    opts.contract.interface,
    opts.contract.signer || opts.contract.provider
  )

  // Now reset Object.defineProperty
  Object.defineProperty = def

  // Override each function call to also `.wait()` so as to simplify the deploy scripts' syntax.
  for (const fnName of Object.keys(contract.functions)) {
    const fn = contract[fnName].bind(contract)
    ;(contract as any)[fnName] = async (...args: any) => {
      // We want to use the configured gas price but we need to set the gas price to zero if we're
      // triggering a static function.
      let gasPrice = opts.gasPrice
      if (contract.interface.getFunction(fnName).constant) {
        gasPrice = 0
      }

      // Now actually trigger the transaction (or call).
      const tx = await fn(...args, {
        gasPrice,
      })

      // Meant for static calls, we don't need to wait for anything, we get the result right away.
      if (typeof tx !== 'object' || typeof tx.wait !== 'function') {
        return tx
      }

      // Wait for the transaction to be included in a block and wait for the specified number of
      // deployment confirmations.
      const maxTimeout = 120
      let timeout = 0
      while (true) {
        await sleep(1000)
        const receipt = await contract.provider.getTransactionReceipt(tx.hash)
        if (receipt === null) {
          timeout++
          if (timeout > maxTimeout) {
            throw new Error('timeout exceeded waiting for txn to be mined')
          }
        } else if (receipt.confirmations >= (opts.confirmations || 0)) {
          return tx
        }
      }
    }
  }

  return contract
}

/**
 * Creates a contract object from a deployed artifact.
 *
 * @param hre HardhatRuntimeEnvironment.
 * @param name Name of the deployed contract to get an object for.
 * @param opts Options for the contract.
 * @param opts.iface Optional interface to use for the contract object.
 * @param opts.signerOrProvider Optional signer or provider to use for the contract object.
 * @returns A contract object.
 */
export const getContractFromArtifact = async (
  hre: HardhatRuntimeEnvironment,
  name: string,
  opts: {
    iface?: string
    signerOrProvider?: Signer | Provider | string
  } = {}
): Promise<ethers.Contract> => {
  const artifact = await hre.deployments.get(name)

  // Get the deployed contract's interface.
  let iface = new hre.ethers.utils.Interface(artifact.abi)
  // Override with optional iface name if requested.
  if (opts.iface) {
    const factory = await hre.ethers.getContractFactory(opts.iface)
    iface = factory.interface
  }

  let signerOrProvider: Signer | Provider = hre.ethers.provider
  if (opts.signerOrProvider) {
    if (typeof opts.signerOrProvider === 'string') {
      signerOrProvider = hre.ethers.provider.getSigner(opts.signerOrProvider)
    } else {
      signerOrProvider = opts.signerOrProvider
    }
  }

  return asAdvancedContract({
    confirmations: hre.deployConfig.numDeployConfirmations,
    contract: new hre.ethers.Contract(
      artifact.address,
      iface,
      signerOrProvider
    ),
  })
}

/**
 * Helper function for asserting that a contract variable is set to the expected value.
 *
 * @param contract Contract object to query.
 * @param variable Name of the variable to query.
 * @param expected Expected value of the variable.
 */
export const assertContractVariable = async (
  contract: ethers.Contract,
  variable: string,
  expected: any
) => {
  // Need to make a copy that doesn't have a signer or we get the error that contracts with
  // signers cannot override the from address.
  const temp = new ethers.Contract(
    contract.address,
    contract.interface,
    contract.provider
  )

  const actual = await temp.callStatic[variable]({
    from: ethers.constants.AddressZero,
  })

  if (ethers.utils.isAddress(expected)) {
    assert(
      actual.toLowerCase() === expected.toLowerCase(),
      `[FATAL] ${variable} is ${actual} but should be ${expected}`
    )
    return
  }

  assert(
    actual === expected || (actual.eq && actual.eq(expected)),
    `[FATAL] ${variable} is ${actual} but should be ${expected}`
  )
}

/**
 * Returns the address for a given deployed contract by name.
 *
 * @param hre HardhatRuntimeEnvironment.
 * @param name Name of the deployed contract.
 * @returns Address of the deployed contract.
 */
export const getDeploymentAddress = async (
  hre: HardhatRuntimeEnvironment,
  name: string
): Promise<string> => {
  const deployment = await hre.deployments.get(name)
  return deployment.address
}

/**
 * Returns whether a given proxy contract has an implementation address.
 *
 * @param hre HardhatRuntimeEnvironment.
 * @param proxyAddress Address of the proxy contract.
 * @returns Whether a proxy contract has an implementation.
 */
export const hasImplementation = async (
  hre: HardhatRuntimeEnvironment,
  proxyAddress: string
): Promise<boolean> => {
  const impl = await hre.ethers.provider.getStorageAt(
    proxyAddress,
    IMPLEMENTATION_SLOT
  )
  return impl !== ethers.constants.HashZero
}
