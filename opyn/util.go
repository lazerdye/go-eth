package opyn

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

const (
	CallType    = "call"
	PutType     = "put"
	UnknownType = "unknown"
)

var (
	USDCContract = common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
)

type ContractInfo struct {
	Address     common.Address
	Name        string
	Currency    common.Address
	Type        string
	StrikePrice decimal.Decimal
	Expiry      *big.Int
}

func (t *OTokenClient) ContractInfo(ctx context.Context) (*ContractInfo, error) {
	name, err := t.Name(ctx)
	if err != nil {
		return nil, err
	}
	expiry, err := t.Expiry(ctx)
	if err != nil {
		return nil, err
	}
	decimals, err := t.Decimals(ctx)
	if err != nil {
		return nil, err
	}
	collateral, err := t.Collateral(ctx)
	if err != nil {
		return nil, err
	}
	underlying, err := t.Underlying(ctx)
	if err != nil {
		return nil, err
	}
	strikePrice, err := t.StrikePrice(ctx)
	if err != nil {
		return nil, err
	}
	strikePrice = strikePrice.Shift(int32(decimals))

	var typ string
	var currency common.Address
	if collateral == USDCContract {
		typ = PutType
		currency = underlying
	} else if underlying == USDCContract {
		typ = CallType
		currency = collateral
		one := decimal.New(1, 0)
		strikePrice = one.Div(strikePrice)
	} else {
		typ = UnknownType
		currency = collateral
	}
	return &ContractInfo{
		t.Contract,
		name,
		currency,
		typ,
		strikePrice,
		expiry,
	}, nil
}
