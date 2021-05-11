package zeroex

import (
	"context"
	"math/big"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/wallet"
	"github.com/lazerdye/go-eth/zeroex/ether_token"
	"github.com/lazerdye/go-eth/zeroex/exchange"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	EtherTokenAddress = common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	ExchangeAddress   = common.HexToAddress("0x61935cbdd02287b511119ddb11aeb42f1593b7ef")

	gasLimit = uint64(300000)
)

type Client struct {
	*client.Client

	etherInstance    *ether_token.EtherToken
	exchangeInstance *exchange.Exchange
}

func NewClient(tokenClient *client.Client) (*Client, error) {
	etherInstance, err := ether_token.NewEtherToken(EtherTokenAddress, tokenClient)
	if err != nil {
		return nil, err
	}
	exchangeInstance, err := exchange.NewExchange(ExchangeAddress, tokenClient)
	if err != nil {
		return nil, err
	}
	return &Client{
		Client:           tokenClient,
		etherInstance:    etherInstance,
		exchangeInstance: exchangeInstance,
	}, nil
}

func (c *Client) EstimateTradeFee(ctx context.Context) (decimal.Decimal, error) {
	// Get gas price for trade.
	gasPrice, err := c.GasPrice(ctx, client.SellGasSpeed)
	if err != nil {
		return decimal.Zero, err
	}
	log.Infof("Gas price: %s", gasPrice)
	// Multiply by gas limit.
	fee := gasPrice.Shift(9).Mul(decimal.NewFromInt(int64(gasLimit)))
	log.Infof("Trade fee: %s", fee)

	return fee.Shift(-18), nil
}

func (c *Client) EtherTokenDeposit(ctx context.Context, account *wallet.Account, amount decimal.Decimal) (*types.Transaction, error) {
	gasPrice, err := c.GasPrice(ctx, client.TransferGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := account.NewTransactor(ctx, client.EthToWei(amount), gasPrice, client.GasLimit)
	if err != nil {
		return nil, err
	}

	return c.etherInstance.Deposit(opts)
}

func (c *Client) EtherTokenBalanceOf(ctx context.Context, account *wallet.Account) (decimal.Decimal, error) {
	opts := &bind.CallOpts{Context: ctx}
	amountInt, err := c.etherInstance.BalanceOf(opts, account.Address())
	if err != nil {
		return decimal.Zero, err
	}
	return client.EthFromWei(amountInt), nil
}

func (c *Client) FillOrder(ctx context.Context, account *wallet.Account, order exchange.LibOrderOrder, amount *big.Int, signature []byte) (*types.Transaction, error) {
	gasPrice, err := c.GasPrice(ctx, client.SellGasSpeed)
	if err != nil {
		return nil, err
	}
	transactOpts, err := account.NewTransactor(ctx, nil, gasPrice, gasLimit)
	if err != nil {
		return nil, err
	}
	log.Infof("Order: %+v amount: %f signature: %x", order, amount, signature)
	return c.exchangeInstance.FillOrder(transactOpts, order, amount, signature)
}

func (c *Client) BatchFillOrders(ctx context.Context, account *wallet.Account, orders []exchange.LibOrderOrder, amounts []*big.Int, signatures [][]byte) (*types.Transaction, error) {
	gasPrice, err := c.GasPrice(ctx, client.SellGasSpeed)
	if err != nil {
		return nil, err
	}
	transactOpts, err := account.NewTransactor(ctx, nil, gasPrice, gasLimit)
	if err != nil {
		return nil, err
	}
	return c.exchangeInstance.BatchFillOrders(transactOpts, orders, amounts, signatures)
}

func (c *Client) FillOrKillOrder(ctx context.Context, account *wallet.Account, order exchange.LibOrderOrder, amount *big.Int, signature []byte) (*types.Transaction, error) {
	gasPrice, err := c.GasPrice(ctx, client.SellGasSpeed)
	if err != nil {
		return nil, err
	}
	transactOpts, err := account.NewTransactor(ctx, nil, gasPrice, gasLimit)
	if err != nil {
		return nil, err
	}
	log.Infof("Order: %+v amount: %f signature: %x", order, amount, signature)
	return c.exchangeInstance.FillOrKillOrder(transactOpts, order, amount, signature)
}

func (c *Client) BatchFillOrKillOrders(ctx context.Context, account *wallet.Account, orders []exchange.LibOrderOrder, amounts []*big.Int, signatures [][]byte) (*types.Transaction, error) {
	gasPrice, err := c.GasPrice(ctx, client.SellGasSpeed)
	if err != nil {
		return nil, err
	}
	transactOpts, err := account.NewTransactor(ctx, nil, gasPrice, gasLimit)
	if err != nil {
		return nil, err
	}
	return c.exchangeInstance.BatchFillOrKillOrders(transactOpts, orders, amounts, signatures)
}

func (c *Client) GetFillEvents(ctx context.Context, hashes [][32]byte, startBlock uint64, endBlock *uint64) ([]exchange.ExchangeFill, error) {
	filterOpts := &bind.FilterOpts{
		Context: ctx,
		Start:   startBlock,
		End:     endBlock,
	}
	fit, err := c.exchangeInstance.FilterFill(filterOpts, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	defer fit.Close()

	var trades []exchange.ExchangeFill
	for {
		if !fit.Next() {
			break
		}
		log.Infof("Trade: %+v", fit.Event)
		trades = append(trades, *fit.Event)
	}

	return trades, nil
}
