# Validations

<!-- All glossary references in this file. -->

[g-l1]: glosarry.md#l1
[g-l2]: glosarry.md#l2
[g-zk-fault-proof]: glossary.md#zk-fault-proof
[g-zktrie]: glossary.md#zk-trie

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Submitting L2 Output Commitments](#submitting-l2-output-commitments)
- [L2 Output Commitment Construction](#l2-output-commitment-construction)
  - [Output Payload(Version 0)](#output-payloadversion-0)
  - [Output Payload(Version 1)](#output-payloadversion-1)
- [The L2 Output Oracle Contract](#the-l2-output-oracle-contract)
  - [Configuration](#configuration)
- [Security Considerations](#security-considerations)
  - [L1 Reorgs](#l1-reorgs)
- [Summary of Definitions](#summary-of-definitions)
  - [Constants](#constants)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

![Validation Overview](assets/verifier-proving-fault-proof.svg)

After processing one or more blocks, the outputs will need to be synchronized with [L1][g-l1] for trustless execution of
L2-to-L1 messaging, such as withdrawals. Outputs are hashed in a tree-structured form which minimizes the cost of
proving any piece of data captured by the outputs.
Validators submit the output roots to L1 and can be contested with a [ZK fault proof][g-zk-fault-proof],
with a bond at stake if the proof is wrong.

## Submitting L2 Output Commitments

The validator's role is to construct and submit output roots, which are commitments made on a configurable interval,
to the `L2OutputOracle` contract running on L1. It does this by running the [validator](../components/validator/),
a service which periodically queries the rollup node's
[`kroma_outputAtBlock` rpc method](./rollup-node.md#l2-output-rpc-method) for the latest output root derived
from the latest [finalized](rollup-node.md#finalization-guarantees) L1 block. The construction of this output root is
described [below](#l2-output-commitment-construction).

If there is no newly finalized output, the service continues querying until it receives one. It then submits this
output, and the appropriate timestamp, to the [L2 Output Oracle](#the-l2-output-oracle-contract) contract's
`submitL2Output()` function. The block number must correspond to the `startingBlockNumber` plus the next
multiple of the `SUBMISSION_INTERVAL` value.

The validator may also delete multiple output roots by calling the `deleteL2Outputs()` function and specifying the
index of the first output to delete, this will also delete all subsequent outputs.

## L2 Output Commitment Construction

The `output_root` is a 32 byte string, which is derived based on the a versioned scheme:

```pseudocode
output_root = keccak256(version_byte || payload)
```

where:

1. `version_byte` (`bytes32`) a simple version string which increments anytime the construction of the output root
   is changed.

2. `payload` (`bytes`) is a byte string of arbitrary length.

### Output Payload(Version 0)

The version 0 payload is defined as:

```pseudocode
payload = state_root || withdrawal_storage_root || block_hash
```

where:

1. The `block_hash` (`bytes32`) is the block hash for the [L2][g-l2] block that the output is generated from.

2. The `state_root` (`bytes32`) is the [ZK-Trie][g-zktrie] root of all execution-layer accounts.
   This value is frequently used and thus elevated closer to the L2 output root, which removes the need to prove its
   inclusion in the pre-image of the `block_hash`. This reduces the merkle proof depth and cost of accessing the
   L2 state root on L1.

3. The `withdrawal_storage_root` (`bytes32`) elevates the ZK Trie root of the
  [L2ToL1MessagePasser contract](./withdrawals.md#the-l2tol1messagepasser-contract) storage. Instead of making a
  [ZKT][g-zktrie] proof for a withdrawal against the state root (proving first the storage root of the
  L2toL1MessagePasser against the state root, then the withdrawal against that storage root), we can prove against the
  L2toL1MessagePasser's storage root directly, thus reducing the verification cost of withdrawals on L1.

### Output Payload(Version 1)

The version 1 payload is defined as:

```pseudocode
payload = state_root || withdrawal_storage_root || block_hash || next_block_hash
```

where:

1. The `next_block_hash` (`bytes32`) is the next block hash for the block that is next to the `block_hash`.

## The L2 Output Oracle Contract

L2 blocks are produced at a constant rate of `L2_BLOCK_TIME` (2 seconds).
A new L2 output MUST be appended to the chain once per `SUBMISSION_INTERVAL` which is based on a number of blocks.

The L2 Output Oracle contract implements the following interface:

```solidity
interface L2OutputOracle {
    function deleteL2Outputs(uint256 _l2OutputIndex) external;

    function submitL2Output(
        bytes32 _outputRoot,
        uint256 _l2BlockNumber,
        bytes32 _l1Blockhash,
        uint256 _l1BlockNumber
    );

    function nextBlockNumber() public view returns (uint256);
}
```

### Configuration

The `startingBlockNumber` must be at least the number of the first recorded L2 block.
The `startingTimestamp` MUST be the same as the timestamp of the first recorded L2 block.

The first `outputRoot` submitted will thus be at height `startingBlockNumber + SUBMISSION_INTERVAL`

## Security Considerations

### L1 Reorgs

If the L1 has a reorg after an output has been generated and submitted, the L2 state and correct output may change
leading to a misbehavior. This is mitigated against by allowing the validator to submit an
L1 block number and hash to the [L2 Output Oracle](#the-l2-output-oracle-contract) when appending a new output;
in the event of a reorg, the block hash will not match that of the block with that number and the call will revert.

## Summary of Definitions

### Constants

| Name                  | Value  | Unit    |
|-----------------------|--------|---------|
| `SUBMISSION_INTERVAL` | `1800` | blocks  |
| `L2_BLOCK_TIME`       | `2`    | seconds |
