// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {
    ReentrancyGuardUpgradeable
} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import { Math } from "@openzeppelin/contracts/utils/math/Math.sol";

import { Constants } from "../libraries/Constants.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { SafeCall } from "../libraries/SafeCall.sol";
import { Types } from "../libraries/Types.sol";
import { Semver } from "../universal/Semver.sol";
import { ValidatorRewardVault } from "../L2/ValidatorRewardVault.sol";
import { KromaPortal } from "./KromaPortal.sol";
import { L2OutputOracle } from "./L2OutputOracle.sol";

/**
 * @custom:proxied
 * @title ValidatorPool
 * @notice The ValidatorPool determines whether the validator is present and manages the validator's deposit.
 */
contract ValidatorPool is ReentrancyGuardUpgradeable, Semver {
    /**
     * @notice The gas limit to use when rewarding validator in the ValidatorRewardVault on L2.
     *         This value is measured through simulation.
     */
    uint64 public constant VAULT_REWARD_GAS_LIMIT = 100000;

    /**
     * @notice The address of the L2OutputOracle contract. Can be updated via upgrade.
     */
    L2OutputOracle public immutable L2_ORACLE;

    /**
     * @notice The address of the KromaPortal contract. Can be updated via upgrade.
     */
    KromaPortal public immutable PORTAL;

    /**
     * @notice The address of the trusted validator. Can be updated via upgrade.
     */
    address public immutable TRUSTED_VALIDATOR;

    /**
     * @notice The minimum bond amount. Can be updated via upgrade.
     */
    uint256 public immutable MIN_BOND_AMOUNT;

    /**
     * @notice The max number of unbonds when trying unbond.
     */
    uint256 public immutable MAX_UNBOND;

    /**
     * @notice The period in a round that the penalty does not apply, after which it does (in seconds).
     */
    uint256 public immutable NON_PENALTY_PERIOD;

    /**
     * @notice The period in a round that the penalty does apply (in seconds).
     */
    uint256 public immutable PENALTY_PERIOD;

    /**
     * @notice The duration of a submission round for one output (in seconds).
     *         Note that there are two submission rounds for an output: PRIORITY ROUND and PUBLIC ROUND.
     */
    uint256 public immutable ROUND_DURATION;

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
     * @notice Address of the next validator with priority for submitting output.
     */
    address internal nextPriorityValidator;

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
     * @param _portal           Address of the KromaPortal.
     * @param _trustedValidator Address of the trusted validator.
     * @param _minBondAmount    The minimum bond amount.
     * @param _maxUnbond        The max number of unbonds when trying unbond.
     * @param _nonPenaltyPeriod The period during a submission round that is not penalized.
     * @param _penaltyPeriod    The period during a submission round when penalties are applied.
     */
    constructor(
        L2OutputOracle _l2OutputOracle,
        KromaPortal _portal,
        address _trustedValidator,
        uint256 _minBondAmount,
        uint256 _maxUnbond,
        uint256 _nonPenaltyPeriod,
        uint256 _penaltyPeriod
    ) Semver(0, 1, 0) {
        L2_ORACLE = _l2OutputOracle;
        PORTAL = _portal;
        TRUSTED_VALIDATOR = _trustedValidator;
        MIN_BOND_AMOUNT = _minBondAmount;
        MAX_UNBOND = _maxUnbond;
        NON_PENALTY_PERIOD = _nonPenaltyPeriod;
        PENALTY_PERIOD = _penaltyPeriod;

        // The submission round duration is the sum of the non-penalty time and the penalty time.
        // Note that this value MUST NOT exceed (SUBMISSION_INTERVAL * L2_BLOCK_TIME) / 2.
        ROUND_DURATION = _nonPenaltyPeriod + _penaltyPeriod;

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

        bool success = SafeCall.call(msg.sender, gasleft(), _amount, "");
        require(success, "ValidatorPool: ETH transfer failed");
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
     * @notice Attempts to unbond starting from nextUnbondOutputIndex and returns whether at least
     *         one unbond is executed. Tries unbond at most MAX_UNBOND number of bonds and sends
     *         a reward message to L2 for each unbond.
     *         Note that it updates the next priority validator using last unbond, and not updates
     *         when no unbond.
     *
     * @return Whether at least one unbond is executed.
     */
    function _tryUnbond() private returns (bool) {
        uint256 outputIndex = nextUnbondOutputIndex;
        uint128 bondAmount;
        Types.Bond storage bond;
        Types.CheckpointOutput memory output;

        uint256 unbondedNum = 0;
        for (; unbondedNum < MAX_UNBOND; ) {
            bond = bonds[outputIndex];
            bondAmount = bond.amount;

            if (block.timestamp >= bond.expiresAt && bondAmount > 0) {
                delete bonds[outputIndex];
                output = L2_ORACLE.getL2Output(outputIndex);
                _increaseBalance(output.submitter, bondAmount);
                emit Unbonded(outputIndex, output.submitter, bondAmount);

                // Send reward message to L2 ValidatorRewardVault.
                _sendRewardMessageToL2Vault(output);

                unchecked {
                    ++unbondedNum;
                    ++outputIndex;
                }
            } else {
                break;
            }
        }

        if (unbondedNum > 0) {
            // Select the next priority validator.
            _updatePriorityValidator(output.outputRoot);

            unchecked {
                nextUnbondOutputIndex = outputIndex;
            }
            return true;
        }

        return false;
    }

    /**
     * @notice Updates next priority validator address.
     *
     * @param _outputRoot The L2 output of the checkpoint block.
     */
    function _updatePriorityValidator(bytes32 _outputRoot) private {
        uint256 len = validators.length;
        if (len > 0) {
            // TODO(pangssu): improve next validator selection
            uint256 validatorIndex = uint256(
                keccak256(abi.encodePacked(_outputRoot, block.number, block.coinbase))
            ) % len;

            nextPriorityValidator = validators[validatorIndex];
        } else {
            nextPriorityValidator = address(0);
        }
    }

    /**
     * @notice Sends reward message to ValidatorRewardVault contract on L2 using Portal.
     *
     * @param _output The finalized output.
     */
    function _sendRewardMessageToL2Vault(Types.CheckpointOutput memory _output) private {
        // Calculate the elapsed time to submitting the output.
        uint256 penaltyNumerator = 0;
        uint256 elapsed = _output.timestamp -
            L2_ORACLE.computeL2Timestamp(_output.l2BlockNumber + 1);
        if (elapsed > ROUND_DURATION) {
            elapsed -= ROUND_DURATION;
        }
        // Count from the time the penalty is applied.
        if (elapsed >= NON_PENALTY_PERIOD) {
            penaltyNumerator = Math.min(elapsed - NON_PENALTY_PERIOD, PENALTY_PERIOD);
        }

        // Pay out rewards via L2 Vault now that the output is finalized.
        PORTAL.depositTransactionByValidatorPool(
            Predeploys.VALIDATOR_REWARD_VAULT,
            VAULT_REWARD_GAS_LIMIT,
            abi.encodeWithSelector(
                ValidatorRewardVault.reward.selector,
                _output.submitter,
                _output.l2BlockNumber,
                penaltyNumerator,
                PENALTY_PERIOD
            )
        );
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
        if (nextPriorityValidator != address(0)) {
            uint256 l2BlockNumber = L2_ORACLE.nextBlockNumber();
            uint256 l2Timestamp = L2_ORACLE.computeL2Timestamp(l2BlockNumber + 1);
            if (block.timestamp >= l2Timestamp) {
                uint256 elapsed = block.timestamp - l2Timestamp;
                // If the current time exceeds one round time, it is a public round.
                if (elapsed > ROUND_DURATION) {
                    return Constants.VALIDATOR_PUBLIC_ROUND_ADDRESS;
                }
            }

            return nextPriorityValidator;
        } else {
            return TRUSTED_VALIDATOR;
        }
    }
}
