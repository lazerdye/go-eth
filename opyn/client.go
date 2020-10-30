package opyn

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/lazerdye/go-eth/client"
)

var (
	OptionsFactoryContract = common.HexToAddress("0xcc5d905b9c2c8c9329eb4e25dc086369d6c7777c")
)

type Client struct {
	*client.Client

	optionsFactory *OptionsFactory
}

func NewClient(client *client.Client) (*Client, error) {
	optionsFactory, err := NewOptionsFactory(OptionsFactoryContract, client)
	if err != nil {
		return nil, err
	}
	return &Client{client, optionsFactory}, nil
}

func (o *Client) GetNumberOfOptionsContracts(ctx context.Context) (*big.Int, error) {
	opts := o.DefaultCallOpts(ctx)
	return o.optionsFactory.GetNumberOfOptionsContracts(opts)
}
