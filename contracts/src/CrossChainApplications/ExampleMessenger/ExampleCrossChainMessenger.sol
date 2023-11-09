// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITeleporterMessenger, TeleporterMessageInput, TeleporterFeeInfo} from "../../Teleporter/ITeleporterMessenger.sol";
import {ITeleporterReceiver} from "../../Teleporter/ITeleporterReceiver.sol";
import {SafeERC20TransferFrom, SafeERC20} from "../../Teleporter/SafeERC20TransferFrom.sol";
import {TeleporterOwnerUpgradeable} from "../../Teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/security/ReentrancyGuard.sol";

/**
 * @dev ExampleCrossChainMessenger is an example contract that demonstrates how to send and receive
 * messages cross chain.
 */
contract ExampleCrossChainMessenger is
    ITeleporterReceiver,
    ReentrancyGuard,
    TeleporterOwnerUpgradeable
{
    using SafeERC20 for IERC20;

    // Messages sent to this contract.
    struct Message {
        address sender;
        string message;
    }

    mapping(bytes32 originChainID => Message message) private _messages;

    /**
     * @dev Emitted when a message is submited to be sent.
     */
    event SendMessage(
        bytes32 indexed destinationChainID,
        address indexed destinationAddress,
        address feeAsset,
        uint256 feeAmount,
        uint256 requiredGasLimit,
        string message
    );

    /**
     * @dev Emitted when a new message is received from a given chain ID.
     */
    event ReceiveMessage(
        bytes32 indexed originChainID,
        address indexed originSenderAddress,
        string message
    );

    constructor(
        address teleporterRegistryAddress
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress) {}

    /**
     * @dev See {ITeleporterReceiver-receiveTeleporterMessage}.
     *
     * Receives a message from another chain.
     */
    function receiveTeleporterMessage(
        bytes32 originChainID,
        address originSenderAddress,
        bytes calldata message
    ) external onlyAllowedTeleporter {
        // Store the message.
        string memory messageString = abi.decode(message, (string));
        _messages[originChainID] = Message(originSenderAddress, messageString);
        emit ReceiveMessage(originChainID, originSenderAddress, messageString);
    }

    /**
     * @dev Sends a message to another chain.
     * @return The message ID of the newly sent message.
     */
    function sendMessage(
        bytes32 destinationChainID,
        address destinationAddress,
        address feeContractAddress,
        uint256 feeAmount,
        uint256 requiredGasLimit,
        string calldata message
    ) external nonReentrant returns (uint256) {
        ITeleporterMessenger teleporterMessenger = teleporterRegistry
            .getLatestTeleporter();
        // For non-zero fee amounts, transfer the fee into the control of this contract first, and then
        // allow the Teleporter contract to spend it.
        uint256 adjustedFeeAmount = 0;
        if (feeAmount > 0) {
            adjustedFeeAmount = SafeERC20TransferFrom.safeTransferFrom(
                IERC20(feeContractAddress),
                feeAmount
            );
            IERC20(feeContractAddress).safeIncreaseAllowance(
                address(teleporterMessenger),
                adjustedFeeAmount
            );
        }

        emit SendMessage({
            destinationChainID: destinationChainID,
            destinationAddress: destinationAddress,
            feeAsset: feeContractAddress,
            feeAmount: adjustedFeeAmount,
            requiredGasLimit: requiredGasLimit,
            message: message
        });
        return
            teleporterMessenger.sendCrossChainMessage(
                TeleporterMessageInput({
                    destinationChainID: destinationChainID,
                    destinationAddress: destinationAddress,
                    feeInfo: TeleporterFeeInfo({
                        contractAddress: feeContractAddress,
                        amount: adjustedFeeAmount
                    }),
                    requiredGasLimit: requiredGasLimit,
                    allowedRelayerAddresses: new address[](0),
                    message: abi.encode(message)
                })
            );
    }

    /**
     * @dev Returns the current message from another chain.
     * @return The sender of the message, and the message itself.
     */
    function getCurrentMessage(
        bytes32 originChainID
    ) external view returns (address, string memory) {
        Message memory messageInfo = _messages[originChainID];
        return (messageInfo.sender, messageInfo.message);
    }
}
