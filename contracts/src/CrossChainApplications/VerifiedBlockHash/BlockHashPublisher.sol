// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "../../Teleporter/ITeleporterMessenger.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./BlockHashReceiver.sol";

/**
 * Contract that publishes the latest block hash of current chain to another chain.
 */
contract BlockHashPublisher {
    ITeleporterMessenger public immutable teleporterMessenger;
    uint256 public constant RECEIVE_BLOCK_HASH_REQUIRED_GAS_LIMIT = 1.5e5;

    /**
     * @dev Emitted when a block hash is submitted to be published to another chain.
     */
    event PublishBlockHash(
        bytes32 indexed destinationChainID,
        address indexed destinationAddress,
        uint256 indexed blockHeight,
        bytes32 blockHash
    );

    constructor(address teleporterMessengerAddress) {
        teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);
    }

    /**
     * @dev Publishes the latest block hash to another chain.
     */
    function publishLatestBlockHash(
        bytes32 destinationChainID,
        address destinationAddress
    ) external returns (uint256 messageID) {
        // Get the latest block info. Note it must the previous block
        // because the current block hash is not available during execution.
        uint256 blockHeight = block.number - 1;
        bytes32 blockHash = blockhash(blockHeight);

        // ABI encode the function arguments to be called on the destination.
        // The originChainID and originSenderAddress arguments of the target function are provided by Warp/Teleporter.
        bytes memory messageData = abi.encode(blockHeight, blockHash);

        emit PublishBlockHash(
            destinationChainID,
            destinationAddress,
            blockHeight,
            blockHash
        );
        messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationChainID: destinationChainID,
                destinationAddress: destinationAddress,
                feeInfo: TeleporterFeeInfo({
                    contractAddress: address(0),
                    amount: 0
                }),
                requiredGasLimit: RECEIVE_BLOCK_HASH_REQUIRED_GAS_LIMIT,
                allowedRelayerAddresses: new address[](0),
                message: messageData
            })
        );
    }
}
