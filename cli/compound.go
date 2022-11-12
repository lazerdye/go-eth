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
	clientCompoundToken   = clientCompoundCommand.Flag("token", "Token for compound").String()
	clientMintCommand     = clientCompoundCommand.Command("mint", "Mint currency")
	clientMintAmount      = clientMintCommand.Arg("amount", "Amount to mint").Required().Float64()
	clientRedeemCommand   = clientCompoundCommand.Command("redeem", "Redeem currency")
	clientRedeemAmount    = clientRedeemCommand.Arg("amount", "Amount to redeem").Required().String()
)

func compoundCommands(ctx context.Context, client *client.Client, reg *token2.Registry, commands []string) error {
	var compoundClient compound.Client
	var err error
	switch *clientCompoundToken {
	case "eth":
		compoundClient, err = compound.NewCEthClient(client)
		if err != nil {
			return err
		}
	case "bat", "dai", "zrx", "uni", "usdc", "usdt":
		tok, err := reg.ByName(*clientCompoundToken)
		if err != nil {
			return err
		}
		compoundClient, err = compound.NewCErc20Client(ctx, client, tok.Address)
		if err != nil {
			return err
		}
	case "":
		return errors.New("--token required")
	default:
		return errors.Errorf("Unsupported token: %s", *clientCompoundToken)
	}

	switch commands[0] {
	case "mint":
		return compoundMint(ctx, compoundClient)
	case "redeem":
		return compoundRedeem(ctx, compoundClient)
	default:
		return errors.Errorf("Unknown compound subcommand: %s", commands[0])
	}
}

func compoundMint(ctx context.Context, compoundClient compound.Client) error {
	account, err := getAccount(true)
	if err != nil {
		return err
	}

	amountBig := decimal.NewFromFloat(*clientMintAmount)
	tx, err := compoundClient.Mint(ctx, account, amountBig)
	if err != nil {
		return err
	}
	fmt.Printf("Mint TX: %s\n", tx.Hash().String())

	return nil
}

func compoundRedeem(ctx context.Context, compoundClient compound.Client) error {
	account, err := getAccount(true)
	if err != nil {
		return err
	}

	amountBig, err := decimal.NewFromString(*clientRedeemAmount)
	if err != nil {
		return err
	}
	tx, err := compoundClient.Redeem(ctx, account, amountBig)
	if err != nil {
		return err
	}
	fmt.Printf("Redeem TX: %s\n", tx.Hash().String())

	return nil
}
