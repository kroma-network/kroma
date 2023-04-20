// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

/* Contract Imports */
import { KromaMintableERC20 } from "../universal/KromaMintableERC20.sol";
import { Semver } from "./Semver.sol";

/**
 * @custom:proxied
 * @custom:predeployed 0x420000000000000000000000000000000000000B
 * @title KromaMintableERC20Factory
 * @notice KromaMintableERC20Factory is a factory contract that generates KromaMintableERC20
 *         contracts on the network it's deployed to. Simplifies the deployment process for users
 *         who may be less familiar with deploying smart contracts. Designed to be backwards
 *         compatible with the older StandardL2ERC20Factory contract.
 */
contract KromaMintableERC20Factory is Semver {
    /**
     * @notice Address of the StandardBridge on this chain.
     */
    address public immutable BRIDGE;

    /**
     * @notice Emitted whenever a new KromaMintableERC20 is created.
     *
     * @param localToken  Address of the created token on the local chain.
     * @param remoteToken Address of the corresponding token on the remote chain.
     * @param deployer    Address of the account that deployed the token.
     */
    event KromaMintableERC20Created(
        address indexed localToken,
        address indexed remoteToken,
        address deployer
    );

    /**
     * @custom:semver 0.1.0
     *
     * @notice The semver MUST be bumped any time that there is a change in
     *         the KromaMintableERC20 token contract since this contract
     *         is responsible for deploying KromaMintableERC20 contracts.
     *
     * @param _bridge Address of the StandardBridge on this chain.
     */
    constructor(address _bridge) Semver(0, 1, 0) {
        BRIDGE = _bridge;
    }

    /**
     * @notice Creates an instance of the KromaMintableERC20 contract.
     *
     * @param _remoteToken Address of the token on the remote chain.
     * @param _name        ERC20 name.
     * @param _symbol      ERC20 symbol.
     *
     * @return Address of the newly created token.
     */
    function createKromaMintableERC20(
        address _remoteToken,
        string memory _name,
        string memory _symbol
    ) public returns (address) {
        require(
            _remoteToken != address(0),
            "KromaMintableERC20Factory: must provide remote token address"
        );

        address localToken = address(
            new KromaMintableERC20(BRIDGE, _remoteToken, _name, _symbol)
        );

        emit KromaMintableERC20Created(localToken, _remoteToken, msg.sender);

        return localToken;
    }
}
