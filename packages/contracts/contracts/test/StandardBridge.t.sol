// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

import { KanvasMintableERC20 } from "../universal/KanvasMintableERC20.sol";
import { StandardBridge } from "../universal/StandardBridge.sol";
import { CommonTest } from "./CommonTest.t.sol";

/**
 * @title StandardBridgeTester
 * @notice Simple wrapper around the StandardBridge contract that exposes
 *         internal functions so they can be more easily tested directly.
 */
contract StandardBridgeTester is StandardBridge {
    constructor(address payable _messenger, address payable _otherBridge)
        StandardBridge(_messenger, _otherBridge)
    {}

    function isKanvasMintableERC20(address _token) external view returns (bool) {
        return _isKanvasMintableERC20(_token);
    }

    function isCorrectTokenPair(address _mintableToken, address _otherToken)
        external
        view
        returns (bool)
    {
        return _isCorrectTokenPair(_mintableToken, _otherToken);
    }

    receive() external payable override {}
}

/**
 * @title StandardBridge_Stateless_Test
 * @notice Tests internal functions that require no existing state or contract
 *         interactions with the messenger.
 */
contract StandardBridge_Stateless_Test is CommonTest {
    StandardBridgeTester internal bridge;
    KanvasMintableERC20 internal mintable;
    ERC20 internal erc20;

    function setUp() public override {
        super.setUp();

        bridge = new StandardBridgeTester({
            _messenger: payable(address(0)),
            _otherBridge: payable(address(0))
        });

        mintable = new KanvasMintableERC20({
            _bridge: address(0),
            _remoteToken: address(0),
            _name: "Stonks",
            _symbol: "STONK"
        });

        erc20 = new ERC20("Altcoin", "ALT");
    }

    /**
     * @notice Test coverage for identifying KanvasMintableERC20 tokens.
     *         This function should return true for
     *         KanvasMintableERC20 tokens and false for any accounts that
     *         do not implement the interface.
     */
    function test_isKanvasMintableERC20_succeeds() external {
        // Both the modern and legacy mintable tokens should return true
        assertTrue(bridge.isKanvasMintableERC20(address(mintable)));
        // A regular ERC20 should return false
        assertFalse(bridge.isKanvasMintableERC20(address(erc20)));
        // Non existent contracts should return false and not revert
        assertEq(address(0x20).code.length, 0);
        assertFalse(bridge.isKanvasMintableERC20(address(0x20)));
    }

    /**
     * @notice Test coverage of isCorrectTokenPair under different types of
     *         tokens.
     */
    function test_isCorrectTokenPair_succeeds() external {
        // Modern + known to be correct remote token
        assertTrue(bridge.isCorrectTokenPair(address(mintable), mintable.REMOTE_TOKEN()));
        // Modern + known to be incorrect remote token
        assertTrue(mintable.REMOTE_TOKEN() != address(0x20));
        assertFalse(bridge.isCorrectTokenPair(address(mintable), address(0x20)));
    }
}
