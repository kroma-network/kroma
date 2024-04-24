# kroma-validator subcommands

This file contains the subcommands for the `kroma-validator` command line tool. It includes
the whole commands used to interact with the validator system. Especially, the commands for
the new validator system with the asset token are included.

## Usage

You can use these subcommands via running kroma-validator container.

```bash
docker compose exec kroma-validator validator [command]
```

The descriptions of whole or particular commands can be checked by running the following command:

```bash
docker compose exec kroma-validator validator --help
docker compose exec kroma-validator validator [command] --help
```

## Commands

### Commands for Validator System V1
The following commands are available in Validator System V1. Note that these commands will be **deprecated** and
removed in the future after the new validator system is introduced. You can still use `withdraw` and `unbond` even
after the transition to the Validator System V2.
- `deposit` - Deposit ETH to `ValidatorPool`.
  - `--amount [value]` - _(Required)_ The amount of ETH to deposit (in Wei).
  ```bash
  docker compose exec kroma-validator validator deposit --amount 100000000
  ```
- `withdraw` - Withdraw ETH from `ValidatorPool`.
  - `--amount [value]` - _(Required)_ The amount of ETH to withdraw (in Wei).
  ```bash
  docker compose exec kroma-validator validator withdraw
  ```
- `unbond` - Manually unbond the bonded ETH from `ValidatorPool` (including reward distribution).
  ```bash
  docker compose exec kroma-validator validator unbond
  ```

### Commands for Validator System V2 (_EXPERIMENTAL_)
The following commands are available in Validator System V2:
- `register` - Register a new validator.
  - `--amount [value]` - _(Required)_ The amount of tokens to delegate initially (in Wei).
  - `--commission-rate [value]` - _(Required)_ The initial commission rate of the validator (in %).
  - `--commission-max-change-rate [value]` - _(Required)_ The max change rate of the commission of the validator (in %).
  ```bash
  docker compose exec kroma-validator validator register --amount 100000000 --commission-rate 5 --commission-max-change-rate 5
  ```
- `activate` - Activate the validator to be eligible to submit output roots and create challenges.
  ```bash
  docker compose exec kroma-validator validator activate
  ```
- `unjail` - Unjail the validator.
  ```bash
  docker compose exec kroma-validator validator unjail
  ```
- `changeCommissionRate` - Change the commission rate of the validator.
  - `--commission-rate [value]` - _(Required)_ The new commission rate of the validator (in %).
  ```bash
  docker compose exec kroma-validator validator changeCommissionRate --commission-rate 5
  ```

- `delegate` - Self-delegate asset tokens to the `AssetManager`.
  - `--amount [value]` - _(Required)_ The amount of tokens to delegate (in Wei).
  ```bash
  docker compose exec kroma-validator validator delegate --amount 100000000
  ```
- `undelegate init` - Initiate the undelegations of asset tokens from the `AssetManager`.
  - `--amount [value]` - _(Required)_ The amount of tokens to undelegate (in Wei).
  ```bash
  docker compose exec kroma-validator validator undelegate init --amount 100000000
  ```

- `undelegate finalize` - Finalize the undelegations of asset tokens from the `AssetManager`.
Should be called after 1 week has passed from the `undelegate init` command.
  ```bash
  docker compose exec kroma-validator validator undelegate finalize
  ```
- `claim init` - Initiate the claim of validator reward from the `AssetManager`.
  - `--amount [value]` - _(Required)_ The amount of tokens to claim (in Wei).
  ```bash
  docker compose exec kroma-validator validator claim init --amount 100000000
  ```
- `claim finalize` - Finalize the claim of validator reward from the `AssetManager`. Should be called after 1 week
has passed from the `claim init` command.
  ```bash
  docker compose exec kroma-validator validator claim finalize
  ```
