package main

import (
	"context"
	"fmt"
	"math/big"

	//"github.com/pkg/errors"
	//"github.com/ethereum/go-ethereum/common"
	"gopkg.in/alecthomas/kingpin.v2"
	//log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/token2"
	"github.com/lazerdye/go-eth/zxmesh"
)

var (
	zxmeshCommand = kingpin.Command("zxmesh", "ZXMesh commands")
	zxmeshServer  = zxmeshCommand.Flag("zxmesh-addr", "Host or host:port for zxmesh graphql server").Envar("ZXMESH_ADDR").Required().String()

	zxmeshStatusCommand = zxmeshCommand.Command("status", "ZXMesh get status")
	zxmeshListCommand   = zxmeshCommand.Command("listen", "ZXMesh listen for orders")

	zxmeshToken2Command = clientToken2Command.Command("zxmesh", "ZXMesh Token Commands")
	zxmeshToken2Server  = zxmeshToken2Command.Flag("zxmesh-addr", "Host or host:port for zxmesh graphql server").Envar("ZXMESH_ADDR").Required().String()
	zxmeshToken2Orders  = zxmeshToken2Command.Command("orders", "ZXMesg Orders")
)

func zxmeshCommands(ctx context.Context, commands []string) (bool, error) {
	zxMesh, err := zxmesh.New(*zxmeshServer)
	if err != nil {
		return false, err
	}
	defer zxMesh.Close()
	if err := zxMesh.Connect(); err != nil {
		return false, err
	}
	switch commands[0] {
	case "status":
		return true, zxmeshStatus(zxMesh)
	case "listen":
		return true, zxmeshListen(zxMesh)
	}
	return false, nil
}

func zxmeshStatus(zx *zxmesh.Client) error {
	stats, err := zx.Stats()
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", stats)
	return nil
}

func zxmeshListen(zx *zxmesh.Client) error {
	c := make(chan []*zxmesh.OrderEvent)
	err := zx.SubscribeToOrders(c)
	if err != nil {
		return err
	}
	for msg := range c {
		for _, event := range msg {
			fmt.Printf("%+v\n", *event)
			if event.SignedOrder == nil {
				fmt.Printf("No order: %+v\n", *event)
				continue
			}
			order := event.SignedOrder
			fmt.Printf("%+v\n", *order)
			fmt.Printf("%s\n", order.MakerAssetData)
		}
	}
	return nil
}

func zxmeshToken2Commands(ctx context.Context, client *client.Client, reg *token2.Registry, commands []string) (bool, error) {
	zxMesh, err := zxmesh.New(*zxmeshToken2Server)
	if err != nil {
		return false, err
	}
	defer zxMesh.Close()
	if err := zxMesh.Connect(); err != nil {
		return false, err
	}
	switch commands[0] {
	case "orders":
		return true, zxmeshOrders(ctx, zxMesh, client, reg)
	}
	return false, nil
}

func zxmeshOrders(ctx context.Context, zx *zxmesh.Client, client *client.Client, reg *token2.Registry) error {
	c := make(chan []*zxmesh.SignedOrder)
	err := zx.Orders(c, 5)
	if err != nil {
		return err
	}
	for msg := range c {
		for _, order := range msg {
			//log.Infof("Order: %+v", order)
			fmt.Printf("== Hash %s\n", order.Hash.String())
			fmt.Printf("-- Maker %s:\n", order.MakerAddress.String())
			fmt.Printf("MakerAssetData: \t%s\n", order.MakerAssetData)
			fmt.Printf("MakerAssetAmount:\t%s\n", order.MakerAssetAmount.String())
			if order.MakerFee.Cmp(new(big.Int)) != 0 {
				fmt.Printf("MakerFee:\t%s\n", order.MakerFee.String())
				if order.MakerFeeAssetData != nil {
					fmt.Printf("MakerFeeAssetData:\t%s\n", order.MakerFeeAssetData)
				}
			}
			fmt.Printf("-- Taker %s:\n", order.TakerAddress.String())
			fmt.Printf("TakerAssetData: \t%s\n", order.TakerAssetData)
			fmt.Printf("TakerAssetAmount:\t%s\n", order.TakerAssetAmount.String())
			if order.TakerFee.Cmp(new(big.Int)) != 0 {
				fmt.Printf("TakerFee:\t%s\n", order.TakerFee.String())
				if order.TakerFeeAssetData != nil {
					fmt.Printf("TakerFeeAssetData:\t%s\n", order.TakerFeeAssetData)
				}
			}
			ercMakerAsset, okMaker := order.MakerAssetData.(zxmesh.ERC20AssetData)
			ercTakerAsset, okTaker := order.TakerAssetData.(zxmesh.ERC20AssetData)
			if okMaker && okTaker {
				makerTokenName, makerToken, err := reg.ByAddress(ercMakerAsset.Address)
				if err != nil {
					continue
				}
				makerTokenAmount := makerToken.FromWei(order.MakerAssetAmount)
				takerTokenName, takerToken, err := reg.ByAddress(ercTakerAsset.Address)
				if err != nil {
					continue
				}
				takerTokenAmount := takerToken.FromWei(order.TakerAssetAmount)
				fmt.Printf("Trade %s %s for %s %s\n", makerTokenAmount, makerTokenName, takerTokenAmount, takerTokenName)
			}
		}
	}
	return nil
}
