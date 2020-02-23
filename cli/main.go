package main

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/etherscan"
	"github.com/lazerdye/go-eth/wallet"
)

var (
	keystore       = kingpin.Flag("keystore", "Location of the wallet keystore").Envar("WALLET_KEYSTORE").String()
	passphrase     = kingpin.Flag("passphrase", "Passphrase for keystore").Envar("WALLET_PASSPHRASE").String()
	address        = kingpin.Flag("address", "Address for account operations").String()
	destAddress    = kingpin.Flag("dest-address", "Destination address for transfer").String()
	transferAmount = kingpin.Flag("amount", "Transfer amount").Float64()

	accountCmd       = kingpin.Command("account", "Account operations")
	listAccounts     = accountCmd.Command("list", "List the accounts")
	newAccount       = accountCmd.Command("new", "Create new account")
	unlockAccountCmd = accountCmd.Command("unlock", "Unlock an account")

	clientCmd        = kingpin.Command("client", "Client operations")
	clientServer     = clientCmd.Flag("server", "URL of the server to connect to").Required().String()
	balanceCmd       = clientCmd.Command("balance", "Get the balance of an account")
	transferCmd      = clientCmd.Command("transfer", "Transfer ethereum")
	transferTransmit = transferCmd.Flag("transmit", "Transmit transaction").Bool()
	tokenCmd         = clientCmd.Command("token", "Token operations")
	tokenBalanceCmd  = tokenCmd.Command("balance", "Get the balance of the token")
	tokenName        = tokenBalanceCmd.Arg("name", "Name of the token").Required().String()

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
	tokenKyberQuantity     = tokenKyberCmd.Flag("quantity", "Source quantity").Float64()
	tokenKyberExpectedRate = tokenKyberCmd.Command("expectedRate", "Expected rate")

	etherscanApikey = kingpin.Flag("etherscan-apikey", "Key for etherscan api").Envar("ETHERSCAN_APIKEY").String()
	etherscanCmd    = kingpin.Command("etherscan", "Etherscan operations")
	etherscanList   = etherscanCmd.Command("list", "List transactions for account")
	etherscanSort   = etherscanList.Arg("sort", "Sort").Default("asc").String()
	etherscanPage   = etherscanList.Arg("page", "Page number").Default("1").Int()
	etherscanOffset = etherscanList.Arg("offset", "Page offset").Default("20").Int()
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

func unlockAccount(keystore, passphrase, address string) (*wallet.Account, error) {
	w, err := wallet.Open(keystore)
	if err != nil {
		return nil, err
	}
	account, err := w.Account(address)
	if err != nil {
		return nil, err
	}
	return account, account.Unlock(passphrase)
}

func doAccountUnlock(keystore string, passphrase string, address string) error {
	_, err := unlockAccount(keystore, passphrase, address)
	return err
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

func doClientTransfer(server string, account *wallet.Account, destAddress string, amount float64) error {
	ctx := context.Background()
	c, err := client.Dial(server)
	if err != nil {
		return err
	}
	tx, err := c.Transfer(ctx, account, destAddress, big.NewFloat(amount), false)
	if err != nil {
		return err
	}
	fmt.Printf("Transaction(unsent): %+v\n", tx)
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
	fmt.Printf("Balance of %s in %s: %g\n", address.String(), tokenName, bal)
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

func doClientTokenKyberExpectedRate(server string, source, dest string, quantityFloat float64) error {
	c, err := client.Dial(server)
	if err != nil {
		return err
	}
	k, err := c.Kyber()
	if err != nil {
		return err
	}
	quantityInt, _ := new(big.Float).Mul(new(big.Float).SetFloat64(quantityFloat), big.NewFloat(math.Pow10(18))).Int(nil)
	expectedRate, slippageRate, err := k.GetExpectedRate(common.HexToAddress(source), common.HexToAddress(dest), quantityInt)
	if err != nil {
		return err
	}
	fmt.Printf("Expected Rate: %s\n", expectedRate.String())
	fmt.Printf("Slippage Rate: %s\n", slippageRate.String())
	return nil
}

func doEtherscanList(apikey string, address string, page, offset int, sort string) error {
	ctx := context.Background()

	c := etherscan.NewClient(apikey)
	transactions, err := c.NormalTransactionsByAddress(ctx, address, page, offset, sort)
	if err != nil {
		return err
	}

	for _, t := range transactions {
		fmt.Printf("%s\t%s\t%s\t%s\t%s\n", t.Hash, t.BlockNumber, t.From, t.To, t.Value)
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
		if *keystore == "" {
			log.Fatal("Parameter --keystore required")
		}
		if *passphrase == "" {
			log.Fatal("Parameter --passphrase required")
		}
		if *address == "" {
			log.Fatal("Parameter --address required")
		}
		if err := doAccountUnlock(*keystore, *passphrase, *address); err != nil {
			log.Fatal(err)
		}
	case "client balance":
		if *address == "" {
			log.Fatal("Parameter --address required")
		}
		if err := doClientBalance(*clientServer, *address); err != nil {
			log.Fatal(err)
		}
	case "client transfer":
		if *keystore == "" || *passphrase == "" {
			log.Fatal("Parameter --keystore and --passphrase required")
		}
		account, err := unlockAccount(*keystore, *passphrase, *tokenTransferSourceAccount)
		if err != nil {
			log.Fatal(err)
		}
		if err := doClientTransfer(*clientServer, account, *destAddress, *transferAmount); err != nil {
			log.Fatal(err)
		}
	case "client token balance":
		if *address == "" {
			log.Fatal("Parameter --address required")
		}
		if err := doClientTokenBalance(*clientServer, *tokenName, *address); err != nil {
			log.Fatal(err)
		}
	case "client token transfer":
		if *keystore == "" || *passphrase == "" {
			log.Fatal("Parameter --keystore and --passphrase required")
		}
		account, err := unlockAccount(*keystore, *passphrase, *tokenTransferSourceAccount)
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
		if *tokenKyberSource == "" || *tokenKyberDest == "" {
			log.Fatal("Both --source and --dest flags required")
		}
		if err := doClientTokenKyberExpectedRate(*clientServer, *tokenKyberSource, *tokenKyberDest, *tokenKyberQuantity); err != nil {
			log.Fatal(err)
		}
	case "etherscan list":
		if *etherscanApikey == "" {
			log.Fatal("--etherscan-api-key required")
		}
		if *address == "" {
			log.Fatal("--address required")
		}
		if err := doEtherscanList(*etherscanApikey, *address, *etherscanPage, *etherscanOffset, *etherscanSort); err != nil {
			log.Fatal(err)
		}
	}
}
