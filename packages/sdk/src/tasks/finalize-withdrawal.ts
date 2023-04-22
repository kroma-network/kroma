import { promises as fs } from 'fs'

import { TransactionReceipt } from '@ethersproject/abstract-provider'
import '@nomiclabs/hardhat-ethers'
import { Wallet, providers } from 'ethers'
import { task, types } from 'hardhat/config'
import { HardhatRuntimeEnvironment } from 'hardhat/types'
import 'hardhat-deploy'

import {
  CONTRACT_ADDRESSES,
  ContractsLike,
  CrossChainMessenger,
  DEFAULT_L2_CONTRACT_ADDRESSES,
  MessageStatus,
  assert,
} from '../'

task('finalize-withdrawal', 'Finalize a withdrawal')
  .addParam(
    'transactionHash',
    'L2 Transaction hash to finalize',
    '',
    types.string
  )
  .addParam(
    'l2ProviderUrl',
    'L2 provider URL',
    'http://localhost:9545',
    types.string
  )
  .addOptionalParam(
    'l1ContractsJsonPath',
    'Path to a JSON with L1 contract addresses in it',
    '',
    types.string
  )
  .setAction(async (args, hre: HardhatRuntimeEnvironment) => {
    const txHash = args.transactionHash
    assert(txHash !== '', 'No tx hash')

    const signers = await hre.ethers.getSigners()
    assert(signers.length > 0, 'No configured signers')
    const signer = signers[0]
    const address = signer.address
    console.log(`Using signer: ${address}`)

    const l2Provider = new providers.StaticJsonRpcProvider(args.l2ProviderUrl)
    const l2Signer = new Wallet(hre.network.config.accounts[0], l2Provider)

    const l2ChainId = await l2Signer.getChainId()
    let contractAddrs = CONTRACT_ADDRESSES[l2ChainId]
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

    const messenger = new CrossChainMessenger({
      l1SignerOrProvider: signer,
      l2SignerOrProvider: l2Signer,
      l1ChainId: await signer.getChainId(),
      l2ChainId,
      contracts: contractAddrs,
    })

    let receipt = await l2Provider.getTransactionReceipt(txHash)

    const currentStatus = await messenger.getMessageStatus(receipt)
    if (currentStatus === MessageStatus.RELAYED) {
      console.log('Withdrawal already proven and finalized')
      return
    }

    if (currentStatus > MessageStatus.READY_TO_PROVE) {
      console.log('Withdrawal already proven')
    } else {
      console.log('Waiting to be able to prove withdrawal')

      const proveInterval = setInterval(async () => {
        const currentProveStatus = await messenger.getMessageStatus(receipt)
        console.log(
          `Message status: ${
            MessageStatus[currentProveStatus]
          } (${await messenger.getLatestBlockNumber()} / ${
            receipt.blockNumber
          })`
        )
      }, 3000)

      try {
        await messenger.waitForMessageStatus(
          receipt,
          MessageStatus.READY_TO_PROVE
        )
      } finally {
        clearInterval(proveInterval)
      }

      let proveReceipt: TransactionReceipt | undefined
      try {
        const proveTx = await messenger.proveMessage(txHash)
        proveReceipt = await proveTx.wait()
        console.log('Prove receipt', proveReceipt)
      } catch (e) {
        console.error(
          'If you run this script on testnet, in most case, this happens when relayer has proven the transaction already.'
        )
        console.error(`Error occurred during proving withdrawal: ${e}`)
      }

      if (proveReceipt) {
        const proveBlock = await hre.ethers.provider.getBlock(
          proveReceipt.blockHash
        )
        const finalizationPeriodSeconds =
          await messenger.getChallengePeriodSeconds()

        console.log(`Withdrawal proven at ${proveBlock.timestamp}`)
        console.log(
          `Finalization period is ${finalizationPeriodSeconds} seconds`
        )

        const proveBlockFinalizedSeconds =
          proveBlock.timestamp + finalizationPeriodSeconds
        console.log(
          `Waiting to be able to finalize withdrawal until ${proveBlockFinalizedSeconds}`
        )
      }
    }

    const finalizeInterval = setInterval(async () => {
      const currentFinalizeStatus = await messenger.getMessageStatus(txHash)
      const lastBlock = await hre.ethers.provider.getBlock(-1)
      console.log(
        `Message status: ${MessageStatus[currentFinalizeStatus]} (${lastBlock.number}:${lastBlock.timestamp})`
      )
    }, 3000)

    try {
      await messenger.waitForMessageStatus(
        txHash,
        MessageStatus.READY_FOR_RELAY
      )
    } finally {
      clearInterval(finalizeInterval)
    }

    try {
      const tx = await messenger.finalizeMessage(txHash)
      receipt = await tx.wait()
      console.log(receipt)
    } catch (e) {
      console.error(
        'If you run this script on testnet, in most case, this happens when relayer has finalized the transaction already.'
      )
      console.error(`Error occurred during finalizing withdrawal: ${e}`)
    }
    console.log('Finalized withdrawal')
  })
