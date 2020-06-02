package main

import (
	"context"

	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasstation"
)

var (
	clientCmd    = kingpin.Command("client", "Client operations")
	clientServer = clientCmd.Flag("server", "URL of the server to connect to").Envar("ETHEREUM_SERVER").Required().String()
)

func newClient() (*client.Client, error) {
	if *clientServer == "" {
		return nil, errors.New("clientServer parameter required")
	}
	return client.Dial(*clientServer, gasstation.NewClient())
}

func clientCommands(ctx context.Context, commands []string) (bool, error) {
	client, err := newClient()
	if err != nil {
		return false, err
	}
	switch commands[0] {
	case "token2":
		return token2Commands(ctx, client, commands[1:])
	}
	return false, nil
}
