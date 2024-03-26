// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { SafeERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { IERC721 } from "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import { IERC721Receiver } from "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import { Math } from "@openzeppelin/contracts/utils/math/Math.sol";

import { Atan2 } from "../libraries/Atan2.sol";
import { BalancedWeightTree } from "../libraries/BalancedWeightTree.sol";
import { Types } from "../libraries/Types.sol";
import { Uint128Math } from "../libraries/Uint128Math.sol";
import { IKGHManager } from "../universal/IKGHManager.sol";
import { L2OutputOracle } from "./L2OutputOracle.sol";

/**
 * @title AssetManager
 * @notice AssetManager is an abstract contract that handles (un)delegations of KRO and KGH, and
 *         the distribution of rewards to the delegators and the validator.
 */
abstract contract AssetManager is IERC721Receiver {
    using SafeERC20 for IERC20;
    using Math for uint256;
    using Uint128Math for uint128;
    using BalancedWeightTree for BalancedWeightTree.Tree;

    /**
     * @notice Represents the parameters of constructor.
     *
     * @custom:field _l2OutputOracle         Address of the L2OutputOracle contract.
     * @custom:field _assetToken             Address of the KRO token.
     * @custom:field _kgh                    Address of the KGH token.
     * @custom:field _kghManager             Address of the KGHManager contract.
     * @custom:field _securityCouncil        Address of the SecurityCouncil contract.
     * @custom:field _maxOutputFinalizations Max number of finalized outputs.
     * @custom:field _baseReward             Base reward for the validator.
     * @custom:field _slashingRateNumerator  Numerator of the slashing rate.
     * @custom:field _minSlashingAmount      Minimum amount to slash.
     * @custom:field _minRegisterAmount      Minimum amount to register as a validator.
     * @custom:field _minStartAmount         Minimum amount to start submitting outputs.
     * @custom:field _undelegationPeriod     Period that should wait to finalize the undelegation.
     */
    struct ConstructorParams {
        L2OutputOracle _l2OutputOracle;
        IERC20 _assetToken;
        IERC721 _kgh;
        IKGHManager _kghManager;
        address _securityCouncil;
        uint128 _maxOutputFinalizations;
        uint128 _baseReward;
        uint128 _slashingRateNumerator;
        uint128 _minSlashingAmount;
        uint128 _minRegisterAmount;
        uint128 _minStartAmount;
        uint256 _undelegationPeriod;
    }

    /**
     * @notice Represents the asset information of the vault of a validator.
     *
     * @custom:field totalKro                   Total amount of KRO in the vault.
     * @custom:field totalKroShares             Total shares for KRO delegation in the vault.
     * @custom:field totalKgh                   Total number of KGH in the vault.
     * @custom:field totalKroInKgh              Total amount of KRO which KGHs in the vault have.
     * @custom:field totalKghShares             Total shares for KGH delegation in the vault.
     * @custom:field validatorKro               Amount of KRO that the validator self-delegated.
     * @custom:field totalPendingAssets         Total pending KRO for undelegation.
     * @custom:field totalPendingBoostedRewards Total pending boosted rewards in KRO for undelegation.
     * @custom:field totalPendingKroShares      Total pending KRO shares for undelegation.
     * @custom:field totalPendingKghShares      Total pending KGH shares for undelegation.
     */
    struct Asset {
        uint128 totalKro;
        uint128 totalKroShares;
        uint128 totalKgh;
        uint128 totalKroInKgh;
        uint128 totalKghShares;
        uint128 validatorKro;
        uint128 totalPendingAssets;
        uint128 totalPendingBoostedRewards;
        uint128 totalPendingKroShares;
        uint128 totalPendingKghShares;
    }

    /**
     * @notice Constructs the reward information of the vault of a validator.
     *
     * @custom:field boostedReward                Cumulated boosted reward for KGH delegators in the vault.
     * @custom:field validatorRewardKro           Cumulated reward for the validator.
     * @custom:field commissionRate               Commission rate of validator.
     * @custom:field commissionMaxChangeRate      Maximum changeable commission rate at once.
     * @custom:field commissionRateChangedAt      Last timestamp when the commission rate was changed.
     * @custom:field totalPendingValidatorRewards Total pending validator rewards.
     * @custom:field claimRequestTimes            Timestamps of validator reward claim requests.
     * @custom:field pendingValidatorRewards      A mapping of timestamp to pending validator rewards.
     */
    struct Reward {
        uint128 boostedReward;
        uint128 validatorRewardKro;
        uint8 commissionRate;
        uint8 commissionMaxChangeRate;
        uint128 commissionRateChangedAt;
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
     * @custom:field kroShares Amount of KRO shares.
     * @custom:field kghShares Amount of KGH shares.
     */
    struct KghDelegatorShares {
        uint128 kroShares;
        uint128 kghShares;
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
     * @custom:field asset             Asset information of the vault.
     * @custom:field reward            Reward information of the vault.
     * @custom:field isInitiated       Whether the vault is initiated.
     * @custom:field noSubmissionCount Number of counts that the validator did not submit the output in priority round.
     * @custom:field kroDelegators     A mapping of validator address to KRO delegator struct.
     * @custom:field kghDelegators     A mapping of validator address to KGH delegator struct.
     */
    struct Vault {
        Asset asset;
        Reward reward;
        bool isInitiated;
        uint8 noSubmissionCount;
        mapping(address => KroDelegator) kroDelegators;
        mapping(address => KghDelegator) kghDelegators;
    }

    /**
     * @notice The denominator for the commission rate.
     */
    uint128 public constant COMMISSION_RATE_DENOM = 100;

    /**
     * @notice The numerator for the boosted reward.
     */
    uint128 public constant BOOSTED_REWARD_NUMERATOR = 40;

    /**
     * @notice The denominator for the boosted reward.
     */
    uint128 public constant BOOSTED_REWARD_DENOM = 100;

    /**
     * @notice The denominator for the slashing rate.
     */
    uint128 public constant SLASHING_RATE_DENOM = 1000;

    /**
     * @notice Virtual KRO amount per KGH.
     */
    uint128 public constant VKRO_PER_KGH = 100e18;

    /**
     * @notice The numerator of the tax.
     */
    uint128 public constant TAX_NUMERATOR = 20;

    /**
     * @notice The denominator of the tax.
     */
    uint128 public constant TAX_DENOMINATOR = 100;

    /**
     * @notice Decimals offset for the KRO and KGH shares.
     */
    uint128 public constant DECIMAL_OFFSET = 10 ** 6;

    /**
     * @notice Address of the KRO token contract.
     */
    IERC20 public immutable ASSET_TOKEN;

    /**
     * @notice Address of the L2OutputOracle contract. Can be updated via upgrade.
     */
    L2OutputOracle public immutable L2_ORACLE;

    /**
     * @notice Address of the KGH token contract.
     */
    IERC721 public immutable KGH;

    /**
     * @notice Address of the KGHManager contract.
     */
    IKGHManager public immutable KGH_MANAGER;

    /**
     * @notice The address of the SecurityCouncil contract. Can be updated via upgrade.
     */
    address public immutable SECURITY_COUNCIL;

    /**
     * @notice The max number of outputs to be finalized at once when distributing rewards.
     */
    uint128 public immutable MAX_OUTPUT_FINALIZATIONS;

    /**
     * @notice Amount of base reward for the validator.
     */
    uint128 public immutable BASE_REWARD;

    /**
     * @notice The numerator of the slashing rate.
     */
    uint128 public immutable SLASHING_RATE_NUMERATOR;

    /**
     * @notice Minimum amount to slash.
     */
    uint128 public immutable MIN_SLASHING_AMOUNT;

    /**
     * @notice Minimum amount to register as a validator.
     */
    uint128 public immutable MIN_REGISTER_AMOUNT;

    /**
     * @notice Minimum amount to start a validator and add it to the validator tree.
     *         Note that only the started validators can submit outputs.
     */
    uint128 public immutable MIN_START_AMOUNT;

    /**
     * @notice Delay for the finalization of undelegation.
     */
    uint256 public immutable UNDELEGATION_PERIOD;

    /**
     * @notice Weighted tree to store and calculate the probability to be selected as an output submitter.
     */
    BalancedWeightTree.Tree internal _validatorTree;

    /**
     * @notice A mapping of validator address to the vault.
     */
    mapping(address => Vault) internal _vaults;

    /**
     * @notice A mapping of output index challenged successfully to pending challenge rewards.
     */
    mapping(uint256 => uint128) internal _pendingChallengeReward;

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
     * @notice Emitted when pending challenge reward for challenge winner is distributed.
     *
     * @param recipient Address of the reward recipient.
     * @param amount    The amount of challenge reward.
     */
    event ChallengeRewardDistributed(address indexed recipient, uint128 amount);

    /**
     * @notice Emitted when the validator is slashed.
     *
     * @param loser  Address of the challenge loser.
     * @param amount The amount of KRO slashed.
     */
    event Slashed(address indexed loser, uint128 amount);

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
     * @notice Emitted when the output reward is distributed.
     *
     * @param validator       Address of the validator whose vault is rewarded.
     * @param validatorReward The amount of validator reward.
     * @param baseReward      The amount of base reward for KRO delegators.
     * @param boostedReward   The amount of boosted reward for KGH delegators.
     */
    event RewardDistributed(
        address indexed validator,
        uint128 validatorReward,
        uint128 baseReward,
        uint128 boostedReward
    );

    /**
     * @notice Modifier to check if the vault is active.
     */
    modifier checkIsActive(address validator) {
        require(
            msg.sender == validator || _vaults[validator].asset.validatorKro >= MIN_REGISTER_AMOUNT,
            "AssetManager: Vault is inactive"
        );
        _;
    }

    /**
     * @notice Constructs the AssetManager contract.
     *
     * @param _constructorParams The parameters of constructor.
     */
    constructor(ConstructorParams memory _constructorParams) {
        L2_ORACLE = _constructorParams._l2OutputOracle;
        ASSET_TOKEN = _constructorParams._assetToken;
        KGH = _constructorParams._kgh;
        KGH_MANAGER = _constructorParams._kghManager;
        SECURITY_COUNCIL = _constructorParams._securityCouncil;
        MAX_OUTPUT_FINALIZATIONS = _constructorParams._maxOutputFinalizations;
        BASE_REWARD = _constructorParams._baseReward;
        SLASHING_RATE_NUMERATOR = _constructorParams._slashingRateNumerator;
        MIN_SLASHING_AMOUNT = _constructorParams._minSlashingAmount;

        require(
            _constructorParams._minRegisterAmount <= _constructorParams._minStartAmount,
            "AssetManager: min register amount should not exceed min start amount"
        );
        MIN_REGISTER_AMOUNT = _constructorParams._minRegisterAmount;
        MIN_START_AMOUNT = _constructorParams._minStartAmount;

        UNDELEGATION_PERIOD = _constructorParams._undelegationPeriod;
    }

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
    ) external view returns (uint128) {
        uint128 shares = _vaults[validator].kroDelegators[delegator].shares;
        return previewUndelegate(validator, shares);
    }

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
    ) external view returns (uint128) {
        uint128 kroInKgh = KGH_MANAGER.totalKroInKgh(tokenId);

        uint128 kghAssets = previewKghUndelegate(
            validator,
            _vaults[validator].kghDelegators[delegator].shares[tokenId].kghShares
        ) - VKRO_PER_KGH;
        uint128 kroAssets = previewUndelegate(
            validator,
            _vaults[validator].kghDelegators[delegator].shares[tokenId].kroShares
        ) - kroInKgh;
        return kghAssets + kroAssets;
    }

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
    ) external view returns (uint128) {
        return _vaults[validator].kroDelegators[delegator].shares;
    }

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
    ) external view returns (uint128, uint128) {
        return (
            _vaults[validator].kghDelegators[delegator].shares[tokenId].kroShares,
            _vaults[validator].kghDelegators[delegator].shares[tokenId].kghShares
        );
    }

    /**
     * @notice Allows an on-chain or off-chain user to simulate the effects of their KRO delegation at the current block.
     *
     * @param validator Address of the validator.
     * @param assets    The amount of assets to delegate.
     *
     * @return The amount of shares that the Vault would exchange for the amount of assets provided.
     */
    function previewDelegate(
        address validator,
        uint128 assets
    ) public view virtual returns (uint128) {
        return _convertToKroShares(validator, assets);
    }

    /**
     * @notice Allows an on-chain or off-chain user to simulate the effects of their KRO undelegation
     *         at the current block.
     *
     * @param validator The address of the validator.
     * @param shares    The amount of shares to undelegate.
     *
     * @return The amount of assets that the Vault would exchange for the amount of shares provided.
     */
    function previewUndelegate(
        address validator,
        uint128 shares
    ) public view virtual returns (uint128) {
        return _convertToKroAssets(validator, shares);
    }

    /**
     * @notice Allows an on-chain or off-chain user to simulate the effects of their KGH delegation
     *         at the current block given current on-chain conditions.
     *
     * @param validator The address of the validator.
     *
     * @return The amount of shares that the Vault would exchange for the amount of assets provided.
     */
    function previewKghDelegate(address validator) public view virtual returns (uint128) {
        return _convertToKghShares(validator, VKRO_PER_KGH);
    }

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
    ) public view virtual returns (uint128) {
        return
            _convertToKghAssets(
                validator,
                _vaults[validator].kghDelegators[msg.sender].shares[tokenId].kghShares
            );
    }

    /**
     * @notice Returns the total amount of KRO assets held by the vault.
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of KRO assets held by the vault.
     */
    function totalKroAssets(address validator) public view virtual returns (uint128) {
        return _vaults[validator].asset.totalKro;
    }

    /**
     * @notice Returns the total number of KGHs held by the vault.
     *
     * @param validator Address of the validator.
     *
     * @return The total number of KGHs held by the vault.
     */
    function totalKghNum(address validator) public view virtual returns (uint128) {
        return _vaults[validator].asset.totalKgh;
    }

    /**
     * @notice Returns the total amount of KGH assets held by the vault.
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of KGH assets held by the vault.
     */
    function _totalKghAssets(address validator) internal view virtual returns (uint128) {
        return
            _vaults[validator].asset.totalKgh *
            VKRO_PER_KGH +
            _vaults[validator].reward.boostedReward;
    }

    /**
     * @notice Returns the total amount of KRO shares held by the vault.
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of shares held by the validator vault.
     */
    function _totalKroShares(address validator) internal view virtual returns (uint128) {
        return _vaults[validator].asset.totalKroShares;
    }

    /**
     * @notice Returns the total amount of KGH shares held by the vault.
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of shares held by the validator vault.
     */
    function _totalKghShares(address validator) internal view virtual returns (uint128) {
        return _vaults[validator].asset.totalKghShares;
    }

    /**
     * @notice Delegate KRO to the validator and returns the amount of shares that the vault would
     *         exchange.
     *
     * @param validator Address of the validator.
     * @param assets    The amount of KRO to delegate.
     *
     * @return The amount of shares that the Vault would exchange for the amount of assets provided.
     */
    function delegate(
        address validator,
        uint128 assets
    ) external checkIsActive(validator) returns (uint128) {
        require(assets > 0, "AssetManager: cannot delegate 0 asset");
        uint128 shares = _delegate(validator, msg.sender, assets, true);
        emit KroDelegated(validator, msg.sender, assets, shares);
        return shares;
    }

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
    function delegateKgh(
        address validator,
        uint256 tokenId
    ) external checkIsActive(validator) returns (uint128, uint128) {
        uint128 kroInKgh = KGH_MANAGER.totalKroInKgh(tokenId);
        uint128 kroShares = previewDelegate(validator, kroInKgh);
        uint128 kghShares = previewKghDelegate(validator);
        _delegateKgh(validator, tokenId, kroInKgh, kroShares, kghShares);

        emit KghDelegated(validator, msg.sender, tokenId, kroInKgh, kroShares, kghShares);
        return (kroShares, kghShares);
    }

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
    ) external checkIsActive(validator) returns (uint128, uint128) {
        require(tokenIds.length > 0, "AssetManager: cannot delegate 0 KGH");

        uint128 kroShares;
        uint128 kghShares = previewKghDelegate(validator);
        uint128 kroInKghs;

        for (uint256 i = 0; i < tokenIds.length; ) {
            KGH.safeTransferFrom(msg.sender, address(this), tokenIds[i]);

            uint128 kroInKgh = KGH_MANAGER.totalKroInKgh(tokenIds[i]);
            uint128 kroSharesForTokenId = previewDelegate(validator, kroInKgh);

            _vaults[validator].kghDelegators[msg.sender].shares[tokenIds[i]] = KghDelegatorShares({
                kroShares: kroSharesForTokenId,
                kghShares: kghShares
            });

            unchecked {
                kroInKghs += kroInKgh;
                kroShares += kroSharesForTokenId;

                ++i;
            }
        }

        kghShares *= uint128(tokenIds.length);
        _delegateKghBatch(validator, uint128(tokenIds.length), kroInKghs, kroShares, kghShares);

        emit KghBatchDelegated(validator, msg.sender, tokenIds, kroShares, kghShares);
        return (kroShares, kghShares);
    }

    /**
     * @notice Initiate KRO undelegation of given shares for given validator.
     *
     * @param validator Address of the validator.
     * @param shares    The amount of shares to undelegate.
     */
    function initUndelegate(address validator, uint128 shares) external {
        require(
            shares > 0 && shares <= _vaults[validator].kroDelegators[msg.sender].shares,
            "AssetManager: Invalid amount of shares to undelegate"
        );

        uint128 assets = previewUndelegate(validator, shares);
        require(assets > 0, "AssetManager: cannot undelegate 0 asset");

        _initUndelegate(validator, msg.sender, assets, shares);
        emit KroUndelegationInitiated(validator, msg.sender, assets, shares);
    }

    /**
     * @notice Initiate KGH undelegation for given validator and tokenId.
     *
     * @param validator Address of the validator.
     * @param tokenId   Token id of KGH to undelegate.
     */
    function initUndelegateKgh(address validator, uint256 tokenId) external {
        uint128 kroShares = _vaults[validator].kghDelegators[msg.sender].shares[tokenId].kroShares;
        uint128 kghShares = _vaults[validator].kghDelegators[msg.sender].shares[tokenId].kghShares;

        require(kghShares != 0, "AssetManager: No shares for the given tokenId");
        _initUndelegateKgh(validator, msg.sender, tokenId, kroShares, kghShares);

        emit KghUndelegationInitiated(validator, msg.sender, tokenId, kroShares, kghShares);
    }

    /**
     * @notice Initiate KGH undelegation for given validator and token ids.
     *
     * @param validator Address of the validator.
     * @param tokenIds  Array of token ids of KGHs to undelegate.
     */
    function initUndelegateKghBatch(address validator, uint256[] calldata tokenIds) external {
        require(tokenIds.length > 0, "AssetManager: cannot undelegate 0 KGH");

        mapping(uint256 => KghDelegatorShares) storage shares = _vaults[validator]
            .kghDelegators[msg.sender]
            .shares;
        uint128 kroShares;
        uint128 kghShares;
        uint128 kroInKghs;
        uint128 kghAssetsToWithdraw;
        for (uint256 i = 0; i < tokenIds.length; ) {
            uint128 kroSharesForTokenId = shares[tokenIds[i]].kroShares;
            uint128 kghSharesForTokenId = shares[tokenIds[i]].kghShares;

            require(kghSharesForTokenId != 0, "AssetManager: No shares for the given tokenId");

            unchecked {
                kroShares += kroSharesForTokenId;
                kghShares += kghSharesForTokenId;
                kroInKghs += KGH_MANAGER.totalKroInKgh(tokenIds[i]);
                kghAssetsToWithdraw += previewKghUndelegate(validator, tokenIds[i]);

                delete shares[tokenIds[i]];

                ++i;
            }
        }

        _initUndelegateKghBatch(
            validator,
            tokenIds,
            kroInKghs,
            kroShares,
            kghShares,
            kghAssetsToWithdraw
        );

        emit KghBatchUndelegationInitiated(validator, msg.sender, tokenIds, kroShares, kghShares);
    }

    /**
     * @notice Claim the reward of the validator.
     *
     * @param amount The amount of reward to claim.
     */
    function initClaimValidatorReward(uint128 amount) external {
        Reward storage reward = _vaults[msg.sender].reward;

        require(
            amount > 0 && amount <= reward.validatorRewardKro,
            "AssetManager: Invalid reward to claim"
        );

        unchecked {
            reward.validatorRewardKro -= amount;
            reward.pendingValidatorRewards[block.timestamp] += amount;
            reward.totalPendingValidatorRewards += amount;
            reward.claimRequestTimes.push(block.timestamp);
        }

        _updateValidatorTree(msg.sender, _vaults[msg.sender], true);

        emit RewardClaimInitiated(msg.sender, amount);
    }

    /**
     * @notice Finalize all pending KRO undelegation and returns the amount of assets that the vault would
     *         exchange for the pending KRO shares.
     *
     * @param validator Address of the validator.
     *
     * @return The amount of assets that the vault would exchange for the pending KRO shares.
     */
    function finalizeUndelegate(address validator) external returns (uint128) {
        Asset storage asset = _vaults[validator].asset;
        require(asset.totalPendingKroShares > 0, "AssetManager: No pending shares to finalize");

        KroDelegator storage delegator = _vaults[validator].kroDelegators[msg.sender];
        uint256[] memory requestTimes = delegator.undelegateRequestTimes;
        require(requestTimes.length > 0, "AssetManager: no undelegation requests exist");

        uint128 assetsToUndelegate;
        uint128 sharesToUndelegate;
        // Loop only while the undelegation request time does exist.
        for (uint256 i = requestTimes.length - 1; i >= 0 && requestTimes[i] > 0; ) {
            if (requestTimes[i] + UNDELEGATION_PERIOD <= block.timestamp) {
                unchecked {
                    sharesToUndelegate += delegator.pendingKroShares[requestTimes[i]];
                }

                delete delegator.pendingKroShares[requestTimes[i]];
                delete delegator.undelegateRequestTimes[i];
            }

            if (i == 0) {
                break;
            }
            unchecked {
                --i;
            }
        }

        require(sharesToUndelegate > 0, "AssetManager: no pending KRO undelegation to finalize");

        unchecked {
            assetsToUndelegate = sharesToUndelegate.mulDiv(
                asset.totalPendingAssets,
                asset.totalPendingKroShares
            );
            asset.totalPendingAssets -= assetsToUndelegate;
            asset.totalPendingKroShares -= sharesToUndelegate;
        }

        ASSET_TOKEN.safeTransfer(msg.sender, assetsToUndelegate);

        emit KroUndelegationFinalized(validator, msg.sender, assetsToUndelegate);
        return assetsToUndelegate;
    }

    /**
     * @notice Finalize all pending KGH undelegation and returns the amount of assets that the vault would
     *         exchange for the pending KRO and KGH shares.
     *
     * @param validator Address of the validator.
     *
     * @return The amount of assets that the vault would exchange for the pending KRO and KGH shares.
     */
    function finalizeUndelegateKgh(address validator) external returns (uint128) {
        KghDelegator storage kghDelegator = _vaults[validator].kghDelegators[msg.sender];
        uint256[] memory requestTimes = kghDelegator.undelegateRequestTimes;

        require(requestTimes.length > 0, "AssetManager: no undelegation requests exist");

        Asset storage asset = _vaults[validator].asset;
        bool rewardExists = asset.totalPendingKghShares > 0;
        uint128 kroSharesToUndelegate;
        uint128 kghSharesToUndelegate;
        uint128 assetsToUndelegate;

        // Loop only while the undelegation request time does exist.
        uint256 finalizedNum;
        for (uint256 i = requestTimes.length - 1; i >= 0 && requestTimes[i] > 0; ) {
            if (requestTimes[i] + UNDELEGATION_PERIOD <= block.timestamp) {
                // Calculate the shares only if reward exists.
                if (rewardExists) {
                    KghDelegatorShares memory pendingShares = kghDelegator.pendingShares[
                        requestTimes[i]
                    ];
                    unchecked {
                        kroSharesToUndelegate += pendingShares.kroShares;
                        kghSharesToUndelegate += pendingShares.kghShares;
                    }
                }

                uint256[] memory tokenIds = kghDelegator.pendingKghIds[requestTimes[i]];
                for (uint256 tokenIdIndex = 0; tokenIdIndex < tokenIds.length; ) {
                    KGH.safeTransferFrom(address(this), msg.sender, tokenIds[tokenIdIndex]);

                    unchecked {
                        ++tokenIdIndex;
                    }
                }

                delete kghDelegator.pendingShares[requestTimes[i]];
                delete kghDelegator.pendingKghIds[requestTimes[i]];
                delete kghDelegator.undelegateRequestTimes[i];

                unchecked {
                    ++finalizedNum;
                }
            }

            if (i == 0) {
                break;
            }
            unchecked {
                --i;
            }
        }

        require(finalizedNum > 0, "AssetManager: no pending KGH undelegation to finalize");

        if (rewardExists) {
            unchecked {
                uint128 kroAssetsToUndelegate = kroSharesToUndelegate.mulDiv(
                    asset.totalPendingAssets,
                    asset.totalPendingKroShares
                );
                uint128 kghAssetsToUndelegate = kghSharesToUndelegate.mulDiv(
                    asset.totalPendingBoostedRewards,
                    asset.totalPendingKghShares
                );
                asset.totalPendingAssets -= kroAssetsToUndelegate;
                asset.totalPendingKroShares -= kroSharesToUndelegate;
                asset.totalPendingBoostedRewards -= kghAssetsToUndelegate;
                asset.totalPendingKghShares -= kghSharesToUndelegate;
                assetsToUndelegate = kroAssetsToUndelegate + kghAssetsToUndelegate;
            }

            ASSET_TOKEN.safeTransfer(msg.sender, assetsToUndelegate);
        }

        emit KghUndelegationFinalized(validator, msg.sender, assetsToUndelegate);
        return assetsToUndelegate;
    }

    /**
     * @notice Finalize the reward claim of the validator.
     */
    function finalizeClaimValidatorReward() external {
        Reward storage reward = _vaults[msg.sender].reward;
        require(
            reward.totalPendingValidatorRewards > 0,
            "AssetManager: no pending validator rewards to finalize"
        );

        uint256[] memory requestTimes = reward.claimRequestTimes;
        uint128 rewardsToClaim;
        for (uint256 i = requestTimes.length - 1; i >= 0 && requestTimes[i] > 0; ) {
            if (requestTimes[i] + UNDELEGATION_PERIOD <= block.timestamp) {
                unchecked {
                    rewardsToClaim += reward.pendingValidatorRewards[requestTimes[i]];
                }

                delete reward.pendingValidatorRewards[requestTimes[i]];
                delete reward.claimRequestTimes[i];
            }

            if (i == 0) {
                break;
            }
            unchecked {
                --i;
            }
        }

        require(rewardsToClaim > 0, "AssetManager: no pending reward claim to finalize yet");

        // To prevent the underflow of the totalPendingValidatorRewards when the validator is slashed.
        if (reward.totalPendingValidatorRewards < rewardsToClaim) {
            rewardsToClaim = reward.totalPendingValidatorRewards;
        }

        unchecked {
            reward.totalPendingValidatorRewards -= rewardsToClaim;
        }
        ASSET_TOKEN.safeTransfer(msg.sender, rewardsToClaim);

        emit RewardClaimFinalized(msg.sender, rewardsToClaim);
    }

    /**
     * @notice Internal function to add base and boosted reward to the vaults of finalized output submitters.
     *
     * @return Whether the reward distribution is done at least once or not.
     */
    function _distributeReward() internal returns (bool) {
        uint256 outputIndex = L2_ORACLE.latestFinalizedOutputIndex() + 1;
        uint256 latestOutputIndex = L2_ORACLE.latestOutputIndex();

        if (!L2_ORACLE.VALIDATOR_POOL().isTerminated(outputIndex)) {
            return false;
        }

        uint128 finalizedOutputNum = 0;
        Types.CheckpointOutput memory output;

        for (
            ;
            finalizedOutputNum < MAX_OUTPUT_FINALIZATIONS && outputIndex <= latestOutputIndex;

        ) {
            if (L2_ORACLE.isFinalized(outputIndex)) {
                output = L2_ORACLE.getL2Output(outputIndex);
                _increaseBalanceWithReward(output.submitter);

                uint128 challengeReward = _pendingChallengeReward[outputIndex];
                if (challengeReward > 0) {
                    _modifyBalanceWithSlashing(_vaults[output.submitter], challengeReward, false);
                    delete _pendingChallengeReward[outputIndex];

                    emit ChallengeRewardDistributed(output.submitter, challengeReward);
                }

                _updateValidatorTree(output.submitter, _vaults[output.submitter], false);

                unchecked {
                    ++outputIndex;
                    ++finalizedOutputNum;
                }
            } else {
                break;
            }
        }

        if (finalizedOutputNum > 0) {
            L2_ORACLE.setLatestFinalizedOutputIndex(outputIndex - 1);

            return true;
        }

        return false;
    }

    /**
     * @notice Internal function to slash KRO at the vault of the challenge loser.
     *
     * @param loser       Address of the challenge loser.
     * @param outputIndex The index of output challenged.
     */
    function _slash(address loser, uint256 outputIndex) internal {
        Vault storage vault = _vaults[loser];
        uint128 amountToSlash = _modifyBalanceWithSlashing(vault, 0, true);
        _pendingChallengeReward[outputIndex] = amountToSlash;

        _updateValidatorTree(loser, vault, true);

        emit Slashed(loser, amountToSlash);
    }

    /**
     * @notice Add pending assets and shares when undelegating KRO.
     *
     * @param vault  Vault of the validator.
     * @param assets The amount of assets to add as pending asset.
     * @param shares The amount of shares to add as pending share.
     */
    function _addPendingKroShares(Vault storage vault, uint128 assets, uint128 shares) internal {
        vault.asset.totalPendingAssets += assets;
        vault.asset.totalPendingKroShares += shares;
        vault.kroDelegators[msg.sender].pendingKroShares[block.timestamp] += shares;
        vault.kroDelegators[msg.sender].undelegateRequestTimes.push(block.timestamp);
    }

    /**
     * @notice Add pending assets and shares when undelegating KGH.
     *
     * @param vault          Vault of the validator.
     * @param tokenIds       Array of token ids of the KGH.
     * @param baseRewards    The amount of base rewards to add as pending asset.
     * @param boostedRewards The amount of boosted rewards to add as pending asset.
     * @param shares         The amount of KRO and KGH shares to add as pending share.
     */
    function _addPendingKghShares(
        Vault storage vault,
        uint256[] memory tokenIds,
        uint128 baseRewards,
        uint128 boostedRewards,
        KghDelegatorShares memory shares
    ) internal {
        for (uint256 i = 0; i < tokenIds.length; i++) {
            vault.kghDelegators[msg.sender].pendingKghIds[block.timestamp].push(tokenIds[i]);
        }

        vault.asset.totalPendingAssets += baseRewards;
        vault.asset.totalPendingKroShares += shares.kroShares;
        vault.asset.totalPendingBoostedRewards += boostedRewards;
        vault.asset.totalPendingKghShares += shares.kghShares;

        vault.kghDelegators[msg.sender].pendingShares[block.timestamp].kroShares += shares
            .kroShares;
        vault.kghDelegators[msg.sender].pendingShares[block.timestamp].kghShares += shares
            .kghShares;
        vault.kghDelegators[msg.sender].undelegateRequestTimes.push(block.timestamp);
    }

    /**
     * @notice Internal conversion function for KRO (from assets to shares).
     *
     * @param validator Address of the validator.
     * @param assets    The amount of assets to convert to shares.
     */
    function _convertToKroShares(
        address validator,
        uint128 assets
    ) internal view virtual returns (uint128) {
        return
            assets.mulDiv(
                _totalKroShares(validator) + DECIMAL_OFFSET,
                totalKroAssets(validator) + 1
            );
    }

    /**
     * @notice Internal conversion function for KRO (from shares to assets).
     *
     * @param validator Address of the validator.
     * @param shares    The amount of shares to convert to assets.
     */
    function _convertToKroAssets(
        address validator,
        uint128 shares
    ) internal view virtual returns (uint128) {
        return
            shares.mulDiv(
                totalKroAssets(validator) + 1,
                _totalKroShares(validator) + DECIMAL_OFFSET
            );
    }

    /**
     * @notice Internal conversion function for KGH (from assets to shares).
     *
     * @param validator Address of the validator.
     * @param assets    The amount of assets to convert to shares.
     */
    function _convertToKghShares(
        address validator,
        uint128 assets
    ) internal view virtual returns (uint128) {
        return
            assets.mulDiv(
                _totalKghShares(validator) + DECIMAL_OFFSET,
                _totalKghAssets(validator) + 1
            );
    }

    /**
     * @notice Internal conversion function for KGH (from shares to assets).
     *
     * @param validator Address of the validator.
     * @param shares    The amount of shares to convert to assets.
     */
    function _convertToKghAssets(
        address validator,
        uint128 shares
    ) internal view virtual returns (uint128) {
        return
            shares.mulDiv(
                _totalKghAssets(validator) + 1,
                _totalKghShares(validator) + DECIMAL_OFFSET
            );
    }

    /**
     * @notice Internal function to delegate KRO to the validator.
     *
     * @param validator  Address of the validator.
     * @param owner      Address of the delegator.
     * @param assets     The amount of KRO to delegate.
     * @param updateTree Flag to update validator tree or not.
     */
    function _delegate(
        address validator,
        address owner,
        uint128 assets,
        bool updateTree
    ) internal virtual returns (uint128) {
        uint128 shares = previewDelegate(validator, assets);
        Vault storage vault = _vaults[validator];

        ASSET_TOKEN.safeTransferFrom(owner, address(this), assets);

        unchecked {
            vault.asset.totalKro += assets;
            vault.asset.totalKroShares += shares;
            vault.kroDelegators[owner].shares += shares;

            if (owner == validator) {
                vault.asset.validatorKro += assets;
            }
        }

        if (updateTree) {
            _updateValidatorTree(validator, vault, false);
        }

        return shares;
    }

    /**
     * @notice Internal function to delegate KGH to the validator.
     *
     * @param validator Address of the validator.
     * @param tokenId   Token Id of the KGH.
     * @param kroInKgh  The amount of KRO in the KGH.
     * @param kroShares The amount of KRO shares to receive.
     * @param kghShares The amount of KGH shares to receive.
     */
    function _delegateKgh(
        address validator,
        uint256 tokenId,
        uint128 kroInKgh,
        uint128 kroShares,
        uint128 kghShares
    ) internal virtual {
        Vault storage vault = _vaults[validator];
        Asset storage asset = vault.asset;

        KGH.safeTransferFrom(msg.sender, address(this), tokenId);

        unchecked {
            asset.totalKro += kroInKgh;
            asset.totalKroInKgh += kroInKgh;
            asset.totalKgh += 1;

            asset.totalKroShares += kroShares;
            asset.totalKghShares += kghShares;
            vault.kghDelegators[msg.sender].shares[tokenId] = KghDelegatorShares(
                kroShares,
                kghShares
            );
        }

        if (kroInKgh > 0) {
            _updateValidatorTree(validator, vault, false);
        }
    }

    /**
     * @notice Internal function to delegate KGHs to the validator.
     *
     * @param validator Address of the validator.
     * @param kghCount  The number of KGHs to delegate.
     * @param kroInKghs The amount of KRO in the KGHs.
     * @param kroShares The amount of KRO shares to receive.
     * @param kghShares The amount of KGH shares to receive.
     */
    function _delegateKghBatch(
        address validator,
        uint128 kghCount,
        uint128 kroInKghs,
        uint128 kroShares,
        uint128 kghShares
    ) internal virtual {
        Asset storage asset = _vaults[validator].asset;

        unchecked {
            asset.totalKro += kroInKghs;
            asset.totalKroInKgh += kroInKghs;
            asset.totalKgh += kghCount;

            asset.totalKroShares += kroShares;
            asset.totalKghShares += kghShares;
        }

        if (kroInKghs > 0) {
            _updateValidatorTree(validator, _vaults[validator], false);
        }
    }

    /**
     * @notice Internal function to undelegate KRO from the validator.
     *
     * @param validator Address of the validator.
     * @param owner     Address of the delegator.
     * @param assets    The amount of KRO to undelegate.
     * @param shares    The amount of shares to undelegate.
     */
    function _initUndelegate(
        address validator,
        address owner,
        uint128 assets,
        uint128 shares
    ) internal virtual {
        Vault storage vault = _vaults[validator];

        unchecked {
            vault.asset.totalKroShares -= shares;
            vault.kroDelegators[owner].shares -= shares;

            vault.asset.totalKro -= assets;
            if (owner == validator) {
                vault.asset.validatorKro -= assets;
            }
            _addPendingKroShares(vault, assets, shares);
        }

        _updateValidatorTree(validator, vault, true);
    }

    /**
     * @notice Internal function to undelegate KGH from the validator.
     *
     * @param validator Address of the validator.
     * @param owner     Address of the delegator.
     * @param tokenId   Token Id of the KGH.
     * @param kroShares The amount of KRO shares to undelegate.
     * @param kghShares The amount of KGH shares to undelegate.
     */
    function _initUndelegateKgh(
        address validator,
        address owner,
        uint256 tokenId,
        uint128 kroShares,
        uint128 kghShares
    ) internal virtual {
        Vault storage vault = _vaults[validator];
        uint128 kroAssetsToWithdraw = previewUndelegate(validator, kroShares);
        uint128 kghAssetsToWithdraw = previewKghUndelegate(validator, tokenId);
        uint128 boostedRewardsToReceive;

        uint256[] memory tokenIds = new uint256[](1);
        tokenIds[0] = tokenId;

        unchecked {
            vault.asset.totalKroShares -= kroShares;
            vault.asset.totalKghShares -= kghShares;
            delete vault.kghDelegators[owner].shares[tokenId];

            uint128 kroInKgh = KGH_MANAGER.totalKroInKgh(tokenId);
            uint128 baseRewardsToReceive = kroAssetsToWithdraw - kroInKgh;
            boostedRewardsToReceive = kghAssetsToWithdraw - VKRO_PER_KGH;

            vault.asset.totalKro -= kroAssetsToWithdraw;
            vault.asset.totalKroInKgh -= kroInKgh;
            vault.asset.totalKgh -= 1;
            vault.reward.boostedReward -= boostedRewardsToReceive;

            KghDelegatorShares memory pendingShares = KghDelegatorShares({
                kroShares: kroShares.mulDiv(baseRewardsToReceive, kroAssetsToWithdraw),
                kghShares: kghShares.mulDiv(boostedRewardsToReceive, kghAssetsToWithdraw)
            });

            _addPendingKghShares(
                vault,
                tokenIds,
                baseRewardsToReceive,
                boostedRewardsToReceive,
                pendingShares
            );
        }

        if (kroAssetsToWithdraw + boostedRewardsToReceive > 0) {
            _updateValidatorTree(validator, vault, true);
        }
    }

    /**
     * @notice Internal function to undelegate KGHs from the validator.
     *
     * @param validator           Address of the validator.
     * @param tokenIds            Array of token ids of KGHs to undelegate.
     * @param kroInKghs           The amount of KRO in the KGHs.
     * @param kroShares           The amount of KRO shares to undelegate.
     * @param kghShares           The amount of KGH shares to undelegate.
     * @param kghAssetsToWithdraw The amount of KGH assets to withdraw.
     */
    function _initUndelegateKghBatch(
        address validator,
        uint256[] calldata tokenIds,
        uint128 kroInKghs,
        uint128 kroShares,
        uint128 kghShares,
        uint128 kghAssetsToWithdraw
    ) internal virtual {
        Vault storage vault = _vaults[validator];
        uint128 kroAssetsToWithdraw = previewUndelegate(validator, kroShares);
        uint128 boostedRewardsToReceive;

        unchecked {
            vault.asset.totalKroShares -= kroShares;
            vault.asset.totalKghShares -= kghShares;

            uint128 baseRewardsToReceive = kroAssetsToWithdraw - kroInKghs;
            boostedRewardsToReceive = kghAssetsToWithdraw - VKRO_PER_KGH * uint128(tokenIds.length);

            vault.asset.totalKro -= kroAssetsToWithdraw;
            vault.asset.totalKroInKgh -= kroInKghs;
            vault.asset.totalKgh -= uint128(tokenIds.length);
            vault.reward.boostedReward -= boostedRewardsToReceive;

            KghDelegatorShares memory pendingShares = KghDelegatorShares({
                kroShares: kroShares.mulDiv(baseRewardsToReceive, kroAssetsToWithdraw),
                kghShares: kghShares.mulDiv(boostedRewardsToReceive, kghAssetsToWithdraw)
            });

            _addPendingKghShares(
                vault,
                tokenIds,
                baseRewardsToReceive,
                boostedRewardsToReceive,
                pendingShares
            );
        }

        if (kroAssetsToWithdraw + boostedRewardsToReceive > 0) {
            _updateValidatorTree(validator, vault, true);
        }
    }

    /**
     * @notice Internal function to distribute the reward and update the weight of the validator.
     *
     * @param validator Address of the validator.
     */
    function _increaseBalanceWithReward(address validator) internal {
        Vault storage vault = _vaults[validator];
        uint128 commissionRate = uint128(vault.reward.commissionRate);
        uint128 boostedReward = _getBoostedReward(vault.asset.totalKgh);
        uint128 baseReward;
        uint128 validatorReward;

        unchecked {
            validatorReward = (BASE_REWARD + boostedReward).mulDiv(
                commissionRate,
                COMMISSION_RATE_DENOM
            );
            baseReward = BASE_REWARD.mulDiv(
                COMMISSION_RATE_DENOM - commissionRate,
                COMMISSION_RATE_DENOM
            );
            boostedReward = boostedReward.mulDiv(
                COMMISSION_RATE_DENOM - commissionRate,
                COMMISSION_RATE_DENOM
            );

            vault.asset.totalKro += baseReward;
            vault.reward.boostedReward += boostedReward;
            vault.reward.validatorRewardKro += validatorReward;
        }

        // TODO - Distribute the reward from a designated vault to the ValidatorManager contract.

        emit RewardDistributed(validator, validatorReward, baseReward, boostedReward);
    }

    /**
     * @notice Internal function to modify the balance of the vault with slashing.
     *
     * @param vault              Vault of the validator.
     * @param amountToSlashOrAdd The amount to slash or add.
     * @param isLoser            True if the validator is the loser at the challenge.
     *
     * @return The amount to slash or add.
     */
    function _modifyBalanceWithSlashing(
        Vault storage vault,
        uint128 amountToSlashOrAdd,
        bool isLoser
    ) internal returns (uint128) {
        uint128 totalAmount = vault.asset.totalKro +
            vault.asset.totalPendingAssets +
            vault.asset.totalPendingBoostedRewards +
            vault.reward.validatorRewardKro +
            vault.reward.totalPendingValidatorRewards +
            vault.reward.boostedReward -
            vault.asset.totalKroInKgh;

        uint128[6] memory arr = [
            (vault.asset.totalKro - vault.asset.totalKroInKgh),
            vault.reward.boostedReward,
            vault.asset.totalPendingAssets,
            vault.asset.totalPendingBoostedRewards,
            vault.reward.validatorRewardKro,
            vault.reward.totalPendingValidatorRewards
        ];

        if (isLoser) {
            amountToSlashOrAdd = totalAmount.mulDiv(SLASHING_RATE_NUMERATOR, SLASHING_RATE_DENOM);
            if (amountToSlashOrAdd < MIN_SLASHING_AMOUNT) {
                amountToSlashOrAdd = MIN_SLASHING_AMOUNT;
            }

            unchecked {
                vault.asset.totalKro -= arr[0].mulDiv(amountToSlashOrAdd, totalAmount);
                vault.reward.boostedReward -= arr[1].mulDiv(amountToSlashOrAdd, totalAmount);
                vault.asset.totalPendingAssets -= arr[2].mulDiv(amountToSlashOrAdd, totalAmount);
                vault.asset.totalPendingBoostedRewards -= arr[3].mulDiv(
                    amountToSlashOrAdd,
                    totalAmount
                );
                vault.reward.validatorRewardKro -= arr[4].mulDiv(amountToSlashOrAdd, totalAmount);
                vault.reward.totalPendingValidatorRewards -= arr[5].mulDiv(
                    amountToSlashOrAdd,
                    totalAmount
                );
            }

            uint128 tax = amountToSlashOrAdd.mulDiv(TAX_NUMERATOR, TAX_DENOMINATOR);
            unchecked {
                amountToSlashOrAdd -= tax;
            }

            ASSET_TOKEN.safeTransfer(SECURITY_COUNCIL, tax);

            return amountToSlashOrAdd;
        } else {
            unchecked {
                vault.asset.totalKro += arr[0].mulDiv(amountToSlashOrAdd, totalAmount);
                vault.reward.boostedReward += arr[1].mulDiv(amountToSlashOrAdd, totalAmount);
                vault.asset.totalPendingAssets += arr[2].mulDiv(amountToSlashOrAdd, totalAmount);
                vault.asset.totalPendingBoostedRewards += arr[3].mulDiv(
                    amountToSlashOrAdd,
                    totalAmount
                );
                vault.reward.validatorRewardKro += arr[4].mulDiv(amountToSlashOrAdd, totalAmount);
                vault.reward.totalPendingValidatorRewards += arr[5].mulDiv(
                    amountToSlashOrAdd,
                    totalAmount
                );
            }

            return amountToSlashOrAdd;
        }
    }

    /**
     * @notice Internal function to get the boosted reward with the number of KGH.
     *
     * @param numKgh The number of KGH.
     *
     * @return The boosted reward with the number of KGH.
     */
    function _getBoostedReward(uint128 numKgh) internal view virtual returns (uint128) {
        uint128 coefficient = BASE_REWARD.mulDiv(BOOSTED_REWARD_NUMERATOR, BOOSTED_REWARD_DENOM);
        return uint128(Atan2.atan2(numKgh, 100).mulDiv(coefficient, 1 << 40));
    }

    /**
     * @notice Internal function to update the weight tree of the validator.
     *
     * @param validator Address of the validator.
     * @param vault     Vault of the validator.
     * @param tryRemove Flag to try remove the validator from weight tree.
     */
    function _updateValidatorTree(address validator, Vault storage vault, bool tryRemove) internal {
        uint128 newWeight = _reflectiveWeight(vault);
        if (tryRemove && newWeight - vault.asset.totalKroInKgh < MIN_START_AMOUNT) {
            _validatorTree.remove(validator);
        } else {
            _validatorTree.update(validator, uint120(newWeight));
        }
    }

    /**
     * @notice Returns the reflective weight of given vault. It can be different from the actual
     *         current weight of the validator in weight tree since it includes all accumulated
     *         rewards.
     *
     * @param vault Vault of the validator.
     *
     * @return The reflective weight of given vault.
     */
    function _reflectiveWeight(Vault storage vault) internal view returns (uint128) {
        return vault.asset.totalKro + vault.reward.boostedReward + vault.reward.validatorRewardKro;
    }

    /**
     * @inheritdoc IERC721Receiver
     */
    function onERC721Received(
        address /* operator */,
        address /* from */,
        uint256 /* tokenId */,
        bytes calldata /* data */
    ) external pure returns (bytes4) {
        return IERC721Receiver.onERC721Received.selector;
    }
}
