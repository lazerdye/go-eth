// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv2

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

// TokenDistributorABI is the input ABI used to generate the binding from.
const TokenDistributorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TokenDistributor is an auto generated Go binding around an Ethereum contract.
type TokenDistributor struct {
	TokenDistributorCaller     // Read-only binding to the contract
	TokenDistributorTransactor // Write-only binding to the contract
	TokenDistributorFilterer   // Log filterer for contract events
}

// TokenDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenDistributorSession struct {
	Contract     *TokenDistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenDistributorCallerSession struct {
	Contract *TokenDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TokenDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenDistributorTransactorSession struct {
	Contract     *TokenDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TokenDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenDistributorRaw struct {
	Contract *TokenDistributor // Generic contract binding to access the raw methods on
}

// TokenDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenDistributorCallerRaw struct {
	Contract *TokenDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// TokenDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenDistributorTransactorRaw struct {
	Contract *TokenDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenDistributor creates a new instance of TokenDistributor, bound to a specific deployed contract.
func NewTokenDistributor(address common.Address, backend bind.ContractBackend) (*TokenDistributor, error) {
	contract, err := bindTokenDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenDistributor{TokenDistributorCaller: TokenDistributorCaller{contract: contract}, TokenDistributorTransactor: TokenDistributorTransactor{contract: contract}, TokenDistributorFilterer: TokenDistributorFilterer{contract: contract}}, nil
}

// NewTokenDistributorCaller creates a new read-only instance of TokenDistributor, bound to a specific deployed contract.
func NewTokenDistributorCaller(address common.Address, caller bind.ContractCaller) (*TokenDistributorCaller, error) {
	contract, err := bindTokenDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenDistributorCaller{contract: contract}, nil
}

// NewTokenDistributorTransactor creates a new write-only instance of TokenDistributor, bound to a specific deployed contract.
func NewTokenDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenDistributorTransactor, error) {
	contract, err := bindTokenDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenDistributorTransactor{contract: contract}, nil
}

// NewTokenDistributorFilterer creates a new log filterer instance of TokenDistributor, bound to a specific deployed contract.
func NewTokenDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenDistributorFilterer, error) {
	contract, err := bindTokenDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenDistributorFilterer{contract: contract}, nil
}

// bindTokenDistributor binds a generic wrapper to an already deployed contract.
func bindTokenDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenDistributorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenDistributor *TokenDistributorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenDistributor.Contract.TokenDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenDistributor *TokenDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenDistributor.Contract.TokenDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenDistributor *TokenDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenDistributor.Contract.TokenDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenDistributor *TokenDistributorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenDistributor *TokenDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenDistributor *TokenDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenDistributor.Contract.contract.Transact(opts, method, params...)
}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 index) view returns(bool)
func (_TokenDistributor *TokenDistributorCaller) IsClaimed(opts *bind.CallOpts, index *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenDistributor.contract.Call(opts, out, "isClaimed", index)
	return *ret0, err
}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 index) view returns(bool)
func (_TokenDistributor *TokenDistributorSession) IsClaimed(index *big.Int) (bool, error) {
	return _TokenDistributor.Contract.IsClaimed(&_TokenDistributor.CallOpts, index)
}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 index) view returns(bool)
func (_TokenDistributor *TokenDistributorCallerSession) IsClaimed(index *big.Int) (bool, error) {
	return _TokenDistributor.Contract.IsClaimed(&_TokenDistributor.CallOpts, index)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_TokenDistributor *TokenDistributorCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TokenDistributor.contract.Call(opts, out, "merkleRoot")
	return *ret0, err
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_TokenDistributor *TokenDistributorSession) MerkleRoot() ([32]byte, error) {
	return _TokenDistributor.Contract.MerkleRoot(&_TokenDistributor.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_TokenDistributor *TokenDistributorCallerSession) MerkleRoot() ([32]byte, error) {
	return _TokenDistributor.Contract.MerkleRoot(&_TokenDistributor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TokenDistributor *TokenDistributorCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenDistributor.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TokenDistributor *TokenDistributorSession) Token() (common.Address, error) {
	return _TokenDistributor.Contract.Token(&_TokenDistributor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TokenDistributor *TokenDistributorCallerSession) Token() (common.Address, error) {
	return _TokenDistributor.Contract.Token(&_TokenDistributor.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x2e7ba6ef.
//
// Solidity: function claim(uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns()
func (_TokenDistributor *TokenDistributorTransactor) Claim(opts *bind.TransactOpts, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _TokenDistributor.contract.Transact(opts, "claim", index, account, amount, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x2e7ba6ef.
//
// Solidity: function claim(uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns()
func (_TokenDistributor *TokenDistributorSession) Claim(index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _TokenDistributor.Contract.Claim(&_TokenDistributor.TransactOpts, index, account, amount, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x2e7ba6ef.
//
// Solidity: function claim(uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns()
func (_TokenDistributor *TokenDistributorTransactorSession) Claim(index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _TokenDistributor.Contract.Claim(&_TokenDistributor.TransactOpts, index, account, amount, merkleProof)
}

// TokenDistributorClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the TokenDistributor contract.
type TokenDistributorClaimedIterator struct {
	Event *TokenDistributorClaimed // Event containing the contract specifics and raw log

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
func (it *TokenDistributorClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenDistributorClaimed)
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
		it.Event = new(TokenDistributorClaimed)
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
func (it *TokenDistributorClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenDistributorClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenDistributorClaimed represents a Claimed event raised by the TokenDistributor contract.
type TokenDistributorClaimed struct {
	Index   *big.Int
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x4ec90e965519d92681267467f775ada5bd214aa92c0dc93d90a5e880ce9ed026.
//
// Solidity: event Claimed(uint256 index, address account, uint256 amount)
func (_TokenDistributor *TokenDistributorFilterer) FilterClaimed(opts *bind.FilterOpts) (*TokenDistributorClaimedIterator, error) {

	logs, sub, err := _TokenDistributor.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &TokenDistributorClaimedIterator{contract: _TokenDistributor.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x4ec90e965519d92681267467f775ada5bd214aa92c0dc93d90a5e880ce9ed026.
//
// Solidity: event Claimed(uint256 index, address account, uint256 amount)
func (_TokenDistributor *TokenDistributorFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *TokenDistributorClaimed) (event.Subscription, error) {

	logs, sub, err := _TokenDistributor.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenDistributorClaimed)
				if err := _TokenDistributor.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x4ec90e965519d92681267467f775ada5bd214aa92c0dc93d90a5e880ce9ed026.
//
// Solidity: event Claimed(uint256 index, address account, uint256 amount)
func (_TokenDistributor *TokenDistributorFilterer) ParseClaimed(log types.Log) (*TokenDistributorClaimed, error) {
	event := new(TokenDistributorClaimed)
	if err := _TokenDistributor.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	return event, nil
}
