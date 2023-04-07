# Standard Bridges

<!-- All glossary references in this file. -->

[g-l1]: glossary.md#layer-1-l1
[g-l2]: glossary.md#layer-2-l2

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Token Depositing](#token-depositing)
- [Upgradeability](#upgradeability)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

The standard bridges are responsible for allowing cross domain
ETH and ERC20 token transfers. They are built on top of the cross domain
messenger contracts and give a standard interface for depositing tokens.

The `L2StandardBridge` is a predeploy contract located at
`0x4200000000000000000000000000000000000009`.

```solidity
interface StandardBridge {
    event ETHBridgeInitiated(
        address indexed from,
        address indexed to,
        uint256 amount,
        bytes extraData
    );

    event ETHBridgeFinalized(
      address indexed from,
      address indexed to,
      uint256 amount,
      bytes extraData
    );

    event ERC20BridgeInitiated(
      address indexed localToken,
      address indexed remoteToken,
      address indexed from,
      address to,
      uint256 amount,
      bytes extraData
    );

    event ERC20BridgeFinalized(
      address indexed localToken,
      address indexed remoteToken,
      address indexed from,
      address to,
      uint256 amount,
      bytes extraData
    );

    function bridgeERC20(
      address _localToken,
      address _remoteToken,
      uint256 _amount,
      uint32 _minGasLimit,
      bytes calldata _extraData
    ) external;

    function bridgeERC20To(
      address _localToken,
      address _remoteToken,
      address _to,
      uint256 _amount,
      uint32 _minGasLimit,
      bytes calldata _extraData
    ) external;

    function bridgeETH(
      uint32 _minGasLimit,
      bytes calldata _extraData
    ) payable external;

    function bridgeETHTo(
      address _to,
      uint32 _minGasLimit,
      bytes calldata _extraData
    ) payable external;

    function deposits(address, address) view external returns (uint256);

    function finalizeBridgeERC20(
      address _localToken,
      address _remoteToken,
      address _from,
      address _to,
      uint256 _amount,
      bytes calldata _extraData
    ) external;

    function finalizeBridgeETH(
      address _from,
      address _to,
      uint256 _amount,
      bytes calldata _extraData
    ) payable external;

    function MESSENGER() view external returns (address);

    function OTHER_BRIDGE() view external returns (address);
}
```

## Token Depositing

The `bridgeERC20` function is used to send a token from one domain to another
domain. An `KanvasMintableERC20` token contract must exist on the remote
domain to be able to deposit tokens to that domain. One of these tokens can be
deployed using the `KanvasMintableERC20Factory` contract.

## Upgradeability

Both the [L1][g-l1] and [L2][g-l2] standard bridges should be behind upgradable proxies.
