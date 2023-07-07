# Validations

<!-- All glossary references in this file. -->

[g-l1]: glosarry.md#l1
[g-l2]: glosarry.md#l2
[g-zk-fault-proof]: glossary.md#zk-fault-proof
[g-zktrie]: glossary.md#zk-trie
[g-l2-output]: glossary.md#l2-output-root
[g-validator]: glossary.md#validator
[g-priority-round]: glossary.md#priority-round
[g-public-round]: glossary.md#public-round

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Submitting L2 Output Commitments](#submitting-l2-output-commitments)
- [L2 Output Commitment Construction](#l2-output-commitment-construction)
  - [Output Payload(Version 0)](#output-payloadversion-0)
- [L2 Output Oracle Smart Contract](#l2-output-oracle-smart-contract)
  - [Configuration of L2OutputOracle](#configuration-of-l2outputoracle)
- [Validator Pool Smart Contract](#validator-pool-smart-contract)
  - [Validation Rewards](#validation-rewards)
  - [Configuration of ValidatorPool](#configuration-of-validatorpool)
- [Security Considerations](#security-considerations)
  - [L1 Reorgs](#l1-reorgs)
- [Summary of Definitions](#summary-of-definitions)
  - [Constants](#constants)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

![Validation Overview](assets/verifier-proving-fault-proof.svg)

After processing one or more blocks, the outputs will need to be synchronized with [L1][g-l1] for trustless execution of
L2-to-L1 messaging, such as withdrawals.
These output proposals act as the bridge's view into the L2 state.
Actors called "Validators" submit the output roots to L1 and can be contested with a [ZK fault proof][g-zk-fault-proof],
with a bond at stake if the proof is wrong.

## Submitting L2 Output Commitments

The validator's role is to construct and submit output roots, which are commitments to the L2's state,
to the `L2OutputOracle` contract on L1. To do this, the validator periodically
queries the [rollup node](./rollup-node.md) for the latest output root derived from the latest
[finalized][finality] L1 block. It then takes the output root and
submits it to the `L2OutputOracle` contract on L1.

[finality]: https://hackmd.io/@prysmaticlabs/finality

The validator that submits the output root is determined by the `ValidatorPool` contract on L1.
The output submission rounds are divided into [Priority Round][g-priority-round] and [Public Round][g-public-round],
and the time limit of each round is configured as `ROUND_DURATION` in the ValidatorPool contract.
A prioritized validator is selected by a random function in the `ValidatorPool` contract, and the prioritized validator
must submit output within the `Priority Round` time.
If the prioritized validator fails to submit within the `Priority Round`, the round moves to the `Public Round`, where
all validators can submit output regardless of priority.

The [validator](../components/validator/) is expected to submit output roots on a deterministic
interval based on the configured `SUBMISSION_INTERVAL` in the `L2OutputOracle`. The larger
the `SUBMISSION_INTERVAL`, the less often L1 transactions need to be sent to the `L2OutputOracle`
contract, but L2 users will need to wait a bit longer for an output root to be included in L1
that includes their intention to withdrawal from the system.

The honest `kroma-validator` algorithm assumes a connection to the `L2OutputOracle` contract to know
the L2 block number that corresponds to the next output root that must be submitted. It also
assumes a connection to an `kroma-node` to be able to query the `kroma_syncStatus` RPC endpoint.

Once your submitted output is [finalized][finality], the submitter becomes eligible for a reward.
For more information on this, see [Validation Rewards](#validation-rewards).

```python
import time

while True:
    next_checkpoint_block = L2OutputOracle.nextBlockNumber()
    rollup_status = kroma_node_client.sync_status()
    if rollup_status.finalized_l2.number >= next_checkpoint_block:
        output = kroma_node_client.output_at_block(next_checkpoint_block)
        tx = send_transaction(output)
    time.sleep(poll_interval)
```

## L2 Output Commitment Construction

The `output_root` is a 32bytes string, which is derived based on a versioned scheme:

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
payload = state_root || withdrawal_storage_root || block_hash || next_block_hash
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

4. The `next_block_hash` (`bytes32`) is the next block hash for the block that is next to the `block_hash`.

The height of the block where the output is submitted has been delayed by one.

## L2 Output Oracle Smart Contract

L2 blocks are produced at a constant rate of `L2_BLOCK_TIME`.
A new L2 output MUST be appended to the chain once per `SUBMISSION_INTERVAL` which is based on the number of L2 blocks.

L2 Output Oracle Smart Contract implements the following interface:

```solidity
interface L2OutputOracle {
  event OutputSubmitted(
    bytes32 indexed outputRoot,
    uint256 indexed l2OutputIndex,
    uint256 indexed l2BlockNumber,
    uint256 l1Timestamp
  );

  event OutputReplaced(uint256 indexed outputIndex, bytes32 newOutputRoot);

  function replaceL2Output(
    uint256 _l2OutputIndex,
    bytes32 _newOutputRoot,
    address _submitter
  ) external;

  function submitL2Output(
    bytes32 _outputRoot,
    uint256 _l2BlockNumber,
    bytes32 _l1BlockHash,
    uint256 _l1BlockNumber,
    uint256 _bondAmount
  ) external payable;
}
```

### Configuration of L2OutputOracle

The `startingBlockNumber` must be at least the number of the first recorded L2 block.
The `startingTimestamp` MUST be the same as the timestamp of the first recorded L2 block.

Thus, the first `outputRoot` submitted will be at height `startingBlockNumber`, and each subsequent one will be at
height incremented by `SUBMISSION_INTERVAL`.

## Validator Pool Smart Contract

Only accounts registered as [Validator][g-validator] can submit [output][g-l2-output] to
the [L2 Output Oracle](#l2-output-oracle-smart-contract).
To register as a [Validator][g-validator], you must deposit at least `MIN_BOND_AMOUNT` of ETH into
the `ValidatorPool` contract.
When submitting the output, the validator must bond at least `MIN_BOND_AMOUNT` of Ethereum, which will be unbonded and
rewarded to the L2 `ValidatorRewardVault` contract when the output is finalized.

Validator Pool Smart Contract implements the following interface:

```solidity
interface ValidatorPool {
  event Bonded(
    address indexed submitter,
    uint256 indexed outputIndex,
    uint128 amount,
    uint128 expiresAt
  );

  event BondIncreased(address indexed challenger, uint256 indexed outputIndex, uint128 amount);
  event Unbonded(uint256 indexed outputIndex, address indexed recipient, uint128 amount);

  function deposit() external payable;

  function withdraw(uint256 _amount) external;

  function createBond(
    uint256 _outputIndex,
    uint128 _amount,
    uint128 _expiresAt
  ) external;

  function increaseBond(address _challenger, uint256 _outputIndex) external;

  function unbond() external;

  function balanceOf(address _addr) external view returns (uint256);

  function nextValidator() public view returns (address);
}
```

### Validation Rewards

A validator who submits an output can receive a reward from L2 `ValidatorRewardVault` contract when the output is
finalized, it is called validation reward. When the output is finalized, the `ValidatorPool` contract sends a message
to pay the reward in the L2 `ValidatorRewardVault` via the `KromaPortal` contract.

The `ROUND_DURATION` time is divided into `NON_PENALTY_PERIOD` and `PENALTY_PERIOD`.
The `NON_PENALTY_PERIOD` is the time to guarantee that transaction will be included in the L1 block.
If the validator submits the output within this time, it can receive the full reward.
The full reward is the balance in the `ValidatorRewardVault` when the output is finalized, divided by `REWARD_DIVIDER`,
which is a value equal to the number of outputs in a week. If the validator submits an output after the
`NON_PENALTY_PERIOD`, the reward will gradually decrease, it is called a penalty. The percentage of the penalty is
calculated using the time elapsed during the `PENALTY_PERIOD`, and the reward is reduced by this percentage.

Rewards received by the validator can be withdrawn to L1 via `withdraw()` in the ValidatorRewardVault.

### Configuration of ValidatorPool

The `NON_PENALTY_PERIOD` is the period during a submission round that is not penalized (in seconds).
The `PENALTY_PERIOD` is the period during a submission round that is penalized (in seconds).

The sum of the two values, `NON_PENALTY_PERIOD` and `PENALTY_PERIOD`, must be equal to `ROUND_DURATION`.
`ROUND_DURATION` is equal to `(L2_BLOCK_TIME * SUBMISSION_INTERVAL) / 2`

## Security Considerations

### L1 Reorgs

If the L1 has a reorg after an output has been generated and submitted, the L2 state and correct output may change
leading to a misbehavior. This is mitigated against by allowing the validator to submit an
L1 block number and hash to the [L2 Output Oracle](#l2-output-oracle-smart-contract) when appending a new output;
in the event of a reorg, the block hash will not match that of the block with that number and the call will revert.

## Summary of Definitions

### Constants

| Name                  | Value  | Unit           |
|-----------------------|--------|----------------|
| `SUBMISSION_INTERVAL` | `1800` | blocks         |
| `L2_BLOCK_TIME`       | `2`    | seconds        |
| `MIN_BOND_AMOUNT`     | TBD    | wei            |
| `REWARD_DIVIDER`      | `168`  | num of outputs |
| `ROUND_DURATION`      | `30`   | minutes        |
| `NON_PENALTY_PERIOD`  | `10`   | minutes        |
| `PENALTY_PERIOD`      | `20`   | minutes        |
