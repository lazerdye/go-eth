package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/wallet"
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

	tokenDutchExchangeCmd            = tokenCmd.Command("dutchx", "DutchX operations")
	tokenDutchExchangeAddress1       = tokenDutchExchangeCmd.Flag("address1", "Address 1").String()
	tokenDutchExchangeAddress2       = tokenDutchExchangeCmd.Flag("address2", "Address 2").String()
	tokenDutchExchangeAccount        = tokenDutchExchangeCmd.Flag("account", "Account").String()
	tokenDutchExchangeGetAuctionInfo = tokenDutchExchangeCmd.Command("getAuctionInfo", "Get auction info")
	tokenDutchExchangeGetFeeRatio    = tokenDutchExchangeCmd.Command("getFeeRatio", "Get fee ratio")
	tokenDutchExchangeListen         = tokenDutchExchangeCmd.Command("listen", "Listen for events")

	tokenKyberCmd          = tokenCmd.Command("kyber", "Kyber Network")
	tokenKyberSource       = tokenKyberCmd.Flag("source", "Source exchange").String()
	tokenKyberDest         = tokenKyberCmd.Flag("dest", "Dest exchange").String()
	tokenKyberQuantity     = tokenKyberCmd.Flag("quantity", "Source quantity").Int64()
	tokenKyberExpectedRate = tokenKyberCmd.Command("expectedRate", "Expected rate")
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

func doClientBalance(server, addressStr string) error {
	address := common.HexToAddress(addressStr)
	c, err := client.Dial(server)
	if err != nil {
		return err
	}
	eth, err := c.BalanceOf(context.Background(), address)
	if err != nil {
		return err
	}
	fmt.Printf("Balance of %s: %g\n", address.String(), eth)

	return nil
}

func doClientTokenBalance(server, tokenName, addressStr string) error {
	c, err := client.Dial(server)
	if err != nil {
		return err
	}
	token, err := c.Token(tokenName)
	if err != nil {
		return err
	}
	address := common.HexToAddress(addressStr)
	bal, err := token.BalanceOf(context.Background(), address)
	if err != nil {
		return err
	}
	fmt.Printf("Balance of %s in %s: %g\n", address, tokenName, bal)
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

func doClientTokenDutchxGetAuctionInfo(server, address1Str, address2Str string) error {
	c, err := client.Dial(server)
	if err != nil {
		return err
	}
	d, err := c.Dutchx()
	if err != nil {
		return err
	}
	address1 := common.HexToAddress(address1Str)
	address2 := common.HexToAddress(address2Str)
	index, err := d.GetAuctionIndex(address1, address2)
	if err != nil {
		return err
	}
	fmt.Printf("Current Auction: %d\n", index)
	start, err := d.GetAuctionStart(address1, address2)
	if err != nil {
		return err
	}
	fmt.Printf("Auction Start: %d\n", start.Int64())
	clear, err := d.GetAuctionClearingTime(address1, address2, index)
	if err != nil {
		return err
	}
	fmt.Printf("Auction Clearing Time: %d\n", clear.Int64())
	prices, err := d.GetCurrentAuctionPrice(address1, address2, index)
	if err != nil {
		return err
	}
	fmt.Printf("Auction Prices: %s (%s)\n", prices.Price.String(), prices.InvPrice.String())
	volume, err := d.GetCurrentAuctionVolume(address1, address2)
	if err != nil {
		return err
	}
	fmt.Printf("Auction Volume: %s / %s\n", volume.Buy.String(), volume.Sell.String())
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

func doClientTokenDutchxGetFeeRatio(server string, address string) error {
	c, err := client.Dial(server)
	if err != nil {
		return err
	}
	d, err := c.Dutchx()
	if err != nil {
		return err
	}
	ratio, err := d.GetFeeRatio(common.HexToAddress(address))
	if err != nil {
		return err
	}
	fmt.Printf("Fee Ratio for account %s: %s\n", address, ratio.Price.String())
	return nil
}

func doClientTokenKyberExpectedRate(server string, source, dest string, quantity int64) error {
	c, err := client.Dial(server)
	if err != nil {
		return err
	}
	k, err := c.Kyber()
	if err != nil {
		return err
	}
	rate, err := k.GetExpectedRate(common.HexToAddress(source), common.HexToAddress(dest), quantity)
	if err != nil {
		return err
	}
	fmt.Printf("Rate: %s\n", rate.String())
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
	case "client token dutchx getAuctionInfo":
		if err := doClientTokenDutchxGetAuctionInfo(*clientServer, *tokenDutchExchangeAddress1, *tokenDutchExchangeAddress2); err != nil {
			log.Fatal(err)
		}
	case "client token dutchx getFeeRatio":
		if err := doClientTokenDutchxGetFeeRatio(*clientServer, *tokenDutchExchangeAccount); err != nil {
			log.Fatal(err)
		}
	case "client token dutchx listen":
		if err := doClientTokenDutchxListen(*clientServer); err != nil {
			log.Fatal(err)
		}
	case "client token kyber expectedRate":
		if err := doClientTokenKyberExpectedRate(*clientServer, *tokenKyberSource, *tokenKyberDest, *tokenKyberQuantity); err != nil {
			log.Fatal(err)
		}
	}
}
