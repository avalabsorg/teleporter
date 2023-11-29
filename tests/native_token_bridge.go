package tests

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	nativetokendestination "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/NativeTokenBridge/NativeTokenDestination"
	nativetokensource "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/NativeTokenBridge/NativeTokenSource"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func NativeTokenBridge(network network.Network) {
	const (

		// Each test file needs a unique deployer that must be funded with tokens to deploy
		deployerKeyStr                     = "aad7440febfc8f9d73a58c3cb1f1754779a566978f9ebffcd4f4698e9b043985"
		NativeTokenSourceByteCodeFile      = "./contracts/out/NativeTokenSource.sol/NativeTokenSource.json"
		NativeTokenDestinationByteCodeFile = "./contracts/out/NativeTokenDestination.sol/NativeTokenDestination.json"
	)
	var (
		initialReserveImbalance = big.NewInt(0).Mul(big.NewInt(1e15), big.NewInt(1e9))
		valueToSend             = big.NewInt(0).Div(initialReserveImbalance, big.NewInt(4))
		valueToReturn           = big.NewInt(0).Div(valueToSend, big.NewInt(4))
		ctx                     = context.Background()
		deployerAddress         = common.HexToAddress("0x1337cfd2dCff6270615B90938aCB1efE79801704")
		tokenReceiverAddress    = common.HexToAddress("0x0123456789012345678901234567890123456789")
		burnedTxFeeAddress      = common.HexToAddress("0x0100000000000000000000000000000000000000")

		emptyDestFeeInfo = nativetokendestination.TeleporterFeeInfo{
			FeeTokenAddress: common.Address{},
			Amount:          common.Big0,
		}
		emptySourceFeeInfo = nativetokensource.TeleporterFeeInfo{
			FeeTokenAddress: common.Address{},
			Amount:          common.Big0,
		}
	)

	subnets := network.GetSubnetsInfo()
	subnetA := subnets[0]
	subnetB := subnets[1]
	teleporterContractAddress := network.GetTeleporterContractAddress()

	// Info we need to calculate for the test
	deployerPK, err := crypto.HexToECDSA(deployerKeyStr)
	Expect(err).Should(BeNil())
	bridgeContractAddress, err := deploymentUtils.DeriveEVMContractAddress(deployerAddress, 0)
	Expect(err).Should(BeNil())
	log.Info("Native Token Bridge Contract Address: " + bridgeContractAddress.Hex())

	{ // Deploy the contracts
		// Both contracts in this test will be deployed to 0xAcB633F5B00099c7ec187eB00156c5cd9D854b5B,
		// though they do not necessarily have to be deployed at the same address, each contract needs
		// to know the address of the other.
		// The nativeTokenDestination contract must be added to "adminAddresses" of "contractNativeMinterConfig"
		// in the genesis file for the subnet. This will allow it to call the native minter precompile.
		erc20TokenSourceAbi, err := nativetokensource.NativeTokenSourceMetaData.GetAbi()
		Expect(err).Should(BeNil())
		utils.DeployContract(
			ctx,
			NativeTokenSourceByteCodeFile,
			deployerPK,
			subnetA,
			erc20TokenSourceAbi,
			teleporterContractAddress,
			subnetB.BlockchainID,
			bridgeContractAddress,
		)
		Expect(err).Should(BeNil())

		nativeTokenDestinationAbi, err := nativetokendestination.NativeTokenDestinationMetaData.GetAbi()
		Expect(err).Should(BeNil())
		utils.DeployContract(
			ctx,
			NativeTokenDestinationByteCodeFile,
			deployerPK,
			subnetB,
			nativeTokenDestinationAbi,
			teleporterContractAddress,
			subnetA.BlockchainID,
			bridgeContractAddress,
			initialReserveImbalance,
		)

		log.Info("Finished deploying Bridge contracts")
	}

	// Create abi objects to call the contract with
	nativeTokenDestination, err := nativetokendestination.NewNativeTokenDestination(
		bridgeContractAddress,
		subnetB.ChainWSClient,
	)
	Expect(err).Should(BeNil())
	nativeTokenSource, err := nativetokensource.NewNativeTokenSource(
		bridgeContractAddress,
		subnetA.ChainWSClient,
	)
	Expect(err).Should(BeNil())

	// Helper function
	sendTokensToSource := func(valueToSend *big.Int, fromKey *ecdsa.PrivateKey, toAddress common.Address) *types.Receipt {
		transactor, err := bind.NewKeyedTransactorWithChainID(fromKey, subnetB.ChainIDInt)
		Expect(err).Should(BeNil())
		transactor.Value = valueToSend

		tx, err := nativeTokenDestination.TransferToSource(
			transactor,
			toAddress,
			emptyDestFeeInfo,
			[]common.Address{},
		)
		Expect(err).Should(BeNil())
		log.Info(
			"Sent TransferToSource transaction on destination chain",
			"sourceBlockchainID",
			subnetA.BlockchainID,
			"txHash",
			tx.Hash().Hex(),
		)

		destChainReceipt := utils.WaitForTransactionSuccess(ctx, tx.Hash(), subnetB)

		transferEvent, err := utils.GetEventFromLogs(
			destChainReceipt.Logs,
			nativeTokenDestination.ParseTransferToSource,
		)
		Expect(err).Should(BeNil())
		utils.ExpectBigEqual(transferEvent.Amount, valueToSend)

		receipt := network.RelayMessage(ctx, destChainReceipt, subnetB, subnetA, true)

		return receipt
	}

	// Helper function
	sendTokensToDestination := func(valueToSend *big.Int, fromKey *ecdsa.PrivateKey, toAddress common.Address) *types.Receipt {
		transactor, err := bind.NewKeyedTransactorWithChainID(fromKey, subnetA.ChainIDInt)
		Expect(err).Should(BeNil())
		transactor.Value = valueToSend

		tx, err := nativeTokenSource.TransferToDestination(
			transactor,
			toAddress,
			emptySourceFeeInfo,
			[]common.Address{},
		)
		Expect(err).Should(BeNil())
		log.Info(
			"Sent TransferToDestination transaction on source chain",
			"destinationBlockchainID",
			subnetB.BlockchainID,
			"txHash",
			tx.Hash().Hex(),
		)

		sourceChainReceipt := utils.WaitForTransactionSuccess(ctx, tx.Hash(), subnetA)

		transferEvent, err := utils.GetEventFromLogs(
			sourceChainReceipt.Logs,
			nativeTokenSource.ParseTransferToDestination,
		)
		Expect(err).Should(BeNil())
		utils.ExpectBigEqual(transferEvent.Amount, valueToSend)

		receipt := network.RelayMessage(ctx, sourceChainReceipt, subnetA, subnetB, true)

		return receipt
	}

	{ // Transfer some tokens A -> B
		// Check starting balance is 0
		utils.CheckBalance(ctx, tokenReceiverAddress, common.Big0, subnetB.ChainWSClient)

		checkReserveImbalance(initialReserveImbalance, nativeTokenDestination)

		destChainReceipt := sendTokensToDestination(valueToSend, deployerPK, tokenReceiverAddress)

		checkCollateralEvent(
			destChainReceipt.Logs,
			nativeTokenDestination,
			valueToSend,
			big.NewInt(0).Sub(initialReserveImbalance, valueToSend),
		)
		checkReserveImbalance(
			big.NewInt(0).Sub(initialReserveImbalance, valueToSend),
			nativeTokenDestination,
		)

		_, err = utils.GetEventFromLogs(
			destChainReceipt.Logs,
			nativeTokenDestination.ParseNativeTokensMinted,
		)
		Expect(err).ShouldNot(BeNil())

		// Check intermediate balance, no tokens should be minted because we haven't collateralized
		utils.CheckBalance(ctx, tokenReceiverAddress, common.Big0, subnetB.ChainWSClient)
	}

	{ // Fail to Transfer tokens B -> A because bridge is not collateralized
		// Check starting balance is 0
		utils.CheckBalance(ctx, tokenReceiverAddress, common.Big0, subnetA.ChainWSClient)

		transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetB.ChainIDInt)
		Expect(err).Should(BeNil())
		transactor.Value = valueToSend

		// This transfer should revert because the bridge isn't collateralized
		_, err = nativeTokenDestination.TransferToSource(
			transactor,
			tokenReceiverAddress,
			emptyDestFeeInfo,
			[]common.Address{},
		)
		Expect(err).ShouldNot(BeNil())

		// Check we should fail to send because we're not collateralized
		utils.CheckBalance(ctx, tokenReceiverAddress, common.Big0, subnetA.ChainWSClient)
	}

	{ // Transfer more tokens A -> B to collateralize the bridge
		// Check starting balance is 0
		utils.CheckBalance(ctx, tokenReceiverAddress, common.Big0, subnetB.ChainWSClient)
		checkReserveImbalance(
			big.NewInt(0).Sub(initialReserveImbalance, valueToSend),
			nativeTokenDestination,
		)

		destChainReceipt := sendTokensToDestination(
			initialReserveImbalance,
			deployerPK,
			tokenReceiverAddress,
		)

		checkCollateralEvent(
			destChainReceipt.Logs,
			nativeTokenDestination,
			big.NewInt(0).Sub(initialReserveImbalance, valueToSend),
			common.Big0,
		)
		checkMintEvent(
			destChainReceipt.Logs,
			nativeTokenDestination,
			tokenReceiverAddress,
			valueToSend,
		)
		checkReserveImbalance(common.Big0, nativeTokenDestination)

		// We should have minted the excess coins after checking the collateral
		utils.CheckBalance(ctx, tokenReceiverAddress, valueToSend, subnetB.ChainWSClient)
	}

	{ // Transfer tokens B -> A
		sourceChainReceipt := sendTokensToSource(valueToReturn, deployerPK, tokenReceiverAddress)

		checkUnlockNativeEvent(
			sourceChainReceipt.Logs,
			nativeTokenSource,
			tokenReceiverAddress,
			valueToReturn,
		)

		utils.CheckBalance(ctx, tokenReceiverAddress, valueToReturn, subnetA.ChainWSClient)
	}

	{ // Check reporting of burned tx fees to Source Chain
		burnedTxFeesBalanceDest, err := subnetB.ChainWSClient.BalanceAt(
			ctx,
			burnedTxFeeAddress,
			nil,
		)
		Expect(err).Should(BeNil())
		Expect(burnedTxFeesBalanceDest.Cmp(common.Big0) > 0).Should(BeTrue())

		transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetB.ChainIDInt)
		Expect(err).Should(BeNil())
		tx, err := nativeTokenDestination.ReportTotalBurnedTxFees(
			transactor,
			emptyDestFeeInfo,
			[]common.Address{},
		)
		Expect(err).Should(BeNil())

		destChainReceipt := utils.WaitForTransactionSuccess(ctx, tx.Hash(), subnetB)

		reportEvent, err := utils.GetEventFromLogs(
			destChainReceipt.Logs,
			nativeTokenDestination.ParseReportTotalBurnedTxFees,
		)
		Expect(err).Should(BeNil())
		utils.ExpectBigEqual(reportEvent.BurnAddressBalance, burnedTxFeesBalanceDest)

		burnedTxFeesBalanceSource, err := subnetA.ChainWSClient.BalanceAt(
			ctx,
			burnedTxFeeAddress,
			nil,
		)
		Expect(err).Should(BeNil())
		Expect(burnedTxFeesBalanceSource.Cmp(common.Big0) > 0).Should(BeTrue())

		sourceChainReceipt := network.RelayMessage(ctx, destChainReceipt, subnetB, subnetA, true)

		burnEvent, err := utils.GetEventFromLogs(
			sourceChainReceipt.Logs,
			nativeTokenSource.ParseBurnTokens,
		)
		Expect(err).Should(BeNil())
		utils.ExpectBigEqual(burnedTxFeesBalanceDest, burnEvent.Amount)

		burnedTxFeesBalanceSource2, err := subnetA.ChainWSClient.BalanceAt(
			ctx,
			burnedTxFeeAddress,
			nil,
		)
		Expect(err).Should(BeNil())
		Expect(
			burnedTxFeesBalanceSource2.Cmp(
				big.NewInt(0).Add(burnedTxFeesBalanceSource, burnEvent.Amount),
			) >= 0,
		).Should(BeTrue())
	}
}

func checkUnlockNativeEvent(
	logs []*types.Log,
	nativeTokenSource *nativetokensource.NativeTokenSource,
	recipient common.Address,
	value *big.Int,
) {
	unlockEvent, err := utils.GetEventFromLogs(logs, nativeTokenSource.ParseUnlockTokens)
	Expect(err).Should(BeNil())
	Expect(unlockEvent.Recipient).Should(Equal(recipient))
	Expect(unlockEvent.Amount.Cmp(value)).Should(BeZero())
}

func checkCollateralEvent(
	logs []*types.Log,
	nativeTokenDestination *nativetokendestination.NativeTokenDestination,
	collateralAdded *big.Int,
	collateralRemaining *big.Int,
) {
	collateralEvent, err := utils.GetEventFromLogs(
		logs,
		nativeTokenDestination.ParseCollateralAdded,
	)
	Expect(err).Should(BeNil())
	Expect(collateralEvent.Amount.Cmp(collateralAdded)).Should(BeZero())
	Expect(collateralEvent.Remaining.Cmp(collateralEvent.Remaining)).Should(BeZero())
}

func checkMintEvent(
	logs []*types.Log,
	nativeTokenDestination *nativetokendestination.NativeTokenDestination,
	recipient common.Address,
	value *big.Int,
) {
	mintEvent, err := utils.GetEventFromLogs(logs, nativeTokenDestination.ParseNativeTokensMinted)
	Expect(err).Should(BeNil())
	Expect(mintEvent.Recipient).Should(Equal(recipient))
	Expect(mintEvent.Amount.Cmp(value)).Should(BeZero())
}

func checkReserveImbalance(
	value *big.Int,
	nativeTokenDestination *nativetokendestination.NativeTokenDestination,
) {
	imbalance, err := nativeTokenDestination.CurrentReserveImbalance(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(imbalance, value)
}
