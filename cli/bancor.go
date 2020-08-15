package main

import (
	"context"
	"fmt"

	"github.com/lazerdye/go-eth/bancor"
	"github.com/lazerdye/go-eth/client"
)

var (
	clientBancorCommand              = clientCmd.Command("bancor", "Bancor Commands")
	clientBancorGetConvertibleTokens = clientBancorCommand.Command("get-convertible-tokens", "Get convertible tokens")
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

func bancorCommands(ctx context.Context, client *client.Client, commands []string) (bool, error) {
	b, err := bancor.NewClient(ctx, client)
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "get-convertible-tokens":
		return true, doGetConvertibleTokens(ctx, b)
	}
	return false, nil
}
