package token

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"

	"gitlab.com/lazerdye/go-eth/wallet"
)

const (
	DGDName     = "DGD"
	DGDContract = "0xe0b7927c4af23765cb51314a0e0521a9645f0e2a"
	DGDDecimals = 9
)

var (
	tokenRepo map[string]tokenInfo
)

type tokenInfo struct {
	contract common.Address
	decimals int
}

func init() {
	tokenRepo = map[string]tokenInfo{
		DGDName: {common.HexToAddress(DGDContract), DGDDecimals},
	}
}

type Client struct {
	client   *ethclient.Client
	instance *Token
	info     tokenInfo
}

func NewClient(name string, client *ethclient.Client) (*Client, error) {
	token, ok := tokenRepo[name]
	if !ok {
		return nil, errors.Errorf("Token not registered: %s", name)
	}
	instance, err := NewToken(token.contract, client)
	if err != nil {
		return nil, err
	}
	return &Client{client, instance, token}, nil
}

func (c *Client) BalanceOf(ctx context.Context, addressHex string) (float64, error) {
	address := common.HexToAddress(addressHex)
	balance, err := c.instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return float64(0), err
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	rawValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(c.info.decimals)))

	f64Value, _ := rawValue.Float64()
	return f64Value, nil
}

func (c *Client) Transfer(sourceAccount *wallet.Account, destAccount string, amount float64, transmit bool) error {
	nonce, err := c.client.PendingNonceAt(context.Background(), sourceAccount.Account.Address)
	if err != nil {
		return err
	}
	toAddress := common.HexToAddress(destAccount)
	transferFnSignature := []byte("transfer(address,uint256)")

	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	log.Infof("Hash: %s", hexutil.Encode(methodID))
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	log.Infof("Padded hash: %s", hexutil.Encode(paddedAddress))

	fAmount := new(big.Float).Mul(big.NewFloat(amount), big.NewFloat(math.Pow10(c.info.decimals)))
	iAmount, _ := fAmount.Int(nil)
	paddedAmount := common.LeftPadBytes(iAmount.Bytes(), 32)
	log.Infof("Padded amount: %s", hexutil.Encode(paddedAmount))
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	//gasLimit, err := c.client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	To:   &toAddress,
	//	Data: data,
	//})
	//if err != nil {
	//	return err
	//}
	gasLimit := uint64(200000)
	log.Infof("Gas limit: %d", gasLimit)

	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	log.Infof("Gas price: %d", gasPrice)

	value := big.NewInt(0) // no ethereum transferred
	tx := types.NewTransaction(nonce, c.info.contract, value, gasLimit, gasPrice, data)
	log.Infof("Transaction: %+v", tx)

	chainID, err := c.client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("ChainID: %+v", chainID)

	txSigned, err := sourceAccount.SignTx(tx, chainID)
	if err != nil {
		return err
	}
	log.Infof("Signed transaction: %+v", txSigned)

	fmt.Printf("Transaction: %s\n", txSigned.Hash().Hex())

	if transmit {
		err = c.client.SendTransaction(context.Background(), txSigned)
		if err != nil {
			return err
		}
	}

	return nil
}
