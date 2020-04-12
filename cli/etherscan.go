package main

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"

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
		feeInt := etherscan.CalculateFee(&t)
		value := t.Value
		feeValue := feeInt.String()
		if *etherscanDecimals != 0 {
			iValue, ok := new(big.Int).SetString(value, 10)
			if !ok {
				return errors.Errorf("Invalid value: %s", value)
			}
			fValue := new(big.Float).Quo(new(big.Float).SetInt(iValue), big.NewFloat(math.Pow10(*etherscanDecimals)))
			value = fmt.Sprintf(fmt.Sprintf("%%.%df", *etherscanDecimals), fValue)

			feeFloat := new(big.Float).Quo(new(big.Float).SetInt(feeInt), big.NewFloat(math.Pow10(18)))
			feeValue = fmt.Sprintf("%.18f", feeFloat)
		}
		if strings.EqualFold(address, t.From) {
			fmt.Printf("%s\tOUT -> %s\n", t.Hash, t.To)
		} else if strings.EqualFold(address, t.To) {
			fmt.Printf("%s\t IN <- %s\n", t.Hash, t.From)
		} else {
			return errors.New("Cannot find from/to")
		}
		fmt.Printf("\t%s (%s)\t%s - %s", t.BlockNumber, t.Timestamp, value, feeValue)
		if t.IsError != "0" && t.IsError != "" {
			fmt.Printf("\tFAILED")
		}
		fmt.Println()
	}

	return nil
}
