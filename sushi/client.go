package sushi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasoracle"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/wallet"
)

var (
	SushiV2FactoryContract  = common.HexToAddress("0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac")
	SushiV2Router02Contract = common.HexToAddress("0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F")

	defaultSwapGasSpeed = gasoracle.Fastest
	defaultSwapGasLimit = uint64(700000)
)

type Client struct {
	*client.Client

	factory *SushiV2Factory
	router  *SushiV2Router02

	swapGasSpeed gasoracle.GasSpeed
	swapGasLimit uint64
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

	return &Client{client, factoryInstance, routerInstance,
		defaultSwapGasSpeed, defaultSwapGasLimit}, nil
}

func (c *Client) SetSwapGasSpeed(gasSpeed gasoracle.GasSpeed) {
	c.swapGasSpeed = gasSpeed
}

func (c *Client) SetSwapGasLimit(gasLimit uint64) {
	c.swapGasLimit = gasLimit
}

func (c *Client) EstimateTradeFee(ctx context.Context) (decimal.Decimal, error) {
	// Get gas price for trade.
	gasPrice, err := c.GasPrice(ctx, c.swapGasSpeed)
	if err != nil {
		return decimal.Zero, err
	}
	log.Infof("Gas price: %s", gasPrice)
	// Multiply by gas limit.
	fee := gasPrice.Shift(9).Mul(decimal.NewFromInt(int64(c.swapGasLimit)))
	log.Infof("SwapGasLimit: %d", c.swapGasLimit)
	log.Infof("Fee: %s", fee)

	return fee.Shift(-18), nil
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

func (c *Client) GetPair(ctx context.Context, token0 *token2.Client, token1 *token2.Client) (*PairClient, error) {
	opts := c.DefaultCallOpts(ctx)
	pair, err := c.factory.GetPair(opts, token0.Address, token1.Address)
	if err != nil {
		return nil, err
	}
	return NewPairClient(c.Client, pair)
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

func (c *Client) SwapETHForExactTokens(ctx context.Context, account *wallet.Account, maxAmountIn decimal.Decimal, amountOut decimal.Decimal, tokenPath []*token2.Client, to common.Address, deadline int64) (*types.Transaction, error) {
	// TODO: Verify tokenPath[len(tokenPath) - 1] == weth
	gasPrice, err := c.GasPrice(ctx, c.swapGasSpeed)
	if err != nil {
		return nil, err
	}
	maxAmountInBig := tokenPath[len(tokenPath)-1].ToWei(maxAmountIn)
	opts, err := account.NewTransactor(ctx, maxAmountInBig, gasPrice, c.swapGasLimit)
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

func (c *Client) SwapExactTokensForETH(ctx context.Context, account *wallet.Account, amountIn decimal.Decimal, amountOutMin decimal.Decimal, tokenPath []*token2.Client, to common.Address, deadline int64) (*types.Transaction, error) {
	// TODO: Verify tokenPath[len(tokenPath)-1] == weth
	gasPrice, err := c.GasPrice(ctx, c.swapGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := account.NewTransactor(ctx, nil, gasPrice, c.swapGasLimit)
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
	if false {
		return nil, errors.New("Not Implemented")
	}
	t, err := c.router.SwapExactTokensForETH(opts, amountInBig, amountOutMinBig, path, to, big.NewInt(deadline))
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (c *Client) SwapExactETHForTokens(ctx context.Context, account *wallet.Account, amountIn decimal.Decimal, amountOutMin decimal.Decimal, tokenPath []*token2.Client, to common.Address, deadline int64) (*types.Transaction, error) {
	// TODO: Verify tokenPath[len(tokenPath)-1] == weth
	gasPrice, err := c.GasPrice(ctx, c.swapGasSpeed)
	if err != nil {
		return nil, err
	}
	path := make([]common.Address, len(tokenPath))
	for i, token := range tokenPath {
		path[i] = token.Address
	}
	amountInBig := tokenPath[0].ToWei(amountIn)
	opts, err := account.NewTransactor(ctx, amountInBig, gasPrice, c.swapGasLimit)
	if err != nil {
		return nil, err
	}
	amountOutMinBig := tokenPath[len(tokenPath)-1].ToWei(amountOutMin)
	if false {
		return nil, errors.New("Not Implemented")
	}
	t, err := c.router.SwapExactETHForTokens(opts, amountOutMinBig, path, to, big.NewInt(deadline))
	if err != nil {
		return nil, err
	}
	return t, nil
}
