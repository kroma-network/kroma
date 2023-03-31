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
- [Upgradeability](#upgradeability)
- [Summary of Definitions](#summary-of-definitions)
  - [Constants](#constants)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

When a [challenger][g-validator] detects that submitted [L2 output root][g-l2-output] contains an invalid state
transitions. It starts a challenge process by triggering [Colosseum contract](#colosseum-contract). This involves the
asserter and the challenger by force and continues until one of either wins. At this moment, only a single
challenge process can exist. In other words, another challenger can't initiate a challenge if the existing challenge is
in progress.

## Colosseum Contract

The Colosseum contract implements the following interface:

```solidity
interface Colosseum {
  event ChallengeCreated(
    uint256 indexed challengeId,
    address indexed challenger,
    uint256 indexed outputIndex,
    uint256 timestamp
  );

  event Bisected(uint256 indexed challengeId, uint256 turn, uint256 timestamp);
  event ProofCompleted(uint256 indexed challengeId, uint256 outputIndex);
  event Closed(uint256 indexed challengeId, uint256 turn, uint256 timestamp);

  function createChallenge(
    uint256 _outputIndex,
    bytes32[] calldata _segments
  ) external payable;

  function bisect(uint256 _pos, bytes32[] calldata _segments) external payable;

  function proveFault(
    uint256 _pos,
    bytes32 _outputRoot,
    uint256[] calldata _proof,
    uint256[] calldata _pair
  ) external payable;

  function asserterTimeout() external;

  function challengerTimeout(uint256 _challengeId) external;
}

```

## Bisection

At this moment, it takes a long time to generate proof for state transition of even a single block.
To resolve this problem, we adopt a way so called `interactive fault proof`.
Unlike other implementations bisecting until both parties find a single step of instruction to execute, we bisect
until the challenger finds a single step of block to prove fault-ness.
The basic idea is derived from [Arbitrum Nitro](https://github.com/OffchainLabs/nitro).

**NOTE:** Someday if the proof generation becomes fast, this process can be removed.

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
2. Then the asserter picks the first invalid segment and submit the next segments(6 outputs) for the picked segment.
   `ASSERTER_TURN` state goes to `CHALLENGER_TURN`.
3. If there's more segments to be interacted with, the challenger picks the first invalid segment and submit the next
   segments(10 outputs) for the picked segment. `CHALLENGER_TURN` state goes to `ASSERTER_TURN` and repeat from step 2.
4. Otherwise, `CHALLENGER_TURN` state goes to `READY_TO_PROVE` automatically. At this state, the challenger is now able
   to pick the first invalid output and submit ZK fault proof.
5. Both `ASSERTER_TURN` and `CHALLENGER_TURN` states have a timeout called `INTERACTION_TIMEOUT` and if it happens, the
   state goes to `ASSERTER_TIMEOUT` and `CHALLENGER_TIMEOUT` respectively. This is to mitigate _liveness attack_.
   This is because we want to give a penalty to one who doesn't respond timely.
6. If the submitted proof is turned out to be invalid, the state stays at `READY_TO_PROVE` until `PROOF_TIMEOUT` is
   occurred.
7. Otherwise, `READY_TO_PROVE` state goes to `PROOF_VERIFIED`.
8. At `PROOF_VERIFIED` state, the challenge waits for proof correction in order to mitigate _ZK soundness attack_.
   Which means there are more than one proof that prove different state transitions. This will be removed once we ensure
   the possibility of soundness is extremely low in the production environment.
9. As `PROOF_VERIFIED` state goes to `CHALLENGE_SUCCESS`, the challenger gets rewarded and the winner gets slashed.
10. At `ASSERTER_TIMEOUT` state, the challenger should do an extra action to close the challenge. then the state goes to
    `CHALLENGE_SUCCESS` and the challenger gets rewarded and the winner gets slashed like step 9.
11. At `ASSERTER_TIMEOUT` state, if the challenger doesn't close the challenge timely, the state goes to
    `CHALLENGER_TIMEOUT`.

**Note:** `CHALLENGER_TIMEOUT` state is treated specially. It is regarded as `CHALLENGER_FAIL` state because there's no
motivation for the asserter to step further.

## Process

We want the validator role to be decentralized. Like how PoS mechanism works, to achieve this,
the validator needs to stake more than `MINIMUM_STAKE` at every output submission. A Validator can deposit stakes
at once for convenience. The qualified validator now obtain the right to submit output.

If outputs are submitted at the same time, only the first output is accepted. If no one submits during
`SUBMISSION_TIMEOUT`, [trusted validator][g-trusted-validator] will submit an output.

Even though the output is challenged, validators still are able to submit output if the asserted output is thought to be
valid. If the asserted output turns out to be invalid, all the proceeding outputs are deleted but the stakes on them
remains untouched. This is because it's impossible to determine whether submitted outputs are invalid without challenge
game.

We'll show the example. Let's say `MINIMUM_STAKE` is 100.

1. At time `t`, alice, bob, and carol are registered as validators and they submitted outputs like followings.

  | Name       | Output | Challenge | Stake | Lock                     |
  |------------|--------|-----------|-------|--------------------------|
  | alice      | O_1800 | N         | 100   | L_{t + 7 days}           |
  | bob        | O_3600 | N         | 100   | L_{t + 7 days + 1 hours} |
  | bob        | O_5400 | N         | 100   | L_{t + 7 days + 2 hours} |
  | carol      | O_7200 | N         | 100   | L_{t + 7 days + 3 hours} |

  **NOTE:** `O_number` denotes the output at specific block `number`. `L_t` denotes "the stake should be locked
  until time `t`".

2. At `t + 3 hours 30 minutes`, david initiates a challenge to output at 5400.

  | Name       | Output | Challenge | Stake | Lock                          |
  |------------|--------|-----------|-------|-------------------------------|
  | alice      | O_1800 | N         | 100   | L_{t + 7 days}                |
  | bob        | O_3600 | N         | 100   | L_{t + 7 days + 1 hours}      |
  | bob        | O_5400 | N         | 100   | L_{t + 7 days + 2 hours}      |
  | carol      | O_7200 | N         | 100   | L_{t + 7 days + 3 hours}      |
  | david      | O_5400 | Y         | 100   | L_{until challenge is closed} |

3. At `t + 4 hours`, emma submits a output at 9000.

  | Name       | Output | Challenge | Stake | Lock                          |
  |------------|--------|-----------|-------|-------------------------------|
  | alice      | O_1800 | N         | 100   | L_{t + 7 days}                |
  | bob        | O_3600 | N         | 100   | L_{t + 7 days + 1 hours}      |
  | bob        | O_5400 | N         | 100   | L_{t + 7 days + 2 hours}      |
  | carol      | O_7200 | N         | 100   | L_{t + 7 days + 3 hours}      |
  | david      | O_5400 | Y         | 100   | L_{until challenge is closed} |
  | emma       | O_9000 | N         | 100   | L_{t + 7 days + 4 hours}      |

4. If the challenger wins:

  | Name       | Output | Challenge | Stake | Lock                     |
  |------------|--------|-----------|-------|--------------------------|
  | alice      | O_1800 | N         | 100   | L_{t + 7 days}           |
  | bob        | O_3600 | N         | 100   | L_{t + 7 days + 1 hours} |
  | bob        |        |           | 0     |                          |
  | carol      |        |           | 100   |                          |
  | david      |        |           | 200   |                          |
  | emma       |        |           | 100   |                          |

5. Otherwise:

  | Name       | Output | Challenge | Stake | Lock                          |
  |------------|--------|-----------|-------|-------------------------------|
  | alice      | O_1800 | N         | 100   | L_{t + 7 days}                |
  | bob        | O_3600 | N         | 100   | L_{t + 7 days + 1 hours}      |
  | bob        | O_5400 | N         | 200   | L_{t + 7 days + 2 hours}      |
  | carol      | O_7200 | N         | 100   | L_{t + 7 days + 3 hours}      |
  | emma       | O_9000 | N         | 100   | L_{t + 7 days + 4 hours}      |

## Upgradeability

Colosseum should be behind upgradable proxies.

## Summary of Definitions

### Constants

| Name                       | Value         | Unit              |
|----------------------------|---------------|-------------------|
| `MINIMUM_STAKE`            | TBD           | gwei              |
| `INTERACTION_TIMEOUT`      | TBD           | seconds           |
| `PROOF_TIMEOUT`            | TBD           | seconds           |
| `PROOF_CORRECTION_TIMEOUT` | TBD           | seconds           |
| `SEGMENTS_LENGTHS`         | [9, 6, 10, 6] | array of integers |
