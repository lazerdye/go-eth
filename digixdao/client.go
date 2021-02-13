package digixdao

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasoracle"
	"github.com/lazerdye/go-eth/wallet"
)

var (
	DGDInterfaceContract = common.HexToAddress("0x23ea10cc1e6ebdb499d24e45369a35f43627062f")

	burnGasSpeed = gasoracle.Average
	burnGasLimit = uint64(300000)
)

type Client struct {
	*client.Client

	dgdInterface *DGDInterface
}

func NewClient(client *client.Client) (*Client, error) {
	dgdInterface, err := NewDGDInterface(DGDInterfaceContract, client)
	if err != nil {
		return nil, err
	}
	return &Client{client, dgdInterface}, nil
}

func (c *Client) Burn(ctx context.Context, account *wallet.Account) (*types.Transaction, error) {
	gasPrice, err := c.GasPrice(ctx, burnGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := account.NewTransactor(ctx, nil, gasPrice, burnGasLimit)
	if err != nil {
		return nil, err
	}
	return c.dgdInterface.Burn(opts)

}
