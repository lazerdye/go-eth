package sushi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/lazerdye/go-eth/client"
)

var (
	SushiV2FactoryContract = common.HexToAddress("0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac")
)

type Client struct {
	*client.Client

	factory *SushiV2Factory
}

func NewClient(client *client.Client) (*Client, error) {
	factoryInstance, err := NewSushiV2Factory(SushiV2FactoryContract, client)
	if err != nil {
		return nil, err
	}

	return &Client{client, factoryInstance}, nil
}

type PairClient struct {
	*client.Client

	Address common.Address
	pair    *Pair
}

func NewPairClient(client *client.Client, address common.Address) (*PairClient, error) {
	pairInstance, err := NewPair(address, client)
	if err != nil {
		return nil, err
	}
	return &PairClient{client, address, pairInstance}, nil
}

func (p *PairClient) Token0(ctx context.Context) (common.Address, error) {
	opts := p.DefaultCallOpts(ctx)
	return p.pair.Token0(opts)
}

func (p *PairClient) Token1(ctx context.Context) (common.Address, error) {
	opts := p.DefaultCallOpts(ctx)
	return p.pair.Token1(opts)
}

func (c *Client) GetPairs(ctx context.Context) ([]*PairClient, error) {
	opts := c.DefaultCallOpts(ctx)
	countBig, err := c.factory.AllPairsLength(opts)
	if err != nil {
		return nil, err
	}
	count := int(countBig.Int64())
	pairs := make([]*PairClient, count)
	for i := 0; i < count; i++ {
		pairAddr, err := c.factory.AllPairs(opts, big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}
		pairs[i], err = NewPairClient(c.Client, pairAddr)
		if err != nil {
			return nil, err
		}
	}
	return pairs, nil
}
