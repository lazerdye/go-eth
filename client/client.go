package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/gasoracle"
	"github.com/lazerdye/go-eth/wallet"
)

type Client struct {
	*ethclient.Client

	overrideBlockNo *big.Int
	gasoracle       gasoracle.GasOracle
}

const (
	ethDecimals = 18

	GasLimit         = uint64(7800000)
	TransferGasSpeed = gasoracle.Fastest // TODO: Make this configurable.
	BuyGasSpeed      = gasoracle.Fastest
	SellGasSpeed     = gasoracle.Fast
)

var (
	dnil = decimal.Decimal{}
)

func Dial(url string, gasoracle gasoracle.GasOracle) (*Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Client{client, nil, gasoracle}, nil
}

func (c *Client) GasPrice(ctx context.Context, speed gasoracle.GasSpeed) (decimal.Decimal, error) {
	return c.gasoracle(ctx, speed)
}

func (c *Client) BalanceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (decimal.Decimal, error) {
	balance, err := c.Client.BalanceAt(ctx, address, blockNumber)
	if err != nil {
		return dnil, err
	}

	return EthFromWei(balance), nil
}

func (c *Client) Transfer(ctx context.Context, sourceAccount *wallet.Account, destAddress common.Address, amount decimal.Decimal, transmit bool) (*types.Transaction, error) {
	// TODO: There's got to be a better way.
	var nonce uint64
	nonceInt := sourceAccount.NextNonceOverride()
	if nonceInt != nil {
		nonce = nonceInt.Uint64()
	} else {
		var err error
		nonce, err = c.Client.PendingNonceAt(ctx, sourceAccount.Account.Address)
		if err != nil {
			return nil, err
		}
	}

	valueInt := EthToWei(amount)

	gasLimit, err := c.Client.EstimateGas(ctx, ethereum.CallMsg{
		To: &destAddress,
	})
	if err != nil {
		return nil, err
	}
	log.Infof("Gas limit: %d", gasLimit)

	gasPrice, err := c.GasPrice(ctx, TransferGasSpeed)
	if err != nil {
		return nil, err
	}
	log.Infof("Gas price: %d", gasPrice)

	tx := types.NewTransaction(nonce, destAddress, valueInt, gasLimit, gasPrice.Shift(9).BigInt(), nil)

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

func (c *Client) SetBlockNumberOverride(block *big.Int) {
	c.overrideBlockNo = block
}

func (c *Client) BlockNumberOverride() *big.Int {
	return c.overrideBlockNo
}

func (c *Client) DefaultCallOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{Context: ctx, BlockNumber: c.overrideBlockNo}
}

func (c *Client) FilterTransferLogs(ctx context.Context, fromBlockNumber *big.Int, toBlockNumber *big.Int) error {
	log.Infof("XXX FilterTransferLogs %s -> %s", fromBlockNumber, toBlockNumber)
	logs, err := c.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: fromBlockNumber,
		ToBlock:   toBlockNumber,
	})
	if err != nil {
		return err
	}
	if len(logs) == 0 {
		log.Infof("No dice")
	}
	for _, l := range logs {
		log.Infof("Log: %+v", l)
		log.Infof("Address: %s", l.Address.String())
		for _, t := range l.Topics {
			log.Infof("%s", t.String())
		}
	}
	return nil
}

func EthToWei(amount decimal.Decimal) *big.Int {
	return amount.Shift(ethDecimals).BigInt()
}

func EthFromWei(iValue *big.Int) decimal.Decimal {
	return decimal.NewFromBigInt(iValue, -int32(ethDecimals))
}
