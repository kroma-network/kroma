# `KromaPortal` Invariants



## Deposits of any value should always succeed unless `_to` = `address(0)` or `_isCreation` = `true`.
**Test:** [`FuzzKromaPortal.sol#L38`](../contracts/echidna/FuzzKromaPortal.sol#L38)

All deposits, barring creation transactions and transactions sent to `address(0)`, should always succeed.
