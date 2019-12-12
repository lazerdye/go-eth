package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	"gitlab.com/lazerdye/go-eth/client"
	"gitlab.com/lazerdye/go-eth/wallet"
)

var (
	//keystore = kingpin.Flag("keystore", "Location of the wallet keystore").String()
	//password = kingpin.Flag("password", "Password for keystore").Envar("WALLET_PASSWORD").String()

	accountCmd      = kingpin.Command("account", "Account operations")
	accountKeystore = accountCmd.Flag("keystore", "Location of the wallet keystore").Required().String()
	list            = accountCmd.Command("list", "List the accounts")

	clientCmd           = kingpin.Command("client", "Client operations")
	clientServer        = clientCmd.Flag("server", "URL of the server to connect to").Required().String()
	balanceCmd          = clientCmd.Command("balance", "Get the balance of an account")
	balanceAccount      = balanceCmd.Arg("account", "Account to get balance of").Required().String()
	tokenCmd            = clientCmd.Command("token", "Token operations")
	tokenBalanceCmd     = tokenCmd.Command("balance", "Get the balance of the token")
	tokenName           = tokenBalanceCmd.Arg("name", "Name of the token").Required().String()
	tokenBalanceAccount = tokenBalanceCmd.Arg("account", "Account to get token balance of").Required().String()
)

func doAccount(keystore string) error {
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

func doClientBalance(server, account string) error {
	c, err := client.Dial(server)
	if err != nil {
		return err
	}
	eth, err := c.EthBalanceOf(context.Background(), account)
	if err != nil {
		return err
	}
	fmt.Printf("Balance of %s: %g\n", account, eth)

	return nil
}

func doClientTokenBalance(server, tokenName, account string) error {
	c, err := client.Dial(server)
	if err != nil {
		return err
	}
	token, err := c.Token(tokenName)
	if err != nil {
		return err
	}
	bal, err := token.BalanceOf(context.Background(), account)
	if err != nil {
		return err
	}
	fmt.Printf("Balance of %s in %s: %g\n", account, tokenName, bal)
    return nil
}

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("0.1").Author("Terence Haddock")
	kingpin.CommandLine.Help = "Ethereum test client"

	switch kingpin.Parse() {
	case "account list":
		if err := doAccount(*accountKeystore); err != nil {
			log.Fatal(err)
		}
	case "client balance":
		if err := doClientBalance(*clientServer, *balanceAccount); err != nil {
			log.Fatal(err)
		}
	case "client token balance":
		if err := doClientTokenBalance(*clientServer, *tokenName, *tokenBalanceAccount); err != nil {
			log.Fatal(err)
		}
	}
}
