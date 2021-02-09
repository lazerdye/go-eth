package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasoracle"
	"github.com/lazerdye/go-eth/gasstation"
)

var (
	clientCmd     = kingpin.Command("client", "Client operations")
	clientBlockNo = clientCmd.Flag("block-number", "Override block number").Int64()
	clientServer  = clientCmd.Flag("server", "URL of the server to connect to").Envar("ETHEREUM_SERVER").Required().String()

	statusCmd             = clientCmd.Command("status", "Get status")
	balanceCmd            = clientCmd.Command("balance", "Get the balance of an account")
	getTransactionCmd     = clientCmd.Command("get-transaction", "Get info about the given transaction hash")
	getTransactionTransId = getTransactionCmd.Arg("transaction-id", "Transaction ID").Required().String()
)

func newClient() (*client.Client, error) {
	if *clientServer == "" {
		return nil, errors.New("clientServer parameter required")
	}
	o := gasoracle.GasOracleFromGasStation(gasstation.NewClient())
	c, err := client.Dial(*clientServer, o)
	if err != nil {
		return nil, err
	}
	if *clientBlockNo != 0 {
		c.SetBlockNumberOverride(big.NewInt(*clientBlockNo))
	}
	return c, err
}

func clientCommands(ctx context.Context, commands []string) (bool, error) {
	client, err := newClient()
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "status":
		return true, doClientStatus(ctx, client)
	case "balance":
		return true, doClientBalance(ctx, client)
	case "get-transaction":
		return true, doClientGetTransaction(ctx, client)
	case "token2":
		return token2Commands(ctx, client, commands[1:])
	case "augur":
		return augurCommands(ctx, client, commands[1:])
	}
	return false, nil
}

func doClientStatus(ctx context.Context, c *client.Client) error {
	stat, err := c.SyncProgress(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", stat)
	return nil
}

func doClientBalance(ctx context.Context, c *client.Client) error {
	if *address == "" {
		return errors.New("--address parameter required")
	}
	queryAddress := common.HexToAddress(*address)
	eth, err := c.BalanceOf(ctx, queryAddress)
	if err != nil {
		return err
	}
	fmt.Printf("Balance of %s: %s\n", queryAddress.String(), eth)
	return nil
}

func doClientGetTransaction(ctx context.Context, c *client.Client) error {
	hash := common.HexToHash(*getTransactionTransId)
	tx, isPending, err := c.TransactionByHash(ctx, hash)
	if err != nil {
		return err
	}
	if isPending {
		fmt.Printf("PENDING Transaction: %+v\n", tx)
	} else {
		fmt.Printf("Transaction: %+v\n", tx)
	}
	return nil
}
