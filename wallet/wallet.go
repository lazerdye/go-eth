package wallet

import (
	"context"
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"
	//log "github.com/sirupsen/logrus"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Wallet struct {
	ks *keystore.KeyStore
}

type Account struct {
	ks            *keystore.KeyStore
	overrideNonce *big.Int
	Account       accounts.Account
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

func (w *Wallet) PrintAccount(account *Account) error {
	fmt.Printf("Account: %s\n", account.Account.Address.Hex())
	return nil
}

func (w *Wallet) Account(address string) (*Account, error) {
	account, err := w.ks.Find(accounts.Account{common.HexToAddress(address), accounts.URL{}})
	if err != nil {
		return nil, err
	}
	return &Account{w.ks, nil, account}, nil
}

func (w *Wallet) NewAccount(passphrase string) (*Account, error) {
	a, err := w.ks.NewAccount(passphrase)
	if err != nil {
		return nil, err
	}
	return &Account{ks: w.ks, Account: a}, nil
}

func (a *Account) OverrideNonce(nonce *big.Int) {
	a.overrideNonce = nonce
}

func (a *Account) NextNonceOverride() *big.Int {
	var nonce *big.Int
	if a.overrideNonce != nil {
		nonce = a.overrideNonce
		a.overrideNonce = new(big.Int).Add(a.overrideNonce, big.NewInt(1))
	}
	return nonce
}

func (a *Account) Address() common.Address {
	return a.Account.Address
}

func (a *Account) Unlock(passphrase string) error {
	return a.ks.Unlock(a.Account, passphrase)
}

func (a *Account) SignTx(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	return a.ks.SignTx(a.Account, tx, chainID)
}

func (a *Account) NewTransactor(ctx context.Context, value *big.Int, gasPrice decimal.Decimal, gasLimit uint64) (*bind.TransactOpts, error) {
	t, err := bind.NewKeyStoreTransactor(a.ks, a.Account)
	if err != nil {
		return nil, err
	}
	t.Context = ctx
	t.Value = value
	t.From = a.Account.Address
	t.GasPrice = gasPrice.Shift(9).BigInt()
	t.GasLimit = gasLimit
	t.Nonce = a.NextNonceOverride() // Will be nil if no nonce override
	return t, err
}
