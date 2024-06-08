// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"param1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int32\",\"name\":\"param2\",\"type\":\"int32\"},{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"param3\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"param4\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"param5\",\"type\":\"uint256\"}],\"name\":\"ExampleEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_param1\",\"type\":\"address\"},{\"internalType\":\"int32\",\"name\":\"_param2\",\"type\":\"int32\"},{\"internalType\":\"int256\",\"name\":\"_param3\",\"type\":\"int256\"},{\"internalType\":\"address\",\"name\":\"_param4\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_param5\",\"type\":\"uint256\"}],\"name\":\"triggerEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506102918061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c806319c94dd01461002d575b5f80fd5b6100476004803603810190610042919061019f565b610049565b005b828460030b8673ffffffffffffffffffffffffffffffffffffffff167f43c5dd2ecf8a04a86d029ffc3601ce4078b42b82e56127fcb95985943b05308f8585604051610096929190610234565b60405180910390a45050505050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100d2826100a9565b9050919050565b6100e2816100c8565b81146100ec575f80fd5b50565b5f813590506100fd816100d9565b92915050565b5f8160030b9050919050565b61011881610103565b8114610122575f80fd5b50565b5f813590506101338161010f565b92915050565b5f819050919050565b61014b81610139565b8114610155575f80fd5b50565b5f8135905061016681610142565b92915050565b5f819050919050565b61017e8161016c565b8114610188575f80fd5b50565b5f8135905061019981610175565b92915050565b5f805f805f60a086880312156101b8576101b76100a5565b5b5f6101c5888289016100ef565b95505060206101d688828901610125565b94505060406101e788828901610158565b93505060606101f8888289016100ef565b92505060806102098882890161018b565b9150509295509295909350565b61021f816100c8565b82525050565b61022e8161016c565b82525050565b5f6040820190506102475f830185610216565b6102546020830184610225565b939250505056fea2646970667358221220e220ae4f311048a75bc6dac5646fda2eccce636d6db1796e950c7febca552c5364736f6c63430008190033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// TriggerEvent is a paid mutator transaction binding the contract method 0x19c94dd0.
//
// Solidity: function triggerEvent(address _param1, int32 _param2, int256 _param3, address _param4, uint256 _param5) returns()
func (_Contract *ContractTransactor) TriggerEvent(opts *bind.TransactOpts, _param1 common.Address, _param2 int32, _param3 *big.Int, _param4 common.Address, _param5 *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "triggerEvent", _param1, _param2, _param3, _param4, _param5)
}

// TriggerEvent is a paid mutator transaction binding the contract method 0x19c94dd0.
//
// Solidity: function triggerEvent(address _param1, int32 _param2, int256 _param3, address _param4, uint256 _param5) returns()
func (_Contract *ContractSession) TriggerEvent(_param1 common.Address, _param2 int32, _param3 *big.Int, _param4 common.Address, _param5 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TriggerEvent(&_Contract.TransactOpts, _param1, _param2, _param3, _param4, _param5)
}

// TriggerEvent is a paid mutator transaction binding the contract method 0x19c94dd0.
//
// Solidity: function triggerEvent(address _param1, int32 _param2, int256 _param3, address _param4, uint256 _param5) returns()
func (_Contract *ContractTransactorSession) TriggerEvent(_param1 common.Address, _param2 int32, _param3 *big.Int, _param4 common.Address, _param5 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TriggerEvent(&_Contract.TransactOpts, _param1, _param2, _param3, _param4, _param5)
}

// ContractExampleEventIterator is returned from FilterExampleEvent and is used to iterate over the raw logs and unpacked data for ExampleEvent events raised by the Contract contract.
type ContractExampleEventIterator struct {
	Event *ContractExampleEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractExampleEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExampleEvent)
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
		it.Event = new(ContractExampleEvent)
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
func (it *ContractExampleEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExampleEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExampleEvent represents a ExampleEvent event raised by the Contract contract.
type ContractExampleEvent struct {
	Param1 common.Address
	Param2 int32
	Param3 *big.Int
	Param4 common.Address
	Param5 *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExampleEvent is a free log retrieval operation binding the contract event 0x43c5dd2ecf8a04a86d029ffc3601ce4078b42b82e56127fcb95985943b05308f.
//
// Solidity: event ExampleEvent(address indexed param1, int32 indexed param2, int256 indexed param3, address param4, uint256 param5)
func (_Contract *ContractFilterer) FilterExampleEvent(opts *bind.FilterOpts, param1 []common.Address, param2 []int32, param3 []*big.Int) (*ContractExampleEventIterator, error) {

	var param1Rule []interface{}
	for _, param1Item := range param1 {
		param1Rule = append(param1Rule, param1Item)
	}
	var param2Rule []interface{}
	for _, param2Item := range param2 {
		param2Rule = append(param2Rule, param2Item)
	}
	var param3Rule []interface{}
	for _, param3Item := range param3 {
		param3Rule = append(param3Rule, param3Item)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ExampleEvent", param1Rule, param2Rule, param3Rule)
	if err != nil {
		return nil, err
	}
	return &ContractExampleEventIterator{contract: _Contract.contract, event: "ExampleEvent", logs: logs, sub: sub}, nil
}

// WatchExampleEvent is a free log subscription operation binding the contract event 0x43c5dd2ecf8a04a86d029ffc3601ce4078b42b82e56127fcb95985943b05308f.
//
// Solidity: event ExampleEvent(address indexed param1, int32 indexed param2, int256 indexed param3, address param4, uint256 param5)
func (_Contract *ContractFilterer) WatchExampleEvent(opts *bind.WatchOpts, sink chan<- *ContractExampleEvent, param1 []common.Address, param2 []int32, param3 []*big.Int) (event.Subscription, error) {

	var param1Rule []interface{}
	for _, param1Item := range param1 {
		param1Rule = append(param1Rule, param1Item)
	}
	var param2Rule []interface{}
	for _, param2Item := range param2 {
		param2Rule = append(param2Rule, param2Item)
	}
	var param3Rule []interface{}
	for _, param3Item := range param3 {
		param3Rule = append(param3Rule, param3Item)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ExampleEvent", param1Rule, param2Rule, param3Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExampleEvent)
				if err := _Contract.contract.UnpackLog(event, "ExampleEvent", log); err != nil {
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

// ParseExampleEvent is a log parse operation binding the contract event 0x43c5dd2ecf8a04a86d029ffc3601ce4078b42b82e56127fcb95985943b05308f.
//
// Solidity: event ExampleEvent(address indexed param1, int32 indexed param2, int256 indexed param3, address param4, uint256 param5)
func (_Contract *ContractFilterer) ParseExampleEvent(log types.Log) (*ContractExampleEvent, error) {
	event := new(ContractExampleEvent)
	if err := _Contract.contract.UnpackLog(event, "ExampleEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
