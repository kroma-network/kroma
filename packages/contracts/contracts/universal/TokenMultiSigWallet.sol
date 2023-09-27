// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "@openzeppelin/contracts-upgradeable/interfaces/IERC5805Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/math/SafeCastUpgradeable.sol";

import { UpgradeGovernor } from "../governance/UpgradeGovernor.sol";
import { SafeCall } from "../libraries/SafeCall.sol";
import { Types } from "../libraries/Types.sol";
import { ITokenMultiSigWallet } from "./ITokenMultiSigWallet.sol";

/**
 * @custom:upgradeable
 * @title TokenMultiSigWallet
 * @notice This contract implements `ITokenMultiSigWallet`.
 *         Allows multiple parties to agree on transactions before execution.
 */
abstract contract TokenMultiSigWallet is ITokenMultiSigWallet, ReentrancyGuardUpgradeable {
    /**
     * @notice The address of the governor contract. Can be updated via upgrade.
     */
    UpgradeGovernor public immutable GOVERNOR;

    /**
     * @notice A mapping of transactions submitted.
     */
    mapping(uint256 => Types.MultiSigTransaction) public transactions;

    /**
     * @notice A mapping of confirmations.
     */
    mapping(uint256 => Types.MultiSigConfirmation) public confirmations;

    /**
     * @notice Spacer for backwards compatibility.
     */
    uint256[3] private spacer_53_0_96;

    /**
     * @notice The number of transactions submitted.
     */
    uint256 public transactionCount;

    /**
     * @notice Only allow the owner of governance token to call the functions.
     *         This ensures that function is only executed by governance.
     */
    modifier onlyTokenOwner(address _address) {
        require(
            getVotes(_address) > 0,
            "TokenMultiSigWallet: only allowed to governance token owner"
        );
        _;
    }

    /**
     * @notice Ensure that the transaction exists.
     *
     * @param _transactionId The ID of submitted transaction requested.
     */
    modifier transactionExists(uint256 _transactionId) {
        require(
            transactions[_transactionId].target != address(0),
            "TokenMultiSigWallet: transaction does not exist"
        );
        _;
    }

    /**
     * @notice Ensure that the transaction not exceuted.
     *
     * @param _transactionId The ID of transaction to check.
     */
    modifier transactionNotExcuted(uint256 _transactionId) {
        require(!transactions[_transactionId].executed, "TokenMultiSigWallet: already executed");
        _;
    }

    /**
     * @notice Ensure that the address is not zero address.
     *
     * @param _address Address resource requested.
     */
    modifier validAddress(address _address) {
        require(_address != address(0), "TokenMultiSigWallet: address is not valid");
        _;
    }

    /**
     * @param _governor Address of the Governor contract.
     */
    constructor(address payable _governor) {
        GOVERNOR = UpgradeGovernor(_governor);
    }

    /**
     * @inheritdoc ITokenMultiSigWallet
     */
    function submitTransaction(
        address _target,
        uint256 _value,
        bytes memory _data
    ) public onlyTokenOwner(msg.sender) returns (uint256) {
        return _submitTransaction(_target, _value, _data);
    }

    function _submitTransaction(
        address _target,
        uint256 _value,
        bytes memory _data
    ) internal validAddress(_target) returns (uint256) {
        uint256 transactionId = generateTransactionId(_target, _value, _data);
        require(
            transactions[transactionId].target == address(0),
            "TokenMultiSigWallet: transaction already exists"
        );

        transactions[transactionId] = Types.MultiSigTransaction({
            target: _target,
            value: _value,
            data: _data,
            executed: false
        });

        unchecked {
            ++transactionCount;
        }

        emit TransactionSubmitted(msg.sender, transactionId);
        return transactionId;
    }

    /**
     * @inheritdoc ITokenMultiSigWallet
     */
    function confirmTransaction(uint256 _transactionId)
        public
        onlyTokenOwner(msg.sender)
        transactionExists(_transactionId)
    {
        Types.MultiSigConfirmation storage confirms = confirmations[_transactionId];
        require(!confirms.confirmedBy[msg.sender], "TokenMultiSigWallet: already confirmed");
        confirms.confirmedBy[msg.sender] = true;
        confirms.confirmationCount += getVotes(msg.sender);
        emit TransactionConfirmed(msg.sender, _transactionId);

        // execute transaction if condition is met.
        if (confirmations[_transactionId].confirmationCount >= quorum()) {
            executeTransaction(_transactionId);
        }
    }

    /**
     * @inheritdoc ITokenMultiSigWallet
     */
    function revokeConfirmation(uint256 _transactionId)
        public
        onlyTokenOwner(msg.sender)
        transactionExists(_transactionId)
        transactionNotExcuted(_transactionId)
    {
        require(
            isConfirmedBy(_transactionId, msg.sender),
            "TokenMultiSigWallet: not confirmed yet"
        );

        Types.MultiSigConfirmation storage confirms = confirmations[_transactionId];
        confirms.confirmedBy[msg.sender] = false;
        confirms.confirmationCount -= getVotes(msg.sender);
        emit ConfirmationRevoked(msg.sender, _transactionId);
    }

    /**
     * @inheritdoc ITokenMultiSigWallet
     */
    function executeTransaction(uint256 _transactionId)
        public
        nonReentrant
        transactionExists(_transactionId)
        transactionNotExcuted(_transactionId)
    {
        require(isConfirmed(_transactionId), "TokenMultiSigWallet: quorum not reached");

        Types.MultiSigTransaction storage txn = transactions[_transactionId];
        txn.executed = true;
        bool success = SafeCall.call(txn.target, gasleft(), txn.value, txn.data);
        require(success, "TokenMultiSigWallet: call transaction failed");
        emit TransactionExecuted(msg.sender, _transactionId);
    }

    /**
     * @inheritdoc ITokenMultiSigWallet
     */
    function isConfirmed(uint256 _transactionId) public view returns (bool) {
        return confirmations[_transactionId].confirmationCount >= quorum();
    }

    /**
     * @inheritdoc ITokenMultiSigWallet
     */
    function quorum() public view returns (uint256) {
        uint256 currentTimepoint = clock() - 1;
        return
            (IERC5805Upgradeable(address(GOVERNOR.token())).getPastTotalSupply(currentTimepoint) *
                GOVERNOR.quorumNumerator(currentTimepoint)) / GOVERNOR.quorumDenominator();
    }

    /**
     * @inheritdoc ITokenMultiSigWallet
     */
    function getVotes(address account) public view returns (uint256) {
        return IERC5805Upgradeable(address(GOVERNOR.token())).getVotes(account);
    }

    /**
     * @inheritdoc ITokenMultiSigWallet
     */
    function isConfirmedBy(uint256 _transactionId, address _account) public view returns (bool) {
        return confirmations[_transactionId].confirmedBy[_account];
    }

    /**
     * @inheritdoc ITokenMultiSigWallet
     */
    function getConfirmationCount(uint256 _transactionId) public view returns (uint256) {
        return confirmations[_transactionId].confirmationCount;
    }

    /**
     * @notice Generate id of the transaction.
     *
     * @param _target Transaction target address.
     * @param _value  Transaction ether value.
     * @param _data   Transaction data payload.
     *
     * @return Generated transaction id.
     */
    function generateTransactionId(
        address _target,
        uint256 _value,
        bytes memory _data
    ) public view validAddress(_target) returns (uint256) {
        return uint256(keccak256(abi.encode(_target, _value, _data, clock())));
    }

    /**
     * @dev Clock (as specified in EIP-6372) is set to match the token's clock.
     *      Fallback to block numbers if the token does not implement EIP-6372.
     */
    function clock() public view returns (uint48) {
        try IERC5805Upgradeable(address(GOVERNOR.token())).clock() returns (uint48 timepoint) {
            return timepoint;
        } catch {
            return SafeCastUpgradeable.toUint48(block.number);
        }
    }
}
