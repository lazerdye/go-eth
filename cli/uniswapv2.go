package main

import (
	"context"
	"fmt"
	"math/big"

	log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/token"
	"github.com/lazerdye/go-eth/uniswapv2"
)

var (
	clientUniswapv2Command            = clientCmd.Command("uniswapv2", "Uniswap V2 operations")
	clientUniswapv2GetPairs           = clientUniswapv2Command.Command("get-pairs", "Get Pairs")
	clientUniswapv2GetAmountOut       = clientUniswapv2Command.Command("get-amount-out", "Get output amount")
	clientUniswapv2GetAmountOutToken0 = clientUniswapv2GetAmountOut.Arg("token0", "Token 0").String()
	clientUniswapv2GetAmountOutToken1 = clientUniswapv2GetAmountOut.Arg("token1", "Token 1").String()
	clientUniswapv2GetAmountOutAmount = clientUniswapv2GetAmountOut.Arg("amount", "Token 1").Float64()
	clientUniswapv2GetAmountIn        = clientUniswapv2Command.Command("get-amount-in", "Get input amount")
	clientUniswapv2GetAmountInToken0  = clientUniswapv2GetAmountIn.Arg("token0", "Token 0").String()
	clientUniswapv2GetAmountInToken1  = clientUniswapv2GetAmountIn.Arg("token1", "Token 1").String()
	clientUniswapv2GetAmountInAmount  = clientUniswapv2GetAmountIn.Arg("amount", "Token 0").Float64()
)

func uniswapV2GetPairs(ctx context.Context, client *uniswapv2.Client) error {
	pairs, err := client.GetPairs(ctx)
	if err != nil {
		return err
	}
	for _, pair := range pairs {
		token0, err := pair.Token0(ctx)
		if err != nil {
			return err
		}
		token0Client, err := token.ByRawAddress(ctx, client.Client, token0)
		if err != nil {
			log.Warnf("Token %s could not be initialized: %s", token0.String(), err)
			continue
		}
		token1, err := pair.Token1(ctx)
		if err != nil {
			return err
		}
		token1Client, err := token.ByRawAddress(ctx, client.Client, token1)
		if err != nil {
			log.Warnf("Token %s could not be initialized: %s", token0.String(), err)
			continue
		}
		fmt.Printf("%+v %+v -> %+v\n", pair.Address.String(), token0Client.Name(), token1Client.Name())
	}
	return nil
}

func uniswapV2GetAmountOut(ctx context.Context, client *uniswapv2.Client) error {
	token0, err := token.ByName(client.Client, *clientUniswapv2GetAmountOutToken0)
	if err != nil {
		return err
	}
	token1, err := token.ByName(client.Client, *clientUniswapv2GetAmountOutToken1)
	if err != nil {
		return err
	}
	amountBig := big.NewFloat(*clientUniswapv2GetAmountOutAmount)
	value, err := client.GetAmountOut(ctx, amountBig, token0, token1)
	if err != nil {
		return err
	}
	fmt.Printf("AmountOut: %s\n", value.String())
	return nil
}

func uniswapV2GetAmountIn(ctx context.Context, client *uniswapv2.Client) error {
	token0, err := token.ByName(client.Client, *clientUniswapv2GetAmountInToken0)
	if err != nil {
		return err
	}
	token1, err := token.ByName(client.Client, *clientUniswapv2GetAmountInToken1)
	if err != nil {
		return err
	}
	amountBig := big.NewFloat(*clientUniswapv2GetAmountInAmount)
	value, err := client.GetAmountIn(ctx, amountBig, token0, token1)
	if err != nil {
		return err
	}
	fmt.Printf("AmountIn: %s\n", value.String())
	return nil
}
