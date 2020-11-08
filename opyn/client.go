package opyn

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasstation"
	"github.com/lazerdye/go-eth/wallet"
)

var (
	OptionsFactoryContract = common.HexToAddress("0xcc5d905b9c2c8c9329eb4e25dc086369d6c7777c")

	EthContract = common.HexToAddress("0x0")
)

const (
	opynGasSpeed      = gasstation.Fast
	openVaultGasLimit = 0
)

type Client struct {
	*client.Client

	optionsFactory *OptionsFactory
}

func NewClient(client *client.Client) (*Client, error) {
	optionsFactory, err := NewOptionsFactory(OptionsFactoryContract, client)
	if err != nil {
		return nil, err
	}
	return &Client{client, optionsFactory}, nil
}

func (o *Client) GetNumberOfOptionsContracts(ctx context.Context) (*big.Int, error) {
	opts := o.DefaultCallOpts(ctx)
	return o.optionsFactory.GetNumberOfOptionsContracts(opts)
}

func (o *Client) OptionsContract(ctx context.Context, index *big.Int) (common.Address, error) {
	opts := o.DefaultCallOpts(ctx)
	return o.optionsFactory.OptionsContracts(opts, index)
}

type OTokenClient struct {
	*client.Client

	otoken *OToken
}

func (o *Client) GetOToken(ctx context.Context, contract common.Address) (*OTokenClient, error) {
	otoken, err := NewOToken(contract, o)
	if err != nil {
		return nil, err
	}
	return &OTokenClient{o.Client, otoken}, nil
}

func (t *OTokenClient) Name(ctx context.Context) (string, error) {
	return t.otoken.Name(t.DefaultCallOpts(ctx))
}

func (t *OTokenClient) Decimals(ctx context.Context) (uint8, error) {
	return t.otoken.Decimals(t.DefaultCallOpts(ctx))
}

func (t *OTokenClient) HasExpired(ctx context.Context) (bool, error) {
	return t.otoken.HasExpired(t.DefaultCallOpts(ctx))
}

func (t *OTokenClient) Underlying(ctx context.Context) (common.Address, error) {
	return t.otoken.Underlying(t.DefaultCallOpts(ctx))
}

func (t *OTokenClient) UnderlyingExp(ctx context.Context) (int32, error) {
	return t.otoken.UnderlyingExp(t.DefaultCallOpts(ctx))
}

func (t *OTokenClient) Collateral(ctx context.Context) (common.Address, error) {
	return t.otoken.Collateral(t.DefaultCallOpts(ctx))
}

func (t *OTokenClient) CollateralExp(ctx context.Context) (int32, error) {
	return t.otoken.CollateralExp(t.DefaultCallOpts(ctx))
}

func (t *OTokenClient) StrikePrice(ctx context.Context) (decimal.Decimal, error) {
	strikePrice, err := t.otoken.StrikePrice(t.DefaultCallOpts(ctx))
	if err != nil {
		return decimal.Zero, nil
	}
	return decimal.NewFromBigInt(strikePrice.Value, strikePrice.Exponent), nil
}

func (t *OTokenClient) OpenVault(ctx context.Context, account *wallet.Account) (*types.Transaction, error) {
	gasPrice, _, err := t.GasPrice(ctx, opynGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := account.NewTransactor(ctx, nil, gasPrice, openVaultGasLimit)
	if err != nil {
		return nil, err
	}

	tx, err := t.otoken.OpenVault(opts)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
