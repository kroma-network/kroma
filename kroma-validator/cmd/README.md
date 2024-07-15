# kroma-validator subcommands

This file contains the subcommands for the `kroma-validator` command line tool. It includes
the whole commands used to interact with the validator system. Especially, the commands for
the new validator system with the asset token are included.

## Usage

You can use these subcommands in the running kroma-validator docker container.

```bash
kroma-validator [command]
```

The descriptions of whole or particular commands can be checked by running the following command:

```bash
kroma-validator --help
kroma-validator [command] --help
```

## Commands

### Commands for Validator System V1

The following commands are available in Validator System V1. Note that these commands will be **deprecated** and
removed in the future after the new validator system is introduced. You can still use `withdraw` and `unbond` even
after the transition to the Validator System V2.

- `deposit` - Deposit ETH to `ValidatorPool`.
  - `--amount [value]` - _(Required)_ The amount of ETH to deposit (in Wei).

  ```bash
  kroma-validator deposit --amount 100000000
  ```

- `withdraw` - Withdraw ETH from `ValidatorPool`.
  - `--amount [value]` - _(Required)_ The amount of ETH to withdraw (in Wei).

  ```bash
  kroma-validator withdraw --amount 100000000
  ```

- `withdrawTo` - Withdraw ETH from `ValidatorPool` to specific address.
  - `--address` - _(Required)_ The address to receive withdrawn ETH.
  - `--amount [value]` - _(Required)_ The amount of ETH to withdraw (in Wei).

  ```bash
  kroma-validator withdrawTo --address 0x0000000000000000000000000000000000000001 --amount 100000000
  ```

- `unbond` - Manually unbond the bonded ETH from `ValidatorPool` (including reward distribution).

  ```bash
  kroma-validator unbond
  ```

### Commands for Validator System V2 (_EXPERIMENTAL_)

The following commands are available in Validator System V2:

- `register` - Register as new validator to `ValidatorManager`.
  - `--amount [value]` - _(Required)_ The amount of tokens to delegate initially (in Wei).
  - `--commission-rate [value]` - _(Required)_ The initial commission rate of the validator (in %).
  - `--withdraw-account [value]` - _(Required)_ The address to withdraw deposited asset token.

  ```bash
  kroma-validator register --amount 100000000 --commission-rate 5 --withdraw-account 0x0000000000000000000000000000000000000001
  ```

- `activate` - Activate the validator in `ValidatorManager` to be eligible to submit output roots and create challenges.

  ```bash
  kroma-validator activate
  ```

- `unjail` - Unjail the validator in `ValidatorManager`.

  ```bash
  kroma-validator unjail
  ```

- `changeCommissionRate init` - Initiate the commission rate change of the validator in `ValidatorManager`.
  - `--commission-rate [value]` - _(Required)_ The new commission rate of the validator (in %).

  ```bash
  kroma-validator changeCommissionRate init --commission-rate 5
  ```

- `changeCommissionRate finalize` - Finalize the commission rate change of the validator in `ValidatorManager`.

  ```bash
  kroma-validator changeCommissionRate finalize
  ```

- `delegate` - Self-delegate asset tokens to the `AssetManager`.
  - `--amount [value]` - _(Required)_ The amount of tokens to delegate (in Wei).

  ```bash
  kroma-validator delegate --amount 100000000
  ```

- `undelegate init` - Initiate the undelegation of asset tokens from the `AssetManager`.
  - `--amount [value]` - _(Required)_ The amount of tokens to undelegate (in Wei).

  ```bash
  kroma-validator undelegate init --amount 100000000
  ```

- `undelegate finalize` - Finalize the undelegations of asset tokens from the `AssetManager`.
  Should be called after 1 week has passed from the `undelegate init` command.

  ```bash
  kroma-validator undelegate finalize
  ```

- `claim init` - Initiate the claim of validator reward from the `AssetManager`.
  - `--amount [value]` - _(Required)_ The amount of tokens to claim (in Wei).

  ```bash
  kroma-validator claim init --amount 100000000
  ```

- `claim finalize` - Finalize the claim of validator reward from the `AssetManager`. Should be called after 1 week
  has passed from the `claim init` command.

  ```bash
  kroma-validator claim finalize
  ```
