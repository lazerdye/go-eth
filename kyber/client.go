package kyber

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/lazerdye/go-eth/token"
	"github.com/lazerdye/go-eth/wallet"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	KyberNetworkProxy = "0x818E6FECD516Ecc3849DAf6845e3EC868087B755"
	EthereumAddress   = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	EthereumDecimals  = 18
)

var (
	gasLimit              = big.NewInt(5000000000)
	gasPrice              = uint64(7800000)
	EthereumCommonAddress = common.HexToAddress(EthereumAddress)
	ethAmounts            = []float64{0.2, 0.5, 1.0, 10.0}
)

type Client struct {
	c        *ethclient.Client
	instance *Kyber
}

func NewClient(client *ethclient.Client) (*Client, error) {
	instance, err := NewKyber(common.HexToAddress(KyberNetworkProxy), client)
	if err != nil {
		return nil, err
	}
	return &Client{
		c:        client,
		instance: instance,
	}, nil
}

func (c *Client) GetExpectedRate(ctx context.Context, source, dest common.Address, quantity *big.Float) (*big.Float, *big.Float, error) {
	var sourceDecimals int
	var destDecimals int
	var ok bool
	if source.Hex() == EthereumCommonAddress.Hex() {
		sourceDecimals = EthereumDecimals
	} else {
		_, sourceDecimals, ok = token.TokenByAddress(source)
		if !ok {
			return nil, nil, errors.Errorf("Error looking up token: %s", source.String())
		}
	}
	if dest.Hex() == EthereumCommonAddress.Hex() {
		destDecimals = EthereumDecimals
	} else {
		_, destDecimals, ok = token.TokenByAddress(dest)
		if !ok {
			return nil, nil, errors.Errorf("Error looking up token: %s", dest.String())
		}
	}

	quantityInt, _ := new(big.Float).Mul(quantity, big.NewFloat(math.Pow10(sourceDecimals))).Int(nil)
	rate, err := c.instance.GetExpectedRate(&bind.CallOpts{Context: ctx}, source, dest, quantityInt)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("Raw: %+v", rate)

	expectedRate := new(big.Float).Quo(new(big.Float).SetInt(rate.ExpectedRate), big.NewFloat(math.Pow10(destDecimals)))
	slippageRate := new(big.Float).Quo(new(big.Float).SetInt(rate.SlippageRate), big.NewFloat(math.Pow10(destDecimals)))
	return expectedRate, slippageRate, nil
}

func (c *Client) SwapEtherToToken(ctx context.Context, account *wallet.Account, token common.Address, amount *big.Float, minRate *big.Float) (*types.Transaction, error) {
	amountInt, _ := new(big.Float).Mul(amount, big.NewFloat(math.Pow10(18))).Int(nil)
	transactOpts, err := account.NewTransactor(ctx, amountInt, gasLimit, gasPrice)
	if err != nil {
		return nil, err
	}
	minRateInt, _ := new(big.Float).Mul(minRate, big.NewFloat(math.Pow10(18))).Int(nil)
	transaction, err := c.instance.SwapEtherToToken(transactOpts, token, minRateInt)
	if err != nil {
		return nil, err
	}
	return transaction, err
}

func (c *Client) SwapTokenToEther(ctx context.Context, account *wallet.Account, token common.Address, amount *big.Float, maxRate *big.Float) (*types.Transaction, error) {
	transactOpts, err := account.NewTransactor(ctx, big.NewInt(0), gasLimit, gasPrice)
	if err != nil {
		return nil, err
	}
	// TODO: What about when decimals != 18?
	amountInt, _ := new(big.Float).Mul(amount, big.NewFloat(math.Pow10(18))).Int(nil)
	maxRateInt, _ := new(big.Float).Mul(maxRate, big.NewFloat(math.Pow10(18))).Int(nil)
	transaction, err := c.instance.SwapTokenToEther(transactOpts, token, amountInt, maxRateInt)
	if err != nil {
		return nil, err
	}
	return transaction, err
}

func (c *Client) GetTradeEvents(ctx context.Context, addresses []common.Address, startBlock uint64, endBlock *uint64) ([]KyberExecuteTrade, error) {
	filterOpts := &bind.FilterOpts{
		Context: ctx,
		Start:   startBlock,
		End:     endBlock,
	}
	fit, err := c.instance.FilterExecuteTrade(filterOpts, addresses)
	if err != nil {
		return nil, err
	}

	var trades []KyberExecuteTrade
	for {
		if !fit.Next() {
			break
		}
		log.Infof("Trade: %+v", fit.Event)
		trades = append(trades, *fit.Event)
	}

	return trades, nil

}
