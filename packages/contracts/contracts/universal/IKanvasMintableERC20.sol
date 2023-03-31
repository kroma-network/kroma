// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { IERC165 } from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

/**
 * @title IKanvasMintableERC20
 * @notice This interface is available on the KanvasMintableERC20 contract. We declare it as a
 *         separate interface so that it can be used in custom implementations of
 *         KanvasMintableERC20.
 */
interface IKanvasMintableERC20 {
    function REMOTE_TOKEN() external view returns (address);

    function BRIDGE() external view returns (address);

    function mint(address _to, uint256 _amount) external;

    function burn(address _from, uint256 _amount) external;
}
