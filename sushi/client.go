package sushi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/token2"
)

var (
	SushiV2FactoryContract  = common.HexToAddress("0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac")
	SushiV2Router02Contract = common.HexToAddress("0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F")
)

type Client struct {
	*client.Client

	factory *SushiV2Factory
	router  *SushiV2Router02
}

func NewClient(client *client.Client) (*Client, error) {
	factoryInstance, err := NewSushiV2Factory(SushiV2FactoryContract, client)
	if err != nil {
		return nil, err
	}
	routerInstance, err := NewSushiV2Router02(SushiV2Router02Contract, client)
	if err != nil {
		return nil, err
	}

	return &Client{client, factoryInstance, routerInstance}, nil
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

func (c *Client) GetAmountsIn(ctx context.Context, amountOut decimal.Decimal, tokenPath []*token2.Client) (decimal.Decimal, error) {
	opts := c.DefaultCallOpts(ctx)
	path := make([]common.Address, len(tokenPath))
	for i, token := range tokenPath {
		path[i] = token.Address
	}
	firstToken := tokenPath[0]
	lastToken := tokenPath[len(tokenPath)-1]
	amounts, err := c.router.GetAmountsIn(opts, lastToken.ToWei(amountOut), path)
	if err != nil {
		return decimal.Zero, err
	}
	if len(amounts) != 2 {
		return decimal.Zero, errors.Errorf("Expected reults of length 2: %+v", amounts)
	}
	return firstToken.FromWei(amounts[0]), nil
}

func (c *Client) GetAmountsOut(ctx context.Context, amountIn decimal.Decimal, tokenPath []*token2.Client) (decimal.Decimal, error) {
	opts := c.DefaultCallOpts(ctx)
	path := make([]common.Address, len(tokenPath))
	for i, token := range tokenPath {
		path[i] = token.Address
	}
	firstToken := tokenPath[0]
	lastToken := tokenPath[len(tokenPath)-1]
	amounts, err := c.router.GetAmountsOut(opts, firstToken.ToWei(amountIn), path)
	if err != nil {
		return decimal.Zero, err
	}
	if len(amounts) != 2 {
		return decimal.Zero, errors.Errorf("Expected reults of length 2: %+v", amounts)
	}
	return lastToken.FromWei(amounts[1]), nil
}
