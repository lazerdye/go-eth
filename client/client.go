package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/gasstation"
	"github.com/lazerdye/go-eth/wallet"
)

type Client struct {
	*ethclient.Client

	gasstation *gasstation.Client
}

const (
	ethDecimals = 18

	GasLimit         = uint64(7800000)
	TransferGasSpeed = gasstation.Fastest // TODO: Make this configurable.
	BuyGasSpeed      = gasstation.Fastest
	SellGasSpeed     = gasstation.Fast
)

var (
	dnil = decimal.Decimal{}
)

func Dial(url string, gasstation *gasstation.Client) (*Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Client{client, gasstation}, nil
}

func (c *Client) GasPrice(ctx context.Context, speed gasstation.Speed) (decimal.Decimal, float64, error) {
	return c.gasstation.GasPrice(ctx, speed)
}

func (c *Client) BalanceOf(ctx context.Context, address common.Address) (decimal.Decimal, error) {
	balance, err := c.Client.BalanceAt(ctx, address, nil)
	if err != nil {
		return dnil, err
	}

	return EthFromWei(balance), nil
}

func (c *Client) Transfer(ctx context.Context, sourceAccount *wallet.Account, dest string, amount decimal.Decimal, transmit bool) (*types.Transaction, error) {
	// TODO: Just use the erc20 contract.
	nonce, err := c.Client.PendingNonceAt(ctx, sourceAccount.Account.Address)
	if err != nil {
		return nil, err
	}

	destAddress := common.HexToAddress(dest)

	valueInt := EthToWei(amount)

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

func EthToWei(amount decimal.Decimal) *big.Int {
	return amount.Shift(ethDecimals).BigInt()
}

func EthFromWei(iValue *big.Int) decimal.Decimal {
	return decimal.NewFromBigInt(iValue, -int32(ethDecimals))
}
