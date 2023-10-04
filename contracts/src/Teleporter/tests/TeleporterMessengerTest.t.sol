// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "forge-std/Test.sol";
import "../TeleporterMessenger.sol";
import "../../Mocks/UnitTestMockERC20.sol";

// Parent contract for TeleporterMessenger tests. Deploys a TeleporterMessenger
// instance in the test setup, and provides helper methods for sending and
// receiving empty mock messages.
contract TeleporterMessengerTest is Test {
    TeleporterMessenger public teleporterMessenger;

    bytes32 public constant MOCK_BLOCK_CHAIN_ID = bytes32(uint256(123456));
    bytes32 public constant DEFAULT_ORIGIN_CHAIN_ID =
        bytes32(
            hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd"
        );
    bytes32 public constant DEFAULT_DESTINATION_CHAIN_ID =
        bytes32(
            hex"1234567812345678123456781234567812345678123456781234567812345678"
        );
    address public constant DEFAULT_DESTINATION_ADDRESS =
        0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    uint256 public constant DEFAULT_REQUIRED_GAS_LIMIT = 1e6;
    address public constant WARP_PRECOMPILE_ADDRESS =
        0x0200000000000000000000000000000000000005;
    address public constant DEFAULT_RELAYER_REWARD_ADDRESS =
        0xa4CEE7d1aF6aDdDD33E3b1cC680AB84fdf1b6d1d;

    UnitTestMockERC20 internal _mockFeeAsset;

    event SendCrossChainMessage(
        bytes32 indexed destinationChainID,
        uint256 indexed messageID,
        TeleporterMessage message,
        TeleporterFeeInfo feeInfo
    );

    event AddFeeAmount(
        bytes32 indexed destinationChainID,
        uint256 indexed messageID,
        TeleporterFeeInfo updatedFeeInfo
    );

    event FailedMessageExecution(
        bytes32 indexed originChainID,
        uint256 indexed messageID,
        TeleporterMessage message
    );

    event MessageExecutionRetried(
        bytes32 indexed originChainID,
        uint256 indexed messageID
    );

    event FailedFeePayment(
        bytes32 indexed destinationChainID,
        uint256 indexed messageID,
        address indexed feeAsset,
        uint256 feeAmount,
        address relayerRewardAddress
    );

    function setUp() public virtual {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(WarpMessenger.getBlockchainID.selector),
            abi.encode(MOCK_BLOCK_CHAIN_ID)
        );

        teleporterMessenger = new TeleporterMessenger();
        _mockFeeAsset = new UnitTestMockERC20();
    }

    function testCannotDeployWithoutWarpPrecompile() public {
        vm.clearMockedCalls();
        vm.expectRevert();
        teleporterMessenger = new TeleporterMessenger();
    }

    /*
     * TEST UTILS / HELPERS
     */

    function _sendTestMessageWithFee(
        bytes32 chainID,
        uint256 feeAmount
    ) internal returns (uint256) {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(WarpMessenger.sendWarpMessage.selector),
            new bytes(0)
        );

        address feeAsset = address(0);
        if (feeAmount > 0) {
            feeAsset = address(_mockFeeAsset);
        }
        TeleporterFeeInfo memory feeInfo = TeleporterFeeInfo({
            contractAddress: feeAsset,
            amount: feeAmount
        });

        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationChainID: chainID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: feeInfo,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        if (feeAmount > 0) {
            // Expect the ERC20 contract transferFrom method to be called to transfer the fee.
            vm.expectCall(
                address(_mockFeeAsset),
                abi.encodeCall(
                    IERC20.transferFrom,
                    (address(this), address(teleporterMessenger), feeAmount)
                )
            );
        }

        return teleporterMessenger.sendCrossChainMessage(messageInput);
    }

    function _sendTestMessageWithNoFee(
        bytes32 chainID
    ) internal returns (uint256) {
        return _sendTestMessageWithFee(chainID, 0);
    }

    function _setUpSuccessGetVerifiedWarpMessageMock(
        uint32 index,
        WarpMessage memory warpMessage
    ) internal {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(WarpMessenger.getVerifiedWarpMessage, (index)),
            abi.encode(warpMessage, true)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(WarpMessenger.getVerifiedWarpMessage, (index))
        );
    }

    function _receiveTestMessage(
        bytes32 originChainID,
        uint256 messageID,
        address relayerRewardAddress,
        TeleporterMessageReceipt[] memory receipts
    ) internal {
        // Construct the test message to be received.
        TeleporterMessage
            memory messageToReceive = _createMockTeleporterMessage(
                messageID,
                new bytes(0)
            );
        messageToReceive.receipts = receipts;

        // Both the sender and destination address should be the teleporter contract address,
        // mocking as if the universal deployer pattern was used.
        WarpMessage memory warpMessage = _createDefaultWarpMessage(
            originChainID,
            abi.encode(messageToReceive)
        );

        // We have to mock the precompile call so that it doesn't revert in the tests.
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Receive the message.
        teleporterMessenger.receiveCrossChainMessage(0, relayerRewardAddress);
    }

    function _receiveMessageSentToEOA()
        internal
        returns (bytes32, address, TeleporterMessage memory)
    {
        // Create valid mock message data, we will mock out the execution of the
        // message itself so it doesn't matter what it is.
        bytes memory messageData = new bytes(1);

        // Construct the test message to be received. By default, the destination
        // address will be the DEFAULT_DESTINATION_ADDRESS.
        TeleporterMessage
            memory messageToReceive = _createMockTeleporterMessage(
                13,
                messageData
            );
        WarpMessage memory warpMessage = _createDefaultWarpMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            abi.encode(messageToReceive)
        );

        // We have to mock the precompile call so that it doesn't revert in the tests.
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        // Mock there being no code at the destination address which will be called
        // when the message is executed. This mimics the default destination address being an EOA.
        vm.etch(DEFAULT_DESTINATION_ADDRESS, new bytes(0));

        // Receive the message - which should fail execution.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit FailedMessageExecution(
            DEFAULT_ORIGIN_CHAIN_ID,
            messageToReceive.messageID,
            messageToReceive
        );
        teleporterMessenger.receiveCrossChainMessage(
            0,
            DEFAULT_RELAYER_REWARD_ADDRESS
        );

        return (
            DEFAULT_ORIGIN_CHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            messageToReceive
        );
    }

    // Create a mock message to be used in tests. It should include no receipts
    // because receipts are only valid if they refer to messages previously send
    // to the origin chain.
    function _createMockTeleporterMessage(
        uint256 messageID,
        bytes memory message
    ) internal view returns (TeleporterMessage memory) {
        return
            TeleporterMessage({
                messageID: messageID,
                senderAddress: address(this),
                destinationAddress: DEFAULT_DESTINATION_ADDRESS,
                requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
                allowedRelayerAddresses: new address[](0),
                receipts: new TeleporterMessageReceipt[](0),
                message: message
            });
    }

    function _createDefaultWarpMessage(
        bytes32 originChainID,
        bytes memory payload
    ) internal view returns (WarpMessage memory) {
        return
            WarpMessage({
                sourceChainID: originChainID,
                originSenderAddress: address(teleporterMessenger),
                destinationChainID: MOCK_BLOCK_CHAIN_ID,
                destinationAddress: address(teleporterMessenger),
                payload: payload
            });
    }
}
