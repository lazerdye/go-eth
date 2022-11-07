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

func wethCommands(ctx context.Context, client *client.Client, commands []string) (bool, error) {
	wethClient, err := weth.NewWethClient(client)
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "deposit":
		return true, doWethDeposit(ctx, wethClient)
	}
	return false, nil
}

func doWethDeposit(ctx context.Context, wethClient *weth.WethClient) error {
	account, unlocked, err := getAccount()
	if err != nil {
		return err
	}
	if !unlocked {
		return errors.New("Wallet Locked")
	}
	amount, _ := decimal.NewFromString(*clientWethDepositAmount)
	tx, err := wethClient.Deposit(ctx, account, amount)
	if err != nil {
		return err
	}
	fmt.Printf("TX: %s\n", tx.Hash())
	return nil
}
