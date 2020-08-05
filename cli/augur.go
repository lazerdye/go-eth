package main

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/lazerdye/go-eth/augur"
	"github.com/lazerdye/go-eth/client"
)

var (
	clientAugurCommand           = clientCmd.Command("augur", "Augur Commands")
	clientAugurMigrateFromLegacy = clientAugurCommand.Command("migrate-from-legacy", "Migrate from Legacy")
)

func doMigrateFromLegacy(ctx context.Context, augur *augur.Client) error {
	account, unlocked, err := getAccount()
	if err != nil {
		return err
	}
	if !unlocked {
		return errors.New("Wallet locked")
	}

	tx, err := augur.MigrateFromLegacyReputationToken(ctx, account)
	if err != nil {
		return err
	}
	fmt.Printf("Transaction: %s\n", tx.Hash().Hex())

	return nil
}

func augurCommands(ctx context.Context, client *client.Client, commands []string) (bool, error) {
	augur, err := augur.NewClient(client)
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "migrate-from-legacy":
		return true, doMigrateFromLegacy(ctx, augur)
	}

	return false, nil
}
