// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bancor

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

// ContractRegistryABI is the input ABI used to generate the binding from.
const ContractRegistryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_CONVERTER_UPGRADER\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BNT_TOKEN\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_contractName\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REGISTRY\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractName\",\"type\":\"bytes32\"}],\"name\":\"unregisterAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contractNames\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_CONVERTER_FACTORY\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BNT_CONVERTER\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractName\",\"type\":\"bytes32\"},{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"registerAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"itemCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_FORMULA\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEATURES\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_NETWORK\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_GAS_PRICE_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_contractName\",\"type\":\"bytes32\"}],\"name\":\"addressOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_X\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_contractName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"AddressUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"

// ContractRegistry is an auto generated Go binding around an Ethereum contract.
type ContractRegistry struct {
	ContractRegistryCaller     // Read-only binding to the contract
	ContractRegistryTransactor // Write-only binding to the contract
	ContractRegistryFilterer   // Log filterer for contract events
}

// ContractRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractRegistrySession struct {
	Contract     *ContractRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractRegistryCallerSession struct {
	Contract *ContractRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ContractRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractRegistryTransactorSession struct {
	Contract     *ContractRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRegistryRaw struct {
	Contract *ContractRegistry // Generic contract binding to access the raw methods on
}

// ContractRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractRegistryCallerRaw struct {
	Contract *ContractRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractRegistryTransactorRaw struct {
	Contract *ContractRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractRegistry creates a new instance of ContractRegistry, bound to a specific deployed contract.
func NewContractRegistry(address common.Address, backend bind.ContractBackend) (*ContractRegistry, error) {
	contract, err := bindContractRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractRegistry{ContractRegistryCaller: ContractRegistryCaller{contract: contract}, ContractRegistryTransactor: ContractRegistryTransactor{contract: contract}, ContractRegistryFilterer: ContractRegistryFilterer{contract: contract}}, nil
}

// NewContractRegistryCaller creates a new read-only instance of ContractRegistry, bound to a specific deployed contract.
func NewContractRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractRegistryCaller, error) {
	contract, err := bindContractRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCaller{contract: contract}, nil
}

// NewContractRegistryTransactor creates a new write-only instance of ContractRegistry, bound to a specific deployed contract.
func NewContractRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractRegistryTransactor, error) {
	contract, err := bindContractRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryTransactor{contract: contract}, nil
}

// NewContractRegistryFilterer creates a new log filterer instance of ContractRegistry, bound to a specific deployed contract.
func NewContractRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractRegistryFilterer, error) {
	contract, err := bindContractRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryFilterer{contract: contract}, nil
}

// bindContractRegistry binds a generic wrapper to an already deployed contract.
func bindContractRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRegistry *ContractRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractRegistry.Contract.ContractRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRegistry *ContractRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistry.Contract.ContractRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRegistry *ContractRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRegistry.Contract.ContractRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRegistry *ContractRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRegistry *ContractRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRegistry *ContractRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRegistry.Contract.contract.Transact(opts, method, params...)
}

// BANCORCONVERTERFACTORY is a free data retrieval call binding the contract method 0x5a46f06c.
//
// Solidity: function BANCOR_CONVERTER_FACTORY() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCaller) BANCORCONVERTERFACTORY(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "BANCOR_CONVERTER_FACTORY")
	return *ret0, err
}

// BANCORCONVERTERFACTORY is a free data retrieval call binding the contract method 0x5a46f06c.
//
// Solidity: function BANCOR_CONVERTER_FACTORY() view returns(bytes32)
func (_ContractRegistry *ContractRegistrySession) BANCORCONVERTERFACTORY() ([32]byte, error) {
	return _ContractRegistry.Contract.BANCORCONVERTERFACTORY(&_ContractRegistry.CallOpts)
}

// BANCORCONVERTERFACTORY is a free data retrieval call binding the contract method 0x5a46f06c.
//
// Solidity: function BANCOR_CONVERTER_FACTORY() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCallerSession) BANCORCONVERTERFACTORY() ([32]byte, error) {
	return _ContractRegistry.Contract.BANCORCONVERTERFACTORY(&_ContractRegistry.CallOpts)
}

// BANCORCONVERTERUPGRADER is a free data retrieval call binding the contract method 0x0c87355e.
//
// Solidity: function BANCOR_CONVERTER_UPGRADER() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCaller) BANCORCONVERTERUPGRADER(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "BANCOR_CONVERTER_UPGRADER")
	return *ret0, err
}

// BANCORCONVERTERUPGRADER is a free data retrieval call binding the contract method 0x0c87355e.
//
// Solidity: function BANCOR_CONVERTER_UPGRADER() view returns(bytes32)
func (_ContractRegistry *ContractRegistrySession) BANCORCONVERTERUPGRADER() ([32]byte, error) {
	return _ContractRegistry.Contract.BANCORCONVERTERUPGRADER(&_ContractRegistry.CallOpts)
}

// BANCORCONVERTERUPGRADER is a free data retrieval call binding the contract method 0x0c87355e.
//
// Solidity: function BANCOR_CONVERTER_UPGRADER() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCallerSession) BANCORCONVERTERUPGRADER() ([32]byte, error) {
	return _ContractRegistry.Contract.BANCORCONVERTERUPGRADER(&_ContractRegistry.CallOpts)
}

// BANCORFORMULA is a free data retrieval call binding the contract method 0x6d7bd3fc.
//
// Solidity: function BANCOR_FORMULA() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCaller) BANCORFORMULA(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "BANCOR_FORMULA")
	return *ret0, err
}

// BANCORFORMULA is a free data retrieval call binding the contract method 0x6d7bd3fc.
//
// Solidity: function BANCOR_FORMULA() view returns(bytes32)
func (_ContractRegistry *ContractRegistrySession) BANCORFORMULA() ([32]byte, error) {
	return _ContractRegistry.Contract.BANCORFORMULA(&_ContractRegistry.CallOpts)
}

// BANCORFORMULA is a free data retrieval call binding the contract method 0x6d7bd3fc.
//
// Solidity: function BANCOR_FORMULA() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCallerSession) BANCORFORMULA() ([32]byte, error) {
	return _ContractRegistry.Contract.BANCORFORMULA(&_ContractRegistry.CallOpts)
}

// BANCORGASPRICELIMIT is a free data retrieval call binding the contract method 0x9249993a.
//
// Solidity: function BANCOR_GAS_PRICE_LIMIT() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCaller) BANCORGASPRICELIMIT(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "BANCOR_GAS_PRICE_LIMIT")
	return *ret0, err
}

// BANCORGASPRICELIMIT is a free data retrieval call binding the contract method 0x9249993a.
//
// Solidity: function BANCOR_GAS_PRICE_LIMIT() view returns(bytes32)
func (_ContractRegistry *ContractRegistrySession) BANCORGASPRICELIMIT() ([32]byte, error) {
	return _ContractRegistry.Contract.BANCORGASPRICELIMIT(&_ContractRegistry.CallOpts)
}

// BANCORGASPRICELIMIT is a free data retrieval call binding the contract method 0x9249993a.
//
// Solidity: function BANCOR_GAS_PRICE_LIMIT() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCallerSession) BANCORGASPRICELIMIT() ([32]byte, error) {
	return _ContractRegistry.Contract.BANCORGASPRICELIMIT(&_ContractRegistry.CallOpts)
}

// BANCORNETWORK is a free data retrieval call binding the contract method 0x9232494e.
//
// Solidity: function BANCOR_NETWORK() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCaller) BANCORNETWORK(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "BANCOR_NETWORK")
	return *ret0, err
}

// BANCORNETWORK is a free data retrieval call binding the contract method 0x9232494e.
//
// Solidity: function BANCOR_NETWORK() view returns(bytes32)
func (_ContractRegistry *ContractRegistrySession) BANCORNETWORK() ([32]byte, error) {
	return _ContractRegistry.Contract.BANCORNETWORK(&_ContractRegistry.CallOpts)
}

// BANCORNETWORK is a free data retrieval call binding the contract method 0x9232494e.
//
// Solidity: function BANCOR_NETWORK() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCallerSession) BANCORNETWORK() ([32]byte, error) {
	return _ContractRegistry.Contract.BANCORNETWORK(&_ContractRegistry.CallOpts)
}

// BANCORX is a free data retrieval call binding the contract method 0xc4a8598e.
//
// Solidity: function BANCOR_X() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCaller) BANCORX(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "BANCOR_X")
	return *ret0, err
}

// BANCORX is a free data retrieval call binding the contract method 0xc4a8598e.
//
// Solidity: function BANCOR_X() view returns(bytes32)
func (_ContractRegistry *ContractRegistrySession) BANCORX() ([32]byte, error) {
	return _ContractRegistry.Contract.BANCORX(&_ContractRegistry.CallOpts)
}

// BANCORX is a free data retrieval call binding the contract method 0xc4a8598e.
//
// Solidity: function BANCOR_X() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCallerSession) BANCORX() ([32]byte, error) {
	return _ContractRegistry.Contract.BANCORX(&_ContractRegistry.CallOpts)
}

// BNTCONVERTER is a free data retrieval call binding the contract method 0x62614ae6.
//
// Solidity: function BNT_CONVERTER() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCaller) BNTCONVERTER(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "BNT_CONVERTER")
	return *ret0, err
}

// BNTCONVERTER is a free data retrieval call binding the contract method 0x62614ae6.
//
// Solidity: function BNT_CONVERTER() view returns(bytes32)
func (_ContractRegistry *ContractRegistrySession) BNTCONVERTER() ([32]byte, error) {
	return _ContractRegistry.Contract.BNTCONVERTER(&_ContractRegistry.CallOpts)
}

// BNTCONVERTER is a free data retrieval call binding the contract method 0x62614ae6.
//
// Solidity: function BNT_CONVERTER() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCallerSession) BNTCONVERTER() ([32]byte, error) {
	return _ContractRegistry.Contract.BNTCONVERTER(&_ContractRegistry.CallOpts)
}

// BNTTOKEN is a free data retrieval call binding the contract method 0x1d000b61.
//
// Solidity: function BNT_TOKEN() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCaller) BNTTOKEN(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "BNT_TOKEN")
	return *ret0, err
}

// BNTTOKEN is a free data retrieval call binding the contract method 0x1d000b61.
//
// Solidity: function BNT_TOKEN() view returns(bytes32)
func (_ContractRegistry *ContractRegistrySession) BNTTOKEN() ([32]byte, error) {
	return _ContractRegistry.Contract.BNTTOKEN(&_ContractRegistry.CallOpts)
}

// BNTTOKEN is a free data retrieval call binding the contract method 0x1d000b61.
//
// Solidity: function BNT_TOKEN() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCallerSession) BNTTOKEN() ([32]byte, error) {
	return _ContractRegistry.Contract.BNTTOKEN(&_ContractRegistry.CallOpts)
}

// CONTRACTFEATURES is a free data retrieval call binding the contract method 0x83315b6e.
//
// Solidity: function CONTRACT_FEATURES() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCaller) CONTRACTFEATURES(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "CONTRACT_FEATURES")
	return *ret0, err
}

// CONTRACTFEATURES is a free data retrieval call binding the contract method 0x83315b6e.
//
// Solidity: function CONTRACT_FEATURES() view returns(bytes32)
func (_ContractRegistry *ContractRegistrySession) CONTRACTFEATURES() ([32]byte, error) {
	return _ContractRegistry.Contract.CONTRACTFEATURES(&_ContractRegistry.CallOpts)
}

// CONTRACTFEATURES is a free data retrieval call binding the contract method 0x83315b6e.
//
// Solidity: function CONTRACT_FEATURES() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCallerSession) CONTRACTFEATURES() ([32]byte, error) {
	return _ContractRegistry.Contract.CONTRACTFEATURES(&_ContractRegistry.CallOpts)
}

// CONTRACTREGISTRY is a free data retrieval call binding the contract method 0x25f9bfef.
//
// Solidity: function CONTRACT_REGISTRY() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCaller) CONTRACTREGISTRY(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "CONTRACT_REGISTRY")
	return *ret0, err
}

// CONTRACTREGISTRY is a free data retrieval call binding the contract method 0x25f9bfef.
//
// Solidity: function CONTRACT_REGISTRY() view returns(bytes32)
func (_ContractRegistry *ContractRegistrySession) CONTRACTREGISTRY() ([32]byte, error) {
	return _ContractRegistry.Contract.CONTRACTREGISTRY(&_ContractRegistry.CallOpts)
}

// CONTRACTREGISTRY is a free data retrieval call binding the contract method 0x25f9bfef.
//
// Solidity: function CONTRACT_REGISTRY() view returns(bytes32)
func (_ContractRegistry *ContractRegistryCallerSession) CONTRACTREGISTRY() ([32]byte, error) {
	return _ContractRegistry.Contract.CONTRACTREGISTRY(&_ContractRegistry.CallOpts)
}

// AddressOf is a free data retrieval call binding the contract method 0xbb34534c.
//
// Solidity: function addressOf(bytes32 _contractName) view returns(address)
func (_ContractRegistry *ContractRegistryCaller) AddressOf(opts *bind.CallOpts, _contractName [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "addressOf", _contractName)
	return *ret0, err
}

// AddressOf is a free data retrieval call binding the contract method 0xbb34534c.
//
// Solidity: function addressOf(bytes32 _contractName) view returns(address)
func (_ContractRegistry *ContractRegistrySession) AddressOf(_contractName [32]byte) (common.Address, error) {
	return _ContractRegistry.Contract.AddressOf(&_ContractRegistry.CallOpts, _contractName)
}

// AddressOf is a free data retrieval call binding the contract method 0xbb34534c.
//
// Solidity: function addressOf(bytes32 _contractName) view returns(address)
func (_ContractRegistry *ContractRegistryCallerSession) AddressOf(_contractName [32]byte) (common.Address, error) {
	return _ContractRegistry.Contract.AddressOf(&_ContractRegistry.CallOpts, _contractName)
}

// ContractNames is a free data retrieval call binding the contract method 0x3ca6bb92.
//
// Solidity: function contractNames(uint256 ) view returns(string)
func (_ContractRegistry *ContractRegistryCaller) ContractNames(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "contractNames", arg0)
	return *ret0, err
}

// ContractNames is a free data retrieval call binding the contract method 0x3ca6bb92.
//
// Solidity: function contractNames(uint256 ) view returns(string)
func (_ContractRegistry *ContractRegistrySession) ContractNames(arg0 *big.Int) (string, error) {
	return _ContractRegistry.Contract.ContractNames(&_ContractRegistry.CallOpts, arg0)
}

// ContractNames is a free data retrieval call binding the contract method 0x3ca6bb92.
//
// Solidity: function contractNames(uint256 ) view returns(string)
func (_ContractRegistry *ContractRegistryCallerSession) ContractNames(arg0 *big.Int) (string, error) {
	return _ContractRegistry.Contract.ContractNames(&_ContractRegistry.CallOpts, arg0)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _contractName) view returns(address)
func (_ContractRegistry *ContractRegistryCaller) GetAddress(opts *bind.CallOpts, _contractName [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "getAddress", _contractName)
	return *ret0, err
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _contractName) view returns(address)
func (_ContractRegistry *ContractRegistrySession) GetAddress(_contractName [32]byte) (common.Address, error) {
	return _ContractRegistry.Contract.GetAddress(&_ContractRegistry.CallOpts, _contractName)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _contractName) view returns(address)
func (_ContractRegistry *ContractRegistryCallerSession) GetAddress(_contractName [32]byte) (common.Address, error) {
	return _ContractRegistry.Contract.GetAddress(&_ContractRegistry.CallOpts, _contractName)
}

// ItemCount is a free data retrieval call binding the contract method 0x6bfb0d01.
//
// Solidity: function itemCount() view returns(uint256)
func (_ContractRegistry *ContractRegistryCaller) ItemCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "itemCount")
	return *ret0, err
}

// ItemCount is a free data retrieval call binding the contract method 0x6bfb0d01.
//
// Solidity: function itemCount() view returns(uint256)
func (_ContractRegistry *ContractRegistrySession) ItemCount() (*big.Int, error) {
	return _ContractRegistry.Contract.ItemCount(&_ContractRegistry.CallOpts)
}

// ItemCount is a free data retrieval call binding the contract method 0x6bfb0d01.
//
// Solidity: function itemCount() view returns(uint256)
func (_ContractRegistry *ContractRegistryCallerSession) ItemCount() (*big.Int, error) {
	return _ContractRegistry.Contract.ItemCount(&_ContractRegistry.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ContractRegistry *ContractRegistryCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "newOwner")
	return *ret0, err
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ContractRegistry *ContractRegistrySession) NewOwner() (common.Address, error) {
	return _ContractRegistry.Contract.NewOwner(&_ContractRegistry.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ContractRegistry *ContractRegistryCallerSession) NewOwner() (common.Address, error) {
	return _ContractRegistry.Contract.NewOwner(&_ContractRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRegistry *ContractRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRegistry *ContractRegistrySession) Owner() (common.Address, error) {
	return _ContractRegistry.Contract.Owner(&_ContractRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRegistry *ContractRegistryCallerSession) Owner() (common.Address, error) {
	return _ContractRegistry.Contract.Owner(&_ContractRegistry.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ContractRegistry *ContractRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistry.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ContractRegistry *ContractRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _ContractRegistry.Contract.AcceptOwnership(&_ContractRegistry.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ContractRegistry *ContractRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ContractRegistry.Contract.AcceptOwnership(&_ContractRegistry.TransactOpts)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0x662de379.
//
// Solidity: function registerAddress(bytes32 _contractName, address _contractAddress) returns()
func (_ContractRegistry *ContractRegistryTransactor) RegisterAddress(opts *bind.TransactOpts, _contractName [32]byte, _contractAddress common.Address) (*types.Transaction, error) {
	return _ContractRegistry.contract.Transact(opts, "registerAddress", _contractName, _contractAddress)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0x662de379.
//
// Solidity: function registerAddress(bytes32 _contractName, address _contractAddress) returns()
func (_ContractRegistry *ContractRegistrySession) RegisterAddress(_contractName [32]byte, _contractAddress common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.RegisterAddress(&_ContractRegistry.TransactOpts, _contractName, _contractAddress)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0x662de379.
//
// Solidity: function registerAddress(bytes32 _contractName, address _contractAddress) returns()
func (_ContractRegistry *ContractRegistryTransactorSession) RegisterAddress(_contractName [32]byte, _contractAddress common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.RegisterAddress(&_ContractRegistry.TransactOpts, _contractName, _contractAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ContractRegistry *ContractRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ContractRegistry.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ContractRegistry *ContractRegistrySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.TransferOwnership(&_ContractRegistry.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ContractRegistry *ContractRegistryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.TransferOwnership(&_ContractRegistry.TransactOpts, _newOwner)
}

// UnregisterAddress is a paid mutator transaction binding the contract method 0x2bbd9530.
//
// Solidity: function unregisterAddress(bytes32 _contractName) returns()
func (_ContractRegistry *ContractRegistryTransactor) UnregisterAddress(opts *bind.TransactOpts, _contractName [32]byte) (*types.Transaction, error) {
	return _ContractRegistry.contract.Transact(opts, "unregisterAddress", _contractName)
}

// UnregisterAddress is a paid mutator transaction binding the contract method 0x2bbd9530.
//
// Solidity: function unregisterAddress(bytes32 _contractName) returns()
func (_ContractRegistry *ContractRegistrySession) UnregisterAddress(_contractName [32]byte) (*types.Transaction, error) {
	return _ContractRegistry.Contract.UnregisterAddress(&_ContractRegistry.TransactOpts, _contractName)
}

// UnregisterAddress is a paid mutator transaction binding the contract method 0x2bbd9530.
//
// Solidity: function unregisterAddress(bytes32 _contractName) returns()
func (_ContractRegistry *ContractRegistryTransactorSession) UnregisterAddress(_contractName [32]byte) (*types.Transaction, error) {
	return _ContractRegistry.Contract.UnregisterAddress(&_ContractRegistry.TransactOpts, _contractName)
}

// ContractRegistryAddressUpdateIterator is returned from FilterAddressUpdate and is used to iterate over the raw logs and unpacked data for AddressUpdate events raised by the ContractRegistry contract.
type ContractRegistryAddressUpdateIterator struct {
	Event *ContractRegistryAddressUpdate // Event containing the contract specifics and raw log

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
func (it *ContractRegistryAddressUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryAddressUpdate)
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
		it.Event = new(ContractRegistryAddressUpdate)
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
func (it *ContractRegistryAddressUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryAddressUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryAddressUpdate represents a AddressUpdate event raised by the ContractRegistry contract.
type ContractRegistryAddressUpdate struct {
	ContractName    [32]byte
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddressUpdate is a free log retrieval operation binding the contract event 0xfc08d1253c81bcd5444fc7056ef1f5a5df4c9220b6fd70d7449267f1f0f29918.
//
// Solidity: event AddressUpdate(bytes32 indexed _contractName, address _contractAddress)
func (_ContractRegistry *ContractRegistryFilterer) FilterAddressUpdate(opts *bind.FilterOpts, _contractName [][32]byte) (*ContractRegistryAddressUpdateIterator, error) {

	var _contractNameRule []interface{}
	for _, _contractNameItem := range _contractName {
		_contractNameRule = append(_contractNameRule, _contractNameItem)
	}

	logs, sub, err := _ContractRegistry.contract.FilterLogs(opts, "AddressUpdate", _contractNameRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryAddressUpdateIterator{contract: _ContractRegistry.contract, event: "AddressUpdate", logs: logs, sub: sub}, nil
}

// WatchAddressUpdate is a free log subscription operation binding the contract event 0xfc08d1253c81bcd5444fc7056ef1f5a5df4c9220b6fd70d7449267f1f0f29918.
//
// Solidity: event AddressUpdate(bytes32 indexed _contractName, address _contractAddress)
func (_ContractRegistry *ContractRegistryFilterer) WatchAddressUpdate(opts *bind.WatchOpts, sink chan<- *ContractRegistryAddressUpdate, _contractName [][32]byte) (event.Subscription, error) {

	var _contractNameRule []interface{}
	for _, _contractNameItem := range _contractName {
		_contractNameRule = append(_contractNameRule, _contractNameItem)
	}

	logs, sub, err := _ContractRegistry.contract.WatchLogs(opts, "AddressUpdate", _contractNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryAddressUpdate)
				if err := _ContractRegistry.contract.UnpackLog(event, "AddressUpdate", log); err != nil {
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

// ParseAddressUpdate is a log parse operation binding the contract event 0xfc08d1253c81bcd5444fc7056ef1f5a5df4c9220b6fd70d7449267f1f0f29918.
//
// Solidity: event AddressUpdate(bytes32 indexed _contractName, address _contractAddress)
func (_ContractRegistry *ContractRegistryFilterer) ParseAddressUpdate(log types.Log) (*ContractRegistryAddressUpdate, error) {
	event := new(ContractRegistryAddressUpdate)
	if err := _ContractRegistry.contract.UnpackLog(event, "AddressUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractRegistryOwnerUpdateIterator is returned from FilterOwnerUpdate and is used to iterate over the raw logs and unpacked data for OwnerUpdate events raised by the ContractRegistry contract.
type ContractRegistryOwnerUpdateIterator struct {
	Event *ContractRegistryOwnerUpdate // Event containing the contract specifics and raw log

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
func (it *ContractRegistryOwnerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryOwnerUpdate)
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
		it.Event = new(ContractRegistryOwnerUpdate)
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
func (it *ContractRegistryOwnerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryOwnerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryOwnerUpdate represents a OwnerUpdate event raised by the ContractRegistry contract.
type ContractRegistryOwnerUpdate struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdate is a free log retrieval operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_ContractRegistry *ContractRegistryFilterer) FilterOwnerUpdate(opts *bind.FilterOpts, _prevOwner []common.Address, _newOwner []common.Address) (*ContractRegistryOwnerUpdateIterator, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ContractRegistry.contract.FilterLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryOwnerUpdateIterator{contract: _ContractRegistry.contract, event: "OwnerUpdate", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdate is a free log subscription operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_ContractRegistry *ContractRegistryFilterer) WatchOwnerUpdate(opts *bind.WatchOpts, sink chan<- *ContractRegistryOwnerUpdate, _prevOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ContractRegistry.contract.WatchLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryOwnerUpdate)
				if err := _ContractRegistry.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
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

// ParseOwnerUpdate is a log parse operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_ContractRegistry *ContractRegistryFilterer) ParseOwnerUpdate(log types.Log) (*ContractRegistryOwnerUpdate, error) {
	event := new(ContractRegistryOwnerUpdate)
	if err := _ContractRegistry.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}
