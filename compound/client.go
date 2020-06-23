package compound

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	//"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasstation"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/wallet"
)

const ()

var (
	CethContractAddress = common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5")
	CbatContractAddress = common.HexToAddress("0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e")

	mintGasSpeed = gasstation.Fast
	mintGasLimit = uint64(300000)
)

type Client interface {
	Mint(ctx context.Context, account *wallet.Account, amount decimal.Decimal) (*types.Transaction, error)
}

type cethClient struct {
	*client.Client

	ceth *Ceth
}

type cerc20Client struct {
	*client.Client

	tok    *token2.Client
	cerc20 *CErc20
}

func NewEthClient(client *client.Client) (Client, error) {
	ceth, err := NewCeth(CethContractAddress, client)
	if err != nil {
		return nil, err
	}
	return &cethClient{client, ceth}, nil
}

func NewBatClient(client *client.Client, tok *token2.Client) (Client, error) {
	cerc20, err := NewCErc20(CbatContractAddress, client)
	if err != nil {
		return nil, err
	}
	return &cerc20Client{client, tok, cerc20}, nil
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

func (c *cerc20Client) Mint(ctx context.Context, account *wallet.Account, amount decimal.Decimal) (*types.Transaction, error) {
	gasPrice, _, err := c.GasPrice(ctx, mintGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := account.NewTransactor(ctx, nil, gasPrice, mintGasLimit)
	if err != nil {
		return nil, err
	}
	amountWei, err := c.tok.ToWeiCapped(ctx, amount, account)
	if err != nil {
		return nil, err
	}

	return c.cerc20.Mint(opts, amountWei)
}
