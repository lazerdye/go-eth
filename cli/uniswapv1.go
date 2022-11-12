package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/uniswapv1"
)

var (
	clientUniswap1Command                              = clientToken2Command.Command("uniswapv1", "Uniswap v1 operations")
	clientUniswap1Address                              = clientUniswap1Command.Flag("token-address", "Token to get exchange for").String()
	clientUniswap1Decimals                             = clientUniswap1Command.Flag("token-decimals", "Token to get exchange for").Uint8()
	clientUniswap1Token                                = clientUniswap1Command.Flag("token", "Token to get exchange for").String()
	clientUniswap1GetExchange                          = clientUniswap1Command.Command("get-exchange", "Get exchange address for token")
	clientUniswap1GetEthToTokenInputPrice              = clientUniswap1Command.Command("eth-to-token-input", "Get eth to token input price")
	clientUniswap1GetEthToTokenInputPriceEthSold       = clientUniswap1GetEthToTokenInputPrice.Arg("eth-sold", "Ethereum sold").Required().Float64()
	clientUniswap1GetEthToTokenOutputPrice             = clientUniswap1Command.Command("eth-to-token-output", "Get eth to token output price")
	clientUniswap1GetEthToTokenOutputPriceTokensBought = clientUniswap1GetEthToTokenOutputPrice.Arg("tokens-bought", "Tokens bought").Required().Float64()
	clientUniswap1GetTokenToEthInputPrice              = clientUniswap1Command.Command("token-to-eth-input", "Get token to eth input price")
	clientUniswap1GetTokenToEthInputPriceEthBought     = clientUniswap1GetTokenToEthInputPrice.Arg("eth-bought", "Ethereum bought").Required().Float64()
	clientUniswap1GetTokenToEthOutputPrice             = clientUniswap1Command.Command("token-to-eth-output", "Get token to eth output price")
	clientUniswap1GetTokenToEthOutputPriceTokensSold   = clientUniswap1GetTokenToEthOutputPrice.Arg("tokens-sold", "Tokens sold").Required().Float64()
	clientUniswap1TokenToEthSwapInput                  = clientUniswap1Command.Command("token-to-eth-swap-input", "Swap to eth the number of tokens")
	clientUniswap1TokenToEthSwapInputTokensSold        = clientUniswap1TokenToEthSwapInput.Arg("tokens-sold", "Tokens sold").Required().Float64()
	clientUniswap1TokenToEthSwapInputMinEth            = clientUniswap1TokenToEthSwapInput.Arg("min-eth", "Minimum eth received").Required().Float64()
	clientUniswap1Graph                                = clientUniswap1Command.Command("graph", "Query the graph")
)

const (
	deadlineOffset = 1200 // 20 minutes
)

func uniswapV1Commands(ctx context.Context, client *client.Client, reg *token2.Registry, commands []string) (bool, error) {
	uniswapv1Client, err := uniswapv1.NewClient(client)
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "get-exchange":
		return true, uniswapGetExchange(ctx, reg, uniswapv1Client)
	case "eth-to-token-input":
		return true, uniswapGetEthToTokenInputPrice(ctx, reg, uniswapv1Client)
	case "eth-to-token-output":
		return true, uniswapGetEthToTokenOutputPrice(ctx, reg, uniswapv1Client)
	case "token-to-eth-input":
		return true, uniswapGetTokenToEthInputPrice(ctx, reg, uniswapv1Client)
	case "token-to-eth-output":
		return true, uniswapGetTokenToEthOutputPrice(ctx, reg, uniswapv1Client)
	case "graph":
		return true, uniswapGraph(ctx, uniswapv1Client)
	case "token-to-eth-swap-input":
		return true, uniswapTokenToEthSwapInput(ctx, reg, uniswapv1Client)
	}
	return false, nil
}

func getExchange(ctx context.Context, reg *token2.Registry, client *uniswapv1.Client) (*uniswapv1.ExchangeClient, error) {
	if *clientUniswap1Token == "" && *clientUniswap1Address == "" {
		return nil, errors.New("Either --token-address/--token-decimals or --token is required")
	}
	var tok *token2.Client
	var err error
	if *clientUniswap1Token != "" {
		tok, err = reg.ByName(*clientUniswap1Token)
		if err != nil {
			return nil, err
		}
	} else {
		if *clientUniswap1Decimals == 0 {
			return nil, errors.New("--decimals is required")
		}
		tok, err = token2.NewClient(client.Client, common.HexToAddress(*clientUniswap1Address), *clientUniswap1Decimals)
		if err != nil {
			return nil, err
		}
	}

	return client.GetExchange(ctx, tok)
}

func uniswapGetExchange(ctx context.Context, reg *token2.Registry, client *uniswapv1.Client) error {
	ex, err := getExchange(ctx, reg, client)
	if err != nil {
		return err
	}

	fmt.Printf("Exchange: %s\n", ex.ContractAddress().String())

	return nil
}

func uniswapGetEthToTokenInputPrice(ctx context.Context, reg *token2.Registry, client *uniswapv1.Client) error {
	ex, err := getExchange(ctx, reg, client)
	if err != nil {
		return err
	}

	ethSold := decimal.NewFromFloat(*clientUniswap1GetEthToTokenInputPriceEthSold)
	tokenBought, err := ex.GetEthToTokenInputPrice(ctx, ethSold)
	if err != nil {
		return err
	}

	fmt.Printf("Token bought: %s\n", tokenBought)

	return nil
}

func uniswapGetEthToTokenOutputPrice(ctx context.Context, reg *token2.Registry, client *uniswapv1.Client) error {
	ex, err := getExchange(ctx, reg, client)
	if err != nil {
		return err
	}

	tokensBought := decimal.NewFromFloat(*clientUniswap1GetEthToTokenOutputPriceTokensBought)
	ethSold, err := ex.GetEthToTokenOutputPrice(ctx, tokensBought)
	if err != nil {
		return err
	}

	fmt.Printf("Eth sold: %s\n", ethSold)

	return nil
}

func uniswapGetTokenToEthInputPrice(ctx context.Context, reg *token2.Registry, client *uniswapv1.Client) error {
	ex, err := getExchange(ctx, reg, client)
	if err != nil {
		return err
	}

	ethBought := decimal.NewFromFloat(*clientUniswap1GetTokenToEthInputPriceEthBought)
	tokenSold, err := ex.GetTokenToEthInputPrice(ctx, ethBought)
	if err != nil {
		return err
	}

	fmt.Printf("Eth bought: %s\n", tokenSold)

	return nil
}

func uniswapGetTokenToEthOutputPrice(ctx context.Context, reg *token2.Registry, client *uniswapv1.Client) error {
	ex, err := getExchange(ctx, reg, client)
	if err != nil {
		return err
	}

	tokensSold := decimal.NewFromFloat(*clientUniswap1GetTokenToEthOutputPriceTokensSold)
	ethBought, err := ex.GetTokenToEthOutputPrice(ctx, tokensSold)
	if err != nil {
		return err
	}

	fmt.Printf("Token sold: %s\n", ethBought)

	return nil
}

func uniswapGraph(ctx context.Context, client *uniswapv1.Client) error {
	graph := uniswapv1.NewGraph()

	exchanges, err := graph.ExchangesByVolume(ctx)
	if err != nil {
		return err
	}

	for _, exchange := range exchanges {
		fmt.Printf("Token %s (%s): %s\n", exchange.TokenName, exchange.TokenSymbol, exchange.TradeVolumeEth)
	}

	return nil
}

func uniswapTokenToEthSwapInput(ctx context.Context, reg *token2.Registry, client *uniswapv1.Client) error {
	ex, err := getExchange(ctx, reg, client)
	if err != nil {
		return err
	}
	tokensSold := decimal.NewFromFloat(*clientUniswap1TokenToEthSwapInputTokensSold)
	minEth := decimal.NewFromFloat(*clientUniswap1TokenToEthSwapInputMinEth)

	account, err := getAccount(true)
	if err != nil {
		return err
	}

	deadline := int(time.Now().Unix() + deadlineOffset)

	tx, err := ex.TokenToEthSwapInput(ctx, account, tokensSold, minEth, deadline)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", tx.Hash().Hex())

	return nil
}
