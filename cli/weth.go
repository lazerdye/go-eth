package main

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/weth"
)

var (
	clientWethCommand       = clientCmd.Command("weth", "Weth Commands")
	clientWethDeposit       = clientWethCommand.Command("deposit", "Weth Deposit")
	clientWethDepositAmount = clientWethDeposit.Arg("amount", "Amount to depoit").Required().String()
)

func wethCommands(ctx context.Context, client *client.Client, commands []string) error {
	wethClient, err := weth.NewWethClient(client)
	if err != nil {
		return err
	}
	switch commands[0] {
	case "deposit":
		return doWethDeposit(ctx, wethClient)
	default:
		return errors.Errorf("Unknown weth subcommand: %s", commands[0])
	}
}

func doWethDeposit(ctx context.Context, wethClient *weth.WethClient) error {
	account, err := getAccount(true)
	if err != nil {
		return err
	}
	amount, _ := decimal.NewFromString(*clientWethDepositAmount)
	tx, err := wethClient.Deposit(ctx, account, amount)
	if err != nil {
		return err
	}
	fmt.Printf("TX: %s\n", tx.Hash())
	return nil
}
