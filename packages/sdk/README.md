# @kroma-network/sdk

The `@kroma-network/sdk` package provides a set of tools for interacting with Kroma.

## Installation

```shell
> npm install @kroma-network/sdk
```

## Using the SDK

### CrossChainMessenger

The [`CrossChainMessenger`] class simplifies the process of moving assets and data between Ethereum and Kroma.
You can use this class to, for example, initiate a withdrawal of ERC20 tokens from Kroma back to Ethereum, accurately
track when the withdrawal is ready to be finalized on Ethereum, and execute the finalization transaction after
the challenge period has elapsed.
The `CrossChainMessenger` can handle deposits and withdrawals of ETH and any ERC20-compatible token.
The `CrossChainMessenger` automatically connects to all relevant contracts so complex configuration is not necessary.

[`CrossChainMessenger`]: ./src/cross-chain-messenger.ts

### L2Provider and related utilities

The Kroma SDK includes [various utilities] for handling Kroma's [transaction fee model].
For instance, [`estimateTotalGasCost`] will estimate the total cost (in wei) to send at transaction on Kroma including
both the L2 execution cost and the L1 data cost. You can also use the [`asL2Provider`] function to wrap an ethers
Provider object into an `L2Provider` which will have all of these helper functions attached.

[various utilities]: ./src/l2-provider.ts
[transaction fee model]: https://community.optimism.io/docs/developers/build/transaction-fees/
[`estimateTotalGasCost`]: https://sdk.optimism.io/modules.html#estimateTotalGasCost
[`asL2Provider`]: https://sdk.optimism.io/modules.html#asL2Provider

### Other utilities

The SDK contains other useful helper functions and constants.
For a complete list, refer to the auto-generated [SDK documentation](https://sdk.optimism.io/)
