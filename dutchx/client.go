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
