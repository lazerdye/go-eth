package main

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/sushi"
	"github.com/lazerdye/go-eth/token2"
)

var (
	clientSushiCommand = clientToken2Command.Command("sushi", "SushiSwap operations")

	clientSushiGetPairs           = clientSushiCommand.Command("get-pairs", "Get Pairs")
	clientSushiAmountsIn          = clientSushiCommand.Command("amounts-in", "Get Amounts In")
	clientSushiAmountsInToken0    = clientSushiAmountsIn.Arg("token0", "Token 0").String()
	clientSushiAmountsInToken1    = clientSushiAmountsIn.Arg("token1", "Token 2").String()
	clientSushiAmountsInAmountOut = clientSushiAmountsIn.Arg("amount-out", "Amount Out").String()
	clientSushiAmountsOut         = clientSushiCommand.Command("amounts-out", "Get Amounts Out")
	clientSushiAmountsOutToken0   = clientSushiAmountsOut.Arg("token0", "Token 0").String()
	clientSushiAmountsOutToken1   = clientSushiAmountsOut.Arg("token1", "Token 2").String()
	clientSushiAmountsOutAmountIn = clientSushiAmountsOut.Arg("amount-in", "Amount In").String()
)

func sushiCommands(ctx context.Context, client *client.Client, reg *token2.Registry, commands []string) (bool, error) {
	sushiClient, err := sushi.NewClient(client)
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "get-pairs":
		return true, sushiGetPairs(ctx, sushiClient)
	case "amounts-in":
		return true, sushiAmountsIn(ctx, sushiClient, reg)
	case "amounts-out":
		return true, sushiAmountsOut(ctx, sushiClient, reg)
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

func sushiAmountsIn(ctx context.Context, client *sushi.Client, reg *token2.Registry) error {
	token0, err := reg.ByName(*clientSushiAmountsInToken0)
	if err != nil {
		return err
	}
	token1, err := reg.ByName(*clientSushiAmountsInToken1)
	if err != nil {
		return err
	}
	amountOut, err := decimal.NewFromString(*clientSushiAmountsInAmountOut)
	if err != nil {
		return err
	}
	value, err := client.GetAmountsIn(ctx, amountOut, []*token2.Client{token0, token1})
	if err != nil {
		return err
	}
	fmt.Printf("AmountOut: %s\n", value.String())
	return nil
}

func sushiAmountsOut(ctx context.Context, client *sushi.Client, reg *token2.Registry) error {
	token0, err := reg.ByName(*clientSushiAmountsOutToken0)
	if err != nil {
		return err
	}
	token1, err := reg.ByName(*clientSushiAmountsOutToken1)
	if err != nil {
		return err
	}
	amountIn, err := decimal.NewFromString(*clientSushiAmountsOutAmountIn)
	if err != nil {
		return err
	}
	value, err := client.GetAmountsOut(ctx, amountIn, []*token2.Client{token0, token1})
	if err != nil {
		return err
	}
	fmt.Printf("AmountIn: %s\n", value.String())
	return nil
}
