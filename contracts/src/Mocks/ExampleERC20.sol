// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

contract ExampleERC20 is ERC20Burnable {
    string private constant _TOKEN_NAME = "Mock Token";
    string private constant _TOKEN_SYMBOL = "EXMP";

    // Errors
    error MaxAmountExceeded(uint256 maxAmount, uint256 mintAmount);

    constructor() ERC20(_TOKEN_NAME, _TOKEN_SYMBOL) {
        _mint(msg.sender, 10_000_000_000_000_000_000_000_000_000);
    }

    function mint(uint256 amount) public {
        if (amount > 10_000_000_000_000_000) {
            revert MaxAmountExceeded(10_000_000_000_000_000, amount);
        }

        _mint(msg.sender, amount);
    }
}
