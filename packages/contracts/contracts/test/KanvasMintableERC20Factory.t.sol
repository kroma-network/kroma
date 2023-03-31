// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Bridge_Initializer } from "./CommonTest.t.sol";
import { LibRLP } from "./RLP.t.sol";

contract KanvasMintableTokenFactory_Test is Bridge_Initializer {
    event KanvasMintableERC20Created(
        address indexed localToken,
        address indexed remoteToken,
        address deployer
    );

    function setUp() public override {
        super.setUp();
    }

    function test_bridge_succeeds() external {
        assertEq(address(L2TokenFactory.BRIDGE()), address(L2Bridge));
    }

    function test_createKanvasMintableERC20_succeeds() external {
        address remote = address(4);
        address local = LibRLP.computeAddress(address(L2TokenFactory), 2);

        vm.expectEmit(true, true, true, true);
        emit KanvasMintableERC20Created(local, remote, alice);

        vm.prank(alice);
        L2TokenFactory.createKanvasMintableERC20(remote, "Beep", "BOOP");
    }

    function test_createKanvasMintableERC20_sameTwice_succeeds() external {
        address remote = address(4);

        vm.prank(alice);
        L2TokenFactory.createKanvasMintableERC20(remote, "Beep", "BOOP");

        address local = LibRLP.computeAddress(address(L2TokenFactory), 3);

        vm.expectEmit(true, true, true, true);
        emit KanvasMintableERC20Created(local, remote, alice);

        vm.prank(alice);
        L2TokenFactory.createKanvasMintableERC20(remote, "Beep", "BOOP");
    }

    function test_createKanvasMintableERC20_remoteIsZero_succeeds() external {
        address remote = address(0);
        vm.expectRevert("KanvasMintableERC20Factory: must provide remote token address");
        L2TokenFactory.createKanvasMintableERC20(remote, "Beep", "BOOP");
    }
}
