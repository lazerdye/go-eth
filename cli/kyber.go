package main

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/kyber"
	"github.com/lazerdye/go-eth/token2"
)

var (
	tokenKyberCmd          = clientCmd.Command("kyber", "Kyber Network")
	tokenKyberTokenFile    = tokenKyberCmd.Flag("token-file", "Token file").Required().String()
	tokenKyberSource       = tokenKyberCmd.Flag("source", "Source exchange").String()
	tokenKyberDest         = tokenKyberCmd.Flag("dest", "Dest exchange").String()
	tokenKyberQuantity     = tokenKyberCmd.Flag("quantity", "Source quantity").Float64()
	tokenKyberExpectedRate = tokenKyberCmd.Command("expectedRate", "Expected rate")
)

func kyberCommands(ctx context.Context, client *client.Client, commands []string) error {
	k, err := kyber.NewClient(client)
	if err != nil {
		return err
	}
	switch commands[0] {
	case "expectedRate":
		return doClientKyberExpectedRate(ctx, k)
	default:
		return errors.Errorf("Unknown kyber subcommand: %s", commands[0])
	}
	return nil
}

func doClientKyberExpectedRate(_ context.Context, k *kyber.Client) error {
	if *tokenKyberSource == "" || *tokenKyberDest == "" {
		return errors.Errorf("Both --source and --dest flags required")
	}
	r, err := token2.RegistryFromFile(k.Client, *tokenKyberTokenFile)
	if err != nil {
		return err
	}
	quantityFloat := decimal.NewFromFloat(*tokenKyberQuantity)
	sourceToken, err := r.ByName(*tokenKyberSource)
	if err != nil {
		return err
	}
	destToken, err := r.ByName(*tokenKyberDest)
	if err != nil {
		return err
	}
	expectedRate, slippageRate, err := k.GetExpectedRate(context.Background(), sourceToken, destToken, quantityFloat)
	if err != nil {
		return err
	}
	fmt.Printf("Expected Rate: %s\n", expectedRate.String())
	fmt.Printf("Slippage Rate: %s\n", slippageRate.String())
	return nil
}
