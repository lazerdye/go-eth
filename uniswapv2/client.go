package uniswapv2

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasstation"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/wallet"
)

var (
	UniswapV2FactoryContract  = common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	UniswapV2Router01Contract = common.HexToAddress("0xf164fc0ec4e93095b804a4795bbe1e041497b92a")

	buyGasSpeed  = gasstation.Fastest
	sellGasSpeed = gasstation.Fast

	swapGasLimit = uint64(500000)
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

func (p *PairClient) ParseSwap(log types.Log) (*PairSwap, error) {
	return p.pair.ParseSwap(log)
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

func (c *Client) PairClient(address common.Address) (*PairClient, error) {
	return NewPairClient(c.Client, address)
}

func (c *Client) GetAmountOut(ctx context.Context, amountIn *big.Float, token0 *token2.Client, token1 *token2.Client) (*big.Float, error) {
	opts := &bind.CallOpts{Context: ctx}
	path := []common.Address{token0.Address, token1.Address}
	amounts, err := c.router.GetAmountsOut(opts, token0.ToWei(amountIn), path)
	if err != nil {
		return nil, err
	}
	if len(amounts) != 2 {
		return nil, errors.Errorf("Expected result of length 1: %+v", amounts)
	}
	return token1.FromWei(amounts[1]), nil
}

func (c *Client) GetAmountIn(ctx context.Context, amountOut *big.Float, token0 *token2.Client, token1 *token2.Client) (*big.Float, error) {
	opts := &bind.CallOpts{Context: ctx}
	path := []common.Address{token0.Address, token1.Address}
	amounts, err := c.router.GetAmountsIn(opts, token1.ToWei(amountOut), path)
	if err != nil {
		return nil, err
	}
	if len(amounts) != 2 {
		return nil, errors.Errorf("Expected result of length 1: %+v", amounts)
	}
	return token0.FromWei(amounts[0]), nil
}

func (c *Client) SwapExactTokensForETH(ctx context.Context, account *wallet.Account, amountIn *big.Float, amountOutMin *big.Float, tokenPath []*token2.Client, to common.Address, deadline int64) (*types.Transaction, error) {
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
	amountInBig := tokenPath[0].ToWei(amountIn)
	amountOutMinBig := tokenPath[len(tokenPath)-1].ToWei(amountOutMin)
	t, err := c.router.SwapExactTokensForETH(opts, amountInBig, amountOutMinBig, path, to, big.NewInt(deadline))
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (c *Client) SwapETHForExactTokens(ctx context.Context, account *wallet.Account, maxAmountIn *big.Float, amountOut *big.Float, tokenPath []*token2.Client, to common.Address, deadline int64) (*types.Transaction, error) {
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