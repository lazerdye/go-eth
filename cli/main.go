package main

import (
	"context"
	"fmt"
	"github.com/lazerdye/go-eth/etherscan"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasoracle"
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

	transferCmd      = clientCmd.Command("transfer", "Transfer ethereum")
	transferTransmit = transferCmd.Flag("transmit", "Transmit transaction").Bool()

	tokenKyberCmd          = clientCmd.Command("kyber", "Kyber Network")
	tokenKyberTokenFile    = tokenKyberCmd.Flag("token-file", "Token file").Required().String()
	tokenKyberSource       = tokenKyberCmd.Flag("source", "Source exchange").String()
	tokenKyberDest         = tokenKyberCmd.Flag("dest", "Dest exchange").String()
	tokenKyberQuantity     = tokenKyberCmd.Flag("quantity", "Source quantity").Float64()
	tokenKyberExpectedRate = tokenKyberCmd.Command("expectedRate", "Expected rate")
)

func doClientTransfer(server string, account *wallet.Account, destAddress string, amount float64) error {
	ctx := context.Background()
	o := gasoracle.GasOracleFromEtherscan(etherscan.NewClient(*etherscanApikey))
	c, err := client.Dial(server, o)
	if err != nil {
		return err
	}
	tx, err := c.Transfer(ctx, account, common.HexToAddress(destAddress), decimal.NewFromFloat(amount), *transferTransmit)
	if err != nil {
		return err
	}
	if *transferTransmit {
		fmt.Printf("Transaction: %s\n", tx.Hash().Hex())
	} else {
		fmt.Printf("Transaction(unsent): %s\n", tx.Hash().Hex())
	}
	return nil
}

func doClientKyberExpectedRate(server string, source, dest string, quantity float64) error {
	o := gasoracle.GasOracleFromEtherscan(etherscan.NewClient(*etherscanApikey))
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

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("0.1").Author("Lazerdye")
	kingpin.CommandLine.Help = "Ethereum test client"

	commands := kingpin.Parse()
	commandsSplit := strings.Split(commands, " ")
	ctx := context.Background()
	switch commandsSplit[0] {
	case "account":
		done, err := accountCommands(ctx, commandsSplit[1:])
		if err != nil {
			log.Fatal(err)
		}
		if done {
			return
		}
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
	case "zxmesh":
		done, err := zxmeshCommands(ctx, commandsSplit[1:])
		if err != nil {
			log.Fatal(err)
		}
		if done {
			return
		}
	}

	// WIP
	switch kingpin.Parse() {
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
	}
}
