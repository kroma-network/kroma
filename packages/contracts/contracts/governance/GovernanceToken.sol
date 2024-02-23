// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol";

import { Predeploys } from "../libraries/Predeploys.sol";
import { KromaMintableERC20 } from "../universal/KromaMintableERC20.sol";

/**
 * @custom:proxied
 * @custom:predeploy 0x42000000000000000000000000000000000000ff
 * @title GovernanceToken
 * @notice The KRO token used in governance and supporting voting and delegation. Implements
 *         EIP 2612 allowing signed approvals. It can be bridged to any other specified chain.
 *         This token has a cap of 50m in total supply, and is minted once every block by the MintManager.
 */
contract GovernanceToken is KromaMintableERC20, ERC20Burnable, ERC20Votes {
    uint256 internal constant MAX_TOTAL_SUPPLY = 50_000_000 ether;

    address public immutable MINT_MANAGER;

    uint256 private _totalMinted;

    /**
     * @custom:semver 1.0.0
     *
     * @param _bridge      Address of the StandardBridge on this network.
     * @param _remoteToken Address of the corresponding version of this token on the remote chain.
     * @param _mintManager Address of the MintManager contract.
     */
    constructor(
        address _bridge,
        address _remoteToken,
        address _mintManager
    ) KromaMintableERC20(_bridge, _remoteToken, "Kroma", "KRO") ERC20Permit("Kroma") {
        MINT_MANAGER = _mintManager;
    }

    /**
     * @notice Allows the owner to mint tokens.
     *
     * @param _account The account receiving minted tokens.
     * @param _amount The amount of tokens to mint.
     */
    function mint(address _account, uint256 _amount) external override {
        require(
            msg.sender == BRIDGE || msg.sender == MINT_MANAGER,
            "GovernanceToken: only minter can mint"
        );

        _mint(_account, _amount);
        emit Mint(_account, _amount);
    }

    /**
     * @notice Returns the total minted amount.
     *
     * @return The total minted amount.
     */
    function totalMinted() public view returns (uint256) {
        return _totalMinted;
    }

    /**
     * @notice Returns the maximum number of tokens that can be minted.
     *
     * @return The maximum number of tokens.
     */
    function cap() public pure returns (uint256) {
        return MAX_TOTAL_SUPPLY;
    }

    /**
     * @notice Callback called after a token transfer.
     *
     * @param _from   The account sending tokens.
     * @param _to     The account receiving tokens.
     * @param _amount The amount of tokens being transferred.
     */
    function _afterTokenTransfer(
        address _from,
        address _to,
        uint256 _amount
    ) internal override(ERC20, ERC20Votes) {
        super._afterTokenTransfer(_from, _to, _amount);
    }

    /**
     * @notice Internal mint function.
     *
     * @param _account The account receiving minted tokens.
     * @param _amount  The amount of tokens to mint.
     */
    function _mint(address _account, uint256 _amount) internal override(ERC20, ERC20Votes) {
        uint256 minted = _totalMinted + _amount;
        require(minted <= MAX_TOTAL_SUPPLY, "GovernanceToken: cap exceeded");
        super._mint(_account, _amount);
        _totalMinted = minted;
        emit Mint(_account, _amount);
    }

    /**
     * @notice Internal burn function.
     *
     * @param _account The account that tokens will be burned from.
     * @param _amount  The amount of tokens that will be burned.
     */
    function _burn(address _account, uint256 _amount) internal override(ERC20, ERC20Votes) {
        super._burn(_account, _amount);
        emit Burn(_account, _amount);
    }
}
