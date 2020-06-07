package uniswapv1

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lazerdye/go-eth/gasstation"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/wallet"
)

var (
	UniswapV1FactoryContract = common.HexToAddress("0xc0a47dFe034B400B47bDaD5FecDa2621de6c4d95")

	buyGasSpeed   = gasstation.Fastest
	sellGasSpeed  = gasstation.Fastest
	tradeGasLimit = uint64(500000)

	dnil = decimal.Decimal{}
)

type Client struct {
	*client.Client

	factory *Factory
}

type ExchangeClient struct {
	*token2.Client

	address  common.Address
	exchange *Exchange
}

func NewClient(client *client.Client) (*Client, error) {
	factoryInstance, err := NewFactory(UniswapV1FactoryContract, client)
	if err != nil {
		return nil, err
	}
	return &Client{client, factoryInstance}, nil
}

func (c *Client) GetExchange(ctx context.Context, tok *token2.Client) (*ExchangeClient, error) {
	opts := &bind.CallOpts{Context: ctx}
	address, err := c.factory.GetExchange(opts, tok.Address)
	if err != nil {
		return nil, err
	}
	exchange, err := NewExchange(address, c.Client)
	if err != nil {
		return nil, err
	}
	return &ExchangeClient{tok, address, exchange}, nil
}

func (e *ExchangeClient) ContractAddress() common.Address {
	return e.address
}

func (e *ExchangeClient) GetEthToTokenInputPrice(ctx context.Context, ethSold decimal.Decimal) (decimal.Decimal, error) {
	opts := &bind.CallOpts{Context: ctx}
	tokenBought, err := e.exchange.GetEthToTokenInputPrice(opts, client.EthToWei(ethSold))
	if err != nil {
		return dnil, err
	}

	return e.FromWei(tokenBought), nil
}

func (e *ExchangeClient) GetEthToTokenOutputPrice(ctx context.Context, tokensBought decimal.Decimal) (decimal.Decimal, error) {
	opts := &bind.CallOpts{Context: ctx}
	ethSold, err := e.exchange.GetEthToTokenOutputPrice(opts, e.ToWei(tokensBought))
	if err != nil {
		return dnil, err
	}

	return client.EthFromWei(ethSold), nil
}

func (e *ExchangeClient) GetTokenToEthInputPrice(ctx context.Context, tokensSold decimal.Decimal) (decimal.Decimal, error) {
	opts := &bind.CallOpts{Context: ctx}
	ethBought, err := e.exchange.GetTokenToEthInputPrice(opts, e.ToWei(tokensSold))
	if err != nil {
		return dnil, err
	}

	return client.EthFromWei(ethBought), nil
}

func (e *ExchangeClient) GetTokenToEthOutputPrice(ctx context.Context, ethBought decimal.Decimal) (decimal.Decimal, error) {
	opts := &bind.CallOpts{Context: ctx}
	tokenSold, err := e.exchange.GetTokenToEthOutputPrice(opts, client.EthToWei(ethBought))
	if err != nil {
		return dnil, err
	}

	return e.FromWei(tokenSold), nil
}

func (e *ExchangeClient) EthToTokenSwapOutput(ctx context.Context, account *wallet.Account, maxEthSold decimal.Decimal, tokensBought decimal.Decimal, deadline int) (*types.Transaction, error) {
	gasPrice, _, err := e.GasPrice(ctx, buyGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := account.NewTransactor(ctx, client.EthToWei(maxEthSold), gasPrice, tradeGasLimit)
	if err != nil {
		return nil, err
	}

	tx, err := e.exchange.EthToTokenSwapOutput(opts, e.ToWei(tokensBought), big.NewInt(int64(deadline)))
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (e *ExchangeClient) TokenToEthSwapInput(ctx context.Context, account *wallet.Account, tokensSold decimal.Decimal, minEth decimal.Decimal, deadline int) (*types.Transaction, error) {
	gasPrice, _, err := e.GasPrice(ctx, sellGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := account.NewTransactor(ctx, nil, gasPrice, tradeGasLimit)
	if err != nil {
		return nil, err
	}

	tokensCapped, err := e.ToWeiCapped(ctx, tokensSold)
	if err != nil {
		return nil, err
	}
	tx, err := e.exchange.TokenToEthSwapInput(opts, tokensCapped, client.EthToWei(minEth), big.NewInt(int64(deadline)))
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (e *ExchangeClient) ParseTokenPurchase(log types.Log) (*ExchangeTokenPurchase, error) {
	return e.exchange.ParseTokenPurchase(log)
}

func (e *ExchangeClient) ParseEthPurchase(log types.Log) (*ExchangeEthPurchase, error) {
	return e.exchange.ParseEthPurchase(log)
}
