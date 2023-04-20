// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { IERC165 } from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

import {
    IKromaMintableERC20
} from "../universal/IKromaMintableERC20.sol";
import { Bridge_Initializer } from "./CommonTest.t.sol";

contract KromaMintableERC20_Test is Bridge_Initializer {
    event Mint(address indexed account, uint256 amount);
    event Burn(address indexed account, uint256 amount);

    function test_semver_succeeds() external {
        assertEq(L2Token.version(), "0.1.0");
    }

    function test_remoteToken_succeeds() external {
        assertEq(L2Token.REMOTE_TOKEN(), address(L1Token));
    }

    function test_bridge_succeeds() external {
        assertEq(L2Token.BRIDGE(), address(L2Bridge));
    }

    function test_mint_succeeds() external {
        vm.expectEmit(true, true, true, true);
        emit Mint(alice, 100);

        vm.prank(address(L2Bridge));
        L2Token.mint(alice, 100);

        assertEq(L2Token.balanceOf(alice), 100);
    }

    function test_mint_notBridge_reverts() external {
        // NOT the bridge
        vm.expectRevert("KromaMintableERC20: only bridge can mint and burn");
        vm.prank(address(alice));
        L2Token.mint(alice, 100);
    }

    function test_burn_succeeds() external {
        vm.prank(address(L2Bridge));
        L2Token.mint(alice, 100);

        vm.expectEmit(true, true, true, true);
        emit Burn(alice, 100);

        vm.prank(address(L2Bridge));
        L2Token.burn(alice, 100);

        assertEq(L2Token.balanceOf(alice), 0);
    }

    function test_burn_notBridge_reverts() external {
        // NOT the bridge
        vm.expectRevert("KromaMintableERC20: only bridge can mint and burn");
        vm.prank(address(alice));
        L2Token.burn(alice, 100);
    }

    function test_erc165_supportsInterface_succeeds() external {
        // The assertEq calls in this test are comparing the manual calculation of the iface,
        // with what is returned by the solidity's type().interfaceId, just to be safe.
        bytes4 iface1 = bytes4(keccak256("supportsInterface(bytes4)"));
        assertEq(iface1, type(IERC165).interfaceId);
        assert(L2Token.supportsInterface(iface1));

        bytes4 iface2 = L2Token.REMOTE_TOKEN.selector ^
            L2Token.BRIDGE.selector ^
            L2Token.mint.selector ^
            L2Token.burn.selector;
        assertEq(iface2, type(IKromaMintableERC20).interfaceId);
        assert(L2Token.supportsInterface(iface2));
    }
}
