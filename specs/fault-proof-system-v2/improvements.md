# Improvements

<!-- All glossary references in this file. -->

[g-checkpoint-output]: ../glossary.md#checkpoint-output
[g-validator]: ../glossary.md#validator
[g-l2-output]: ../glossary.md#l2-output-root

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Overview](#overview)
- [Guaranteed Win of Valid Output](#guaranteed-win-of-valid-output)
- [Guaranteed Data Verifiability for ZK Proof Verification](#guaranteed-data-verifiability-for-zk-proof-verification)
- [Robust to Delay Attacks](#robust-to-delay-attacks)
- [Accurate Identification of First Disagreeing Block](#accurate-identification-of-first-disagreeing-block)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

Fault Proof System V2 has been designed to address the key limitations identified in
[Fault Proof System V1](../challenge.md), as outlined in the [Background](./README.md#background). This document details
how each of the four main limitations of V1 has been effectively resolved in V2.

## Guaranteed Win of Valid Output

In V2, the [output branching model](./output-proposal.md#output-branching-model) has been introduced. This model allows
[validators][g-validator] to follow and propose [outputs][g-checkpoint-output] along the branch they believe to be
correct, thereby eliminating the need to delete outputs deemed invalid as a result of a challenge. This directly
addresses a critical V1 issue where a challenger could always win if the previous output was deleted, as it was
impossible to verify the last point of agreement between the asserter and challenger on-chain. Now, the previous output
remains intact, and validators can branch out from the first output they dispute.

Also, in V2, challenge creation and dissection are based on [historical commitments](./definitions.md#claim) of outputs,
unlike V1, which was based on individual outputs. This change allows any validators who concur with the entire
[history](./definitions.md#history) to participate in an ongoing challenge, ensuring that they don't lose due to
timeouts.

## Guaranteed Data Verifiability for ZK Proof Verification

[Transaction Data Commitment](./transaction-data-commitment.md) in V2 ensures that the transaction data required for ZK
proof verification is verifiable on-chain. Validators can now conclusively prove on-chain that they have executed the
correct transactions in the target block to compute the state transition and claimed the correct post-state.

## Robust to Delay Attacks

In V1, delay attacks were possible through successive malicious challenges and deletions of outputs in the case where a
previous output was deleted. V2's output branching model solves this effectively.

## Accurate Identification of First Disagreeing Block

Because [transaction data is verifiable on-chain](./transaction-data-commitment.md), the
[version 1 output root payload](./output-proposal.md#output-root-payload-version-1) no longer relies on the information
from the next block, which leads to the elimination of `next_block_hash` from [output root][g-l2-output] payload. This
allows accurate identification of the first disagreeing block after dissection.
