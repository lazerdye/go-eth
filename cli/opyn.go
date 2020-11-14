package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/opyn"
	"github.com/lazerdye/go-eth/token2"
)

var (
	clientOpynCommand                    = clientToken2Command.Command("opyn", "Opyn operations")
	clientOpynOTokenAddress              = clientToken2Command.Flag("otoken", "OToken Address").String()
	clientOpynListOptionsContracts       = clientOpynCommand.Command("list-options-contracts", "List options contracts")
	clientOpynListOptionsContractsActive = clientOpynListOptionsContracts.Flag("active", "Only show active contracts").Bool()
	clientOpynOpenVaultCommand           = clientOpynCommand.Command("open-vault", "Open a vault")
)

func opynCommands(ctx context.Context, client *client.Client, reg *token2.Registry, commands []string) (bool, error) {
	opynClient, err := opyn.NewClient(client)
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "list-options-contracts":
		return true, opynListOptionsContracts(ctx, opynClient, reg)
	case "open-vault":
		return true, opynOpenVault(ctx, opynClient)
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
		name, err := otoken.Name(ctx)
		if err != nil {
			return err
		}
		if *clientOpynListOptionsContractsActive && (hasExpired || name == "") {
			continue
		}
		info, err := otoken.ContractInfo(ctx)
		if err != nil {
			return err
		}
		currencyName := info.Currency.Hex()
		if info.Currency == opyn.EthContract {
			currencyName = "eth"
		} else {
			name, _, err := reg.ByAddress(info.Currency)
			if err == nil {
				currencyName = name
			}
		}
		fmt.Printf("%d %s: %s %s %s (%s)\n", i, info.Type, info.Name, info.StrikePrice, info.Expiry, currencyName)
	}
	return nil
}

func opynOpenVault(ctx context.Context, opynClient *opyn.Client) error {
	if *clientOpynOTokenAddress == "" {
		return errors.New("Flag --otoken required")
	}
	account, unlocked, err := getAccount()
	if err != nil {
		return err
	}
	if !unlocked {
		return errors.New("Wallet is locked")
	}
	otoken, err := opynClient.GetOToken(ctx, common.HexToAddress(*clientOpynOTokenAddress))
	if err != nil {
		return err
	}
	tx, err := otoken.OpenVault(ctx, account)
	if err != nil {
		return err
	}
	fmt.Printf("Transaction: %s\n", tx.Hash())
	return nil
}
