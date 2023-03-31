# `Hashing` Invariants

## `hashCrossDomainMessage` reverts if `version` is > `0`.
**Test:** [`FuzzHashing.sol#L81`](../contracts/echidna/FuzzHashing.sol#L81)

The `hashCrossDomainMessage` function should always revert if the `version` passed is > `0`. 


## `version` = `0`: `hashCrossDomainMessage` and `hashCrossDomainMessageV0` are equivalent.
**Test:** [`FuzzHashing.sol#L93`](../contracts/echidna/FuzzHashing.sol#L93)

If the version passed is 0, `hashCrossDomainMessage` and `hashCrossDomainMessageV0` should be equivalent. 
