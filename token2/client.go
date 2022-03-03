package token2

import (
	"context"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/util"
	"github.com/lazerdye/go-eth/wallet"
)

var (
	approveGasLimit  = uint64(300000)
	transferGasLimit = uint64(300000)

	minDiff, _ = decimal.NewFromString("1e-4")
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

func (c *Client) ToWeiCapped(ctx context.Context, f decimal.Decimal, account *wallet.Account) (*big.Int, error) {
	balance, err := c.BalanceOf(ctx, account.Address())
	if err != nil {
		return nil, err
	}
	log.Infof("Balance of %s in account %s is: %s", c.Address.Hex(), account.Address().Hex(), balance)
	diff := f.Sub(balance).Abs()
	if f.Cmp(balance) > 0 && diff.Cmp(minDiff) > 0 {
		return nil, errors.Errorf("Cannot cap, too much of a difference: %s-%s=%s", f, balance, diff)
	} else if !diff.IsZero() && diff.Cmp(minDiff) <= 0 {
		log.Infof("Capping withdraw to %s", balance)
		return c.ToWei(balance), nil
	}
	return c.ToWei(f), nil
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
	gasFeeCap, gasTipCap, err := c.GasPrice(ctx, client.TransferGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := from.NewTransactor(ctx, nil, gasFeeCap, gasTipCap, approveGasLimit)
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
	gasFeeCap, gasTipCap, err := c.GasPrice(ctx, client.TransferGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := sourceAccount.NewTransactor(ctx, nil, gasFeeCap, gasTipCap, transferGasLimit)
	if err != nil {
		return nil, err
	}
	amountCapped, err := c.ToWeiCapped(ctx, amount, sourceAccount)
	if err != nil {
		return nil, err
	}
	return c.instance.Transfer(opts, destAccount, amountCapped)
}

type Registry struct {
	http   util.HttpClient
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

type TokenList struct {
	Name      string `json:"name"`
	Timestamp string `json:"timestamp"`
	Version   struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	} `json:"version"`
	Tags     interface{} `json:"tags"` // TODO: Define
	LogoURI  string      `json:"logoURI"`
	Keywords []string    `json:"keywords"`
	Tokens   []struct {
		Name     string `json:"name"`
		Address  string `json:"address"`
		Symbol   string `json:"symbol"`
		Decimals uint8  `json:"decimals"`
		ChainID  int    `json:"chainID"`
		LogoURI  string `json:"logoURI"`
	} `json:"tokens"`
}

func (r *Registry) ImportTokenList(ctx context.Context, etherClient *client.Client, url string, chainID int) error {
	var tokenList TokenList
	err := r.http.Get(ctx, url, nil, &tokenList)
	if err != nil {
		return err
	}
	for _, token := range tokenList.Tokens {
		if token.ChainID != chainID {
			continue
		}
		name := strings.ToLower(token.Symbol)
		addressHex := common.HexToAddress(token.Address)
		log.Infof("%s (%s)\n", name, addressHex.String())
		if client, ok := r.tokens[name]; ok {
			if client.Address != addressHex {
				log.Warnf("Token %s (%s) conflicts with pre-registered address %s", name, addressHex.String(), client.Address.String())
			}
			continue
		}
		client, err := NewClient(etherClient, addressHex, token.Decimals)
		if err != nil {
			return err
		}
		if err := r.Register(name, client); err != nil {
			return err
		}

	}
	return nil
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
