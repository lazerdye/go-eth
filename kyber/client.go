package kyber

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasoracle"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/wallet"
)

const (
	KyberNetworkProxyAddressString = "0x818E6FECD516Ecc3849DAf6845e3EC868087B755"
	EthereumAddressString          = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	EthereumDecimals               = 18
)

var (
	DefaultTradeGasLimit     = uint64(1000000)
	KyberNetworkProxyAddress = common.HexToAddress(KyberNetworkProxyAddressString)
	EthereumAddress          = common.HexToAddress(EthereumAddressString)

	tradeGasSpeed = gasoracle.Fastest

	dnil = decimal.Decimal{}
)

type Client struct {
	*client.Client

	instance      *Kyber
	tradeGasLimit uint64
}

func NewClient(client *client.Client) (*Client, error) {
	instance, err := NewKyber(KyberNetworkProxyAddress, client)
	if err != nil {
		return nil, err
	}
	return &Client{
		Client:        client,
		instance:      instance,
		tradeGasLimit: DefaultTradeGasLimit,
	}, nil
}

func (c *Client) SetTradeGasLimit(tradeGasLimit uint64) {
	c.tradeGasLimit = tradeGasLimit
}

func (c *Client) EstimateTradeFee(ctx context.Context) (decimal.Decimal, error) {
	// Get gas price for trade.
	gasFeeCap, gasTipCap, err := c.GasPrice(ctx, tradeGasSpeed)
	if err != nil {
		return decimal.Zero, err
	}
	log.Infof("GasFeeCap: %d, GasTipCap: %d", gasFeeCap, gasTipCap)
	// Calculate gas price.
	gasPrice := gasFeeCap.Add(gasTipCap)
	// Multiply by gas limit.
	fee := gasPrice.Shift(9).Mul(decimal.NewFromInt(int64(c.tradeGasLimit)))
	log.Infof("Trade fee: %s", fee)

	return fee.Shift(-18), nil
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
	gasFeeCap, gasTipCap, err := c.GasPrice(ctx, tradeGasSpeed)
	if err != nil {
		return nil, err
	}
	log.Infof("GasFeeCap: %f, GasTipCap: %f", gasFeeCap, gasTipCap)

	amountInt := client.EthToWei(amount)
	log.Infof("Amount: %s, gasFeeCap: %s, gasTipCap: %s, gasLimit: %s", amountInt.String(), gasFeeCap, gasTipCap, c.tradeGasLimit)
	transactOpts, err := account.NewTransactor(ctx, amountInt, gasFeeCap, gasTipCap, c.tradeGasLimit)
	if err != nil {
		return nil, err
	}
	log.Infof("Transact Opts: %+v", transactOpts)
	// It's confusing, but rate always uses 18 decimals.
	transaction, err := c.instance.SwapEtherToToken(transactOpts, tok.Address, client.EthToWei(minRate))
	if err != nil {
		return nil, err
	}
	return transaction, err
}

func (c *Client) SwapTokenToEther(ctx context.Context, account *wallet.Account, tok *token2.Client, amount decimal.Decimal, maxRate decimal.Decimal) (*types.Transaction, error) {
	gasFeeCap, gasTipCap, err := c.GasPrice(ctx, tradeGasSpeed)
	if err != nil {
		return nil, err
	}
	transactOpts, err := account.NewTransactor(ctx, nil, gasFeeCap, gasTipCap, c.tradeGasLimit)
	if err != nil {
		return nil, err
	}
	amountInt, err := tok.ToWeiCapped(ctx, amount, account)
	if err != nil {
		return nil, err
	}

	// We always use 18 for the rate.
	transaction, err := c.instance.SwapTokenToEther(transactOpts, tok.Address, amountInt, client.EthToWei(maxRate))
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
