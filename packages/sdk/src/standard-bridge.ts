/* eslint-disable @typescript-eslint/no-unused-vars */
import {
  BlockTag,
  TransactionRequest,
  TransactionResponse,
} from '@ethersproject/abstract-provider'
import { getContractInterface } from '@kroma-network/contracts'
import { hexStringEquals } from '@kroma-network/core-utils'
import {
  BigNumber,
  CallOverrides,
  Contract,
  Overrides,
  PayableOverrides,
  Signer,
  ethers,
} from 'ethers'

import { CrossChainMessenger } from './cross-chain-messenger'
import {
  AddressLike,
  ERC20BridgeMessage,
  MessageDirection,
  NumberLike,
} from './interfaces'
import { toAddress } from './utils'

/**
 * Bridge adapter for any token bridge that uses the standard token bridge interface.
 */
export class StandardBridge {
  /**
   * Provider used to make queries related to cross-chain interactions.
   */
  public messenger: CrossChainMessenger

  /**
   * L1 bridge contract.
   */
  public l1Bridge: Contract

  /**
   * L2 bridge contract.
   */
  public l2Bridge: Contract

  /**
   * Creates a StandardBridgeAdapter instance.
   *
   * @param opts Options for the adapter.
   * @param opts.messenger Provider used to make queries related to cross-chain interactions.
   * @param opts.l1Bridge L1 bridge contract.
   * @param opts.l2Bridge L2 bridge contract.
   */
  constructor(opts: {
    messenger: CrossChainMessenger
    l1Bridge: AddressLike
    l2Bridge: AddressLike
  }) {
    this.messenger = opts.messenger
    this.l1Bridge = new Contract(
      toAddress(opts.l1Bridge),
      getContractInterface('L1StandardBridge'),
      this.messenger.l1Provider
    )
    this.l2Bridge = new Contract(
      toAddress(opts.l2Bridge),
      getContractInterface('L2StandardBridge'),
      this.messenger.l2Provider
    )
  }

  /**
   * Gets all deposits for a given address.
   *
   * @param address Address to search for messages from.
   * @param opts Options object.
   * @param opts.fromBlock Block to start searching for messages from. If not provided, will start
   * from the first block (block #0).
   * @param opts.toBlock Block to stop searching for messages at. If not provided, will stop at the
   * latest known block ("latest").
   * @returns All deposit token bridge messages sent by the given address.
   */
  public async getDepositsByAddress(
    address: AddressLike,
    opts?: {
      fromBlock?: BlockTag
      toBlock?: BlockTag
    }
  ): Promise<ERC20BridgeMessage[]> {
    const events = await this.l1Bridge.queryFilter(
      this.l1Bridge.filters.ERC20BridgeInitiated(
        /*localToken=*/ undefined,
        /*remoteToken=*/ undefined,
        /*from=*/ address
      ),
      opts?.fromBlock,
      opts?.toBlock
    )

    return events
      .map((event) => {
        return {
          direction: MessageDirection.L1_TO_L2,
          from: event.args.from,
          to: event.args.to,
          l1Token: event.args.localToken,
          l2Token: event.args.remoteToken,
          amount: event.args.amount,
          data: event.args.extraData,
          logIndex: event.logIndex,
          blockNumber: event.blockNumber,
          transactionHash: event.transactionHash,
        }
      })
      .sort((a, b) => {
        // Sort descending by block number
        return b.blockNumber - a.blockNumber
      })
  }

  /**
   * Gets all withdrawals for a given address.
   *
   * @param address Address to search for messages from.
   * @param opts Options object.
   * @param opts.fromBlock Block to start searching for messages from. If not provided, will start
   * from the first block (block #0).
   * @param opts.toBlock Block to stop searching for messages at. If not provided, will stop at the
   * latest known block ("latest").
   * @returns All withdrawal token bridge messages sent by the given address.
   */
  public async getWithdrawalsByAddress(
    address: AddressLike,
    opts?: {
      fromBlock?: BlockTag
      toBlock?: BlockTag
    }
  ): Promise<ERC20BridgeMessage[]> {
    const events = await this.l2Bridge.queryFilter(
      this.l2Bridge.filters.ERC20BridgeInitiated(
        /*localToken=*/ undefined,
        /*remoteToken=*/ undefined,
        /*from=*/ address
      ),
      opts?.fromBlock,
      opts?.toBlock
    )

    return events
      .map((event) => {
        return {
          direction: MessageDirection.L2_TO_L1,
          from: event.args.from,
          to: event.args.to,
          l1Token: event.args.remoteToken,
          l2Token: event.args.localToken,
          amount: event.args.amount,
          data: event.args.extraData,
          logIndex: event.logIndex,
          blockNumber: event.blockNumber,
          transactionHash: event.transactionHash,
        }
      })
      .sort((a, b) => {
        // Sort descending by block number
        return b.blockNumber - a.blockNumber
      })
  }

  /**
   * Checks whether the given token pair is supported by the bridge.
   *
   * @param l1Token The L1 token address.
   * @param l2Token The L2 token address.
   * @returns Whether the given token pair is supported by the bridge.
   */
  public async supportsTokenPair(
    l1Token: AddressLike,
    l2Token: AddressLike
  ): Promise<boolean> {
    try {
      const contract = new Contract(
        toAddress(l2Token),
        getContractInterface('KromaMintableERC20'),
        this.messenger.l2Provider
      )

      // Make sure the L1 token matches.
      const remoteL1Token = await contract.REMOTE_TOKEN()

      if (!hexStringEquals(remoteL1Token, toAddress(l1Token))) {
        return false
      }

      // Make sure the L2 bridge matches.
      const remoteL2Bridge = await contract.BRIDGE()
      if (!hexStringEquals(remoteL2Bridge, this.l2Bridge.address)) {
        return false
      }

      return true
    } catch (err) {
      // If the L2 token is not an L2StandardERC20, it may throw an error. If there's a call
      // exception then we assume that the token is not supported. Other errors are thrown. Since
      // the JSON-RPC API is not well-specified, we need to handle multiple possible error codes.
      if (
        err.message.toString().includes('CALL_EXCEPTION') ||
        err.stack.toString().includes('execution reverted')
      ) {
        return false
      } else {
        throw err
      }
    }
  }

  /**
   * Queries the account's approval amount for a given L1 token.
   *
   * @param l1Token The L1 token address.
   * @param l2Token The L2 token address.
   * @param signer Signer to query the approval for.
   * @returns Amount of tokens approved for deposits from the account.
   */
  public async approval(
    l1Token: AddressLike,
    l2Token: AddressLike,
    signer: ethers.Signer
  ): Promise<BigNumber> {
    if (!(await this.supportsTokenPair(l1Token, l2Token))) {
      throw new Error(`token pair not supported by bridge`)
    }

    const token = new Contract(
      toAddress(l1Token),
      getContractInterface('KromaMintableERC20'), // Any ERC20 will do
      this.messenger.l1Provider
    )

    return token.allowance(await signer.getAddress(), this.l1Bridge.address)
  }

  /**
   * Approves a deposit into the L2 chain.
   *
   * @param l1Token The L1 token address.
   * @param l2Token The L2 token address.
   * @param amount Amount of the token to approve.
   * @param signer Signer used to sign and send the transaction.
   * @param opts Additional options.
   * @param opts.overrides Optional transaction overrides.
   * @returns Transaction response for the approval transaction.
   */
  public async approve(
    l1Token: AddressLike,
    l2Token: AddressLike,
    amount: NumberLike,
    signer: Signer,
    opts?: {
      overrides?: Overrides
    }
  ): Promise<TransactionResponse> {
    return signer.sendTransaction(
      await this.populateTransaction.approve(l1Token, l2Token, amount, opts)
    )
  }

  /**
   * Deposits some ETHs into the L2 chain.
   *
   * @param amount Amount of the token to deposit.
   * @param signer Signer used to sign and send the transaction.
   * @param opts Additional options.
   * @param opts.recipient Optional address to receive the funds on L2. Defaults to sender.
   * @param opts.l2GasLimit Optional gas limit to use for the transaction on L2.
   * @param opts.overrides Optional transaction overrides.
   * @returns Transaction response for the deposit transaction.
   */
  public async depositETH(
    amount: NumberLike,
    signer: Signer,
    opts?: {
      recipient?: AddressLike
      l2GasLimit?: NumberLike
      overrides?: PayableOverrides
    }
  ): Promise<TransactionResponse> {
    return signer.sendTransaction(
      await this.populateTransaction.depositETH(amount, opts)
    )
  }

  /**
   * Deposits some ERC20 tokens into the L2 chain.
   *
   * @param l1Token The L1 token address.
   * @param l2Token The L2 token address.
   * @param amount Amount of the token to deposit.
   * @param signer Signer used to sign and send the transaction.
   * @param opts Additional options.
   * @param opts.recipient Optional address to receive the funds on L2. Defaults to sender.
   * @param opts.l2GasLimit Optional gas limit to use for the transaction on L2.
   * @param opts.overrides Optional transaction overrides.
   * @returns Transaction response for the deposit transaction.
   */
  public async depositERC20(
    l1Token: AddressLike,
    l2Token: AddressLike,
    amount: NumberLike,
    signer: Signer,
    opts?: {
      recipient?: AddressLike
      l2GasLimit?: NumberLike
      overrides?: Overrides
    }
  ): Promise<TransactionResponse> {
    return signer.sendTransaction(
      await this.populateTransaction.depositERC20(
        l1Token,
        l2Token,
        amount,
        opts
      )
    )
  }

  /**
   * Withdraws some ETHs back to the L1 chain.
   *
   * @param amount Amount of the token to withdraw.
   * @param signer Signer used to sign and send the transaction.
   * @param opts Additional options.
   * @param opts.recipient Optional address to receive the funds on L1. Defaults to sender.
   * @param opts.overrides Optional transaction overrides.
   * @returns Transaction response for the withdraw transaction.
   */
  public async withdrawETH(
    amount: NumberLike,
    signer: Signer,
    opts?: {
      recipient?: AddressLike
      overrides?: Overrides
    }
  ): Promise<TransactionResponse> {
    return signer.sendTransaction(
      await this.populateTransaction.withdrawETH(amount, opts)
    )
  }

  /**
   * Withdraws some ERC20 tokens back to the L1 chain.
   *
   * @param l1Token The L1 token address.
   * @param l2Token The L2 token address.
   * @param amount Amount of the token to withdraw.
   * @param signer Signer used to sign and send the transaction.
   * @param opts Additional options.
   * @param opts.recipient Optional address to receive the funds on L1. Defaults to sender.
   * @param opts.overrides Optional transaction overrides.
   * @returns Transaction response for the withdraw transaction.
   */
  public async withdrawERC20(
    l1Token: AddressLike,
    l2Token: AddressLike,
    amount: NumberLike,
    signer: Signer,
    opts?: {
      recipient?: AddressLike
      overrides?: Overrides
    }
  ): Promise<TransactionResponse> {
    return signer.sendTransaction(
      await this.populateTransaction.withdrawERC20(
        l1Token,
        l2Token,
        amount,
        opts
      )
    )
  }

  /**
   * Object that holds the functions that generate transactions to be signed by the user.
   * Follows the pattern used by ethers.js.
   */
  populateTransaction = {
    /**
     * Generates a transaction for approving some tokens to deposit into the L2 chain.
     *
     * @param l1Token The L1 token address.
     * @param l2Token The L2 token address.
     * @param amount Amount of the token to approve.
     * @param opts Additional options.
     * @param opts.overrides Optional transaction overrides.
     * @returns Transaction that can be signed and executed to deposit the tokens.
     */
    approve: async (
      l1Token: AddressLike,
      l2Token: AddressLike,
      amount: NumberLike,
      opts?: {
        overrides?: Overrides
      }
    ): Promise<TransactionRequest> => {
      if (!(await this.supportsTokenPair(l1Token, l2Token))) {
        throw new Error(`token pair not supported by bridge`)
      }

      const token = new Contract(
        toAddress(l1Token),
        getContractInterface('KromaMintableERC20'), // Any ERC20 will do
        this.messenger.l1Provider
      )

      return token.populateTransaction.approve(
        this.l1Bridge.address,
        amount,
        opts?.overrides || {}
      )
    },

    /**
     * Generates a transaction for depositing some ETHs into the L2 chain.
     *
     * @param amount Amount of the token to deposit.
     * @param opts Additional options.
     * @param opts.recipient Optional address to receive the funds on L2. Defaults to sender.
     * @param opts.l2GasLimit Optional gas limit to use for the transaction on L2.
     * @param opts.overrides Optional transaction overrides.
     * @returns Transaction that can be signed and executed to deposit the ETHs.
     */
    depositETH: async (
      amount: NumberLike,
      opts?: {
        recipient?: AddressLike
        l2GasLimit?: NumberLike
        overrides?: Overrides
      }
    ): Promise<TransactionRequest> => {
      if (opts?.recipient === undefined) {
        return this.l1Bridge.populateTransaction.bridgeETH(
          opts?.l2GasLimit || 200_000, // Default to 200k gas limit.
          '0x', // No data.
          {
            value: amount,
            ...opts?.overrides,
          }
        )
      } else {
        return this.l1Bridge.populateTransaction.bridgeETHTo(
          toAddress(opts.recipient),
          opts?.l2GasLimit || 200_000, // Default to 200k gas limit.
          '0x', // No data.
          {
            value: amount,
            ...opts?.overrides,
          }
        )
      }
    },

    /**
     * Generates a transaction for depositing some tokens into the L2 chain.
     *
     * @param l1Token The L1 token address.
     * @param l2Token The L2 token address.
     * @param amount Amount of the token to deposit.
     * @param opts Additional options.
     * @param opts.recipient Optional address to receive the funds on L2. Defaults to sender.
     * @param opts.l2GasLimit Optional gas limit to use for the transaction on L2. Defaults to 200k.
     * @param opts.overrides Optional transaction overrides.
     * @returns Transaction that can be signed and executed to deposit the tokens.
     */
    depositERC20: async (
      l1Token: AddressLike,
      l2Token: AddressLike,
      amount: NumberLike,
      opts?: {
        recipient?: AddressLike
        l2GasLimit?: NumberLike
        overrides?: Overrides
      }
    ): Promise<TransactionRequest> => {
      if (!(await this.supportsTokenPair(l1Token, l2Token))) {
        throw new Error(`token pair not supported by bridge`)
      }

      if (opts?.recipient === undefined) {
        return this.l1Bridge.populateTransaction.bridgeERC20(
          toAddress(l1Token),
          toAddress(l2Token),
          amount,
          opts?.l2GasLimit || 200_000, // Default to 200k gas limit.
          '0x', // No data.
          opts?.overrides || {}
        )
      } else {
        return this.l1Bridge.populateTransaction.bridgeERC20To(
          toAddress(l1Token),
          toAddress(l2Token),
          toAddress(opts.recipient),
          amount,
          opts?.l2GasLimit || 200_000, // Default to 200k gas limit.
          '0x', // No data.
          opts?.overrides || {}
        )
      }
    },

    /**
     * Generates a transaction for withdrawing some ETHs back to the L1 chain.
     *
     * @param amount Amount of the token to withdraw.
     * @param opts Additional options.
     * @param opts.recipient Optional address to receive the funds on L1. Defaults to sender.
     * @param opts.overrides Optional transaction overrides.
     * @returns Transaction that can be signed and executed to withdraw the ETHs.
     */
    withdrawETH: async (
      amount: NumberLike,
      opts?: {
        recipient?: AddressLike
        overrides?: Overrides
      }
    ): Promise<TransactionRequest> => {
      if (opts?.recipient === undefined) {
        return this.l2Bridge.populateTransaction.bridgeETH(
          0, // L1 gas not required.
          '0x', // No data.
          {
            value: amount,
            ...opts?.overrides,
          }
        )
      } else {
        return this.l2Bridge.populateTransaction.bridgeETHTo(
          toAddress(opts.recipient),
          0, // L1 gas not required.
          '0x', // No data.
          {
            value: amount,
            ...opts?.overrides,
          }
        )
      }
    },

    /**
     * Generates a transaction for withdrawing some ERC20 tokens back to the L1 chain.
     *
     * @param l1Token The L1 token address.
     * @param l2Token The L2 token address.
     * @param amount Amount of the token to withdraw.
     * @param opts Additional options.
     * @param opts.recipient Optional address to receive the funds on L1. Defaults to sender.
     * @param opts.overrides Optional transaction overrides.
     * @returns Transaction that can be signed and executed to withdraw the ERC20 tokens.
     */
    withdrawERC20: async (
      l1Token: AddressLike,
      l2Token: AddressLike,
      amount: NumberLike,
      opts?: {
        recipient?: AddressLike
        overrides?: Overrides
      }
    ): Promise<TransactionRequest> => {
      if (!(await this.supportsTokenPair(l1Token, l2Token))) {
        throw new Error(`token pair not supported by bridge`)
      }

      if (opts?.recipient === undefined) {
        return this.l2Bridge.populateTransaction.bridgeERC20(
          toAddress(l2Token),
          toAddress(l1Token),
          amount,
          0, // L1 gas not required.
          '0x', // No data.
          opts?.overrides || {}
        )
      } else {
        return this.l2Bridge.populateTransaction.bridgeERC20To(
          toAddress(l2Token),
          toAddress(l1Token),
          toAddress(opts.recipient),
          amount,
          0, // L1 gas not required.
          '0x', // No data.
          opts?.overrides || {}
        )
      }
    },
  }

  /**
   * Object that holds the functions that estimates the gas required for a given transaction.
   * Follows the pattern used by ethers.js.
   */
  estimateGas = {
    /**
     * Estimates gas required to approve some tokens to deposit into the L2 chain.
     *
     * @param l1Token The L1 token address.
     * @param l2Token The L2 token address.
     * @param amount Amount of the token to approve.
     * @param opts Additional options.
     * @param opts.overrides Optional transaction overrides.
     * @returns Gas estimate for the transaction.
     */
    approve: async (
      l1Token: AddressLike,
      l2Token: AddressLike,
      amount: NumberLike,
      opts?: {
        overrides?: CallOverrides
      }
    ): Promise<BigNumber> => {
      return this.messenger.l1Provider.estimateGas(
        await this.populateTransaction.approve(l1Token, l2Token, amount, opts)
      )
    },

    /**
     * Estimates gas required to deposit some ETHs into the L2 chain.
     *
     * @param amount Amount of the token to deposit.
     * @param opts Additional options.
     * @param opts.recipient Optional address to receive the funds on L2. Defaults to sender.
     * @param opts.l2GasLimit Optional gas limit to use for the transaction on L2.
     * @param opts.overrides Optional transaction overrides.
     * @returns Gas estimate for the transaction.
     */
    depositETH: async (
      amount: NumberLike,
      opts?: {
        recipient?: AddressLike
        l2GasLimit?: NumberLike
        overrides?: CallOverrides
      }
    ): Promise<BigNumber> => {
      return this.messenger.l1Provider.estimateGas(
        await this.populateTransaction.depositETH(amount, opts)
      )
    },

    /**
     * Estimates gas required to deposit some ERC20 tokens into the L2 chain.
     *
     * @param l1Token The L1 token address.
     * @param l2Token The L2 token address.
     * @param amount Amount of the token to deposit.
     * @param opts Additional options.
     * @param opts.recipient Optional address to receive the funds on L2. Defaults to sender.
     * @param opts.l2GasLimit Optional gas limit to use for the transaction on L2.
     * @param opts.overrides Optional transaction overrides.
     * @returns Gas estimate for the transaction.
     */
    depositERC20: async (
      l1Token: AddressLike,
      l2Token: AddressLike,
      amount: NumberLike,
      opts?: {
        recipient?: AddressLike
        l2GasLimit?: NumberLike
        overrides?: CallOverrides
      }
    ): Promise<BigNumber> => {
      return this.messenger.l1Provider.estimateGas(
        await this.populateTransaction.depositERC20(
          l1Token,
          l2Token,
          amount,
          opts
        )
      )
    },

    /**
     * Estimates gas required to withdraw some tokens back to the L1 chain.
     *
     * @param amount Amount of the token to withdraw.
     * @param opts Additional options.
     * @param opts.recipient Optional address to receive the funds on L1. Defaults to sender.
     * @param opts.overrides Optional transaction overrides.
     * @returns Gas estimate for the transaction.
     */
    withdrawETH: async (
      amount: NumberLike,
      opts?: {
        recipient?: AddressLike
        overrides?: CallOverrides
      }
    ): Promise<BigNumber> => {
      return this.messenger.l2Provider.estimateGas(
        await this.populateTransaction.withdrawETH(amount, opts)
      )
    },

    /**
     * Estimates gas required to withdraw some ERC20 tokens back to the L1 chain.
     *
     * @param l1Token The L1 token address.
     * @param l2Token The L2 token address.
     * @param amount Amount of the token to withdraw.
     * @param opts Additional options.
     * @param opts.recipient Optional address to receive the funds on L1. Defaults to sender.
     * @param opts.overrides Optional transaction overrides.
     * @returns Gas estimate for the transaction.
     */
    withdrawERC20: async (
      l1Token: AddressLike,
      l2Token: AddressLike,
      amount: NumberLike,
      opts?: {
        recipient?: AddressLike
        overrides?: CallOverrides
      }
    ): Promise<BigNumber> => {
      return this.messenger.l2Provider.estimateGas(
        await this.populateTransaction.withdrawERC20(
          l1Token,
          l2Token,
          amount,
          opts
        )
      )
    },
  }
}
