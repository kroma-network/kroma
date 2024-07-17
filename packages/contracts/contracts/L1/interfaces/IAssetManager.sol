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
     * @custom:field validatorKro         Total amount of KRO that deposited by the validator and
     *                                    accumulated as validator reward (including validatorKroReserved).
     * @custom:field validatorKroReserved Total amount of validator KRO that reserved during output
     *                                    submission or challenge creation.
     * @custom:field totalKro             Total amount of KRO that delegated by the delegators and
     *                                    accumulated as KRO delegation reward (including totalKroInKgh).
     * @custom:field totalKroShares       Total shares for KRO delegation in the vault.
     * @custom:field totalKgh             Total number of KGH in the vault.
     * @custom:field totalKroInKgh        Total amount of KRO which KGHs in the vault have.
     * @custom:field rewardPerKghStored   Accumulated boosted reward per 1 KGH.
     */
    struct Asset {
        uint128 validatorKro;
        uint128 validatorKroReserved;
        uint128 totalKro;
        uint128 totalKroShares;
        uint128 totalKgh;
        uint128 totalKroInKgh;
        uint128 rewardPerKghStored;
    }

    /**
     * @notice Constructs the delegator of KRO in the vault of a validator.
     *
     * @custom:field shares          Amount of shares for KRO delegation.
     * @custom:field lastDelegatedAt Last timestamp when the delegator delegated. The delegator can
     *                               undelegate after UNDELEGATION_DELAY_SECONDS elapsed.
     */
    struct KroDelegator {
        uint128 shares;
        uint128 lastDelegatedAt;
    }

    /**
     * @notice Constructs the delegator of KGH in the vault of a validator.
     *
     * @custom:field lastDelegatedAt   Last timestamp when the delegator delegated. The delegator
     *                                 can undelegate after UNDELEGATION_DELAY_SECONDS elapsed.
     * @custom:field rewardPerKghPaid  Accumulated paid boosted reward per 1 KGH.
     * @custom:field kghNum            Total number of KGH delegated.
     * @custom:field delegationHistory A mapping of tokenId to the delegation timestamp.
     * @custom:field kroShares         A mapping of tokenId to the amount of shares for KRO in KGH.
     */
    struct KghDelegator {
        uint128 lastDelegatedAt;
        uint128 rewardPerKghPaid;
        uint256 kghNum;
        mapping(uint256 => uint128) delegationHistory;
        mapping(uint256 => uint128) kroShares;
    }

    /**
     * @notice Constructs the vault of a validator.
     *
     * @custom:field withdrawAccount An account where assets can be withdrawn to. Only this account
     *                               can withdraw the assets.
     * @custom:field lastDepositedAt Last timestamp when the validator deposited. The validator can
     *                               withdraw after UNDELEGATION_DELAY_SECONDS elapsed.
     * @custom:field asset           Asset information of the vault.
     * @custom:field kroDelegators   A mapping of validator address to KRO delegator struct.
     * @custom:field kghDelegators   A mapping of validator address to KGH delegator struct.
     */
    struct Vault {
        address withdrawAccount;
        uint128 lastDepositedAt;
        Asset asset;
        mapping(address => KroDelegator) kroDelegators;
        mapping(address => KghDelegator) kghDelegators;
    }

    /**
     * @notice Emitted when validator deposited KROs.
     *
     * @param validator Address of the validator.
     * @param amount    The amount of KRO deposited.
     */
    event Deposited(address indexed validator, uint128 amount);

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
     * @param tokenId   Token id of the KGH.
     * @param kroInKgh  The amount of KRO in the KGH.
     * @param kroShares The amount of KRO shares received.
     */
    event KghDelegated(
        address indexed validator,
        address indexed delegator,
        uint256 tokenId,
        uint128 kroInKgh,
        uint128 kroShares
    );

    /**
     * @notice Emitted when KGHs are delegated in batch.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param tokenIds  Array of token ids of the KGHs.
     * @param kroInKghs The amount of KRO in the KGHs.
     * @param kroShares The amount of KRO shares received.
     */
    event KghBatchDelegated(
        address indexed validator,
        address indexed delegator,
        uint256[] tokenIds,
        uint128 kroInKghs,
        uint128 kroShares
    );

    /**
     * @notice Emitted when validator withdrew KRO.
     *
     * @param validator Address of the validator.
     * @param amount    The amount of KRO the validator withdrew.
     */
    event Withdrawn(address indexed validator, uint128 amount);

    /**
     * @notice Emitted when KRO is undelegated.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param amount    The amount of KRO to undelegate.
     * @param shares    The amount of shares to be burnt.
     */
    event KroUndelegated(
        address indexed validator,
        address indexed delegator,
        uint128 amount,
        uint128 shares
    );

    /**
     * @notice Emitted when KGH is undelegated.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param tokenId   Token id of the KGH.
     * @param amount    The amount of KRO claimed as reward.
     * @param kroShares The amount of KRO shares to be burnt.
     */
    event KghUndelegated(
        address indexed validator,
        address indexed delegator,
        uint256 tokenId,
        uint128 amount,
        uint128 kroShares
    );

    /**
     * @notice Emitted when KGHs are undelegated in batch.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param tokenIds  Array of token ids of the KGHs.
     * @param amount    The amount of KRO claimed as reward.
     * @param kroShares The amount of KRO shares to be burnt.
     */
    event KghBatchUndelegated(
        address indexed validator,
        address indexed delegator,
        uint256[] tokenIds,
        uint128 amount,
        uint128 kroShares
    );

    /**
     * @notice Emitted when accumulated rewards of KGH delegation are claimed.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param amount    The amount of KRO claimed as reward.
     * @param kroShares The amount of KRO shares to be burnt.
     */
    event KghRewardClaimed(
        address indexed validator,
        address indexed delegator,
        uint128 amount,
        uint128 kroShares
    );

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
     * @notice Returns the address of withdraw account of given validator.
     *
     * @param validator Address of the validator.
     *
     * @return The address of withdraw account of given validator.
     */
    function getWithdrawAccount(address validator) external view returns (address);

    /**
     * @notice Returns when the validator can withdraw KRO. The validator can withdraw after
     *         UNDELEGATION_DELAY_SECONDS elapsed from lastDepositedAt.
     *
     * @param validator Address of the validator.
     *
     * @return When the validator can withdraw KRO.
     */
    function canWithdrawAt(address validator) external view returns (uint256);

    /**
     * @notice Returns the total amount of KRO a validator has deposited and rewarded.
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of KRO a validator has deposited and rewarded.
     */
    function totalValidatorKro(address validator) external view returns (uint128);

    /**
     * @notice Returns the total amount of KRO that delegated by the delegators and accumulated as
     *         KRO delegation reward (including totalKroInKgh).
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of KRO that delegated by the delegators and accumulated as KRO
     *         delegation reward (including totalKroInKgh).
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
     * @notice Returns the total amount of KRO in KGHs held by the vault.
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of KRO in KGHs held by the vault.
     */
    function totalKroInKgh(address validator) external view returns (uint128);

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
     * @notice Returns when the KRO delegators can undelegate KRO. The delegators can undelegate
     *         after UNDELEGATION_DELAY_SECONDS elapsed from lastDelegatedAt.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the KRO delegator.
     *
     * @return When the KRO delegators can undelegate KRO.
     */
    function canUndelegateKroAt(
        address validator,
        address delegator
    ) external view returns (uint256);

    /**
     * @notice Returns the number of KGH delegated by the given delegator.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     *
     * @return The number of KGH delegated by the given delegator.
     */
    function getKghNum(address validator, address delegator) external view returns (uint128);

    /**
     * @notice Returns the amount of KRO shares that the KGH delegator has.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param tokenId   Token id of the KGH.
     *
     * @return The amount of KRO shares that the KGH delegator has.
     */
    function getKghTotalShareBalance(
        address validator,
        address delegator,
        uint256 tokenId
    ) external view returns (uint128);

    /**
     * @notice Returns when the KGH delegators can undelegate KGH. The delegators can undelegate
     *         after UNDELEGATION_DELAY_SECONDS elapsed from lastDelegatedAt.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the KGH delegator.
     *
     * @return When the KGH delegators can undelegate KGH.
     */
    function canUndelegateKghAt(
        address validator,
        address delegator
    ) external view returns (uint256);

    /**
     * @notice Allows an on-chain or off-chain user to simulate the effects of their KRO delegation
     *         at the current block.
     *
     * @param validator Address of the validator.
     * @param assets    The amount of assets to delegate.
     *
     * @return The amount of shares that the Vault would exchange for the amount of assets provided.
     */
    function previewDelegate(address validator, uint128 assets) external view returns (uint128);

    /**
     * @notice Allows an on-chain or off-chain user to simulate the effects of their KRO
     *         undelegation at the current block.
     *
     * @param validator The address of the validator.
     * @param shares    The amount of shares to undelegate.
     *
     * @return The amount of assets that the Vault would exchange for the amount of shares provided.
     */
    function previewUndelegate(address validator, uint128 shares) external view returns (uint128);

    /**
     * @notice Returns the claimable reward of KGH delegation.
     *
     * @param validator The address of the validator.
     * @param delegator The address of the delegator.
     *
     * @return The amount of claimable reward of KGH delegation.
     */
    function getKghReward(address validator, address delegator) external view returns (uint128);

    /**
     * @notice Deposit KRO. To deposit KRO, the validator should be initiated.
     *
     * @param assets The amount of KRO to deposit.
     */
    function deposit(uint128 assets) external;

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
     * @notice Delegate KGH to the validator and returns the amount of KRO shares that the vault
     *         would exchange.
     *
     * @param validator Address of the validator.
     * @param tokenId   The token id of KGH to delegate.
     *
     * @return The amount of KRO shares that the Vault would exchange for the KGH provided.
     */
    function delegateKgh(address validator, uint256 tokenId) external returns (uint128);

    /**
     * @notice Delegate KGHs to the validator and returns the amount of KRO shares that the vault
     *         would exchange.
     *
     * @param validator Address of the validator.
     * @param tokenIds  The token ids of KGHs to delegate.
     *
     * @return The amount of KRO shares that the Vault would exchange for the KGHs provided.
     */
    function delegateKghBatch(
        address validator,
        uint256[] calldata tokenIds
    ) external returns (uint128);

    /**
     * @notice Withdraw KRO by the validator.
     *
     * @param assets The amount of KRO to withdraw.
     */
    function withdraw(uint128 assets) external;

    /**
     * @notice Undelegate the KRO of given shares for the given validator.
     *
     * @param validator Address of the validator.
     * @param shares    The amount of shares to undelegate.
     */
    function undelegate(address validator, uint128 shares) external;

    /**
     * @notice Undelegate KGH for given validator and tokenId.
     *
     * @param validator Address of the validator.
     * @param tokenId   Token id of KGH to undelegate.
     */
    function undelegateKgh(address validator, uint256 tokenId) external;

    /**
     * @notice Undelegate KGHs for given validator and token ids.
     *
     * @param validator Address of the validator.
     * @param tokenIds  Array of token ids of KGHs to undelegate.
     */
    function undelegateKghBatch(address validator, uint256[] calldata tokenIds) external;

    /**
     * @notice Claim the reward of the KGH delegator from the given validator vault.
     *
     * @param validator Address of the validator.
     */
    function claimKghReward(address validator) external;
}
