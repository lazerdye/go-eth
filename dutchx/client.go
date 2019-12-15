package dutchx

import (
	log "github.com/sirupsen/logrus"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	DutchXProxy  = "0xb9812e2fa995ec53b5b6df34d21f9304762c5497"
	DutchXMaster = "0x2bae491b065032a76be1db9e9ecf5738afae203e"
)

type Client struct {
	instance *DutchExchange
}

func NewClient(client *ethclient.Client) (*Client, error) {
	instance, err := NewDutchExchange(common.HexToAddress(DutchXProxy), client)
	if err != nil {
		return nil, err
	}
	return &Client{instance}, nil
}

func (c *Client) GetAuctionIndex(address1Str, address2Str string) error {
	address1 := common.HexToAddress(address1Str)
	address2 := common.HexToAddress(address2Str)
	value, err := c.instance.GetAuctionIndex(&bind.CallOpts{}, address1, address2)
	if err != nil {
		return err
	}
	log.Infof("Value: %+v", value)

	//price, err := c.instance.GetCurrentAuctionPrice(&bind.CallOpts{}, address1, address2, big.NewInt(517))
	price, err := c.instance.GetCurrentAuctionPrice(&bind.CallOpts{}, address1, address2, value)
	if err != nil {
		return err
	}
	log.Infof("Value: %+v", price)

	if price.Den.Int64() != int64(0) {
		n := new(big.Float).SetInt(price.Num)
		d := new(big.Float).SetInt(price.Den)
		r := new(big.Float).Quo(n, d)
		rz := new(big.Float).Quo(new(big.Float).SetInt64(1), r)
		log.Infof("R: %+v", r.String())
		log.Infof("RZ: %+v", rz.String())
	}

	start, err := c.instance.GetAuctionStart(&bind.CallOpts{}, address1, address2)
	if err != nil {
		return err
	}
	log.Infof("Value: %+v", start)

	bvol, err := c.instance.BuyVolumes(&bind.CallOpts{}, address1, address2)
	if err != nil {
		return err
	}
	log.Infof("BuyVolumes: %+v", bvol)

	svolc, err := c.instance.SellVolumesCurrent(&bind.CallOpts{}, address1, address2)
	if err != nil {
		return err
	}
	log.Infof("SellVolumesCurrent: %+v", svolc)

	svoln, err := c.instance.SellVolumesNext(&bind.CallOpts{}, address1, address2)
	if err != nil {
		return err
	}
	log.Infof("SellVolumesNext: %+v", svoln)

	return nil
}
