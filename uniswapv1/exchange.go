// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv1

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ExchangeABI is the input ABI used to generate the binding from.
const ExchangeABI = "[{\"name\":\"TokenPurchase\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"eth_sold\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"tokens_bought\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"EthPurchase\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"tokens_sold\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"eth_bought\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"eth_amount\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"token_amount\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"eth_amount\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"token_amount\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"_from\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"setup\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":175875},{\"name\":\"addLiquidity\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"min_liquidity\"},{\"type\":\"uint256\",\"name\":\"max_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":82616},{\"name\":\"removeLiquidity\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount\"},{\"type\":\"uint256\",\"name\":\"min_eth\"},{\"type\":\"uint256\",\"name\":\"min_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":116814},{\"name\":\"__default__\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":true,\"type\":\"function\"},{\"name\":\"ethToTokenSwapInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"min_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":12757},{\"name\":\"ethToTokenTransferInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"min_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":12965},{\"name\":\"ethToTokenSwapOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":50463},{\"name\":\"ethToTokenTransferOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":50671},{\"name\":\"tokenToEthSwapInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_eth\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":47503},{\"name\":\"tokenToEthTransferInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_eth\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":47712},{\"name\":\"tokenToEthSwapOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"eth_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":50175},{\"name\":\"tokenToEthTransferOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"eth_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":50384},{\"name\":\"tokenToTokenSwapInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_tokens_bought\"},{\"type\":\"uint256\",\"name\":\"min_eth_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":51007},{\"name\":\"tokenToTokenTransferInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_tokens_bought\"},{\"type\":\"uint256\",\"name\":\"min_eth_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":51098},{\"name\":\"tokenToTokenSwapOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens_sold\"},{\"type\":\"uint256\",\"name\":\"max_eth_sold\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":54928},{\"name\":\"tokenToTokenTransferOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens_sold\"},{\"type\":\"uint256\",\"name\":\"max_eth_sold\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":55019},{\"name\":\"tokenToExchangeSwapInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_tokens_bought\"},{\"type\":\"uint256\",\"name\":\"min_eth_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"exchange_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":49342},{\"name\":\"tokenToExchangeTransferInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_tokens_bought\"},{\"type\":\"uint256\",\"name\":\"min_eth_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"address\",\"name\":\"exchange_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":49532},{\"name\":\"tokenToExchangeSwapOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens_sold\"},{\"type\":\"uint256\",\"name\":\"max_eth_sold\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"exchange_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":53233},{\"name\":\"tokenToExchangeTransferOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens_sold\"},{\"type\":\"uint256\",\"name\":\"max_eth_sold\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"address\",\"name\":\"exchange_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":53423},{\"name\":\"getEthToTokenInputPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"eth_sold\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":5542},{\"name\":\"getEthToTokenOutputPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6872},{\"name\":\"getTokenToEthInputPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":5637},{\"name\":\"getTokenToEthOutputPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"eth_bought\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6897},{\"name\":\"tokenAddress\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1413},{\"name\":\"factoryAddress\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1443},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1645},{\"name\":\"transfer\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":75034},{\"name\":\"transferFrom\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":110907},{\"name\":\"approve\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":38769},{\"name\":\"allowance\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"},{\"type\":\"address\",\"name\":\"_spender\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1925},{\"name\":\"name\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1623},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1653},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1683},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1713}]"

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
}

// bindExchange binds a generic wrapper to an already deployed contract.
func bindExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 out)
func (_Exchange *ExchangeCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 out)
func (_Exchange *ExchangeSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Exchange.Contract.Allowance(&_Exchange.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 out)
func (_Exchange *ExchangeCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Exchange.Contract.Allowance(&_Exchange.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 out)
func (_Exchange *ExchangeCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 out)
func (_Exchange *ExchangeSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Exchange.Contract.BalanceOf(&_Exchange.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 out)
func (_Exchange *ExchangeCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Exchange.Contract.BalanceOf(&_Exchange.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256 out)
func (_Exchange *ExchangeCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256 out)
func (_Exchange *ExchangeSession) Decimals() (*big.Int, error) {
	return _Exchange.Contract.Decimals(&_Exchange.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256 out)
func (_Exchange *ExchangeCallerSession) Decimals() (*big.Int, error) {
	return _Exchange.Contract.Decimals(&_Exchange.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() constant returns(address out)
func (_Exchange *ExchangeCaller) FactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "factoryAddress")
	return *ret0, err
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() constant returns(address out)
func (_Exchange *ExchangeSession) FactoryAddress() (common.Address, error) {
	return _Exchange.Contract.FactoryAddress(&_Exchange.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() constant returns(address out)
func (_Exchange *ExchangeCallerSession) FactoryAddress() (common.Address, error) {
	return _Exchange.Contract.FactoryAddress(&_Exchange.CallOpts)
}

// GetEthToTokenInputPrice is a free data retrieval call binding the contract method 0xcd7724c3.
//
// Solidity: function getEthToTokenInputPrice(uint256 eth_sold) constant returns(uint256 out)
func (_Exchange *ExchangeCaller) GetEthToTokenInputPrice(opts *bind.CallOpts, eth_sold *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "getEthToTokenInputPrice", eth_sold)
	return *ret0, err
}

// GetEthToTokenInputPrice is a free data retrieval call binding the contract method 0xcd7724c3.
//
// Solidity: function getEthToTokenInputPrice(uint256 eth_sold) constant returns(uint256 out)
func (_Exchange *ExchangeSession) GetEthToTokenInputPrice(eth_sold *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetEthToTokenInputPrice(&_Exchange.CallOpts, eth_sold)
}

// GetEthToTokenInputPrice is a free data retrieval call binding the contract method 0xcd7724c3.
//
// Solidity: function getEthToTokenInputPrice(uint256 eth_sold) constant returns(uint256 out)
func (_Exchange *ExchangeCallerSession) GetEthToTokenInputPrice(eth_sold *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetEthToTokenInputPrice(&_Exchange.CallOpts, eth_sold)
}

// GetEthToTokenOutputPrice is a free data retrieval call binding the contract method 0x59e94862.
//
// Solidity: function getEthToTokenOutputPrice(uint256 tokens_bought) constant returns(uint256 out)
func (_Exchange *ExchangeCaller) GetEthToTokenOutputPrice(opts *bind.CallOpts, tokens_bought *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "getEthToTokenOutputPrice", tokens_bought)
	return *ret0, err
}

// GetEthToTokenOutputPrice is a free data retrieval call binding the contract method 0x59e94862.
//
// Solidity: function getEthToTokenOutputPrice(uint256 tokens_bought) constant returns(uint256 out)
func (_Exchange *ExchangeSession) GetEthToTokenOutputPrice(tokens_bought *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetEthToTokenOutputPrice(&_Exchange.CallOpts, tokens_bought)
}

// GetEthToTokenOutputPrice is a free data retrieval call binding the contract method 0x59e94862.
//
// Solidity: function getEthToTokenOutputPrice(uint256 tokens_bought) constant returns(uint256 out)
func (_Exchange *ExchangeCallerSession) GetEthToTokenOutputPrice(tokens_bought *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetEthToTokenOutputPrice(&_Exchange.CallOpts, tokens_bought)
}

// GetTokenToEthInputPrice is a free data retrieval call binding the contract method 0x95b68fe7.
//
// Solidity: function getTokenToEthInputPrice(uint256 tokens_sold) constant returns(uint256 out)
func (_Exchange *ExchangeCaller) GetTokenToEthInputPrice(opts *bind.CallOpts, tokens_sold *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "getTokenToEthInputPrice", tokens_sold)
	return *ret0, err
}

// GetTokenToEthInputPrice is a free data retrieval call binding the contract method 0x95b68fe7.
//
// Solidity: function getTokenToEthInputPrice(uint256 tokens_sold) constant returns(uint256 out)
func (_Exchange *ExchangeSession) GetTokenToEthInputPrice(tokens_sold *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetTokenToEthInputPrice(&_Exchange.CallOpts, tokens_sold)
}

// GetTokenToEthInputPrice is a free data retrieval call binding the contract method 0x95b68fe7.
//
// Solidity: function getTokenToEthInputPrice(uint256 tokens_sold) constant returns(uint256 out)
func (_Exchange *ExchangeCallerSession) GetTokenToEthInputPrice(tokens_sold *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetTokenToEthInputPrice(&_Exchange.CallOpts, tokens_sold)
}

// GetTokenToEthOutputPrice is a free data retrieval call binding the contract method 0x2640f62c.
//
// Solidity: function getTokenToEthOutputPrice(uint256 eth_bought) constant returns(uint256 out)
func (_Exchange *ExchangeCaller) GetTokenToEthOutputPrice(opts *bind.CallOpts, eth_bought *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "getTokenToEthOutputPrice", eth_bought)
	return *ret0, err
}

// GetTokenToEthOutputPrice is a free data retrieval call binding the contract method 0x2640f62c.
//
// Solidity: function getTokenToEthOutputPrice(uint256 eth_bought) constant returns(uint256 out)
func (_Exchange *ExchangeSession) GetTokenToEthOutputPrice(eth_bought *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetTokenToEthOutputPrice(&_Exchange.CallOpts, eth_bought)
}

// GetTokenToEthOutputPrice is a free data retrieval call binding the contract method 0x2640f62c.
//
// Solidity: function getTokenToEthOutputPrice(uint256 eth_bought) constant returns(uint256 out)
func (_Exchange *ExchangeCallerSession) GetTokenToEthOutputPrice(eth_bought *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetTokenToEthOutputPrice(&_Exchange.CallOpts, eth_bought)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32 out)
func (_Exchange *ExchangeCaller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32 out)
func (_Exchange *ExchangeSession) Name() ([32]byte, error) {
	return _Exchange.Contract.Name(&_Exchange.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32 out)
func (_Exchange *ExchangeCallerSession) Name() ([32]byte, error) {
	return _Exchange.Contract.Name(&_Exchange.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32 out)
func (_Exchange *ExchangeCaller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32 out)
func (_Exchange *ExchangeSession) Symbol() ([32]byte, error) {
	return _Exchange.Contract.Symbol(&_Exchange.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32 out)
func (_Exchange *ExchangeCallerSession) Symbol() ([32]byte, error) {
	return _Exchange.Contract.Symbol(&_Exchange.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address out)
func (_Exchange *ExchangeCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "tokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address out)
func (_Exchange *ExchangeSession) TokenAddress() (common.Address, error) {
	return _Exchange.Contract.TokenAddress(&_Exchange.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address out)
func (_Exchange *ExchangeCallerSession) TokenAddress() (common.Address, error) {
	return _Exchange.Contract.TokenAddress(&_Exchange.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 out)
func (_Exchange *ExchangeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 out)
func (_Exchange *ExchangeSession) TotalSupply() (*big.Int, error) {
	return _Exchange.Contract.TotalSupply(&_Exchange.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 out)
func (_Exchange *ExchangeCallerSession) TotalSupply() (*big.Int, error) {
	return _Exchange.Contract.TotalSupply(&_Exchange.CallOpts)
}

// Default is a paid mutator transaction binding the contract method 0x89402a72.
//
// Solidity: function __default__() returns()
func (_Exchange *ExchangeTransactor) Default(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "__default__")
}

// Default is a paid mutator transaction binding the contract method 0x89402a72.
//
// Solidity: function __default__() returns()
func (_Exchange *ExchangeSession) Default() (*types.Transaction, error) {
	return _Exchange.Contract.Default(&_Exchange.TransactOpts)
}

// Default is a paid mutator transaction binding the contract method 0x89402a72.
//
// Solidity: function __default__() returns()
func (_Exchange *ExchangeTransactorSession) Default() (*types.Transaction, error) {
	return _Exchange.Contract.Default(&_Exchange.TransactOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x422f1043.
//
// Solidity: function addLiquidity(uint256 min_liquidity, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Exchange *ExchangeTransactor) AddLiquidity(opts *bind.TransactOpts, min_liquidity *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "addLiquidity", min_liquidity, max_tokens, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x422f1043.
//
// Solidity: function addLiquidity(uint256 min_liquidity, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Exchange *ExchangeSession) AddLiquidity(min_liquidity *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.AddLiquidity(&_Exchange.TransactOpts, min_liquidity, max_tokens, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x422f1043.
//
// Solidity: function addLiquidity(uint256 min_liquidity, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) AddLiquidity(min_liquidity *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.AddLiquidity(&_Exchange.TransactOpts, min_liquidity, max_tokens, deadline)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool out)
func (_Exchange *ExchangeTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool out)
func (_Exchange *ExchangeSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Approve(&_Exchange.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool out)
func (_Exchange *ExchangeTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Approve(&_Exchange.TransactOpts, _spender, _value)
}

// EthToTokenSwapInput is a paid mutator transaction binding the contract method 0xf39b5b9b.
//
// Solidity: function ethToTokenSwapInput(uint256 min_tokens, uint256 deadline) returns(uint256 out)
func (_Exchange *ExchangeTransactor) EthToTokenSwapInput(opts *bind.TransactOpts, min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "ethToTokenSwapInput", min_tokens, deadline)
}

// EthToTokenSwapInput is a paid mutator transaction binding the contract method 0xf39b5b9b.
//
// Solidity: function ethToTokenSwapInput(uint256 min_tokens, uint256 deadline) returns(uint256 out)
func (_Exchange *ExchangeSession) EthToTokenSwapInput(min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.EthToTokenSwapInput(&_Exchange.TransactOpts, min_tokens, deadline)
}

// EthToTokenSwapInput is a paid mutator transaction binding the contract method 0xf39b5b9b.
//
// Solidity: function ethToTokenSwapInput(uint256 min_tokens, uint256 deadline) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) EthToTokenSwapInput(min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.EthToTokenSwapInput(&_Exchange.TransactOpts, min_tokens, deadline)
}

// EthToTokenSwapOutput is a paid mutator transaction binding the contract method 0x6b1d4db7.
//
// Solidity: function ethToTokenSwapOutput(uint256 tokens_bought, uint256 deadline) returns(uint256 out)
func (_Exchange *ExchangeTransactor) EthToTokenSwapOutput(opts *bind.TransactOpts, tokens_bought *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "ethToTokenSwapOutput", tokens_bought, deadline)
}

// EthToTokenSwapOutput is a paid mutator transaction binding the contract method 0x6b1d4db7.
//
// Solidity: function ethToTokenSwapOutput(uint256 tokens_bought, uint256 deadline) returns(uint256 out)
func (_Exchange *ExchangeSession) EthToTokenSwapOutput(tokens_bought *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.EthToTokenSwapOutput(&_Exchange.TransactOpts, tokens_bought, deadline)
}

// EthToTokenSwapOutput is a paid mutator transaction binding the contract method 0x6b1d4db7.
//
// Solidity: function ethToTokenSwapOutput(uint256 tokens_bought, uint256 deadline) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) EthToTokenSwapOutput(tokens_bought *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.EthToTokenSwapOutput(&_Exchange.TransactOpts, tokens_bought, deadline)
}

// EthToTokenTransferInput is a paid mutator transaction binding the contract method 0xad65d76d.
//
// Solidity: function ethToTokenTransferInput(uint256 min_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Exchange *ExchangeTransactor) EthToTokenTransferInput(opts *bind.TransactOpts, min_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "ethToTokenTransferInput", min_tokens, deadline, recipient)
}

// EthToTokenTransferInput is a paid mutator transaction binding the contract method 0xad65d76d.
//
// Solidity: function ethToTokenTransferInput(uint256 min_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Exchange *ExchangeSession) EthToTokenTransferInput(min_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.EthToTokenTransferInput(&_Exchange.TransactOpts, min_tokens, deadline, recipient)
}

// EthToTokenTransferInput is a paid mutator transaction binding the contract method 0xad65d76d.
//
// Solidity: function ethToTokenTransferInput(uint256 min_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) EthToTokenTransferInput(min_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.EthToTokenTransferInput(&_Exchange.TransactOpts, min_tokens, deadline, recipient)
}

// EthToTokenTransferOutput is a paid mutator transaction binding the contract method 0x0b573638.
//
// Solidity: function ethToTokenTransferOutput(uint256 tokens_bought, uint256 deadline, address recipient) returns(uint256 out)
func (_Exchange *ExchangeTransactor) EthToTokenTransferOutput(opts *bind.TransactOpts, tokens_bought *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "ethToTokenTransferOutput", tokens_bought, deadline, recipient)
}

// EthToTokenTransferOutput is a paid mutator transaction binding the contract method 0x0b573638.
//
// Solidity: function ethToTokenTransferOutput(uint256 tokens_bought, uint256 deadline, address recipient) returns(uint256 out)
func (_Exchange *ExchangeSession) EthToTokenTransferOutput(tokens_bought *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.EthToTokenTransferOutput(&_Exchange.TransactOpts, tokens_bought, deadline, recipient)
}

// EthToTokenTransferOutput is a paid mutator transaction binding the contract method 0x0b573638.
//
// Solidity: function ethToTokenTransferOutput(uint256 tokens_bought, uint256 deadline, address recipient) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) EthToTokenTransferOutput(tokens_bought *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.EthToTokenTransferOutput(&_Exchange.TransactOpts, tokens_bought, deadline, recipient)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xf88bf15a.
//
// Solidity: function removeLiquidity(uint256 amount, uint256 min_eth, uint256 min_tokens, uint256 deadline) returns(uint256 out, uint256 out)
func (_Exchange *ExchangeTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int, min_eth *big.Int, min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "removeLiquidity", amount, min_eth, min_tokens, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xf88bf15a.
//
// Solidity: function removeLiquidity(uint256 amount, uint256 min_eth, uint256 min_tokens, uint256 deadline) returns(uint256 out, uint256 out)
func (_Exchange *ExchangeSession) RemoveLiquidity(amount *big.Int, min_eth *big.Int, min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveLiquidity(&_Exchange.TransactOpts, amount, min_eth, min_tokens, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xf88bf15a.
//
// Solidity: function removeLiquidity(uint256 amount, uint256 min_eth, uint256 min_tokens, uint256 deadline) returns(uint256 out, uint256 out)
func (_Exchange *ExchangeTransactorSession) RemoveLiquidity(amount *big.Int, min_eth *big.Int, min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.RemoveLiquidity(&_Exchange.TransactOpts, amount, min_eth, min_tokens, deadline)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(address token_addr) returns()
func (_Exchange *ExchangeTransactor) Setup(opts *bind.TransactOpts, token_addr common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "setup", token_addr)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(address token_addr) returns()
func (_Exchange *ExchangeSession) Setup(token_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.Setup(&_Exchange.TransactOpts, token_addr)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(address token_addr) returns()
func (_Exchange *ExchangeTransactorSession) Setup(token_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.Setup(&_Exchange.TransactOpts, token_addr)
}

// TokenToEthSwapInput is a paid mutator transaction binding the contract method 0x95e3c50b.
//
// Solidity: function tokenToEthSwapInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline) returns(uint256 out)
func (_Exchange *ExchangeTransactor) TokenToEthSwapInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "tokenToEthSwapInput", tokens_sold, min_eth, deadline)
}

// TokenToEthSwapInput is a paid mutator transaction binding the contract method 0x95e3c50b.
//
// Solidity: function tokenToEthSwapInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline) returns(uint256 out)
func (_Exchange *ExchangeSession) TokenToEthSwapInput(tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToEthSwapInput(&_Exchange.TransactOpts, tokens_sold, min_eth, deadline)
}

// TokenToEthSwapInput is a paid mutator transaction binding the contract method 0x95e3c50b.
//
// Solidity: function tokenToEthSwapInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) TokenToEthSwapInput(tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToEthSwapInput(&_Exchange.TransactOpts, tokens_sold, min_eth, deadline)
}

// TokenToEthSwapOutput is a paid mutator transaction binding the contract method 0x013efd8b.
//
// Solidity: function tokenToEthSwapOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Exchange *ExchangeTransactor) TokenToEthSwapOutput(opts *bind.TransactOpts, eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "tokenToEthSwapOutput", eth_bought, max_tokens, deadline)
}

// TokenToEthSwapOutput is a paid mutator transaction binding the contract method 0x013efd8b.
//
// Solidity: function tokenToEthSwapOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Exchange *ExchangeSession) TokenToEthSwapOutput(eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToEthSwapOutput(&_Exchange.TransactOpts, eth_bought, max_tokens, deadline)
}

// TokenToEthSwapOutput is a paid mutator transaction binding the contract method 0x013efd8b.
//
// Solidity: function tokenToEthSwapOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) TokenToEthSwapOutput(eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToEthSwapOutput(&_Exchange.TransactOpts, eth_bought, max_tokens, deadline)
}

// TokenToEthTransferInput is a paid mutator transaction binding the contract method 0x7237e031.
//
// Solidity: function tokenToEthTransferInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline, address recipient) returns(uint256 out)
func (_Exchange *ExchangeTransactor) TokenToEthTransferInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "tokenToEthTransferInput", tokens_sold, min_eth, deadline, recipient)
}

// TokenToEthTransferInput is a paid mutator transaction binding the contract method 0x7237e031.
//
// Solidity: function tokenToEthTransferInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline, address recipient) returns(uint256 out)
func (_Exchange *ExchangeSession) TokenToEthTransferInput(tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToEthTransferInput(&_Exchange.TransactOpts, tokens_sold, min_eth, deadline, recipient)
}

// TokenToEthTransferInput is a paid mutator transaction binding the contract method 0x7237e031.
//
// Solidity: function tokenToEthTransferInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline, address recipient) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) TokenToEthTransferInput(tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToEthTransferInput(&_Exchange.TransactOpts, tokens_sold, min_eth, deadline, recipient)
}

// TokenToEthTransferOutput is a paid mutator transaction binding the contract method 0xd4e4841d.
//
// Solidity: function tokenToEthTransferOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Exchange *ExchangeTransactor) TokenToEthTransferOutput(opts *bind.TransactOpts, eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "tokenToEthTransferOutput", eth_bought, max_tokens, deadline, recipient)
}

// TokenToEthTransferOutput is a paid mutator transaction binding the contract method 0xd4e4841d.
//
// Solidity: function tokenToEthTransferOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Exchange *ExchangeSession) TokenToEthTransferOutput(eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToEthTransferOutput(&_Exchange.TransactOpts, eth_bought, max_tokens, deadline, recipient)
}

// TokenToEthTransferOutput is a paid mutator transaction binding the contract method 0xd4e4841d.
//
// Solidity: function tokenToEthTransferOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) TokenToEthTransferOutput(eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToEthTransferOutput(&_Exchange.TransactOpts, eth_bought, max_tokens, deadline, recipient)
}

// TokenToExchangeSwapInput is a paid mutator transaction binding the contract method 0xb1cb43bf.
//
// Solidity: function tokenToExchangeSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Exchange *ExchangeTransactor) TokenToExchangeSwapInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "tokenToExchangeSwapInput", tokens_sold, min_tokens_bought, min_eth_bought, deadline, exchange_addr)
}

// TokenToExchangeSwapInput is a paid mutator transaction binding the contract method 0xb1cb43bf.
//
// Solidity: function tokenToExchangeSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Exchange *ExchangeSession) TokenToExchangeSwapInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToExchangeSwapInput(&_Exchange.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, exchange_addr)
}

// TokenToExchangeSwapInput is a paid mutator transaction binding the contract method 0xb1cb43bf.
//
// Solidity: function tokenToExchangeSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) TokenToExchangeSwapInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToExchangeSwapInput(&_Exchange.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, exchange_addr)
}

// TokenToExchangeSwapOutput is a paid mutator transaction binding the contract method 0xea650c7d.
//
// Solidity: function tokenToExchangeSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Exchange *ExchangeTransactor) TokenToExchangeSwapOutput(opts *bind.TransactOpts, tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "tokenToExchangeSwapOutput", tokens_bought, max_tokens_sold, max_eth_sold, deadline, exchange_addr)
}

// TokenToExchangeSwapOutput is a paid mutator transaction binding the contract method 0xea650c7d.
//
// Solidity: function tokenToExchangeSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Exchange *ExchangeSession) TokenToExchangeSwapOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToExchangeSwapOutput(&_Exchange.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, exchange_addr)
}

// TokenToExchangeSwapOutput is a paid mutator transaction binding the contract method 0xea650c7d.
//
// Solidity: function tokenToExchangeSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) TokenToExchangeSwapOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToExchangeSwapOutput(&_Exchange.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, exchange_addr)
}

// TokenToExchangeTransferInput is a paid mutator transaction binding the contract method 0xec384a3e.
//
// Solidity: function tokenToExchangeTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Exchange *ExchangeTransactor) TokenToExchangeTransferInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "tokenToExchangeTransferInput", tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferInput is a paid mutator transaction binding the contract method 0xec384a3e.
//
// Solidity: function tokenToExchangeTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Exchange *ExchangeSession) TokenToExchangeTransferInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToExchangeTransferInput(&_Exchange.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferInput is a paid mutator transaction binding the contract method 0xec384a3e.
//
// Solidity: function tokenToExchangeTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) TokenToExchangeTransferInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToExchangeTransferInput(&_Exchange.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferOutput is a paid mutator transaction binding the contract method 0x981a1327.
//
// Solidity: function tokenToExchangeTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Exchange *ExchangeTransactor) TokenToExchangeTransferOutput(opts *bind.TransactOpts, tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "tokenToExchangeTransferOutput", tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferOutput is a paid mutator transaction binding the contract method 0x981a1327.
//
// Solidity: function tokenToExchangeTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Exchange *ExchangeSession) TokenToExchangeTransferOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToExchangeTransferOutput(&_Exchange.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferOutput is a paid mutator transaction binding the contract method 0x981a1327.
//
// Solidity: function tokenToExchangeTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) TokenToExchangeTransferOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToExchangeTransferOutput(&_Exchange.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, exchange_addr)
}

// TokenToTokenSwapInput is a paid mutator transaction binding the contract method 0xddf7e1a7.
//
// Solidity: function tokenToTokenSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address token_addr) returns(uint256 out)
func (_Exchange *ExchangeTransactor) TokenToTokenSwapInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "tokenToTokenSwapInput", tokens_sold, min_tokens_bought, min_eth_bought, deadline, token_addr)
}

// TokenToTokenSwapInput is a paid mutator transaction binding the contract method 0xddf7e1a7.
//
// Solidity: function tokenToTokenSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address token_addr) returns(uint256 out)
func (_Exchange *ExchangeSession) TokenToTokenSwapInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToTokenSwapInput(&_Exchange.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, token_addr)
}

// TokenToTokenSwapInput is a paid mutator transaction binding the contract method 0xddf7e1a7.
//
// Solidity: function tokenToTokenSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address token_addr) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) TokenToTokenSwapInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToTokenSwapInput(&_Exchange.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, token_addr)
}

// TokenToTokenSwapOutput is a paid mutator transaction binding the contract method 0xb040d545.
//
// Solidity: function tokenToTokenSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address token_addr) returns(uint256 out)
func (_Exchange *ExchangeTransactor) TokenToTokenSwapOutput(opts *bind.TransactOpts, tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "tokenToTokenSwapOutput", tokens_bought, max_tokens_sold, max_eth_sold, deadline, token_addr)
}

// TokenToTokenSwapOutput is a paid mutator transaction binding the contract method 0xb040d545.
//
// Solidity: function tokenToTokenSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address token_addr) returns(uint256 out)
func (_Exchange *ExchangeSession) TokenToTokenSwapOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToTokenSwapOutput(&_Exchange.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, token_addr)
}

// TokenToTokenSwapOutput is a paid mutator transaction binding the contract method 0xb040d545.
//
// Solidity: function tokenToTokenSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address token_addr) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) TokenToTokenSwapOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToTokenSwapOutput(&_Exchange.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, token_addr)
}

// TokenToTokenTransferInput is a paid mutator transaction binding the contract method 0xf552d91b.
//
// Solidity: function tokenToTokenTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Exchange *ExchangeTransactor) TokenToTokenTransferInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "tokenToTokenTransferInput", tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, token_addr)
}

// TokenToTokenTransferInput is a paid mutator transaction binding the contract method 0xf552d91b.
//
// Solidity: function tokenToTokenTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Exchange *ExchangeSession) TokenToTokenTransferInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToTokenTransferInput(&_Exchange.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, token_addr)
}

// TokenToTokenTransferInput is a paid mutator transaction binding the contract method 0xf552d91b.
//
// Solidity: function tokenToTokenTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) TokenToTokenTransferInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToTokenTransferInput(&_Exchange.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, token_addr)
}

// TokenToTokenTransferOutput is a paid mutator transaction binding the contract method 0xf3c0efe9.
//
// Solidity: function tokenToTokenTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Exchange *ExchangeTransactor) TokenToTokenTransferOutput(opts *bind.TransactOpts, tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "tokenToTokenTransferOutput", tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, token_addr)
}

// TokenToTokenTransferOutput is a paid mutator transaction binding the contract method 0xf3c0efe9.
//
// Solidity: function tokenToTokenTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Exchange *ExchangeSession) TokenToTokenTransferOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToTokenTransferOutput(&_Exchange.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, token_addr)
}

// TokenToTokenTransferOutput is a paid mutator transaction binding the contract method 0xf3c0efe9.
//
// Solidity: function tokenToTokenTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Exchange *ExchangeTransactorSession) TokenToTokenTransferOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TokenToTokenTransferOutput(&_Exchange.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, token_addr)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool out)
func (_Exchange *ExchangeTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool out)
func (_Exchange *ExchangeSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Transfer(&_Exchange.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool out)
func (_Exchange *ExchangeTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Transfer(&_Exchange.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool out)
func (_Exchange *ExchangeTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool out)
func (_Exchange *ExchangeSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.TransferFrom(&_Exchange.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool out)
func (_Exchange *ExchangeTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.TransferFrom(&_Exchange.TransactOpts, _from, _to, _value)
}

// ExchangeAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Exchange contract.
type ExchangeAddLiquidityIterator struct {
	Event *ExchangeAddLiquidity // Event containing the contract specifics and raw log

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
func (it *ExchangeAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeAddLiquidity)
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
		it.Event = new(ExchangeAddLiquidity)
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
func (it *ExchangeAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeAddLiquidity represents a AddLiquidity event raised by the Exchange contract.
type ExchangeAddLiquidity struct {
	Provider    common.Address
	EthAmount   *big.Int
	TokenAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca.
//
// Solidity: event AddLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_Exchange *ExchangeFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address, eth_amount []*big.Int, token_amount []*big.Int) (*ExchangeAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var eth_amountRule []interface{}
	for _, eth_amountItem := range eth_amount {
		eth_amountRule = append(eth_amountRule, eth_amountItem)
	}
	var token_amountRule []interface{}
	for _, token_amountItem := range token_amount {
		token_amountRule = append(token_amountRule, token_amountItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "AddLiquidity", providerRule, eth_amountRule, token_amountRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeAddLiquidityIterator{contract: _Exchange.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca.
//
// Solidity: event AddLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_Exchange *ExchangeFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *ExchangeAddLiquidity, provider []common.Address, eth_amount []*big.Int, token_amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var eth_amountRule []interface{}
	for _, eth_amountItem := range eth_amount {
		eth_amountRule = append(eth_amountRule, eth_amountItem)
	}
	var token_amountRule []interface{}
	for _, token_amountItem := range token_amount {
		token_amountRule = append(token_amountRule, token_amountItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "AddLiquidity", providerRule, eth_amountRule, token_amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeAddLiquidity)
				if err := _Exchange.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca.
//
// Solidity: event AddLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_Exchange *ExchangeFilterer) ParseAddLiquidity(log types.Log) (*ExchangeAddLiquidity, error) {
	event := new(ExchangeAddLiquidity)
	if err := _Exchange.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExchangeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Exchange contract.
type ExchangeApprovalIterator struct {
	Event *ExchangeApproval // Event containing the contract specifics and raw log

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
func (it *ExchangeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeApproval)
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
		it.Event = new(ExchangeApproval)
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
func (it *ExchangeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeApproval represents a Approval event raised by the Exchange contract.
type ExchangeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Exchange *ExchangeFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*ExchangeApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeApprovalIterator{contract: _Exchange.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Exchange *ExchangeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ExchangeApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeApproval)
				if err := _Exchange.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Exchange *ExchangeFilterer) ParseApproval(log types.Log) (*ExchangeApproval, error) {
	event := new(ExchangeApproval)
	if err := _Exchange.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExchangeEthPurchaseIterator is returned from FilterEthPurchase and is used to iterate over the raw logs and unpacked data for EthPurchase events raised by the Exchange contract.
type ExchangeEthPurchaseIterator struct {
	Event *ExchangeEthPurchase // Event containing the contract specifics and raw log

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
func (it *ExchangeEthPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeEthPurchase)
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
		it.Event = new(ExchangeEthPurchase)
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
func (it *ExchangeEthPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeEthPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeEthPurchase represents a EthPurchase event raised by the Exchange contract.
type ExchangeEthPurchase struct {
	Buyer      common.Address
	TokensSold *big.Int
	EthBought  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEthPurchase is a free log retrieval operation binding the contract event 0x7f4091b46c33e918a0f3aa42307641d17bb67029427a5369e54b353984238705.
//
// Solidity: event EthPurchase(address indexed buyer, uint256 indexed tokens_sold, uint256 indexed eth_bought)
func (_Exchange *ExchangeFilterer) FilterEthPurchase(opts *bind.FilterOpts, buyer []common.Address, tokens_sold []*big.Int, eth_bought []*big.Int) (*ExchangeEthPurchaseIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokens_soldRule []interface{}
	for _, tokens_soldItem := range tokens_sold {
		tokens_soldRule = append(tokens_soldRule, tokens_soldItem)
	}
	var eth_boughtRule []interface{}
	for _, eth_boughtItem := range eth_bought {
		eth_boughtRule = append(eth_boughtRule, eth_boughtItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "EthPurchase", buyerRule, tokens_soldRule, eth_boughtRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeEthPurchaseIterator{contract: _Exchange.contract, event: "EthPurchase", logs: logs, sub: sub}, nil
}

// WatchEthPurchase is a free log subscription operation binding the contract event 0x7f4091b46c33e918a0f3aa42307641d17bb67029427a5369e54b353984238705.
//
// Solidity: event EthPurchase(address indexed buyer, uint256 indexed tokens_sold, uint256 indexed eth_bought)
func (_Exchange *ExchangeFilterer) WatchEthPurchase(opts *bind.WatchOpts, sink chan<- *ExchangeEthPurchase, buyer []common.Address, tokens_sold []*big.Int, eth_bought []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokens_soldRule []interface{}
	for _, tokens_soldItem := range tokens_sold {
		tokens_soldRule = append(tokens_soldRule, tokens_soldItem)
	}
	var eth_boughtRule []interface{}
	for _, eth_boughtItem := range eth_bought {
		eth_boughtRule = append(eth_boughtRule, eth_boughtItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "EthPurchase", buyerRule, tokens_soldRule, eth_boughtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeEthPurchase)
				if err := _Exchange.contract.UnpackLog(event, "EthPurchase", log); err != nil {
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

// ParseEthPurchase is a log parse operation binding the contract event 0x7f4091b46c33e918a0f3aa42307641d17bb67029427a5369e54b353984238705.
//
// Solidity: event EthPurchase(address indexed buyer, uint256 indexed tokens_sold, uint256 indexed eth_bought)
func (_Exchange *ExchangeFilterer) ParseEthPurchase(log types.Log) (*ExchangeEthPurchase, error) {
	event := new(ExchangeEthPurchase)
	if err := _Exchange.contract.UnpackLog(event, "EthPurchase", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExchangeRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Exchange contract.
type ExchangeRemoveLiquidityIterator struct {
	Event *ExchangeRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *ExchangeRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeRemoveLiquidity)
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
		it.Event = new(ExchangeRemoveLiquidity)
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
func (it *ExchangeRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeRemoveLiquidity represents a RemoveLiquidity event raised by the Exchange contract.
type ExchangeRemoveLiquidity struct {
	Provider    common.Address
	EthAmount   *big.Int
	TokenAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x0fbf06c058b90cb038a618f8c2acbf6145f8b3570fd1fa56abb8f0f3f05b36e8.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_Exchange *ExchangeFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address, eth_amount []*big.Int, token_amount []*big.Int) (*ExchangeRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var eth_amountRule []interface{}
	for _, eth_amountItem := range eth_amount {
		eth_amountRule = append(eth_amountRule, eth_amountItem)
	}
	var token_amountRule []interface{}
	for _, token_amountItem := range token_amount {
		token_amountRule = append(token_amountRule, token_amountItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "RemoveLiquidity", providerRule, eth_amountRule, token_amountRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeRemoveLiquidityIterator{contract: _Exchange.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x0fbf06c058b90cb038a618f8c2acbf6145f8b3570fd1fa56abb8f0f3f05b36e8.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_Exchange *ExchangeFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *ExchangeRemoveLiquidity, provider []common.Address, eth_amount []*big.Int, token_amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var eth_amountRule []interface{}
	for _, eth_amountItem := range eth_amount {
		eth_amountRule = append(eth_amountRule, eth_amountItem)
	}
	var token_amountRule []interface{}
	for _, token_amountItem := range token_amount {
		token_amountRule = append(token_amountRule, token_amountItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "RemoveLiquidity", providerRule, eth_amountRule, token_amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeRemoveLiquidity)
				if err := _Exchange.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x0fbf06c058b90cb038a618f8c2acbf6145f8b3570fd1fa56abb8f0f3f05b36e8.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_Exchange *ExchangeFilterer) ParseRemoveLiquidity(log types.Log) (*ExchangeRemoveLiquidity, error) {
	event := new(ExchangeRemoveLiquidity)
	if err := _Exchange.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExchangeTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the Exchange contract.
type ExchangeTokenPurchaseIterator struct {
	Event *ExchangeTokenPurchase // Event containing the contract specifics and raw log

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
func (it *ExchangeTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeTokenPurchase)
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
		it.Event = new(ExchangeTokenPurchase)
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
func (it *ExchangeTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeTokenPurchase represents a TokenPurchase event raised by the Exchange contract.
type ExchangeTokenPurchase struct {
	Buyer        common.Address
	EthSold      *big.Int
	TokensBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0xcd60aa75dea3072fbc07ae6d7d856b5dc5f4eee88854f5b4abf7b680ef8bc50f.
//
// Solidity: event TokenPurchase(address indexed buyer, uint256 indexed eth_sold, uint256 indexed tokens_bought)
func (_Exchange *ExchangeFilterer) FilterTokenPurchase(opts *bind.FilterOpts, buyer []common.Address, eth_sold []*big.Int, tokens_bought []*big.Int) (*ExchangeTokenPurchaseIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var eth_soldRule []interface{}
	for _, eth_soldItem := range eth_sold {
		eth_soldRule = append(eth_soldRule, eth_soldItem)
	}
	var tokens_boughtRule []interface{}
	for _, tokens_boughtItem := range tokens_bought {
		tokens_boughtRule = append(tokens_boughtRule, tokens_boughtItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "TokenPurchase", buyerRule, eth_soldRule, tokens_boughtRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeTokenPurchaseIterator{contract: _Exchange.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0xcd60aa75dea3072fbc07ae6d7d856b5dc5f4eee88854f5b4abf7b680ef8bc50f.
//
// Solidity: event TokenPurchase(address indexed buyer, uint256 indexed eth_sold, uint256 indexed tokens_bought)
func (_Exchange *ExchangeFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *ExchangeTokenPurchase, buyer []common.Address, eth_sold []*big.Int, tokens_bought []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var eth_soldRule []interface{}
	for _, eth_soldItem := range eth_sold {
		eth_soldRule = append(eth_soldRule, eth_soldItem)
	}
	var tokens_boughtRule []interface{}
	for _, tokens_boughtItem := range tokens_bought {
		tokens_boughtRule = append(tokens_boughtRule, tokens_boughtItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "TokenPurchase", buyerRule, eth_soldRule, tokens_boughtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeTokenPurchase)
				if err := _Exchange.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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

// ParseTokenPurchase is a log parse operation binding the contract event 0xcd60aa75dea3072fbc07ae6d7d856b5dc5f4eee88854f5b4abf7b680ef8bc50f.
//
// Solidity: event TokenPurchase(address indexed buyer, uint256 indexed eth_sold, uint256 indexed tokens_bought)
func (_Exchange *ExchangeFilterer) ParseTokenPurchase(log types.Log) (*ExchangeTokenPurchase, error) {
	event := new(ExchangeTokenPurchase)
	if err := _Exchange.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExchangeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Exchange contract.
type ExchangeTransferIterator struct {
	Event *ExchangeTransfer // Event containing the contract specifics and raw log

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
func (it *ExchangeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeTransfer)
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
		it.Event = new(ExchangeTransfer)
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
func (it *ExchangeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeTransfer represents a Transfer event raised by the Exchange contract.
type ExchangeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Exchange *ExchangeFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ExchangeTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransferIterator{contract: _Exchange.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Exchange *ExchangeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ExchangeTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeTransfer)
				if err := _Exchange.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Exchange *ExchangeFilterer) ParseTransfer(log types.Log) (*ExchangeTransfer, error) {
	event := new(ExchangeTransfer)
	if err := _Exchange.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
