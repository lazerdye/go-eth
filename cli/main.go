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
	keystore = kingpin.Flag("keystore", "Location of the wallet keystore").String()
	password = kingpin.Flag("password", "Password for keystore").Envar("WALLET_PASSWORD").String()

	accountCmd = kingpin.Command("account", "Account operations")
	list       = accountCmd.Command("list", "List the accounts")

	clientCmd           = kingpin.Command("client", "Client operations")
	clientServer        = clientCmd.Flag("server", "URL of the server to connect to").Required().String()
	balanceCmd          = clientCmd.Command("balance", "Get the balance of an account")
	balanceAccount      = balanceCmd.Arg("account", "Account to get balance of").Required().String()
	tokenCmd            = clientCmd.Command("token", "Token operations")
	tokenBalanceCmd     = tokenCmd.Command("balance", "Get the balance of the token")
	tokenName           = tokenBalanceCmd.Arg("name", "Name of the token").Required().String()
	tokenBalanceAccount = tokenBalanceCmd.Arg("account", "Account to get token balance of").Required().String()

	tokenTransferCmd           = tokenCmd.Command("transfer", "Transfer a token")
	tokenTransferName          = tokenTransferCmd.Arg("name", "Name of the token to transfer").Required().String()
	tokenTransferAmount        = tokenTransferCmd.Arg("amount", "Amount of the token to transfer").Required().Float64()
	tokenTransferSourceAccount = tokenTransferCmd.Arg("source-account", "Account to transfer token from").Required().String()
	tokenTransferDestAccount   = tokenTransferCmd.Arg("dest-account", "Account to transfer token to").Required().String()
	tokenTransferTransmitFlag  = tokenTransferCmd.Flag("transmit", "Transmit transaction").Bool()

	tokenDutchExchangeCmd             = tokenCmd.Command("dutchx", "DutchX operations")
	tokenDutchExchangeAddress1        = tokenDutchExchangeCmd.Flag("address1", "Address 1").String()
	tokenDutchExchangeAddress2        = tokenDutchExchangeCmd.Flag("address2", "Address 1").String()
	tokenDutchExchangeGetAuctionIndex = tokenDutchExchangeCmd.Command("getAuctionIndex", "Get auction index")
	tokenDutchExchangeListen          = tokenDutchExchangeCmd.Command("listen", "Listen for events")
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

func unlockAccount(keystore, password, address string) (*wallet.Account, error) {
	w, err := wallet.Open(keystore)
	if err != nil {
		return nil, err
	}
	account, err := w.Account(address)
	if err != nil {
		return nil, err
	}
	return account, account.Unlock(password)
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

func doClientTokenTransfer(server string, account *wallet.Account, tokenName, sourceAccount, destAccount string,
	amount float64, transmit bool) error {
	c, err := client.Dial(server)
	if err != nil {
		return err
	}
	token, err := c.Token(tokenName)
	if err != nil {
		return err
	}
	if err := token.Transfer(account, destAccount, amount, transmit); err != nil {
		return err
	}
	return nil
}

func doClientTokenDutchxGetAuctionIndex(server, address1, address2 string) error {
	c, err := client.Dial(server)
	if err != nil {
		return err
	}
	d, err := c.Dutchx()
	if err != nil {
		return err
	}
	err = d.GetAuctionIndex(address1, address2)
	if err != nil {
		return err
	}
	return nil
}

func doClientTokenDutchxListen(server string) error {
	c, err := client.Dial(server)
	if err != nil {
		return err
	}
	d, err := c.Dutchx()
	if err != nil {
		return err
	}
	err = d.Listen()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("0.1").Author("Terence Haddock")
	kingpin.CommandLine.Help = "Ethereum test client"

	switch kingpin.Parse() {
	case "account list":
		if *keystore == "" {
			log.Fatal("Parameter --keystore required")
		}
		if err := doAccount(*keystore); err != nil {
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
	case "client token transfer":
		fmt.Printf("XXX %s %s\n", *keystore, password)
		if *keystore == "" || *password == "" {
			log.Fatal("Parameter --keystore and --password required")
		}
		account, err := unlockAccount(*keystore, *password, *tokenTransferSourceAccount)
		if err != nil {
			log.Fatal(err)
		}
		if err := doClientTokenTransfer(*clientServer, account, *tokenTransferName, *tokenTransferSourceAccount,
			*tokenTransferDestAccount, *tokenTransferAmount, *tokenTransferTransmitFlag); err != nil {
			log.Fatal(err)
		}
	case "client token dutchx getAuctionIndex":
		if err := doClientTokenDutchxGetAuctionIndex(*clientServer, *tokenDutchExchangeAddress1, *tokenDutchExchangeAddress2); err != nil {
			log.Fatal(err)
		}
	case "client token dutchx listen":
		if err := doClientTokenDutchxListen(*clientServer); err != nil {
			log.Fatal(err)
		}
	}
}
