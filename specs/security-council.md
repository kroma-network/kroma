# Security Council

[g-l1]: glossary.md#layer-1-l1

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Overview](#overview)
- [Guardian of Fault Proof System](#guardian-of-fault-proof-system)
- [Guardian of Bridge](#guardian-of-bridge)
- [Contract Upgrades](#contract-upgrades)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

The Security Council comprises a group of trusted parties with the responsibility of safeguarding the blockchain's
security. They are minimally eligible for ensuring the blockchain's safety concerning the Fault Proof System, bridge
operations, and contract upgrades. No single member of the Security Council can unilaterally execute actions. Instead,
all actions are carried out through multi-sig transaction or governance processes.

## Guardian of Fault Proof System

When an undeniable bug occurs within the [ZK Fault Proof System](./glossary.md#zk-fault-proof), assets locked in
Layer 2 may be exposed to potential risks. To prevent this, the Security Council has the authority to rectify such
issues. The Security Council intervenes in cases where two valid and contradictory ZK proofs exist
([ZK soundness error](challenge.md#dismiss-challenge)) or fail to prove with a valid proof
([ZK completeness error](challenge.md#force-delete-output)). Their intervention aims to prevent invalid outputs from
being finalized, thereby safeguarding the assets locked in Layer 2.

## Guardian of Bridge

If potential threats exposing Layer 2 assets within the Bridge, the Security Council possesses the authority to promptly
pause/unpause the bridge through a multi-sig transaction. The `GUARDIAN` in the
[`KromaPortal.sol`](../packages/contracts/contracts/L1/KromaPortal.sol) is configured to be Security Council.

```solidity
    /**
     * @notice Pause deposits and withdrawals.
     */
    function pause() external {
        require(msg.sender == GUARDIAN, "KromaPortal: only guardian can pause");
        paused = true;
        emit Paused(msg.sender);
    }

    /**
     * @notice Unpause deposits and withdrawals.
     */
    function unpause() external {
        require(msg.sender == GUARDIAN, "KromaPortal: only guardian can unpause");
        paused = false;
        emit Unpaused(msg.sender);
    }
```

## Contract Upgrades

All contract upgrades deployed on [Layer 1][g-l1] are conducted by the
[governance of the Security Council](contract-upgrades.md#upgrade-by-governance). These upgrades are proposed by a
member of Security Council and are determined through the voting of these members. If a proposal passes, it typically
has a 30-day timelock delay for execution.

When an urgent upgrade is needed for a blockchain security, the Security Council has the authority to perform an
[emergency upgrade](contract-upgrades.md#upgrade-by-governance).
