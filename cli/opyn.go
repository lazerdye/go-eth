package main

import (
	"context"
	"fmt"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/opyn"
	"github.com/lazerdye/go-eth/token2"
)

var (
	clientOpynCommand              = clientToken2Command.Command("opyn", "Opyn operations")
	clientOpynListOptionsContracts = clientOpynCommand.Command("list-options-contracts", "List options contracts")
)

func opynCommands(ctx context.Context, client *client.Client, reg *token2.Registry, commands []string) (bool, error) {
	opynClient, err := opyn.NewClient(client)
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "list-options-contracts":
		return true, opynListOptionsContracts(ctx, opynClient)
	}
	return false, nil
}

func opynListOptionsContracts(ctx context.Context, opynClient *opyn.Client) error {
	numberOfOptionsContracts, err := opynClient.GetNumberOfOptionsContracts(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Number of options contracts: %+v\n", numberOfOptionsContracts)
	return nil
}
