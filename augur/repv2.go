// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package augur

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

// Repv2ABI is the input ABI used to generate the binding from.
const Repv2ABI = "[{\"inputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"_augur\",\"type\":\"address\"},{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"contractIUniverse\",\"name\":\"_parentUniverse\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"augur\",\"outputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountToBurn\",\"type\":\"uint256\"}],\"name\":\"burnForMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLegacyRepToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalMigrated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalTheoreticalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUniverse\",\"outputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"legacyRepToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"migrateFromLegacyReputationToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"migrateIn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"migrateOutByPayout\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountMigrated\",\"type\":\"uint256\"}],\"name\":\"mintForReportingParticipant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountToMint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"mintForWarpSync\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"parentUniverse\",\"outputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedDisputeWindowTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedMarketTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedReportingParticipantTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_attotokens\",\"type\":\"uint256\"}],\"name\":\"trustedUniverseTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"warpSync\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Repv2 is an auto generated Go binding around an Ethereum contract.
type Repv2 struct {
	Repv2Caller     // Read-only binding to the contract
	Repv2Transactor // Write-only binding to the contract
	Repv2Filterer   // Log filterer for contract events
}

// Repv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Repv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Repv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Repv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Repv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Repv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Repv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Repv2Session struct {
	Contract     *Repv2            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Repv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Repv2CallerSession struct {
	Contract *Repv2Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Repv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Repv2TransactorSession struct {
	Contract     *Repv2Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Repv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Repv2Raw struct {
	Contract *Repv2 // Generic contract binding to access the raw methods on
}

// Repv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Repv2CallerRaw struct {
	Contract *Repv2Caller // Generic read-only contract binding to access the raw methods on
}

// Repv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Repv2TransactorRaw struct {
	Contract *Repv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRepv2 creates a new instance of Repv2, bound to a specific deployed contract.
func NewRepv2(address common.Address, backend bind.ContractBackend) (*Repv2, error) {
	contract, err := bindRepv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Repv2{Repv2Caller: Repv2Caller{contract: contract}, Repv2Transactor: Repv2Transactor{contract: contract}, Repv2Filterer: Repv2Filterer{contract: contract}}, nil
}

// NewRepv2Caller creates a new read-only instance of Repv2, bound to a specific deployed contract.
func NewRepv2Caller(address common.Address, caller bind.ContractCaller) (*Repv2Caller, error) {
	contract, err := bindRepv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Repv2Caller{contract: contract}, nil
}

// NewRepv2Transactor creates a new write-only instance of Repv2, bound to a specific deployed contract.
func NewRepv2Transactor(address common.Address, transactor bind.ContractTransactor) (*Repv2Transactor, error) {
	contract, err := bindRepv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Repv2Transactor{contract: contract}, nil
}

// NewRepv2Filterer creates a new log filterer instance of Repv2, bound to a specific deployed contract.
func NewRepv2Filterer(address common.Address, filterer bind.ContractFilterer) (*Repv2Filterer, error) {
	contract, err := bindRepv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Repv2Filterer{contract: contract}, nil
}

// bindRepv2 binds a generic wrapper to an already deployed contract.
func bindRepv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Repv2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Repv2 *Repv2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Repv2.Contract.Repv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Repv2 *Repv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Repv2.Contract.Repv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Repv2 *Repv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Repv2.Contract.Repv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Repv2 *Repv2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Repv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Repv2 *Repv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Repv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Repv2 *Repv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Repv2.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Repv2 *Repv2Caller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Repv2.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Repv2 *Repv2Session) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Repv2.Contract.Allowance(&_Repv2.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Repv2 *Repv2CallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Repv2.Contract.Allowance(&_Repv2.CallOpts, _owner, _spender)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_Repv2 *Repv2Caller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Repv2.contract.Call(opts, out, "allowances", arg0, arg1)
	return *ret0, err
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_Repv2 *Repv2Session) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Repv2.Contract.Allowances(&_Repv2.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_Repv2 *Repv2CallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Repv2.Contract.Allowances(&_Repv2.CallOpts, arg0, arg1)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_Repv2 *Repv2Caller) Augur(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Repv2.contract.Call(opts, out, "augur")
	return *ret0, err
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_Repv2 *Repv2Session) Augur() (common.Address, error) {
	return _Repv2.Contract.Augur(&_Repv2.CallOpts)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_Repv2 *Repv2CallerSession) Augur() (common.Address, error) {
	return _Repv2.Contract.Augur(&_Repv2.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Repv2 *Repv2Caller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Repv2.contract.Call(opts, out, "balanceOf", _account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Repv2 *Repv2Session) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Repv2.Contract.BalanceOf(&_Repv2.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Repv2 *Repv2CallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Repv2.Contract.BalanceOf(&_Repv2.CallOpts, _account)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Repv2 *Repv2Caller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Repv2.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Repv2 *Repv2Session) Balances(arg0 common.Address) (*big.Int, error) {
	return _Repv2.Contract.Balances(&_Repv2.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Repv2 *Repv2CallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Repv2.Contract.Balances(&_Repv2.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Repv2 *Repv2Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Repv2.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Repv2 *Repv2Session) Decimals() (uint8, error) {
	return _Repv2.Contract.Decimals(&_Repv2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Repv2 *Repv2CallerSession) Decimals() (uint8, error) {
	return _Repv2.Contract.Decimals(&_Repv2.CallOpts)
}

// GetLegacyRepToken is a free data retrieval call binding the contract method 0x77469275.
//
// Solidity: function getLegacyRepToken() view returns(address)
func (_Repv2 *Repv2Caller) GetLegacyRepToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Repv2.contract.Call(opts, out, "getLegacyRepToken")
	return *ret0, err
}

// GetLegacyRepToken is a free data retrieval call binding the contract method 0x77469275.
//
// Solidity: function getLegacyRepToken() view returns(address)
func (_Repv2 *Repv2Session) GetLegacyRepToken() (common.Address, error) {
	return _Repv2.Contract.GetLegacyRepToken(&_Repv2.CallOpts)
}

// GetLegacyRepToken is a free data retrieval call binding the contract method 0x77469275.
//
// Solidity: function getLegacyRepToken() view returns(address)
func (_Repv2 *Repv2CallerSession) GetLegacyRepToken() (common.Address, error) {
	return _Repv2.Contract.GetLegacyRepToken(&_Repv2.CallOpts)
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() view returns(uint256)
func (_Repv2 *Repv2Caller) GetTotalMigrated(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Repv2.contract.Call(opts, out, "getTotalMigrated")
	return *ret0, err
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() view returns(uint256)
func (_Repv2 *Repv2Session) GetTotalMigrated() (*big.Int, error) {
	return _Repv2.Contract.GetTotalMigrated(&_Repv2.CallOpts)
}

// GetTotalMigrated is a free data retrieval call binding the contract method 0x91d76bbb.
//
// Solidity: function getTotalMigrated() view returns(uint256)
func (_Repv2 *Repv2CallerSession) GetTotalMigrated() (*big.Int, error) {
	return _Repv2.Contract.GetTotalMigrated(&_Repv2.CallOpts)
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() view returns(uint256)
func (_Repv2 *Repv2Caller) GetTotalTheoreticalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Repv2.contract.Call(opts, out, "getTotalTheoreticalSupply")
	return *ret0, err
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() view returns(uint256)
func (_Repv2 *Repv2Session) GetTotalTheoreticalSupply() (*big.Int, error) {
	return _Repv2.Contract.GetTotalTheoreticalSupply(&_Repv2.CallOpts)
}

// GetTotalTheoreticalSupply is a free data retrieval call binding the contract method 0x238d3590.
//
// Solidity: function getTotalTheoreticalSupply() view returns(uint256)
func (_Repv2 *Repv2CallerSession) GetTotalTheoreticalSupply() (*big.Int, error) {
	return _Repv2.Contract.GetTotalTheoreticalSupply(&_Repv2.CallOpts)
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() view returns(address)
func (_Repv2 *Repv2Caller) GetUniverse(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Repv2.contract.Call(opts, out, "getUniverse")
	return *ret0, err
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() view returns(address)
func (_Repv2 *Repv2Session) GetUniverse() (common.Address, error) {
	return _Repv2.Contract.GetUniverse(&_Repv2.CallOpts)
}

// GetUniverse is a free data retrieval call binding the contract method 0x870c426d.
//
// Solidity: function getUniverse() view returns(address)
func (_Repv2 *Repv2CallerSession) GetUniverse() (common.Address, error) {
	return _Repv2.Contract.GetUniverse(&_Repv2.CallOpts)
}

// LegacyRepToken is a free data retrieval call binding the contract method 0xd5466777.
//
// Solidity: function legacyRepToken() view returns(address)
func (_Repv2 *Repv2Caller) LegacyRepToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Repv2.contract.Call(opts, out, "legacyRepToken")
	return *ret0, err
}

// LegacyRepToken is a free data retrieval call binding the contract method 0xd5466777.
//
// Solidity: function legacyRepToken() view returns(address)
func (_Repv2 *Repv2Session) LegacyRepToken() (common.Address, error) {
	return _Repv2.Contract.LegacyRepToken(&_Repv2.CallOpts)
}

// LegacyRepToken is a free data retrieval call binding the contract method 0xd5466777.
//
// Solidity: function legacyRepToken() view returns(address)
func (_Repv2 *Repv2CallerSession) LegacyRepToken() (common.Address, error) {
	return _Repv2.Contract.LegacyRepToken(&_Repv2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Repv2 *Repv2Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Repv2.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Repv2 *Repv2Session) Name() (string, error) {
	return _Repv2.Contract.Name(&_Repv2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Repv2 *Repv2CallerSession) Name() (string, error) {
	return _Repv2.Contract.Name(&_Repv2.CallOpts)
}

// ParentUniverse is a free data retrieval call binding the contract method 0x183636c5.
//
// Solidity: function parentUniverse() view returns(address)
func (_Repv2 *Repv2Caller) ParentUniverse(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Repv2.contract.Call(opts, out, "parentUniverse")
	return *ret0, err
}

// ParentUniverse is a free data retrieval call binding the contract method 0x183636c5.
//
// Solidity: function parentUniverse() view returns(address)
func (_Repv2 *Repv2Session) ParentUniverse() (common.Address, error) {
	return _Repv2.Contract.ParentUniverse(&_Repv2.CallOpts)
}

// ParentUniverse is a free data retrieval call binding the contract method 0x183636c5.
//
// Solidity: function parentUniverse() view returns(address)
func (_Repv2 *Repv2CallerSession) ParentUniverse() (common.Address, error) {
	return _Repv2.Contract.ParentUniverse(&_Repv2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Repv2 *Repv2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Repv2.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Repv2 *Repv2Session) Symbol() (string, error) {
	return _Repv2.Contract.Symbol(&_Repv2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Repv2 *Repv2CallerSession) Symbol() (string, error) {
	return _Repv2.Contract.Symbol(&_Repv2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Repv2 *Repv2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Repv2.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Repv2 *Repv2Session) TotalSupply() (*big.Int, error) {
	return _Repv2.Contract.TotalSupply(&_Repv2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Repv2 *Repv2CallerSession) TotalSupply() (*big.Int, error) {
	return _Repv2.Contract.TotalSupply(&_Repv2.CallOpts)
}

// WarpSync is a free data retrieval call binding the contract method 0x5bd5ea71.
//
// Solidity: function warpSync() view returns(address)
func (_Repv2 *Repv2Caller) WarpSync(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Repv2.contract.Call(opts, out, "warpSync")
	return *ret0, err
}

// WarpSync is a free data retrieval call binding the contract method 0x5bd5ea71.
//
// Solidity: function warpSync() view returns(address)
func (_Repv2 *Repv2Session) WarpSync() (common.Address, error) {
	return _Repv2.Contract.WarpSync(&_Repv2.CallOpts)
}

// WarpSync is a free data retrieval call binding the contract method 0x5bd5ea71.
//
// Solidity: function warpSync() view returns(address)
func (_Repv2 *Repv2CallerSession) WarpSync() (common.Address, error) {
	return _Repv2.Contract.WarpSync(&_Repv2.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Repv2 *Repv2Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Repv2.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Repv2 *Repv2Session) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.Approve(&_Repv2.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Repv2 *Repv2TransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.Approve(&_Repv2.TransactOpts, _spender, _amount)
}

// BurnForMarket is a paid mutator transaction binding the contract method 0xae3816ee.
//
// Solidity: function burnForMarket(uint256 _amountToBurn) returns(bool)
func (_Repv2 *Repv2Transactor) BurnForMarket(opts *bind.TransactOpts, _amountToBurn *big.Int) (*types.Transaction, error) {
	return _Repv2.contract.Transact(opts, "burnForMarket", _amountToBurn)
}

// BurnForMarket is a paid mutator transaction binding the contract method 0xae3816ee.
//
// Solidity: function burnForMarket(uint256 _amountToBurn) returns(bool)
func (_Repv2 *Repv2Session) BurnForMarket(_amountToBurn *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.BurnForMarket(&_Repv2.TransactOpts, _amountToBurn)
}

// BurnForMarket is a paid mutator transaction binding the contract method 0xae3816ee.
//
// Solidity: function burnForMarket(uint256 _amountToBurn) returns(bool)
func (_Repv2 *Repv2TransactorSession) BurnForMarket(_amountToBurn *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.BurnForMarket(&_Repv2.TransactOpts, _amountToBurn)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_Repv2 *Repv2Transactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Repv2.contract.Transact(opts, "decreaseAllowance", _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_Repv2 *Repv2Session) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.DecreaseAllowance(&_Repv2.TransactOpts, _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_Repv2 *Repv2TransactorSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.DecreaseAllowance(&_Repv2.TransactOpts, _spender, _subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_Repv2 *Repv2Transactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Repv2.contract.Transact(opts, "increaseAllowance", _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_Repv2 *Repv2Session) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.IncreaseAllowance(&_Repv2.TransactOpts, _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_Repv2 *Repv2TransactorSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.IncreaseAllowance(&_Repv2.TransactOpts, _spender, _addedValue)
}

// MigrateFromLegacyReputationToken is a paid mutator transaction binding the contract method 0x75d9aa1a.
//
// Solidity: function migrateFromLegacyReputationToken() returns(bool)
func (_Repv2 *Repv2Transactor) MigrateFromLegacyReputationToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Repv2.contract.Transact(opts, "migrateFromLegacyReputationToken")
}

// MigrateFromLegacyReputationToken is a paid mutator transaction binding the contract method 0x75d9aa1a.
//
// Solidity: function migrateFromLegacyReputationToken() returns(bool)
func (_Repv2 *Repv2Session) MigrateFromLegacyReputationToken() (*types.Transaction, error) {
	return _Repv2.Contract.MigrateFromLegacyReputationToken(&_Repv2.TransactOpts)
}

// MigrateFromLegacyReputationToken is a paid mutator transaction binding the contract method 0x75d9aa1a.
//
// Solidity: function migrateFromLegacyReputationToken() returns(bool)
func (_Repv2 *Repv2TransactorSession) MigrateFromLegacyReputationToken() (*types.Transaction, error) {
	return _Repv2.Contract.MigrateFromLegacyReputationToken(&_Repv2.TransactOpts)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(address _reporter, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2Transactor) MigrateIn(opts *bind.TransactOpts, _reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.contract.Transact(opts, "migrateIn", _reporter, _attotokens)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(address _reporter, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2Session) MigrateIn(_reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.MigrateIn(&_Repv2.TransactOpts, _reporter, _attotokens)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(address _reporter, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2TransactorSession) MigrateIn(_reporter common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.MigrateIn(&_Repv2.TransactOpts, _reporter, _attotokens)
}

// MigrateOutByPayout is a paid mutator transaction binding the contract method 0x5b66876a.
//
// Solidity: function migrateOutByPayout(uint256[] _payoutNumerators, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2Transactor) MigrateOutByPayout(opts *bind.TransactOpts, _payoutNumerators []*big.Int, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.contract.Transact(opts, "migrateOutByPayout", _payoutNumerators, _attotokens)
}

// MigrateOutByPayout is a paid mutator transaction binding the contract method 0x5b66876a.
//
// Solidity: function migrateOutByPayout(uint256[] _payoutNumerators, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2Session) MigrateOutByPayout(_payoutNumerators []*big.Int, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.MigrateOutByPayout(&_Repv2.TransactOpts, _payoutNumerators, _attotokens)
}

// MigrateOutByPayout is a paid mutator transaction binding the contract method 0x5b66876a.
//
// Solidity: function migrateOutByPayout(uint256[] _payoutNumerators, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2TransactorSession) MigrateOutByPayout(_payoutNumerators []*big.Int, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.MigrateOutByPayout(&_Repv2.TransactOpts, _payoutNumerators, _attotokens)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(uint256 _amountMigrated) returns(bool)
func (_Repv2 *Repv2Transactor) MintForReportingParticipant(opts *bind.TransactOpts, _amountMigrated *big.Int) (*types.Transaction, error) {
	return _Repv2.contract.Transact(opts, "mintForReportingParticipant", _amountMigrated)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(uint256 _amountMigrated) returns(bool)
func (_Repv2 *Repv2Session) MintForReportingParticipant(_amountMigrated *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.MintForReportingParticipant(&_Repv2.TransactOpts, _amountMigrated)
}

// MintForReportingParticipant is a paid mutator transaction binding the contract method 0xdb054134.
//
// Solidity: function mintForReportingParticipant(uint256 _amountMigrated) returns(bool)
func (_Repv2 *Repv2TransactorSession) MintForReportingParticipant(_amountMigrated *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.MintForReportingParticipant(&_Repv2.TransactOpts, _amountMigrated)
}

// MintForWarpSync is a paid mutator transaction binding the contract method 0xe1897b59.
//
// Solidity: function mintForWarpSync(uint256 _amountToMint, address _target) returns(bool)
func (_Repv2 *Repv2Transactor) MintForWarpSync(opts *bind.TransactOpts, _amountToMint *big.Int, _target common.Address) (*types.Transaction, error) {
	return _Repv2.contract.Transact(opts, "mintForWarpSync", _amountToMint, _target)
}

// MintForWarpSync is a paid mutator transaction binding the contract method 0xe1897b59.
//
// Solidity: function mintForWarpSync(uint256 _amountToMint, address _target) returns(bool)
func (_Repv2 *Repv2Session) MintForWarpSync(_amountToMint *big.Int, _target common.Address) (*types.Transaction, error) {
	return _Repv2.Contract.MintForWarpSync(&_Repv2.TransactOpts, _amountToMint, _target)
}

// MintForWarpSync is a paid mutator transaction binding the contract method 0xe1897b59.
//
// Solidity: function mintForWarpSync(uint256 _amountToMint, address _target) returns(bool)
func (_Repv2 *Repv2TransactorSession) MintForWarpSync(_amountToMint *big.Int, _target common.Address) (*types.Transaction, error) {
	return _Repv2.Contract.MintForWarpSync(&_Repv2.TransactOpts, _amountToMint, _target)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Repv2 *Repv2Transactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Repv2.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Repv2 *Repv2Session) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.Transfer(&_Repv2.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Repv2 *Repv2TransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.Transfer(&_Repv2.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_Repv2 *Repv2Transactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Repv2.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_Repv2 *Repv2Session) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.TransferFrom(&_Repv2.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_Repv2 *Repv2TransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.TransferFrom(&_Repv2.TransactOpts, _sender, _recipient, _amount)
}

// TrustedDisputeWindowTransfer is a paid mutator transaction binding the contract method 0x721bb170.
//
// Solidity: function trustedDisputeWindowTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2Transactor) TrustedDisputeWindowTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.contract.Transact(opts, "trustedDisputeWindowTransfer", _source, _destination, _attotokens)
}

// TrustedDisputeWindowTransfer is a paid mutator transaction binding the contract method 0x721bb170.
//
// Solidity: function trustedDisputeWindowTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2Session) TrustedDisputeWindowTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.TrustedDisputeWindowTransfer(&_Repv2.TransactOpts, _source, _destination, _attotokens)
}

// TrustedDisputeWindowTransfer is a paid mutator transaction binding the contract method 0x721bb170.
//
// Solidity: function trustedDisputeWindowTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2TransactorSession) TrustedDisputeWindowTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.TrustedDisputeWindowTransfer(&_Repv2.TransactOpts, _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2Transactor) TrustedMarketTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.contract.Transact(opts, "trustedMarketTransfer", _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2Session) TrustedMarketTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.TrustedMarketTransfer(&_Repv2.TransactOpts, _source, _destination, _attotokens)
}

// TrustedMarketTransfer is a paid mutator transaction binding the contract method 0xf22b258a.
//
// Solidity: function trustedMarketTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2TransactorSession) TrustedMarketTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.TrustedMarketTransfer(&_Repv2.TransactOpts, _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2Transactor) TrustedReportingParticipantTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.contract.Transact(opts, "trustedReportingParticipantTransfer", _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2Session) TrustedReportingParticipantTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.TrustedReportingParticipantTransfer(&_Repv2.TransactOpts, _source, _destination, _attotokens)
}

// TrustedReportingParticipantTransfer is a paid mutator transaction binding the contract method 0xb873e9a7.
//
// Solidity: function trustedReportingParticipantTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2TransactorSession) TrustedReportingParticipantTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.TrustedReportingParticipantTransfer(&_Repv2.TransactOpts, _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2Transactor) TrustedUniverseTransfer(opts *bind.TransactOpts, _source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.contract.Transact(opts, "trustedUniverseTransfer", _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2Session) TrustedUniverseTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.TrustedUniverseTransfer(&_Repv2.TransactOpts, _source, _destination, _attotokens)
}

// TrustedUniverseTransfer is a paid mutator transaction binding the contract method 0xfe98184d.
//
// Solidity: function trustedUniverseTransfer(address _source, address _destination, uint256 _attotokens) returns(bool)
func (_Repv2 *Repv2TransactorSession) TrustedUniverseTransfer(_source common.Address, _destination common.Address, _attotokens *big.Int) (*types.Transaction, error) {
	return _Repv2.Contract.TrustedUniverseTransfer(&_Repv2.TransactOpts, _source, _destination, _attotokens)
}

// Repv2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Repv2 contract.
type Repv2ApprovalIterator struct {
	Event *Repv2Approval // Event containing the contract specifics and raw log

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
func (it *Repv2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Repv2Approval)
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
		it.Event = new(Repv2Approval)
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
func (it *Repv2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Repv2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Repv2Approval represents a Approval event raised by the Repv2 contract.
type Repv2Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Repv2 *Repv2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Repv2ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Repv2.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Repv2ApprovalIterator{contract: _Repv2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Repv2 *Repv2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Repv2Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Repv2.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Repv2Approval)
				if err := _Repv2.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Repv2 *Repv2Filterer) ParseApproval(log types.Log) (*Repv2Approval, error) {
	event := new(Repv2Approval)
	if err := _Repv2.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Repv2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Repv2 contract.
type Repv2TransferIterator struct {
	Event *Repv2Transfer // Event containing the contract specifics and raw log

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
func (it *Repv2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Repv2Transfer)
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
		it.Event = new(Repv2Transfer)
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
func (it *Repv2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Repv2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Repv2Transfer represents a Transfer event raised by the Repv2 contract.
type Repv2Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Repv2 *Repv2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Repv2TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Repv2.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Repv2TransferIterator{contract: _Repv2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Repv2 *Repv2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Repv2Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Repv2.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Repv2Transfer)
				if err := _Repv2.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Repv2 *Repv2Filterer) ParseTransfer(log types.Log) (*Repv2Transfer, error) {
	event := new(Repv2Transfer)
	if err := _Repv2.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
