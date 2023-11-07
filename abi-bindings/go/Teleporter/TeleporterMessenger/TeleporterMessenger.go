// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teleportermessenger

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TeleporterFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type TeleporterFeeInfo struct {
	ContractAddress common.Address
	Amount          *big.Int
}

// TeleporterMessage is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessage struct {
	MessageID               *big.Int
	SenderAddress           common.Address
	DestinationChainID      [32]byte
	DestinationAddress      common.Address
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Receipts                []TeleporterMessageReceipt
	Message                 []byte
}

// TeleporterMessageInput is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageInput struct {
	DestinationChainID      [32]byte
	DestinationAddress      common.Address
	FeeInfo                 TeleporterFeeInfo
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Message                 []byte
}

// TeleporterMessageReceipt is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageReceipt struct {
	ReceivedMessageID    *big.Int
	RelayerRewardAddress common.Address
}

// TeleporterMessengerMetaData contains all meta data concerning the TeleporterMessenger contract.
var TeleporterMessengerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"updatedFeeInfo\",\"type\":\"tuple\"}],\"name\":\"AddFeeAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"MessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MessageExecutionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deliverer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardRedeemer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ReceiveCrossChainMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RelayerRewardsRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"SendCrossChainMessage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalFeeAmount\",\"type\":\"uint256\"}],\"name\":\"addFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"checkRelayerRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"getFeeInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chainID\",\"type\":\"bytes32\"}],\"name\":\"getNextMessageID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getReceiptAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"chainID\",\"type\":\"bytes32\"}],\"name\":\"getReceiptQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"getRelayerRewardAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"}],\"name\":\"latestMessageIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"messageReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"}],\"name\":\"receiptQueues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"first\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"name\":\"receiveCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"receivedFailedMessageHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"redeemRelayerRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retryMessageExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retrySendCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageInput\",\"name\":\"messageInput\",\"type\":\"tuple\"}],\"name\":\"sendCrossChainMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"messageIDs\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"sendSpecifiedReceipts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"name\":\"sentMessageInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TeleporterMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleporterMessengerMetaData.ABI instead.
var TeleporterMessengerABI = TeleporterMessengerMetaData.ABI

// TeleporterMessenger is an auto generated Go binding around an Ethereum contract.
type TeleporterMessenger struct {
	TeleporterMessengerCaller     // Read-only binding to the contract
	TeleporterMessengerTransactor // Write-only binding to the contract
	TeleporterMessengerFilterer   // Log filterer for contract events
}

// TeleporterMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeleporterMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeleporterMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeleporterMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeleporterMessengerSession struct {
	Contract     *TeleporterMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TeleporterMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeleporterMessengerCallerSession struct {
	Contract *TeleporterMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TeleporterMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeleporterMessengerTransactorSession struct {
	Contract     *TeleporterMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TeleporterMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeleporterMessengerRaw struct {
	Contract *TeleporterMessenger // Generic contract binding to access the raw methods on
}

// TeleporterMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeleporterMessengerCallerRaw struct {
	Contract *TeleporterMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// TeleporterMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeleporterMessengerTransactorRaw struct {
	Contract *TeleporterMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeleporterMessenger creates a new instance of TeleporterMessenger, bound to a specific deployed contract.
func NewTeleporterMessenger(address common.Address, backend bind.ContractBackend) (*TeleporterMessenger, error) {
	contract, err := bindTeleporterMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessenger{TeleporterMessengerCaller: TeleporterMessengerCaller{contract: contract}, TeleporterMessengerTransactor: TeleporterMessengerTransactor{contract: contract}, TeleporterMessengerFilterer: TeleporterMessengerFilterer{contract: contract}}, nil
}

// NewTeleporterMessengerCaller creates a new read-only instance of TeleporterMessenger, bound to a specific deployed contract.
func NewTeleporterMessengerCaller(address common.Address, caller bind.ContractCaller) (*TeleporterMessengerCaller, error) {
	contract, err := bindTeleporterMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerCaller{contract: contract}, nil
}

// NewTeleporterMessengerTransactor creates a new write-only instance of TeleporterMessenger, bound to a specific deployed contract.
func NewTeleporterMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*TeleporterMessengerTransactor, error) {
	contract, err := bindTeleporterMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerTransactor{contract: contract}, nil
}

// NewTeleporterMessengerFilterer creates a new log filterer instance of TeleporterMessenger, bound to a specific deployed contract.
func NewTeleporterMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*TeleporterMessengerFilterer, error) {
	contract, err := bindTeleporterMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerFilterer{contract: contract}, nil
}

// bindTeleporterMessenger binds a generic wrapper to an already deployed contract.
func bindTeleporterMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterMessenger *TeleporterMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterMessenger.Contract.TeleporterMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterMessenger *TeleporterMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.TeleporterMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterMessenger *TeleporterMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.TeleporterMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterMessenger *TeleporterMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterMessenger *TeleporterMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterMessenger *TeleporterMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.contract.Transact(opts, method, params...)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterMessenger *TeleporterMessengerSession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterMessenger.Contract.WARPMESSENGER(&_TeleporterMessenger.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterMessenger.Contract.WARPMESSENGER(&_TeleporterMessenger.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) BlockchainID() ([32]byte, error) {
	return _TeleporterMessenger.Contract.BlockchainID(&_TeleporterMessenger.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) BlockchainID() ([32]byte, error) {
	return _TeleporterMessenger.Contract.BlockchainID(&_TeleporterMessenger.CallOpts)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCaller) CheckRelayerRewardAmount(opts *bind.CallOpts, relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "checkRelayerRewardAmount", relayer, feeAsset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) CheckRelayerRewardAmount(relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	return _TeleporterMessenger.Contract.CheckRelayerRewardAmount(&_TeleporterMessenger.CallOpts, relayer, feeAsset)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) CheckRelayerRewardAmount(relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	return _TeleporterMessenger.Contract.CheckRelayerRewardAmount(&_TeleporterMessenger.CallOpts, relayer, feeAsset)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0x82f2c43a.
//
// Solidity: function getFeeInfo(bytes32 destinationChainID, uint256 messageID) view returns(address, uint256)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetFeeInfo(opts *bind.CallOpts, destinationChainID [32]byte, messageID *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getFeeInfo", destinationChainID, messageID)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetFeeInfo is a free data retrieval call binding the contract method 0x82f2c43a.
//
// Solidity: function getFeeInfo(bytes32 destinationChainID, uint256 messageID) view returns(address, uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) GetFeeInfo(destinationChainID [32]byte, messageID *big.Int) (common.Address, *big.Int, error) {
	return _TeleporterMessenger.Contract.GetFeeInfo(&_TeleporterMessenger.CallOpts, destinationChainID, messageID)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0x82f2c43a.
//
// Solidity: function getFeeInfo(bytes32 destinationChainID, uint256 messageID) view returns(address, uint256)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetFeeInfo(destinationChainID [32]byte, messageID *big.Int) (common.Address, *big.Int, error) {
	return _TeleporterMessenger.Contract.GetFeeInfo(&_TeleporterMessenger.CallOpts, destinationChainID, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x220c9568.
//
// Solidity: function getMessageHash(bytes32 destinationChainID, uint256 messageID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetMessageHash(opts *bind.CallOpts, destinationChainID [32]byte, messageID *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getMessageHash", destinationChainID, messageID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x220c9568.
//
// Solidity: function getMessageHash(bytes32 destinationChainID, uint256 messageID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) GetMessageHash(destinationChainID [32]byte, messageID *big.Int) ([32]byte, error) {
	return _TeleporterMessenger.Contract.GetMessageHash(&_TeleporterMessenger.CallOpts, destinationChainID, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x220c9568.
//
// Solidity: function getMessageHash(bytes32 destinationChainID, uint256 messageID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetMessageHash(destinationChainID [32]byte, messageID *big.Int) ([32]byte, error) {
	return _TeleporterMessenger.Contract.GetMessageHash(&_TeleporterMessenger.CallOpts, destinationChainID, messageID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 chainID) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetNextMessageID(opts *bind.CallOpts, chainID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getNextMessageID", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 chainID) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) GetNextMessageID(chainID [32]byte) (*big.Int, error) {
	return _TeleporterMessenger.Contract.GetNextMessageID(&_TeleporterMessenger.CallOpts, chainID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 chainID) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetNextMessageID(chainID [32]byte) (*big.Int, error) {
	return _TeleporterMessenger.Contract.GetNextMessageID(&_TeleporterMessenger.CallOpts, chainID)
}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 chainID, uint256 index) view returns((uint256,address))
func (_TeleporterMessenger *TeleporterMessengerCaller) GetReceiptAtIndex(opts *bind.CallOpts, chainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getReceiptAtIndex", chainID, index)

	if err != nil {
		return *new(TeleporterMessageReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(TeleporterMessageReceipt)).(*TeleporterMessageReceipt)

	return out0, err

}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 chainID, uint256 index) view returns((uint256,address))
func (_TeleporterMessenger *TeleporterMessengerSession) GetReceiptAtIndex(chainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	return _TeleporterMessenger.Contract.GetReceiptAtIndex(&_TeleporterMessenger.CallOpts, chainID, index)
}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 chainID, uint256 index) view returns((uint256,address))
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetReceiptAtIndex(chainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	return _TeleporterMessenger.Contract.GetReceiptAtIndex(&_TeleporterMessenger.CallOpts, chainID, index)
}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 chainID) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetReceiptQueueSize(opts *bind.CallOpts, chainID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getReceiptQueueSize", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 chainID) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) GetReceiptQueueSize(chainID [32]byte) (*big.Int, error) {
	return _TeleporterMessenger.Contract.GetReceiptQueueSize(&_TeleporterMessenger.CallOpts, chainID)
}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 chainID) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetReceiptQueueSize(chainID [32]byte) (*big.Int, error) {
	return _TeleporterMessenger.Contract.GetReceiptQueueSize(&_TeleporterMessenger.CallOpts, chainID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x33e890fe.
//
// Solidity: function getRelayerRewardAddress(bytes32 originChainID, uint256 messageID) view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetRelayerRewardAddress(opts *bind.CallOpts, originChainID [32]byte, messageID *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getRelayerRewardAddress", originChainID, messageID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x33e890fe.
//
// Solidity: function getRelayerRewardAddress(bytes32 originChainID, uint256 messageID) view returns(address)
func (_TeleporterMessenger *TeleporterMessengerSession) GetRelayerRewardAddress(originChainID [32]byte, messageID *big.Int) (common.Address, error) {
	return _TeleporterMessenger.Contract.GetRelayerRewardAddress(&_TeleporterMessenger.CallOpts, originChainID, messageID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x33e890fe.
//
// Solidity: function getRelayerRewardAddress(bytes32 originChainID, uint256 messageID) view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetRelayerRewardAddress(originChainID [32]byte, messageID *big.Int) (common.Address, error) {
	return _TeleporterMessenger.Contract.GetRelayerRewardAddress(&_TeleporterMessenger.CallOpts, originChainID, messageID)
}

// LatestMessageIDs is a free data retrieval call binding the contract method 0x29ec9beb.
//
// Solidity: function latestMessageIDs(bytes32 destinationChainID) view returns(uint256 messageID)
func (_TeleporterMessenger *TeleporterMessengerCaller) LatestMessageIDs(opts *bind.CallOpts, destinationChainID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "latestMessageIDs", destinationChainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestMessageIDs is a free data retrieval call binding the contract method 0x29ec9beb.
//
// Solidity: function latestMessageIDs(bytes32 destinationChainID) view returns(uint256 messageID)
func (_TeleporterMessenger *TeleporterMessengerSession) LatestMessageIDs(destinationChainID [32]byte) (*big.Int, error) {
	return _TeleporterMessenger.Contract.LatestMessageIDs(&_TeleporterMessenger.CallOpts, destinationChainID)
}

// LatestMessageIDs is a free data retrieval call binding the contract method 0x29ec9beb.
//
// Solidity: function latestMessageIDs(bytes32 destinationChainID) view returns(uint256 messageID)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) LatestMessageIDs(destinationChainID [32]byte) (*big.Int, error) {
	return _TeleporterMessenger.Contract.LatestMessageIDs(&_TeleporterMessenger.CallOpts, destinationChainID)
}

// MessageReceived is a free data retrieval call binding the contract method 0xe03555df.
//
// Solidity: function messageReceived(bytes32 originChainID, uint256 messageID) view returns(bool)
func (_TeleporterMessenger *TeleporterMessengerCaller) MessageReceived(opts *bind.CallOpts, originChainID [32]byte, messageID *big.Int) (bool, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "messageReceived", originChainID, messageID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MessageReceived is a free data retrieval call binding the contract method 0xe03555df.
//
// Solidity: function messageReceived(bytes32 originChainID, uint256 messageID) view returns(bool)
func (_TeleporterMessenger *TeleporterMessengerSession) MessageReceived(originChainID [32]byte, messageID *big.Int) (bool, error) {
	return _TeleporterMessenger.Contract.MessageReceived(&_TeleporterMessenger.CallOpts, originChainID, messageID)
}

// MessageReceived is a free data retrieval call binding the contract method 0xe03555df.
//
// Solidity: function messageReceived(bytes32 originChainID, uint256 messageID) view returns(bool)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) MessageReceived(originChainID [32]byte, messageID *big.Int) (bool, error) {
	return _TeleporterMessenger.Contract.MessageReceived(&_TeleporterMessenger.CallOpts, originChainID, messageID)
}

// ReceiptQueues is a free data retrieval call binding the contract method 0xe6e67bd5.
//
// Solidity: function receiptQueues(bytes32 sourceChainID) view returns(uint256 first, uint256 last)
func (_TeleporterMessenger *TeleporterMessengerCaller) ReceiptQueues(opts *bind.CallOpts, sourceChainID [32]byte) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "receiptQueues", sourceChainID)

	outstruct := new(struct {
		First *big.Int
		Last  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.First = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Last = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ReceiptQueues is a free data retrieval call binding the contract method 0xe6e67bd5.
//
// Solidity: function receiptQueues(bytes32 sourceChainID) view returns(uint256 first, uint256 last)
func (_TeleporterMessenger *TeleporterMessengerSession) ReceiptQueues(sourceChainID [32]byte) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	return _TeleporterMessenger.Contract.ReceiptQueues(&_TeleporterMessenger.CallOpts, sourceChainID)
}

// ReceiptQueues is a free data retrieval call binding the contract method 0xe6e67bd5.
//
// Solidity: function receiptQueues(bytes32 sourceChainID) view returns(uint256 first, uint256 last)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) ReceiptQueues(sourceChainID [32]byte) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	return _TeleporterMessenger.Contract.ReceiptQueues(&_TeleporterMessenger.CallOpts, sourceChainID)
}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0xc9bb1143.
//
// Solidity: function receivedFailedMessageHashes(bytes32 sourceChainID, uint256 messageID) view returns(bytes32 messageHash)
func (_TeleporterMessenger *TeleporterMessengerCaller) ReceivedFailedMessageHashes(opts *bind.CallOpts, sourceChainID [32]byte, messageID *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "receivedFailedMessageHashes", sourceChainID, messageID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0xc9bb1143.
//
// Solidity: function receivedFailedMessageHashes(bytes32 sourceChainID, uint256 messageID) view returns(bytes32 messageHash)
func (_TeleporterMessenger *TeleporterMessengerSession) ReceivedFailedMessageHashes(sourceChainID [32]byte, messageID *big.Int) ([32]byte, error) {
	return _TeleporterMessenger.Contract.ReceivedFailedMessageHashes(&_TeleporterMessenger.CallOpts, sourceChainID, messageID)
}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0xc9bb1143.
//
// Solidity: function receivedFailedMessageHashes(bytes32 sourceChainID, uint256 messageID) view returns(bytes32 messageHash)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) ReceivedFailedMessageHashes(sourceChainID [32]byte, messageID *big.Int) ([32]byte, error) {
	return _TeleporterMessenger.Contract.ReceivedFailedMessageHashes(&_TeleporterMessenger.CallOpts, sourceChainID, messageID)
}

// SentMessageInfo is a free data retrieval call binding the contract method 0x66533d12.
//
// Solidity: function sentMessageInfo(bytes32 destinationChainID, uint256 messageID) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerCaller) SentMessageInfo(opts *bind.CallOpts, destinationChainID [32]byte, messageID *big.Int) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "sentMessageInfo", destinationChainID, messageID)

	outstruct := new(struct {
		MessageHash [32]byte
		FeeInfo     TeleporterFeeInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MessageHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.FeeInfo = *abi.ConvertType(out[1], new(TeleporterFeeInfo)).(*TeleporterFeeInfo)

	return *outstruct, err

}

// SentMessageInfo is a free data retrieval call binding the contract method 0x66533d12.
//
// Solidity: function sentMessageInfo(bytes32 destinationChainID, uint256 messageID) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerSession) SentMessageInfo(destinationChainID [32]byte, messageID *big.Int) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	return _TeleporterMessenger.Contract.SentMessageInfo(&_TeleporterMessenger.CallOpts, destinationChainID, messageID)
}

// SentMessageInfo is a free data retrieval call binding the contract method 0x66533d12.
//
// Solidity: function sentMessageInfo(bytes32 destinationChainID, uint256 messageID) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) SentMessageInfo(destinationChainID [32]byte, messageID *big.Int) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	return _TeleporterMessenger.Contract.SentMessageInfo(&_TeleporterMessenger.CallOpts, destinationChainID, messageID)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x19570c74.
//
// Solidity: function addFeeAmount(bytes32 destinationChainID, uint256 messageID, address feeContractAddress, uint256 additionalFeeAmount) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) AddFeeAmount(opts *bind.TransactOpts, destinationChainID [32]byte, messageID *big.Int, feeContractAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "addFeeAmount", destinationChainID, messageID, feeContractAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x19570c74.
//
// Solidity: function addFeeAmount(bytes32 destinationChainID, uint256 messageID, address feeContractAddress, uint256 additionalFeeAmount) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) AddFeeAmount(destinationChainID [32]byte, messageID *big.Int, feeContractAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.AddFeeAmount(&_TeleporterMessenger.TransactOpts, destinationChainID, messageID, feeContractAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x19570c74.
//
// Solidity: function addFeeAmount(bytes32 destinationChainID, uint256 messageID, address feeContractAddress, uint256 additionalFeeAmount) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) AddFeeAmount(destinationChainID [32]byte, messageID *big.Int, feeContractAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.AddFeeAmount(&_TeleporterMessenger.TransactOpts, destinationChainID, messageID, feeContractAddress, additionalFeeAmount)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) ReceiveCrossChainMessage(opts *bind.TransactOpts, messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "receiveCrossChainMessage", messageIndex, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) ReceiveCrossChainMessage(messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.ReceiveCrossChainMessage(&_TeleporterMessenger.TransactOpts, messageIndex, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) ReceiveCrossChainMessage(messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.ReceiveCrossChainMessage(&_TeleporterMessenger.TransactOpts, messageIndex, relayerRewardAddress)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) RedeemRelayerRewards(opts *bind.TransactOpts, feeAsset common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "redeemRelayerRewards", feeAsset)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) RedeemRelayerRewards(feeAsset common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RedeemRelayerRewards(&_TeleporterMessenger.TransactOpts, feeAsset)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) RedeemRelayerRewards(feeAsset common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RedeemRelayerRewards(&_TeleporterMessenger.TransactOpts, feeAsset)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xfc2d6197.
//
// Solidity: function retryMessageExecution(bytes32 originChainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) RetryMessageExecution(opts *bind.TransactOpts, originChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "retryMessageExecution", originChainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xfc2d6197.
//
// Solidity: function retryMessageExecution(bytes32 originChainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) RetryMessageExecution(originChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RetryMessageExecution(&_TeleporterMessenger.TransactOpts, originChainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xfc2d6197.
//
// Solidity: function retryMessageExecution(bytes32 originChainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) RetryMessageExecution(originChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RetryMessageExecution(&_TeleporterMessenger.TransactOpts, originChainID, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0xaf402850.
//
// Solidity: function retrySendCrossChainMessage(bytes32 destinationChainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) RetrySendCrossChainMessage(opts *bind.TransactOpts, destinationChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "retrySendCrossChainMessage", destinationChainID, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0xaf402850.
//
// Solidity: function retrySendCrossChainMessage(bytes32 destinationChainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) RetrySendCrossChainMessage(destinationChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RetrySendCrossChainMessage(&_TeleporterMessenger.TransactOpts, destinationChainID, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0xaf402850.
//
// Solidity: function retrySendCrossChainMessage(bytes32 destinationChainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) RetrySendCrossChainMessage(destinationChainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RetrySendCrossChainMessage(&_TeleporterMessenger.TransactOpts, destinationChainID, message)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerTransactor) SendCrossChainMessage(opts *bind.TransactOpts, messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "sendCrossChainMessage", messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.SendCrossChainMessage(&_TeleporterMessenger.TransactOpts, messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.SendCrossChainMessage(&_TeleporterMessenger.TransactOpts, messageInput)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0x191eb698.
//
// Solidity: function sendSpecifiedReceipts(bytes32 originChainID, uint256[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerTransactor) SendSpecifiedReceipts(opts *bind.TransactOpts, originChainID [32]byte, messageIDs []*big.Int, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "sendSpecifiedReceipts", originChainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0x191eb698.
//
// Solidity: function sendSpecifiedReceipts(bytes32 originChainID, uint256[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) SendSpecifiedReceipts(originChainID [32]byte, messageIDs []*big.Int, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.SendSpecifiedReceipts(&_TeleporterMessenger.TransactOpts, originChainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0x191eb698.
//
// Solidity: function sendSpecifiedReceipts(bytes32 originChainID, uint256[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) SendSpecifiedReceipts(originChainID [32]byte, messageIDs []*big.Int, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.SendSpecifiedReceipts(&_TeleporterMessenger.TransactOpts, originChainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// TeleporterMessengerAddFeeAmountIterator is returned from FilterAddFeeAmount and is used to iterate over the raw logs and unpacked data for AddFeeAmount events raised by the TeleporterMessenger contract.
type TeleporterMessengerAddFeeAmountIterator struct {
	Event *TeleporterMessengerAddFeeAmount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterMessengerAddFeeAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerAddFeeAmount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeleporterMessengerAddFeeAmount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeleporterMessengerAddFeeAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerAddFeeAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerAddFeeAmount represents a AddFeeAmount event raised by the TeleporterMessenger contract.
type TeleporterMessengerAddFeeAmount struct {
	DestinationChainID [32]byte
	MessageID          *big.Int
	UpdatedFeeInfo     TeleporterFeeInfo
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAddFeeAmount is a free log retrieval operation binding the contract event 0x28fe05eedf0479c9159e5b6dd2a28c93fa1a408eba22dc801fd9bc493a7fc0c2.
//
// Solidity: event AddFeeAmount(bytes32 indexed destinationChainID, uint256 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterAddFeeAmount(opts *bind.FilterOpts, destinationChainID [][32]byte, messageID []*big.Int) (*TeleporterMessengerAddFeeAmountIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "AddFeeAmount", destinationChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerAddFeeAmountIterator{contract: _TeleporterMessenger.contract, event: "AddFeeAmount", logs: logs, sub: sub}, nil
}

// WatchAddFeeAmount is a free log subscription operation binding the contract event 0x28fe05eedf0479c9159e5b6dd2a28c93fa1a408eba22dc801fd9bc493a7fc0c2.
//
// Solidity: event AddFeeAmount(bytes32 indexed destinationChainID, uint256 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchAddFeeAmount(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerAddFeeAmount, destinationChainID [][32]byte, messageID []*big.Int) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "AddFeeAmount", destinationChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerAddFeeAmount)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddFeeAmount is a log parse operation binding the contract event 0x28fe05eedf0479c9159e5b6dd2a28c93fa1a408eba22dc801fd9bc493a7fc0c2.
//
// Solidity: event AddFeeAmount(bytes32 indexed destinationChainID, uint256 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseAddFeeAmount(log types.Log) (*TeleporterMessengerAddFeeAmount, error) {
	event := new(TeleporterMessengerAddFeeAmount)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerMessageExecutedIterator is returned from FilterMessageExecuted and is used to iterate over the raw logs and unpacked data for MessageExecuted events raised by the TeleporterMessenger contract.
type TeleporterMessengerMessageExecutedIterator struct {
	Event *TeleporterMessengerMessageExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterMessengerMessageExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerMessageExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeleporterMessengerMessageExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeleporterMessengerMessageExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerMessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerMessageExecuted represents a MessageExecuted event raised by the TeleporterMessenger contract.
type TeleporterMessengerMessageExecuted struct {
	OriginChainID [32]byte
	MessageID     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMessageExecuted is a free log retrieval operation binding the contract event 0x5ad362d54cba0e49d358be9ce586a7136d10a2533579c4460b7e48ec273083ef.
//
// Solidity: event MessageExecuted(bytes32 indexed originChainID, uint256 indexed messageID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterMessageExecuted(opts *bind.FilterOpts, originChainID [][32]byte, messageID []*big.Int) (*TeleporterMessengerMessageExecutedIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "MessageExecuted", originChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerMessageExecutedIterator{contract: _TeleporterMessenger.contract, event: "MessageExecuted", logs: logs, sub: sub}, nil
}

// WatchMessageExecuted is a free log subscription operation binding the contract event 0x5ad362d54cba0e49d358be9ce586a7136d10a2533579c4460b7e48ec273083ef.
//
// Solidity: event MessageExecuted(bytes32 indexed originChainID, uint256 indexed messageID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchMessageExecuted(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerMessageExecuted, originChainID [][32]byte, messageID []*big.Int) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "MessageExecuted", originChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerMessageExecuted)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageExecuted is a log parse operation binding the contract event 0x5ad362d54cba0e49d358be9ce586a7136d10a2533579c4460b7e48ec273083ef.
//
// Solidity: event MessageExecuted(bytes32 indexed originChainID, uint256 indexed messageID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseMessageExecuted(log types.Log) (*TeleporterMessengerMessageExecuted, error) {
	event := new(TeleporterMessengerMessageExecuted)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerMessageExecutionFailedIterator is returned from FilterMessageExecutionFailed and is used to iterate over the raw logs and unpacked data for MessageExecutionFailed events raised by the TeleporterMessenger contract.
type TeleporterMessengerMessageExecutionFailedIterator struct {
	Event *TeleporterMessengerMessageExecutionFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterMessengerMessageExecutionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerMessageExecutionFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeleporterMessengerMessageExecutionFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeleporterMessengerMessageExecutionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerMessageExecutionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerMessageExecutionFailed represents a MessageExecutionFailed event raised by the TeleporterMessenger contract.
type TeleporterMessengerMessageExecutionFailed struct {
	OriginChainID [32]byte
	MessageID     *big.Int
	Message       TeleporterMessage
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMessageExecutionFailed is a free log retrieval operation binding the contract event 0xbedbbe6103cef0a6c9ecbf6aa23da414542c42d7918bea18aab8b601b2c3a449.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed originChainID, uint256 indexed messageID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterMessageExecutionFailed(opts *bind.FilterOpts, originChainID [][32]byte, messageID []*big.Int) (*TeleporterMessengerMessageExecutionFailedIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "MessageExecutionFailed", originChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerMessageExecutionFailedIterator{contract: _TeleporterMessenger.contract, event: "MessageExecutionFailed", logs: logs, sub: sub}, nil
}

// WatchMessageExecutionFailed is a free log subscription operation binding the contract event 0xbedbbe6103cef0a6c9ecbf6aa23da414542c42d7918bea18aab8b601b2c3a449.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed originChainID, uint256 indexed messageID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchMessageExecutionFailed(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerMessageExecutionFailed, originChainID [][32]byte, messageID []*big.Int) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "MessageExecutionFailed", originChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerMessageExecutionFailed)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "MessageExecutionFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageExecutionFailed is a log parse operation binding the contract event 0xbedbbe6103cef0a6c9ecbf6aa23da414542c42d7918bea18aab8b601b2c3a449.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed originChainID, uint256 indexed messageID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseMessageExecutionFailed(log types.Log) (*TeleporterMessengerMessageExecutionFailed, error) {
	event := new(TeleporterMessengerMessageExecutionFailed)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "MessageExecutionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerReceiveCrossChainMessageIterator is returned from FilterReceiveCrossChainMessage and is used to iterate over the raw logs and unpacked data for ReceiveCrossChainMessage events raised by the TeleporterMessenger contract.
type TeleporterMessengerReceiveCrossChainMessageIterator struct {
	Event *TeleporterMessengerReceiveCrossChainMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterMessengerReceiveCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerReceiveCrossChainMessage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeleporterMessengerReceiveCrossChainMessage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeleporterMessengerReceiveCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerReceiveCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerReceiveCrossChainMessage represents a ReceiveCrossChainMessage event raised by the TeleporterMessenger contract.
type TeleporterMessengerReceiveCrossChainMessage struct {
	OriginChainID  [32]byte
	MessageID      *big.Int
	Deliverer      common.Address
	RewardRedeemer common.Address
	Message        TeleporterMessage
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReceiveCrossChainMessage is a free log retrieval operation binding the contract event 0x6b013241f9192863bc66c1f1e9a01dc592c94592bfed5e1ed380808525679575.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed originChainID, uint256 indexed messageID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterReceiveCrossChainMessage(opts *bind.FilterOpts, originChainID [][32]byte, messageID []*big.Int, deliverer []common.Address) (*TeleporterMessengerReceiveCrossChainMessageIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var delivererRule []interface{}
	for _, delivererItem := range deliverer {
		delivererRule = append(delivererRule, delivererItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "ReceiveCrossChainMessage", originChainIDRule, messageIDRule, delivererRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerReceiveCrossChainMessageIterator{contract: _TeleporterMessenger.contract, event: "ReceiveCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchReceiveCrossChainMessage is a free log subscription operation binding the contract event 0x6b013241f9192863bc66c1f1e9a01dc592c94592bfed5e1ed380808525679575.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed originChainID, uint256 indexed messageID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchReceiveCrossChainMessage(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerReceiveCrossChainMessage, originChainID [][32]byte, messageID []*big.Int, deliverer []common.Address) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var delivererRule []interface{}
	for _, delivererItem := range deliverer {
		delivererRule = append(delivererRule, delivererItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "ReceiveCrossChainMessage", originChainIDRule, messageIDRule, delivererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerReceiveCrossChainMessage)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceiveCrossChainMessage is a log parse operation binding the contract event 0x6b013241f9192863bc66c1f1e9a01dc592c94592bfed5e1ed380808525679575.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed originChainID, uint256 indexed messageID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseReceiveCrossChainMessage(log types.Log) (*TeleporterMessengerReceiveCrossChainMessage, error) {
	event := new(TeleporterMessengerReceiveCrossChainMessage)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerRelayerRewardsRedeemedIterator is returned from FilterRelayerRewardsRedeemed and is used to iterate over the raw logs and unpacked data for RelayerRewardsRedeemed events raised by the TeleporterMessenger contract.
type TeleporterMessengerRelayerRewardsRedeemedIterator struct {
	Event *TeleporterMessengerRelayerRewardsRedeemed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterMessengerRelayerRewardsRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerRelayerRewardsRedeemed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeleporterMessengerRelayerRewardsRedeemed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeleporterMessengerRelayerRewardsRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerRelayerRewardsRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerRelayerRewardsRedeemed represents a RelayerRewardsRedeemed event raised by the TeleporterMessenger contract.
type TeleporterMessengerRelayerRewardsRedeemed struct {
	Redeemer common.Address
	Asset    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRelayerRewardsRedeemed is a free log retrieval operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterRelayerRewardsRedeemed(opts *bind.FilterOpts, redeemer []common.Address, asset []common.Address) (*TeleporterMessengerRelayerRewardsRedeemedIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "RelayerRewardsRedeemed", redeemerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerRelayerRewardsRedeemedIterator{contract: _TeleporterMessenger.contract, event: "RelayerRewardsRedeemed", logs: logs, sub: sub}, nil
}

// WatchRelayerRewardsRedeemed is a free log subscription operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchRelayerRewardsRedeemed(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerRelayerRewardsRedeemed, redeemer []common.Address, asset []common.Address) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "RelayerRewardsRedeemed", redeemerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerRelayerRewardsRedeemed)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "RelayerRewardsRedeemed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerRewardsRedeemed is a log parse operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseRelayerRewardsRedeemed(log types.Log) (*TeleporterMessengerRelayerRewardsRedeemed, error) {
	event := new(TeleporterMessengerRelayerRewardsRedeemed)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "RelayerRewardsRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerSendCrossChainMessageIterator is returned from FilterSendCrossChainMessage and is used to iterate over the raw logs and unpacked data for SendCrossChainMessage events raised by the TeleporterMessenger contract.
type TeleporterMessengerSendCrossChainMessageIterator struct {
	Event *TeleporterMessengerSendCrossChainMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeleporterMessengerSendCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerSendCrossChainMessage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeleporterMessengerSendCrossChainMessage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeleporterMessengerSendCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerSendCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerSendCrossChainMessage represents a SendCrossChainMessage event raised by the TeleporterMessenger contract.
type TeleporterMessengerSendCrossChainMessage struct {
	DestinationChainID [32]byte
	MessageID          *big.Int
	Message            TeleporterMessage
	FeeInfo            TeleporterFeeInfo
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSendCrossChainMessage is a free log retrieval operation binding the contract event 0x0563d357b89128d5a0c37c9b06420836e35d193eaf17f7960fc88e47d1e02f57.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed destinationChainID, uint256 indexed messageID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterSendCrossChainMessage(opts *bind.FilterOpts, destinationChainID [][32]byte, messageID []*big.Int) (*TeleporterMessengerSendCrossChainMessageIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "SendCrossChainMessage", destinationChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerSendCrossChainMessageIterator{contract: _TeleporterMessenger.contract, event: "SendCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchSendCrossChainMessage is a free log subscription operation binding the contract event 0x0563d357b89128d5a0c37c9b06420836e35d193eaf17f7960fc88e47d1e02f57.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed destinationChainID, uint256 indexed messageID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchSendCrossChainMessage(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerSendCrossChainMessage, destinationChainID [][32]byte, messageID []*big.Int) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "SendCrossChainMessage", destinationChainIDRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerSendCrossChainMessage)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSendCrossChainMessage is a log parse operation binding the contract event 0x0563d357b89128d5a0c37c9b06420836e35d193eaf17f7960fc88e47d1e02f57.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed destinationChainID, uint256 indexed messageID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseSendCrossChainMessage(log types.Log) (*TeleporterMessengerSendCrossChainMessage, error) {
	event := new(TeleporterMessengerSendCrossChainMessage)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
