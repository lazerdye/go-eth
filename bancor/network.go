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

// NetworkABI is the input ABI used to generate the binding from.
const NetworkABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_onlyOwnerCanUpdateRegistry\",\"type\":\"bool\"}],\"name\":\"restrictRegistryUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_register\",\"type\":\"bool\"}],\"name\":\"registerEtherToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getReturnByPath\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"claimAndConvertFor2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"onlyOwnerCanUpdateRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"convert2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAffiliateFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"prevRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"rateByPath\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"etherTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_bancorX\",\"type\":\"address\"},{\"name\":\"_conversionId\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"completeXConversion\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"convertFor2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"claimAndConvertFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restoreRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"convertByPath\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_targetBlockchain\",\"type\":\"bytes32\"},{\"name\":\"_targetAccount\",\"type\":\"bytes32\"},{\"name\":\"_conversionId\",\"type\":\"uint256\"}],\"name\":\"xConvert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"claimAndConvert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"convertFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_targetBlockchain\",\"type\":\"bytes32\"},{\"name\":\"_targetAccount\",\"type\":\"bytes32\"},{\"name\":\"_conversionId\",\"type\":\"uint256\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"xConvert2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sourceToken\",\"type\":\"address\"},{\"name\":\"_targetToken\",\"type\":\"address\"}],\"name\":\"conversionPath\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"claimAndConvert2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"convert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxAffiliateFee\",\"type\":\"uint256\"}],\"name\":\"setMaxAffiliateFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_registry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_smartToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_fromToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_toToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"Conversion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"

// Network is an auto generated Go binding around an Ethereum contract.
type Network struct {
	NetworkCaller     // Read-only binding to the contract
	NetworkTransactor // Write-only binding to the contract
	NetworkFilterer   // Log filterer for contract events
}

// NetworkCaller is an auto generated read-only Go binding around an Ethereum contract.
type NetworkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NetworkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NetworkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NetworkSession struct {
	Contract     *Network          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NetworkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NetworkCallerSession struct {
	Contract *NetworkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// NetworkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NetworkTransactorSession struct {
	Contract     *NetworkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NetworkRaw is an auto generated low-level Go binding around an Ethereum contract.
type NetworkRaw struct {
	Contract *Network // Generic contract binding to access the raw methods on
}

// NetworkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NetworkCallerRaw struct {
	Contract *NetworkCaller // Generic read-only contract binding to access the raw methods on
}

// NetworkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NetworkTransactorRaw struct {
	Contract *NetworkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNetwork creates a new instance of Network, bound to a specific deployed contract.
func NewNetwork(address common.Address, backend bind.ContractBackend) (*Network, error) {
	contract, err := bindNetwork(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Network{NetworkCaller: NetworkCaller{contract: contract}, NetworkTransactor: NetworkTransactor{contract: contract}, NetworkFilterer: NetworkFilterer{contract: contract}}, nil
}

// NewNetworkCaller creates a new read-only instance of Network, bound to a specific deployed contract.
func NewNetworkCaller(address common.Address, caller bind.ContractCaller) (*NetworkCaller, error) {
	contract, err := bindNetwork(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkCaller{contract: contract}, nil
}

// NewNetworkTransactor creates a new write-only instance of Network, bound to a specific deployed contract.
func NewNetworkTransactor(address common.Address, transactor bind.ContractTransactor) (*NetworkTransactor, error) {
	contract, err := bindNetwork(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkTransactor{contract: contract}, nil
}

// NewNetworkFilterer creates a new log filterer instance of Network, bound to a specific deployed contract.
func NewNetworkFilterer(address common.Address, filterer bind.ContractFilterer) (*NetworkFilterer, error) {
	contract, err := bindNetwork(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NetworkFilterer{contract: contract}, nil
}

// bindNetwork binds a generic wrapper to an already deployed contract.
func bindNetwork(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Network *NetworkRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Network.Contract.NetworkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Network *NetworkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Network.Contract.NetworkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Network *NetworkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Network.Contract.NetworkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Network *NetworkCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Network.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Network *NetworkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Network.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Network *NetworkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Network.Contract.contract.Transact(opts, method, params...)
}

// ConversionPath is a free data retrieval call binding the contract method 0xd734fa19.
//
// Solidity: function conversionPath(address _sourceToken, address _targetToken) view returns(address[])
func (_Network *NetworkCaller) ConversionPath(opts *bind.CallOpts, _sourceToken common.Address, _targetToken common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Network.contract.Call(opts, out, "conversionPath", _sourceToken, _targetToken)
	return *ret0, err
}

// ConversionPath is a free data retrieval call binding the contract method 0xd734fa19.
//
// Solidity: function conversionPath(address _sourceToken, address _targetToken) view returns(address[])
func (_Network *NetworkSession) ConversionPath(_sourceToken common.Address, _targetToken common.Address) ([]common.Address, error) {
	return _Network.Contract.ConversionPath(&_Network.CallOpts, _sourceToken, _targetToken)
}

// ConversionPath is a free data retrieval call binding the contract method 0xd734fa19.
//
// Solidity: function conversionPath(address _sourceToken, address _targetToken) view returns(address[])
func (_Network *NetworkCallerSession) ConversionPath(_sourceToken common.Address, _targetToken common.Address) ([]common.Address, error) {
	return _Network.Contract.ConversionPath(&_Network.CallOpts, _sourceToken, _targetToken)
}

// EtherTokens is a free data retrieval call binding the contract method 0x8077ccf7.
//
// Solidity: function etherTokens(address ) view returns(bool)
func (_Network *NetworkCaller) EtherTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Network.contract.Call(opts, out, "etherTokens", arg0)
	return *ret0, err
}

// EtherTokens is a free data retrieval call binding the contract method 0x8077ccf7.
//
// Solidity: function etherTokens(address ) view returns(bool)
func (_Network *NetworkSession) EtherTokens(arg0 common.Address) (bool, error) {
	return _Network.Contract.EtherTokens(&_Network.CallOpts, arg0)
}

// EtherTokens is a free data retrieval call binding the contract method 0x8077ccf7.
//
// Solidity: function etherTokens(address ) view returns(bool)
func (_Network *NetworkCallerSession) EtherTokens(arg0 common.Address) (bool, error) {
	return _Network.Contract.EtherTokens(&_Network.CallOpts, arg0)
}

// GetReturnByPath is a free data retrieval call binding the contract method 0x0c8496cc.
//
// Solidity: function getReturnByPath(address[] _path, uint256 _amount) view returns(uint256, uint256)
func (_Network *NetworkCaller) GetReturnByPath(opts *bind.CallOpts, _path []common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Network.contract.Call(opts, out, "getReturnByPath", _path, _amount)
	return *ret0, *ret1, err
}

// GetReturnByPath is a free data retrieval call binding the contract method 0x0c8496cc.
//
// Solidity: function getReturnByPath(address[] _path, uint256 _amount) view returns(uint256, uint256)
func (_Network *NetworkSession) GetReturnByPath(_path []common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _Network.Contract.GetReturnByPath(&_Network.CallOpts, _path, _amount)
}

// GetReturnByPath is a free data retrieval call binding the contract method 0x0c8496cc.
//
// Solidity: function getReturnByPath(address[] _path, uint256 _amount) view returns(uint256, uint256)
func (_Network *NetworkCallerSession) GetReturnByPath(_path []common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _Network.Contract.GetReturnByPath(&_Network.CallOpts, _path, _amount)
}

// MaxAffiliateFee is a free data retrieval call binding the contract method 0x5d732ff2.
//
// Solidity: function maxAffiliateFee() view returns(uint256)
func (_Network *NetworkCaller) MaxAffiliateFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Network.contract.Call(opts, out, "maxAffiliateFee")
	return *ret0, err
}

// MaxAffiliateFee is a free data retrieval call binding the contract method 0x5d732ff2.
//
// Solidity: function maxAffiliateFee() view returns(uint256)
func (_Network *NetworkSession) MaxAffiliateFee() (*big.Int, error) {
	return _Network.Contract.MaxAffiliateFee(&_Network.CallOpts)
}

// MaxAffiliateFee is a free data retrieval call binding the contract method 0x5d732ff2.
//
// Solidity: function maxAffiliateFee() view returns(uint256)
func (_Network *NetworkCallerSession) MaxAffiliateFee() (*big.Int, error) {
	return _Network.Contract.MaxAffiliateFee(&_Network.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Network *NetworkCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Network.contract.Call(opts, out, "newOwner")
	return *ret0, err
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Network *NetworkSession) NewOwner() (common.Address, error) {
	return _Network.Contract.NewOwner(&_Network.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Network *NetworkCallerSession) NewOwner() (common.Address, error) {
	return _Network.Contract.NewOwner(&_Network.CallOpts)
}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_Network *NetworkCaller) OnlyOwnerCanUpdateRegistry(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Network.contract.Call(opts, out, "onlyOwnerCanUpdateRegistry")
	return *ret0, err
}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_Network *NetworkSession) OnlyOwnerCanUpdateRegistry() (bool, error) {
	return _Network.Contract.OnlyOwnerCanUpdateRegistry(&_Network.CallOpts)
}

// OnlyOwnerCanUpdateRegistry is a free data retrieval call binding the contract method 0x2fe8a6ad.
//
// Solidity: function onlyOwnerCanUpdateRegistry() view returns(bool)
func (_Network *NetworkCallerSession) OnlyOwnerCanUpdateRegistry() (bool, error) {
	return _Network.Contract.OnlyOwnerCanUpdateRegistry(&_Network.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Network *NetworkCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Network.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Network *NetworkSession) Owner() (common.Address, error) {
	return _Network.Contract.Owner(&_Network.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Network *NetworkCallerSession) Owner() (common.Address, error) {
	return _Network.Contract.Owner(&_Network.CallOpts)
}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_Network *NetworkCaller) PrevRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Network.contract.Call(opts, out, "prevRegistry")
	return *ret0, err
}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_Network *NetworkSession) PrevRegistry() (common.Address, error) {
	return _Network.Contract.PrevRegistry(&_Network.CallOpts)
}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_Network *NetworkCallerSession) PrevRegistry() (common.Address, error) {
	return _Network.Contract.PrevRegistry(&_Network.CallOpts)
}

// RateByPath is a free data retrieval call binding the contract method 0x7f9c0ecd.
//
// Solidity: function rateByPath(address[] _path, uint256 _amount) view returns(uint256)
func (_Network *NetworkCaller) RateByPath(opts *bind.CallOpts, _path []common.Address, _amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Network.contract.Call(opts, out, "rateByPath", _path, _amount)
	return *ret0, err
}

// RateByPath is a free data retrieval call binding the contract method 0x7f9c0ecd.
//
// Solidity: function rateByPath(address[] _path, uint256 _amount) view returns(uint256)
func (_Network *NetworkSession) RateByPath(_path []common.Address, _amount *big.Int) (*big.Int, error) {
	return _Network.Contract.RateByPath(&_Network.CallOpts, _path, _amount)
}

// RateByPath is a free data retrieval call binding the contract method 0x7f9c0ecd.
//
// Solidity: function rateByPath(address[] _path, uint256 _amount) view returns(uint256)
func (_Network *NetworkCallerSession) RateByPath(_path []common.Address, _amount *big.Int) (*big.Int, error) {
	return _Network.Contract.RateByPath(&_Network.CallOpts, _path, _amount)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Network *NetworkCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Network.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Network *NetworkSession) Registry() (common.Address, error) {
	return _Network.Contract.Registry(&_Network.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Network *NetworkCallerSession) Registry() (common.Address, error) {
	return _Network.Contract.Registry(&_Network.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Network *NetworkTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Network *NetworkSession) AcceptOwnership() (*types.Transaction, error) {
	return _Network.Contract.AcceptOwnership(&_Network.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Network *NetworkTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Network.Contract.AcceptOwnership(&_Network.TransactOpts)
}

// ClaimAndConvert is a paid mutator transaction binding the contract method 0xc7ba24bc.
//
// Solidity: function claimAndConvert(address[] _path, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_Network *NetworkTransactor) ClaimAndConvert(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "claimAndConvert", _path, _amount, _minReturn)
}

// ClaimAndConvert is a paid mutator transaction binding the contract method 0xc7ba24bc.
//
// Solidity: function claimAndConvert(address[] _path, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_Network *NetworkSession) ClaimAndConvert(_path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _Network.Contract.ClaimAndConvert(&_Network.TransactOpts, _path, _amount, _minReturn)
}

// ClaimAndConvert is a paid mutator transaction binding the contract method 0xc7ba24bc.
//
// Solidity: function claimAndConvert(address[] _path, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_Network *NetworkTransactorSession) ClaimAndConvert(_path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _Network.Contract.ClaimAndConvert(&_Network.TransactOpts, _path, _amount, _minReturn)
}

// ClaimAndConvert2 is a paid mutator transaction binding the contract method 0xe57738e5.
//
// Solidity: function claimAndConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_Network *NetworkTransactor) ClaimAndConvert2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "claimAndConvert2", _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// ClaimAndConvert2 is a paid mutator transaction binding the contract method 0xe57738e5.
//
// Solidity: function claimAndConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_Network *NetworkSession) ClaimAndConvert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.Contract.ClaimAndConvert2(&_Network.TransactOpts, _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// ClaimAndConvert2 is a paid mutator transaction binding the contract method 0xe57738e5.
//
// Solidity: function claimAndConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_Network *NetworkTransactorSession) ClaimAndConvert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.Contract.ClaimAndConvert2(&_Network.TransactOpts, _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// ClaimAndConvertFor is a paid mutator transaction binding the contract method 0xb1e9932b.
//
// Solidity: function claimAndConvertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_Network *NetworkTransactor) ClaimAndConvertFor(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "claimAndConvertFor", _path, _amount, _minReturn, _beneficiary)
}

// ClaimAndConvertFor is a paid mutator transaction binding the contract method 0xb1e9932b.
//
// Solidity: function claimAndConvertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_Network *NetworkSession) ClaimAndConvertFor(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Network.Contract.ClaimAndConvertFor(&_Network.TransactOpts, _path, _amount, _minReturn, _beneficiary)
}

// ClaimAndConvertFor is a paid mutator transaction binding the contract method 0xb1e9932b.
//
// Solidity: function claimAndConvertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_Network *NetworkTransactorSession) ClaimAndConvertFor(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Network.Contract.ClaimAndConvertFor(&_Network.TransactOpts, _path, _amount, _minReturn, _beneficiary)
}

// ClaimAndConvertFor2 is a paid mutator transaction binding the contract method 0x2978c10e.
//
// Solidity: function claimAndConvertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_Network *NetworkTransactor) ClaimAndConvertFor2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "claimAndConvertFor2", _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ClaimAndConvertFor2 is a paid mutator transaction binding the contract method 0x2978c10e.
//
// Solidity: function claimAndConvertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_Network *NetworkSession) ClaimAndConvertFor2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.Contract.ClaimAndConvertFor2(&_Network.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ClaimAndConvertFor2 is a paid mutator transaction binding the contract method 0x2978c10e.
//
// Solidity: function claimAndConvertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_Network *NetworkTransactorSession) ClaimAndConvertFor2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.Contract.ClaimAndConvertFor2(&_Network.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// CompleteXConversion is a paid mutator transaction binding the contract method 0x89f9cc61.
//
// Solidity: function completeXConversion(address[] _path, address _bancorX, uint256 _conversionId, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_Network *NetworkTransactor) CompleteXConversion(opts *bind.TransactOpts, _path []common.Address, _bancorX common.Address, _conversionId *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "completeXConversion", _path, _bancorX, _conversionId, _minReturn, _beneficiary)
}

// CompleteXConversion is a paid mutator transaction binding the contract method 0x89f9cc61.
//
// Solidity: function completeXConversion(address[] _path, address _bancorX, uint256 _conversionId, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_Network *NetworkSession) CompleteXConversion(_path []common.Address, _bancorX common.Address, _conversionId *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Network.Contract.CompleteXConversion(&_Network.TransactOpts, _path, _bancorX, _conversionId, _minReturn, _beneficiary)
}

// CompleteXConversion is a paid mutator transaction binding the contract method 0x89f9cc61.
//
// Solidity: function completeXConversion(address[] _path, address _bancorX, uint256 _conversionId, uint256 _minReturn, address _beneficiary) returns(uint256)
func (_Network *NetworkTransactorSession) CompleteXConversion(_path []common.Address, _bancorX common.Address, _conversionId *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Network.Contract.CompleteXConversion(&_Network.TransactOpts, _path, _bancorX, _conversionId, _minReturn, _beneficiary)
}

// Convert is a paid mutator transaction binding the contract method 0xf3898a97.
//
// Solidity: function convert(address[] _path, uint256 _amount, uint256 _minReturn) payable returns(uint256)
func (_Network *NetworkTransactor) Convert(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "convert", _path, _amount, _minReturn)
}

// Convert is a paid mutator transaction binding the contract method 0xf3898a97.
//
// Solidity: function convert(address[] _path, uint256 _amount, uint256 _minReturn) payable returns(uint256)
func (_Network *NetworkSession) Convert(_path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _Network.Contract.Convert(&_Network.TransactOpts, _path, _amount, _minReturn)
}

// Convert is a paid mutator transaction binding the contract method 0xf3898a97.
//
// Solidity: function convert(address[] _path, uint256 _amount, uint256 _minReturn) payable returns(uint256)
func (_Network *NetworkTransactorSession) Convert(_path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _Network.Contract.Convert(&_Network.TransactOpts, _path, _amount, _minReturn)
}

// Convert2 is a paid mutator transaction binding the contract method 0x569706eb.
//
// Solidity: function convert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Network *NetworkTransactor) Convert2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "convert2", _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// Convert2 is a paid mutator transaction binding the contract method 0x569706eb.
//
// Solidity: function convert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Network *NetworkSession) Convert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.Contract.Convert2(&_Network.TransactOpts, _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// Convert2 is a paid mutator transaction binding the contract method 0x569706eb.
//
// Solidity: function convert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Network *NetworkTransactorSession) Convert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.Contract.Convert2(&_Network.TransactOpts, _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// ConvertByPath is a paid mutator transaction binding the contract method 0xb77d239b.
//
// Solidity: function convertByPath(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Network *NetworkTransactor) ConvertByPath(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "convertByPath", _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ConvertByPath is a paid mutator transaction binding the contract method 0xb77d239b.
//
// Solidity: function convertByPath(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Network *NetworkSession) ConvertByPath(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.Contract.ConvertByPath(&_Network.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ConvertByPath is a paid mutator transaction binding the contract method 0xb77d239b.
//
// Solidity: function convertByPath(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Network *NetworkTransactorSession) ConvertByPath(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.Contract.ConvertByPath(&_Network.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ConvertFor is a paid mutator transaction binding the contract method 0xc98fefed.
//
// Solidity: function convertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) payable returns(uint256)
func (_Network *NetworkTransactor) ConvertFor(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "convertFor", _path, _amount, _minReturn, _beneficiary)
}

// ConvertFor is a paid mutator transaction binding the contract method 0xc98fefed.
//
// Solidity: function convertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) payable returns(uint256)
func (_Network *NetworkSession) ConvertFor(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Network.Contract.ConvertFor(&_Network.TransactOpts, _path, _amount, _minReturn, _beneficiary)
}

// ConvertFor is a paid mutator transaction binding the contract method 0xc98fefed.
//
// Solidity: function convertFor(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary) payable returns(uint256)
func (_Network *NetworkTransactorSession) ConvertFor(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _Network.Contract.ConvertFor(&_Network.TransactOpts, _path, _amount, _minReturn, _beneficiary)
}

// ConvertFor2 is a paid mutator transaction binding the contract method 0xab6214ce.
//
// Solidity: function convertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Network *NetworkTransactor) ConvertFor2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "convertFor2", _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ConvertFor2 is a paid mutator transaction binding the contract method 0xab6214ce.
//
// Solidity: function convertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Network *NetworkSession) ConvertFor2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.Contract.ConvertFor2(&_Network.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// ConvertFor2 is a paid mutator transaction binding the contract method 0xab6214ce.
//
// Solidity: function convertFor2(address[] _path, uint256 _amount, uint256 _minReturn, address _beneficiary, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Network *NetworkTransactorSession) ConvertFor2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _beneficiary common.Address, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.Contract.ConvertFor2(&_Network.TransactOpts, _path, _amount, _minReturn, _beneficiary, _affiliateAccount, _affiliateFee)
}

// RegisterEtherToken is a paid mutator transaction binding the contract method 0x02ef521e.
//
// Solidity: function registerEtherToken(address _token, bool _register) returns()
func (_Network *NetworkTransactor) RegisterEtherToken(opts *bind.TransactOpts, _token common.Address, _register bool) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "registerEtherToken", _token, _register)
}

// RegisterEtherToken is a paid mutator transaction binding the contract method 0x02ef521e.
//
// Solidity: function registerEtherToken(address _token, bool _register) returns()
func (_Network *NetworkSession) RegisterEtherToken(_token common.Address, _register bool) (*types.Transaction, error) {
	return _Network.Contract.RegisterEtherToken(&_Network.TransactOpts, _token, _register)
}

// RegisterEtherToken is a paid mutator transaction binding the contract method 0x02ef521e.
//
// Solidity: function registerEtherToken(address _token, bool _register) returns()
func (_Network *NetworkTransactorSession) RegisterEtherToken(_token common.Address, _register bool) (*types.Transaction, error) {
	return _Network.Contract.RegisterEtherToken(&_Network.TransactOpts, _token, _register)
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_Network *NetworkTransactor) RestoreRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "restoreRegistry")
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_Network *NetworkSession) RestoreRegistry() (*types.Transaction, error) {
	return _Network.Contract.RestoreRegistry(&_Network.TransactOpts)
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_Network *NetworkTransactorSession) RestoreRegistry() (*types.Transaction, error) {
	return _Network.Contract.RestoreRegistry(&_Network.TransactOpts)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_Network *NetworkTransactor) RestrictRegistryUpdate(opts *bind.TransactOpts, _onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "restrictRegistryUpdate", _onlyOwnerCanUpdateRegistry)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_Network *NetworkSession) RestrictRegistryUpdate(_onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _Network.Contract.RestrictRegistryUpdate(&_Network.TransactOpts, _onlyOwnerCanUpdateRegistry)
}

// RestrictRegistryUpdate is a paid mutator transaction binding the contract method 0x024c7ec7.
//
// Solidity: function restrictRegistryUpdate(bool _onlyOwnerCanUpdateRegistry) returns()
func (_Network *NetworkTransactorSession) RestrictRegistryUpdate(_onlyOwnerCanUpdateRegistry bool) (*types.Transaction, error) {
	return _Network.Contract.RestrictRegistryUpdate(&_Network.TransactOpts, _onlyOwnerCanUpdateRegistry)
}

// SetMaxAffiliateFee is a paid mutator transaction binding the contract method 0xf3bc7d2a.
//
// Solidity: function setMaxAffiliateFee(uint256 _maxAffiliateFee) returns()
func (_Network *NetworkTransactor) SetMaxAffiliateFee(opts *bind.TransactOpts, _maxAffiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "setMaxAffiliateFee", _maxAffiliateFee)
}

// SetMaxAffiliateFee is a paid mutator transaction binding the contract method 0xf3bc7d2a.
//
// Solidity: function setMaxAffiliateFee(uint256 _maxAffiliateFee) returns()
func (_Network *NetworkSession) SetMaxAffiliateFee(_maxAffiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.Contract.SetMaxAffiliateFee(&_Network.TransactOpts, _maxAffiliateFee)
}

// SetMaxAffiliateFee is a paid mutator transaction binding the contract method 0xf3bc7d2a.
//
// Solidity: function setMaxAffiliateFee(uint256 _maxAffiliateFee) returns()
func (_Network *NetworkTransactorSession) SetMaxAffiliateFee(_maxAffiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.Contract.SetMaxAffiliateFee(&_Network.TransactOpts, _maxAffiliateFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Network *NetworkTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Network *NetworkSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Network.Contract.TransferOwnership(&_Network.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Network *NetworkTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Network.Contract.TransferOwnership(&_Network.TransactOpts, _newOwner)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_Network *NetworkTransactor) UpdateRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "updateRegistry")
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_Network *NetworkSession) UpdateRegistry() (*types.Transaction, error) {
	return _Network.Contract.UpdateRegistry(&_Network.TransactOpts)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_Network *NetworkTransactorSession) UpdateRegistry() (*types.Transaction, error) {
	return _Network.Contract.UpdateRegistry(&_Network.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_Network *NetworkTransactor) WithdrawTokens(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "withdrawTokens", _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_Network *NetworkSession) WithdrawTokens(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Network.Contract.WithdrawTokens(&_Network.TransactOpts, _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_Network *NetworkTransactorSession) WithdrawTokens(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Network.Contract.WithdrawTokens(&_Network.TransactOpts, _token, _to, _amount)
}

// XConvert is a paid mutator transaction binding the contract method 0xc52173de.
//
// Solidity: function xConvert(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId) payable returns(uint256)
func (_Network *NetworkTransactor) XConvert(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "xConvert", _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId)
}

// XConvert is a paid mutator transaction binding the contract method 0xc52173de.
//
// Solidity: function xConvert(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId) payable returns(uint256)
func (_Network *NetworkSession) XConvert(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int) (*types.Transaction, error) {
	return _Network.Contract.XConvert(&_Network.TransactOpts, _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId)
}

// XConvert is a paid mutator transaction binding the contract method 0xc52173de.
//
// Solidity: function xConvert(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId) payable returns(uint256)
func (_Network *NetworkTransactorSession) XConvert(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int) (*types.Transaction, error) {
	return _Network.Contract.XConvert(&_Network.TransactOpts, _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId)
}

// XConvert2 is a paid mutator transaction binding the contract method 0xcb32564e.
//
// Solidity: function xConvert2(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Network *NetworkTransactor) XConvert2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.contract.Transact(opts, "xConvert2", _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId, _affiliateAccount, _affiliateFee)
}

// XConvert2 is a paid mutator transaction binding the contract method 0xcb32564e.
//
// Solidity: function xConvert2(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Network *NetworkSession) XConvert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.Contract.XConvert2(&_Network.TransactOpts, _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId, _affiliateAccount, _affiliateFee)
}

// XConvert2 is a paid mutator transaction binding the contract method 0xcb32564e.
//
// Solidity: function xConvert2(address[] _path, uint256 _amount, uint256 _minReturn, bytes32 _targetBlockchain, bytes32 _targetAccount, uint256 _conversionId, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_Network *NetworkTransactorSession) XConvert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _targetBlockchain [32]byte, _targetAccount [32]byte, _conversionId *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _Network.Contract.XConvert2(&_Network.TransactOpts, _path, _amount, _minReturn, _targetBlockchain, _targetAccount, _conversionId, _affiliateAccount, _affiliateFee)
}

// NetworkConversionIterator is returned from FilterConversion and is used to iterate over the raw logs and unpacked data for Conversion events raised by the Network contract.
type NetworkConversionIterator struct {
	Event *NetworkConversion // Event containing the contract specifics and raw log

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
func (it *NetworkConversionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkConversion)
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
		it.Event = new(NetworkConversion)
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
func (it *NetworkConversionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkConversionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkConversion represents a Conversion event raised by the Network contract.
type NetworkConversion struct {
	SmartToken common.Address
	FromToken  common.Address
	ToToken    common.Address
	FromAmount *big.Int
	ToAmount   *big.Int
	Trader     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConversion is a free log retrieval operation binding the contract event 0x7154b38b5dd31bb3122436a96d4e09aba5b323ae1fd580025fab55074334c095.
//
// Solidity: event Conversion(address indexed _smartToken, address indexed _fromToken, address indexed _toToken, uint256 _fromAmount, uint256 _toAmount, address _trader)
func (_Network *NetworkFilterer) FilterConversion(opts *bind.FilterOpts, _smartToken []common.Address, _fromToken []common.Address, _toToken []common.Address) (*NetworkConversionIterator, error) {

	var _smartTokenRule []interface{}
	for _, _smartTokenItem := range _smartToken {
		_smartTokenRule = append(_smartTokenRule, _smartTokenItem)
	}
	var _fromTokenRule []interface{}
	for _, _fromTokenItem := range _fromToken {
		_fromTokenRule = append(_fromTokenRule, _fromTokenItem)
	}
	var _toTokenRule []interface{}
	for _, _toTokenItem := range _toToken {
		_toTokenRule = append(_toTokenRule, _toTokenItem)
	}

	logs, sub, err := _Network.contract.FilterLogs(opts, "Conversion", _smartTokenRule, _fromTokenRule, _toTokenRule)
	if err != nil {
		return nil, err
	}
	return &NetworkConversionIterator{contract: _Network.contract, event: "Conversion", logs: logs, sub: sub}, nil
}

// WatchConversion is a free log subscription operation binding the contract event 0x7154b38b5dd31bb3122436a96d4e09aba5b323ae1fd580025fab55074334c095.
//
// Solidity: event Conversion(address indexed _smartToken, address indexed _fromToken, address indexed _toToken, uint256 _fromAmount, uint256 _toAmount, address _trader)
func (_Network *NetworkFilterer) WatchConversion(opts *bind.WatchOpts, sink chan<- *NetworkConversion, _smartToken []common.Address, _fromToken []common.Address, _toToken []common.Address) (event.Subscription, error) {

	var _smartTokenRule []interface{}
	for _, _smartTokenItem := range _smartToken {
		_smartTokenRule = append(_smartTokenRule, _smartTokenItem)
	}
	var _fromTokenRule []interface{}
	for _, _fromTokenItem := range _fromToken {
		_fromTokenRule = append(_fromTokenRule, _fromTokenItem)
	}
	var _toTokenRule []interface{}
	for _, _toTokenItem := range _toToken {
		_toTokenRule = append(_toTokenRule, _toTokenItem)
	}

	logs, sub, err := _Network.contract.WatchLogs(opts, "Conversion", _smartTokenRule, _fromTokenRule, _toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkConversion)
				if err := _Network.contract.UnpackLog(event, "Conversion", log); err != nil {
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

// ParseConversion is a log parse operation binding the contract event 0x7154b38b5dd31bb3122436a96d4e09aba5b323ae1fd580025fab55074334c095.
//
// Solidity: event Conversion(address indexed _smartToken, address indexed _fromToken, address indexed _toToken, uint256 _fromAmount, uint256 _toAmount, address _trader)
func (_Network *NetworkFilterer) ParseConversion(log types.Log) (*NetworkConversion, error) {
	event := new(NetworkConversion)
	if err := _Network.contract.UnpackLog(event, "Conversion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NetworkOwnerUpdateIterator is returned from FilterOwnerUpdate and is used to iterate over the raw logs and unpacked data for OwnerUpdate events raised by the Network contract.
type NetworkOwnerUpdateIterator struct {
	Event *NetworkOwnerUpdate // Event containing the contract specifics and raw log

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
func (it *NetworkOwnerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkOwnerUpdate)
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
		it.Event = new(NetworkOwnerUpdate)
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
func (it *NetworkOwnerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkOwnerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkOwnerUpdate represents a OwnerUpdate event raised by the Network contract.
type NetworkOwnerUpdate struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdate is a free log retrieval operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_Network *NetworkFilterer) FilterOwnerUpdate(opts *bind.FilterOpts, _prevOwner []common.Address, _newOwner []common.Address) (*NetworkOwnerUpdateIterator, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _Network.contract.FilterLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NetworkOwnerUpdateIterator{contract: _Network.contract, event: "OwnerUpdate", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdate is a free log subscription operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_Network *NetworkFilterer) WatchOwnerUpdate(opts *bind.WatchOpts, sink chan<- *NetworkOwnerUpdate, _prevOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _Network.contract.WatchLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkOwnerUpdate)
				if err := _Network.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
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
func (_Network *NetworkFilterer) ParseOwnerUpdate(log types.Log) (*NetworkOwnerUpdate, error) {
	event := new(NetworkOwnerUpdate)
	if err := _Network.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}
