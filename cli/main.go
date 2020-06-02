package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/compound"
	"github.com/lazerdye/go-eth/gasstation"
	"github.com/lazerdye/go-eth/kyber"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/uniswapv1"
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

	balanceCmd       = clientCmd.Command("balance", "Get the balance of an account")
	transferCmd      = clientCmd.Command("transfer", "Transfer ethereum")
	transferTransmit = transferCmd.Flag("transmit", "Transmit transaction").Bool()

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

func doClientBalance(server, addressStr string) error {
	address := common.HexToAddress(addressStr)
	c, err := client.Dial(server, gasstation.NewClient())
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
	c, err := client.Dial(server, gasstation.NewClient())
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

func doClientKyberExpectedRate(server string, source, dest string, quantity float64) error {
	c, err := client.Dial(server, gasstation.NewClient())
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

func newCompoundClient() (*compound.Client, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}
	return compound.NewClient(client)
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

func newUniswapV1Client() (*uniswapv1.Client, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}
	return uniswapv1.NewClient(client)
}

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("0.1").Author("Terence Haddock")
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
	case "client balance":
		if *address == "" {
			log.Fatal("Parameter --address required")
		}
		if err := doClientBalance(*clientServer, *address); err != nil {
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
	case "etherscan list":
		if *address == "" {
			log.Fatal("--address required")
		}
		if err := doEtherscanList(context.Background(), *address); err != nil {
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
	case "client uniswapv1 get-exchange":
		uClient, err := newUniswapV1Client()
		if err != nil {
			log.Fatal(err)
		}
		if err := uniswapGetExchange(context.Background(), uClient); err != nil {
			log.Fatal(err)
		}
	case "client uniswapv1 eth-to-token-input":
		uClient, err := newUniswapV1Client()
		if err != nil {
			log.Fatal(err)
		}
		if err := uniswapGetEthToTokenInputPrice(context.Background(), uClient); err != nil {
			log.Fatal(err)
		}
	case "client uniswapv1 eth-to-token-output":
		uClient, err := newUniswapV1Client()
		if err != nil {
			log.Fatal(err)
		}
		if err := uniswapGetEthToTokenOutputPrice(context.Background(), uClient); err != nil {
			log.Fatal(err)
		}
	case "client uniswapv1 token-to-eth-input":
		uClient, err := newUniswapV1Client()
		if err != nil {
			log.Fatal(err)
		}
		if err := uniswapGetTokenToEthInputPrice(context.Background(), uClient); err != nil {
			log.Fatal(err)
		}
	case "client uniswapv1 token-to-eth-output":
		uClient, err := newUniswapV1Client()
		if err != nil {
			log.Fatal(err)
		}
		if err := uniswapGetTokenToEthOutputPrice(context.Background(), uClient); err != nil {
			log.Fatal(err)
		}
	case "client compound mint":
		cClient, err := newCompoundClient()
		if err != nil {
			log.Fatal(err)
		}
		account, _, err := getAccount()
		if err != nil {
			log.Fatal(err)
		}
		if err := compoundMint(context.Background(), cClient, account); err != nil {
			log.Fatal(err)
		}
	case "gasstation":
		client := gasstation.NewClient()
		if err := gasstationCommand(context.Background(), client); err != nil {
			log.Fatal(err)
		}
	}
}
