package maker

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasoracle"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/wallet"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

type collateralInfo struct {
	ContractAddress         common.Address
	CollateralTokenAddress  common.Address
	CollateralTokenDecimals uint8
}

var (
	// TODO: Take from chainlog.makerdao.com
	GemJoinContractAddresses = map[string]collateralInfo{
		"ETH-A": {
			ContractAddress:         common.HexToAddress("0x2F0b23f53734252Bda2277357e97e1517d6B042A"),
			CollateralTokenAddress:  common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
			CollateralTokenDecimals: 18,
		},
	}

	defaultJoinGasLimit = uint64(1000000)
	joinGasSpeed        = gasoracle.Fastest
)

type GemJoinClient struct {
	*client.Client
	address      common.Address
	token        *token2.Client
	joinGasSpeed uint64
	gemJoin      *GemJoin
}

func NewGemJoinClient(c *client.Client, ilk string) (*GemJoinClient, error) {
	info, ok := GemJoinContractAddresses[ilk]
	if !ok {
		return nil, errors.Errorf("Unsupported maker token: %s", ilk)
	}
	token, err := token2.NewClient(c, info.CollateralTokenAddress, info.CollateralTokenDecimals)
	if err != nil {
		return nil, err
	}
	gemJoin, err := NewGemJoin(info.ContractAddress, c)
	if err != nil {
		return nil, err
	}
	return &GemJoinClient{
		Client:       c,
		address:      info.ContractAddress,
		token:        token,
		joinGasSpeed: defaultJoinGasLimit,
		gemJoin:      gemJoin,
	}, nil
}

func (c *GemJoinClient) Join(ctx context.Context, account *wallet.Account, usr common.Address, amount decimal.Decimal) (*types.Transaction, error) {
	gasFeeCap, gasTipCap, err := c.GasPrice(ctx, joinGasSpeed)
	if err != nil {
		return nil, err
	}
	wad, err := c.token.ToWeiCapped(ctx, amount, account)
	if err != nil {
		return nil, err
	}
	opts, err := account.NewTransactor(ctx, nil, gasFeeCap, gasTipCap, c.joinGasSpeed)
	if err != nil {
		return nil, err
	}
	allowance, err := c.token.Allowance(ctx, account.Address(), c.address)
	if err != nil {
		return nil, err
	}
	if allowance.LessThan(amount) {
		return nil, errors.Errorf("Allowance for token %s on contract %s is too small: %s",
			c.token.Address, c.address, allowance)
	}
	log.Infof("Allowance: %+v", allowance)

	return c.gemJoin.Join(opts, usr, wad)
}

func (c *GemJoinClient) Exit(ctx context.Context, account *wallet.Account, usr common.Address, amount decimal.Decimal) (*types.Transaction, error) {
	gasFeeCap, gasTipCap, err := c.GasPrice(ctx, joinGasSpeed)
	if err != nil {
		return nil, err
	}
	wad := c.token.ToWei(amount)
	opts, err := account.NewTransactor(ctx, nil, gasFeeCap, gasTipCap, c.joinGasSpeed)
	if err != nil {
		return nil, err
	}

	return c.gemJoin.Exit(opts, usr, wad)
}
