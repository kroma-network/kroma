// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

/**
 * @title ITokenMultiSigWallet
 * @notice Interface for contracts of a token based multi-signature wallet.
 */
interface ITokenMultiSigWallet {
    /**
     * @notice Emitted when anyone submit a transaction.
     *
     * @param sender        Address of submitter.
     * @param transactionId The ID of transaction submitted.
     */
    event TransactionSubmitted(address indexed sender, uint256 indexed transactionId);

    /**
     * @notice Emitted when anyone confirm a transaction.
     *
     * @param sender        Owner of address that confirm a transaction.
     * @param transactionId The ID of transaction confirmed.
     */
    event TransactionConfirmed(address indexed sender, uint256 indexed transactionId);

    /**
     * @notice Emitted when transaction is executed.
     *
     * @param sender        Owner of address that execute a transaction.
     * @param transactionId The ID of transaction executed.
     */
    event TransactionExecuted(address indexed sender, uint256 indexed transactionId);

    /**
     * @notice Emitted when anyone revoke a confirmation.
     *
     * @param sender        Owner of address that revoke a transaction.
     * @param transactionId The ID of transaction to revoke.
     */
    event ConfirmationRevoked(address indexed sender, uint256 indexed transactionId);

    /**
     * @notice Allows an owner to submit and confirm a transaction.
     *
     * @param _target Transaction target address.
     * @param _value  Transaction ether value.
     * @param _data   Transaction data payload.
     *
     * @return Returns transaction ID.
     */
    function submitTransaction(
        address _target,
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
     * @notice Allows an owner to revoke a transaction.
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
     * @notice Returns the current quorum, in terms of number of votes.
     *
     * @return Current quorum, in terms of number of votes: `supply * quorumNumerator / quorumDenominator`.
     */
    function quorum() external view returns (uint256);

    /**
     * @notice Returns the number of votes.
     *
     * @param _account Account to check votes.
     *
     * @return Number of votes.
     */
    function getVotes(address _account) external view returns (uint256);

    /**
     * @notice Returns the number of confirmations that account has confirmed.
     *
     * @param _transactionId Transaction id to check.
     *
     * @return The number of confirmations.
     */
    function getConfirmationCount(uint256 _transactionId) external view returns (uint256);

    /**
     * @notice Returns whether the account has confirmed the transaction.
     *
     * @param _transactionId Transaction id to check.
     * @param _account       Address to check.
     *
     * @return Confirmed status.
     */
    function isConfirmedBy(uint256 _transactionId, address _account) external view returns (bool);
}
