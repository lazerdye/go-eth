package main

import (
	"context"
	"fmt"
	"syscall"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/term"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/lazerdye/go-eth/wallet"
)

var (
	accountCmd             = kingpin.Command("account", "Account operations")
	listAccounts           = accountCmd.Command("list", "List the accounts")
	newAccountCmd          = accountCmd.Command("new", "Create new account")
	unlockAccountCmd       = accountCmd.Command("unlock", "Unlock an account")
	restoreAccountCmd      = accountCmd.Command("restore", "Restore account")
	restoreAccountMnemonic = restoreAccountCmd.Arg("mnemonic", "Restore Mnemonic").String()
)

func getPassphrase(repeat bool) (string, error) {
	if *passphrase != "" {
		return *passphrase, nil
	}
	fmt.Print("Enter Passphrase: ")
	passwordBytes1, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	fmt.Print("\n")
	if repeat {
		fmt.Print("Repeat: ")
		passwordBytes2, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return "", err
		}
		fmt.Print("\n")
		if string(passwordBytes1) != string(passwordBytes2) {
			return "", errors.New("Passwords do not match")
		}
	}
	return string(passwordBytes1), nil
}

func getAccount(needUnlocked bool) (*wallet.Account, error) {
	if *address == "" {
		return nil, errors.New("address parameter required")
	}
	if *keystore == "" {
		return nil, errors.New("keystore parameter required")
	}
	w, err := wallet.Open(*keystore)
	if err != nil {
		return nil, err
	}
	account, err := w.Account(*address)
	if err != nil {
		return nil, err
	}
	unlockPassphrase := *passphrase
	// If we do not have a password set by CLI, try getting one from the user.
	if unlockPassphrase == "" && needUnlocked {
		var err error
		unlockPassphrase, err = getPassphrase(false)
		if err != nil {
			return nil, err
		}
		if unlockPassphrase == "" {
			return nil, errors.New("Passphrase required")
		}
	}
	if unlockPassphrase != "" {
		if err := account.Unlock(unlockPassphrase); err != nil {
			return nil, err
		}
	}
	return account, nil
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
	passphrase, err := getPassphrase(true)
	if err != nil {
		return err
	}
	w, err := wallet.Open(*keystore)
	if err != nil {
		return err
	}
	account, err := w.NewAccount(passphrase)
	if err != nil {
		return err
	}
	return w.PrintAccount(account)
}

func doAccountUnlock() error {
	_, err := getAccount(true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Unlocked %s\n", *address)
	return nil
}
