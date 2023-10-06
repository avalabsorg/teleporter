package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/ava-labs/avalanche-network-runner/rpcpb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/plugin/evm"
	"github.com/ava-labs/subnet-evm/rpc"
	"github.com/ava-labs/subnet-evm/tests/utils/runner"
	teleporterutilities "github.com/ava-labs/teleporter/go/teleporter-utilities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

const (
	fundedKeyStr = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
)

var (
	teleporterContractAddress      common.Address
	subnetA, subnetB               ids.ID
	blockchainIDA, blockchainIDB   ids.ID
	chainANodeURIs, chainBNodeURIs []string
	fundedAddress                  = common.HexToAddress("0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC")
	fundedKey                      *ecdsa.PrivateKey
	chainAWSClient, chainBWSClient ethclient.Client
	chainAWSURI, chainBWSURI       string
	chainAIDInt, chainBIDInt       *big.Int

	// Internal vars only used to set up the local network
	anrConfig           = runner.NewDefaultANRConfig()
	manager             = runner.NewNetworkManager(anrConfig)
	warpChainConfigPath string
)

//
// Global test state getters. Should be called within a test spec, after SetupNetwork has been called
//

type SubnetTestInfo struct {
	SubnetID      ids.ID
	BlockchainID  ids.ID
	ChainNodeURIs []string
	ChainWSClient ethclient.Client
	ChainWSURI    string
	ChainIDInt    *big.Int
}

func GetSubnetsInfo() []SubnetTestInfo {
	return []SubnetTestInfo{
		GetSubnetATestInfo(),
		GetSubnetBTestInfo(),
	}
}

func GetSubnetATestInfo() SubnetTestInfo {
	return SubnetTestInfo{
		SubnetID:      subnetA,
		BlockchainID:  blockchainIDA,
		ChainNodeURIs: chainANodeURIs,
		ChainWSClient: chainAWSClient,
		ChainWSURI:    chainAWSURI,
		ChainIDInt:    big.NewInt(0).Set(chainAIDInt),
	}
}
func GetSubnetBTestInfo() SubnetTestInfo {
	return SubnetTestInfo{
		SubnetID:      subnetB,
		BlockchainID:  blockchainIDB,
		ChainNodeURIs: chainBNodeURIs,
		ChainWSClient: chainBWSClient,
		ChainWSURI:    chainBWSURI,
		ChainIDInt:    big.NewInt(0).Set(chainBIDInt),
	}
}
func GetTeleporterContractAddress() common.Address {
	return teleporterContractAddress
}
func GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	key, err := crypto.ToECDSA(crypto.FromECDSA(fundedKey))
	Expect(err).Should(BeNil())
	return fundedAddress, key
}

// SetupNetwork starts the default network and adds 10 new nodes as validators with BLS keys
// registered on the P-Chain.
// Adds two disjoint sets of 5 of the new validator nodes to validate two new subnets with a
// a single Subnet-EVM blockchain.
func SetupNetwork(warpGenesisFile string) {
	ctx := context.Background()
	var err error

	// Name 10 new validators (which should have BLS key registered)
	subnetANodeNames := []string{}
	subnetBNodeNames := []string{}
	for i := 1; i <= 10; i++ {
		n := fmt.Sprintf("node%d-bls", i)
		if i <= 5 {
			subnetANodeNames = append(subnetANodeNames, n)
		} else {
			subnetBNodeNames = append(subnetBNodeNames, n)
		}
	}
	f, err := os.CreateTemp(os.TempDir(), "config.json")
	Expect(err).Should(BeNil())
	_, err = f.Write([]byte(`{"warp-api-enabled": true}`))
	Expect(err).Should(BeNil())
	warpChainConfigPath = f.Name()

	// Make sure that the warp genesis file exists
	_, err = os.Stat(warpGenesisFile)
	Expect(err).Should(BeNil())

	// Construct the network using the avalanche-network-runner
	_, err = manager.StartDefaultNetwork(ctx)
	Expect(err).Should(BeNil())
	err = manager.SetupNetwork(
		ctx,
		anrConfig.AvalancheGoExecPath,
		[]*rpcpb.BlockchainSpec{
			{
				VmName:      evm.IDStr,
				Genesis:     warpGenesisFile,
				ChainConfig: warpChainConfigPath,
				SubnetSpec: &rpcpb.SubnetSpec{
					SubnetConfig: "",
					Participants: subnetANodeNames,
				},
			},
			{
				VmName:      evm.IDStr,
				Genesis:     warpGenesisFile,
				ChainConfig: warpChainConfigPath,
				SubnetSpec: &rpcpb.SubnetSpec{
					SubnetConfig: "",
					Participants: subnetBNodeNames,
				},
			},
		},
	)
	Expect(err).Should(BeNil())

	// Issue transactions to activate the proposerVM fork on the chains
	fundedKey, err = crypto.HexToECDSA(fundedKeyStr)
	Expect(err).Should(BeNil())
	SetupProposerVM(ctx, fundedKey, manager, 0)
	SetupProposerVM(ctx, fundedKey, manager, 1)

	// Set up subnet URIs
	subnetIDs := manager.GetSubnets()
	Expect(len(subnetIDs)).Should(Equal(2))

	subnetA = subnetIDs[0]
	subnetADetails, ok := manager.GetSubnet(subnetA)
	Expect(ok).Should(BeTrue())
	Expect(len(subnetADetails.ValidatorURIs)).Should(Equal(5))
	blockchainIDA = subnetADetails.BlockchainID
	chainANodeURIs = append(chainANodeURIs, subnetADetails.ValidatorURIs...)

	subnetB = subnetIDs[1]
	subnetBDetails, ok := manager.GetSubnet(subnetB)
	Expect(ok).Should(BeTrue())
	Expect(len(subnetBDetails.ValidatorURIs)).Should(Equal(5))
	blockchainIDB = subnetBDetails.BlockchainID
	chainBNodeURIs = append(chainBNodeURIs, subnetBDetails.ValidatorURIs...)

	log.Info(
		"Created URIs for subnets",
		"chainAURIs", chainANodeURIs,
		"chainBURIs", chainBNodeURIs,
		"blockchainIDA", blockchainIDA,
		"blockchainIDB", blockchainIDB,
	)

	chainAWSURI := HttpToWebsocketURI(chainANodeURIs[0], blockchainIDA.String())
	log.Info("Creating ethclient for blockchainA", "wsURI", chainAWSURI)
	chainAWSClient, err = ethclient.Dial(chainAWSURI)
	Expect(err).Should(BeNil())

	chainAIDInt, err = chainAWSClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	chainBWSURI := HttpToWebsocketURI(chainBNodeURIs[0], blockchainIDB.String())
	log.Info("Creating ethclient for blockchainB", "wsURI", chainBWSURI)
	chainBWSClient, err = ethclient.Dial(chainBWSURI)
	Expect(err).Should(BeNil())

	chainBIDInt, err = chainBWSClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	log.Info("Finished setting up e2e test subnet variables")
}

// DeployTeleporterContracts deploys the Teleporter contract to the two subnets. The caller is responsible for generating the
// deployment transaction information
func DeployTeleporterContracts(transactionBytes []byte, deployerAddress common.Address, contractAddress common.Address) {
	log.Info("Deploying Teleporter contract to subnets")

	subnetsInfo := GetSubnetsInfo()

	// Set the package level teleporterContractAddress
	teleporterContractAddress = contractAddress

	ctx := context.Background()

	for _, subnetInfo := range subnetsInfo {
		client := subnetInfo.ChainWSClient

		nonce, err := client.NonceAt(ctx, fundedAddress, nil)
		Expect(err).Should(BeNil())
		gasTipCap, err := client.SuggestGasTipCap(context.Background())
		Expect(err).Should(BeNil())
		baseFee, err := client.EstimateBaseFee(context.Background())
		Expect(err).Should(BeNil())
		gasFeeCap := baseFee.Mul(baseFee, big.NewInt(teleporterutilities.BaseFeeFactor))
		gasFeeCap.Add(gasFeeCap, big.NewInt(teleporterutilities.MaxPriorityFeePerGas))
		// Fund the deployer address
		{
			value := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10)) // 10eth
			txA := types.NewTx(&types.DynamicFeeTx{
				ChainID:   subnetInfo.ChainIDInt,
				Nonce:     nonce,
				To:        &deployerAddress,
				Gas:       DefaultTeleporterTransactionGas,
				GasFeeCap: gasFeeCap,
				GasTipCap: gasTipCap,
				Value:     value,
			})
			txSigner := types.LatestSignerForChainID(subnetInfo.ChainIDInt)
			triggerTx, err := types.SignTx(txA, txSigner, fundedKey)
			Expect(err).Should(BeNil())

			SendTransactionAndWaitForAcceptance(ctx, subnetInfo.ChainWSClient, triggerTx)
		}
		log.Info("Finished funding Teleporter deployer")

		// Deploy Teleporter
		{
			rpcClient, err := rpc.DialContext(ctx, HttpToRPCURI(subnetInfo.ChainNodeURIs[0], subnetInfo.BlockchainID.String()))
			Expect(err).Should(BeNil())
			defer rpcClient.Close()

			newHeads := make(chan *types.Header, 10)
			subA, err := subnetInfo.ChainWSClient.SubscribeNewHead(ctx, newHeads)
			Expect(err).Should(BeNil())
			defer subA.Unsubscribe()

			err = rpcClient.CallContext(ctx, nil, "eth_sendRawTransaction", hexutil.Encode(transactionBytes))
			Expect(err).Should(BeNil())

			<-newHeads
			teleporterCode, err := client.CodeAt(ctx, teleporterContractAddress, nil)
			Expect(err).Should(BeNil())
			Expect(len(teleporterCode)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode
		}
		log.Info("Finished deploying Teleporter contracts")
	}
}

func TearDownNetwork() {
	log.Info("Tearing down network")
	Expect(manager).ShouldNot(BeNil())
	Expect(manager.TeardownNetwork()).Should(BeNil())
	Expect(os.Remove(warpChainConfigPath)).Should(BeNil())
}
