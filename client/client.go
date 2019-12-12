package client

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"gitlab.com/lazerdye/go-eth/token"
)

type Client struct {
	c *ethclient.Client
}

func Dial(url string) (*Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func (c *Client) Token(name string) (*token.Client, error) {
	return token.NewClient(name, c.c)
}

func (c *Client) EthBalanceOf(ctx context.Context, address string) (float64, error) {
	account := common.HexToAddress(address)
	balance, err := c.c.BalanceAt(ctx, account, nil)
	if err != nil {
		return float64(0), err
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	retValue, _ := ethValue.Float64()
	return retValue, nil
}
