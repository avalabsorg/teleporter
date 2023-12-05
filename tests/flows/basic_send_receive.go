package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

// Tests basic one-way send from Subnet A to Subnet B and vice versa
func BasicSendReceive(network interfaces.Network) {
	subnetAInfo, subnetBInfo, _ := utils.GetThreeSubnets(network)
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	//
	ctx := context.Background()

	feeAmount := big.NewInt(1)
	feeTokenAddress, feeToken := utils.DeployExampleERC20(
		ctx,
		fundedKey,
		subnetAInfo,
	)
	utils.ERC20Approve(
		ctx,
		feeToken,
		teleporterContractAddress,
		big.NewInt(0).Mul(big.NewInt(1e18),
			big.NewInt(10)),
		subnetAInfo,
		fundedKey,
	)

	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: subnetBInfo.BlockchainID,
		DestinationAddress:      fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: feeTokenAddress,
			Amount:          feeAmount,
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	receipt, teleporterMessageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		subnetAInfo,
		subnetBInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)

	//
	// Relay the message to the destination
	//
	deliveryReceipt := network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)
	receiveEvent, err := utils.GetEventFromLogs(
		deliveryReceipt.Logs,
		subnetBInfo.TeleporterMessenger.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())

	//
	// Check Teleporter message received on the destination
	//
	delivered, err := subnetBInfo.TeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, subnetAInfo.BlockchainID, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	//
	// Send a transaction to Subnet B to issue a Warp Message from the Teleporter contract to Subnet A
	//
	sendCrossChainMessageInput.DestinationBlockchainID = subnetAInfo.BlockchainID
	sendCrossChainMessageInput.FeeInfo.Amount = big.NewInt(0)
	receipt, teleporterMessageID = utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		subnetBInfo,
		subnetAInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)

	//
	// Relay the message to the destination
	//
	network.RelayMessage(ctx, receipt, subnetBInfo, subnetAInfo, true)

	//
	// Check Teleporter message received on the destination
	//
	delivered, err = subnetAInfo.TeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, subnetBInfo.BlockchainID, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// If the reward address of the message from A->B is the funded address, which is able to send
	// transactions on subnet A, then redeem the rewards.
	if receiveEvent.RewardRedeemer == fundedAddress {
		utils.RedeemRelayerRewardsAndConfirm(
			ctx, subnetAInfo, feeToken, feeTokenAddress, fundedKey, feeAmount,
		)
	}
}
