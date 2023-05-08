# `L2OutputOracle` Invariants

## The block number of the checkpoint output should monotonically increase.
**Test:** [`L2OutputOracle.t.sol#L64`](../contracts/test/invariants/L2OutputOracle.t.sol#L64)

When a new output is submitted, it should never be allowed to correspond to a block number that is less than the current output. 
