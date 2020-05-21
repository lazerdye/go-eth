package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/compound"
	"github.com/lazerdye/go-eth/gasstation"
	"github.com/lazerdye/go-eth/kyber"
	"github.com/lazerdye/go-eth/token"
	"github.com/lazerdye/go-eth/uniswapv1"
	"github.com/lazerdye/go-eth/uniswapv2"
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

	clientCmd              = kingpin.Command("client", "Client operations")
	clientServer           = clientCmd.Flag("server", "URL of the server to connect to").Envar("ETHEREUM_SERVER").Required().String()
	balanceCmd             = clientCmd.Command("balance", "Get the balance of an account")
	transferCmd            = clientCmd.Command("transfer", "Transfer ethereum")
	transferTransmit       = transferCmd.Flag("transmit", "Transmit transaction").Bool()
	tokenCmd               = clientCmd.Command("token", "Token operations")
	tokenBalanceCmd        = tokenCmd.Command("balance", "Get the balance of the token")
	tokenBalanceName       = tokenBalanceCmd.Arg("name", "Name of the token").Required().String()
	tokenApproveCmd        = tokenCmd.Command("approve", "Approve a contract to act on the accounts behalf")
	tokenApproveName       = tokenApproveCmd.Arg("name", "Name of the token").Required().String()
	tokenApproveContract   = tokenApproveCmd.Arg("contract", "Contract to approve").Required().String()
	tokenApproveAmount     = tokenApproveCmd.Arg("amount", "Amount (float) to approve").Required().String()
	tokenAllowanceCmd      = tokenCmd.Command("allowance", "Get the allowance for the given contract")
	tokenAllowanceName     = tokenAllowanceCmd.Arg("name", "Name of token to allow access to contract").Required().String()
	tokenAllowanceContract = tokenAllowanceCmd.Arg("contract", "Contract to check allowance of").Required().String()

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
	tx, err := c.Transfer(ctx, account, destAddress, big.NewFloat(amount), false)
	if err != nil {
		return err
	}
	fmt.Printf("Transaction(unsent): %+v\n", tx)
	return nil
}

func doClientTokenBalance(server, tokenName, addressStr string) error {
	c, err := client.Dial(server, gasstation.NewClient())
	if err != nil {
		return err
	}
	token, err := token.ByName(c, tokenName)
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

func doClientTokenAllowance(server, tokenName, contractStr, addressStr string) error {
	c, err := client.Dial(server, gasstation.NewClient())
	if err != nil {
		return err
	}
	token, err := token.ByName(c, tokenName)
	if err != nil {
		return err
	}
	contract := common.HexToAddress(contractStr)
	address := common.HexToAddress(addressStr)
	fmt.Printf("%s %s\n", contract.String(), address.String())
	allowance, err := token.Allowance(context.Background(), address, contract)
	if err != nil {
		return err
	}
	fmt.Printf("%s allowance for contract %s on %s is: %s\n", tokenName, contract.String(), address.String(), allowance.String())
	return nil
}

func doClientTokenTransfer(server string, account *wallet.Account, tokenName, sourceAccount, destAccount string,
	amount float64, transmit bool) error {
	if !transmit {
		return errors.New("Sending without transmitting is currently not supported")
	}
	c, err := client.Dial(server, gasstation.NewClient())
	if err != nil {
		return err
	}
	token, err := token.ByName(c, tokenName)
	if err != nil {
		return err
	}
	tx, err := token.Transfer(context.Background(), account, common.HexToAddress(destAccount), big.NewFloat(amount))
	if err != nil {
		return err
	}
	fmt.Printf("Transfer transaction: %s\n", tx.Hash().Hex())
	return nil
}

func doClientTokenApprove(server string, account *wallet.Account, tokenName, contractStr, amountStr string) error {
	c, err := client.Dial(server, gasstation.NewClient())
	if err != nil {
		return err
	}
	token, err := token.ByName(c, tokenName)
	if err != nil {
		return err
	}
	contract := common.HexToAddress(contractStr)
	amount, _, err := new(big.Float).Parse(amountStr, 10)
	if err != nil {
		return err
	}
	tx, err := token.Approve(context.Background(), account, contract, amount)
	if err != nil {
		return err
	}
	fmt.Printf("Transaction: %s\n", tx.Hash().String())
	return nil
}

func doClientTokenDutchxGetAuctionInfo(server, address1Str, address2Str string) error {
	c, err := client.Dial(server, gasstation.NewClient())
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
	c, err := client.Dial(server, gasstation.NewClient())
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
	c, err := client.Dial(server, gasstation.NewClient())
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

func doClientTokenKyberExpectedRate(server string, source, dest string, quantity float64) error {
	c, err := client.Dial(server, gasstation.NewClient())
	if err != nil {
		return err
	}
	k, err := kyber.NewClient(c)
	if err != nil {
		return err
	}
	quantityFloat := big.NewFloat(quantity)
	sourceToken, ok := token.DefaultRegistry.ByName(source)
	if !ok {
		return errors.Errorf("Unknown Token: %s", source)
	}
	destToken, ok := token.DefaultRegistry.ByName(dest)
	if !ok {
		return errors.Errorf("Unknown Token: %s", dest)
	}
	expectedRate, slippageRate, err := k.GetExpectedRate(context.Background(), sourceToken, destToken, quantityFloat)
	if err != nil {
		return err
	}
	fmt.Printf("Expected Rate: %s\n", expectedRate.String())
	fmt.Printf("Slippage Rate: %s\n", slippageRate.String())
	return nil
}

func newClient() (*client.Client, error) {
	if *clientServer == "" {
		return nil, errors.New("clientServer parameter required")
	}
	return client.Dial(*clientServer, gasstation.NewClient())
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

func newUniswapV2Client() (*uniswapv2.Client, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}
	return uniswapv2.NewClient(client)
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
	case "client token balance":
		if *address == "" {
			log.Fatal("Parameter --address required")
		}
		if err := doClientTokenBalance(*clientServer, *tokenBalanceName, *address); err != nil {
			log.Fatal(err)
		}
	case "client token allowance":
		if *address == "" {
			log.Fatal("Parameter --address required")
		}
		if err := doClientTokenAllowance(*clientServer, *tokenAllowanceName, *tokenAllowanceContract, *address); err != nil {
			log.Fatal(err)
		}
	case "client token approve":
		if *keystore == "" || *passphrase == "" {
			log.Fatal("Parameter --keystore and --passphrase required")
		}
		if *address == "" {
			log.Fatal("Parameter --address required")
		}
		account, unlocked, err := getAccount()
		if err != nil {
			log.Fatal(err)
		}
		if !unlocked {
			log.Fatal("Passphrase required")
		}
		if err := doClientTokenApprove(*clientServer, account, *tokenApproveName, *tokenApproveContract, *tokenApproveAmount); err != nil {
			log.Fatal(err)
		}
	case "client token transfer":
		if *keystore == "" || *passphrase == "" {
			log.Fatal("Parameter --keystore and --passphrase required")
		}
		account, unlocked, err := getAccount()
		if err != nil {
			log.Fatal(err)
		}
		if !unlocked {
			log.Fatal("Passphrase required")
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
	case "client uniswapv2 get-pairs":
		uClient, err := newUniswapV2Client()
		if err != nil {
			log.Fatal(err)
		}
		if err := uniswapV2GetPairs(context.Background(), uClient); err != nil {
			log.Fatal(err)
		}
	case "client uniswapv2 get-amount-out":
		uClient, err := newUniswapV2Client()
		if err != nil {
			log.Fatal(err)
		}
		if err := uniswapV2GetAmountOut(context.Background(), uClient); err != nil {
			log.Fatal(err)
		}
	case "client uniswapv2 get-amount-in":
		uClient, err := newUniswapV2Client()
		if err != nil {
			log.Fatal(err)
		}
		if err := uniswapV2GetAmountIn(context.Background(), uClient); err != nil {
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
