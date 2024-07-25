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
     * @custom:field validatorKro       Total amount of KRO that deposited by the validator and
     *                                  accumulated as validator reward (including validatorKroBonded).
     * @custom:field validatorKroBonded Total amount of validator KRO that bonded during output
     *                                  submission or challenge creation.
     * @custom:field totalKro           Total amount of KRO that delegated by the delegators and
     *                                  accumulated as KRO delegation reward.
     * @custom:field totalKroShares     Total shares for KRO delegation in the vault.
     * @custom:field totalKgh           Total number of KGH in the vault.
     * @custom:field rewardPerKghStored Accumulated boosted reward per 1 KGH.
     */
    struct Asset {
        uint128 validatorKro;
        uint128 validatorKroBonded;
        uint128 totalKro;
        uint128 totalKroShares;
        uint128 totalKgh;
        uint128 rewardPerKghStored;
    }

    /**
     * @notice Constructs the delegator of KRO in the vault of a validator.
     *
     * @custom:field shares          Amount of shares for KRO delegation.
     * @custom:field lastDelegatedAt Last timestamp when the delegator delegated. The delegator can
     *                               undelegate after MIN_DELEGATION_PERIOD elapsed.
     */
    struct KroDelegator {
        uint128 shares;
        uint128 lastDelegatedAt;
    }

    /**
     * @notice Constructs the delegator of KGH in the vault of a validator.
     *
     * @custom:field rewardPerKghPaid Accumulated paid boosted reward per 1 KGH.
     * @custom:field kghNum           Total number of KGH delegated.
     * @custom:field delegatedAt      A mapping of tokenId to the delegation timestamp. The
     *                                delegator can undelegate after MIN_DELEGATION_PERIOD
     *                                elapsed from each delegation timestamp.
     */
    struct KghDelegator {
        uint128 rewardPerKghPaid;
        uint128 kghNum;
        mapping(uint256 => uint128) delegatedAt;
    }

    /**
     * @notice Constructs the vault of a validator.
     *
     * @custom:field withdrawAccount An account where assets can be withdrawn to. Only this account
     *                               can withdraw the assets.
     * @custom:field lastDepositedAt Last timestamp when the validator deposited. The validator can
     *                               withdraw after MIN_DELEGATION_PERIOD elapsed.
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
     */
    event KghDelegated(address indexed validator, address indexed delegator, uint256 tokenId);

    /**
     * @notice Emitted when KGHs are delegated in batch.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param tokenIds  Array of token ids of the KGHs.
     */
    event KghBatchDelegated(
        address indexed validator,
        address indexed delegator,
        uint256[] tokenIds
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
     * @param amount    The amount of KRO claimed as boosted reward.
     */
    event KghUndelegated(
        address indexed validator,
        address indexed delegator,
        uint256 tokenId,
        uint128 amount
    );

    /**
     * @notice Emitted when KGHs are undelegated in batch.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param tokenIds  Array of token ids of the KGHs.
     * @param amount    The amount of KRO claimed as boosted reward.
     */
    event KghBatchUndelegated(
        address indexed validator,
        address indexed delegator,
        uint256[] tokenIds,
        uint128 amount
    );

    /**
     * @notice Emitted when accumulated rewards of KGH delegation are claimed.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param amount    The amount of KRO claimed as boosted reward.
     */
    event KghRewardClaimed(address indexed validator, address indexed delegator, uint128 amount);

    /**
     * @notice Emitted when validator KRO is bonded during output submission or challenge creation.
     *
     * @param validator Address of the validator.
     * @param amount    The amount of KRO bonded.
     * @param remainder The remaining amount of validator KRO excluding bonded KRO.
     */
    event ValidatorKroBonded(address indexed validator, uint128 amount, uint128 remainder);

    /**
     * @notice Emitted when validator KRO is unbonded during output finalization or slashing.
     *
     * @param validator Address of the validator.
     * @param amount    The amount of KRO unbonded.
     * @param remainder The remaining amount of validator KRO excluding bonded KRO.
     */
    event ValidatorKroUnbonded(address indexed validator, uint128 amount, uint128 remainder);

    /**
     * @notice Reverts when caller is not allowed.
     */
    error NotAllowedCaller();

    /**
     * @notice Reverts when the status of validator is improper.
     */
    error ImproperValidatorStatus();

    /**
     * @notice Reverts when try to input zero.
     */
    error NotAllowedZeroInput();

    /**
     * @notice Reverts when the address is zero address.
     */
    error ZeroAddress();

    /**
     * @notice Reverts when the asset is insufficient.
     */
    error InsufficientAsset();

    /**
     * @notice Reverts when the share is insufficient.
     */
    error InsufficientShare();

    /**
     * @notice Reverts when the minimum delegation period is not elapsed.
     */
    error NotElapsedMinDelegationPeriod();

    /**
     * @notice Reverts when the given token ids are invalid.
     */
    error InvalidTokenIdsInput();

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
     *         MIN_DELEGATION_PERIOD elapsed from lastDepositedAt.
     *
     * @param validator Address of the validator.
     *
     * @return When the validator can withdraw KRO.
     */
    function canWithdrawAt(address validator) external view returns (uint128);

    /**
     * @notice Returns the total amount of KRO a validator has deposited and been rewarded.
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of KRO a validator has deposited and been rewarded.
     */
    function totalValidatorKro(address validator) external view returns (uint128);

    /**
     * @notice Returns the total amount of validator KRO that bonded during output submission or
     *         challenge creation.
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of validator KRO bonded.
     */
    function totalValidatorKroBonded(address validator) external view returns (uint128);

    /**
     * @notice Returns the total amount of validator balance excluding the bond amount.
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of validator balance excluding the bond amount.
     */
    function totalValidatorKroNotBonded(address validator) external view returns (uint128);

    /**
     * @notice Returns the total amount of KRO that delegated by the delegators and accumulated as
     *         KRO delegation reward.
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of KRO that delegated by the delegators and accumulated as KRO
     *         delegation reward.
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
     * @notice Returns the amount of KRO assets delegated to the given validator by the delegator.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     *
     * @return The amount of KRO assets that the delegator delegated to the validator.
     */
    function getKroAssets(address validator, address delegator) external view returns (uint128);

    /**
     * @notice Returns when the KRO delegators can undelegate KRO. The delegators can undelegate
     *         after MIN_DELEGATION_PERIOD elapsed from lastDelegatedAt.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the KRO delegator.
     *
     * @return When the KRO delegators can undelegate KRO.
     */
    function canUndelegateKroAt(
        address validator,
        address delegator
    ) external view returns (uint128);

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
     * @notice Returns when the KGH delegators can undelegate KGH. The delegators can undelegate KGH
     *         for the given token id after MIN_DELEGATION_PERIOD elapsed from delegation
     *         timestamp.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the KGH delegator.
     * @param tokenId   The token id of KGH to undelegate.
     *
     * @return When the KGH delegators can undelegate KGH for the given token id.
     */
    function canUndelegateKghAt(
        address validator,
        address delegator,
        uint256 tokenId
    ) external view returns (uint128);

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
     * @notice Withdraw KRO. To withdraw KRO, the validator should be initiated and MIN_DELEGATION_PERIOD
     *         should be passed after the last deposit time. Only withdrawAccount of the validator can call
     *         this function.
     *
     * @param validator Address of the validator.
     * @param assets    The amount of KRO to withdraw.
     */
    function withdraw(address validator, uint128 assets) external;

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
     * @notice Delegate KGH to the validator.
     *
     * @param validator Address of the validator.
     * @param tokenId   The token id of KGH to delegate.
     */
    function delegateKgh(address validator, uint256 tokenId) external;

    /**
     * @notice Delegate KGHs to the validator.
     *
     * @param validator Address of the validator.
     * @param tokenIds  The token ids of KGHs to delegate.
     */
    function delegateKghBatch(address validator, uint256[] calldata tokenIds) external;

    /**
     * @notice Undelegate the KRO of given assets for the given validator.
     *
     * @param validator Address of the validator.
     * @param assets    The amount of assets to undelegate.
     */
    function undelegate(address validator, uint128 assets) external;

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
     * @notice Claim the boosted reward of the KGH delegator from the given validator vault.
     *
     * @param validator Address of the validator.
     */
    function claimKghReward(address validator) external;
}
