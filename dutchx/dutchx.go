// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dutchx

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

// DutchExchangeABI is the input ABI used to generate the binding from.
const DutchExchangeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellVolume\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buyVolume\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"}],\"name\":\"AuctionCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionStart\",\"type\":\"uint256\"}],\"name\":\"AuctionStartScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"primaryToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"secondarToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Fee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NewBuyOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"frtsIssued\",\"type\":\"uint256\"}],\"name\":\"NewBuyerFundsClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NewDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newMasterCopy\",\"type\":\"address\"}],\"name\":\"NewMasterCopyProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractPriceOracleInterface\",\"name\":\"priceOracleInterface\",\"type\":\"address\"}],\"name\":\"NewOracleProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NewSellOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"frtsIssued\",\"type\":\"uint256\"}],\"name\":\"NewSellerFundsClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"}],\"name\":\"NewTokenPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NewWithdrawal\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"token1Funding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token2Funding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialClosingPriceNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialClosingPriceDen\",\"type\":\"uint256\"}],\"name\":\"addTokenPair\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"}],\"name\":\"atleastZero\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"auctionStarts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"auctioneer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"buyVolumes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"buyerBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimAndWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"frtsIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newBal\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"auctionSellTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"auctionBuyTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"auctionIndices\",\"type\":\"uint256[]\"}],\"name\":\"claimAndWithdrawTokensFromSeveralAuctionsAsBuyer\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"frtsIssued\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"auctionSellTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"auctionBuyTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"auctionIndices\",\"type\":\"uint256[]\"}],\"name\":\"claimAndWithdrawTokensFromSeveralAuctionsAsSeller\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"frtsIssued\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"}],\"name\":\"claimBuyerFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"frtsIssued\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"}],\"name\":\"claimSellerFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"frtsIssued\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"auctionSellTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"auctionBuyTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"auctionIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"claimTokensFromSeveralAuctionsAsBuyer\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"auctionSellTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"auctionBuyTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"auctionIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"claimTokensFromSeveralAuctionsAsSeller\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"clearingTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"}],\"name\":\"closeTheoreticalClosedAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"closingPrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"den\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositAndSell\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newBal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newSellerBal\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethUSDOracle\",\"outputs\":[{\"internalType\":\"contractPriceOracleInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"extraTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"frtToken\",\"outputs\":[{\"internalType\":\"contractTokenFRT\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addressesToCheck\",\"type\":\"address[]\"}],\"name\":\"getApprovedAddressesOfList\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token2\",\"type\":\"address\"}],\"name\":\"getAuctionIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token2\",\"type\":\"address\"}],\"name\":\"getAuctionStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionStart\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"}],\"name\":\"getClearingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"}],\"name\":\"getCurrentAuctionPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"den\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getFeeRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"den\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMasterCopy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"}],\"name\":\"getPriceInPastAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"den\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPriceOfTokenInLastAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"den\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token2\",\"type\":\"address\"}],\"name\":\"getTokenOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"}],\"name\":\"getUnclaimedBuyerFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"unclaimedBuyerFunds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"den\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractPriceOracleInterface\",\"name\":\"_ethUSDOracle\",\"type\":\"address\"}],\"name\":\"initiateEthUsdOracleUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"latestAuctionIndices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"masterCopy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"masterCopyCountdown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"min\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"mul\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newMasterCopy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newProposalEthUSDOracle\",\"outputs\":[{\"internalType\":\"contractPriceOracleInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracleInterfaceCountdown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owlToken\",\"outputs\":[{\"internalType\":\"contractTokenOWL\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"postBuyOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newBuyerBal\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"auctionIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"postSellOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"safeToAdd\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"safeToMul\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"safeToSub\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sellVolumesCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sellVolumesNext\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sellerBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractTokenFRT\",\"name\":\"_frtToken\",\"type\":\"address\"},{\"internalType\":\"contractTokenOWL\",\"name\":\"_owlToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_auctioneer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethToken\",\"type\":\"address\"},{\"internalType\":\"contractPriceOracleInterface\",\"name\":\"_ethUSDOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_thresholdNewTokenPair\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_thresholdNewAuction\",\"type\":\"uint256\"}],\"name\":\"setupDutchExchange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_masterCopy\",\"type\":\"address\"}],\"name\":\"startMasterCopyCountdown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"thresholdNewAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"thresholdNewTokenPair\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"updateApprovalOfToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auctioneer\",\"type\":\"address\"}],\"name\":\"updateAuctioneer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateEthUSDOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateMasterCopy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_thresholdNewAuction\",\"type\":\"uint256\"}],\"name\":\"updateThresholdNewAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_thresholdNewTokenPair\",\"type\":\"uint256\"}],\"name\":\"updateThresholdNewTokenPair\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DutchExchange is an auto generated Go binding around an Ethereum contract.
type DutchExchange struct {
	DutchExchangeCaller     // Read-only binding to the contract
	DutchExchangeTransactor // Write-only binding to the contract
	DutchExchangeFilterer   // Log filterer for contract events
}

// DutchExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type DutchExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DutchExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DutchExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DutchExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DutchExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DutchExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DutchExchangeSession struct {
	Contract     *DutchExchange    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DutchExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DutchExchangeCallerSession struct {
	Contract *DutchExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DutchExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DutchExchangeTransactorSession struct {
	Contract     *DutchExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DutchExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type DutchExchangeRaw struct {
	Contract *DutchExchange // Generic contract binding to access the raw methods on
}

// DutchExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DutchExchangeCallerRaw struct {
	Contract *DutchExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// DutchExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DutchExchangeTransactorRaw struct {
	Contract *DutchExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDutchExchange creates a new instance of DutchExchange, bound to a specific deployed contract.
func NewDutchExchange(address common.Address, backend bind.ContractBackend) (*DutchExchange, error) {
	contract, err := bindDutchExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DutchExchange{DutchExchangeCaller: DutchExchangeCaller{contract: contract}, DutchExchangeTransactor: DutchExchangeTransactor{contract: contract}, DutchExchangeFilterer: DutchExchangeFilterer{contract: contract}}, nil
}

// NewDutchExchangeCaller creates a new read-only instance of DutchExchange, bound to a specific deployed contract.
func NewDutchExchangeCaller(address common.Address, caller bind.ContractCaller) (*DutchExchangeCaller, error) {
	contract, err := bindDutchExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DutchExchangeCaller{contract: contract}, nil
}

// NewDutchExchangeTransactor creates a new write-only instance of DutchExchange, bound to a specific deployed contract.
func NewDutchExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*DutchExchangeTransactor, error) {
	contract, err := bindDutchExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DutchExchangeTransactor{contract: contract}, nil
}

// NewDutchExchangeFilterer creates a new log filterer instance of DutchExchange, bound to a specific deployed contract.
func NewDutchExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*DutchExchangeFilterer, error) {
	contract, err := bindDutchExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DutchExchangeFilterer{contract: contract}, nil
}

// bindDutchExchange binds a generic wrapper to an already deployed contract.
func bindDutchExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DutchExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DutchExchange *DutchExchangeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DutchExchange.Contract.DutchExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DutchExchange *DutchExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DutchExchange.Contract.DutchExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DutchExchange *DutchExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DutchExchange.Contract.DutchExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DutchExchange *DutchExchangeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DutchExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DutchExchange *DutchExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DutchExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DutchExchange *DutchExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DutchExchange.Contract.contract.Transact(opts, method, params...)
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 a, uint256 b) constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) Add(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "add", a, b)
	return *ret0, err
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 a, uint256 b) constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) Add(a *big.Int, b *big.Int) (*big.Int, error) {
	return _DutchExchange.Contract.Add(&_DutchExchange.CallOpts, a, b)
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 a, uint256 b) constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) Add(a *big.Int, b *big.Int) (*big.Int, error) {
	return _DutchExchange.Contract.Add(&_DutchExchange.CallOpts, a, b)
}

// ApprovedTokens is a free data retrieval call binding the contract method 0x6d1ea3fa.
//
// Solidity: function approvedTokens(address ) constant returns(bool)
func (_DutchExchange *DutchExchangeCaller) ApprovedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "approvedTokens", arg0)
	return *ret0, err
}

// ApprovedTokens is a free data retrieval call binding the contract method 0x6d1ea3fa.
//
// Solidity: function approvedTokens(address ) constant returns(bool)
func (_DutchExchange *DutchExchangeSession) ApprovedTokens(arg0 common.Address) (bool, error) {
	return _DutchExchange.Contract.ApprovedTokens(&_DutchExchange.CallOpts, arg0)
}

// ApprovedTokens is a free data retrieval call binding the contract method 0x6d1ea3fa.
//
// Solidity: function approvedTokens(address ) constant returns(bool)
func (_DutchExchange *DutchExchangeCallerSession) ApprovedTokens(arg0 common.Address) (bool, error) {
	return _DutchExchange.Contract.ApprovedTokens(&_DutchExchange.CallOpts, arg0)
}

// AtleastZero is a free data retrieval call binding the contract method 0x30690468.
//
// Solidity: function atleastZero(int256 a) constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) AtleastZero(opts *bind.CallOpts, a *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "atleastZero", a)
	return *ret0, err
}

// AtleastZero is a free data retrieval call binding the contract method 0x30690468.
//
// Solidity: function atleastZero(int256 a) constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) AtleastZero(a *big.Int) (*big.Int, error) {
	return _DutchExchange.Contract.AtleastZero(&_DutchExchange.CallOpts, a)
}

// AtleastZero is a free data retrieval call binding the contract method 0x30690468.
//
// Solidity: function atleastZero(int256 a) constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) AtleastZero(a *big.Int) (*big.Int, error) {
	return _DutchExchange.Contract.AtleastZero(&_DutchExchange.CallOpts, a)
}

// AuctionStarts is a free data retrieval call binding the contract method 0x1006a41f.
//
// Solidity: function auctionStarts(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) AuctionStarts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "auctionStarts", arg0, arg1)
	return *ret0, err
}

// AuctionStarts is a free data retrieval call binding the contract method 0x1006a41f.
//
// Solidity: function auctionStarts(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) AuctionStarts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.AuctionStarts(&_DutchExchange.CallOpts, arg0, arg1)
}

// AuctionStarts is a free data retrieval call binding the contract method 0x1006a41f.
//
// Solidity: function auctionStarts(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) AuctionStarts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.AuctionStarts(&_DutchExchange.CallOpts, arg0, arg1)
}

// Auctioneer is a free data retrieval call binding the contract method 0x5ec2c7bf.
//
// Solidity: function auctioneer() constant returns(address)
func (_DutchExchange *DutchExchangeCaller) Auctioneer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "auctioneer")
	return *ret0, err
}

// Auctioneer is a free data retrieval call binding the contract method 0x5ec2c7bf.
//
// Solidity: function auctioneer() constant returns(address)
func (_DutchExchange *DutchExchangeSession) Auctioneer() (common.Address, error) {
	return _DutchExchange.Contract.Auctioneer(&_DutchExchange.CallOpts)
}

// Auctioneer is a free data retrieval call binding the contract method 0x5ec2c7bf.
//
// Solidity: function auctioneer() constant returns(address)
func (_DutchExchange *DutchExchangeCallerSession) Auctioneer() (common.Address, error) {
	return _DutchExchange.Contract.Auctioneer(&_DutchExchange.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "balances", arg0, arg1)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) Balances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.Balances(&_DutchExchange.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xc23f001f.
//
// Solidity: function balances(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) Balances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.Balances(&_DutchExchange.CallOpts, arg0, arg1)
}

// BuyVolumes is a free data retrieval call binding the contract method 0xb8beafd6.
//
// Solidity: function buyVolumes(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) BuyVolumes(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "buyVolumes", arg0, arg1)
	return *ret0, err
}

// BuyVolumes is a free data retrieval call binding the contract method 0xb8beafd6.
//
// Solidity: function buyVolumes(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) BuyVolumes(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.BuyVolumes(&_DutchExchange.CallOpts, arg0, arg1)
}

// BuyVolumes is a free data retrieval call binding the contract method 0xb8beafd6.
//
// Solidity: function buyVolumes(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) BuyVolumes(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.BuyVolumes(&_DutchExchange.CallOpts, arg0, arg1)
}

// BuyerBalances is a free data retrieval call binding the contract method 0x37775807.
//
// Solidity: function buyerBalances(address , address , uint256 , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) BuyerBalances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "buyerBalances", arg0, arg1, arg2, arg3)
	return *ret0, err
}

// BuyerBalances is a free data retrieval call binding the contract method 0x37775807.
//
// Solidity: function buyerBalances(address , address , uint256 , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) BuyerBalances(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.BuyerBalances(&_DutchExchange.CallOpts, arg0, arg1, arg2, arg3)
}

// BuyerBalances is a free data retrieval call binding the contract method 0x37775807.
//
// Solidity: function buyerBalances(address , address , uint256 , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) BuyerBalances(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.BuyerBalances(&_DutchExchange.CallOpts, arg0, arg1, arg2, arg3)
}

// ClaimedAmounts is a free data retrieval call binding the contract method 0x9fec8e96.
//
// Solidity: function claimedAmounts(address , address , uint256 , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) ClaimedAmounts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "claimedAmounts", arg0, arg1, arg2, arg3)
	return *ret0, err
}

// ClaimedAmounts is a free data retrieval call binding the contract method 0x9fec8e96.
//
// Solidity: function claimedAmounts(address , address , uint256 , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) ClaimedAmounts(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.ClaimedAmounts(&_DutchExchange.CallOpts, arg0, arg1, arg2, arg3)
}

// ClaimedAmounts is a free data retrieval call binding the contract method 0x9fec8e96.
//
// Solidity: function claimedAmounts(address , address , uint256 , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) ClaimedAmounts(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.ClaimedAmounts(&_DutchExchange.CallOpts, arg0, arg1, arg2, arg3)
}

// ClearingTimes is a free data retrieval call binding the contract method 0xb04c0239.
//
// Solidity: function clearingTimes(address , address , uint256 ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) ClearingTimes(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "clearingTimes", arg0, arg1, arg2)
	return *ret0, err
}

// ClearingTimes is a free data retrieval call binding the contract method 0xb04c0239.
//
// Solidity: function clearingTimes(address , address , uint256 ) constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) ClearingTimes(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _DutchExchange.Contract.ClearingTimes(&_DutchExchange.CallOpts, arg0, arg1, arg2)
}

// ClearingTimes is a free data retrieval call binding the contract method 0xb04c0239.
//
// Solidity: function clearingTimes(address , address , uint256 ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) ClearingTimes(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _DutchExchange.Contract.ClearingTimes(&_DutchExchange.CallOpts, arg0, arg1, arg2)
}

// ClosingPrices is a free data retrieval call binding the contract method 0xebcc0de1.
//
// Solidity: function closingPrices(address , address , uint256 ) constant returns(uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeCaller) ClosingPrices(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (struct {
	Num *big.Int
	Den *big.Int
}, error) {
	ret := new(struct {
		Num *big.Int
		Den *big.Int
	})
	out := ret
	err := _DutchExchange.contract.Call(opts, out, "closingPrices", arg0, arg1, arg2)
	return *ret, err
}

// ClosingPrices is a free data retrieval call binding the contract method 0xebcc0de1.
//
// Solidity: function closingPrices(address , address , uint256 ) constant returns(uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeSession) ClosingPrices(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (struct {
	Num *big.Int
	Den *big.Int
}, error) {
	return _DutchExchange.Contract.ClosingPrices(&_DutchExchange.CallOpts, arg0, arg1, arg2)
}

// ClosingPrices is a free data retrieval call binding the contract method 0xebcc0de1.
//
// Solidity: function closingPrices(address , address , uint256 ) constant returns(uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeCallerSession) ClosingPrices(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (struct {
	Num *big.Int
	Den *big.Int
}, error) {
	return _DutchExchange.Contract.ClosingPrices(&_DutchExchange.CallOpts, arg0, arg1, arg2)
}

// EthToken is a free data retrieval call binding the contract method 0x7bf1a627.
//
// Solidity: function ethToken() constant returns(address)
func (_DutchExchange *DutchExchangeCaller) EthToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "ethToken")
	return *ret0, err
}

// EthToken is a free data retrieval call binding the contract method 0x7bf1a627.
//
// Solidity: function ethToken() constant returns(address)
func (_DutchExchange *DutchExchangeSession) EthToken() (common.Address, error) {
	return _DutchExchange.Contract.EthToken(&_DutchExchange.CallOpts)
}

// EthToken is a free data retrieval call binding the contract method 0x7bf1a627.
//
// Solidity: function ethToken() constant returns(address)
func (_DutchExchange *DutchExchangeCallerSession) EthToken() (common.Address, error) {
	return _DutchExchange.Contract.EthToken(&_DutchExchange.CallOpts)
}

// EthUSDOracle is a free data retrieval call binding the contract method 0x706eb3ab.
//
// Solidity: function ethUSDOracle() constant returns(address)
func (_DutchExchange *DutchExchangeCaller) EthUSDOracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "ethUSDOracle")
	return *ret0, err
}

// EthUSDOracle is a free data retrieval call binding the contract method 0x706eb3ab.
//
// Solidity: function ethUSDOracle() constant returns(address)
func (_DutchExchange *DutchExchangeSession) EthUSDOracle() (common.Address, error) {
	return _DutchExchange.Contract.EthUSDOracle(&_DutchExchange.CallOpts)
}

// EthUSDOracle is a free data retrieval call binding the contract method 0x706eb3ab.
//
// Solidity: function ethUSDOracle() constant returns(address)
func (_DutchExchange *DutchExchangeCallerSession) EthUSDOracle() (common.Address, error) {
	return _DutchExchange.Contract.EthUSDOracle(&_DutchExchange.CallOpts)
}

// ExtraTokens is a free data retrieval call binding the contract method 0xf79710fd.
//
// Solidity: function extraTokens(address , address , uint256 ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) ExtraTokens(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "extraTokens", arg0, arg1, arg2)
	return *ret0, err
}

// ExtraTokens is a free data retrieval call binding the contract method 0xf79710fd.
//
// Solidity: function extraTokens(address , address , uint256 ) constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) ExtraTokens(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _DutchExchange.Contract.ExtraTokens(&_DutchExchange.CallOpts, arg0, arg1, arg2)
}

// ExtraTokens is a free data retrieval call binding the contract method 0xf79710fd.
//
// Solidity: function extraTokens(address , address , uint256 ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) ExtraTokens(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _DutchExchange.Contract.ExtraTokens(&_DutchExchange.CallOpts, arg0, arg1, arg2)
}

// FrtToken is a free data retrieval call binding the contract method 0x8261eb1b.
//
// Solidity: function frtToken() constant returns(address)
func (_DutchExchange *DutchExchangeCaller) FrtToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "frtToken")
	return *ret0, err
}

// FrtToken is a free data retrieval call binding the contract method 0x8261eb1b.
//
// Solidity: function frtToken() constant returns(address)
func (_DutchExchange *DutchExchangeSession) FrtToken() (common.Address, error) {
	return _DutchExchange.Contract.FrtToken(&_DutchExchange.CallOpts)
}

// FrtToken is a free data retrieval call binding the contract method 0x8261eb1b.
//
// Solidity: function frtToken() constant returns(address)
func (_DutchExchange *DutchExchangeCallerSession) FrtToken() (common.Address, error) {
	return _DutchExchange.Contract.FrtToken(&_DutchExchange.CallOpts)
}

// GetApprovedAddressesOfList is a free data retrieval call binding the contract method 0x02ffc0b0.
//
// Solidity: function getApprovedAddressesOfList(address[] addressesToCheck) constant returns(bool[])
func (_DutchExchange *DutchExchangeCaller) GetApprovedAddressesOfList(opts *bind.CallOpts, addressesToCheck []common.Address) ([]bool, error) {
	var (
		ret0 = new([]bool)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "getApprovedAddressesOfList", addressesToCheck)
	return *ret0, err
}

// GetApprovedAddressesOfList is a free data retrieval call binding the contract method 0x02ffc0b0.
//
// Solidity: function getApprovedAddressesOfList(address[] addressesToCheck) constant returns(bool[])
func (_DutchExchange *DutchExchangeSession) GetApprovedAddressesOfList(addressesToCheck []common.Address) ([]bool, error) {
	return _DutchExchange.Contract.GetApprovedAddressesOfList(&_DutchExchange.CallOpts, addressesToCheck)
}

// GetApprovedAddressesOfList is a free data retrieval call binding the contract method 0x02ffc0b0.
//
// Solidity: function getApprovedAddressesOfList(address[] addressesToCheck) constant returns(bool[])
func (_DutchExchange *DutchExchangeCallerSession) GetApprovedAddressesOfList(addressesToCheck []common.Address) ([]bool, error) {
	return _DutchExchange.Contract.GetApprovedAddressesOfList(&_DutchExchange.CallOpts, addressesToCheck)
}

// GetAuctionIndex is a free data retrieval call binding the contract method 0x14584a9d.
//
// Solidity: function getAuctionIndex(address token1, address token2) constant returns(uint256 auctionIndex)
func (_DutchExchange *DutchExchangeCaller) GetAuctionIndex(opts *bind.CallOpts, token1 common.Address, token2 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "getAuctionIndex", token1, token2)
	return *ret0, err
}

// GetAuctionIndex is a free data retrieval call binding the contract method 0x14584a9d.
//
// Solidity: function getAuctionIndex(address token1, address token2) constant returns(uint256 auctionIndex)
func (_DutchExchange *DutchExchangeSession) GetAuctionIndex(token1 common.Address, token2 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.GetAuctionIndex(&_DutchExchange.CallOpts, token1, token2)
}

// GetAuctionIndex is a free data retrieval call binding the contract method 0x14584a9d.
//
// Solidity: function getAuctionIndex(address token1, address token2) constant returns(uint256 auctionIndex)
func (_DutchExchange *DutchExchangeCallerSession) GetAuctionIndex(token1 common.Address, token2 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.GetAuctionIndex(&_DutchExchange.CallOpts, token1, token2)
}

// GetAuctionStart is a free data retrieval call binding the contract method 0xdae595e5.
//
// Solidity: function getAuctionStart(address token1, address token2) constant returns(uint256 auctionStart)
func (_DutchExchange *DutchExchangeCaller) GetAuctionStart(opts *bind.CallOpts, token1 common.Address, token2 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "getAuctionStart", token1, token2)
	return *ret0, err
}

// GetAuctionStart is a free data retrieval call binding the contract method 0xdae595e5.
//
// Solidity: function getAuctionStart(address token1, address token2) constant returns(uint256 auctionStart)
func (_DutchExchange *DutchExchangeSession) GetAuctionStart(token1 common.Address, token2 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.GetAuctionStart(&_DutchExchange.CallOpts, token1, token2)
}

// GetAuctionStart is a free data retrieval call binding the contract method 0xdae595e5.
//
// Solidity: function getAuctionStart(address token1, address token2) constant returns(uint256 auctionStart)
func (_DutchExchange *DutchExchangeCallerSession) GetAuctionStart(token1 common.Address, token2 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.GetAuctionStart(&_DutchExchange.CallOpts, token1, token2)
}

// GetClearingTime is a free data retrieval call binding the contract method 0x7420a32f.
//
// Solidity: function getClearingTime(address token1, address token2, uint256 auctionIndex) constant returns(uint256 time)
func (_DutchExchange *DutchExchangeCaller) GetClearingTime(opts *bind.CallOpts, token1 common.Address, token2 common.Address, auctionIndex *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "getClearingTime", token1, token2, auctionIndex)
	return *ret0, err
}

// GetClearingTime is a free data retrieval call binding the contract method 0x7420a32f.
//
// Solidity: function getClearingTime(address token1, address token2, uint256 auctionIndex) constant returns(uint256 time)
func (_DutchExchange *DutchExchangeSession) GetClearingTime(token1 common.Address, token2 common.Address, auctionIndex *big.Int) (*big.Int, error) {
	return _DutchExchange.Contract.GetClearingTime(&_DutchExchange.CallOpts, token1, token2, auctionIndex)
}

// GetClearingTime is a free data retrieval call binding the contract method 0x7420a32f.
//
// Solidity: function getClearingTime(address token1, address token2, uint256 auctionIndex) constant returns(uint256 time)
func (_DutchExchange *DutchExchangeCallerSession) GetClearingTime(token1 common.Address, token2 common.Address, auctionIndex *big.Int) (*big.Int, error) {
	return _DutchExchange.Contract.GetClearingTime(&_DutchExchange.CallOpts, token1, token2, auctionIndex)
}

// GetCurrentAuctionPrice is a free data retrieval call binding the contract method 0xfdab1b7b.
//
// Solidity: function getCurrentAuctionPrice(address sellToken, address buyToken, uint256 auctionIndex) constant returns(uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeCaller) GetCurrentAuctionPrice(opts *bind.CallOpts, sellToken common.Address, buyToken common.Address, auctionIndex *big.Int) (struct {
	Num *big.Int
	Den *big.Int
}, error) {
	ret := new(struct {
		Num *big.Int
		Den *big.Int
	})
	out := ret
	err := _DutchExchange.contract.Call(opts, out, "getCurrentAuctionPrice", sellToken, buyToken, auctionIndex)
	return *ret, err
}

// GetCurrentAuctionPrice is a free data retrieval call binding the contract method 0xfdab1b7b.
//
// Solidity: function getCurrentAuctionPrice(address sellToken, address buyToken, uint256 auctionIndex) constant returns(uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeSession) GetCurrentAuctionPrice(sellToken common.Address, buyToken common.Address, auctionIndex *big.Int) (struct {
	Num *big.Int
	Den *big.Int
}, error) {
	return _DutchExchange.Contract.GetCurrentAuctionPrice(&_DutchExchange.CallOpts, sellToken, buyToken, auctionIndex)
}

// GetCurrentAuctionPrice is a free data retrieval call binding the contract method 0xfdab1b7b.
//
// Solidity: function getCurrentAuctionPrice(address sellToken, address buyToken, uint256 auctionIndex) constant returns(uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeCallerSession) GetCurrentAuctionPrice(sellToken common.Address, buyToken common.Address, auctionIndex *big.Int) (struct {
	Num *big.Int
	Den *big.Int
}, error) {
	return _DutchExchange.Contract.GetCurrentAuctionPrice(&_DutchExchange.CallOpts, sellToken, buyToken, auctionIndex)
}

// GetFeeRatio is a free data retrieval call binding the contract method 0xedd0b5cb.
//
// Solidity: function getFeeRatio(address user) constant returns(uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeCaller) GetFeeRatio(opts *bind.CallOpts, user common.Address) (struct {
	Num *big.Int
	Den *big.Int
}, error) {
	ret := new(struct {
		Num *big.Int
		Den *big.Int
	})
	out := ret
	err := _DutchExchange.contract.Call(opts, out, "getFeeRatio", user)
	return *ret, err
}

// GetFeeRatio is a free data retrieval call binding the contract method 0xedd0b5cb.
//
// Solidity: function getFeeRatio(address user) constant returns(uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeSession) GetFeeRatio(user common.Address) (struct {
	Num *big.Int
	Den *big.Int
}, error) {
	return _DutchExchange.Contract.GetFeeRatio(&_DutchExchange.CallOpts, user)
}

// GetFeeRatio is a free data retrieval call binding the contract method 0xedd0b5cb.
//
// Solidity: function getFeeRatio(address user) constant returns(uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeCallerSession) GetFeeRatio(user common.Address) (struct {
	Num *big.Int
	Den *big.Int
}, error) {
	return _DutchExchange.Contract.GetFeeRatio(&_DutchExchange.CallOpts, user)
}

// GetMasterCopy is a free data retrieval call binding the contract method 0x04e80e90.
//
// Solidity: function getMasterCopy() constant returns(address)
func (_DutchExchange *DutchExchangeCaller) GetMasterCopy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "getMasterCopy")
	return *ret0, err
}

// GetMasterCopy is a free data retrieval call binding the contract method 0x04e80e90.
//
// Solidity: function getMasterCopy() constant returns(address)
func (_DutchExchange *DutchExchangeSession) GetMasterCopy() (common.Address, error) {
	return _DutchExchange.Contract.GetMasterCopy(&_DutchExchange.CallOpts)
}

// GetMasterCopy is a free data retrieval call binding the contract method 0x04e80e90.
//
// Solidity: function getMasterCopy() constant returns(address)
func (_DutchExchange *DutchExchangeCallerSession) GetMasterCopy() (common.Address, error) {
	return _DutchExchange.Contract.GetMasterCopy(&_DutchExchange.CallOpts)
}

// GetPriceInPastAuction is a free data retrieval call binding the contract method 0x4bf8e7a2.
//
// Solidity: function getPriceInPastAuction(address token1, address token2, uint256 auctionIndex) constant returns(uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeCaller) GetPriceInPastAuction(opts *bind.CallOpts, token1 common.Address, token2 common.Address, auctionIndex *big.Int) (struct {
	Num *big.Int
	Den *big.Int
}, error) {
	ret := new(struct {
		Num *big.Int
		Den *big.Int
	})
	out := ret
	err := _DutchExchange.contract.Call(opts, out, "getPriceInPastAuction", token1, token2, auctionIndex)
	return *ret, err
}

// GetPriceInPastAuction is a free data retrieval call binding the contract method 0x4bf8e7a2.
//
// Solidity: function getPriceInPastAuction(address token1, address token2, uint256 auctionIndex) constant returns(uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeSession) GetPriceInPastAuction(token1 common.Address, token2 common.Address, auctionIndex *big.Int) (struct {
	Num *big.Int
	Den *big.Int
}, error) {
	return _DutchExchange.Contract.GetPriceInPastAuction(&_DutchExchange.CallOpts, token1, token2, auctionIndex)
}

// GetPriceInPastAuction is a free data retrieval call binding the contract method 0x4bf8e7a2.
//
// Solidity: function getPriceInPastAuction(address token1, address token2, uint256 auctionIndex) constant returns(uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeCallerSession) GetPriceInPastAuction(token1 common.Address, token2 common.Address, auctionIndex *big.Int) (struct {
	Num *big.Int
	Den *big.Int
}, error) {
	return _DutchExchange.Contract.GetPriceInPastAuction(&_DutchExchange.CallOpts, token1, token2, auctionIndex)
}

// GetPriceOfTokenInLastAuction is a free data retrieval call binding the contract method 0xf41d97fc.
//
// Solidity: function getPriceOfTokenInLastAuction(address token) constant returns(uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeCaller) GetPriceOfTokenInLastAuction(opts *bind.CallOpts, token common.Address) (struct {
	Num *big.Int
	Den *big.Int
}, error) {
	ret := new(struct {
		Num *big.Int
		Den *big.Int
	})
	out := ret
	err := _DutchExchange.contract.Call(opts, out, "getPriceOfTokenInLastAuction", token)
	return *ret, err
}

// GetPriceOfTokenInLastAuction is a free data retrieval call binding the contract method 0xf41d97fc.
//
// Solidity: function getPriceOfTokenInLastAuction(address token) constant returns(uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeSession) GetPriceOfTokenInLastAuction(token common.Address) (struct {
	Num *big.Int
	Den *big.Int
}, error) {
	return _DutchExchange.Contract.GetPriceOfTokenInLastAuction(&_DutchExchange.CallOpts, token)
}

// GetPriceOfTokenInLastAuction is a free data retrieval call binding the contract method 0xf41d97fc.
//
// Solidity: function getPriceOfTokenInLastAuction(address token) constant returns(uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeCallerSession) GetPriceOfTokenInLastAuction(token common.Address) (struct {
	Num *big.Int
	Den *big.Int
}, error) {
	return _DutchExchange.Contract.GetPriceOfTokenInLastAuction(&_DutchExchange.CallOpts, token)
}

// GetTokenOrder is a free data retrieval call binding the contract method 0x00599e65.
//
// Solidity: function getTokenOrder(address token1, address token2) constant returns(address, address)
func (_DutchExchange *DutchExchangeCaller) GetTokenOrder(opts *bind.CallOpts, token1 common.Address, token2 common.Address) (common.Address, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DutchExchange.contract.Call(opts, out, "getTokenOrder", token1, token2)
	return *ret0, *ret1, err
}

// GetTokenOrder is a free data retrieval call binding the contract method 0x00599e65.
//
// Solidity: function getTokenOrder(address token1, address token2) constant returns(address, address)
func (_DutchExchange *DutchExchangeSession) GetTokenOrder(token1 common.Address, token2 common.Address) (common.Address, common.Address, error) {
	return _DutchExchange.Contract.GetTokenOrder(&_DutchExchange.CallOpts, token1, token2)
}

// GetTokenOrder is a free data retrieval call binding the contract method 0x00599e65.
//
// Solidity: function getTokenOrder(address token1, address token2) constant returns(address, address)
func (_DutchExchange *DutchExchangeCallerSession) GetTokenOrder(token1 common.Address, token2 common.Address) (common.Address, common.Address, error) {
	return _DutchExchange.Contract.GetTokenOrder(&_DutchExchange.CallOpts, token1, token2)
}

// GetUnclaimedBuyerFunds is a free data retrieval call binding the contract method 0xdf6af7d1.
//
// Solidity: function getUnclaimedBuyerFunds(address sellToken, address buyToken, address user, uint256 auctionIndex) constant returns(uint256 unclaimedBuyerFunds, uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeCaller) GetUnclaimedBuyerFunds(opts *bind.CallOpts, sellToken common.Address, buyToken common.Address, user common.Address, auctionIndex *big.Int) (struct {
	UnclaimedBuyerFunds *big.Int
	Num                 *big.Int
	Den                 *big.Int
}, error) {
	ret := new(struct {
		UnclaimedBuyerFunds *big.Int
		Num                 *big.Int
		Den                 *big.Int
	})
	out := ret
	err := _DutchExchange.contract.Call(opts, out, "getUnclaimedBuyerFunds", sellToken, buyToken, user, auctionIndex)
	return *ret, err
}

// GetUnclaimedBuyerFunds is a free data retrieval call binding the contract method 0xdf6af7d1.
//
// Solidity: function getUnclaimedBuyerFunds(address sellToken, address buyToken, address user, uint256 auctionIndex) constant returns(uint256 unclaimedBuyerFunds, uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeSession) GetUnclaimedBuyerFunds(sellToken common.Address, buyToken common.Address, user common.Address, auctionIndex *big.Int) (struct {
	UnclaimedBuyerFunds *big.Int
	Num                 *big.Int
	Den                 *big.Int
}, error) {
	return _DutchExchange.Contract.GetUnclaimedBuyerFunds(&_DutchExchange.CallOpts, sellToken, buyToken, user, auctionIndex)
}

// GetUnclaimedBuyerFunds is a free data retrieval call binding the contract method 0xdf6af7d1.
//
// Solidity: function getUnclaimedBuyerFunds(address sellToken, address buyToken, address user, uint256 auctionIndex) constant returns(uint256 unclaimedBuyerFunds, uint256 num, uint256 den)
func (_DutchExchange *DutchExchangeCallerSession) GetUnclaimedBuyerFunds(sellToken common.Address, buyToken common.Address, user common.Address, auctionIndex *big.Int) (struct {
	UnclaimedBuyerFunds *big.Int
	Num                 *big.Int
	Den                 *big.Int
}, error) {
	return _DutchExchange.Contract.GetUnclaimedBuyerFunds(&_DutchExchange.CallOpts, sellToken, buyToken, user, auctionIndex)
}

// LatestAuctionIndices is a free data retrieval call binding the contract method 0xfac7abe3.
//
// Solidity: function latestAuctionIndices(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) LatestAuctionIndices(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "latestAuctionIndices", arg0, arg1)
	return *ret0, err
}

// LatestAuctionIndices is a free data retrieval call binding the contract method 0xfac7abe3.
//
// Solidity: function latestAuctionIndices(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) LatestAuctionIndices(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.LatestAuctionIndices(&_DutchExchange.CallOpts, arg0, arg1)
}

// LatestAuctionIndices is a free data retrieval call binding the contract method 0xfac7abe3.
//
// Solidity: function latestAuctionIndices(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) LatestAuctionIndices(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.LatestAuctionIndices(&_DutchExchange.CallOpts, arg0, arg1)
}

// MasterCopy is a free data retrieval call binding the contract method 0xa619486e.
//
// Solidity: function masterCopy() constant returns(address)
func (_DutchExchange *DutchExchangeCaller) MasterCopy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "masterCopy")
	return *ret0, err
}

// MasterCopy is a free data retrieval call binding the contract method 0xa619486e.
//
// Solidity: function masterCopy() constant returns(address)
func (_DutchExchange *DutchExchangeSession) MasterCopy() (common.Address, error) {
	return _DutchExchange.Contract.MasterCopy(&_DutchExchange.CallOpts)
}

// MasterCopy is a free data retrieval call binding the contract method 0xa619486e.
//
// Solidity: function masterCopy() constant returns(address)
func (_DutchExchange *DutchExchangeCallerSession) MasterCopy() (common.Address, error) {
	return _DutchExchange.Contract.MasterCopy(&_DutchExchange.CallOpts)
}

// MasterCopyCountdown is a free data retrieval call binding the contract method 0x0e7c0f80.
//
// Solidity: function masterCopyCountdown() constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) MasterCopyCountdown(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "masterCopyCountdown")
	return *ret0, err
}

// MasterCopyCountdown is a free data retrieval call binding the contract method 0x0e7c0f80.
//
// Solidity: function masterCopyCountdown() constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) MasterCopyCountdown() (*big.Int, error) {
	return _DutchExchange.Contract.MasterCopyCountdown(&_DutchExchange.CallOpts)
}

// MasterCopyCountdown is a free data retrieval call binding the contract method 0x0e7c0f80.
//
// Solidity: function masterCopyCountdown() constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) MasterCopyCountdown() (*big.Int, error) {
	return _DutchExchange.Contract.MasterCopyCountdown(&_DutchExchange.CallOpts)
}

// Min is a free data retrieval call binding the contract method 0x7ae2b5c7.
//
// Solidity: function min(uint256 a, uint256 b) constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) Min(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "min", a, b)
	return *ret0, err
}

// Min is a free data retrieval call binding the contract method 0x7ae2b5c7.
//
// Solidity: function min(uint256 a, uint256 b) constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) Min(a *big.Int, b *big.Int) (*big.Int, error) {
	return _DutchExchange.Contract.Min(&_DutchExchange.CallOpts, a, b)
}

// Min is a free data retrieval call binding the contract method 0x7ae2b5c7.
//
// Solidity: function min(uint256 a, uint256 b) constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) Min(a *big.Int, b *big.Int) (*big.Int, error) {
	return _DutchExchange.Contract.Min(&_DutchExchange.CallOpts, a, b)
}

// Mul is a free data retrieval call binding the contract method 0xc8a4ac9c.
//
// Solidity: function mul(uint256 a, uint256 b) constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) Mul(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "mul", a, b)
	return *ret0, err
}

// Mul is a free data retrieval call binding the contract method 0xc8a4ac9c.
//
// Solidity: function mul(uint256 a, uint256 b) constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) Mul(a *big.Int, b *big.Int) (*big.Int, error) {
	return _DutchExchange.Contract.Mul(&_DutchExchange.CallOpts, a, b)
}

// Mul is a free data retrieval call binding the contract method 0xc8a4ac9c.
//
// Solidity: function mul(uint256 a, uint256 b) constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) Mul(a *big.Int, b *big.Int) (*big.Int, error) {
	return _DutchExchange.Contract.Mul(&_DutchExchange.CallOpts, a, b)
}

// NewMasterCopy is a free data retrieval call binding the contract method 0x6ea68360.
//
// Solidity: function newMasterCopy() constant returns(address)
func (_DutchExchange *DutchExchangeCaller) NewMasterCopy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "newMasterCopy")
	return *ret0, err
}

// NewMasterCopy is a free data retrieval call binding the contract method 0x6ea68360.
//
// Solidity: function newMasterCopy() constant returns(address)
func (_DutchExchange *DutchExchangeSession) NewMasterCopy() (common.Address, error) {
	return _DutchExchange.Contract.NewMasterCopy(&_DutchExchange.CallOpts)
}

// NewMasterCopy is a free data retrieval call binding the contract method 0x6ea68360.
//
// Solidity: function newMasterCopy() constant returns(address)
func (_DutchExchange *DutchExchangeCallerSession) NewMasterCopy() (common.Address, error) {
	return _DutchExchange.Contract.NewMasterCopy(&_DutchExchange.CallOpts)
}

// NewProposalEthUSDOracle is a free data retrieval call binding the contract method 0xcd04ccfc.
//
// Solidity: function newProposalEthUSDOracle() constant returns(address)
func (_DutchExchange *DutchExchangeCaller) NewProposalEthUSDOracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "newProposalEthUSDOracle")
	return *ret0, err
}

// NewProposalEthUSDOracle is a free data retrieval call binding the contract method 0xcd04ccfc.
//
// Solidity: function newProposalEthUSDOracle() constant returns(address)
func (_DutchExchange *DutchExchangeSession) NewProposalEthUSDOracle() (common.Address, error) {
	return _DutchExchange.Contract.NewProposalEthUSDOracle(&_DutchExchange.CallOpts)
}

// NewProposalEthUSDOracle is a free data retrieval call binding the contract method 0xcd04ccfc.
//
// Solidity: function newProposalEthUSDOracle() constant returns(address)
func (_DutchExchange *DutchExchangeCallerSession) NewProposalEthUSDOracle() (common.Address, error) {
	return _DutchExchange.Contract.NewProposalEthUSDOracle(&_DutchExchange.CallOpts)
}

// OracleInterfaceCountdown is a free data retrieval call binding the contract method 0xa48cef4a.
//
// Solidity: function oracleInterfaceCountdown() constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) OracleInterfaceCountdown(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "oracleInterfaceCountdown")
	return *ret0, err
}

// OracleInterfaceCountdown is a free data retrieval call binding the contract method 0xa48cef4a.
//
// Solidity: function oracleInterfaceCountdown() constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) OracleInterfaceCountdown() (*big.Int, error) {
	return _DutchExchange.Contract.OracleInterfaceCountdown(&_DutchExchange.CallOpts)
}

// OracleInterfaceCountdown is a free data retrieval call binding the contract method 0xa48cef4a.
//
// Solidity: function oracleInterfaceCountdown() constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) OracleInterfaceCountdown() (*big.Int, error) {
	return _DutchExchange.Contract.OracleInterfaceCountdown(&_DutchExchange.CallOpts)
}

// OwlToken is a free data retrieval call binding the contract method 0xcd94a2a4.
//
// Solidity: function owlToken() constant returns(address)
func (_DutchExchange *DutchExchangeCaller) OwlToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "owlToken")
	return *ret0, err
}

// OwlToken is a free data retrieval call binding the contract method 0xcd94a2a4.
//
// Solidity: function owlToken() constant returns(address)
func (_DutchExchange *DutchExchangeSession) OwlToken() (common.Address, error) {
	return _DutchExchange.Contract.OwlToken(&_DutchExchange.CallOpts)
}

// OwlToken is a free data retrieval call binding the contract method 0xcd94a2a4.
//
// Solidity: function owlToken() constant returns(address)
func (_DutchExchange *DutchExchangeCallerSession) OwlToken() (common.Address, error) {
	return _DutchExchange.Contract.OwlToken(&_DutchExchange.CallOpts)
}

// SafeToAdd is a free data retrieval call binding the contract method 0x4e30a66c.
//
// Solidity: function safeToAdd(uint256 a, uint256 b) constant returns(bool)
func (_DutchExchange *DutchExchangeCaller) SafeToAdd(opts *bind.CallOpts, a *big.Int, b *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "safeToAdd", a, b)
	return *ret0, err
}

// SafeToAdd is a free data retrieval call binding the contract method 0x4e30a66c.
//
// Solidity: function safeToAdd(uint256 a, uint256 b) constant returns(bool)
func (_DutchExchange *DutchExchangeSession) SafeToAdd(a *big.Int, b *big.Int) (bool, error) {
	return _DutchExchange.Contract.SafeToAdd(&_DutchExchange.CallOpts, a, b)
}

// SafeToAdd is a free data retrieval call binding the contract method 0x4e30a66c.
//
// Solidity: function safeToAdd(uint256 a, uint256 b) constant returns(bool)
func (_DutchExchange *DutchExchangeCallerSession) SafeToAdd(a *big.Int, b *big.Int) (bool, error) {
	return _DutchExchange.Contract.SafeToAdd(&_DutchExchange.CallOpts, a, b)
}

// SafeToMul is a free data retrieval call binding the contract method 0xcb10fa76.
//
// Solidity: function safeToMul(uint256 a, uint256 b) constant returns(bool)
func (_DutchExchange *DutchExchangeCaller) SafeToMul(opts *bind.CallOpts, a *big.Int, b *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "safeToMul", a, b)
	return *ret0, err
}

// SafeToMul is a free data retrieval call binding the contract method 0xcb10fa76.
//
// Solidity: function safeToMul(uint256 a, uint256 b) constant returns(bool)
func (_DutchExchange *DutchExchangeSession) SafeToMul(a *big.Int, b *big.Int) (bool, error) {
	return _DutchExchange.Contract.SafeToMul(&_DutchExchange.CallOpts, a, b)
}

// SafeToMul is a free data retrieval call binding the contract method 0xcb10fa76.
//
// Solidity: function safeToMul(uint256 a, uint256 b) constant returns(bool)
func (_DutchExchange *DutchExchangeCallerSession) SafeToMul(a *big.Int, b *big.Int) (bool, error) {
	return _DutchExchange.Contract.SafeToMul(&_DutchExchange.CallOpts, a, b)
}

// SafeToSub is a free data retrieval call binding the contract method 0xe31c71c4.
//
// Solidity: function safeToSub(uint256 a, uint256 b) constant returns(bool)
func (_DutchExchange *DutchExchangeCaller) SafeToSub(opts *bind.CallOpts, a *big.Int, b *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "safeToSub", a, b)
	return *ret0, err
}

// SafeToSub is a free data retrieval call binding the contract method 0xe31c71c4.
//
// Solidity: function safeToSub(uint256 a, uint256 b) constant returns(bool)
func (_DutchExchange *DutchExchangeSession) SafeToSub(a *big.Int, b *big.Int) (bool, error) {
	return _DutchExchange.Contract.SafeToSub(&_DutchExchange.CallOpts, a, b)
}

// SafeToSub is a free data retrieval call binding the contract method 0xe31c71c4.
//
// Solidity: function safeToSub(uint256 a, uint256 b) constant returns(bool)
func (_DutchExchange *DutchExchangeCallerSession) SafeToSub(a *big.Int, b *big.Int) (bool, error) {
	return _DutchExchange.Contract.SafeToSub(&_DutchExchange.CallOpts, a, b)
}

// SellVolumesCurrent is a free data retrieval call binding the contract method 0xb64c4905.
//
// Solidity: function sellVolumesCurrent(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) SellVolumesCurrent(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "sellVolumesCurrent", arg0, arg1)
	return *ret0, err
}

// SellVolumesCurrent is a free data retrieval call binding the contract method 0xb64c4905.
//
// Solidity: function sellVolumesCurrent(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) SellVolumesCurrent(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.SellVolumesCurrent(&_DutchExchange.CallOpts, arg0, arg1)
}

// SellVolumesCurrent is a free data retrieval call binding the contract method 0xb64c4905.
//
// Solidity: function sellVolumesCurrent(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) SellVolumesCurrent(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.SellVolumesCurrent(&_DutchExchange.CallOpts, arg0, arg1)
}

// SellVolumesNext is a free data retrieval call binding the contract method 0xb3c2083f.
//
// Solidity: function sellVolumesNext(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) SellVolumesNext(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "sellVolumesNext", arg0, arg1)
	return *ret0, err
}

// SellVolumesNext is a free data retrieval call binding the contract method 0xb3c2083f.
//
// Solidity: function sellVolumesNext(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) SellVolumesNext(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.SellVolumesNext(&_DutchExchange.CallOpts, arg0, arg1)
}

// SellVolumesNext is a free data retrieval call binding the contract method 0xb3c2083f.
//
// Solidity: function sellVolumesNext(address , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) SellVolumesNext(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.SellVolumesNext(&_DutchExchange.CallOpts, arg0, arg1)
}

// SellerBalances is a free data retrieval call binding the contract method 0xc1a21bf3.
//
// Solidity: function sellerBalances(address , address , uint256 , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) SellerBalances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "sellerBalances", arg0, arg1, arg2, arg3)
	return *ret0, err
}

// SellerBalances is a free data retrieval call binding the contract method 0xc1a21bf3.
//
// Solidity: function sellerBalances(address , address , uint256 , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) SellerBalances(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.SellerBalances(&_DutchExchange.CallOpts, arg0, arg1, arg2, arg3)
}

// SellerBalances is a free data retrieval call binding the contract method 0xc1a21bf3.
//
// Solidity: function sellerBalances(address , address , uint256 , address ) constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) SellerBalances(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 common.Address) (*big.Int, error) {
	return _DutchExchange.Contract.SellerBalances(&_DutchExchange.CallOpts, arg0, arg1, arg2, arg3)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) Sub(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "sub", a, b)
	return *ret0, err
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _DutchExchange.Contract.Sub(&_DutchExchange.CallOpts, a, b)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _DutchExchange.Contract.Sub(&_DutchExchange.CallOpts, a, b)
}

// ThresholdNewAuction is a free data retrieval call binding the contract method 0xee93114c.
//
// Solidity: function thresholdNewAuction() constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) ThresholdNewAuction(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "thresholdNewAuction")
	return *ret0, err
}

// ThresholdNewAuction is a free data retrieval call binding the contract method 0xee93114c.
//
// Solidity: function thresholdNewAuction() constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) ThresholdNewAuction() (*big.Int, error) {
	return _DutchExchange.Contract.ThresholdNewAuction(&_DutchExchange.CallOpts)
}

// ThresholdNewAuction is a free data retrieval call binding the contract method 0xee93114c.
//
// Solidity: function thresholdNewAuction() constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) ThresholdNewAuction() (*big.Int, error) {
	return _DutchExchange.Contract.ThresholdNewAuction(&_DutchExchange.CallOpts)
}

// ThresholdNewTokenPair is a free data retrieval call binding the contract method 0x6e6260fa.
//
// Solidity: function thresholdNewTokenPair() constant returns(uint256)
func (_DutchExchange *DutchExchangeCaller) ThresholdNewTokenPair(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DutchExchange.contract.Call(opts, out, "thresholdNewTokenPair")
	return *ret0, err
}

// ThresholdNewTokenPair is a free data retrieval call binding the contract method 0x6e6260fa.
//
// Solidity: function thresholdNewTokenPair() constant returns(uint256)
func (_DutchExchange *DutchExchangeSession) ThresholdNewTokenPair() (*big.Int, error) {
	return _DutchExchange.Contract.ThresholdNewTokenPair(&_DutchExchange.CallOpts)
}

// ThresholdNewTokenPair is a free data retrieval call binding the contract method 0x6e6260fa.
//
// Solidity: function thresholdNewTokenPair() constant returns(uint256)
func (_DutchExchange *DutchExchangeCallerSession) ThresholdNewTokenPair() (*big.Int, error) {
	return _DutchExchange.Contract.ThresholdNewTokenPair(&_DutchExchange.CallOpts)
}

// AddTokenPair is a paid mutator transaction binding the contract method 0xe9f8cd70.
//
// Solidity: function addTokenPair(address token1, address token2, uint256 token1Funding, uint256 token2Funding, uint256 initialClosingPriceNum, uint256 initialClosingPriceDen) returns()
func (_DutchExchange *DutchExchangeTransactor) AddTokenPair(opts *bind.TransactOpts, token1 common.Address, token2 common.Address, token1Funding *big.Int, token2Funding *big.Int, initialClosingPriceNum *big.Int, initialClosingPriceDen *big.Int) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "addTokenPair", token1, token2, token1Funding, token2Funding, initialClosingPriceNum, initialClosingPriceDen)
}

// AddTokenPair is a paid mutator transaction binding the contract method 0xe9f8cd70.
//
// Solidity: function addTokenPair(address token1, address token2, uint256 token1Funding, uint256 token2Funding, uint256 initialClosingPriceNum, uint256 initialClosingPriceDen) returns()
func (_DutchExchange *DutchExchangeSession) AddTokenPair(token1 common.Address, token2 common.Address, token1Funding *big.Int, token2Funding *big.Int, initialClosingPriceNum *big.Int, initialClosingPriceDen *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.AddTokenPair(&_DutchExchange.TransactOpts, token1, token2, token1Funding, token2Funding, initialClosingPriceNum, initialClosingPriceDen)
}

// AddTokenPair is a paid mutator transaction binding the contract method 0xe9f8cd70.
//
// Solidity: function addTokenPair(address token1, address token2, uint256 token1Funding, uint256 token2Funding, uint256 initialClosingPriceNum, uint256 initialClosingPriceDen) returns()
func (_DutchExchange *DutchExchangeTransactorSession) AddTokenPair(token1 common.Address, token2 common.Address, token1Funding *big.Int, token2Funding *big.Int, initialClosingPriceNum *big.Int, initialClosingPriceDen *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.AddTokenPair(&_DutchExchange.TransactOpts, token1, token2, token1Funding, token2Funding, initialClosingPriceNum, initialClosingPriceDen)
}

// ClaimAndWithdraw is a paid mutator transaction binding the contract method 0x06d58f2a.
//
// Solidity: function claimAndWithdraw(address sellToken, address buyToken, address user, uint256 auctionIndex, uint256 amount) returns(uint256 returned, uint256 frtsIssued, uint256 newBal)
func (_DutchExchange *DutchExchangeTransactor) ClaimAndWithdraw(opts *bind.TransactOpts, sellToken common.Address, buyToken common.Address, user common.Address, auctionIndex *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "claimAndWithdraw", sellToken, buyToken, user, auctionIndex, amount)
}

// ClaimAndWithdraw is a paid mutator transaction binding the contract method 0x06d58f2a.
//
// Solidity: function claimAndWithdraw(address sellToken, address buyToken, address user, uint256 auctionIndex, uint256 amount) returns(uint256 returned, uint256 frtsIssued, uint256 newBal)
func (_DutchExchange *DutchExchangeSession) ClaimAndWithdraw(sellToken common.Address, buyToken common.Address, user common.Address, auctionIndex *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.ClaimAndWithdraw(&_DutchExchange.TransactOpts, sellToken, buyToken, user, auctionIndex, amount)
}

// ClaimAndWithdraw is a paid mutator transaction binding the contract method 0x06d58f2a.
//
// Solidity: function claimAndWithdraw(address sellToken, address buyToken, address user, uint256 auctionIndex, uint256 amount) returns(uint256 returned, uint256 frtsIssued, uint256 newBal)
func (_DutchExchange *DutchExchangeTransactorSession) ClaimAndWithdraw(sellToken common.Address, buyToken common.Address, user common.Address, auctionIndex *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.ClaimAndWithdraw(&_DutchExchange.TransactOpts, sellToken, buyToken, user, auctionIndex, amount)
}

// ClaimAndWithdrawTokensFromSeveralAuctionsAsBuyer is a paid mutator transaction binding the contract method 0x0c57cfba.
//
// Solidity: function claimAndWithdrawTokensFromSeveralAuctionsAsBuyer(address[] auctionSellTokens, address[] auctionBuyTokens, uint256[] auctionIndices) returns(uint256[], uint256 frtsIssued)
func (_DutchExchange *DutchExchangeTransactor) ClaimAndWithdrawTokensFromSeveralAuctionsAsBuyer(opts *bind.TransactOpts, auctionSellTokens []common.Address, auctionBuyTokens []common.Address, auctionIndices []*big.Int) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "claimAndWithdrawTokensFromSeveralAuctionsAsBuyer", auctionSellTokens, auctionBuyTokens, auctionIndices)
}

// ClaimAndWithdrawTokensFromSeveralAuctionsAsBuyer is a paid mutator transaction binding the contract method 0x0c57cfba.
//
// Solidity: function claimAndWithdrawTokensFromSeveralAuctionsAsBuyer(address[] auctionSellTokens, address[] auctionBuyTokens, uint256[] auctionIndices) returns(uint256[], uint256 frtsIssued)
func (_DutchExchange *DutchExchangeSession) ClaimAndWithdrawTokensFromSeveralAuctionsAsBuyer(auctionSellTokens []common.Address, auctionBuyTokens []common.Address, auctionIndices []*big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.ClaimAndWithdrawTokensFromSeveralAuctionsAsBuyer(&_DutchExchange.TransactOpts, auctionSellTokens, auctionBuyTokens, auctionIndices)
}

// ClaimAndWithdrawTokensFromSeveralAuctionsAsBuyer is a paid mutator transaction binding the contract method 0x0c57cfba.
//
// Solidity: function claimAndWithdrawTokensFromSeveralAuctionsAsBuyer(address[] auctionSellTokens, address[] auctionBuyTokens, uint256[] auctionIndices) returns(uint256[], uint256 frtsIssued)
func (_DutchExchange *DutchExchangeTransactorSession) ClaimAndWithdrawTokensFromSeveralAuctionsAsBuyer(auctionSellTokens []common.Address, auctionBuyTokens []common.Address, auctionIndices []*big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.ClaimAndWithdrawTokensFromSeveralAuctionsAsBuyer(&_DutchExchange.TransactOpts, auctionSellTokens, auctionBuyTokens, auctionIndices)
}

// ClaimAndWithdrawTokensFromSeveralAuctionsAsSeller is a paid mutator transaction binding the contract method 0x82a57b43.
//
// Solidity: function claimAndWithdrawTokensFromSeveralAuctionsAsSeller(address[] auctionSellTokens, address[] auctionBuyTokens, uint256[] auctionIndices) returns(uint256[], uint256 frtsIssued)
func (_DutchExchange *DutchExchangeTransactor) ClaimAndWithdrawTokensFromSeveralAuctionsAsSeller(opts *bind.TransactOpts, auctionSellTokens []common.Address, auctionBuyTokens []common.Address, auctionIndices []*big.Int) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "claimAndWithdrawTokensFromSeveralAuctionsAsSeller", auctionSellTokens, auctionBuyTokens, auctionIndices)
}

// ClaimAndWithdrawTokensFromSeveralAuctionsAsSeller is a paid mutator transaction binding the contract method 0x82a57b43.
//
// Solidity: function claimAndWithdrawTokensFromSeveralAuctionsAsSeller(address[] auctionSellTokens, address[] auctionBuyTokens, uint256[] auctionIndices) returns(uint256[], uint256 frtsIssued)
func (_DutchExchange *DutchExchangeSession) ClaimAndWithdrawTokensFromSeveralAuctionsAsSeller(auctionSellTokens []common.Address, auctionBuyTokens []common.Address, auctionIndices []*big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.ClaimAndWithdrawTokensFromSeveralAuctionsAsSeller(&_DutchExchange.TransactOpts, auctionSellTokens, auctionBuyTokens, auctionIndices)
}

// ClaimAndWithdrawTokensFromSeveralAuctionsAsSeller is a paid mutator transaction binding the contract method 0x82a57b43.
//
// Solidity: function claimAndWithdrawTokensFromSeveralAuctionsAsSeller(address[] auctionSellTokens, address[] auctionBuyTokens, uint256[] auctionIndices) returns(uint256[], uint256 frtsIssued)
func (_DutchExchange *DutchExchangeTransactorSession) ClaimAndWithdrawTokensFromSeveralAuctionsAsSeller(auctionSellTokens []common.Address, auctionBuyTokens []common.Address, auctionIndices []*big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.ClaimAndWithdrawTokensFromSeveralAuctionsAsSeller(&_DutchExchange.TransactOpts, auctionSellTokens, auctionBuyTokens, auctionIndices)
}

// ClaimBuyerFunds is a paid mutator transaction binding the contract method 0xb0293850.
//
// Solidity: function claimBuyerFunds(address sellToken, address buyToken, address user, uint256 auctionIndex) returns(uint256 returned, uint256 frtsIssued)
func (_DutchExchange *DutchExchangeTransactor) ClaimBuyerFunds(opts *bind.TransactOpts, sellToken common.Address, buyToken common.Address, user common.Address, auctionIndex *big.Int) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "claimBuyerFunds", sellToken, buyToken, user, auctionIndex)
}

// ClaimBuyerFunds is a paid mutator transaction binding the contract method 0xb0293850.
//
// Solidity: function claimBuyerFunds(address sellToken, address buyToken, address user, uint256 auctionIndex) returns(uint256 returned, uint256 frtsIssued)
func (_DutchExchange *DutchExchangeSession) ClaimBuyerFunds(sellToken common.Address, buyToken common.Address, user common.Address, auctionIndex *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.ClaimBuyerFunds(&_DutchExchange.TransactOpts, sellToken, buyToken, user, auctionIndex)
}

// ClaimBuyerFunds is a paid mutator transaction binding the contract method 0xb0293850.
//
// Solidity: function claimBuyerFunds(address sellToken, address buyToken, address user, uint256 auctionIndex) returns(uint256 returned, uint256 frtsIssued)
func (_DutchExchange *DutchExchangeTransactorSession) ClaimBuyerFunds(sellToken common.Address, buyToken common.Address, user common.Address, auctionIndex *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.ClaimBuyerFunds(&_DutchExchange.TransactOpts, sellToken, buyToken, user, auctionIndex)
}

// ClaimSellerFunds is a paid mutator transaction binding the contract method 0x65054e55.
//
// Solidity: function claimSellerFunds(address sellToken, address buyToken, address user, uint256 auctionIndex) returns(uint256 returned, uint256 frtsIssued)
func (_DutchExchange *DutchExchangeTransactor) ClaimSellerFunds(opts *bind.TransactOpts, sellToken common.Address, buyToken common.Address, user common.Address, auctionIndex *big.Int) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "claimSellerFunds", sellToken, buyToken, user, auctionIndex)
}

// ClaimSellerFunds is a paid mutator transaction binding the contract method 0x65054e55.
//
// Solidity: function claimSellerFunds(address sellToken, address buyToken, address user, uint256 auctionIndex) returns(uint256 returned, uint256 frtsIssued)
func (_DutchExchange *DutchExchangeSession) ClaimSellerFunds(sellToken common.Address, buyToken common.Address, user common.Address, auctionIndex *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.ClaimSellerFunds(&_DutchExchange.TransactOpts, sellToken, buyToken, user, auctionIndex)
}

// ClaimSellerFunds is a paid mutator transaction binding the contract method 0x65054e55.
//
// Solidity: function claimSellerFunds(address sellToken, address buyToken, address user, uint256 auctionIndex) returns(uint256 returned, uint256 frtsIssued)
func (_DutchExchange *DutchExchangeTransactorSession) ClaimSellerFunds(sellToken common.Address, buyToken common.Address, user common.Address, auctionIndex *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.ClaimSellerFunds(&_DutchExchange.TransactOpts, sellToken, buyToken, user, auctionIndex)
}

// ClaimTokensFromSeveralAuctionsAsBuyer is a paid mutator transaction binding the contract method 0xd3cc8d1c.
//
// Solidity: function claimTokensFromSeveralAuctionsAsBuyer(address[] auctionSellTokens, address[] auctionBuyTokens, uint256[] auctionIndices, address user) returns(uint256[], uint256[])
func (_DutchExchange *DutchExchangeTransactor) ClaimTokensFromSeveralAuctionsAsBuyer(opts *bind.TransactOpts, auctionSellTokens []common.Address, auctionBuyTokens []common.Address, auctionIndices []*big.Int, user common.Address) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "claimTokensFromSeveralAuctionsAsBuyer", auctionSellTokens, auctionBuyTokens, auctionIndices, user)
}

// ClaimTokensFromSeveralAuctionsAsBuyer is a paid mutator transaction binding the contract method 0xd3cc8d1c.
//
// Solidity: function claimTokensFromSeveralAuctionsAsBuyer(address[] auctionSellTokens, address[] auctionBuyTokens, uint256[] auctionIndices, address user) returns(uint256[], uint256[])
func (_DutchExchange *DutchExchangeSession) ClaimTokensFromSeveralAuctionsAsBuyer(auctionSellTokens []common.Address, auctionBuyTokens []common.Address, auctionIndices []*big.Int, user common.Address) (*types.Transaction, error) {
	return _DutchExchange.Contract.ClaimTokensFromSeveralAuctionsAsBuyer(&_DutchExchange.TransactOpts, auctionSellTokens, auctionBuyTokens, auctionIndices, user)
}

// ClaimTokensFromSeveralAuctionsAsBuyer is a paid mutator transaction binding the contract method 0xd3cc8d1c.
//
// Solidity: function claimTokensFromSeveralAuctionsAsBuyer(address[] auctionSellTokens, address[] auctionBuyTokens, uint256[] auctionIndices, address user) returns(uint256[], uint256[])
func (_DutchExchange *DutchExchangeTransactorSession) ClaimTokensFromSeveralAuctionsAsBuyer(auctionSellTokens []common.Address, auctionBuyTokens []common.Address, auctionIndices []*big.Int, user common.Address) (*types.Transaction, error) {
	return _DutchExchange.Contract.ClaimTokensFromSeveralAuctionsAsBuyer(&_DutchExchange.TransactOpts, auctionSellTokens, auctionBuyTokens, auctionIndices, user)
}

// ClaimTokensFromSeveralAuctionsAsSeller is a paid mutator transaction binding the contract method 0x7895dd21.
//
// Solidity: function claimTokensFromSeveralAuctionsAsSeller(address[] auctionSellTokens, address[] auctionBuyTokens, uint256[] auctionIndices, address user) returns(uint256[], uint256[])
func (_DutchExchange *DutchExchangeTransactor) ClaimTokensFromSeveralAuctionsAsSeller(opts *bind.TransactOpts, auctionSellTokens []common.Address, auctionBuyTokens []common.Address, auctionIndices []*big.Int, user common.Address) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "claimTokensFromSeveralAuctionsAsSeller", auctionSellTokens, auctionBuyTokens, auctionIndices, user)
}

// ClaimTokensFromSeveralAuctionsAsSeller is a paid mutator transaction binding the contract method 0x7895dd21.
//
// Solidity: function claimTokensFromSeveralAuctionsAsSeller(address[] auctionSellTokens, address[] auctionBuyTokens, uint256[] auctionIndices, address user) returns(uint256[], uint256[])
func (_DutchExchange *DutchExchangeSession) ClaimTokensFromSeveralAuctionsAsSeller(auctionSellTokens []common.Address, auctionBuyTokens []common.Address, auctionIndices []*big.Int, user common.Address) (*types.Transaction, error) {
	return _DutchExchange.Contract.ClaimTokensFromSeveralAuctionsAsSeller(&_DutchExchange.TransactOpts, auctionSellTokens, auctionBuyTokens, auctionIndices, user)
}

// ClaimTokensFromSeveralAuctionsAsSeller is a paid mutator transaction binding the contract method 0x7895dd21.
//
// Solidity: function claimTokensFromSeveralAuctionsAsSeller(address[] auctionSellTokens, address[] auctionBuyTokens, uint256[] auctionIndices, address user) returns(uint256[], uint256[])
func (_DutchExchange *DutchExchangeTransactorSession) ClaimTokensFromSeveralAuctionsAsSeller(auctionSellTokens []common.Address, auctionBuyTokens []common.Address, auctionIndices []*big.Int, user common.Address) (*types.Transaction, error) {
	return _DutchExchange.Contract.ClaimTokensFromSeveralAuctionsAsSeller(&_DutchExchange.TransactOpts, auctionSellTokens, auctionBuyTokens, auctionIndices, user)
}

// CloseTheoreticalClosedAuction is a paid mutator transaction binding the contract method 0x821b98f3.
//
// Solidity: function closeTheoreticalClosedAuction(address sellToken, address buyToken, uint256 auctionIndex) returns()
func (_DutchExchange *DutchExchangeTransactor) CloseTheoreticalClosedAuction(opts *bind.TransactOpts, sellToken common.Address, buyToken common.Address, auctionIndex *big.Int) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "closeTheoreticalClosedAuction", sellToken, buyToken, auctionIndex)
}

// CloseTheoreticalClosedAuction is a paid mutator transaction binding the contract method 0x821b98f3.
//
// Solidity: function closeTheoreticalClosedAuction(address sellToken, address buyToken, uint256 auctionIndex) returns()
func (_DutchExchange *DutchExchangeSession) CloseTheoreticalClosedAuction(sellToken common.Address, buyToken common.Address, auctionIndex *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.CloseTheoreticalClosedAuction(&_DutchExchange.TransactOpts, sellToken, buyToken, auctionIndex)
}

// CloseTheoreticalClosedAuction is a paid mutator transaction binding the contract method 0x821b98f3.
//
// Solidity: function closeTheoreticalClosedAuction(address sellToken, address buyToken, uint256 auctionIndex) returns()
func (_DutchExchange *DutchExchangeTransactorSession) CloseTheoreticalClosedAuction(sellToken common.Address, buyToken common.Address, auctionIndex *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.CloseTheoreticalClosedAuction(&_DutchExchange.TransactOpts, sellToken, buyToken, auctionIndex)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address tokenAddress, uint256 amount) returns(uint256)
func (_DutchExchange *DutchExchangeTransactor) Deposit(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "deposit", tokenAddress, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address tokenAddress, uint256 amount) returns(uint256)
func (_DutchExchange *DutchExchangeSession) Deposit(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.Deposit(&_DutchExchange.TransactOpts, tokenAddress, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address tokenAddress, uint256 amount) returns(uint256)
func (_DutchExchange *DutchExchangeTransactorSession) Deposit(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.Deposit(&_DutchExchange.TransactOpts, tokenAddress, amount)
}

// DepositAndSell is a paid mutator transaction binding the contract method 0x657a37ad.
//
// Solidity: function depositAndSell(address sellToken, address buyToken, uint256 amount) returns(uint256 newBal, uint256 auctionIndex, uint256 newSellerBal)
func (_DutchExchange *DutchExchangeTransactor) DepositAndSell(opts *bind.TransactOpts, sellToken common.Address, buyToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "depositAndSell", sellToken, buyToken, amount)
}

// DepositAndSell is a paid mutator transaction binding the contract method 0x657a37ad.
//
// Solidity: function depositAndSell(address sellToken, address buyToken, uint256 amount) returns(uint256 newBal, uint256 auctionIndex, uint256 newSellerBal)
func (_DutchExchange *DutchExchangeSession) DepositAndSell(sellToken common.Address, buyToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.DepositAndSell(&_DutchExchange.TransactOpts, sellToken, buyToken, amount)
}

// DepositAndSell is a paid mutator transaction binding the contract method 0x657a37ad.
//
// Solidity: function depositAndSell(address sellToken, address buyToken, uint256 amount) returns(uint256 newBal, uint256 auctionIndex, uint256 newSellerBal)
func (_DutchExchange *DutchExchangeTransactorSession) DepositAndSell(sellToken common.Address, buyToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.DepositAndSell(&_DutchExchange.TransactOpts, sellToken, buyToken, amount)
}

// InitiateEthUsdOracleUpdate is a paid mutator transaction binding the contract method 0x403fbf54.
//
// Solidity: function initiateEthUsdOracleUpdate(address _ethUSDOracle) returns()
func (_DutchExchange *DutchExchangeTransactor) InitiateEthUsdOracleUpdate(opts *bind.TransactOpts, _ethUSDOracle common.Address) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "initiateEthUsdOracleUpdate", _ethUSDOracle)
}

// InitiateEthUsdOracleUpdate is a paid mutator transaction binding the contract method 0x403fbf54.
//
// Solidity: function initiateEthUsdOracleUpdate(address _ethUSDOracle) returns()
func (_DutchExchange *DutchExchangeSession) InitiateEthUsdOracleUpdate(_ethUSDOracle common.Address) (*types.Transaction, error) {
	return _DutchExchange.Contract.InitiateEthUsdOracleUpdate(&_DutchExchange.TransactOpts, _ethUSDOracle)
}

// InitiateEthUsdOracleUpdate is a paid mutator transaction binding the contract method 0x403fbf54.
//
// Solidity: function initiateEthUsdOracleUpdate(address _ethUSDOracle) returns()
func (_DutchExchange *DutchExchangeTransactorSession) InitiateEthUsdOracleUpdate(_ethUSDOracle common.Address) (*types.Transaction, error) {
	return _DutchExchange.Contract.InitiateEthUsdOracleUpdate(&_DutchExchange.TransactOpts, _ethUSDOracle)
}

// PostBuyOrder is a paid mutator transaction binding the contract method 0x5e7f22c2.
//
// Solidity: function postBuyOrder(address sellToken, address buyToken, uint256 auctionIndex, uint256 amount) returns(uint256 newBuyerBal)
func (_DutchExchange *DutchExchangeTransactor) PostBuyOrder(opts *bind.TransactOpts, sellToken common.Address, buyToken common.Address, auctionIndex *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "postBuyOrder", sellToken, buyToken, auctionIndex, amount)
}

// PostBuyOrder is a paid mutator transaction binding the contract method 0x5e7f22c2.
//
// Solidity: function postBuyOrder(address sellToken, address buyToken, uint256 auctionIndex, uint256 amount) returns(uint256 newBuyerBal)
func (_DutchExchange *DutchExchangeSession) PostBuyOrder(sellToken common.Address, buyToken common.Address, auctionIndex *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.PostBuyOrder(&_DutchExchange.TransactOpts, sellToken, buyToken, auctionIndex, amount)
}

// PostBuyOrder is a paid mutator transaction binding the contract method 0x5e7f22c2.
//
// Solidity: function postBuyOrder(address sellToken, address buyToken, uint256 auctionIndex, uint256 amount) returns(uint256 newBuyerBal)
func (_DutchExchange *DutchExchangeTransactorSession) PostBuyOrder(sellToken common.Address, buyToken common.Address, auctionIndex *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.PostBuyOrder(&_DutchExchange.TransactOpts, sellToken, buyToken, auctionIndex, amount)
}

// PostSellOrder is a paid mutator transaction binding the contract method 0x59f96ae5.
//
// Solidity: function postSellOrder(address sellToken, address buyToken, uint256 auctionIndex, uint256 amount) returns(uint256, uint256)
func (_DutchExchange *DutchExchangeTransactor) PostSellOrder(opts *bind.TransactOpts, sellToken common.Address, buyToken common.Address, auctionIndex *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "postSellOrder", sellToken, buyToken, auctionIndex, amount)
}

// PostSellOrder is a paid mutator transaction binding the contract method 0x59f96ae5.
//
// Solidity: function postSellOrder(address sellToken, address buyToken, uint256 auctionIndex, uint256 amount) returns(uint256, uint256)
func (_DutchExchange *DutchExchangeSession) PostSellOrder(sellToken common.Address, buyToken common.Address, auctionIndex *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.PostSellOrder(&_DutchExchange.TransactOpts, sellToken, buyToken, auctionIndex, amount)
}

// PostSellOrder is a paid mutator transaction binding the contract method 0x59f96ae5.
//
// Solidity: function postSellOrder(address sellToken, address buyToken, uint256 auctionIndex, uint256 amount) returns(uint256, uint256)
func (_DutchExchange *DutchExchangeTransactorSession) PostSellOrder(sellToken common.Address, buyToken common.Address, auctionIndex *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.PostSellOrder(&_DutchExchange.TransactOpts, sellToken, buyToken, auctionIndex, amount)
}

// SetupDutchExchange is a paid mutator transaction binding the contract method 0xacb10351.
//
// Solidity: function setupDutchExchange(address _frtToken, address _owlToken, address _auctioneer, address _ethToken, address _ethUSDOracle, uint256 _thresholdNewTokenPair, uint256 _thresholdNewAuction) returns()
func (_DutchExchange *DutchExchangeTransactor) SetupDutchExchange(opts *bind.TransactOpts, _frtToken common.Address, _owlToken common.Address, _auctioneer common.Address, _ethToken common.Address, _ethUSDOracle common.Address, _thresholdNewTokenPair *big.Int, _thresholdNewAuction *big.Int) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "setupDutchExchange", _frtToken, _owlToken, _auctioneer, _ethToken, _ethUSDOracle, _thresholdNewTokenPair, _thresholdNewAuction)
}

// SetupDutchExchange is a paid mutator transaction binding the contract method 0xacb10351.
//
// Solidity: function setupDutchExchange(address _frtToken, address _owlToken, address _auctioneer, address _ethToken, address _ethUSDOracle, uint256 _thresholdNewTokenPair, uint256 _thresholdNewAuction) returns()
func (_DutchExchange *DutchExchangeSession) SetupDutchExchange(_frtToken common.Address, _owlToken common.Address, _auctioneer common.Address, _ethToken common.Address, _ethUSDOracle common.Address, _thresholdNewTokenPair *big.Int, _thresholdNewAuction *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.SetupDutchExchange(&_DutchExchange.TransactOpts, _frtToken, _owlToken, _auctioneer, _ethToken, _ethUSDOracle, _thresholdNewTokenPair, _thresholdNewAuction)
}

// SetupDutchExchange is a paid mutator transaction binding the contract method 0xacb10351.
//
// Solidity: function setupDutchExchange(address _frtToken, address _owlToken, address _auctioneer, address _ethToken, address _ethUSDOracle, uint256 _thresholdNewTokenPair, uint256 _thresholdNewAuction) returns()
func (_DutchExchange *DutchExchangeTransactorSession) SetupDutchExchange(_frtToken common.Address, _owlToken common.Address, _auctioneer common.Address, _ethToken common.Address, _ethUSDOracle common.Address, _thresholdNewTokenPair *big.Int, _thresholdNewAuction *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.SetupDutchExchange(&_DutchExchange.TransactOpts, _frtToken, _owlToken, _auctioneer, _ethToken, _ethUSDOracle, _thresholdNewTokenPair, _thresholdNewAuction)
}

// StartMasterCopyCountdown is a paid mutator transaction binding the contract method 0xf625ee28.
//
// Solidity: function startMasterCopyCountdown(address _masterCopy) returns()
func (_DutchExchange *DutchExchangeTransactor) StartMasterCopyCountdown(opts *bind.TransactOpts, _masterCopy common.Address) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "startMasterCopyCountdown", _masterCopy)
}

// StartMasterCopyCountdown is a paid mutator transaction binding the contract method 0xf625ee28.
//
// Solidity: function startMasterCopyCountdown(address _masterCopy) returns()
func (_DutchExchange *DutchExchangeSession) StartMasterCopyCountdown(_masterCopy common.Address) (*types.Transaction, error) {
	return _DutchExchange.Contract.StartMasterCopyCountdown(&_DutchExchange.TransactOpts, _masterCopy)
}

// StartMasterCopyCountdown is a paid mutator transaction binding the contract method 0xf625ee28.
//
// Solidity: function startMasterCopyCountdown(address _masterCopy) returns()
func (_DutchExchange *DutchExchangeTransactorSession) StartMasterCopyCountdown(_masterCopy common.Address) (*types.Transaction, error) {
	return _DutchExchange.Contract.StartMasterCopyCountdown(&_DutchExchange.TransactOpts, _masterCopy)
}

// UpdateApprovalOfToken is a paid mutator transaction binding the contract method 0x65b0d711.
//
// Solidity: function updateApprovalOfToken(address[] token, bool approved) returns()
func (_DutchExchange *DutchExchangeTransactor) UpdateApprovalOfToken(opts *bind.TransactOpts, token []common.Address, approved bool) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "updateApprovalOfToken", token, approved)
}

// UpdateApprovalOfToken is a paid mutator transaction binding the contract method 0x65b0d711.
//
// Solidity: function updateApprovalOfToken(address[] token, bool approved) returns()
func (_DutchExchange *DutchExchangeSession) UpdateApprovalOfToken(token []common.Address, approved bool) (*types.Transaction, error) {
	return _DutchExchange.Contract.UpdateApprovalOfToken(&_DutchExchange.TransactOpts, token, approved)
}

// UpdateApprovalOfToken is a paid mutator transaction binding the contract method 0x65b0d711.
//
// Solidity: function updateApprovalOfToken(address[] token, bool approved) returns()
func (_DutchExchange *DutchExchangeTransactorSession) UpdateApprovalOfToken(token []common.Address, approved bool) (*types.Transaction, error) {
	return _DutchExchange.Contract.UpdateApprovalOfToken(&_DutchExchange.TransactOpts, token, approved)
}

// UpdateAuctioneer is a paid mutator transaction binding the contract method 0x796a8076.
//
// Solidity: function updateAuctioneer(address _auctioneer) returns()
func (_DutchExchange *DutchExchangeTransactor) UpdateAuctioneer(opts *bind.TransactOpts, _auctioneer common.Address) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "updateAuctioneer", _auctioneer)
}

// UpdateAuctioneer is a paid mutator transaction binding the contract method 0x796a8076.
//
// Solidity: function updateAuctioneer(address _auctioneer) returns()
func (_DutchExchange *DutchExchangeSession) UpdateAuctioneer(_auctioneer common.Address) (*types.Transaction, error) {
	return _DutchExchange.Contract.UpdateAuctioneer(&_DutchExchange.TransactOpts, _auctioneer)
}

// UpdateAuctioneer is a paid mutator transaction binding the contract method 0x796a8076.
//
// Solidity: function updateAuctioneer(address _auctioneer) returns()
func (_DutchExchange *DutchExchangeTransactorSession) UpdateAuctioneer(_auctioneer common.Address) (*types.Transaction, error) {
	return _DutchExchange.Contract.UpdateAuctioneer(&_DutchExchange.TransactOpts, _auctioneer)
}

// UpdateEthUSDOracle is a paid mutator transaction binding the contract method 0xf4279d1f.
//
// Solidity: function updateEthUSDOracle() returns()
func (_DutchExchange *DutchExchangeTransactor) UpdateEthUSDOracle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "updateEthUSDOracle")
}

// UpdateEthUSDOracle is a paid mutator transaction binding the contract method 0xf4279d1f.
//
// Solidity: function updateEthUSDOracle() returns()
func (_DutchExchange *DutchExchangeSession) UpdateEthUSDOracle() (*types.Transaction, error) {
	return _DutchExchange.Contract.UpdateEthUSDOracle(&_DutchExchange.TransactOpts)
}

// UpdateEthUSDOracle is a paid mutator transaction binding the contract method 0xf4279d1f.
//
// Solidity: function updateEthUSDOracle() returns()
func (_DutchExchange *DutchExchangeTransactorSession) UpdateEthUSDOracle() (*types.Transaction, error) {
	return _DutchExchange.Contract.UpdateEthUSDOracle(&_DutchExchange.TransactOpts)
}

// UpdateMasterCopy is a paid mutator transaction binding the contract method 0x2cef4dac.
//
// Solidity: function updateMasterCopy() returns()
func (_DutchExchange *DutchExchangeTransactor) UpdateMasterCopy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "updateMasterCopy")
}

// UpdateMasterCopy is a paid mutator transaction binding the contract method 0x2cef4dac.
//
// Solidity: function updateMasterCopy() returns()
func (_DutchExchange *DutchExchangeSession) UpdateMasterCopy() (*types.Transaction, error) {
	return _DutchExchange.Contract.UpdateMasterCopy(&_DutchExchange.TransactOpts)
}

// UpdateMasterCopy is a paid mutator transaction binding the contract method 0x2cef4dac.
//
// Solidity: function updateMasterCopy() returns()
func (_DutchExchange *DutchExchangeTransactorSession) UpdateMasterCopy() (*types.Transaction, error) {
	return _DutchExchange.Contract.UpdateMasterCopy(&_DutchExchange.TransactOpts)
}

// UpdateThresholdNewAuction is a paid mutator transaction binding the contract method 0xc6af43f9.
//
// Solidity: function updateThresholdNewAuction(uint256 _thresholdNewAuction) returns()
func (_DutchExchange *DutchExchangeTransactor) UpdateThresholdNewAuction(opts *bind.TransactOpts, _thresholdNewAuction *big.Int) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "updateThresholdNewAuction", _thresholdNewAuction)
}

// UpdateThresholdNewAuction is a paid mutator transaction binding the contract method 0xc6af43f9.
//
// Solidity: function updateThresholdNewAuction(uint256 _thresholdNewAuction) returns()
func (_DutchExchange *DutchExchangeSession) UpdateThresholdNewAuction(_thresholdNewAuction *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.UpdateThresholdNewAuction(&_DutchExchange.TransactOpts, _thresholdNewAuction)
}

// UpdateThresholdNewAuction is a paid mutator transaction binding the contract method 0xc6af43f9.
//
// Solidity: function updateThresholdNewAuction(uint256 _thresholdNewAuction) returns()
func (_DutchExchange *DutchExchangeTransactorSession) UpdateThresholdNewAuction(_thresholdNewAuction *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.UpdateThresholdNewAuction(&_DutchExchange.TransactOpts, _thresholdNewAuction)
}

// UpdateThresholdNewTokenPair is a paid mutator transaction binding the contract method 0xe1c95bb9.
//
// Solidity: function updateThresholdNewTokenPair(uint256 _thresholdNewTokenPair) returns()
func (_DutchExchange *DutchExchangeTransactor) UpdateThresholdNewTokenPair(opts *bind.TransactOpts, _thresholdNewTokenPair *big.Int) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "updateThresholdNewTokenPair", _thresholdNewTokenPair)
}

// UpdateThresholdNewTokenPair is a paid mutator transaction binding the contract method 0xe1c95bb9.
//
// Solidity: function updateThresholdNewTokenPair(uint256 _thresholdNewTokenPair) returns()
func (_DutchExchange *DutchExchangeSession) UpdateThresholdNewTokenPair(_thresholdNewTokenPair *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.UpdateThresholdNewTokenPair(&_DutchExchange.TransactOpts, _thresholdNewTokenPair)
}

// UpdateThresholdNewTokenPair is a paid mutator transaction binding the contract method 0xe1c95bb9.
//
// Solidity: function updateThresholdNewTokenPair(uint256 _thresholdNewTokenPair) returns()
func (_DutchExchange *DutchExchangeTransactorSession) UpdateThresholdNewTokenPair(_thresholdNewTokenPair *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.UpdateThresholdNewTokenPair(&_DutchExchange.TransactOpts, _thresholdNewTokenPair)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address tokenAddress, uint256 amount) returns(uint256)
func (_DutchExchange *DutchExchangeTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.contract.Transact(opts, "withdraw", tokenAddress, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address tokenAddress, uint256 amount) returns(uint256)
func (_DutchExchange *DutchExchangeSession) Withdraw(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.Withdraw(&_DutchExchange.TransactOpts, tokenAddress, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address tokenAddress, uint256 amount) returns(uint256)
func (_DutchExchange *DutchExchangeTransactorSession) Withdraw(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DutchExchange.Contract.Withdraw(&_DutchExchange.TransactOpts, tokenAddress, amount)
}

// DutchExchangeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DutchExchange contract.
type DutchExchangeApprovalIterator struct {
	Event *DutchExchangeApproval // Event containing the contract specifics and raw log

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
func (it *DutchExchangeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchExchangeApproval)
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
		it.Event = new(DutchExchangeApproval)
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
func (it *DutchExchangeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchExchangeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchExchangeApproval represents a Approval event raised by the DutchExchange contract.
type DutchExchangeApproval struct {
	Token    common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xc091bf3abd3a42f670f8ad1a6ad5b849311210403e1d85d6ac31f43114d5ca6e.
//
// Solidity: event Approval(address indexed token, bool approved)
func (_DutchExchange *DutchExchangeFilterer) FilterApproval(opts *bind.FilterOpts, token []common.Address) (*DutchExchangeApprovalIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DutchExchange.contract.FilterLogs(opts, "Approval", tokenRule)
	if err != nil {
		return nil, err
	}
	return &DutchExchangeApprovalIterator{contract: _DutchExchange.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xc091bf3abd3a42f670f8ad1a6ad5b849311210403e1d85d6ac31f43114d5ca6e.
//
// Solidity: event Approval(address indexed token, bool approved)
func (_DutchExchange *DutchExchangeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DutchExchangeApproval, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DutchExchange.contract.WatchLogs(opts, "Approval", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchExchangeApproval)
				if err := _DutchExchange.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0xc091bf3abd3a42f670f8ad1a6ad5b849311210403e1d85d6ac31f43114d5ca6e.
//
// Solidity: event Approval(address indexed token, bool approved)
func (_DutchExchange *DutchExchangeFilterer) ParseApproval(log types.Log) (*DutchExchangeApproval, error) {
	event := new(DutchExchangeApproval)
	if err := _DutchExchange.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DutchExchangeAuctionClearedIterator is returned from FilterAuctionCleared and is used to iterate over the raw logs and unpacked data for AuctionCleared events raised by the DutchExchange contract.
type DutchExchangeAuctionClearedIterator struct {
	Event *DutchExchangeAuctionCleared // Event containing the contract specifics and raw log

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
func (it *DutchExchangeAuctionClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchExchangeAuctionCleared)
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
		it.Event = new(DutchExchangeAuctionCleared)
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
func (it *DutchExchangeAuctionClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchExchangeAuctionClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchExchangeAuctionCleared represents a AuctionCleared event raised by the DutchExchange contract.
type DutchExchangeAuctionCleared struct {
	SellToken    common.Address
	BuyToken     common.Address
	SellVolume   *big.Int
	BuyVolume    *big.Int
	AuctionIndex *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuctionCleared is a free log retrieval operation binding the contract event 0xb5806f8610464e96807c2b147620cc721c65309647f16cfccdf9fb7bd95152ac.
//
// Solidity: event AuctionCleared(address indexed sellToken, address indexed buyToken, uint256 sellVolume, uint256 buyVolume, uint256 indexed auctionIndex)
func (_DutchExchange *DutchExchangeFilterer) FilterAuctionCleared(opts *bind.FilterOpts, sellToken []common.Address, buyToken []common.Address, auctionIndex []*big.Int) (*DutchExchangeAuctionClearedIterator, error) {

	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}
	var buyTokenRule []interface{}
	for _, buyTokenItem := range buyToken {
		buyTokenRule = append(buyTokenRule, buyTokenItem)
	}

	var auctionIndexRule []interface{}
	for _, auctionIndexItem := range auctionIndex {
		auctionIndexRule = append(auctionIndexRule, auctionIndexItem)
	}

	logs, sub, err := _DutchExchange.contract.FilterLogs(opts, "AuctionCleared", sellTokenRule, buyTokenRule, auctionIndexRule)
	if err != nil {
		return nil, err
	}
	return &DutchExchangeAuctionClearedIterator{contract: _DutchExchange.contract, event: "AuctionCleared", logs: logs, sub: sub}, nil
}

// WatchAuctionCleared is a free log subscription operation binding the contract event 0xb5806f8610464e96807c2b147620cc721c65309647f16cfccdf9fb7bd95152ac.
//
// Solidity: event AuctionCleared(address indexed sellToken, address indexed buyToken, uint256 sellVolume, uint256 buyVolume, uint256 indexed auctionIndex)
func (_DutchExchange *DutchExchangeFilterer) WatchAuctionCleared(opts *bind.WatchOpts, sink chan<- *DutchExchangeAuctionCleared, sellToken []common.Address, buyToken []common.Address, auctionIndex []*big.Int) (event.Subscription, error) {

	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}
	var buyTokenRule []interface{}
	for _, buyTokenItem := range buyToken {
		buyTokenRule = append(buyTokenRule, buyTokenItem)
	}

	var auctionIndexRule []interface{}
	for _, auctionIndexItem := range auctionIndex {
		auctionIndexRule = append(auctionIndexRule, auctionIndexItem)
	}

	logs, sub, err := _DutchExchange.contract.WatchLogs(opts, "AuctionCleared", sellTokenRule, buyTokenRule, auctionIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchExchangeAuctionCleared)
				if err := _DutchExchange.contract.UnpackLog(event, "AuctionCleared", log); err != nil {
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

// ParseAuctionCleared is a log parse operation binding the contract event 0xb5806f8610464e96807c2b147620cc721c65309647f16cfccdf9fb7bd95152ac.
//
// Solidity: event AuctionCleared(address indexed sellToken, address indexed buyToken, uint256 sellVolume, uint256 buyVolume, uint256 indexed auctionIndex)
func (_DutchExchange *DutchExchangeFilterer) ParseAuctionCleared(log types.Log) (*DutchExchangeAuctionCleared, error) {
	event := new(DutchExchangeAuctionCleared)
	if err := _DutchExchange.contract.UnpackLog(event, "AuctionCleared", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DutchExchangeAuctionStartScheduledIterator is returned from FilterAuctionStartScheduled and is used to iterate over the raw logs and unpacked data for AuctionStartScheduled events raised by the DutchExchange contract.
type DutchExchangeAuctionStartScheduledIterator struct {
	Event *DutchExchangeAuctionStartScheduled // Event containing the contract specifics and raw log

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
func (it *DutchExchangeAuctionStartScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchExchangeAuctionStartScheduled)
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
		it.Event = new(DutchExchangeAuctionStartScheduled)
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
func (it *DutchExchangeAuctionStartScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchExchangeAuctionStartScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchExchangeAuctionStartScheduled represents a AuctionStartScheduled event raised by the DutchExchange contract.
type DutchExchangeAuctionStartScheduled struct {
	SellToken    common.Address
	BuyToken     common.Address
	AuctionIndex *big.Int
	AuctionStart *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuctionStartScheduled is a free log retrieval operation binding the contract event 0x20017e7b1ef8e7882103f55ff346ca3135c4afe13dff1da2f01b482aece766a5.
//
// Solidity: event AuctionStartScheduled(address indexed sellToken, address indexed buyToken, uint256 indexed auctionIndex, uint256 auctionStart)
func (_DutchExchange *DutchExchangeFilterer) FilterAuctionStartScheduled(opts *bind.FilterOpts, sellToken []common.Address, buyToken []common.Address, auctionIndex []*big.Int) (*DutchExchangeAuctionStartScheduledIterator, error) {

	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}
	var buyTokenRule []interface{}
	for _, buyTokenItem := range buyToken {
		buyTokenRule = append(buyTokenRule, buyTokenItem)
	}
	var auctionIndexRule []interface{}
	for _, auctionIndexItem := range auctionIndex {
		auctionIndexRule = append(auctionIndexRule, auctionIndexItem)
	}

	logs, sub, err := _DutchExchange.contract.FilterLogs(opts, "AuctionStartScheduled", sellTokenRule, buyTokenRule, auctionIndexRule)
	if err != nil {
		return nil, err
	}
	return &DutchExchangeAuctionStartScheduledIterator{contract: _DutchExchange.contract, event: "AuctionStartScheduled", logs: logs, sub: sub}, nil
}

// WatchAuctionStartScheduled is a free log subscription operation binding the contract event 0x20017e7b1ef8e7882103f55ff346ca3135c4afe13dff1da2f01b482aece766a5.
//
// Solidity: event AuctionStartScheduled(address indexed sellToken, address indexed buyToken, uint256 indexed auctionIndex, uint256 auctionStart)
func (_DutchExchange *DutchExchangeFilterer) WatchAuctionStartScheduled(opts *bind.WatchOpts, sink chan<- *DutchExchangeAuctionStartScheduled, sellToken []common.Address, buyToken []common.Address, auctionIndex []*big.Int) (event.Subscription, error) {

	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}
	var buyTokenRule []interface{}
	for _, buyTokenItem := range buyToken {
		buyTokenRule = append(buyTokenRule, buyTokenItem)
	}
	var auctionIndexRule []interface{}
	for _, auctionIndexItem := range auctionIndex {
		auctionIndexRule = append(auctionIndexRule, auctionIndexItem)
	}

	logs, sub, err := _DutchExchange.contract.WatchLogs(opts, "AuctionStartScheduled", sellTokenRule, buyTokenRule, auctionIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchExchangeAuctionStartScheduled)
				if err := _DutchExchange.contract.UnpackLog(event, "AuctionStartScheduled", log); err != nil {
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

// ParseAuctionStartScheduled is a log parse operation binding the contract event 0x20017e7b1ef8e7882103f55ff346ca3135c4afe13dff1da2f01b482aece766a5.
//
// Solidity: event AuctionStartScheduled(address indexed sellToken, address indexed buyToken, uint256 indexed auctionIndex, uint256 auctionStart)
func (_DutchExchange *DutchExchangeFilterer) ParseAuctionStartScheduled(log types.Log) (*DutchExchangeAuctionStartScheduled, error) {
	event := new(DutchExchangeAuctionStartScheduled)
	if err := _DutchExchange.contract.UnpackLog(event, "AuctionStartScheduled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DutchExchangeFeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the DutchExchange contract.
type DutchExchangeFeeIterator struct {
	Event *DutchExchangeFee // Event containing the contract specifics and raw log

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
func (it *DutchExchangeFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchExchangeFee)
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
		it.Event = new(DutchExchangeFee)
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
func (it *DutchExchangeFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchExchangeFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchExchangeFee represents a Fee event raised by the DutchExchange contract.
type DutchExchangeFee struct {
	PrimaryToken  common.Address
	SecondarToken common.Address
	User          common.Address
	AuctionIndex  *big.Int
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0x30c4d3fe752442ffa2415fd4e6398cb9e378bab963f042ce30ef4363b6ad93b6.
//
// Solidity: event Fee(address indexed primaryToken, address indexed secondarToken, address indexed user, uint256 auctionIndex, uint256 fee)
func (_DutchExchange *DutchExchangeFilterer) FilterFee(opts *bind.FilterOpts, primaryToken []common.Address, secondarToken []common.Address, user []common.Address) (*DutchExchangeFeeIterator, error) {

	var primaryTokenRule []interface{}
	for _, primaryTokenItem := range primaryToken {
		primaryTokenRule = append(primaryTokenRule, primaryTokenItem)
	}
	var secondarTokenRule []interface{}
	for _, secondarTokenItem := range secondarToken {
		secondarTokenRule = append(secondarTokenRule, secondarTokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DutchExchange.contract.FilterLogs(opts, "Fee", primaryTokenRule, secondarTokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DutchExchangeFeeIterator{contract: _DutchExchange.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0x30c4d3fe752442ffa2415fd4e6398cb9e378bab963f042ce30ef4363b6ad93b6.
//
// Solidity: event Fee(address indexed primaryToken, address indexed secondarToken, address indexed user, uint256 auctionIndex, uint256 fee)
func (_DutchExchange *DutchExchangeFilterer) WatchFee(opts *bind.WatchOpts, sink chan<- *DutchExchangeFee, primaryToken []common.Address, secondarToken []common.Address, user []common.Address) (event.Subscription, error) {

	var primaryTokenRule []interface{}
	for _, primaryTokenItem := range primaryToken {
		primaryTokenRule = append(primaryTokenRule, primaryTokenItem)
	}
	var secondarTokenRule []interface{}
	for _, secondarTokenItem := range secondarToken {
		secondarTokenRule = append(secondarTokenRule, secondarTokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DutchExchange.contract.WatchLogs(opts, "Fee", primaryTokenRule, secondarTokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchExchangeFee)
				if err := _DutchExchange.contract.UnpackLog(event, "Fee", log); err != nil {
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

// ParseFee is a log parse operation binding the contract event 0x30c4d3fe752442ffa2415fd4e6398cb9e378bab963f042ce30ef4363b6ad93b6.
//
// Solidity: event Fee(address indexed primaryToken, address indexed secondarToken, address indexed user, uint256 auctionIndex, uint256 fee)
func (_DutchExchange *DutchExchangeFilterer) ParseFee(log types.Log) (*DutchExchangeFee, error) {
	event := new(DutchExchangeFee)
	if err := _DutchExchange.contract.UnpackLog(event, "Fee", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DutchExchangeNewBuyOrderIterator is returned from FilterNewBuyOrder and is used to iterate over the raw logs and unpacked data for NewBuyOrder events raised by the DutchExchange contract.
type DutchExchangeNewBuyOrderIterator struct {
	Event *DutchExchangeNewBuyOrder // Event containing the contract specifics and raw log

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
func (it *DutchExchangeNewBuyOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchExchangeNewBuyOrder)
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
		it.Event = new(DutchExchangeNewBuyOrder)
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
func (it *DutchExchangeNewBuyOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchExchangeNewBuyOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchExchangeNewBuyOrder represents a NewBuyOrder event raised by the DutchExchange contract.
type DutchExchangeNewBuyOrder struct {
	SellToken    common.Address
	BuyToken     common.Address
	User         common.Address
	AuctionIndex *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewBuyOrder is a free log retrieval operation binding the contract event 0xf1751a362067564d5feb9ed26f1898bb14c17e1254e3724d454bc2ae80195c25.
//
// Solidity: event NewBuyOrder(address indexed sellToken, address indexed buyToken, address indexed user, uint256 auctionIndex, uint256 amount)
func (_DutchExchange *DutchExchangeFilterer) FilterNewBuyOrder(opts *bind.FilterOpts, sellToken []common.Address, buyToken []common.Address, user []common.Address) (*DutchExchangeNewBuyOrderIterator, error) {

	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}
	var buyTokenRule []interface{}
	for _, buyTokenItem := range buyToken {
		buyTokenRule = append(buyTokenRule, buyTokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DutchExchange.contract.FilterLogs(opts, "NewBuyOrder", sellTokenRule, buyTokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DutchExchangeNewBuyOrderIterator{contract: _DutchExchange.contract, event: "NewBuyOrder", logs: logs, sub: sub}, nil
}

// WatchNewBuyOrder is a free log subscription operation binding the contract event 0xf1751a362067564d5feb9ed26f1898bb14c17e1254e3724d454bc2ae80195c25.
//
// Solidity: event NewBuyOrder(address indexed sellToken, address indexed buyToken, address indexed user, uint256 auctionIndex, uint256 amount)
func (_DutchExchange *DutchExchangeFilterer) WatchNewBuyOrder(opts *bind.WatchOpts, sink chan<- *DutchExchangeNewBuyOrder, sellToken []common.Address, buyToken []common.Address, user []common.Address) (event.Subscription, error) {

	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}
	var buyTokenRule []interface{}
	for _, buyTokenItem := range buyToken {
		buyTokenRule = append(buyTokenRule, buyTokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DutchExchange.contract.WatchLogs(opts, "NewBuyOrder", sellTokenRule, buyTokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchExchangeNewBuyOrder)
				if err := _DutchExchange.contract.UnpackLog(event, "NewBuyOrder", log); err != nil {
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

// ParseNewBuyOrder is a log parse operation binding the contract event 0xf1751a362067564d5feb9ed26f1898bb14c17e1254e3724d454bc2ae80195c25.
//
// Solidity: event NewBuyOrder(address indexed sellToken, address indexed buyToken, address indexed user, uint256 auctionIndex, uint256 amount)
func (_DutchExchange *DutchExchangeFilterer) ParseNewBuyOrder(log types.Log) (*DutchExchangeNewBuyOrder, error) {
	event := new(DutchExchangeNewBuyOrder)
	if err := _DutchExchange.contract.UnpackLog(event, "NewBuyOrder", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DutchExchangeNewBuyerFundsClaimIterator is returned from FilterNewBuyerFundsClaim and is used to iterate over the raw logs and unpacked data for NewBuyerFundsClaim events raised by the DutchExchange contract.
type DutchExchangeNewBuyerFundsClaimIterator struct {
	Event *DutchExchangeNewBuyerFundsClaim // Event containing the contract specifics and raw log

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
func (it *DutchExchangeNewBuyerFundsClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchExchangeNewBuyerFundsClaim)
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
		it.Event = new(DutchExchangeNewBuyerFundsClaim)
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
func (it *DutchExchangeNewBuyerFundsClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchExchangeNewBuyerFundsClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchExchangeNewBuyerFundsClaim represents a NewBuyerFundsClaim event raised by the DutchExchange contract.
type DutchExchangeNewBuyerFundsClaim struct {
	SellToken    common.Address
	BuyToken     common.Address
	User         common.Address
	AuctionIndex *big.Int
	Amount       *big.Int
	FrtsIssued   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewBuyerFundsClaim is a free log retrieval operation binding the contract event 0x4d1c39fd1a9c74f88b9f90c7b439b7e5dc6f26b6ff280fd497fdec5c538aaf52.
//
// Solidity: event NewBuyerFundsClaim(address indexed sellToken, address indexed buyToken, address indexed user, uint256 auctionIndex, uint256 amount, uint256 frtsIssued)
func (_DutchExchange *DutchExchangeFilterer) FilterNewBuyerFundsClaim(opts *bind.FilterOpts, sellToken []common.Address, buyToken []common.Address, user []common.Address) (*DutchExchangeNewBuyerFundsClaimIterator, error) {

	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}
	var buyTokenRule []interface{}
	for _, buyTokenItem := range buyToken {
		buyTokenRule = append(buyTokenRule, buyTokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DutchExchange.contract.FilterLogs(opts, "NewBuyerFundsClaim", sellTokenRule, buyTokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DutchExchangeNewBuyerFundsClaimIterator{contract: _DutchExchange.contract, event: "NewBuyerFundsClaim", logs: logs, sub: sub}, nil
}

// WatchNewBuyerFundsClaim is a free log subscription operation binding the contract event 0x4d1c39fd1a9c74f88b9f90c7b439b7e5dc6f26b6ff280fd497fdec5c538aaf52.
//
// Solidity: event NewBuyerFundsClaim(address indexed sellToken, address indexed buyToken, address indexed user, uint256 auctionIndex, uint256 amount, uint256 frtsIssued)
func (_DutchExchange *DutchExchangeFilterer) WatchNewBuyerFundsClaim(opts *bind.WatchOpts, sink chan<- *DutchExchangeNewBuyerFundsClaim, sellToken []common.Address, buyToken []common.Address, user []common.Address) (event.Subscription, error) {

	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}
	var buyTokenRule []interface{}
	for _, buyTokenItem := range buyToken {
		buyTokenRule = append(buyTokenRule, buyTokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DutchExchange.contract.WatchLogs(opts, "NewBuyerFundsClaim", sellTokenRule, buyTokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchExchangeNewBuyerFundsClaim)
				if err := _DutchExchange.contract.UnpackLog(event, "NewBuyerFundsClaim", log); err != nil {
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

// ParseNewBuyerFundsClaim is a log parse operation binding the contract event 0x4d1c39fd1a9c74f88b9f90c7b439b7e5dc6f26b6ff280fd497fdec5c538aaf52.
//
// Solidity: event NewBuyerFundsClaim(address indexed sellToken, address indexed buyToken, address indexed user, uint256 auctionIndex, uint256 amount, uint256 frtsIssued)
func (_DutchExchange *DutchExchangeFilterer) ParseNewBuyerFundsClaim(log types.Log) (*DutchExchangeNewBuyerFundsClaim, error) {
	event := new(DutchExchangeNewBuyerFundsClaim)
	if err := _DutchExchange.contract.UnpackLog(event, "NewBuyerFundsClaim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DutchExchangeNewDepositIterator is returned from FilterNewDeposit and is used to iterate over the raw logs and unpacked data for NewDeposit events raised by the DutchExchange contract.
type DutchExchangeNewDepositIterator struct {
	Event *DutchExchangeNewDeposit // Event containing the contract specifics and raw log

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
func (it *DutchExchangeNewDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchExchangeNewDeposit)
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
		it.Event = new(DutchExchangeNewDeposit)
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
func (it *DutchExchangeNewDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchExchangeNewDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchExchangeNewDeposit represents a NewDeposit event raised by the DutchExchange contract.
type DutchExchangeNewDeposit struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewDeposit is a free log retrieval operation binding the contract event 0x2cb77763bc1e8490c1a904905c4d74b4269919aca114464f4bb4d911e60de364.
//
// Solidity: event NewDeposit(address indexed token, uint256 amount)
func (_DutchExchange *DutchExchangeFilterer) FilterNewDeposit(opts *bind.FilterOpts, token []common.Address) (*DutchExchangeNewDepositIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DutchExchange.contract.FilterLogs(opts, "NewDeposit", tokenRule)
	if err != nil {
		return nil, err
	}
	return &DutchExchangeNewDepositIterator{contract: _DutchExchange.contract, event: "NewDeposit", logs: logs, sub: sub}, nil
}

// WatchNewDeposit is a free log subscription operation binding the contract event 0x2cb77763bc1e8490c1a904905c4d74b4269919aca114464f4bb4d911e60de364.
//
// Solidity: event NewDeposit(address indexed token, uint256 amount)
func (_DutchExchange *DutchExchangeFilterer) WatchNewDeposit(opts *bind.WatchOpts, sink chan<- *DutchExchangeNewDeposit, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DutchExchange.contract.WatchLogs(opts, "NewDeposit", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchExchangeNewDeposit)
				if err := _DutchExchange.contract.UnpackLog(event, "NewDeposit", log); err != nil {
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

// ParseNewDeposit is a log parse operation binding the contract event 0x2cb77763bc1e8490c1a904905c4d74b4269919aca114464f4bb4d911e60de364.
//
// Solidity: event NewDeposit(address indexed token, uint256 amount)
func (_DutchExchange *DutchExchangeFilterer) ParseNewDeposit(log types.Log) (*DutchExchangeNewDeposit, error) {
	event := new(DutchExchangeNewDeposit)
	if err := _DutchExchange.contract.UnpackLog(event, "NewDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DutchExchangeNewMasterCopyProposalIterator is returned from FilterNewMasterCopyProposal and is used to iterate over the raw logs and unpacked data for NewMasterCopyProposal events raised by the DutchExchange contract.
type DutchExchangeNewMasterCopyProposalIterator struct {
	Event *DutchExchangeNewMasterCopyProposal // Event containing the contract specifics and raw log

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
func (it *DutchExchangeNewMasterCopyProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchExchangeNewMasterCopyProposal)
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
		it.Event = new(DutchExchangeNewMasterCopyProposal)
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
func (it *DutchExchangeNewMasterCopyProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchExchangeNewMasterCopyProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchExchangeNewMasterCopyProposal represents a NewMasterCopyProposal event raised by the DutchExchange contract.
type DutchExchangeNewMasterCopyProposal struct {
	NewMasterCopy common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewMasterCopyProposal is a free log retrieval operation binding the contract event 0x36dceb79f427eda3edba9ac3c1d4db7a6e4d0b8637d97320847d93fa8e2f7a04.
//
// Solidity: event NewMasterCopyProposal(address newMasterCopy)
func (_DutchExchange *DutchExchangeFilterer) FilterNewMasterCopyProposal(opts *bind.FilterOpts) (*DutchExchangeNewMasterCopyProposalIterator, error) {

	logs, sub, err := _DutchExchange.contract.FilterLogs(opts, "NewMasterCopyProposal")
	if err != nil {
		return nil, err
	}
	return &DutchExchangeNewMasterCopyProposalIterator{contract: _DutchExchange.contract, event: "NewMasterCopyProposal", logs: logs, sub: sub}, nil
}

// WatchNewMasterCopyProposal is a free log subscription operation binding the contract event 0x36dceb79f427eda3edba9ac3c1d4db7a6e4d0b8637d97320847d93fa8e2f7a04.
//
// Solidity: event NewMasterCopyProposal(address newMasterCopy)
func (_DutchExchange *DutchExchangeFilterer) WatchNewMasterCopyProposal(opts *bind.WatchOpts, sink chan<- *DutchExchangeNewMasterCopyProposal) (event.Subscription, error) {

	logs, sub, err := _DutchExchange.contract.WatchLogs(opts, "NewMasterCopyProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchExchangeNewMasterCopyProposal)
				if err := _DutchExchange.contract.UnpackLog(event, "NewMasterCopyProposal", log); err != nil {
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

// ParseNewMasterCopyProposal is a log parse operation binding the contract event 0x36dceb79f427eda3edba9ac3c1d4db7a6e4d0b8637d97320847d93fa8e2f7a04.
//
// Solidity: event NewMasterCopyProposal(address newMasterCopy)
func (_DutchExchange *DutchExchangeFilterer) ParseNewMasterCopyProposal(log types.Log) (*DutchExchangeNewMasterCopyProposal, error) {
	event := new(DutchExchangeNewMasterCopyProposal)
	if err := _DutchExchange.contract.UnpackLog(event, "NewMasterCopyProposal", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DutchExchangeNewOracleProposalIterator is returned from FilterNewOracleProposal and is used to iterate over the raw logs and unpacked data for NewOracleProposal events raised by the DutchExchange contract.
type DutchExchangeNewOracleProposalIterator struct {
	Event *DutchExchangeNewOracleProposal // Event containing the contract specifics and raw log

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
func (it *DutchExchangeNewOracleProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchExchangeNewOracleProposal)
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
		it.Event = new(DutchExchangeNewOracleProposal)
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
func (it *DutchExchangeNewOracleProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchExchangeNewOracleProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchExchangeNewOracleProposal represents a NewOracleProposal event raised by the DutchExchange contract.
type DutchExchangeNewOracleProposal struct {
	PriceOracleInterface common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewOracleProposal is a free log retrieval operation binding the contract event 0x3f174cfd713408ca6e4620d1efcc241b23fc39aa7d4694bd98610d3384dc001c.
//
// Solidity: event NewOracleProposal(address priceOracleInterface)
func (_DutchExchange *DutchExchangeFilterer) FilterNewOracleProposal(opts *bind.FilterOpts) (*DutchExchangeNewOracleProposalIterator, error) {

	logs, sub, err := _DutchExchange.contract.FilterLogs(opts, "NewOracleProposal")
	if err != nil {
		return nil, err
	}
	return &DutchExchangeNewOracleProposalIterator{contract: _DutchExchange.contract, event: "NewOracleProposal", logs: logs, sub: sub}, nil
}

// WatchNewOracleProposal is a free log subscription operation binding the contract event 0x3f174cfd713408ca6e4620d1efcc241b23fc39aa7d4694bd98610d3384dc001c.
//
// Solidity: event NewOracleProposal(address priceOracleInterface)
func (_DutchExchange *DutchExchangeFilterer) WatchNewOracleProposal(opts *bind.WatchOpts, sink chan<- *DutchExchangeNewOracleProposal) (event.Subscription, error) {

	logs, sub, err := _DutchExchange.contract.WatchLogs(opts, "NewOracleProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchExchangeNewOracleProposal)
				if err := _DutchExchange.contract.UnpackLog(event, "NewOracleProposal", log); err != nil {
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

// ParseNewOracleProposal is a log parse operation binding the contract event 0x3f174cfd713408ca6e4620d1efcc241b23fc39aa7d4694bd98610d3384dc001c.
//
// Solidity: event NewOracleProposal(address priceOracleInterface)
func (_DutchExchange *DutchExchangeFilterer) ParseNewOracleProposal(log types.Log) (*DutchExchangeNewOracleProposal, error) {
	event := new(DutchExchangeNewOracleProposal)
	if err := _DutchExchange.contract.UnpackLog(event, "NewOracleProposal", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DutchExchangeNewSellOrderIterator is returned from FilterNewSellOrder and is used to iterate over the raw logs and unpacked data for NewSellOrder events raised by the DutchExchange contract.
type DutchExchangeNewSellOrderIterator struct {
	Event *DutchExchangeNewSellOrder // Event containing the contract specifics and raw log

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
func (it *DutchExchangeNewSellOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchExchangeNewSellOrder)
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
		it.Event = new(DutchExchangeNewSellOrder)
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
func (it *DutchExchangeNewSellOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchExchangeNewSellOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchExchangeNewSellOrder represents a NewSellOrder event raised by the DutchExchange contract.
type DutchExchangeNewSellOrder struct {
	SellToken    common.Address
	BuyToken     common.Address
	User         common.Address
	AuctionIndex *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewSellOrder is a free log retrieval operation binding the contract event 0x3681d6f6ad159bac260c32828859f6df545bbf841c6e70787bcf0acbc390512a.
//
// Solidity: event NewSellOrder(address indexed sellToken, address indexed buyToken, address indexed user, uint256 auctionIndex, uint256 amount)
func (_DutchExchange *DutchExchangeFilterer) FilterNewSellOrder(opts *bind.FilterOpts, sellToken []common.Address, buyToken []common.Address, user []common.Address) (*DutchExchangeNewSellOrderIterator, error) {

	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}
	var buyTokenRule []interface{}
	for _, buyTokenItem := range buyToken {
		buyTokenRule = append(buyTokenRule, buyTokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DutchExchange.contract.FilterLogs(opts, "NewSellOrder", sellTokenRule, buyTokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DutchExchangeNewSellOrderIterator{contract: _DutchExchange.contract, event: "NewSellOrder", logs: logs, sub: sub}, nil
}

// WatchNewSellOrder is a free log subscription operation binding the contract event 0x3681d6f6ad159bac260c32828859f6df545bbf841c6e70787bcf0acbc390512a.
//
// Solidity: event NewSellOrder(address indexed sellToken, address indexed buyToken, address indexed user, uint256 auctionIndex, uint256 amount)
func (_DutchExchange *DutchExchangeFilterer) WatchNewSellOrder(opts *bind.WatchOpts, sink chan<- *DutchExchangeNewSellOrder, sellToken []common.Address, buyToken []common.Address, user []common.Address) (event.Subscription, error) {

	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}
	var buyTokenRule []interface{}
	for _, buyTokenItem := range buyToken {
		buyTokenRule = append(buyTokenRule, buyTokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DutchExchange.contract.WatchLogs(opts, "NewSellOrder", sellTokenRule, buyTokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchExchangeNewSellOrder)
				if err := _DutchExchange.contract.UnpackLog(event, "NewSellOrder", log); err != nil {
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

// ParseNewSellOrder is a log parse operation binding the contract event 0x3681d6f6ad159bac260c32828859f6df545bbf841c6e70787bcf0acbc390512a.
//
// Solidity: event NewSellOrder(address indexed sellToken, address indexed buyToken, address indexed user, uint256 auctionIndex, uint256 amount)
func (_DutchExchange *DutchExchangeFilterer) ParseNewSellOrder(log types.Log) (*DutchExchangeNewSellOrder, error) {
	event := new(DutchExchangeNewSellOrder)
	if err := _DutchExchange.contract.UnpackLog(event, "NewSellOrder", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DutchExchangeNewSellerFundsClaimIterator is returned from FilterNewSellerFundsClaim and is used to iterate over the raw logs and unpacked data for NewSellerFundsClaim events raised by the DutchExchange contract.
type DutchExchangeNewSellerFundsClaimIterator struct {
	Event *DutchExchangeNewSellerFundsClaim // Event containing the contract specifics and raw log

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
func (it *DutchExchangeNewSellerFundsClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchExchangeNewSellerFundsClaim)
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
		it.Event = new(DutchExchangeNewSellerFundsClaim)
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
func (it *DutchExchangeNewSellerFundsClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchExchangeNewSellerFundsClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchExchangeNewSellerFundsClaim represents a NewSellerFundsClaim event raised by the DutchExchange contract.
type DutchExchangeNewSellerFundsClaim struct {
	SellToken    common.Address
	BuyToken     common.Address
	User         common.Address
	AuctionIndex *big.Int
	Amount       *big.Int
	FrtsIssued   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewSellerFundsClaim is a free log retrieval operation binding the contract event 0xa3ac9b53d029621ef95693b5f9b1d0b0da75029fe8530389271be02715e24c13.
//
// Solidity: event NewSellerFundsClaim(address indexed sellToken, address indexed buyToken, address indexed user, uint256 auctionIndex, uint256 amount, uint256 frtsIssued)
func (_DutchExchange *DutchExchangeFilterer) FilterNewSellerFundsClaim(opts *bind.FilterOpts, sellToken []common.Address, buyToken []common.Address, user []common.Address) (*DutchExchangeNewSellerFundsClaimIterator, error) {

	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}
	var buyTokenRule []interface{}
	for _, buyTokenItem := range buyToken {
		buyTokenRule = append(buyTokenRule, buyTokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DutchExchange.contract.FilterLogs(opts, "NewSellerFundsClaim", sellTokenRule, buyTokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DutchExchangeNewSellerFundsClaimIterator{contract: _DutchExchange.contract, event: "NewSellerFundsClaim", logs: logs, sub: sub}, nil
}

// WatchNewSellerFundsClaim is a free log subscription operation binding the contract event 0xa3ac9b53d029621ef95693b5f9b1d0b0da75029fe8530389271be02715e24c13.
//
// Solidity: event NewSellerFundsClaim(address indexed sellToken, address indexed buyToken, address indexed user, uint256 auctionIndex, uint256 amount, uint256 frtsIssued)
func (_DutchExchange *DutchExchangeFilterer) WatchNewSellerFundsClaim(opts *bind.WatchOpts, sink chan<- *DutchExchangeNewSellerFundsClaim, sellToken []common.Address, buyToken []common.Address, user []common.Address) (event.Subscription, error) {

	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}
	var buyTokenRule []interface{}
	for _, buyTokenItem := range buyToken {
		buyTokenRule = append(buyTokenRule, buyTokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DutchExchange.contract.WatchLogs(opts, "NewSellerFundsClaim", sellTokenRule, buyTokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchExchangeNewSellerFundsClaim)
				if err := _DutchExchange.contract.UnpackLog(event, "NewSellerFundsClaim", log); err != nil {
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

// ParseNewSellerFundsClaim is a log parse operation binding the contract event 0xa3ac9b53d029621ef95693b5f9b1d0b0da75029fe8530389271be02715e24c13.
//
// Solidity: event NewSellerFundsClaim(address indexed sellToken, address indexed buyToken, address indexed user, uint256 auctionIndex, uint256 amount, uint256 frtsIssued)
func (_DutchExchange *DutchExchangeFilterer) ParseNewSellerFundsClaim(log types.Log) (*DutchExchangeNewSellerFundsClaim, error) {
	event := new(DutchExchangeNewSellerFundsClaim)
	if err := _DutchExchange.contract.UnpackLog(event, "NewSellerFundsClaim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DutchExchangeNewTokenPairIterator is returned from FilterNewTokenPair and is used to iterate over the raw logs and unpacked data for NewTokenPair events raised by the DutchExchange contract.
type DutchExchangeNewTokenPairIterator struct {
	Event *DutchExchangeNewTokenPair // Event containing the contract specifics and raw log

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
func (it *DutchExchangeNewTokenPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchExchangeNewTokenPair)
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
		it.Event = new(DutchExchangeNewTokenPair)
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
func (it *DutchExchangeNewTokenPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchExchangeNewTokenPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchExchangeNewTokenPair represents a NewTokenPair event raised by the DutchExchange contract.
type DutchExchangeNewTokenPair struct {
	SellToken common.Address
	BuyToken  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTokenPair is a free log retrieval operation binding the contract event 0x6f4b2adffa0c3e90e47fdcd9d2c36f48b57eb3271dce519997271073dac17be9.
//
// Solidity: event NewTokenPair(address indexed sellToken, address indexed buyToken)
func (_DutchExchange *DutchExchangeFilterer) FilterNewTokenPair(opts *bind.FilterOpts, sellToken []common.Address, buyToken []common.Address) (*DutchExchangeNewTokenPairIterator, error) {

	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}
	var buyTokenRule []interface{}
	for _, buyTokenItem := range buyToken {
		buyTokenRule = append(buyTokenRule, buyTokenItem)
	}

	logs, sub, err := _DutchExchange.contract.FilterLogs(opts, "NewTokenPair", sellTokenRule, buyTokenRule)
	if err != nil {
		return nil, err
	}
	return &DutchExchangeNewTokenPairIterator{contract: _DutchExchange.contract, event: "NewTokenPair", logs: logs, sub: sub}, nil
}

// WatchNewTokenPair is a free log subscription operation binding the contract event 0x6f4b2adffa0c3e90e47fdcd9d2c36f48b57eb3271dce519997271073dac17be9.
//
// Solidity: event NewTokenPair(address indexed sellToken, address indexed buyToken)
func (_DutchExchange *DutchExchangeFilterer) WatchNewTokenPair(opts *bind.WatchOpts, sink chan<- *DutchExchangeNewTokenPair, sellToken []common.Address, buyToken []common.Address) (event.Subscription, error) {

	var sellTokenRule []interface{}
	for _, sellTokenItem := range sellToken {
		sellTokenRule = append(sellTokenRule, sellTokenItem)
	}
	var buyTokenRule []interface{}
	for _, buyTokenItem := range buyToken {
		buyTokenRule = append(buyTokenRule, buyTokenItem)
	}

	logs, sub, err := _DutchExchange.contract.WatchLogs(opts, "NewTokenPair", sellTokenRule, buyTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchExchangeNewTokenPair)
				if err := _DutchExchange.contract.UnpackLog(event, "NewTokenPair", log); err != nil {
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

// ParseNewTokenPair is a log parse operation binding the contract event 0x6f4b2adffa0c3e90e47fdcd9d2c36f48b57eb3271dce519997271073dac17be9.
//
// Solidity: event NewTokenPair(address indexed sellToken, address indexed buyToken)
func (_DutchExchange *DutchExchangeFilterer) ParseNewTokenPair(log types.Log) (*DutchExchangeNewTokenPair, error) {
	event := new(DutchExchangeNewTokenPair)
	if err := _DutchExchange.contract.UnpackLog(event, "NewTokenPair", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DutchExchangeNewWithdrawalIterator is returned from FilterNewWithdrawal and is used to iterate over the raw logs and unpacked data for NewWithdrawal events raised by the DutchExchange contract.
type DutchExchangeNewWithdrawalIterator struct {
	Event *DutchExchangeNewWithdrawal // Event containing the contract specifics and raw log

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
func (it *DutchExchangeNewWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DutchExchangeNewWithdrawal)
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
		it.Event = new(DutchExchangeNewWithdrawal)
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
func (it *DutchExchangeNewWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DutchExchangeNewWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DutchExchangeNewWithdrawal represents a NewWithdrawal event raised by the DutchExchange contract.
type DutchExchangeNewWithdrawal struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewWithdrawal is a free log retrieval operation binding the contract event 0x6e2e05fb6a732995d6952d9158ca6b75f11cc6bf5a4af943aa1eb475a249440b.
//
// Solidity: event NewWithdrawal(address indexed token, uint256 amount)
func (_DutchExchange *DutchExchangeFilterer) FilterNewWithdrawal(opts *bind.FilterOpts, token []common.Address) (*DutchExchangeNewWithdrawalIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DutchExchange.contract.FilterLogs(opts, "NewWithdrawal", tokenRule)
	if err != nil {
		return nil, err
	}
	return &DutchExchangeNewWithdrawalIterator{contract: _DutchExchange.contract, event: "NewWithdrawal", logs: logs, sub: sub}, nil
}

// WatchNewWithdrawal is a free log subscription operation binding the contract event 0x6e2e05fb6a732995d6952d9158ca6b75f11cc6bf5a4af943aa1eb475a249440b.
//
// Solidity: event NewWithdrawal(address indexed token, uint256 amount)
func (_DutchExchange *DutchExchangeFilterer) WatchNewWithdrawal(opts *bind.WatchOpts, sink chan<- *DutchExchangeNewWithdrawal, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DutchExchange.contract.WatchLogs(opts, "NewWithdrawal", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DutchExchangeNewWithdrawal)
				if err := _DutchExchange.contract.UnpackLog(event, "NewWithdrawal", log); err != nil {
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

// ParseNewWithdrawal is a log parse operation binding the contract event 0x6e2e05fb6a732995d6952d9158ca6b75f11cc6bf5a4af943aa1eb475a249440b.
//
// Solidity: event NewWithdrawal(address indexed token, uint256 amount)
func (_DutchExchange *DutchExchangeFilterer) ParseNewWithdrawal(log types.Log) (*DutchExchangeNewWithdrawal, error) {
	event := new(DutchExchangeNewWithdrawal)
	if err := _DutchExchange.contract.UnpackLog(event, "NewWithdrawal", log); err != nil {
		return nil, err
	}
	return event, nil
}
