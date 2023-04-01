import {
  Provider,
  TransactionReceipt,
  TransactionResponse,
} from '@ethersproject/abstract-provider'
import { Signer } from '@ethersproject/abstract-signer'
import { BigNumber, Contract } from 'ethers'

/**
 * L1 network chain IDs
 */
export enum L1ChainID {
  SEPOLIA = 11155111,
  KANVAS_EASEL = 7789,
  LOCAL_DEVNET = 900,
}

/**
 * L2 network chain IDs
 */
export enum L2ChainID {
  KANVAS_AQUA = 2357,
  KANVAS_SAIL = 7790,
  KANVAS_LOCAL_DEVNET = 901,
}

/**
 * L1 contract references.
 */
export interface L1Contracts {
  KanvasPortal: Contract
  L1CrossDomainMessenger: Contract
  L1StandardBridge: Contract
  L2OutputOracle: Contract
}

/**
 * L2 contract references.
 */
export interface L2Contracts {
  L2CrossDomainMessenger: Contract
  L2StandardBridge: Contract
  L2ToL1MessagePasser: Contract
  WETH9: Contract
}

/**
 * Represents Kanvas contracts, assumed to be connected to their appropriate
 * providers and addresses.
 */
export interface Contracts {
  l1: L1Contracts
  l2: L2Contracts
}

/**
 * Convenience type for something that looks like the L1 contract interface but could be
 * addresses instead of actual contract objects.
 */
export type L1ContractsLike = {
  [K in keyof L1Contracts]: AddressLike
}

/**
 * Convenience type for something that looks like the L2 contract interface but could be
 * addresses instead of actual contract objects.
 */
export type L2ContractsLike = {
  [K in keyof L2Contracts]: AddressLike
}

/**
 * Convenience type for something that looks like the contract interface but could be
 * addresses instead of actual contract objects.
 */
export interface ContractsLike {
  l1: L1ContractsLike
  l2: L2ContractsLike
}

/**
 * Enum describing the status of a message.
 */
export enum MessageStatus {
  /**
   * Message is an L1 to L2 message and has not been processed by the L2.
   */
  UNCONFIRMED_L1_TO_L2_MESSAGE,

  /**
   * Message is an L1 to L2 message and the transaction to execute the message failed.
   * When this status is returned, you will need to resend the L1 to L2 message, probably with a
   * higher gas limit.
   */
  FAILED_L1_TO_L2_MESSAGE,

  /**
   * Message is an L2 to L1 message and no output root has been published yet.
   */
  OUTPUT_ROOT_NOT_PUBLISHED,

  /**
   * Message is ready to be proved on L1 to initiate the challenge period.
   */
  READY_TO_PROVE,

  /**
   * Message is a proved L2 to L1 message and is undergoing the challenge period.
   */
  IN_CHALLENGE_PERIOD,

  /**
   * Message is ready to be relayed.
   */
  READY_FOR_RELAY,

  /**
   * Message has been relayed.
   */
  RELAYED,
}

/**
 * Enum describing the direction of a message.
 */
export enum MessageDirection {
  L1_TO_L2,
  L2_TO_L1,
}

/**
 * Partial message that needs to be signed and executed by a specific signer.
 */
export interface CrossChainMessageRequest {
  direction: MessageDirection
  target: string
  message: string
}

/**
 * Core components of a cross chain message.
 */
export interface CoreCrossChainMessage {
  sender: string
  target: string
  message: string
  messageNonce: BigNumber
  value: BigNumber
  minGasLimit: BigNumber
}

/**
 * Describes a message that is sent between L1 and L2. Direction determines where the message was
 * sent from and where it's being sent to.
 */
export interface CrossChainMessage extends CoreCrossChainMessage {
  direction: MessageDirection
  logIndex: number
  blockNumber: number
  transactionHash: string
}

/**
 * Describes messages sent inside the L2ToL1MessagePasser on L2. Happens to be the same structure
 * as the CoreCrossChainMessage so we'll reuse the type for now.
 */
export type LowLevelMessage = CoreCrossChainMessage

/**
 * Describes an ETH withdrawal or deposit, along with the underlying raw cross chain message
 * behind the deposit or withdrawal.
 */
export interface ETHBridgeMessage {
  direction: MessageDirection
  from: string
  to: string
  amount: BigNumber
  data: string
  logIndex: number
  blockNumber: number
  transactionHash: string
}

/**
 * Describes an ERC20 withdrawal or deposit, along with the underlying raw cross chain message
 * behind the deposit or withdrawal.
 */
export interface ERC20BridgeMessage {
  direction: MessageDirection
  from: string
  to: string
  l1Token: string
  l2Token: string
  amount: BigNumber
  data: string
  logIndex: number
  blockNumber: number
  transactionHash: string
}

export type BridgeMessage = ETHBridgeMessage | ERC20BridgeMessage

/**
 * Represents a withdrawal entry within the logs of a L2 to L1
 * CrossChainMessage
 */
export interface WithdrawalEntry {
  MessagePassed: any
}

/**
 * Enum describing the status of a CrossDomainMessage message receipt.
 */
export enum MessageReceiptStatus {
  RELAYED_FAILED,
  RELAYED_SUCCEEDED,
}

/**
 * CrossDomainMessage receipt.
 */
export interface MessageReceipt {
  receiptStatus: MessageReceiptStatus
  transactionReceipt: TransactionReceipt
}

/**
 * ProvenWithdrawal in KanvasPortal
 */
export interface ProvenWithdrawal {
  outputRoot: string
  timestamp: BigNumber
  l2BlockNumber: BigNumber
}

/**
 * Stuff that can be coerced into a transaction.
 */
export type TransactionLike = string | TransactionReceipt | TransactionResponse

/**
 * Stuff that can be coerced into a CrossChainMessage.
 */
export type MessageLike =
  | CrossChainMessage
  | TransactionLike
  | ERC20BridgeMessage

/**
 * Stuff that can be coerced into a CrossChainMessageRequest.
 */
export type MessageRequestLike =
  | CrossChainMessageRequest
  | CrossChainMessage
  | TransactionLike
  | ERC20BridgeMessage

/**
 * Stuff that can be coerced into a provider.
 */
export type ProviderLike = string | Provider

/**
 * Stuff that can be coerced into a signer.
 */
export type SignerLike = string | Signer

/**
 * Stuff that can be coerced into a signer or provider.
 */
export type SignerOrProviderLike = SignerLike | ProviderLike

/**
 * Stuff that can be coerced into an address.
 */
export type AddressLike = string | Contract

/**
 * Stuff that can be coerced into a number.
 */
export type NumberLike = string | number | BigNumber
