package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/lazerdye/go-eth/token"
	"github.com/lazerdye/go-eth/uniswapv1"
)

var (
	clientUniswap1Command                              = clientCmd.Command("uniswapv1", "Uniswap v1 operations")
	clientUniswap1Token                                = clientUniswap1Command.Flag("token", "Token to get exchange for").Required().String()
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
	tok, err := token.ByName(client.Client, *clientUniswap1Token)
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

	ethSold := big.NewFloat(*clientUniswap1GetEthToTokenInputPriceEthSold)
	tokenBought, err := ex.GetEthToTokenInputPrice(ctx, ethSold)
	if err != nil {
		return err
	}

	fmt.Printf("Token bought: %.18f\n", tokenBought)

	return nil
}

func uniswapGetEthToTokenOutputPrice(ctx context.Context, client *uniswapv1.Client) error {
	ex, err := getExchange(ctx, client)
	if err != nil {
		return err
	}

	tokensBought := big.NewFloat(*clientUniswap1GetEthToTokenOutputPriceTokensBought)
	ethSold, err := ex.GetEthToTokenOutputPrice(ctx, tokensBought)
	if err != nil {
		return err
	}

	fmt.Printf("Eth sold: %.18f\n", ethSold)

	return nil
}

func uniswapGetTokenToEthInputPrice(ctx context.Context, client *uniswapv1.Client) error {
	ex, err := getExchange(ctx, client)
	if err != nil {
		return err
	}

	ethBought := big.NewFloat(*clientUniswap1GetTokenToEthInputPriceEthBought)
	tokenSold, err := ex.GetTokenToEthInputPrice(ctx, ethBought)
	if err != nil {
		return err
	}

	fmt.Printf("Eth bought: %.18f\n", tokenSold)

	return nil
}

func uniswapGetTokenToEthOutputPrice(ctx context.Context, client *uniswapv1.Client) error {
	ex, err := getExchange(ctx, client)
	if err != nil {
		return err
	}

	tokensSold := big.NewFloat(*clientUniswap1GetTokenToEthOutputPriceTokensSold)
	ethBought, err := ex.GetTokenToEthOutputPrice(ctx, tokensSold)
	if err != nil {
		return err
	}

	fmt.Printf("Token sold: %.18f\n", ethBought)

	return nil
}