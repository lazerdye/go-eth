package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/sushi"
	"github.com/lazerdye/go-eth/token2"
)

var (
	clientSushiCommand = clientToken2Command.Command("sushi", "SushiSwap operations")

	clientSushiGetPairs = clientSushiCommand.Command("get-pairs", "Get Pairs")
)

func sushiCommands(ctx context.Context, client *client.Client, reg *token2.Registry, commands []string) (bool, error) {
	sushiClient, err := sushi.NewClient(client)
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "get-pairs":
		return true, sushiGetPairs(ctx, sushiClient)
	}
	return false, nil
}

func sushiGetPairs(ctx context.Context, client *sushi.Client) error {
	pairs, err := client.GetPairs(ctx)
	if err != nil {
		return err
	}
	for _, pair := range pairs {
		token0, err := pair.Token0(ctx)
		if err != nil {
			return err
		}
		token0Symbol := "?"
		token0Client, err := token2.ByAddress(ctx, client.Client, token0)
		if err == nil {
			token0Symbol, err = token0Client.Symbol(ctx)
			if err != nil {
				log.Warnf("Token %s could not be initialized: %s", token0.String(), err)
			}
		}

		token1, err := pair.Token1(ctx)
		if err != nil {
			return err
		}
		token1Symbol := "?"
		token1Client, err := token2.ByAddress(ctx, client.Client, token1)
		if err == nil {
			token1Symbol, err = token1Client.Symbol(ctx)
			if err != nil {
				log.Warnf("Token %s could not be initialized: %s", token1.String(), err)
			}
		}
		fmt.Printf("%+v %s(%+v) -> %s(%+v)\n", pair.Address.String(), token0Symbol, token0.String(), token1Symbol, token1.String())
	}
	return nil
}
