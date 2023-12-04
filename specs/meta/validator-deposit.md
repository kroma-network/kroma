<!-- DOCTOC SKIP -->

# Validator Deposit Guide

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Deposit into `ValidatorPool`](#deposit-into-validatorpool)
- [Withdraw from `ValidatorPool`](#withdraw-from-validatorpool)
- [Try unbond in `ValidatorPool`](#try-unbond-in-validatorpool)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

We want the [validator](../../kroma-validator/) role to be decentralized. Like how the PoS mechanism works, to
achieve this, a validator needs to bond ETH at every [output
submission](../validator.md#submitting-l2-output-commitments). When submitting an output, the amount of bond
specified by the validator is automatically bonded from [the ETH the validator has deposited into the
`ValidatorPool`](#deposit-into-validatorpool). The bonded ETH is automatically unbonded when the submitted output is
finalized. The finalization of the output is checked when the next outputs are submitted, or if the finalization period
of the submitted output has passed, you can directly [trigger unbond by using the `unbond` command](
  #try-unbond-in-validatorpool). For more details about submitting an output as a validator, see
[here](../validator.md).

This guide teaches you how to deposit, withdraw, or try to unbond in `ValidatorPool` via CLI. You can find the proxy
address of `ValidatorPool` on Sepolia [here](../../packages/contracts/deployments/sepolia/ValidatorPoolProxy.json),
on Mainnet TBD.

## Deposit into `ValidatorPool`

```shell
> cd kroma-validator
```

```shell
> go run ./cmd/main.go \
  --valpool-address <validator-pool-address> \ # must be set
  --l1-eth-rpc <l1-eth-rpc> \
  --mnemonic <mnemonic> \
  --hd-path <hd-path> \
  --rollup-rpc "" \ # empty required flags
  --l2oo-address "" \
  --colosseum-address "" \
  --challenger.poll-interval 0s \
  deposit \
  --amount <amount-wei> # must be set
```

## Withdraw from `ValidatorPool`

```shell
> cd kroma-validator
```

```shell
> go run ./cmd/main.go \
  --valpool-address <validator-pool-address> \ # must be set
  --l1-eth-rpc <l1-eth-rpc> \
  --mnemonic <mnemonic> \
  --hd-path <hd-path> \
  --rollup-rpc "" \ # empty required flags
  --l2oo-address "" \
  --colosseum-address "" \
  --challenger.poll-interval 0s \
  withdraw \
  --amount <amount-wei> # must be set
```

## Try unbond in `ValidatorPool`

```shell
> cd kroma-validator
```

```shell
> go run ./cmd/main.go \
  --valpool-address <validator-pool-address> \ # must be set
  --l1-eth-rpc <l1-eth-rpc> \
  --mnemonic <mnemonic> \
  --hd-path <hd-path> \
  --rollup-rpc "" \ # empty required flags
  --l2oo-address "" \
  --colosseum-address "" \
  --challenger.poll-interval 0s \
  unbond
```
