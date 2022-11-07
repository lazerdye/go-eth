// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package maker

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
)

// GemJoinMetaData contains all meta data concerning the GemJoin contract.
var GemJoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"gem_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"contractGemLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ilk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"join\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GemJoinABI is the input ABI used to generate the binding from.
// Deprecated: Use GemJoinMetaData.ABI instead.
var GemJoinABI = GemJoinMetaData.ABI

// GemJoin is an auto generated Go binding around an Ethereum contract.
type GemJoin struct {
	GemJoinCaller     // Read-only binding to the contract
	GemJoinTransactor // Write-only binding to the contract
	GemJoinFilterer   // Log filterer for contract events
}

// GemJoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type GemJoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GemJoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GemJoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GemJoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GemJoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GemJoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GemJoinSession struct {
	Contract     *GemJoin          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GemJoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GemJoinCallerSession struct {
	Contract *GemJoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GemJoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GemJoinTransactorSession struct {
	Contract     *GemJoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GemJoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type GemJoinRaw struct {
	Contract *GemJoin // Generic contract binding to access the raw methods on
}

// GemJoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GemJoinCallerRaw struct {
	Contract *GemJoinCaller // Generic read-only contract binding to access the raw methods on
}

// GemJoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GemJoinTransactorRaw struct {
	Contract *GemJoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGemJoin creates a new instance of GemJoin, bound to a specific deployed contract.
func NewGemJoin(address common.Address, backend bind.ContractBackend) (*GemJoin, error) {
	contract, err := bindGemJoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GemJoin{GemJoinCaller: GemJoinCaller{contract: contract}, GemJoinTransactor: GemJoinTransactor{contract: contract}, GemJoinFilterer: GemJoinFilterer{contract: contract}}, nil
}

// NewGemJoinCaller creates a new read-only instance of GemJoin, bound to a specific deployed contract.
func NewGemJoinCaller(address common.Address, caller bind.ContractCaller) (*GemJoinCaller, error) {
	contract, err := bindGemJoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GemJoinCaller{contract: contract}, nil
}

// NewGemJoinTransactor creates a new write-only instance of GemJoin, bound to a specific deployed contract.
func NewGemJoinTransactor(address common.Address, transactor bind.ContractTransactor) (*GemJoinTransactor, error) {
	contract, err := bindGemJoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GemJoinTransactor{contract: contract}, nil
}

// NewGemJoinFilterer creates a new log filterer instance of GemJoin, bound to a specific deployed contract.
func NewGemJoinFilterer(address common.Address, filterer bind.ContractFilterer) (*GemJoinFilterer, error) {
	contract, err := bindGemJoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GemJoinFilterer{contract: contract}, nil
}

// bindGemJoin binds a generic wrapper to an already deployed contract.
func bindGemJoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GemJoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GemJoin *GemJoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GemJoin.Contract.GemJoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GemJoin *GemJoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GemJoin.Contract.GemJoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GemJoin *GemJoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GemJoin.Contract.GemJoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GemJoin *GemJoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GemJoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GemJoin *GemJoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GemJoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GemJoin *GemJoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GemJoin.Contract.contract.Transact(opts, method, params...)
}

// Dec is a free data retrieval call binding the contract method 0xb3bcfa82.
//
// Solidity: function dec() view returns(uint256)
func (_GemJoin *GemJoinCaller) Dec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GemJoin.contract.Call(opts, &out, "dec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dec is a free data retrieval call binding the contract method 0xb3bcfa82.
//
// Solidity: function dec() view returns(uint256)
func (_GemJoin *GemJoinSession) Dec() (*big.Int, error) {
	return _GemJoin.Contract.Dec(&_GemJoin.CallOpts)
}

// Dec is a free data retrieval call binding the contract method 0xb3bcfa82.
//
// Solidity: function dec() view returns(uint256)
func (_GemJoin *GemJoinCallerSession) Dec() (*big.Int, error) {
	return _GemJoin.Contract.Dec(&_GemJoin.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_GemJoin *GemJoinCaller) Gem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GemJoin.contract.Call(opts, &out, "gem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_GemJoin *GemJoinSession) Gem() (common.Address, error) {
	return _GemJoin.Contract.Gem(&_GemJoin.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_GemJoin *GemJoinCallerSession) Gem() (common.Address, error) {
	return _GemJoin.Contract.Gem(&_GemJoin.CallOpts)
}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_GemJoin *GemJoinCaller) Ilk(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GemJoin.contract.Call(opts, &out, "ilk")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_GemJoin *GemJoinSession) Ilk() ([32]byte, error) {
	return _GemJoin.Contract.Ilk(&_GemJoin.CallOpts)
}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_GemJoin *GemJoinCallerSession) Ilk() ([32]byte, error) {
	return _GemJoin.Contract.Ilk(&_GemJoin.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_GemJoin *GemJoinCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GemJoin.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_GemJoin *GemJoinSession) Live() (*big.Int, error) {
	return _GemJoin.Contract.Live(&_GemJoin.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_GemJoin *GemJoinCallerSession) Live() (*big.Int, error) {
	return _GemJoin.Contract.Live(&_GemJoin.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_GemJoin *GemJoinCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GemJoin.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_GemJoin *GemJoinSession) Vat() (common.Address, error) {
	return _GemJoin.Contract.Vat(&_GemJoin.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_GemJoin *GemJoinCallerSession) Vat() (common.Address, error) {
	return _GemJoin.Contract.Vat(&_GemJoin.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_GemJoin *GemJoinCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GemJoin.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_GemJoin *GemJoinSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _GemJoin.Contract.Wards(&_GemJoin.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_GemJoin *GemJoinCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _GemJoin.Contract.Wards(&_GemJoin.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_GemJoin *GemJoinTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GemJoin.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_GemJoin *GemJoinSession) Cage() (*types.Transaction, error) {
	return _GemJoin.Contract.Cage(&_GemJoin.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_GemJoin *GemJoinTransactorSession) Cage() (*types.Transaction, error) {
	return _GemJoin.Contract.Cage(&_GemJoin.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_GemJoin *GemJoinTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _GemJoin.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_GemJoin *GemJoinSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _GemJoin.Contract.Deny(&_GemJoin.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_GemJoin *GemJoinTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _GemJoin.Contract.Deny(&_GemJoin.TransactOpts, usr)
}

// Exit is a paid mutator transaction binding the contract method 0xef693bed.
//
// Solidity: function exit(address usr, uint256 wad) returns()
func (_GemJoin *GemJoinTransactor) Exit(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _GemJoin.contract.Transact(opts, "exit", usr, wad)
}

// Exit is a paid mutator transaction binding the contract method 0xef693bed.
//
// Solidity: function exit(address usr, uint256 wad) returns()
func (_GemJoin *GemJoinSession) Exit(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _GemJoin.Contract.Exit(&_GemJoin.TransactOpts, usr, wad)
}

// Exit is a paid mutator transaction binding the contract method 0xef693bed.
//
// Solidity: function exit(address usr, uint256 wad) returns()
func (_GemJoin *GemJoinTransactorSession) Exit(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _GemJoin.Contract.Exit(&_GemJoin.TransactOpts, usr, wad)
}

// Join is a paid mutator transaction binding the contract method 0x3b4da69f.
//
// Solidity: function join(address usr, uint256 wad) returns()
func (_GemJoin *GemJoinTransactor) Join(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _GemJoin.contract.Transact(opts, "join", usr, wad)
}

// Join is a paid mutator transaction binding the contract method 0x3b4da69f.
//
// Solidity: function join(address usr, uint256 wad) returns()
func (_GemJoin *GemJoinSession) Join(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _GemJoin.Contract.Join(&_GemJoin.TransactOpts, usr, wad)
}

// Join is a paid mutator transaction binding the contract method 0x3b4da69f.
//
// Solidity: function join(address usr, uint256 wad) returns()
func (_GemJoin *GemJoinTransactorSession) Join(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _GemJoin.Contract.Join(&_GemJoin.TransactOpts, usr, wad)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_GemJoin *GemJoinTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _GemJoin.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_GemJoin *GemJoinSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _GemJoin.Contract.Rely(&_GemJoin.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_GemJoin *GemJoinTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _GemJoin.Contract.Rely(&_GemJoin.TransactOpts, usr)
}
