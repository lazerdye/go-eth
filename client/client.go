package client

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/dutchx"
	"github.com/lazerdye/go-eth/kyber"
	"github.com/lazerdye/go-eth/token"
	"github.com/lazerdye/go-eth/wallet"
)

type Client struct {
	*ethclient.Client
}

const (
	ethDigits = 18
)

var (
	GasLimit = big.NewInt(5000000000)
	GasPrice = uint64(7800000)
)

func Dial(url string) (*Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Client{Client: client}, nil
}

func (c *Client) Token(tokenName string) (*token.Client, error) {
	t, ok := token.DefaultRegistry.ByName(tokenName)
	if !ok {
		return nil, errors.Errorf("Unknown token: %s", tokenName)
	}
	return token.NewClient(t, c.Client)
}

func (c *Client) Dutchx() (*dutchx.Client, error) {
	return dutchx.NewClient(c.Client)
}

func (c *Client) Kyber() (*kyber.Client, error) {
	return kyber.NewClient(c.Client)
}

func (c *Client) BalanceOf(ctx context.Context, address common.Address) (*big.Float, error) {
	balance, err := c.Client.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, err
	}

	balanceFloat := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(ethDigits)))
	return balanceFloat, nil
}

func (c *Client) Transfer(ctx context.Context, sourceAccount *wallet.Account, dest string, amount *big.Float, transmit bool) (*types.Transaction, error) {
	// TODO: Just use the erc20 contract.
	nonce, err := c.Client.PendingNonceAt(ctx, sourceAccount.Account.Address)
	if err != nil {
		return nil, err
	}

	destAddress := common.HexToAddress(dest)

	value := new(big.Float).Mul(amount, big.NewFloat(math.Pow10(18)))
	valueInt, _ := value.Int(nil)

	gasLimit, err := c.Client.EstimateGas(ctx, ethereum.CallMsg{
		To: &destAddress,
	})
	if err != nil {
		return nil, err
	}
	log.Infof("Gas limit: %d", gasLimit)

	gasPrice, err := c.Client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	log.Infof("Gas price: %d", gasPrice)

	tx := types.NewTransaction(nonce, destAddress, valueInt, gasLimit, gasPrice, nil)

	chainID, err := c.Client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	log.Infof("ChainID: %+v", chainID)

	txSigned, err := sourceAccount.SignTx(tx, chainID)
	if err != nil {
		return nil, err
	}

	if transmit {
		err = c.Client.SendTransaction(ctx, txSigned)
		if err != nil {
			return nil, err
		}
	}
	return txSigned, nil
}
