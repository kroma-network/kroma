// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { GovernanceToken } from "../governance/GovernanceToken.sol";
import { MintManager } from "../governance/MintManager.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { L1Block } from "../L2/L1Block.sol";
import { CommonTest } from "./CommonTest.t.sol";

contract L1BlockTest is CommonTest {
    L1Block lb;
    address depositor;
    bytes32 immutable NON_ZERO_HASH = keccak256(abi.encode(1));

    GovernanceToken governanceToken;
    MintManager mintManager;

    function setUp() public override {
        super.setUp();

        // Deploy GovernanceToken
        governanceToken = new GovernanceToken(address(1), address(2), Predeploys.MINT_MANAGER);
        vm.etch(Predeploys.GOVERNANCE_TOKEN, address(governanceToken).code);
        governanceToken = GovernanceToken(Predeploys.GOVERNANCE_TOKEN);
        // Deploy MintManager
        mintManager = new MintManager(0, 1 ether, 3888000, 92224);
        vm.etch(Predeploys.MINT_MANAGER, address(mintManager).code);
        mintManager = MintManager(Predeploys.MINT_MANAGER);
        // Initialize MintManager
        address[] memory recipients = new address[](40);
        uint256[] memory shares = new uint256[](recipients.length);
        for (uint256 i = 0; i < recipients.length; i++) {
            recipients[i] = makeAddr(string(abi.encodePacked("recipient", i)));
            shares[i] = 100000 / recipients.length;
        }
        MintManager(Predeploys.MINT_MANAGER).initialize(recipients, shares);

        vm.etch(Predeploys.L1_BLOCK_ATTRIBUTES, address(new L1Block()).code);
        lb = L1Block(Predeploys.L1_BLOCK_ATTRIBUTES);
        depositor = lb.DEPOSITOR_ACCOUNT();
        vm.prank(depositor);
        lb.setL1BlockValues({
            _number: uint64(1),
            _timestamp: uint64(2),
            _basefee: 3,
            _hash: NON_ZERO_HASH,
            _sequenceNumber: uint64(4),
            _batcherHash: bytes32(0),
            _l1FeeOverhead: 2,
            _l1FeeScalar: 3,
            _validatorRewardScalar: 1
        });

        // Governance tokens can only be minted once per block, so increase the block number
        // after setting the L1block attributes.
        vm.roll(block.number + 1);
    }

    function testFuzz_updatesValues_succeeds(
        uint64 n,
        uint64 t,
        uint256 b,
        bytes32 h,
        uint64 s,
        bytes32 bt,
        uint256 fo,
        uint256 fs,
        uint256 vrr
    ) external {
        uint256 prevSupply = governanceToken.totalSupply();
        uint256 mintPerBlock = mintManager.mintAmountPerBlock(block.number);

        vrr = bound(vrr, 0, 10000);
        vm.prank(depositor);
        lb.setL1BlockValues(n, t, b, h, s, bt, fo, fs, vrr);
        assertEq(lb.number(), n);
        assertEq(lb.timestamp(), t);
        assertEq(lb.basefee(), b);
        assertEq(lb.hash(), h);
        assertEq(lb.sequenceNumber(), s);
        assertEq(lb.batcherHash(), bt);
        assertEq(lb.l1FeeOverhead(), fo);
        assertEq(lb.l1FeeScalar(), fs);
        assertEq(lb.validatorRewardScalar(), vrr);

        assertEq(governanceToken.totalSupply(), prevSupply + mintPerBlock);
    }

    function test_number_succeeds() external {
        assertEq(lb.number(), uint64(1));
    }

    function test_timestamp_succeeds() external {
        assertEq(lb.timestamp(), uint64(2));
    }

    function test_basefee_succeeds() external {
        assertEq(lb.basefee(), 3);
    }

    function test_hash_succeeds() external {
        assertEq(lb.hash(), NON_ZERO_HASH);
    }

    function test_sequenceNumber_succeeds() external {
        assertEq(lb.sequenceNumber(), uint64(4));
    }

    function test_updateValues_succeeds() external {
        uint256 prevSupply = governanceToken.totalSupply();
        uint256 mintPerBlock = mintManager.mintAmountPerBlock(block.number);

        vm.prank(depositor);
        lb.setL1BlockValues({
            _number: type(uint64).max,
            _timestamp: type(uint64).max,
            _basefee: type(uint256).max,
            _hash: keccak256(abi.encode(1)),
            _sequenceNumber: type(uint64).max,
            _batcherHash: bytes32(type(uint256).max),
            _l1FeeOverhead: type(uint256).max,
            _l1FeeScalar: type(uint256).max,
            _validatorRewardScalar: 10000
        });

        assertEq(governanceToken.totalSupply(), prevSupply + mintPerBlock);
    }

    function test_updateValues_updateTwice_succeeds() external {
        // Rollback to the previous block to test in the same block as before.
        vm.roll(block.number - 1);
        uint256 prevSupply = governanceToken.totalSupply();

        vm.prank(depositor);
        lb.setL1BlockValues({
            _number: type(uint64).max,
            _timestamp: type(uint64).max,
            _basefee: type(uint256).max,
            _hash: keccak256(abi.encode(1)),
            _sequenceNumber: type(uint64).max,
            _batcherHash: bytes32(type(uint256).max),
            _l1FeeOverhead: type(uint256).max,
            _l1FeeScalar: type(uint256).max,
            _validatorRewardScalar: 10000
        });

        // Tokens should not be minted even if executed twice in the same block.
        assertEq(governanceToken.totalSupply(), prevSupply);
    }
}
