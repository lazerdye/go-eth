package main

import (
	"context"
	"fmt"

	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/uniswapv1"
	"github.com/shopspring/decimal"
)

var (
	clientUniswap1Command                              = clientCmd.Command("uniswapv1", "Uniswap v1 operations")
	clientUniswap1Token                                = clientUniswap1Command.Flag("token", "Token to get exchange for").Required().String()
	clientUniswap1TokenFile                            = clientUniswap1Command.Flag("token-file", "File for token registry").Envar("TOKEN_FILE").Required().String()
	clientUniswap1GetExchange                          = clientUniswap1Command.Command("get-exchange", "Get exchange address for token")
	clientUniswap1GetEthToTokenInputPrice              = clientUniswap1Command.Command("eth-to-token-input", "Get eth to token input price")
	clientUniswap1GetEthToTokenInputPriceEthSold       = clientUniswap1GetEthToTokenInputPrice.Arg("eth-sold", "Ethereum sold").Required().Float64()
	clientUniswap1GetEthToTokenOutputPrice             = clientUniswap1Command.Command("eth-to-token-output", "Get eth to token output price")
	clientUniswap1GetEthToTokenOutputPriceTokensBought = clientUniswap1GetEthToTokenOutputPrice.Arg("tokens-bought", "Tokens bought").Required().Float64()
	clientUniswap1GetTokenToEthInputPrice              = clientUniswap1Command.Command("token-to-eth-input", "Get token to eth input price")
	clientUniswap1GetTokenToEthInputPriceEthBought     = clientUniswap1GetTokenToEthInputPrice.Arg("eth-bought", "Ethereum bought").Required().Float64()
	clientUniswap1GetTokenToEthOutputPrice             = clientUniswap1Command.Command("token-to-eth-output", "Get token to eth output price")
	clientUniswap1GetTokenToEthOutputPriceTokensSold   = clientUniswap1GetTokenToEthOutputPrice.Arg("tokens-sold", "Tokens sold").Required().Float64()
)

func getExchange(ctx context.Context, client *uniswapv1.Client) (*uniswapv1.ExchangeClient, error) {
	registry, err := token2.RegistryFromFile(client.Client, *clientUniswap1TokenFile)
	if err != nil {
		return nil, err
	}
	tok, err := registry.ByName(*clientUniswap1Token)
	if err != nil {
		return nil, err
	}

	return client.GetExchange(ctx, tok)
}

func uniswapGetExchange(ctx context.Context, client *uniswapv1.Client) error {
	ex, err := getExchange(ctx, client)
	if err != nil {
		return err
	}

	fmt.Printf("Exchange: %s\n", ex.ContractAddress().String())

	return nil
}

func uniswapGetEthToTokenInputPrice(ctx context.Context, client *uniswapv1.Client) error {
	ex, err := getExchange(ctx, client)
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

func uniswapGetEthToTokenOutputPrice(ctx context.Context, client *uniswapv1.Client) error {
	ex, err := getExchange(ctx, client)
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

func uniswapGetTokenToEthInputPrice(ctx context.Context, client *uniswapv1.Client) error {
	ex, err := getExchange(ctx, client)
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

func uniswapGetTokenToEthOutputPrice(ctx context.Context, client *uniswapv1.Client) error {
	ex, err := getExchange(ctx, client)
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
