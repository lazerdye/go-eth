package main

import (
	"context"
	"fmt"

	"github.com/lazerdye/go-eth/token"
	"github.com/lazerdye/go-eth/uniswapv1"
)

var (
	clientUniswapCommand          = clientCmd.Command("uniswap", "Uniswap operations")
	clientUniswapGetExchange      = clientUniswapCommand.Command("get-exchange", "Get exchange address for token")
	clientUniswapGetExchangeToken = clientUniswapGetExchange.Arg("token", "Token to get exchange for").Required().String()
)

func uniswapGetExchange(ctx context.Context, client *uniswapv1.Client) error {
	tok, err := token.ByName(client.Client, *clientUniswapGetExchangeToken)
	if err != nil {
		return err
	}

	ex, err := client.GetExchange(ctx, tok)
	if err != nil {
		return err
	}

	fmt.Printf("Exchange: %s\n", ex.String())

	return nil
}
