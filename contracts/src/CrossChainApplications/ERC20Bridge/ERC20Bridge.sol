// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "./IERC20Bridge.sol";
import "./BridgeToken.sol";
import "../../Teleporter/ITeleporterMessenger.sol";
import "../../Teleporter/ITeleporterReceiver.sol";
import "../../Teleporter/SafeERC20TransferFrom.sol";
import "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

struct TokenID {
    bytes32 chainID;
    address bridgeContract;
    address asset;
}

/**
 * @dev Implementation of the {IERC20Bridge} interface.
 *
 * This implementation uses the {BridgeToken} contract to represent tokens on this chain, and uses
 * {ITeleporterMessenger} to send and receive messages to other chains.
 */
contract ERC20Bridge is IERC20Bridge, ITeleporterReceiver, ReentrancyGuard {
    using SafeERC20 for IERC20;

    struct WrappedTokenTransferInfo {
        bytes32 destinationChainID;
        address destinationBridgeAddress;
        address wrappedContractAddress;
        address recipient;
        uint256 totalAmount;
        uint256 primaryFeeAmount;
        uint256 secondaryFeeAmount;
    }

    address public constant WARP_PRECOMPILE_ADDRESS =
        0x0200000000000000000000000000000000000005;
    bytes32 public immutable currentChainID;

    // Used for sending an receiving Teleporter messages.
    ITeleporterMessenger public immutable teleporterMessenger;

    // Tracks which bridge tokens have been submitted to be created other bridge instances.
    // (destinationChainID, destinationBridgeAddress) -> nativeTokenContract -> tokenCreationSubmitted
    // Note that the existence of a bridge token in this mapping does not ensure that it exists on
    // the destination bridge because the message to create the new token may not have been
    // successfully delivered yet.
    mapping(bytes32 => mapping(address => mapping(address => bool)))
        public submittedBridgeTokenCreations;

    // Tracks the balances of native tokens sent to other bridge instances.
    // Bridges are not allowed to unwrap more than has been sent to them.
    // (destinationChainID, destinationBridgeAddress) -> nativeTokenContract -> balance
    mapping(bytes32 => mapping(address => mapping(address => uint256)))
        public bridgedBalances;

    // Set of bridge tokens created by this bridge instance.
    mapping(address => bool) public wrappedTokenContracts;

    // Tracks the wrapped bridge token contract address for each native token bridged to this bridge instance.
    // (nativeChainID, nativeBridgeAddress, nativeTokenAddress) -> bridgeTokenAddress
    mapping(bytes32 => mapping(address => mapping(address => address)))
        public nativeToWrappedTokens;

    uint256 public constant CREATE_BRIDGE_TOKENS_REQUIRED_GAS = 2_000_000;
    uint256 public constant MINT_BRIDGE_TOKENS_REQUIRED_GAS = 200_000;
    uint256 public constant TRANSFER_BRIDGE_TOKENS_REQUIRED_GAS = 300_000;

    /**
     * @dev Initializes the Teleporter messenger used for sending and receiving messages,
     * and initializes the current chain ID.
     */
    constructor(address teleporterMessengerAddress) {
        require(
            teleporterMessengerAddress != address(0),
            "Invalid teleporter messenger address"
        );
        teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);
        currentChainID = WarpMessenger(WARP_PRECOMPILE_ADDRESS)
            .getBlockchainID();
    }

    /**
     * @dev See {IERC20Bridge-bridgeTokens}.
     *
     * Requirements:
     *
     * - `destinationChainID` cannot be the same as the current chain ID.
     * - For wrapped tokens, `totalAmount` must be greater than the sum of the primary and secondary fee amounts.
     * - For native tokens, `adjustedAmount` after safe transfer must be greater than the primary fee amount.
     */
    function bridgeTokens(
        bytes32 destinationChainID,
        address destinationBridgeAddress,
        address tokenContractAddress,
        address recipient,
        uint256 totalAmount,
        uint256 primaryFeeAmount,
        uint256 secondaryFeeAmount
    ) external nonReentrant {
        // Bridging tokens within a single chain is not allowed.
        require(
            destinationChainID != currentChainID,
            "Cannot bridge token within same chain."
        );

        // Neither the recipient nor the destination bridge can be the zero address.
        require(
            recipient != address(0),
            "Recipient address cannot be zero address."
        );

        require(
            destinationBridgeAddress != address(0),
            "Destination bridge address cannot be zero address."
        );

        // If the token to be bridged is an existing wrapped token of this bridge,
        // then handle it as an "unwrap" by burning the tokens, and sending a message
        // back to the native chain of the token.
        // Otherwise, handle it as a "wrap" by locking the tokens in this bridge instance,
        // and sending a message to the destination to mint new tokens.
        if (wrappedTokenContracts[tokenContractAddress]) {
            // The fee amounts are taken out of the total amount to be transferred.
            // In the wrapped token case, we know that the bridgeToken to be burned
            // is not a "fee/burn on transfer" token, since it was deployed by this
            // contract itself.
            require(
                totalAmount > primaryFeeAmount + secondaryFeeAmount,
                "Fee amounts more than total."
            );

            return
                _processWrappedTokenTransfer(
                    WrappedTokenTransferInfo({
                        destinationChainID: destinationChainID,
                        destinationBridgeAddress: destinationBridgeAddress,
                        wrappedContractAddress: tokenContractAddress,
                        recipient: recipient,
                        totalAmount: totalAmount,
                        primaryFeeAmount: primaryFeeAmount,
                        secondaryFeeAmount: secondaryFeeAmount
                    })
                );
        }

        // Otherwise, this is a token "native" to this chain.
        require(
            submittedBridgeTokenCreations[destinationChainID][
                destinationBridgeAddress
            ][tokenContractAddress],
            "Bridge token has not been previously submitted for creation."
        );

        // Lock tokens in this bridge instance. Supports "fee/burn on transfer" ERC20 token
        // implementations by only bridging the actual balance increase reflected by the call
        // to transferFrom.
        uint256 adjustedAmount = SafeERC20TransferFrom.safeTransferFrom(
            IERC20(tokenContractAddress),
            totalAmount
        );

        // Ensure that the adjusted amount is greater than the fee to be paid.
        // The secondary fee amount is not used in this case (and can assumed to be 0) since bridging
        // a native token to another chain only ever involves a single cross-chain message.
        require(
            adjustedAmount > primaryFeeAmount,
            "Fee amount more than adjusted transfer amount."
        );

        return
            _processNativeTokenTransfer({
                destinationChainID: destinationChainID,
                destinationBridgeAddress: destinationBridgeAddress,
                nativeContractAddress: tokenContractAddress,
                recipient: recipient,
                totalAmount: adjustedAmount,
                feeAmount: primaryFeeAmount
            });
    }

    /**
     * @dev See {IERC20Bridge-submitCreateBridgeToken}.
     *
     * We allow for `submitCreateBridgeToken` to be called multiple times with the same bridge and token
     * information because a previous message may have been dropped or otherwise selectively not delivered.
     * If the bridge token already exists on the destination, we are sending a message that will
     * simply have no effect on the destination.
     *
     * Emits a {SubmitCreateBridgeToken} event.
     */
    function submitCreateBridgeToken(
        bytes32 destinationChainID,
        address destinationBridgeAddress,
        ERC20 nativeToken,
        address messageFeeAsset,
        uint256 messageFeeAmount
    ) external nonReentrant {
        require(
            destinationBridgeAddress != address(0),
            "Destination bridge address cannot be zero address."
        );

        // For non-zero fee amounts, transfer the fee into the control of this contract first, and then
        // allow the Teleporter contract to spend it.
        uint256 adjustedFeeAmount = 0;
        if (messageFeeAmount > 0) {
            adjustedFeeAmount = SafeERC20TransferFrom.safeTransferFrom(
                IERC20(messageFeeAsset),
                messageFeeAmount
            );
            IERC20(messageFeeAsset).safeIncreaseAllowance(
                address(teleporterMessenger),
                adjustedFeeAmount
            );
        }

        // Create the calldata to create the bridge token on the destination chain.
        bytes memory messageData = encodeCreateBridgeTokenData(
            address(nativeToken),
            nativeToken.name(),
            nativeToken.symbol(),
            nativeToken.decimals()
        );

        // Send Teleporter message.
        uint256 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationChainID: destinationChainID,
                destinationAddress: destinationBridgeAddress,
                feeInfo: TeleporterFeeInfo({
                    contractAddress: messageFeeAsset,
                    amount: adjustedFeeAmount
                }),
                requiredGasLimit: CREATE_BRIDGE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: new address[](0),
                message: messageData
            })
        );

        submittedBridgeTokenCreations[destinationChainID][
            destinationBridgeAddress
        ][address(nativeToken)] = true;

        emit SubmitCreateBridgeToken(
            destinationChainID,
            destinationBridgeAddress,
            address(nativeToken),
            messageID
        );
    }

    /**
     * @dev See {ITeleporterReceiver-receiveTeleporterMessage}.
     *
     * Receives a Teleporter message and routes to the appropriate internal function call.
     */
    function receiveTeleporterMessage(
        bytes32 nativeChainID,
        address nativeBridgeAddress,
        bytes calldata message
    ) external {
        // Only allow the Teleporter messenger to deliver messages.
        require(msg.sender == address(teleporterMessenger), "Unauthorized.");

        // Decode the payload to recover the action and corresponding function parameters
        (BridgeAction action, bytes memory actionData) = abi.decode(
            message,
            (BridgeAction, bytes)
        );

        // Route to the appropriate function.
        if (action == BridgeAction.Create) {
            (
                address nativeContractAddress,
                string memory nativeName,
                string memory nativeSymbol,
                uint8 nativeDecimals
            ) = abi.decode(actionData, (address, string, string, uint8));
            _createBridgeToken({
                nativeChainID: nativeChainID,
                nativeBridgeAddress: nativeBridgeAddress,
                nativeContractAddress: nativeContractAddress,
                nativeName: nativeName,
                nativeSymbol: nativeSymbol,
                nativeDecimals: nativeDecimals
            });
        } else if (action == BridgeAction.Mint) {
            (
                address nativeContractAddress,
                address recipient,
                uint256 amount
            ) = abi.decode(actionData, (address, address, uint256));
            _mintBridgeTokens(
                nativeChainID,
                nativeBridgeAddress,
                nativeContractAddress,
                recipient,
                amount
            );
        } else if (action == BridgeAction.Transfer) {
            (
                bytes32 destinationChainID,
                address destinationBridgeAddress,
                address nativeContractAddress,
                address recipient,
                uint256 totalAmount,
                uint256 secondaryFeeAmount
            ) = abi.decode(
                    actionData,
                    (bytes32, address, address, address, uint256, uint256)
                );
            _transferBridgeTokens({
                sourceChainID: nativeChainID,
                sourceBridgeAddress: nativeBridgeAddress,
                destinationChainID: destinationChainID,
                destinationBridgeAddress: destinationBridgeAddress,
                nativeContractAddress: nativeContractAddress,
                recipient: recipient,
                totalAmount: totalAmount,
                secondaryFeeAmount: secondaryFeeAmount
            });
        } else {
            revert("Invalid action.");
        }
    }

    /**
     * @dev Encodes the parameters for the Create action to be decoded and executed on the destination.
     */
    function encodeCreateBridgeTokenData(
        address nativeContractAddress,
        string memory nativeName,
        string memory nativeSymbol,
        uint8 nativeDecimals
    ) public pure returns (bytes memory) {
        // ABI encode the Create action and corresponding parameters for the createBridgeToken
        // call to to be decoded and executed on the destination.
        bytes memory paramsData = abi.encode(
            nativeContractAddress,
            nativeName,
            nativeSymbol,
            nativeDecimals
        );
        return abi.encode(BridgeAction.Create, paramsData);
    }

    /**
     * @dev Encodes the parameters for the Mint action to be decoded and executed on the destination.
     */
    function encodeMintBridgeTokensData(
        address nativeContractAddress,
        address recipient,
        uint256 bridgeAmount
    ) public pure returns (bytes memory) {
        // ABI encode the Mint action and corresponding parameters for the mintBridgeTokens
        // call to to be decoded and executed on the destination.
        bytes memory paramsData = abi.encode(
            nativeContractAddress,
            recipient,
            bridgeAmount
        );
        return abi.encode(BridgeAction.Mint, paramsData);
    }

    /**
     * @dev Encodes the parameters for the Transfer action to be decoded and executed on the destination.
     */
    function encodeTransferBridgeTokensData(
        bytes32 destinationChainID,
        address destinationBridgeAddress,
        address nativeContractAddress,
        address recipient,
        uint256 amount,
        uint256 feeAmount
    ) public pure returns (bytes memory) {
        // ABI encode the Transfer action and corresponding parameters for the transferBridgeToken
        // call to to be decoded and executed on the destination.
        bytes memory paramsData = abi.encode(
            destinationChainID,
            destinationBridgeAddress,
            nativeContractAddress,
            recipient,
            amount,
            feeAmount
        );
        return abi.encode(BridgeAction.Transfer, paramsData);
    }

    /**
     * @dev Teleporter message receiver for creating a new bridge token on this chain.
     *
     * Emits a {CreateBridgeToken} event.
     */
    function _createBridgeToken(
        bytes32 nativeChainID,
        address nativeBridgeAddress,
        address nativeContractAddress,
        string memory nativeName,
        string memory nativeSymbol,
        uint8 nativeDecimals
    ) private {
        // Check that the bridge token doesn't already exist.
        require(
            nativeToWrappedTokens[nativeChainID][nativeBridgeAddress][
                nativeContractAddress
            ] == address(0),
            "Bridge token already exists."
        );

        address bridgeTokenAddress = address(
            new BridgeToken({
                sourceChainID: nativeChainID,
                sourceBridge: nativeBridgeAddress,
                sourceAsset: nativeContractAddress,
                tokenName: nativeName,
                tokenSymbol: nativeSymbol,
                tokenDecimals: nativeDecimals
            })
        );

        wrappedTokenContracts[bridgeTokenAddress] = true;
        nativeToWrappedTokens[nativeChainID][nativeBridgeAddress][
            nativeContractAddress
        ] = bridgeTokenAddress;

        emit CreateBridgeToken(
            nativeChainID,
            nativeBridgeAddress,
            nativeContractAddress,
            bridgeTokenAddress
        );
    }

    /**
     * @dev Teleporter message receiver for minting of an existing bridge token on this chain.
     *
     * Emits a {MintBridgeTokens} event.
     */
    function _mintBridgeTokens(
        bytes32 nativeChainID,
        address nativeBridgeAddress,
        address nativeContractAddress,
        address recipient,
        uint256 amount
    ) private nonReentrant {
        // Only allow the Teleporter messenger to deliver messages.
        require(msg.sender == address(teleporterMessenger), "Unauthorized.");

        // The recipient cannot be the zero address.
        require(
            recipient != address(0),
            "Recipient address cannot be zero address."
        );

        // Check that a bridge token exists for this native asset.
        // If not, one needs to be created by the delivery of a "createBridgeToken" message first
        // before this mint can be processed. Once the bridge token is create, this message
        // could then be retried to mint the tokens.
        address bridgeTokenAddress = nativeToWrappedTokens[nativeChainID][
            nativeBridgeAddress
        ][nativeContractAddress];
        require(
            bridgeTokenAddress != address(0),
            "Bridge token does not yet exist."
        );

        // Mint the wrapped tokens.
        BridgeToken(bridgeTokenAddress).mint(recipient, amount);
        emit MintBridgeTokens(bridgeTokenAddress, recipient, amount);
    }

    /**
     * @dev Teleporter message receiver for handling bridge tokens transfers back from another chain
     * and optionally routing them to a different third chain.
     */
    function _transferBridgeTokens(
        bytes32 sourceChainID,
        address sourceBridgeAddress,
        bytes32 destinationChainID,
        address destinationBridgeAddress,
        address nativeContractAddress,
        address recipient,
        uint256 totalAmount,
        uint256 secondaryFeeAmount
    ) private nonReentrant {
        // Only allow the teleporter messenger to deliver messages.
        require(msg.sender == address(teleporterMessenger), "Unauthorized.");

        // Neither the recipient nor the destination bridge can be the zero address.
        require(
            recipient != address(0),
            "Recipient address cannot be zero address."
        );

        require(
            destinationBridgeAddress != address(0),
            "Destination bridge address cannot be zero address."
        );

        // Check that the bridge returning the tokens has sufficient balance to do so.
        uint256 currentBalance = bridgedBalances[sourceChainID][
            sourceBridgeAddress
        ][nativeContractAddress];
        require(
            currentBalance >= totalAmount,
            "Insufficient wrapped token balance."
        );
        bridgedBalances[sourceChainID][sourceBridgeAddress][
            nativeContractAddress
        ] = currentBalance - totalAmount;

        // If the destination chain ID and bridge is this bridge instance, then release the tokens back to the recipient.
        // In this case, since there is no secondary Teleporter message, the secondary fee amount is not used.
        if (destinationChainID == currentChainID) {
            require(
                destinationBridgeAddress == address(this),
                "Destination bridge address must be this bridge instance."
            );

            // Transfer tokens to the recipient.
            // We don't need have a special case for handling "fee/burn on transfer" ERC20 token implementations
            // here because the amount actually transfered to the user in the ERC20 contract (whether or not
            // it's less than totalAmount) is the amount the user receives from this completed bridge transfer,
            // which is out of control of the bridge contract itself.
            IERC20(nativeContractAddress).safeTransfer(recipient, totalAmount);
            return;
        }

        // Otherwise, re-bridge the tokens on to their ultimate destination.
        // The tokens are already locked in this contract from when they were previously bridged.
        // We deduct the balance from bridge instance that sent this message, and now will increment the
        // balance of the destination bridge instance.
        return
            _processNativeTokenTransfer({
                destinationChainID: destinationChainID,
                destinationBridgeAddress: destinationBridgeAddress,
                nativeContractAddress: nativeContractAddress,
                recipient: recipient,
                totalAmount: totalAmount,
                feeAmount: secondaryFeeAmount
            });
    }

    /**
     * @dev Increments the balance of the native tokens bridged to the specified bridge instance and
     * sends a Teleporter message to have the destination bridge mint the new tokens. The tokens to be
     * bridge must already be locked in this contract before calling.
     *
     * Emits a {BridgeTokens} event.
     * Requirements:
     *
     * - `destinationChainID` cannot be the same as the current chain ID.
     * - can not do nested bridging of wrapped tokens.
     */
    function _processNativeTokenTransfer(
        bytes32 destinationChainID,
        address destinationBridgeAddress,
        address nativeContractAddress,
        address recipient,
        uint256 totalAmount,
        uint256 feeAmount
    ) private {
        // Do not allow nested bridging of wrapped tokens.
        require(
            !wrappedTokenContracts[nativeContractAddress],
            "Cannot bridge wrapped token."
        );

        // Bridging tokens within a single chain is not allowed.
        // This function is called by bridgeTokens and transferBridgeTokens which both already make this check,
        // so this check is redundant but left in for clarity.
        require(
            destinationChainID != currentChainID,
            "Cannot bridge token within same chain."
        );

        // Allow the Teleporter messenger to spend the fee amount.
        if (feeAmount > 0) {
            IERC20(nativeContractAddress).safeIncreaseAllowance(
                address(teleporterMessenger),
                feeAmount
            );
        }

        // Update balances.
        uint256 bridgeAmount = totalAmount - feeAmount;
        bridgedBalances[destinationChainID][destinationBridgeAddress][
            nativeContractAddress
        ] += bridgeAmount;

        // Send Teleporter message.
        bytes memory messageData = encodeMintBridgeTokensData(
            nativeContractAddress,
            recipient,
            bridgeAmount
        );

        uint256 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationChainID: destinationChainID,
                destinationAddress: destinationBridgeAddress,
                feeInfo: TeleporterFeeInfo({
                    contractAddress: nativeContractAddress,
                    amount: feeAmount
                }),
                requiredGasLimit: MINT_BRIDGE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: new address[](0),
                message: messageData
            })
        );

        emit BridgeTokens({
            tokenContractAddress: nativeContractAddress,
            destinationChainID: destinationChainID,
            teleporterMessageID: messageID,
            destinationBridgeAddress: destinationBridgeAddress,
            recipient: recipient,
            amount: bridgeAmount
        });
    }

    /**
     * @dev Processes a wrapped token transfer by burning the tokens and sending a Teleporter message
     * to the native chain and bridge of the wrapped asset that was burned.
     *
     * It is the caller's responsibility to ensure that the wrapped token contract is supported by this bridge instance.
     * Emits a {BridgeTokens} event.
     */
    function _processWrappedTokenTransfer(
        WrappedTokenTransferInfo memory wrappedTransferInfo
    ) private {
        // If necessary, transfer the primary fee amount to this contract and approve the
        // Teleporter messenger to spend it when the first message back to the native subnet
        // is submitted. The secondary fee amount is then handled by the native subnet when
        // submitting a message to the destination chain, if applicable.
        uint256 adjustedPrimaryFeeAmount = 0;
        if (wrappedTransferInfo.primaryFeeAmount > 0) {
            // We know that the ERC20 contract is not a "fee on transfer" or "burn on transfer" contract
            // because it is a BridgeToken contract instance that was deployed by this contract itself.
            // However, we still use safeTransferFrom for completeness.
            adjustedPrimaryFeeAmount = SafeERC20TransferFrom.safeTransferFrom(
                IERC20(wrappedTransferInfo.wrappedContractAddress),
                wrappedTransferInfo.primaryFeeAmount
            );
            IERC20(wrappedTransferInfo.wrappedContractAddress)
                .safeIncreaseAllowance(
                    address(teleporterMessenger),
                    adjustedPrimaryFeeAmount
                );
        }

        // Burn the wrapped tokens to be bridged.
        // The bridge amount is the total amount minus the original fee amount. Even if the adjusted fee amount
        // is less than the original fee amount, the original amount is the portion that is spent out of the total
        // amount. We know that the burnFrom call will decrease the total supply by bridgeAmount because the
        // bridgeToken contract was deployed by this contract itself and does not implement "fee on burn" functionality.
        uint256 bridgeAmount = wrappedTransferInfo.totalAmount -
            wrappedTransferInfo.primaryFeeAmount;
        BridgeToken bridgeToken = BridgeToken(
            wrappedTransferInfo.wrappedContractAddress
        );
        bridgeToken.burnFrom(msg.sender, bridgeAmount);

        // If the destination chain ID is the native chain ID for the wrapped token, the bridge address must also match.
        // This is because you are not allowed to bridge a token within its native chain.
        bytes32 nativeChainID = bridgeToken.nativeChainID();
        address nativeBridgeAddress = bridgeToken.nativeBridge();
        if (wrappedTransferInfo.destinationChainID == nativeChainID) {
            require(
                wrappedTransferInfo.destinationBridgeAddress ==
                    nativeBridgeAddress,
                "Invalid destination bridge address for native chain."
            );
        }

        // Send a message to the native chain and bridge of the wrapped asset that was burned.
        // The message includes the destination chain ID  and bridge contract, which will differ from the native
        // ones in the event that the tokens are being bridge from one non-native chain to another with two hops.
        bytes memory messageData = encodeTransferBridgeTokensData({
            destinationChainID: wrappedTransferInfo.destinationChainID,
            destinationBridgeAddress: wrappedTransferInfo
                .destinationBridgeAddress,
            nativeContractAddress: bridgeToken.nativeAsset(),
            recipient: wrappedTransferInfo.recipient,
            amount: bridgeAmount,
            feeAmount: wrappedTransferInfo.secondaryFeeAmount
        });

        uint256 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationChainID: nativeChainID,
                destinationAddress: nativeBridgeAddress,
                feeInfo: TeleporterFeeInfo({
                    contractAddress: wrappedTransferInfo.wrappedContractAddress,
                    amount: adjustedPrimaryFeeAmount
                }),
                requiredGasLimit: TRANSFER_BRIDGE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: new address[](0),
                message: messageData
            })
        );
        emit BridgeTokens({
            tokenContractAddress: wrappedTransferInfo.wrappedContractAddress,
            destinationChainID: wrappedTransferInfo.destinationChainID,
            teleporterMessageID: messageID,
            destinationBridgeAddress: wrappedTransferInfo
                .destinationBridgeAddress,
            recipient: wrappedTransferInfo.recipient,
            amount: bridgeAmount
        });
    }
}
