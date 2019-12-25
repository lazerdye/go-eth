package kyber

import (
	"fmt"
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

func (c *Client) GetExpectedRate(source, dest common.Address, quantity int64) (*big.Int, error) {
	quantityBig := new(big.Int).SetInt64(quantity)
	rate, err := c.instance.GetExpectedRate(&bind.CallOpts{}, source, dest, quantityBig)
	if err != nil {
		return nil, err
	}
	fmt.Println("Rate: %+v", rate)
	return rate.ExpectedRate, nil
}
