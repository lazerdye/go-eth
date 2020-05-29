package token2

import (
	"context"
	"io/ioutil"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v2"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/wallet"
)

var (
	approveGasLimit  = uint64(200000)
	transferGasLimit = uint64(200000)
)

type TokenInfo struct {
	Address  string `json:"address"`
	Decimals uint8  `json:"decimals"`
}

type Tokens struct {
	Tokens map[string]TokenInfo `json:"tokens"`
}

type Client struct {
	*client.Client

	instance *Erc20
	Address  common.Address
	Decimals uint8
}

func NewClient(etherClient *client.Client, address common.Address, decimals uint8) (*Client, error) {
	instance, err := NewErc20(address, etherClient)
	if err != nil {
		return nil, err
	}
	return &Client{etherClient, instance, address, decimals}, nil
}

func ByAddress(ctx context.Context, client *client.Client, address common.Address) (*Client, error) {
	opts := bind.CallOpts{Context: ctx}
	instance, err := NewErc20(address, client)
	if err != nil {
		return nil, err
	}
	decimals, err := instance.Decimals(&opts)
	if err != nil {
		return nil, err
	}
	return &Client{client, instance, address, decimals}, nil
}

func (c *Client) Symbol(ctx context.Context) (string, error) {
	opts := &bind.CallOpts{Context: ctx}
	return c.instance.Symbol(opts)
}

func (c *Client) FromWei(i *big.Int) decimal.Decimal {
	return decimal.NewFromBigInt(i, -int32(c.Decimals))
}

func (c *Client) ToWei(f decimal.Decimal) *big.Int {
	return f.Shift(int32(c.Decimals)).BigInt()
}

func (c *Client) Allowance(ctx context.Context, address, contract common.Address) (decimal.Decimal, error) {
	opts := bind.CallOpts{Context: ctx}
	iAmount, err := c.instance.Allowance(&opts, address, contract)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return c.FromWei(iAmount), nil
}

func (c *Client) Approve(ctx context.Context, from *wallet.Account, contract common.Address, value decimal.Decimal) (*types.Transaction, error) {
	gasPrice, _, err := c.GasPrice(ctx, client.TransferGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := from.NewTransactor(ctx, nil, gasPrice, approveGasLimit)
	if err != nil {
		return nil, err
	}

	iAmount := c.ToWei(value)
	return c.instance.Approve(opts, contract, iAmount)
}

func (c *Client) BalanceOf(ctx context.Context, account common.Address) (decimal.Decimal, error) {
	opts := &bind.CallOpts{Context: ctx}
	balance, err := c.instance.BalanceOf(opts, account)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return c.FromWei(balance), nil
}

func (c *Client) Transfer(ctx context.Context, sourceAccount *wallet.Account, destAccount common.Address, amount decimal.Decimal) (*types.Transaction, error) {
	gasPrice, _, err := c.GasPrice(ctx, client.TransferGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := sourceAccount.NewTransactor(ctx, nil, gasPrice, transferGasLimit)
	if err != nil {
		return nil, err
	}
	return c.instance.Transfer(opts, destAccount, c.ToWei(amount))
}

type Registry struct {
	tokens map[string]*Client
}

func RegistryFromFile(etherClient *client.Client, filename string) (*Registry, error) {
	var tokens Tokens
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(yamlFile, &tokens); err != nil {
		return nil, err
	}
	registry := &Registry{}
	registry.tokens = make(map[string]*Client, len(tokens.Tokens))
	for name, tokenInfo := range tokens.Tokens {
		client, err := NewClient(etherClient, common.HexToAddress(tokenInfo.Address), tokenInfo.Decimals)
		if err != nil {
			return nil, err
		}
		if err := registry.Register(name, client); err != nil {
			return nil, err
		}
	}
	return registry, nil
}

func (r *Registry) Register(name string, client *Client) error {
	if _, ok := r.tokens[name]; ok != false {
		return errors.Errorf("Token %s already defined", name)
	}
	r.tokens[name] = client
	return nil
}

func (r *Registry) AllNames() []string {
	names := make([]string, 0, len(r.tokens))
	for name, _ := range r.tokens {
		names = append(names, name)
	}
	return names
}

func (r *Registry) ByName(name string) (*Client, error) {
	client, ok := r.tokens[name]
	if !ok {
		return nil, errors.Errorf("Token %s not found", name)
	}
	return client, nil
}

func (r *Registry) ByAddress(address common.Address) (string, *Client, error) {
	for name, client := range r.tokens {
		if client.Address == address {
			return name, client, nil
		}
	}
	return "", nil, errors.Errorf("Client with address %s not found", address.String())
}

func (r *Registry) Client(etherClient *client.Client, tokenName string) (*Client, error) {
	token, ok := r.tokens[tokenName]
	if !ok {
		return nil, errors.Errorf("Unknown token: %s", tokenName)
	}
	return token, nil
}
