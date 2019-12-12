package token

import (
    "context"
    "math"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

const (
	DGDName     = "DGD"
	DGDContract = "0xe0b7927c4af23765cb51314a0e0521a9645f0e2a"
	DGDDecimals = 9
)

var (
	tokenRepo map[string]tokenInfo
)

type tokenInfo struct {
	contract common.Address
	decimals int
}

func init() {
	tokenRepo = map[string]tokenInfo{
		DGDName: {common.HexToAddress(DGDContract), DGDDecimals},
	}
}

type Client struct {
	client *ethclient.Client
    instance *Token
	info   tokenInfo
}

func NewClient(name string, client *ethclient.Client) (*Client, error) {
	token, ok := tokenRepo[name]
	if !ok {
		return nil, errors.Errorf("Token not registered: %s", name)
	}
    instance, err := NewToken(token.contract, client)
    if err != nil {
        return nil, err
    }
	return &Client{client, instance, token}, nil
}

func (c *Client) BalanceOf(ctx context.Context, addressHex string) (float64, error) {
    address := common.HexToAddress(addressHex)
    balance, err := c.instance.BalanceOf(&bind.CallOpts{}, address)
    if err != nil {
        return float64(0), err
    }

    fbalance := new(big.Float)
    fbalance.SetString(balance.String())
    rawValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(c.info.decimals)))

    f64Value, _ := rawValue.Float64()
    return f64Value, nil
}
