// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Types } from "../../libraries/Types.sol";

interface IValidatorPool {
    event BondIncreased(uint256 indexed outputIndex, address indexed challenger, uint128 amount);
    event Bonded(
        address indexed submitter,
        uint256 indexed outputIndex,
        uint128 amount,
        uint128 expiresAt
    );
    event Initialized(uint8 version);
    event PendingBondAdded(uint256 indexed outputIndex, address indexed challenger, uint128 amount);
    event PendingBondReleased(
        uint256 indexed outputIndex,
        address indexed challenger,
        address indexed recipient,
        uint128 amount
    );
    event Unbonded(uint256 indexed outputIndex, address indexed recipient, uint128 amount);

    function L2_ORACLE() external view returns (address);

    function MAX_UNBOND() external view returns (uint256);

    function PORTAL() external view returns (address);

    function REQUIRED_BOND_AMOUNT() external view returns (uint128);

    function ROUND_DURATION() external view returns (uint256);

    function SECURITY_COUNCIL() external view returns (address);

    function TAX_DENOMINATOR() external view returns (uint128);

    function TAX_NUMERATOR() external view returns (uint128);

    function TERMINATE_OUTPUT_INDEX() external view returns (uint256);

    function TRUSTED_VALIDATOR() external view returns (address);

    function VAULT_REWARD_GAS_LIMIT() external view returns (uint64);

    function addPendingBond(uint256 _outputIndex, address _challenger) external;

    function balanceOf(address _addr) external view returns (uint256);

    function createBond(uint256 _outputIndex, uint128 _expiresAt) external;

    function deposit() external payable;

    function getBond(uint256 _outputIndex) external view returns (Types.Bond memory);

    function getPendingBond(
        uint256 _outputIndex,
        address _challenger
    ) external view returns (uint128);

    function increaseBond(uint256 _outputIndex, address _challenger) external;

    function initialize() external;

    function isTerminated(uint256 _outputIndex) external view returns (bool);

    function isValidator(address _addr) external view returns (bool);

    function nextValidator() external view returns (address);

    function releasePendingBond(
        uint256 _outputIndex,
        address _challenger,
        address _recipient
    ) external;

    function unbond() external;

    function validatorCount() external view returns (uint256);

    function version() external view returns (string memory);

    function withdraw(uint256 _amount) external;

    function withdrawTo(address _to, uint256 _amount) external;
}
