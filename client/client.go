package client

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/lazerdye/go-eth/dutchx"
	"github.com/lazerdye/go-eth/kyber"
	"github.com/lazerdye/go-eth/token"
)

type Client struct {
	c *ethclient.Client
}

const (
	ethDigits = 18
)

func Dial(url string) (*Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func (c *Client) Token(tokenName string) (*token.Client, error) {
	return token.NewClient(tokenName, c.c)
}

func (c *Client) Dutchx() (*dutchx.Client, error) {
	return dutchx.NewClient(c.c)
}

func (c *Client) Kyber() (*kyber.Client, error) {
	return kyber.NewClient(c.c)
}

func (c *Client) BalanceOf(ctx context.Context, address common.Address) (*big.Float, error) {
	balance, err := c.c.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, err
	}

	balanceFloat := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(ethDigits)))
	return balanceFloat, nil
}
