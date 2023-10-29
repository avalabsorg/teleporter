// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import "./IERC20TokenSource.sol";
import "../../Teleporter/ITeleporterMessenger.sol";
import "../../Teleporter/ITeleporterReceiver.sol";
import "../../Teleporter/SafeERC20TransferFrom.sol";

contract ERC20TokenSource is
    ITeleporterReceiver,
    IERC20TokenSource,
    ReentrancyGuard
{
    using SafeERC20 for IERC20;

    uint256 public constant MINT_NATIVE_TOKENS_REQUIRED_GAS = 150_000; // TODO this is a placeholder
    bytes32 public immutable currentBlockchainID;
    bytes32 public immutable destinationBlockchainID;
    address public immutable nativeTokenDestinationAddress;

    // Used for sending an receiving Teleporter messages.
    ITeleporterMessenger public immutable teleporterMessenger;

    constructor(
        address teleporterMessengerAddress,
        bytes32 destinationBlockchainID_,
        address nativeTokenDestinationAddress_
    ) {
        currentBlockchainID = WarpMessenger(
            0x0200000000000000000000000000000000000005
        ).getBlockchainID();

        require(
            teleporterMessengerAddress != address(0),
            "Invalid TeleporterMessenger Address"
        );
        teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);

        require(
            destinationBlockchainID_ != bytes32(0),
            "Invalid Destination Chain ID"
        );
        require(
            destinationBlockchainID_ != currentBlockchainID,
            "Cannot Bridge With Same Blockchain"
        );
        destinationBlockchainID = destinationBlockchainID_;

        require(
            nativeTokenDestinationAddress_ != address(0),
            "Invalid Destination Contract Address"
        );
        nativeTokenDestinationAddress = nativeTokenDestinationAddress_;
    }

    /**
     * @dev See {ITeleporterReceiver-receiveTeleporterMessage}.
     *
     * Receives a Teleporter message and routes to the appropriate internal function call.
     */
    function receiveTeleporterMessage(
        bytes32 senderBlockchainID,
        address senderAddress,
        bytes calldata message
    ) external nonReentrant {
        // Only allow the Teleporter messenger to deliver messages.
        require(
            msg.sender == address(teleporterMessenger),
            "Unauthorized TeleporterMessenger contract"
        );

        // Only allow messages from the destination chain.
        require(
            senderBlockchainID == destinationBlockchainID,
            "Invalid Destination Chain"
        );

        // Only allow the partner contract to send messages.
        require(
            senderAddress == nativeTokenDestinationAddress,
            "Unauthorized Sender"
        );

        (address recipient, uint256 amount) = abi.decode(
            message,
            (address, uint256)
        );
        require(recipient != address(0), "Invalid Recipient Address");

        // Send to recipient
        payable(recipient).transfer(amount);

        emit UnlockTokens(recipient, amount);
    }

    /**
     * @dev See {IERC20TokenSource-transferToDestination}.
     */
    function transferToDestination(
        address recipient,
        address ERC20ContractAddress,
        uint256 totalAmount,
        uint256 feeAmount,
        address[] calldata allowedRelayerAddresses
    ) external nonReentrant {
        // The recipient cannot be the zero address.
        require(recipient != address(0), "Invalid Recipient Address");

        // Lock tokens in this contract. Supports "fee/burn on transfer" ERC20 token
        // implementations by only bridging the actual balance increase reflected by the call
        // to transferFrom.
        uint256 adjustedAmount = SafeERC20TransferFrom.safeTransferFrom(
            IERC20(ERC20ContractAddress),
            totalAmount
        );

        // Ensure that the adjusted amount is greater than the fee to be paid.
        require(
            adjustedAmount > feeAmount,
            "ERC20TokenSource: insufficient adjusted amount"
        );

        // Allow the Teleporter messenger to spend the fee amount.
        if (feeAmount > 0) {
            IERC20(ERC20ContractAddress).safeIncreaseAllowance(
                address(teleporterMessenger),
                feeAmount
            );
        }

        uint256 transferAmount = totalAmount - feeAmount;

        uint256 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationChainID: destinationBlockchainID,
                destinationAddress: nativeTokenDestinationAddress,
                feeInfo: TeleporterFeeInfo({
                    contractAddress: ERC20ContractAddress,
                    amount: feeAmount
                }),
                requiredGasLimit: MINT_NATIVE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: allowedRelayerAddresses,
                message: abi.encode(recipient, transferAmount)
            })
        );

        emit TransferToDestination({
            sender: msg.sender,
            recipient: recipient,
            ERC20ContractAddress: ERC20ContractAddress,
            transferAmount: totalAmount,
            feeAmount: feeAmount,
            teleporterMessageID: messageID
        });
    }
}
