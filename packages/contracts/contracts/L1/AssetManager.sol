// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { IERC721 } from "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import { IERC721Receiver } from "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import { Math } from "@openzeppelin/contracts/utils/math/Math.sol";
import { SafeERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import { Atan2 } from "../libraries/Atan2.sol";
import { BalancedWeightTree } from "../libraries/BalancedWeightTree.sol";
import { Types } from "../libraries/Types.sol";
import { Uint128Math } from "../libraries/Uint128Math.sol";
import { IKGHManager } from "../universal/IKGHManager.sol";
import { Colosseum } from "./Colosseum.sol";
import { L2OutputOracle } from "./L2OutputOracle.sol";

/**
 * @title AssetManager
 * @notice AssetManager is an abstract contract that handles (un)delegations of KRO and KGH, and
 *         the distribution of rewards to the delegators and the validator.
 */
abstract contract AssetManager is IERC721Receiver {
    using Math for uint256;
    using Uint128Math for uint128;
    using BalancedWeightTree for BalancedWeightTree.Tree;

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
     * @notice Represents the amounts of KRO and KGH shares that requested for undelegation when undelegating KGH.
     *
     * @custom:field kroShares Amount of KRO shares to undelegate.
     * @custom:field kghShares Amount of KGH shares to undelegate.
     */
    struct UndelegateShares {
        uint128 kroShares;
        uint128 kghShares;
    }

    /**
     * @notice Constructs the delegator of KGH in the vault of a validator.
     *
     * @custom:field kroShares              Amount of shares for KRO in KGH with a corresponding tokenId.
     * @custom:field kghShares              Amount of shares for delegation of KGH with a corresponding tokenId.
     * @custom:field undelegateRequestTimes Timestamps of undelegation requests.
     * @custom:field pendingKghIds          A mapping of timestamp undelegations are requested to
     *                                      token ids of KGHs for undelegation.
     * @custom:field pendingShares          A mapping of timestamp undelegations are requested to
     *                                      pending KRO and KGH shares for undelegation.
     */
    struct KghDelegator {
        mapping(uint256 => uint128) kroShares;
        mapping(uint256 => uint128) kghShares;
        uint256[] undelegateRequestTimes;
        mapping(uint256 => uint256[]) pendingKghIds;
        mapping(uint256 => UndelegateShares) pendingShares;
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
     * @notice The numerator for the boosted reward.
     */
    uint128 public constant BOOSTED_REWARD_NUMERATOR = 40;

    /**
     * @notice The denominator for the boosted reward.
     */
    uint128 public constant BOOSTED_REWARD_DENOM = 100;

    /**
     * @notice The denominator for per-mille used at slashing and initialization of undelegation.
     */
    uint128 public constant PER_MILLE_DENOM = 1000;

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
     * @notice The max number of outputs to be finalized at once.
     */
    uint128 public immutable MAX_OUTPUT_FINALIZATIONS;

    /**
     * @notice Amount of base reward for the validator.
     */
    uint128 public immutable BASE_REWARD;

    /**
     * @notice The numerator of the slashing rate. It will be divided by PER_MILLE_DENOM.
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
     * @notice Minimum amount to start submitting outputs.
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
     * @notice The output index to distribute reward next.
     */
    uint256 internal _nextRewardOutputIndex;

    /**
     * @notice A mapping of validator address to the vault.
     */
    mapping(address => Vault) internal _vaults;

    /**
     * @notice A mapping of the submitted output index to finalization timestamp.
     */
    mapping(uint256 => uint256) internal _submittedOutputs;

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
     * @param loser  Address of the loser at the challenge.
     * @param winner Address of the winner at the challenge.
     * @param amount The amount of KRO slashed.
     */
    event Slashed(address indexed loser, address indexed winner, uint128 amount);

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
     * @param recipient Address of the reward recipient.
     * @param amount    The amount of reward.
     */
    event RewardDistributed(address indexed recipient, uint128 amount);

    /**
     * @notice Modifier to check if the caller is the Colosseum contract.
     */
    modifier onlyColosseum() {
        require(
            msg.sender == L2_ORACLE.COLOSSEUM(),
            "AssetManager: Only Colosseum can call this function"
        );
        _;
    }

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
     * @param _l2OutputOracle         Address of the L2OutputOracle contract.
     * @param _assetToken             Address of the KRO token.
     * @param _kgh                    Address of the KGH token.
     * @param _kghManager             Address of the KGHManager contract.
     * @param _securityCouncil        Address of the SecurityCouncil contract.
     * @param _maxOutputFinalizations Max number of finalized outputs.
     * @param _baseReward             Base reward for the validator.
     * @param _slashingRateNumerator  Numerator of the slashing rate.
     * @param _minSlashingAmount      Minimum amount to slash.
     * @param _minRegisterAmount      Minimum amount to register as a validator.
     * @param _minStartAmount         Minimum amount to start submitting outputs.
     * @param _undelegationPeriod     Period that should wait to finalize the undelegation.
     */
    constructor(
        L2OutputOracle _l2OutputOracle,
        IERC20 _assetToken,
        IERC721 _kgh,
        IKGHManager _kghManager,
        address _securityCouncil,
        uint128 _maxOutputFinalizations,
        uint128 _baseReward,
        uint128 _slashingRateNumerator,
        uint128 _minSlashingAmount,
        uint128 _minRegisterAmount,
        uint128 _minStartAmount,
        uint256 _undelegationPeriod
    ) {
        L2_ORACLE = _l2OutputOracle;
        ASSET_TOKEN = _assetToken;
        KGH = _kgh;
        KGH_MANAGER = _kghManager;
        SECURITY_COUNCIL = _securityCouncil;
        MAX_OUTPUT_FINALIZATIONS = _maxOutputFinalizations;
        BASE_REWARD = _baseReward;
        SLASHING_RATE_NUMERATOR = _slashingRateNumerator;
        MIN_SLASHING_AMOUNT = _minSlashingAmount;
        MIN_REGISTER_AMOUNT = _minRegisterAmount;
        MIN_START_AMOUNT = _minStartAmount;
        UNDELEGATION_PERIOD = _undelegationPeriod;
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
        KghDelegator storage kghDelegator = _vaults[validator].kghDelegators[delegator];
        uint128 kghShares = kghDelegator.kghShares[tokenId];
        uint128 kroShares = kghDelegator.kroShares[tokenId];
        uint128 kroInKgh = KGH_MANAGER.totalKroInKgh(tokenId);

        uint128 kghAssets = previewKghUndelegate(validator, kghShares) - VKRO_PER_KGH;
        uint128 kroAssets = previewUndelegate(validator, kroShares) - kroInKgh;
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
        KghDelegator storage kghDelegator = _vaults[validator].kghDelegators[delegator];
        return (kghDelegator.kroShares[tokenId], kghDelegator.kghShares[tokenId]);
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
                _vaults[validator].kghDelegators[msg.sender].kghShares[tokenId]
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
        Vault storage vault = _vaults[validator];
        return vault.asset.totalKgh * VKRO_PER_KGH + vault.reward.boostedReward;
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
        uint128 shares = previewDelegate(validator, assets);
        _delegate(validator, msg.sender, assets, shares);
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
        uint128 kroShares;
        uint128 kghShares = previewKghDelegate(validator);
        uint128 kroInKghs;

        for (uint256 i = 0; i < tokenIds.length; ) {
            KGH.safeTransferFrom(msg.sender, address(this), tokenIds[i]);

            uint128 kroInKgh = KGH_MANAGER.totalKroInKgh(tokenIds[i]);
            uint128 kroSharesForTokenId = previewDelegate(validator, kroInKgh);

            unchecked {
                kroInKghs += kroInKgh;
                kroShares += kroSharesForTokenId;

                _vaults[validator].kghDelegators[msg.sender].kroShares[
                    tokenIds[i]
                ] = kroSharesForTokenId;
                _vaults[validator].kghDelegators[msg.sender].kghShares[tokenIds[i]] = kghShares;

                ++i;
            }
        }

        kghShares *= uint128(tokenIds.length);
        _delegateKghBatch(validator, uint128(tokenIds.length), kroInKghs, kroShares, kghShares);

        emit KghBatchDelegated(validator, msg.sender, tokenIds, kroShares, kghShares);
        return (kroShares, kghShares);
    }

    /**
     * @notice Initiate KRO undelegation for given validator and per-mille.
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
        uint128 kroShares = _vaults[validator].kghDelegators[msg.sender].kroShares[tokenId];
        uint128 kghShares = _vaults[validator].kghDelegators[msg.sender].kghShares[tokenId];

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
        KghDelegator storage kghDelegator = _vaults[validator].kghDelegators[msg.sender];
        uint128 kroShares;
        uint128 kghShares;
        uint128 kroInKghs;
        uint128 kghAssetsToWithdraw;
        for (uint256 i = 0; i < tokenIds.length; ) {
            uint128 kroSharesForTokenId = kghDelegator.kroShares[tokenIds[i]];
            uint128 kghShareForTokenId = kghDelegator.kghShares[tokenIds[i]];

            require(kghShareForTokenId != 0, "AssetManager: No shares for the given tokenId");

            unchecked {
                kroShares += kroSharesForTokenId;
                kghShares += kghShareForTokenId;
                kroInKghs += KGH_MANAGER.totalKroInKgh(tokenIds[i]);
                kghAssetsToWithdraw += previewKghUndelegate(validator, tokenIds[i]);

                delete kghDelegator.kroShares[tokenIds[i]];
                delete kghDelegator.kghShares[tokenIds[i]];

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
     * @param amount The amount of reward to be claimed.
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

        _updateWeightTree(msg.sender, _vaults[msg.sender]);

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
    function finalizeUndelegate(address validator) public returns (uint128) {
        require(
            _vaults[validator].asset.totalPendingKroShares > 0,
            "AssetManager: No pending shares to finalize"
        );

        Asset storage asset = _vaults[validator].asset;
        KroDelegator storage delegator = _vaults[validator].kroDelegators[msg.sender];
        uint256[] memory requestTimes = delegator.undelegateRequestTimes;
        uint128 assetsToUndelegate;
        uint128 sharesToUndelegate;

        for (uint256 index = requestTimes.length; index > 0; ) {
            uint256 i = index - 1;
            if (requestTimes[i] + UNDELEGATION_PERIOD <= block.timestamp) {
                // If the latest undelegation request has no pending shares, then break the loop.
                if (delegator.pendingKroShares[requestTimes[i]] > 0) {
                    uint128 sharesToBurn = delegator.pendingKroShares[requestTimes[i]];

                    unchecked {
                        assetsToUndelegate += sharesToBurn.mulDiv(
                            asset.totalPendingAssets,
                            asset.totalPendingKroShares
                        );
                        sharesToUndelegate += sharesToBurn;
                    }

                    delete delegator.pendingKroShares[requestTimes[i]];
                    delete delegator.undelegateRequestTimes[i];
                } else {
                    break;
                }
            }

            unchecked {
                --index;
            }
        }

        unchecked {
            asset.totalPendingAssets -= assetsToUndelegate;
            asset.totalPendingKroShares -= sharesToUndelegate;
        }

        SafeERC20.safeTransfer(ASSET_TOKEN, msg.sender, assetsToUndelegate);

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
        Asset storage asset = _vaults[validator].asset;
        KghDelegator storage kghDelegator = _vaults[validator].kghDelegators[msg.sender];
        uint256[] memory requestTimes = kghDelegator.undelegateRequestTimes;

        uint128 kroAssetsToUndelegate;
        uint128 kroSharesToUndelegate;
        uint128 kghAssetsToUndelegate;
        uint128 kghSharesToUndelegate;

        for (uint256 index = requestTimes.length; index > 0; ) {
            uint256 i = index - 1;
            if (requestTimes[i] + UNDELEGATION_PERIOD <= block.timestamp) {
                // If the latest undelegation request has no pending shares or no pending KGHs,
                // then break the loop.
                if (
                    kghDelegator.pendingShares[requestTimes[i]].kroShares > 0 ||
                    kghDelegator.pendingKghIds[requestTimes[i]].length > 0
                ) {
                    uint256[] memory tokenId = kghDelegator.pendingKghIds[requestTimes[i]];

                    if (asset.totalPendingKghShares > 0) {
                        UndelegateShares memory pendingShares = kghDelegator.pendingShares[
                            requestTimes[i]
                        ];
                        uint128 kroSharesToBurn = pendingShares.kroShares;
                        uint128 kghSharesToBurn = pendingShares.kghShares;
                        unchecked {
                            kroAssetsToUndelegate += kroSharesToBurn.mulDiv(
                                asset.totalPendingAssets,
                                asset.totalPendingKroShares
                            );
                            kghAssetsToUndelegate += kghSharesToBurn.mulDiv(
                                asset.totalPendingBoostedRewards,
                                asset.totalPendingKghShares
                            );
                            kroSharesToUndelegate += kroSharesToBurn;
                            kghSharesToUndelegate += kghSharesToBurn;
                        }
                    }

                    for (uint256 tokenIdIndex = 0; tokenIdIndex < tokenId.length; ) {
                        KGH.safeTransferFrom(address(this), msg.sender, tokenId[tokenIdIndex]);

                        unchecked {
                            ++tokenIdIndex;
                        }
                    }

                    delete kghDelegator.pendingShares[requestTimes[i]];
                    delete kghDelegator.pendingKghIds[requestTimes[i]];
                    delete kghDelegator.undelegateRequestTimes[i];
                } else {
                    break;
                }
            }

            unchecked {
                --index;
            }
        }

        unchecked {
            asset.totalPendingAssets -= kroAssetsToUndelegate;
            asset.totalPendingKroShares -= kroSharesToUndelegate;
            asset.totalPendingBoostedRewards -= kghAssetsToUndelegate;
            asset.totalPendingKghShares -= kghSharesToUndelegate;
        }

        uint128 assetsToUndelegate = kroAssetsToUndelegate + kghAssetsToUndelegate;

        SafeERC20.safeTransfer(ASSET_TOKEN, msg.sender, assetsToUndelegate);

        emit KghUndelegationFinalized(validator, msg.sender, assetsToUndelegate);
        return assetsToUndelegate;
    }

    /**
     * @notice Finalize the reward claim of the validator.
     */
    function finalizeClaimValidatorReward() external {
        Reward storage reward = _vaults[msg.sender].reward;
        uint256[] memory requestTimes = reward.claimRequestTimes;

        uint128 assetsToWithdraw;
        for (uint256 index = requestTimes.length; index > 0; ) {
            uint256 i = index - 1;
            if (requestTimes[i] + UNDELEGATION_PERIOD <= block.timestamp) {
                if (reward.pendingValidatorRewards[requestTimes[i]] > 0) {
                    assetsToWithdraw += reward.pendingValidatorRewards[requestTimes[i]];

                    delete reward.pendingValidatorRewards[requestTimes[i]];
                    delete reward.claimRequestTimes[i];
                } else {
                    break;
                }
            }

            unchecked {
                --index;
            }
        }

        // To prevent the underflow of the totalPendingValidatorRewards when the validator is slashed.
        if (reward.totalPendingValidatorRewards < assetsToWithdraw) {
            assetsToWithdraw = reward.totalPendingValidatorRewards;
        }

        reward.totalPendingValidatorRewards -= assetsToWithdraw;
        SafeERC20.safeTransfer(ASSET_TOKEN, msg.sender, assetsToWithdraw);

        emit RewardClaimFinalized(msg.sender, assetsToWithdraw);
    }

    /**
     * @notice Internal function to add base and boosted reward to the vaults of finalized output submitters.
     *
     * @return Whether the reward distribution is done at least once or not.
     */
    function _distributeReward() internal returns (bool) {
        // TODO - Query the index from L2OutputOracle
        uint256 outputIndex = _nextRewardOutputIndex;
        uint128 finalizedOutputNum = 0;
        Types.CheckpointOutput memory output;

        for (; finalizedOutputNum < MAX_OUTPUT_FINALIZATIONS; ) {
            uint256 expiryTime = _submittedOutputs[outputIndex];
            if (block.timestamp >= expiryTime && expiryTime != 0) {
                delete _submittedOutputs[outputIndex];
                output = L2_ORACLE.getL2Output(outputIndex);
                _increaseBalanceWithReward(output.submitter);

                uint128 challengeReward = _pendingChallengeReward[outputIndex];
                if (challengeReward > 0) {
                    _modifyBalanceWithSlashing(_vaults[output.submitter], challengeReward, false);
                    delete _pendingChallengeReward[outputIndex];

                    emit ChallengeRewardDistributed(output.submitter, challengeReward);
                }

                _updateWeightTree(output.submitter, _vaults[output.submitter]);

                unchecked {
                    ++outputIndex;
                    ++finalizedOutputNum;
                }
            } else {
                break;
            }
        }

        if (finalizedOutputNum > 0) {
            unchecked {
                _nextRewardOutputIndex = outputIndex;
            }
            return true;
        }

        return false;
    }

    /**
     * @notice Slash KRO at the vault of the validator.
     *
     * @param loser       Address of the loser at the challenge.
     * @param winner      Address of the winner at the challenge.
     * @param outputIndex The index of output challenged.
     */
    function slash(address loser, address winner, uint256 outputIndex) external onlyColosseum {
        Vault storage loserVault = _vaults[loser];
        uint128 amountToSlash = _modifyBalanceWithSlashing(loserVault, 0, true);
        _pendingChallengeReward[outputIndex] = amountToSlash;

        _updateWeightTree(loser, loserVault);

        emit Slashed(loser, winner, amountToSlash);
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
     * @param basedRewards   The amount of based rewards to add as pending asset.
     * @param boostedRewards The amount of boosted rewards to add as pending asset.
     * @param shares         The amount of KRO and KGH shares to add as pending share.
     */
    function _addPendingKghShares(
        Vault storage vault,
        uint256[] memory tokenIds,
        uint128 basedRewards,
        uint128 boostedRewards,
        UndelegateShares memory shares
    ) internal {
        for (uint256 i = 0; i < tokenIds.length; ) {
            vault.kghDelegators[msg.sender].pendingKghIds[block.timestamp].push(tokenIds[i]);

            unchecked {
                ++i;
            }
        }

        unchecked {
            vault.asset.totalPendingAssets += basedRewards;
            vault.asset.totalPendingKroShares += shares.kroShares;
            vault.asset.totalPendingBoostedRewards += boostedRewards;
            vault.asset.totalPendingKghShares += shares.kghShares;

            vault.kghDelegators[msg.sender].pendingShares[block.timestamp].kroShares += shares
                .kroShares;
            vault.kghDelegators[msg.sender].pendingShares[block.timestamp].kghShares += shares
                .kghShares;
            vault.kghDelegators[msg.sender].undelegateRequestTimes.push(block.timestamp);
        }
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
     * @param validator Address of the validator.
     * @param owner     Address of the delegator.
     * @param assets    The amount of KRO to delegate.
     * @param shares    The amount of shares to receive.
     */
    function _delegate(
        address validator,
        address owner,
        uint128 assets,
        uint128 shares
    ) internal virtual {
        Vault storage vault = _vaults[validator];
        SafeERC20.safeTransferFrom(ASSET_TOKEN, owner, address(this), assets);

        unchecked {
            vault.asset.totalKro += assets;
            vault.asset.totalKroShares += shares;
            vault.kroDelegators[owner].shares += shares;

            if (owner == validator) {
                vault.asset.validatorKro += assets;
            }
        }

        _updateWeightTree(validator, vault);
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
            vault.kghDelegators[msg.sender].kroShares[tokenId] = kroShares;
            vault.kghDelegators[msg.sender].kghShares[tokenId] = kghShares;
        }

        _updateWeightTree(validator, vault);
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

        _updateWeightTree(validator, _vaults[validator]);
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

        _updateWeightTree(validator, vault);
        _tryRemoveFromWeightTree(validator, vault);
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

        uint256[] memory tokenIds = new uint256[](1);
        tokenIds[0] = tokenId;

        unchecked {
            vault.asset.totalKroShares -= kroShares;
            vault.asset.totalKghShares -= kghShares;
            delete vault.kghDelegators[owner].kroShares[tokenId];
            delete vault.kghDelegators[owner].kghShares[tokenId];

            uint128 kroInKgh = KGH_MANAGER.totalKroInKgh(tokenId);
            uint128 baseRewardsToReceive = kroAssetsToWithdraw - kroInKgh;
            uint128 boostedRewardsToReceive = kghAssetsToWithdraw - VKRO_PER_KGH;

            vault.asset.totalKro -= kroAssetsToWithdraw;
            vault.asset.totalKroInKgh -= kroInKgh;
            vault.asset.totalKgh -= 1;
            vault.reward.boostedReward -= boostedRewardsToReceive;

            UndelegateShares memory pendingShares = UndelegateShares({
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

        _updateWeightTree(validator, vault);
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

        unchecked {
            vault.asset.totalKroShares -= kroShares;
            vault.asset.totalKghShares -= kghShares;

            uint128 baseRewardsToReceive = kroAssetsToWithdraw - kroInKghs;
            uint128 boostedRewardsToReceive = kghAssetsToWithdraw -
                VKRO_PER_KGH *
                uint128(tokenIds.length);

            vault.asset.totalKro -= kroAssetsToWithdraw;
            vault.asset.totalKroInKgh -= kroInKghs;
            vault.asset.totalKgh -= uint128(tokenIds.length);
            vault.reward.boostedReward -= boostedRewardsToReceive;

            UndelegateShares memory pendingShares = UndelegateShares({
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

        _updateWeightTree(validator, vault);
    }

    /**
     * @notice Internal function to distribute the reward and update the weight of the validator.
     *
     * @param validator Address of the validator.
     */
    function _increaseBalanceWithReward(address validator) internal {
        Vault storage vault = _vaults[validator];
        uint128 boostedReward = _getBoostedReward(vault.asset.totalKgh);
        uint128 commissionRate = uint128(vault.reward.commissionRate);
        uint128 validatorReward;

        unchecked {
            validatorReward = (BASE_REWARD + boostedReward).mulDiv(commissionRate, 100);

            vault.asset.totalKro += BASE_REWARD.mulDiv(100 - commissionRate, 100);
            vault.reward.boostedReward += boostedReward.mulDiv(100 - commissionRate, 100);
            vault.reward.validatorRewardKro += validatorReward;
        }

        // TODO - Distribute the reward from a designated vault to the ValidatorManager contract.

        emit RewardDistributed(validator, validatorReward);
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
            amountToSlashOrAdd = totalAmount.mulDiv(SLASHING_RATE_NUMERATOR, PER_MILLE_DENOM);
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

            SafeERC20.safeTransfer(ASSET_TOKEN, SECURITY_COUNCIL, tax);

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
     * @notice Internal function to try to remove the validator from the weight tree.
     *
     * @param validator Address of the validator.
     * @param vault     Vault of the validator.
     */
    function _tryRemoveFromWeightTree(address validator, Vault storage vault) internal {
        if (vault.asset.validatorKro < MIN_START_AMOUNT) {
            _validatorTree.remove(validator);
        }
    }

    /**
     * @notice Internal function to update the weight tree of the validator.
     *
     * @param validator Address of the validator.
     * @param vault     Vault of the validator.
     */
    function _updateWeightTree(address validator, Vault storage vault) internal {
        _validatorTree.update(
            validator,
            uint120(
                (vault.asset.totalKro +
                    vault.reward.boostedReward +
                    vault.reward.validatorRewardKro)
            )
        );
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
