// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title Predeploys
 * @notice Contains constant addresses for contracts that are pre-deployed to the L2 system.
 */
library Predeploys {
    /**
     * @notice Address of the ProxyAdmin predeploy.
     */
    address internal constant PROXY_ADMIN = 0x4200000000000000000000000000000000000000;

    /**
     * @notice Address of the L1Block predeploy.
     */
    address internal constant L1_BLOCK_ATTRIBUTES = 0x4200000000000000000000000000000000000002;

    /**
     * @notice Address of the L2ToL1MessagePasser predeploy.
     */
    address internal constant L2_TO_L1_MESSAGE_PASSER = 0x4200000000000000000000000000000000000003;

    /**
     * @notice Address of the L2CrossDomainMessenger predeploy.
     */
    address internal constant L2_CROSS_DOMAIN_MESSENGER =
        0x4200000000000000000000000000000000000004;

    /**
     * @notice Address of the GasPriceOracle predeploy. Includes fee information
     *         and helpers for computing the L1 portion of the transaction fee.
     */
    address internal constant GAS_PRICE_ORACLE = 0x4200000000000000000000000000000000000005;

    /**
     * @notice Address of the ProtocolVault predeploy.
     */
    address internal constant PROTOCOL_VAULT = 0x4200000000000000000000000000000000000006;

    /**
     * @notice Address of the ProposerRewardVault predeploy.
     */
    address internal constant PROPOSER_REWARD_VAULT = 0x4200000000000000000000000000000000000007;

    /**
     * @notice Address of the ValidatorRewardVault predeploy.
     */
    address internal constant VALIDATOR_REWARD_VAULT = 0x4200000000000000000000000000000000000008;

    /**
     * @notice Address of the L2StandardBridge predeploy.
     */
    address internal constant L2_STANDARD_BRIDGE = 0x4200000000000000000000000000000000000009;

    /**
     * @notice Address of the L2ERC721Bridge predeploy.
     */
    address internal constant L2_ERC721_BRIDGE = 0x420000000000000000000000000000000000000A;

    /**
     * @notice Address of the KromaMintableERC20Factory predeploy.
     */
    address internal constant KROMA_MINTABLE_ERC20_FACTORY =
        0x420000000000000000000000000000000000000B;

    /**
     * @notice Address of the KromaMintableERC721Factory predeploy.
     */
    address internal constant KROMA_MINTABLE_ERC721_FACTORY =
        0x420000000000000000000000000000000000000c;
}
