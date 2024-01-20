# Differences from Optimism

<!-- All glossary references in this file. -->

[g-l2-output-root]: glossary.md#l2-output-root
[g-mpt]: glossary.md#merkle-patricia-trie
[g-zktrie]: glossary.md#zk-trie
[g-zk-fault-proof]: glossary.md#zk-fault-proof
[g-system-config]: glossary.md#system-configuration
[g-validation-rewards]: validator.md#validation-rewards
[g-output-payload-v0]: validator.md#output-payloadversion-0

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Nodes](#nodes)
  - [Verifier -> Validator](#verifier---validator)
  - [Compositions](#compositions)
  - [Adding field to System Configuration](#adding-field-to-system-configuration)
  - [Adding field to Output Payload](#adding-field-to-output-payload)
- [Geth](#geth)
- [Validator](#validator)
  - [ZK fault proof](#zk-fault-proof)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Nodes

There are two types of network participants in the OP Stack:

- [Sequencers](https://github.com/ethereum-optimism/optimism/blob/develop/specs/introduction.md#sequencers) consolidate
 users' on/off chain transactions into blocks. They submit checkpoint outputs as well as batch transactions.
- [Verifiers](https://github.com/ethereum-optimism/optimism/blob/develop/specs/introduction.md#verifiers) verify rollup
   integrity and dispute invalid assertions.

It is crucial to have at least one honest verifier who can verify the integrity of the rollup chain to ensure the
ongoing security of the network. However, there exists a well-known obstacle known as the 'Verifier's Dilemma' that
poses a threat to the security of optimistic rollups by introducing disincentives in such scenarios.

To resolve the 'Verifier's Dilemma', we have devised an incentive mechanism that motivates node operators to actively
participate in the Kroma network. As part of this redesign, we have separated the responsibility of submitting
checkpoint outputs from `sequencers` and assigned it to `verifiers`. As a result, these participants have been renamed
to reflect these role changes and the future direction of Kroma's decentralization. For more detailed information about
our decentralization scheme, please refer to
[this article](https://medium.com/@kroma-network/the-road-to-kromas-decentralization-38f8e46df442)
on the Kroma blog.

### Verifier -> Validator

We utilize the term `validator` to denote a participant who is responsible for submitting the
[L2 output root][g-l2-output-root] and validating its accuracy by either submitting dispute challenges (during the
optimistic rollup phase) or providing ZK validity proofs (during the ZK rollup phase). This concept bears a resemblance
to how L1 validators cast FFG votes at each epoch.

### Compositions

Kroma maintains the modular architecture of the OP Stack, with various components communicating through Json RPC calls.
As part of the transition from `verifier` to `validator`, we have renamed the `proposer` (op-proposer) component of the
OP Stack to `validator` (kroma-validator) and made necessary modifications to the code to handle the dispute challenge
processes.

The followings are components that are used to run different types of nodes:

| Node        | Components                                                           |
|-------------|----------------------------------------------------------------------|
| `Sequencer` | `L2 EL client` + `L2 CL client` + `kroma-batcher`                    |
| `Validator` | `L2 EL client` + `L2 CL client` + `kroma-validator` + `kroma-prover` |
| `Full node` | `L2 EL client` + `L2 CL client`                                      |

**NOTE:** Here `L2 EL client` means `kroma-geth` and `L2 CL client` means `kroma-node`. `L2 EL client` can
be expanded to other clients for pragmatic decentralization.

### Adding field to System Configuration

The `ValidatorRewardScalar` field was added to [system configuration][g-system-config].

```
type L1BlockInfo struct {
    Number    uint64
    Time      uint64
    BaseFee   *big.Int
    BlockHash common.Hash
    // Not strictly a piece of L1 information. Represents the number of L2 blocks since the start of the epoch,
    // i.e. when the actual L1 info was first introduced.
    SequenceNumber uint64
    // BatcherHash version 0 is just the address with 0 padding to the left.
    BatcherAddr   common.Address
    L1FeeOverhead eth.Bytes32
    L1FeeScalar   eth.Bytes32
    // [Kroma: START]
    ValidatorRewardScalar eth.Bytes32
    // [Kroma: END]
}
```

<pre>
<a href="https://github.com/kroma-network/kroma/blob/dev/op-node/rollup/derive/l1_block_info.go">Code link here</a>
</pre>
This value is set via the `SystemConfig` contract on L1 and passed through the L2 derivation process and used as an
ingredient in the reward calculation. (Detailed calculations : [Validation Rewards][g-validation-rewards])

### Adding field to Output Payload

The `next_block_hash` field was added to [Output Payload][g-output-payload-v0].

```
type OutputV0 struct {
    StateRoot                Bytes32
    MessagePasserStorageRoot Bytes32
    BlockHash                common.Hash
    // [Kroma: START]
    NextBlockHash common.Hash
    // [Kroma: END]
}
```

<pre>
<a href="https://github.com/kroma-network/kroma/blob/dev/op-service/eth/output.go">Code here</a>
</pre>
This value is used as an additional material for the [verification process][g-zk-fault-proof] of the fault
proof system.
It is used to validate the relationship between the Source OutputRootProof and Dest OutputRootProof, and the validation
of the public input.

## Geth

To prepare for migration to ZK Rollup, we use a [ZK Trie][g-zktrie] to represent state. Currently, this makes
the chain slower than [Merkle Patrica Trie][g-mpt]. As the bottleneck is the time to produce ZK proof right now,
we adopt it from [Scroll]. When we overcome the proof generation time problem, we will smoothly migrate state
without a hard fork or huge changes. Thus, you might face an unexpected result when retrieving JSON-RPC such as
`eth_getProof`.

Additionally, to produce a zkEVM proof, geth should return so called `execution trace` via JSON-RPC
`kroma_getBlockTraceByNumberOrHash` which provides zkEVM prover with data as a witness.

[scroll]: https://scroll.io/

## Validator

### ZK fault proof

Instead of [cannon], Kroma uses zkEVM for [ZK fault proof][g-zk-fault-proof].

[cannon]: https://github.com/ethereum-optimism/cannon
