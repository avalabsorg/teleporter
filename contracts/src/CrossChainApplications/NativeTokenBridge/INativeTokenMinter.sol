// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

/**
 * @dev Interface that describes functionalities for a cross-chain ERC20 bridge.
 */
interface INativeTokenMinter {

    /**
     * @dev Emitted when tokens are locked in this bridge contract to be bridged to another chain.
     */
    event BridgeTokens(
        address indexed tokenContractAddress,
        uint256 indexed teleporterMessageID,
        bytes32 destinationChainID,
        address destinationBridgeAddress,
        address recipient,
        uint256 transferAmount,
        uint256 feeAmount
    );

    /**
     * @dev Emitted when minting bridge tokens.
     */
    event MintNativeTokens(
        address recipient,
        uint256 amount
    );

    /**
     * @dev Transfers ERC20 tokens to another chain.
     *
     * This can be wrapping, unwrapping, and transferring a wrapped token between two non-native chains.
     */
    function bridgeTokens(
        address recipient,
        address feeTokenContractAddress,
        uint256 feeAmount
    ) external payable;
}
