# zkEVM Prover

<!-- All glossary references in this file. -->

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

A Prover produces a [ZK fault proof][g-zk-fault-proof] that states that a [state][g-state] transition from `S` to
`S'` is valid. It sounds like there are no big differences from validity proofs. That's true. But the point is this is
used to prove the state transition `S` to `S''` is wrong by showing a valid state transition `S` to `S'`.

## zkEVM Proof

See [zkevm-circuits](https://github.com/kroma-network/zkevm-circuits) and
[zkevm-specs](https://github.com/kroma-network/zkevm-specs) for details.

zkEVM proof is a proof that proves the relation between public input and witness. This proof can be verified by the
[ZK Verifier contract](#the-zk-verifier-contract).

Public input and witness are followings:

- Public input: the 2 hashes needed by `PublicInputCircuit`. You can compute them like below:

  ```ts
  import { DataOptions, hexlify } from '@ethersproject/bytes';
  import { Wallet, constants } from 'ethers';
  import { keccak256 } from 'ethers/lib/utils';

  function strip0x(str: string): string {
    if (str.startsWith('0x')) {
      return str.slice(2);
    }
    return str;
  }

  function toFixedBuffer(
    value: string | number,
    length,
    padding = '0',
  ): Buffer {
    const options: DataOptions = {
      hexPad: 'left',
    };
    return hexToBuffer(
      strip0x(hexlify(value, options)).padStart(length * 2, padding),
    );
  }

  async function getDummyTxHash(chainId: number): Promise<string> {
    const sk = hex.toFixedBuffer(1, 32);
    const signer = new Wallet(sk);
    const rlp = await signer.signTransaction({
      nonce: 0,
      gasLimit: 0,
      gasPrice: 0,
      to: constants.AddressZero,
      value: 0,
      data: '0x',
      chainId,
    });
    return keccak256(rlp);
  }

  async function computePublicInput(block: RPCBlock, chainId: number): Promise<[string, string]> {
    const maxTxs = 25;

    const buf = Buffer.concat([
      hex.toFixedBuffer(prevStateRoot, 32),
      hex.toFixedBuffer(block.stateRoot, 32),
      hex.toFixedBuffer(block.withdrawalsRoot ?? 0, 32),
      hex.toFixedBuffer(block.hash, 32),
      hex.toFixedBuffer(block.parentHash, 32),
      hex.toFixedBuffer(block.number, 8),
      hex.toFixedBuffer(block.timestamp, 8),
      hex.toFixedBuffer(block.baseFeePerGas ?? 0, 32),
      hex.toFixedBuffer(block.gasLimit, 8),
      hex.toFixedBuffer(block.transactions.length, 2),
      Buffer.concat(
        block.transactions.map((txHash: string) => {
            return toFixedBuffer(txHash, 32);
        }),
      ),
      Buffer.concat(
        Array(maxTxs - block.transactions.length).fill(
          toFixedBuffer(await getDummyTxHash(chainId), 32),
        ),
      ),
    ]);
    const h = hex.toFixedBuffer(keccak256(buf), 32);
    return [
      '0x' + h.subarray(0, 16).toString('hex'),
      '0x' + h.subarray(16, 32).toString('hex'),
    ];
  }
  ```

- Witness: TODO

> These relations are changed soon after [super circuit] integration.

[super circuit]: https://github.com/kroma-network/zkevm-specs/blob/dev/specs/super_circuit.png

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

See [kroma-prover](https://github.com/kroma-network/kroma-prover) for details.

## RPC

Currently, to request a proof generation, a validator needs to communicate via [gRPC](https://grpc.io/). A validator acts
as a client and prover acts as a server. Because proof generation takes time, it must wait for `FETCHING_TIMEOUT`
seconds.

### Protobuf

See [prover-grpc-proto](https://github.com/kroma-network/prover-grpc-proto) for details.

## Summary of Definitions

### Constants

| Name               | Value | Unit    |
| ------------------ | ----- | ------- |
| `FETCHING_TIMEOUT` | TBD   | seconds |
