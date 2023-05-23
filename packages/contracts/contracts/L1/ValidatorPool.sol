// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {
    ReentrancyGuardUpgradeable
} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import { Types } from "../libraries/Types.sol";
import { Semver } from "../universal/Semver.sol";
import { L2OutputOracle } from "./L2OutputOracle.sol";

/**
 * @custom:proxied
 * @title ValidatorPool
 * @notice The ValidatorPool determines whether the validator is present and manages the validator's deposit.
 */
contract ValidatorPool is ReentrancyGuardUpgradeable, Semver {
    /**
     * @notice The address of the L2OutputOracle contract. Can be updated via upgrade.
     */
    L2OutputOracle public immutable L2_ORACLE;

    /**
     * @notice The address of the trusted validator. Can be updated via upgrade.
     */
    address public immutable TRUSTED_VALIDATOR;

    /**
     * @notice The minimum bond amount. Can be updated via upgrade.
     */
    uint256 public immutable MIN_BOND_AMOUNT;

    /**
     * @notice A mapping of balances.
     */
    mapping(address => uint256) internal balances;

    /**
     * @notice The bond corresponding to a specific output index.
     */
    mapping(uint256 => Types.Bond) internal bonds;

    /**
     * @notice The output index to unbond next.
     */
    uint256 internal nextUnbondOutputIndex;

    /**
     * @notice An array of validator addresses.
     */
    address[] internal validators;

    /**
     * @notice The index of the specific address in the validator array.
     */
    mapping(address => uint256) internal validatorIndexes;

    /**
     * @notice Emitted when a validator bonds.
     *
     * @param submitter   Address of submitter.
     * @param outputIndex Index of the L2 checkpoint output index.
     * @param amount      Amount of bonded.
     * @param expiresAt   The expiration timestamp of bond.
     */
    event Bonded(
        address indexed submitter,
        uint256 indexed outputIndex,
        uint128 amount,
        uint128 expiresAt
    );

    /**
     * @notice Emitted when the bond amount is increased.
     *
     * @param challenger  Address of the challenger.
     * @param outputIndex Index of the L2 checkpoint output.
     * @param amount      Amount of bond increased.
     */
    event BondIncreased(address indexed challenger, uint256 indexed outputIndex, uint128 amount);

    /**
     * @notice Emitted when a validator unbonds.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param recipient   Address of the recipient.
     * @param amount      Amount of unbonded.
     */
    event Unbonded(uint256 indexed outputIndex, address indexed recipient, uint128 amount);

    /**
     * @custom:semver 0.1.0
     *
     * @param _l2OutputOracle   Address of the L2OutputOracle.
     * @param _trustedValidator Address of the trusted validator.
     * @param _minBondAmount    The minimum bond amount.
     */
    constructor(
        L2OutputOracle _l2OutputOracle,
        address _trustedValidator,
        uint256 _minBondAmount
    ) Semver(0, 1, 0) {
        L2_ORACLE = _l2OutputOracle;
        TRUSTED_VALIDATOR = _trustedValidator;
        MIN_BOND_AMOUNT = _minBondAmount;
        initialize();
    }

    /**
     * @notice Initializer.
     */
    function initialize() public initializer {}

    /**
     * @notice Deposit ETH to be used as bond.
     */
    function deposit() external payable {
        _increaseBalance(msg.sender, msg.value);
    }

    /**
     * @notice Withdraw a given amount.
     *
     * @param _amount Amount to withdraw.
     */
    function withdraw(uint256 _amount) external nonReentrant {
        _decreaseBalance(msg.sender, _amount);
        (bool success, ) = payable(msg.sender).call{ value: _amount }("");
        require(success);
    }

    /**
     * @notice Bond asset corresponding to the given output index. This function is called when submitting output.
     *         Note that the amount to bond should be at least the minimum bond amount,
     *         and the output root must also be given for later output root comparison.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     * @param _amount      Amount of bond.
     * @param _expiresAt   The expiration timestamp of bond.
     */
    function createBond(
        uint256 _outputIndex,
        uint128 _amount,
        uint128 _expiresAt
    ) external {
        require(msg.sender == address(L2_ORACLE), "ValidatorPool: sender is not L2OutputOracle");
        require(_amount >= MIN_BOND_AMOUNT, "ValidatorPool: the bond amount is too small");

        Types.Bond storage bond = bonds[_outputIndex];
        require(
            bond.expiresAt == 0,
            "ValidatorPool: bond of the given output index already exists"
        );

        // Unbond the bond of nextUnbondOutputIndex if available.
        _tryUnbond();

        address submitter = L2_ORACLE.getSubmitter(_outputIndex);
        _decreaseBalance(submitter, _amount);

        bond.amount = _amount;
        bond.expiresAt = _expiresAt;

        emit Bonded(submitter, _outputIndex, _amount, _expiresAt);
    }

    /**
     * @notice Increases the bond amount corresponding to the given output index.
     *
     * @param _challenger  Address of the challenger.
     * @param _outputIndex Index of the L2 checkpoint output.
     */
    function increaseBond(address _challenger, uint256 _outputIndex) external {
        Types.Bond storage bond = bonds[_outputIndex];
        require(bond.expiresAt > 0, "ValidatorPool: the bond does not exist");

        uint128 bonded = bond.amount;
        _decreaseBalance(_challenger, bonded);
        bond.amount = bonded << 1;

        emit BondIncreased(_challenger, _outputIndex, bonded);
    }

    /**
     * @notice Attempt to unbond. Reverts if unbond is not possible.
     */
    function unbond() external {
        bool released = _tryUnbond();
        require(released, "ValidatorPool: no bond that can be unbond");
    }

    /**
     * @notice Attempts to unbond corresponding to nextUnbondOutputIndex and returns whether the unbond was successful.
     *
     * @return Whether the bond has been successfully unbonded.
     */
    function _tryUnbond() private returns (bool) {
        uint256 outputIndex = nextUnbondOutputIndex;
        address recipient = L2_ORACLE.getSubmitter(outputIndex);

        Types.Bond storage bond = bonds[outputIndex];
        if (block.timestamp >= bond.expiresAt && bond.amount > 0) {
            _increaseBalance(recipient, bond.amount);
            nextUnbondOutputIndex = nextUnbondOutputIndex + 1;
            emit Unbonded(outputIndex, recipient, bond.amount);
            bond.amount = 0;

            return true;
        }

        return false;
    }

    /**
     * @notice Increases the balance of the given address, If the increased balance is at lease
     *         the minimum bond amount, add the given address to the validator set.
     *
     * @param _validator Address to increase the balance.
     * @param _amount    Amount of balance increased.
     */
    function _increaseBalance(address _validator, uint256 _amount) private {
        uint256 balance = balances[_validator] + _amount;

        if (balance >= MIN_BOND_AMOUNT && !isValidator(_validator)) {
            validatorIndexes[_validator] = validators.length;
            validators.push(_validator);
        }

        balances[_validator] = balance;
    }

    /**
     * @notice Deceases the balance of the given address. If the decreased balance is less than
     *         the minimum bond amount, remove the given address from the validator set.
     *
     * @param _validator Address to decrease the balance.
     * @param _amount    Amount of balance decreased.
     */
    function _decreaseBalance(address _validator, uint256 _amount) private {
        uint256 balance = balances[_validator];
        require(balance >= _amount, "ValidatorPool: insufficient balances");
        balance = balance - _amount;

        if (balance < MIN_BOND_AMOUNT && isValidator(_validator)) {
            uint256 lastValidatorIndex = validators.length - 1;
            if (lastValidatorIndex > 0) {
                uint256 validatorIndex = validatorIndexes[_validator];
                address lastValidator = validators[lastValidatorIndex];

                validators[validatorIndex] = lastValidator;
                validatorIndexes[lastValidator] = validatorIndex;
            }
            delete validatorIndexes[_validator];
            validators.pop();
        }

        balances[_validator] = balance;
    }

    /**
     * @notice Returns the bond corresponding to the output index. Reverts if the bond does not exist.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     *
     * @return The bond data.
     */
    function getBond(uint256 _outputIndex) external view returns (Types.Bond memory) {
        Types.Bond storage bond = bonds[_outputIndex];
        require(bond.amount > 0 && bond.expiresAt > 0, "ValidatorPool: the bond does not exist");
        return bond;
    }

    /**
     * @notice Returns the balance of given address.
     *
     * @param _addr Address of validator.
     *
     * @return Balance of given address.
     */
    function balanceOf(address _addr) external view returns (uint256) {
        return balances[_addr];
    }

    /**
     * @notice Determines whether the given address is an active validator.
     *
     * @param _addr Address of validator.
     *
     * @return Whether the given address is an active validator.
     */
    function isValidator(address _addr) public view returns (bool) {
        if (validators.length == 0) {
            return false;
        } else if (_addr == address(0)) {
            return false;
        }

        uint256 index = validatorIndexes[_addr];
        return validators[index] == _addr;
    }

    /**
     * @notice Returns the number of validators.
     *
     * @return The number of validators.
     */
    function validatorCount() external view returns (uint256) {
        return validators.length;
    }

    /**
     * @notice Determines who can submit the L2 output next.
     *
     * @return The address of the validator.
     */
    function nextValidator() public view returns (address) {
        uint256 nextOutputIndex = L2_ORACLE.nextOutputIndex();
        if (nextOutputIndex > 0) {
            // TODO(pangssu): Make sure to use the finalized output root.
            Types.CheckpointOutput memory output = L2_ORACLE.getL2Output(nextOutputIndex - 1);
            uint256 rand = uint256(
                keccak256(
                    abi.encodePacked(
                        validators.length,
                        output.outputRoot,
                        output.submitter,
                        output.timestamp,
                        output.l2BlockNumber
                    )
                )
            ) % validators.length;

            return validators[rand];
        } else {
            return TRUSTED_VALIDATOR;
        }
    }
}
