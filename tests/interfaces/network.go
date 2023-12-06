package interfaces

import (
	"context"
	"crypto/ecdsa"

	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ethereum/go-ethereum/common"
)

// Defines the interface for the network setup functions used in the E2E tests
type Network interface {
	// Returns all of the subnets support by this network.
	GetSubnetsInfo() []SubnetTestInfo

	// Returns the Teleporter contract address for all subnets in this network.
	GetTeleporterContractAddress() common.Address

	// An address and corresponding key that has native tokens on each of the subnets in this network.
	GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey)

	// Whether or not the Avalanche network being used is available outside the scope of the test application.
	IsExternalNetwork() bool

	// Whether or not the funded wallet is capable of relaying messages between subnets in this network.
	// Intended to be true for local networks where all nodes are querable by the test application for their
	// BLS signatures, and false for testnet networks where test application does not necessarily have
	// connections with each validator.
	SupportsIndependentRelaying() bool

	// For implementations where SupportsIndependentRelaying() is true, relays the specified message between the
	// two subnets,and returns the receipt of the transaction the message was delivered in.
	// For implementations where SupportsIndependentRelaying() is false, waits for the specific message to be relayed
	// by an external relayer, and returns the receipt of the transaction the message was delivered in.
	RelayMessage(
		ctx context.Context,
		sourceReceipt *types.Receipt,
		source SubnetTestInfo,
		destination SubnetTestInfo,
		expectSuccess bool) *types.Receipt
}
