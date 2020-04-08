package uniswapv1

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/token"
)

var (
	UniswapV1FactoryContract = common.HexToAddress("0xc0a47dFe034B400B47bDaD5FecDa2621de6c4d95")
)

type Client struct {
	*client.Client

	factory *Factory
}

type ExchangeClient struct {
	*token.Client

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

func (c *Client) GetExchange(ctx context.Context, tok *token.Client) (*ExchangeClient, error) {
	opts := &bind.CallOpts{Context: ctx}
	address, err := c.factory.GetExchange(opts, tok.ContractAddress())
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

func (e *ExchangeClient) GetEthToTokenInputPrice(ctx context.Context, ethSold *big.Float) (*big.Float, error) {
	// Convert ethSold to int.
	intValue, _ := new(big.Float).Mul(ethSold, big.NewFloat(math.Pow10(18))).Int(nil)

	opts := &bind.CallOpts{Context: ctx}
	tokenBought, err := e.exchange.GetEthToTokenInputPrice(opts, intValue)
	if err != nil {
		return nil, err
	}

	return e.FromGwei(tokenBought), nil
}

func (e *ExchangeClient) GetEthToTokenOutputPrice(ctx context.Context, tokensBought *big.Float) (*big.Float, error) {
	// Convert tokensBought to int.

	opts := &bind.CallOpts{Context: ctx}
	ethSold, err := e.exchange.GetEthToTokenOutputPrice(opts, e.ToGwei(tokensBought))
	if err != nil {
		return nil, err
	}

	value := new(big.Float).Quo(new(big.Float).SetInt(ethSold), big.NewFloat(math.Pow10(18)))

	return value, nil
}

func (e *ExchangeClient) GetTokenToEthInputPrice(ctx context.Context, tokensSold *big.Float) (*big.Float, error) {
	opts := &bind.CallOpts{Context: ctx}
	ethBought, err := e.exchange.GetTokenToEthInputPrice(opts, e.ToGwei(tokensSold))
	if err != nil {
		return nil, err
	}

	value := new(big.Float).Quo(new(big.Float).SetInt(ethBought), big.NewFloat(math.Pow10(18)))

	return value, nil
}

func (e *ExchangeClient) GetTokenToEthOutputPrice(ctx context.Context, ethBought *big.Float) (*big.Float, error) {
	intValue, _ := new(big.Float).Mul(ethBought, big.NewFloat(math.Pow10(18))).Int(nil)

	opts := &bind.CallOpts{Context: ctx}
	tokenSold, err := e.exchange.GetTokenToEthOutputPrice(opts, intValue)
	if err != nil {
		return nil, err
	}

	return e.FromGwei(tokenSold), nil
}
