package tests

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

func RetrySuccessfulExecutionGinkgo() {
	RetrySuccessfulExecution(&network.LocalNetwork{})
}

func RetrySuccessfulExecution(network network.Network) {
	var (
		teleporterMessageID *big.Int
	)

	subnets := network.GetSubnetsInfo()
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	subnetATeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetAInfo.ChainRPCClient)
	Expect(err).Should(BeNil())
	subnetBTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetBInfo.ChainRPCClient)
	Expect(err).Should(BeNil())

	//
	// Deploy ExampleMessenger to Subnets A and B
	//
	ctx := context.Background()

	_, subnetAExampleMessenger := utils.DeployExampleCrossChainMessenger(ctx, fundedAddress, fundedKey, subnetAInfo)
	exampleMessengerContractB, subnetBExampleMessenger := utils.DeployExampleCrossChainMessenger(ctx, fundedAddress, fundedKey, subnetBInfo)

	//
	// Call the example messenger contract on Subnet A
	//
	message := "Hello, world!"
	optsA := utils.CreateTransactorOpts(ctx, subnetAInfo, fundedAddress, fundedKey)
	tx, err := subnetAExampleMessenger.SendMessage(optsA, subnetBInfo.BlockchainID, exampleMessengerContractB, fundedAddress, big.NewInt(0), big.NewInt(300000), message)

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, subnetAInfo.ChainRPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	event, err := utils.GetEventFromLogs(receipt.Logs, subnetATeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationChainID[:]).Should(Equal(subnetBInfo.BlockchainID[:]))

	teleporterMessageID = event.Message.MessageID

	//
	// Relay the message to the destination
	//

	receipt = network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, false, true)
	receiveEvent, err := utils.GetEventFromLogs(receipt.Logs, subnetBTeleporterMessenger.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())
	deliveredTeleporterMessage := receiveEvent.Message

	//
	// Check Teleporter message received on the destination
	//
	delivered, err := subnetBTeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, teleporterMessageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	//
	// Verify we received the expected string
	//
	_, currMessage, err := subnetBExampleMessenger.GetCurrentMessage(&bind.CallOpts{}, subnetAInfo.BlockchainID)
	Expect(currMessage).Should(Equal(message))

	//
	// Attempt to retry message execution, which should fail
	//
	optsB := utils.CreateTransactorOpts(ctx, subnetBInfo, fundedAddress, fundedKey)
	tx, err = subnetBTeleporterMessenger.RetryMessageExecution(optsB, subnetAInfo.BlockchainID, deliveredTeleporterMessage)
	Expect(err).Should(Not(BeNil()))
	Expect(tx).Should(BeNil())
}
