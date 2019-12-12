package wallet

import (
	"fmt"
	//"io/ioutil"

	//log "github.com/sirupsen/logrus"
	//"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type Wallet struct {
	ks *keystore.KeyStore
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
