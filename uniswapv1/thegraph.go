package uniswapv1

import (
	"context"

	"github.com/lazerdye/go-eth/util"
	//"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

const (
	graphRoot = "https://api.thegraph.com/subgraphs/name/graphprotocol/uniswap"
)

type Graph struct {
	http util.HttpClient
}

type exchange struct {
	Id             string          `json:"id"`
	TokenName      string          `json:"tokenName"`
	TokenAddress   string          `json:"tokenAddress"`
	TokenSymbol    string          `json:"tokenSymbol"`
	TokenDecimals  int64           `json:"tokenDecimals"`
	TradeVolumeEth decimal.Decimal `json:"tradeVolumeEth"`
	// TODO: More fields
}

type exchangesResponse struct {
	Data struct {
		Exchanges []exchange `json:"exchanges"`
	} `json:"data"`
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) ExchangesByVolume(ctx context.Context) ([]exchange, error) {
	exchangesByVolumeQuery := `
query exchangesByVolume($limit: Int!) {
  exchanges(orderBy: tradeVolumeEth, orderDirection: desc, first: $limit) {
    id, tokenName, tokenAddress, tokenSymbol, tokenName, tokenDecimals, tradeVolumeEth
  }
}
`
	var result exchangesResponse
	body := make(map[string]interface{})
	body["query"] = exchangesByVolumeQuery
	variables := make(map[string]interface{})
	variables["limit"] = 100
	body["variables"] = variables

	err := g.http.Post(ctx, graphRoot, nil, nil, body, &result)
	if err != nil {
		return nil, err
	}

	log.Infof("XXX Result %+v", result)

	return result.Data.Exchanges, nil
}
