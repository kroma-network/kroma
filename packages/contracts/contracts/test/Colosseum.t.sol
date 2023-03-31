// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Math } from "@openzeppelin/contracts/utils/math/Math.sol";

import { Types } from "../libraries/Types.sol";
import { Colosseum } from "../L1/Colosseum.sol";
import { Colosseum_Initializer } from "./CommonTest.t.sol";
import { MockProofData } from "./MockProofData.sol";

// Test the implementations of the Colosseum
contract Colosseum_Test is Colosseum_Initializer {
    mapping(uint256 => bytes32) private _actualOutputs;
    mapping(uint256 => bytes32) private _invalidOutputs;

    uint256 startBlockNumber = 0;

    function _toOutputRoot(
        uint256 _blockNumber,
        bytes32 _prevOutput
    ) private pure returns (bytes32) {
        return keccak256(abi.encodePacked(_blockNumber, _prevOutput));
    }

    function setUp() public virtual override {
        super.setUp();

        uint256 interval = oracle.SUBMISSION_INTERVAL();
        startBlockNumber = oracle.latestBlockNumber();
        uint256 nextBlockNumber = startBlockNumber + interval;

        uint256 invalidIndex = uint256(keccak256(abi.encodePacked(block.timestamp))) % interval;

        for (uint256 i = 1; i <= interval * 2; i++) {
            uint256 blockNumber = startBlockNumber + i;

            _actualOutputs[blockNumber] = _toOutputRoot(
                blockNumber,
                _actualOutputs[blockNumber - 1]
            );

            if (i == (interval + invalidIndex)) {
                // insert invalid output root
                _invalidOutputs[blockNumber] = _toOutputRoot(blockNumber, bytes32(0));
            } else {
                _invalidOutputs[blockNumber] = _toOutputRoot(
                    blockNumber,
                    _invalidOutputs[blockNumber - 1]
                );
            }
        }

        // Roll to after the block number we'll submit
        warpToSubmitTime(nextBlockNumber);
        vm.prank(asserter);
        oracle.submitL2Output(_actualOutputs[nextBlockNumber], nextBlockNumber, 0, 0);

        // submit again
        startBlockNumber = oracle.latestBlockNumber();
        nextBlockNumber = startBlockNumber + interval;
        warpToSubmitTime(nextBlockNumber);
        vm.prank(asserter);
        oracle.submitL2Output(_invalidOutputs[nextBlockNumber], nextBlockNumber, 0, 0);
    }

    function _newSegments(
        address _sender,
        uint256 _turn,
        uint256 _segStart,
        uint256 _segSize
    ) private view returns (bytes32[] memory) {
        uint256 segLen = colosseum.getSegmentsLength(_turn);

        bytes32[] memory arr = new bytes32[](segLen);

        for (uint256 i = 0; i < segLen; i++) {
            uint256 n = _segStart + i * (_segSize / (segLen - 1));
            arr[i] = _sender == asserter ? _invalidOutputs[n] : _actualOutputs[n];
        }

        return arr;
    }

    function _detectFault(
        Types.Challenge memory _challenge,
        address _sender
    ) private view returns (uint256) {
        if (_sender == challenger && _challenge.current == _sender) {
            return 0;
        }

        uint256 segLen = colosseum.getSegmentsLength(_challenge.turn);
        uint256 start = _challenge.segStart;
        uint256 degree = _challenge.segSize / (segLen - 1);
        uint256 current = start + degree;

        for (uint256 i = 1; i < segLen; i++) {
            bytes32 output = _sender == asserter
                ? _invalidOutputs[current]
                : _actualOutputs[current];

            if (_challenge.segments[i] != output) {
                return i - 1;
            }

            current += degree;
        }

        revert("failed to select");
    }

    function _createChallenge() private returns (uint256) {
        uint256 challengeId = colosseum.latestChallengeId();

        uint256 end = oracle.latestBlockNumber();
        uint256 start = end - oracle.SUBMISSION_INTERVAL();
        uint256 outputIndex = oracle.getL2OutputIndexAfter(end);
        Types.CheckpointOutput memory targetOutput = oracle.getL2Output(outputIndex);

        assertTrue(_actualOutputs[end] != targetOutput.outputRoot, "not an invalid output");

        bytes32[] memory segments = _newSegments(challenger, 1, start, end - start);

        vm.prank(challenger);
        colosseum.createChallenge(outputIndex, segments);

        // check if challengeId is increased
        assertEq(colosseum.latestChallengeId(), challengeId + 1);
        challengeId = challengeId + 1;

        Types.Challenge memory challenge = colosseum.getChallengeInProgress();

        assertEq(challenge.outputIndex, outputIndex);
        assertEq(challenge.current, challenger);
        assertEq(challenge.next, oracle.VALIDATOR());
        assertEq(challenge.timeoutAt, block.timestamp + colosseum.CHALLENGE_TIMEOUT());
        assertEq(challenge.segments.length, colosseum.getSegmentsLength(1));
        assertEq(challenge.segStart, start);
        assertEq(challenge.segSize, end - start);
        assertEq(challenge.turn, 1);

        return challengeId;
    }

    function _bisect(address _sender) private {
        Types.Challenge memory challenge = colosseum.getChallengeInProgress();

        uint256 position = _detectFault(challenge, _sender);
        uint256 segSize = challenge.segSize / (colosseum.getSegmentsLength(challenge.turn) - 1);
        uint256 segStart = challenge.segStart + position * segSize;

        bytes32[] memory segments = _newSegments(_sender, challenge.turn + 1, segStart, segSize);

        vm.prank(_sender);
        colosseum.bisect(position, segments);

        Types.Challenge memory newChallenge = colosseum.getChallengeInProgress();
        assertEq(newChallenge.turn, challenge.turn + 1);
        assertEq(newChallenge.current, challenge.next);
        assertEq(newChallenge.next, challenge.current);
        assertEq(newChallenge.segments.length, segments.length);
        assertEq(newChallenge.segStart, segStart);
        assertEq(newChallenge.segSize, segSize);
    }

    function _proveFault() private {
        Types.Challenge memory challenge = colosseum.getChallengeInProgress();

        uint256 position = _detectFault(challenge, challenger);
        uint256 segSize = challenge.segSize / (colosseum.getSegmentsLength(challenge.turn) - 1);
        uint256 segStart = challenge.segStart + position * segSize;

        vm.prank(challenger);
        colosseum.proveFault(
            position,
            _actualOutputs[segStart + 1],
            MockProofData.getProof(),
            MockProofData.getPair()
        );

        assertEq(uint256(colosseum.getStatusInProgress()), uint256(Colosseum.ChallengeStatus.NONE));
        assertEq(colosseum.isInProgress(), false);
        assertEq(oracle.latestBlockNumber(), startBlockNumber);
    }

    function test_constructor() external {
        assertEq(address(colosseum.L2_ORACLE()), address(oracle), "oracle address not matched");
        assertEq(address(colosseum.ZK_VERIFIER()), address(zkVerifier), "zkVerifier address not matched");
        assertEq(colosseum.CHALLENGE_TIMEOUT(), timeout);
    }

    function test_createChallenge() external {
        _createChallenge();
    }

    function testCannot_createChallengeWithBadSegments() external {
        uint256 end = oracle.latestBlockNumber();
        uint256 start = end - oracle.SUBMISSION_INTERVAL();
        uint256 outputIndex = oracle.getL2OutputIndexAfter(end);
        uint256 segLen = colosseum.getSegmentsLength(1);

        vm.startPrank(challenger);

        // invalid segments length
        vm.expectRevert("Colosseum: invalid segments length");
        colosseum.createChallenge(outputIndex, new bytes32[](2));

        bytes32[] memory segments = new bytes32[](segLen);

        // invalid output of first segment
        for (uint256 i = 0; i < segments.length; i++) {
            segments[i] = keccak256(abi.encodePacked("wrong hash", i));
        }
        vm.expectRevert("Colosseum: the first segment must be matched");
        colosseum.createChallenge(outputIndex, segments);

        // invalid output of last segments
        segments[0] = _actualOutputs[start];
        segments[segments.length - 1] = oracle.getL2OutputAfter(end).outputRoot;
        vm.expectRevert("Colosseum: the last segment must not be matched");
        colosseum.createChallenge(outputIndex, segments);

        vm.stopPrank();
    }

    function testCannot_createChallengeWhenPrevInProgress() external {
        _createChallenge();

        uint256 end = oracle.latestBlockNumber();
        uint256 start = end - oracle.SUBMISSION_INTERVAL();
        uint256 outputIndex = oracle.getL2OutputIndexAfter(end);

        bytes32[] memory segments = _newSegments(challenger, 1, start, end - start);

        vm.prank(challenger);
        vm.expectRevert("Colosseum: previous challenge is in progress");
        colosseum.createChallenge(outputIndex, segments);
    }

    function test_bisect() external {
        _createChallenge();

        Types.Challenge memory challenge = colosseum.getChallengeInProgress();

        assertEq(colosseum.isInProgress(), true);
        assertEq(challenge.next, asserter);

        _bisect(asserter);
    }

    function test_bisectIfNotYourTurn() external {
        _createChallenge();
        Types.Challenge memory challenge = colosseum.getChallengeInProgress();

        assertEq(colosseum.isInProgress(), true);
        assertEq(challenge.next, asserter);

        vm.prank(challenger);
        vm.expectRevert("Colosseum: not your turn");
        colosseum.bisect(0, new bytes32[](1));
    }

    function test_asserterBisectAfterTimedOut() external {
        _createChallenge();
        Types.Challenge memory challenge = colosseum.getChallengeInProgress();

        assertEq(colosseum.isInProgress(), true);
        assertEq(challenge.next, asserter);

        uint256 segLen = colosseum.getSegmentsLength(challenge.turn + 1);

        vm.warp(challenge.timeoutAt + 1);
        vm.prank(asserter);
        vm.expectRevert("Colosseum: not your turn");
        colosseum.bisect(0, new bytes32[](segLen));

        assertEq(
            uint256(colosseum.getStatusInProgress()),
            uint256(Colosseum.ChallengeStatus.ASSERTER_TIMEOUT)
        );
    }

    function test_challengerBisectAfterTimedOut() external {
        _createChallenge();
        Types.Challenge memory challenge = colosseum.getChallengeInProgress();

        assertEq(colosseum.isInProgress(), true);
        assertEq(challenge.next, asserter);

        _bisect(asserter);

        // update challenge
        challenge = colosseum.getChallengeInProgress();

        uint256 segLen = colosseum.getSegmentsLength(challenge.turn + 1);

        vm.warp(challenge.timeoutAt + 1);
        vm.prank(challenger);
        vm.expectRevert("Colosseum: not your turn");
        colosseum.bisect(0, new bytes32[](segLen));

        assertEq(
            uint256(colosseum.getStatusInProgress()),
            uint256(Colosseum.ChallengeStatus.CHALLENGER_TIMEOUT)
        );
    }

    function test_proveFault() external {
        _createChallenge();

        while (colosseum.isAbleToBisect()) {
            Types.Challenge memory challenge = colosseum.getChallengeInProgress();
            _bisect(challenge.next);
        }

        assertEq(
            uint256(colosseum.getStatusInProgress()),
            uint256(Colosseum.ChallengeStatus.READY_TO_PROVE)
        );
        assertEq(colosseum.isInProgress(), true);
        assertEq(colosseum.getChallengeInProgress().next, challenger);

        _proveFault();
    }

    function test_timeoutWhenTimedOutAsserterTurn() external {
        _createChallenge();
        Types.Challenge memory challenge = colosseum.getChallengeInProgress();

        assertEq(colosseum.isInProgress(), true);
        assertEq(challenge.next, asserter);

        vm.warp(challenge.timeoutAt + 1);
        // check the asserter timeout
        assertEq(colosseum.isInProgress(), true);
        assertEq(
            uint256(colosseum.getStatusInProgress()),
            uint256(Colosseum.ChallengeStatus.ASSERTER_TIMEOUT)
        );

        vm.prank(challenger);
        colosseum.asserterTimeout();
    }

    function test_timeoutWhenTimedOutChallengerTurn() external {
        uint256 challengeId = _createChallenge();
        Types.Challenge memory challenge = colosseum.getChallengeInProgress();

        assertEq(colosseum.isInProgress(), true);
        assertEq(challenge.next, asserter);

        _bisect(asserter);

        challenge = colosseum.getChallengeInProgress();
        vm.warp(challenge.timeoutAt + 1);
        // check the challenger timeout
        assertEq(colosseum.isInProgress(), false);
        assertEq(challenge.next, challenger);
        assertEq(
            uint256(colosseum.getStatusInProgress()),
            uint256(Colosseum.ChallengeStatus.CHALLENGER_TIMEOUT)
        );

        vm.prank(asserter);
        colosseum.challengerTimeout(challengeId);
    }

    function test_challengerNotCloseWhenAsserterTimeout() external {
        _createChallenge();
        Types.Challenge memory challenge = colosseum.getChallengeInProgress();

        assertEq(colosseum.isInProgress(), true);
        assertEq(challenge.next, asserter);

        vm.warp(challenge.timeoutAt + 1);
        // check the asserter timeout
        assertEq(colosseum.isInProgress(), true);
        assertEq(
            uint256(colosseum.getStatusInProgress()),
            uint256(Colosseum.ChallengeStatus.ASSERTER_TIMEOUT)
        );
        // then challenger do not anything

        vm.warp(challenge.timeoutAt + colosseum.CHALLENGE_TIMEOUT() + 1);
        // check the challenger timeout
        assertEq(colosseum.isInProgress(), false);
        assertEq(
            uint256(colosseum.getStatusInProgress()),
            uint256(Colosseum.ChallengeStatus.CHALLENGER_TIMEOUT)
        );
    }
}
