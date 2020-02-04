package client

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/dutchx"
	"github.com/lazerdye/go-eth/kyber"
	"github.com/lazerdye/go-eth/token"
	"github.com/lazerdye/go-eth/wallet"
)

type Client struct {
	c *ethclient.Client
}

const (
	ethDigits = 18
)

func Dial(url string) (*Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func (c *Client) Token(tokenName string) (*token.Client, error) {
	return token.NewClient(tokenName, c.c)
}

func (c *Client) Dutchx() (*dutchx.Client, error) {
	return dutchx.NewClient(c.c)
}

func (c *Client) Kyber() (*kyber.Client, error) {
	return kyber.NewClient(c.c)
}

func (c *Client) BalanceOf(ctx context.Context, address common.Address) (*big.Float, error) {
	balance, err := c.c.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, err
	}

	balanceFloat := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(ethDigits)))
	return balanceFloat, nil
}

func (c *Client) Transfer(ctx context.Context, sourceAccount *wallet.Account, dest string, amount *big.Float, transmit bool) (*types.Transaction, error) {
	nonce, err := c.c.PendingNonceAt(ctx, sourceAccount.Account.Address)
	if err != nil {
		return nil, err
	}

	destAddress := common.HexToAddress(dest)

	value := new(big.Float).Mul(amount, big.NewFloat(math.Pow10(18)))
	valueInt, _ := value.Int(nil)

	gasLimit, err := c.c.EstimateGas(ctx, ethereum.CallMsg{
		To: &destAddress,
	})
	if err != nil {
		return nil, err
	}
	log.Infof("Gas limit: %d", gasLimit)

	gasPrice, err := c.c.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	log.Infof("Gas price: %d", gasPrice)

	tx := types.NewTransaction(nonce, destAddress, valueInt, gasLimit, gasPrice, nil)

	chainID, err := c.c.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	log.Infof("ChainID: %+v", chainID)

	txSigned, err := sourceAccount.SignTx(tx, chainID)
	if err != nil {
		return nil, err
	}

	if transmit {
		err = c.c.SendTransaction(ctx, txSigned)
		if err != nil {
			return nil, err
		}
	}
	return txSigned, nil
}
