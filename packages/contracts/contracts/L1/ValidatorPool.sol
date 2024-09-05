// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {
    ReentrancyGuardUpgradeable
} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import { Constants } from "../libraries/Constants.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { SafeCall } from "../libraries/SafeCall.sol";
import { Types } from "../libraries/Types.sol";
import { ISemver } from "../universal/ISemver.sol";
import { ValidatorRewardVault } from "../L2/ValidatorRewardVault.sol";
import { KromaPortal } from "./KromaPortal.sol";
import { L2OutputOracle } from "./L2OutputOracle.sol";

/**
 * @custom:proxied
 * @title ValidatorPool
 * @notice The ValidatorPool determines whether the validator is present and manages the validator's deposit.
 */
contract ValidatorPool is ReentrancyGuardUpgradeable, ISemver {
    /**
     * @notice The gas limit to use when rewarding validator in the ValidatorRewardVault on L2.
     *         This value is measured through simulation.
     */
    uint64 public constant VAULT_REWARD_GAS_LIMIT = 100000;

    /**
     * @notice The numerator of the tax.
     */
    uint128 public constant TAX_NUMERATOR = 20;

    /**
     * @notice The denominator of the tax.
     */
    uint128 public constant TAX_DENOMINATOR = 100;

    /**
     * @notice The address of the L2OutputOracle contract. Can be updated via upgrade.
     */
    L2OutputOracle public immutable L2_ORACLE;

    /**
     * @notice The address of the KromaPortal contract. Can be updated via upgrade.
     */
    KromaPortal public immutable PORTAL;

    /**
     * @notice The address of the SecurityCouncil contract. Can be updated via upgrade.
     */
    address public immutable SECURITY_COUNCIL;

    /**
     * @notice The address of the trusted validator. Can be updated via upgrade.
     */
    address public immutable TRUSTED_VALIDATOR;

    /**
     * @notice The required bond amount. Can be updated via upgrade.
     */
    uint128 public immutable REQUIRED_BOND_AMOUNT;

    /**
     * @notice The max number of unbonds when trying unbond.
     */
    uint256 public immutable MAX_UNBOND;

    /**
     * @notice The duration of a submission round for one output (in seconds).
     *         Note that there are two submission rounds for an output: PRIORITY ROUND and PUBLIC ROUND.
     */
    uint256 public immutable ROUND_DURATION;

    /**
     * @notice The output index where ValidatorPool is terminated after. Validator system is
     *         upgraded to ValidatorManager.
     */
    uint256 public immutable TERMINATE_OUTPUT_INDEX;

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
     * @notice A mapping of pending bonds that have not yet been included in a bond.
     */
    mapping(uint256 => mapping(address => uint128)) internal pendingBonds;

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
     * @notice Emitted when the pending bond is added.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param challenger  Address of the challenger.
     * @param amount      Amount of bond added.
     */
    event PendingBondAdded(uint256 indexed outputIndex, address indexed challenger, uint128 amount);

    /**
     * @notice Emitted when the bond is increased.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param challenger  Address of the challenger.
     * @param amount      Amount of bond increased.
     */
    event BondIncreased(uint256 indexed outputIndex, address indexed challenger, uint128 amount);

    /**
     * @notice Emitted when the pending bond is released(refunded).
     *
     * @param outputIndex  Index of the L2 checkpoint output.
     * @param challenger   Address of the challenger.
     * @param recipient    Address to receive amount from a pending bond.
     * @param amount       Amount of bond released.
     */
    event PendingBondReleased(
        uint256 indexed outputIndex,
        address indexed challenger,
        address indexed recipient,
        uint128 amount
    );

    /**
     * @notice Emitted when a validator unbonds.
     *
     * @param outputIndex Index of the L2 checkpoint output.
     * @param recipient   Address of the recipient.
     * @param amount      Amount of unbonded.
     */
    event Unbonded(uint256 indexed outputIndex, address indexed recipient, uint128 amount);

    /**
     * @notice A modifier that only allows the Colosseum contract to call
     */
    modifier onlyColosseum() {
        require(msg.sender == L2_ORACLE.COLOSSEUM(), "ValidatorPool: sender is not Colosseum");
        _;
    }
    /**
     * @notice Semantic version.
     * @custom:semver 1.1.0
     */
    string public constant version = "1.1.0";

    /**
     * @notice Constructs the ValidatorPool contract.
     *
     * @param _l2OutputOracle       Address of the L2OutputOracle.
     * @param _portal               Address of the KromaPortal.
     * @param _securityCouncil      Address of the security council.
     * @param _trustedValidator     Address of the trusted validator.
     * @param _requiredBondAmount   The required bond amount.
     * @param _maxUnbond            The max number of unbonds when trying unbond.
     * @param _roundDuration        The duration of one submission round in seconds.
     * @param _terminateOutputIndex The output index where ValidatorPool is terminated after.
     */
    constructor(
        L2OutputOracle _l2OutputOracle,
        KromaPortal _portal,
        address _securityCouncil,
        address _trustedValidator,
        uint256 _requiredBondAmount,
        uint256 _maxUnbond,
        uint256 _roundDuration,
        uint256 _terminateOutputIndex
    ) {
        L2_ORACLE = _l2OutputOracle;
        PORTAL = _portal;
        SECURITY_COUNCIL = _securityCouncil;
        TRUSTED_VALIDATOR = _trustedValidator;
        REQUIRED_BOND_AMOUNT = uint128(_requiredBondAmount);
        MAX_UNBOND = _maxUnbond;
        TERMINATE_OUTPUT_INDEX = _terminateOutputIndex;

        // Note that this value MUST be (SUBMISSION_INTERVAL * L2_BLOCK_TIME) / 2.
        ROUND_DURATION = _roundDuration;

        initialize();
    }

    /**
     * @notice Initializer.
     */
    function initialize() public initializer {
        __ReentrancyGuard_init_unchained();
    }

    /**
     * @notice Deposit ETH to be used as bond. Note that deposit after termination is not allowed.
     */
    function deposit() external payable {
        require(
            L2_ORACLE.nextOutputIndex() < TERMINATE_OUTPUT_INDEX + 1,
            "ValidatorPool: only can deposit to ValidatorPool before terminated"
        );

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
     * @notice Withdraw a given amount.
     *
     * @param _to     Address to withdraw asset to.
     * @param _amount Amount to withdraw.
     */
    function withdrawTo(address _to, uint256 _amount) external nonReentrant {
        require(_to != address(0), "ValidatorPool: cannot withdraw to the zero address");

        _decreaseBalance(msg.sender, _amount);

        bool success = SafeCall.call(_to, gasleft(), _amount, "");
        require(success, "ValidatorPool: ETH transfer failed");
    }

    /**
     * @notice Bond asset corresponding to the given output index.
     *         This function is called when submitting output.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     * @param _expiresAt   The expiration timestamp of bond.
     */
    function createBond(uint256 _outputIndex, uint128 _expiresAt) external {
        require(msg.sender == address(L2_ORACLE), "ValidatorPool: sender is not L2OutputOracle");

        Types.Bond storage bond = bonds[_outputIndex];
        require(
            bond.expiresAt == 0,
            "ValidatorPool: bond of the given output index already exists"
        );

        // Unbond the bond of nextUnbondOutputIndex if available.
        _tryUnbond();

        address submitter = L2_ORACLE.getSubmitter(_outputIndex);
        _decreaseBalance(submitter, REQUIRED_BOND_AMOUNT);

        bond.amount = REQUIRED_BOND_AMOUNT;
        bond.expiresAt = _expiresAt;

        emit Bonded(submitter, _outputIndex, REQUIRED_BOND_AMOUNT, _expiresAt);

        // Select the next priority validator
        _updatePriorityValidator();
    }

    /**
     * @notice Adds a pending bond to the challenge corresponding to the given output index and challenger address.
     *         The pending bond is added to the bond when the challenge is proven or challenger is timed out,
     *         or refunded when the challenge is canceled.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     * @param _challenger  Address of the challenger.
     */
    function addPendingBond(uint256 _outputIndex, address _challenger) external onlyColosseum {
        Types.Bond storage bond = bonds[_outputIndex];
        require(
            bond.expiresAt >= block.timestamp,
            "ValidatorPool: the output is already finalized"
        );

        _decreaseBalance(_challenger, REQUIRED_BOND_AMOUNT);
        pendingBonds[_outputIndex][_challenger] = REQUIRED_BOND_AMOUNT;

        emit PendingBondAdded(_outputIndex, _challenger, REQUIRED_BOND_AMOUNT);
    }

    /**
     * @notice Releases the corresponding pending bond to the given output index and challenger address
     *         if a challenge is canceled.
     *
     * @param _outputIndex  Index of the L2 checkpoint output.
     * @param _challenger   Address of the challenger.
     * @param _recipient    Address to receive amount from a pending bond.
     */
    function releasePendingBond(
        uint256 _outputIndex,
        address _challenger,
        address _recipient
    ) external onlyColosseum {
        uint128 bonded = pendingBonds[_outputIndex][_challenger];
        require(bonded > 0, "ValidatorPool: the pending bond does not exist");
        delete pendingBonds[_outputIndex][_challenger];

        _increaseBalance(_recipient, bonded);
        emit PendingBondReleased(_outputIndex, _challenger, _recipient, bonded);
    }

    /**
     * @notice Increases the bond amount corresponding to the given output index by the pending bond amount.
     *         This is when taxes are charged, and note that taxes are a means of preventing collusive attacks by
     *         the asserter and challenger.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     * @param _challenger  Address of the challenger.
     */
    function increaseBond(uint256 _outputIndex, address _challenger) external onlyColosseum {
        Types.Bond storage bond = bonds[_outputIndex];
        require(
            bond.expiresAt >= block.timestamp,
            "ValidatorPool: the output is already finalized"
        );

        uint128 pendingBond = pendingBonds[_outputIndex][_challenger];
        require(pendingBond > 0, "ValidatorPool: the pending bond does not exist");
        uint128 tax = (pendingBond * TAX_NUMERATOR) / TAX_DENOMINATOR;
        uint128 increased = pendingBond - tax;
        delete pendingBonds[_outputIndex][_challenger];

        unchecked {
            bond.amount += increased;
            balances[SECURITY_COUNCIL] += tax;
        }

        emit BondIncreased(_outputIndex, _challenger, increased);
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
            unchecked {
                nextUnbondOutputIndex = outputIndex;
            }

            // Set the next output index to be finalized in L2OutputOracle.
            L2_ORACLE.setNextFinalizeOutputIndex(outputIndex);

            return true;
        }

        return false;
    }

    /**
     * @notice Updates next priority validator address.
     */
    function _updatePriorityValidator() private {
        uint256 len = validators.length;
        if (len > 0 && nextUnbondOutputIndex > 0) {
            // TODO(pangssu): improve next validator selection
            Types.CheckpointOutput memory output = L2_ORACLE.getL2Output(nextUnbondOutputIndex - 1);
            uint256 validatorIndex = uint256(
                keccak256(
                    abi.encodePacked(
                        output.outputRoot,
                        block.number,
                        block.coinbase,
                        block.difficulty,
                        blockhash(block.number - 1)
                    )
                )
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
        // Pay out rewards via L2 Vault now that the output is finalized.
        PORTAL.depositTransactionByValidatorPool(
            Predeploys.VALIDATOR_REWARD_VAULT,
            VAULT_REWARD_GAS_LIMIT,
            abi.encodeWithSelector(
                ValidatorRewardVault.reward.selector,
                _output.submitter,
                _output.l2BlockNumber
            )
        );
    }

    /**
     * @notice Increases the balance of the given address. If the balance is greater than the required bond amount,
     *         add the given address to the validator set.
     *
     * @param _validator Address to increase the balance.
     * @param _amount    Amount of balance increased.
     */
    function _increaseBalance(address _validator, uint256 _amount) private {
        uint256 balance = balances[_validator] + _amount;

        if (balance >= REQUIRED_BOND_AMOUNT && !isValidator(_validator)) {
            if (_validator != SECURITY_COUNCIL) {
                validatorIndexes[_validator] = validators.length;
                validators.push(_validator);
            }
        }

        balances[_validator] = balance;
    }

    /**
     * @notice Deceases the balance of the given address. If the balance is less than the required bond amount,
     *         remove the given address from the validator set.
     *
     * @param _validator Address to decrease the balance.
     * @param _amount    Amount of balance decreased.
     */
    function _decreaseBalance(address _validator, uint256 _amount) private {
        uint256 balance = balances[_validator];
        require(balance >= _amount, "ValidatorPool: insufficient balances");
        balance = balance - _amount;

        if (balance < REQUIRED_BOND_AMOUNT && isValidator(_validator)) {
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
     * @notice Returns the pending bond corresponding to the output index and challenger address.
     *         Reverts if the pending bond does not exist.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     * @param _challenger  Address of the challenger.
     *
     * @return Amount of the pending bond.
     */
    function getPendingBond(
        uint256 _outputIndex,
        address _challenger
    ) external view returns (uint128) {
        uint128 pendingBond = pendingBonds[_outputIndex][_challenger];
        require(pendingBond > 0, "ValidatorPool: the pending bond does not exist");
        return pendingBond;
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
            uint256 l2Timestamp = L2_ORACLE.nextOutputMinL2Timestamp();
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

    /**
     * @notice Determines whether the given output index must not interact with ValidatorPool.
     *
     * @param _outputIndex Index of the L2 checkpoint output.
     *
     * @return Whether the given output index must not interact with ValidatorPool.
     */
    function isTerminated(uint256 _outputIndex) external view returns (bool) {
        return _outputIndex > TERMINATE_OUTPUT_INDEX;
    }
}
