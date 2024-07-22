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
     * @notice The numerator of the tax.
     */
    uint128 public constant TAX_NUMERATOR = 20;

    /**
     * @notice The denominator of the tax.
     */
    uint128 public constant TAX_DENOMINATOR = 100;

    /**
     * @notice Decimals offset for the KRO shares.
     */
    uint128 public constant DECIMAL_OFFSET = 10 ** 6;

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
     * @notice The amount to bond.
     */
    uint128 public immutable BOND_AMOUNT;

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
     * @notice Modifier to check if the validator is registered and not in jail.
     */
    modifier isRegistered(address validator) {
        if (
            VALIDATOR_MANAGER.getStatus(validator) < IValidatorManager.ValidatorStatus.REGISTERED ||
            VALIDATOR_MANAGER.inJail(validator)
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
     * @param _bondAmount         Amount to bond.
     */
    constructor(
        IERC20 _assetToken,
        IERC721 _kgh,
        IKGHManager _kghManager,
        address _securityCouncil,
        IValidatorManager _validatorManager,
        uint128 _undelegationPeriod,
        uint128 _bondAmount
    ) {
        ASSET_TOKEN = _assetToken;
        KGH = _kgh;
        KGH_MANAGER = _kghManager;
        SECURITY_COUNCIL = _securityCouncil;
        VALIDATOR_MANAGER = _validatorManager;
        UNDELEGATION_PERIOD = _undelegationPeriod;
        BOND_AMOUNT = _bondAmount;
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
    function getKghNum(address validator, address delegator) external view returns (uint128) {
        return _vaults[validator].kghDelegators[delegator].kghNum;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function getKghTotalShareBalance(
        address validator,
        address delegator,
        uint256 tokenId
    ) external view returns (uint128) {
        return _vaults[validator].kghDelegators[delegator].kroShares[tokenId];
    }

    /**
     * @inheritdoc IAssetManager
     */
    function previewDelegate(address validator, uint128 assets) external view returns (uint128) {
        return _convertToKroShares(validator, assets);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function previewUndelegate(address validator, uint128 shares) external view returns (uint128) {
        return _convertToKroAssets(validator, shares);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function getKghReward(
        address validator,
        address delegator,
        uint256[] calldata tokenIds
    ) external view returns (uint128) {
        Vault storage vault = _vaults[validator];
        KghDelegator storage kghDelegator = vault.kghDelegators[delegator];

        uint128 rewardPerKghStored = vault.asset.rewardPerKghStored;
        uint128 totalBoostedReward = kghDelegator.kghNum *
            (rewardPerKghStored - kghDelegator.rewardPerKghPaid);

        (, uint128 kghBaseReward) = _calculateBaseRewardForKgh(validator, kghDelegator, tokenIds);

        return totalBoostedReward + kghBaseReward;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function getWithdrawAccount(address validator) external view returns (address) {
        return _vaults[validator].withdrawAccount;
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
    function totalValidatorKroBonded(address validator) external view returns (uint128) {
        return _vaults[validator].asset.validatorKroBonded;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function totalValidatorReward(address validator) external view returns (uint128) {
        return _vaults[validator].asset.validatorRewardKro;
    }

    /**
     * @notice Returns the reflective weight of given validator. It can be different from the actual
     *         current weight of the validator in validator tree since it includes all accumulated
     *         rewards.
     *
     * @param validator Address of the validator.
     *
     * @return The reflective weight of given validator.
     */
    // TODO: modify total weight
    function reflectiveWeight(address validator) external view returns (uint128) {
        return
            _vaults[validator].asset.totalKro +
            _vaults[validator].asset.validatorKro +
            _vaults[validator].asset.boostedReward +
            _vaults[validator].asset.validatorRewardKro;
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
     * @notice Deposit KRO to register as a validator. This function is only called by the
     *         ValidatorManager contract.
     *
     * @param validator       Address of the validator.
     * @param assets          The amount of KRO to deposit.
     * @param withdrawAccount An account where assets can be withdrawn to. Only this account can
     *                        withdraw the assets.
     */
    function depositToRegister(
        address validator,
        uint128 assets,
        address withdrawAccount
    ) external onlyValidatorManager {
        if (assets == 0) revert NotAllowedZeroInput();
        if (withdrawAccount == address(0)) revert ZeroAddress();

        _vaults[validator].withdrawAccount = withdrawAccount;
        _deposit(validator, assets, false);
        emit Deposited(validator, assets);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function deposit(uint128 assets) external {
        if (assets == 0) revert NotAllowedZeroInput();
        if (VALIDATOR_MANAGER.getStatus(msg.sender) == IValidatorManager.ValidatorStatus.NONE)
            revert ImproperValidatorStatus();

        _deposit(msg.sender, assets, true);
        emit Deposited(msg.sender, assets);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function delegate(
        address validator,
        uint128 assets
    ) external isRegistered(validator) returns (uint128) {
        if (assets == 0) revert NotAllowedZeroInput();
        ASSET_TOKEN.safeTransferFrom(msg.sender, address(this), assets);
        uint128 shares = _delegate(validator, msg.sender, assets);

        emit KroDelegated(validator, msg.sender, assets, shares);
        return shares;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function delegateKgh(
        address validator,
        uint256 tokenId
    ) external isRegistered(validator) returns (uint128) {
        // TODO: claim boosted reward and delegate to KRO pool

        KGH.safeTransferFrom(msg.sender, address(this), tokenId);

        uint128 kroInKgh = KGH_MANAGER.totalKroInKgh(tokenId);
        uint128 kroShares = _convertToKroShares(validator, kroInKgh);

        _delegateKgh(validator, msg.sender, tokenId, kroInKgh, kroShares);

        if (kroInKgh > 0) {
            VALIDATOR_MANAGER.updateValidatorTree(validator, false);
        }

        emit KghDelegated(validator, msg.sender, tokenId, kroInKgh, kroShares);
        return kroShares;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function delegateKghBatch(
        address validator,
        uint256[] calldata tokenIds
    ) external isRegistered(validator) returns (uint128) {
        if (tokenIds.length == 0) revert NotAllowedZeroInput();

        // TODO: claim boosted reward and delegate to KRO pool

        KghDelegator storage kghDelegator = _vaults[validator].kghDelegators[msg.sender];
        uint128 kroShares;
        uint128 kroInKghs;

        for (uint256 i = 0; i < tokenIds.length; ) {
            KGH.safeTransferFrom(msg.sender, address(this), tokenIds[i]);

            uint128 kroInKgh = KGH_MANAGER.totalKroInKgh(tokenIds[i]);
            uint128 kroSharesForTokenId = _convertToKroShares(validator, kroInKgh);

            kghDelegator.kroShares[tokenIds[i]] = kroSharesForTokenId;
            kghDelegator.delegatedAt[tokenIds[i]] = uint128(block.timestamp);

            unchecked {
                kroInKghs += kroInKgh;
                kroShares += kroSharesForTokenId;

                ++i;
            }
        }

        _delegateKghBatch(validator, msg.sender, uint128(tokenIds.length), kroInKghs, kroShares);

        if (kroInKghs > 0) {
            VALIDATOR_MANAGER.updateValidatorTree(validator, false);
        }

        emit KghBatchDelegated(validator, msg.sender, tokenIds, kroInKghs, kroShares);
        return kroShares;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function initUndelegate(address validator, uint128 shares) external {
        if (shares == 0) revert NotAllowedZeroInput();
        if (shares > _vaults[validator].kroDelegators[msg.sender].shares)
            revert InsufficientShare();

        uint128 assets = _convertToKroAssets(validator, shares);
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
                kghAssetsToWithdraw += _convertToKghAssets(validator, msg.sender, tokenIds[i]);

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
     * @inheritdoc IAssetManager
     */
    function claimKghReward(address validator, uint256[] calldata tokenIds) external {
        if (validator == address(0)) revert ZeroAddress();
        if (tokenIds.length == 0) revert NotAllowedZeroInput();

        Vault storage vault = _vaults[validator];
        KghDelegator storage kghDelegator = vault.kghDelegators[msg.sender];
        if (kghDelegator.kghNum != uint128(tokenIds.length)) revert InvalidTokenIdsInput();

        uint128 claimedRewards = _claimBoostedReward(validator, msg.sender);
        if (claimedRewards == 0) revert InsufficientAsset();

        (uint128 kroSharesToBurn, uint128 kghBaseReward) = _calculateBaseRewardForKgh(
            validator,
            kghDelegator,
            tokenIds
        );

        unchecked {
            claimedRewards += kghBaseReward;
            vault.asset.totalKroShares -= kroSharesToBurn;
            vault.asset.totalKro -= kghBaseReward;
        }

        ASSET_TOKEN.safeTransfer(msg.sender, claimedRewards);

        emit KghRewardClaimed(validator, msg.sender, claimedRewards, kroSharesToBurn);
    }

    /**
     * @notice Bond KRO from validator KRO during output submission or challenge creation. This
     *         function is only called by the ValidatorManager contract.
     *
     * @param validator Address of the validator.
     */
    function bondValidatorKro(address validator) external onlyValidatorManager {
        Asset storage asset = _vaults[validator].asset;
        uint128 remainder = asset.validatorKro - asset.validatorKroBonded;
        if (remainder < BOND_AMOUNT) revert InsufficientAsset();

        unchecked {
            asset.validatorKroBonded += BOND_AMOUNT;
        }

        emit ValidatorKroBonded(validator, BOND_AMOUNT, remainder - BOND_AMOUNT);
    }

    /**
     * @notice Unbond KRO from validator KRO during output finalization or challenge slashing. This
     *         function is only called by the ValidatorManager contract.
     *
     * @param validator Address of the validator.
     */
    function unbondValidatorKro(address validator) external onlyValidatorManager {
        Asset storage asset = _vaults[validator].asset;
        _unbondValidatorKro(validator, asset);
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
            Asset storage asset = _vaults[validator].asset;
            unchecked {
                asset.totalKro += baseReward;
                // TODO: handle reward for boosted reward
            }

            _unbondValidatorKro(validator, asset);
        }

        // TODO - Distribute the reward from a designated vault to the AssetManager contract.
    }

    /**
     * @notice Update the vault of challenge winner with the challenge reward. This function is only
     *         called by the ValidatorManager contract.
     *
     * @param winner          Address of the challenge winner.
     * @param challengeReward The challenge reward to be added to the winner's asset after excluding
     *                        tax.
     *
     * @return The challenge reward added to winner's asset.
     */
    function increaseBalanceWithChallenge(
        address winner,
        uint128 challengeReward
    ) external onlyValidatorManager returns (uint128) {
        Asset storage asset = _vaults[winner].asset;

        // If challenge reward is distributed to SECURITY_COUNCIL, transfer it directly.
        if (winner == SECURITY_COUNCIL) {
            ASSET_TOKEN.safeTransfer(SECURITY_COUNCIL, challengeReward);
            return;
        }

        uint128 tax = challengeReward.mulDiv(TAX_NUMERATOR, TAX_DENOMINATOR);
        ASSET_TOKEN.safeTransfer(SECURITY_COUNCIL, tax);

        unchecked {
            challengeReward -= tax;
            asset.validatorKro += challengeReward;
        }

        return challengeReward;
    }

    /**
     * @notice Update the vault of challenge loser with the challenge reward. This function is only
     *         called by the ValidatorManager contract.
     *
     * @param loser Address of the challenge loser.
     *
     * @return The challenge reward slashed from loser's asset.
     */
    function decreaseBalanceWithChallenge(
        address loser
    ) external onlyValidatorManager returns (uint128) {
        Asset storage asset = _vaults[loser].asset;

        unchecked {
            asset.validatorKroBonded -= BOND_AMOUNT;
            asset.validatorKro -= BOND_AMOUNT;
        }

        return BOND_AMOUNT;
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
     * @notice Internal function to calculate base rewards from the given validator address,
     *         delegator struct, and token ids of KGHs.
     *
     * @param validator Address of the validator.
     * @param delegator KghDelegator struct of the delegator.
     * @param tokenIds  Array of token ids of the KGH to calculate base rewards.
     *
     * @return The amount of shares corresponding to the reward.
     * @return The amount of base rewards.
     */
    function _calculateBaseRewardForKgh(
        address validator,
        KghDelegator storage delegator,
        uint256[] calldata tokenIds
    ) internal view returns (uint128, uint128) {
        uint128 kroSharesToBurn;
        for (uint256 i = 0; i < tokenIds.length; ) {
            if (kghDelegator.delegatedAt[tokenIds[i]] == 0) revert InvalidTokenIdsInput();
            uint128 kroShares = kghDelegator.kroShares[tokenIds[i]] -
                _convertToKroShares(validator, KGH_MANAGER.totalKroInKgh(tokenIds[i]));

            unchecked {
                kroSharesToBurn += kroShares;
                ++i;
            }
        }

        uint128 kghBaseReward = _convertToKroAssets(validator, kroSharesToBurn);
        return (kroSharesToBurn, kghBaseReward);
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
     * @notice Internal function to deposit KRO by the validator.
     *
     * @param validator  Address of the validator.
     * @param assets     The amount of KRO to deposit.
     * @param updateTree Flag to update the validator tree.
     */
    function _deposit(address validator, uint128 assets, bool updateTree) internal {
        Asset storage asset = _vaults[validator].asset;
        ASSET_TOKEN.safeTransferFrom(validator, address(this), assets);

        unchecked {
            asset.validatorKro += assets;
        }

        if (updateTree) {
            VALIDATOR_MANAGER.updateValidatorTree(validator, false);
        }
    }

    /**
     * @notice Internal function to delegate KRO to the validator.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param assets    The amount of KRO to delegate.
     *
     * @return The amount of shares that the Vault would exchange for the amount of assets provided.
     */
    function _delegate(
        address validator,
        address delegator,
        uint128 assets
    ) internal returns (uint128) {
        uint128 shares = _convertToKroShares(validator, assets);
        Vault storage vault = _vaults[validator];

        unchecked {
            vault.asset.totalKro += assets;
            vault.asset.totalKroShares += shares;
            vault.kroDelegators[delegator].shares += shares;
        }

        VALIDATOR_MANAGER.updateValidatorTree(validator, false);

        return shares;
    }

    /**
     * @notice Internal function to delegate KGH to the validator.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param tokenId   Token Id of the KGH.
     * @param kroInKgh  The amount of KRO in the KGH.
     * @param kroShares The amount of KRO shares to receive.
     */
    function _delegateKgh(
        address validator,
        address delegator,
        uint256 tokenId,
        uint128 kroInKgh,
        uint128 kroShares
    ) internal {
        Asset storage asset = _vaults[validator].asset;
        KghDelegator storage kghDelegator = _vaults[validator].kghDelegators[delegator];

        unchecked {
            asset.totalKro += kroInKgh;
            asset.totalKroInKgh += kroInKgh;
            asset.totalKgh += 1;
            asset.totalKroShares += kroShares;

            ++kghDelegator.kghNum;
            kghDelegator.kroShares[tokenId] = kroShares;
            kghDelegator.delegatedAt[tokenId] = uint128(block.timestamp);
        }
    }

    /**
     * @notice Internal function to delegate KGHs to the validator.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param kghCount  The number of KGHs to delegate.
     * @param kroInKghs The amount of KRO in the KGHs.
     * @param kroShares The amount of KRO shares to receive.
     */
    function _delegateKghBatch(
        address validator,
        address delegator,
        uint128 kghCount,
        uint128 kroInKghs,
        uint128 kroShares
    ) internal {
        Asset storage asset = _vaults[validator].asset;
        KghDelegator storage kghDelegator = _vaults[validator].kghDelegators[delegator];

        unchecked {
            asset.totalKro += kroInKghs;
            asset.totalKroInKgh += kroInKghs;
            asset.totalKgh += kghCount;
            asset.totalKroShares += kroShares;

            kghDelegator.kghNum += kghCount;
        }
    }

    /**
     * @notice Internal function to undelegate KRO from the validator.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param assets    The amount of KRO to undelegate.
     * @param shares    The amount of shares to undelegate.
     */
    function _initUndelegate(
        address validator,
        address delegator,
        uint128 assets,
        uint128 shares
    ) internal {
        Vault storage vault = _vaults[validator];

        unchecked {
            vault.asset.totalKroShares -= shares;
            vault.kroDelegators[delegator].shares -= shares;

            vault.asset.totalKro -= assets;
            _addPendingKroShares(vault, assets, shares);
        }

        VALIDATOR_MANAGER.updateValidatorTree(validator, true);
    }

    /**
     * @notice Internal function to undelegate KGH from the validator.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param tokenId   Token Id of the KGH.
     * @param kroShares The amount of KRO shares to undelegate.
     * @param kghShares The amount of KGH shares to undelegate.
     */
    function _initUndelegateKgh(
        address validator,
        address delegator,
        uint256 tokenId,
        uint128 kroShares,
        uint128 kghShares
    ) internal {
        Vault storage vault = _vaults[validator];
        uint128 kroAssetsToWithdraw = _convertToKroAssets(validator, kroShares);
        uint128 kghAssetsToWithdraw = _convertToKghAssets(validator, delegator, tokenId);
        uint128 boostedRewardsToReceive;

        uint256[] memory tokenIds = new uint256[](1);
        tokenIds[0] = tokenId;

        unchecked {
            vault.asset.totalKroShares -= kroShares;
            vault.asset.totalKghShares -= kghShares;
            delete vault.kghDelegators[delegator].shares[tokenId];

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
        uint128 kroAssetsToWithdraw = _convertToKroAssets(validator, kroShares);
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
     * @notice Internal function to claim the boosted reward of the delegator.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     *
     * @return The amount of the claimed boosted reward.
     */
    function _claimBoostedReward(address validator, address delegator) internal returns (uint128) {
        Vault storage vault = _vaults[validator];
        KghDelegator storage kghDelegator = vault.kghDelegators[delegator];

        uint128 rewardPerKghStored = vault.asset.rewardPerKghStored;
        uint128 totalBoostedReward = kghDelegator.kghNum *
            (rewardPerKghStored - kghDelegator.rewardPerKghPaid);

        kghDelegator.rewardPerKghPaid = rewardPerKghStored;

        return totalBoostedReward;
    }

    /**
     * @notice Internal function to unbond KRO from validator KRO during output finalization or
     *         challenge slashing.
     *
     * @param validator Address of the validator.
     * @param asset     Asset information of the vault of the validator.
     */
    function _unbondValidatorKro(address validator, Asset storage asset) internal {
        unchecked {
            asset.validatorKroBonded -= BOND_AMOUNT;

            emit ValidatorKroUnbonded(
                validator,
                BOND_AMOUNT,
                asset.validatorKro - asset.validatorKroBonded
            );
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
