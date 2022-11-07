package maker

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"math/big"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasoracle"
	"github.com/lazerdye/go-eth/wallet"
)

var (
	// TODO: Take from https://chainlog.makerdao.com/api/mainnet/active.json
	CDPManagerContract = common.HexToAddress("0x5ef30b9986345249bc32d8928B7ee64DE9435E39")

	defaultOpenGasLimit = uint64(1000000)
	openGasSpeed        = gasoracle.Fastest
	defaultFrobGasLimit = uint64(1000000)
	frobGasSpeed        = gasoracle.Fastest
)

type CDPManagerClient struct {
	*client.Client

	openGasLimit uint64
	frobGasLimit uint64
	cdpManager   *CDPManager
}

func NewCDPManagerClient(c *client.Client) (*CDPManagerClient, error) {
	cdpManager, err := NewCDPManager(CDPManagerContract, c)
	if err != nil {
		return nil, err
	}
	return &CDPManagerClient{
		Client:       c,
		openGasLimit: defaultOpenGasLimit,
		frobGasLimit: defaultFrobGasLimit,
		cdpManager:   cdpManager,
	}, nil
}

func (c *CDPManagerClient) Open(ctx context.Context, ilk string, account *wallet.Account) (*types.Transaction, error) {
	gasFeeCap, gasTipCap, err := c.GasPrice(ctx, openGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := account.NewTransactor(ctx, nil, gasFeeCap, gasTipCap, c.openGasLimit)
	if err != nil {
		return nil, err
	}

	ilkBytes := AsciiToByte(ilk)
	return c.cdpManager.Open(opts, ilkBytes, account.Address())
}

func (c *CDPManagerClient) Last(ctx context.Context, account *wallet.Account) (*big.Int, error) {
	opts := c.DefaultCallOpts(ctx)
	return c.cdpManager.Last(opts, account.Address())
}

func (c *CDPManagerClient) Urns(ctx context.Context, cdpId *big.Int) (common.Address, error) {
	opts := c.DefaultCallOpts(ctx)
	return c.cdpManager.Urns(opts, cdpId)
}

func (c *CDPManagerClient) Frob(ctx context.Context, account *wallet.Account, cdpId *big.Int, dink decimal.Decimal, dart decimal.Decimal) (*types.Transaction, error) {
	gasFeeCap, gasTipCap, err := c.GasPrice(ctx, frobGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := account.NewTransactor(ctx, nil, gasFeeCap, gasTipCap, c.openGasLimit)
	if err != nil {
		return nil, err
	}
	// TODO: Other token decimals?
	dinkWei := client.EthToWei(dink)
	dartWei := client.EthToWei(dart)
	return c.cdpManager.Frob(opts, cdpId, dinkWei, dartWei)
}
