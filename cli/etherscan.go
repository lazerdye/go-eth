package main

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/lazerdye/go-eth/etherscan"
)

var (
	etherscanApikey   = kingpin.Flag("etherscan-apikey", "Key for etherscan api").Envar("ETHERSCAN_APIKEY").String()
	etherscanCmd      = kingpin.Command("etherscan", "Etherscan operations")
	etherscanList     = etherscanCmd.Command("list", "List transactions for account")
	etherscanToken    = etherscanCmd.Flag("token", "List for given token").String()
	etherscanDecimals = etherscanCmd.Flag("decimals", "Adjust value by given decimal value").Int()
	etherscanSort     = etherscanList.Arg("sort", "Sort").Default("asc").String()
	etherscanPage     = etherscanList.Arg("page", "Page number").Default("1").Int()
	etherscanOffset   = etherscanList.Arg("offset", "Page offset").Default("20").Int()
)

func doEtherscanList(ctx context.Context, address string) error {
	if *etherscanApikey == "" {
		return errors.New("--etherscan-apikey required")
	}
	c := etherscan.NewClient(*etherscanApikey)
	var transactions []etherscan.Transaction
	var err error
	if *etherscanToken == "" {
		transactions, err = c.NormalTransactionsByAddress(ctx, address, *etherscanPage, *etherscanOffset, *etherscanSort)
		if err != nil {
			return err
		}
	} else {
		transactions, err = c.TokenTransactionsByAddress(ctx, address, *etherscanToken, *etherscanPage, *etherscanOffset, *etherscanSort)
		if err != nil {
			return err
		}
	}

	for _, t := range transactions {
		value := t.Value
		if *etherscanDecimals != 0 {
			iValue, ok := new(big.Int).SetString(value, 10)
			if !ok {
				return errors.Errorf("Invalid value: %s", value)
			}
			fValue := new(big.Float).Quo(new(big.Float).SetInt(iValue), big.NewFloat(math.Pow10(*etherscanDecimals)))
			value = fmt.Sprintf(fmt.Sprintf("%%.%df", *etherscanDecimals), fValue)
		}
		fmt.Printf("%s\t%s -> %s\n", t.Hash, t.From, t.To)
		fmt.Printf("\t%s (%s)\t%s\n", t.BlockNumber, t.Timestamp, value)
	}

	return nil
}
