# kroma-validator subcommands

This file contains the subcommands for the `kroma-validator` command line tool. It includes
the whole commands used to interact with the validator system. Especially, the commands for
the new validator system with the governance token are included.

## Usage

```bash
kroma-validator [command]
```

The descriptions of whole or particular commands can be checked by running the following command:

```bash
kroma-validator --help
kroma-validator [command] --help
```

## Commands

### Commands for Validator System V2
The following commands are available in Validator System V2:
- `approve` - Approve a contract to spend governance tokens. Default spender is the `ValidatorManager`.
  - `--amount [value]` - _(Required)_ The amount of tokens to approve (in Wei).
  ```bash
  kroma-validator approve --amount 100000000
  ```
- `delegate` - Self-delegate governance tokens to the `ValidatorManager`.
  - `--amount [value]` - _(Required)_ The amount of tokens to delegate (in Wei).
  ```bash
  kroma-validator delegate --amount 100000000
  ```
- `undelegate init` - Initiate the undelegation of governance tokens from the `AssetManager`.
  - `--amount [value]` - _(Required)_ The amount of tokens to undelegate (in Wei).
  ```bash
  kroma-validator undelegate init --amount 100000000
  ```

- `undelegate finalize` - Finalize the undelegation of governance tokens from the `AssetManager`.
Should be called after 1 week has passed from the `undelegate init` command.
  ```bash
  kroma-validator undelegate finalize
  ```
- `claim init` - Initiate the claim of governance tokens from the `AssetManager`.
  - `--amount [value]` - _(Required)_ The amount of tokens to claim (in Wei).
  ```bash
  kroma-validator claim init --amount 100000000
  ```
- `claim finalize` - Finalize the claim of governance tokens from the `AssetManager`. Should be called after 1 week
has passed from the `claim init` command.
  ```bash
  kroma-validator claim finalize
  ```
- `register` - Register a new validator.
  - `--amount [value]` - _(Required)_ The amount of tokens to delegate initially (in Wei).
  - `--commission-rate [value]` - _(Required)_ The initial commission rate of the validator (in %).
  - `--commission-max-rate [value]` - _(Required)_ The max change rate of the commission of the validator (in %).
  ```bash
  kroma-validator register --amount 100000000 --commission-rate 5 --commission-max-rate 5
  ```
- `unjail` - Unjail the validator.
  ```bash
  kroma-validator unjail
  ```
- `changeCommissionRate` - Change the commission rate of the validator.
  - `--commission-rate [value]` - _(Required)_ The new commission rate of the validator (in %).
  ```bash
  kroma-validator changeCommissionRate --commission-rate 5
  ```

### Commands for Validator System V1
The following commands are available in Validator System V1. Note that these commands are **deprecated** and will be removed
in the future after the new validator system is introduced.
- `deposit` - Deposit ETH to `ValidatorPool`.
  - `--amount [value]` - _(Required)_ The amount of tokens to deposit (in Wei).
  ```bash
  kroma-validator deposit --amount 100000000
  ```
- `withdraw` - Withdraw ETH from `ValidatorPool`.
  ```bash
  kroma-validator withdraw
  ```
- `unbond` - Unbond the ETH (including reward) from `ValidatorPool`.
  ```bash
  kroma-validator unbond
  ```
