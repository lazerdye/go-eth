package kyber

import (
	"context"
	"math"
	"math/big"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasstation"
	"github.com/lazerdye/go-eth/token"
	"github.com/lazerdye/go-eth/wallet"
	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	KyberNetworkProxyAddressString = "0x818E6FECD516Ecc3849DAf6845e3EC868087B755"
	EthereumAddressString          = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	EthereumDecimals               = 18
)

var (
	tradeGasLimit            = uint64(400000)
	KyberNetworkProxyAddress = common.HexToAddress(KyberNetworkProxyAddressString)
	EthereumAddress          = common.HexToAddress(EthereumAddressString)
	ethAmounts               = []float64{0.2, 0.5, 1.0, 10.0}

	tradeGasSpeed = gasstation.Fast
)

type Client struct {
	*client.Client

	instance *Kyber
}

func NewClient(client *client.Client) (*Client, error) {
	instance, err := NewKyber(KyberNetworkProxyAddress, client)
	if err != nil {
		return nil, err
	}
	return &Client{
		Client:   client,
		instance: instance,
	}, nil
}

func (c *Client) GetExpectedRate(ctx context.Context, source, dest *token.Token, quantity *big.Float) (*big.Float, *big.Float, error) {
	quantityInt := source.ToGwei(quantity)
	log.Infof("GetExpectedRate input %s %s %s", source.Contract().String(), dest.Contract().String(), quantityInt.String())
	rate, err := c.instance.GetExpectedRate(&bind.CallOpts{Context: ctx}, source.Contract(), dest.Contract(), quantityInt)
	if err != nil {
		return nil, nil, err
	}

	expectedRate := new(big.Float).Quo(new(big.Float).SetInt(rate.ExpectedRate), big.NewFloat(math.Pow10(18)))
	slippageRate := new(big.Float).Quo(new(big.Float).SetInt(rate.SlippageRate), big.NewFloat(math.Pow10(18)))
	return expectedRate, slippageRate, nil
}

func (c *Client) SwapEtherToToken(ctx context.Context, account *wallet.Account, token common.Address, amount *big.Float, minRate *big.Float) (*types.Transaction, error) {
	gasPrice, waitTime, err := c.GasPrice(ctx, tradeGasSpeed)
	if err != nil {
		return nil, err
	}
	log.Infof("Gas price: %f - wait time: %fs", gasPrice, waitTime)

	amountInt, _ := new(big.Float).Mul(amount, big.NewFloat(math.Pow10(18))).Int(nil)
	log.Infof("Amount: %s, gasPrice: %s, gasLimit: %s", amountInt.String(), gasPrice, tradeGasLimit)
	transactOpts, err := account.NewTransactor(ctx, amountInt, gasPrice, tradeGasLimit)
	if err != nil {
		return nil, err
	}
	log.Infof("Transact Opts: %+v", transactOpts)
	minRateInt, _ := new(big.Float).Mul(minRate, big.NewFloat(math.Pow10(18))).Int(nil)
	transaction, err := c.instance.SwapEtherToToken(transactOpts, token, minRateInt)
	if err != nil {
		return nil, err
	}
	return transaction, err
}

func (c *Client) SwapTokenToEther(ctx context.Context, account *wallet.Account, token common.Address, amount *big.Float, maxRate *big.Float) (*types.Transaction, error) {
	gasPrice, _, err := c.GasPrice(ctx, tradeGasSpeed)
	if err != nil {
		return nil, err
	}
	transactOpts, err := account.NewTransactor(ctx, big.NewInt(0), gasPrice, tradeGasLimit)
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
	defer fit.Close()

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
