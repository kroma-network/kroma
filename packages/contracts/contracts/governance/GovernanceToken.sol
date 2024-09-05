// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol";
import "@openzeppelin/contracts-upgradeable/access/Ownable2StepUpgradeable.sol";

import { KromaMintableERC20 } from "../universal/KromaMintableERC20.sol";

/**
 * @custom:proxied
 * @title GovernanceToken
 * @notice The KRO token used in governance, supporting voting and delegation. Implements
 *         EIP 2612 allowing signed approvals. `mint` function is only allowed to the owner or
 *         `Bridge`, and the total supply amount is minted at once (TGE). `Bridge` has the
 *         permission to `mint` and `burn`, for the purpose of bridging KRO to the remote chain.
 */
contract GovernanceToken is KromaMintableERC20, ERC20Votes, Ownable2StepUpgradeable {
    /**
     * @notice Constructs the GovernanceToken contract.
     *
     * @param _bridge      Address of the StandardBridge contract on this network.
     * @param _remoteToken Address of the corresponding token on the remote chain.
     */
    constructor(
        address _bridge,
        address _remoteToken
    ) KromaMintableERC20(_bridge, _remoteToken, "", "") ERC20Permit("Kroma") {
        _disableInitializers();
    }

    /**
     * @notice Initializer.
     *
     * @param _owner The owner of this contract.
     */
    function initialize(address _owner) public initializer {
        __Ownable2Step_init();
        transferOwnership(_owner);
    }

    /**
     * @notice Allows StandardBridge or the owner to mint tokens.
     *
     * @param _to     Address to mint tokens to.
     * @param _amount Amount of tokens to mint.
     */
    function mint(address _to, uint256 _amount) external override {
        require(
            msg.sender == BRIDGE || msg.sender == owner(),
            "GovernanceToken: only bridge or owner can mint"
        );

        _mint(_to, _amount);
    }

    /**
     * @inheritdoc KromaMintableERC20
     */
    function burn(address _from, uint256 _amount) external override onlyBridge {
        _burn(_from, _amount);
    }

    /**
     * @inheritdoc ERC20
     */
    function name() public pure override returns (string memory) {
        return "Kroma";
    }

    /**
     * @inheritdoc ERC20
     */
    function symbol() public pure override returns (string memory) {
        return "KRO";
    }

    /**
     * @notice Callback called after a token transfer.
     *
     * @param from   The account sending tokens.
     * @param to     The account receiving tokens.
     * @param amount The amount of tokens being transferred.
     */
    function _afterTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal override(ERC20, ERC20Votes) {
        super._afterTokenTransfer(from, to, amount);
    }

    /**
     * @notice Internal mint function.
     *
     * @param account The account receiving minted tokens.
     * @param amount  The amount of tokens to mint.
     */
    function _mint(address account, uint256 amount) internal override(ERC20, ERC20Votes) {
        super._mint(account, amount);
        emit Mint(account, amount);
    }

    /**
     * @notice Internal burn function.
     *
     * @param account The account that tokens will be burned from.
     * @param amount  The amount of tokens that will be burned.
     */
    function _burn(address account, uint256 amount) internal override(ERC20, ERC20Votes) {
        super._burn(account, amount);
        emit Burn(account, amount);
    }

    /**
     * @notice Override function.
     */
    function _msgSender() internal view override(Context, ContextUpgradeable) returns (address) {
        return super._msgSender();
    }

    /**
     * @notice Override function.
     */
    function _msgData()
        internal
        view
        override(Context, ContextUpgradeable)
        returns (bytes calldata)
    {
        return super._msgData();
    }
}
