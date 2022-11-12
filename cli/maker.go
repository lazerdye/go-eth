package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/maker"
)

var (
	clientMakerCommand        = clientCmd.Command("maker", "Maker Commands")
	clientMakerOpenVault      = clientMakerCommand.Command("open-vault", "Open vault")
	clientMakerOpenVaultIlk   = clientMakerOpenVault.Arg("ilk", "ILK to open").Required().String()
	clientMakerLastVault      = clientMakerCommand.Command("last-vault", "Get last vault id")
	clientMakerUrnsVault      = clientMakerCommand.Command("urns-vault", "Get URNs for vault")
	clientMakerUrnsVaultCdpId = clientMakerUrnsVault.Arg("cdp-id", "CDP ID to get URNs for").Required().String()
	clientMakerFrobVault      = clientMakerCommand.Command("frob-vault", "'Frob' the collateral in the vault")
	clientMakerFrobVaultCdpId = clientMakerFrobVault.Arg("cdp-id", "CDP ID to move").Required().String()
	clientMakerFrobVaultDink  = clientMakerFrobVault.Arg("dink", "Dink collateral to add").Required().String()
	clientMakerFrobVaultDart  = clientMakerFrobVault.Arg("dart", "Dart vollateral to remove").Required().String()
	clientMakerGemJoin        = clientMakerCommand.Command("join", "Gem join")
	clientMakerGemJoinIlk     = clientMakerGemJoin.Arg("ilk", "Vault ilk").Required().String()
	clientMakerGemJoinUrn     = clientMakerGemJoin.Arg("urn", "Vault urn").Required().String()
	clientMakerGemJoinAmount  = clientMakerGemJoin.Arg("amount", "Join amount").Required().String()
	clientMakerGemExit        = clientMakerCommand.Command("exit", "Gem exit")
	clientMakerGemExitIlk     = clientMakerGemExit.Arg("ilk", "Vault ilk").Required().String()
	clientMakerGemExitUrn     = clientMakerGemExit.Arg("urn", "Vault urn").Required().String()
	clientMakerGemExitAmount  = clientMakerGemExit.Arg("amount", "Vault urn").Required().String()
)

func makerCommands(ctx context.Context, client *client.Client, commands []string) error {
	cdpClient, err := maker.NewCDPManagerClient(client)
	if err != nil {
		return err
	}
	switch commands[0] {
	case "open-vault":
		return doMakerOpen(ctx, cdpClient)
	case "last-vault":
		return doMakerLastVault(ctx, cdpClient)
	case "urns-vault":
		return doMakerUrnsVault(ctx, cdpClient)
	case "frob-vault":
		return doMakerFrobVault(ctx, cdpClient)
	case "join":
		return doMakerGemJoin(ctx, cdpClient)
	case "exit":
		return doMakerGemExit(ctx, cdpClient)
	default:
		return errors.Errorf("Unknown maker subcommand: %s", commands[0])
	}
}

func doMakerOpen(ctx context.Context, cdpClient *maker.CDPManagerClient) error {
	account, err := getAccount(true)
	if err != nil {
		return err
	}

	tx, err := cdpClient.Open(ctx, *clientMakerOpenVaultIlk, account)
	if err != nil {
		return err
	}
	fmt.Printf("Transaction: %s\n", tx.Hash())
	return nil
}

func doMakerLastVault(ctx context.Context, cdpClient *maker.CDPManagerClient) error {
	account, err := getAccount(false)
	if err != nil {
		return err
	}
	id, err := cdpClient.Last(ctx, account)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", id)
	return nil
}

func doMakerUrnsVault(ctx context.Context, cdpClient *maker.CDPManagerClient) error {
	cdpId, _ := new(big.Int).SetString(*clientMakerUrnsVaultCdpId, 10)

	address, err := cdpClient.Urns(ctx, cdpId)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", address)
	return nil
}

func doMakerFrobVault(ctx context.Context, cdpClient *maker.CDPManagerClient) error {
	account, err := getAccount(true)
	if err != nil {
		return err
	}
	cdpId, _ := new(big.Int).SetString(*clientMakerFrobVaultCdpId, 10)
	dink, _ := decimal.NewFromString(*clientMakerFrobVaultDink)
	dart, _ := decimal.NewFromString(*clientMakerFrobVaultDart)

	tx, err := cdpClient.Frob(ctx, account, cdpId, dink, dart)
	if err != nil {
		return err
	}
	fmt.Printf("TX: %s\n", tx.Hash())
	return nil
}

func doMakerGemJoin(ctx context.Context, cdpClient *maker.CDPManagerClient) error {
	account, err := getAccount(true)
	if err != nil {
		return err
	}
	gem, err := maker.NewGemJoinClient(cdpClient.Client, *clientMakerGemJoinIlk)
	if err != nil {
		return err
	}
	amount, _ := decimal.NewFromString(*clientMakerGemJoinAmount)
	usr := common.HexToAddress(*clientMakerGemJoinUrn)
	tx, err := gem.Join(ctx, account, usr, amount)
	if err != nil {
		return err
	}
	fmt.Printf("TX: %s\n", tx.Hash())
	return nil
}

func doMakerGemExit(ctx context.Context, cdpClient *maker.CDPManagerClient) error {
	account, err := getAccount(true)
	if err != nil {
		return err
	}
	gem, err := maker.NewGemJoinClient(cdpClient.Client, *clientMakerGemExitIlk)
	if err != nil {
		return err
	}
	amount, _ := decimal.NewFromString(*clientMakerGemExitAmount)
	usr := common.HexToAddress(*clientMakerGemExitUrn)
	tx, err := gem.Exit(ctx, account, usr, amount)
	if err != nil {
		return err
	}
	fmt.Printf("TX: %s\n", tx.Hash())
	return nil
}
