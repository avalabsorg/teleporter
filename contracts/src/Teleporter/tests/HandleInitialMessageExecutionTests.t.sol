// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: BSD-3-Clause

pragma solidity 0.8.18;

import "./TeleporterMessengerTest.t.sol";
import "../ITeleporterReceiver.sol";

enum SampleMessageReceiverAction {
    Receive,
    ReceiveRecursive
}

contract SampleMessageReceiver is ITeleporterReceiver {
    address public immutable teleporterContract;
    string public latestMessage;
    bytes32 public latestMessageSenderSubnetID;
    address public latestMessageSenderAddress;

    constructor(address teleporterContractAddress) {
        teleporterContract = teleporterContractAddress;
    }

    function receiveTeleporterMessage(
        bytes32 originChainID,
        address originSenderAddress,
        bytes calldata message
    ) external {
        require(msg.sender == teleporterContract, "Unauthorized.");
        // Decode the payload to recover the action and corresponding function parameters
        (SampleMessageReceiverAction action, bytes memory actionData) = abi
            .decode(message, (SampleMessageReceiverAction, bytes));
        if (action == SampleMessageReceiverAction.Receive) {
            (string memory messageString, bool succeed) = abi.decode(
                actionData,
                (string, bool)
            );
            _receiveMessage(
                originChainID,
                originSenderAddress,
                messageString,
                succeed
            );
        } else if (action == SampleMessageReceiverAction.ReceiveRecursive) {
            string memory messageString = abi.decode(actionData, (string));
            _receiveMessageRecursive(
                originChainID,
                originSenderAddress,
                messageString
            );
        } else {
            revert("Invalid action.");
        }
    }

    // Stores the message in this contract to be fetched by anyone.
    function _receiveMessage(
        bytes32 originChainID,
        address originSenderAddress,
        string memory message,
        bool succeed
    ) internal {
        require(msg.sender == teleporterContract, "Unauthorized.");
        require(succeed, "Intended to fail.");
        latestMessage = message;
        latestMessageSenderSubnetID = originChainID;
        latestMessageSenderAddress = originSenderAddress;
    }

    // Tries to recursively call the teleporterContract to receive a message, which should always fail.
    function _receiveMessageRecursive(
        bytes32 originChainID,
        address originSenderAddress,
        string memory message
    ) internal {
        require(msg.sender == teleporterContract, "Unauthorized.");
        ITeleporterMessenger messenger = ITeleporterMessenger(
            teleporterContract
        );
        messenger.receiveCrossChainMessage(address(42));
        latestMessage = message;
        latestMessageSenderSubnetID = originChainID;
        latestMessageSenderAddress = originSenderAddress;
    }
}

contract HandleInitialMessageExecutionTest is TeleporterMessengerTest {
    SampleMessageReceiver public destinationContract;

    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
        destinationContract = new SampleMessageReceiver(
            address(teleporterMessenger)
        );
    }

    function testSuccess() public {
        // Construct the mock message to be received.
        string memory messageString = "Testing successful message";
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageID: 42,
            senderAddress: address(this),
            destinationAddress: address(destinationContract),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: abi.encode(
                SampleMessageReceiverAction.Receive,
                abi.encode(messageString, true)
            )
        });
        WarpMessage memory warpMessage = _createDefaultWarpMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            abi.encode(messageToReceive)
        );

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(warpMessage);

        // Receive the message.
        teleporterMessenger.receiveCrossChainMessage(
            DEFAULT_RELAYER_REWARD_ADDRESS
        );

        // Check that the message had the proper affect on the destination contract.
        assertEq(destinationContract.latestMessage(), messageString);
        assertEq(
            destinationContract.latestMessageSenderSubnetID(),
            DEFAULT_ORIGIN_CHAIN_ID
        );
        assertEq(
            destinationContract.latestMessageSenderAddress(),
            address(this)
        );
        assertEq(
            teleporterMessenger.getRelayerRewardAddress(
                DEFAULT_ORIGIN_CHAIN_ID,
                messageToReceive.messageID
            ),
            DEFAULT_RELAYER_REWARD_ADDRESS
        );
    }

    function testInsufficientGasProvided() public {
        // Construct the mock message to be received.
        string memory messageString = "Testing successful message";
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageID: 42,
            senderAddress: address(this),
            destinationAddress: address(destinationContract),
            requiredGasLimit: uint256(
                bytes32(
                    hex"FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
                ) // UINT256_MAX
            ),
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: abi.encode(
                SampleMessageReceiverAction.Receive,
                abi.encode(messageString, true)
            )
        });
        WarpMessage memory warpMessage = _createDefaultWarpMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            abi.encode(messageToReceive)
        );

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(warpMessage);

        // Receive the message.
        vm.expectRevert("Insufficient gas provided to execute message.");
        teleporterMessenger.receiveCrossChainMessage(
            DEFAULT_RELAYER_REWARD_ADDRESS
        );
    }

    function testCannotReceiveMessageRecursively() public {
        // Construct the mock message to be received.
        string memory messageString = "Testing successful message";
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageID: 42,
            senderAddress: address(this),
            destinationAddress: address(destinationContract),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: abi.encode(
                SampleMessageReceiverAction.ReceiveRecursive,
                abi.encode(messageString)
            )
        });
        WarpMessage memory warpMessage = _createDefaultWarpMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            abi.encode(messageToReceive)
        );

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(warpMessage);

        // Receive the message - this does not revert because the recursive call
        // is considered a failed message execution, but the message itself is
        // still successfully delivered.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit FailedMessageExecution(
            DEFAULT_ORIGIN_CHAIN_ID,
            messageToReceive.messageID,
            messageToReceive
        );
        teleporterMessenger.receiveCrossChainMessage(
            DEFAULT_RELAYER_REWARD_ADDRESS
        );

        // Check that the message hash was stored in state and the message did not have any affect on the destination.
        assertEq(destinationContract.latestMessage(), "");
        assertEq(destinationContract.latestMessageSenderSubnetID(), bytes32(0));
        assertEq(destinationContract.latestMessageSenderAddress(), address(0));
        assertEq(
            teleporterMessenger.getRelayerRewardAddress(
                DEFAULT_ORIGIN_CHAIN_ID,
                messageToReceive.messageID
            ),
            DEFAULT_RELAYER_REWARD_ADDRESS
        );
        vm.expectRevert("Message retry execution attempt failed.");
        teleporterMessenger.retryMessageExecution(
            DEFAULT_ORIGIN_CHAIN_ID,
            messageToReceive
        );
    }

    function testStoreHashOfFailedMessageExecution() public {
        // Construct the mock message to be received.
        string memory messageString = "Testing successful message";
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageID: 42,
            senderAddress: address(this),
            destinationAddress: address(destinationContract),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: abi.encode(
                SampleMessageReceiverAction.Receive,
                abi.encode(messageString, false)
            )
        });
        WarpMessage memory warpMessage = _createDefaultWarpMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            abi.encode(messageToReceive)
        );

        // Mock the call to the warp precompile to get the message.
        _setUpSuccessGetVerifiedWarpMessageMock(warpMessage);

        // Receive the message.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit FailedMessageExecution(
            DEFAULT_ORIGIN_CHAIN_ID,
            messageToReceive.messageID,
            messageToReceive
        );
        teleporterMessenger.receiveCrossChainMessage(
            DEFAULT_RELAYER_REWARD_ADDRESS
        );

        // Check that the message hash was stored in state and the message did not have any affect on the destination.
        assertEq(destinationContract.latestMessage(), "");
        assertEq(destinationContract.latestMessageSenderSubnetID(), bytes32(0));
        assertEq(destinationContract.latestMessageSenderAddress(), address(0));
        assertEq(
            teleporterMessenger.getRelayerRewardAddress(
                DEFAULT_ORIGIN_CHAIN_ID,
                messageToReceive.messageID
            ),
            DEFAULT_RELAYER_REWARD_ADDRESS
        );
        vm.expectRevert("Message retry execution attempt failed.");
        teleporterMessenger.retryMessageExecution(
            DEFAULT_ORIGIN_CHAIN_ID,
            messageToReceive
        );
    }
}
