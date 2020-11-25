package eth2

import (
	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/token2"
)

type Client struct {
	*client.Client

	deposit *Deposit
}

func NewClient(client *client.Client, reg *token2.Registry) (*Client, error) {
	depositToken, err := reg.ByName("_eth2_deposit")
	if err != nil {
		return nil, err
	}
	deposit, err := NewDeposit(depositToken.Address, client)
	if err != nil {
		return nil, err
	}

	return &Client{client, deposit}, nil
}
