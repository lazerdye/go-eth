package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/uniswapv2"
)

var (
	clientUniswapv2Command            = clientToken2Command.Command("uniswapv2", "Uniswap V2 operations")
	clientUniswapv2GetPairs           = clientUniswapv2Command.Command("get-pairs", "Get Pairs")
	clientUniswapv2GetPair            = clientUniswapv2Command.Command("get-pair", "Get Pairs")
	clientUniswapv2GetPairToken0      = clientUniswapv2GetPair.Arg("token0", "Token 0").String()
	clientUniswapv2GetPairToken1      = clientUniswapv2GetPair.Arg("token1", "Token 1").String()
	clientUniswapv2GetAmountOut       = clientUniswapv2Command.Command("get-amount-out", "Get output amount")
	clientUniswapv2GetAmountOutToken0 = clientUniswapv2GetAmountOut.Arg("token0", "Token 0").String()
	clientUniswapv2GetAmountOutToken1 = clientUniswapv2GetAmountOut.Arg("token1", "Token 1").String()
	clientUniswapv2GetAmountOutAmount = clientUniswapv2GetAmountOut.Arg("amount", "Token 1").Float64()
	clientUniswapv2GetAmountIn        = clientUniswapv2Command.Command("get-amount-in", "Get input amount")
	clientUniswapv2GetAmountInToken0  = clientUniswapv2GetAmountIn.Arg("token0", "Token 0").String()
	clientUniswapv2GetAmountInToken1  = clientUniswapv2GetAmountIn.Arg("token1", "Token 1").String()
	clientUniswapv2GetAmountInAmount  = clientUniswapv2GetAmountIn.Arg("amount", "Token 0").Float64()
	clientUniswapv2GetReserves        = clientUniswapv2Command.Command("get-reserves", "Get reserves")
	clientUniswapv2GetReservesToken0  = clientUniswapv2GetReserves.Arg("token0", "Token 0").String()
	clientUniswapv2GetReservesToken1  = clientUniswapv2GetReserves.Arg("token1", "Token 1").String()
	clientUniswapv2Claim              = clientUniswapv2Command.Command("claim-token", "Claim UNI token")
	clientUniswapv2ClaimIndex         = clientUniswapv2Claim.Arg("index", "Token claim index").Int64()
	clientUniswapv2ClaimAmount        = clientUniswapv2Claim.Arg("amount", "Token claim amount").String()
	clientUniswapv2ClaimMerkleProof   = clientUniswapv2Claim.Arg("merkle-proof", "Token claim merkle proof - comma separated hex").String()
)

func uniswapV2Commands(ctx context.Context, client *client.Client, reg *token2.Registry, commands []string) (bool, error) {
	uniswapv2Client, err := uniswapv2.NewClient(client)
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "get-pairs":
		return true, uniswapV2GetPairs(ctx, uniswapv2Client)
	case "get-pair":
		return true, uniswapV2GetPair(ctx, reg, uniswapv2Client)
	case "get-amount-out":
		return true, uniswapV2GetAmountOut(ctx, reg, uniswapv2Client)
	case "get-amount-in":
		return true, uniswapV2GetAmountIn(ctx, reg, uniswapv2Client)
	case "get-reserves":
		return true, uniswapV2GetReserves(ctx, reg, uniswapv2Client)
	case "claim-token":
		return true, uniswapV2Claim(ctx, reg, uniswapv2Client)
	}
	return false, nil
}

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
		token0Client, err := token2.ByAddress(ctx, client.Client, token0)
		if err != nil {
			log.Warnf("Token %s could not be initialized: %s", token0.String(), err)
			continue
		}
		token1, err := pair.Token1(ctx)
		if err != nil {
			return err
		}
		token1Client, err := token2.ByAddress(ctx, client.Client, token1)
		if err != nil {
			log.Warnf("Token %s could not be initialized: %s", token0.String(), err)
			continue
		}
		fmt.Printf("%+v %+v -> %+v\n", pair.Address.String(), token0Client.Address.String(), token1Client.Address.String())
	}
	return nil
}

func uniswapV2GetPair(ctx context.Context, registry *token2.Registry, client *uniswapv2.Client) error {
	token0, err := registry.ByName(*clientUniswapv2GetPairToken0)
	if err != nil {
		return err
	}
	token1, err := registry.ByName(*clientUniswapv2GetPairToken1)
	if err != nil {
		return err
	}
	pair, err := client.GetPair(ctx, token0, token1)
	if err != nil {
		return err
	}
	fmt.Printf("Pair address: %s\n", pair.Address.String())
	return nil
}

func uniswapV2GetAmountOut(ctx context.Context, registry *token2.Registry, client *uniswapv2.Client) error {
	token0, err := registry.ByName(*clientUniswapv2GetAmountOutToken0)
	if err != nil {
		return err
	}
	token1, err := registry.ByName(*clientUniswapv2GetAmountOutToken1)
	if err != nil {
		return err
	}
	amountBig := decimal.NewFromFloat(*clientUniswapv2GetAmountOutAmount)
	value, err := client.GetAmountOut(ctx, amountBig, token0, token1)
	if err != nil {
		return err
	}
	fmt.Printf("AmountOut: %s\n", value.String())
	return nil
}

func uniswapV2GetAmountIn(ctx context.Context, registry *token2.Registry, client *uniswapv2.Client) error {
	token0, err := registry.ByName(*clientUniswapv2GetAmountInToken0)
	if err != nil {
		return err
	}
	token1, err := registry.ByName(*clientUniswapv2GetAmountInToken1)
	if err != nil {
		return err
	}
	amountBig := decimal.NewFromFloat(*clientUniswapv2GetAmountInAmount)
	value, err := client.GetAmountIn(ctx, amountBig, token0, token1)
	if err != nil {
		return err
	}
	fmt.Printf("AmountIn: %s\n", value.String())
	return nil
}

func uniswapV2GetReserves(ctx context.Context, registry *token2.Registry, client *uniswapv2.Client) error {
	token0, err := registry.ByName(*clientUniswapv2GetReservesToken0)
	if err != nil {
		return err
	}
	token1, err := registry.ByName(*clientUniswapv2GetReservesToken1)
	if err != nil {
		return err
	}
	pair, err := client.GetPair(ctx, token0, token1)
	if err != nil {
		return err
	}
	reserves, err := pair.GetReserves(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Reserves: %s\n", reserves)
	return nil
}

func uniswapV2Claim(ctx context.Context, registry *token2.Registry, client *uniswapv2.Client) error {
	uniToken, err := registry.ByName("uni")
	if err != nil {
		return err
	}
	account, unlocked, err := getAccount()
	if err != nil {
		return err
	}
	if !unlocked {
		return errors.New("Wallet locked")
	}
	proofs := strings.Split(*clientUniswapv2ClaimMerkleProof, ",")
	proofsByte := make([][32]byte, len(proofs))
	for i, proof := range proofs {
		proofsByte[i] = common.HexToHash(proof)
	}
	amount, err := decimal.NewFromString(*clientUniswapv2ClaimAmount)
	if err != nil {
		return err
	}
	tx, err := client.ClaimToken(ctx, account, uniToken, *clientUniswapv2ClaimIndex, amount, proofsByte)
	if err != nil {
		return err
	}
	fmt.Printf("Result: %s\n", tx.Hash().String())
	return nil
}
