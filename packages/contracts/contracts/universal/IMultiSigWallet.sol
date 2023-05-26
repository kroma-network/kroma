// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

/**
 * @title IMultiSigWallet
 * @notice Interface for contracts that are compatible with the Safe legacy MultiSigWallet.
 */
interface IMultiSigWallet {
    /**
     * @notice Emitted when anyone confirm a transaction.
     *
     * @param sender        Owner of address that confirm a transaction.
     * @param transactionId Index of transaction confirmed.
     */
    event Confirmation(address indexed sender, uint256 indexed transactionId);

    /**
     * @notice Emitted when anyone revoke a transaction.
     *
     * @param sender        Owner of address that revoke a transaction.
     * @param transactionId Index of transaction revoked.
     */
    event Revocation(address indexed sender, uint256 indexed transactionId);

    /**
     * @notice Emitted when anyone submit a transaction.
     *
     * @param transactionId Index of transaction submitted.
     */
    event Submission(uint256 indexed transactionId);

    /**
     * @notice Emitted when transaction is executed.
     *
     * @param transactionId Index of transaction executed.
     */
    event Execution(uint256 indexed transactionId);

    /**
     * @notice Emitted when transaction is executed but failed.
     *
     * @param transactionId Index of transaction failed to execute.
     */
    event ExecutionFailure(uint256 indexed transactionId);

    /**
     * @notice Emitted when an owner address is added.
     *
     * @param owner Owner address that added.
     */
    event OwnerAddition(address indexed owner);

    /**
     * @notice Emitted when an owner address is removed.
     *
     * @param owner Owner address that removed.
     */
    event OwnerRemoval(address indexed owner);

    /**
     * @notice Emitted when a requirement changed
     *
     * @param required Required value that changed.
     */
    event RequirementChange(uint256 required);

    /**
     * @notice Allows to add a new owner. Transaction has to be sent by wallet.
     *
     * @param _owner Address of new owner.
     */
    function addOwner(address _owner) external;

    /**
     * @notice Allows to remove an owner. Transaction has to be sent by wallet.
     *
     * @param _owner Address of owner.
     */
    function removeOwner(address _owner) external;

    /**
     * @notice Allows to replace an owner with a new owner. Transaction has to be sent by wallet.
     *
     * @param _owner    Address of owner to be replaced.
     * @param _newOwner Address of new owner.
     */
    function replaceOwner(address _owner, address _newOwner) external;

    /**
     * @notice Allows to change the number of required confirmations.
     *         Transaction has to be sent by wallet.
     *
     * @param _required Number of required confirmations.
     */
    function changeRequirement(uint256 _required) external;

    /**
     * @notice Allows an owner to submit and confirm a transaction.
     *
     * @param _destination Transaction target address.
     * @param _value       Transaction ether value.
     * @param _data        Transaction data payload.
     *
     * @return Returns transaction ID.
     */
    function submitTransaction(
        address _destination,
        uint256 _value,
        bytes memory _data
    ) external returns (uint256);

    /**
     * @notice Allows an owner to confirm a transaction.
     *
     * @param _transactionId Transaction ID.
     */
    function confirmTransaction(uint256 _transactionId) external;

    /**
     * @notice Allows an owner to revoke a confirmation for a transaction.
     *
     * @param _transactionId Transaction ID.
     */
    function revokeConfirmation(uint256 _transactionId) external;

    /**
     * @notice Allows anyone to execute a confirmed transaction.
     *
     * @param _transactionId Transaction ID.
     */
    function executeTransaction(uint256 _transactionId) external;

    /**
     * @notice Returns the confirmation status of a transaction.
     *
     * @param _transactionId Transaction ID.
     *
     * @return Confirmation status.
     */
    function isConfirmed(uint256 _transactionId) external view returns (bool);

    /**
     * @notice Returns number of confirmations of a transaction.
     *
     * @param _transactionId Transaction ID.
     *
     * @return Number of confirmations.
     */
    function getConfirmationCount(uint256 _transactionId) external view returns (uint256);

    /**
     * @notice Returns total number of transactions after filters are applied.
     *
     * @param _pending  Whether include pending transactions.
     * @param _executed Whether include executed transactions.
     *
     * @return Total number of transactions after filters are applied.
     */
    function getTransactionCount(bool _pending, bool _executed) external view returns (uint256);

    /**
     * @notice Returns list of owners.
     *
     * @return The list of owner addresses.
     */
    function getOwners() external view returns (address[] memory);

    /**
     * @notice Returns a list of owners who have confirmed the transaction.
     *
     * @param _transactionId Transaction ID.
     *
     * @return Returns array of owner addresses.
     */
    function getConfirmations(uint256 _transactionId) external view returns (address[] memory);

    /**
     * @notice Returns the list of transaction IDs in defined range.
     *
     * @param _from     The starting index of transaction array.
     * @param _to       The ending index of the transaction array.
     * @param _pending  Whether include pending transactions.
     * @param _executed Whether include executed transactions.
     *
     * @return List of the transaction IDs in a defined range.
     */
    function getTransactionIds(
        uint256 _from,
        uint256 _to,
        bool _pending,
        bool _executed
    ) external view returns (uint256[] memory);
}
