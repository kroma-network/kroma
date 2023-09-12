# Challenge

<!-- All glossary references in this file. -->

[g-l1]: glossary.md#layer-1-l1
[g-l2]: glossary.md#layer-2-l2
[g-l2-output]: glossary.md#l2-output-root
[g-trusted-validator]: glossary.md#trusted-validator
[g-validator]: glossary.md#validator
[g-zk-fault-proof]: glossary.md#zk-fault-proof
[g-security-council]: glossary.md#security-council

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Overview](#overview)
- [State Diagram](#state-diagram)
- [Challenge Creation](#challenge-creation)
- [Bisection](#bisection)
- [Proving Fault](#proving-fault)
- [Dismiss Challenge](#dismiss-challenge)
- [Force Delete Output](#force-delete-output)
- [Contract Interface](#contract-interface)
- [Upgradeability](#upgradeability)
- [Summary of Definitions](#summary-of-definitions)
  - [Constants](#constants)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

When a [validator][g-validator] detects that a submitted [L2 output root][g-l2-output] contains an invalid state
transition, it can start a dispute challenge process by triggering the [Colosseum contract](#colosseum-contract). We
refer to a validator who submits a dispute challenge as a "challenger" and a validator who initially submitted
an L2 output as an "asserter." A dispute challenge entails a confrontational interaction between an asserter and a
challenger, which persists until one of them emerges victorious. If the challenger wins, the corresponding L2 output 
will be deleted.

A single output can be subject to multiple challenges. Challengers also need to stake their bonds equivalent to those 
staked by asserter when submitting L2 outputs to generate challenges. Should the asserter emerge victorious in a 
challenge, they receive the staked bonds of all the challengers as reward. On the other hand, if a challenger prevails, 
the one who submitted the first valid ZK fault proof is given the asserter's staked bond. As a preventive measure 
against collusion between asserters and challengers, tax is imposed. If there are any ongoing challenges, 
the challenges are canceled, and staked bonds are refunded to the respective challengers.

In the ZK fault-proof challenge process, the following undeniable bug might arise, prompting the intervention of the 
[Security Council][g-security-council]:

- The deletion of a valid output due to two valid and contradictory zk proofs
- The failure to delete an invalid output due to the bugs in prover/verifier or ZK completeness error
- The deletion of a valid output due to two valid and contradictory ZK proofs
- The failure to delete an invalid output due to the bugs in prover/verifier or ZK completeness error

In the former case, the Security Council validates the legitimacy of the deleted output and, if the aforementioned 
error is identified, dismisses the challenge and initiates a rollback of the deleted output. 
In the latter scenario, all challengers will fail in proving the fault. In such cases, the Security Council verifies 
the output and, if deemed invalid, delete the output forcibly. All interventions by the Security Council are executed 
through multi-sig transactions.

## State Diagram

![state-diagram](assets/colosseum-state-diagram.svg)

1. If the challenge is created, at the same time, the challenger needs to submit the first segments(9 outputs).
   The state is set to `ASSERTER_TURN`.
2. Then the asserter picks the first invalid segment and submits the next segments(6 outputs) for the picked segment.
   `ASSERTER_TURN` state goes to `CHALLENGER_TURN`.
3. If there's more segments to be interacted with, the challenger picks the first invalid segment and submits the next
   segments(10 outputs) for the picked segment. `CHALLENGER_TURN` state goes to `ASSERTER_TURN` and repeat from step 2.
   If the output has already been deleted by other challenger, the challenger cancel challenge and refund bond.
4. Both `ASSERTER_TURN` and `CHALLENGER_TURN` states have a timeout called `BISECTION_TIMEOUT` and if it happens, the
   state goes to `ASSERTER_TIMEOUT` and `CHALLENGER_TIMEOUT` respectively. This is to mitigate _liveness attack_.
   This is because we want to give a penalty to one who doesn't respond timely.
5. When the asserter timed out or bisection is completed, the state of challenge will be `READY_TO_PROVE` automatically.
   At this state, the challenger is now able to pick the first invalid output and submit ZK fault proof.
   Likewise, the challenge is canceled if the output is already been deleted.
6. If the submitted proof is turned out to be invalid, the state stays at `READY_TO_PROVE` until `PROVING_TIMEOUT` is
   occurred.
7. Otherwise, `READY_TO_PROVE` state goes to `PROVEN`, and the L2 output is deleted.
8. The deleted output would be validated by the **[Security Council][g-security-council]** to mitigate ZK soundness 
   attack.
9. If the deleted output was invalid, so it should have been, the Security Council do nothing.
10. Otherwise, the **Security council** will dismiss the challenge and rollback the valid output.

## Challenge Creation

Validators can initiate challenges when they suspect that an invalid output has been submitted. In their role as 
challengers, they start the challenge process with initial segments for interactive fault proof.

> **Note** Challenges can only be initiated within the `CREATION_PERIOD` (< `FINALIZATION_PERIOD`) since the output 
> is submitted. This restriction aims to prevent malicious challengers from deleting outputs just before finalization, 
> causing a delay attack. 

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

When the challenge process is completed and the corresponding output is deleted by other challenger during bisection,
the challenge will be canceled automatically.

## Proving Fault

Since Colosseum verifies public input along with [zkEVM-proof](./zkevm-prover.md#zkevm-proof), challengers should
calculate as below and enclose the public input to the `proveFault` transaction.

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
  const maxTxs = 100;

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

The following is the verification process of invalid output by
[ZK Verifier Contract](./zkevm-prover.md#the-zk-verifier-contract):

1. Check whether the challenge is ready to prove. The status of challenge should be `READY_TO_PROVE`
   or `ASSERTER_TIMEOUT`.
2. Check whether `srcOutputRootProof` is the preimage of the first output root of the segment.
3. Check whether `dstOutputRootProof` is the preimage of the next output root of the segment.
4. Verify that the `nextBlockHash` in `srcOutputRootProof` matches the `blockHash` in `dstOutputRootProof`.
5. Verify that the `stateRoot` in `publicInput` matches the `stateRoot` in `dstOutputRootProof`.
6. Verify that the `nextBlockHash` in `srcOutputRootProof` matches the block hash derived from `publicInput` and `rlps`.
7. Verify that the `withdrawalStorageRoot` in `dstOutputRootProof` is contained in `stateRoot` in `dstOutputRootProof`
   using `merkleProof`.
8. If the length of transaction hashes in `publicInput` is less than `MAX_TXS`, fill it with `DUMMY_HASH`.
9. Verify the `_zkproof` using `_pair` and `publicInputHash`. The `publicInputHash` is derived from the `publicInput`
   and `stateRoot` of `srcOutputRootProof`, while `_zkproof` and `_pair` are submitted by the challenger directly.
10. Delete the output and request validation of the challenge to [Security Council][g-security-council] if there is any 
    undeniable bugs such as soundness error.
11. If the deleted output was valid so the challenge has an undeniable bug, Security Council will 
    [dismiss](#dismiss-challenge) the challenge and roll back the output.

## Dismiss Challenge

Upon a successful challenge resulting in output deletion, the Security Council will verify the genuineness of the
deleted output(two valid contradicting ZK proofs). Given that the deletion of output introduces withdrawal delays, 
the Security Council conducts a thorough investigation into this issue. Upon validation of the legitimate nature of the 
output deletion, the Security Council will dismiss the challenge and initiate the process of output rollback.
This can only be executed through the multi-sig transaction of the Security Council.

## Force Delete Output

In the event that an undeniable bug within the ZK fault-proof system, such as a ZK completeness error, is detected, it
becomes necessary to remove outputs deemed invalid. To address this, the Security Council is tasked with inspecting 
outputs that have completed the bisect process but have failed the fault-proof verification. If an invalid output is 
submitted and is determined to be associated with an undeniable bug, the Security Council holds the authority to delete
the output through a multi-sig transaction.

## Contract Interface

The Colosseum contract implements the following interface:

```solidity
interface Colosseum {
    /**
       * @notice Emitted when the challenge is created.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param asserter    Address of the asserter.
     * @param challenger  Address of the challenger.
     * @param timestamp   The timestamp when created.
     */
    event ChallengeCreated(
        uint256 indexed outputIndex,
        address indexed asserter,
        address indexed challenger,
        uint256 timestamp
    );

    /**
     * @notice Emitted when segments are bisected.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param challenger  Address of the challenger.
     * @param turn        The current turn.
     * @param timestamp   The timestamp when bisected.
     */
    event Bisected(
        uint256 indexed outputIndex,
        address indexed challenger,
        uint8 turn,
        uint256 timestamp
    );

    /**
     * @notice Emitted when it is ready to be proved.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param challenger  Address of the challenger.
     */
    event ReadyToProve(uint256 indexed outputIndex, address indexed challenger);

    /**
     * @notice Emitted when proven fault.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param challenger  Address of the challenger.
     * @param timestamp   The timestamp when proven.
     */
    event Proven(uint256 indexed outputIndex, address indexed challenger, uint256 timestamp);

    /**
     * @notice Emitted when challenge is dismissed.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param challenger  Address of the challenger.
     * @param timestamp   The timestamp when dismissed.
     */
    event ChallengeDismissed(
        uint256 indexed outputIndex,
        address indexed challenger,
        uint256 timestamp
    );

    /**
     * @notice Emitted when challenge is canceled.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param challenger  Address of the challenger.
     * @param timestamp   The timestamp when canceled.
     */
    event ChallengeCanceled(
        uint256 indexed outputIndex,
        address indexed challenger,
        uint256 timestamp
    );

    /**
     * @notice Emitted when challenger timed out.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param challenger  Address of the challenger.
     * @param timestamp   The timestamp when deleted.
     */
    event ChallengerTimedOut(
        uint256 indexed outputIndex,
        address indexed challenger,
        uint256 timestamp
    );

    /**
       * @notice Creates a challenge against an invalid output.
     *
     * @param _outputIndex   Index of the invalid L2 checkpoint output.
     * @param _l1BlockHash   The block hash of L1 at the time the output L2 block was created.
     * @param _l1BlockNumber The block number of L1 with the specified L1 block hash.
     * @param _segments      Array of the segment. A segment is the first output root of a specific range.
     */
    function createChallenge(
        uint256 _outputIndex,
        bytes32 _l1BlockHash,
        uint256 _l1BlockNumber,
        bytes32[] calldata _segments
    ) external;

    /**
       * @notice Selects an invalid section and submit segments of that section.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     * @param _challenger  Address of the challenger.
     * @param _pos         Position of the last valid segment.
     * @param _segments    Array of the segment. A segment is the first output root of a specific range.
     */
    function bisect(
        uint256 _outputIndex,
        address _challenger,
        uint256 _pos,
        bytes32[] calldata _segments
    ) external;

    /**
       * @notice Proves that a specific output is invalid using ZKP.
     *         This function can only be called in the READY_TO_PROVE and ASSERTER_TIMEOUT states.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     * @param _pos         Position of the last valid segment.
     * @param _proof       Proof for public input validation.
     * @param _zkproof     Halo2 proofs composed of points and scalars.
     *                     See https://zcash.github.io/halo2/design/implementation/proofs.html.
     * @param _pair        Aggregated multi-opening proofs and public inputs. (Currently only 2 public inputs)
     */
    function proveFault(
        uint256 _outputIndex,
        uint256 _pos,
        Types.PublicInputProof calldata _proof,
        uint256[] calldata _zkproof,
        uint256[] calldata _pair
    ) external;

    /**
      * @notice Calls a private function that deletes the challenge because the challenger has timed out.
     *         Reverts if the challenger hasn't timed out.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     * @param _challenger  Address of the challenger.
     */
    function challengerTimeout(uint256 _outputIndex, address _challenger) external;

    /**
      * @notice Cancels the challenge.
     *         Reverts if is not possible to cancel the sender's challenge for the given output index.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     */
    function cancelChallenge(uint256 _outputIndex) external;

    /**
      * @notice Dismisses the challenge and rollback l2 output.
     *         This function can only be called by Security Council contract.
     *
     * @param _outputIndex      Index of the L2 checkpoint output.
     * @param _challenger       Address of the challenger.
     * @param _asserter         Address of the asserter.
     * @param _outputRoot       The L2 output root to rollback.
     * @param _publicInputHash  Hash of public input.
     */
    function dismissChallenge(
        uint256 _outputIndex,
        address _challenger,
        address _asserter,
        bytes32 _outputRoot,
        bytes32 _publicInputHash
    ) external;

    /**
      * @notice Deletes the L2 output root forcefully by the Security Council
     *         when zk-proving is not possible due to an undeniable bug.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     */
    function forceDeleteOutput(uint256 _outputIndex) external;
}

```

## Upgradeability

Colosseum contract should be deployed behind upgradable proxies.

## Summary of Definitions

### Constants

| Name                          | Value                                                              | Unit              |
|-------------------------------|--------------------------------------------------------------------|-------------------|
| `REQUIRED_BOND_AMOUNT`        | 200000000000000000 (0.2 ETH)                                       | wei               |
| `FINALIZATION_PERIOD_SECONDS` | 604800                                                             | seconds           |
| `CREATION_PERIOD_SECONDS`     | 518400                                                             | seconds           |           
| `BISECTION_TIMEOUT`           | 3600                                                               | seconds           |
| `PROVING_TIMEOUT`             | 28800                                                              | seconds           |
| `SEGMENTS_LENGTHS`            | [9, 6, 10, 6]                                                      | array of integers |
| `MAX_TXS`                     | 100                                                                | uint256           |
| `DUMMY_HASH`                  | 0xedf1ae3da135c124658e215a9bf53477facb442a1dcd5a92388332cb6193237f | bytes32           |
