package zeroex

import (
	"context"
	"math"
	"math/big"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/wallet"
	"github.com/lazerdye/go-eth/zeroex/ether_token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	EtherTokenAddress = common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
)

type Client struct {
	*client.Client

	etherInstance *ether_token.EtherToken
}

func NewClient(tokenClient *client.Client) (*Client, error) {
	etherInstance, err := ether_token.NewEtherToken(EtherTokenAddress, tokenClient)
	if err != nil {
		return nil, err
	}
	return &Client{Client: tokenClient, etherInstance: etherInstance}, nil
}

func (c *Client) EtherTokenDeposit(ctx context.Context, account *wallet.Account, amount *big.Float) (*types.Transaction, error) {
	value, _ := new(big.Float).Mul(amount, big.NewFloat(math.Pow10(18))).Int(nil)
	opts, err := account.NewTransactor(ctx, value, client.GasLimit, client.GasPrice)
	if err != nil {
		return nil, err
	}

	return c.etherInstance.Deposit(opts)
}

func (c *Client) EtherTokenBalanceOf(ctx context.Context, account *wallet.Account) (*big.Float, error) {
	opts := &bind.CallOpts{Context: ctx}
	amountInt, err := c.etherInstance.BalanceOf(opts, account.Address())
	if err != nil {
		return nil, err
	}
	value := new(big.Float).Quo(new(big.Float).SetInt(amountInt), big.NewFloat(math.Pow10(18)))
	return value, nil
}
