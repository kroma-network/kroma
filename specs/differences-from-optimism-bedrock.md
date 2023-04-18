# Differences from Optimism Bedrock

<!-- All glossary references in this file. -->

[g-l2-output-root]: glossary.md#l2-output-root
[g-mpt]: glossary.md#merkle-patricia-trie
[g-zktrie]: glossary.md#zk-trie
[g-zk-fault-proof]: glossary.md#zk-fault-proof

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Nodes](#nodes)
  - [Overview](#overview)
  - [Sequencer -> Proposer](#sequencer---proposer)
  - [Proposer -> Validator](#proposer---validator)
- [Geth](#geth)
- [Validator](#validator)
  - [ZK fault proof](#zk-fault-proof)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Nodes

### Overview

Kanvas is composed of 3 nodes: `proposer`, `validator` and `vanilla`.  The following are components that
are used to run each node:

| Node        | Components                                                                          |
|-------------|-------------------------------------------------------------------------------------|
| `proposer`  | `L2 EL client` + `L2 CL client` + `kanvas-batcher`                                  |
| `validator` | `L2 EL client` + `L2 CL client` + `kanvas-validator` + (optionally `kanvas-prover`) |
| `vanilla`   | `L2 EL client` + `L2 CL client`                                                     |

`kanvas-prover` is only needed when `kanvas-validator` is turned on with `challenger` mode.

**NOTE:** Here `L2 EE client` means `kanvas-geth` and `L2 CL client` means `kanvas-node`. `L2 EE client` can
be expanded to other clients for pragmatic decentralization.

### Sequencer -> Proposer

We use `sequencing` for transaction inclusion, exclusion and ordering. Currently, `sequencer` does not only
this but also does block building. Thinking of Ethereum PoS, this is what `proposer` exactly does, which is building
and proposing a block. To take this into consideration, we rename the actor `sequencer` to `proposer` and plan to
separate block building ability from `proposer` in the future. We will make the results of our research public.

### Proposer -> Validator

For the same reason above, we use the term `validator` to represent an entity that submits
[L2 output root][g-l2-output-root]. Because this resembles how L1 validator casts FFG vote at every epoch.

## Geth

To prepare for migration to ZK Rollup, we use a [ZK Trie][g-zktrie] to represent state. Currently, this makes
the chain slower than [Merkle Patrica Trie][g-mpt]. As the bottleneck is the time to produce ZK proof right now,
we adopt it from [Scroll]. When we overcome the proof generation time problem, we will smoothly migrate state
without a hard fork or huge changes. Thus, you might face an unexpected result when retrieving JSON-RPC such as
`eth_getProof`.

Additionally, to produce a zkEVM proof, geth should return so called `execution trace` via JSON-RPC
`kanvas_getBlockTraceByNumberOrHash` which provides zkEVM prover with data as a witness.

[scroll]: https://scroll.io/

## Validator

### ZK fault proof

Instead of [cannon], Kanvas uses zkEVM for [ZK fault proof][g-zk-fault-proof].

[cannon]: https://github.com/ethereum-optimism/cannon
