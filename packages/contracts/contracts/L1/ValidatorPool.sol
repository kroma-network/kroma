// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {
    ReentrancyGuardUpgradeable
} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import { Types } from "../libraries/Types.sol";
import { Semver } from "../universal/Semver.sol";

/**
 * @custom:proxied
 * @title ValidatorPool
 * @notice TBD
 */
contract ValidatorPool is ReentrancyGuardUpgradeable, Semver {
    /**
     * @notice The address of the colosseum. Can be updated via upgrade.
     */
    address public immutable COLOSSEUM_ADDRESS;

    /**
     * @notice The address of the trusted validator. Can be updated via upgrade.
     */
    address public immutable TRUSTED_VALIDATOR;

    /**
     * @notice The minimum deposit amount. Can be updated via upgrade.
     */
    uint256 public immutable MIN_DEPOSIT_AMOUNT;

    /**
     * @notice The slashing amount. Can be updated via upgrade.
     */
    uint256 public immutable SLASHING_AMOUNT;

    /**
     * @notice Mapping of registered validator addresses to boolean status.
     */
    mapping(address => bool) internal validators;

    /**
     * @notice Mapping of deposits of validators.
     */
    mapping(address => uint256) internal depositsOf;

    /**
     * @notice Emitted when a new validator is registered.
     *
     * @param validator The address of new validator.
     */
    event NewValidator(address validator);

    /**
     * @notice Emitted when the validator deposits.
     *
     * @param validator The address of new validator.
     * @param amount    The deposit amount.
     * @param balance   The deposit balance of validator.
     */
    event Deposit(address validator, uint256 amount, uint256 balance);

    /**
     * @notice Emitted when the validator withdraws.
     *
     * @param validator The address of new validator.
     * @param amount    The withdrawal amount.
     * @param balance   The deposit balance of validator.
     */
    event Withdraw(address validator, uint256 amount, uint256 balance);

    /**
     * @notice Emitted when the validator deposits is slashed.
     *
     * @param validator  The address of new validator.
     * @param challenger The address of challenger.
     * @param balance    The deposit balance of validator.
     */
    event Slashing(address validator, address challenger, uint256 balance);

    /**
     * @custom:semver 0.1.0
     *
     * @param _colosseumAddress The address of colosseum contract.
     * @param _trustedValidator The address of trusted validator.
     * @param _minDepositAmount The minimum deposit amount.
     * @param _slashingAmount   The slashing amount.
     */
    constructor(
        address _colosseumAddress,
        address _trustedValidator,
        uint256 _minDepositAmount,
        uint256 _slashingAmount
    ) Semver(0, 1, 0) {
        COLOSSEUM_ADDRESS = _colosseumAddress;
        TRUSTED_VALIDATOR = _trustedValidator;
        MIN_DEPOSIT_AMOUNT = _minDepositAmount;
        SLASHING_AMOUNT = _slashingAmount;
        initialize();
    }

    /**
     * @notice Initializer.
     */
    function initialize() public initializer {}

    function deposit() external payable nonReentrant {
        uint256 newDeposits = depositsOf[msg.sender] + msg.value;
        require(msg.sender != TRUSTED_VALIDATOR, "ValidatorPool: trusted validator cannot deposit");
        require(
            newDeposits >= MIN_DEPOSIT_AMOUNT,
            "ValidatorPool: the total deposit is less than minimum amount"
        );

        depositsOf[msg.sender] = newDeposits;

        if (!validators[msg.sender]) {
            validators[msg.sender] = true;
            emit NewValidator(msg.sender);
        }

        emit Deposit(msg.sender, msg.value, newDeposits);
    }

    function withdraw(uint256 amount) external nonReentrant {
        uint256 deposits = depositsOf[msg.sender];
        require(deposits >= amount, "ValidatorPool: insufficient deposits");

        (bool success, ) = address(this).call{ value: amount }("");
        require(success, "ValidatorPool: failed to withdraw");

        emit Withdraw(msg.sender, amount, deposits - amount);
    }

    function slash(address validator, address challenger) external nonReentrant {
        require(
            msg.sender == COLOSSEUM_ADDRESS,
            "ValidatorPool: sender is not a colosseum contract"
        );
        require(
            validator != TRUSTED_VALIDATOR,
            "ValidatorPool: cannot slash deposits of trusted validator"
        );

        uint256 validatorDeposits = depositsOf[validator];
        uint256 slashing = SLASHING_AMOUNT;

        if (validatorDeposits < SLASHING_AMOUNT) {
            slashing = validatorDeposits - SLASHING_AMOUNT;
        }

        depositsOf[validator] -= slashing;
        depositsOf[challenger] += slashing;

        emit Slashing(validator, challenger, validatorDeposits);
    }

    /**
     * @notice Returns the deposit balance of validator.
     *
     * @param _validator The address of validator
     *
     * @return Deposit balance of validator.
     */
    function balanceOf(address _validator) public view returns (uint256) {
        return depositsOf[_validator];
    }

    function isValidator(address _validator) public view returns (bool) {
        return validators[_validator] || _validator == TRUSTED_VALIDATOR;
    }

    function isValidatorAlive(address _validator) public view returns (bool) {
        return depositsOf[_validator] > 0 || _validator == TRUSTED_VALIDATOR;
    }
}
