// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/finance/VestingWalletUpgradeable.sol";

/**
 * @custom:proxied
 * @title KromaVestingWallet
 * @notice KromaVestingWallet vests funds equally every `vestingCycle` for the duration from the
 *         start timestamp. It has contract owner who can change the beneficiary and migrate funds
 *         to a new wallet. `totalAllocation / cliffDivider` amount of the tokens will be vested at
 *         the start timestamp, and after that remaining tokens will be vested every `vestingCycle`
 *         for the duration.
 *         Adapted from https://github.com/ArbitrumFoundation/governance.
 */
contract KromaVestingWallet is VestingWalletUpgradeable, OwnableUpgradeable {
    using SafeERC20 for IERC20;

    /**
     * @notice The (re)declared variable overriding the one in VestingWalletUpgradeable so that it
     *         can be changed later.
     */
    address private _beneficiary;

    /**
     * @notice The divider of total allocation to calculate the cliff amount to be vested at the
     *         start timestamp.
     */
    uint64 private _cliffDivider;

    /**
     * @notice The cycle that represents how often the funds are vested (unit: seconds).
     */
    uint64 private _vestingCycle;

    /**
     * @notice Emitted when beneficiary address is changed.
     *
     * @param newBeneficiary Address of the new beneficiary.
     * @param caller         Address that called beneficiary-setter; either current beneficiary or
     *                       contract owner.
     */
    event BeneficiarySet(address indexed newBeneficiary, address indexed caller);

    /**
     * @notice Emitted when tokens are migrated to a new wallet.
     *
     * @param token       Address of token being migrated.
     * @param destination New wallet address.
     * @param amount      Amount of tokens migrated.
     */
    event TokenMigrated(address indexed token, address indexed destination, uint256 amount);

    /**
     * @notice Emitted when ETH is migrated to a new wallet.
     *
     * @param destination New wallet address.
     * @param amount      Amount of ETH migrated.
     */
    event EthMigrated(address indexed destination, uint256 amount);

    /**
     * @notice A modifier that only allows beneficiary to call.
     */
    modifier onlyBeneficiary() {
        require(msg.sender == beneficiary(), "KromaVestingWallet: caller is not beneficiary");
        _;
    }

    /**
     * @notice A modifier that only allows beneficiary or owner to call.
     */
    modifier onlyBeneficiaryOrOwner() {
        require(
            msg.sender == beneficiary() || msg.sender == owner(),
            "KromaVestingWallet: caller is not beneficiary or owner"
        );
        _;
    }

    /**
     * @notice Constructs the KromaVestingWallet contract.
     */
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initializer.
     *
     * @param _beneficiaryAddress  Address that can release funds and receives released funds.
     * @param _startTimestamp      The timestamp to start vesting.
     * @param _durationSeconds     The time period for funds to fully vest.
     * @param _cliffDividerValue   The divider to calculate the cliff amount at start timestamp.
     * @param _vestingCycleSeconds The cycle that represents how often funds are vested.
     * @param _owner               The owner of this contract.
     */
    function initialize(
        address _beneficiaryAddress,
        uint64 _startTimestamp,
        uint64 _durationSeconds,
        uint64 _cliffDividerValue,
        uint64 _vestingCycleSeconds,
        address _owner
    ) public initializer {
        require(_cliffDividerValue > 0, "KromaVestingWallet: cliff divider is zero");
        require(_vestingCycleSeconds > 0, "KromaVestingWallet: vesting cycle is zero");
        require(_owner != address(0), "KromaVestingWallet: owner is zero address");

        _cliffDivider = _cliffDividerValue;
        _vestingCycle = _vestingCycleSeconds;

        // Initiate vesting wallet.
        // The first argument (beneficiary) is unused by contract; a dummy value is provided.
        __VestingWallet_init(address(0xdead), _startTimestamp, _durationSeconds);
        _setBeneficiary(_beneficiaryAddress);

        // Set owner of this contract.
        __Ownable_init();
        _transferOwnership(_owner);
    }

    /**
     * @notice Returns the divider of total allocation to calculate the cliff amount at the start
     *         timestamp.
     */
    function cliffDivider() external view returns (uint256) {
        return _cliffDivider;
    }

    /**
     * @notice Returns the cycle that represents how often the funds are vested in seconds.
     */
    function vestingCycle() external view returns (uint256) {
        return _vestingCycle;
    }

    /**
     * @notice Returns the beneficiary of this contract.
     * @dev Overrides the one of OZ VestingWalletUpgradeable contract since it has private
     *      `_beneficiary` variable with no setter.
     *
     * @return The beneficiary of this contract.
     */
    function beneficiary() public view override returns (address) {
        return _beneficiary;
    }

    /**
     * @notice Sets new beneficiary; only the owner or current beneficiary can call.
     *
     * @param newBeneficiary New address to receive proceeds from the vesting contract.
     */
    function setBeneficiary(address newBeneficiary) external onlyBeneficiaryOrOwner {
        _setBeneficiary(newBeneficiary);
    }

    /**
     * @notice Internal function to set new beneficiary.
     *
     * @param newBeneficiary New address to receive proceeds from the vesting contract.
     */
    function _setBeneficiary(address newBeneficiary) internal {
        require(newBeneficiary != address(0), "KromaVestingWallet: beneficiary is zero address");

        _beneficiary = newBeneficiary;

        emit BeneficiarySet(newBeneficiary, msg.sender);
    }

    /**
     * @notice Releases the tokens that have already vested; only beneficiary can call.
     *
     * @param token Address of ERC20 token to release.
     */
    function release(address token) public override onlyBeneficiary {
        super.release(token);
    }

    /**
     * @notice Releases the native token (ether) that have already vested; only beneficiary can call.
     */
    function release() public override onlyBeneficiary {
        super.release();
    }

    /**
     * @notice Overrides vesting formula. This returns the amount vested, as a function of time, for
     *         an asset given its total historical allocation.
     *
     * @param totalAllocation Total historical allocation of the asset.
     * @param timestamp       Timestamp to be used to calculate the vested amount.
     *
     * @return The amount vested, as a function of time, for an asset given its total historical
     *         allocation.
     */
    function _vestingSchedule(
        uint256 totalAllocation,
        uint64 timestamp
    ) internal view override returns (uint256) {
        if (timestamp < start()) {
            return 0;
        } else if (timestamp > start() + duration()) {
            return totalAllocation;
        } else {
            // At the start date, cliff amount of the assets are immediately vested.
            // After, the assets will be vested proportionally every vesting cycle for the duration.
            uint256 cliff = totalAllocation / _cliffDivider;

            // Since vested in units of cycle, remove any seconds over the end of last cycle.
            uint256 vestedTimeSeconds = timestamp - start();
            uint256 vestedTimeSecondsFloored = vestedTimeSeconds -
                (vestedTimeSeconds % _vestingCycle);
            uint256 remaining = ((totalAllocation - cliff) * vestedTimeSecondsFloored) / duration();

            return cliff + remaining;
        }
    }

    /**
     * @notice Migrates unvested (as well as vested but not yet claimed) tokens to a new wallet.
     *         e.g. one with a different vesting schedule. Only owner can call.
     *
     * @param token     Address of token to be migrated.
     * @param newWallet Address of new wallet to receive tokens.
     */
    function migrateTokensToNewWallet(address token, address newWallet) external onlyOwner {
        require(Address.isContract(newWallet), "KromaVestingWallet: new wallet must be a contract");

        IERC20 _token = IERC20(token);
        uint256 tokenBalance = _token.balanceOf(address(this));
        _token.safeTransfer(newWallet, tokenBalance);

        emit TokenMigrated(token, newWallet, tokenBalance);
    }

    /**
     * @notice Migrates unvested (as well as vested but not yet claimed) ETH to a new wallet.
     *         e.g. one with a different vesting schedule. Only owner can call.
     *
     * @param newWallet Address of new wallet to receive ETH.
     */
    function migrateEthToNewWallet(address newWallet) external onlyOwner {
        require(Address.isContract(newWallet), "KromaVestingWallet: new wallet must be a contract");

        uint256 ethBalance = address(this).balance;
        (bool success, ) = newWallet.call{ value: ethBalance }("");
        require(success, "KromaVestingWallet: eth transfer failed");

        emit EthMigrated(newWallet, ethBalance);
    }
}
