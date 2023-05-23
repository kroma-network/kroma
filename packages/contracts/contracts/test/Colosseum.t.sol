// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Math } from "@openzeppelin/contracts/utils/math/Math.sol";

import { Types } from "../libraries/Types.sol";
import { Colosseum } from "../L1/Colosseum.sol";
import { Colosseum_Initializer } from "./CommonTest.t.sol";
import { ColosseumTestData } from "./testdata/ColosseumTestData.sol";

// Test the implementations of the Colosseum
contract ColosseumTest is Colosseum_Initializer {
    uint256 startBlockNumber = 0;

    function nextSender(Types.Challenge memory _challenge) internal pure returns (address) {
        return _challenge.turn % 2 == 0 ? _challenge.challenger : _challenge.asserter;
    }

    function setUp() public virtual override {
        super.setUp();

        vm.prank(trusted);
        pool.deposit{ value: trusted.balance }();
        vm.prank(asserter);
        pool.deposit{ value: asserter.balance }();

        uint256 interval = oracle.SUBMISSION_INTERVAL();
        startBlockNumber = oracle.latestBlockNumber();
        uint256 nextBlockNumber = startBlockNumber + interval;

        // Roll to after the block number we'll submit
        warpToSubmitTime(nextBlockNumber);
        vm.prank(pool.nextValidator());
        oracle.submitL2Output(bytes32(nextBlockNumber), nextBlockNumber, 0, 0, minBond);

        // Submit invalid output
        startBlockNumber = oracle.latestBlockNumber();
        nextBlockNumber = startBlockNumber + interval;
        warpToSubmitTime(nextBlockNumber);
        vm.prank(pool.nextValidator());
        oracle.submitL2Output(keccak256(abi.encode()), nextBlockNumber, 0, 0, minBond);

        vm.prank(challenger);
        pool.deposit{ value: challenger.balance }();
    }

    function _getOutputRoot(address _sender, uint256 _blockNumber) private view returns (bytes32) {
        uint256 targetBlockNumber = ColosseumTestData.INVALID_BLOCK_NUMBER;
        if (_blockNumber == targetBlockNumber - 1) {
            return ColosseumTestData.PREV_OUTPUT_ROOT;
        }

        if (_sender == challenger) {
            if (_blockNumber == targetBlockNumber) {
                return ColosseumTestData.TARGET_OUTPUT_ROOT;
            }
        } else if (_blockNumber >= targetBlockNumber) {
            return keccak256(abi.encode(_blockNumber));
        }

        return bytes32(_blockNumber);
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
            arr[i] = _getOutputRoot(_sender, n);
        }

        return arr;
    }

    function _detectFault(Types.Challenge memory _challenge, address _sender)
        private
        view
        returns (uint256)
    {
        if (_sender == _challenge.challenger && _sender != nextSender(_challenge)) {
            return 0;
        }

        uint256 segLen = colosseum.getSegmentsLength(_challenge.turn);
        uint256 start = _challenge.segStart;
        uint256 degree = _challenge.segSize / (segLen - 1);
        uint256 current = start + degree;

        for (uint256 i = 1; i < segLen; i++) {
            bytes32 output = _getOutputRoot(_sender, current);

            if (_challenge.segments[i] != output) {
                return i - 1;
            }

            current += degree;
        }

        revert("failed to select");
    }

    function _createChallenge(uint256 outputIndex) private returns (uint256) {
        uint256 end = oracle.latestBlockNumber();
        uint256 start = end - oracle.SUBMISSION_INTERVAL();
        Types.CheckpointOutput memory targetOutput = oracle.getL2Output(outputIndex);

        assertTrue(
            _getOutputRoot(targetOutput.submitter, end) != targetOutput.outputRoot,
            "not an invalid output"
        );

        bytes32[] memory segments = _newSegments(challenger, 1, start, end - start);

        vm.prank(challenger);
        colosseum.createChallenge(outputIndex, segments);

        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex);

        assertEq(challenge.asserter, targetOutput.submitter);
        assertEq(challenge.challenger, challenger);
        assertEq(challenge.timeoutAt, block.timestamp + colosseum.BISECTION_TIMEOUT());
        assertEq(challenge.segments.length, colosseum.getSegmentsLength(1));
        assertEq(challenge.segStart, start);
        assertEq(challenge.segSize, end - start);
        assertEq(challenge.turn, 1);

        return outputIndex;
    }

    function _bisect(uint256 _outputIndex, address _sender) private {
        Types.Challenge memory challenge = colosseum.getChallenge(_outputIndex);

        uint256 position = _detectFault(challenge, _sender);
        uint256 segSize = challenge.segSize / (colosseum.getSegmentsLength(challenge.turn) - 1);
        uint256 segStart = challenge.segStart + position * segSize;

        bytes32[] memory segments = _newSegments(_sender, challenge.turn + 1, segStart, segSize);

        vm.prank(_sender);
        colosseum.bisect(_outputIndex, position, segments);

        Types.Challenge memory newChallenge = colosseum.getChallenge(_outputIndex);
        assertEq(newChallenge.turn, challenge.turn + 1);
        assertEq(newChallenge.segments.length, segments.length);
        assertEq(newChallenge.segStart, segStart);
        assertEq(newChallenge.segSize, segSize);
    }

    function _proveFault(uint256 _outputIndex) private {
        Types.Challenge memory challenge = colosseum.getChallenge(_outputIndex);
        Types.CheckpointOutput memory output = oracle.getL2Output(_outputIndex);

        uint256 position = _detectFault(challenge, challenge.challenger);
        bytes32 newOutputRoot = _getOutputRoot(challenger, output.l2BlockNumber);
        assertTrue(newOutputRoot != output.outputRoot);

        (
            Types.OutputRootProof memory srcOutputRootProof,
            Types.OutputRootProof memory dstOutputRootProof
        ) = ColosseumTestData.outputRootProof();
        Types.PublicInput memory publicInput = ColosseumTestData.publicInput();
        Types.BlockHeaderRLP memory rlps = ColosseumTestData.blockHeaderRLP();

        bytes32 piHash = ColosseumTestData.publicInputHash(colosseum);
        ColosseumTestData.ProofPair memory pp = ColosseumTestData.proofAndPair(piHash);

        vm.prank(challenge.challenger);
        colosseum.proveFault(
            _outputIndex,
            newOutputRoot,
            position,
            srcOutputRootProof,
            dstOutputRootProof,
            publicInput,
            rlps,
            pp.proof,
            pp.pair
        );

        assertEq(
            uint256(colosseum.getStatus(_outputIndex)),
            uint256(Colosseum.ChallengeStatus.PROVEN)
        );
    }

    function test_constructor() external {
        assertEq(address(colosseum.L2_ORACLE()), address(oracle), "oracle address not matched");
        assertEq(
            address(colosseum.ZK_VERIFIER()),
            address(zkVerifier),
            "zk verifier address not matched"
        );
        assertEq(colosseum.CHAIN_ID(), CHAIN_ID);
        assertEq(colosseum.DUMMY_HASH(), DUMMY_HASH);
        assertEq(colosseum.MAX_TXS(), MAX_TXS);
        assertEq(colosseum.GUARDIAN(), guardian);
    }

    function test_createChallenge_succeeds() external {
        uint256 outputIndex = oracle.latestOutputIndex();
        _createChallenge(outputIndex);
    }

    function test_createChallenge_sameOutput_reverts() external {
        uint256 outputIndex = oracle.latestOutputIndex();
        _createChallenge(outputIndex);

        uint256 segLen = colosseum.getSegmentsLength(1);

        vm.expectRevert(
            "Colosseum: the challenge for given output index is already in progress or approved."
        );
        colosseum.createChallenge(outputIndex, new bytes32[](segLen));
    }

    function test_createChallenge_WithBadSegments_reverts() external {
        uint256 latestBlockNumber = oracle.latestBlockNumber();
        uint256 outputIndex = oracle.getL2OutputIndexAfter(latestBlockNumber);
        uint256 segLen = colosseum.getSegmentsLength(1);

        vm.startPrank(challenger);

        // invalid segments length
        vm.expectRevert("Colosseum: invalid segments length");
        colosseum.createChallenge(outputIndex, new bytes32[](segLen + 1));

        bytes32[] memory segments = new bytes32[](segLen);

        // invalid output of first segment
        for (uint256 i = 0; i < segments.length; i++) {
            segments[i] = keccak256(abi.encodePacked("wrong hash", i));
        }
        vm.expectRevert("Colosseum: the first segment must be matched");
        colosseum.createChallenge(outputIndex, segments);

        // invalid output of last segments
        segments[0] = oracle.getL2Output(outputIndex - 1).outputRoot;
        segments[segLen - 1] = oracle.getL2Output(outputIndex).outputRoot;
        vm.expectRevert("Colosseum: the last segment must not be matched");
        colosseum.createChallenge(outputIndex, segments);

        vm.stopPrank();
    }

    function test_createChallenge_ongoingChallenge_reverts() external {
        uint256 outputIndex = oracle.latestOutputIndex();
        _createChallenge(outputIndex);

        assertEq(
            uint256(colosseum.getStatus(outputIndex)),
            uint256(Colosseum.ChallengeStatus.ASSERTER_TURN)
        );

        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex);
        vm.prank(challenger);
        vm.expectRevert(
            "Colosseum: the challenge for given output index is already in progress or approved."
        );
        colosseum.createChallenge(outputIndex, challenge.segments);
    }

    function test_createChallenge_afterChallengeApproved_reverts() external {
        uint256 outputIndex = oracle.latestOutputIndex();
        test_approveChallenge_succeeds();

        assertEq(
            uint256(colosseum.getStatus(outputIndex)),
            uint256(Colosseum.ChallengeStatus.APPROVED)
        );

        uint256 segLen = colosseum.getSegmentsLength(1);

        vm.prank(challenger);
        vm.expectRevert(
            "Colosseum: the challenge for given output index is already in progress or approved."
        );
        colosseum.createChallenge(outputIndex, new bytes32[](segLen));
    }

    function test_createChallenge_afterChallengerTimedOut_succeeds() external {
        uint256 outputIndex = oracle.latestOutputIndex();
        _createChallenge(outputIndex);

        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex);

        _bisect(outputIndex, challenge.asserter);
        challenge = colosseum.getChallenge(outputIndex);
        vm.warp(challenge.timeoutAt + 1);

        assertEq(
            uint256(colosseum.getStatus(outputIndex)),
            uint256(Colosseum.ChallengeStatus.CHALLENGER_TIMEOUT)
        );

        _createChallenge(outputIndex);
        assertEq(
            uint256(colosseum.getStatus(outputIndex)),
            uint256(Colosseum.ChallengeStatus.ASSERTER_TURN)
        );
    }

    function test_createChallenge_x2BondAfterChallengerTimedOut_succeeds() external {
        uint256 outputIndex = oracle.latestOutputIndex();
        test_challengerTimeout_succeeds();

        Types.Bond memory bond = pool.getBond(outputIndex);
        assertEq(bond.amount, minBond * 2);

        _createChallenge(outputIndex);

        Types.Bond memory newBond = pool.getBond(outputIndex);
        assertEq(newBond.amount, bond.amount * 2);
    }

    function test_bisect_succeeds() external {
        uint256 outputIndex = oracle.latestOutputIndex();
        _createChallenge(outputIndex);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex);

        assertEq(colosseum.isInProgress(outputIndex), true);
        assertEq(nextSender(challenge), challenge.asserter);

        _bisect(outputIndex, challenge.asserter);
    }

    function test_bisect_ifNotYourTurn_reverts() external {
        uint256 outputIndex = oracle.latestOutputIndex();
        _createChallenge(outputIndex);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex);

        assertEq(colosseum.isInProgress(outputIndex), true);
        assertEq(nextSender(challenge), challenge.asserter);

        uint256 segLen = colosseum.getSegmentsLength(challenge.turn + 1);

        vm.prank(challenger);
        vm.expectRevert("Colosseum: not your turn");
        colosseum.bisect(outputIndex, 0, new bytes32[](segLen));
    }

    function test_bisect_whenAsserterTimedOut_reverts() external {
        uint256 outputIndex = oracle.latestOutputIndex();
        _createChallenge(outputIndex);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex);

        assertEq(colosseum.isInProgress(outputIndex), true);
        assertEq(nextSender(challenge), challenge.asserter);

        uint256 segLen = colosseum.getSegmentsLength(challenge.turn + 1);

        vm.warp(challenge.timeoutAt + 1);
        vm.prank(challenge.asserter);
        vm.expectRevert("Colosseum: not your turn");
        colosseum.bisect(outputIndex, 0, new bytes32[](segLen));

        assertEq(
            uint256(colosseum.getStatus(outputIndex)),
            uint256(Colosseum.ChallengeStatus.ASSERTER_TIMEOUT)
        );
    }

    function test_bisect_whenChallengerTimedOut() external {
        uint256 outputIndex = oracle.latestOutputIndex();
        _createChallenge(outputIndex);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex);

        assertEq(colosseum.isInProgress(outputIndex), true);
        assertEq(nextSender(challenge), challenge.asserter);

        _bisect(outputIndex, challenge.asserter);

        // update challenge
        challenge = colosseum.getChallenge(outputIndex);

        uint256 segLen = colosseum.getSegmentsLength(challenge.turn + 1);

        vm.warp(challenge.timeoutAt + 1);
        vm.prank(challenge.challenger);
        vm.expectRevert("Colosseum: not your turn");
        colosseum.bisect(outputIndex, 0, new bytes32[](segLen));

        assertEq(
            uint256(colosseum.getStatus(outputIndex)),
            uint256(Colosseum.ChallengeStatus.CHALLENGER_TIMEOUT)
        );
    }

    function test_proveFault_succeeds() public {
        uint256 outputIndex = oracle.latestOutputIndex();
        _createChallenge(outputIndex);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex);

        while (colosseum.isAbleToBisect(outputIndex)) {
            challenge = colosseum.getChallenge(outputIndex);
            _bisect(outputIndex, nextSender(challenge));
        }

        assertEq(
            uint256(colosseum.getStatus(outputIndex)),
            uint256(Colosseum.ChallengeStatus.READY_TO_PROVE)
        );
        assertEq(colosseum.isInProgress(outputIndex), true);

        _proveFault(outputIndex);
    }

    function test_proveFault_whenAsserterTimedOut_succeeds() external {
        uint256 outputIndex = oracle.latestOutputIndex();
        _createChallenge(outputIndex);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex);

        assertEq(colosseum.isInProgress(outputIndex), true);
        assertEq(nextSender(challenge), challenge.asserter);

        vm.warp(challenge.timeoutAt + 1);
        // check the asserter timeout
        assertEq(colosseum.isInProgress(outputIndex), true);
        assertEq(
            uint256(colosseum.getStatus(outputIndex)),
            uint256(Colosseum.ChallengeStatus.ASSERTER_TIMEOUT)
        );

        _proveFault(outputIndex);
    }

    function test_approveChallenge_succeeds() public {
        test_proveFault_succeeds();

        uint256 outputIndex = oracle.latestOutputIndex();
        Types.CheckpointOutput memory prevOutput = oracle.getL2Output(outputIndex);

        vm.prank(guardian);
        colosseum.approveChallenge(outputIndex);

        Types.CheckpointOutput memory newOutput = oracle.getL2Output(outputIndex);

        assertTrue(prevOutput.outputRoot != newOutput.outputRoot);
        assertEq(prevOutput.timestamp, newOutput.timestamp);
        assertEq(prevOutput.l2BlockNumber, newOutput.l2BlockNumber);
        assertEq(newOutput.submitter, challenger);
        assertEq(outputIndex, oracle.latestOutputIndex());
    }

    function test_approveChallenge_notGuardian_reverts() external {
        test_proveFault_succeeds();

        uint256 outputIndex = oracle.latestOutputIndex();

        vm.prank(trusted);
        vm.expectRevert("Colosseum: sender is not the guardian");
        colosseum.approveChallenge(outputIndex);
    }

    function test_approveChallenge_notProven_reverts() external {
        uint256 outputIndex = oracle.latestOutputIndex();

        vm.prank(guardian);
        vm.expectRevert("Colosseum: this challenge is not proven");
        colosseum.approveChallenge(outputIndex);
    }

    function test_challengerTimeout_succeeds() public {
        uint256 outputIndex = oracle.latestOutputIndex();
        _createChallenge(outputIndex);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex);

        assertEq(colosseum.isInProgress(outputIndex), true);
        assertEq(nextSender(challenge), challenge.asserter);

        _bisect(outputIndex, challenge.asserter);

        challenge = colosseum.getChallenge(outputIndex);
        vm.warp(challenge.timeoutAt + 1);
        // check the challenger timeout
        assertEq(colosseum.isInProgress(outputIndex), false);
        assertEq(nextSender(challenge), challenge.challenger);
        assertEq(
            uint256(colosseum.getStatus(outputIndex)),
            uint256(Colosseum.ChallengeStatus.CHALLENGER_TIMEOUT)
        );

        vm.prank(challenge.asserter);
        colosseum.challengerTimeout(outputIndex);
    }

    function test_challengerNotCloseWhenAsserterTimeout_succeeds() external {
        uint256 outputIndex = oracle.latestOutputIndex();
        _createChallenge(outputIndex);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex);

        assertEq(colosseum.isInProgress(outputIndex), true);
        assertEq(nextSender(challenge), challenge.asserter);

        vm.warp(challenge.timeoutAt + 1);
        // check the asserter timeout
        assertEq(colosseum.isInProgress(outputIndex), true);
        assertEq(
            uint256(colosseum.getStatus(outputIndex)),
            uint256(Colosseum.ChallengeStatus.ASSERTER_TIMEOUT)
        );
        // then challenger do not anything

        vm.warp(challenge.timeoutAt + colosseum.PROVING_TIMEOUT() + 1);
        // check the challenger timeout
        assertEq(colosseum.isInProgress(outputIndex), false);
        assertEq(
            uint256(colosseum.getStatus(outputIndex)),
            uint256(Colosseum.ChallengeStatus.CHALLENGER_TIMEOUT)
        );
    }
}
