// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Test } from "forge-std/Test.sol";

import { Proxy } from "../universal/Proxy.sol";
import { ProxyAdmin } from "../universal/ProxyAdmin.sol";
import { SimpleStorage } from "./Proxy.t.sol";

contract ProxyAdmin_Test is Test {
    address alice = address(64);

    Proxy proxy;

    ProxyAdmin admin;

    SimpleStorage implementation;

    function setUp() external {
        // Deploy the proxy admin
        admin = new ProxyAdmin(alice);
        // Deploy the standard proxy
        proxy = new Proxy(address(admin));

        implementation = new SimpleStorage();
    }

    function test_owner_succeeds() external {
        assertEq(admin.owner(), alice);
    }

    function test_getProxyImplementation_succeeds() external {
        {
            address impl = admin.getProxyImplementation(payable(proxy));
            assertEq(impl, address(0));
        }

        vm.prank(alice);
        admin.upgrade(payable(proxy), address(implementation));

        {
            address impl = admin.getProxyImplementation(payable(proxy));
            assertEq(impl, address(implementation));
        }
    }

    function test_getProxyAdmin_succeeds() external {
        address owner = admin.getProxyAdmin(payable(proxy));
        assertEq(owner, address(admin));
    }

    function test_changeProxyAdmin_succeeds() external {
        vm.prank(alice);
        admin.changeProxyAdmin(payable(proxy), address(128));

        // The proxy is no longer the admin and can
        // no longer call the proxy interface.
        vm.expectRevert("Proxy: implementation not initialized");
        admin.getProxyAdmin(payable(proxy));

        // Call the proxy contract directly to get the admin.
        // Different proxy types have different interfaces.
        vm.prank(address(128));
        assertEq(Proxy(payable(proxy)).admin(), address(128));
    }

    function test_upgrade_succeeds() external {
        vm.prank(alice);
        admin.upgrade(payable(proxy), address(implementation));

        address impl = admin.getProxyImplementation(payable(proxy));
        assertEq(impl, address(implementation));
    }

    function test_upgradeAndCall_succeeds() external {
        vm.prank(alice);
        admin.upgradeAndCall(
            payable(proxy),
            address(implementation),
            abi.encodeWithSelector(SimpleStorage.set.selector, 1, 1)
        );

        address impl = admin.getProxyImplementation(payable(proxy));
        assertEq(impl, address(implementation));

        uint256 got = SimpleStorage(address(proxy)).get(1);
        assertEq(got, 1);
    }

    function test_onlyOwner_notOwner_reverts() external {
        vm.expectRevert("Ownable: caller is not the owner");
        admin.changeProxyAdmin(payable(proxy), address(0));

        vm.expectRevert("Ownable: caller is not the owner");
        admin.upgrade(payable(proxy), address(implementation));

        vm.expectRevert("Ownable: caller is not the owner");
        admin.upgradeAndCall(payable(proxy), address(implementation), hex"");
    }
}
