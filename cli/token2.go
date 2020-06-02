package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/token2"
)

var (
	clientToken2Command           = clientCmd.Command("token2", "Token2 Commands")
	clientToken2TokenFile         = clientToken2Command.Flag("token-file", "File for token registry").Envar("TOKEN_FILE").Required().String()
	clientToken2BalanceOf         = clientToken2Command.Command("balance-of", "Balance of token in other contract")
	clientToken2BalanceOfToken    = clientToken2BalanceOf.Arg("token", "Token for balance-of").String()
	clientToken2BalanceOfAddress  = clientToken2BalanceOf.Arg("address", "Address for balance-of").String()
	clientToken2AllowanceCommand  = clientToken2Command.Command("allowance", "Display token allowance")
	clientToken2AllowanceToken    = clientToken2AllowanceCommand.Arg("token", "Token for allowance").Required().String()
	clientToken2AllowanceContract = clientToken2AllowanceCommand.Arg("contract", "Contract to check allowance of").Required().String()
	clientToken2ApproveCommand    = clientToken2Command.Command("approve", "Approve token for contract")
	clientToken2ApproveToken      = clientToken2ApproveCommand.Arg("token", "Token for allowance").Required().String()
	clientToken2ApproveContract   = clientToken2ApproveCommand.Arg("contract", "Contract to approve").Required().String()
	clientToken2ApproveAmount     = clientToken2ApproveCommand.Arg("amount", "Amount to approve").Required().Float64()
)

func newTokenRegistry(ctx context.Context, ethClient *client.Client) (*token2.Registry, error) {
	return token2.RegistryFromFile(ethClient, *clientToken2TokenFile)
}

func doTokenBalanceOf(ctx context.Context, reg *token2.Registry) error {
	token, err := reg.ByName(*clientToken2BalanceOfToken)
	if err != nil {
		return err
	}
	address := common.HexToAddress(*clientToken2BalanceOfAddress)
	balance, err := token.BalanceOf(ctx, address)
	if err != nil {
		return err
	}
	fmt.Printf("Balance of %s on address %s is: %s\n", *clientToken2BalanceOfToken, address.String(), balance)
	return nil
}

func doTokenAllowance(ctx context.Context, reg *token2.Registry) error {
	token, err := reg.ByName(*clientToken2AllowanceToken)
	if err != nil {
		return err
	}
	contract := common.HexToAddress(*clientToken2AllowanceContract)
	address := common.HexToAddress(*address)
	allowance, err := token.Allowance(ctx, address, contract)
	if err != nil {
		return err
	}
	fmt.Printf("Allowance: %s\n", allowance.String())
	return nil
}

func doTokenApprove(ctx context.Context, reg *token2.Registry) error {
	token, err := reg.ByName(*clientToken2ApproveToken)
	if err != nil {
		return err
	}
	account, unlocked, err := getAccount()
	if err != nil {
		return err
	}
	if !unlocked {
		return errors.New("Wallet locked")
	}
	contract := common.HexToAddress(*clientToken2ApproveContract)
	amount := decimal.NewFromFloat(*clientToken2ApproveAmount)
	transaction, err := token.Approve(ctx, account, contract, amount)
	if err != nil {
		return err
	}
	fmt.Printf("Transaction: %s\n", transaction.Hash().Hex())
	return nil
}

func token2Commands(ctx context.Context, client *client.Client, commands []string) (bool, error) {
	reg, err := newTokenRegistry(ctx, client)
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "balance-of":
		return true, doTokenBalanceOf(ctx, reg)
	case "allowance":
		return true, doTokenAllowance(ctx, reg)
	case "approve":
		return true, doTokenApprove(ctx, reg)
	case "uniswapv2":
		return uniswapV2Commands(ctx, client, reg, commands[1:])
	}

	return false, nil
}
