# Security Council Intervention

<!-- All glossary references in this file. -->

[g-checkpoint-output]: ../glossary.md#checkpoint-output
[g-zk-fault-proof]: ../glossary.md#zk-fault-proof
[g-security-council]: ../glossary.md#security-council

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Overview](#overview)
- [Output Finalization By Security Council](#output-finalization-by-security-council)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

In Fault Proof System V2, the objective is to ensure that a valid [outputs][g-checkpoint-output] are always finalized,
while preventing the finalization of any invalid outputs. However, unique scenarios such as ZK soundness errors, where
contradictory [ZK proofs][g-zk-fault-proof] might both appear valid, can lead to the finalization of an invalid output.
These rare instances are detectable on-chain when two [rival edges](./definitions.md#rival-edge) are
[confirmed by ZK proof](./challenge.md#confirm-by-zk-proof). This is where the [Security Council][g-security-council]
steps in.

## Output Finalization By Security Council

The procedure for Security Council intervention is as follows:

- **Scanning Period**: The system allows a maximum period, `challengePeriodBlocks`, to resolve a challenge. Following
  this, there is an additional time window for the Security Council to act, calculated as
  `finalizationPeriodBlocks - challengePeriodBlocks`. During this period, council members scan for redundant
  confirmations by ZK proof among rival edges before any associated outputs are finalized.
- **Detection of Redundant Confirmations**: If redundant confirmations are detected, and the depth zero
  [edge](./definitions.md#edge) connected to the invalid output has confirmed before the edge linked to the valid
  output, the invalid output is at risk of being finalized.
- **Forceful Output Finalization**: To prevent this, council members proactively finalize the valid output within the
  scanning period using a multi-signature transaction. This action ensures that the invalid output, despite its
  connected edge being confirmed, cannot be finalized thereafter.

```solidity
/**
 * @notice Finalize a valid output by Security Council in the case of ZK soundness error.
 *
 * @param outputRoot           The output root of the checkpoint output to finalize forcefully.
 * @param confirmedByZKEdgeIds The ids of the rival edges confirmed by ZK proof.
 */
function finalizeOutputBySecurityCouncil(
    bytes32 outputRoot,
    bytes32[] confirmedByZKEdgeIds
) external onlySecurityCouncil {
    // 1. Ensure that the output corresponding to outputRoot exists in l2Outputs and is in the PENDING status
    // 2. Check if challengePeriodBlocks has elapsed since the proposal of the output
    // 3. Check if finalizationPeriodBlocks has not elapsed since the proposal of the output
    // 4. Ensure that the previous output connected to the output is the most recently finalized output
    // 5. Check if there is another output connected to the same previous output to ensure a challenge has
    //    created
    // 6. Ensure that the edges corresponding to confirmedByZKEdgeIds exist, and both of them have the status
    //    of CONFIRMED, and one of them is associated with outputRoot
    // 7. Ensure that the edges are rival edges
    // 8. Change the status of the output to FINALIZED
}
```
