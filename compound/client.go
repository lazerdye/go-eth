package compound

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasstation"
	"github.com/lazerdye/go-eth/wallet"
)

const ()

var (
	CethContractAddress = common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5")

	mintGasSpeed = gasstation.Average
	mintGasLimit = uint64(150000)
)

type Client interface {
	Mint(ctx context.Context, account *wallet.Account, amount decimal.Decimal) (*types.Transaction, error)
}

type cethClient struct {
	*client.Client

	ceth *Ceth
}

func NewClient(client *client.Client, currency string) (Client, error) { 
	if currency != "eth" {
		return nil, errors.Errorf("Unsupported currency: %s", currency)
	}
	ceth, err := NewCeth(CethContractAddress, client)
	if err != nil {
		return nil, err
	}
	return &cethClient{client, ceth}, nil
}

func (c *cethClient) Mint(ctx context.Context, account *wallet.Account, amount decimal.Decimal) (*types.Transaction, error) {
	gasPrice, _, err := c.GasPrice(ctx, mintGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := account.NewTransactor(ctx, client.EthToWei(amount), gasPrice, mintGasLimit)
	if err != nil {
		return nil, err
	}

	return c.ceth.Mint(opts)
}
