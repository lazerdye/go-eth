package oneinchv3

import (
	"context"
	"math/big"
	"net/url"

	//"github.com/pkg/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/util"
)

const (
	oneInchApi = "https://api.1inch.exchange/v3.0/1"
)

var (
	ethTokenAddress = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
)

type Client struct {
	util.HttpClient
}

func NewClient() *Client {
	return &Client{}
}

type TokenInfo struct {
	Symbol   string         `json:"symbol"`
	Name     string         `json:"name"`
	Decimals int8           `json:"decimals"`
	Address  common.Address `json:"address"`
	LogoURI  string         `json:"logoURI"`
}

type TokenResponse struct {
	Tokens map[common.Address]TokenInfo `json:"tokens"`
}

func (c *Client) Tokens(ctx context.Context) (map[common.Address]TokenInfo, error) {
	var tokens TokenResponse

	uri := oneInchApi + "/tokens"
	if err := c.Get(ctx, uri, nil, &tokens); err != nil {
		return nil, err
	}
	return tokens.Tokens, nil
}

type QuoteResponse struct {
	FromToken       TokenInfo   `json:"fromToken"`
	FromTokenAmount string      `json:"fromTokenAmount"`
	ToToken         TokenInfo   `json:"toToken"`
	ToTokenAmount   string      `json:"toTokenAmount"`
	Protocols       interface{} `json:"protocols"`
	EstimatedGas    int64       `json:"estimatedgas"`
}

func (c *Client) Quote(ctx context.Context, fromToken *token2.Client, toToken *token2.Client, amount decimal.Decimal) (*QuoteResponse, decimal.Decimal, error) {
	var response QuoteResponse

	var fromTokenAddress, toTokenAddress string
	var fromAmount string
	if fromToken != nil {
		fromTokenAddress = fromToken.Address.String()
		fromAmount = fromToken.ToWei(amount).String()
	} else {
		fromTokenAddress = ethTokenAddress.String()
		fromAmount = client.EthToWei(amount).String()
	}
	if toToken != nil {
		toTokenAddress = toToken.Address.String()
	} else {
		toTokenAddress = ethTokenAddress.String()
	}

	uri := oneInchApi + "/quote"
	params := url.Values{}
	params.Set("fromTokenAddress", fromTokenAddress)
	params.Set("toTokenAddress", toTokenAddress)
	params.Set("amount", fromAmount)
	if err := c.Get(ctx, uri, params, &response); err != nil {
		return nil, decimal.Zero, err
	}
	var toAmount decimal.Decimal
	toTokenAmount, _ := new(big.Int).SetString(response.ToTokenAmount, 10)
	if toToken != nil {
		toAmount = toToken.FromWei(toTokenAmount)
	} else {
		toAmount = client.EthFromWei(toTokenAmount)
	}
	return &response, toAmount, nil
}
