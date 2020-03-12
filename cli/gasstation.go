package main

import (
	"context"
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/lazerdye/go-eth/gasstation"
)

var (
	gasstationcmd   = kingpin.Command("gasstation", "Eth Gas Station commands")
	gasstationspeed = gasstationcmd.Arg("speed", "Speed of transaction").Required().String()
)

func gasstationCommand(ctx context.Context, client *gasstation.Client) error {
	gas, speed, err := client.GasPrice(ctx, gasstation.Speed(*gasstationspeed))
	if err != nil {
		return err
	}
	fmt.Printf("Gas: %f\n", gas)
	fmt.Printf("Speed: %fm\n", speed)

	return nil
}
