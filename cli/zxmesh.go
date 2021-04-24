package main

import (
	"context"
	"fmt"

	//"github.com/pkg/errors"
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/lazerdye/go-eth/zxmesh"
)

var (
	zxmeshCommand = kingpin.Command("zxmesh", "ZXMesh commands")
	zxmeshServer  = zxmeshCommand.Flag("zxmesh-addr", "Host or host:port for zxmesh graphql server").Envar("ZXMESH_ADDR").Required().String()

	zxmeshStatusCommand = zxmeshCommand.Command("status", "ZXMesh get status")
	zxmeshListCommand   = zxmeshCommand.Command("listen", "ZXMesh listen for orders")
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
	case "listen":
		return true, zxmeshListen(zxMesh)
	case "status":
		return true, zxmeshStatus(zxMesh)
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
	decoder := zxmesh.NewAssetDataDecoder()
	for msg := range c {
		for _, event := range msg {
			fmt.Printf("%+v\n", *event)
			if event.SignedOrder == nil {
				fmt.Printf("No order: %+v\n", *event)
				continue
			}
			order := event.SignedOrder
			fmt.Printf("%+v\n", *order)
			makerAssetData := order.MakerAssetData
			makerAssetName, err := decoder.GetName(makerAssetData)
			if err != nil {
				return err
			}
			fmt.Printf("MakerAssetName: %s\n", makerAssetName)
			if makerAssetName == "ERC20Token" {
				args, err := decoder.Decode(makerAssetData)
				if err != nil {
					return err
				}
				argList := args.([]interface{})
				type erc20 struct {
					Address common.Address
				}
				fmt.Printf("%+v\n", argList[0].(common.Address))
			}
		}
	}
	return nil
}
