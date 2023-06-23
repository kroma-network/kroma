// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {
    ReentrancyGuardUpgradeable
} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import { SafeCall } from "../libraries/SafeCall.sol";
import { Types } from "../libraries/Types.sol";
import { IMultiSigWallet } from "./IMultiSigWallet.sol";

/**
 * @custom:upgradeable
 * @title MultiSigWallet
 * @notice This contract implements `IMultiSigWallet`.
 *         Allows multiple parties to agree on transactions before execution.
 */
abstract contract MultiSigWallet is IMultiSigWallet, ReentrancyGuardUpgradeable {
    /**
     * @notice A mapping of transactions submitted.
     */
    mapping(uint256 => Types.MultiSigTransaction) public transactions;

    /**
     * @notice A mapping of confirmations.
     */
    mapping(uint256 => mapping(address => bool)) public confirmations;

    /**
     * @notice A mapping that indicates whether someone is an owner or not.
     */
    mapping(address => bool) public isOwner;

    /**
     * @notice A list of owners.
     */
    address[] public owners;

    /**
     * @notice The number of confirmations required to execute a transaction.
     */
    uint256 public numConfirmationsRequired;

    /**
     * @notice The number of transactions submitted.
     */
    uint256 public transactionCount;

    /**
     * @notice Only allow this contract to call the functions.
     *         This ensures that function is only executed through a multisig-based process.
     */
    modifier onlyWallet() {
        require(
            msg.sender == address(this),
            "MultiSigWallet: only allow this contract to call the functions"
        );
        _;
    }

    /**
     * @notice Ensure that the caller is not owner.
     *
     * @param _addr Address resource requested.
     */
    modifier ownerDoesNotExist(address _addr) {
        require(!isOwner[_addr], "MultiSigWallet: owner exists");
        _;
    }

    /**
     * @notice Ensure that the caller is owner.
     *
     * @param _addr Address resource requested.
     */
    modifier ownerExists(address _addr) {
        require(isOwner[_addr], "MultiSigWallet: owner does not exist");
        _;
    }

    /**
     * @notice Ensure that the transaction exists.
     *
     * @param _transactionId Index of submitted transaction requested.
     */
    modifier transactionExists(uint256 _transactionId) {
        require(
            transactions[_transactionId].destination != address(0),
            "MultiSigWallet: transaction does not exist"
        );
        _;
    }

    /**
     * @notice Ensure that the transaction with id and owner is confirmed.
     *
     * @param _transactionId Index of submitted transaction requested.
     * @param _owner         Address resource requested.
     */
    modifier confirmed(uint256 _transactionId, address _owner) {
        require(
            confirmations[_transactionId][_owner],
            "MultiSigWallet: transaction with id and owner is not confirmed"
        );
        _;
    }

    /**
     * @notice Ensure that the transaction with id and owner is not confirmed.
     *
     * @param _transactionId Index of submitted transaction requested.
     * @param _owner         Address resource requested.
     */
    modifier notConfirmed(uint256 _transactionId, address _owner) {
        require(
            !confirmations[_transactionId][_owner],
            "MultiSigWallet: transaction with id and owner is confirmed"
        );
        _;
    }

    /**
     * @notice Ensure that the transaction is not executed.
     *
     * @param _transactionId Index of submitted transaction requested.
     */
    modifier notExecuted(uint256 _transactionId) {
        require(
            !transactions[_transactionId].executed,
            "MultiSigWallet: transaction with id is already executed"
        );
        _;
    }

    /**
     * @notice Ensure that the address is not zero address.
     *
     * @param _addr Address resource requested.
     */
    modifier validAddress(address _addr) {
        require(_addr != address(0), "MultiSigWallet: address is not valid");
        _;
    }

    /**
     * @notice Ensure that the number of confirmations required is valid.
     *
     * @param _ownerCount               Number of owners.
     * @param _numConfirmationsRequired Number of required confirmations.
     */
    modifier validNumConfirmations(uint256 _ownerCount, uint256 _numConfirmationsRequired) {
        require(
            _numConfirmationsRequired <= _ownerCount &&
                _numConfirmationsRequired != 0 &&
                _ownerCount != 0,
            "MultiSigWallet: number of required confirmation is not valid"
        );
        _;
    }

    /**
     * @notice Initializer.
     *
     * @param _owners                   List of initial owners.
     * @param _numConfirmationsRequired Number of required confirmations.
     */
    function initialize(address[] memory _owners, uint256 _numConfirmationsRequired)
        public
        onlyInitializing
        validNumConfirmations(_owners.length, _numConfirmationsRequired)
    {
        for (uint256 i = 0; i < _owners.length; ) {
            address owner = _owners[i];
            require(!isOwner[owner], "MultiSigWallet: owner already exists");
            require(owner != address(0), "MultiSigWallet: invalid owner address");
            isOwner[owner] = true;

            unchecked {
                ++i;
            }
        }
        owners = _owners;
        numConfirmationsRequired = _numConfirmationsRequired;
    }

    /**
     * @inheritdoc IMultiSigWallet
     */
    function addOwner(address _owner)
        external
        validAddress(_owner)
        onlyWallet
        ownerDoesNotExist(_owner)
    {
        isOwner[_owner] = true;
        owners.push(_owner);
        emit OwnerAddition(_owner);
    }

    /**
     * @inheritdoc IMultiSigWallet
     */
    function removeOwner(address _owner) external onlyWallet ownerExists(_owner) {
        isOwner[_owner] = false;
        // find & delete item
        for (uint256 i = 0; i < owners.length - 1; ) {
            if (owners[i] == _owner) {
                owners[i] = owners[owners.length - 1];
                owners.pop();
                break;
            }

            unchecked {
                ++i;
            }
        }

        if (numConfirmationsRequired > owners.length) {
            _changeNumConfirmationRequirement(owners.length);
        }
        emit OwnerRemoval(_owner);
    }

    /**
     * @inheritdoc IMultiSigWallet
     */
    function replaceOwner(address _owner, address _newOwner)
        external
        onlyWallet
        validAddress(_newOwner)
        ownerExists(_owner)
        ownerDoesNotExist(_newOwner)
    {
        for (uint256 i = 0; i < owners.length; ) {
            if (owners[i] == _owner) {
                owners[i] = _newOwner;
                break;
            }

            unchecked {
                ++i;
            }
        }

        isOwner[_owner] = false;
        isOwner[_newOwner] = true;
        emit OwnerRemoval(_owner);
        emit OwnerAddition(_newOwner);
    }

    /**
     * @notice Allows to change number of confirmations required.
     *
     * @param _numConfirmationsRequired Number of required confirmations.
     */
    function _changeNumConfirmationRequirement(uint256 _numConfirmationsRequired)
        internal
        validNumConfirmations(owners.length, _numConfirmationsRequired)
    {
        numConfirmationsRequired = _numConfirmationsRequired;
    }

    /**
     * @inheritdoc IMultiSigWallet
     */
    function changeRequirement(uint256 _numConfirmationsRequired) external onlyWallet {
        _changeNumConfirmationRequirement(_numConfirmationsRequired);
        emit RequirementChange(_numConfirmationsRequired);
    }

    /**
     * @inheritdoc IMultiSigWallet
     */
    function submitTransaction(
        address _destination,
        uint256 _value,
        bytes memory _data
    ) public virtual ownerExists(msg.sender) returns (uint256) {
        uint256 transactionId = _addTransaction(_destination, _value, _data);
        _confirmTransaction(transactionId);
        return transactionId;
    }

    /**
     * @inheritdoc IMultiSigWallet
     */
    function confirmTransaction(uint256 _transactionId) public virtual ownerExists(msg.sender) {
        _confirmTransaction(_transactionId);
    }

    /**
     * @notice Allows an owner to confirm a transaction.
     *
     * @param _transactionId Transaction ID.
     */
    function _confirmTransaction(uint256 _transactionId)
        internal
        transactionExists(_transactionId)
        notConfirmed(_transactionId, msg.sender)
    {
        confirmations[_transactionId][msg.sender] = true;
        emit Confirmation(msg.sender, _transactionId);
        _executeTransaction(_transactionId);
    }

    /**
     * @inheritdoc IMultiSigWallet
     */
    function revokeConfirmation(uint256 _transactionId)
        external
        virtual
        ownerExists(msg.sender)
        confirmed(_transactionId, msg.sender)
        notExecuted(_transactionId)
    {
        confirmations[_transactionId][msg.sender] = false;
        emit Revocation(msg.sender, _transactionId);
    }

    /**
     * @notice Internal functions. Execute a confirmed transaction.
     *
     * @param _transactionId Transaction ID.
     */
    function _executeTransaction(uint256 _transactionId)
        internal
        notExecuted(_transactionId)
        nonReentrant
    {
        if (_isConfirmed(_transactionId)) {
            Types.MultiSigTransaction storage txn = transactions[_transactionId];
            txn.executed = true;
            bool success = SafeCall.call(txn.destination, gasleft(), txn.value, txn.data);
            require(success, "MultiSigWallet: call transaction failed");
            emit Execution(_transactionId);
        }
    }

    /**
     * @inheritdoc IMultiSigWallet
     */
    function executeTransaction(uint256 _transactionId) external ownerExists(msg.sender) {
        _executeTransaction(_transactionId);
    }

    /**
     * @notice Internal functions. Returns the confirmation status of a transaction.
     *
     * @param _transactionId Transaction ID.
     *
     * @return Confirmation status.
     */
    function _isConfirmed(uint256 _transactionId) internal view returns (bool) {
        uint256 count = 0;
        mapping(address => bool) storage confirmation = confirmations[_transactionId];
        for (uint256 i = 0; i < owners.length; ) {
            if (confirmation[owners[i]]) {
                count += 1;
            }
            if (count == numConfirmationsRequired) {
                return true;
            }

            unchecked {
                ++i;
            }
        }
        return false;
    }

    /**
     * @inheritdoc IMultiSigWallet
     */
    function isConfirmed(uint256 _transactionId) external view returns (bool) {
        return _isConfirmed(_transactionId);
    }

    /**
     * @notice Adds a new transaction to the transaction mapping, if transaction does not exist yet.
     *
     * @param _destination Transaction target address.
     * @param _value       Transaction ether value.
     * @param _data        Transaction data payload.
     *
     * @return transactionId Returns transaction ID.
     */
    function _addTransaction(
        address _destination,
        uint256 _value,
        bytes memory _data
    ) internal validAddress(_destination) returns (uint256 transactionId) {
        transactionId = transactionCount;
        transactions[transactionId] = Types.MultiSigTransaction({
            destination: _destination,
            value: _value,
            data: _data,
            executed: false
        });
        transactionCount += 1;
        emit Submission(transactionId);
    }

    /**
     * @inheritdoc IMultiSigWallet
     */
    function getConfirmationCount(uint256 _transactionId) external view returns (uint256) {
        uint256 count = 0;
        mapping(address => bool) storage confirmation = confirmations[_transactionId];
        for (uint256 i = 0; i < owners.length; ) {
            if (confirmation[owners[i]]) {
                unchecked {
                    ++count;
                }
            }

            unchecked {
                ++i;
            }
        }
        return count;
    }

    /**
     * @inheritdoc IMultiSigWallet
     */
    function getTransactionCount(bool _pending, bool _executed) external view returns (uint256) {
        bool executed;
        uint256 count = 0;
        for (uint256 i = 0; i < transactionCount; ) {
            executed = transactions[i].executed;
            if ((_pending && !executed) || (_executed && executed)) {
                unchecked {
                    ++count;
                }
            }

            unchecked {
                ++i;
            }
        }
        return count;
    }

    /**
     * @inheritdoc IMultiSigWallet
     */
    function getOwners() external view returns (address[] memory) {
        return owners;
    }

    /**
     * @inheritdoc IMultiSigWallet
     */
    function getConfirmations(uint256 _transactionId) external view returns (address[] memory) {
        address[] memory confirmationsTemp = new address[](owners.length);
        uint256 count = 0;
        uint256 i;
        mapping(address => bool) storage confirmation = confirmations[_transactionId];
        for (i = 0; i < owners.length; ) {
            if (confirmation[owners[i]]) {
                confirmationsTemp[count] = owners[i];
                unchecked {
                    ++count;
                }
            }

            unchecked {
                ++i;
            }
        }
        address[] memory _confirmations = new address[](count);
        for (i = 0; i < count; ) {
            _confirmations[i] = confirmationsTemp[i];

            unchecked {
                ++i;
            }
        }
        return _confirmations;
    }

    /**
     * @inheritdoc IMultiSigWallet
     */
    function getTransactionIds(
        uint256 _from,
        uint256 _to,
        bool _pending,
        bool _executed
    ) external view returns (uint256[] memory) {
        bool executed;
        uint256 count = 0;
        uint256 i = 0;
        uint256[] memory _transactionIdsTemp = new uint256[](_to - _from);
        Types.MultiSigTransaction memory transaction;
        for (i = _from; i < _to; ) {
            transaction = transactions[i];
            executed = transaction.executed;
            if ((_pending && !executed) || (_executed && executed)) {
                _transactionIdsTemp[count] = i;
                unchecked {
                    ++count;
                }
            }
            unchecked {
                ++i;
            }
        }

        uint256[] memory _transactionIds = new uint256[](count);
        for (i = 0; i < count; ) {
            _transactionIds[i] = _transactionIdsTemp[i];

            unchecked {
                ++i;
            }
        }
        return _transactionIds;
    }
}
