package main

import (
	"context"
	"fmt"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/eth2"
	"github.com/lazerdye/go-eth/token2"
)

var (
	clientEth2Command            = clientToken2Command.Command("eth2", "ETH2 Commands")
	clientEth2DepositCommand     = clientEth2Command.Command("deposit", "Deposit into eth2 constract")
	clientEth2DepositDepositData = clientEth2DepositCommand.Arg("deposit-data", "JSON deposit data").Required().String()
)

func eth2Commands(ctx context.Context, client *client.Client, reg *token2.Registry, commands []string) (bool, error) {
	eth2Client, err := eth2.NewClient(client, reg)
	if err != nil {
		return false, err
	}

	switch commands[0] {
	case "deposit":
		return true, doDeposit(ctx, eth2Client)
	}

	return false, nil
}

func doDeposit(ctx context.Context, eth2Client *eth2.Client) error {
	account, err := getAccount(true)
	if err != nil {
		return err
	}

	txs, err := eth2Client.Deposit(ctx, account, *clientEth2DepositDepositData)
	if err != nil {
		return err
	}

	for i, tx := range txs {
		fmt.Printf("Transaction %d: %s", i, tx.Hash().String())
	}
	return nil
}
