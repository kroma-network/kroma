// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

/* Contract Imports */
import { KanvasMintableERC20 } from "../universal/KanvasMintableERC20.sol";
import { Semver } from "./Semver.sol";

/**
 * @custom:proxied
 * @custom:predeployed 0x420000000000000000000000000000000000000B
 * @title KanvasMintableERC20Factory
 * @notice KanvasMintableERC20Factory is a factory contract that generates KanvasMintableERC20
 *         contracts on the network it's deployed to. Simplifies the deployment process for users
 *         who may be less familiar with deploying smart contracts. Designed to be backwards
 *         compatible with the older StandardL2ERC20Factory contract.
 */
contract KanvasMintableERC20Factory is Semver {
    /**
     * @notice Address of the StandardBridge on this chain.
     */
    address public immutable BRIDGE;

    /**
     * @notice Emitted whenever a new KanvasMintableERC20 is created.
     *
     * @param localToken  Address of the created token on the local chain.
     * @param remoteToken Address of the corresponding token on the remote chain.
     * @param deployer    Address of the account that deployed the token.
     */
    event KanvasMintableERC20Created(
        address indexed localToken,
        address indexed remoteToken,
        address deployer
    );

    /**
     * @custom:semver 0.1.0
     *
     * @notice The semver MUST be bumped any time that there is a change in
     *         the KanvasMintableERC20 token contract since this contract
     *         is responsible for deploying KanvasMintableERC20 contracts.
     *
     * @param _bridge Address of the StandardBridge on this chain.
     */
    constructor(address _bridge) Semver(0, 1, 0) {
        BRIDGE = _bridge;
    }

    /**
     * @notice Creates an instance of the KanvasMintableERC20 contract.
     *
     * @param _remoteToken Address of the token on the remote chain.
     * @param _name        ERC20 name.
     * @param _symbol      ERC20 symbol.
     *
     * @return Address of the newly created token.
     */
    function createKanvasMintableERC20(
        address _remoteToken,
        string memory _name,
        string memory _symbol
    ) public returns (address) {
        require(
            _remoteToken != address(0),
            "KanvasMintableERC20Factory: must provide remote token address"
        );

        address localToken = address(
            new KanvasMintableERC20(BRIDGE, _remoteToken, _name, _symbol)
        );

        emit KanvasMintableERC20Created(localToken, _remoteToken, msg.sender);

        return localToken;
    }
}
