package etherscan

import (
	"context"
	"net/url"
	"strconv"

	"github.com/pkg/errors"

	"github.com/lazerdye/go-eth/util"
)

const (
	etherscanApi = "https://api.etherscan.io/api"
)

type Client struct {
	apikey string

	httpClient *util.HttpClient
}

func NewClient(apikey string) *Client {
	return &Client{apikey: apikey, httpClient: util.NewHttpClient()}
}

type transaction struct {
	BlockNumber string `json:"blockNumber"`
}

type transactionResult struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Result  []transaction `json:"result"`
}

func (c *Client) NormalTransactionsByAddress(ctx context.Context, address string, page, offset int) ([]transaction, error) {
	var transactionResult transactionResult

	params := url.Values{}
	params.Set("module", "account")
	params.Set("action", "txlist")
	params.Set("address", address)
	params.Set("page", strconv.Itoa(page))
	params.Set("offset", strconv.Itoa(offset))
	params.Set("apikey", c.apikey)

	if err := c.httpClient.Get(ctx, etherscanApi, params, &transactionResult); err != nil {
		return nil, err
	}

	if transactionResult.Status != "1" {
		return nil, errors.Errorf("Error with transaction list: %+v", transactionResult)
	}

	return transactionResult.Result, nil
}
