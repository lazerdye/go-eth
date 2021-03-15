package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/lazerdye/go-eth/wallet"
)

const (
	defaultHdPath = "m/44'/60'/0'/0/0"
)

var (
	accountCmd             = kingpin.Command("account", "Account operations")
	listAccounts           = accountCmd.Command("list", "List the accounts")
	newAccountCmd          = accountCmd.Command("new", "Create new account")
	newAccountMnemonic     = newAccountCmd.Flag("mnemonic", "Show Mnemonic").Bool()
	newAccountHdPath       = newAccountCmd.Flag("hdpath", "HD derivition path").Default(defaultHdPath).String()
	unlockAccountCmd       = accountCmd.Command("unlock", "Unlock an account")
	restoreAccountCmd      = accountCmd.Command("restore", "Restore account")
	restoreAccountMnemonic = restoreAccountCmd.Arg("mnemonic", "Restore Mnemonic").String()
	restoreAccountHdPath   = restoreAccountCmd.Flag("hdpath", "HD derivition path").Default(defaultHdPath).String()
)

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

func accountCommands(ctx context.Context, commands []string) (bool, error) {
	switch commands[0] {
	case "list":
		return true, doAccountList()
	case "new":
		return true, doAccountNew()
	case "unlock":
		return true, doAccountUnlock()
	case "restore":
		return true, doAccountRestore()
	}
	return false, nil
}

func doAccountList() error {
	if *keystore == "" {
		log.Fatal("Parameter --keystore required")
	}
	w, err := wallet.Open(*keystore)
	if err != nil {
		return err
	}
	err = w.PrintAccounts()
	if err != nil {
		return err
	}
	return nil
}

func doAccountNew() error {
	if *keystore == "" {
		log.Fatal("Parameter --keystore required")
	}
	if *passphrase == "" {
		log.Fatal("Parameter --passphrase required")
	}
	// Create a new phrase for the account.
	mnemonic, err := hdwallet.NewMnemonic(256)
	if err != nil {
		return err
	}
	if *newAccountMnemonic {
		fmt.Println("Please record this string to restore your new account:")
		fmt.Println(mnemonic)
	}
	return addAccountFromMnemonic(mnemonic, *newAccountHdPath)
}

func doAccountRestore() error {
	var mnemonic string
	if *restoreAccountMnemonic != "" {
		mnemonic = *restoreAccountMnemonic
	} else {
		fmt.Println("Enter the phrase for the account:")
		//var err error
		//in := bufio.NewReader(os.Stdin)
		//mnemonic, err = in.ReadString('\n')
		//if err != nil {
		//	return err
		//}
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		mnemonic = scanner.Text()
	}
	return addAccountFromMnemonic(mnemonic, *restoreAccountHdPath)
}

func addAccountFromMnemonic(mnemonic string, hdpath string) error {
	tmpWallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return err
	}
	path := hdwallet.MustParseDerivationPath(hdpath)
	tmpAccount, err := tmpWallet.Derive(path, false)
	if err != nil {
		return err
	}
	w, err := wallet.Open(*keystore)
	if err != nil {
		return err
	}
	privKey, err := tmpWallet.PrivateKey(tmpAccount)
	if err != nil {
		return err
	}
	account, err := w.ImportAccount(privKey, *passphrase)
	if err != nil {
		return err
	}
	err = w.PrintAccount(account)
	if err != nil {
		return err
	}
	return nil
}

func doAccountUnlock() error {
	_, unlocked, err := getAccount()
	if err != nil {
		log.Fatal(err)
	}
	if !unlocked {
		log.Fatal("Passphrase required")
	}
	fmt.Printf("Unlocked %s\n", *address)
	return nil
}
