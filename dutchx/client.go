package dutchx

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/big"

	//"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	DutchXProxy  = "0xb9812e2fa995ec53b5b6df34d21f9304762c5497"
	DutchXMaster = "0x2bae491b065032a76be1db9e9ecf5738afae203e"
)

type Client struct {
	c        *ethclient.Client
	instance *DutchExchange
}

func NewClient(client *ethclient.Client) (*Client, error) {
	instance, err := NewDutchExchange(common.HexToAddress(DutchXProxy), client)
	if err != nil {
		return nil, err
	}
	return &Client{client, instance}, nil
}

func (c *Client) GetAuctionIndex(address1, address2 common.Address) (int64, error) {
	value, err := c.instance.GetAuctionIndex(&bind.CallOpts{}, address1, address2)
	if err != nil {
		return int64(0), err
	}
	log.Infof("Value: %+v", value)
	return value.Int64(), nil
}

type AuctionPriceInfo struct {
	Price    big.Float
	InvPrice big.Float
}

func newAuctionPriceInfo(price struct {
	Num *big.Int
	Den *big.Int
}) *AuctionPriceInfo {
	var priceInfo AuctionPriceInfo
	if price.Den.Int64() == int64(0) {
		return &priceInfo
	}
	n := new(big.Float).SetInt(price.Num)
	d := new(big.Float).SetInt(price.Den)

	priceInfo.Price.Quo(n, d)
	priceInfo.InvPrice.Quo(new(big.Float).SetInt64(1), &priceInfo.Price)

	return &priceInfo
}

func (c *Client) GetCurrentAuctionPrice(address1, address2 common.Address) (*AuctionPriceInfo, error) {
	auctionIndex, err := c.instance.GetAuctionIndex(&bind.CallOpts{}, address1, address2)
	if err != nil {
		return nil, err
	}

	price, err := c.instance.GetCurrentAuctionPrice(&bind.CallOpts{}, address1, address2, auctionIndex)
	if err != nil {
		return nil, err
	}
	log.Infof("Price for auction %+v: %+v", auctionIndex, price)

	return newAuctionPriceInfo(price), nil
}

func (c *Client) GetPriceOfTokenInLastAuction(address common.Address) (*AuctionPriceInfo, error) {
	price, err := c.instance.GetPriceOfTokenInLastAuction(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}
	return newAuctionPriceInfo(price), nil
}

//	start, err := c.instance.GetAuctionStart(&bind.CallOpts{}, address1, address2)
//	if err != nil {
//		return err
//	}
//	log.Infof("Value: %+v", start)
//
//	bvol, err := c.instance.BuyVolumes(&bind.CallOpts{}, address1, address2)
//	if err != nil {
//		return err
//	}
//	log.Infof("BuyVolumes: %+v", bvol)
//
//	svolc, err := c.instance.SellVolumesCurrent(&bind.CallOpts{}, address1, address2)
//	if err != nil {
//		return err
//	}
//	log.Infof("SellVolumesCurrent: %+v", svolc)
//
//	svoln, err := c.instance.SellVolumesNext(&bind.CallOpts{}, address1, address2)
//	if err != nil {
//		return err
//	}
//	log.Infof("SellVolumesNext: %+v", svoln)

func (c *Client) Listen() error {
	//query := ethereum.FilterQuery{
	//	Addresses: []common.Address{
	//		common.HexToAddress(DutchXProxy),
	//		common.HexToAddress(DutchXMaster),
	//	},
	//}
	approvalChan := make(chan *DutchExchangeApproval)
	approvalSub, err := c.instance.WatchApproval(&bind.WatchOpts{}, approvalChan, nil)
	if err != nil {
		return err
	}
	auctionClearedChan := make(chan *DutchExchangeAuctionCleared)
	auctionClearedSub, err := c.instance.WatchAuctionCleared(&bind.WatchOpts{}, auctionClearedChan, nil, nil, nil)
	if err != nil {
		return err
	}
	auctionStartScheduledChan := make(chan *DutchExchangeAuctionStartScheduled)
	auctionStartScheduledSub, err := c.instance.WatchAuctionStartScheduled(&bind.WatchOpts{}, auctionStartScheduledChan, nil, nil, nil)
	if err != nil {
		return err
	}
	feeChan := make(chan *DutchExchangeFee)
	feeSub, err := c.instance.WatchFee(&bind.WatchOpts{}, feeChan, nil, nil, nil)
	if err != nil {
		return err
	}
	newBuyOrderChan := make(chan *DutchExchangeNewBuyOrder)
	newBuyOrderSub, err := c.instance.WatchNewBuyOrder(&bind.WatchOpts{}, newBuyOrderChan, nil, nil, nil)
	if err != nil {
		return err
	}
	newBuyerFundsClaimChan := make(chan *DutchExchangeNewBuyerFundsClaim)
	newBuyerFundsClaimSub, err := c.instance.WatchNewBuyerFundsClaim(&bind.WatchOpts{}, newBuyerFundsClaimChan, nil, nil, nil)
	if err != nil {
		return err
	}
	newDepositChan := make(chan *DutchExchangeNewDeposit)
	newDepositSub, err := c.instance.WatchNewDeposit(&bind.WatchOpts{}, newDepositChan, nil)
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-approvalSub.Err():
			return err
		case approval := <-approvalChan:
			fmt.Printf("Approval: %+v\n", *approval)
		case err := <-auctionClearedSub.Err():
			return err
		case auctionCleared := <-auctionClearedChan:
			fmt.Printf("AuctionCleared: %s %s %+v\n", auctionCleared.SellToken.Hex(), auctionCleared.BuyToken.Hex(), *auctionCleared)
		case err := <-auctionStartScheduledSub.Err():
			return err
		case auctionStartScheduled := <-auctionStartScheduledChan:
			fmt.Printf("AuctionStartScheduled: %s %s %+v\n", auctionStartScheduled.SellToken.Hex(), auctionStartScheduled.BuyToken.Hex(), *auctionStartScheduled)
		case err := <-feeSub.Err():
			return err
		case fee := <-feeChan:
			fmt.Printf("Fee: %+v\n", *fee)
		case err := <-newBuyOrderSub.Err():
			return err
		case newBuyOrder := <-newBuyOrderChan:
			fmt.Printf("NewBuyOrder: %s %s %+v\n", newBuyOrder.SellToken.Hex(), newBuyOrder.BuyToken.Hex(), *newBuyOrder)
		case err := <-newBuyerFundsClaimSub.Err():
			return err
		case newBuyerFundsClaim := <-newBuyerFundsClaimChan:
			fmt.Printf("NewBuyerFundsClaim: %s %s %+v\n", newBuyerFundsClaim.SellToken.Hex(), newBuyerFundsClaim.BuyToken.Hex(), *newBuyerFundsClaim)
		case err := <-newDepositSub.Err():
			return err
		case newDeposit := <-newDepositChan:
			fmt.Printf("NewDeposit: %+v\n", *newDeposit)
		}
	}
}
