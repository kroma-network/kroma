// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;


contract DestructForTest {
        uint256 a = 0;
    constructor() {
        a =123;
    }

    function transfer() public {
        selfdestruct(payable(msg.sender));
    }

}
