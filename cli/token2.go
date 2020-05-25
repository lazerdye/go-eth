package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/token2"
)

var (
	clientToken2Command           = clientCmd.Command("token2", "Token2 Commands")
	clientToken2Token             = clientToken2Command.Flag("token", "Token to operate on").Required().String()
	clientToken2TokenFile         = clientCmd.Flag("token-file", "File for token registry").Envar("TOKEN_FILE").Required().String()
	clientToken2AllowanceCommand  = clientToken2Command.Command("allowance", "Display token allowance")
	clientToken2AllowanceContract = clientToken2AllowanceCommand.Arg("contract", "Contract to check allowance of").Required().String()
	clientToken2ApproveCommand    = clientToken2Command.Command("approve", "Approve token for contract")
	clientToken2ApproveContract   = clientToken2ApproveCommand.Arg("contract", "Contract to approve").Required().String()
	clientToken2ApproveAmount     = clientToken2ApproveCommand.Arg("amount", "Amount to approve").Required().Float64()
)

func newTokenRegistry(ctx context.Context, ethClient *client.Client) (*token2.Registry, error) {
	return token2.RegistryFromFile(ethClient, *clientToken2TokenFile)
}

func doTokenAllowance(ctx context.Context, token *token2.Client) error {
	contract := common.HexToAddress(*clientToken2AllowanceContract)
	address := common.HexToAddress(*address)
	allowance, err := token.Allowance(ctx, address, contract)
	if err != nil {
		return err
	}
	fmt.Printf("Allowance: %s\n", allowance.String())
	return nil
}

func doTokenApprove(ctx context.Context, token *token2.Client) error {
	account, unlocked, err := getAccount()
	if err != nil {
		return err
	}
	if !unlocked {
		return errors.New("Wallet locked")
	}
	contract := common.HexToAddress(*clientToken2ApproveContract)
	amount := big.NewFloat(*clientToken2ApproveAmount)
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
	token, err := reg.ByName(*clientToken2Token)
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "allowance":
		return true, doTokenAllowance(ctx, token)
	case "approve":
		return true, doTokenApprove(ctx, token)
	}

	return false, nil
}
