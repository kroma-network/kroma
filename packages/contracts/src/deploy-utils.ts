import assert from 'assert'

import { Provider } from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'
import { sleep } from '@kroma/core-utils'
import '@kroma/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import { Contract, ethers } from 'ethers'
import { keccak256 } from 'ethers/lib/utils'
import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { ArtifactData } from 'hardhat-deploy/dist/types'
import 'hardhat-deploy'

import { predeploys } from './constants'

const PROXY_IMPLEMENTATION_SLOT =
  '0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc'
const PROXY_OWNER_SLOT =
  '0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103'

interface DeployOptions {
  contract?: string | ArtifactData
  args?: any[]
  postDeployAction?: (contract: Contract) => Promise<void>
  isProxyImpl?: boolean
  initArgs?: any[]
  initializer?: string
}

/**
 * Deploys implementation contract and upgrades proxy to the deployed implementation contract.
 * Upgrade is processed via ProxyAdmin contract.
 *
 * @param hre HardhatRuntimeEnvironment.
 * @param name Name to use for the deployment file.
 * @param opts Parameters for the deployment.
 * @param opts.isProxyImpl Whether to update the implementation of the proxy.
 * @param opts.initArgs Arguments to pass to the proxy initializer.
 * @returns A deployed contract object.
 */
export const deploy = async (
  hre: HardhatRuntimeEnvironment,
  name: string,
  opts: DeployOptions = {}
): Promise<Contract | null> => {
  const [created, newlyDeployed] = await deployImpl(hre, name, opts)
  if (!newlyDeployed) {
    return created
  }

  if (opts.isProxyImpl) {
    const { deployer } = await hre.getNamedAccounts()
    const proxyName = name + 'Proxy'
    const proxy = await getContractFromArtifact(hre, proxyName, {
      signerOrProvider: deployer,
    })
    const hasImpl = await hasImplementation(hre, proxy.address)
    const admin = await getProxyAdmin(hre, proxy.address)

    let proxyAdmin = await hre.ethers.getContractAt('ProxyAdmin', admin)
    const proxyOwner = await proxyAdmin.owner()
    proxyAdmin = proxyAdmin.connect(hre.ethers.provider.getSigner(proxyOwner))

    if (!opts.initArgs || hasImpl) {
      console.log(`upgrading "${proxyName}" to ${created.address}`)
      const tx = await proxyAdmin.upgrade(proxy.address, created.address)
      await hre.ethers.provider.waitForTransaction(tx.hash)
    } else {
      console.log(
        `upgrading "${proxyName}" to ${created.address} and initializing`
      )

      if (!opts.initializer) {
        opts.initializer = 'initialize'
      }

      // Ensure that the contract has the initialize function.
      try {
        created.interface.getFunction(opts.initializer)
      } catch (error) {
        throw new Error(
          `deployed "${name}" does not have the function "${opts.initializer}"`
        )
      }

      const tx = await proxyAdmin.upgradeAndCall(
        proxy.address,
        created.address,
        created.interface.encodeFunctionData(opts.initializer, opts.initArgs)
      )
      await hre.ethers.provider.waitForTransaction(tx.hash)
    }
  }

  return created
}

/**
 * Deploys implementation contract and upgrades proxy to the deployed implementation contract.
 * Only used when proxy should be upgraded by deployer.
 *
 * @param hre HardhatRuntimeEnvironment.
 * @param name Name to use for the deployment file.
 * @param opts Parameters for the deployment.
 * @param opts.initArgs Arguments to pass to the proxy initializer.
 * @returns A deployed contract object.
 */
export const deployAndUpgradeByDeployer = async (
  hre: HardhatRuntimeEnvironment,
  name: string,
  opts: DeployOptions = {}
): Promise<Contract | null> => {
  const [created, newlyDeployed] = await deployImpl(hre, name, opts)
  if (!newlyDeployed) {
    return created
  }

  const { deployer } = await hre.getNamedAccounts()
  const proxyName = name + 'Proxy'
  const proxy = await getContractFromArtifact(hre, proxyName, {
    signerOrProvider: deployer,
  })
  const hasImpl = await hasImplementation(hre, proxy.address)

  if (!opts.initArgs || hasImpl) {
    console.log(`upgrading "${proxyName}" to ${created.address}`)
    const tx = await proxy.upgradeTo(created.address)
    await hre.ethers.provider.waitForTransaction(tx.hash)
  } else {
    console.log(
      `upgrading "${proxyName}" to ${created.address} and initializing`
    )

    if (!opts.initializer) {
      opts.initializer = 'initialize'
    }

    // Ensure that the contract has the initialize function.
    try {
      created.interface.getFunction(opts.initializer)
    } catch (error) {
      throw new Error(
        `deployed "${name}" does not have the function "${opts.initializer}"`
      )
    }

    const tx = await proxy.upgradeToAndCall(
      created.address,
      created.interface.encodeFunctionData(opts.initializer, opts.initArgs)
    )
    await hre.ethers.provider.waitForTransaction(tx.hash)
  }

  return created
}

/**
 * Wrapper around hardhat-deploy with some extra features.
 *
 * @param hre HardhatRuntimeEnvironment.
 * @param name Name to use for the deployment file.
 * @param opts Parameters for the deployment.
 * @param opts.contract Name of the contract to deploy.
 * @param opts.args Arguments to pass to the contract constructor.
 * @param opts.postDeployAction Action to perform after the contract is deployed.
 * @returns A deployed contract object.
 * @returns If the contract is newly deployed or not.
 */
const deployImpl = async (
  hre: HardhatRuntimeEnvironment,
  name: string,
  opts: DeployOptions = {}
): Promise<[Contract | null, boolean]> => {
  const { deployer } = await hre.getNamedAccounts()

  // Wrap in a try/catch in case there is not a deployConfig for the current network.
  let numDeployConfirmations: number
  try {
    numDeployConfirmations = hre.deployConfig.numDeployConfirmations
  } catch (e) {
    numDeployConfirmations = 1
  }

  const result = await hre.deployments.deploy(name, {
    contract: opts.contract,
    from: deployer,
    args: opts.args,
    log: true,
    waitConfirmations: numDeployConfirmations,
  })

  // Create the contract object to return.
  const created = asAdvancedContract({
    confirmations: numDeployConfirmations,
    contract: new Contract(
      result.address,
      result.abi,
      hre.ethers.provider.getSigner(deployer)
    ),
  })

  // If the contract is not newly deployed, do not proceed further.
  if (!result.newlyDeployed) {
    return [created, false]
  }

  // Always wait for the transaction to be mined, just in case.
  await hre.ethers.provider.waitForTransaction(result.transactionHash)
  console.log(`deployed "${name}" at ${result.address}`)

  // Check to make sure there is code
  const code = await hre.ethers.provider.getCode(result.address)
  if (code === '0x') {
    throw new Error(`no code for ${result.address}`)
  }

  if (typeof opts.postDeployAction === 'function') {
    await opts.postDeployAction(created)
  }

  return [created, true]
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
 * Deploys proxy contract to deterministic address using CREATE2.
 * Proxy name must end with "Proxy"
 *
 * @param hre HardhatRuntimeEnvironment.
 * @param name Name to use for the proxy.
 * @param admin Admin address of the proxy.
 * @param salt Salt to determine the deployment address.
 * @returns A deployed contract object.
 */
export const deployDeterministicProxy = async (
  hre: HardhatRuntimeEnvironment,
  name: string,
  admin: string,
  salt: string
): Promise<null> => {
  if (!name.endsWith('Proxy')) {
    throw new Error('proxy contract name must end with "Proxy"')
  }

  // Wrap in a try/catch in case there is not a deployConfig for the current network.
  let numDeployConfirmations: number
  try {
    numDeployConfirmations = hre.deployConfig.numDeployConfirmations
  } catch (e) {
    numDeployConfirmations = 1
  }

  // Calculate the address of proxy using deployer address, salt, initCode.
  const proxy = await hre.ethers.getContractFactory('Proxy')
  const simulateTx = proxy.getDeployTransaction(admin)

  const create2Inputs = [
    '0xff',
    predeploys.Create2Deployer,
    salt,
    keccak256(simulateTx.data),
  ].map((i) => (i.startsWith('0x') ? i : `0x${i}`))
  const create2Input = '0x' + create2Inputs.map((i) => i.slice(2)).join('')

  const create2Hash = keccak256(create2Input)
  const create2Address = hre.ethers.utils.getAddress(
    `0x${create2Hash.slice(-40)}`
  )

  // Ensure there is not code at the address.
  let code = await hre.ethers.provider.getCode(create2Address)
  if (code !== '0x') {
    throw new Error(
      `existing contract code found at ${create2Address}. Use a different salt or verify the intended deployment address.`
    )
  }

  // Ensure there is code at the Create2Deployer address.
  code = await hre.ethers.provider.getCode(predeploys.Create2Deployer)
  if (code === '0x') {
    throw new Error(`no code at ${predeploys.Create2Deployer}`)
  }

  const { deployer } = await hre.getNamedAccounts()
  const signer = hre.ethers.provider.getSigner(deployer)

  const create2DeployerAbi = [
    'function deploy(uint256 value,bytes32 salt,bytes memory code) public',
  ]
  let create2Deployer = new hre.ethers.Contract(
    predeploys.Create2Deployer,
    create2DeployerAbi,
    hre.ethers.provider
  )
  create2Deployer = create2Deployer.connect(signer)

  // Call deploy function of Create2Deployer contract.
  const deployTx = await create2Deployer.deploy(0, salt, simulateTx.data)
  await deployTx.wait(numDeployConfirmations)
  console.log(`deployed "${name}" at ${create2Address}`)

  // Save the deployment.
  const proxyAbi = JSON.parse(proxy.interface.format('json') as string)
  const proxyBuildInfo = await hre.artifacts.getBuildInfo('Proxy.sol:Proxy')
  const proxyArtifact = await hre.artifacts.readArtifact('Proxy')
  const proxyCompiledOutput: any =
    proxyBuildInfo.output.contracts[proxyArtifact.sourceName]['Proxy']
  let metadata: string
  try {
    metadata = JSON.stringify(proxyCompiledOutput.metadata)
  } catch (error) {
    console.log(
      `compiled output of Proxy contract does not have metadata field: ${error}`
    )
    metadata = ''
  }
  const deployedBytecode = await hre.ethers.provider.getCode(create2Address)
  const deployTxReceipt = await hre.ethers.provider.getTransactionReceipt(
    deployTx.hash
  )

  const proxyDeployment = {
    address: create2Address,
    abi: proxyAbi,
    transactionHash: deployTx.hash,
    receipt: deployTxReceipt,
    args: [admin],
    metadata,
    bytecode: proxy.bytecode,
    deployedBytecode,
  }
  await hre.deployments.save(name, proxyDeployment)

  return null
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

  let numDeployConfirmations: number
  try {
    numDeployConfirmations = hre.deployConfig.numDeployConfirmations
  } catch (e) {
    numDeployConfirmations = 1
  }

  return asAdvancedContract({
    confirmations: numDeployConfirmations,
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
 * Returns the implementation address for a given proxy address.
 *
 * @param hre HardhatRuntimeEnvironment.
 * @param proxyAddress Address of the proxy contract.
 * @returns Address of the implementation.
 */
export const getImplementation = async (
  hre: HardhatRuntimeEnvironment,
  proxyAddress: string
): Promise<string> => {
  const slotValue = await hre.ethers.provider.getStorageAt(
    proxyAddress,
    PROXY_IMPLEMENTATION_SLOT
  )
  return ethers.utils.getAddress(ethers.utils.hexDataSlice(slotValue, 12))
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
  const impl = await getImplementation(hre, proxyAddress)
  return impl !== ethers.constants.AddressZero
}

/**
 * Returns the admin address for a given proxy address.
 *
 * @param hre HardhatRuntimeEnvironment.
 * @param proxyAddress Address of the proxy contract.
 * @returns Address of the proxy admin.
 */
export const getProxyAdmin = async (
  hre: HardhatRuntimeEnvironment,
  proxyAddress: string
): Promise<string> => {
  const slotValue = await hre.ethers.provider.getStorageAt(
    proxyAddress,
    PROXY_OWNER_SLOT
  )
  return ethers.utils.getAddress(ethers.utils.hexDataSlice(slotValue, 12))
}
