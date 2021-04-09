package main

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/oneinchv3"
	"github.com/lazerdye/go-eth/token2"
)

var (
	clientOneInchV3Command        = clientToken2Command.Command("1inchv3", "1inch v3 commands")
	clientOneInchV3TokensCommand  = clientOneInchV3Command.Command("tokens", "1inch list tokens")
	clientOneInchV3QuoteCommand   = clientOneInchV3Command.Command("quote", "1inch quote")
	clientOneInchV3QuoteFromToken = clientOneInchV3QuoteCommand.Arg("from-token", "From Token").Required().String()
	clientOneInchV3QuoteToToken   = clientOneInchV3QuoteCommand.Arg("to-token", "To Token").Required().String()
	clientOneInchV3QuoteAmount    = clientOneInchV3QuoteCommand.Arg("amount", "Amount").Required().String()
)

func oneInchV3Commands(ctx context.Context, client *client.Client, reg *token2.Registry, commands []string) (bool, error) {
	oneInchV3 := oneinchv3.NewClient()
	switch commands[0] {
	case "tokens":
		return true, oneInchV3Tokens(ctx, oneInchV3)
	case "quote":
		return true, oneInchV3Quote(ctx, reg, oneInchV3)
	}
	return false, nil
}

func oneInchV3Tokens(ctx context.Context, oneInchV3 *oneinchv3.Client) error {
	tokens, err := oneInchV3.Tokens(ctx)
	if err != nil {
		return err
	}
	for _, token := range tokens {
		fmt.Printf("%s %s (%s) dec:%d\n", token.Address.String(), token.Symbol, token.Name, token.Decimals)
	}
	return nil
}

func oneInchV3Quote(ctx context.Context, reg *token2.Registry, oneInchV3 *oneinchv3.Client) error {
	var fromToken, toToken *token2.Client
	var err error
	if *clientOneInchV3QuoteFromToken != "eth" {
		fromToken, err = reg.ByName(*clientOneInchV3QuoteFromToken)
		if err != nil {
			return err
		}
	}
	if *clientOneInchV3QuoteToToken != "eth" {
		toToken, err = reg.ByName(*clientOneInchV3QuoteToToken)
		if err != nil {
			return err
		}
	}
	amount, err := decimal.NewFromString(*clientOneInchV3QuoteAmount)
	if err != nil {
		return err
	}
	quote, amount, err := oneInchV3.Quote(ctx, fromToken, toToken, amount)
	if err != nil {
		return err
	}
	fmt.Printf("%s %+v\n", amount.String(), *quote)
	return nil
}
