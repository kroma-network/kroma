import { promises as fs } from 'fs'

import '@nomiclabs/hardhat-ethers'
import { predeploys, getContractDefinition } from '@wemixkanvas/contracts'
import { providers, utils } from 'ethers'
import { task, types } from 'hardhat/config'
import 'hardhat-deploy'

import {
  CONTRACT_ADDRESSES,
  ContractsLike,
  CrossChainMessenger,
  DEFAULT_L2_CONTRACT_ADDRESSES,
  MessageStatus,
  assert,
} from '../'

const { formatEther, parseEther } = utils

task('deposit-eth', 'Deposits ether to L2.')
  .addParam(
    'l2ProviderUrl',
    'L2 provider URL',
    'http://localhost:9545',
    types.string
  )
  .addOptionalParam('to', 'Recipient of the ether', '', types.string)
  .addOptionalParam(
    'amount',
    'Amount of ether to send (in ETH)',
    '',
    types.string
  )
  .addOptionalParam(
    'withdraw',
    'Follow up with a withdrawal',
    true,
    types.boolean
  )
  .addOptionalParam(
    'l1ContractsJsonPath',
    'Path to a JSON with L1 contract addresses in it',
    '',
    types.string
  )
  .addOptionalParam('withdrawAmount', 'Amount to withdraw', '', types.string)
  .addFlag(
    'checkBalanceMismatch',
    'Whether to check balance after deposit and withdrawal'
  )
  .setAction(async (args, hre) => {
    const signers = await hre.ethers.getSigners()
    assert(signers.length > 0, 'No configured signers')
    // Use the first configured signer for simplicity
    const signer = signers[0]
    const address = signer.address
    console.log(`Using signer ${address}`)

    // Ensure that the signer has a balance before trying to
    // do anything
    const balance = await signer.getBalance()
    assert(balance.gt(0), 'Singer has no balance')
    console.log(`Signer balance: ${formatEther(balance.toString())} ETH`)

    const l2Provider = new providers.StaticJsonRpcProvider(args.l2ProviderUrl)

    // send to self if not specified
    const to = args.to ? args.to : address
    const amount = parseEther(args.amount ?? '1')
    const withdrawAmount = args.withdrawAmount
      ? parseEther(args.withdrawAmount)
      : amount.div(2)

    const l2Signer = new hre.ethers.Wallet(
      hre.network.config.accounts[0],
      l2Provider
    )

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

    const Artifact__L2ToL1MessagePasser = await getContractDefinition(
      'L2ToL1MessagePasser'
    )

    const Artifact__L2CrossDomainMessenger = await getContractDefinition(
      'L2CrossDomainMessenger'
    )

    const Artifact__L2StandardBridge = await getContractDefinition(
      'L2StandardBridge'
    )

    const Artifact__KanvasPortal = await getContractDefinition('KanvasPortal')

    const Artifact__L1CrossDomainMessenger = await getContractDefinition(
      'L1CrossDomainMessenger'
    )

    const Artifact__L1StandardBridge = await getContractDefinition(
      'L1StandardBridge'
    )

    const KanvasPortal = new hre.ethers.Contract(
      contractAddrs.l1.KanvasPortal,
      Artifact__KanvasPortal.abi,
      signer
    )

    const L1CrossDomainMessenger = new hre.ethers.Contract(
      contractAddrs.l1.L1CrossDomainMessenger,
      Artifact__L1CrossDomainMessenger.abi,
      signer
    )

    const L1StandardBridge = new hre.ethers.Contract(
      contractAddrs.l1.L1StandardBridge,
      Artifact__L1StandardBridge.abi,
      signer
    )

    const L2ToL1MessagePasser = new hre.ethers.Contract(
      predeploys.L2ToL1MessagePasser,
      Artifact__L2ToL1MessagePasser.abi
    )

    const L2CrossDomainMessenger = new hre.ethers.Contract(
      predeploys.L2CrossDomainMessenger,
      Artifact__L2CrossDomainMessenger.abi
    )

    const L2StandardBridge = new hre.ethers.Contract(
      predeploys.L2StandardBridge,
      Artifact__L2StandardBridge.abi
    )

    const messenger = new CrossChainMessenger({
      l1SignerOrProvider: signer,
      l2SignerOrProvider: l2Signer,
      l1ChainId: await signer.getChainId(),
      l2ChainId,
      contracts: contractAddrs,
    })

    const kanvasBalanceBefore = await signer.provider.getBalance(
      KanvasPortal.address
    )

    const l1BridgeBalanceBefore = await signer.provider.getBalance(
      L1StandardBridge.address
    )

    // Deposit ETH
    console.log('Depositing ETH through StandardBridge')
    console.log(`Sending ${formatEther(amount)} ETH`)
    const ethDeposit = await messenger.depositETH(amount, { recipient: to })
    console.log(`Transaction hash: ${ethDeposit.hash}`)
    const depositMessageReceipt = await messenger.waitForMessageReceipt(
      ethDeposit
    )
    assert(depositMessageReceipt.receiptStatus === 1, 'deposit failed')
    console.log(
      `Deposit complete - included in block ${depositMessageReceipt.transactionReceipt.blockNumber}`
    )

    const kanvasBalanceAfter = await signer.provider.getBalance(
      KanvasPortal.address
    )

    const l1BridgeBalanceAfter = await signer.provider.getBalance(
      L1StandardBridge.address
    )

    console.log(
      `L1StandardBridge balance before: ${formatEther(
        l1BridgeBalanceBefore
      )} ETH`
    )

    console.log(
      `L1StandardBridge balance after: ${formatEther(l1BridgeBalanceAfter)} ETH`
    )

    console.log(
      `KanvasPortal balance before: ${formatEther(kanvasBalanceBefore)} ETH`
    )
    console.log(
      `KanvasPortal balance after: ${formatEther(kanvasBalanceAfter)} ETH`
    )

    if (args.checkBalanceMismatch) {
      assert(
        kanvasBalanceBefore.add(amount).eq(kanvasBalanceAfter),
        'KanvasPortal balance mismatch'
      )
    }

    const l2Balance = await l2Provider.getBalance(to)
    console.log(
      `L2 balance of deposit recipient: ${formatEther(
        l2Balance.toString()
      )} ETH`
    )

    if (!args.withdraw) {
      return
    }

    console.log('Withdrawing ETH')
    const ethWithdraw = await messenger.withdrawETH(withdrawAmount)
    console.log(`Transaction hash: ${ethWithdraw.hash}`)
    const ethWithdrawReceipt = await ethWithdraw.wait()
    console.log(
      `ETH withdrawn on L2 - included in block ${ethWithdrawReceipt.blockNumber}`
    )

    {
      // check the logs
      for (const log of ethWithdrawReceipt.logs) {
        switch (log.address) {
          case L2ToL1MessagePasser.address: {
            const parsed = L2ToL1MessagePasser.interface.parseLog(log)
            console.log(parsed.name)
            console.log(parsed.args)
            console.log()
            break
          }
          case L2StandardBridge.address: {
            const parsed = L2StandardBridge.interface.parseLog(log)
            console.log(parsed.name)
            console.log(parsed.args)
            console.log()
            break
          }
          case L2CrossDomainMessenger.address: {
            const parsed = L2CrossDomainMessenger.interface.parseLog(log)
            console.log(parsed.name)
            console.log(parsed.args)
            console.log()
            break
          }
          default: {
            console.log(`Unknown log from ${log.address} - ${log.topics[0]}`)
          }
        }
      }
    }

    console.log('Waiting to be able to prove withdrawal')

    const proveInterval = setInterval(async () => {
      const currentStatus = await messenger.getMessageStatus(ethWithdrawReceipt)
      console.log(
        `Message status: ${
          MessageStatus[currentStatus]
        } (${await messenger.getLatestBlockNumber()} / ${
          ethWithdrawReceipt.blockNumber
        })`
      )
    }, 3000)

    try {
      await messenger.waitForMessageStatus(
        ethWithdrawReceipt,
        MessageStatus.READY_TO_PROVE
      )
    } finally {
      clearInterval(proveInterval)
    }

    console.log('Proving eth withdrawal...')
    const ethProve = await messenger.proveMessage(ethWithdrawReceipt)
    console.log(`Transaction hash: ${ethProve.hash}`)
    const ethProveReceipt = await ethProve.wait()
    assert(
      ethProveReceipt.status === 1,
      'Prove withdrawal transaction reverted'
    )
    console.log('Successfully proved withdrawal')

    const ethProveBlock = await hre.ethers.provider.getBlock(
      ethProveReceipt.blockHash
    )
    const finalizationPeriodSeconds =
      await messenger.getChallengePeriodSeconds()

    console.log(`Withdrawal proven at ${ethProveBlock.timestamp}`)
    console.log(`Finalization period is ${finalizationPeriodSeconds} seconds`)

    const ethProveBlockFinalizedSeconds =
      ethProveBlock.timestamp + finalizationPeriodSeconds
    console.log(
      `Waiting to be able to finalize withdrawal until ${ethProveBlockFinalizedSeconds}`
    )

    const finalizeInterval = setInterval(async () => {
      const currentStatus = await messenger.getMessageStatus(ethWithdrawReceipt)
      const lastBlock = await hre.ethers.provider.getBlock(-1)
      console.log(
        `Message status: ${MessageStatus[currentStatus]} (${lastBlock.number}:${lastBlock.timestamp})`
      )
    }, 3000)

    try {
      await messenger.waitForMessageStatus(
        ethWithdrawReceipt,
        MessageStatus.READY_FOR_RELAY
      )
    } finally {
      clearInterval(finalizeInterval)
    }

    console.log('Finalizing eth withdrawal...')
    const ethFinalize = await messenger.finalizeMessage(ethWithdrawReceipt)
    console.log(`Transaction hash: ${ethFinalize.hash}`)
    const ethFinalizeReceipt = await ethFinalize.wait()
    assert(ethFinalizeReceipt.status === 1, 'Finalize withdrawal reverted')

    console.log(
      `ETH withdrawal complete - included in block ${ethFinalizeReceipt.blockNumber}`
    )
    {
      // Check that the logs are correct
      for (const log of ethFinalizeReceipt.logs) {
        switch (log.address) {
          case L1StandardBridge.address: {
            const parsed = L1StandardBridge.interface.parseLog(log)
            console.log(parsed.name)
            console.log(parsed.args)
            console.log()
            assert(
              parsed.name === 'ETHBridgeFinalized' ||
                parsed.name === 'ETHWithdrawalFinalized',
              'Wrong event name from L1StandardBridge'
            )
            assert(
              parsed.args.amount.eq(withdrawAmount),
              'Wrong amount in event'
            )
            assert(parsed.args.from === address, 'Wrong to in event')
            assert(parsed.args.to === address, 'Wrong from in event')
            break
          }
          case L1CrossDomainMessenger.address: {
            const parsed = L1CrossDomainMessenger.interface.parseLog(log)
            console.log(parsed.name)
            console.log(parsed.args)
            console.log()
            assert(
              parsed.name === 'RelayedMessage',
              'Wrong event from L1CrossDomainMessenger'
            )
            break
          }
          case KanvasPortal.address: {
            const parsed = KanvasPortal.interface.parseLog(log)
            console.log(parsed.name)
            console.log(parsed.args)
            console.log()
            // TODO: remove this if check
            if (parsed.name === 'WithdrawalFinalized') {
              assert(parsed.args.success, 'Unsuccessful withdrawal call')
            }
            break
          }
          default: {
            console.log(`Unknown log from ${log.address} - ${log.topics[0]}`)
          }
        }
      }
    }

    if (args.checkBalanceMismatch) {
      const kanvasBalanceFinally = await signer.provider.getBalance(
        KanvasPortal.address
      )
      assert(
        kanvasBalanceAfter.sub(withdrawAmount).eq(kanvasBalanceFinally),
        'KanvasPortal balance mismatch'
      )
    }

    console.log('Withdraw success')
  })
