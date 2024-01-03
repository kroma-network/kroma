# Fault Proof System V2

<!-- All glossary references in this file. -->

[g-validator]: ../glossary.md#validator
[g-checkpoint-output]: ../glossary.md#checkpoint-output
[g-zk-fault-proof]: ../glossary.md#zk-fault-proof
[g-sequencer-batch]: ../glossary.md#sequencer-batch
[g-l1-attr-deposit]: ../glossary.md#l1-attributes-deposited-transaction
[g-user-deposited]: ../glossary.md#user-deposited-transaction
[g-withdrawals]: ../glossary.md#withdrawals
[g-l2-output]: ../glossary.md#l2-output-root
[g-state-root]: ../glossary.md#state-root
[g-security-council]: ../glossary.md#security-council
[g-state]: ../glossary.md#state

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Background](#background)
- [Overview](#overview)
- [Assumptions](#assumptions)
- [Contents](#contents)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Background

The initial version, [Fault Proof System V1](../challenge.md) introduced a permissionless [validator][g-validator]
system. However, it encountered several limitations:

- **Challenge Invalidity**: There was no assurance that a valid [output][g-checkpoint-output] would always win in
  challenges. Particularly, if a previous output was deleted after a challenge, a challenger could invariably win, even
  against a valid output. Moreover, a valid output could lose if the asserter failed to respond during a challenge.
- **Unverified Proof Inputs**: The [ZK fault proof][g-zk-fault-proof] submitted by the challenger did not necessarily
  ensure that the proof was generated from the [batch data][g-sequencer-batch] posted to L1. It also neglected to
  include [L1 attributes deposited transaction][g-l1-attr-deposit] and [user-deposited transactions][g-user-deposited]
  in the ZK proof verification.
- **Delay Attacks**: The system was prone to delay attacks. Once an output is deleted after a successful challenge, an
  attacker could create continuous challenges against valid outputs following the deleted output, causing them to be
  deleted and leading to indefinite delays in [withdrawals][g-withdrawals].
- **Bisection Inaccuracy**: The addition of the `next_block_hash` field to the [output root][g-l2-output] could
  misdirect the bisection process, incorrectly identifying the first disagreeing block. There is no problem when the
  transaction data is not manipulated but only the [state root][g-state-root] that makes up the output root is
  manipulated. However, if the transaction data itself is manipulated, the block hash will be different and the
  bisection will find an earlier block as the first disagreeing block.

In V1, the [Security Council][g-security-council] will intervene and correct any issues that arise due to the above
limitations. To address these limitations and make complete system, Fault Proof System V2 is designed. This document
provides an overview of Fault Proof System V2 and explains what V2 aims to accomplish by improving upon the limitations
of V1.

## Overview

**Fault Proof System V2** is inspired by Arbitrum's dispute protocol,
[BoLD (Bounded Liquidity Delay)](https://github.com/OffchainLabs/bold). Its design allows parallel challenge engagement
among multiple participants, enabling all-vs-all challenge. In addition, valid output can be finalized in max upper
bound time, assuming at least one honest validator exists. The main differences from BoLD are that the ZK fault proof is
used to prove the validity of [state][g-state] transition on a block-by-block basis, and that dissection is used instead
of bisection. These two differences have the advantage of allowing the challenge to proceed with a lot fewer
interactions.

Fault Proof System V2 also features 100% on-chain verifiability. By storing the commitment of transaction data on-chain,
it is possible to verify that the ZK fault proof was generated as a result of executing the correct transaction.

## Assumptions

Fault Proof System V2 guarantees that valid outputs will always win the challenge within a max upper bound time under
the following assumptions:

- The existence of at least one honest validator.
- L2 batch data submitted to L1 is consistently reliable.
- A dishonest party can censor honest party's transactions for a maximum of 7 days.

*Note*: In the presence of a ZK soundness error, an invalid output may win the challenge. The Security Council will
intervene only in this case.

## Contents

First, we define new terms introduced in Fault Proof System V2. Then, we dive into how the commitments of transaction
data are stored and utilized on-chain to enforce data availability in the ZK proof verification. Following this, the
changes in the output proposal process are detailed, introducing the output branching model. Subsequently, the new
challenge process in Fault Proof System V2 is outlined - challenge creation, the first disagreement point identification
through dissection, and conclusion of challenge through ZK proof verification. Lastly, the documentation examines the
edge case where the Security Council must intervene, and concludes with how V2 addresses the limitations of V1.

- Specifications
  - [Definitions](./definitions.md)
  - [Transaction Data Commitment](./transaction-data-commitment.md)
  - [Output Proposal](./output-proposal.md)
  - [Challenge](./challenge.md)
- [Security Council Intervention](./security-council-intervention.md)
- [Improvements](./improvements.md)
