package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/bancor"
	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/token2"
)

var (
	clientBancorCommand                      = clientToken2Command.Command("bancor", "Bancor Commands")
	clientBancorGetConvertibleTokens         = clientBancorCommand.Command("get-convertible-tokens", "Get convertible tokens")
	clientBancorGetConversionPath            = clientBancorCommand.Command("get-conversion-path", "Get conversion path")
	clientBancorGetConversionPathSourceToken = clientBancorGetConversionPath.Arg("source-token", "Source Token").Required().String()
	clientBancorGetConversionPathTargetToken = clientBancorGetConversionPath.Arg("target-token", "Target Token").Required().String()
	clientBancorRatePath                     = clientBancorCommand.Command("rate-path", "Rate path from source to dest token")
	clientBancorRatePathSourceToken          = clientBancorRatePath.Arg("source-token", "Source Token").Required().String()
	clientBancorRatePathTargetToken          = clientBancorRatePath.Arg("target-token", "Target Token").Required().String()
	clientBancorRatePathAmount               = clientBancorRatePath.Arg("amount", "Input amount for rate").Required().String()
)

func doGetConvertibleTokens(ctx context.Context, reg *token2.Registry, b *bancor.Client) error {
	tokens, err := b.GetConvertibleTokens(ctx)
	if err != nil {
		return err
	}
	for _, t := range tokens {
		tokName, _, err := reg.ByAddress(t)
		if err != nil {
			tokName = "-"
		}
		fmt.Printf("%s\t%s\n", tokName, t.String())
	}
	return nil
}

func getAddress(reg *token2.Registry, tokenOrAddress string) (common.Address, error) {
	if tokenOrAddress == "eth" {
		return bancor.EthAddress, nil
	} else if strings.HasPrefix(tokenOrAddress, "0x") {
		return common.HexToAddress(tokenOrAddress), nil
	} else {
		tokClient, err := reg.ByName(tokenOrAddress)
		if err != nil {
			return common.HexToAddress("0x"), err
		}
		return tokClient.Address, nil
	}
}

func getName(reg *token2.Registry, address common.Address) (string, error) {
	if address == bancor.EthAddress {
		return "eth", nil
	}
	tokName, _, err := reg.ByAddress(address)
	if err != nil {
		return "-", err
	}
	return tokName, nil
}

func doGetConversionPath(ctx context.Context, reg *token2.Registry, b *bancor.Client) error {
	sourceAddress, err := getAddress(reg, *clientBancorGetConversionPathSourceToken)
	if err != nil {
		return err
	}
	targetAddress, err := getAddress(reg, *clientBancorGetConversionPathTargetToken)
	if err != nil {
		return err
	}
	path, err := b.GetConversionPath(ctx, sourceAddress, targetAddress)
	if err != nil {
		return err
	}
	if len(path) == 0 {
		return errors.New("No Path")
	}
	for _, t := range path {
		tokName, err := getName(reg, t)
		if err != nil {
			tokName = "-"
		}
		fmt.Printf("%s\t%s\n", tokName, t.String())
	}
	return nil
}

func doRatePath(ctx context.Context, reg *token2.Registry, b *bancor.Client) error {
	sourceAddress, err := getAddress(reg, *clientBancorRatePathSourceToken)
	if err != nil {
		return err
	}
	targetAddress, err := getAddress(reg, *clientBancorRatePathTargetToken)
	if err != nil {
		return err
	}
	path, err := b.GetConversionPath(ctx, sourceAddress, targetAddress)
	if err != nil {
		return err
	}
	if len(path) == 0 {
		return errors.New("No Path")
	}
	amount, err := decimal.NewFromString(*clientBancorRatePathAmount)
	if err != nil {
		return err
	}
	amountWei, err := b.ToWei(reg, sourceAddress, amount)
	if err != nil {
		return err
	}
	targetAmountWei, err := b.RateByPath(ctx, path, amountWei)
	if err != nil {
		return err
	}
	targetAmount, err := b.FromWei(reg, targetAddress, targetAmountWei)
	if err != nil {
		return err
	}
	fmt.Printf("Rate for %s from %s to %s: %s\n", amount, sourceAddress.String(), targetAddress.String(), targetAmount)

	return nil
}

func bancorCommands(ctx context.Context, client *client.Client, reg *token2.Registry, commands []string) (bool, error) {
	b, err := bancor.NewClient(ctx, client)
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "get-convertible-tokens":
		return true, doGetConvertibleTokens(ctx, reg, b)
	case "get-conversion-path":
		return true, doGetConversionPath(ctx, reg, b)
	case "rate-path":
		return true, doRatePath(ctx, reg, b)
	}
	return false, nil
}
