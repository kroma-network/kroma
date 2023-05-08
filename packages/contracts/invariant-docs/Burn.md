# `Burn` Invariants

## `eth(uint256)` always burns the exact amount of eth passed.
**Test:** [`FuzzBurn.sol#L36`](../contracts/echidna/FuzzBurn.sol#L36)

Asserts that when `Burn.eth(uint256)` is called, it always burns the exact amount of ETH passed to the function. 


## `gas(uint256)` always burns at least the amount of gas passed.
**Test:** [`FuzzBurn.sol#L78`](../contracts/echidna/FuzzBurn.sol#L78)

Asserts that when `Burn.gas(uint256)` is called, it always burns at least the amount of gas passed to the function. 
