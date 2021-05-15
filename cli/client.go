package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

	statusCmd                   = clientCmd.Command("status", "Get status")
	balanceCmd                  = clientCmd.Command("balance", "Get the balance of an account")
	balanceAtCmd                = clientCmd.Command("balance-at", "Get the balance of an account")
	balanceAtBlock              = balanceAtCmd.Arg("block-number", "Block Number").Required().Int64()
	getTransactionCmd           = clientCmd.Command("get-transaction", "Get info about the given transaction hash")
	getTransactionTransId       = getTransactionCmd.Arg("transaction-id", "Transaction ID").Required().String()
	resubmitTransactionCmd      = clientCmd.Command("resubmit-transaction", "Resubmit a transaction")
	resubmitTransactionTransId  = resubmitTransactionCmd.Arg("transaction-id", "Transaction ID").Required().String()
	resubmitTransactionGasLimit = resubmitTransactionCmd.Flag("gas-limit", "Use this gas limit instead of increasing by 10%").Int64()
	cancelTransactionCmd        = clientCmd.Command("cancel-transaction", "Cancel a transaction")
	cancelTransactionTransId    = cancelTransactionCmd.Arg("transaction-id", "Transaction ID").Required().String()
	submitTransactionCmd        = clientCmd.Command("submit-transaction", "Submit a transaction from json")
	submitTransactionTrans      = submitTransactionCmd.Arg("transaction", "transaction json").Required().String()
	filterLogsCmd               = clientCmd.Command("filter-logs", "Filter Logs")
	filterLogsQuery             = filterLogsCmd.Arg("query", "Query").String() // a|b,c|d
	filterLogsFromBlockNumber   = filterLogsCmd.Arg("from", "From Block Number").Int64()
	filterLogsToBlockNumber     = filterLogsCmd.Arg("to", "From Block Number").Int64()
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
	case "balance-at":
		return true, doClientBalanceAt(ctx, client)
	case "get-transaction":
		return true, doClientGetTransaction(ctx, client)
	case "submit-transaction":
		return true, doClientSubmitTransaction(ctx, client)
	case "resubmit-transaction":
		return true, doClientResubmitTransaction(ctx, client)
	case "cancel-transaction":
		return true, doClientCancelTransaction(ctx, client)
	case "filter-logs":
		return true, doClientFilterLogs(ctx, client)
	case "token2":
		return token2Commands(ctx, client, commands[1:])
	case "augur":
		return augurCommands(ctx, client, commands[1:])
	case "digixdao":
		return digixDaoCommands(ctx, client, commands[1:])
	}
	return false, nil
}

func doClientStatus(ctx context.Context, c *client.Client) error {
	stat, err := c.SyncProgress(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", stat)
	if stat != nil {
		return errors.New("Sync in progress")
	}
	return nil
}

func doClientBalance(ctx context.Context, c *client.Client) error {
	if *address == "" {
		return errors.New("--address parameter required")
	}
	queryAddress := common.HexToAddress(*address)
	eth, err := c.BalanceAt(ctx, queryAddress, nil)
	if err != nil {
		return err
	}
	fmt.Printf("Balance of %s: %s\n", queryAddress.String(), eth)
	return nil
}

func doClientBalanceAt(ctx context.Context, c *client.Client) error {
	if *address == "" {
		return errors.New("--address parameter required")
	}
	queryAddress := common.HexToAddress(*address)
	eth, err := c.BalanceAt(ctx, queryAddress, big.NewInt(*balanceAtBlock))
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
	json, err := tx.MarshalJSON()
	if err != nil {
		return err
	}
	if isPending {
		fmt.Printf("PENDING Transaction: %s\n", string(json))
	} else {
		fmt.Printf("Transaction: %s\n", string(json))
	}
	return nil
}

func doClientSubmitTransaction(ctx context.Context, c *client.Client) error {
	account, unlocked, err := getAccount()
	if err != nil {
		return err
	}
	if !unlocked {
		return errors.New("Wallet locked")
	}
	txValue := *submitTransactionTrans
	var tx types.Transaction
	if err := tx.UnmarshalJSON([]byte(txValue)); err != nil {
		return err
	}
	fmt.Printf("%+v\n", tx)
	chainID, err := c.Client.NetworkID(ctx)
	if err != nil {
		return err
	}
	txSigned, err := account.SignTx(&tx, chainID)
	if err != nil {
		return err
	}

	return c.SendTransaction(ctx, txSigned)
}

func doClientResubmitTransaction(ctx context.Context, c *client.Client) error {
	account, unlocked, err := getAccount()
	if err != nil {
		return err
	}
	if !unlocked {
		return errors.New("Wallet locked")
	}

	hash := common.HexToHash(*resubmitTransactionTransId)
	tx, _, err := c.TransactionByHash(ctx, hash)
	if err != nil {
		return err
	}
	pAddress := tx.To()
	var price *big.Int
	if *resubmitTransactionGasLimit == 0 {
		incPrice := new(big.Int).Div(tx.GasPrice(), big.NewInt(10))
		price = new(big.Int).Add(tx.GasPrice(), incPrice)
		fmt.Printf("Increasing gas price from %d to %d\n", tx.GasPrice(), price)
	} else {
		price = new(big.Int).Mul(big.NewInt(*resubmitTransactionGasLimit), big.NewInt(1e9))
		fmt.Printf("Using gas price %d\n", price)
	}
	newTx := types.NewTransaction(tx.Nonce(), *pAddress, tx.Value(), tx.Gas(), price, tx.Data())
	chainID, err := c.Client.NetworkID(ctx)
	if err != nil {
		return err
	}
	txSigned, err := account.SignTx(newTx, chainID)
	if err != nil {
		return err
	}
	json, err := txSigned.MarshalJSON()
	if err != nil {
		return err
	}

	fmt.Printf("Transaction %s: %s\n", txSigned.Hash().String(), string(json))

	return c.SendTransaction(ctx, txSigned)
}

func doClientCancelTransaction(ctx context.Context, c *client.Client) error {
	account, unlocked, err := getAccount()
	if err != nil {
		return err
	}
	if !unlocked {
		return errors.New("Wallet locked")
	}

	hash := common.HexToHash(*cancelTransactionTransId)
	tx, _, err := c.TransactionByHash(ctx, hash)
	if err != nil {
		return err
	}
	fmt.Printf("Clearing valid and data for transaction\n")
	newTx := types.NewTransaction(tx.Nonce(), account.Address(), nil, tx.Gas(), tx.GasPrice(), nil)
	chainID, err := c.Client.NetworkID(ctx)
	if err != nil {
		return err
	}
	txSigned, err := account.SignTx(newTx, chainID)
	if err != nil {
		return err
	}
	json, err := txSigned.MarshalJSON()
	if err != nil {
		return err
	}

	fmt.Printf("Transaction %s: %s\n", txSigned.Hash().String(), string(json))

	return c.SendTransaction(ctx, txSigned)
}

func doClientFilterLogs(ctx context.Context, c *client.Client) error {
	items := strings.Split(*filterLogsQuery, ",")
	fmt.Printf("%+v\n", items)
	hashes := make([][]common.Hash, len(items))
	for i, item := range items {
		subItems := strings.Split(item, "|")
		subHashes := make([]common.Hash, len(subItems))
		for j, hash := range subItems {
			subHashes[j] = common.HexToHash(hash)
		}
		hashes[i] = subHashes
	}
	fmt.Printf("%+v\n", hashes)

	var fromBlockNumber *big.Int
	if *filterLogsFromBlockNumber != 0 {
		fromBlockNumber = big.NewInt(*filterLogsFromBlockNumber)
	}
	var toBlockNumber *big.Int
	if *filterLogsToBlockNumber != 0 {
		toBlockNumber = big.NewInt(*filterLogsToBlockNumber)
	}

	if err := c.FilterTransferLogs(ctx, fromBlockNumber, toBlockNumber); err != nil {
		return err
	}
	return nil
}
