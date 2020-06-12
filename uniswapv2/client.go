package uniswapv2

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasstation"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/wallet"
)

var (
	UniswapV2FactoryContract  = common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	UniswapV2Router02Contract = common.HexToAddress("0x7a250d5630b4cf539739df2c5dacb4c659f2488d")

	buyGasSpeed  = gasstation.Fastest
	sellGasSpeed = gasstation.Fastest

	swapGasLimit = uint64(700000)
)

type Client struct {
	*client.Client

	factory *Factory
	router  *Router02
}

func NewClient(client *client.Client) (*Client, error) {
	factoryInstance, err := NewFactory(UniswapV2FactoryContract, client)
	if err != nil {
		return nil, err
	}
	routerInstance, err := NewRouter02(UniswapV2Router02Contract, client)
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

func (p *PairClient) GetReserves(ctx context.Context) (string, error) {
	opts := p.DefaultCallOpts(ctx)
	info, err := p.pair.GetReserves(opts)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%+v", info), err
}

func (p *PairClient) ParseSwap(log types.Log) (*PairSwap, error) {
	return p.pair.ParseSwap(log)
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

func (c *Client) GetPair(ctx context.Context, token0 *token2.Client, token1 *token2.Client) (*PairClient, error) {
	opts := c.DefaultCallOpts(ctx)
	address, err := c.factory.GetPair(opts, token0.Address, token1.Address)
	if err != nil {
		return nil, err
	}
	return c.PairClient(address)
}

func (c *Client) PairClient(address common.Address) (*PairClient, error) {
	return NewPairClient(c.Client, address)
}

func (c *Client) GetAmountOut(ctx context.Context, amountIn decimal.Decimal, token0 *token2.Client, token1 *token2.Client) (decimal.Decimal, error) {
	opts := c.DefaultCallOpts(ctx)
	path := []common.Address{token0.Address, token1.Address}
	amounts, err := c.router.GetAmountsOut(opts, token0.ToWei(amountIn), path)
	if err != nil {
		return decimal.Decimal{}, err
	}
	if len(amounts) != 2 {
		return decimal.Decimal{}, errors.Errorf("Expected result of length 1: %+v", amounts)
	}
	return token1.FromWei(amounts[1]), nil
}

func (c *Client) GetAmountIn(ctx context.Context, amountOut decimal.Decimal, token0 *token2.Client, token1 *token2.Client) (decimal.Decimal, error) {
	opts := c.DefaultCallOpts(ctx)
	path := []common.Address{token0.Address, token1.Address}
	amounts, err := c.router.GetAmountsIn(opts, token1.ToWei(amountOut), path)
	if err != nil {
		return decimal.Decimal{}, err
	}
	if len(amounts) != 2 {
		return decimal.Decimal{}, errors.Errorf("Expected result of length 1: %+v", amounts)
	}
	return token0.FromWei(amounts[0]), nil
}

func (c *Client) SwapExactTokensForETH(ctx context.Context, account *wallet.Account, amountIn decimal.Decimal, amountOutMin decimal.Decimal, tokenPath []*token2.Client, to common.Address, deadline int64) (*types.Transaction, error) {
	// TODO: Verify tokenPath[len(tokenPath)-1] == weth
	gasPrice, _, err := c.GasPrice(ctx, sellGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := account.NewTransactor(ctx, nil, gasPrice, swapGasLimit)
	if err != nil {
		return nil, err
	}
	path := make([]common.Address, len(tokenPath))
	for i, token := range tokenPath {
		path[i] = token.Address
	}
	amountInBig, err := tokenPath[0].ToWeiCapped(ctx, amountIn, account)
	if err != nil {
		return nil, err
	}
	amountOutMinBig := tokenPath[len(tokenPath)-1].ToWei(amountOutMin)
	fmt.Printf("amountInBig: %d - amountOutMinBig: %d - opts: %+v\n", amountInBig, amountOutMinBig, opts)
	if false {
		return nil, errors.New("Not Implemented")
	}
	t, err := c.router.SwapExactTokensForETH(opts, amountInBig, amountOutMinBig, path, to, big.NewInt(deadline))
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (c *Client) SwapETHForExactTokens(ctx context.Context, account *wallet.Account, maxAmountIn decimal.Decimal, amountOut decimal.Decimal, tokenPath []*token2.Client, to common.Address, deadline int64) (*types.Transaction, error) {
	// TODO: Verify tokenPath[len(tokenPath) - 1] == weth
	gasPrice, _, err := c.GasPrice(ctx, buyGasSpeed)
	if err != nil {
		return nil, err
	}
	maxAmountInBig := tokenPath[len(tokenPath)-1].ToWei(maxAmountIn)
	opts, err := account.NewTransactor(ctx, maxAmountInBig, gasPrice, swapGasLimit)
	if err != nil {
		return nil, err
	}
	path := make([]common.Address, len(tokenPath))
	for i, token := range tokenPath {
		path[len(tokenPath)-1-i] = token.Address
	}
	amountOutBig := tokenPath[0].ToWei(amountOut)
	t, err := c.router.SwapETHForExactTokens(opts, amountOutBig, path, to, big.NewInt(deadline))
	if err != nil {
		return nil, err
	}
	return t, nil
}
