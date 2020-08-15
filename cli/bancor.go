package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/lazerdye/go-eth/bancor"
	"github.com/lazerdye/go-eth/client"
)

var (
	clientBancorCommand                      = clientCmd.Command("bancor", "Bancor Commands")
	clientBancorGetConvertibleTokens         = clientBancorCommand.Command("get-convertible-tokens", "Get convertible tokens")
	clientBancorGetConversionPath            = clientBancorCommand.Command("get-conversion-path", "Get conversion path")
	clientBancorGetConversionPathSourceToken = clientBancorGetConversionPath.Arg("source-token", "Source Token").Required().String()
	clientBancorGetConversionPathTargetToken = clientBancorGetConversionPath.Arg("target-token", "Target Token").Required().String()
)

func doGetConvertibleTokens(ctx context.Context, b *bancor.Client) error {
	tokens, err := b.GetConvertibleTokens(ctx)
	if err != nil {
		return err
	}
	for _, t := range tokens {
		fmt.Printf("%s\n", t.String())
	}
	return nil
}

func doGetConversionPath(ctx context.Context, b *bancor.Client) error {
	path, err := b.GetConversionPath(ctx,
		common.HexToAddress(*clientBancorGetConversionPathSourceToken),
		common.HexToAddress(*clientBancorGetConversionPathTargetToken))
	if err != nil {
		return err
	}
	if len(path) == 0 {
		fmt.Printf("No path\n")
	}
	for _, t := range path {
		fmt.Printf("%s\n", t.String())
	}
	return nil
}

func bancorCommands(ctx context.Context, client *client.Client, commands []string) (bool, error) {
	b, err := bancor.NewClient(ctx, client)
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "get-convertible-tokens":
		return true, doGetConvertibleTokens(ctx, b)
	case "get-conversion-path":
		return true, doGetConversionPath(ctx, b)
	}
	return false, nil
}
