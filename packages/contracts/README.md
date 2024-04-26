# Kroma Smart Contracts

This package contains the smart contracts that compose the on-chain component of Kroma.
You can find detailed specifications for the contracts contained within this package [here](https://specs.kroma.network).

## Contracts Overview

### Contracts deployed to L1

| Name                                                                                                          | Proxy Type                                 | Description |
|---------------------------------------------------------------------------------------------------------------|--------------------------------------------|-------------|
| [`L1CrossDomainMessenger`](https://specs.kroma.network/protocol/messengers.html)                              | [`Proxy`](./contracts/universal/Proxy.sol) |             |
| High-level interface for sending messages to and receiving messages from Kroma                                |                                            |             |
| [`L1StandardBridge`](https://specs.kroma.network/protocol/bridges.html)                                       | [`Proxy`](./contracts/universal/Proxy.sol) |             |
| Standardized system for transfering ERC20 tokens to/from Kroma                                                |                                            |             |
| [`L2OutputOracle`](https://specs.kroma.network/protocol/validator.html#l2-output-oracle-smart-contract)       | [`Proxy`](./contracts/universal/Proxy.sol) |             |
| Stores commitments to the state of Kroma which can be used by contracts on L1 to access L2 state              |                                            |             |
| [`KromaPortal`](https://specs.kroma.network/protocol/deposits.html#deposit-contract)                          | [`Proxy`](./contracts/universal/Proxy.sol) |             |
| Low-level message passing interface                                                                           |                                            |             |
| [`KromaMintableERC20Factory`](https://specs.kroma.network/protocol/predeploys.html#kromamintableerc20factory) | [`Proxy`](./contracts/universal/Proxy.sol) |             |
| Deploys standard `KromaMintableERC20` tokens that are compatible with either `StandardBridge`                 |                                            |             |
| [`ProxyAdmin`](https://specs.kroma.network/protocol/predeploys.html#proxyadmin)                               | -                                          |             |
| Contract that can upgrade L1 contracts                                                                        |                                            |             |

### Contracts deployed to L2

| Name                                                                                                          | Proxy Type                                 | Description                                                                                   |
|---------------------------------------------------------------------------------------------------------------|--------------------------------------------|-----------------------------------------------------------------------------------------------|
| [`GasPriceOracle`](https://specs.kroma.network/protocol/predeploys.html#gaspriceoracle)                       | [`Proxy`](./contracts/universal/Proxy.sol) | Stores L2 gas price configuration values                                                      |
| [`L1Block`](https://specs.kroma.network/protocol/predeploys.html#l1block)                                     | [`Proxy`](./contracts/universal/Proxy.sol) | Stores L1 block context information (e.g., latest known L1 block hash)                        |
| [`L2CrossDomainMessenger`](https://specs.kroma.network/protocol/predeploys.html#l2crossdomainmessenger)       | [`Proxy`](./contracts/universal/Proxy.sol) | High-level interface for sending messages to and receiving messages from L1                   |
| [`L2StandardBridge`](https://specs.kroma.network/protocol/predeploys.html#l2standardbridge)                   | [`Proxy`](./contracts/universal/Proxy.sol) | Standardized system for transferring ERC20 tokens to/from L1                                  |
| [`L2ToL1MessagePasser`](https://specs.kroma.network/protocol/predeploys.html#l2tol1messagepasser)             | [`Proxy`](./contracts/universal/Proxy.sol) | Low-level message passing interface                                                           |
| [`ValidatorRewardVault`](https://specs.kroma.network/protocol/predeploys.html#validatorrewardvault)           | [`Proxy`](./contracts/universal/Proxy.sol) | Vault for L2 transaction fees                                                                 |
| [`KromaMintableERC20Factory`](https://specs.kroma.network/protocol/predeploys.html#kromamintableerc20factory) | [`Proxy`](./contracts/universal/Proxy.sol) | Deploys standard `KromaMintableERC20` tokens that are compatible with either `StandardBridge` |
| [`L2ProxyAdmin`](https://specs.kroma.network/protocol/predeploys.html#proxyadmin)                             | -                                          | Contract that can upgrade L2 contracts when sent a transaction from L1                        |

## Installation

We export contract ABIs, contract source code, and contract deployment information for this package via `npm`:

```shell
> npm install @kroma/contracts
```

## Development

### Dependencies

We work on this repository with a combination of [Hardhat](https://hardhat.org) and [Foundry](https://getfoundry.sh/).

1. Install Foundry by following [the instructions located here](https://getfoundry.sh/).
   A specific version must be used.

   ```shell
   > foundryup --version nightly-f625d0fa7c51e65b4bf1e8f7931cd1c6e2e285e9
   ```

2. Install node modules with pnpm and Node.js:

   ```shell
   > pnpm i
   ```

### Build

```shell
> pnpm build
```

### Test

```shell
> pnpm test
```

#### Running Echidna tests

You must have [Echidna](https://github.com/crytic/echidna) installed.

Contracts targeted for Echidna testing are located in `./contracts/echidna`
Each target contract is tested with a separate pnpm command, for example:

```shell
> pnpm echidna:aliasing
```

### Deployment

#### Configuration

1. Create or modify a file `<network-name>.ts` inside of the [`deploy-config`](./deploy-config/) folder.
2. Fill out this file according to the `deployConfigSpec` located inside of the
   [`hardhat.config.ts`](./hardhat.config.ts).
3. Optionally: Run `npx hardhat generate-deploy-config --network <network-name>` to generate the associated JSON
   file. This is required if using `utils/chain-ops`.

#### Execution

1. Copy `.env.example` into `.env`
2. Fill out the `L1_RPC` and `PRIVATE_KEY_DEPLOYER` environment variables in `.env`
3. Run `npx hardhat deploy --network <network-name>` to deploy the L1 contracts
4. Run `npx hardhat etherscan-verify --network <network-name> --sleep` to verify contracts on Etherscan

## Tools

### Layout Locking

We use a system called "layout locking" as a safety mechanism to prevent certain contract variables from being moved to
different storage slots accidentally.
To lock a contract variable, add it to the `layout-lock.json` file which has the following format:

```json
{
  "MyContractName": {
    "myVariableName": {
      "slot": 1,
      "offset": 0,
      "length": 32
    }
  }
}
```

With the above config, the `validate-spacers` hardhat task will check that we have a contract called `MyContractName`,
that the contract has a variable named `myVariableName`, and that the variable is in the correct position
as defined in the lock file.
You should add things to the `layout-lock.json` file when you want those variables to **never** change.
Layout locking should be used in combination with diffing the `.storage-layout` file in CI.

## Standards and Conventions

### Style

#### Comments

We use [Seaport](https://github.com/ProjectOpenSea/seaport/blob/main/contracts/Seaport.sol)-style comments with
some minor modifications.
Some basic rules:

- Always use `@notice` since it has the same general effect as `@dev` but avoids confusion about when to use one over
  the other.
- Include a newline between `@notice` and the first `@param`.
- Include a newline between `@param` and the first `@return`.
- Use a line-length of 100 characters.

We also have the following custom tags:

- `@custom:proxied`: Add to a contract whenever it's meant to live behind a proxy.
- `@custom:upgradeable`: Add to a contract whenever it's meant to be used in an upgradeable contract.
- `@custom:semver`: Add to a constructor to indicate the version of a contract.

#### Errors

- Use `require` statements when making simple assertions.
- Use `revert` if throwing an error where an assertion is not being made (no custom errors).
- Error strings MUST have the format `"{ContractName}: {message}"` where `message` is a lower case string.

#### Function Parameters

- Function parameters should be prefixed with an underscore.

#### Event Parameters

- Event parameters should NOT be prefixed with an underscore.

#### Spacers

We use spacer variables to account for old storage slots that are no longer being used.
The name of a spacer variable MUST be in the format `spacer_<slot>_<offset>_<length>` where `<slot>` is the original
storage slot number, `<offset>` is the original offset position within the storage slot,
and `<length>` is the original size of the variable.
Spacers MUST be `private`.

### Proxy by Default

All contracts should be assumed to live behind proxies (except in certain special circumstances).
This means that new contracts MUST be built under the assumption of upgradeability.
We use a minimal [`Proxy`](./contracts/universal/Proxy.sol) contract designed to be owned
by a corresponding [`ProxyAdmin`](./contracts/universal/ProxyAdmin.sol)
which follow the interfaces of OpenZeppelin's `Proxy` and `ProxyAdmin` contracts, respectively.

Unless explicitly discussed otherwise, you MUST include the following basic upgradeability pattern for each new
implementation contract:

1. Extend OpenZeppelin's `Initializable` base contract.
2. Include a `uint8 public constant VERSION = X` at the TOP of your contract.
3. Include a function `initialize` with the modifier `reinitializer(VERSION)`.
4. In the `constructor`, set any `immutable` variables and call the `initialize` function for setting mutables.

### Versioning

All (non-library and non-abstract) contracts MUST extend the `Semver` base contract which exposes a `version()` function
that returns a semver-compliant version string.
Before mainnet release, the development process the `Semver` value for all contracts SHOULD return `0.1.0`

After the mainnet release, contracts MUST use the following versioning scheme:

- `patch` releases are to be used only for changes that do NOT modify contract bytecode (such as updating comments).
- `minor` releases are to be used for changes that modify bytecode OR changes that expand the contract ABI provided that
  these changes do NOT break the existing interface.
- `major` releases are to be used for changes that break the existing contract interface OR changes
  that modify the security model of a contract.

#### Exceptions

We have made an exception to the `Semver` rule for the `WETH` contract to avoid making changes to a well-known, simple,
and recognizable contract.

### Tests

Tests are written using Foundry.

All test contracts and functions should be organized and named according to the following guidelines.

These guidelines are also encoded in a script which can be run with:

```shell
> ts-node scripts/forge-test-names.ts
```

*Note: This is a work in progress, not all test files are compliant with these guidelines.*

#### Organizing Principles

- Solidity `contract`s are used to organize the test suite similar to how mocha uses describe.
- Every non-trivial state changing function should have a separate contract for happy and sad path
  tests. This helps to make it very obvious where there are not yet sad path tests.
- Simpler functions like getters and setters are grouped together into test contracts.

#### Test function naming convention

Test function names are split by underscores, into 3 or 4 parts.
An example function name is `test_onlyOwner_callerIsNotOwner_reverts()`.

The parts are: `[method]_[FunctionName]_[reason]_[success]`, where:

- `[method]` is either `test`, `testFuzz`, or `testDiff`
- `[FunctionName]` is the name of the function or higher level behavior being tested.
- `[reason]` is an optional description for the behavior being tested.
- `[status]` must be one of:
  - `succeeds`: used for most happy path cases
  - `reverts`: used for most sad path cases
  - `works`: used for tests which include a mix of happy and sad assertions (these should be broken up if possible)
  - `fails`: used for tests which 'fail' in some way other than reverting
  - `benchmark`: used for tests intended to establish gas costs

#### Contract Naming Conventions

Test contracts should be named one of the following according to their use:

- `TargetContract_Init` for contracts that perform basic setup to be reused in other test contracts.
- `TargetContract_Function_Test` for contracts containing happy path tests for a given function.
- `TargetContract_Function_TestFail` for contracts containing sad path tests for a given function.

## Withdrawing From Fee Vaults

See the file `scripts/FeeVaultWithdrawal.s.sol` to withdraw from the L2 fee vaults. It includes
instructions on how to run it. `foundry` is required.
