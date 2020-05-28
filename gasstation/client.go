package gasstation

import (
	"context"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/util"
)

type Speed string

const (
	ethgasstationApi = "https://ethgasstation.info/json/ethgasAPI.json"

	Fast    = Speed("fast")
	Fastest = Speed("fastest")
	SafeLow = Speed("safeLow")
	Average = Speed("average")
)

var (
	WaitValues = make(map[Speed]string)

	dnil = decimal.Decimal{}
)

func init() {
	WaitValues[Fast] = "fastWait"
	WaitValues[Fastest] = "fastestWait"
	WaitValues[Average] = "avgWait"
	WaitValues[SafeLow] = "safeLowWait"
}

type Client struct {
	httpClient *util.HttpClient
}

func NewClient() *Client {
	return &Client{httpClient: util.NewHttpClient()}
}

func (c *Client) GasPrice(ctx context.Context, speed Speed) (decimal.Decimal, float64, error) {
	var response map[string]interface{}

	if err := c.httpClient.Get(ctx, ethgasstationApi, nil, &response); err != nil {
		return dnil, 0, err
	}

	log.Debugf("Response: %s", response)

	waitValues, ok := WaitValues[speed]
	if !ok {
		return dnil, 0, errors.Errorf("Unknown speed value: %s", speed)
	}
	gas, ok := response[string(speed)].(float64)
	if !ok {
		return dnil, 0, errors.Errorf("Internal error getting gas value")
	}
	wait, ok := response[waitValues].(float64)
	if !ok {
		return dnil, 0, errors.Errorf("Internal error getting wait value")
	}
	gasVal := decimal.NewFromFloat(gas).Shift(-10)

	return gasVal, wait, nil
}
