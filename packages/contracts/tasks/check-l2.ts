import assert from 'assert'

import '@nomiclabs/hardhat-ethers'
import { Contract, Signer, Wallet, providers } from 'ethers'
import { task, types } from 'hardhat/config'
import { HardhatRuntimeEnvironment } from 'hardhat/types'
import 'hardhat-deploy'

import { predeploys } from '../src'

const expectedSemver = '0.1.0'
const implSlot =
  '0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc'
const adminSlot =
  '0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103'
const prefix = '0x420000000000000000000000000000000000'

const logLoud = () => {
  console.log('   !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!')
}

const yell = (msg: string) => {
  logLoud()
  console.log(msg)
  logLoud()
}

// checkPredeploys will ensure that all of the predeploys are set
const checkPredeploys = async (
  hre: HardhatRuntimeEnvironment,
  provider: providers.Provider
) => {
  console.log('Checking predeploys are configured correctly')
  for (let i = 0; i < 256; i++) {
    const num = hre.ethers.utils.hexZeroPad('0x' + i.toString(16), 2)
    const addr = hre.ethers.utils.getAddress(
      hre.ethers.utils.hexConcat([prefix, num])
    )

    const code = await provider.getCode(addr)
    if (code === '0x') {
      throw new Error(`no code found at ${addr}`)
    }

    if (addr === predeploys.ProxyAdmin || addr === predeploys.WETH9) {
      continue
    }

    const slot = await provider.getStorageAt(addr, adminSlot)
    const admin = hre.ethers.utils.hexConcat([
      '0x000000000000000000000000',
      predeploys.ProxyAdmin,
    ])

    if (admin !== slot) {
      throw new Error(`incorrect admin slot in ${addr}`)
    }

    if (i % 200 === 0) {
      console.log(`Checked through ${addr}`)
    }
  }
}

// assertSemver will ensure that the semver is the correct version
const assertSemver = async (
  contract: Contract,
  name: string,
  override?: string
) => {
  const version = await contract.version()
  let target = expectedSemver
  if (override) {
    target = override
  }
  if (version !== target) {
    throw new Error(
      `${name}: version mismatch. Got ${version}, expected ${target}`
    )
  }
  console.log(`  - version: ${version}`)
}

// checkProxy will print out the proxy slots
const checkProxy = async (
  _hre: HardhatRuntimeEnvironment,
  name: string,
  provider: providers.Provider
) => {
  const address = predeploys[name]
  if (!address) {
    throw new Error(`unknown contract name: ${name}`)
  }

  const impl = await provider.getStorageAt(address, implSlot)
  const admin = await provider.getStorageAt(address, adminSlot)

  console.log(`  - EIP-1967 implementation slot: ${impl}`)
  console.log(`  - EIP-1967 admin slot: ${admin}`)
}

// assertProxy will require the proxy is set
const assertProxy = async (
  hre: HardhatRuntimeEnvironment,
  name: string,
  provider: providers.Provider
) => {
  const address = predeploys[name]
  if (!address) {
    throw new Error(`unknown contract name: ${name}`)
  }

  const code = await provider.getCode(address)
  const deployInfo = await hre.artifacts.readArtifact('Proxy')

  if (code !== deployInfo.deployedBytecode) {
    throw new Error(`${address}: code mismatch`)
  }

  const impl = await provider.getStorageAt(address, implSlot)
  const implAddress = '0x' + impl.slice(26)
  const implCode = await provider.getCode(implAddress)
  if (implCode === '0x') {
    throw new Error('No code at implementation')
  }
}

const check = {
  // L2CrossDomainMessenger
  // - check version
  // - check OTHER_MESSENGER
  // - is behind a proxy
  // - check owner
  // - check initialized
  L2CrossDomainMessenger: async (
    hre: HardhatRuntimeEnvironment,
    signer: Signer
  ) => {
    const L2CrossDomainMessenger = await hre.ethers.getContractAt(
      'L2CrossDomainMessenger',
      predeploys.L2CrossDomainMessenger,
      signer
    )

    await assertSemver(L2CrossDomainMessenger, 'L2CrossDomainMessenger')

    const xDomainMessageSenderSlot = await signer.provider.getStorageAt(
      predeploys.L2CrossDomainMessenger,
      204
    )

    const xDomainMessageSender = '0x' + xDomainMessageSenderSlot.slice(26)
    assert(
      xDomainMessageSender === '0x000000000000000000000000000000000000dead'
    )

    const otherMessenger = await L2CrossDomainMessenger.OTHER_MESSENGER()
    assert(otherMessenger !== hre.ethers.constants.AddressZero)
    yell(`  - OTHER_MESSENGER: ${otherMessenger}`)

    await checkProxy(hre, 'L2CrossDomainMessenger', signer.provider)
    await assertProxy(hre, 'L2CrossDomainMessenger', signer.provider)

    const owner = await L2CrossDomainMessenger.owner()
    assert(owner !== hre.ethers.constants.AddressZero)
    yell(`  - owner: ${owner}`)

    const MESSAGE_VERSION = await L2CrossDomainMessenger.MESSAGE_VERSION()
    console.log(`  - MESSAGE_VERSION: ${MESSAGE_VERSION}`)
    const MIN_GAS_CALLDATA_OVERHEAD =
      await L2CrossDomainMessenger.MIN_GAS_CALLDATA_OVERHEAD()
    console.log(`  - MIN_GAS_CALLDATA_OVERHEAD: ${MIN_GAS_CALLDATA_OVERHEAD}`)
    const MIN_GAS_CONSTANT_OVERHEAD =
      await L2CrossDomainMessenger.MIN_GAS_CONSTANT_OVERHEAD()
    console.log(`  - MIN_GAS_CONSTANT_OVERHEAD: ${MIN_GAS_CONSTANT_OVERHEAD}`)
    const MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR =
      await L2CrossDomainMessenger.MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR()
    console.log(
      `  - MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR: ${MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR}`
    )
    const MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR =
      await L2CrossDomainMessenger.MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR()
    console.log(
      `  - MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR: ${MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR}`
    )

    const slot = await signer.provider.getStorageAt(
      predeploys.L2CrossDomainMessenger,
      0
    )

    const spacer = '0x' + slot.slice(26)
    console.log(`  - legacy spacer: ${spacer}`)

    const initialized = '0x' + slot.slice(24, 26)
    assert(initialized === '0x01')
    console.log(`  - initialized: ${initialized}`)
  },
  // GasPriceOracle
  // - check version
  // - check decimals
  GasPriceOracle: async (hre: HardhatRuntimeEnvironment, signer: Signer) => {
    const GasPriceOracle = await hre.ethers.getContractAt(
      'GasPriceOracle',
      predeploys.GasPriceOracle,
      signer
    )

    await assertSemver(GasPriceOracle, 'GasPriceOracle')

    const decimals = await GasPriceOracle.decimals()
    assert(decimals.eq(6))
    console.log(`  - decimals: ${decimals.toNumber()}`)

    await checkProxy(hre, 'GasPriceOracle', signer.provider)
    await assertProxy(hre, 'GasPriceOracle', signer.provider)
  },
  // L2StandardBridge
  // - check version
  L2StandardBridge: async (hre: HardhatRuntimeEnvironment, signer: Signer) => {
    const L2StandardBridge = await hre.ethers.getContractAt(
      'L2StandardBridge',
      predeploys.L2StandardBridge,
      signer
    )

    await assertSemver(L2StandardBridge, 'L2StandardBridge', '1.1.0')

    const OTHER_BRIDGE = await L2StandardBridge.OTHER_BRIDGE()
    assert(OTHER_BRIDGE !== hre.ethers.constants.AddressZero)
    yell(`  - OTHER_BRIDGE: ${OTHER_BRIDGE}`)

    const MESSENGER = await L2StandardBridge.MESSENGER()
    assert(MESSENGER === predeploys.L2CrossDomainMessenger)

    await checkProxy(hre, 'L2StandardBridge', signer.provider)
    await assertProxy(hre, 'L2StandardBridge', signer.provider)
  },
  // ValidatorRewardVault
  // - check version
  // - check RECIPIENT
  ValidatorRewardVault: async (hre: HardhatRuntimeEnvironment, signer: Signer) => {
    const ValidatorRewardVault = await hre.ethers.getContractAt(
      'ValidatorRewardVault',
      predeploys.ValidatorRewardVault,
      signer
    )

    await assertSemver(ValidatorRewardVault, 'ValidatorRewardVault')

    const RECIPIENT = await ValidatorRewardVault.RECIPIENT()
    assert(RECIPIENT !== hre.ethers.constants.AddressZero)
    yell(`  - RECIPIENT: ${RECIPIENT}`)

    const MIN_WITHDRAWAL_AMOUNT = await ValidatorRewardVault.MIN_WITHDRAWAL_AMOUNT()
    console.log(`  - MIN_WITHDRAWAL_AMOUNT: ${MIN_WITHDRAWAL_AMOUNT}`)

    await checkProxy(hre, 'ValidatorRewardVault', signer.provider)
    await assertProxy(hre, 'ValidatorRewardVault', signer.provider)
  },
  // KromaMintableERC20Factory
  // - check version
  KromaMintableERC20Factory: async (
    hre: HardhatRuntimeEnvironment,
    signer: Signer
  ) => {
    const KromaMintableERC20Factory = await hre.ethers.getContractAt(
      'KromaMintableERC20Factory',
      predeploys.KromaMintableERC20Factory,
      signer
    )

    await assertSemver(KromaMintableERC20Factory, 'KromaMintableERC20Factory')

    const BRIDGE = await KromaMintableERC20Factory.BRIDGE()
    assert(BRIDGE !== hre.ethers.constants.AddressZero)

    await checkProxy(hre, 'KromaMintableERC20Factory', signer.provider)
    await assertProxy(hre, 'KromaMintableERC20Factory', signer.provider)
  },
  // L1Block
  // - check version
  L1Block: async (hre: HardhatRuntimeEnvironment, signer: Signer) => {
    const L1Block = await hre.ethers.getContractAt(
      'L1Block',
      predeploys.L1Block,
      signer
    )

    await assertSemver(L1Block, 'L1Block')

    await checkProxy(hre, 'L1Block', signer.provider)
    await assertProxy(hre, 'L1Block', signer.provider)
  },
  // WETH9
  // - check name
  // - check symbol
  // - check decimals
  WETH9: async (hre: HardhatRuntimeEnvironment, signer: Signer) => {
    const WETH9 = await hre.ethers.getContractAt(
      'WETH9',
      predeploys.WETH9,
      signer
    )

    const name = await WETH9.name()
    assert(name === 'Wrapped Ether')
    console.log(`  - name: ${name}`)

    const symbol = await WETH9.symbol()
    assert(symbol === 'WETH')
    console.log(`  - symbol: ${symbol}`)

    const decimals = await WETH9.decimals()
    assert(decimals === 18)
    console.log(`  - decimals: ${decimals}`)
  },
  // L2ERC721Bridge
  // - check version
  L2ERC721Bridge: async (hre: HardhatRuntimeEnvironment, signer: Signer) => {
    const L2ERC721Bridge = await hre.ethers.getContractAt(
      'L2ERC721Bridge',
      predeploys.L2ERC721Bridge,
      signer
    )

    await assertSemver(L2ERC721Bridge, 'L2ERC721Bridge')

    const MESSENGER = await L2ERC721Bridge.MESSENGER()
    assert(MESSENGER !== hre.ethers.constants.AddressZero)
    console.log(`  - MESSENGER: ${MESSENGER}`)

    const OTHER_BRIDGE = await L2ERC721Bridge.OTHER_BRIDGE()
    assert(OTHER_BRIDGE !== hre.ethers.constants.AddressZero)
    yell(`  - OTHER_BRIDGE: ${OTHER_BRIDGE}`)

    await checkProxy(hre, 'L2ERC721Bridge', signer.provider)
    await assertProxy(hre, 'L2ERC721Bridge', signer.provider)
  },
  // KromaMintableERC721Factory
  // - check version
  KromaMintableERC721Factory: async (
    hre: HardhatRuntimeEnvironment,
    signer: Signer
  ) => {
    const KromaMintableERC721Factory = await hre.ethers.getContractAt(
      'KromaMintableERC721Factory',
      predeploys.KromaMintableERC721Factory,
      signer
    )

    await assertSemver(KromaMintableERC721Factory, 'KromaMintableERC721Factory')

    const BRIDGE = await KromaMintableERC721Factory.BRIDGE()
    assert(BRIDGE !== hre.ethers.constants.AddressZero)
    console.log(`  - BRIDGE: ${BRIDGE}`)

    const REMOTE_CHAIN_ID = await KromaMintableERC721Factory.REMOTE_CHAIN_ID()
    assert(REMOTE_CHAIN_ID !== 0)
    console.log(`  - REMOTE_CHAIN_ID: ${REMOTE_CHAIN_ID}`)

    await checkProxy(hre, 'KromaMintableERC721Factory', signer.provider)
    await assertProxy(hre, 'KromaMintableERC721Factory', signer.provider)
  },
  // ProxyAdmin
  // - check owner
  ProxyAdmin: async (hre: HardhatRuntimeEnvironment, signer: Signer) => {
    const ProxyAdmin = await hre.ethers.getContractAt(
      'ProxyAdmin',
      predeploys.ProxyAdmin,
      signer
    )

    const owner = await ProxyAdmin.owner()
    assert(owner !== hre.ethers.constants.AddressZero)
    yell(`  - owner: ${owner}`)

    await checkProxy(hre, 'ProxyAdmin', signer.provider)
    await assertProxy(hre, 'ProxyAdmin', signer.provider)
  },
  // ProtocolVault
  // - check version
  // - check MIN_WITHDRAWAL_AMOUNT
  // - check RECIPIENT
  ProtocolVault: async (hre: HardhatRuntimeEnvironment, signer: Signer) => {
    const ProtocolVault = await hre.ethers.getContractAt(
      'ProtocolVault',
      predeploys.ProtocolVault,
      signer
    )

    await assertSemver(ProtocolVault, 'ProtocolVault')

    const MIN_WITHDRAWAL_AMOUNT = await ProtocolVault.MIN_WITHDRAWAL_AMOUNT()
    console.log(`  - MIN_WITHDRAWAL_AMOUNT: ${MIN_WITHDRAWAL_AMOUNT}`)

    const RECIPIENT = await ProtocolVault.RECIPIENT()
    assert(RECIPIENT !== hre.ethers.constants.AddressZero)
    yell(`  - RECIPIENT: ${RECIPIENT}`)

    await checkProxy(hre, 'ProtocolVault', signer.provider)
    await assertProxy(hre, 'ProtocolVault', signer.provider)
  },
  // ProposerRewardVault
  // - check version
  // - check MIN_WITHDRAWAL_AMOUNT
  // - check RECIPIENT
  ProposerRewardVault: async (hre: HardhatRuntimeEnvironment, signer: Signer) => {
    const ProposerRewardVault = await hre.ethers.getContractAt(
      'ProposerRewardVault',
      predeploys.ProposerRewardVault,
      signer
    )

    await assertSemver(ProposerRewardVault, 'ProposerRewardVault')

    const MIN_WITHDRAWAL_AMOUNT = await ProposerRewardVault.MIN_WITHDRAWAL_AMOUNT()
    console.log(`  - MIN_WITHDRAWAL_AMOUNT: ${MIN_WITHDRAWAL_AMOUNT}`)

    const RECIPIENT = await ProposerRewardVault.RECIPIENT()
    assert(RECIPIENT !== hre.ethers.constants.AddressZero)
    yell(`  - RECIPIENT: ${RECIPIENT}`)

    await checkProxy(hre, 'ProposerRewardVault', signer.provider)
    await assertProxy(hre, 'ProposerRewardVault', signer.provider)
  },
  // L2ToL1MessagePasser
  // - check version
  L2ToL1MessagePasser: async (
    hre: HardhatRuntimeEnvironment,
    signer: Signer
  ) => {
    const L2ToL1MessagePasser = await hre.ethers.getContractAt(
      'L2ToL1MessagePasser',
      predeploys.L2ToL1MessagePasser,
      signer
    )

    await assertSemver(L2ToL1MessagePasser, 'L2ToL1MessagePasser')

    const MESSAGE_VERSION = await L2ToL1MessagePasser.MESSAGE_VERSION()
    console.log(`  - MESSAGE_VERSION: ${MESSAGE_VERSION}`)

    const messageNonce = await L2ToL1MessagePasser.messageNonce()
    console.log(`  - messageNonce: ${messageNonce}`)

    await checkProxy(hre, 'L2ToL1MessagePasser', signer.provider)
    await assertProxy(hre, 'L2ToL1MessagePasser', signer.provider)
  },
}

task('check-l2', 'Checks a freshly migrated L2 system for correct migration')
  .addOptionalParam('l1RpcUrl', 'L1 RPC URL of node', '', types.string)
  .addOptionalParam('l2RpcUrl', 'L2 RPC URL of node', '', types.string)
  .addOptionalParam('chainId', 'Expected chain id', 0, types.int)
  .addOptionalParam(
    'l2OutputOracleAddress',
    'Address of the L2OutputOracle oracle',
    '',
    types.string
  )
  .addOptionalParam(
    'skipPredeployCheck',
    'Skip long check',
    false,
    types.boolean
  )
  .setAction(async (args, hre: HardhatRuntimeEnvironment) => {
    yell('Manually check values wrapped in !!!!')
    console.log()

    let signer: Signer = hre.ethers.provider.getSigner()

    if (args.l2RpcUrl !== '') {
      console.log('Using CLI URL for provider instead of hardhat network')
      const provider = new hre.ethers.providers.JsonRpcProvider(args.l2RpcUrl)
      signer = Wallet.createRandom().connect(provider)
    }

    if (args.chainId !== 0) {
      const chainId = await signer.getChainId()
      if (chainId !== args.chainId) {
        throw new Error(
          `Unexpected Chain ID. Got ${chainId}, expected ${args.chainId}`
        )
      }
      console.log(`Verified Chain ID: ${chainId}`)
    } else {
      console.log(`Skipping Chain ID validation...`)
    }

    // Ensure that all the predeploys exist, including the not
    // currently configured ones
    if (!args.skipPredeployCheck) {
      await checkPredeploys(hre, signer.provider)
    }

    console.log()
    // Check the currently configured predeploys
    for (const [name, fn] of Object.entries(check)) {
      const address = predeploys[name]
      console.log(`${name}: ${address}`)
      await fn(hre, signer)
    }
  })
