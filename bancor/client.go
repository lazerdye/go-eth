package bancor

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasoracle"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/wallet"
)

var (
	ContractRegistryContract = common.HexToAddress("0x52Ae12ABe5D8BD778BD5397F99cA900624CfADD4")
	EthAddress               = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
)

const (
	convertGasSpeed = gasoracle.Fastest
	convertGasLimit = uint64(700000)

	tradeGasSpeed = gasoracle.Fastest
	tradeGasLimit = uint64(300000)
)

type Client struct {
	*client.Client

	contractRegistryInstance  *ContractRegistry
	converterRegistryInstance *ConverterRegistry
	networkInstance           *Network
}

func getContractByName(opts *bind.CallOpts, registry *ContractRegistry, contractName string) (common.Address, error) {
	b := make([]byte, 32)
	copy(b[:], []byte(contractName))
	hash := common.BytesToHash(b)
	return registry.AddressOf(opts, hash)
}

func NewClient(ctx context.Context, client *client.Client) (*Client, error) {
	contractRegistryInstance, err := NewContractRegistry(ContractRegistryContract, client)
	if err != nil {
		return nil, err
	}
	converterRegistryAddress, err := getContractByName(client.DefaultCallOpts(ctx), contractRegistryInstance, "BancorConverterRegistry")
	if err != nil {
		return nil, err
	}
	converterRegistryInstance, err := NewConverterRegistry(converterRegistryAddress, client)
	if err != nil {
		return nil, err
	}
	networkAddress, err := getContractByName(client.DefaultCallOpts(ctx), contractRegistryInstance, "BancorNetwork")
	if err != nil {
		return nil, err
	}
	networkInstance, err := NewNetwork(networkAddress, client)
	if err != nil {
		return nil, err
	}
	return &Client{client, contractRegistryInstance, converterRegistryInstance, networkInstance}, nil
}

func (c *Client) BancorNetworkContract(ctx context.Context) (common.Address, error) {
	return getContractByName(c.DefaultCallOpts(ctx), c.contractRegistryInstance, "BancorNetwork")
}

func (c *Client) EstimateTradeFee(ctx context.Context) (decimal.Decimal, error) {
	// Get gas price for trade.
	gasPrice, err := c.GasPrice(ctx, convertGasSpeed)
	if err != nil {
		return decimal.Zero, err
	}
	log.Infof("Gas price: %s", gasPrice)
	// Multiply by gas limit.
	fee := gasPrice.Shift(9).Mul(decimal.NewFromInt(int64(convertGasLimit)))
	log.Infof("Trade fee: %s", fee)

	return fee.Shift(-18), nil
}

func (c *Client) ToWei(reg *token2.Registry, address common.Address, amount decimal.Decimal) (*big.Int, error) {
	if address == EthAddress {
		return client.EthToWei(amount), nil
	}
	_, tokClient, err := reg.ByAddress(address)
	if err != nil {
		return nil, err
	}
	return tokClient.ToWei(amount), nil
}

func (c *Client) FromWei(reg *token2.Registry, address common.Address, amount *big.Int) (decimal.Decimal, error) {
	if address == EthAddress {
		return client.EthFromWei(amount), nil
	}
	_, tokClient, err := reg.ByAddress(address)
	if err != nil {
		return decimal.Zero, err
	}
	return tokClient.FromWei(amount), nil
}

func (c *Client) GetConvertibleTokens(ctx context.Context) ([]common.Address, error) {
	return c.converterRegistryInstance.GetConvertibleTokens(c.DefaultCallOpts(ctx))
}

func (c *Client) GetConversionPath(ctx context.Context, sourceToken, targetToken common.Address) ([]common.Address, error) {
	return c.networkInstance.ConversionPath(c.DefaultCallOpts(ctx), sourceToken, targetToken)
}

func (c *Client) RateByPath(ctx context.Context, path []common.Address, amount *big.Int) (*big.Int, error) {
	return c.networkInstance.RateByPath(c.DefaultCallOpts(ctx), path, amount)
}

func (c *Client) ConvertByPath(ctx context.Context, account *wallet.Account, path []common.Address, amount *big.Int, minReturn *big.Int) (*types.Transaction, error) {
	var ethAmount *big.Int
	if path[0] == EthAddress {
		ethAmount = amount
	}
	gasPrice, err := c.GasPrice(ctx, convertGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := account.NewTransactor(ctx, ethAmount, gasPrice, convertGasLimit)
	if err != nil {
		return nil, err
	}
	return c.networkInstance.ConvertByPath(opts, path, amount, minReturn, common.HexToAddress(""), common.HexToAddress(""), big.NewInt(0))
}

func (c *Client) ParseNetworkConversion(log types.Log) (*NetworkConversion, error) {
	return c.networkInstance.ParseConversion(log)
}

func (c *Client) FilterConversion(ctx context.Context, fromBlock uint64, toBlock *uint64, smartToken, fromToken, toToken []common.Address) (*NetworkConversionIterator, error) {
	opts := bind.FilterOpts{Start: fromBlock, End: toBlock, Context: ctx}
	return c.networkInstance.FilterConversion(&opts, smartToken, fromToken, toToken)
}
