package augur

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasstation"
	"github.com/lazerdye/go-eth/wallet"
)

var (
	Repv2Contract = common.HexToAddress("0x221657776846890989a759BA2973e427DfF5C9bB")

	tradeGasLimit = uint64(900000)
	tradeGasSpeed = gasstation.Fast
)

type Client struct {
	*client.Client

	repv2Instance *Repv2
}

func NewClient(client *client.Client) (*Client, error) {
	repv2Instance, err := NewRepv2(Repv2Contract, client)
	if err != nil {
		return nil, err
	}
	return &Client{client, repv2Instance}, nil

}

func (c *Client) MigrateFromLegacyReputationToken(ctx context.Context, account *wallet.Account) (*types.Transaction, error) {
	gasPrice, _, err := c.GasPrice(ctx, tradeGasSpeed)
	if err != nil {
		return nil, err
	}
	transactOpts, err := account.NewTransactor(ctx, nil, gasPrice, tradeGasLimit)
	if err != nil {
		return nil, err
	}
	return c.repv2Instance.MigrateFromLegacyReputationToken(transactOpts)
}
