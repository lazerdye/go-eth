package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasoracle"
	"github.com/lazerdye/go-eth/gasstation"
	"github.com/lazerdye/go-eth/kyber"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/wallet"
	"github.com/lazerdye/go-eth/zeroex"
)

var (
	keystore       = kingpin.Flag("keystore", "Location of the wallet keystore").Envar("WALLET_KEYSTORE").String()
	passphrase     = kingpin.Flag("passphrase", "Passphrase for keystore").Envar("WALLET_PASSPHRASE").String()
	address        = kingpin.Flag("address", "Address for account operations").Envar("ETHEREUM_ADDRESS").String()
	destAddress    = kingpin.Flag("dest-address", "Destination address for transfer").String()
	transferAmount = kingpin.Flag("amount", "Transfer amount").Float64()
	contract       = kingpin.Flag("contract", "Contract for approval/allowance").String()

	accountCmd       = kingpin.Command("account", "Account operations")
	listAccounts     = accountCmd.Command("list", "List the accounts")
	newAccount       = accountCmd.Command("new", "Create new account")
	unlockAccountCmd = accountCmd.Command("unlock", "Unlock an account")

	statusCmd                 = clientCmd.Command("status", "Get status")
	balanceCmd                = clientCmd.Command("balance", "Get the balance of an account")
	transferCmd               = clientCmd.Command("transfer", "Transfer ethereum")
	transferTransmit          = transferCmd.Flag("transmit", "Transmit transaction").Bool()
	filterLogsCmd             = clientCmd.Command("filterLogs", "Filter Logs")
	filterLogsFromBlockNumber = filterLogsCmd.Flag("from", "From Block Number").Int64()
	filterLogsToBlockNumber   = filterLogsCmd.Flag("to", "From Block Number").Int64()

	tokenKyberCmd          = clientCmd.Command("kyber", "Kyber Network")
	tokenKyberTokenFile    = tokenKyberCmd.Flag("token-file", "Token file").Required().String()
	tokenKyberSource       = tokenKyberCmd.Flag("source", "Source exchange").String()
	tokenKyberDest         = tokenKyberCmd.Flag("dest", "Dest exchange").String()
	tokenKyberQuantity     = tokenKyberCmd.Flag("quantity", "Source quantity").Float64()
	tokenKyberExpectedRate = tokenKyberCmd.Command("expectedRate", "Expected rate")
)

func doAccountList(keystore string) error {
	w, err := wallet.Open(keystore)
	if err != nil {
		return err
	}
	err = w.PrintAccounts()
	if err != nil {
		return err
	}
	return nil
}

func doAccountNew(keystore string, passphrase string) error {
	w, err := wallet.Open(keystore)
	if err != nil {
		return err
	}
	account, err := w.NewAccount(passphrase)
	if err != nil {
		return err
	}
	err = w.PrintAccount(account)
	if err != nil {
		return err
	}
	return nil
}

func doClientStatus(ctx context.Context, server string) error {
	o := gasoracle.GasOracleFromGasStation(gasstation.NewClient())
	c, err := client.Dial(server, o)
	if err != nil {
		return err
	}
	stat, err := c.SyncProgress(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", stat)
	return nil
}

func doClientBalance(server, addressStr string) error {
	address := common.HexToAddress(addressStr)
	o := gasoracle.GasOracleFromGasStation(gasstation.NewClient())
	c, err := client.Dial(server, o)
	if err != nil {
		return err
	}
	eth, err := c.BalanceOf(context.Background(), address)
	if err != nil {
		return err
	}
	fmt.Printf("Balance of %s: %s\n", address.String(), eth)

	return nil
}

func doClientTransfer(server string, account *wallet.Account, destAddress string, amount float64) error {
	ctx := context.Background()
	o := gasoracle.GasOracleFromGasStation(gasstation.NewClient())
	c, err := client.Dial(server, o)
	if err != nil {
		return err
	}
	tx, err := c.Transfer(ctx, account, common.HexToAddress(destAddress), decimal.NewFromFloat(amount), false)
	if err != nil {
		return err
	}
	fmt.Printf("Transaction(unsent): %+v\n", tx)
	return nil
}

func doClientFilterLogs(server string) error {
	ctx := context.Background()
	o := gasoracle.GasOracleFromGasStation(gasstation.NewClient())
	c, err := client.Dial(server, o)
	if err != nil {
		return err
	}
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

func doClientKyberExpectedRate(server string, source, dest string, quantity float64) error {
	o := gasoracle.GasOracleFromGasStation(gasstation.NewClient())
	c, err := client.Dial(server, o)
	if err != nil {
		return err
	}
	k, err := kyber.NewClient(c)
	if err != nil {
		return err
	}
	r, err := token2.RegistryFromFile(c, *tokenKyberTokenFile)
	if err != nil {
		return err
	}
	quantityFloat := decimal.NewFromFloat(quantity)
	sourceToken, err := r.ByName(source)
	if err != nil {
		return err
	}
	destToken, err := r.ByName(dest)
	if err != nil {
		return err
	}
	expectedRate, slippageRate, err := k.GetExpectedRate(context.Background(), sourceToken, destToken, quantityFloat)
	if err != nil {
		return err
	}
	fmt.Printf("Expected Rate: %s\n", expectedRate.String())
	fmt.Printf("Slippage Rate: %s\n", slippageRate.String())
	return nil
}

func newZeroexClient() (*zeroex.Client, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}
	return zeroex.NewClient(client)
}

func getAccount() (*wallet.Account, bool, error) {
	if *address == "" {
		return nil, false, errors.New("address parameter required")
	}
	if *keystore == "" {
		return nil, false, errors.New("keystore parameter required")
	}
	w, err := wallet.Open(*keystore)
	if err != nil {
		return nil, false, err
	}
	account, err := w.Account(*address)
	if err != nil {
		return nil, false, err
	}
	if *passphrase != "" {
		if err := account.Unlock(*passphrase); err != nil {
			return nil, false, err
		}
		return account, true, nil
	} else {
		return account, false, nil
	}
}

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("0.1").Author("Lazerdye")
	kingpin.CommandLine.Help = "Ethereum test client"

	commands := kingpin.Parse()
	commandsSplit := strings.Split(commands, " ")
	ctx := context.Background()
	switch commandsSplit[0] {
	case "client":
		done, err := clientCommands(ctx, commandsSplit[1:])
		if err != nil {
			log.Fatal(err)
		}
		if done {
			return
		}
	case "etherscan":
		done, err := etherscanCommands(ctx, commandsSplit[1:])
		if err != nil {
			log.Fatal(err)
		}
		if done {
			return
		}
	}

	// WIP
	switch kingpin.Parse() {
	case "account list":
		if *keystore == "" {
			log.Fatal("Parameter --keystore required")
		}
		if err := doAccountList(*keystore); err != nil {
			log.Fatal(err)
		}
	case "account new":
		if *keystore == "" {
			log.Fatal("Parameter --keystore required")
		}
		if *passphrase == "" {
			log.Fatal("Parameter --passphrase required")
		}
		if err := doAccountNew(*keystore, *passphrase); err != nil {
			log.Fatal(err)
		}
	case "account unlock":
		_, unlocked, err := getAccount()
		if err != nil {
			log.Fatal(err)
		}
		if !unlocked {
			log.Fatal("Passphrase required")
		}
		fmt.Printf("Unlocked %s\n", *address)
	case "client status":
		if err := doClientStatus(ctx, *clientServer); err != nil {
			log.Fatal(err)
		}
	case "client balance":
		if *address == "" {
			log.Fatal("Parameter --address required")
		}
		if err := doClientBalance(*clientServer, *address); err != nil {
			log.Fatal(err)
		}
	case "client filterLogs":
		if *address == "" {
			log.Fatal("Parameter --address required")
		}
		if err := doClientFilterLogs(*clientServer); err != nil {
			log.Fatal(err)
		}
	case "client transfer":
		account, unlocked, err := getAccount()
		if err != nil {
			log.Fatal(err)
		}
		if !unlocked {
			log.Fatal("Passphrase required")
		}
		if err := doClientTransfer(*clientServer, account, *destAddress, *transferAmount); err != nil {
			log.Fatal(err)
		}
	case "client kyber expectedRate":
		if *tokenKyberSource == "" || *tokenKyberDest == "" {
			log.Fatal("Both --source and --dest flags required")
		}
		if err := doClientKyberExpectedRate(*clientServer, *tokenKyberSource, *tokenKyberDest, *tokenKyberQuantity); err != nil {
			log.Fatal(err)
		}
	case "client zeroex deposit":
		zClient, err := newZeroexClient()
		if err != nil {
			log.Fatal(err)
		}
		account, unlocked, err := getAccount()
		if err != nil {
			log.Fatal(err)
		}
		if !unlocked {
			log.Fatal("Passphrase required")
		}
		if err := zeroexDepositCommand(zClient, account); err != nil {
			log.Fatal(err)
		}
	case "client zeroex balanceof":
		zClient, err := newZeroexClient()
		if err != nil {
			log.Fatal(err)
		}
		account, _, err := getAccount()
		if err != nil {
			log.Fatal(err)
		}
		if err := zeroexBalanceOfCommand(zClient, account); err != nil {
			log.Fatal(err)
		}
	case "gasstation":
		client := gasstation.NewClient()
		if err := gasstationCommand(context.Background(), client); err != nil {
			log.Fatal(err)
		}
	}
}
