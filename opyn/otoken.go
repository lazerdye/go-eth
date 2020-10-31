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

// OTokenABI is the input ABI used to generate the binding from.
const OTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"vaultOwner\",\"type\":\"address\"},{\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"addERC20Collateral\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVaultOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"hasVault\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isExerciseWindow\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"vaultOwner\",\"type\":\"address\"}],\"name\":\"getVault\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"oTokensToIssue\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"issueOTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amtToCreate\",\"type\":\"uint256\"},{\"name\":\"amtCollateral\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"addAndSellERC20CollateralOption\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amtToRemove\",\"type\":\"uint256\"}],\"name\":\"removeCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquidationFactor\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"exponent\",\"type\":\"int32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amtToCreate\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"createAndSellETHCollateralOption\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"optionsExchange\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amtToCreate\",\"type\":\"uint256\"},{\"name\":\"amtCollateral\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"createERC20CollateralOption\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"oTokensToExercise\",\"type\":\"uint256\"},{\"name\":\"vaultsToExerciseFrom\",\"type\":\"address[]\"}],\"name\":\"exercise\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amtToCreate\",\"type\":\"uint256\"},{\"name\":\"amtCollateral\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"addERC20CollateralOption\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"collateralAmt\",\"type\":\"uint256\"}],\"name\":\"maxOTokensIssuable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"oTokensToExercise\",\"type\":\"uint256\"}],\"name\":\"underlyingRequiredToExercise\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"openVault\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"COMPOUND_ORACLE\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquidationIncentive\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"exponent\",\"type\":\"int32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasExpired\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vaultOwner\",\"type\":\"address\"}],\"name\":\"addETHCollateral\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionFee\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"exponent\",\"type\":\"int32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"strike\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlyingExp\",\"outputs\":[{\"name\":\"\",\"type\":\"int32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralExp\",\"outputs\":[{\"name\":\"\",\"type\":\"int32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oTokenExchangeRate\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"exponent\",\"type\":\"int32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"redeemVaultBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"setDetails\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amtToCreate\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"addETHCollateralOption\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minCollateralizationRatio\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"exponent\",\"type\":\"int32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vaultOwner\",\"type\":\"address\"},{\"name\":\"oTokensToLiquidate\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"strikePrice\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"exponent\",\"type\":\"int32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amtToCreate\",\"type\":\"uint256\"},{\"name\":\"amtCollateral\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"createAndSellERC20CollateralOption\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"vaultOwner\",\"type\":\"address\"}],\"name\":\"isUnsafe\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amtToCreate\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"addAndSellETHCollateralOption\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"vaultOwner\",\"type\":\"address\"}],\"name\":\"maxOTokensLiquidatable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"expiry\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"transferFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amtToBurn\",\"type\":\"uint256\"}],\"name\":\"burnOTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amtToCreate\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"createETHCollateralOption\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_liquidationIncentive\",\"type\":\"uint256\"},{\"name\":\"_liquidationFactor\",\"type\":\"uint256\"},{\"name\":\"_transactionFee\",\"type\":\"uint256\"},{\"name\":\"_minCollateralizationRatio\",\"type\":\"uint256\"}],\"name\":\"updateParameters\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ierc20\",\"type\":\"address\"}],\"name\":\"isETH\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeUnderlying\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_collateral\",\"type\":\"address\"},{\"name\":\"_collExp\",\"type\":\"int32\"},{\"name\":\"_underlying\",\"type\":\"address\"},{\"name\":\"_underlyingExp\",\"type\":\"int32\"},{\"name\":\"_oTokenExchangeExp\",\"type\":\"int32\"},{\"name\":\"_strikePrice\",\"type\":\"uint256\"},{\"name\":\"_strikeExp\",\"type\":\"int32\"},{\"name\":\"_strike\",\"type\":\"address\"},{\"name\":\"_expiry\",\"type\":\"uint256\"},{\"name\":\"_optionsExchange\",\"type\":\"address\"},{\"name\":\"_oracleAddress\",\"type\":\"address\"},{\"name\":\"_windowSize\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vaultOwner\",\"type\":\"address\"}],\"name\":\"VaultOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vaultOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"payer\",\"type\":\"address\"}],\"name\":\"ETHCollateralAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vaultOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"payer\",\"type\":\"address\"}],\"name\":\"ERC20CollateralAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"issuedTo\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oTokensIssued\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"vaultOwner\",\"type\":\"address\"}],\"name\":\"IssuedOTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amtCollateralToPay\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"vaultOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"liquidator\",\"type\":\"address\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amtUnderlyingToPay\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amtCollateralToPay\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"exerciser\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"vaultExercisedFrom\",\"type\":\"address\"}],\"name\":\"Exercise\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amtCollateralRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amtUnderlyingRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"vaultOwner\",\"type\":\"address\"}],\"name\":\"RedeemVaultBalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vaultOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oTokensBurned\",\"type\":\"uint256\"}],\"name\":\"BurnOTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amtRemoved\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"vaultOwner\",\"type\":\"address\"}],\"name\":\"RemoveCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"liquidationIncentive\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"liquidationFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"transactionFee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"minCollateralizationRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"UpdateParameters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"TransferFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amountUnderlying\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"vaultOwner\",\"type\":\"address\"}],\"name\":\"RemoveUnderlying\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OToken is an auto generated Go binding around an Ethereum contract.
type OToken struct {
	OTokenCaller     // Read-only binding to the contract
	OTokenTransactor // Write-only binding to the contract
	OTokenFilterer   // Log filterer for contract events
}

// OTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type OTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OTokenSession struct {
	Contract     *OToken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OTokenCallerSession struct {
	Contract *OTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OTokenTransactorSession struct {
	Contract     *OTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type OTokenRaw struct {
	Contract *OToken // Generic contract binding to access the raw methods on
}

// OTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OTokenCallerRaw struct {
	Contract *OTokenCaller // Generic read-only contract binding to access the raw methods on
}

// OTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OTokenTransactorRaw struct {
	Contract *OTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOToken creates a new instance of OToken, bound to a specific deployed contract.
func NewOToken(address common.Address, backend bind.ContractBackend) (*OToken, error) {
	contract, err := bindOToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OToken{OTokenCaller: OTokenCaller{contract: contract}, OTokenTransactor: OTokenTransactor{contract: contract}, OTokenFilterer: OTokenFilterer{contract: contract}}, nil
}

// NewOTokenCaller creates a new read-only instance of OToken, bound to a specific deployed contract.
func NewOTokenCaller(address common.Address, caller bind.ContractCaller) (*OTokenCaller, error) {
	contract, err := bindOToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OTokenCaller{contract: contract}, nil
}

// NewOTokenTransactor creates a new write-only instance of OToken, bound to a specific deployed contract.
func NewOTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*OTokenTransactor, error) {
	contract, err := bindOToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OTokenTransactor{contract: contract}, nil
}

// NewOTokenFilterer creates a new log filterer instance of OToken, bound to a specific deployed contract.
func NewOTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*OTokenFilterer, error) {
	contract, err := bindOToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OTokenFilterer{contract: contract}, nil
}

// bindOToken binds a generic wrapper to an already deployed contract.
func bindOToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OToken *OTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OToken.Contract.OTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OToken *OTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OToken.Contract.OTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OToken *OTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OToken.Contract.OTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OToken *OTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OToken *OTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OToken *OTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OToken.Contract.contract.Transact(opts, method, params...)
}

// COMPOUNDORACLE is a free data retrieval call binding the contract method 0x8a5e8cc7.
//
// Solidity: function COMPOUND_ORACLE() view returns(address)
func (_OToken *OTokenCaller) COMPOUNDORACLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "COMPOUND_ORACLE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COMPOUNDORACLE is a free data retrieval call binding the contract method 0x8a5e8cc7.
//
// Solidity: function COMPOUND_ORACLE() view returns(address)
func (_OToken *OTokenSession) COMPOUNDORACLE() (common.Address, error) {
	return _OToken.Contract.COMPOUNDORACLE(&_OToken.CallOpts)
}

// COMPOUNDORACLE is a free data retrieval call binding the contract method 0x8a5e8cc7.
//
// Solidity: function COMPOUND_ORACLE() view returns(address)
func (_OToken *OTokenCallerSession) COMPOUNDORACLE() (common.Address, error) {
	return _OToken.Contract.COMPOUNDORACLE(&_OToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_OToken *OTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_OToken *OTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _OToken.Contract.Allowance(&_OToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_OToken *OTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _OToken.Contract.Allowance(&_OToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OToken *OTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OToken *OTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _OToken.Contract.BalanceOf(&_OToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OToken *OTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _OToken.Contract.BalanceOf(&_OToken.CallOpts, account)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_OToken *OTokenCaller) Collateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "collateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_OToken *OTokenSession) Collateral() (common.Address, error) {
	return _OToken.Contract.Collateral(&_OToken.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_OToken *OTokenCallerSession) Collateral() (common.Address, error) {
	return _OToken.Contract.Collateral(&_OToken.CallOpts)
}

// CollateralExp is a free data retrieval call binding the contract method 0xb6e61c08.
//
// Solidity: function collateralExp() view returns(int32)
func (_OToken *OTokenCaller) CollateralExp(opts *bind.CallOpts) (int32, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "collateralExp")

	if err != nil {
		return *new(int32), err
	}

	out0 := *abi.ConvertType(out[0], new(int32)).(*int32)

	return out0, err

}

// CollateralExp is a free data retrieval call binding the contract method 0xb6e61c08.
//
// Solidity: function collateralExp() view returns(int32)
func (_OToken *OTokenSession) CollateralExp() (int32, error) {
	return _OToken.Contract.CollateralExp(&_OToken.CallOpts)
}

// CollateralExp is a free data retrieval call binding the contract method 0xb6e61c08.
//
// Solidity: function collateralExp() view returns(int32)
func (_OToken *OTokenCallerSession) CollateralExp() (int32, error) {
	return _OToken.Contract.CollateralExp(&_OToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OToken *OTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OToken *OTokenSession) Decimals() (uint8, error) {
	return _OToken.Contract.Decimals(&_OToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OToken *OTokenCallerSession) Decimals() (uint8, error) {
	return _OToken.Contract.Decimals(&_OToken.CallOpts)
}

// Expiry is a free data retrieval call binding the contract method 0xe184c9be.
//
// Solidity: function expiry() view returns(uint256)
func (_OToken *OTokenCaller) Expiry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "expiry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Expiry is a free data retrieval call binding the contract method 0xe184c9be.
//
// Solidity: function expiry() view returns(uint256)
func (_OToken *OTokenSession) Expiry() (*big.Int, error) {
	return _OToken.Contract.Expiry(&_OToken.CallOpts)
}

// Expiry is a free data retrieval call binding the contract method 0xe184c9be.
//
// Solidity: function expiry() view returns(uint256)
func (_OToken *OTokenCallerSession) Expiry() (*big.Int, error) {
	return _OToken.Contract.Expiry(&_OToken.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x0eb9af38.
//
// Solidity: function getVault(address vaultOwner) view returns(uint256, uint256, uint256, bool)
func (_OToken *OTokenCaller) GetVault(opts *bind.CallOpts, vaultOwner common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "getVault", vaultOwner)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)

	return out0, out1, out2, out3, err

}

// GetVault is a free data retrieval call binding the contract method 0x0eb9af38.
//
// Solidity: function getVault(address vaultOwner) view returns(uint256, uint256, uint256, bool)
func (_OToken *OTokenSession) GetVault(vaultOwner common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	return _OToken.Contract.GetVault(&_OToken.CallOpts, vaultOwner)
}

// GetVault is a free data retrieval call binding the contract method 0x0eb9af38.
//
// Solidity: function getVault(address vaultOwner) view returns(uint256, uint256, uint256, bool)
func (_OToken *OTokenCallerSession) GetVault(vaultOwner common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	return _OToken.Contract.GetVault(&_OToken.CallOpts, vaultOwner)
}

// GetVaultOwners is a free data retrieval call binding the contract method 0x060ab2ea.
//
// Solidity: function getVaultOwners() view returns(address[])
func (_OToken *OTokenCaller) GetVaultOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "getVaultOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetVaultOwners is a free data retrieval call binding the contract method 0x060ab2ea.
//
// Solidity: function getVaultOwners() view returns(address[])
func (_OToken *OTokenSession) GetVaultOwners() ([]common.Address, error) {
	return _OToken.Contract.GetVaultOwners(&_OToken.CallOpts)
}

// GetVaultOwners is a free data retrieval call binding the contract method 0x060ab2ea.
//
// Solidity: function getVaultOwners() view returns(address[])
func (_OToken *OTokenCallerSession) GetVaultOwners() ([]common.Address, error) {
	return _OToken.Contract.GetVaultOwners(&_OToken.CallOpts)
}

// HasExpired is a free data retrieval call binding the contract method 0x90e64d13.
//
// Solidity: function hasExpired() view returns(bool)
func (_OToken *OTokenCaller) HasExpired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "hasExpired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasExpired is a free data retrieval call binding the contract method 0x90e64d13.
//
// Solidity: function hasExpired() view returns(bool)
func (_OToken *OTokenSession) HasExpired() (bool, error) {
	return _OToken.Contract.HasExpired(&_OToken.CallOpts)
}

// HasExpired is a free data retrieval call binding the contract method 0x90e64d13.
//
// Solidity: function hasExpired() view returns(bool)
func (_OToken *OTokenCallerSession) HasExpired() (bool, error) {
	return _OToken.Contract.HasExpired(&_OToken.CallOpts)
}

// HasVault is a free data retrieval call binding the contract method 0x0d453efb.
//
// Solidity: function hasVault(address owner) view returns(bool)
func (_OToken *OTokenCaller) HasVault(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "hasVault", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVault is a free data retrieval call binding the contract method 0x0d453efb.
//
// Solidity: function hasVault(address owner) view returns(bool)
func (_OToken *OTokenSession) HasVault(owner common.Address) (bool, error) {
	return _OToken.Contract.HasVault(&_OToken.CallOpts, owner)
}

// HasVault is a free data retrieval call binding the contract method 0x0d453efb.
//
// Solidity: function hasVault(address owner) view returns(bool)
func (_OToken *OTokenCallerSession) HasVault(owner common.Address) (bool, error) {
	return _OToken.Contract.HasVault(&_OToken.CallOpts, owner)
}

// IsETH is a free data retrieval call binding the contract method 0xf70a2508.
//
// Solidity: function isETH(address _ierc20) pure returns(bool)
func (_OToken *OTokenCaller) IsETH(opts *bind.CallOpts, _ierc20 common.Address) (bool, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "isETH", _ierc20)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsETH is a free data retrieval call binding the contract method 0xf70a2508.
//
// Solidity: function isETH(address _ierc20) pure returns(bool)
func (_OToken *OTokenSession) IsETH(_ierc20 common.Address) (bool, error) {
	return _OToken.Contract.IsETH(&_OToken.CallOpts, _ierc20)
}

// IsETH is a free data retrieval call binding the contract method 0xf70a2508.
//
// Solidity: function isETH(address _ierc20) pure returns(bool)
func (_OToken *OTokenCallerSession) IsETH(_ierc20 common.Address) (bool, error) {
	return _OToken.Contract.IsETH(&_OToken.CallOpts, _ierc20)
}

// IsExerciseWindow is a free data retrieval call binding the contract method 0x0d6cd8aa.
//
// Solidity: function isExerciseWindow() view returns(bool)
func (_OToken *OTokenCaller) IsExerciseWindow(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "isExerciseWindow")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExerciseWindow is a free data retrieval call binding the contract method 0x0d6cd8aa.
//
// Solidity: function isExerciseWindow() view returns(bool)
func (_OToken *OTokenSession) IsExerciseWindow() (bool, error) {
	return _OToken.Contract.IsExerciseWindow(&_OToken.CallOpts)
}

// IsExerciseWindow is a free data retrieval call binding the contract method 0x0d6cd8aa.
//
// Solidity: function isExerciseWindow() view returns(bool)
func (_OToken *OTokenCallerSession) IsExerciseWindow() (bool, error) {
	return _OToken.Contract.IsExerciseWindow(&_OToken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_OToken *OTokenCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_OToken *OTokenSession) IsOwner() (bool, error) {
	return _OToken.Contract.IsOwner(&_OToken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_OToken *OTokenCallerSession) IsOwner() (bool, error) {
	return _OToken.Contract.IsOwner(&_OToken.CallOpts)
}

// IsUnsafe is a free data retrieval call binding the contract method 0xcdb4b5c2.
//
// Solidity: function isUnsafe(address vaultOwner) view returns(bool)
func (_OToken *OTokenCaller) IsUnsafe(opts *bind.CallOpts, vaultOwner common.Address) (bool, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "isUnsafe", vaultOwner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUnsafe is a free data retrieval call binding the contract method 0xcdb4b5c2.
//
// Solidity: function isUnsafe(address vaultOwner) view returns(bool)
func (_OToken *OTokenSession) IsUnsafe(vaultOwner common.Address) (bool, error) {
	return _OToken.Contract.IsUnsafe(&_OToken.CallOpts, vaultOwner)
}

// IsUnsafe is a free data retrieval call binding the contract method 0xcdb4b5c2.
//
// Solidity: function isUnsafe(address vaultOwner) view returns(bool)
func (_OToken *OTokenCallerSession) IsUnsafe(vaultOwner common.Address) (bool, error) {
	return _OToken.Contract.IsUnsafe(&_OToken.CallOpts, vaultOwner)
}

// LiquidationFactor is a free data retrieval call binding the contract method 0x352ade55.
//
// Solidity: function liquidationFactor() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenCaller) LiquidationFactor(opts *bind.CallOpts) (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "liquidationFactor")

	outstruct := new(struct {
		Value    *big.Int
		Exponent int32
	})

	outstruct.Value = out[0].(*big.Int)
	outstruct.Exponent = out[1].(int32)

	return *outstruct, err

}

// LiquidationFactor is a free data retrieval call binding the contract method 0x352ade55.
//
// Solidity: function liquidationFactor() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenSession) LiquidationFactor() (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	return _OToken.Contract.LiquidationFactor(&_OToken.CallOpts)
}

// LiquidationFactor is a free data retrieval call binding the contract method 0x352ade55.
//
// Solidity: function liquidationFactor() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenCallerSession) LiquidationFactor() (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	return _OToken.Contract.LiquidationFactor(&_OToken.CallOpts)
}

// LiquidationIncentive is a free data retrieval call binding the contract method 0x8c765e94.
//
// Solidity: function liquidationIncentive() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenCaller) LiquidationIncentive(opts *bind.CallOpts) (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "liquidationIncentive")

	outstruct := new(struct {
		Value    *big.Int
		Exponent int32
	})

	outstruct.Value = out[0].(*big.Int)
	outstruct.Exponent = out[1].(int32)

	return *outstruct, err

}

// LiquidationIncentive is a free data retrieval call binding the contract method 0x8c765e94.
//
// Solidity: function liquidationIncentive() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenSession) LiquidationIncentive() (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	return _OToken.Contract.LiquidationIncentive(&_OToken.CallOpts)
}

// LiquidationIncentive is a free data retrieval call binding the contract method 0x8c765e94.
//
// Solidity: function liquidationIncentive() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenCallerSession) LiquidationIncentive() (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	return _OToken.Contract.LiquidationIncentive(&_OToken.CallOpts)
}

// MaxOTokensIssuable is a free data retrieval call binding the contract method 0x686c1e21.
//
// Solidity: function maxOTokensIssuable(uint256 collateralAmt) view returns(uint256)
func (_OToken *OTokenCaller) MaxOTokensIssuable(opts *bind.CallOpts, collateralAmt *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "maxOTokensIssuable", collateralAmt)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxOTokensIssuable is a free data retrieval call binding the contract method 0x686c1e21.
//
// Solidity: function maxOTokensIssuable(uint256 collateralAmt) view returns(uint256)
func (_OToken *OTokenSession) MaxOTokensIssuable(collateralAmt *big.Int) (*big.Int, error) {
	return _OToken.Contract.MaxOTokensIssuable(&_OToken.CallOpts, collateralAmt)
}

// MaxOTokensIssuable is a free data retrieval call binding the contract method 0x686c1e21.
//
// Solidity: function maxOTokensIssuable(uint256 collateralAmt) view returns(uint256)
func (_OToken *OTokenCallerSession) MaxOTokensIssuable(collateralAmt *big.Int) (*big.Int, error) {
	return _OToken.Contract.MaxOTokensIssuable(&_OToken.CallOpts, collateralAmt)
}

// MaxOTokensLiquidatable is a free data retrieval call binding the contract method 0xdec44c0b.
//
// Solidity: function maxOTokensLiquidatable(address vaultOwner) view returns(uint256)
func (_OToken *OTokenCaller) MaxOTokensLiquidatable(opts *bind.CallOpts, vaultOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "maxOTokensLiquidatable", vaultOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxOTokensLiquidatable is a free data retrieval call binding the contract method 0xdec44c0b.
//
// Solidity: function maxOTokensLiquidatable(address vaultOwner) view returns(uint256)
func (_OToken *OTokenSession) MaxOTokensLiquidatable(vaultOwner common.Address) (*big.Int, error) {
	return _OToken.Contract.MaxOTokensLiquidatable(&_OToken.CallOpts, vaultOwner)
}

// MaxOTokensLiquidatable is a free data retrieval call binding the contract method 0xdec44c0b.
//
// Solidity: function maxOTokensLiquidatable(address vaultOwner) view returns(uint256)
func (_OToken *OTokenCallerSession) MaxOTokensLiquidatable(vaultOwner common.Address) (*big.Int, error) {
	return _OToken.Contract.MaxOTokensLiquidatable(&_OToken.CallOpts, vaultOwner)
}

// MinCollateralizationRatio is a free data retrieval call binding the contract method 0xba1be554.
//
// Solidity: function minCollateralizationRatio() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenCaller) MinCollateralizationRatio(opts *bind.CallOpts) (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "minCollateralizationRatio")

	outstruct := new(struct {
		Value    *big.Int
		Exponent int32
	})

	outstruct.Value = out[0].(*big.Int)
	outstruct.Exponent = out[1].(int32)

	return *outstruct, err

}

// MinCollateralizationRatio is a free data retrieval call binding the contract method 0xba1be554.
//
// Solidity: function minCollateralizationRatio() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenSession) MinCollateralizationRatio() (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	return _OToken.Contract.MinCollateralizationRatio(&_OToken.CallOpts)
}

// MinCollateralizationRatio is a free data retrieval call binding the contract method 0xba1be554.
//
// Solidity: function minCollateralizationRatio() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenCallerSession) MinCollateralizationRatio() (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	return _OToken.Contract.MinCollateralizationRatio(&_OToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OToken *OTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OToken *OTokenSession) Name() (string, error) {
	return _OToken.Contract.Name(&_OToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OToken *OTokenCallerSession) Name() (string, error) {
	return _OToken.Contract.Name(&_OToken.CallOpts)
}

// OTokenExchangeRate is a free data retrieval call binding the contract method 0xb7365540.
//
// Solidity: function oTokenExchangeRate() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenCaller) OTokenExchangeRate(opts *bind.CallOpts) (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "oTokenExchangeRate")

	outstruct := new(struct {
		Value    *big.Int
		Exponent int32
	})

	outstruct.Value = out[0].(*big.Int)
	outstruct.Exponent = out[1].(int32)

	return *outstruct, err

}

// OTokenExchangeRate is a free data retrieval call binding the contract method 0xb7365540.
//
// Solidity: function oTokenExchangeRate() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenSession) OTokenExchangeRate() (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	return _OToken.Contract.OTokenExchangeRate(&_OToken.CallOpts)
}

// OTokenExchangeRate is a free data retrieval call binding the contract method 0xb7365540.
//
// Solidity: function oTokenExchangeRate() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenCallerSession) OTokenExchangeRate() (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	return _OToken.Contract.OTokenExchangeRate(&_OToken.CallOpts)
}

// OptionsExchange is a free data retrieval call binding the contract method 0x3bd33f62.
//
// Solidity: function optionsExchange() view returns(address)
func (_OToken *OTokenCaller) OptionsExchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "optionsExchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptionsExchange is a free data retrieval call binding the contract method 0x3bd33f62.
//
// Solidity: function optionsExchange() view returns(address)
func (_OToken *OTokenSession) OptionsExchange() (common.Address, error) {
	return _OToken.Contract.OptionsExchange(&_OToken.CallOpts)
}

// OptionsExchange is a free data retrieval call binding the contract method 0x3bd33f62.
//
// Solidity: function optionsExchange() view returns(address)
func (_OToken *OTokenCallerSession) OptionsExchange() (common.Address, error) {
	return _OToken.Contract.OptionsExchange(&_OToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OToken *OTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OToken *OTokenSession) Owner() (common.Address, error) {
	return _OToken.Contract.Owner(&_OToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OToken *OTokenCallerSession) Owner() (common.Address, error) {
	return _OToken.Contract.Owner(&_OToken.CallOpts)
}

// Strike is a free data retrieval call binding the contract method 0xad8f5008.
//
// Solidity: function strike() view returns(address)
func (_OToken *OTokenCaller) Strike(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "strike")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Strike is a free data retrieval call binding the contract method 0xad8f5008.
//
// Solidity: function strike() view returns(address)
func (_OToken *OTokenSession) Strike() (common.Address, error) {
	return _OToken.Contract.Strike(&_OToken.CallOpts)
}

// Strike is a free data retrieval call binding the contract method 0xad8f5008.
//
// Solidity: function strike() view returns(address)
func (_OToken *OTokenCallerSession) Strike() (common.Address, error) {
	return _OToken.Contract.Strike(&_OToken.CallOpts)
}

// StrikePrice is a free data retrieval call binding the contract method 0xc52987cf.
//
// Solidity: function strikePrice() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenCaller) StrikePrice(opts *bind.CallOpts) (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "strikePrice")

	outstruct := new(struct {
		Value    *big.Int
		Exponent int32
	})

	outstruct.Value = out[0].(*big.Int)
	outstruct.Exponent = out[1].(int32)

	return *outstruct, err

}

// StrikePrice is a free data retrieval call binding the contract method 0xc52987cf.
//
// Solidity: function strikePrice() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenSession) StrikePrice() (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	return _OToken.Contract.StrikePrice(&_OToken.CallOpts)
}

// StrikePrice is a free data retrieval call binding the contract method 0xc52987cf.
//
// Solidity: function strikePrice() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenCallerSession) StrikePrice() (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	return _OToken.Contract.StrikePrice(&_OToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OToken *OTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OToken *OTokenSession) Symbol() (string, error) {
	return _OToken.Contract.Symbol(&_OToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OToken *OTokenCallerSession) Symbol() (string, error) {
	return _OToken.Contract.Symbol(&_OToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OToken *OTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OToken *OTokenSession) TotalSupply() (*big.Int, error) {
	return _OToken.Contract.TotalSupply(&_OToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OToken *OTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _OToken.Contract.TotalSupply(&_OToken.CallOpts)
}

// TransactionFee is a free data retrieval call binding the contract method 0x9ed3edf0.
//
// Solidity: function transactionFee() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenCaller) TransactionFee(opts *bind.CallOpts) (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "transactionFee")

	outstruct := new(struct {
		Value    *big.Int
		Exponent int32
	})

	outstruct.Value = out[0].(*big.Int)
	outstruct.Exponent = out[1].(int32)

	return *outstruct, err

}

// TransactionFee is a free data retrieval call binding the contract method 0x9ed3edf0.
//
// Solidity: function transactionFee() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenSession) TransactionFee() (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	return _OToken.Contract.TransactionFee(&_OToken.CallOpts)
}

// TransactionFee is a free data retrieval call binding the contract method 0x9ed3edf0.
//
// Solidity: function transactionFee() view returns(uint256 value, int32 exponent)
func (_OToken *OTokenCallerSession) TransactionFee() (struct {
	Value    *big.Int
	Exponent int32
}, error) {
	return _OToken.Contract.TransactionFee(&_OToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_OToken *OTokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_OToken *OTokenSession) Underlying() (common.Address, error) {
	return _OToken.Contract.Underlying(&_OToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_OToken *OTokenCallerSession) Underlying() (common.Address, error) {
	return _OToken.Contract.Underlying(&_OToken.CallOpts)
}

// UnderlyingExp is a free data retrieval call binding the contract method 0xb2c2b13f.
//
// Solidity: function underlyingExp() view returns(int32)
func (_OToken *OTokenCaller) UnderlyingExp(opts *bind.CallOpts) (int32, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "underlyingExp")

	if err != nil {
		return *new(int32), err
	}

	out0 := *abi.ConvertType(out[0], new(int32)).(*int32)

	return out0, err

}

// UnderlyingExp is a free data retrieval call binding the contract method 0xb2c2b13f.
//
// Solidity: function underlyingExp() view returns(int32)
func (_OToken *OTokenSession) UnderlyingExp() (int32, error) {
	return _OToken.Contract.UnderlyingExp(&_OToken.CallOpts)
}

// UnderlyingExp is a free data retrieval call binding the contract method 0xb2c2b13f.
//
// Solidity: function underlyingExp() view returns(int32)
func (_OToken *OTokenCallerSession) UnderlyingExp() (int32, error) {
	return _OToken.Contract.UnderlyingExp(&_OToken.CallOpts)
}

// UnderlyingRequiredToExercise is a free data retrieval call binding the contract method 0x6fd865f9.
//
// Solidity: function underlyingRequiredToExercise(uint256 oTokensToExercise) view returns(uint256)
func (_OToken *OTokenCaller) UnderlyingRequiredToExercise(opts *bind.CallOpts, oTokensToExercise *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OToken.contract.Call(opts, &out, "underlyingRequiredToExercise", oTokensToExercise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingRequiredToExercise is a free data retrieval call binding the contract method 0x6fd865f9.
//
// Solidity: function underlyingRequiredToExercise(uint256 oTokensToExercise) view returns(uint256)
func (_OToken *OTokenSession) UnderlyingRequiredToExercise(oTokensToExercise *big.Int) (*big.Int, error) {
	return _OToken.Contract.UnderlyingRequiredToExercise(&_OToken.CallOpts, oTokensToExercise)
}

// UnderlyingRequiredToExercise is a free data retrieval call binding the contract method 0x6fd865f9.
//
// Solidity: function underlyingRequiredToExercise(uint256 oTokensToExercise) view returns(uint256)
func (_OToken *OTokenCallerSession) UnderlyingRequiredToExercise(oTokensToExercise *big.Int) (*big.Int, error) {
	return _OToken.Contract.UnderlyingRequiredToExercise(&_OToken.CallOpts, oTokensToExercise)
}

// AddAndSellERC20CollateralOption is a paid mutator transaction binding the contract method 0x3226052d.
//
// Solidity: function addAndSellERC20CollateralOption(uint256 amtToCreate, uint256 amtCollateral, address receiver) returns()
func (_OToken *OTokenTransactor) AddAndSellERC20CollateralOption(opts *bind.TransactOpts, amtToCreate *big.Int, amtCollateral *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "addAndSellERC20CollateralOption", amtToCreate, amtCollateral, receiver)
}

// AddAndSellERC20CollateralOption is a paid mutator transaction binding the contract method 0x3226052d.
//
// Solidity: function addAndSellERC20CollateralOption(uint256 amtToCreate, uint256 amtCollateral, address receiver) returns()
func (_OToken *OTokenSession) AddAndSellERC20CollateralOption(amtToCreate *big.Int, amtCollateral *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.AddAndSellERC20CollateralOption(&_OToken.TransactOpts, amtToCreate, amtCollateral, receiver)
}

// AddAndSellERC20CollateralOption is a paid mutator transaction binding the contract method 0x3226052d.
//
// Solidity: function addAndSellERC20CollateralOption(uint256 amtToCreate, uint256 amtCollateral, address receiver) returns()
func (_OToken *OTokenTransactorSession) AddAndSellERC20CollateralOption(amtToCreate *big.Int, amtCollateral *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.AddAndSellERC20CollateralOption(&_OToken.TransactOpts, amtToCreate, amtCollateral, receiver)
}

// AddAndSellETHCollateralOption is a paid mutator transaction binding the contract method 0xcfbea789.
//
// Solidity: function addAndSellETHCollateralOption(uint256 amtToCreate, address receiver) payable returns()
func (_OToken *OTokenTransactor) AddAndSellETHCollateralOption(opts *bind.TransactOpts, amtToCreate *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "addAndSellETHCollateralOption", amtToCreate, receiver)
}

// AddAndSellETHCollateralOption is a paid mutator transaction binding the contract method 0xcfbea789.
//
// Solidity: function addAndSellETHCollateralOption(uint256 amtToCreate, address receiver) payable returns()
func (_OToken *OTokenSession) AddAndSellETHCollateralOption(amtToCreate *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.AddAndSellETHCollateralOption(&_OToken.TransactOpts, amtToCreate, receiver)
}

// AddAndSellETHCollateralOption is a paid mutator transaction binding the contract method 0xcfbea789.
//
// Solidity: function addAndSellETHCollateralOption(uint256 amtToCreate, address receiver) payable returns()
func (_OToken *OTokenTransactorSession) AddAndSellETHCollateralOption(amtToCreate *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.AddAndSellETHCollateralOption(&_OToken.TransactOpts, amtToCreate, receiver)
}

// AddERC20Collateral is a paid mutator transaction binding the contract method 0x01b4a3c1.
//
// Solidity: function addERC20Collateral(address vaultOwner, uint256 amt) returns(uint256)
func (_OToken *OTokenTransactor) AddERC20Collateral(opts *bind.TransactOpts, vaultOwner common.Address, amt *big.Int) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "addERC20Collateral", vaultOwner, amt)
}

// AddERC20Collateral is a paid mutator transaction binding the contract method 0x01b4a3c1.
//
// Solidity: function addERC20Collateral(address vaultOwner, uint256 amt) returns(uint256)
func (_OToken *OTokenSession) AddERC20Collateral(vaultOwner common.Address, amt *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.AddERC20Collateral(&_OToken.TransactOpts, vaultOwner, amt)
}

// AddERC20Collateral is a paid mutator transaction binding the contract method 0x01b4a3c1.
//
// Solidity: function addERC20Collateral(address vaultOwner, uint256 amt) returns(uint256)
func (_OToken *OTokenTransactorSession) AddERC20Collateral(vaultOwner common.Address, amt *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.AddERC20Collateral(&_OToken.TransactOpts, vaultOwner, amt)
}

// AddERC20CollateralOption is a paid mutator transaction binding the contract method 0x5ca7c8a6.
//
// Solidity: function addERC20CollateralOption(uint256 amtToCreate, uint256 amtCollateral, address receiver) returns()
func (_OToken *OTokenTransactor) AddERC20CollateralOption(opts *bind.TransactOpts, amtToCreate *big.Int, amtCollateral *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "addERC20CollateralOption", amtToCreate, amtCollateral, receiver)
}

// AddERC20CollateralOption is a paid mutator transaction binding the contract method 0x5ca7c8a6.
//
// Solidity: function addERC20CollateralOption(uint256 amtToCreate, uint256 amtCollateral, address receiver) returns()
func (_OToken *OTokenSession) AddERC20CollateralOption(amtToCreate *big.Int, amtCollateral *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.AddERC20CollateralOption(&_OToken.TransactOpts, amtToCreate, amtCollateral, receiver)
}

// AddERC20CollateralOption is a paid mutator transaction binding the contract method 0x5ca7c8a6.
//
// Solidity: function addERC20CollateralOption(uint256 amtToCreate, uint256 amtCollateral, address receiver) returns()
func (_OToken *OTokenTransactorSession) AddERC20CollateralOption(amtToCreate *big.Int, amtCollateral *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.AddERC20CollateralOption(&_OToken.TransactOpts, amtToCreate, amtCollateral, receiver)
}

// AddETHCollateral is a paid mutator transaction binding the contract method 0x9ce07251.
//
// Solidity: function addETHCollateral(address vaultOwner) payable returns(uint256)
func (_OToken *OTokenTransactor) AddETHCollateral(opts *bind.TransactOpts, vaultOwner common.Address) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "addETHCollateral", vaultOwner)
}

// AddETHCollateral is a paid mutator transaction binding the contract method 0x9ce07251.
//
// Solidity: function addETHCollateral(address vaultOwner) payable returns(uint256)
func (_OToken *OTokenSession) AddETHCollateral(vaultOwner common.Address) (*types.Transaction, error) {
	return _OToken.Contract.AddETHCollateral(&_OToken.TransactOpts, vaultOwner)
}

// AddETHCollateral is a paid mutator transaction binding the contract method 0x9ce07251.
//
// Solidity: function addETHCollateral(address vaultOwner) payable returns(uint256)
func (_OToken *OTokenTransactorSession) AddETHCollateral(vaultOwner common.Address) (*types.Transaction, error) {
	return _OToken.Contract.AddETHCollateral(&_OToken.TransactOpts, vaultOwner)
}

// AddETHCollateralOption is a paid mutator transaction binding the contract method 0xb96661ba.
//
// Solidity: function addETHCollateralOption(uint256 amtToCreate, address receiver) payable returns()
func (_OToken *OTokenTransactor) AddETHCollateralOption(opts *bind.TransactOpts, amtToCreate *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "addETHCollateralOption", amtToCreate, receiver)
}

// AddETHCollateralOption is a paid mutator transaction binding the contract method 0xb96661ba.
//
// Solidity: function addETHCollateralOption(uint256 amtToCreate, address receiver) payable returns()
func (_OToken *OTokenSession) AddETHCollateralOption(amtToCreate *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.AddETHCollateralOption(&_OToken.TransactOpts, amtToCreate, receiver)
}

// AddETHCollateralOption is a paid mutator transaction binding the contract method 0xb96661ba.
//
// Solidity: function addETHCollateralOption(uint256 amtToCreate, address receiver) payable returns()
func (_OToken *OTokenTransactorSession) AddETHCollateralOption(amtToCreate *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.AddETHCollateralOption(&_OToken.TransactOpts, amtToCreate, receiver)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_OToken *OTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_OToken *OTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.Approve(&_OToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_OToken *OTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.Approve(&_OToken.TransactOpts, spender, amount)
}

// BurnOTokens is a paid mutator transaction binding the contract method 0xeaa376b5.
//
// Solidity: function burnOTokens(uint256 amtToBurn) returns()
func (_OToken *OTokenTransactor) BurnOTokens(opts *bind.TransactOpts, amtToBurn *big.Int) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "burnOTokens", amtToBurn)
}

// BurnOTokens is a paid mutator transaction binding the contract method 0xeaa376b5.
//
// Solidity: function burnOTokens(uint256 amtToBurn) returns()
func (_OToken *OTokenSession) BurnOTokens(amtToBurn *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.BurnOTokens(&_OToken.TransactOpts, amtToBurn)
}

// BurnOTokens is a paid mutator transaction binding the contract method 0xeaa376b5.
//
// Solidity: function burnOTokens(uint256 amtToBurn) returns()
func (_OToken *OTokenTransactorSession) BurnOTokens(amtToBurn *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.BurnOTokens(&_OToken.TransactOpts, amtToBurn)
}

// CreateAndSellERC20CollateralOption is a paid mutator transaction binding the contract method 0xc56749ce.
//
// Solidity: function createAndSellERC20CollateralOption(uint256 amtToCreate, uint256 amtCollateral, address receiver) returns()
func (_OToken *OTokenTransactor) CreateAndSellERC20CollateralOption(opts *bind.TransactOpts, amtToCreate *big.Int, amtCollateral *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "createAndSellERC20CollateralOption", amtToCreate, amtCollateral, receiver)
}

// CreateAndSellERC20CollateralOption is a paid mutator transaction binding the contract method 0xc56749ce.
//
// Solidity: function createAndSellERC20CollateralOption(uint256 amtToCreate, uint256 amtCollateral, address receiver) returns()
func (_OToken *OTokenSession) CreateAndSellERC20CollateralOption(amtToCreate *big.Int, amtCollateral *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.CreateAndSellERC20CollateralOption(&_OToken.TransactOpts, amtToCreate, amtCollateral, receiver)
}

// CreateAndSellERC20CollateralOption is a paid mutator transaction binding the contract method 0xc56749ce.
//
// Solidity: function createAndSellERC20CollateralOption(uint256 amtToCreate, uint256 amtCollateral, address receiver) returns()
func (_OToken *OTokenTransactorSession) CreateAndSellERC20CollateralOption(amtToCreate *big.Int, amtCollateral *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.CreateAndSellERC20CollateralOption(&_OToken.TransactOpts, amtToCreate, amtCollateral, receiver)
}

// CreateAndSellETHCollateralOption is a paid mutator transaction binding the contract method 0x3667429f.
//
// Solidity: function createAndSellETHCollateralOption(uint256 amtToCreate, address receiver) payable returns()
func (_OToken *OTokenTransactor) CreateAndSellETHCollateralOption(opts *bind.TransactOpts, amtToCreate *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "createAndSellETHCollateralOption", amtToCreate, receiver)
}

// CreateAndSellETHCollateralOption is a paid mutator transaction binding the contract method 0x3667429f.
//
// Solidity: function createAndSellETHCollateralOption(uint256 amtToCreate, address receiver) payable returns()
func (_OToken *OTokenSession) CreateAndSellETHCollateralOption(amtToCreate *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.CreateAndSellETHCollateralOption(&_OToken.TransactOpts, amtToCreate, receiver)
}

// CreateAndSellETHCollateralOption is a paid mutator transaction binding the contract method 0x3667429f.
//
// Solidity: function createAndSellETHCollateralOption(uint256 amtToCreate, address receiver) payable returns()
func (_OToken *OTokenTransactorSession) CreateAndSellETHCollateralOption(amtToCreate *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.CreateAndSellETHCollateralOption(&_OToken.TransactOpts, amtToCreate, receiver)
}

// CreateERC20CollateralOption is a paid mutator transaction binding the contract method 0x52f89fe3.
//
// Solidity: function createERC20CollateralOption(uint256 amtToCreate, uint256 amtCollateral, address receiver) returns()
func (_OToken *OTokenTransactor) CreateERC20CollateralOption(opts *bind.TransactOpts, amtToCreate *big.Int, amtCollateral *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "createERC20CollateralOption", amtToCreate, amtCollateral, receiver)
}

// CreateERC20CollateralOption is a paid mutator transaction binding the contract method 0x52f89fe3.
//
// Solidity: function createERC20CollateralOption(uint256 amtToCreate, uint256 amtCollateral, address receiver) returns()
func (_OToken *OTokenSession) CreateERC20CollateralOption(amtToCreate *big.Int, amtCollateral *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.CreateERC20CollateralOption(&_OToken.TransactOpts, amtToCreate, amtCollateral, receiver)
}

// CreateERC20CollateralOption is a paid mutator transaction binding the contract method 0x52f89fe3.
//
// Solidity: function createERC20CollateralOption(uint256 amtToCreate, uint256 amtCollateral, address receiver) returns()
func (_OToken *OTokenTransactorSession) CreateERC20CollateralOption(amtToCreate *big.Int, amtCollateral *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.CreateERC20CollateralOption(&_OToken.TransactOpts, amtToCreate, amtCollateral, receiver)
}

// CreateETHCollateralOption is a paid mutator transaction binding the contract method 0xed1f41c3.
//
// Solidity: function createETHCollateralOption(uint256 amtToCreate, address receiver) payable returns()
func (_OToken *OTokenTransactor) CreateETHCollateralOption(opts *bind.TransactOpts, amtToCreate *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "createETHCollateralOption", amtToCreate, receiver)
}

// CreateETHCollateralOption is a paid mutator transaction binding the contract method 0xed1f41c3.
//
// Solidity: function createETHCollateralOption(uint256 amtToCreate, address receiver) payable returns()
func (_OToken *OTokenSession) CreateETHCollateralOption(amtToCreate *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.CreateETHCollateralOption(&_OToken.TransactOpts, amtToCreate, receiver)
}

// CreateETHCollateralOption is a paid mutator transaction binding the contract method 0xed1f41c3.
//
// Solidity: function createETHCollateralOption(uint256 amtToCreate, address receiver) payable returns()
func (_OToken *OTokenTransactorSession) CreateETHCollateralOption(amtToCreate *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.CreateETHCollateralOption(&_OToken.TransactOpts, amtToCreate, receiver)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_OToken *OTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_OToken *OTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.DecreaseAllowance(&_OToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_OToken *OTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.DecreaseAllowance(&_OToken.TransactOpts, spender, subtractedValue)
}

// Exercise is a paid mutator transaction binding the contract method 0x58b36dac.
//
// Solidity: function exercise(uint256 oTokensToExercise, address[] vaultsToExerciseFrom) payable returns()
func (_OToken *OTokenTransactor) Exercise(opts *bind.TransactOpts, oTokensToExercise *big.Int, vaultsToExerciseFrom []common.Address) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "exercise", oTokensToExercise, vaultsToExerciseFrom)
}

// Exercise is a paid mutator transaction binding the contract method 0x58b36dac.
//
// Solidity: function exercise(uint256 oTokensToExercise, address[] vaultsToExerciseFrom) payable returns()
func (_OToken *OTokenSession) Exercise(oTokensToExercise *big.Int, vaultsToExerciseFrom []common.Address) (*types.Transaction, error) {
	return _OToken.Contract.Exercise(&_OToken.TransactOpts, oTokensToExercise, vaultsToExerciseFrom)
}

// Exercise is a paid mutator transaction binding the contract method 0x58b36dac.
//
// Solidity: function exercise(uint256 oTokensToExercise, address[] vaultsToExerciseFrom) payable returns()
func (_OToken *OTokenTransactorSession) Exercise(oTokensToExercise *big.Int, vaultsToExerciseFrom []common.Address) (*types.Transaction, error) {
	return _OToken.Contract.Exercise(&_OToken.TransactOpts, oTokensToExercise, vaultsToExerciseFrom)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_OToken *OTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_OToken *OTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.IncreaseAllowance(&_OToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_OToken *OTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.IncreaseAllowance(&_OToken.TransactOpts, spender, addedValue)
}

// IssueOTokens is a paid mutator transaction binding the contract method 0x1a0e21bd.
//
// Solidity: function issueOTokens(uint256 oTokensToIssue, address receiver) returns()
func (_OToken *OTokenTransactor) IssueOTokens(opts *bind.TransactOpts, oTokensToIssue *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "issueOTokens", oTokensToIssue, receiver)
}

// IssueOTokens is a paid mutator transaction binding the contract method 0x1a0e21bd.
//
// Solidity: function issueOTokens(uint256 oTokensToIssue, address receiver) returns()
func (_OToken *OTokenSession) IssueOTokens(oTokensToIssue *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.IssueOTokens(&_OToken.TransactOpts, oTokensToIssue, receiver)
}

// IssueOTokens is a paid mutator transaction binding the contract method 0x1a0e21bd.
//
// Solidity: function issueOTokens(uint256 oTokensToIssue, address receiver) returns()
func (_OToken *OTokenTransactorSession) IssueOTokens(oTokensToIssue *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OToken.Contract.IssueOTokens(&_OToken.TransactOpts, oTokensToIssue, receiver)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbcbaf487.
//
// Solidity: function liquidate(address vaultOwner, uint256 oTokensToLiquidate) returns()
func (_OToken *OTokenTransactor) Liquidate(opts *bind.TransactOpts, vaultOwner common.Address, oTokensToLiquidate *big.Int) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "liquidate", vaultOwner, oTokensToLiquidate)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbcbaf487.
//
// Solidity: function liquidate(address vaultOwner, uint256 oTokensToLiquidate) returns()
func (_OToken *OTokenSession) Liquidate(vaultOwner common.Address, oTokensToLiquidate *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.Liquidate(&_OToken.TransactOpts, vaultOwner, oTokensToLiquidate)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbcbaf487.
//
// Solidity: function liquidate(address vaultOwner, uint256 oTokensToLiquidate) returns()
func (_OToken *OTokenTransactorSession) Liquidate(vaultOwner common.Address, oTokensToLiquidate *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.Liquidate(&_OToken.TransactOpts, vaultOwner, oTokensToLiquidate)
}

// OpenVault is a paid mutator transaction binding the contract method 0x86f54712.
//
// Solidity: function openVault() returns(bool)
func (_OToken *OTokenTransactor) OpenVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "openVault")
}

// OpenVault is a paid mutator transaction binding the contract method 0x86f54712.
//
// Solidity: function openVault() returns(bool)
func (_OToken *OTokenSession) OpenVault() (*types.Transaction, error) {
	return _OToken.Contract.OpenVault(&_OToken.TransactOpts)
}

// OpenVault is a paid mutator transaction binding the contract method 0x86f54712.
//
// Solidity: function openVault() returns(bool)
func (_OToken *OTokenTransactorSession) OpenVault() (*types.Transaction, error) {
	return _OToken.Contract.OpenVault(&_OToken.TransactOpts)
}

// RedeemVaultBalance is a paid mutator transaction binding the contract method 0xb76fdb6c.
//
// Solidity: function redeemVaultBalance() returns()
func (_OToken *OTokenTransactor) RedeemVaultBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "redeemVaultBalance")
}

// RedeemVaultBalance is a paid mutator transaction binding the contract method 0xb76fdb6c.
//
// Solidity: function redeemVaultBalance() returns()
func (_OToken *OTokenSession) RedeemVaultBalance() (*types.Transaction, error) {
	return _OToken.Contract.RedeemVaultBalance(&_OToken.TransactOpts)
}

// RedeemVaultBalance is a paid mutator transaction binding the contract method 0xb76fdb6c.
//
// Solidity: function redeemVaultBalance() returns()
func (_OToken *OTokenTransactorSession) RedeemVaultBalance() (*types.Transaction, error) {
	return _OToken.Contract.RedeemVaultBalance(&_OToken.TransactOpts)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0x3237c158.
//
// Solidity: function removeCollateral(uint256 amtToRemove) returns()
func (_OToken *OTokenTransactor) RemoveCollateral(opts *bind.TransactOpts, amtToRemove *big.Int) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "removeCollateral", amtToRemove)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0x3237c158.
//
// Solidity: function removeCollateral(uint256 amtToRemove) returns()
func (_OToken *OTokenSession) RemoveCollateral(amtToRemove *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.RemoveCollateral(&_OToken.TransactOpts, amtToRemove)
}

// RemoveCollateral is a paid mutator transaction binding the contract method 0x3237c158.
//
// Solidity: function removeCollateral(uint256 amtToRemove) returns()
func (_OToken *OTokenTransactorSession) RemoveCollateral(amtToRemove *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.RemoveCollateral(&_OToken.TransactOpts, amtToRemove)
}

// RemoveUnderlying is a paid mutator transaction binding the contract method 0xfaa2041f.
//
// Solidity: function removeUnderlying() returns()
func (_OToken *OTokenTransactor) RemoveUnderlying(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "removeUnderlying")
}

// RemoveUnderlying is a paid mutator transaction binding the contract method 0xfaa2041f.
//
// Solidity: function removeUnderlying() returns()
func (_OToken *OTokenSession) RemoveUnderlying() (*types.Transaction, error) {
	return _OToken.Contract.RemoveUnderlying(&_OToken.TransactOpts)
}

// RemoveUnderlying is a paid mutator transaction binding the contract method 0xfaa2041f.
//
// Solidity: function removeUnderlying() returns()
func (_OToken *OTokenTransactorSession) RemoveUnderlying() (*types.Transaction, error) {
	return _OToken.Contract.RemoveUnderlying(&_OToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OToken *OTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OToken *OTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _OToken.Contract.RenounceOwnership(&_OToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OToken *OTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OToken.Contract.RenounceOwnership(&_OToken.TransactOpts)
}

// SetDetails is a paid mutator transaction binding the contract method 0xb7b090ee.
//
// Solidity: function setDetails(string _name, string _symbol) returns()
func (_OToken *OTokenTransactor) SetDetails(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "setDetails", _name, _symbol)
}

// SetDetails is a paid mutator transaction binding the contract method 0xb7b090ee.
//
// Solidity: function setDetails(string _name, string _symbol) returns()
func (_OToken *OTokenSession) SetDetails(_name string, _symbol string) (*types.Transaction, error) {
	return _OToken.Contract.SetDetails(&_OToken.TransactOpts, _name, _symbol)
}

// SetDetails is a paid mutator transaction binding the contract method 0xb7b090ee.
//
// Solidity: function setDetails(string _name, string _symbol) returns()
func (_OToken *OTokenTransactorSession) SetDetails(_name string, _symbol string) (*types.Transaction, error) {
	return _OToken.Contract.SetDetails(&_OToken.TransactOpts, _name, _symbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_OToken *OTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_OToken *OTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.Transfer(&_OToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_OToken *OTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.Transfer(&_OToken.TransactOpts, recipient, amount)
}

// TransferFee is a paid mutator transaction binding the contract method 0xea8c4bcf.
//
// Solidity: function transferFee(address _address) returns()
func (_OToken *OTokenTransactor) TransferFee(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "transferFee", _address)
}

// TransferFee is a paid mutator transaction binding the contract method 0xea8c4bcf.
//
// Solidity: function transferFee(address _address) returns()
func (_OToken *OTokenSession) TransferFee(_address common.Address) (*types.Transaction, error) {
	return _OToken.Contract.TransferFee(&_OToken.TransactOpts, _address)
}

// TransferFee is a paid mutator transaction binding the contract method 0xea8c4bcf.
//
// Solidity: function transferFee(address _address) returns()
func (_OToken *OTokenTransactorSession) TransferFee(_address common.Address) (*types.Transaction, error) {
	return _OToken.Contract.TransferFee(&_OToken.TransactOpts, _address)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_OToken *OTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_OToken *OTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.TransferFrom(&_OToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_OToken *OTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.TransferFrom(&_OToken.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OToken *OTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OToken *OTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OToken.Contract.TransferOwnership(&_OToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OToken *OTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OToken.Contract.TransferOwnership(&_OToken.TransactOpts, newOwner)
}

// UpdateParameters is a paid mutator transaction binding the contract method 0xee1eab4f.
//
// Solidity: function updateParameters(uint256 _liquidationIncentive, uint256 _liquidationFactor, uint256 _transactionFee, uint256 _minCollateralizationRatio) returns()
func (_OToken *OTokenTransactor) UpdateParameters(opts *bind.TransactOpts, _liquidationIncentive *big.Int, _liquidationFactor *big.Int, _transactionFee *big.Int, _minCollateralizationRatio *big.Int) (*types.Transaction, error) {
	return _OToken.contract.Transact(opts, "updateParameters", _liquidationIncentive, _liquidationFactor, _transactionFee, _minCollateralizationRatio)
}

// UpdateParameters is a paid mutator transaction binding the contract method 0xee1eab4f.
//
// Solidity: function updateParameters(uint256 _liquidationIncentive, uint256 _liquidationFactor, uint256 _transactionFee, uint256 _minCollateralizationRatio) returns()
func (_OToken *OTokenSession) UpdateParameters(_liquidationIncentive *big.Int, _liquidationFactor *big.Int, _transactionFee *big.Int, _minCollateralizationRatio *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.UpdateParameters(&_OToken.TransactOpts, _liquidationIncentive, _liquidationFactor, _transactionFee, _minCollateralizationRatio)
}

// UpdateParameters is a paid mutator transaction binding the contract method 0xee1eab4f.
//
// Solidity: function updateParameters(uint256 _liquidationIncentive, uint256 _liquidationFactor, uint256 _transactionFee, uint256 _minCollateralizationRatio) returns()
func (_OToken *OTokenTransactorSession) UpdateParameters(_liquidationIncentive *big.Int, _liquidationFactor *big.Int, _transactionFee *big.Int, _minCollateralizationRatio *big.Int) (*types.Transaction, error) {
	return _OToken.Contract.UpdateParameters(&_OToken.TransactOpts, _liquidationIncentive, _liquidationFactor, _transactionFee, _minCollateralizationRatio)
}

// OTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the OToken contract.
type OTokenApprovalIterator struct {
	Event *OTokenApproval // Event containing the contract specifics and raw log

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
func (it *OTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTokenApproval)
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
		it.Event = new(OTokenApproval)
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
func (it *OTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTokenApproval represents a Approval event raised by the OToken contract.
type OTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OToken *OTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*OTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OTokenApprovalIterator{contract: _OToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OToken *OTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTokenApproval)
				if err := _OToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_OToken *OTokenFilterer) ParseApproval(log types.Log) (*OTokenApproval, error) {
	event := new(OTokenApproval)
	if err := _OToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OTokenBurnOTokensIterator is returned from FilterBurnOTokens and is used to iterate over the raw logs and unpacked data for BurnOTokens events raised by the OToken contract.
type OTokenBurnOTokensIterator struct {
	Event *OTokenBurnOTokens // Event containing the contract specifics and raw log

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
func (it *OTokenBurnOTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTokenBurnOTokens)
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
		it.Event = new(OTokenBurnOTokens)
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
func (it *OTokenBurnOTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTokenBurnOTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTokenBurnOTokens represents a BurnOTokens event raised by the OToken contract.
type OTokenBurnOTokens struct {
	VaultOwner    common.Address
	OTokensBurned *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBurnOTokens is a free log retrieval operation binding the contract event 0xdf8cebdea6ef1fd20576b80bc951377c0e61e2a8169153a1f836673ccce80e62.
//
// Solidity: event BurnOTokens(address vaultOwner, uint256 oTokensBurned)
func (_OToken *OTokenFilterer) FilterBurnOTokens(opts *bind.FilterOpts) (*OTokenBurnOTokensIterator, error) {

	logs, sub, err := _OToken.contract.FilterLogs(opts, "BurnOTokens")
	if err != nil {
		return nil, err
	}
	return &OTokenBurnOTokensIterator{contract: _OToken.contract, event: "BurnOTokens", logs: logs, sub: sub}, nil
}

// WatchBurnOTokens is a free log subscription operation binding the contract event 0xdf8cebdea6ef1fd20576b80bc951377c0e61e2a8169153a1f836673ccce80e62.
//
// Solidity: event BurnOTokens(address vaultOwner, uint256 oTokensBurned)
func (_OToken *OTokenFilterer) WatchBurnOTokens(opts *bind.WatchOpts, sink chan<- *OTokenBurnOTokens) (event.Subscription, error) {

	logs, sub, err := _OToken.contract.WatchLogs(opts, "BurnOTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTokenBurnOTokens)
				if err := _OToken.contract.UnpackLog(event, "BurnOTokens", log); err != nil {
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

// ParseBurnOTokens is a log parse operation binding the contract event 0xdf8cebdea6ef1fd20576b80bc951377c0e61e2a8169153a1f836673ccce80e62.
//
// Solidity: event BurnOTokens(address vaultOwner, uint256 oTokensBurned)
func (_OToken *OTokenFilterer) ParseBurnOTokens(log types.Log) (*OTokenBurnOTokens, error) {
	event := new(OTokenBurnOTokens)
	if err := _OToken.contract.UnpackLog(event, "BurnOTokens", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OTokenERC20CollateralAddedIterator is returned from FilterERC20CollateralAdded and is used to iterate over the raw logs and unpacked data for ERC20CollateralAdded events raised by the OToken contract.
type OTokenERC20CollateralAddedIterator struct {
	Event *OTokenERC20CollateralAdded // Event containing the contract specifics and raw log

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
func (it *OTokenERC20CollateralAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTokenERC20CollateralAdded)
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
		it.Event = new(OTokenERC20CollateralAdded)
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
func (it *OTokenERC20CollateralAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTokenERC20CollateralAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTokenERC20CollateralAdded represents a ERC20CollateralAdded event raised by the OToken contract.
type OTokenERC20CollateralAdded struct {
	VaultOwner common.Address
	Amount     *big.Int
	Payer      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterERC20CollateralAdded is a free log retrieval operation binding the contract event 0x2199418ea9428ed3ff7d460860e1edaf5831452fa4ea0f8d1a60d63c60348782.
//
// Solidity: event ERC20CollateralAdded(address vaultOwner, uint256 amount, address payer)
func (_OToken *OTokenFilterer) FilterERC20CollateralAdded(opts *bind.FilterOpts) (*OTokenERC20CollateralAddedIterator, error) {

	logs, sub, err := _OToken.contract.FilterLogs(opts, "ERC20CollateralAdded")
	if err != nil {
		return nil, err
	}
	return &OTokenERC20CollateralAddedIterator{contract: _OToken.contract, event: "ERC20CollateralAdded", logs: logs, sub: sub}, nil
}

// WatchERC20CollateralAdded is a free log subscription operation binding the contract event 0x2199418ea9428ed3ff7d460860e1edaf5831452fa4ea0f8d1a60d63c60348782.
//
// Solidity: event ERC20CollateralAdded(address vaultOwner, uint256 amount, address payer)
func (_OToken *OTokenFilterer) WatchERC20CollateralAdded(opts *bind.WatchOpts, sink chan<- *OTokenERC20CollateralAdded) (event.Subscription, error) {

	logs, sub, err := _OToken.contract.WatchLogs(opts, "ERC20CollateralAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTokenERC20CollateralAdded)
				if err := _OToken.contract.UnpackLog(event, "ERC20CollateralAdded", log); err != nil {
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

// ParseERC20CollateralAdded is a log parse operation binding the contract event 0x2199418ea9428ed3ff7d460860e1edaf5831452fa4ea0f8d1a60d63c60348782.
//
// Solidity: event ERC20CollateralAdded(address vaultOwner, uint256 amount, address payer)
func (_OToken *OTokenFilterer) ParseERC20CollateralAdded(log types.Log) (*OTokenERC20CollateralAdded, error) {
	event := new(OTokenERC20CollateralAdded)
	if err := _OToken.contract.UnpackLog(event, "ERC20CollateralAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OTokenETHCollateralAddedIterator is returned from FilterETHCollateralAdded and is used to iterate over the raw logs and unpacked data for ETHCollateralAdded events raised by the OToken contract.
type OTokenETHCollateralAddedIterator struct {
	Event *OTokenETHCollateralAdded // Event containing the contract specifics and raw log

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
func (it *OTokenETHCollateralAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTokenETHCollateralAdded)
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
		it.Event = new(OTokenETHCollateralAdded)
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
func (it *OTokenETHCollateralAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTokenETHCollateralAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTokenETHCollateralAdded represents a ETHCollateralAdded event raised by the OToken contract.
type OTokenETHCollateralAdded struct {
	VaultOwner common.Address
	Amount     *big.Int
	Payer      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterETHCollateralAdded is a free log retrieval operation binding the contract event 0xf24ce6016de57e90501829715846e26ac283a0aabfc160647e0ae8b05e0f433d.
//
// Solidity: event ETHCollateralAdded(address vaultOwner, uint256 amount, address payer)
func (_OToken *OTokenFilterer) FilterETHCollateralAdded(opts *bind.FilterOpts) (*OTokenETHCollateralAddedIterator, error) {

	logs, sub, err := _OToken.contract.FilterLogs(opts, "ETHCollateralAdded")
	if err != nil {
		return nil, err
	}
	return &OTokenETHCollateralAddedIterator{contract: _OToken.contract, event: "ETHCollateralAdded", logs: logs, sub: sub}, nil
}

// WatchETHCollateralAdded is a free log subscription operation binding the contract event 0xf24ce6016de57e90501829715846e26ac283a0aabfc160647e0ae8b05e0f433d.
//
// Solidity: event ETHCollateralAdded(address vaultOwner, uint256 amount, address payer)
func (_OToken *OTokenFilterer) WatchETHCollateralAdded(opts *bind.WatchOpts, sink chan<- *OTokenETHCollateralAdded) (event.Subscription, error) {

	logs, sub, err := _OToken.contract.WatchLogs(opts, "ETHCollateralAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTokenETHCollateralAdded)
				if err := _OToken.contract.UnpackLog(event, "ETHCollateralAdded", log); err != nil {
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

// ParseETHCollateralAdded is a log parse operation binding the contract event 0xf24ce6016de57e90501829715846e26ac283a0aabfc160647e0ae8b05e0f433d.
//
// Solidity: event ETHCollateralAdded(address vaultOwner, uint256 amount, address payer)
func (_OToken *OTokenFilterer) ParseETHCollateralAdded(log types.Log) (*OTokenETHCollateralAdded, error) {
	event := new(OTokenETHCollateralAdded)
	if err := _OToken.contract.UnpackLog(event, "ETHCollateralAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OTokenExerciseIterator is returned from FilterExercise and is used to iterate over the raw logs and unpacked data for Exercise events raised by the OToken contract.
type OTokenExerciseIterator struct {
	Event *OTokenExercise // Event containing the contract specifics and raw log

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
func (it *OTokenExerciseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTokenExercise)
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
		it.Event = new(OTokenExercise)
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
func (it *OTokenExerciseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTokenExerciseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTokenExercise represents a Exercise event raised by the OToken contract.
type OTokenExercise struct {
	AmtUnderlyingToPay *big.Int
	AmtCollateralToPay *big.Int
	Exerciser          common.Address
	VaultExercisedFrom common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterExercise is a free log retrieval operation binding the contract event 0xfa7bab37479e50a9b24a9412b879d400de9bcaa1e3a2b343e90bb370d85bbaa7.
//
// Solidity: event Exercise(uint256 amtUnderlyingToPay, uint256 amtCollateralToPay, address exerciser, address vaultExercisedFrom)
func (_OToken *OTokenFilterer) FilterExercise(opts *bind.FilterOpts) (*OTokenExerciseIterator, error) {

	logs, sub, err := _OToken.contract.FilterLogs(opts, "Exercise")
	if err != nil {
		return nil, err
	}
	return &OTokenExerciseIterator{contract: _OToken.contract, event: "Exercise", logs: logs, sub: sub}, nil
}

// WatchExercise is a free log subscription operation binding the contract event 0xfa7bab37479e50a9b24a9412b879d400de9bcaa1e3a2b343e90bb370d85bbaa7.
//
// Solidity: event Exercise(uint256 amtUnderlyingToPay, uint256 amtCollateralToPay, address exerciser, address vaultExercisedFrom)
func (_OToken *OTokenFilterer) WatchExercise(opts *bind.WatchOpts, sink chan<- *OTokenExercise) (event.Subscription, error) {

	logs, sub, err := _OToken.contract.WatchLogs(opts, "Exercise")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTokenExercise)
				if err := _OToken.contract.UnpackLog(event, "Exercise", log); err != nil {
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

// ParseExercise is a log parse operation binding the contract event 0xfa7bab37479e50a9b24a9412b879d400de9bcaa1e3a2b343e90bb370d85bbaa7.
//
// Solidity: event Exercise(uint256 amtUnderlyingToPay, uint256 amtCollateralToPay, address exerciser, address vaultExercisedFrom)
func (_OToken *OTokenFilterer) ParseExercise(log types.Log) (*OTokenExercise, error) {
	event := new(OTokenExercise)
	if err := _OToken.contract.UnpackLog(event, "Exercise", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OTokenIssuedOTokensIterator is returned from FilterIssuedOTokens and is used to iterate over the raw logs and unpacked data for IssuedOTokens events raised by the OToken contract.
type OTokenIssuedOTokensIterator struct {
	Event *OTokenIssuedOTokens // Event containing the contract specifics and raw log

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
func (it *OTokenIssuedOTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTokenIssuedOTokens)
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
		it.Event = new(OTokenIssuedOTokens)
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
func (it *OTokenIssuedOTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTokenIssuedOTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTokenIssuedOTokens represents a IssuedOTokens event raised by the OToken contract.
type OTokenIssuedOTokens struct {
	IssuedTo      common.Address
	OTokensIssued *big.Int
	VaultOwner    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterIssuedOTokens is a free log retrieval operation binding the contract event 0x5e5aaabf04e3760968ffb551bdf9708f4dbf95d53ad98539e91a56b125e88f08.
//
// Solidity: event IssuedOTokens(address issuedTo, uint256 oTokensIssued, address vaultOwner)
func (_OToken *OTokenFilterer) FilterIssuedOTokens(opts *bind.FilterOpts) (*OTokenIssuedOTokensIterator, error) {

	logs, sub, err := _OToken.contract.FilterLogs(opts, "IssuedOTokens")
	if err != nil {
		return nil, err
	}
	return &OTokenIssuedOTokensIterator{contract: _OToken.contract, event: "IssuedOTokens", logs: logs, sub: sub}, nil
}

// WatchIssuedOTokens is a free log subscription operation binding the contract event 0x5e5aaabf04e3760968ffb551bdf9708f4dbf95d53ad98539e91a56b125e88f08.
//
// Solidity: event IssuedOTokens(address issuedTo, uint256 oTokensIssued, address vaultOwner)
func (_OToken *OTokenFilterer) WatchIssuedOTokens(opts *bind.WatchOpts, sink chan<- *OTokenIssuedOTokens) (event.Subscription, error) {

	logs, sub, err := _OToken.contract.WatchLogs(opts, "IssuedOTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTokenIssuedOTokens)
				if err := _OToken.contract.UnpackLog(event, "IssuedOTokens", log); err != nil {
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

// ParseIssuedOTokens is a log parse operation binding the contract event 0x5e5aaabf04e3760968ffb551bdf9708f4dbf95d53ad98539e91a56b125e88f08.
//
// Solidity: event IssuedOTokens(address issuedTo, uint256 oTokensIssued, address vaultOwner)
func (_OToken *OTokenFilterer) ParseIssuedOTokens(log types.Log) (*OTokenIssuedOTokens, error) {
	event := new(OTokenIssuedOTokens)
	if err := _OToken.contract.UnpackLog(event, "IssuedOTokens", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OTokenLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the OToken contract.
type OTokenLiquidateIterator struct {
	Event *OTokenLiquidate // Event containing the contract specifics and raw log

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
func (it *OTokenLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTokenLiquidate)
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
		it.Event = new(OTokenLiquidate)
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
func (it *OTokenLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTokenLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTokenLiquidate represents a Liquidate event raised by the OToken contract.
type OTokenLiquidate struct {
	AmtCollateralToPay *big.Int
	VaultOwner         common.Address
	Liquidator         common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0xcab8e1abb9f8235c6db895cf185336dc9461aecf477b98c1be83687ee549e66a.
//
// Solidity: event Liquidate(uint256 amtCollateralToPay, address vaultOwner, address liquidator)
func (_OToken *OTokenFilterer) FilterLiquidate(opts *bind.FilterOpts) (*OTokenLiquidateIterator, error) {

	logs, sub, err := _OToken.contract.FilterLogs(opts, "Liquidate")
	if err != nil {
		return nil, err
	}
	return &OTokenLiquidateIterator{contract: _OToken.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0xcab8e1abb9f8235c6db895cf185336dc9461aecf477b98c1be83687ee549e66a.
//
// Solidity: event Liquidate(uint256 amtCollateralToPay, address vaultOwner, address liquidator)
func (_OToken *OTokenFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *OTokenLiquidate) (event.Subscription, error) {

	logs, sub, err := _OToken.contract.WatchLogs(opts, "Liquidate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTokenLiquidate)
				if err := _OToken.contract.UnpackLog(event, "Liquidate", log); err != nil {
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

// ParseLiquidate is a log parse operation binding the contract event 0xcab8e1abb9f8235c6db895cf185336dc9461aecf477b98c1be83687ee549e66a.
//
// Solidity: event Liquidate(uint256 amtCollateralToPay, address vaultOwner, address liquidator)
func (_OToken *OTokenFilterer) ParseLiquidate(log types.Log) (*OTokenLiquidate, error) {
	event := new(OTokenLiquidate)
	if err := _OToken.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OToken contract.
type OTokenOwnershipTransferredIterator struct {
	Event *OTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTokenOwnershipTransferred)
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
		it.Event = new(OTokenOwnershipTransferred)
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
func (it *OTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTokenOwnershipTransferred represents a OwnershipTransferred event raised by the OToken contract.
type OTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OToken *OTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OTokenOwnershipTransferredIterator{contract: _OToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OToken *OTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTokenOwnershipTransferred)
				if err := _OToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OToken *OTokenFilterer) ParseOwnershipTransferred(log types.Log) (*OTokenOwnershipTransferred, error) {
	event := new(OTokenOwnershipTransferred)
	if err := _OToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OTokenRedeemVaultBalanceIterator is returned from FilterRedeemVaultBalance and is used to iterate over the raw logs and unpacked data for RedeemVaultBalance events raised by the OToken contract.
type OTokenRedeemVaultBalanceIterator struct {
	Event *OTokenRedeemVaultBalance // Event containing the contract specifics and raw log

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
func (it *OTokenRedeemVaultBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTokenRedeemVaultBalance)
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
		it.Event = new(OTokenRedeemVaultBalance)
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
func (it *OTokenRedeemVaultBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTokenRedeemVaultBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTokenRedeemVaultBalance represents a RedeemVaultBalance event raised by the OToken contract.
type OTokenRedeemVaultBalance struct {
	AmtCollateralRedeemed *big.Int
	AmtUnderlyingRedeemed *big.Int
	VaultOwner            common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRedeemVaultBalance is a free log retrieval operation binding the contract event 0xe481532a3f7d078365ca0145442ed0a0a3e0443f3c0bae0c29cff13111267838.
//
// Solidity: event RedeemVaultBalance(uint256 amtCollateralRedeemed, uint256 amtUnderlyingRedeemed, address vaultOwner)
func (_OToken *OTokenFilterer) FilterRedeemVaultBalance(opts *bind.FilterOpts) (*OTokenRedeemVaultBalanceIterator, error) {

	logs, sub, err := _OToken.contract.FilterLogs(opts, "RedeemVaultBalance")
	if err != nil {
		return nil, err
	}
	return &OTokenRedeemVaultBalanceIterator{contract: _OToken.contract, event: "RedeemVaultBalance", logs: logs, sub: sub}, nil
}

// WatchRedeemVaultBalance is a free log subscription operation binding the contract event 0xe481532a3f7d078365ca0145442ed0a0a3e0443f3c0bae0c29cff13111267838.
//
// Solidity: event RedeemVaultBalance(uint256 amtCollateralRedeemed, uint256 amtUnderlyingRedeemed, address vaultOwner)
func (_OToken *OTokenFilterer) WatchRedeemVaultBalance(opts *bind.WatchOpts, sink chan<- *OTokenRedeemVaultBalance) (event.Subscription, error) {

	logs, sub, err := _OToken.contract.WatchLogs(opts, "RedeemVaultBalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTokenRedeemVaultBalance)
				if err := _OToken.contract.UnpackLog(event, "RedeemVaultBalance", log); err != nil {
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

// ParseRedeemVaultBalance is a log parse operation binding the contract event 0xe481532a3f7d078365ca0145442ed0a0a3e0443f3c0bae0c29cff13111267838.
//
// Solidity: event RedeemVaultBalance(uint256 amtCollateralRedeemed, uint256 amtUnderlyingRedeemed, address vaultOwner)
func (_OToken *OTokenFilterer) ParseRedeemVaultBalance(log types.Log) (*OTokenRedeemVaultBalance, error) {
	event := new(OTokenRedeemVaultBalance)
	if err := _OToken.contract.UnpackLog(event, "RedeemVaultBalance", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OTokenRemoveCollateralIterator is returned from FilterRemoveCollateral and is used to iterate over the raw logs and unpacked data for RemoveCollateral events raised by the OToken contract.
type OTokenRemoveCollateralIterator struct {
	Event *OTokenRemoveCollateral // Event containing the contract specifics and raw log

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
func (it *OTokenRemoveCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTokenRemoveCollateral)
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
		it.Event = new(OTokenRemoveCollateral)
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
func (it *OTokenRemoveCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTokenRemoveCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTokenRemoveCollateral represents a RemoveCollateral event raised by the OToken contract.
type OTokenRemoveCollateral struct {
	AmtRemoved *big.Int
	VaultOwner common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemoveCollateral is a free log retrieval operation binding the contract event 0x5a945309b3c58e9bb259128c2a530a6579dc75ac1d7d61b3db4c0b8305a16821.
//
// Solidity: event RemoveCollateral(uint256 amtRemoved, address vaultOwner)
func (_OToken *OTokenFilterer) FilterRemoveCollateral(opts *bind.FilterOpts) (*OTokenRemoveCollateralIterator, error) {

	logs, sub, err := _OToken.contract.FilterLogs(opts, "RemoveCollateral")
	if err != nil {
		return nil, err
	}
	return &OTokenRemoveCollateralIterator{contract: _OToken.contract, event: "RemoveCollateral", logs: logs, sub: sub}, nil
}

// WatchRemoveCollateral is a free log subscription operation binding the contract event 0x5a945309b3c58e9bb259128c2a530a6579dc75ac1d7d61b3db4c0b8305a16821.
//
// Solidity: event RemoveCollateral(uint256 amtRemoved, address vaultOwner)
func (_OToken *OTokenFilterer) WatchRemoveCollateral(opts *bind.WatchOpts, sink chan<- *OTokenRemoveCollateral) (event.Subscription, error) {

	logs, sub, err := _OToken.contract.WatchLogs(opts, "RemoveCollateral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTokenRemoveCollateral)
				if err := _OToken.contract.UnpackLog(event, "RemoveCollateral", log); err != nil {
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

// ParseRemoveCollateral is a log parse operation binding the contract event 0x5a945309b3c58e9bb259128c2a530a6579dc75ac1d7d61b3db4c0b8305a16821.
//
// Solidity: event RemoveCollateral(uint256 amtRemoved, address vaultOwner)
func (_OToken *OTokenFilterer) ParseRemoveCollateral(log types.Log) (*OTokenRemoveCollateral, error) {
	event := new(OTokenRemoveCollateral)
	if err := _OToken.contract.UnpackLog(event, "RemoveCollateral", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OTokenRemoveUnderlyingIterator is returned from FilterRemoveUnderlying and is used to iterate over the raw logs and unpacked data for RemoveUnderlying events raised by the OToken contract.
type OTokenRemoveUnderlyingIterator struct {
	Event *OTokenRemoveUnderlying // Event containing the contract specifics and raw log

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
func (it *OTokenRemoveUnderlyingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTokenRemoveUnderlying)
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
		it.Event = new(OTokenRemoveUnderlying)
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
func (it *OTokenRemoveUnderlyingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTokenRemoveUnderlyingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTokenRemoveUnderlying represents a RemoveUnderlying event raised by the OToken contract.
type OTokenRemoveUnderlying struct {
	AmountUnderlying *big.Int
	VaultOwner       common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRemoveUnderlying is a free log retrieval operation binding the contract event 0xea0bff65fa9380b944e9a761f9c6a665ad2d31e74706a52773ddb45c8a57c83d.
//
// Solidity: event RemoveUnderlying(uint256 amountUnderlying, address vaultOwner)
func (_OToken *OTokenFilterer) FilterRemoveUnderlying(opts *bind.FilterOpts) (*OTokenRemoveUnderlyingIterator, error) {

	logs, sub, err := _OToken.contract.FilterLogs(opts, "RemoveUnderlying")
	if err != nil {
		return nil, err
	}
	return &OTokenRemoveUnderlyingIterator{contract: _OToken.contract, event: "RemoveUnderlying", logs: logs, sub: sub}, nil
}

// WatchRemoveUnderlying is a free log subscription operation binding the contract event 0xea0bff65fa9380b944e9a761f9c6a665ad2d31e74706a52773ddb45c8a57c83d.
//
// Solidity: event RemoveUnderlying(uint256 amountUnderlying, address vaultOwner)
func (_OToken *OTokenFilterer) WatchRemoveUnderlying(opts *bind.WatchOpts, sink chan<- *OTokenRemoveUnderlying) (event.Subscription, error) {

	logs, sub, err := _OToken.contract.WatchLogs(opts, "RemoveUnderlying")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTokenRemoveUnderlying)
				if err := _OToken.contract.UnpackLog(event, "RemoveUnderlying", log); err != nil {
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

// ParseRemoveUnderlying is a log parse operation binding the contract event 0xea0bff65fa9380b944e9a761f9c6a665ad2d31e74706a52773ddb45c8a57c83d.
//
// Solidity: event RemoveUnderlying(uint256 amountUnderlying, address vaultOwner)
func (_OToken *OTokenFilterer) ParseRemoveUnderlying(log types.Log) (*OTokenRemoveUnderlying, error) {
	event := new(OTokenRemoveUnderlying)
	if err := _OToken.contract.UnpackLog(event, "RemoveUnderlying", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the OToken contract.
type OTokenTransferIterator struct {
	Event *OTokenTransfer // Event containing the contract specifics and raw log

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
func (it *OTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTokenTransfer)
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
		it.Event = new(OTokenTransfer)
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
func (it *OTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTokenTransfer represents a Transfer event raised by the OToken contract.
type OTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OToken *OTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OTokenTransferIterator{contract: _OToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OToken *OTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTokenTransfer)
				if err := _OToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_OToken *OTokenFilterer) ParseTransfer(log types.Log) (*OTokenTransfer, error) {
	event := new(OTokenTransfer)
	if err := _OToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OTokenTransferFeeIterator is returned from FilterTransferFee and is used to iterate over the raw logs and unpacked data for TransferFee events raised by the OToken contract.
type OTokenTransferFeeIterator struct {
	Event *OTokenTransferFee // Event containing the contract specifics and raw log

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
func (it *OTokenTransferFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTokenTransferFee)
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
		it.Event = new(OTokenTransferFee)
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
func (it *OTokenTransferFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTokenTransferFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTokenTransferFee represents a TransferFee event raised by the OToken contract.
type OTokenTransferFee struct {
	To   common.Address
	Fees *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransferFee is a free log retrieval operation binding the contract event 0x88b171bb78d3ac5e1caa8e729dddce4e1322e84c80c093ebbe52507b62c77d98.
//
// Solidity: event TransferFee(address to, uint256 fees)
func (_OToken *OTokenFilterer) FilterTransferFee(opts *bind.FilterOpts) (*OTokenTransferFeeIterator, error) {

	logs, sub, err := _OToken.contract.FilterLogs(opts, "TransferFee")
	if err != nil {
		return nil, err
	}
	return &OTokenTransferFeeIterator{contract: _OToken.contract, event: "TransferFee", logs: logs, sub: sub}, nil
}

// WatchTransferFee is a free log subscription operation binding the contract event 0x88b171bb78d3ac5e1caa8e729dddce4e1322e84c80c093ebbe52507b62c77d98.
//
// Solidity: event TransferFee(address to, uint256 fees)
func (_OToken *OTokenFilterer) WatchTransferFee(opts *bind.WatchOpts, sink chan<- *OTokenTransferFee) (event.Subscription, error) {

	logs, sub, err := _OToken.contract.WatchLogs(opts, "TransferFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTokenTransferFee)
				if err := _OToken.contract.UnpackLog(event, "TransferFee", log); err != nil {
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

// ParseTransferFee is a log parse operation binding the contract event 0x88b171bb78d3ac5e1caa8e729dddce4e1322e84c80c093ebbe52507b62c77d98.
//
// Solidity: event TransferFee(address to, uint256 fees)
func (_OToken *OTokenFilterer) ParseTransferFee(log types.Log) (*OTokenTransferFee, error) {
	event := new(OTokenTransferFee)
	if err := _OToken.contract.UnpackLog(event, "TransferFee", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OTokenUpdateParametersIterator is returned from FilterUpdateParameters and is used to iterate over the raw logs and unpacked data for UpdateParameters events raised by the OToken contract.
type OTokenUpdateParametersIterator struct {
	Event *OTokenUpdateParameters // Event containing the contract specifics and raw log

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
func (it *OTokenUpdateParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTokenUpdateParameters)
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
		it.Event = new(OTokenUpdateParameters)
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
func (it *OTokenUpdateParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTokenUpdateParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTokenUpdateParameters represents a UpdateParameters event raised by the OToken contract.
type OTokenUpdateParameters struct {
	LiquidationIncentive      *big.Int
	LiquidationFactor         *big.Int
	TransactionFee            *big.Int
	MinCollateralizationRatio *big.Int
	Owner                     common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterUpdateParameters is a free log retrieval operation binding the contract event 0x3450d20c21ea671871fed271900cc8ff03badafa9b6fe2ff7f86991950e86b6b.
//
// Solidity: event UpdateParameters(uint256 liquidationIncentive, uint256 liquidationFactor, uint256 transactionFee, uint256 minCollateralizationRatio, address owner)
func (_OToken *OTokenFilterer) FilterUpdateParameters(opts *bind.FilterOpts) (*OTokenUpdateParametersIterator, error) {

	logs, sub, err := _OToken.contract.FilterLogs(opts, "UpdateParameters")
	if err != nil {
		return nil, err
	}
	return &OTokenUpdateParametersIterator{contract: _OToken.contract, event: "UpdateParameters", logs: logs, sub: sub}, nil
}

// WatchUpdateParameters is a free log subscription operation binding the contract event 0x3450d20c21ea671871fed271900cc8ff03badafa9b6fe2ff7f86991950e86b6b.
//
// Solidity: event UpdateParameters(uint256 liquidationIncentive, uint256 liquidationFactor, uint256 transactionFee, uint256 minCollateralizationRatio, address owner)
func (_OToken *OTokenFilterer) WatchUpdateParameters(opts *bind.WatchOpts, sink chan<- *OTokenUpdateParameters) (event.Subscription, error) {

	logs, sub, err := _OToken.contract.WatchLogs(opts, "UpdateParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTokenUpdateParameters)
				if err := _OToken.contract.UnpackLog(event, "UpdateParameters", log); err != nil {
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

// ParseUpdateParameters is a log parse operation binding the contract event 0x3450d20c21ea671871fed271900cc8ff03badafa9b6fe2ff7f86991950e86b6b.
//
// Solidity: event UpdateParameters(uint256 liquidationIncentive, uint256 liquidationFactor, uint256 transactionFee, uint256 minCollateralizationRatio, address owner)
func (_OToken *OTokenFilterer) ParseUpdateParameters(log types.Log) (*OTokenUpdateParameters, error) {
	event := new(OTokenUpdateParameters)
	if err := _OToken.contract.UnpackLog(event, "UpdateParameters", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OTokenVaultOpenedIterator is returned from FilterVaultOpened and is used to iterate over the raw logs and unpacked data for VaultOpened events raised by the OToken contract.
type OTokenVaultOpenedIterator struct {
	Event *OTokenVaultOpened // Event containing the contract specifics and raw log

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
func (it *OTokenVaultOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OTokenVaultOpened)
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
		it.Event = new(OTokenVaultOpened)
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
func (it *OTokenVaultOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OTokenVaultOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OTokenVaultOpened represents a VaultOpened event raised by the OToken contract.
type OTokenVaultOpened struct {
	VaultOwner common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVaultOpened is a free log retrieval operation binding the contract event 0x66a872561db77eb92ef3079a44a5af00c68c3a09e0976814a95bd91721f57c2f.
//
// Solidity: event VaultOpened(address vaultOwner)
func (_OToken *OTokenFilterer) FilterVaultOpened(opts *bind.FilterOpts) (*OTokenVaultOpenedIterator, error) {

	logs, sub, err := _OToken.contract.FilterLogs(opts, "VaultOpened")
	if err != nil {
		return nil, err
	}
	return &OTokenVaultOpenedIterator{contract: _OToken.contract, event: "VaultOpened", logs: logs, sub: sub}, nil
}

// WatchVaultOpened is a free log subscription operation binding the contract event 0x66a872561db77eb92ef3079a44a5af00c68c3a09e0976814a95bd91721f57c2f.
//
// Solidity: event VaultOpened(address vaultOwner)
func (_OToken *OTokenFilterer) WatchVaultOpened(opts *bind.WatchOpts, sink chan<- *OTokenVaultOpened) (event.Subscription, error) {

	logs, sub, err := _OToken.contract.WatchLogs(opts, "VaultOpened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OTokenVaultOpened)
				if err := _OToken.contract.UnpackLog(event, "VaultOpened", log); err != nil {
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

// ParseVaultOpened is a log parse operation binding the contract event 0x66a872561db77eb92ef3079a44a5af00c68c3a09e0976814a95bd91721f57c2f.
//
// Solidity: event VaultOpened(address vaultOwner)
func (_OToken *OTokenFilterer) ParseVaultOpened(log types.Log) (*OTokenVaultOpened, error) {
	event := new(OTokenVaultOpened)
	if err := _OToken.contract.UnpackLog(event, "VaultOpened", log); err != nil {
		return nil, err
	}
	return event, nil
}
