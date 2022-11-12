package main

import (
	"context"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

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
)

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
		err := accountCommands(ctx, commandsSplit[1:])
		if err != nil {
			log.Fatal(err)
		}
	case "client":
		err := clientCommands(ctx, commandsSplit[1:])
		if err != nil {
			log.Fatal(err)
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
		err := zxmeshCommands(ctx, commandsSplit[1:])
		if err != nil {
			log.Fatal(err)
		}
	}
}
