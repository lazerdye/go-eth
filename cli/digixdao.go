package main

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/digixdao"
)

var (
	clientDigixDaoCommand = clientCmd.Command("digixdao", "Digix Dao commands")
	clientDigixDaoBurnCmd = clientDigixDaoCommand.Command("burn", "Burn Digix tokens")
)

func digixDaoCommands(ctx context.Context, client *client.Client, commands []string) error {
	dgd, err := digixdao.NewClient(client)
	if err != nil {
		return err
	}
	switch commands[0] {
	case "burn":
		return doDigixDaoBurn(ctx, dgd)
	default:
		return errors.Errorf("Unknown digixDao subcommand: %s", commands[0])
	}
}

func doDigixDaoBurn(ctx context.Context, dgd *digixdao.Client) error {
	account, err := getAccount(true)
	if err != nil {
		return err
	}

	tx, err := dgd.Burn(ctx, account)
	if err != nil {
		return err
	}
	fmt.Printf("Transaction: %s\n", tx.Hash().Hex())

	return nil

}
