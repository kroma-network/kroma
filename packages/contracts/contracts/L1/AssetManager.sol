// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { SafeERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { IERC721 } from "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import { IERC721Receiver } from "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";

import { Uint128Math } from "../libraries/Uint128Math.sol";
import { IKGHManager } from "../universal/IKGHManager.sol";
import { ISemver } from "../universal/ISemver.sol";
import { IAssetManager } from "./interfaces/IAssetManager.sol";
import { IValidatorManager } from "./interfaces/IValidatorManager.sol";

/**
 * @title AssetManager
 * @notice AssetManager is an abstract contract that handles (un)delegations of KRO and KGH, and
 *         the distribution of rewards to the delegators and the validator.
 */
contract AssetManager is ISemver, IERC721Receiver, IAssetManager {
    using SafeERC20 for IERC20;
    using Uint128Math for uint128;

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
     * @notice The denominator for the slashing rate.
     */
    uint128 public constant SLASHING_RATE_DENOM = 1000;

    /**
     * @notice Address of the KRO token contract.
     */
    IERC20 public immutable ASSET_TOKEN;

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
     * @notice Address of ValidatorManager contract. Can be updated via upgrade.
     */
    IValidatorManager public immutable VALIDATOR_MANAGER;

    /**
     * @notice Delay for the finalization of undelegation.
     */
    uint256 public immutable UNDELEGATION_PERIOD;

    /**
     * @notice The numerator of the slashing rate.
     */
    uint128 public immutable SLASHING_RATE;

    /**
     * @notice Minimum amount to slash. It should be equal or less than
     *         ValidatorManager.MIN_START_AMOUNT.
     */
    uint128 public immutable MIN_SLASHING_AMOUNT;

    /**
     * @notice A mapping of validator address to the vault.
     */
    mapping(address => Vault) internal _vaults;

    /**
     * @notice Modifier to check if the caller is the ValidatorManager contract.
     */
    modifier onlyValidatorManager() {
        if (msg.sender != address(VALIDATOR_MANAGER)) revert NotAllowedCaller();
        _;
    }

    /**
     * @notice Modifier to check if the vault is active.
     */
    modifier checkIsActive(address validator) {
        if (
            msg.sender != validator &&
            (VALIDATOR_MANAGER.getStatus(validator) < IValidatorManager.ValidatorStatus.ACTIVE ||
                VALIDATOR_MANAGER.inJail(validator))
        ) revert ImproperValidatorStatus();
        _;
    }

    /**
     * @notice Semantic version.
     * @custom:semver 1.0.0
     */
    string public constant version = "1.0.0";

    /**
     * @notice Constructs the AssetManager contract.
     *
     * @param _assetToken         Address of the KRO token.
     * @param _kgh                Address of the KGH token.
     * @param _kghManager         Address of the KGHManager contract.
     * @param _securityCouncil    Address of the SecurityCouncil contract.
     * @param _validatorManager   Address of the ValidatorManager contract.
     * @param _undelegationPeriod Period that should wait to finalize the undelegation.
     * @param _slashingRate       Numerator of the slashing rate.
     * @param _minSlashingAmount  Minimum amount to slash.
     */
    constructor(
        IERC20 _assetToken,
        IERC721 _kgh,
        IKGHManager _kghManager,
        address _securityCouncil,
        IValidatorManager _validatorManager,
        uint128 _undelegationPeriod,
        uint128 _slashingRate,
        uint128 _minSlashingAmount
    ) {
        if (_slashingRate > SLASHING_RATE_DENOM) revert InvalidConstructorParams();

        ASSET_TOKEN = _assetToken;
        KGH = _kgh;
        KGH_MANAGER = _kghManager;
        SECURITY_COUNCIL = _securityCouncil;
        VALIDATOR_MANAGER = _validatorManager;
        UNDELEGATION_PERIOD = _undelegationPeriod;
        SLASHING_RATE = _slashingRate;
        MIN_SLASHING_AMOUNT = _minSlashingAmount;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function getKroTotalBalance(
        address validator,
        address delegator
    ) external view returns (uint128) {
        uint128 shares = _vaults[validator].kroDelegators[delegator].shares;
        return previewUndelegate(validator, shares);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function getKghTotalBalance(
        address validator,
        address delegator,
        uint256 tokenId
    ) external view returns (uint128) {
        uint128 kroInKgh = KGH_MANAGER.totalKroInKgh(tokenId);

        uint128 kghAssets = previewKghUndelegate(
            validator,
            _vaults[validator].kghDelegators[delegator].shares[tokenId].kgh
        ) - VKRO_PER_KGH;
        uint128 kroAssets = previewUndelegate(
            validator,
            _vaults[validator].kghDelegators[delegator].shares[tokenId].kro
        ) - kroInKgh;
        return kghAssets + kroAssets;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function getKroTotalShareBalance(
        address validator,
        address delegator
    ) external view returns (uint128) {
        return _vaults[validator].kroDelegators[delegator].shares;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function getKghTotalShareBalance(
        address validator,
        address delegator,
        uint256 tokenId
    ) external view returns (uint128, uint128) {
        return (
            _vaults[validator].kghDelegators[delegator].shares[tokenId].kro,
            _vaults[validator].kghDelegators[delegator].shares[tokenId].kgh
        );
    }

    /**
     * @inheritdoc IAssetManager
     */
    function previewDelegate(address validator, uint128 assets) public view returns (uint128) {
        return _convertToKroShares(validator, assets);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function previewUndelegate(address validator, uint128 shares) public view returns (uint128) {
        return _convertToKroAssets(validator, shares);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function previewKghDelegate(address validator) public view returns (uint128) {
        return _convertToKghShares(validator, VKRO_PER_KGH);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function previewKghUndelegate(
        address validator,
        uint256 tokenId
    ) public view returns (uint128) {
        return
            _convertToKghAssets(
                validator,
                _vaults[validator].kghDelegators[msg.sender].shares[tokenId].kgh
            );
    }

    /**
     * @inheritdoc IAssetManager
     */
    function totalKroAssets(address validator) public view returns (uint128) {
        return _vaults[validator].asset.totalKro;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function totalKghNum(address validator) external view returns (uint128) {
        return _vaults[validator].asset.totalKgh;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function totalKroInKgh(address validator) external view returns (uint128) {
        return _vaults[validator].asset.totalKroInKgh;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function totalValidatorKro(address validator) external view returns (uint128) {
        return _vaults[validator].asset.validatorKro;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function reflectiveWeight(address validator) external view returns (uint128) {
        return
            _vaults[validator].asset.totalKro +
            _vaults[validator].asset.boostedReward +
            _vaults[validator].asset.validatorRewardKro;
    }

    /**
     * @notice Returns the total amount of KGH assets held by the vault.
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of KGH assets held by the vault.
     */
    function _totalKghAssets(address validator) internal view returns (uint128) {
        return
            _vaults[validator].asset.totalKgh *
            VKRO_PER_KGH +
            _vaults[validator].asset.boostedReward;
    }

    /**
     * @notice Returns the total amount of KRO shares held by the vault.
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of shares held by the validator vault.
     */
    function _totalKroShares(address validator) internal view returns (uint128) {
        return _vaults[validator].asset.totalKroShares;
    }

    /**
     * @notice Returns the total amount of KGH shares held by the vault.
     *
     * @param validator Address of the validator.
     *
     * @return The total amount of shares held by the validator vault.
     */
    function _totalKghShares(address validator) internal view returns (uint128) {
        return _vaults[validator].asset.totalKghShares;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function delegate(
        address validator,
        uint128 assets
    ) external checkIsActive(validator) returns (uint128) {
        if (assets == 0) revert NotAllowedZeroInput();
        uint128 shares = _delegate(validator, msg.sender, assets, true);
        emit KroDelegated(validator, msg.sender, assets, shares);
        return shares;
    }

    /**
     * @notice Delegate KRO to the validator and returns the amount of shares that the vault would
     *        exchange. This function is only called by the ValidatorManager contract.
     *
     * @param validator Address of the validator.
     * @param assets    The amount of KRO to delegate.
     *
     * @return The amount of shares that the Vault would exchange for the amount of assets provided.
     */
    function delegateToRegister(
        address validator,
        uint128 assets
    ) external onlyValidatorManager returns (uint128) {
        uint128 shares = _delegate(validator, validator, assets, false);
        return shares;
    }

    /**
     * @inheritdoc IAssetManager
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
     * @inheritdoc IAssetManager
     */
    function delegateKghBatch(
        address validator,
        uint256[] calldata tokenIds
    ) external checkIsActive(validator) returns (uint128, uint128) {
        if (tokenIds.length == 0) revert NotAllowedZeroInput();

        uint128 kroShares;
        uint128 kghShares = previewKghDelegate(validator);
        uint128 kroInKghs;

        for (uint256 i = 0; i < tokenIds.length; ) {
            KGH.safeTransferFrom(msg.sender, address(this), tokenIds[i]);

            uint128 kroInKgh = KGH_MANAGER.totalKroInKgh(tokenIds[i]);
            uint128 kroSharesForTokenId = previewDelegate(validator, kroInKgh);

            _vaults[validator].kghDelegators[msg.sender].shares[tokenIds[i]] = KghDelegatorShares({
                kro: kroSharesForTokenId,
                kgh: kghShares
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
     * @inheritdoc IAssetManager
     */
    function initUndelegate(address validator, uint128 shares) external {
        if (shares == 0) revert NotAllowedZeroInput();
        if (shares > _vaults[validator].kroDelegators[msg.sender].shares)
            revert InsufficientShare();

        uint128 assets = previewUndelegate(validator, shares);
        if (assets == 0) revert InsufficientAsset();

        _initUndelegate(validator, msg.sender, assets, shares);
        emit KroUndelegationInitiated(validator, msg.sender, assets, shares);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function initUndelegateKgh(address validator, uint256 tokenId) external {
        uint128 kroShares = _vaults[validator].kghDelegators[msg.sender].shares[tokenId].kro;
        uint128 kghShares = _vaults[validator].kghDelegators[msg.sender].shares[tokenId].kgh;

        if (kghShares == 0) revert ShareNotExists();
        _initUndelegateKgh(validator, msg.sender, tokenId, kroShares, kghShares);

        emit KghUndelegationInitiated(validator, msg.sender, tokenId, kroShares, kghShares);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function initUndelegateKghBatch(address validator, uint256[] calldata tokenIds) external {
        if (tokenIds.length == 0) revert NotAllowedZeroInput();

        mapping(uint256 => KghDelegatorShares) storage shares = _vaults[validator]
            .kghDelegators[msg.sender]
            .shares;
        uint128 kroShares;
        uint128 kghShares;
        uint128 kroInKghs;
        uint128 kghAssetsToWithdraw;
        for (uint256 i = 0; i < tokenIds.length; ) {
            uint128 kroSharesForTokenId = shares[tokenIds[i]].kro;
            uint128 kghSharesForTokenId = shares[tokenIds[i]].kgh;

            if (kghSharesForTokenId == 0) revert ShareNotExists();

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
     * @inheritdoc IAssetManager
     */
    function initClaimValidatorReward(uint128 amount) external {
        Asset storage asset = _vaults[msg.sender].asset;
        Pending storage pendingAsset = _vaults[msg.sender].pending;

        if (amount == 0) revert NotAllowedZeroInput();
        if (amount > asset.validatorRewardKro) revert InsufficientAsset();

        unchecked {
            asset.validatorRewardKro -= amount;
            pendingAsset.pendingValidatorRewards[block.timestamp] += amount;
            pendingAsset.totalPendingValidatorRewards += amount;
            pendingAsset.claimRequestTimes.push(block.timestamp);
        }

        VALIDATOR_MANAGER.updateValidatorTree(msg.sender, true);

        emit RewardClaimInitiated(msg.sender, amount);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function finalizeUndelegate(address validator) external returns (uint128) {
        Pending storage pendingAsset = _vaults[validator].pending;
        if (pendingAsset.totalPendingKroShares == 0) revert PendingNotExists();

        KroDelegator storage delegator = _vaults[validator].kroDelegators[msg.sender];
        uint256[] memory requestTimes = delegator.undelegateRequestTimes;
        if (requestTimes.length == 0) revert RequestNotExists();

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

        if (sharesToUndelegate == 0) revert FinalizedPendingNotExists();

        unchecked {
            assetsToUndelegate = sharesToUndelegate.mulDiv(
                pendingAsset.totalPendingAssets,
                pendingAsset.totalPendingKroShares
            );
            pendingAsset.totalPendingAssets -= assetsToUndelegate;
            pendingAsset.totalPendingKroShares -= sharesToUndelegate;
        }

        ASSET_TOKEN.safeTransfer(msg.sender, assetsToUndelegate);

        emit KroUndelegationFinalized(validator, msg.sender, assetsToUndelegate);
        return assetsToUndelegate;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function finalizeUndelegateKgh(address validator) external returns (uint128) {
        KghDelegator storage kghDelegator = _vaults[validator].kghDelegators[msg.sender];
        uint256[] memory requestTimes = kghDelegator.undelegateRequestTimes;

        if (requestTimes.length == 0) revert RequestNotExists();

        Pending storage pendingAsset = _vaults[validator].pending;
        bool rewardExists = pendingAsset.totalPendingKghShares > 0;
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
                        kroSharesToUndelegate += pendingShares.kro;
                        kghSharesToUndelegate += pendingShares.kgh;
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

        if (finalizedNum == 0) revert FinalizedPendingNotExists();

        if (rewardExists) {
            unchecked {
                uint128 kroAssetsToUndelegate = kroSharesToUndelegate.mulDiv(
                    pendingAsset.totalPendingAssets,
                    pendingAsset.totalPendingKroShares
                );
                uint128 kghAssetsToUndelegate = kghSharesToUndelegate.mulDiv(
                    pendingAsset.totalPendingBoostedRewards,
                    pendingAsset.totalPendingKghShares
                );
                pendingAsset.totalPendingAssets -= kroAssetsToUndelegate;
                pendingAsset.totalPendingKroShares -= kroSharesToUndelegate;
                pendingAsset.totalPendingBoostedRewards -= kghAssetsToUndelegate;
                pendingAsset.totalPendingKghShares -= kghSharesToUndelegate;
                assetsToUndelegate = kroAssetsToUndelegate + kghAssetsToUndelegate;
            }

            ASSET_TOKEN.safeTransfer(msg.sender, assetsToUndelegate);
        }

        emit KghUndelegationFinalized(validator, msg.sender, assetsToUndelegate);
        return assetsToUndelegate;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function finalizeClaimValidatorReward() external {
        Pending storage pendingAsset = _vaults[msg.sender].pending;
        if (pendingAsset.totalPendingValidatorRewards == 0) revert PendingNotExists();

        uint256[] memory requestTimes = pendingAsset.claimRequestTimes;
        uint128 rewardsToClaim;
        for (uint256 i = requestTimes.length - 1; i >= 0 && requestTimes[i] > 0; ) {
            if (requestTimes[i] + UNDELEGATION_PERIOD <= block.timestamp) {
                unchecked {
                    rewardsToClaim += pendingAsset.pendingValidatorRewards[requestTimes[i]];
                }

                delete pendingAsset.pendingValidatorRewards[requestTimes[i]];
                delete pendingAsset.claimRequestTimes[i];
            }

            if (i == 0) {
                break;
            }
            unchecked {
                --i;
            }
        }

        if (rewardsToClaim == 0) revert FinalizedPendingNotExists();

        // To prevent the underflow of the totalPendingValidatorRewards when the validator is slashed.
        if (pendingAsset.totalPendingValidatorRewards < rewardsToClaim) {
            rewardsToClaim = pendingAsset.totalPendingValidatorRewards;
        }

        unchecked {
            pendingAsset.totalPendingValidatorRewards -= rewardsToClaim;
        }
        ASSET_TOKEN.safeTransfer(msg.sender, rewardsToClaim);

        emit RewardClaimFinalized(msg.sender, rewardsToClaim);
    }

    /**
     * @notice Update the vault of validator with the distributed reward. This function is only
     *         called by the ValidatorManager contract.
     *
     * @param validator       Address of the validator.
     * @param baseReward      The base reward to distribute.
     * @param boostedReward   The boosted reward to distribute.
     * @param validatorReward The validator reward to distribute.
     */
    function increaseBalanceWithReward(
        address validator,
        uint128 baseReward,
        uint128 boostedReward,
        uint128 validatorReward
    ) external onlyValidatorManager {
        // If reward is distributed to SECURITY_COUNCIL, transfer it directly.
        if (validator == SECURITY_COUNCIL) {
            ASSET_TOKEN.safeTransfer(
                SECURITY_COUNCIL,
                baseReward + boostedReward + validatorReward
            );
        } else {
            Vault storage vault = _vaults[validator];
            unchecked {
                vault.asset.totalKro += baseReward;
                vault.asset.boostedReward += boostedReward;
                vault.asset.validatorRewardKro += validatorReward;
            }
        }

        // TODO - Distribute the reward from a designated vault to the AssetManager contract.
    }

    /**
     * @notice Modify the balance of the vault with slashing. This function is only called by the
     *         ValidatorManager contract.
     *
     * @param validator          Address of the validator.
     * @param amountToSlashOrAdd The amount to slash or add.
     * @param isLoser            True if the validator is the loser at the challenge.
     *
     * @return The amount to slash or add.
     */
    function modifyBalanceWithSlashing(
        address validator,
        uint128 amountToSlashOrAdd,
        bool isLoser
    ) external onlyValidatorManager returns (uint128) {
        Vault storage vault = _vaults[validator];

        uint128 totalAmount = vault.asset.totalKro +
            vault.pending.totalPendingAssets +
            vault.pending.totalPendingBoostedRewards +
            vault.asset.validatorRewardKro +
            vault.pending.totalPendingValidatorRewards +
            vault.asset.boostedReward -
            vault.asset.totalKroInKgh;

        uint128[6] memory arr = [
            (vault.asset.totalKro - vault.asset.totalKroInKgh),
            vault.asset.boostedReward,
            vault.pending.totalPendingAssets,
            vault.pending.totalPendingBoostedRewards,
            vault.asset.validatorRewardKro,
            vault.pending.totalPendingValidatorRewards
        ];

        if (isLoser) {
            amountToSlashOrAdd = totalAmount.mulDiv(SLASHING_RATE, SLASHING_RATE_DENOM);
            amountToSlashOrAdd = amountToSlashOrAdd > MIN_SLASHING_AMOUNT
                ? amountToSlashOrAdd
                : MIN_SLASHING_AMOUNT;

            unchecked {
                vault.asset.totalKro -= arr[0].mulDiv(amountToSlashOrAdd, totalAmount);
                vault.asset.boostedReward -= arr[1].mulDiv(amountToSlashOrAdd, totalAmount);
                vault.pending.totalPendingAssets -= arr[2].mulDiv(amountToSlashOrAdd, totalAmount);
                vault.pending.totalPendingBoostedRewards -= arr[3].mulDiv(
                    amountToSlashOrAdd,
                    totalAmount
                );
                vault.asset.validatorRewardKro -= arr[4].mulDiv(amountToSlashOrAdd, totalAmount);
                vault.pending.totalPendingValidatorRewards -= arr[5].mulDiv(
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
            // If slashing amount is distributed to SECURITY_COUNCIL, transfer it directly.
            if (validator == SECURITY_COUNCIL) {
                ASSET_TOKEN.safeTransfer(SECURITY_COUNCIL, amountToSlashOrAdd);
                return amountToSlashOrAdd;
            }

            unchecked {
                vault.asset.totalKro += arr[0].mulDiv(amountToSlashOrAdd, totalAmount);
                vault.asset.boostedReward += arr[1].mulDiv(amountToSlashOrAdd, totalAmount);
                vault.pending.totalPendingAssets += arr[2].mulDiv(amountToSlashOrAdd, totalAmount);
                vault.pending.totalPendingBoostedRewards += arr[3].mulDiv(
                    amountToSlashOrAdd,
                    totalAmount
                );
                vault.asset.validatorRewardKro += arr[4].mulDiv(amountToSlashOrAdd, totalAmount);
                vault.pending.totalPendingValidatorRewards += arr[5].mulDiv(
                    amountToSlashOrAdd,
                    totalAmount
                );
            }

            return amountToSlashOrAdd;
        }
    }

    /**
     * @notice Add pending assets and shares when undelegating KRO.
     *
     * @param vault  Vault of the validator.
     * @param assets The amount of assets to add as pending asset.
     * @param shares The amount of shares to add as pending share.
     */
    function _addPendingKroShares(Vault storage vault, uint128 assets, uint128 shares) internal {
        vault.pending.totalPendingAssets += assets;
        vault.pending.totalPendingKroShares += shares;
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

        vault.pending.totalPendingAssets += baseRewards;
        vault.pending.totalPendingKroShares += shares.kro;
        vault.pending.totalPendingBoostedRewards += boostedRewards;
        vault.pending.totalPendingKghShares += shares.kgh;

        vault.kghDelegators[msg.sender].pendingShares[block.timestamp].kro += shares.kro;
        vault.kghDelegators[msg.sender].pendingShares[block.timestamp].kgh += shares.kgh;
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
    ) internal view returns (uint128) {
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
    ) internal view returns (uint128) {
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
    ) internal view returns (uint128) {
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
    ) internal view returns (uint128) {
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
     * @param updateTree Flag to update the validator tree.
     *
     * @return The amount of shares that the Vault would exchange for the amount of assets provided.
     */
    function _delegate(
        address validator,
        address owner,
        uint128 assets,
        bool updateTree
    ) internal returns (uint128) {
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
            VALIDATOR_MANAGER.updateValidatorTree(validator, false);
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
    ) internal {
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
            VALIDATOR_MANAGER.updateValidatorTree(validator, false);
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
    ) internal {
        Asset storage asset = _vaults[validator].asset;

        unchecked {
            asset.totalKro += kroInKghs;
            asset.totalKroInKgh += kroInKghs;
            asset.totalKgh += kghCount;

            asset.totalKroShares += kroShares;
            asset.totalKghShares += kghShares;
        }

        if (kroInKghs > 0) {
            VALIDATOR_MANAGER.updateValidatorTree(validator, false);
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
    ) internal {
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

        VALIDATOR_MANAGER.updateValidatorTree(validator, true);
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
    ) internal {
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
            vault.asset.boostedReward -= boostedRewardsToReceive;

            KghDelegatorShares memory pendingShares = KghDelegatorShares({
                kro: kroShares.mulDiv(baseRewardsToReceive, kroAssetsToWithdraw),
                kgh: kghShares.mulDiv(boostedRewardsToReceive, kghAssetsToWithdraw)
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
            VALIDATOR_MANAGER.updateValidatorTree(validator, true);
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
    ) internal {
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
            vault.asset.boostedReward -= boostedRewardsToReceive;

            KghDelegatorShares memory pendingShares = KghDelegatorShares({
                kro: kroShares.mulDiv(baseRewardsToReceive, kroAssetsToWithdraw),
                kgh: kghShares.mulDiv(boostedRewardsToReceive, kghAssetsToWithdraw)
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
            VALIDATOR_MANAGER.updateValidatorTree(validator, true);
        }
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
