package gasoracle

import (
	"context"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/etherscan"
	"github.com/lazerdye/go-eth/gasstation"
)

type GasSpeed string

const (
	Low     = GasSpeed("low")
	Average = GasSpeed("average")
	Fast    = GasSpeed("fast")
	Fastest = GasSpeed("fastest")
)

type GasOracle func(ctx context.Context, speed GasSpeed) (decimal.Decimal, error)

func GasOracleFromGasStation(station *gasstation.Client) GasOracle {
	return func(ctx context.Context, speed GasSpeed) (decimal.Decimal, error) {
		var stationSpeed gasstation.Speed
		switch speed {
		case Low:
			stationSpeed = gasstation.SafeLow
		case Average:
			stationSpeed = gasstation.Average
		case Fast:
			stationSpeed = gasstation.Fast
		case Fastest:
			stationSpeed = gasstation.Fastest
		default:
			return decimal.Zero, errors.Errorf("Unknown gas speed: %s", speed)
		}
		amount, _, err := station.GasPrice(ctx, stationSpeed)
		return amount, err
	}
}

func GasOracleFromEtherscan(etherscan *etherscan.Client) GasOracle {
	return func(ctx context.Context, speed GasSpeed) (decimal.Decimal, error) {
		etherGasOracle, err := etherscan.GasOracle(ctx)
		if err != nil {
			return decimal.Zero, err
		}
		var decimalGasOracle decimal.Decimal
		switch speed {
		case Low:
			decimalGasOracle, err = decimal.NewFromString(etherGasOracle.SafeGasPrice)
			if err != nil {
				return decimal.Zero, err
			}
		case Average:
			decimalGasOracle, err = decimal.NewFromString(etherGasOracle.ProposeGasPrice)
			if err != nil {
				return decimal.Zero, err
			}
		case Fast:
			decimalGasOracle, err = decimal.NewFromString(etherGasOracle.FastGasPrice)
			if err != nil {
				return decimal.Zero, err
			}
		case Fastest:
			// There is no 'fastest' with etherscan...
			decimalGasOracle, err = decimal.NewFromString(etherGasOracle.FastGasPrice)
			if err != nil {
				return decimal.Zero, err
			}
		default:
			return decimal.Zero, errors.Errorf("Unknown gas speed: %s", speed)
		}
		return decimalGasOracle, nil
	}
}
