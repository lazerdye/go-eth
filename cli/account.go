package main

import (
	"context"
	"fmt"

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
	unlockAccountCmd       = accountCmd.Command("unlock", "Unlock an account")
	restoreAccountCmd      = accountCmd.Command("restore", "Restore account")
	restoreAccountMnemonic = restoreAccountCmd.Arg("mnemonic", "Restore Mnemonic").String()
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
	w, err := wallet.Open(*keystore)
	if err != nil {
		return err
	}
	account, err := w.NewAccount(*passphrase)
	if err != nil {
		return err
	}
	return w.PrintAccount(account)
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
