package bancor

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/token2"
)

var (
	ContractRegistryContract = common.HexToAddress("0x52Ae12ABe5D8BD778BD5397F99cA900624CfADD4")
	EthAddress               = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
)

type Client struct {
	*client.Client

	contractRegistryInstance  *ContractRegistry
	converterRegistryInstance *ConverterRegistry
	networkInstance           *Network
}

func getContractByName(opts *bind.CallOpts, registry *ContractRegistry, contractName string) (common.Address, error) {
	b := make([]byte, 32)
	copy(b[:], []byte(contractName))
	hash := common.BytesToHash(b)
	return registry.AddressOf(opts, hash)
}

func NewClient(ctx context.Context, client *client.Client) (*Client, error) {
	contractRegistryInstance, err := NewContractRegistry(ContractRegistryContract, client)
	if err != nil {
		return nil, err
	}
	converterRegistryAddress, err := getContractByName(client.DefaultCallOpts(ctx), contractRegistryInstance, "BancorConverterRegistry")
	if err != nil {
		return nil, err
	}
	converterRegistryInstance, err := NewConverterRegistry(converterRegistryAddress, client)
	if err != nil {
		return nil, err
	}
	networkAddress, err := getContractByName(client.DefaultCallOpts(ctx), contractRegistryInstance, "BancorNetwork")
	if err != nil {
		return nil, err
	}
	networkInstance, err := NewNetwork(networkAddress, client)
	if err != nil {
		return nil, err
	}
	return &Client{client, contractRegistryInstance, converterRegistryInstance, networkInstance}, nil
}

func (c *Client) ToWei(reg *token2.Registry, address common.Address, amount decimal.Decimal) (*big.Int, error) {
	if address == EthAddress {
		return client.EthToWei(amount), nil
	}
	_, tokClient, err := reg.ByAddress(address)
	if err != nil {
		return nil, err
	}
	return tokClient.ToWei(amount), nil
}

func (c *Client) FromWei(reg *token2.Registry, address common.Address, amount *big.Int) (decimal.Decimal, error) {
	if address == EthAddress {
		return client.EthFromWei(amount), nil
	}
	_, tokClient, err := reg.ByAddress(address)
	if err != nil {
		return decimal.Zero, err
	}
	return tokClient.FromWei(amount), nil
}

func (c *Client) GetConvertibleTokens(ctx context.Context) ([]common.Address, error) {
	return c.converterRegistryInstance.GetConvertibleTokens(c.DefaultCallOpts(ctx))
}

func (c *Client) GetConversionPath(ctx context.Context, sourceToken, targetToken common.Address) ([]common.Address, error) {
	return c.networkInstance.ConversionPath(c.DefaultCallOpts(ctx), sourceToken, targetToken)
}

func (c *Client) RateByPath(ctx context.Context, path []common.Address, amount *big.Int) (*big.Int, error) {
	return c.networkInstance.RateByPath(c.DefaultCallOpts(ctx), path, amount)
}
