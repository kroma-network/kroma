// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { MultiSigWallet } from "../universal/MultiSigWallet.sol";
import { Semver } from "../universal/Semver.sol";

/**
 * @custom:proxied
 * @title SecurityCouncil
 * @notice SecurityCouncil receives validation requests for specific output data,
 *         and allows security council parties to validate & agree on transactions before execution.
 */
contract SecurityCouncil is MultiSigWallet, Semver {
    /**
     * @notice The address of the colosseum contract. Can be updated via upgrade.
     */
    address public immutable COLOSSEUM;

    /**
     * @notice Emitted when a validation request is submitted.
     *
     * @param transactionId Index of the submitted transaction.
     * @param outputRoot    L2 output root that was proven against.
     * @param l2BlockNumber The L2 block number of the output root.
     */
    event ValidationRequested(
        uint256 indexed transactionId,
        bytes32 outputRoot,
        uint128 l2BlockNumber
    );

    /**
     * @notice Disallow calls from anyone except Colosseum.
     */
    modifier onlyColosseum() {
        require(
            msg.sender == COLOSSEUM,
            "SecurityCouncil: only the colosseum contract can be a sender"
        );
        _;
    }

    /**
     * @custom:semver 0.1.0
     *
     * @param _colosseum Address of the Colosseum contract.
     */
    constructor(address _colosseum) Semver(0, 1, 0) {
        COLOSSEUM = _colosseum;
    }

    /**
     * @notice Initializer.
     *
     * @param ``                        Not used. Dummy parameter to prevent override.
     * @param _owners                   List of initial owners.
     * @param _numConfirmationsRequired Number of required confirmations.
     *
     */
    function initialize(
        bool,
        address[] memory _owners,
        uint256 _numConfirmationsRequired
    ) public initializer {
        MultiSigWallet.initialize(_owners, _numConfirmationsRequired);
    }

    /**
     * @notice Allows the Colosseum to request for validate output data.
     *
     * @param _outputRoot    Output root byte data.
     * @param _l2BlockNumber L2 block number corresponding to output.
     * @param _data          Calldata for callback purpose.
     */
    function requestValidation(
        bytes32 _outputRoot,
        uint128 _l2BlockNumber,
        bytes memory _data
    ) public onlyColosseum {
        uint256 transactionId = _addTransaction(msg.sender, 0, _data);
        emit ValidationRequested(transactionId, _outputRoot, _l2BlockNumber);
    }
}
