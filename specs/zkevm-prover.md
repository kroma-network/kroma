# zkEVM Prover

<!-- All glossary references in this file. -->

[g-l2-output]: glossary.md#l2-output-root
[g-state]: glossary.md#state
[g-zk-fault-proof]: glossary.md#zk-fault-proof

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Overview](#overview)
- [zkEVM Proof](#zkevm-proof)
- [the ZK Verifier Contract](#the-zk-verifier-contract)
- [Implementation](#implementation)
- [RPC](#rpc)
  - [Protobuf](#protobuf)
- [Summary of Definitions](#summary-of-definitions)
  - [Constants](#constants)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

Prover produces a so called [ZK fault proof][g-zk-fault-proof] that states a [state][g-state] transition from `S` to
`S'` is valid. It sounds like there's no big differences from validity proof. That's true. But the point is this is used
to prove the state transition `S` to `S''` is wrong by showing a valid state transition `S` to `S'`.

## zkEVM Proof

See [zkevm-circuits](https://github.com/wemixkanvas/zkevm-circuits) and
[zkevm-specs](https://github.com/wemixkanvas/zkevm-specs) for details.

zkEVM proof is a proof that proves the relation between public input and witness. This proof can be verified by
[ZK Verifier contract](#the-zk-verifier-contract).

Public input and witness are followings:

- Public input
  - `TODO`: given from calldata.
  - `previous state root`: [L2 output root][g-l2-output] is recorded in
    [L2OutputOracle.sol](../packages/contracts/contracts/L1/L2OutputOracle.sol) and retrieved state root by submitting
    preimage of L2 output.
  - `state root`: same as above.
- Witness
  - `transactions`: L2 transactions.
  - `states`: L2 states that comprises state root.

> These relations are changed soon after [super circuit] integration.

[super circuit]: https://github.com/wemixkanvas/zkevm-specs/blob/dev/specs/super_circuit.png

## the ZK Verifier Contract

The ZK Verifier contract implements `verify` function like followings:

```solidity
interface ZKVerifier {
    function verify(
        uint256[] calldata proof,
        uint256[] calldata target_circuit_final_pair
    ) public view returns (bool);
}
```

## Implementation

See [kanvas-prover](https://github.com/wemixkanvas/kanvas-prover) for details.

## RPC

Currently, to request a proof generation, validator needs to communicate via [gRPC](https://grpc.io/). A validator acts
as a client and prover acts as a server. Because proof generation takes too long, it must wait for `FETCHING_TIMEOUT`
seconds.

### Protobuf

See [kanvas-grpc-proto](https://github.com/wemixkanvas/kanvas-grpc-proto) for details.

## Summary of Definitions

### Constants

| Name               | Value | Unit    |
| ------------------ | ----- | ------- |
| `FETCHING_TIMEOUT` | TBD   | seconds |
