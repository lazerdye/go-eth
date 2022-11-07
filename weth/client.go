package weth

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasoracle"
	"github.com/lazerdye/go-eth/wallet"
)

var (
	WethContractAddress = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")

	defaultDepositGasLimit = uint64(1000000)
	depositGasSpeed        = gasoracle.Fastest
)

type WethClient struct {
	*client.Client

	depositGasLimit uint64
	weth            *Weth
}

func NewWethClient(c *client.Client) (*WethClient, error) {
	weth, err := NewWeth(WethContractAddress, c)
	if err != nil {
		return nil, err
	}
	return &WethClient{
		Client:          c,
		depositGasLimit: defaultDepositGasLimit,
		weth:            weth,
	}, nil
}

func (c *WethClient) Deposit(ctx context.Context, account *wallet.Account, amount decimal.Decimal) (*types.Transaction, error) {
	gasFeeCap, gasTipCap, err := c.GasPrice(ctx, depositGasSpeed)
	if err != nil {
		return nil, err
	}
	wei := client.EthToWei(amount)
	opts, err := account.NewTransactor(ctx, wei, gasFeeCap, gasTipCap, c.depositGasLimit)
	if err != nil {
		return nil, err
	}
	return c.weth.Deposit(opts)
}
