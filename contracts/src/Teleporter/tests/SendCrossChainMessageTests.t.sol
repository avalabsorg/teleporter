// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "./TeleporterMessengerTest.t.sol";

contract SendCrossChainMessageTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testSendMessageNoFee() public {
        // Arrange
        TeleporterMessage memory expectedMessage = _createMockTeleporterMessage(
            1,
            hex"deadbeef"
        );
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationChainID: hex"11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff",
            destinationAddress: expectedMessage.destinationAddress,
            feeInfo: TeleporterFeeInfo(address(0), 0),
            requiredGasLimit: expectedMessage.requiredGasLimit,
            allowedRelayerAddresses: expectedMessage.allowedRelayerAddresses,
            message: expectedMessage.message
        });

        // We have to mock the precompile call so that the test does not revert.
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(WarpMessenger.sendWarpMessage.selector),
            new bytes(0)
        );

        // Expect the exact message to be passed to the precompile.
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.sendWarpMessage,
                (
                    messageInput.destinationChainID,
                    address(teleporterMessenger),
                    abi.encode(expectedMessage)
                )
            )
        );

        // Expect the SendCrossChainMessage event to be emitted.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            messageInput.destinationChainID,
            expectedMessage.messageID,
            expectedMessage
        );

        // Act
        uint256 messageID = teleporterMessenger.sendCrossChainMessage(
            messageInput
        );

        // Assert
        assertEq(messageID, 1);
    }

    function testSendMessageWithFee() public {
        // Arrange
        // Construct the message to submit.
        TeleporterMessage memory expectedMessage = _createMockTeleporterMessage(
            1,
            hex"deadbeef"
        );
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationChainID: hex"11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff",
            destinationAddress: expectedMessage.destinationAddress,
            feeInfo: TeleporterFeeInfo(
                address(_mockFeeAsset),
                13131313131313131313
            ),
            requiredGasLimit: expectedMessage.requiredGasLimit,
            allowedRelayerAddresses: expectedMessage.allowedRelayerAddresses,
            message: expectedMessage.message
        });

        // We have to mock the precompile call so that the test does not revert.
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(WarpMessenger.sendWarpMessage.selector),
            new bytes(0)
        );

        // Expect the exact message to be passed to the precompile.
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(
                WarpMessenger.sendWarpMessage,
                (
                    messageInput.destinationChainID,
                    address(teleporterMessenger),
                    abi.encode(expectedMessage)
                )
            )
        );

        // Expect the SendCrossChainMessage event to be emitted.
        vm.expectEmit(true, true, true, true, address(teleporterMessenger));
        emit SendCrossChainMessage(
            messageInput.destinationChainID,
            expectedMessage.messageID,
            expectedMessage
        );

        // Expect the ERC20 contract transferFrom method to be called to transfer the fee.
        vm.expectCall(
            messageInput.feeInfo.contractAddress,
            abi.encodeCall(
                IERC20.transferFrom,
                (
                    address(this),
                    address(teleporterMessenger),
                    messageInput.feeInfo.amount
                )
            )
        );

        // Act
        uint256 messageID = teleporterMessenger.sendCrossChainMessage(
            messageInput
        );

        // Assert
        assertEq(messageID, 1);
    }

    function testFeeAssetDoesNotExist() public {
        address invalidFeeAsset = 0xb8be9140D8717f4a8fd7e8ae23C5668bc3A4B39c;
        uint256 feeAmount = 4567;
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationChainID: DEFAULT_DESTINATION_CHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: TeleporterFeeInfo(invalidFeeAsset, feeAmount),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        // We expect a call to the fee asset address, but since we did not
        // mock it there is no code there to execute.
        vm.expectCall(
            address(invalidFeeAsset),
            abi.encodeCall(IERC20.balanceOf, (address(teleporterMessenger)))
        );
        vm.expectRevert();
        teleporterMessenger.sendCrossChainMessage(messageInput);
    }

    function testFeeTransferFailure() public {
        uint256 feeAmount = 4567;
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationChainID: DEFAULT_DESTINATION_CHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: TeleporterFeeInfo(address(_mockFeeAsset), feeAmount),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        vm.mockCall(
            address(_mockFeeAsset),
            abi.encodeCall(
                IERC20.transferFrom,
                (address(this), address(teleporterMessenger), feeAmount)
            ),
            abi.encode(false)
        );
        vm.expectCall(
            address(_mockFeeAsset),
            abi.encodeCall(
                IERC20.transferFrom,
                (address(this), address(teleporterMessenger), feeAmount)
            )
        );
        vm.expectRevert("SafeERC20: ERC20 operation did not succeed");
        teleporterMessenger.sendCrossChainMessage(messageInput);
    }

    function testInvalidFeeAssetFailure() public {
        address invalidFeeAsset = address(0);
        uint256 feeAmount = 4567;
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationChainID: DEFAULT_DESTINATION_CHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: TeleporterFeeInfo(invalidFeeAsset, feeAmount),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        vm.expectRevert("Invalid fee asset contract address.");
        teleporterMessenger.sendCrossChainMessage(messageInput);
    }
}
