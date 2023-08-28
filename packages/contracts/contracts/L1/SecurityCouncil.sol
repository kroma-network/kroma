// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { MultiSigWallet } from "../universal/MultiSigWallet.sol";
import { Semver } from "../universal/Semver.sol";
import { Colosseum } from "./Colosseum.sol";

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
     * @notice A mapping of outputs requested to be deleted.
     */
    mapping(uint256 => bool) public outputsDeleteRequested;

    /**
     * @notice Emitted when a validation request is submitted.
     *
     * @param transactionId Index of the submitted transaction.
     * @param outputRoot    The L2 output of the checkpoint block to be validated.
     * @param l2BlockNumber The L2 block number to be validated.
     */
    event ValidationRequested(
        uint256 indexed transactionId,
        bytes32 outputRoot,
        uint256 l2BlockNumber
    );

    /**
     * @notice Emitted when an output is requested to be deleted.
     *
     * @param transactionId Index of the requested transaction.
     * @param outputIndex   Index of output to be deleted.
     */
    event DeletionRequested(uint256 indexed transactionId, uint256 indexed outputIndex);

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
     * @custom:semver 1.0.0
     *
     * @param _colosseum Address of the Colosseum contract.
     * @param _governor  Address of governor contract.
     */
    constructor(address _colosseum, address _governor) MultiSigWallet(_governor) Semver(1, 0, 0) {
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
     * @param _outputRoot    The L2 output of the checkpoint block to be validated.
     * @param _l2BlockNumber The L2 block number to be validated.
     * @param _data          Calldata for callback purpose.
     */
    function requestValidation(
        bytes32 _outputRoot,
        uint256 _l2BlockNumber,
        bytes memory _data
    ) public onlyColosseum {
        uint256 transactionId = _addTransaction(msg.sender, 0, _data);
        emit ValidationRequested(transactionId, _outputRoot, _l2BlockNumber);
    }

    /**
     * @notice Requests to delete an output to Colosseum forcefully.
     *         This should only be called by one of the Security Council when undeniable bugs occur.
     *
     * @param _outputIndex Index of output to be deleted.
     * @param _force       Option to forcibly make a request to delete the output.
     */
    function requestDeletion(uint256 _outputIndex, bool _force) public ownerExists(msg.sender) {
        require(
            !outputsDeleteRequested[_outputIndex] || _force,
            "SecurityCouncil: the output has already been requested to be deleted"
        );
        bytes memory message = abi.encodeWithSelector(
            Colosseum.forceDeleteOutput.selector,
            _outputIndex
        );
        uint256 transactionId = submitTransaction(COLOSSEUM, 0, message);
        outputsDeleteRequested[_outputIndex] = true;
        emit DeletionRequested(transactionId, _outputIndex);
    }
}
