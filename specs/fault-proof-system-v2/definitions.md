# Definitions

<!-- All glossary references in this file. -->

[g-l2-output]: ../glossary.md#l2-output-root
[g-validator]: ../glossary.md#validator
[g-zk-fault-proof]: ../glossary.md#zk-fault-proof
[g-security-council]: ../glossary.md#security-council

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Overview](#overview)
- [Participants](#participants)
  - [Asserter](#asserter)
  - [Challenger](#challenger)
- [History](#history)
- [Claim](#claim)
  - [History Root](#history-root)
- [Edge](#edge)
  - [Edge Dissection](#edge-dissection)
  - [Edge Confirmation](#edge-confirmation)
  - [Rival Edge](#rival-edge)
  - [Presumptive Edge](#presumptive-edge)
  - [Presumptive Timer](#presumptive-timer)
  - [Edge Confirmation Condition](#edge-confirmation-condition)
- [Constants](#constants)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

This document provides definitions of new terms introduced in Fault Proof System V2.

## Participants

### Asserter

An asserter refers to a validator who proposes the first output for a particular index or a validator proceeding a
challenge in the same party as the one proposing the first output.

### Challenger

A challenger refers to a validator who creates a challenge by proposing a different output when an output already exists
on a particular index, or a validator proceeding the challenge in the same party as the one creating the challenge.

## History

The history $H$ represents the sequence of [output roots][g-l2-output] $O$. The first output root in the sequence $O_0$
is the agreed-upon initial state. The height of the history $n$ can be at most `PROPOSAL_INTERVAL`.

```math
H_n = (O_0, O_1, ..., O_n) \ \ \ \ \ (n \leq \verb#PROPOSAL_INTERVAL#)
```

## Claim

A claim $C$ is a commitment to the entire sequence of the history, represented as a Merkle root of a Merkle tree with
all outputs in the history as leaf nodes. When a challenger creates a challenge and submits a claim of the history that
the challenger believes to be correct, any [validators][g-validator] who agree with the history can participate in the
challenge and take over the remaining challenge process on behalf of the challenger who initiated the challenge.
Likewise, after the asserter submits a claim to respond to a challenge, anyone who agrees with the asserter's claim can
continue the challenge on behalf of the asserter. This mechanism allows to turn the one-vs-one challenge of V1 into an
all-vs-all challenge.

### History Root

History root is essentially the same as claim, it denotes the Merkle root of historical outputs.

## Edge

An edge is consisted of two claims $(C_n, C_m) \ (n < m)$. The length of an edge is $m - n$, which can be at least 1 and
at most `PROPOSAL_INTERVAL`. The challenger must submit $(C_0, C_\verb#PROPOSAL_INTERVAL#)$ when creating a challenge,
where $C_\verb#PROPOSAL_INTERVAL#$ is the claim corresponding to the challenge target output and $C_0$ is the claim
corresponding to the previous output.

### Edge Dissection

Edges are dissected to find the first disagreement point between asserters and challengers. Each dissection deepens the
edge and necessitates the submission of intermediate claims, and these claims can also be represented as edges in pairs
of two. In other words, a challenge can be considered as a contest between edges. The depth of an edge indicates the
number of times that the edge has been dissected from the first edge submitted when creating the challenge. The number
of claims that should be submitted for each depth as a result of dissection is determined by a configuration.

### Edge Confirmation

Edge confirmation determines the challenge outcome. A valid edge can be confirmed, and the associated output is
considered to be valid.

### Rival Edge

A rival edge represents one of the conflicting edges with the same length and starting claim but different ending
claims. Only a pair of rival edges exists at each depth because the first claim must be the same and the last claim must
be different. The rival edge is dissected in the next dissection, narrowing down to an edge with a length of 1. As the
rival edges represent different histories, only one of them can be confirmed.

### Presumptive Edge

An edge without a rival. It can be confirmed after a set period if no rival edge submitted.

### Presumptive Timer

An edge's presumptive timer tracks the duration an edge and its parent edges have been in presumptive status. Edge
confirmation is possible once the presumptive timer passes a predefined deadline.

### Edge Confirmation Condition

Edge confirmation can be accomplished in three ways.

- If all child edges are confirmed, the parent edge can be confirmed.
- A presumptive edge can be confirmed after its presumptive timer expires.
- A rival edge of length 1 can be confirmed with [ZK fault proof][g-zk-fault-proof] verification.

If one of the first two conditions is met, only one of the rival edges can be confirmed on a first-come, first-served
basis. If the third condition is met, redundant confirmations are allowed in order to detect ZK soundness errors
on-chain. In this case, the [Security Council][g-security-council] must intervene to handle the redundant confirmations.

## Constants

| Name                | Value  | Unit   |
|---------------------|--------|--------|
| `PROPOSAL_INTERVAL` | `1800` | blocks |
