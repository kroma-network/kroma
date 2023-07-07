pragma solidity 0.8.15;

import { Vm } from "forge-std/Vm.sol";

import { L2OutputOracle } from "../../L1/L2OutputOracle.sol";
import { L2OutputOracle_Initializer } from "../CommonTest.t.sol";

contract L2OutputOracle_Validator {
    L2OutputOracle internal oracle;
    Vm internal vm;

    constructor(L2OutputOracle _oracle, Vm _vm) {
        oracle = _oracle;
        vm = _vm;
    }

    /**
     * @dev Allows the actor to submit an L2 output to the `L2OutputOracle`
     */
    function submitL2Output(
        bytes32 _outputRoot,
        uint256 _l2BlockNumber,
        bytes32 _l1BlockHash,
        uint256 _l1BlockNumber,
        uint256 _bondAmount
    ) external {
        // Act as the validator and submit a new output.
        vm.prank(oracle.VALIDATOR_POOL().nextValidator());
        oracle.submitL2Output(_outputRoot, _l2BlockNumber, _l1BlockHash, _l1BlockNumber, _bondAmount);
    }
}

contract L2OutputOracle_MonotonicBlockNumIncrease_Invariant is L2OutputOracle_Initializer {
    L2OutputOracle_Validator internal actor;

    function setUp() public override {
        super.setUp();

        // Create a proposer actor.
        actor = new L2OutputOracle_Validator(oracle, vm);

        // Set the target contract to the validator actor
        targetContract(address(actor));

        // Set the target selector for `submitL2Output`
        // `submitL2Output` is the only function we care about, as it is the only function
        // that can modify the `l2Outputs` array in the oracle.
        bytes4[] memory selectors = new bytes4[](1);
        selectors[0] = actor.submitL2Output.selector;
        FuzzSelector memory selector = FuzzSelector({ addr: address(actor), selectors: selectors });
        targetSelector(selector);
    }

    /**
     * @custom:invariant The block number of the checkpoint output should monotonically
     * increase.
     *
     * When a new output is submitted, it should never be allowed to correspond to a block
     * number that is less than the current output.
     */
    function invariant_monotonicBlockNumIncrease() external {
        // Assert that the block number of checkpoint output must monotonically increase.
        assertTrue(oracle.nextBlockNumber() >= oracle.latestBlockNumber());
    }
}
