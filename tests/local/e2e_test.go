// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package local

import (
	"os"
	"testing"

	"github.com/ava-labs/teleporter/tests/flows"
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	teleporterByteCodeFile = "./contracts/out/TeleporterMessenger.sol/TeleporterMessenger.json"
	warpGenesisFile        = "./tests/utils/warp-genesis.json"
)

var (
	localNetworkInstance *localNetwork
)

func TestE2E(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Teleporter e2e test")
}

// Define the Teleporter before and after suite functions.
var _ = ginkgo.BeforeSuite(func() {
	localNetworkInstance = newLocalNetwork(warpGenesisFile)
	// Generate the Teleporter deployment values
	teleporterDeployerTransaction, teleporterDeployerAddress, teleporterContractAddress, err :=
		deploymentUtils.ConstructKeylessTransaction(teleporterByteCodeFile, false)
	Expect(err).Should(BeNil())

	_, fundedKey := localNetworkInstance.GetFundedAccountInfo()
	localNetworkInstance.deployTeleporterContracts(
		teleporterDeployerTransaction,
		teleporterDeployerAddress,
		teleporterContractAddress,
		fundedKey,
	)
	localNetworkInstance.deployTeleporterRegistryContracts(teleporterContractAddress, fundedKey)
	log.Info("Set up ginkgo before suite")
})

var _ = ginkgo.AfterSuite(func() {
	localNetworkInstance.tearDownNetwork()
})

var _ = ginkgo.Describe("[Teleporter integration tests]", func() {
	// Cross-chain application tests
	ginkgo.It("Example cross chain messenger", func() {
		flows.ExampleMessenger(localNetworkInstance)
	})
	ginkgo.It("ERC20 bridge multihop", func() {
		flows.ERC20BridgeMultihop(localNetworkInstance)
	})

	// Teleporter tests
	ginkgo.It("Send a message from Subnet A to Subnet B, and one from B to A", func() {
		flows.BasicSendReceive(localNetworkInstance)
	})
	ginkgo.It("Deliver to the wrong chain", func() {
		flows.DeliverToWrongChain(localNetworkInstance)
	})
	ginkgo.It("Deliver to non-existent contract", func() {
		flows.DeliverToNonExistentContract(localNetworkInstance)
	})
	ginkgo.It("Retry successful execution", func() {
		flows.RetrySuccessfulExecution(localNetworkInstance)
	})
	ginkgo.It("Unallowed relayer", func() {
		flows.UnallowedRelayer(localNetworkInstance)
	})
	ginkgo.It("Relay message twice", func() {
		flows.RelayMessageTwice(localNetworkInstance)
	})
	ginkgo.It("Add additional fee amount", func() {
		flows.AddFeeAmount(localNetworkInstance)
	})
	ginkgo.It("Send specific receipts", func() {
		flows.SendSpecificReceipts(localNetworkInstance)
	})
	ginkgo.It("Insufficient gas", func() {
		flows.InsufficientGas(localNetworkInstance)
	})
	ginkgo.It("Resubmit altered message", func() {
		flows.ResubmitAlteredMessage(localNetworkInstance)
	})
	ginkgo.It("Block hash publish and receive", func() {
		flows.BlockHashPublishReceive(localNetworkInstance)
	})

	// The following tests require special behavior by the relayer, so we only run them on a local network
	ginkgo.It("Relayer modifies message", func() {
		flows.RelayerModifiesMessage(localNetworkInstance)
	})
	ginkgo.It("Validator churn", func() {
		flows.ValidatorChurn(localNetworkInstance)
	})
	// Since the validator churn test modifies the network topology, we put it last for now.
	// It should not affect the other tests, but we get some errors if we run it before the other tests.
	// TODO: we should fix this so that the order of the tests does not matter.
})
