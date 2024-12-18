import { KromaNodeProvider, sleep } from '@kroma/core-utils'
import '@nomiclabs/hardhat-ethers'
import { task, types } from 'hardhat/config'
import 'hardhat-deploy'

import { predeploys } from '../src'

task('watch', 'Watch a Kroma System')
  .addParam(
    'l1ProviderUrl',
    'L1 provider URL.',
    'http://localhost:8545',
    types.string
  )
  .addParam(
    'l2ProviderUrl',
    'L2 provider URL.',
    'http://localhost:9545',
    types.string
  )
  .addParam(
    'kromaNodeProviderUrl',
    'kroma-node provider URL',
    'http://localhost:7545',
    types.string
  )
  .setAction(async (args, hre) => {
    const { utils } = hre.ethers

    const l1Provider = new hre.ethers.providers.StaticJsonRpcProvider(
      args.l1Provider
    )
    const l2Provider = new hre.ethers.providers.StaticJsonRpcProvider(
      args.l2ProviderUrl
    )

    const contracts = {}
    const deployments = await hre.deployments.all()
    for (const [contract, deployment] of Object.entries(deployments)) {
      contracts[contract] = deployment.address
    }
    console.log('Deployed Contracts')
    console.table(contracts)

    const kromaNodeProvider = new KromaNodeProvider(args.kromaNodeProviderUrl)
    const kromaNodeConfig = await kromaNodeProvider.rollupConfig()
    console.log('kroma-node config')
    console.table({
      'layer-one-hash': kromaNodeConfig.genesis.l1.hash,
      'layer-one-number': kromaNodeConfig.genesis.l1.number,
      'layer-two-hash': kromaNodeConfig.genesis.l2.hash,
      'layer-two-number': kromaNodeConfig.genesis.l2.number,
      'layer-two-time': kromaNodeConfig.genesis.l2_time,
      'block-time': kromaNodeConfig.block_time,
      'max-sequencer-drift': kromaNodeConfig.max_sequencer_drift,
      'sequencer-window-size': kromaNodeConfig.sequencer_window_size,
      'channel-timeout': kromaNodeConfig.channel_timeout,
      'l1-chain-id': kromaNodeConfig.l1_chain_id,
      'l2-chain-id': kromaNodeConfig.l2_chain_id,
      'p2p-sequencer-address': kromaNodeConfig.p2p_sequencer_address,
      'fee-recipient-address': kromaNodeConfig.fee_recipient_address,
      'batch-inbox-address': kromaNodeConfig.batch_inbox_address,
      'batch-sender-address': kromaNodeConfig.batch_sender_address,
      'deposit-contract-address': kromaNodeConfig.deposit_contract_address,
    })

    const Deployment__L2OutputOracle = await hre.deployments.get(
      'L2OutputOracle'
    )

    const Deployment__L2OutputOracleProxy = await hre.deployments.get(
      'L2OutputOracleProxy'
    )

    const L2OutputOracle = new hre.ethers.Contract(
      Deployment__L2OutputOracleProxy.address,
      Deployment__L2OutputOracle.abi,
      l1Provider
    )

    const validator = await L2OutputOracle.validator()
    console.log(`L2OutputOracle validator ${validator}`)
    console.log()

    setInterval(async () => {
      const latestBlockNumber = await L2OutputOracle.latestBlockNumber()
      console.log(
        `L2OutputOracle latest block number: ${latestBlockNumber.toString()}`
      )
      console.log()
    }, 10000)

    l1Provider.on('block', async (num) => {
      const block = await l1Provider.getBlockWithTransactions(num)
      for (const txn of block.transactions) {
        const to = utils.getAddress(txn.to || hre.ethers.constants.AddressZero)
        const from = utils.getAddress(txn.from)
        const isBatchSender =
          utils.getAddress(txn.from) ===
          utils.getAddress(kromaNodeConfig.batch_sender_address)
        const isBatchInbox =
          to === utils.getAddress(kromaNodeConfig.batch_inbox_address)

        const isOutputOracle =
          to === utils.getAddress(L2OutputOracle.address) &&
          from === utils.getAddress(validator)

        if (isBatchSender && isBatchInbox) {
          console.log('Batch submitted:')
          console.log(`  tx hash: ${txn.hash}`)
          console.log(`  tx data: ${txn.data}`)
          console.log()
        }

        if (isOutputOracle) {
          console.log('L2 Output Submitted:')
          const data = L2OutputOracle.interface.parseTransaction(txn)
          console.log(`  tx hash: ${txn.hash}`)
          console.log(`    output root:   ${data.args._outputRoot}`)
          console.log(`    l2 blocknum:   ${data.args._l2BlockNumber}`)
          console.log(`    l1 blockhash:  ${data.args._l1Blockhash}`)
          console.log(`    l1 blocknum:   ${data.args._l1BlockNumber}`)
          console.log()
        }
      }
    })

    const KromaL1Block = await hre.ethers.getContractAt(
      'KromaL1Block',
      predeploys.KromaL1Block
    )
    const L1Block = await hre.ethers.getContractAt(
      'L1Block',
      predeploys.L1Block
    )

    l2Provider.on('block', async (num) => {
      const block = await l2Provider.getBlockWithTransactions(num)
      for (const txn of block.transactions) {
        const to = utils.getAddress(txn.to || hre.ethers.constants.AddressZero)

        if (to === utils.getAddress(predeploys.KromaL1Block)) {
          const data = KromaL1Block.interface.parseTransaction(txn)
          console.log('KromaL1Block values updated')
          console.log(`  tx hash: ${txn.hash}`)
          console.log(`    number:         ${data.args._number}`)
          console.log(`    timestamp:      ${data.args._timestamp}`)
          console.log(`    basefee:        ${data.args._basefee}`)
          console.log(`    hash:           ${data.args._hash}`)
          console.log(`    sequenceNumber: ${data.args._sequenceNumber}`)
          console.log()
        } else if (to === utils.getAddress(predeploys.L1Block)) {
          const data = L1Block.interface.parseTransaction(txn)
          console.log('L1Block values updated')
          console.log(`  tx hash: ${txn.hash}`)
          console.log(`    number:         ${data.args._number}`)
          console.log(`    timestamp:      ${data.args._timestamp}`)
          console.log(`    basefee:        ${data.args._basefee}`)
          console.log(`    hash:           ${data.args._hash}`)
          console.log(`    sequenceNumber: ${data.args._sequenceNumber}`)
          console.log()
        }
      }
    })

    setInterval(async () => {
      await sleep(100000)
    })
    await sleep(100000)
  })
