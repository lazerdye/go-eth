package main

import (
	//"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"gitlab.com/lazerdye/go-eth/token"
)

func main() {
	//client, err := ethclient.Dial("http://10.4.0.45:31282")
	client, err := ethclient.Dial("ws://10.4.0.45:31744")
	if err != nil {
		log.Fatal(err)
	}

	dgdAddress := common.HexToAddress("0xe0b7927c4af23765cb51314a0e0521a9645f0e2a")
	instance, err := token.NewToken(dgdAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x2a563f94d39966d6d24fa08e84633948b92ec3d9")
	balance, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(9)))
	fmt.Println(ethValue)
}
