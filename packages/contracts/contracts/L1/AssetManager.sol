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
     * @notice Minimum delegation period. Can be updated via upgrade.
     */
    uint256 public immutable MIN_DELEGATION_PERIOD;

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
     * @param _assetToken          Address of the KRO token.
     * @param _kgh                 Address of the KGH token.
     * @param _kghManager          Address of the KGHManager contract.
     * @param _securityCouncil     Address of the SecurityCouncil contract.
     * @param _validatorManager    Address of the ValidatorManager contract.
     * @param _minDelegationPeriod Minimum delegation period.
     * @param _bondAmount          Amount to bond.
     */
    constructor(
        IERC20 _assetToken,
        IERC721 _kgh,
        IKGHManager _kghManager,
        address _securityCouncil,
        IValidatorManager _validatorManager,
        uint128 _minDelegationPeriod,
        uint128 _bondAmount
    ) {
        ASSET_TOKEN = _assetToken;
        KGH = _kgh;
        KGH_MANAGER = _kghManager;
        SECURITY_COUNCIL = _securityCouncil;
        VALIDATOR_MANAGER = _validatorManager;
        MIN_DELEGATION_PERIOD = _minDelegationPeriod;
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
    function canUndelegateKroAt(
        address validator,
        address delegator
    ) public view returns (uint128) {
        return
            _vaults[validator].kroDelegators[delegator].lastDelegatedAt +
            uint128(MIN_DELEGATION_PERIOD);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function canUndelegateKghAt(
        address validator,
        address delegator,
        uint256 tokenId
    ) public view returns (uint128) {
        return
            _vaults[validator].kghDelegators[delegator].delegatedAt[tokenId] +
            uint128(MIN_DELEGATION_PERIOD);
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
        uint128 shares = _convertToKroShares(validator, assets);
        _delegate(validator, msg.sender, assets, shares);
        VALIDATOR_MANAGER.updateValidatorTree(validator, false);

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
        // claim boost reward and auto compounding it
        uint128 boostRewardAssets = _claimBoostedReward(validator, msg.sender);
        if (boostRewardAssets > 0) {
            uint128 boostRewardShares = _convertToKroShares(validator, boostRewardAssets);
            emit KghRewardClaimed(validator, msg.sender, boostRewardAssets, boostRewardShares);
            _delegate(validator, msg.sender, boostRewardAssets, boostRewardShares);
        }

        KGH.safeTransferFrom(msg.sender, address(this), tokenId);

        uint128 kroInKgh = KGH_MANAGER.totalKroInKgh(tokenId);
        uint128 kroShares = _convertToKroShares(validator, kroInKgh);

        _delegateKgh(validator, msg.sender, tokenId, kroInKgh, kroShares);

        if (boostRewardAssets + kroInKgh > 0) {
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

        // claim boost reward and auto compounding it
        uint128 boostRewardAssets = _claimBoostedReward(validator, msg.sender);
        if (boostRewardAssets > 0) {
            uint128 boostRewardShares = _convertToKroShares(validator, boostRewardAssets);
            emit KghRewardClaimed(validator, msg.sender, boostRewardAssets, boostRewardShares);
            _delegate(validator, msg.sender, boostRewardAssets, boostRewardShares);
        }

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

        if (boostRewardAssets + kroInKghs > 0) {
            VALIDATOR_MANAGER.updateValidatorTree(validator, false);
        }

        emit KghBatchDelegated(validator, msg.sender, tokenIds, kroInKghs, kroShares);
        return kroShares;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function undelegate(address validator, uint128 assets) external {
        if (assets == 0) revert NotAllowedZeroInput();
        uint128 shares = _convertToKroShares(validator, assets);
        if (shares == 0) revert InsufficientShare();
        if (shares > _vaults[validator].kroDelegators[msg.sender].shares)
            revert InsufficientShare();

        if (canUndelegateKroAt(validator, msg.sender) > block.timestamp)
            revert NotElapsedMinDelegationPeriod();

        _undelegate(validator, msg.sender, assets, shares);
        VALIDATOR_MANAGER.updateValidatorTree(validator, true);
        ASSET_TOKEN.safeTransfer(msg.sender, assets);

        emit KroUndelegated(validator, msg.sender, assets, shares);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function undelegateKgh(address validator, uint256 tokenId) external {
        KghDelegator storage kghDelegator = _vaults[validator].kghDelegators[msg.sender];
        if (kghDelegator.delegatedAt[tokenId] == 0) revert InvalidTokenIdsInput();
        if (canUndelegateKghAt(validator, msg.sender, tokenId) > block.timestamp)
            revert NotElapsedMinDelegationPeriod();

        // KRO in KGH including base reward
        uint128 kroShares = kghDelegator.kroShares[tokenId];
        uint128 kroAssets = _convertToKroAssets(validator, kroShares);

        // boost reward of KGH
        uint128 boostReward = _claimBoostedReward(validator, msg.sender);

        // update storage
        _undelegateKgh(validator, msg.sender, tokenId, kroAssets, kroShares);

        // update validator tree
        VALIDATOR_MANAGER.updateValidatorTree(validator, true);

        // transfer KGH
        KGH.safeTransfer(msg.sender, tokenId);

        // transfer KRO
        ASSET_TOKEN.safeTransfer(msg.sender, boostReward + kroAssets);

        emit KghUndelegated(validator, msg.sender, tokenId, kroAssets + boostReward, kroShares);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function undelegateKghBatch(address validator, uint256[] calldata tokenIds) external {
        if (tokenIds.length == 0) revert NotAllowedZeroInput();

        KghDelegator storage kghDelegator = _vaults[validator].kghDelegators[msg.sender];

        // KRO in KGH including base reward
        uint128 kroShares;
        uint128 kroInKghs;
        for (uint256 i = 0; i < tokenIds.length; ) {
            if (kghDelegator.delegatedAt[tokenIds[i]] == 0) revert InvalidTokenIdsInput();
            if (canUndelegateKghAt(validator, msg.sender, tokenIds[i]) > block.timestamp)
                revert NotElapsedMinDelegationPeriod();

            uint128 kroSharesForTokenId = kghDelegator.kroShares[tokenIds[i]];

            delete kghDelegator.delegatedAt[tokenIds[i]];
            delete kghDelegator.kroShares[tokenIds[i]];

            unchecked {
                kroShares += kroSharesForTokenId;
                kroInKghs += KGH_MANAGER.totalKroInKgh(tokenIds[i]);

                ++i;
            }
        }
        uint128 kroAssets = _convertToKroAssets(validator, kroShares);

        // boost reward of KGH
        uint128 boostReward = _claimBoostedReward(validator, msg.sender);

        // update storage
        _undelegateKghBatch(
            validator,
            msg.sender,
            uint128(tokenIds.length),
            kroInKghs,
            kroAssets,
            kroShares
        );

        // update validator tree
        VALIDATOR_MANAGER.updateValidatorTree(validator, true);

        // transfer KGH
        for (uint256 i = 0; i < tokenIds.length; i++) {
            KGH.safeTransfer(msg.sender, tokenIds[i]);

            unchecked {
                ++i;
            }
        }

        // transfer KRO
        uint128 totalKroAsset = boostReward + kroAssets;
        ASSET_TOKEN.safeTransfer(msg.sender, totalKroAsset);

        emit KghBatchUndelegated(validator, msg.sender, tokenIds, totalKroAsset, kroShares);
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
        Vault storage vault = _vaults[validator];
        if (vault.asset.validatorKro - vault.asset.validatorKroBonded < BOND_AMOUNT)
            revert InsufficientAsset();

        unchecked {
            vault.asset.validatorKroBonded += BOND_AMOUNT;
        }

        emit ValidatorKroBonded(
            validator,
            BOND_AMOUNT,
            vault.asset.validatorKro - vault.asset.validatorKroBonded
        );
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
                // TODO: handle reward for boosted reward
            }
        }

        // TODO - Distribute the reward from a designated vault to the AssetManager contract.
    }

    /**
     * @notice Modify the balance of the vault during challenge. This function is only called by the
     *         ValidatorManager contract.
     *
     * @param validator       Address of the validator.
     * @param challengeReward The challenge reward to be added to the winner's asset. If 0, the
     *                        challenge reward which will be slashed from loser's asset is
     *                        determined in this function.
     * @param isLoser         True if the given validator is the loser at the challenge.
     *
     * @return The tax amount.
     * @return The challenge reward to be added to the winner's asset.
     */
    // TODO: change this according to the new design regarding to bond
    function modifyBalanceWithChallenge(
        address validator,
        uint128 challengeReward,
        bool isLoser
    ) external onlyValidatorManager returns (uint128, uint128) {
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
            challengeReward = totalAmount.mulDiv(SLASHING_RATE, SLASHING_RATE_DENOM); // TODO
            challengeReward = challengeReward > MIN_SLASHING_AMOUNT
                ? challengeReward
                : MIN_SLASHING_AMOUNT;

            unchecked {
                // Since validatorKro is included in totalKro, it does not need to be included in
                // totalAmount calculation, but we should update this value as well.
                vault.asset.validatorKro -= vault.asset.validatorKro.mulDiv(
                    challengeReward,
                    totalAmount
                );

                vault.asset.totalKro -= arr[0].mulDiv(challengeReward, totalAmount);
                vault.asset.boostedReward -= arr[1].mulDiv(challengeReward, totalAmount);
                vault.pending.totalPendingAssets -= arr[2].mulDiv(challengeReward, totalAmount);
                vault.pending.totalPendingBoostedRewards -= arr[3].mulDiv(
                    challengeReward,
                    totalAmount
                );
                vault.asset.validatorRewardKro -= arr[4].mulDiv(challengeReward, totalAmount);
                vault.pending.totalPendingValidatorRewards -= arr[5].mulDiv(
                    challengeReward,
                    totalAmount
                );
            }

            uint128 tax = challengeReward.mulDiv(TAX_NUMERATOR, TAX_DENOMINATOR);
            unchecked {
                challengeReward -= tax;
            }

            ASSET_TOKEN.safeTransfer(SECURITY_COUNCIL, tax);

            return (tax, challengeReward);
        } else {
            // If challenge reward is distributed to SECURITY_COUNCIL, transfer it directly.
            if (validator == SECURITY_COUNCIL) {
                ASSET_TOKEN.safeTransfer(SECURITY_COUNCIL, challengeReward);
                return (0, challengeReward);
            }

            unchecked {
                vault.asset.validatorKro += vault.asset.validatorKro.mulDiv(
                    challengeReward,
                    totalAmount
                );

                vault.asset.totalKro += arr[0].mulDiv(challengeReward, totalAmount);
                vault.asset.boostedReward += arr[1].mulDiv(challengeReward, totalAmount);
                vault.pending.totalPendingAssets += arr[2].mulDiv(challengeReward, totalAmount);
                vault.pending.totalPendingBoostedRewards += arr[3].mulDiv(
                    challengeReward,
                    totalAmount
                );
                vault.asset.validatorRewardKro += arr[4].mulDiv(challengeReward, totalAmount);
                vault.pending.totalPendingValidatorRewards += arr[5].mulDiv(
                    challengeReward,
                    totalAmount
                );
            }

            return (0, challengeReward);
        }
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
        Vault storage vault = _vaults[validator];
        ASSET_TOKEN.safeTransferFrom(validator, address(this), assets);

        unchecked {
            vault.asset.validatorKro += assets;
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
     * @param shares    The amount of shares to delegate.
     */
    function _delegate(
        address validator,
        address delegator,
        uint128 assets,
        uint128 shares
    ) internal {
        Vault storage vault = _vaults[validator];

        unchecked {
            vault.asset.totalKro += assets;
            vault.asset.totalKroShares += shares;
            vault.kroDelegators[delegator].shares += shares;
            vault.kroDelegators[delegator].lastDelegatedAt = uint128(block.timestamp);
        }
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
    function _undelegate(
        address validator,
        address delegator,
        uint128 assets,
        uint128 shares
    ) internal {
        Vault storage vault = _vaults[validator];

        unchecked {
            vault.asset.totalKroShares -= shares;
            vault.asset.totalKro -= assets;
            vault.kroDelegators[delegator].shares -= shares;
        }
    }

    /**
     * @notice Internal function to undelegate KGH from the validator.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param tokenId   Token Id of the KGH.
     * @param kroAssets The amount of KRO assets to undelegate.
     * @param kroShares The amount of KRO shares to undelegate.
     */
    function _undelegateKgh(
        address validator,
        address delegator,
        uint256 tokenId,
        uint128 kroAssets,
        uint128 kroShares
    ) internal {
        Asset storage asset = _vaults[validator].asset;
        KghDelegator storage kghDelegator = _vaults[validator].kghDelegators[delegator];

        uint128 kroInKgh = KGH_MANAGER.totalKroInKgh(tokenId);

        unchecked {
            // asset
            asset.totalKroShares -= kroShares;
            asset.totalKgh -= 1;
            asset.totalKro -= kroAssets;
            asset.totalKroInKgh -= kroInKgh;

            // delegator
            kghDelegator.kghNum -= 1;
            delete kghDelegator.delegatedAt[tokenId];
            delete kghDelegator.kroShares[tokenId];
        }
    }

    /**
     * @notice Internal function to undelegate KGHs from the validator.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param kghCount  The number of KGH token to undelegate.
     * @param kroInKghs The amount of KRO in the KGHs.
     * @param kroAssets The amount of KRO assets to undelegate.
     * @param kroShares The amount of KRO shares to undelegate.
     */
    function _undelegateKghBatch(
        address validator,
        address delegator,
        uint128 kghCount,
        uint128 kroInKghs,
        uint128 kroAssets,
        uint128 kroShares
    ) internal {
        Asset storage asset = _vaults[validator].asset;

        unchecked {
            // asset
            asset.totalKroShares -= kroShares;
            asset.totalKgh -= kghCount;
            asset.totalKro -= kroAssets;
            asset.totalKroInKgh -= kroInKghs;

            // delegator
            _vaults[validator].kghDelegators[delegator].kghNum -= kghCount;
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
