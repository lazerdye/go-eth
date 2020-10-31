package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/opyn"
	"github.com/lazerdye/go-eth/token2"
)

var (
	clientOpynCommand                    = clientToken2Command.Command("opyn", "Opyn operations")
	clientOpynListOptionsContracts       = clientOpynCommand.Command("list-options-contracts", "List options contracts")
	clientOpynListOptionsContractsActive = clientOpynListOptionsContracts.Flag("active", "Only show active contracts").Bool()
)

func opynCommands(ctx context.Context, client *client.Client, reg *token2.Registry, commands []string) (bool, error) {
	opynClient, err := opyn.NewClient(client)
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "list-options-contracts":
		return true, opynListOptionsContracts(ctx, opynClient, reg)
	}
	return false, nil
}

func opynListOptionsContracts(ctx context.Context, opynClient *opyn.Client, reg *token2.Registry) error {
	numberOfOptionsContracts, err := opynClient.GetNumberOfOptionsContracts(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Number of options contracts: %v\n", numberOfOptionsContracts)
	for i := int64(0); i < numberOfOptionsContracts.Int64(); i++ {
		optionsContractAddress, err := opynClient.OptionsContract(ctx, big.NewInt(i))
		if err != nil {
			return err
		}
		otoken, err := opynClient.GetOToken(ctx, optionsContractAddress)
		if err != nil {
			return err
		}
		hasExpired, err := otoken.HasExpired(ctx)
		if err != nil {
			return err
		}
		if *clientOpynListOptionsContractsActive && hasExpired {
			continue
		}
		name, err := otoken.Name(ctx)
		if err != nil {
			return err
		}
		decimals, err := otoken.Decimals(ctx)
		if err != nil {
			return err
		}
		expiredStr := ""
		if hasExpired {
			expiredStr = " expired"
		}
		fmt.Printf("%d: %s (%v %d)%s\n", i, name, optionsContractAddress.Hex(), decimals, expiredStr)
		underlying, err := otoken.Underlying(ctx)
		if err != nil {
			return err
		}
		underlyingExp, err := otoken.UnderlyingExp(ctx)
		if err != nil {
			return err
		}
		underlyingName := underlying.Hex()
		if underlying == opyn.EthContract {
			underlyingName = "eth"
		} else {
			name, _, err := reg.ByAddress(underlying)
			if err == nil {
				underlyingName = name
			}
		}
		collateral, err := otoken.Collateral(ctx)
		if err != nil {
			return err
		}
		collateralExp, err := otoken.CollateralExp(ctx)
		if err != nil {
			return err
		}
		collateralName := collateral.Hex()
		if collateral == opyn.EthContract {
			collateralName = "eth"
		} else {
			name, _, err := reg.ByAddress(collateral)
			if err == nil {
				collateralName = name
			}
		}
		strikePrice, err := otoken.StrikePrice(ctx)
		if err != nil {
			return err
		}

		fmt.Printf("    %s %d -> %s %d (%s)\n",
			underlyingName, underlyingExp, collateralName, collateralExp, strikePrice.Shift(int32(decimals)))
	}
	return nil
}
