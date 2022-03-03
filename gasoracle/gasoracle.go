package gasoracle

import (
	"context"

	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/etherscan"
)

type GasSpeed string

const (
	Low     = GasSpeed("low")
	Average = GasSpeed("average")
	Fast    = GasSpeed("fast")
	Fastest = GasSpeed("fastest")
)

type GasOracleResponse struct {
	GasFeeCap decimal.Decimal
	GasTipCap map[GasSpeed]decimal.Decimal
}

type GasOracle func(ctx context.Context) (*GasOracleResponse, error)

func GasOracleFromEtherscan(etherscan *etherscan.Client) GasOracle {
	return func(ctx context.Context) (*GasOracleResponse, error) {
		etherGasOracle, err := etherscan.GasOracle(ctx)
		if err != nil {
			return nil, err
		}

		return &GasOracleResponse{
			GasFeeCap: etherGasOracle.SuggestBaseFee,
			GasTipCap: map[GasSpeed]decimal.Decimal{
				Low:     etherGasOracle.SafeGasPrice.Sub(etherGasOracle.SuggestBaseFee),
				Average: etherGasOracle.ProposeGasPrice.Sub(etherGasOracle.SuggestBaseFee),
				Fast:    etherGasOracle.FastGasPrice.Sub(etherGasOracle.SuggestBaseFee),
				Fastest: etherGasOracle.FastGasPrice.Sub(etherGasOracle.SuggestBaseFee),
			},
		}, nil
	}
}
