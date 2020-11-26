package eth2

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	//"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasstation"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/wallet"
)

const (
	depositGasSpeed = gasstation.Fast
	depositGasLimit = 0
)

type Client struct {
	*client.Client

	deposit *Deposit
}

func NewClient(client *client.Client, reg *token2.Registry) (*Client, error) {
	depositToken, err := reg.ByName("_eth2_deposit")
	if err != nil {
		return nil, err
	}
	deposit, err := NewDeposit(depositToken.Address, client)
	if err != nil {
		return nil, err
	}

	return &Client{client, deposit}, nil
}

type deposit struct {
	PubKey              string          `json:"pubkey"`
	WithdrawCredentials string          `json:"withdrawal_credentials"`
	Amount              decimal.Decimal `json:"amount"`
	Signature           string          `json:"signature"`
	DepositDataRoot     string          `json:"deposit_data_root"`
}

func (c *Client) Deposit(ctx context.Context, account *wallet.Account, depositFile string) ([]*types.Transaction, error) {

	gasPrice, _, err := c.GasPrice(ctx, depositGasSpeed)
	if err != nil {
		return nil, err
	}

	depFile, err := os.Open(depositFile)
	if err != nil {
		return nil, err
	}
	defer depFile.Close()

	var deposits []deposit

	depBytes, err := ioutil.ReadAll(depFile)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(depBytes, &deposits); err != nil {
		return nil, err
	}

	txs := make([]*types.Transaction, len(deposits))
	for i, deposit := range deposits {
		pubkey := common.Hex2Bytes(deposit.PubKey)
		withdrawCredentials := common.Hex2Bytes(deposit.WithdrawCredentials)
		signature := common.Hex2Bytes(deposit.Signature)
		depositDataRoot := common.HexToHash(deposit.DepositDataRoot)

		fmt.Printf("%d: %x / %x / %x / %x\n", i, pubkey, withdrawCredentials, signature, depositDataRoot)

		opts, err := account.NewTransactor(ctx, client.EthToWei(decimal.NewFromInt(32)), gasPrice, depositGasLimit)
		if err != nil {
			return txs, err
		}
		txs[i], err = c.deposit.Deposit(opts, pubkey, withdrawCredentials, signature, depositDataRoot)
		if err != nil {
			return txs, err
		}
	}

	return txs, nil
}
