package uniswapv2

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/token"
)

var (
	UniswapV2FactoryContract  = common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	UniswapV2Router01Contract = common.HexToAddress("0xf164fc0ec4e93095b804a4795bbe1e041497b92a")
)

type Client struct {
	*client.Client

	factory *Factory
	router  *Router01
}

func NewClient(client *client.Client) (*Client, error) {
	factoryInstance, err := NewFactory(UniswapV2FactoryContract, client)
	if err != nil {
		return nil, err
	}
	routerInstance, err := NewRouter01(UniswapV2Router01Contract, client)
	if err != nil {
		return nil, err
	}
	return &Client{client, factoryInstance, routerInstance}, nil
}

type PairClient struct {
	Address common.Address

	pair *Pair
}

func NewPairClient(client *client.Client, address common.Address) (*PairClient, error) {
	pairInstance, err := NewPair(address, client)
	if err != nil {
		return nil, err
	}
	return &PairClient{address, pairInstance}, nil
}

func (p *PairClient) Token0(ctx context.Context) (common.Address, error) {
	opts := &bind.CallOpts{Context: ctx}
	return p.pair.Token0(opts)
}

func (p *PairClient) Token1(ctx context.Context) (common.Address, error) {
	opts := &bind.CallOpts{Context: ctx}
	return p.pair.Token1(opts)
}

func (c *Client) GetPairs(ctx context.Context) ([]*PairClient, error) {
	opts := &bind.CallOpts{Context: ctx}
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

func (c *Client) GetAmountOut(ctx context.Context, amountIn *big.Float, token0 *token.Client, token1 *token.Client) (*big.Float, error) {
	opts := &bind.CallOpts{Context: ctx}
	path := []common.Address{token0.ContractAddress(), token1.ContractAddress()}
	amounts, err := c.router.GetAmountsOut(opts, token0.ToWei(amountIn), path)
	if err != nil {
		return nil, err
	}
	if len(amounts) != 2 {
		return nil, errors.Errorf("Expected result of length 1: %+v", amounts)
	}
	return token1.FromWei(amounts[1]), nil
}

func (c *Client) GetAmountIn(ctx context.Context, amountOut *big.Float, token0 *token.Client, token1 *token.Client) (*big.Float, error) {
	opts := &bind.CallOpts{Context: ctx}
	path := []common.Address{token0.ContractAddress(), token1.ContractAddress()}
	amounts, err := c.router.GetAmountsIn(opts, token1.ToWei(amountOut), path)
	if err != nil {
		return nil, err
	}
	if len(amounts) != 2 {
		return nil, errors.Errorf("Expected result of length 1: %+v", amounts)
	}
	return token0.FromWei(amounts[0]), nil
}
