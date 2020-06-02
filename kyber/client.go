package kyber

import (
	"context"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasstation"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/wallet"
	"github.com/shopspring/decimal"
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
	tradeGasLimit            = uint64(700000)
	KyberNetworkProxyAddress = common.HexToAddress(KyberNetworkProxyAddressString)
	EthereumAddress          = common.HexToAddress(EthereumAddressString)
	ethAmounts               = []float64{0.2, 0.5, 1.0, 10.0}

	tradeGasSpeed = gasstation.Fastest

	dnil = decimal.Decimal{}
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

func (c *Client) GetExpectedRate(ctx context.Context, source, dest *token2.Client, quantity decimal.Decimal) (decimal.Decimal, decimal.Decimal, error) {
	quantityInt := source.ToWei(quantity)
	log.Infof("GetExpectedRate input %s %s %s", source.Address.String(), dest.Address.String(), quantityInt.String())
	rate, err := c.instance.GetExpectedRate(&bind.CallOpts{Context: ctx}, source.Address, dest.Address, quantityInt)
	if err != nil {
		return dnil, dnil, err
	}

	// It's confusing, but the 'rate' ignores token decimals.
	expectedRate := client.EthFromWei(rate.ExpectedRate)
	slippageRate := client.EthFromWei(rate.SlippageRate)
	return expectedRate, slippageRate, nil
}

func (c *Client) SwapEtherToToken(ctx context.Context, account *wallet.Account, tok *token2.Client, amount decimal.Decimal, minRate decimal.Decimal) (*types.Transaction, error) {
	gasPrice, waitTime, err := c.GasPrice(ctx, tradeGasSpeed)
	if err != nil {
		return nil, err
	}
	log.Infof("Gas price: %f - wait time: %fs", gasPrice, waitTime)

	amountInt := client.EthToWei(amount)
	log.Infof("Amount: %s, gasPrice: %s, gasLimit: %s", amountInt.String(), gasPrice, tradeGasLimit)
	transactOpts, err := account.NewTransactor(ctx, amountInt, gasPrice, tradeGasLimit)
	if err != nil {
		return nil, err
	}
	log.Infof("Transact Opts: %+v", transactOpts)
	minRateInt := tok.ToWei(minRate)
	transaction, err := c.instance.SwapEtherToToken(transactOpts, tok.Address, minRateInt)
	if err != nil {
		return nil, err
	}
	return transaction, err
}

func (c *Client) SwapTokenToEther(ctx context.Context, account *wallet.Account, tok *token2.Client, amount decimal.Decimal, maxRate decimal.Decimal) (*types.Transaction, error) {
	gasPrice, _, err := c.GasPrice(ctx, tradeGasSpeed)
	if err != nil {
		return nil, err
	}
	transactOpts, err := account.NewTransactor(ctx, nil, gasPrice, tradeGasLimit)
	if err != nil {
		return nil, err
	}
	// TODO: Balance check
	//amountInt, err := tok.ToWeiCapped(amount, account.Address())
	//if err != nil {
	//	return nil, err
	//}

	// TODO: Is 18 right for rate?
	transaction, err := c.instance.SwapTokenToEther(transactOpts, tok.Address, tok.ToWei(amount), client.EthToWei(maxRate))
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
