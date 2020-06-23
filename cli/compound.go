package main

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/compound"
	"github.com/lazerdye/go-eth/token2"
)

var (
	clientCompoundCommand = clientToken2Command.Command("compound", "Compound operations")
	clientMintCommand     = clientCompoundCommand.Command("mint", "Mint currency")
	clientMintToken       = clientMintCommand.Arg("token", "Token to mint").Required().String()
	clientMintAmount      = clientMintCommand.Arg("amount", "Amount to mint").Required().Float64()
)

func compoundCommands(ctx context.Context, client *client.Client, reg *token2.Registry, commands []string) (bool, error) {
	switch commands[0] {
	case "mint":
		return true, compoundMint(ctx, client)
	}
	return false, nil
}

func compoundMint(ctx context.Context, client *client.Client) error {
	compoundClient, err := compound.NewClient(client, *clientMintToken)
	if err != nil {
		return err
	}
	account, opened, err := getAccount()
	if err != nil {
		return err
	}
	if !opened {
		return errors.New("Account not unlocked")
	}


	amountBig := decimal.NewFromFloat(*clientMintAmount)
	tx, err := compoundClient.Mint(ctx, account, amountBig)
	if err != nil {
		return err
	}
	fmt.Printf("Mint TX: %s\n", tx.Hash().String())

	return nil
}
