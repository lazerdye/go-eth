package etherscan

import (
	"context"
	"math/big"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/util"
)

const (
	etherscanApi      = "https://api.etherscan.io/api"
	etherscanWaitTime = 500 * time.Millisecond
)

var (
	lastReq *time.Time
	limitMu sync.Mutex
)

func checkRate() {
	limitMu.Lock()
	defer limitMu.Unlock()

	now := time.Now()
	if lastReq != nil {
		nextReq := lastReq.Add(etherscanWaitTime)
		if now.Before(nextReq) {
			sleepTime := nextReq.Sub(now)
			log.Infof("Etherscan wait %s", sleepTime)
			time.Sleep(sleepTime)
		}
	}
	lastReq = &now
}

type Client struct {
	apikey string

	httpClient *util.HttpClient
}

func NewClient(apikey string) *Client {
	return &Client{apikey: apikey, httpClient: util.NewHttpClient()}
}

type blockNumberResult struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

func (c *Client) BlockNumberByTimestamp(ctx context.Context, timestamp string, closest string) (string, error) {
	var blockNumberResult blockNumberResult

	params := url.Values{}
	params.Set("module", "block")
	params.Set("action", "getblocknobytime")
	params.Set("timestamp", timestamp)
	params.Set("closest", closest)
	params.Set("apikey", c.apikey)

	if err := c.httpClient.Get(ctx, etherscanApi, params, &blockNumberResult); err != nil {
		return "", err
	}

	if blockNumberResult.Status != "1" {
		return "", errors.Errorf("Error with BlockNumberByTimestamp: %+v", blockNumberResult)
	}

	return blockNumberResult.Result, nil
}

type Transaction struct {
	BlockNumber       string `json:"blockNumber"`
	Timestamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	TxReceiptStatus   string `json:"txreceipt_status"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
}

type transactionResult struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Result  []Transaction `json:"result"`
}

func (c *Client) NormalTransactionsByAddress(ctx context.Context, address string, startBlock, endBlock *string, page, offset int, sort string) ([]Transaction, error) {
	checkRate()

	var transactionResult transactionResult

	params := url.Values{}
	params.Set("module", "account")
	params.Set("action", "txlist")
	params.Set("address", address)
	if startBlock != nil {
		params.Set("startblock", *startBlock)
	}
	if endBlock != nil {
		params.Set("endblock", *endBlock)
	}
	params.Set("page", strconv.Itoa(page))
	params.Set("offset", strconv.Itoa(offset))
	params.Set("sort", sort)
	params.Set("apikey", c.apikey)

	if err := c.httpClient.Get(ctx, etherscanApi, params, &transactionResult); err != nil {
		return nil, err
	}

	if transactionResult.Status != "1" {
		if transactionResult.Message != "No transactions found" {
			return nil, errors.Errorf("Error with TransactionsByAddress: %+v", transactionResult)
		}
	}

	return transactionResult.Result, nil
}

func (c *Client) TokenTransactionsByAddress(ctx context.Context, address, contractaddress string, startblock, endblock *string, page, offset int, sort string) ([]Transaction, error) {
	checkRate()

	var transactionResult transactionResult

	params := url.Values{}
	params.Set("module", "account")
	params.Set("action", "tokentx")
	params.Set("address", address)
	if contractaddress != "" {
		params.Set("contractaddress", contractaddress)
	}
	if startblock != nil {
		params.Set("startblock", *startblock)
	}
	if endblock != nil {
		params.Set("endblock", *endblock)
	}
	params.Set("page", strconv.Itoa(page))
	params.Set("offset", strconv.Itoa(offset))
	params.Set("sort", sort)
	params.Set("apikey", c.apikey)

	if err := c.httpClient.Get(ctx, etherscanApi, params, &transactionResult); err != nil {
		return nil, err
	}

	if transactionResult.Status != "1" {
		if transactionResult.Message != "No transactions found" {
			return nil, errors.Errorf("Error with TokenTransactionsByAddress: %+v", transactionResult)
		}
	}

	return transactionResult.Result, nil
}

type contractExecutionStatus struct {
	IsError        string `json:"isError"`
	ErrDescription string `json:"errDescription"`
}

type contractExecutionResult struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Result  contractExecutionStatus `json:"result"`
}

func (c *Client) ContractExecutionStatus(ctx context.Context, transaction string) (bool, string, error) {
	checkRate()
	var contractExecutionResult contractExecutionResult

	params := url.Values{}
	params.Set("module", "transaction")
	params.Set("action", "getstatus")
	params.Set("txhash", transaction)
	params.Set("apikey", c.apikey)

	if err := c.httpClient.Get(ctx, etherscanApi, params, &contractExecutionResult); err != nil {
		return false, "", err
	}

	if contractExecutionResult.Status != "1" {
		return false, "", errors.Errorf("Error with transaction list: %+v", contractExecutionResult)
	}

	return contractExecutionResult.Result.IsError != "0", contractExecutionResult.Result.ErrDescription, nil
}

type GasOracle struct {
	LastBlock       string
	SafeGasPrice    string
	ProposeGasPrice string
	FastGasPrice    string
}

type gasOracleResponse struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Result  GasOracle `json:"result"`
}

func (c *Client) GasOracle(ctx context.Context) (*GasOracle, error) {
	checkRate()
	var gasOracleResponse gasOracleResponse

	params := url.Values{}
	params.Set("module", "gastracker")
	params.Set("action", "gasoracle")
	params.Set("apikey", c.apikey)

	if err := c.httpClient.Get(ctx, etherscanApi, params, &gasOracleResponse); err != nil {
		return nil, err
	}

	if gasOracleResponse.Status != "1" {
		return nil, errors.Errorf("Error with transaction list: %+v", gasOracleResponse)
	}

	return &gasOracleResponse.Result, nil
}

func CalculateFee(t *Transaction) *big.Int {
	gasPriceInt, _ := new(big.Int).SetString(t.GasPrice, 10)
	gasUsedInt, _ := new(big.Int).SetString(t.GasUsed, 10)
	return new(big.Int).Mul(gasPriceInt, gasUsedInt)
}

func (c *Client) EventLogs(ctx context.Context, fromBlock, toBlock, address string, topics []string) error {
	// NOTE: This doesn't seem to work...
	checkRate()
	var eventLogsResponse interface{}

	params := url.Values{}
	params.Set("module", "logs")
	params.Set("action", "getLogs")
	params.Set("apikey", c.apikey)
	params.Set("fromBlock", fromBlock)
	params.Set("toBlock", toBlock)
	params.Set("address", address)
	if len(topics) > 0 {
		return errors.New("Topics not yet supported")
	}
	log.Infof("%+v", params)
	if err := c.httpClient.Get(ctx, etherscanApi, params, &eventLogsResponse); err != nil {
		return err
	}
	log.Infof("RESP: %+v", eventLogsResponse)
	return errors.New("Not Implemented")
}
