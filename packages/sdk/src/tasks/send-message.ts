import { promises as fs } from 'fs'

import { providers, utils } from 'ethers'
import '@nomiclabs/hardhat-ethers'
import { task, types } from 'hardhat/config'
import { HardhatRuntimeEnvironment } from 'hardhat/types'
import 'hardhat-deploy'

import {
  CONTRACT_ADDRESSES,
  ContractsLike,
  CrossChainMessenger,
  DEFAULT_L2_CONTRACT_ADDRESSES,
  MessageDirection,
  assert,
  getAllContracts,
} from '../'

const { formatEther  } = utils

task('send-message', 'Deposits message to L2.')
  .addParam(
    'l2ProviderUrl',
    'L2 provider URL',
    'http://localhost:9545',
    types.string
  )
  .addOptionalParam('to', 'Recipient of message', '', types.string)
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
    console.log(`Signer balance: ${formatEther(balance.toString())} ETH`)

    const l2Provider = new providers.StaticJsonRpcProvider(args.l2ProviderUrl)


    // const to = args.to ? args.to : address

    // const amount = parseEther('100000')
    // const withdrawAmount = args.withdrawAmount
    //   ? parseEther(args.withdrawAmount)
    //   : amount.div(2)

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

    const {
      l1: {  L1CrossDomainMessenger, },
      l2: { L2CrossDomainMessenger},
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


    console.log("send to :", L1CrossDomainMessenger.address)
    console.log('current L1 BLock Number : ', await signer.provider.getBlockNumber())
    console.log('target address : ', L2CrossDomainMessenger.address)


    // const to = args.to ? args.to : address

      // 엉뚱한 컨트랙트한테 전송해서 실패할수밖에없는 deposit tx를 보낸다..
    await messenger.sendMessage(
      {
        direction : MessageDirection.L1_TO_L2,
        target :'0xc0D3c0D3c0D3c0d3c0d3C0d3C0d3c0D3C0D30004',
        message : "0x8a4068dd"
      },
      {
        l2GasLimit: '10000000'
      },
    )

    console.log('Completed sendMessage!')

  })
