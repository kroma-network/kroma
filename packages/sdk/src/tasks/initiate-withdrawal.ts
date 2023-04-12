import { promises as fs } from 'fs'

import '@nomiclabs/hardhat-ethers'
import { providers, utils, Wallet } from 'ethers'
import { task, types } from 'hardhat/config'
import 'hardhat-deploy'

import {
  CONTRACT_ADDRESSES,
  CrossChainMessenger,
  DEFAULT_L2_CONTRACT_ADDRESSES,
  ContractsLike,
} from '../'

const { formatEther } = utils

task('initiate-withdrawal', 'Initiate a withdrawal.')
  .addOptionalParam('to', 'Recipient of the ether', '', types.string)
  .addOptionalParam(
    'amount',
    'Amount of ether to send (in ETH)',
    '',
    types.string
  )
  .addParam('l1Url', 'L1 HTTP URL', 'http://localhost:8545', types.string)
  .addOptionalParam(
    'l1ContractsJsonPath',
    'Path to a JSON with L1 contract addresses in it',
    '',
    types.string
  )
  .setAction(async (args, hre) => {
    const signers = await hre.ethers.getSigners()
    if (signers.length === 0) {
      throw new Error('No configured signers')
    }
    // Use the first configured signer for simplicity
    const signer = signers[0]
    const address = await signer.getAddress()
    console.log(`Using signer ${address}`)

    // Ensure that the signer has a balance before trying to
    // do anything
    const balance = await signer.getBalance()
    if (balance.eq(0)) {
      throw new Error('Signer has no balance')
    }
    console.log(`Signer balance: ${formatEther(balance.toString())}`)

    // send to self if not specified
    const to = args.to ? args.to : address
    const amount = args.amount
      ? utils.parseEther(args.amount)
      : utils.parseEther('1')

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

      const Deployment__KanvasPortal = await hre.deployments.get(
        'KanvasPortalProxy'
      )

      const Deployment__L2OutputOracle = await hre.deployments.get(
        'L2OutputOracleProxy'
      )

      contractAddrs = {
        l1: {
          L1CrossDomainMessenger: Deployment__L1CrossDomainMessenger,
          L1StandardBridge: Deployment__L1StandardBridge,
          KanvasPortal: Deployment__KanvasPortal.address,
          L2OutputOracle: Deployment__L2OutputOracle.address,
        },
        l2: DEFAULT_L2_CONTRACT_ADDRESSES,
      }
    }

    const l1Provider = new providers.StaticJsonRpcProvider(args.l1Url)
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
