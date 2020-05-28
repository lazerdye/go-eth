package main

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/wallet"
	"github.com/lazerdye/go-eth/zeroex"
)

var (
	clientZeroexCommand       = clientCmd.Command("zeroex", "Zeroex operations")
	clientZeroexDeposit       = clientZeroexCommand.Command("deposit", "Deposit eth into zeroex")
	clientZeroexDepositAmount = clientZeroexDeposit.Arg("amount", "Amount to deposit").Required().String()
	clientZeroexBalanceOf     = clientZeroexCommand.Command("balanceof", "Eth (weth) balance for transactions")
)

func zeroexDepositCommand(client *zeroex.Client, account *wallet.Account) error {
	amount, err := decimal.NewFromString(*clientZeroexDepositAmount)
	if err != nil {
		return err
	}
	log.Infof("Amount: %s", amount)
	tx, err := client.EtherTokenDeposit(context.Background(), account, amount)
	if err != nil {
		return err
	}
	fmt.Printf("Result: %s\n", tx.Hash().String())
	return nil
}

func zeroexBalanceOfCommand(client *zeroex.Client, account *wallet.Account) error {
	amount, err := client.EtherTokenBalanceOf(context.Background(), account)
	if err != nil {
		return err
	}
	fmt.Printf("Balance of account %s: %s\n", account.Address().String(), amount.String())
	return nil
}
