// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package opyn

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

// OptionsFactoryABI is the input ABI used to generate the binding from.
const OptionsFactoryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"string\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_asset\",\"type\":\"string\"},{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"changeAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"optionsExchange\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfOptionsContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_collateralType\",\"type\":\"string\"},{\"name\":\"_collateralExp\",\"type\":\"int32\"},{\"name\":\"_underlyingType\",\"type\":\"string\"},{\"name\":\"_underlyingExp\",\"type\":\"int32\"},{\"name\":\"_oTokenExchangeExp\",\"type\":\"int32\"},{\"name\":\"_strikePrice\",\"type\":\"uint256\"},{\"name\":\"_strikeExp\",\"type\":\"int32\"},{\"name\":\"_strikeAsset\",\"type\":\"string\"},{\"name\":\"_expiry\",\"type\":\"uint256\"},{\"name\":\"_windowSize\",\"type\":\"uint256\"}],\"name\":\"createOptionsContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracleAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_asset\",\"type\":\"string\"},{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"addAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_asset\",\"type\":\"string\"}],\"name\":\"supportsAsset\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_asset\",\"type\":\"string\"}],\"name\":\"deleteAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"optionsContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_optionsExchangeAddr\",\"type\":\"address\"},{\"name\":\"_oracleAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"OptionsContractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"asset\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AssetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"asset\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AssetChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"asset\",\"type\":\"string\"}],\"name\":\"AssetDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OptionsFactory is an auto generated Go binding around an Ethereum contract.
type OptionsFactory struct {
	OptionsFactoryCaller     // Read-only binding to the contract
	OptionsFactoryTransactor // Write-only binding to the contract
	OptionsFactoryFilterer   // Log filterer for contract events
}

// OptionsFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptionsFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptionsFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptionsFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptionsFactorySession struct {
	Contract     *OptionsFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OptionsFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptionsFactoryCallerSession struct {
	Contract *OptionsFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OptionsFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptionsFactoryTransactorSession struct {
	Contract     *OptionsFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OptionsFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptionsFactoryRaw struct {
	Contract *OptionsFactory // Generic contract binding to access the raw methods on
}

// OptionsFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptionsFactoryCallerRaw struct {
	Contract *OptionsFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// OptionsFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptionsFactoryTransactorRaw struct {
	Contract *OptionsFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptionsFactory creates a new instance of OptionsFactory, bound to a specific deployed contract.
func NewOptionsFactory(address common.Address, backend bind.ContractBackend) (*OptionsFactory, error) {
	contract, err := bindOptionsFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptionsFactory{OptionsFactoryCaller: OptionsFactoryCaller{contract: contract}, OptionsFactoryTransactor: OptionsFactoryTransactor{contract: contract}, OptionsFactoryFilterer: OptionsFactoryFilterer{contract: contract}}, nil
}

// NewOptionsFactoryCaller creates a new read-only instance of OptionsFactory, bound to a specific deployed contract.
func NewOptionsFactoryCaller(address common.Address, caller bind.ContractCaller) (*OptionsFactoryCaller, error) {
	contract, err := bindOptionsFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptionsFactoryCaller{contract: contract}, nil
}

// NewOptionsFactoryTransactor creates a new write-only instance of OptionsFactory, bound to a specific deployed contract.
func NewOptionsFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*OptionsFactoryTransactor, error) {
	contract, err := bindOptionsFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptionsFactoryTransactor{contract: contract}, nil
}

// NewOptionsFactoryFilterer creates a new log filterer instance of OptionsFactory, bound to a specific deployed contract.
func NewOptionsFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*OptionsFactoryFilterer, error) {
	contract, err := bindOptionsFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptionsFactoryFilterer{contract: contract}, nil
}

// bindOptionsFactory binds a generic wrapper to an already deployed contract.
func bindOptionsFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OptionsFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptionsFactory *OptionsFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptionsFactory.Contract.OptionsFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptionsFactory *OptionsFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsFactory.Contract.OptionsFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptionsFactory *OptionsFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptionsFactory.Contract.OptionsFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptionsFactory *OptionsFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptionsFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptionsFactory *OptionsFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptionsFactory *OptionsFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptionsFactory.Contract.contract.Transact(opts, method, params...)
}

// GetNumberOfOptionsContracts is a free data retrieval call binding the contract method 0x857c26ff.
//
// Solidity: function getNumberOfOptionsContracts() view returns(uint256)
func (_OptionsFactory *OptionsFactoryCaller) GetNumberOfOptionsContracts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OptionsFactory.contract.Call(opts, &out, "getNumberOfOptionsContracts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfOptionsContracts is a free data retrieval call binding the contract method 0x857c26ff.
//
// Solidity: function getNumberOfOptionsContracts() view returns(uint256)
func (_OptionsFactory *OptionsFactorySession) GetNumberOfOptionsContracts() (*big.Int, error) {
	return _OptionsFactory.Contract.GetNumberOfOptionsContracts(&_OptionsFactory.CallOpts)
}

// GetNumberOfOptionsContracts is a free data retrieval call binding the contract method 0x857c26ff.
//
// Solidity: function getNumberOfOptionsContracts() view returns(uint256)
func (_OptionsFactory *OptionsFactoryCallerSession) GetNumberOfOptionsContracts() (*big.Int, error) {
	return _OptionsFactory.Contract.GetNumberOfOptionsContracts(&_OptionsFactory.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_OptionsFactory *OptionsFactoryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OptionsFactory.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_OptionsFactory *OptionsFactorySession) IsOwner() (bool, error) {
	return _OptionsFactory.Contract.IsOwner(&_OptionsFactory.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_OptionsFactory *OptionsFactoryCallerSession) IsOwner() (bool, error) {
	return _OptionsFactory.Contract.IsOwner(&_OptionsFactory.CallOpts)
}

// OptionsContracts is a free data retrieval call binding the contract method 0xeeafd00b.
//
// Solidity: function optionsContracts(uint256 ) view returns(address)
func (_OptionsFactory *OptionsFactoryCaller) OptionsContracts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OptionsFactory.contract.Call(opts, &out, "optionsContracts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptionsContracts is a free data retrieval call binding the contract method 0xeeafd00b.
//
// Solidity: function optionsContracts(uint256 ) view returns(address)
func (_OptionsFactory *OptionsFactorySession) OptionsContracts(arg0 *big.Int) (common.Address, error) {
	return _OptionsFactory.Contract.OptionsContracts(&_OptionsFactory.CallOpts, arg0)
}

// OptionsContracts is a free data retrieval call binding the contract method 0xeeafd00b.
//
// Solidity: function optionsContracts(uint256 ) view returns(address)
func (_OptionsFactory *OptionsFactoryCallerSession) OptionsContracts(arg0 *big.Int) (common.Address, error) {
	return _OptionsFactory.Contract.OptionsContracts(&_OptionsFactory.CallOpts, arg0)
}

// OptionsExchange is a free data retrieval call binding the contract method 0x3bd33f62.
//
// Solidity: function optionsExchange() view returns(address)
func (_OptionsFactory *OptionsFactoryCaller) OptionsExchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptionsFactory.contract.Call(opts, &out, "optionsExchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptionsExchange is a free data retrieval call binding the contract method 0x3bd33f62.
//
// Solidity: function optionsExchange() view returns(address)
func (_OptionsFactory *OptionsFactorySession) OptionsExchange() (common.Address, error) {
	return _OptionsFactory.Contract.OptionsExchange(&_OptionsFactory.CallOpts)
}

// OptionsExchange is a free data retrieval call binding the contract method 0x3bd33f62.
//
// Solidity: function optionsExchange() view returns(address)
func (_OptionsFactory *OptionsFactoryCallerSession) OptionsExchange() (common.Address, error) {
	return _OptionsFactory.Contract.OptionsExchange(&_OptionsFactory.CallOpts)
}

// OracleAddress is a free data retrieval call binding the contract method 0xa89ae4ba.
//
// Solidity: function oracleAddress() view returns(address)
func (_OptionsFactory *OptionsFactoryCaller) OracleAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptionsFactory.contract.Call(opts, &out, "oracleAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleAddress is a free data retrieval call binding the contract method 0xa89ae4ba.
//
// Solidity: function oracleAddress() view returns(address)
func (_OptionsFactory *OptionsFactorySession) OracleAddress() (common.Address, error) {
	return _OptionsFactory.Contract.OracleAddress(&_OptionsFactory.CallOpts)
}

// OracleAddress is a free data retrieval call binding the contract method 0xa89ae4ba.
//
// Solidity: function oracleAddress() view returns(address)
func (_OptionsFactory *OptionsFactoryCallerSession) OracleAddress() (common.Address, error) {
	return _OptionsFactory.Contract.OracleAddress(&_OptionsFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OptionsFactory *OptionsFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptionsFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OptionsFactory *OptionsFactorySession) Owner() (common.Address, error) {
	return _OptionsFactory.Contract.Owner(&_OptionsFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OptionsFactory *OptionsFactoryCallerSession) Owner() (common.Address, error) {
	return _OptionsFactory.Contract.Owner(&_OptionsFactory.CallOpts)
}

// SupportsAsset is a free data retrieval call binding the contract method 0xd37aefe8.
//
// Solidity: function supportsAsset(string _asset) view returns(bool)
func (_OptionsFactory *OptionsFactoryCaller) SupportsAsset(opts *bind.CallOpts, _asset string) (bool, error) {
	var out []interface{}
	err := _OptionsFactory.contract.Call(opts, &out, "supportsAsset", _asset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsAsset is a free data retrieval call binding the contract method 0xd37aefe8.
//
// Solidity: function supportsAsset(string _asset) view returns(bool)
func (_OptionsFactory *OptionsFactorySession) SupportsAsset(_asset string) (bool, error) {
	return _OptionsFactory.Contract.SupportsAsset(&_OptionsFactory.CallOpts, _asset)
}

// SupportsAsset is a free data retrieval call binding the contract method 0xd37aefe8.
//
// Solidity: function supportsAsset(string _asset) view returns(bool)
func (_OptionsFactory *OptionsFactoryCallerSession) SupportsAsset(_asset string) (bool, error) {
	return _OptionsFactory.Contract.SupportsAsset(&_OptionsFactory.CallOpts, _asset)
}

// Tokens is a free data retrieval call binding the contract method 0x04c2320b.
//
// Solidity: function tokens(string ) view returns(address)
func (_OptionsFactory *OptionsFactoryCaller) Tokens(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _OptionsFactory.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x04c2320b.
//
// Solidity: function tokens(string ) view returns(address)
func (_OptionsFactory *OptionsFactorySession) Tokens(arg0 string) (common.Address, error) {
	return _OptionsFactory.Contract.Tokens(&_OptionsFactory.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x04c2320b.
//
// Solidity: function tokens(string ) view returns(address)
func (_OptionsFactory *OptionsFactoryCallerSession) Tokens(arg0 string) (common.Address, error) {
	return _OptionsFactory.Contract.Tokens(&_OptionsFactory.CallOpts, arg0)
}

// AddAsset is a paid mutator transaction binding the contract method 0xadc02a64.
//
// Solidity: function addAsset(string _asset, address _addr) returns()
func (_OptionsFactory *OptionsFactoryTransactor) AddAsset(opts *bind.TransactOpts, _asset string, _addr common.Address) (*types.Transaction, error) {
	return _OptionsFactory.contract.Transact(opts, "addAsset", _asset, _addr)
}

// AddAsset is a paid mutator transaction binding the contract method 0xadc02a64.
//
// Solidity: function addAsset(string _asset, address _addr) returns()
func (_OptionsFactory *OptionsFactorySession) AddAsset(_asset string, _addr common.Address) (*types.Transaction, error) {
	return _OptionsFactory.Contract.AddAsset(&_OptionsFactory.TransactOpts, _asset, _addr)
}

// AddAsset is a paid mutator transaction binding the contract method 0xadc02a64.
//
// Solidity: function addAsset(string _asset, address _addr) returns()
func (_OptionsFactory *OptionsFactoryTransactorSession) AddAsset(_asset string, _addr common.Address) (*types.Transaction, error) {
	return _OptionsFactory.Contract.AddAsset(&_OptionsFactory.TransactOpts, _asset, _addr)
}

// ChangeAsset is a paid mutator transaction binding the contract method 0x3670ffbe.
//
// Solidity: function changeAsset(string _asset, address _addr) returns()
func (_OptionsFactory *OptionsFactoryTransactor) ChangeAsset(opts *bind.TransactOpts, _asset string, _addr common.Address) (*types.Transaction, error) {
	return _OptionsFactory.contract.Transact(opts, "changeAsset", _asset, _addr)
}

// ChangeAsset is a paid mutator transaction binding the contract method 0x3670ffbe.
//
// Solidity: function changeAsset(string _asset, address _addr) returns()
func (_OptionsFactory *OptionsFactorySession) ChangeAsset(_asset string, _addr common.Address) (*types.Transaction, error) {
	return _OptionsFactory.Contract.ChangeAsset(&_OptionsFactory.TransactOpts, _asset, _addr)
}

// ChangeAsset is a paid mutator transaction binding the contract method 0x3670ffbe.
//
// Solidity: function changeAsset(string _asset, address _addr) returns()
func (_OptionsFactory *OptionsFactoryTransactorSession) ChangeAsset(_asset string, _addr common.Address) (*types.Transaction, error) {
	return _OptionsFactory.Contract.ChangeAsset(&_OptionsFactory.TransactOpts, _asset, _addr)
}

// CreateOptionsContract is a paid mutator transaction binding the contract method 0xa1b72d8a.
//
// Solidity: function createOptionsContract(string _collateralType, int32 _collateralExp, string _underlyingType, int32 _underlyingExp, int32 _oTokenExchangeExp, uint256 _strikePrice, int32 _strikeExp, string _strikeAsset, uint256 _expiry, uint256 _windowSize) returns(address)
func (_OptionsFactory *OptionsFactoryTransactor) CreateOptionsContract(opts *bind.TransactOpts, _collateralType string, _collateralExp int32, _underlyingType string, _underlyingExp int32, _oTokenExchangeExp int32, _strikePrice *big.Int, _strikeExp int32, _strikeAsset string, _expiry *big.Int, _windowSize *big.Int) (*types.Transaction, error) {
	return _OptionsFactory.contract.Transact(opts, "createOptionsContract", _collateralType, _collateralExp, _underlyingType, _underlyingExp, _oTokenExchangeExp, _strikePrice, _strikeExp, _strikeAsset, _expiry, _windowSize)
}

// CreateOptionsContract is a paid mutator transaction binding the contract method 0xa1b72d8a.
//
// Solidity: function createOptionsContract(string _collateralType, int32 _collateralExp, string _underlyingType, int32 _underlyingExp, int32 _oTokenExchangeExp, uint256 _strikePrice, int32 _strikeExp, string _strikeAsset, uint256 _expiry, uint256 _windowSize) returns(address)
func (_OptionsFactory *OptionsFactorySession) CreateOptionsContract(_collateralType string, _collateralExp int32, _underlyingType string, _underlyingExp int32, _oTokenExchangeExp int32, _strikePrice *big.Int, _strikeExp int32, _strikeAsset string, _expiry *big.Int, _windowSize *big.Int) (*types.Transaction, error) {
	return _OptionsFactory.Contract.CreateOptionsContract(&_OptionsFactory.TransactOpts, _collateralType, _collateralExp, _underlyingType, _underlyingExp, _oTokenExchangeExp, _strikePrice, _strikeExp, _strikeAsset, _expiry, _windowSize)
}

// CreateOptionsContract is a paid mutator transaction binding the contract method 0xa1b72d8a.
//
// Solidity: function createOptionsContract(string _collateralType, int32 _collateralExp, string _underlyingType, int32 _underlyingExp, int32 _oTokenExchangeExp, uint256 _strikePrice, int32 _strikeExp, string _strikeAsset, uint256 _expiry, uint256 _windowSize) returns(address)
func (_OptionsFactory *OptionsFactoryTransactorSession) CreateOptionsContract(_collateralType string, _collateralExp int32, _underlyingType string, _underlyingExp int32, _oTokenExchangeExp int32, _strikePrice *big.Int, _strikeExp int32, _strikeAsset string, _expiry *big.Int, _windowSize *big.Int) (*types.Transaction, error) {
	return _OptionsFactory.Contract.CreateOptionsContract(&_OptionsFactory.TransactOpts, _collateralType, _collateralExp, _underlyingType, _underlyingExp, _oTokenExchangeExp, _strikePrice, _strikeExp, _strikeAsset, _expiry, _windowSize)
}

// DeleteAsset is a paid mutator transaction binding the contract method 0xe24aa37c.
//
// Solidity: function deleteAsset(string _asset) returns()
func (_OptionsFactory *OptionsFactoryTransactor) DeleteAsset(opts *bind.TransactOpts, _asset string) (*types.Transaction, error) {
	return _OptionsFactory.contract.Transact(opts, "deleteAsset", _asset)
}

// DeleteAsset is a paid mutator transaction binding the contract method 0xe24aa37c.
//
// Solidity: function deleteAsset(string _asset) returns()
func (_OptionsFactory *OptionsFactorySession) DeleteAsset(_asset string) (*types.Transaction, error) {
	return _OptionsFactory.Contract.DeleteAsset(&_OptionsFactory.TransactOpts, _asset)
}

// DeleteAsset is a paid mutator transaction binding the contract method 0xe24aa37c.
//
// Solidity: function deleteAsset(string _asset) returns()
func (_OptionsFactory *OptionsFactoryTransactorSession) DeleteAsset(_asset string) (*types.Transaction, error) {
	return _OptionsFactory.Contract.DeleteAsset(&_OptionsFactory.TransactOpts, _asset)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OptionsFactory *OptionsFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OptionsFactory *OptionsFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _OptionsFactory.Contract.RenounceOwnership(&_OptionsFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OptionsFactory *OptionsFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OptionsFactory.Contract.RenounceOwnership(&_OptionsFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OptionsFactory *OptionsFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OptionsFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OptionsFactory *OptionsFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OptionsFactory.Contract.TransferOwnership(&_OptionsFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OptionsFactory *OptionsFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OptionsFactory.Contract.TransferOwnership(&_OptionsFactory.TransactOpts, newOwner)
}

// OptionsFactoryAssetAddedIterator is returned from FilterAssetAdded and is used to iterate over the raw logs and unpacked data for AssetAdded events raised by the OptionsFactory contract.
type OptionsFactoryAssetAddedIterator struct {
	Event *OptionsFactoryAssetAdded // Event containing the contract specifics and raw log

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
func (it *OptionsFactoryAssetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptionsFactoryAssetAdded)
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
		it.Event = new(OptionsFactoryAssetAdded)
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
func (it *OptionsFactoryAssetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptionsFactoryAssetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptionsFactoryAssetAdded represents a AssetAdded event raised by the OptionsFactory contract.
type OptionsFactoryAssetAdded struct {
	Asset common.Hash
	Addr  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetAdded is a free log retrieval operation binding the contract event 0x77780d19d6f850ec861d508f6035068d0b68fb9d334887c55393910e9b2951ea.
//
// Solidity: event AssetAdded(string indexed asset, address indexed addr)
func (_OptionsFactory *OptionsFactoryFilterer) FilterAssetAdded(opts *bind.FilterOpts, asset []string, addr []common.Address) (*OptionsFactoryAssetAddedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _OptionsFactory.contract.FilterLogs(opts, "AssetAdded", assetRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &OptionsFactoryAssetAddedIterator{contract: _OptionsFactory.contract, event: "AssetAdded", logs: logs, sub: sub}, nil
}

// WatchAssetAdded is a free log subscription operation binding the contract event 0x77780d19d6f850ec861d508f6035068d0b68fb9d334887c55393910e9b2951ea.
//
// Solidity: event AssetAdded(string indexed asset, address indexed addr)
func (_OptionsFactory *OptionsFactoryFilterer) WatchAssetAdded(opts *bind.WatchOpts, sink chan<- *OptionsFactoryAssetAdded, asset []string, addr []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _OptionsFactory.contract.WatchLogs(opts, "AssetAdded", assetRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptionsFactoryAssetAdded)
				if err := _OptionsFactory.contract.UnpackLog(event, "AssetAdded", log); err != nil {
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

// ParseAssetAdded is a log parse operation binding the contract event 0x77780d19d6f850ec861d508f6035068d0b68fb9d334887c55393910e9b2951ea.
//
// Solidity: event AssetAdded(string indexed asset, address indexed addr)
func (_OptionsFactory *OptionsFactoryFilterer) ParseAssetAdded(log types.Log) (*OptionsFactoryAssetAdded, error) {
	event := new(OptionsFactoryAssetAdded)
	if err := _OptionsFactory.contract.UnpackLog(event, "AssetAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OptionsFactoryAssetChangedIterator is returned from FilterAssetChanged and is used to iterate over the raw logs and unpacked data for AssetChanged events raised by the OptionsFactory contract.
type OptionsFactoryAssetChangedIterator struct {
	Event *OptionsFactoryAssetChanged // Event containing the contract specifics and raw log

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
func (it *OptionsFactoryAssetChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptionsFactoryAssetChanged)
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
		it.Event = new(OptionsFactoryAssetChanged)
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
func (it *OptionsFactoryAssetChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptionsFactoryAssetChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptionsFactoryAssetChanged represents a AssetChanged event raised by the OptionsFactory contract.
type OptionsFactoryAssetChanged struct {
	Asset common.Hash
	Addr  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetChanged is a free log retrieval operation binding the contract event 0x6ce5d5a8378fe2e66e2858afeaf063b9f2d38f7e46f1313d95b102c635e11cf2.
//
// Solidity: event AssetChanged(string indexed asset, address indexed addr)
func (_OptionsFactory *OptionsFactoryFilterer) FilterAssetChanged(opts *bind.FilterOpts, asset []string, addr []common.Address) (*OptionsFactoryAssetChangedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _OptionsFactory.contract.FilterLogs(opts, "AssetChanged", assetRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &OptionsFactoryAssetChangedIterator{contract: _OptionsFactory.contract, event: "AssetChanged", logs: logs, sub: sub}, nil
}

// WatchAssetChanged is a free log subscription operation binding the contract event 0x6ce5d5a8378fe2e66e2858afeaf063b9f2d38f7e46f1313d95b102c635e11cf2.
//
// Solidity: event AssetChanged(string indexed asset, address indexed addr)
func (_OptionsFactory *OptionsFactoryFilterer) WatchAssetChanged(opts *bind.WatchOpts, sink chan<- *OptionsFactoryAssetChanged, asset []string, addr []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _OptionsFactory.contract.WatchLogs(opts, "AssetChanged", assetRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptionsFactoryAssetChanged)
				if err := _OptionsFactory.contract.UnpackLog(event, "AssetChanged", log); err != nil {
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

// ParseAssetChanged is a log parse operation binding the contract event 0x6ce5d5a8378fe2e66e2858afeaf063b9f2d38f7e46f1313d95b102c635e11cf2.
//
// Solidity: event AssetChanged(string indexed asset, address indexed addr)
func (_OptionsFactory *OptionsFactoryFilterer) ParseAssetChanged(log types.Log) (*OptionsFactoryAssetChanged, error) {
	event := new(OptionsFactoryAssetChanged)
	if err := _OptionsFactory.contract.UnpackLog(event, "AssetChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OptionsFactoryAssetDeletedIterator is returned from FilterAssetDeleted and is used to iterate over the raw logs and unpacked data for AssetDeleted events raised by the OptionsFactory contract.
type OptionsFactoryAssetDeletedIterator struct {
	Event *OptionsFactoryAssetDeleted // Event containing the contract specifics and raw log

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
func (it *OptionsFactoryAssetDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptionsFactoryAssetDeleted)
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
		it.Event = new(OptionsFactoryAssetDeleted)
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
func (it *OptionsFactoryAssetDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptionsFactoryAssetDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptionsFactoryAssetDeleted represents a AssetDeleted event raised by the OptionsFactory contract.
type OptionsFactoryAssetDeleted struct {
	Asset common.Hash
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetDeleted is a free log retrieval operation binding the contract event 0xa34d93a1951aca62c967cf6d17f05049fd27b41672a16e301805c4d93995b163.
//
// Solidity: event AssetDeleted(string indexed asset)
func (_OptionsFactory *OptionsFactoryFilterer) FilterAssetDeleted(opts *bind.FilterOpts, asset []string) (*OptionsFactoryAssetDeletedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _OptionsFactory.contract.FilterLogs(opts, "AssetDeleted", assetRule)
	if err != nil {
		return nil, err
	}
	return &OptionsFactoryAssetDeletedIterator{contract: _OptionsFactory.contract, event: "AssetDeleted", logs: logs, sub: sub}, nil
}

// WatchAssetDeleted is a free log subscription operation binding the contract event 0xa34d93a1951aca62c967cf6d17f05049fd27b41672a16e301805c4d93995b163.
//
// Solidity: event AssetDeleted(string indexed asset)
func (_OptionsFactory *OptionsFactoryFilterer) WatchAssetDeleted(opts *bind.WatchOpts, sink chan<- *OptionsFactoryAssetDeleted, asset []string) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _OptionsFactory.contract.WatchLogs(opts, "AssetDeleted", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptionsFactoryAssetDeleted)
				if err := _OptionsFactory.contract.UnpackLog(event, "AssetDeleted", log); err != nil {
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

// ParseAssetDeleted is a log parse operation binding the contract event 0xa34d93a1951aca62c967cf6d17f05049fd27b41672a16e301805c4d93995b163.
//
// Solidity: event AssetDeleted(string indexed asset)
func (_OptionsFactory *OptionsFactoryFilterer) ParseAssetDeleted(log types.Log) (*OptionsFactoryAssetDeleted, error) {
	event := new(OptionsFactoryAssetDeleted)
	if err := _OptionsFactory.contract.UnpackLog(event, "AssetDeleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OptionsFactoryOptionsContractCreatedIterator is returned from FilterOptionsContractCreated and is used to iterate over the raw logs and unpacked data for OptionsContractCreated events raised by the OptionsFactory contract.
type OptionsFactoryOptionsContractCreatedIterator struct {
	Event *OptionsFactoryOptionsContractCreated // Event containing the contract specifics and raw log

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
func (it *OptionsFactoryOptionsContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptionsFactoryOptionsContractCreated)
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
		it.Event = new(OptionsFactoryOptionsContractCreated)
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
func (it *OptionsFactoryOptionsContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptionsFactoryOptionsContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptionsFactoryOptionsContractCreated represents a OptionsContractCreated event raised by the OptionsFactory contract.
type OptionsFactoryOptionsContractCreated struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOptionsContractCreated is a free log retrieval operation binding the contract event 0x4fa983340a1832a8aff635de38049f79c112e361f89a49cf77f7d49501cabca6.
//
// Solidity: event OptionsContractCreated(address addr)
func (_OptionsFactory *OptionsFactoryFilterer) FilterOptionsContractCreated(opts *bind.FilterOpts) (*OptionsFactoryOptionsContractCreatedIterator, error) {

	logs, sub, err := _OptionsFactory.contract.FilterLogs(opts, "OptionsContractCreated")
	if err != nil {
		return nil, err
	}
	return &OptionsFactoryOptionsContractCreatedIterator{contract: _OptionsFactory.contract, event: "OptionsContractCreated", logs: logs, sub: sub}, nil
}

// WatchOptionsContractCreated is a free log subscription operation binding the contract event 0x4fa983340a1832a8aff635de38049f79c112e361f89a49cf77f7d49501cabca6.
//
// Solidity: event OptionsContractCreated(address addr)
func (_OptionsFactory *OptionsFactoryFilterer) WatchOptionsContractCreated(opts *bind.WatchOpts, sink chan<- *OptionsFactoryOptionsContractCreated) (event.Subscription, error) {

	logs, sub, err := _OptionsFactory.contract.WatchLogs(opts, "OptionsContractCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptionsFactoryOptionsContractCreated)
				if err := _OptionsFactory.contract.UnpackLog(event, "OptionsContractCreated", log); err != nil {
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

// ParseOptionsContractCreated is a log parse operation binding the contract event 0x4fa983340a1832a8aff635de38049f79c112e361f89a49cf77f7d49501cabca6.
//
// Solidity: event OptionsContractCreated(address addr)
func (_OptionsFactory *OptionsFactoryFilterer) ParseOptionsContractCreated(log types.Log) (*OptionsFactoryOptionsContractCreated, error) {
	event := new(OptionsFactoryOptionsContractCreated)
	if err := _OptionsFactory.contract.UnpackLog(event, "OptionsContractCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OptionsFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OptionsFactory contract.
type OptionsFactoryOwnershipTransferredIterator struct {
	Event *OptionsFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OptionsFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptionsFactoryOwnershipTransferred)
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
		it.Event = new(OptionsFactoryOwnershipTransferred)
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
func (it *OptionsFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptionsFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptionsFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the OptionsFactory contract.
type OptionsFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OptionsFactory *OptionsFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OptionsFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OptionsFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OptionsFactoryOwnershipTransferredIterator{contract: _OptionsFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OptionsFactory *OptionsFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OptionsFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OptionsFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptionsFactoryOwnershipTransferred)
				if err := _OptionsFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OptionsFactory *OptionsFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*OptionsFactoryOwnershipTransferred, error) {
	event := new(OptionsFactoryOwnershipTransferred)
	if err := _OptionsFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
