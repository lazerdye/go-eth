// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package digixdao

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DGDInterfaceABI is the input ABI used to generate the binding from.
const DGDInterfaceABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dgds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dgdTokenContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_weiPerNanoDGD\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dgdTokenContract\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiPerNanoDGD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DGDInterface is an auto generated Go binding around an Ethereum contract.
type DGDInterface struct {
	DGDInterfaceCaller     // Read-only binding to the contract
	DGDInterfaceTransactor // Write-only binding to the contract
	DGDInterfaceFilterer   // Log filterer for contract events
}

// DGDInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DGDInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DGDInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DGDInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DGDInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DGDInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DGDInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DGDInterfaceSession struct {
	Contract     *DGDInterface     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DGDInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DGDInterfaceCallerSession struct {
	Contract *DGDInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DGDInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DGDInterfaceTransactorSession struct {
	Contract     *DGDInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DGDInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DGDInterfaceRaw struct {
	Contract *DGDInterface // Generic contract binding to access the raw methods on
}

// DGDInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DGDInterfaceCallerRaw struct {
	Contract *DGDInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// DGDInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DGDInterfaceTransactorRaw struct {
	Contract *DGDInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDGDInterface creates a new instance of DGDInterface, bound to a specific deployed contract.
func NewDGDInterface(address common.Address, backend bind.ContractBackend) (*DGDInterface, error) {
	contract, err := bindDGDInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DGDInterface{DGDInterfaceCaller: DGDInterfaceCaller{contract: contract}, DGDInterfaceTransactor: DGDInterfaceTransactor{contract: contract}, DGDInterfaceFilterer: DGDInterfaceFilterer{contract: contract}}, nil
}

// NewDGDInterfaceCaller creates a new read-only instance of DGDInterface, bound to a specific deployed contract.
func NewDGDInterfaceCaller(address common.Address, caller bind.ContractCaller) (*DGDInterfaceCaller, error) {
	contract, err := bindDGDInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DGDInterfaceCaller{contract: contract}, nil
}

// NewDGDInterfaceTransactor creates a new write-only instance of DGDInterface, bound to a specific deployed contract.
func NewDGDInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*DGDInterfaceTransactor, error) {
	contract, err := bindDGDInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DGDInterfaceTransactor{contract: contract}, nil
}

// NewDGDInterfaceFilterer creates a new log filterer instance of DGDInterface, bound to a specific deployed contract.
func NewDGDInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*DGDInterfaceFilterer, error) {
	contract, err := bindDGDInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DGDInterfaceFilterer{contract: contract}, nil
}

// bindDGDInterface binds a generic wrapper to an already deployed contract.
func bindDGDInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DGDInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DGDInterface *DGDInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DGDInterface.Contract.DGDInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DGDInterface *DGDInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DGDInterface.Contract.DGDInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DGDInterface *DGDInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DGDInterface.Contract.DGDInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DGDInterface *DGDInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DGDInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DGDInterface *DGDInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DGDInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DGDInterface *DGDInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DGDInterface.Contract.contract.Transact(opts, method, params...)
}

// DgdTokenContract is a free data retrieval call binding the contract method 0x41232f6b.
//
// Solidity: function dgdTokenContract() view returns(address)
func (_DGDInterface *DGDInterfaceCaller) DgdTokenContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DGDInterface.contract.Call(opts, &out, "dgdTokenContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DgdTokenContract is a free data retrieval call binding the contract method 0x41232f6b.
//
// Solidity: function dgdTokenContract() view returns(address)
func (_DGDInterface *DGDInterfaceSession) DgdTokenContract() (common.Address, error) {
	return _DGDInterface.Contract.DgdTokenContract(&_DGDInterface.CallOpts)
}

// DgdTokenContract is a free data retrieval call binding the contract method 0x41232f6b.
//
// Solidity: function dgdTokenContract() view returns(address)
func (_DGDInterface *DGDInterfaceCallerSession) DgdTokenContract() (common.Address, error) {
	return _DGDInterface.Contract.DgdTokenContract(&_DGDInterface.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_DGDInterface *DGDInterfaceCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DGDInterface.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_DGDInterface *DGDInterfaceSession) IsInitialized() (bool, error) {
	return _DGDInterface.Contract.IsInitialized(&_DGDInterface.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_DGDInterface *DGDInterfaceCallerSession) IsInitialized() (bool, error) {
	return _DGDInterface.Contract.IsInitialized(&_DGDInterface.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DGDInterface *DGDInterfaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DGDInterface.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DGDInterface *DGDInterfaceSession) Owner() (common.Address, error) {
	return _DGDInterface.Contract.Owner(&_DGDInterface.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DGDInterface *DGDInterfaceCallerSession) Owner() (common.Address, error) {
	return _DGDInterface.Contract.Owner(&_DGDInterface.CallOpts)
}

// WeiPerNanoDGD is a free data retrieval call binding the contract method 0x456a7cd8.
//
// Solidity: function weiPerNanoDGD() view returns(uint256)
func (_DGDInterface *DGDInterfaceCaller) WeiPerNanoDGD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DGDInterface.contract.Call(opts, &out, "weiPerNanoDGD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiPerNanoDGD is a free data retrieval call binding the contract method 0x456a7cd8.
//
// Solidity: function weiPerNanoDGD() view returns(uint256)
func (_DGDInterface *DGDInterfaceSession) WeiPerNanoDGD() (*big.Int, error) {
	return _DGDInterface.Contract.WeiPerNanoDGD(&_DGDInterface.CallOpts)
}

// WeiPerNanoDGD is a free data retrieval call binding the contract method 0x456a7cd8.
//
// Solidity: function weiPerNanoDGD() view returns(uint256)
func (_DGDInterface *DGDInterfaceCallerSession) WeiPerNanoDGD() (*big.Int, error) {
	return _DGDInterface.Contract.WeiPerNanoDGD(&_DGDInterface.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns(bool _success)
func (_DGDInterface *DGDInterfaceTransactor) Burn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DGDInterface.contract.Transact(opts, "burn")
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns(bool _success)
func (_DGDInterface *DGDInterfaceSession) Burn() (*types.Transaction, error) {
	return _DGDInterface.Contract.Burn(&_DGDInterface.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns(bool _success)
func (_DGDInterface *DGDInterfaceTransactorSession) Burn() (*types.Transaction, error) {
	return _DGDInterface.Contract.Burn(&_DGDInterface.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xb792e6ec.
//
// Solidity: function init(uint256 _weiPerNanoDGD, address _dgdTokenContract) returns(bool _success)
func (_DGDInterface *DGDInterfaceTransactor) Init(opts *bind.TransactOpts, _weiPerNanoDGD *big.Int, _dgdTokenContract common.Address) (*types.Transaction, error) {
	return _DGDInterface.contract.Transact(opts, "init", _weiPerNanoDGD, _dgdTokenContract)
}

// Init is a paid mutator transaction binding the contract method 0xb792e6ec.
//
// Solidity: function init(uint256 _weiPerNanoDGD, address _dgdTokenContract) returns(bool _success)
func (_DGDInterface *DGDInterfaceSession) Init(_weiPerNanoDGD *big.Int, _dgdTokenContract common.Address) (*types.Transaction, error) {
	return _DGDInterface.Contract.Init(&_DGDInterface.TransactOpts, _weiPerNanoDGD, _dgdTokenContract)
}

// Init is a paid mutator transaction binding the contract method 0xb792e6ec.
//
// Solidity: function init(uint256 _weiPerNanoDGD, address _dgdTokenContract) returns(bool _success)
func (_DGDInterface *DGDInterfaceTransactorSession) Init(_weiPerNanoDGD *big.Int, _dgdTokenContract common.Address) (*types.Transaction, error) {
	return _DGDInterface.Contract.Init(&_DGDInterface.TransactOpts, _weiPerNanoDGD, _dgdTokenContract)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DGDInterface *DGDInterfaceTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _DGDInterface.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DGDInterface *DGDInterfaceSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DGDInterface.Contract.Fallback(&_DGDInterface.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DGDInterface *DGDInterfaceTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DGDInterface.Contract.Fallback(&_DGDInterface.TransactOpts, calldata)
}

// DGDInterfaceRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the DGDInterface contract.
type DGDInterfaceRefundIterator struct {
	Event *DGDInterfaceRefund // Event containing the contract specifics and raw log

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
func (it *DGDInterfaceRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DGDInterfaceRefund)
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
		it.Event = new(DGDInterfaceRefund)
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
func (it *DGDInterfaceRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DGDInterfaceRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DGDInterfaceRefund represents a Refund event raised by the DGDInterface contract.
type DGDInterfaceRefund struct {
	User         common.Address
	Dgds         *big.Int
	RefundAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0x73f04af9dcc582a923ec15d3eea990fe34adabfff2879e28d44572e01a54abb6.
//
// Solidity: event Refund(address indexed user, uint256 indexed dgds, uint256 refundAmount)
func (_DGDInterface *DGDInterfaceFilterer) FilterRefund(opts *bind.FilterOpts, user []common.Address, dgds []*big.Int) (*DGDInterfaceRefundIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var dgdsRule []interface{}
	for _, dgdsItem := range dgds {
		dgdsRule = append(dgdsRule, dgdsItem)
	}

	logs, sub, err := _DGDInterface.contract.FilterLogs(opts, "Refund", userRule, dgdsRule)
	if err != nil {
		return nil, err
	}
	return &DGDInterfaceRefundIterator{contract: _DGDInterface.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0x73f04af9dcc582a923ec15d3eea990fe34adabfff2879e28d44572e01a54abb6.
//
// Solidity: event Refund(address indexed user, uint256 indexed dgds, uint256 refundAmount)
func (_DGDInterface *DGDInterfaceFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *DGDInterfaceRefund, user []common.Address, dgds []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var dgdsRule []interface{}
	for _, dgdsItem := range dgds {
		dgdsRule = append(dgdsRule, dgdsItem)
	}

	logs, sub, err := _DGDInterface.contract.WatchLogs(opts, "Refund", userRule, dgdsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DGDInterfaceRefund)
				if err := _DGDInterface.contract.UnpackLog(event, "Refund", log); err != nil {
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

// ParseRefund is a log parse operation binding the contract event 0x73f04af9dcc582a923ec15d3eea990fe34adabfff2879e28d44572e01a54abb6.
//
// Solidity: event Refund(address indexed user, uint256 indexed dgds, uint256 refundAmount)
func (_DGDInterface *DGDInterfaceFilterer) ParseRefund(log types.Log) (*DGDInterfaceRefund, error) {
	event := new(DGDInterfaceRefund)
	if err := _DGDInterface.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
