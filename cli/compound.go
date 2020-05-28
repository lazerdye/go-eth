package main

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/compound"
	"github.com/lazerdye/go-eth/wallet"
)

var (
	clientCompoundCommand = clientCmd.Command("compound", "Compound operations")
	clientMintCommand     = clientCompoundCommand.Command("mint", "Mint currency")
	clientMintToken       = clientMintCommand.Arg("token", "Token to mint").Required().String()
	clientMintAmount      = clientMintCommand.Arg("amount", "Amount to mint").Required().Float64()
)

func compoundMint(ctx context.Context, client *compound.Client, account *wallet.Account) error {
	tokName := *clientMintToken
	if tokName != "eth" {
		return errors.New("Currently only eth is supported")
	}

	amountBig := decimal.NewFromFloat(*clientMintAmount)
	tx, err := client.Mint(ctx, account, amountBig)
	if err != nil {
		return err
	}
	fmt.Printf("Mint TX: %s\n", tx.Hash().String())

	return nil
}
