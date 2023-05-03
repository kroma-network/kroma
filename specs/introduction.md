# Introduction

<!-- All glossary references in this file. -->

[g-checkpoint-output]: glossary.md#checkpoint-output
[g-zk-fault-proof]: glossary.md#zk-fault-proof

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Foundations](#foundations)
  - [Blockchain Trilemma](#blockchain-trilemma)
  - [What is Ethereum scalability?](#what-is-ethereum-scalability)
  - [What is a Layer2?](#what-is-a-layer2)
  - [What is a Rollup?](#what-is-a-rollup)
  - [What is EVM Equivalence?](#what-is-evm-equivalence)
  - [🎶 All together now 🎶](#-all-together-now-)
- [Protocol Guarantees](#protocol-guarantees)
- [Network Participants](#network-participants)
  - [Users](#users)
  - [The Proposer](#the-proposer)
  - [Validators](#validators)
- [Key Interaction Diagrams](#key-interaction-diagrams)
  - [Depositing and Sending Transactions](#depositing-and-sending-transactions)
  - [Withdrawing](#withdrawing)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

Kroma is an _EVM equivalent_, _optimistic rollup_ protocol designed to _scale Ethereum_ while remaining maximally
compatible with existing Ethereum infrastructure. This document provides an overview of the protocol to provide context
for the rest of the specification.

## Foundations

### Blockchain Trilemma

The core value of the blockchain is _decentralization_. It means anyone can take participation in the network without
any permission. Then an issue arises that because participants of the network can't be trusted. Thus, there needs
so-called consensus among the participants. To do consensus, more than 2/3 participants' agreements are required.
In other words, to re-org the chain, you need this much stake. In reality, it is very difficult to perform this attack.
That's why blockchain is _secure_. What matters is an individual participant needs to re-execute transaction(a.k.a state
transition) to do consensus, which requires state synchronization. This is an obstacle for blockchain to be _scalable_.
You can’t simply reduce the block time or increase the block size to achieve scalability. Because the block time gets
reduced or the block size gets increased abruptly, it raises the huddle. As a result, only some of centralized parties
will remain as participants and this will hinder _decentralization_ that is the reason blockchain exists for.

### What is Ethereum scalability?

Scaling Ethereum means increasing the number of useful transactions the Ethereum network can process. Ethereum's
limited resources, specifically bandwidth, computation, and storage, constrain the number of transactions which can be
processed on the network. Of the three resources, computation and storage are currently the most significant
bottlenecks. These bottlenecks limit the supply of transactions, leading to extremely high fees. Scaling ethereum and
reducing fees can be achieved by better utilizing bandwidth, computation and storage.

### What is a Layer2?

Layer 2 is a solution for Ethereum scalability. The key is to do state transition at not onchain but offchain.
This means the transaction execution is not part of the consensus. Why? because the more participants re-execute
transactions, the more expensive and the slower the computation will be. So small part of participants will do
state transition and they submit the result of state transition to Ethereum(or called Layer 1). Therefore,
Layer 2 inherits the security of Layer 1. To re-org L2 chain, you need to perform attack on Layer 1.

### What is a Rollup?

[Rollup](https://vitalik.ca/general/2021/01/05/rollup.html) is one kind of Layer 2 solutions. What differs from others
is to submit transactions to Layer 1. This is called data is available. So anyone can derive state and know how much
balance they own. Depending on how to prove state transition, the solution is divided into
Optimistic Rollup(ORU) and ZK Rollup(ZKR).

Optimistic rollup is a rollup that needs a so-called fault proof to prove a valid state transition. As the name
suggests, it's based on the assumption that most of the state transitions are valid. Something is needed when
the state transition is wrong. This is where a fault proof kicks in. A fault proof is designed to be able to verify
that submitted state transition is invalid. This can reduce a tremendous fee as a result. But the only weakness of
optimistic rollup is the _dispute period_ that is enough to be long to accept any challenges. That's why this needs
7 days to withdraw.

Opposite to Optimistic Rollup, ZK rollup is a rollup that needs a so-called validity proof. This proof is mathematically
very difficult to deceive verifier. Therefore, as soon as the proof is verified, the state transition can be immediately
finalized.

### What is EVM Equivalence?

[EVM Equivalence](https://medium.com/ethereum-optimism/introducing-evm-equivalence-5c2021deb306) is complete compliance
with the state transition function described in the Ethereum yellow paper, the formal definition of the protocol. By
conforming to the Ethereum standard across EVM equivalent rollups, smart contract developers can write once and deploy
anywhere.

**NOTE:** At this moment, `OP_SELFDESTRUCT` is disabled. We are actively trying to fully cover the evm spec. But like
[eip-4758] suggests, we don't recommend users to rely on this opcode.

[eip-4758]: https://eips.ethereum.org/EIPS/eip-4758

### 🎶 All together now 🎶

**Kroma is an _EVM equivalent_, _optimistic rollup_ protocol designed to _scale Ethereum_.**

## Protocol Guarantees

In order to scale Ethereum without sacrificing security, we must preserve 3 critical properties of Ethereum layer 1:
liveness, availability, and validity.

1. **Liveness** - Anyone must be able to extend the rollup chain by sending transactions at any time.
   - There are two ways transactions can be sent to the rollup chain: 1) via the proposer, and 2) directly on layer 1.
     The proposer provides low latency & low cost transactions, while sending transactions directly to layer 1 provides
     censorship resistance.
2. **Availability** - Anyone must be able to download the rollup chain.
   - All information required to derive the chain is embedded into layer 1 blocks. That way as long as the layer 1 chain
     is available, so is the rollup.
3. **Validity** - All transactions must be correctly executed and all withdrawals correctly processed.
   - The rollup state and withdrawals are managed on an L1 contract called the `L2OutputOracle`. This oracle is
     guaranteed to _only_ finalize correct (ie. valid) rollup block hashes given a **single honest validator**
     assumption. If there is ever an invalid block hash asserted on layer 1, an honest validator will prove it is
     invalid and win a bond.

## Network Participants

There are three actors in Kroma: users, proposers, and validators.

![Network Overview](./assets/network-participants-overview.svg)

### Users

At the heart of the network are users (us!). Users can:

1. Deposit or withdraw arbitrary transactions on L2 by sending data to a contract on Ethereum mainnet.
2. Use EVM smart contracts on layer 2 by sending transactions to the proposers.
3. View the status of transactions using block explorers provided by network validators.

### The Proposer

The proposer is the primary block producer.
(At this moment, there is a single proposer that is operated by Lightscale.
We have been actively discussing about how to open the role and how to separate block building role
from proposer's role.)
There may be one proposer **or** many using a consensus protocol.
In general, specifications may use "the proposer" to be a stand-in term
for the consensus protocol operated by multiple proposers.

The proposer:

1. Accepts user off-chain transactions.
2. Observes on-chain transactions (primarily, deposit events coming from L1).
3. Consolidates both kinds of transactions into L2 blocks with a specific ordering.
4. Propagates consolidated L2 blocks to L1, by submitting two things as calldata to L1:
   - The pending off-chain transactions accepted in step 1.
   - Sufficient information about the ordering of the on-chain transactions to successfully reconstruct the blocks.
from step 3., purely by watching L1.

The proposer also provides access to block data as early as step 3., so that users may access real-time state in
advance of L1 confirmation if they so choose.

### Validators

Validators serve two purposes:

1. Verifying rollup integrity by asserting [checkpoint output][g-checkpoint-output].
2. Disputing invalid assertions by submitting [ZK fault proof][g-zk-fault-proof].

In order for the network to remain secure there must be **at least** one honest validator who is able to verify the
integrity of the rollup chain & serve blockchain data to users.

## Key Interaction Diagrams

The following diagrams demonstrate how protocol components are utilized during key user interactions in order to
provide context when diving into any particular component specification.

### Depositing and Sending Transactions

Users will often begin their L2 journey by depositing ETH from L1.
Once they have ETH to pay fees, they'll start sending transactions on L2.
The following diagram demonstrates this interaction and all key Kroma components which are or should be utilized:

![Diagram of Depositing and Sending Transactions](./assets/proposer-handling-deposits-and-transactions.svg)

Links to components mentioned in this diagram:

- [Rollup Node](./rollup-node.md)
- [Execution Engine](./exec-engine.md)
- [Batcher](./batcher.md)
- [L2 Output Oracle](./validations.md#the-l2-output-oracle-contract)
- [Validator](./validations.md#submitting-l2-output-commitments)

### Withdrawing

Just as important as depositing, it is critical that users can withdraw from the rollup. Withdrawals are initiated by
normal transactions on L2, but then completed using a transaction on L1 after the dispute period has elapsed.

![Diagram of Withdrawing](./assets/user-withdrawing-to-l1.svg)

Links to components mentioned in this diagram:

- [Kroma Portal](./deposits.md#deposit-contract)
- [L2 Output Oracle](./validations.md#l2-output-oracle-smart-contract)
