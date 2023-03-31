# kanvas-node

This is the reference implementation of the [rollup-node spec](../specs/rollup-node.md).

## Compiling

Compile a binary:

```shell
> cd kanvas-node
> go build -o bin/kanvas-node ./cmd
```

## Testing

Run kanvas-node unit tests:

```shell
> cd kanvas-node
> go test ./...
```

Run end-to-end tests:

```shell
> cd e2e
> go test ./...
```

## Running

Options can be reviewed with:

```shell
> ./bin/kanvas-node --help
```

To start syncing the rollup:

Connect to at least one L1 RPC and L2 execution engine:

- L1: use any L1 node / RPC (websocket connection path may differ)
- L2: run the Kanvas fork of geth: [`kanvas-geth`](https://github.com/wemixkanvas/go-ethereum)

```shell
# websockets or IPC preferred for event notifications to improve sync, http RPC works with adaptive polling.
> ./bin/kanvas-node \
  --l1=ws://localhost:8546 --l2=ws//localhost:9001 \
  --rollup.config=./path-to-network-config/rollup.json \
  --rpc.addr=127.0.0.1 \
  --rpc.port=7000
```

## Devnet Genesis Generation

The `kanvas-node` can generate geth compatible `genesis.json` files. These files
can be used with `geth init` to initialize the `StateDB` with accounts, storage,
code and balances. The L2 state must be initialized with predeploy contracts
that exist in the state and act as system level contracts. The `kanvas-node` can
generate a genesis file with these predeploys configured correctly given
hardhat compilation artifacts, hardhat deployment artifacts, a L1 RPC URL
and a deployment config.

The hardhat compilation artifacts are produced by `hardhat compile`. The native
hardhat compiler toolchain produces them by default and the
`@foundry-rs/hardhat` plugin can also produce them when using the foundry
compiler toolchain. They can usually be found in an `artifacts` directory.

The hardhat deployment artifacts are produced by running `hardhat deploy`. These
exist to make it easy to track deployments of smart contract systems over time.
They can usually be found in a `deployments` directory.

The deployment config contains all of the information required to deploy the
system. It can be found in `packages/contracts/deploy-config`. Each
deploy config file can be JSON or TypeScript, although only JSON files are
supported by the `kanvas-node`. The network name must match the name of the file
in the deploy config directory.

Example usage:

```shell
> ./bin/kanvas-node genesis devnet-l2 \
   --artifacts $CONTRACTS/artifacts \
   --network $NETWORK \
   --deployments $CONTRACTS/deployments \
   --deploy-config $CONTRACTS/deploy-config \
   --rpc-url http://localhost:8545
```
