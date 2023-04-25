// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Semver } from "../universal/Semver.sol";

/**
 * @custom:proxied
 * @custom:predeploy 0x4200000000000000000000000000000000000002
 * @title L1Block
 * @notice The L1Block predeploy gives users access to information about the last known L1 block.
 *         Values within this contract are updated once per epoch (every L1 block) and can only be
 *         set by the "depositor" account, a special system address. Depositor account transactions
 *         are created by the protocol whenever we move to a new epoch.
 */
contract L1Block is Semver {
    /**
     * @notice Address of the special depositor account.
     */
    address public constant DEPOSITOR_ACCOUNT = 0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001;

    /**
     * @notice The latest L1 block number known by the L2 system.
     */
    uint64 public number;

    /**
     * @notice The latest L1 timestamp known by the L2 system.
     */
    uint64 public timestamp;

    /**
     * @notice The latest L1 basefee.
     */
    uint256 public basefee;

    /**
     * @notice The latest L1 blockhash.
     */
    bytes32 public hash;

    /**
     * @notice The number of L2 blocks in the same epoch.
     */
    uint64 public sequenceNumber;

    /**
     * @notice The versioned hash to authenticate the batcher by.
     */
    bytes32 public batcherHash;

    /**
     * @notice The overhead value applied to the L1 portion of the transaction
     *         fee.
     */
    uint256 public l1FeeOverhead;

    /**
     * @notice The scalar value applied to the L1 portion of the transaction fee.
     */
    uint256 public l1FeeScalar;

    /**
     * @notice The ratio to distribute transaction fees as validator reward. 4 decimal.
     */
    uint256 public validatorRewardRatio;

    /**
     * @custom:semver 0.1.0
     */
    constructor() Semver(0, 1, 0) {}

    /**
     * @notice Updates the L1 block values.
     *
     * @param _number         L1 blocknumber.
     * @param _timestamp      L1 timestamp.
     * @param _basefee        L1 basefee.
     * @param _hash           L1 blockhash.
     * @param _sequenceNumber Number of L2 blocks since epoch start.
     * @param _batcherHash    Versioned hash to authenticate batcher by.
     * @param _l1FeeOverhead  L1 fee overhead.
     * @param _l1FeeScalar    L1 fee scalar.
     * @param _vRewardRatio   Validator reward ratio.
     */
    function setL1BlockValues(
        uint64 _number,
        uint64 _timestamp,
        uint256 _basefee,
        bytes32 _hash,
        uint64 _sequenceNumber,
        bytes32 _batcherHash,
        uint256 _l1FeeOverhead,
        uint256 _l1FeeScalar,
        uint256 _vRewardRatio
    ) external {
        require(
            msg.sender == DEPOSITOR_ACCOUNT,
            "L1Block: only the depositor account can set L1 block values"
        );
        require(
            _vRewardRatio <= 10000,
            "L1Block: the max value of validation reward ratio has been exceeded"
        );

        number = _number;
        timestamp = _timestamp;
        basefee = _basefee;
        hash = _hash;
        sequenceNumber = _sequenceNumber;
        batcherHash = _batcherHash;
        l1FeeOverhead = _l1FeeOverhead;
        l1FeeScalar = _l1FeeScalar;
        validatorRewardRatio = _vRewardRatio;
    }
}
