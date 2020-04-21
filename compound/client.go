package compound

import (
	"context"
	"math/big"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasstation"
	"github.com/lazerdye/go-eth/wallet"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const ()

var (
	CompoundContractAddress = common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5")

	mintGasSpeed = gasstation.Average
	mintGasLimit = uint64(150000)
)

type Client struct {
	*client.Client

	instance *Compound
}

func NewClient(client *client.Client) (*Client, error) {
	instance, err := NewCompound(CompoundContractAddress, client)
	if err != nil {
		return nil, err
	}
	return &Client{client, instance}, nil
}

func (c *Client) Mint(ctx context.Context, account *wallet.Account, amount *big.Float) (*types.Transaction, error) {
	gasPrice, _, err := c.GasPrice(ctx, mintGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := account.NewTransactor(ctx, client.EthToWei(amount), gasPrice, mintGasLimit)
	if err != nil {
		return nil, err
	}

	return c.instance.Mint(opts)
}
