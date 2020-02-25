package kyber

import (
	"fmt"
	"math"
	"math/big"

	"github.com/lazerdye/go-eth/wallet"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	KyberNetworkProxy = "0x818E6FECD516Ecc3849DAf6845e3EC868087B755"
)

var (
	gasLimit *big.Int
	gasPrice uint64
)

func init() {
	// TODO: Make this configurable.
	gasLimit = big.NewInt(5000000000)
	gasPrice = 780000
}

type Client struct {
	c        *ethclient.Client
	instance *Kyber
}

func NewClient(client *ethclient.Client) (*Client, error) {
	instance, err := NewKyber(common.HexToAddress(KyberNetworkProxy), client)
	if err != nil {
		return nil, err
	}
	return &Client{
		c:        client,
		instance: instance,
	}, nil
}

func (c *Client) GetExpectedRate(source, dest common.Address, quantity *big.Int) (*big.Float, *big.Float, error) {
	rate, err := c.instance.GetExpectedRate(&bind.CallOpts{}, source, dest, quantity)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("Raw: %+v", rate)

	expectedRate := new(big.Float).Quo(new(big.Float).SetInt(rate.ExpectedRate), big.NewFloat(math.Pow10(18)))
	slippageRate := new(big.Float).Quo(new(big.Float).SetInt(rate.SlippageRate), big.NewFloat(math.Pow10(18)))
	return expectedRate, slippageRate, nil
}

func (c *Client) SwapEtherToToken(account *wallet.Account, token common.Address, minRate *big.Float) (*types.Transaction, error) {
	transactOpts, err := account.NewTransactor(gasLimit, gasPrice)
	if err != nil {
		return nil, err
	}
	minRateInt, _ := new(big.Float).Mul(minRate, big.NewFloat(math.Pow10(18))).Int(nil)
	transaction, err := c.instance.SwapEtherToToken(transactOpts, token, minRateInt)
	if err != nil {
		return nil, err
	}
	return transaction, err
}

func (c *Client) SwapTokenToEther(account *wallet.Account, token common.Address, amount *big.Float, maxRate *big.Float) (*types.Transaction, error) {
	transactOpts, err := account.NewTransactor(gasLimit, gasPrice)
	if err != nil {
		return nil, err
	}
	// TODO: What about when decimals != 18?
	amountInt, _ := new(big.Float).Mul(amount, big.NewFloat(math.Pow10(18))).Int(nil)
	maxRateInt, _ := new(big.Float).Mul(maxRate, big.NewFloat(math.Pow10(18))).Int(nil)
	transaction, err := c.instance.SwapTokenToEther(transactOpts, token, amountInt, maxRateInt)
	if err != nil {
		return nil, err
	}
	return transaction, err
}
