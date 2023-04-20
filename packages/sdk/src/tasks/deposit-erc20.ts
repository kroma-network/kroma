import { promises as fs } from 'fs'

import { getContractDefinition, predeploys } from '@kroma-network/contracts'
import { sleep } from '@kroma-network/core-utils'
import '@nomiclabs/hardhat-ethers'
import { BigNumber, Contract, Event, Wallet, providers, utils } from 'ethers'
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
  getAllContracts,
} from '../'

const oneEtherInWei = utils.parseEther('1')

const deployWETH9 = async (
  hre: HardhatRuntimeEnvironment,
  wrap: boolean
): Promise<Contract> => {
  const signers = await hre.ethers.getSigners()
  const signer = signers[0]

  const Artifact__WETH9 = await getContractDefinition('WETH9')
  const Factory__WETH9 = new hre.ethers.ContractFactory(
    Artifact__WETH9.abi,
    Artifact__WETH9.bytecode,
    signer
  )

  console.log('Sending deployment transaction')
  const WETH9 = await Factory__WETH9.deploy()
  const receipt = await WETH9.deployTransaction.wait()
  console.log(`WETH9 deployed: ${receipt.transactionHash}`)

  if (wrap) {
    const deposit = await signer.sendTransaction({
      value: oneEtherInWei,
      to: WETH9.address,
    })
    await deposit.wait()
  }

  return WETH9
}

const createKromaMintableERC20 = async (
  L1ERC20: Contract,
  l2Signer: Wallet
): Promise<Contract> => {
  const Artifact__KromaMintableERC20Token = await getContractDefinition(
    'KromaMintableERC20'
  )

  const Artifact__KromaMintableERC20TokenFactory = await getContractDefinition(
    'KromaMintableERC20Factory'
  )

  const KromaMintableERC20TokenFactory = new Contract(
    predeploys.KromaMintableERC20Factory,
    Artifact__KromaMintableERC20TokenFactory.abi,
    l2Signer
  )

  const name = await L1ERC20.name()
  const symbol = await L1ERC20.symbol()

  const tx = await KromaMintableERC20TokenFactory.createKromaMintableERC20(
    L1ERC20.address,
    `L2 ${name}`,
    `L2-${symbol}`
  )

  const receipt = await tx.wait()
  const event = receipt.events.find(
    (e: Event) => e.event === 'KromaMintableERC20Created'
  )

  assert(event, 'Unable to find KromaMintableERC20Created event')

  const l2WethAddress = event.args.localToken
  console.log(`Deployed to ${l2WethAddress}`)

  return new Contract(
    l2WethAddress,
    Artifact__KromaMintableERC20Token.abi,
    l2Signer
  )
}

// TODO(tynes): this task could be modularized in the future
// so that it can deposit an arbitrary token. Right now it
// deploys a WETH9 contract, mints some WETH9 and then
// deposits that into L2 through the StandardBridge.
task('deposit-erc20', 'Deposits WETH9 onto L2.')
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
    assert(balance.gt(0), 'Signer has no balance')

    const l2Provider = new providers.StaticJsonRpcProvider(args.l2ProviderUrl)

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
    }

    const {
      l1: { KromaPortal, L1CrossDomainMessenger, L1StandardBridge },
      l2: { L2CrossDomainMessenger, L2StandardBridge, L2ToL1MessagePasser },
    } = getAllContracts(l2ChainId, {
      l1SignerOrProvider: signer,
      overrides: contractAddrs,
    })

    const messenger = new CrossChainMessenger({
      l1SignerOrProvider: signer,
      l2SignerOrProvider: l2Signer,
      l1ChainId: await signer.getChainId(),
      l2ChainId,
      contracts: contractAddrs,
    })

    console.log('Deploying WETH9 to L1')
    const WETH9 = await deployWETH9(hre, true)
    console.log(`Deployed to ${WETH9.address}`)

    console.log('Creating L2 WETH9')
    const KromaMintableERC20 = await createKromaMintableERC20(WETH9, l2Signer)

    console.log(`Approving WETH9 for deposit`)
    const approvalTx = await messenger.approveERC20(
      WETH9.address,
      KromaMintableERC20.address,
      hre.ethers.constants.MaxUint256
    )
    await approvalTx.wait()
    console.log('WETH9 approved')

    console.log('Depositing WETH9 to L2')
    const depositTx = await messenger.depositERC20(
      WETH9.address,
      KromaMintableERC20.address,
      oneEtherInWei
    )
    await depositTx.wait()
    console.log(`ERC20 deposited - ${depositTx.hash}`)

    // Deposit might get reorged, wait 30s and also log for reorgs.
    let prevBlockHash = ''
    for (let i = 0; i < 30; i++) {
      const messageReceipt = await messenger.waitForMessageReceipt(depositTx)
      if (messageReceipt.receiptStatus !== 1) {
        console.log(`Deposit failed, retrying...`)
      }

      if (
        prevBlockHash !== '' &&
        messageReceipt.transactionReceipt.blockHash !== prevBlockHash
      ) {
        console.log(
          `Block hash changed from ${prevBlockHash} to ${messageReceipt.transactionReceipt.blockHash}`
        )

        // Wait for stability, we want at least 30 seconds after any reorg
        i = 0
      }

      prevBlockHash = messageReceipt.transactionReceipt.blockHash
      await sleep(1000)
    }

    if (args.checkBalanceMismatch) {
      const l2Balance = await KromaMintableERC20.balanceOf(address)
      assert(l2Balance.gte(oneEtherInWei), 'bad deposit')
    }

    console.log(`Deposit success`)

    let preBalance: BigNumber | undefined
    if (args.checkBalanceMismatch) {
      preBalance = await WETH9.balanceOf(address)
    }

    console.log('Starting withdrawal')
    const withdraw = await messenger.withdrawERC20(
      WETH9.address,
      KromaMintableERC20.address,
      oneEtherInWei
    )
    const withdrawalReceipt = await withdraw.wait()
    for (const log of withdrawalReceipt.logs) {
      switch (log.address) {
        case L2ToL1MessagePasser.address: {
          const parsed = L2ToL1MessagePasser.interface.parseLog(log)
          console.log(`Log ${parsed.name} from ${log.address}`)
          console.log(parsed.args)
          console.log()
          break
        }
        case L2StandardBridge.address: {
          const parsed = L2StandardBridge.interface.parseLog(log)
          console.log(`Log ${parsed.name} from ${log.address}`)
          console.log(parsed.args)
          console.log()
          break
        }
        case L2CrossDomainMessenger.address: {
          const parsed = L2CrossDomainMessenger.interface.parseLog(log)
          console.log(`Log ${parsed.name} from ${log.address}`)
          console.log(parsed.args)
          console.log()
          break
        }
        default: {
          console.log(`Unknown log from ${log.address} - ${log.topics[0]}`)
        }
      }
    }

    setInterval(async () => {
      const currentStatus = await messenger.getMessageStatus(withdraw)
      console.log(`Message status: ${MessageStatus[currentStatus]}`)
      console.log(
        `Message status: ${
          MessageStatus[currentStatus]
        } (${await messenger.getLatestBlockNumber()} / ${
          withdrawalReceipt.blockNumber
        })`
      )
    }, 3000)

    const now = Math.floor(Date.now() / 1000)

    console.log('Waiting for message to be able to be proved')
    await messenger.waitForMessageStatus(withdraw, MessageStatus.READY_TO_PROVE)

    console.log('Proving withdrawal...')
    const prove = await messenger.proveMessage(withdraw)
    const proveReceipt = await prove.wait()
    console.log(proveReceipt)
    assert(proveReceipt.status === 1, 'Prove withdrawal transaction reverted')

    const proveBlock = await hre.ethers.provider.getBlock(
      proveReceipt.blockHash
    )
    const finalizationPeriodSeconds =
      await messenger.getChallengePeriodSeconds()

    console.log(`Withdrawal proven at ${proveBlock.timestamp}`)
    console.log(`Finalization period is ${finalizationPeriodSeconds} seconds`)

    const proveBlockFinalizedSeconds =
      proveBlock.timestamp + finalizationPeriodSeconds
    console.log(
      `Waiting to be able to finalize withdrawal until ${proveBlockFinalizedSeconds}`
    )

    const finalizeInterval = setInterval(async () => {
      const currentStatus = await messenger.getMessageStatus(withdraw)
      const lastBlock = await hre.ethers.provider.getBlock(-1)
      console.log(
        `Message status: ${MessageStatus[currentStatus]} (${lastBlock.number}:${lastBlock.timestamp})`
      )
    }, 3000)

    try {
      await messenger.waitForMessageStatus(
        withdraw,
        MessageStatus.READY_FOR_RELAY
      )
    } finally {
      clearInterval(finalizeInterval)
    }

    console.log('Finalizing withdrawal...')
    // TODO: Update SDK to properly estimate gas
    const finalize = await messenger.finalizeMessage(withdraw, {
      overrides: { gasLimit: 500_000 },
    })
    const finalizeReceipt = await finalize.wait()
    console.log('finalizeReceipt:', finalizeReceipt)
    console.log(`Took ${Math.floor(Date.now() / 1000) - now} seconds`)

    for (const log of finalizeReceipt.logs) {
      switch (log.address) {
        case KromaPortal.address: {
          const parsed = KromaPortal.interface.parseLog(log)
          console.log(`Log ${parsed.name} from KromaPortal (${log.address})`)
          console.log(parsed.args)
          console.log()
          break
        }
        case L1CrossDomainMessenger.address: {
          const parsed = L1CrossDomainMessenger.interface.parseLog(log)
          console.log(
            `Log ${parsed.name} from L1CrossDomainMessenger (${log.address})`
          )
          console.log(parsed.args)
          console.log()
          break
        }
        case L1StandardBridge.address: {
          const parsed = L1StandardBridge.interface.parseLog(log)
          console.log(
            `Log ${parsed.name} from L1StandardBridge (${log.address})`
          )
          console.log(parsed.args)
          console.log()
          break
        }
        case WETH9.address: {
          const parsed = WETH9.interface.parseLog(log)
          console.log(`Log ${parsed.name} from WETH9 (${log.address})`)
          console.log(parsed.args)
          console.log()
          break
        }
        default:
          console.log(
            `Unknown log emitted from ${log.address} - ${log.topics[0]}`
          )
      }
    }

    if (args.checkBalanceMismatch) {
      const postBalance = await WETH9.balanceOf(address)
      const expectedBalance = preBalance.add(oneEtherInWei)
      assert(
        expectedBalance.eq(postBalance),
        `Balance mismatch, expected: ${expectedBalance}, actual: ${postBalance}`
      )
    }

    console.log('Withdrawal success')
  })
