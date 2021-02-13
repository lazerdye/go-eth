package main

import (
	"context"
	"fmt"

	"github.com/lazerdye/go-eth/digixdao"
	"github.com/pkg/errors"

	"github.com/lazerdye/go-eth/client"
)

var (
	clientDigixDaoCommand = clientCmd.Command("digixdao", "Digix Dao commands")
	clientDigixDaoBurnCmd = clientDigixDaoCommand.Command("burn", "Burn Digix tokens")
)

func digixDaoCommands(ctx context.Context, client *client.Client, commands []string) (bool, error) {
	dgd, err := digixdao.NewClient(client)
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "burn":
		return true, doDigixDaoBurn(ctx, dgd)
	}

	return false, nil
}

func doDigixDaoBurn(ctx context.Context, dgd *digixdao.Client) error {
	account, unlocked, err := getAccount()
	if err != nil {
		return err
	}
	if !unlocked {
		return errors.New("Wallet locked")
	}

	tx, err := dgd.Burn(ctx, account)
	if err != nil {
		return err
	}
	fmt.Printf("Transaction: %s\n", tx.Hash().Hex())

	return nil

}
