package opyn

import (
	"context"
	"strings"

	"github.com/lazerdye/go-eth/util"
	log "github.com/sirupsen/logrus"
)

const (
	graphRoot = "https://api.thegraph.com/subgraphs/name/aparnakr/opyn"
)

type Graph struct {
	http util.HttpClient
}

type optionsContract struct {
	Address string `json:"address"`
	// TODO: More fields
}

type vaultOpenedAction struct {
	Id              string          `json:"id"`
	OptionsContract optionsContract `json:"optionsContract"`
	Owner           string          `json:"owner"`
	// TODO: More fields
}

type vaultsResponse struct {
	Data struct {
		VaultOpenedActions []vaultOpenedAction `json:"vaultOpenedActions"`
	} `json:"data"`
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) OpenedVaults(ctx context.Context, contract string) ([]vaultOpenedAction, error) {
	vaultsOpenedQuery := `
query vaultOpenedActions($limit: Int!, $contract: String) {
  vaultOpenedActions(first: $limit, where: { optionsContract: $contract }) {
    optionsContract {
      address
    },
    owner
  }
}
`
	body := make(map[string]interface{})
	body["query"] = vaultsOpenedQuery
	variables := make(map[string]interface{})
	variables["limit"] = 50
	variables["contract"] = strings.ToLower(contract)
	log.Infof("%+v", variables)
	body["variables"] = variables

	var result vaultsResponse
	//var result interface{}

	err := g.http.Post(ctx, graphRoot, nil, nil, body, &result)
	if err != nil {
		return nil, err
	}

	log.Infof("XXX Result %+v", result)

	return result.Data.VaultOpenedActions, nil
}
