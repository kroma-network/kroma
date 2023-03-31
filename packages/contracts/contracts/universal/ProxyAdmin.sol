// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";

import { Proxy } from "./Proxy.sol";

/**
 * @title IStaticERC1967Proxy
 * @notice IStaticERC1967Proxy is a static version of the ERC1967 proxy interface.
 */
interface IStaticERC1967Proxy {
    function implementation() external view returns (address);

    function admin() external view returns (address);
}

/**
 * @title ProxyAdmin
 * @notice This is an auxiliary contract meant to be assigned as the admin of an ERC1967 Proxy,
 *         based on the OpenZeppelin implementation.
 */
contract ProxyAdmin is Ownable {
    /**
     * @param _owner Address of the initial owner of this contract.
     */
    constructor(address _owner) Ownable() {
        _transferOwnership(_owner);
    }

    /**
     * @notice Returns the implementation of the given proxy address.
     *
     * @param _proxy Address of the proxy to get the implementation of.
     *
     * @return Address of the implementation of the proxy.
     */
    function getProxyImplementation(address _proxy) external view returns (address) {
        return IStaticERC1967Proxy(_proxy).implementation();
    }

    /**
     * @notice Returns the admin of the given proxy address.
     *
     * @param _proxy Address of the proxy to get the admin of.
     *
     * @return Address of the admin of the proxy.
     */
    function getProxyAdmin(address payable _proxy) external view returns (address) {
        return IStaticERC1967Proxy(_proxy).admin();
    }

    /**
     * @notice Updates the admin of the given proxy address.
     *
     * @param _proxy    Address of the proxy to update.
     * @param _newAdmin Address of the new proxy admin.
     */
    function changeProxyAdmin(address payable _proxy, address _newAdmin) external onlyOwner {
        Proxy(_proxy).changeAdmin(_newAdmin);
    }

    /**
     * @notice Changes a proxy's implementation contract.
     *
     * @param _proxy          Address of the proxy to upgrade.
     * @param _implementation Address of the new implementation address.
     */
    function upgrade(address payable _proxy, address _implementation) public onlyOwner {
        Proxy(_proxy).upgradeTo(_implementation);
    }

    /**
     * @notice Changes a proxy's implementation contract and delegatecalls the new implementation
     *         with some given data. Useful for atomic upgrade-and-initialize calls.
     *
     * @param _proxy          Address of the proxy to upgrade.
     * @param _implementation Address of the new implementation address.
     * @param _data           Data to trigger the new implementation with.
     */
    function upgradeAndCall(
        address payable _proxy,
        address _implementation,
        bytes memory _data
    ) external payable onlyOwner {
        Proxy(_proxy).upgradeToAndCall{ value: msg.value }(_implementation, _data);
    }
}
