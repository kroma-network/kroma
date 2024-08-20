// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { SafeERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { IERC721 } from "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import { IERC721Receiver } from "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";

import { Uint128Math } from "../libraries/Uint128Math.sol";
import { ISemver } from "../universal/ISemver.sol";
import { IAssetManager } from "./interfaces/IAssetManager.sol";
import { IValidatorManager } from "./interfaces/IValidatorManager.sol";

/**
 * @title AssetManager
 * @notice AssetManager is a contract that handles (un)delegations of KRO and KGH, and the
 *         distribution of rewards to the delegators and the validator.
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
     * @notice The address of the SecurityCouncil contract. Can be updated via upgrade.
     */
    address public immutable SECURITY_COUNCIL;

    /**
     * @notice The address of Validator Reward Vault. Can be updated via upgrade.
     */
    address public immutable VALIDATOR_REWARD_VAULT;

    /**
     * @notice Address of ValidatorManager contract. Can be updated via upgrade.
     */
    IValidatorManager public immutable VALIDATOR_MANAGER;

    /**
     * @notice Minimum delegation period. Can be updated via upgrade.
     */
    uint128 public immutable MIN_DELEGATION_PERIOD;

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
     * @notice Modifier to check if the caller is the withdraw account of the validator.
     */
    modifier onlyWithdrawAccount(address validator) {
        if (msg.sender != _vaults[validator].withdrawAccount) revert NotAllowedCaller();
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
     * @param _assetToken           Address of the KRO token.
     * @param _kgh                  Address of the KGH token.
     * @param _securityCouncil      Address of the SecurityCouncil contract.
     * @param _validatorRewardVault Address of the Validator Reward Vault.
     * @param _validatorManager     Address of the ValidatorManager contract.
     * @param _minDelegationPeriod  Minimum delegation period.
     * @param _bondAmount           Amount to bond.
     */
    constructor(
        IERC20 _assetToken,
        IERC721 _kgh,
        address _securityCouncil,
        address _validatorRewardVault,
        IValidatorManager _validatorManager,
        uint128 _minDelegationPeriod,
        uint128 _bondAmount
    ) {
        ASSET_TOKEN = _assetToken;
        KGH = _kgh;
        SECURITY_COUNCIL = _securityCouncil;
        VALIDATOR_REWARD_VAULT = _validatorRewardVault;
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
    function getKroAssets(address validator, address delegator) external view returns (uint128) {
        return _convertToKroAssets(validator, _vaults[validator].kroDelegators[delegator].shares);
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
        return _vaults[validator].kroDelegators[delegator].lastDelegatedAt + MIN_DELEGATION_PERIOD;
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
            MIN_DELEGATION_PERIOD;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function canWithdrawAt(address validator) public view returns (uint128) {
        return _vaults[validator].lastDepositedAt + MIN_DELEGATION_PERIOD;
    }

    /**
     * @inheritdoc IAssetManager
     */
    function getKghReward(address validator, address delegator) external view returns (uint128) {
        Vault storage vault = _vaults[validator];
        KghDelegator storage kghDelegator = vault.kghDelegators[delegator];

        uint128 rewardPerKghStored = vault.asset.rewardPerKghStored;
        uint128 totalBoostedReward = kghDelegator.kghNum *
            (rewardPerKghStored - kghDelegator.rewardPerKghPaid);

        return totalBoostedReward;
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
    function totalValidatorKroNotBonded(address validator) external view returns (uint128) {
        return _vaults[validator].asset.validatorKro - _vaults[validator].asset.validatorKroBonded;
    }

    /**
     * @notice Returns the reflective weight of given validator.
     *
     * @param validator Address of the validator.
     *
     * @return The reflective weight of given validator.
     */
    function reflectiveWeight(address validator) external view returns (uint128) {
        return _vaults[validator].asset.totalKro + _vaults[validator].asset.validatorKro;
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

        VALIDATOR_MANAGER.tryActivateValidator(msg.sender);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function withdraw(address validator, uint128 assets) external onlyWithdrawAccount(validator) {
        if (assets == 0) revert NotAllowedZeroInput();
        if (canWithdrawAt(validator) > block.timestamp) {
            revert NotElapsedMinDelegationPeriod();
        }
        if (VALIDATOR_MANAGER.jailExpiresAt(validator) > block.timestamp)
            revert ImproperValidatorStatus();

        _withdraw(validator, assets);

        VALIDATOR_MANAGER.updateValidatorTree(validator, true);

        ASSET_TOKEN.safeTransfer(_vaults[validator].withdrawAccount, assets);

        emit Withdrawn(validator, assets);
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
    function delegateKgh(address validator, uint256 tokenId) external isRegistered(validator) {
        // claim boosted reward
        uint128 boostedReward = _claimBoostedReward(validator, msg.sender);
        if (boostedReward > 0) {
            ASSET_TOKEN.safeTransfer(msg.sender, boostedReward);
            emit KghRewardClaimed(validator, msg.sender, boostedReward);
        }

        KGH.safeTransferFrom(msg.sender, address(this), tokenId);
        _delegateKgh(validator, msg.sender, tokenId);

        emit KghDelegated(validator, msg.sender, tokenId);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function delegateKghBatch(
        address validator,
        uint256[] calldata tokenIds
    ) external isRegistered(validator) {
        if (tokenIds.length == 0) revert NotAllowedZeroInput();

        // claim boosted reward
        uint128 boostedReward = _claimBoostedReward(validator, msg.sender);
        if (boostedReward > 0) {
            ASSET_TOKEN.safeTransfer(msg.sender, boostedReward);
            emit KghRewardClaimed(validator, msg.sender, boostedReward);
        }

        KghDelegator storage kghDelegator = _vaults[validator].kghDelegators[msg.sender];
        for (uint256 i = 0; i < tokenIds.length; ) {
            KGH.safeTransferFrom(msg.sender, address(this), tokenIds[i]);
            kghDelegator.delegatedAt[tokenIds[i]] = uint128(block.timestamp);

            unchecked {
                ++i;
            }
        }

        _delegateKghBatch(validator, msg.sender, uint128(tokenIds.length));

        emit KghBatchDelegated(validator, msg.sender, tokenIds);
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

        // boosted reward of KGH
        uint128 boostedReward = _claimBoostedReward(validator, msg.sender);

        // update storage
        _undelegateKgh(validator, msg.sender, tokenId);

        // transfer KGH
        KGH.safeTransferFrom(address(this), msg.sender, tokenId);

        // transfer KRO
        if (boostedReward > 0) {
            ASSET_TOKEN.safeTransfer(msg.sender, boostedReward);
        }

        emit KghUndelegated(validator, msg.sender, tokenId, boostedReward);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function undelegateKghBatch(address validator, uint256[] calldata tokenIds) external {
        if (tokenIds.length == 0) revert NotAllowedZeroInput();

        KghDelegator storage kghDelegator = _vaults[validator].kghDelegators[msg.sender];

        for (uint256 i = 0; i < tokenIds.length; ) {
            if (kghDelegator.delegatedAt[tokenIds[i]] == 0) revert InvalidTokenIdsInput();
            if (canUndelegateKghAt(validator, msg.sender, tokenIds[i]) > block.timestamp)
                revert NotElapsedMinDelegationPeriod();

            delete kghDelegator.delegatedAt[tokenIds[i]];

            unchecked {
                ++i;
            }
        }

        // boosted reward of KGHs
        uint128 boostedReward = _claimBoostedReward(validator, msg.sender);

        // update storage
        _undelegateKghBatch(validator, msg.sender, uint128(tokenIds.length));

        // transfer KGHs
        for (uint256 i = 0; i < tokenIds.length; ) {
            KGH.safeTransferFrom(address(this), msg.sender, tokenIds[i]);

            unchecked {
                ++i;
            }
        }

        // transfer KRO
        if (boostedReward > 0) {
            ASSET_TOKEN.safeTransfer(msg.sender, boostedReward);
        }

        emit KghBatchUndelegated(validator, msg.sender, tokenIds, boostedReward);
    }

    /**
     * @inheritdoc IAssetManager
     */
    function claimKghReward(address validator) external {
        uint128 boostedReward = _claimBoostedReward(validator, msg.sender);
        if (boostedReward == 0) revert InsufficientAsset();

        ASSET_TOKEN.safeTransfer(msg.sender, boostedReward);

        emit KghRewardClaimed(validator, msg.sender, boostedReward);
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

        unchecked {
            asset.validatorKroBonded -= BOND_AMOUNT;
        }

        emit ValidatorKroUnbonded(
            validator,
            BOND_AMOUNT,
            asset.validatorKro - asset.validatorKroBonded
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
        // Distribute the reward from a designated vault to the AssetManager contract.
        ASSET_TOKEN.safeTransferFrom(
            VALIDATOR_REWARD_VAULT,
            address(this),
            baseReward + boostedReward + validatorReward
        );

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
                asset.validatorKro += validatorReward;
                if (asset.totalKgh != 0) {
                    asset.rewardPerKghStored += boostedReward / asset.totalKgh;
                }
                asset.validatorKroBonded -= BOND_AMOUNT;
            }

            emit ValidatorKroUnbonded(
                validator,
                BOND_AMOUNT,
                asset.validatorKro - asset.validatorKroBonded
            );
        }
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
            return challengeReward;
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
     * @notice Revert the changes of decreaseBalanceWithChallenge. This function is only called by
     *         the ValidatorManager contract.
     *
     * @param loser Address of the challenge original loser.
     *
     * @return The challenge reward refunded to loser's asset.
     */
    function revertDecreaseBalanceWithChallenge(
        address loser
    ) external onlyValidatorManager returns (uint128) {
        Asset storage asset = _vaults[loser].asset;

        unchecked {
            asset.validatorKroBonded += BOND_AMOUNT;
            asset.validatorKro += BOND_AMOUNT;
        }

        return BOND_AMOUNT;
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
            vault.lastDepositedAt = uint128(block.timestamp);
        }

        if (updateTree) {
            VALIDATOR_MANAGER.updateValidatorTree(validator, false);
        }
    }

    /**
     * @notice Internal function to withdraw KRO by the validator.
     *
     * @param validator Address of the validator.
     * @param assets    The amount of KRO to withdraw.
     */
    function _withdraw(address validator, uint128 assets) internal {
        Asset storage asset = _vaults[validator].asset;
        if (assets > asset.validatorKro - asset.validatorKroBonded) revert InsufficientAsset();

        unchecked {
            asset.validatorKro -= assets;
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
     */
    function _delegateKgh(address validator, address delegator, uint256 tokenId) internal {
        Vault storage vault = _vaults[validator];
        KghDelegator storage kghDelegator = vault.kghDelegators[delegator];

        unchecked {
            vault.asset.totalKgh += 1;

            ++kghDelegator.kghNum;
            kghDelegator.delegatedAt[tokenId] = uint128(block.timestamp);
        }
    }

    /**
     * @notice Internal function to delegate KGHs to the validator.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param kghCount  The number of KGHs to delegate.
     */
    function _delegateKghBatch(address validator, address delegator, uint128 kghCount) internal {
        Vault storage vault = _vaults[validator];

        unchecked {
            // asset
            vault.asset.totalKgh += kghCount;

            // delegator
            vault.kghDelegators[delegator].kghNum += kghCount;
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
     */
    function _undelegateKgh(address validator, address delegator, uint256 tokenId) internal {
        Vault storage vault = _vaults[validator];
        KghDelegator storage kghDelegator = vault.kghDelegators[delegator];

        unchecked {
            // asset
            vault.asset.totalKgh -= 1;

            // delegator
            kghDelegator.kghNum -= 1;
            delete kghDelegator.delegatedAt[tokenId];
        }
    }

    /**
     * @notice Internal function to undelegate KGHs from the validator.
     *
     * @param validator Address of the validator.
     * @param delegator Address of the delegator.
     * @param kghCount  The number of KGH token to undelegate.
     */
    function _undelegateKghBatch(address validator, address delegator, uint128 kghCount) internal {
        Vault storage vault = _vaults[validator];

        unchecked {
            // asset
            vault.asset.totalKgh -= kghCount;

            // delegator
            vault.kghDelegators[delegator].kghNum -= kghCount;
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
