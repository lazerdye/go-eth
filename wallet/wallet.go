package wallet

import (
	"fmt"
	//"io/ioutil"
	"math/big"

	//log "github.com/sirupsen/logrus"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Wallet struct {
	ks *keystore.KeyStore
}

type Account struct {
	ks      *keystore.KeyStore
	Account accounts.Account
}

func Open(dir string) (*Wallet, error) {
	ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	return &Wallet{ks: ks}, nil
}

func (w *Wallet) PrintAccounts() error {
	for _, account := range w.ks.Accounts() {
		fmt.Printf("Account: %s\n", account.Address.Hex())
	}
	return nil
}

func (w *Wallet) Account(address string) (*Account, error) {
	account, err := w.ks.Find(accounts.Account{common.HexToAddress(address), accounts.URL{}})
	if err != nil {
		fmt.Printf("XXX %s\n", address)
		return nil, err
	}
	return &Account{w.ks, account}, nil
}

func (a *Account) Unlock(password string) error {
	return a.ks.Unlock(a.Account, password)
}

func (a *Account) SignTx(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	return a.ks.SignTx(a.Account, tx, chainID)
}
