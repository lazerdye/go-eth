package main

import (
	"context"

	"github.com/pkg/errors"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/eth2"
	"github.com/lazerdye/go-eth/token2"
)

var (
	clientEth2Command        = clientToken2Command.Command("eth2", "ETH2 Commands")
	clientEth2DepositCommand = clientEth2Command.Command("deposit", "Deposit into eth2 constract")
)

func eth2Commands(ctx context.Context, client *client.Client, reg *token2.Registry, commands []string) (bool, error) {
	_, err := eth2.NewClient(client, reg)
	if err != nil {
		return false, err
	}

	return false, errors.New("Not Implemented")
}
