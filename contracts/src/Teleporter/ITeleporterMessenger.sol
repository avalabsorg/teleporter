// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

struct TeleporterMessageReceipt {
    uint256 receivedMessageID;
    address relayerRewardAddress;
}

struct TeleporterMessageInput {
    bytes32 destinationChainID;
    address destinationAddress;
    TeleporterFeeInfo feeInfo;
    uint256 requiredGasLimit;
    address[] allowedRelayerAddresses;
    bytes message;
}

struct TeleporterMessage {
    uint256 messageID;
    address senderAddress;
    address destinationAddress;
    uint256 requiredGasLimit;
    address[] allowedRelayerAddresses;
    TeleporterMessageReceipt[] receipts;
    bytes message;
}

struct TeleporterFeeInfo {
    address contractAddress;
    uint256 amount;
}

/**
 * @dev Interface that describes functionalities for a cross chain messenger.
 */
interface ITeleporterMessenger {
    /**
     * @dev Emitted when sending a Teleporter message cross chain.
     */
    event SendCrossChainMessage(
        bytes32 indexed destinationChainID,
        uint256 indexed messageID,
        TeleporterMessage message
    );

    /**
     * @dev Emitted when an additional fee amount is added to a Teleporter message that had previously
     * been sent, but receipt not yet received.
     */
    event AddFeeAmount(
        bytes32 indexed destinationChainID,
        uint256 indexed messageID,
        TeleporterFeeInfo updatedFeeInfo
    );

    /**
     * @dev Emitted when Teleporter message is being delivered on destination chain and address,
     * but message execution fails. Failed messages can then be retried.
     */
    event FailedMessageExecution(
        bytes32 indexed originChainID,
        uint256 indexed messageID,
        TeleporterMessage message
    );

    /**
     * @dev Emitted when a Teleporter message that previously failed to be executed properly
     * is retried and is successfully executed. The message is then no longer able to be retried again
     * in the future since each message will be successfully executed at most once.
     */
    event MessageExecutionRetried(
        bytes32 indexed originChainID,
        uint256 indexed messageID
    );

    /**
     * @dev Emitted when a TeleporterMessage is successfully received.
     */
    event ReceiveCrossChainMessage(
        bytes32 indexed originChainID,
        uint256 indexed messageID,
        TeleporterMessage message
    );

    /**
     * @dev Called by transactions to initiate the sending of a cross subnet message.
     */
    function sendCrossChainMessage(
        TeleporterMessageInput calldata messageInput
    ) external returns (uint256 messageID);

    /**
     * @dev Called by transactions to retry the sending of a cross subnet message.
     *
     * Retriggers the sending of a message previously emitted by sendCrossChainMessage that has not yet been acknowledged
     * with a receipt from the destination chain. This may be necessary in the unlikely event that less than the required
     * threshold of stake weight successfully inserted the message in their messages DB at the time of the first submission.
     * The message is checked to have already been previously submitted by comparing it's message hash against those kept in
     * state until a receipt is received for the message.
     */
    function retrySendCrossChainMessage(
        bytes32 destinationChainID,
        TeleporterMessage calldata message
    ) external;

    /**
     * @dev Adds the additional fee amount to the amount to be paid to the relayer that delivers
     * the given message ID to the destination subnet.
     *
     * The fee contract address must be the same asset type as the fee asset specified in the original
     * call to sendCrossChainMessage. Returns a failure if the message doesn't exist or there is already
     * receipt of delivery of the message.
     */
    function addFeeAmount(
        bytes32 destinationChainID,
        uint256 messageID,
        address feeContractAddress,
        uint256 additionalFeeAmount
    ) external;

    /**
     * @dev Receives a cross chain message, and marks the `relayerRewardAddress` for fee reward for a successful delivery.
     *
     * The message specified by `messageIndex` must be provided at that index in the access list storage slots of the transaction,
     * and is verified in the precompile predicate.
     */
    function receiveCrossChainMessage(
        uint32 messageIndex,
        address relayerRewardAddress
    ) external;

    /**
     * @dev Retries the execution of a previously delivered message by verifying the payload matches
     * the hash of the payload originally delivered, and calling the destination address again.
     *
     * Intended to be used if the original required gas limit was not sufficient for the message
     * execution. Messages are ensured to be successfully executed at most once.
     */
    function retryMessageExecution(
        bytes32 originChainID,
        TeleporterMessage calldata message
    ) external;

    /**
     * @dev Retries the sending of receipts for the given `messageIDs`.
     *
     * Sends the specified message receipts in a new message (with an empty payload) back to the origin chain.
     * This is intended to be used if the message receipts were originally included in messages that were dropped
     * or otherwise not delivered in a timely manner.
     */
    function retryReceipts(
        bytes32 originChainID,
        uint256[] calldata messageIDs,
        TeleporterFeeInfo calldata feeInfo,
        address[] calldata allowedRelayerAddresses
    ) external returns (uint256 messageID);

    /**
     * @dev Sends any fee amount rewards for the given fee asset out to the caller.
     */
    function redeemRelayerRewards(address feeAsset) external;

    /**
     * @dev Gets the hash of a given message stored in the EVM state, if the message exists.
     */
    function getMessageHash(
        bytes32 destinationChainID,
        uint256 messageID
    ) external view returns (bytes32 messageHash);

    /**
     * @dev Checks whether or not the given message has been received by this chain.
     */
    function messageReceived(
        bytes32 originChainID,
        uint256 messageID
    ) external view returns (bool delivered);

    /**
     * @dev Returns the address the relayer reward should be sent to on the origin subnet
     * for a given message, assuming that the message has already been delivered.
     */
    function getRelayerRewardAddress(
        bytes32 originChainID,
        uint256 messageID
    ) external view returns (address relayerRewardAddress);

    /**
     * Gets the current reward amount of a given fee asset that is redeemable by the given relayer.
     */
    function checkRelayerRewardAmount(
        address relayer,
        address feeAsset
    ) external view returns (uint256);

    /**
     * @dev Gets the fee asset and amount for a given message.
     */
    function getFeeInfo(
        bytes32 destinationChainID,
        uint256 messageID
    ) external view returns (address feeAsset, uint256 feeAmount);
}
