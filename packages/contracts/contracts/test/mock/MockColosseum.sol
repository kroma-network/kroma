// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Colosseum } from "contracts/L1/Colosseum.sol";
import { L2OutputOracle } from "contracts/L1/L2OutputOracle.sol";
import { ZKProofVerifier } from "contracts/L1/ZKProofVerifier.sol";

contract MockColosseum is Colosseum {
    constructor(
        address _l2Oracle,
        ZKProofVerifier _zkProofVerifier,
        uint256 _submissionInterval,
        address _securityCouncil,
        uint256 _guardianPeriod,
        uint256 _maxClockDuration
    )
        Colosseum(
            _l2Oracle,
            _zkProofVerifier,
            _submissionInterval,
            _securityCouncil,
            _guardianPeriod,
            _maxClockDuration
        )
    {}

    function isAbleToBisect(
        uint256 _outputIndex,
        address _challenger
    ) external view returns (bool) {
        return _isAbleToBisect(challenges[_outputIndex][_challenger]);
    }

    function setL1Head(uint256 _outputIndex, address _challenger, bytes32 _l1Head) external {
        challenges[_outputIndex][_challenger].l1Head = _l1Head;
    }
}
