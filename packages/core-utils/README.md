# @wemixkanvas/core-utils

## What is this?

`@wemixkanvas/core-utils` contains the Kanvas core utilities.

## Getting started

### Building and usage

After cloning and switching to the repository, install dependencies:

```shell
> yarn
```

Use the following commands to build, use, test, and lint:

```shell
> yarn build
> yarn start
> yarn test
> yarn lint
```

### L2 Fees

`TxGasLimit` can be used to `encode` and `decode` the L2 Gas Limit
locally.

```typescript
import { TxGasLimit } from '@wemixkanvas/core-utils'
import { JsonRpcProvider } from 'ethers'

const L2Provider = new JsonRpcProvider(L2_JSON_RPC_URL)
const L1Provider = new JsonRpcProvider(L1_JSON_RPC_URL)

const l2GasLimit = await L2Provider.send('eth_estimateExecutionGas', [tx])
const l1GasPrice = await L1Provider.getGasPrice()

const encoded = TxGasLimit.encode({
  data: '0x',
  l1GasPrice,
  l2GasLimit,
  l2GasPrice: 10000000,
})

const decoded = TxGasLimit.decode(encoded)
assert(decoded.eq(gasLimit))

const estimate = await L2Provider.estimateGas()
assert(estimate.eq(encoded))
```
