package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/lazerdye/go-eth/token"
	"github.com/lazerdye/go-eth/uniswapv1"
	"github.com/lazerdye/go-eth/wallet"
)

var (
	clientUniswapCommand                              = clientCmd.Command("uniswap", "Uniswap operations")
	clientUniswapToken                                = clientUniswapCommand.Flag("token", "Token to get exchange for").Required().String()
	clientUniswapGetExchange                          = clientUniswapCommand.Command("get-exchange", "Get exchange address for token")
	clientUniswapGetEthToTokenInputPrice              = clientUniswapCommand.Command("eth-to-token-input", "Get eth to token input price")
	clientUniswapGetEthToTokenInputPriceEthSold       = clientUniswapGetEthToTokenInputPrice.Arg("eth-sold", "Ethereum sold").Required().Float64()
	clientUniswapGetEthToTokenOutputPrice             = clientUniswapCommand.Command("eth-to-token-output", "Get eth to token output price")
	clientUniswapGetEthToTokenOutputPriceTokensBought = clientUniswapGetEthToTokenOutputPrice.Arg("tokens-bought", "Tokens bought").Required().Float64()
	clientUniswapGetTokenToEthInputPrice              = clientUniswapCommand.Command("token-to-eth-input", "Get token to eth input price")
	clientUniswapGetTokenToEthInputPriceEthBought     = clientUniswapGetTokenToEthInputPrice.Arg("eth-bought", "Ethereum bought").Required().Float64()
	clientUniswapGetTokenToEthOutputPrice             = clientUniswapCommand.Command("token-to-eth-output", "Get token to eth output price")
	clientUniswapGetTokenToEthOutputPriceTokensSold   = clientUniswapGetTokenToEthOutputPrice.Arg("tokens-sold", "Tokens sold").Required().Float64()
	clientUniswapApprove                              = clientUniswapCommand.Command("approve", "Approve exchange")
	clientUniswapApproveAmount                        = clientUniswapApprove.Arg("amount", "Amount to approve").Required().Float64()
)

func getExchange(ctx context.Context, client *uniswapv1.Client) (*uniswapv1.ExchangeClient, error) {
	tok, err := token.ByName(client.Client, *clientUniswapToken)
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

	ethSold := big.NewFloat(*clientUniswapGetEthToTokenInputPriceEthSold)
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

	tokensBought := big.NewFloat(*clientUniswapGetEthToTokenOutputPriceTokensBought)
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

	ethBought := big.NewFloat(*clientUniswapGetTokenToEthInputPriceEthBought)
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

	tokensSold := big.NewFloat(*clientUniswapGetTokenToEthOutputPriceTokensSold)
	ethBought, err := ex.GetTokenToEthOutputPrice(ctx, tokensSold)
	if err != nil {
		return err
	}

	fmt.Printf("Token sold: %.18f\n", ethBought)

	return nil
}

func uniswapApprove(ctx context.Context, client *uniswapv1.Client, account *wallet.Account) error {
	ex, err := getExchange(ctx, client)
	if err != nil {
		return err
	}

	amountToApprove := big.NewFloat(*clientUniswapApproveAmount)
	tx, err := ex.Approve(ctx, account, amountToApprove)
	if err != nil {
		return err
	}

	fmt.Printf("Transaction: %+v\n", tx)

	return nil
}
