package gasoracle

import (
	"context"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

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
