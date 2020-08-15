package bancor

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/lazerdye/go-eth/client"
)

var (
	ContractRegistryContract = common.HexToAddress("0x52Ae12ABe5D8BD778BD5397F99cA900624CfADD4")
)

type Client struct {
	*client.Client

	contractRegistryInstance  *ContractRegistry
	converterRegistryInstance *ConverterRegistry
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
	return &Client{client, contractRegistryInstance, converterRegistryInstance}, nil
}

func (c *Client) GetConvertibleTokens(ctx context.Context) ([]common.Address, error) {
	return c.converterRegistryInstance.GetConvertibleTokens(c.DefaultCallOpts(ctx))
}
