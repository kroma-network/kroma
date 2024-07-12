// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "@openzeppelin/contracts-upgradeable/finance/VestingWalletUpgradeable.sol";

/**
 * @custom:proxied
 * @title KromaVestingWallet
 * @notice KromaVestingWallet vests funds equally every `vestingCycle` for the duration from the
 *         start timestamp. `totalAllocation / cliffDivider` amount of the tokens will be vested at
 *         the start timestamp, and after that remaining tokens will be vested every `vestingCycle`
 *         for the duration.
 */
contract KromaVestingWallet is VestingWalletUpgradeable {
    /**
     * @notice The divider of total allocation to calculate the cliff amount to be vested at the
     *         start timestamp.
     */
    uint256 public immutable CLIFF_DIVIDER;

    /**
     * @notice The cycle that represents how often the funds are vested (unit: seconds).
     */
    uint256 public immutable VESTING_CYCLE;

    /**
     * @notice A modifier that only allows beneficiary to call.
     */
    modifier onlyBeneficiary() {
        require(msg.sender == beneficiary(), "KromaVestingWallet: caller is not beneficiary");
        _;
    }

    /**
     * @notice Constructs the KromaVestingWallet contract.
     *
     * @param _cliffDivider        The divider to calculate the cliff amount at start timestamp.
     * @param _vestingCycleSeconds The cycle that represents how often funds are vested.
     */
    constructor(uint64 _cliffDivider, uint64 _vestingCycleSeconds) {
        require(_cliffDivider > 0, "KromaVestingWallet: cliff divider is zero");
        require(_vestingCycleSeconds > 0, "KromaVestingWallet: vesting cycle is zero");

        CLIFF_DIVIDER = _cliffDivider;
        VESTING_CYCLE = _vestingCycleSeconds;

        _disableInitializers();
    }

    /**
     * @notice Initializer.
     *
     * @param _beneficiary     Address that can release funds and receives released funds.
     * @param _startTimestamp  The timestamp to start vesting.
     * @param _durationSeconds The time period for funds to fully vest.
     */
    function initialize(
        address _beneficiary,
        uint64 _startTimestamp,
        uint64 _durationSeconds
    ) public initializer {
        require(
            _durationSeconds % VESTING_CYCLE == 0,
            "KromaVestingWallet: duration should be multiple of vesting cycle"
        );

        __VestingWallet_init(_beneficiary, _startTimestamp, _durationSeconds);
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
        } else if (timestamp >= start() + duration()) {
            return totalAllocation;
        } else {
            // At the start date, cliff amount of the assets are immediately vested.
            // After, the assets will be vested proportionally every vesting cycle for the duration.
            uint256 cliffAmount = totalAllocation / CLIFF_DIVIDER;

            // Since vested in units of cycle, remove any seconds over the end of last cycle.
            uint256 vestedSeconds = timestamp - start();
            uint256 vestedSecondsFloored = vestedSeconds - (vestedSeconds % VESTING_CYCLE);
            uint256 afterVestedAmount = ((totalAllocation - cliffAmount) * vestedSecondsFloored) /
                duration();

            return cliffAmount + afterVestedAmount;
        }
    }
}
