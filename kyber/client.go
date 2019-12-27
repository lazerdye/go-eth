package kyber

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	KyberNetworkProxy = "0x818E6FECD516Ecc3849DAf6845e3EC868087B755"
)

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
