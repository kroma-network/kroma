import { promises as fs } from 'fs'

import '@nomiclabs/hardhat-ethers'
import { Wallet, providers, utils } from 'ethers'
import { task, types } from 'hardhat/config'
import { HardhatRuntimeEnvironment } from 'hardhat/types'
import 'hardhat-deploy'

import {
  CONTRACT_ADDRESSES,
  ContractsLike,
  CrossChainMessenger,
  DEFAULT_L2_CONTRACT_ADDRESSES,
  assert,
} from '../'

const { formatEther, parseEther } = utils

task('initiate-withdrawal', 'Initiate a withdrawal.')
  .addOptionalParam('to', 'Recipient of the ether', '', types.string)
  .addOptionalParam(
    'amount',
    'Amount of ether to send (in ETH)',
    '',
    types.string
  )
  .addParam(
    'l1ProviderUrl',
    'L1 Provider URL',
    'http://localhost:8545',
    types.string
  )
  .addOptionalParam(
    'l1ContractsJsonPath',
    'Path to a JSON with L1 contract addresses in it',
    '',
    types.string
  )
  .setAction(async (args, hre: HardhatRuntimeEnvironment) => {
    const signers = await hre.ethers.getSigners()
    assert(signers.length > 0, 'No configured signers')
    // Use the first configured signer for simplicity
    const signer = signers[0]
    const address = signer.address
    console.log(`Using signer ${address}`)

    // Ensure that the signer has a balance before trying to
    // do anything
    const balance = await signer.getBalance()
    assert(balance.gt(0), 'Signer has no balance')
    console.log(`Signer balance: ${formatEther(balance.toString())} ETH`)

    // send to self if not specified
    const to = args.to ? args.to : address
    const amount = parseEther(args.amount ?? '1')

    const chainId = await signer.getChainId()
    let contractAddrs = CONTRACT_ADDRESSES[chainId]
    if (args.l1ContractsJsonPath) {
      const data = await fs.readFile(args.l1ContractsJsonPath)
      contractAddrs = {
        l1: JSON.parse(data.toString()),
        l2: DEFAULT_L2_CONTRACT_ADDRESSES,
      } as ContractsLike
    } else if (!contractAddrs) {
      const Deployment__L1CrossDomainMessenger = await hre.deployments.get(
        'L1CrossDomainMessengerProxy'
      )

      const Deployment__L1StandardBridge = await hre.deployments.get(
        'L1StandardBridgeProxy'
      )

      const Deployment__KromaPortal = await hre.deployments.get(
        'KromaPortalProxy'
      )

      const Deployment__L2OutputOracle = await hre.deployments.get(
        'L2OutputOracleProxy'
      )

      contractAddrs = {
        l1: {
          L1CrossDomainMessenger: Deployment__L1CrossDomainMessenger,
          L1StandardBridge: Deployment__L1StandardBridge,
          KromaPortal: Deployment__KromaPortal.address,
          L2OutputOracle: Deployment__L2OutputOracle.address,
        },
        l2: DEFAULT_L2_CONTRACT_ADDRESSES,
      }
    }

    const l1Provider = new providers.StaticJsonRpcProvider(args.l1ProviderUrl)
    const l1Signer = new Wallet(hre.network.config.accounts[0], l1Provider)

    const messenger = new CrossChainMessenger({
      l1SignerOrProvider: l1Signer,
      l2SignerOrProvider: signer,
      l1ChainId: await l1Signer.getChainId(),
      l2ChainId: chainId,
      contracts: contractAddrs,
    })

    console.log('Withdrawing ETH')
    const withdraw = await messenger.withdrawETH(amount, {
      recipient: to,
    })
    console.log(`Transaction hash: ${withdraw.hash}`)
    const withdrawReceipt = await withdraw.wait()
    console.log('Withdraw receipt', withdrawReceipt)

    console.log('Initiated withdrawal')
  })
