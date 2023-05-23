# Challenge

<!-- All glossary references in this file. -->

[g-l1]: glossary.md#layer-1-l1

[g-l2]: glossary.md#layer-2-l2

[g-l2-output]: glossary.md#l2-output-root

[g-trusted-validator]: glossary.md#trusted-validator

[g-validator]: glossary.md#validator

[g-zk-fault-proof]: glossary.md#zk-fault-proof

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Overview](#overview)
- [Colosseum Contract](#colosseum-contract)
- [Bisection](#bisection)
- [State Diagram](#state-diagram)
- [Process](#process)
- [Public Input Verification](#public-input-verification)
- [Upgradeability](#upgradeability)
- [Summary of Definitions](#summary-of-definitions)
  - [Constants](#constants)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

When a [challenger][g-validator] detects that a submitted [L2 output root][g-l2-output] contains an invalid state
transition, it starts a challenge process by triggering the [Colosseum contract](#colosseum-contract). This involves the
asserter and the challenger by force and continues until either one wins. Only one challenge can be on an output,
and if the challenger wins, the existing output is replaced with the new output claimed by the challenger.
All challenges must be approved by the Security Council. If the content of the challenge is untrue,
the challenge may not be approved, even if the proof is verified by the contract.

## Colosseum Contract

The Colosseum contract implements the following interface:

```solidity
interface Colosseum {
  event ChallengeCreated(
    uint256 indexed outputIndex,
    address indexed asserter,
    address indexed challenger,
    uint256 timestamp
  );

  event Bisected(uint256 indexed outputIndex, uint256 turn, uint256 timestamp);
  event Proven(uint256 indexed outputIndex, bytes32 newOutputRoot);
  event Approved(uint256 indexed outputIndex, uint256 timestamp);
  event Deleted(uint256 indexed outputIndex, uint256 timestamp);

  function createChallenge(
    uint256 _outputIndex,
    bytes32[] calldata _segments
  ) external;

  function bisect(uint256 _outputIndex, uint256 _pos, bytes32[] calldata _segments) external;

  function proveFault(
    uint256 _outputIndex,
    bytes32 _outputRoot,
    uint256 _pos,
    Types.OutputRootProof calldata _srcOutputRootProof,
    Types.OutputRootProof calldata _dstOutputRootProof,
    Types.PublicInput calldata _publicInput,
    Types.BlockHeaderRLP calldata _rlps,
    uint256[] calldata _proof,
    uint256[] calldata _pair
  ) external;

  function challengerTimeout(uint256 _outputIndex) external;

  function approveChallenge(uint256 _outputIndex) external;
}

```

## Bisection

At this moment, it takes a long time to generate a proof for state transition of even a single block.
To resolve this problem, we adopt an `interactive fault proof`.
Unlike other implementations bisecting until both parties find a single step of instruction to execute, we bisect
until the challenger finds a single step of block to prove fault-ness.
The basic idea is derived from [Arbitrum Nitro](https://github.com/OffchainLabs/nitro).

**NOTE:** Someday if the proof generation becomes fast, this process will be removed.

In conclusion, we have two requirements. One is bisecting to find a single block, the other is the last turn should be
challenger's turn. By making use of the fact that [L2 output root][g-l2-output] is always submitted at every 1800
[L2][g-l2] blocks, for example, we can decompose 1800 into 45 and 40. We let the challenger submit 46(45 + 1) segments
and the asserter do 41(40 + 1). The reason why the challenger submits more than the asserter is to prevent challenge
abusing. As a result, after one interaction between parties, [ZK fault proof][g-zk-fault-proof] can be ready to be
verified on [L1][g-l1]. In reality, 1800 blocks are segmented using `SEGMENTS_LENGTHS`.

Here we give a simple example with a small number to show you how it works. Let's say there are 11 blocks and the 3rd
block is the block a challenger want to argue against and this number is decomposed into 5 and 2. Also, let's assume
that both of them agree the state transitions to the 2nd block.

| Turn       | Action          | Segment Start | Segment Length | Segments        | Condition          |
|------------|-----------------|---------------|----------------|-----------------|--------------------|
| Challenger | createChallenge | 0             | 6              | [0, 2, ..., 10] | No                 |
| Asserter   | bisect(2)       | 2             | 3              | [2, 3', 4']     | 2 = 2 && 4 != 4'   |
| Challenger | proveFault(2)   | 2             | 2              | [2, 3'']        | 2 = 2 && 3' != 3'' |

You can notice that in each turn, the first element of the segments must be same with the element at the same index of
the previous segments. Whereas, the last element of the segments must be different from the element at the same index of
the previous segments. In this way, both parties are able to agree with a single step of block.

## State Diagram

![state-diagram](assets/colosseum-state-diagram.svg)

1. If the challenge is created, at the same time, the challenger needs to submit the first segments(9 outputs).
   The state is set to `ASSERTER_TURN`.
2. Then the asserter picks the first invalid segment and submits the next segments(6 outputs) for the picked segment.
   `ASSERTER_TURN` state goes to `CHALLENGER_TURN`.
3. If there's more segments to be interacted with, the challenger picks the first invalid segment and submits the next
   segments(10 outputs) for the picked segment. `CHALLENGER_TURN` state goes to `ASSERTER_TURN` and repeat from step 2.
4. Otherwise, `CHALLENGER_TURN` state goes to `READY_TO_PROVE` automatically. At this state, the challenger is now able
   to pick the first invalid output and submit ZK fault proof.
5. Both `ASSERTER_TURN` and `CHALLENGER_TURN` states have a timeout called `BISECTION_TIMEOUT` and if it happens, the
   state goes to `ASSERTER_TIMEOUT` and `CHALLENGER_TIMEOUT` respectively. This is to mitigate _liveness attack_.
   This is because we want to give a penalty to one who doesn't respond timely.
6. If the submitted proof is turned out to be invalid, the state stays at `READY_TO_PROVE` until `PROVING_TIMEOUT` is
   occurred.
7. Otherwise, `READY_TO_PROVE` state goes to `PROVEN`.
8. At `PROVEN` state, the challenge must be approved by the **Security Council** to mitigate _ZK soundness attack_.
   Which means there are more than one proof that prove different state transitions. This will be removed once we ensure
   the possibility of soundness is extremely low in the production environment.
9. As `PROVEN` state goes to `APPROVED`, The L2 output root is replaced by the one claimed by the challenger,
   and the challenger takes all the bonds for that output.
10. The `ASSERTER_TIMEOUT` state is similar to `READY_TO_PROVE`, it requires the proof to be submitted and verified as
    in step 6 to complete the challenge.
11. At `ASSERTER_TIMEOUT` state, if the challenger doesn't prove the fault within the timeout called `PROVING_TIMEOUT`,
    the state goes to `CHALLENGER_TIMEOUT`.
12. At `PROVEN` state, the **Security Council** verifies the authenticity of the challenge and approves it.
    If the challenge is incorrect, it will not be approved and the challenge will fail.

**Note:** `CHALLENGER_TIMEOUT` state is treated specially. It is regarded as `CHALLENGE_FAIL` state because there's no
motivation for the asserter to step further.

## Process

We want the validator role to be decentralized. Like how the PoS mechanism works, to achieve this,
the validator needs to bond more than `MIN_BOND_AMOUNT` at every output submission. A Validator can deposit at once
for convenience. The qualified validator now obtains the right to submit output.

If outputs are submitted at the same time, only the first output is accepted. If no one submits during
`SUBMISSION_TIMEOUT`, [trusted validator][g-trusted-validator] will submit an output.

Even though the output is challenged, validators still are able to submit an output if the asserted output is thought
to be valid. If the asserted output turns out to be invalid, it is replaced, but the bond for that remains untouched.
This is because it's impossible to determine whether submitted outputs are invalid without a challenge game.

We'll show an example. Let's say `MIN_BOND_AMOUNT` is 100.

1. At time `t`, alice, bob, and carol are registered as validators, and they submitted outputs like following:

| Name  | Output | Challenge | Bond | Lock                     |
|-------|--------|-----------|------|--------------------------|
| alice | O_1800 | N         | 100  | L_{t + 7 days}           |
| bob   | O_3600 | N         | 100  | L_{t + 7 days + 1 hours} |
| bob   | O_5400 | N         | 100  | L_{t + 7 days + 2 hours} |
| carol | O_7200 | N         | 100  | L_{t + 7 days + 3 hours} |

**NOTE:** `O_number` denotes the output at specific block `number`. `L_t` denotes "the bond should be locked
until time `t`".

2. At `t + 3 hours 30 minutes`, david initiates a challenge to the output at 5400.

| Name  | Output | Challenge    | Bond | Lock                     |
|-------|--------|--------------|------|--------------------------|
| alice | O_1800 | N            | 100  | L_{t + 7 days}           |
| bob   | O_3600 | N            | 100  | L_{t + 7 days + 1 hours} |
| bob   | O_5400 | Y (by david) | 200  | L_{t + 7 days + 2 hours} |
| carol | O_7200 | N            | 100  | L_{t + 7 days + 3 hours} |

3. At `t + 4 hours`, emma submits a output at 9000.

| Name  | Output | Challenge    | Bond | Lock                     |
|-------|--------|--------------|------|--------------------------|
| alice | O_1800 | N            | 100  | L_{t + 7 days}           |
| bob   | O_3600 | N            | 100  | L_{t + 7 days + 1 hours} |
| bob   | O_5400 | Y (by david) | 200  | L_{t + 7 days + 2 hours} |
| carol | O_7200 | N            | 100  | L_{t + 7 days + 3 hours} |
| emma  | O_9000 | N            | 100  | L_{t + 7 days + 4 hours} |

4. If the challenger wins:

| Name  | Output | Challenge | Bond | Lock                     |
|-------|--------|-----------|------|--------------------------|
| alice | O_1800 | N         | 100  | L_{t + 7 days}           |
| bob   | O_3600 | N         | 100  | L_{t + 7 days + 1 hours} |
| david | O_5400 | N         | 200  | L_{t + 7 days + 2 hours} |
| carol | O_7200 | N         | 100  | L_{t + 7 days + 3 hours} |
| emma  | O_9000 | N         | 100  | L_{t + 7 days + 4 hours} |

5. Otherwise:

| Name  | Output | Challenge | Bond | Lock                     |
|-------|--------|-----------|------|--------------------------|
| alice | O_1800 | N         | 100  | L_{t + 7 days}           |
| bob   | O_3600 | N         | 100  | L_{t + 7 days + 1 hours} |
| bob   | O_5400 | N         | 200  | L_{t + 7 days + 2 hours} |
| carol | O_7200 | N         | 100  | L_{t + 7 days + 3 hours} |
| emma  | O_9000 | N         | 100  | L_{t + 7 days + 4 hours} |

## Public Input Verification

The following verification process applies to output version 1 and later:

The `_pair[4]` contains the [public input](./zkevm-prover.md#zkevm-proof), which must be processed before
verification by [ZK Verifier Contract](./zkevm-prover.md#the-zk-verifier-contract) can be performed.

1. Check whether `_srcOutputRootProof` is the preimage of the first output root of the segment.
2. Check whether `_dstOutputRootProof` is the preimage of the next output root of the segment.
3. Verify that the `nextBlockHash` in `_srcOutputRootProof` matches the `blockHash` in `_dstOutputRootProof`.
4. Verify that the `nextBlockHash` in `_srcOutputRootProof` matches the block hash derived from `_publicInput` and `_rlps`.
5. Verify that the `transactionsRoot` in `_publicInput` matches the transaction root derived from the transaction
   hashes.
6. If the transaction hash in `_publicInput` is less than `MAX_TXS`, fill it with `DUMMY_HASH`.

## Upgradeability

Colosseum should be behind upgradable proxies.

## Summary of Definitions

### Constants

| Name                  | Value                                                              | Unit              |
|-----------------------|--------------------------------------------------------------------|-------------------|
| `MIN_BOND_AMOUNT`     | TBD                                                                | wei               |
| `SUBMISSION_TIMEOUT`  | TBD                                                                | seconds           |
| `BISECTION_TIMEOUT`   | TBD                                                                | seconds           |
| `PROVING_TIMEOUT`     | TBD                                                                | seconds           |
| `SEGMENTS_LENGTHS`    | [9, 6, 10, 6]                                                      | array of integers |
| `MAX_TXS`             | 25                                                                 | uint256           |
| `DUMMY_HASH`(sepolia) | 0xe3c0fb45c84ee6608b3ee3a7016c505f46ff23736038a4344abf62156e2b21be | bytes32           |
| `DUMMY_HASH`(easel)   | 0xec42d5de5b086e5922e6b0b65ff579305bca5681eed40133209c86cfbc2c7d48 | bytes32           |
