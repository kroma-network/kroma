# `KanvasPortal` Invariants



## Deposits of any value should always succeed unless `_to` = `address(0)` or `_isCreation` = `true`.
**Test:** [`FuzzKanvasPortal.sol#L38`](../contracts/echidna/FuzzKanvasPortal.sol#L38)

All deposits, barring creation transactions and transactions sent to `address(0)`, should always succeed. 
