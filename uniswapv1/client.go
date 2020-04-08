package uniswapv1

import (
	"context"

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

func NewClient(client *client.Client) (*Client, error) {
	factoryInstance, err := NewFactory(UniswapV1FactoryContract, client)
	if err != nil {
		return nil, err
	}
	return &Client{client, factoryInstance}, nil
}

func (c *Client) GetExchange(ctx context.Context, tok *token.Client) (common.Address, error) {
	opts := &bind.CallOpts{Context: ctx}
	address, err := c.factory.GetExchange(opts, tok.ContractAddress())
	if err != nil {
		return common.Address{}, err
	}
	return address, nil
}
