// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

/**
 * @title IAssetManager
 * @notice Interface for AssetManager contract.
 */
interface IAssetManager {
    /**
     * @notice Represents the asset information of the vault of a validator.
     *
     * @custom:field totalKro           Total amount of KRO in the vault.
     * @custom:field totalKroShares     Total shares for KRO delegation in the vault.
     * @custom:field totalKgh           Total number of KGH in the vault.
     * @custom:field totalKroInKgh      Total amount of KRO which KGHs in the vault have.
     * @custom:field totalKghShares     Total shares for KGH delegation in the vault.
     * @custom:field validatorKro       Amount of KRO that the validator self-delegated.
     * @custom:field boostedReward      Cumulated boosted reward for KGH delegators in the vault.
     * @custom:field validatorRewardKro Cumulated reward for the validator.
     */
    struct Asset {
        uint128 totalKro;
        uint128 totalKroShares;
        uint128 totalKgh;
        uint128 totalKroInKgh;
        uint128 totalKghShares;
        uint128 validatorKro;
        uint128 boostedReward;
        uint128 validatorRewardKro;
    }

    /**
     * @notice Constructs the pending asset information of the vault of a validator.
     *
     * @custom:field totalPendingAssets           Total pending KRO for undelegation.
     * @custom:field totalPendingBoostedRewards   Total pending boosted rewards in KRO for undelegation.
     * @custom:field totalPendingKroShares        Total pending KRO shares for undelegation.
     * @custom:field totalPendingKghShares        Total pending KGH shares for undelegation.
     * @custom:field totalPendingValidatorRewards Total pending validator rewards.
     * @custom:field claimRequestTimes            Timestamps of validator reward claim requests.
     * @custom:field pendingValidatorRewards      A mapping of timestamp to pending validator rewards.
     */
    struct Pending {
        uint128 totalPendingAssets;
        uint128 totalPendingBoostedRewards;
        uint128 totalPendingKroShares;
        uint128 totalPendingKghShares;
        uint128 totalPendingValidatorRewards;
        uint256[] claimRequestTimes;
        mapping(uint256 => uint128) pendingValidatorRewards;
    }

    /**
     * @notice Constructs the delegator of KRO in the vault of a validator.
     *
     * @custom:field shares                 Amount of shares for KRO delegation.
     * @custom:field undelegateRequestTimes Timestamps of undelegation requests.
     * @custom:field pendingKroShares       A mapping of timestamp undelegations are requested to
     *                                      pending KRO shares for undelegation.
     */
    struct KroDelegator {
        uint128 shares;
        uint256[] undelegateRequestTimes;
        mapping(uint256 => uint128) pendingKroShares;
    }

    /**
     * @notice Represents the amounts of KRO and KGH shares for KGH delegator.
     *
     * @custom:field kro Amount of KRO shares.
     * @custom:field kgh Amount of KGH shares.
     */
    struct KghDelegatorShares {
        uint128 kro;
        uint128 kgh;
    }

    /**
     * @notice Constructs the delegator of KGH in the vault of a validator.
     *
     * @custom:field shares                 Amount of shares for KRO in KGH and KGH itself to the
     *                                      corresponding tokenId.
     * @custom:field undelegateRequestTimes Timestamps of undelegation requests.
     * @custom:field pendingKghIds          A mapping of timestamp undelegations are requested to
     *                                      token ids of KGHs for undelegation.
     * @custom:field pendingShares          A mapping of timestamp undelegations are requested to
     *                                      pending KRO and KGH shares for undelegation.
     */
    struct KghDelegator {
        mapping(uint256 => KghDelegatorShares) shares;
        uint256[] undelegateRequestTimes;
        mapping(uint256 => uint256[]) pendingKghIds;
        mapping(uint256 => KghDelegatorShares) pendingShares;
    }

    /**
     * @notice Constructs the vault of a validator.
     *
     * @custom:field asset         Asset information of the vault.
     * @custom:field pending       Pending asset information of the vault.
     * @custom:field kroDelegators A mapping of validator address to KRO delegator struct.
     * @custom:field kghDelegators A mapping of validator address to KGH delegator struct.
     */
    struct Vault {
        Asset asset;
        Pending pending;
        mapping(address => KroDelegator) kroDelegators;
        mapping(address => KghDelegator) kghDelegators;
    }

    /**
     * @notice Emitted when KROs are delegated.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param amount    The amount of KRO delegated.
     * @param shares    The amount of shares received.
     */
    event KroDelegated(
        address indexed validator,
        address indexed delegator,
        uint128 amount,
        uint128 shares
    );

    /**
     * @notice Emitted when a KGH is delegated.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param tokenId   Token Id of the KGH.
     * @param kroInKgh  The amount of KRO in the KGH.
     * @param kroShares The amount of KRO shares received.
     * @param kghShares The amount of KGH shares received.
     */
    event KghDelegated(
        address indexed validator,
        address indexed delegator,
        uint256 tokenId,
        uint128 kroInKgh,
        uint128 kroShares,
        uint128 kghShares
    );

    /**
     * @notice Emitted when KGHs are delegated in batch.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param tokenIds  Array of token ids of the KGHs.
     * @param kroShares The amount of KRO shares received.
     * @param kghShares The amount of KGH shares received.
     */
    event KghBatchDelegated(
        address indexed validator,
        address indexed delegator,
        uint256[] tokenIds,
        uint128 kroShares,
        uint128 kghShares
    );

    /**
     * @notice Emitted when KRO undelegation is initiated.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param amount    The amount of KRO to undelegate.
     * @param shares    The amount of shares to undelegate.
     */
    event KroUndelegationInitiated(
        address indexed validator,
        address indexed delegator,
        uint128 amount,
        uint128 shares
    );

    /**
     * @notice Emitted when KGH undelegation is initiated.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param tokenId   Token Id of the KGH.
     * @param kroShares The amount of KRO shares to undelegate.
     * @param kghShares The amount of KGH shares to undelegate.
     */
    event KghUndelegationInitiated(
        address indexed validator,
        address indexed delegator,
        uint256 tokenId,
        uint128 kroShares,
        uint128 kghShares
    );

    /**
     * @notice Emitted when KGHs undelegation is initiated in batch.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param tokenIds  Array of token ids of the KGHs.
     * @param kroShares The amount of KRO shares to undelegate.
     * @param kghShares The amount of KGH shares to undelegate.
     */
    event KghBatchUndelegationInitiated(
        address indexed validator,
        address indexed delegator,
        uint256[] tokenIds,
        uint128 kroShares,
        uint128 kghShares
    );

    /**
     * @notice Emitted when KRO undelegation is finalized.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param amount    The amount of KRO undelegated.
     */
    event KroUndelegationFinalized(
        address indexed validator,
        address indexed delegator,
        uint128 amount
    );

    /**
     * @notice Emitted when KGH undelegation is finalized.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param amount    The amount of KRO undelegated during the KGH undelegation.
     */
    event KghUndelegationFinalized(
        address indexed validator,
        address indexed delegator,
        uint128 amount
    );

    /**
     * @notice Emitted when validator reward claim is initiated.
     *
     * @param validator Address of the validator.
     * @param amount    The amount of validator reward requested for claim.
     */
    event RewardClaimInitiated(address indexed validator, uint128 amount);

    /**
     * @notice Emitted when validator reward claim is finalized.
     *
     * @param validator Address of the validator.
     * @param amount    The amount of validator reward claimed.
     */
    event RewardClaimFinalized(address indexed validator, uint128 amount);

    /**
     * @notice Reverts when caller is not allowed.
     */
    error NotAllowedCaller();

    /**
     * @notice Reverts when constructor parameters are invalid.
     */
    error InvalidConstructorParams();

    /**
     * @notice Reverts when the status of validator is improper.
     */
    error ImproperValidatorStatus();

    /**
     * @notice Reverts when try to input zero.
     */
    error NotAllowedZeroInput();

    /**
     * @notice Reverts when the asset is insufficient.
     */
    error InsufficientAsset();

    /**
     * @notice Reverts when the share is insufficient.
     */
    error InsufficientShare();

    /**
     * @notice Reverts when the share does not exist.
     */
    error ShareNotExists();

    /**
     * @notice Reverts when the pending does not exist.
     */
    error PendingNotExists();

    /**
     * @notice Reverts when request does not exist.
     */
    error RequestNotExists();

    /**
     * @notice Reverts when the finalized pending does not exist.
     */
    error FinalizedPendingNotExists();

    /**
     * @notice Returns the max amount that a KRO delegator can undelegate.
     *
     * @param validator Address of validator.
     * @param delegator Address of KRO delegator.
     *
     * @return The max amount that a KRO delegator can undelegate.
     */
    function getKroTotalBalance(
        address validator,
        address delegator
    ) external view returns (uint128);

    /**
     * @notice Returns the max amount of KRO that a KGH delegator can undelegate.
     *
     * @param validator Address of validator.
     * @param delegator Address of KGH delegator.
     * @param tokenId   Token Id of the KGH.
     *
     * @return The max amount that a KGH delegator can undelegate.
     */
    function getKghTotalBalance(
        address validator,
        address delegator,
        uint256 tokenId
    ) external view returns (uint128);

    /**
     * @notice Returns the amount of KRO shares that the KRO delegator has.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     *
     * @return The amount of KRO shares that the KRO delegator has.
     */
    function getKroTotalShareBalance(
        address validator,
        address delegator
    ) external view returns (uint128);

    /**
     * @notice Returns the amount of KGH shares that the KGH delegator has.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param tokenId   Token Id of the KGH.
     *
     * @return The amount of KRO shares that the KGH delegator has.
     * @return The amount of KGH shares that the KGH delegator has.
     */
    function getKghTotalShareBalance(
        address validator,
        address delegator,
        uint256 tokenId
    ) external view returns (uint128, uint128);

    /**
     * @notice Allows an on-chain or off-chain user to simulate the effects of their KRO delegation at the current block.
     *
     * @param validator Address of the validator.
     * @param assets    The amount of assets to delegate.
     *
     * @return The amount of shares that the Vault would exchange for the amount of assets provided.
     */
    function previewDelegate(address validator, uint128 assets) external view returns (uint128);

    /**
     * @notice Allows an on-chain or off-chain user to simulate the effects of their KRO undelegation
     *         at the current block.
     *
     * @param validator The address of the validator.
     * @param shares    The amount of shares to undelegate.
     *
     * @return The amount of assets that the Vault would exchange for the amount of shares provided.
     */
    function previewUndelegate(address validator, uint128 shares) external view returns (uint128);

    /**
     * @notice Allows an on-chain or off-chain user to simulate the effects of their KGH delegation
     *         at the current block given current on-chain conditions.
     *
     * @param validator The address of the validator.
     *
     * @return The amount of shares that the Vault would exchange for the amount of assets provided.
     */
    function previewKghDelegate(address validator) external view returns (uint128);

    /**
     * @notice Allows an on-chain or off-chain user to simulate the effects of their KGH undelegation
     *         at the current block, given current on-chain conditions.
     *
     * @param validator The address of the validator.
     * @param tokenId   The tokenId of KGH to undelegate.
     *
     * @return The amount of assets that the Vault would exchange for the KGH of given tokenId.
     */
    function previewKghUndelegate(
        address validator,
        uint256 tokenId
    ) external view returns (uint128);

    /**
     * @notice Returns the total amount of KRO assets held by the vault.
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of KRO assets held by the vault.
     */
    function totalKroAssets(address validator) external view returns (uint128);

    /**
     * @notice Returns the total number of KGHs held by the vault.
     *
     * @param validator Address of the validator.
     *
     * @return The total number of KGHs held by the vault.
     */
    function totalKghNum(address validator) external view returns (uint128);

    /**
     * @notice Returns the total amount of KRO in KGH held by the vault.
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of KRO in KGH held by the vault.
     */
    function totalKroInKgh(address validator) external view returns (uint128);

    /**
     * @notice Returns the total amount of KRO a validator has self-delegated.
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of KRO a validator has self-delegated.
     */
    function totalValidatorKro(address validator) external view returns (uint128);

    /**
     * @notice Returns the reflective weight of given validator. It can be different from the actual
     *         current weight of the validator in validator tree since it includes all accumulated
     *         rewards.
     *
     * @param validator Address of the validator.
     *
     * @return The reflective weight of given validator.
     */
    function reflectiveWeight(address validator) external view returns (uint128);

    /**
     * @notice Delegate KRO to the validator and returns the amount of shares that the vault would
     *         exchange.
     *
     * @param validator Address of the validator.
     * @param assets    The amount of KRO to delegate.
     *
     * @return The amount of shares that the Vault would exchange for the amount of assets provided.
     */
    function delegate(address validator, uint128 assets) external returns (uint128);

    /**
     * @notice Delegate KGH to the validator and returns the amount of shares that the vault would
     *         exchange.
     *
     * @param validator Address of the validator.
     * @param tokenId   The tokenId of KGH to delegate.
     *
     * @return The amount of KRO shares that the Vault would exchange for the KGH provided.
     * @return The amount of KGH shares that the Vault would exchange for the KGH provided.
     */
    function delegateKgh(address validator, uint256 tokenId) external returns (uint128, uint128);

    /**
     * @notice Delegate KGHs to the validator and returns the amount of shares that the vault would
     *         exchange.
     *
     * @param validator Address of the validator.
     * @param tokenIds  The token ids of KGHs to delegate.
     *
     * @return The amount of KRO shares that the Vault would exchange for the KGHs provided.
     * @return The amount of KGH shares that the Vault would exchange for the KGHs provided.
     */
    function delegateKghBatch(
        address validator,
        uint256[] calldata tokenIds
    ) external returns (uint128, uint128);

    /**
     * @notice Initiate the KRO undelegation of given shares for the given validator.
     *
     * @param validator Address of the validator.
     * @param shares    The amount of shares to undelegate.
     */
    function initUndelegate(address validator, uint128 shares) external;

    /**
     * @notice Initiate KGH undelegation for given validator and tokenId.
     *
     * @param validator Address of the validator.
     * @param tokenId   Token id of KGH to undelegate.
     */
    function initUndelegateKgh(address validator, uint256 tokenId) external;

    /**
     * @notice Initiate KGH undelegation for given validator and token ids.
     *
     * @param validator Address of the validator.
     * @param tokenIds  Array of token ids of KGHs to undelegate.
     */
    function initUndelegateKghBatch(address validator, uint256[] calldata tokenIds) external;

    /**
     * @notice Claim the reward of the validator.
     *
     * @param amount The amount of reward to claim.
     */
    function initClaimValidatorReward(uint128 amount) external;

    /**
     * @notice Finalize all pending KRO undelegation and returns the amount of assets that the vault would
     *         exchange for the pending KRO shares.
     *
     * @param validator Address of the validator.
     *
     * @return The amount of assets that the vault would exchange for the pending KRO shares.
     */
    function finalizeUndelegate(address validator) external returns (uint128);

    /**
     * @notice Finalize all pending KGH undelegation and returns the amount of assets that the vault would
     *         exchange for the pending KRO and KGH shares.
     *
     * @param validator Address of the validator.
     *
     * @return The amount of assets that the vault would exchange for the pending KRO and KGH shares.
     */
    function finalizeUndelegateKgh(address validator) external returns (uint128);

    /**
     * @notice Finalize the reward claim of the validator.
     */
    function finalizeClaimValidatorReward() external;
}
