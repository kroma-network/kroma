// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { UpgradeGovernor_Initializer } from "./CommonTest.t.sol";
import { Types } from "../libraries/Types.sol";
import { Proxy } from "../universal/Proxy.sol";

contract UpgradeGovernorTest is UpgradeGovernor_Initializer {
    uint256 blockNumber = 0;
    enum ProposalState {
        Pending,
        Active,
        Canceled,
        Defeated,
        Succeeded,
        Queued,
        Expired,
        Executed
    }

    function setUp() public virtual override {
        super.setUp();
        _roll(1);
        // minting to guardians
        vm.prank(owner);
        securityCouncilToken.safeMint(guardian1, baseUri);
        vm.prank(owner);
        securityCouncilToken.safeMint(guardian2, baseUri);

        vm.prank(guardian1);
        securityCouncilToken.delegate(guardian1);
        vm.prank(guardian2);
        securityCouncilToken.delegate(guardian2);

        assertEq(securityCouncilToken.balanceOf(guardian1), 1);
        assertEq(securityCouncilToken.balanceOf(guardian2), 1);
        assertEq(securityCouncilToken.getVotes(guardian1), 1);
        assertEq(securityCouncilToken.getVotes(guardian2), 1);
        assertEq(securityCouncilToken.owner(), owner);

        _roll(1);
    }

    function test_initialize_succeeds() external {
        // check securityCouncilToken
        assertEq(securityCouncilToken.symbol(), "KSC");

        // check timeLock
        bytes32 PROPOSER_ROLE = timeLock.PROPOSER_ROLE();
        bytes32 EXECUTOR_ROLE = timeLock.EXECUTOR_ROLE();
        bytes32 ADMIN_ROLE = timeLock.TIMELOCK_ADMIN_ROLE();
        assertTrue(timeLock.hasRole(PROPOSER_ROLE, address(upgradeGovernor)));
        assertTrue(timeLock.hasRole(EXECUTOR_ROLE, address(upgradeGovernor)));
        assertTrue(timeLock.hasRole(ADMIN_ROLE, address(upgradeGovernor)));

        // check upgradeGovernor
        assertEq(address(upgradeGovernor.token()), address(securityCouncilToken));
        assertEq(upgradeGovernor.timelock(), address(timeLock));
        assertEq(upgradeGovernor.votingDelay(), initialVotingDelay);
        assertEq(upgradeGovernor.votingPeriod(), initialVotingPeriod);
        assertEq(upgradeGovernor.proposalThreshold(), initialProposalThreshold);

        // check proxy admin
        vm.startPrank(address(timeLock));
        assertEq(Proxy(payable(address(upgradeGovernor))).admin(), address(timeLock));
        assertEq(Proxy(payable(address(securityCouncilToken))).admin(), address(timeLock));
        assertEq(Proxy(payable(address(timeLock))).admin(), address(timeLock));
        vm.stopPrank();
    }

    function _createProposal() private returns (uint256) {
        address[] memory targetContracts = new address[](1);
        targetContracts[0] = address(upgradeGovernor);
        uint256[] memory values = new uint256[](1);
        values[0] = 0;
        bytes[] memory callDatas = new bytes[](1);
        callDatas[0] = abi.encodeCall(upgradeGovernor.setVotingDelay, 10);
        vm.prank(guardian1);
        uint256 proposalId = upgradeGovernor.propose(
            targetContracts,
            values,
            callDatas,
            "test proposal"
        );
        _roll(1);
        return proposalId;
    }

    function _queue() private returns (uint256) {
        address[] memory targetContracts = new address[](1);
        targetContracts[0] = address(upgradeGovernor);
        uint256[] memory values = new uint256[](1);
        values[0] = 0;
        bytes[] memory callDatas = new bytes[](1);
        callDatas[0] = abi.encodeCall(upgradeGovernor.setVotingDelay, 10);
        vm.prank(guardian1);
        uint256 queuedId = upgradeGovernor.queue(
            targetContracts,
            values,
            callDatas,
            _hashDescription("test proposal")
        );
        return queuedId;
    }

    function _execute() private returns (uint256) {
        address[] memory targetContracts = new address[](1);
        targetContracts[0] = address(upgradeGovernor);
        uint256[] memory values = new uint256[](1);
        values[0] = 0;
        bytes[] memory callDatas = new bytes[](1);
        callDatas[0] = abi.encodeCall(upgradeGovernor.setVotingDelay, 10);
        vm.prank(guardian1);
        uint256 executedId = upgradeGovernor.execute(
            targetContracts,
            values,
            callDatas,
            _hashDescription("test proposal")
        );
        return executedId;
    }

    function _hashDescription(string memory desc) private pure returns (bytes32) {
        return keccak256(bytes(desc));
    }

    function _roll(uint256 addBlockNumber) private {
        blockNumber += addBlockNumber;
        vm.roll(blockNumber);
    }

    function test_createProposal_tokenThreshold_reverts() external {
        address[] memory targetContracts = new address[](1);
        targetContracts[0] = address(upgradeGovernor);
        uint256[] memory values = new uint256[](1);
        values[0] = 0;
        bytes[] memory callDatas = new bytes[](1);
        callDatas[0] = abi.encodeCall(upgradeGovernor.setVotingDelay, 10);
        vm.expectRevert("Governor: proposer votes below proposal threshold");
        vm.prank(notGuardian);
        upgradeGovernor.propose(targetContracts, values, callDatas, "test proposal");
    }

    function test_createProposal_succeeds() external {
        address[] memory targetContracts = new address[](1);
        targetContracts[0] = address(upgradeGovernor);
        uint256[] memory values = new uint256[](1);
        values[0] = 0;
        bytes[] memory callDatas = new bytes[](1);
        callDatas[0] = abi.encodeCall(upgradeGovernor.setVotingDelay, 10);
        vm.prank(guardian1);
        uint256 proposalId = upgradeGovernor.propose(
            targetContracts,
            values,
            callDatas,
            "test proposal"
        );
        uint8 state = uint8(upgradeGovernor.state(proposalId));
        assertEq(state, uint8(ProposalState.Pending));
    }

    function test_voteProposal_overPeriod_reverts() external {
        uint256 proposalId = _createProposal();
        _roll(initialVotingPeriod);
        vm.prank(guardian1);
        vm.expectRevert("Governor: vote not currently active");
        upgradeGovernor.castVote(proposalId, 1);
        uint8 state = uint8(upgradeGovernor.state(proposalId));
        assertEq(state, uint8(ProposalState.Defeated));
    }

    function test_voteProposal_succeeds() external {
        uint256 proposalId = _createProposal();
        vm.prank(guardian1);
        uint256 voted = upgradeGovernor.castVote(proposalId, 1);
        assertEq(voted, securityCouncilToken.balanceOf(guardian1));

        uint8 state = uint8(upgradeGovernor.state(proposalId));
        assertEq(state, uint8(ProposalState.Active));

        _roll(initialVotingPeriod);
        state = uint8(upgradeGovernor.state(proposalId));
        assertEq(state, uint8(ProposalState.Succeeded));
    }

    function test_queueProposal_succeeds() external {
        uint256 proposalId = _createProposal();

        // vote
        vm.prank(guardian1);
        upgradeGovernor.castVote(proposalId, 1);
        _roll(initialVotingPeriod);

        uint256 queuedId = _queue();
        assertEq(proposalId, queuedId);

        uint8 state = uint8(upgradeGovernor.state(proposalId));
        assertEq(state, uint8(ProposalState.Queued));
    }

    function test_executeProposal_succeeds() external {
        uint256 proposalId = _createProposal();

        // vote
        vm.prank(guardian1);
        upgradeGovernor.castVote(proposalId, 1);
        _roll(initialVotingPeriod);

        _queue();
        vm.warp(minDelaySeconds + 1);
        uint256 executedId = _execute();
        assertEq(proposalId, executedId);

        uint8 state = uint8(upgradeGovernor.state(proposalId));
        assertEq(state, uint8(ProposalState.Executed));

        uint256 votingDelay = upgradeGovernor.votingDelay();
        assertEq(votingDelay, 10);
    }
}
