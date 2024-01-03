# Challenge

<!-- All glossary references in this file. -->

[g-checkpoint-output]: ../glossary.md#checkpoint-output
[g-state]: ../glossary.md#state
[g-zk-fault-proof]: ../glossary.md#zk-fault-proof
[g-l2-output]: ../glossary.md#l2-output-root
[g-sequencer-batch]: ../glossary.md#sequencer-batch
[g-user-deposited]: ../glossary.md#user-deposited-transaction
[g-batcher-transaction]: ../glossary.md#batcher-transaction
[g-channel]: ../glossary.md#channel
[g-channel-frame]: ../glossary.md#channel-frame
[g-sequencer]: ../glossary.md#sequencer
[g-security-council]: ../glossary.md#security-council

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Overview](#overview)
- [Challenge Process](#challenge-process)
- [`Colosseum` Interface](#colosseum-interface)
  - [Edge Struct](#edge-struct)
  - [Depth Zero Edge Creation](#depth-zero-edge-creation)
  - [Edge Dissection](#edge-dissection)
  - [Edge Confirmation](#edge-confirmation)
    - [Confirm By Children](#confirm-by-children)
    - [Confirm By Time](#confirm-by-time)
    - [Confirm By ZK Proof](#confirm-by-zk-proof)
- [ZK Proving Scheme](#zk-proving-scheme)
- [Constants](#constants)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

Similar to [V1](../challenge.md), the challenge process in Fault Proof System V2 revolves around interactively
dissecting [outputs][g-checkpoint-output] to find the first disagreeing block, thereby proving the validity of the
[state][g-state] transition.

Key differences in V2, as opposed to V1, include the followings:

- **Commitment Submission**: Asserters and challengers submit a [history](./definitions.md#history) commitment, termed a
  [_claim_](./definitions.md#claim), during challenge creation or dissection.
- **Participation Expansion**: Any participant agreeing with the commitment can join the dissection of challenges
  initiated by others.
- **Challenge Conclusion**: After dissection, a length 1 [edge](./definitions.md#edge) is
  [confirmed](./definitions.md#edge-confirmation) by submitting a [ZK fault proof][g-zk-fault-proof] for state
  transition verification. Upon the challenge period expiration, [presumptive edges](./definitions.md#presumptive-edge)
  are also confirmed. Parent edges are finally confirmed when all child edges are confirmed, ultimately concluding the
  challenge. The output linked to the confirmed depth zero edge is deemed valid.

## Challenge Process

The challenge process proceeds as follows:

1. Output Proposal
    - A malicious [asserter](./definitions.md#asserter) proposes an output.
2. Challenge Creation
    - An honest [challenger](./definitions.md#challenger) proposes a new, valid output and creates a challenge with a
      corresponding depth zero edge.
3. Rival Depth Zero Edge Submission
    - The asserter submits a depth zero edge as a rival.
4. Dissection by Challenger
    - The challenger dissects the depth zero edge, submitting dissected edges at depth 1.
5. Dissection by Asserter
    - The asserter also submits depth 1 edges, including a [rival edge](./definitions.md#rival-edge).
6. Further Dissections
    - Steps 4 and 5 are repeated, increasing the depth by one each time, until the edge length becomes 1.
7. ZK Fault Proof Submission
    - The challenger submits a ZK fault proof for the deepest rival edge. (The malicious asserter can also submit ZK
      proof, but it will never be verified)
8. Deepest Rival Edge Confirmation
    - The edge is confirmed upon successful ZK proof verification. (`confirmEdgeByZkProof`)
9. Presumptive Edges Confirmation
    - If the [presumptive timer](./definitions.md#presumptive-timer) exceeds, each presumptive edge is confirmed.
      (`confirmEdgeByTime`)
10. Parent Edges Confirmation
    - From the deepest depth to zero, if all child edges are confirmed, the parent edge can is also confirmed.
      (`confirmEdgeByChildren`)
11. Challenge Conclusion
    - The depth zero edge that the honest challenger submitted is finally confirmed.

The dissection process, depicted in below diagram, starts from a depth zero edge with a length of 1800 (assuming
`PROPOSAL_INTERVAL` = `1800`). In this example, dissection creates `[8, 5, 9, 5]` edges for each depth. The number of
depths and edges for each depth are configurable settings. In every depth, there is only one pair of rival edges, while
all others are presumptive edges.

As mentioned in [Edge Confirmation Condition](./definitions.md#edge-confirmation-condition), at the deepest depth, depth
4, the presumptive edges submitted by the malicious asserter can all be confirmed in time. Otherwise, the rival edge can
never be confirmed because the asserter will not be able to verify the ZK proof for correct execution. Therefore, the
rival edge at the higher depth cannot be confirmed because one of its children has not been confirmed, also the edge at
depth zero submitted by the malicious asserter can never be confirmed.

![edge-dissection](../assets/edge-dissection.svg)

While [BoLD](https://github.com/OffchainLabs/bold) bisects edges, Kroma dissects them. The choice for dissection lies in
its ability to divide the range into more sections simultaneously, enabling asserters and challengers to identify the
first point of disagreement with fewer interactions. Although a single dissection transaction may incur higher gas
usage, the entire challenge process can be proceeded at a significantly lower expense due to the substantially reduced
number of required transactions. Additionally, in contrast to BoLD, which does not have a fixed L2 block interval for
assertion submissions, Kroma operates with a fixed interval known as `PROPOSAL_INTERVAL` for proposing outputs.
Consequently, the number of edges to be dissected at each depth remains constant, facilitating dissection instead of
bisection.

## `Colosseum` Interface

### Edge Struct

Edges, being central to the challenge process, require an efficient storage and management scheme within contracts. For
efficient management, we define the mutual ID shared by rival edges and the unique ID of each edge as follows:

```solidity
bytes32 mutualId = keccak256(
    abi.encodePacked(
        edge.prevOutputRoot,
        edge.startHeight,
        edge.startHistoryRoot,
        edge.endHeight
    )
);

bytes32 edgeId = keccak256(abi.encodePacked(mutualId, edge.endHistoryRoot));
```

Below is the definition of an edge struct.

```solidity
/**
 * @notice EdgeStatus represents the current status of an edge.
 *
 * @custom:value PENDING   Yet to be confirmed. Not all edges can be confirmed.
 * @custom:value CONFIRMED Once confirmed it cannot transition back to pending.
 */
enum EdgeStatus {
    PENDING,
    CONFIRMED
}

/**
 * @notice ChallengeEdge represents an edge committing to a range of states. These edges will be dissected,
 *         slowly reducing them in length until they reach length 1.
 *
 * @custom:field prevOutputRoot   The previous output root is a link from the edge to an output at a lower
 *                                level. Intuitively all edges with the same previous output root agree on
 *                                the state committed to the previous output root. The purpose of the previous
 *                                output root is to ensure that only edges that agree on a common start
 *                                position are being compared against one another.
 * @custom:field outputRoot       The output root that this edge claims to be true.
 * @custom:field startHistoryRoot A root of all the states in the history up to the start height.
 * @custom:field startHeight      The height of the start history root.
 * @custom:field endHistoryRoot   A root of all the states in the history up to the end height. Since
 *                                endHeight > startHeight, the start history root must commit to a prefix of
 *                                the states committed to by the end history root.
 * @custom:field endHeight        The height of the end history root.
 * @custom:field childIds         Edges can be dissected into multiple children. If this edge has been
 *                                dissected, the ids of the children are populated here. Until that time this
 *                                array length is 0.
 * @custom:field createdAtBlock   The L1 block number when this edge was created.
 * @custom:field confirmedAtBlock The L1 block number at which this edge was confirmed. Zero if not confirmed.
 * @custom:field status           Current status of this edge. All edges are created PENDING, and may be updated
 *                                to CONFIRMED. Once CONFIRMED they cannot transition back to PENDING.
 * @custom:field depth            The depth of this edge. Next dissection history roots length will be
 *                                determined by the depth.
 */
struct ChallengeEdge {
    bytes32 prevOutputRoot;
    bytes32 outputRoot;
    bytes32 startHistoryRoot;
    uint256 startHeight;
    bytes32 endHistoryRoot;
    uint256 endHeight;
    bytes32[] childIds;
    uint64 createdAtBlock;
    uint64 confirmedAtBlock;
    EdgeStatus status;
    uint8 depth;
}
```

### Depth Zero Edge Creation

The challenger must submit an edge of depth zero to create a challenge. The depth zero edge used to create a challenge
consists of:

- the root of a Merkle tree of size 1 with the previous output of the challenge target output as a leaf node
- the root of a Merkle tree with all outputs between the previous output and the target output as leaf nodes

Therefore, to verify that the edge submitted by the challenger is correctly computed with the corresponding outputs, the
process includes verifying it through inclusion proof and prefix proof. The prefix proof is used to construct the post
tree based on the Merkle tree composed of the previous output, ensuring that the post tree has the same root as the
Merkle tree composed of all outputs up to the target output.

```solidity
/**
 * @notice CreateEdgeArgs represents the data for creating a depth zero edge.
 *
 * @custom:field prevOutputRoot The previous output root of the challenged output root.
 * @custom:field outputRoot     The output root that is being claimed correct by the newly created edge.
 * @custom:field endHistoryRoot The end history root of the edge to be created.
 * @custom:field endHeight      The end height of the edge to be created. End height is
 *                              (PROPOSAL_INTERVAL * challenged output index).
 * @custom:field prefixProof    Proof that the start history root commits to a prefix of the states that
 *                              end history root commits to.
 * @custom:field inclusionProof Proof to show that the end state is the last state in the end history root.
 * @custom:field l1BlockHash    The L1 block hash which must be included in the current chain.
 * @custom:field l1BlockNumber  The L1 block number with the specified L1 block hash.
 */
struct CreateEdgeArgs {
    bytes32 prevOutputRoot;
    bytes32 outputRoot;
    bytes32 endHistoryRoot;
    uint256 endHeight;
    bytes prefixProof;
    bytes32[] inclusionProof;
    bytes32 l1BlockHash;
    uint256 l1BlockNumber;
}

/**
 * @notice Performs necessary checks and creates a new depth zero edge.
 *
 * @param args Edge creation args.
 */
function createDepthZeroEdge(CreateEdgeArgs calldata args) external onlyValidator {
    // 1. Check if args.outputRoot is correctly linked to args.prevOutputRoot
    // 2. Check if the status of args.outputRoot is PENDING
    // 3. Check if there is another output connected to args.prevOutputRoot (i.e., if there are no other
    //    rival outputs, no need to create a challenge)
    // 4. Calculate the root of the Merkle tree with args.outputRoot as the (PROPOSAL_INTERVAL)-th leaf
    //    node using the args.inclusionProof, and verify if this value is identical to args.endHistoryRoot
    // 5. Calculate the root of the Merkle tree with size 1 using args.prevOutputRoot, and verify if this
    //    value as startHistoryRoot is the prefix of args.endHistoryRoot using args.prefixProof
    // 6. Ensure that the edge does not already exist using edge id
    // 7. Check if the block hash of args.l1BlockNumber is the same as args.l1BlockHash (sanity check for
    //    reorg)
    // 8. Store the edge
}
```

### Edge Dissection

The edge dissection process submits claims to construct multiple child edges as a result of the dissection. The claims
are Merkle roots corresponding to the intermediate points that evenly divide the entire range represented by the parent
edge. Therefore, the process involves verifying these claims with a prefix proof to ensure that they are computed
correctly.

```solidity
/**
 * @notice Mapping of edge depth to the number of claims to be submitted.
 */
mapping(uint256 => uint256) internal dissectionHistoryRootsNum;

/**
 * @notice Dissects an edge. This creates multiple child edges:
 *         The lowest child has the same start root and height as this edge, but a different end root and
 *         height. The highest child has the same end root and height as this edge, but a different start
 *         root and height.
 *
 * @param edgeId                 Edge id to dissect.
 * @param dissectionHistoryRoots The new history roots to be the children of this edge.
 * @param prefixProof            A proof to show that the dissection history roots commit to prefixes of
 *                               the current end history root.
 */
function dissectEdge(
    bytes32 edgeId,
    bytes32[] memory dissectionHistoryRoots,
    bytes calldata prefixProof
) external onlyValidator {
    // 1. Check if the status of the edge with edgeId is PENDING
    // 2. Check if there is a rival edge for the edge (No need to dissect for presumptive edges)
    // 3. Check if the dissectionHistoryRootsNum corresponding to edge.depth matches the length of
    //    dissectionHistoryRoots
    // 4. Validate dissectionHistoryRoots using prefixProof if they are the prefixes of the current
    //    edge.endHistoryRoot. prefixProof consists of preExpansion representing leaf nodes of the Merkle
    //    tree represented by the last dissectionHistoryRoot and a proof verifying that the last
    //    dissectionHistoryRoot is a sub-tree of edge.endHistoryRoot. Using this, verify if each
    //    dissectionHistoryRoot is computed correctly by prefix nodes of preExpansion, and finally,
    //    validate if the post tree constructed using the proof matches the edge.endHistoryRoot
    (bytes32[] memory preExpansion, bytes32[] memory proof) = abi.decode(
        prefixProof,
        (bytes32[], bytes32[])
    );
    // 5. For each element in dissectionHistoryRoots, create a new edge and store it if it does not already
    //    exist
    // 6. Store dissectionHistoryRoots in edge.childIds
    // 7. Update edge.depth to (edge.depth + 1)
}
```

### Edge Confirmation

#### Confirm By Children

If all children edges are confirmed, the parent edge can be confirmed.

```solidity
/**
 * @notice Confirm an edge if all of its children are already confirmed.
 *
 * @param edgeId The id of the edge to confirm.
 */
function confirmEdgeByChildren(bytes32 edgeId) external onlyValidator {
    // 1. Check if the edge with edgeId exists
    // 2. Iterate through edge.childIds, checking if each child exists and has a status of CONFIRMED
    // 3. Ensure that there are no rival edges that have already been confirmed
    // 4. Check if the status of the edge is PENDING and change it to CONFIRMED
}
```

#### Confirm By Time

For a presumptive edge, it can be confirmed if the presumptive timer exceeds the time below.

```math
T = 7 \ \text{days} + ((1 + \verb#number_of_dissections#) * 2 + \verb#number_of_confirmations#) * δ \ \ \ \ \ (δ = \verb#transaction_delay#)
```

The $7 \ \text{days}$ in the first term is to prevent censorship, and the second term is the time it takes to
successfully submit the entire transactions that must be submitted during the challenge process. The added $1$
represents the transaction creating the depth zero edge, and the multiplied $2$ is because the asserters and challengers
each need to submit all the transactions. Thus, a challenge can be confirmed in a maximum of $T$ time. Below,
`challengePeriodBlocks` is the variable that represents $T$.

```solidity
/**
 * @notice The number of blocks to be elapsed to confirm an edge by time.
 */
uint64 public challengePeriodBlocks;

/**
 * @notice An edge can be confirmed if the total amount of time it and a single chain of its direct
 *         ancestors has spent unrivaled is greater than the challenge period.
 *
 * @param edgeId          The id of the edge to confirm.
 * @param ancestorEdgeIds The ids of the direct ancestors of an edge. These are ordered from the parent
 *                        first, then going to grand-parent, great-grandparent etc.
 */
function confirmEdgeByTime(bytes32 edgeId, bytes32[] calldata ancestorEdgeIds) external onlyValidator {
    // 1. Retrieve the depth zero edge corresponding to the last element of ancestorEdgeIds
    // 2. Check if the edge with edgeId exists
    // 3. Add the time elapsed until a rival output is created for the output root connected to the edge
    //    to the timer
    // 4. Add the time elapsed until a rival edge is created for the edge to the timer
    // 5. Iterate through ancestorEdgeIds, calculating the time elapsed until a rival edge is created for
    //    each edge, and add it to the timer. In other words, the total timer represents the cumulative
    //    time during which the connected output root remained unrivaled plus the time the chain of edges,
    //    from the depth zero edge to the edge being confirmed, remained unrivaled after the rival output
    //    root was created
    // 6. Ensure that the total timer exceeds challengePeriodBlocks and there are no rival edges that have
    //    already confirmed
    // 7. Ensure that edge.status is PENDING, and change it to CONFIRMED
}
```

#### Confirm By ZK Proof

A rival edge of length 1 can be confirmed when verified by a ZK fault proof. This process verifies that the edge and the
proof data submitted by the asserter or challenger correspond correctly. It involves retrieving the commitments of
transaction data in the target block on-chain and including it in the public input of the ZK proof to verify that the
asserter or challenger executed the correct transactions, as described in
[Transaction Data Commitment](./transaction-data-commitment.md).

The `OutputRootProof` and `PublicInputProof` in the previous [proving fault](../challenge.md#proving-fault) have been
changed as follows. The `OutputRootProof` was changed accordingly by removing `nextBlockHash` from the
[output root][g-l2-output] payload as described in
[Output Root Payload (Version 1)](./output-proposal.md#output-root-payload-version-1), and some fields for on-chain
verification of [batch data][g-sequencer-batch] and [user-deposited transaction][g-user-deposited] data have been added
to the `PublicInputProof`.

```solidity
/**
 * @notice OutputRootProof represents the elements that are hashed together to generate an output root which
 *         itself represents a snapshot of the L2 state.
 *
 * @custom:field version                  Version of the output root.
 * @custom:field stateRoot                Root of the state trie at the block of this output.
 * @custom:field messagePasserStorageRoot Root of the message passer storage trie.
 * @custom:field blockHash                Hash of the block this output was generated from.
 */
struct OutputRootProof {
    bytes32 version;
    bytes32 stateRoot;
    bytes32 messagePasserStorageRoot;
    bytes32 blockHash;
}

/**
 * @notice PublicInputProof represents the data for verifying the public input.
 *
 * @custom:field srcOutputRootProof          Proof of the source output root.
 * @custom:field dstOutputRootProof          Proof of the destination output root.
 * @custom:field publicInput                 Ingredients to compute the public input used by ZK proof
 *                                           verification.
 * @custom:field rlps                        Pre-encoded RLPs to compute the block hash of the destination
 *                                           output root proof.
 * @custom:field l2ToL1MessagePasserBalance  Balance of the L2ToL1MessagePasser account.
 * @custom:field l2ToL1MessagePasserCodeHash Codehash of the L2ToL1MessagePasser account.
 * @custom:field merkleProof                 Merkle proof of L2ToL1MessagePasser account against the state
 *                                           root.
 * @custom:field dstL1BlockRefNum            The L1 block number corresponding to the destination block.
 * @custom:field versionedHashes             The array of versioned hashes of the L2 batch blobs containing
 *                                           the destination block.
 */
struct PublicInputProof {
    OutputRootProof srcOutputRootProof;
    OutputRootProof dstOutputRootProof;
    PublicInput publicInput;
    BlockHeaderRLP rlps;
    bytes32 l2ToL1MessagePasserBalance;
    bytes32 l2ToL1MessagePasserCodeHash;
    bytes[] merkleProof;
    uint256 dstL1BlockRefNum;
    bytes32[] versionedHashes;
}
```

To verify that the ZK proof is generated using the correct L2 transaction data submitted on-chain, all blobs in the
[batcher transaction][g-batcher-transaction] containing the target block should be checked. This can be verified by the
following steps:

1. reconstructing all the [channels][g-channel] associated with those blobs using all the relevant L2 blocks
2. splitting them into [channel frames][g-channel-frame]
3. creating blobs from the frames in the batcher transaction
4. calculating the `versionedHashes` from the blobs
5. checking if the `versionedHashes` are stored on-chain by the [sequencer][g-sequencer]

This is why `versionedHashes` is added to `PublicInputProof` as an array and the process to ensure that all the elements
of `versionedHashes` are stored on-chain is also added. Because the batch data contains L2 block timestamp information,
we can extract the transaction data of the target block from the blobs even though the blobs contain multiple L2 blocks.

```solidity
/**
 * @notice Confirm an edge by ZK proof verification. Only the edges that have length 1 can be confirmed by
 *         ZK proof, and redundant confirmations between rival edges are allowed.
 *
 * @param edgeId                      The id of the edge to confirm.
 * @param beforeHistoryInclusionProof Proof that the state which is the start of the edge is committed to the
 *                                    start history root.
 * @param afterHistoryInclusionProof  Proof that the state which is the end of the edge is committed to the
 *                                    end history root.
 * @param proof                       Proof for the public input validation.
 * @param zkproof                     Halo2 proofs composed of points and scalars.
 * @param pair                        Aggregated multi-opening proofs and public inputs. (Currently only
 *                                    2 public inputs)
 */
function confirmEdgeByZkProof(
    bytes32 edgeId,
    bytes32[] calldata beforeHistoryInclusionProof,
    bytes32[] calldata afterHistoryInclusionProof,
    PublicInputProof calldata proof,
    uint256[] calldata zkproof,
    uint256[] calldata pair
) external onlyValidator {
    // 1. Check if the length of the edge corresponding to the edgeId is 1
    // 2. Verify if edge.startHistoryRoot has the output root calculated with proof.srcOutputRootProof as a
    //    leaf node at edge.startHeight using beforeHistoryInclusionProof
    // 3. Verify if edge.endHistoryRoot has the output root calculated with proof.dstOutputRootProof as a
    //    leaf node at edge.endHeight using afterHistoryInclusionProof
    // 4. Ensure that proof.publicInput.stateRoot matches proof.dstOutputRootProof.stateRoot
    // 5. Verify if the blockHash calculated with proof.publicInput and proof.rlps matches
    //    proof.dstOutputRootProof.blockHash
    // 6. Verify if the proof.L2ToL1MessagePasser account is included in the proof.dstOutputRootProof.stateRoot
    //    using proof.merkleProof
    // 7. Verify if all the elements of proof.versionedHashes are the values stored on-chain and include them
    //    in the public input
    // 8. Retrieve userDepositedTxAccHash corresponding to proof.dstL1BlockRefNum from on-chain data and
    //    include it in the public input
    bytes32 publicInputHash = hashPublicInput(
        proof.srcOutputRootProof.stateRoot,
        proof.publicInput,
        proof.versionedHashes,
        userDepositedTxAccHash
    );
    // 9. Check if the publicInputHash has not been verified before
    // 10. Verify the ZK proof using zkproof, pair, and publicInputHash
    // 11. Mark that the publicInputHash has been verified
    // 12. Ensure that the edge.status is PENDING, and change it to CONFIRMED
}
```

Note that when confirmed by ZK proofs, redundant confirmations between rival edges are allowed to detect ZK soundness
errors on-chain. However, when confirmed by other conditions, only one of the rival edges can be confirmed and redundant
confirmations are not allowed. In the case of redundant confirmations, the [Security Council][g-security-council] will
step in to guarantee the finalization of the valid output.
([Security Council Intervention](./security-council-intervention.md))

## ZK Proving Scheme

If what the prover is trying to prove is $C(x, w) == true$, then the mathematical representation of this would be like
below.

```math
\begin{array}{l}
C(x, w) =
\\[0.2cm]
ComputeBlockHash(w.Block_{dst}) == x.BlockHash_{dst} \ \ \land
\\[0.2cm]
MerkleVerify(x.StateRoot_{src}, w.State_{src}, w.\pi_{src}) == true \ \ \land
\\[0.2cm]
MerkleVerify(x.StateRoot_{dst}, w.State_{dst}, w.\pi_{dst}) == true \ \ \land
\\[0.2cm]
ComputeAccHash(GetUserDepositTxs(w.Block_{dst})) == x.H_{deposited} \ \ \land
\\[0.2cm]
GetTargetTxs(w.Batches[:], ComputeTimestamp(w.Block_{dst}.n)) == GetL2Txs(w.Block_{dst}) \ \ \land
\\[0.2cm]
\forall j \in \{0, ..., length(TargetBlobs)\}: \verb#VERSIONED_HASH_VERSION_KZG# \ || \ Hash(KZGCommit(TargetBlobs[j]))[1:] == x.H_{versioned}[j] \ \ \land
\\[0.2cm]
STFVerify(x.StateRoot_{src}, x.StateRoot_{dst}, ComputeTxRoot(w.Block_{dst}.transactions); w.State_{src}, w.State_{dst}, w.Block_{dst}.transactions) == true
\\[0.5cm]
\text{where} \ \forall i \in \{0, ..., length(w.Batches)\}: Channel[i] = Compress(rlp(w.Batches[i])),
\\[0.2cm]
\forall i \in \{0, ..., length(w.Batches)\}: Blobs[i] = ConvertChannelToBlobs(Channel[i]),
\\[0.2cm]
TargetBlobs = GetTargetBlobs(Blobs[:], w.TargetBlobIndices)
\end{array}
```

## Constants

| Name                | Value  | Unit   |
|---------------------|--------|--------|
| `PROPOSAL_INTERVAL` | `1800` | blocks |
