// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Types } from "../libraries/Types.sol";
import { ISP1Verifier } from "../vendor/ISP1Verifier.sol";
import { IValidatorManager } from "../L1/interfaces/IValidatorManager.sol";
import { Colosseum } from "../L1/Colosseum.sol";
import { L2OutputOracle } from "../L1/L2OutputOracle.sol";
import { ValidatorPool } from "../L1/ValidatorPool.sol";
import { ValidatorManager } from "../L1/ValidatorManager.sol";
import { ZKProofVerifier } from "../L1/ZKProofVerifier.sol";
import { ZKVerifier } from "../L1/ZKVerifier.sol";
import { Proxy } from "../universal/Proxy.sol";
import { MockColosseum } from "./mock/MockColosseum.sol";
import { ZkEvmTestData } from "./testdata/ZkEvmTestData.sol";
import { ZkVmTestData } from "./testdata/ZkVmTestData.sol";
import { Colosseum_Initializer } from "./CommonTest.t.sol";
import { MockL2OutputOracle, MockValidatorManager } from "./ValidatorManager.t.sol";

contract MockZKProofVerifier is ZKProofVerifier {
    constructor(
        ZKVerifier _zkVerifier,
        bytes32 _dummyHash,
        uint256 _maxTxs,
        address _zkMerkleTrie,
        ISP1Verifier _sp1Verifier,
        bytes32 _zkVmProgramVKey
    )
        ZKProofVerifier(
            _zkVerifier,
            _dummyHash,
            _maxTxs,
            _zkMerkleTrie,
            _sp1Verifier,
            _zkVmProgramVKey
        )
    {}

    function hashZkEvmPublicInput(
        Types.PublicInputProof calldata _proof
    ) external view returns (bytes32) {
        return _hashZkEvmPublicInput(_proof.srcOutputRootProof.stateRoot, _proof.publicInput);
    }

    function hashZkVmPublicInput(bytes calldata _publicValues) external pure returns (bytes32) {
        return keccak256(_publicValues);
    }
}

// Test the implementations of the Colosseum
contract ColosseumTest is Colosseum_Initializer {
    uint256 internal constant CHALLENGER_TURN = 1;
    uint256 internal constant ASSERTER_TURN = 2;
    uint256 internal constant READY_TO_PROVE = 3;

    MockColosseum mockColosseum;
    MockZKProofVerifier mockZKProofVerifier;
    uint256 internal targetOutputIndex;
    mapping(address => bool) internal isChallenger;

    event ReadyToProve(uint256 indexed outputIndex, address indexed challenger);
    event ChallengeCreated(
        uint256 indexed outputIndex,
        address indexed asserter,
        address indexed challenger,
        uint256 timestamp
    );
    event ChallengerTimedOut(
        uint256 indexed outputIndex,
        address indexed challenger,
        uint256 timestamp
    );
    event ChallengeCanceled(
        uint256 indexed outputIndex,
        address indexed challenger,
        uint256 timestamp
    );

    function nextSender(Types.Challenge memory _challenge) internal pure returns (address) {
        return _challenge.turn % 2 == 0 ? _challenge.challenger : _challenge.asserter;
    }

    function setUp() public virtual override {
        super.setUp();

        MockColosseum mockColosseumImpl = new MockColosseum(
            address(oracle),
            zkProofVerifier,
            submissionInterval,
            address(securityCouncil),
            guardianPeriod,
            maxClockDuration,
            challengeGracePeriod
        );
        vm.prank(multisig);
        Proxy(payable(address(colosseum))).upgradeTo(address(mockColosseumImpl));
        mockColosseum = MockColosseum(address(colosseum));
        colosseum = Colosseum(address(colosseum));

        MockZKProofVerifier mockVerifierImpl = new MockZKProofVerifier({
            _zkVerifier: zkVerifier,
            _dummyHash: DUMMY_HASH,
            _maxTxs: MAX_TXS,
            _zkMerkleTrie: address(zkMerkleTrie),
            _sp1Verifier: sp1Verifier,
            _zkVmProgramVKey: ZKVM_PROGRAM_V_KEY
        });
        vm.prank(multisig);
        Proxy(payable(address(zkProofVerifier))).upgradeTo(address(mockVerifierImpl));
        mockZKProofVerifier = MockZKProofVerifier(address(zkProofVerifier));
        vm.prank(trusted);
        pool.deposit{ value: trusted.balance }();
        vm.prank(asserter);
        pool.deposit{ value: asserter.balance }();

        // Submit genesis output
        uint256 nextBlockNumber = oracle.nextBlockNumber();
        // Roll to after the block number we'll submit
        warpToSubmitTime();
        vm.prank(pool.nextValidator());
        oracle.submitL2Output(bytes32(nextBlockNumber), nextBlockNumber, 0, 0);

        // Submit invalid output
        nextBlockNumber = oracle.nextBlockNumber();
        warpToSubmitTime();
        vm.prank(pool.nextValidator());
        oracle.submitL2Output(keccak256(abi.encode()), nextBlockNumber, 0, 0);

        vm.prank(challenger);
        pool.deposit{ value: challenger.balance }();
        isChallenger[challenger] = true;

        _createAssertion();

        targetOutputIndex = oracle.latestOutputIndex();
    }

    function _getOutputRoot(address _sender, uint256 _blockNumber) private view returns (bytes32) {
        uint256 targetBlockNumber;

        targetBlockNumber = ZkVmTestData.INVALID_BLOCK_NUMBER;

        if (_blockNumber == targetBlockNumber - 1) {
            return ZkVmTestData.PREV_OUTPUT_ROOT;
        }

        if (isChallenger[_sender]) {
            if (_blockNumber == targetBlockNumber) {
                return ZkVmTestData.TARGET_OUTPUT_ROOT;
            }
        } else if (_blockNumber >= targetBlockNumber) {
            return keccak256(abi.encode(_blockNumber));
        }

        return bytes32(_blockNumber);
    }

    function _detectFault(
        Types.Challenge memory _challenge,
        address _sender
    ) private view returns (uint256) {
        uint256 start;
        uint256 end;
        if (_challenge.segment.output == _getOutputRoot(_sender, _challenge.segment.pos)) {
            start = _challenge.segment.pos;
            end = _challenge.segment.end;
        } else {
            start = _challenge.segment.start;
            end = _challenge.segment.pos;
        }

        if (start + 1 == end) {
            if (_challenge.segment.output == _getOutputRoot(_sender, start)) {
                return start - 1;
            } else {
                return start;
            }
        }

        uint256 pos = (start + end) / 2;
        return pos;
    }

    function _newChallenger(string memory name) private returns (address) {
        address newAddr = makeAddr(name);

        vm.deal(newAddr, 10 ether);
        vm.prank(newAddr);
        pool.deposit{ value: newAddr.balance }();
        isChallenger[newAddr] = true;

        return newAddr;
    }

    function _createAssertion() private {
        uint256 nextBlockNumber = oracle.nextBlockNumber();
        // Roll to after the block number we'll submit
        warpToSubmitTime();

        address submitter = pool.nextValidator();

        uint256 _targetOutputIndex = oracle.latestOutputIndex() + 1;

        // Expect a burn event.
        vm.expectEmit(true, true, true, true);
        emit OutputSubmitted(
            bytes32(nextBlockNumber),
            _targetOutputIndex,
            nextBlockNumber,
            block.timestamp
        );

        vm.prank(submitter);
        oracle.submitL2Output(bytes32(nextBlockNumber), nextBlockNumber, 0, 0);
    }

    function _createChallenge(uint256 _outputIndex, address _challenger) private {
        Types.CheckpointOutput memory latestFinalizedOutput = oracle.getLatestFinalizeOutput();
        Types.CheckpointOutput memory targetOutput = oracle.getL2Output(_outputIndex);
        uint256 end = targetOutput.l2BlockNumber;
        uint256 start = end - oracle.SUBMISSION_INTERVAL();

        assertTrue(
            _getOutputRoot(targetOutput.submitter, end) != targetOutput.outputRoot,
            "not an invalid output"
        );

        Types.Assertion memory assertion = colosseum.getAssertion(_outputIndex);
        // Expect a ChallengeCreated event.
        vm.expectEmit(true, true, true, true);
        emit ChallengeCreated(_outputIndex, assertion.asserter, _challenger, block.timestamp);
        vm.prank(_challenger);
        colosseum.createChallenge(_outputIndex, bytes32(0), 0);

        Types.Challenge memory challenge = colosseum.getChallenge(_outputIndex, _challenger);
        assertEq(challenge.challenger, _challenger);
        assertEq(challenge.asserter, targetOutput.submitter);
        assertEq(challenge.turn, 1);
        assertEq(
            challenge.challengerTimeLeft,
            colosseum.MAX_CLOCK_DURATION_SECONDS() - (block.timestamp - assertion.assertedAt)
        );
        assertEq(challenge.asserterTimeLeft, colosseum.MAX_CLOCK_DURATION_SECONDS());
        assertEq(challenge.updatedAt, block.timestamp);
        assertEq(challenge.segment.start, latestFinalizedOutput.l2BlockNumber);
        assertEq(challenge.segment.end, targetOutput.l2BlockNumber);
        assertEq(
            challenge.segment.pos,
            (targetOutput.l2BlockNumber + latestFinalizedOutput.l2BlockNumber) / 2
        );
        assertEq(challenge.l1Head, blockhash(block.number - 1));
    }

    function _bisect(uint256 _outputIndex, address _challenger, address _sender) private {
        Types.Challenge memory challenge = colosseum.getChallenge(_outputIndex, _challenger);

        uint256 position = _detectFault(challenge, _sender);
        bytes32 output = _getOutputRoot(_sender, position);

        vm.prank(_sender);
        // check that ReadyToProve event was emitted on the last bisection.
        if ((challenge.segment.end - challenge.segment.start) / 2 == 2) {
            vm.expectEmit(true, true, false, false);
            emit ReadyToProve(_outputIndex, _challenger);
        }

        colosseum.bisect(_outputIndex, challenge.challenger, position, output);

        Types.Challenge memory newChallenge = colosseum.getChallenge(_outputIndex, _challenger);
        uint256 expectedChallengerTimeLeft;
        uint256 expectedAsserterTimeLeft;
        uint256 elapsed = newChallenge.updatedAt - challenge.updatedAt;

        if (challenge.turn % 2 == 0) {
            expectedChallengerTimeLeft = challenge.challengerTimeLeft - elapsed;
            expectedAsserterTimeLeft = challenge.asserterTimeLeft;
        } else {
            expectedAsserterTimeLeft = challenge.asserterTimeLeft - elapsed;
            expectedChallengerTimeLeft = challenge.challengerTimeLeft;
        }

        assertEq(newChallenge.turn, challenge.turn + 1);
        assertEq(newChallenge.segment.output, output);
        assertEq(newChallenge.updatedAt, block.timestamp);
        assertEq(newChallenge.asserterTimeLeft, expectedAsserterTimeLeft);
        assertEq(newChallenge.challengerTimeLeft, expectedChallengerTimeLeft);
    }

    function _proveFault(
        uint256 _outputIndex,
        address _challenger
    ) private returns (bytes32 publicInputHash) {
        // get previous snapshot
        Types.CheckpointOutput memory prevOutput = oracle.getL2Output(_outputIndex);

        Types.Challenge memory challenge = colosseum.getChallenge(_outputIndex, _challenger);

        _detectFault(challenge, challenge.challenger);

        publicInputHash = _doProveFault(challenge.challenger, _outputIndex);

        assertEq(
            uint256(colosseum.getStatus(_outputIndex, challenge.challenger)),
            uint256(Colosseum.ChallengeStatus.NONE)
        );

        Types.CheckpointOutput memory newOutput = oracle.getL2Output(_outputIndex);

        assertEq(newOutput.submitter, _challenger);
        assertEq(newOutput.outputRoot, bytes32(0));
        assertEq(prevOutput.timestamp, newOutput.timestamp);
        assertEq(prevOutput.l2BlockNumber, newOutput.l2BlockNumber);
    }

    function _doProveFault(address _challenger, uint256 _outputIndex) private returns (bytes32) {
        Types.ZkVmProof memory zkVmProof = ZkVmTestData.zkVmProof();

        vm.prank(_challenger);
        colosseum.proveFaultWithZkVm(_outputIndex, zkVmProof);

        return mockZKProofVerifier.hashZkVmPublicInput(zkVmProof.publicValues);
    }

    function _dismissChallenge(uint256 txId) private {
        // confirm transaction without check condition
        vm.prank(guardian1);
        securityCouncil.confirmTransaction(txId);

        vm.prank(guardian2);
        securityCouncil.confirmTransaction(txId);
    }

    function test_constructor_succeeds() external {
        assertEq(address(colosseum.L2_ORACLE()), address(oracle), "oracle address not matched");
        assertEq(
            address(colosseum.ZK_PROOF_VERIFIER()),
            address(zkProofVerifier),
            "zk proof verifier address not matched"
        );
        assertEq(colosseum.L2_ORACLE_SUBMISSION_INTERVAL(), submissionInterval);
        assertEq(colosseum.SECURITY_COUNCIL(), address(securityCouncil));
        assertEq(colosseum.L2_ORACLE_SUBMISSION_INTERVAL(), submissionInterval);
        assertEq(colosseum.MAX_CLOCK_DURATION_SECONDS(), maxClockDuration);
        assertEq(colosseum.GUARDIAN_PERIOD(), guardianPeriod);
    }

    function test_initialize_succeeds() external {}

    function test_createChallenge_succeeds() external {
        _createChallenge(targetOutputIndex, challenger);
    }

    function test_createChallenge_otherChallenger_succeeds() external {
        uint256 outputIndex = targetOutputIndex;
        _createChallenge(outputIndex, challenger);

        address otherChallenger = makeAddr("other challenger");

        vm.deal(otherChallenger, 1 ether);
        vm.prank(otherChallenger);
        pool.deposit{ value: requiredBondAmount }();

        _createChallenge(outputIndex, otherChallenger);

        // ensure that both challenges are enabled.
        assertEq(
            uint256(colosseum.getStatus(outputIndex, challenger)),
            uint256(Colosseum.ChallengeStatus.ASSERTER_TURN)
        );
        assertEq(
            uint256(colosseum.getStatus(outputIndex, otherChallenger)),
            uint256(Colosseum.ChallengeStatus.ASSERTER_TURN)
        );
    }

    function test_createChallenge_genesisOutput_reverts() external {
        vm.prank(challenger);
        vm.expectRevert(Colosseum.NotAllowedGenesisOutput.selector);
        colosseum.createChallenge(0, bytes32(0), 0);
    }

    function test_createChallenge_asAsserter_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        Types.CheckpointOutput memory targetOutput = oracle.getL2Output(outputIndex);

        vm.prank(targetOutput.submitter);
        vm.expectRevert(Colosseum.NotAllowedCaller.selector);
        colosseum.createChallenge(outputIndex, bytes32(0), 0);
    }

    function test_createChallenge_existedChallenge_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        _createChallenge(outputIndex, challenger);

        assertEq(
            uint256(colosseum.getStatus(outputIndex, challenger)),
            uint256(Colosseum.ChallengeStatus.ASSERTER_TURN)
        );

        vm.prank(challenger);
        vm.expectRevert(Colosseum.ImproperChallengeStatus.selector);
        colosseum.createChallenge(outputIndex, bytes32(0), 0);
    }

    function test_bisect_afterChallengerTimedOut_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        _createChallenge(outputIndex, challenger);

        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);
        _bisect(outputIndex, challenge.challenger, challenge.asserter);

        uint256 position = _detectFault(challenge, challenge.challenger);
        bytes32 output = _getOutputRoot(challenge.challenger, position);

        vm.warp(block.timestamp + challenge.challengerTimeLeft);
        vm.prank(challenge.challenger);
        vm.expectRevert(Colosseum.ChallengerTimeout.selector);
        colosseum.bisect(outputIndex, challenge.challenger, position, output);
    }

    function test_bisect_afterAsserterTimedOut_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        _createChallenge(outputIndex, challenger);

        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);
        _bisect(outputIndex, challenge.challenger, challenge.asserter);
        _bisect(outputIndex, challenge.challenger, challenge.challenger);

        challenge = colosseum.getChallenge(outputIndex, challenger);
        uint256 position = _detectFault(challenge, challenge.asserter);
        bytes32 output = _getOutputRoot(challenge.asserter, position);

        vm.warp(block.timestamp + challenge.asserterTimeLeft);
        vm.prank(challenge.asserter);
        vm.expectRevert(Colosseum.AsserterTimeout.selector);
        colosseum.bisect(outputIndex, challenge.challenger, position, output);
    }

    function test_challengerTimeout_succeeds() public {
        uint256 outputIndex = targetOutputIndex;
        _createChallenge(outputIndex, challenger);

        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);
        _bisect(outputIndex, challenge.challenger, challenge.asserter);

        uint256 position = _detectFault(challenge, challenge.challenger);
        bytes32 output = _getOutputRoot(challenge.challenger, position);

        vm.warp(block.timestamp + challenge.challengerTimeLeft);

        vm.expectEmit(true, true, false, true);
        emit ChallengerTimedOut(outputIndex, challenge.challenger, block.timestamp);
        vm.prank(challenge.asserter);
        colosseum.challengerTimeout(outputIndex, challenge.challenger);
    }

    function test_challengerTimeout_whenReadyToProve_succeeds() public {
        uint256 outputIndex = targetOutputIndex;
        _createChallenge(outputIndex, challenger);

        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);

        while (mockColosseum.isAbleToBisect(outputIndex, challenge.challenger)) {
            challenge = colosseum.getChallenge(outputIndex, challenge.challenger);
            _bisect(outputIndex, challenge.challenger, nextSender(challenge));
        }

        assertEq(
            uint256(colosseum.getStatus(outputIndex, challenger)),
            uint256(Colosseum.ChallengeStatus.READY_TO_PROVE)
        );

        vm.warp(block.timestamp + challenge.challengerTimeLeft + challengeGracePeriod);

        vm.expectEmit(true, true, false, true);
        emit ChallengerTimedOut(outputIndex, challenge.challenger, block.timestamp);
        vm.prank(challenge.asserter);
        colosseum.challengerTimeout(outputIndex, challenge.challenger);
    }

    function test_challengerTimeout_whenReadyToProve_reverts() public {
        uint256 outputIndex = targetOutputIndex;
        _createChallenge(outputIndex, challenger);

        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);

        while (mockColosseum.isAbleToBisect(outputIndex, challenge.challenger)) {
            challenge = colosseum.getChallenge(outputIndex, challenge.challenger);
            _bisect(outputIndex, challenge.challenger, nextSender(challenge));
        }

        assertEq(
            uint256(colosseum.getStatus(outputIndex, challenger)),
            uint256(Colosseum.ChallengeStatus.READY_TO_PROVE)
        );

        vm.warp(block.timestamp + challenge.challengerTimeLeft + challengeGracePeriod - 1);

        vm.prank(challenge.asserter);
        vm.expectRevert(Colosseum.ImproperChallengeStatus.selector);
        colosseum.challengerTimeout(outputIndex, challenge.challenger);
    }

    function test_createChallenge_notSubmittedOutput_reverts() external {
        uint256 outputIndex = targetOutputIndex;

        vm.prank(challenger);
        vm.expectRevert();
        colosseum.createChallenge(outputIndex + 1, bytes32(0), 0);
    }

    function test_createChallenge_afterChallengeProven_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        test_proveFaultWithZkVm_succeeds();

        assertEq(
            uint256(colosseum.getStatus(outputIndex, challenger)),
            uint256(Colosseum.ChallengeStatus.NONE)
        );

        vm.prank(challenger);
        vm.expectRevert(Colosseum.NotChallengeable.selector);
        colosseum.createChallenge(outputIndex, bytes32(0), 0);
    }

    function test_challengerTimeout_reverts() public {
        uint256 outputIndex = targetOutputIndex;
        _createChallenge(outputIndex, challenger);

        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);
        _bisect(outputIndex, challenge.challenger, challenge.asserter);

        uint256 position = _detectFault(challenge, challenge.challenger);
        bytes32 output = _getOutputRoot(challenge.challenger, position);

        vm.warp(block.timestamp + challenge.challengerTimeLeft - 1);

        vm.expectRevert(Colosseum.ImproperChallengeStatus.selector);
        vm.prank(challenge.asserter);
        colosseum.challengerTimeout(outputIndex, challenge.challenger);
    }

    function test_createChallenge_afterDismissed_reverts() external {
        uint256 outputIndex = targetOutputIndex;

        test_dismissChallenge_succeeds();

        vm.expectRevert(Colosseum.NotChallengeable.selector);
        colosseum.createChallenge(outputIndex, bytes32(0), 0);
    }

    function test_createChallenge_wrongFork_reverts() external {
        uint256 outputIndex = targetOutputIndex;

        vm.prank(challenger);
        vm.expectRevert(Colosseum.L1Reorged.selector);
        colosseum.createChallenge(outputIndex, bytes32(uint256(0x01)), block.number - 1);
    }

    function test_bisect_succeeds() external {
        uint256 outputIndex = targetOutputIndex;
        _createChallenge(outputIndex, challenger);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);

        assertEq(nextSender(challenge), challenge.asserter);

        _bisect(outputIndex, challenge.challenger, challenge.asserter);
    }

    function test_bisect_withBadPos_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        _createChallenge(outputIndex, challenger);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);

        assertEq(nextSender(challenge), challenge.asserter);

        vm.prank(challenge.asserter);

        // invalid output of the first segment
        uint256 invalid_position = _detectFault(challenge, challenge.challenger) + 1;
        bytes32 output = _getOutputRoot(challenge.challenger, invalid_position);

        vm.expectRevert(Colosseum.InvalidPos.selector);
        colosseum.bisect(outputIndex, challenge.challenger, invalid_position, output);
    }

    function test_bisect_ifNotYourTurn_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        _createChallenge(outputIndex, challenger);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);

        uint256 position = _detectFault(challenge, challenge.asserter);
        bytes32 output = _getOutputRoot(challenge.asserter, position);

        assertEq(nextSender(challenge), challenge.asserter);

        vm.prank(challenge.challenger);
        vm.expectRevert(Colosseum.NotAllowedCaller.selector);
        colosseum.bisect(outputIndex, challenge.challenger, position, output);
    }

    function test_bisect_whenAsserterTimedOut_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        _createChallenge(outputIndex, challenger);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);

        assertEq(nextSender(challenge), challenge.asserter);

        vm.warp(block.timestamp + challenge.asserterTimeLeft + 1);
        vm.prank(challenge.asserter);
        vm.expectRevert(Colosseum.AsserterTimeout.selector);
        colosseum.bisect(outputIndex, challenge.challenger, 0, 0);

        assertEq(
            uint256(colosseum.getStatus(outputIndex, challenge.challenger)),
            uint256(Colosseum.ChallengeStatus.ASSERTER_TIMEOUT)
        );
    }

    function test_bisect_whenChallengerTimedOut_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        _createChallenge(outputIndex, challenger);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);

        assertEq(nextSender(challenge), challenge.asserter);

        _bisect(outputIndex, challenge.challenger, challenge.asserter);

        // update challenge
        challenge = colosseum.getChallenge(outputIndex, challenge.challenger);

        vm.warp(block.timestamp + challenge.challengerTimeLeft + 1);
        vm.prank(challenge.challenger);
        vm.expectRevert(Colosseum.ChallengerTimeout.selector);
        colosseum.bisect(outputIndex, challenge.challenger, 0, 0);

        assertEq(
            uint256(colosseum.getStatus(outputIndex, challenger)),
            uint256(Colosseum.ChallengeStatus.CHALLENGER_TIMEOUT)
        );
    }

    function test_bisect_cancelChallenge_succeeds() external {
        uint256 outputIndex = targetOutputIndex;
        address otherChallenger = _newChallenger("other challenger");

        _createChallenge(outputIndex, otherChallenger);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, otherChallenger);
        // Make it the challenger turn
        _bisect(outputIndex, otherChallenger, challenge.asserter);

        // The output root of the target output index was replaced by another challenge.
        test_proveFaultWithZkVm_succeeds();

        uint256 prevDeposit = pool.balanceOf(otherChallenger);
        uint256 pendingBond = pool.getPendingBond(outputIndex, otherChallenger);

        vm.prank(otherChallenger);
        vm.expectEmit(true, true, false, true);
        emit ChallengeCanceled(outputIndex, otherChallenger, block.timestamp);

        colosseum.bisect(outputIndex, otherChallenger, 0, 0);

        // Ensure that the challenge has been deleted.
        assertEq(
            uint256(colosseum.getStatus(outputIndex, otherChallenger)),
            uint256(Colosseum.ChallengeStatus.NONE)
        );
        // Ensure that the pending bond has been refunded.
        vm.expectRevert("ValidatorPool: the pending bond does not exist");
        pool.getPendingBond(outputIndex, otherChallenger);
        assertEq(pool.balanceOf(otherChallenger), prevDeposit + pendingBond);
    }

    function test_bisect_cancelChallenge_senderNotChallenger_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        address otherChallenger = _newChallenger("other challenger");

        _createChallenge(outputIndex, otherChallenger);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, otherChallenger);
        // Make it the challenger turn
        _bisect(outputIndex, otherChallenger, challenge.asserter);

        // The output root of the target output index was replaced by another challenge.
        test_proveFaultWithZkVm_succeeds();

        vm.prank(challenger);
        vm.expectRevert(Colosseum.OnlyChallengerCanCancel.selector);
        colosseum.bisect(outputIndex, otherChallenger, 0, 0);
    }

    function test_proveFaultWithZkVm_succeeds() public returns (bytes32 publicInputHash) {
        uint256 outputIndex = targetOutputIndex;
        Types.CheckpointOutput memory targetOutput = oracle.getL2Output(outputIndex);

        _createChallenge(outputIndex, challenger);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);

        // Replace challenge.l1Head with test data.
        mockColosseum.setL1Head(outputIndex, challenge.challenger, ZkVmTestData.L1_HEAD);

        while (mockColosseum.isAbleToBisect(outputIndex, challenge.challenger)) {
            challenge = colosseum.getChallenge(outputIndex, challenge.challenger);
            _bisect(outputIndex, challenge.challenger, nextSender(challenge));
        }

        assertEq(
            uint256(colosseum.getStatus(outputIndex, challenger)),
            uint256(Colosseum.ChallengeStatus.READY_TO_PROVE)
        );

        publicInputHash = _proveFault(outputIndex, challenge.challenger);

        (, bytes32 outputRoot, , ) = colosseum.deletedOutputs(outputIndex);
        assertEq(outputRoot, targetOutput.outputRoot);
        assertTrue(colosseum.verifiedPublicInputs(publicInputHash));
    }

    function test_proveFault_finalizedOutput_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        _createChallenge(outputIndex, challenger);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);

        while (mockColosseum.isAbleToBisect(outputIndex, challenge.challenger)) {
            challenge = colosseum.getChallenge(outputIndex, challenge.challenger);
            _bisect(outputIndex, challenge.challenger, nextSender(challenge));
        }

        assertEq(
            uint256(colosseum.getStatus(outputIndex, challenger)),
            uint256(Colosseum.ChallengeStatus.READY_TO_PROVE)
        );

        Types.CheckpointOutput memory targetOutput = oracle.getL2Output(outputIndex);
        vm.warp(targetOutput.timestamp + 7 days + 1);

        // Expect a revert because the challenger has timed out.
        vm.expectRevert(Colosseum.ImproperChallengeStatus.selector);
        _doProveFault(challenger, outputIndex);
    }

    // TODO(pangssu): Testing is impossible in the current state. It must be fixed without fail.
    // function test_proveFault_whenAsserterTimedOut_succeeds() external {
    //     uint256 outputIndex = targetOutputIndex;
    //
    //     _createChallenge(outputIndex, challenger);
    //
    //     Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);
    //
    //     assertEq(nextSender(challenge), challenge.asserter);
    //
    //     vm.warp(challenge.timeoutAt + 1);
    //     // check the asserter timeout
    //     assertEq(
    //         uint256(colosseum.getStatus(outputIndex, challenge.challenger)),
    //         uint256(Colosseum.ChallengeStatus.ASSERTER_TIMEOUT)
    //     );
    //
    //     _proveFault(outputIndex, challenge.challenger);
    // }

    function test_proveFault_cancelChallenge_succeeds() external {
        uint256 outputIndex = targetOutputIndex;
        address otherChallenger = _newChallenger("other challenger");

        _createChallenge(outputIndex, otherChallenger);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, otherChallenger);

        while (mockColosseum.isAbleToBisect(outputIndex, otherChallenger)) {
            challenge = colosseum.getChallenge(outputIndex, otherChallenger);
            _bisect(outputIndex, otherChallenger, nextSender(challenge));
        }

        // The output root of the target output index was replaced by another challenge.
        test_proveFaultWithZkVm_succeeds();

        uint256 prevDeposit = pool.balanceOf(otherChallenger);
        uint256 pendingBond = pool.getPendingBond(outputIndex, otherChallenger);
        Types.ZkVmProof memory emptyZkVmProof;

        vm.prank(otherChallenger);
        colosseum.proveFaultWithZkVm(outputIndex, emptyZkVmProof);

        // Ensure that the challenge has been deleted.
        assertEq(
            uint256(colosseum.getStatus(outputIndex, otherChallenger)),
            uint256(Colosseum.ChallengeStatus.NONE)
        );

        // Ensure that the pending bond has been refunded.
        vm.expectRevert("ValidatorPool: the pending bond does not exist");
        pool.getPendingBond(outputIndex, otherChallenger);
        assertEq(pool.balanceOf(otherChallenger), prevDeposit + pendingBond);
    }

    function test_dismissChallenge_succeeds() public {
        uint256 outputIndex = targetOutputIndex;
        Types.CheckpointOutput memory output = oracle.getL2Output(outputIndex);

        bytes32 publicInputHash = test_proveFaultWithZkVm_succeeds();
        Types.CheckpointOutput memory newOutput = oracle.getL2Output(outputIndex);

        vm.prank(address(securityCouncil));
        colosseum.dismissChallenge(
            outputIndex,
            newOutput.submitter,
            output.submitter,
            output.outputRoot,
            publicInputHash
        );

        (, bytes32 outputRoot, , ) = colosseum.deletedOutputs(outputIndex);
        assertEq(outputRoot, bytes32(0));
        assertFalse(colosseum.verifiedPublicInputs(publicInputHash));
    }

    function test_dismissChallenge_notSecurityCouncil_reverts() external {
        test_proveFaultWithZkVm_succeeds();

        vm.prank(makeAddr("not_security_council"));
        vm.expectRevert(Colosseum.NotAllowedCaller.selector);
        colosseum.dismissChallenge(0, address(0), address(0), bytes32(0), bytes32(0));
    }

    function test_dismissChallenge_outputNotDeleted_reverts() external {
        vm.prank(address(securityCouncil));
        vm.expectRevert(Colosseum.OutputNotDeleted.selector);
        colosseum.dismissChallenge(
            targetOutputIndex,
            address(0),
            address(0),
            bytes32(0),
            bytes32(0)
        );
    }

    function test_dismissChallenge_finalizedOutput_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        _createChallenge(outputIndex, challenger);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);

        while (mockColosseum.isAbleToBisect(outputIndex, challenge.challenger)) {
            challenge = colosseum.getChallenge(outputIndex, challenge.challenger);
            _bisect(outputIndex, challenge.challenger, nextSender(challenge));
        }

        Types.CheckpointOutput memory targetOutput = oracle.getL2Output(outputIndex);
        vm.warp(targetOutput.timestamp + 7 days + 1);

        vm.prank(address(securityCouncil));
        vm.expectRevert(Colosseum.OutputAlreadyFinalized.selector);
        colosseum.dismissChallenge(0, address(0), address(0), bytes32(0), bytes32(0));
    }

    function test_dismissChallenge_invalidOutputGiven_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        Types.CheckpointOutput memory output = oracle.getL2Output(outputIndex);

        bytes32 publicInputHash = test_proveFaultWithZkVm_succeeds();
        Types.CheckpointOutput memory newOutput = oracle.getL2Output(outputIndex);

        vm.prank(address(securityCouncil));
        vm.expectRevert(Colosseum.InvalidOutputGiven.selector);
        colosseum.dismissChallenge(
            outputIndex,
            newOutput.submitter,
            output.submitter,
            bytes32(0),
            publicInputHash
        );
    }

    function test_dismissChallenge_invalidAddressGiven_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        Types.CheckpointOutput memory output = oracle.getL2Output(outputIndex);

        bytes32 publicInputHash = test_proveFaultWithZkVm_succeeds();
        Types.CheckpointOutput memory newOutput = oracle.getL2Output(outputIndex);

        vm.prank(address(securityCouncil));
        vm.expectRevert(Colosseum.InvalidAddressGiven.selector);
        colosseum.dismissChallenge(
            outputIndex,
            address(0),
            output.submitter,
            output.outputRoot,
            publicInputHash
        );

        vm.prank(address(securityCouncil));
        vm.expectRevert(Colosseum.InvalidAddressGiven.selector);
        colosseum.dismissChallenge(
            outputIndex,
            newOutput.submitter,
            address(0),
            output.outputRoot,
            publicInputHash
        );
    }

    function test_dismissChallenge_invalidPublicInput_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        Types.CheckpointOutput memory output = oracle.getL2Output(outputIndex);

        test_proveFaultWithZkVm_succeeds();
        Types.CheckpointOutput memory newOutput = oracle.getL2Output(outputIndex);

        vm.prank(address(securityCouncil));
        vm.expectRevert(Colosseum.InvalidPublicInputHash.selector);
        colosseum.dismissChallenge(
            outputIndex,
            newOutput.submitter,
            output.submitter,
            output.outputRoot,
            bytes32(0)
        );
    }

    function test_cancelChallenge_succeeds() external {
        uint256 outputIndex = targetOutputIndex;
        address otherChallenger = _newChallenger("other challenger");

        _createChallenge(outputIndex, otherChallenger);

        assertEq(
            uint256(colosseum.getStatus(outputIndex, otherChallenger)),
            uint256(Colosseum.ChallengeStatus.ASSERTER_TURN)
        );

        // The output root of the target output index was replaced by another challenge.
        test_proveFaultWithZkVm_succeeds();

        assertEq(
            uint256(colosseum.getStatus(outputIndex, otherChallenger)),
            uint256(Colosseum.ChallengeStatus.ASSERTER_TURN)
        );

        uint256 prevDeposit = pool.balanceOf(otherChallenger);
        uint256 pendingBond = pool.getPendingBond(outputIndex, otherChallenger);

        vm.prank(otherChallenger);
        colosseum.cancelChallenge(outputIndex);

        // Ensure that the pending bond has been refunded.
        vm.expectRevert("ValidatorPool: the pending bond does not exist");
        pool.getPendingBond(outputIndex, otherChallenger);
        assertEq(pool.balanceOf(otherChallenger), prevDeposit + pendingBond);
    }

    function test_cancelChallenge_noChallenge_reverts() external {
        vm.expectRevert(Colosseum.InvalidOutputGiven.selector);
        colosseum.cancelChallenge(0);
    }

    function test_cancelChallenge_outputNotDeleted_reverts() external {
        uint256 outputIndex = targetOutputIndex;

        _createChallenge(outputIndex, challenger);

        vm.prank(challenger);
        vm.expectRevert(Colosseum.CannotCancelChallenge.selector);
        colosseum.cancelChallenge(outputIndex);
    }

    function test_cancelChallenge_senderNotChallenger_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        address otherChallenger = _newChallenger("other challenger");

        _createChallenge(outputIndex, otherChallenger);

        // The output root of the target output index was replaced by another challenge.
        test_proveFaultWithZkVm_succeeds();

        vm.prank(challenger);
        vm.expectRevert(Colosseum.OnlyChallengerCanCancel.selector);
        colosseum.cancelChallenge(outputIndex);
    }

    function test_cancelChallenge_whenChallengerTimedOut_reverts() external {
        uint256 outputIndex = targetOutputIndex;
        address otherChallenger = _newChallenger("other challenger");
        _createChallenge(outputIndex, otherChallenger);
        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, otherChallenger);
        _bisect(outputIndex, otherChallenger, challenge.asserter);

        vm.warp(block.timestamp + 2);

        // The output root of the target output index was replaced by another challenge.
        test_proveFaultWithZkVm_succeeds();

        vm.warp(block.timestamp + challenge.challengerTimeLeft);
        vm.prank(otherChallenger);
        vm.expectRevert(Colosseum.ImproperChallengeStatusToCancel.selector);

        colosseum.cancelChallenge(outputIndex);
    }

    function test_forceDeleteOutput_succeeds() external {
        uint256 outputIndex = targetOutputIndex;

        _createChallenge(outputIndex, challenger);

        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);

        while (mockColosseum.isAbleToBisect(outputIndex, challenge.challenger)) {
            challenge = colosseum.getChallenge(outputIndex, challenge.challenger);
            _bisect(outputIndex, challenge.challenger, nextSender(challenge));
        }

        vm.prank(address(securityCouncil));
        colosseum.forceDeleteOutput(outputIndex);
    }

    function test_forceDeleteOutput_notSecurityCouncil_reverts() external {
        uint256 outputIndex = targetOutputIndex;

        vm.prank(address(1));
        vm.expectRevert(Colosseum.NotAllowedCaller.selector);
        colosseum.forceDeleteOutput(outputIndex);
    }

    function test_forceDeleteOutput_finalizedOutput_reverts() external {
        uint256 outputIndex = targetOutputIndex;

        vm.warp(block.timestamp + finalizationPeriod + 1);

        vm.prank(address(securityCouncil));
        vm.expectRevert(Colosseum.OutputAlreadyFinalized.selector);
        colosseum.forceDeleteOutput(outputIndex);
    }

    function test_forceDeleteOutput_alreadyDeletedOutput_reverts() external {
        uint256 outputIndex = targetOutputIndex;

        _createChallenge(outputIndex, challenger);

        Types.Challenge memory challenge = colosseum.getChallenge(outputIndex, challenger);

        while (mockColosseum.isAbleToBisect(outputIndex, challenge.challenger)) {
            challenge = colosseum.getChallenge(outputIndex, challenge.challenger);
            _bisect(outputIndex, challenge.challenger, nextSender(challenge));
        }

        vm.prank(address(securityCouncil));
        colosseum.forceDeleteOutput(outputIndex);

        vm.prank(address(securityCouncil));
        vm.expectRevert(Colosseum.OutputAlreadyDeleted.selector);
        colosseum.forceDeleteOutput(outputIndex);
    }
}
