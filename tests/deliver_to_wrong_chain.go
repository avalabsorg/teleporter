package tests

import (
	"context"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func DeliverToWrongChainGinkgo() {
	DeliverToWrongChain(&network.LocalNetwork{})
}

func DeliverToWrongChain(network network.Network) {
	subnets := network.GetSubnetsInfo()
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	subnetATeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress, subnetAInfo.ChainRPCClient,
	)
	Expect(err).Should(BeNil())
	subnetBTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress, subnetBInfo.ChainRPCClient,
	)
	Expect(err).Should(BeNil())

	//
	// Get the expected teleporter message ID for Subnet B
	//
	teleporterMessageID, err := subnetATeleporterMessenger.GetNextMessageID(&bind.CallOpts{}, subnetBInfo.BlockchainID)
	Expect(err).Should(BeNil())

	//
	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to a chain other than Subnet B
	//
	ctx := context.Background()

	destinationAddressKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	destinationAddress := crypto.PubkeyToAddress(destinationAddressKey.PublicKey)

	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationChainID: ids.Empty, // Some other chain ID
		DestinationAddress: destinationAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			ContractAddress: fundedAddress,
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	log.Info(
		"Sending Teleporter transaction on source chain",
		"destinationChainID", subnetBInfo.BlockchainID,
	)

	receipt, _ := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		subnetAInfo,
		subnetBInfo,
		sendCrossChainMessageInput,
		fundedAddress,
		fundedKey,
		subnetATeleporterMessenger,
	)

	//
	// Relay the message to the destination
	//

	network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, false, false)

	//
	// Check that the message was not received on the Subnet B
	//
	delivered, err := subnetBTeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, subnetAInfo.BlockchainID, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeFalse())
}
