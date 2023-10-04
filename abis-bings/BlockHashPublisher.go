// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abisbings

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

// BlockHashPublisherMetaData contains all meta data concerning the BlockHashPublisher contract.
var BlockHashPublisherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"PublishBlockHash\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RECEIVE_BLOCK_HASH_REQUIRED_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"}],\"name\":\"publishLatestBlockHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BlockHashPublisherABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockHashPublisherMetaData.ABI instead.
var BlockHashPublisherABI = BlockHashPublisherMetaData.ABI

// BlockHashPublisher is an auto generated Go binding around an Ethereum contract.
type BlockHashPublisher struct {
	BlockHashPublisherCaller     // Read-only binding to the contract
	BlockHashPublisherTransactor // Write-only binding to the contract
	BlockHashPublisherFilterer   // Log filterer for contract events
}

// BlockHashPublisherCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockHashPublisherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockHashPublisherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockHashPublisherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockHashPublisherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockHashPublisherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockHashPublisherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockHashPublisherSession struct {
	Contract     *BlockHashPublisher // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BlockHashPublisherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockHashPublisherCallerSession struct {
	Contract *BlockHashPublisherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BlockHashPublisherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockHashPublisherTransactorSession struct {
	Contract     *BlockHashPublisherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BlockHashPublisherRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockHashPublisherRaw struct {
	Contract *BlockHashPublisher // Generic contract binding to access the raw methods on
}

// BlockHashPublisherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockHashPublisherCallerRaw struct {
	Contract *BlockHashPublisherCaller // Generic read-only contract binding to access the raw methods on
}

// BlockHashPublisherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockHashPublisherTransactorRaw struct {
	Contract *BlockHashPublisherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockHashPublisher creates a new instance of BlockHashPublisher, bound to a specific deployed contract.
func NewBlockHashPublisher(address common.Address, backend bind.ContractBackend) (*BlockHashPublisher, error) {
	contract, err := bindBlockHashPublisher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockHashPublisher{BlockHashPublisherCaller: BlockHashPublisherCaller{contract: contract}, BlockHashPublisherTransactor: BlockHashPublisherTransactor{contract: contract}, BlockHashPublisherFilterer: BlockHashPublisherFilterer{contract: contract}}, nil
}

// NewBlockHashPublisherCaller creates a new read-only instance of BlockHashPublisher, bound to a specific deployed contract.
func NewBlockHashPublisherCaller(address common.Address, caller bind.ContractCaller) (*BlockHashPublisherCaller, error) {
	contract, err := bindBlockHashPublisher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockHashPublisherCaller{contract: contract}, nil
}

// NewBlockHashPublisherTransactor creates a new write-only instance of BlockHashPublisher, bound to a specific deployed contract.
func NewBlockHashPublisherTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockHashPublisherTransactor, error) {
	contract, err := bindBlockHashPublisher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockHashPublisherTransactor{contract: contract}, nil
}

// NewBlockHashPublisherFilterer creates a new log filterer instance of BlockHashPublisher, bound to a specific deployed contract.
func NewBlockHashPublisherFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockHashPublisherFilterer, error) {
	contract, err := bindBlockHashPublisher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockHashPublisherFilterer{contract: contract}, nil
}

// bindBlockHashPublisher binds a generic wrapper to an already deployed contract.
func bindBlockHashPublisher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockHashPublisherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockHashPublisher *BlockHashPublisherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockHashPublisher.Contract.BlockHashPublisherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockHashPublisher *BlockHashPublisherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockHashPublisher.Contract.BlockHashPublisherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockHashPublisher *BlockHashPublisherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockHashPublisher.Contract.BlockHashPublisherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockHashPublisher *BlockHashPublisherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockHashPublisher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockHashPublisher *BlockHashPublisherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockHashPublisher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockHashPublisher *BlockHashPublisherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockHashPublisher.Contract.contract.Transact(opts, method, params...)
}

// RECEIVEBLOCKHASHREQUIREDGASLIMIT is a free data retrieval call binding the contract method 0x18aef19b.
//
// Solidity: function RECEIVE_BLOCK_HASH_REQUIRED_GAS_LIMIT() view returns(uint256)
func (_BlockHashPublisher *BlockHashPublisherCaller) RECEIVEBLOCKHASHREQUIREDGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockHashPublisher.contract.Call(opts, &out, "RECEIVE_BLOCK_HASH_REQUIRED_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RECEIVEBLOCKHASHREQUIREDGASLIMIT is a free data retrieval call binding the contract method 0x18aef19b.
//
// Solidity: function RECEIVE_BLOCK_HASH_REQUIRED_GAS_LIMIT() view returns(uint256)
func (_BlockHashPublisher *BlockHashPublisherSession) RECEIVEBLOCKHASHREQUIREDGASLIMIT() (*big.Int, error) {
	return _BlockHashPublisher.Contract.RECEIVEBLOCKHASHREQUIREDGASLIMIT(&_BlockHashPublisher.CallOpts)
}

// RECEIVEBLOCKHASHREQUIREDGASLIMIT is a free data retrieval call binding the contract method 0x18aef19b.
//
// Solidity: function RECEIVE_BLOCK_HASH_REQUIRED_GAS_LIMIT() view returns(uint256)
func (_BlockHashPublisher *BlockHashPublisherCallerSession) RECEIVEBLOCKHASHREQUIREDGASLIMIT() (*big.Int, error) {
	return _BlockHashPublisher.Contract.RECEIVEBLOCKHASHREQUIREDGASLIMIT(&_BlockHashPublisher.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_BlockHashPublisher *BlockHashPublisherCaller) TeleporterMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockHashPublisher.contract.Call(opts, &out, "teleporterMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_BlockHashPublisher *BlockHashPublisherSession) TeleporterMessenger() (common.Address, error) {
	return _BlockHashPublisher.Contract.TeleporterMessenger(&_BlockHashPublisher.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_BlockHashPublisher *BlockHashPublisherCallerSession) TeleporterMessenger() (common.Address, error) {
	return _BlockHashPublisher.Contract.TeleporterMessenger(&_BlockHashPublisher.CallOpts)
}

// PublishLatestBlockHash is a paid mutator transaction binding the contract method 0x82ab2b86.
//
// Solidity: function publishLatestBlockHash(bytes32 destinationChainID, address destinationAddress) returns(uint256 messageID)
func (_BlockHashPublisher *BlockHashPublisherTransactor) PublishLatestBlockHash(opts *bind.TransactOpts, destinationChainID [32]byte, destinationAddress common.Address) (*types.Transaction, error) {
	return _BlockHashPublisher.contract.Transact(opts, "publishLatestBlockHash", destinationChainID, destinationAddress)
}

// PublishLatestBlockHash is a paid mutator transaction binding the contract method 0x82ab2b86.
//
// Solidity: function publishLatestBlockHash(bytes32 destinationChainID, address destinationAddress) returns(uint256 messageID)
func (_BlockHashPublisher *BlockHashPublisherSession) PublishLatestBlockHash(destinationChainID [32]byte, destinationAddress common.Address) (*types.Transaction, error) {
	return _BlockHashPublisher.Contract.PublishLatestBlockHash(&_BlockHashPublisher.TransactOpts, destinationChainID, destinationAddress)
}

// PublishLatestBlockHash is a paid mutator transaction binding the contract method 0x82ab2b86.
//
// Solidity: function publishLatestBlockHash(bytes32 destinationChainID, address destinationAddress) returns(uint256 messageID)
func (_BlockHashPublisher *BlockHashPublisherTransactorSession) PublishLatestBlockHash(destinationChainID [32]byte, destinationAddress common.Address) (*types.Transaction, error) {
	return _BlockHashPublisher.Contract.PublishLatestBlockHash(&_BlockHashPublisher.TransactOpts, destinationChainID, destinationAddress)
}

// BlockHashPublisherPublishBlockHashIterator is returned from FilterPublishBlockHash and is used to iterate over the raw logs and unpacked data for PublishBlockHash events raised by the BlockHashPublisher contract.
type BlockHashPublisherPublishBlockHashIterator struct {
	Event *BlockHashPublisherPublishBlockHash // Event containing the contract specifics and raw log

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
func (it *BlockHashPublisherPublishBlockHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockHashPublisherPublishBlockHash)
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
		it.Event = new(BlockHashPublisherPublishBlockHash)
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
func (it *BlockHashPublisherPublishBlockHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockHashPublisherPublishBlockHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockHashPublisherPublishBlockHash represents a PublishBlockHash event raised by the BlockHashPublisher contract.
type BlockHashPublisherPublishBlockHash struct {
	DestinationChainID [32]byte
	DestinationAddress common.Address
	BlockHeight        *big.Int
	BlockHash          [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPublishBlockHash is a free log retrieval operation binding the contract event 0xe13623d33d18131ce960c33b1282ceac1fe7b5ccfcf7f8c0f6dad32dd61e3bdd.
//
// Solidity: event PublishBlockHash(bytes32 indexed destinationChainID, address indexed destinationAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_BlockHashPublisher *BlockHashPublisherFilterer) FilterPublishBlockHash(opts *bind.FilterOpts, destinationChainID [][32]byte, destinationAddress []common.Address, blockHeight []*big.Int) (*BlockHashPublisherPublishBlockHashIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}
	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	logs, sub, err := _BlockHashPublisher.contract.FilterLogs(opts, "PublishBlockHash", destinationChainIDRule, destinationAddressRule, blockHeightRule)
	if err != nil {
		return nil, err
	}
	return &BlockHashPublisherPublishBlockHashIterator{contract: _BlockHashPublisher.contract, event: "PublishBlockHash", logs: logs, sub: sub}, nil
}

// WatchPublishBlockHash is a free log subscription operation binding the contract event 0xe13623d33d18131ce960c33b1282ceac1fe7b5ccfcf7f8c0f6dad32dd61e3bdd.
//
// Solidity: event PublishBlockHash(bytes32 indexed destinationChainID, address indexed destinationAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_BlockHashPublisher *BlockHashPublisherFilterer) WatchPublishBlockHash(opts *bind.WatchOpts, sink chan<- *BlockHashPublisherPublishBlockHash, destinationChainID [][32]byte, destinationAddress []common.Address, blockHeight []*big.Int) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}
	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	logs, sub, err := _BlockHashPublisher.contract.WatchLogs(opts, "PublishBlockHash", destinationChainIDRule, destinationAddressRule, blockHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockHashPublisherPublishBlockHash)
				if err := _BlockHashPublisher.contract.UnpackLog(event, "PublishBlockHash", log); err != nil {
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

// ParsePublishBlockHash is a log parse operation binding the contract event 0xe13623d33d18131ce960c33b1282ceac1fe7b5ccfcf7f8c0f6dad32dd61e3bdd.
//
// Solidity: event PublishBlockHash(bytes32 indexed destinationChainID, address indexed destinationAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_BlockHashPublisher *BlockHashPublisherFilterer) ParsePublishBlockHash(log types.Log) (*BlockHashPublisherPublishBlockHash, error) {
	event := new(BlockHashPublisherPublishBlockHash)
	if err := _BlockHashPublisher.contract.UnpackLog(event, "PublishBlockHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
