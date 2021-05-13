package zxmesh

import (
	"encoding/json"
	"math/big"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var (
	heartbeatTimeout = time.Second * 20
)

type Client struct {
	c             *websocket.Conn
	listeners     map[string]chan []byte // TODO: MUTEX
	connectionAck bool                   // Was the connection acked?
}

func New(hostport string) (*Client, error) {
	u := url.URL{Scheme: "ws", Host: hostport, Path: "/graphql"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	z := &Client{c: c, listeners: make(map[string]chan []byte)}
	go z.listen()
	return z, nil
}

func (z *Client) Connect() error {
	c := make(chan []byte)
	z.Listen("_connection_ack", c)

	if err := z.c.WriteMessage(websocket.TextMessage, []byte("{\"type\":\"connection_init\"}")); err != nil {
		return err
	}

	log.Infof("Waiting for connection_ack")
	<-c

	return nil
}

func (z *Client) Listen(id string, c chan []byte) {
	z.listeners[id] = c
}

type Message struct {
	Type    string           `json:"type"`
	Id      *string          `json:"id"`
	Payload *json.RawMessage `json:"payload"`
}

func (z *Client) Start(id string, message []byte) error {
	sendMsg := Message{
		Type:    "start",
		Id:      &id,
		Payload: (*json.RawMessage)(&message),
	}
	bytes, err := json.Marshal(sendMsg)
	if err != nil {
		return err
	}
	return z.c.WriteMessage(websocket.TextMessage, bytes)
}

func (z *Client) listen() {
	log.Infof("Starting listen loop...")
	for {
		z.c.SetReadDeadline(time.Now().Add(heartbeatTimeout))
		_, message, err := z.c.ReadMessage()
		if err != nil {
			// TODO: Handle this more cleanly.
			for _, c := range z.listeners {
				close(c)
			}
			log.Fatalf("error: %v", err)
			return
		}
		log.Infof("Message: %+v", string(message))
		var responseMessage Message
		if err := json.Unmarshal(message, &responseMessage); err != nil {
			log.Warnf("Invalid message: %s", string(message))
			continue
		}
		switch responseMessage.Type {
		case "connection_ack":
			log.Info("Connection acked")
			z.handleConnectionAck()
			z.connectionAck = true
		case "ka":
			log.Info("Heartbeat")
		case "data":
			z.handleData(responseMessage)
		case "complete":
			z.handleComplete(responseMessage)
		default:
			log.Warnf("Unknown type: %s", responseMessage.Type)
		}
	}
}

func (z *Client) handleConnectionAck() {
	c, ok := z.listeners["_connection_ack"]
	if !ok {
		log.Warnf("Received connection ack, but nobody listening")
		return
	}
	c <- []byte("")
}

func (z *Client) handleData(msg Message) {
	if msg.Id == nil || msg.Payload == nil {
		log.Warnf("Missing ID or Payload: %+v", msg)
		return
	}
	c, ok := z.listeners[*msg.Id]
	if !ok {
		log.Warnf("Unknown listener: %s", *msg.Id)
		return
	}
	c <- ([]byte)(*msg.Payload)
}

func (z *Client) handleComplete(msg Message) {
	if msg.Id == nil {
		log.Warnf("Missing ID: %+v", msg)
		return
	}
	c, ok := z.listeners[*msg.Id]
	if !ok {
		log.Warnf("Unknown listener: %s", *msg.Id)
		return
	}
	close(c)
	delete(z.listeners, *msg.Id)
}

func (z *Client) Close() {
	z.c.Close()
}

type StatsResponse struct {
	Data struct {
		Stats struct {
			Version         string `json:"version"`
			PubSubTopic     string `json:"pubSubTopic"`
			Rendezvous      string `json:"rendezvous"`
			PeerID          string `json:"peerID"`
			EthereumChainID int    `json:"ethereumChainID"`
			LatestBlock     struct {
				Number string `json:"number"`
				Hash   string `json:"hash"`
			} `json:"latestBlock"`
			NumPeers                          int    `json:"numPeers"`
			NumOrders                         int    `json:"numOrders"`
			NumOrdersIncludingRemoved         int    `json:"numOrdersIncludingRemoved"`
			StartOfCurrentUTCDay              string `json:"startOfCurrentUTCDay"`
			EthRPCRequestsSentInCurrentUTCDay int    `json:"ethRCPRequestsSentInCurrentUTCDay"`
			EthRPCRateLimitExpiredRequests    int    `json:"ethRPCRateLimitExpiredRequests"`
			MaxExpirationTime                 string `json:"maxExpirationTime"`
		} `json:"stats"`
	} `json:"data"`
}

type UnknownAssetData struct {
	AssetName string
}

func (decoder AssetDataDecoder) DecodeToInterface(assetData []byte) (interface{}, error) {
	if len(assetData) == 0 {
		return nil, nil
	}
	assetName, err := decoder.GetName(assetData)
	if err != nil {
		return "", err
	}
	switch assetName {
	case "ERC20Token":
		args, err := decoder.Decode(assetData)
		if err != nil {
			return "", err
		}
		argList := args.([]interface{})
		return ERC20AssetData{
			Address: argList[0].(common.Address),
		}, nil
	default:
		return UnknownAssetData{
			AssetName: assetName,
		}, nil
	}
}

type Query struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables,omitempty"`
}

const statsQuery = `
query {
        stats{
                version
                pubSubTopic
                rendezvous
                peerID
                ethereumChainID
                latestBlock {
                    number
                    hash
                }
                numPeers
                numOrders
                numOrdersIncludingRemoved
                startOfCurrentUTCDay
                ethRPCRequestsSentInCurrentUTCDay
                ethRPCRateLimitExpiredRequests
                maxExpirationTime
        }
}`

func (z *Client) Stats() (*StatsResponse, error) {
	id := "status"
	c := make(chan []byte)
	z.Listen(id, c)
	q, err := json.Marshal(Query{Query: statsQuery})
	if err != nil {
		return nil, err
	}
	if err := z.Start(id, q); err != nil {
		return nil, err
	}
	resp := <-c
	var response StatsResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

const subscribeQuery = `
subscription {
    orderEvents {
            timestamp
            endState
            order {
		hash
                makerAddress
                takerAddress
                feeRecipientAddress
                senderAddress
                makerAssetAmount
                takerAssetAmount
                makerFee
                takerFee
                expirationTimeSeconds
                salt
                makerAssetData
                takerAssetData
                makerFeeAssetData
                takerFeeAssetData
                hash
                fillableTakerAssetAmount
                signature
            }
    }
}
`

type SignedOrderPayload struct {
	Hash                     string         `json:"hash"`
	SenderAddress            common.Address `json:"senderAddress"`
	MakerAddress             common.Address `json:"makerAddress"`
	MakerAssetData           string         `json:"makerAssetData"`
	MakerAssetAmount         string         `json:"makerAssetAmount"`
	MakerFee                 string         `json:"makerFee"`
	MakerFeeAssetData        string         `json:"makerFeeAssetData"`
	TakerAddress             common.Address `json:"takerAddress"`
	TakerAssetData           string         `json:"takerAssetData"`
	TakerAssetAmount         string         `json:"takerAssetAmount"`
	TakerFee                 string         `json:"takerFee"`
	TakerFeeAssetData        string         `json:"takerFeeAssetData"`
	FeeRecipientAddress      common.Address `json:"feeRecipientAddress"`
	ExpirationTimeSeconds    string         `json:"expirationTimeSeconds"`
	Salt                     string         `json:"salt"`
	Signature                string         `json:"signature"`
	FillableTakerAssetAmount *string        `json:"fillableTakerAssetAmount,omitempty"`
}

type OrderEventPayload struct {
	Timestamp   string              `json:"timestamp"`
	EndState    string              `json:"endState"`
	SignedOrder *SignedOrderPayload `json:"order"`
}

type SignedOrder struct {
	Hash                     common.Hash
	SenderAddress            common.Address
	MakerAddress             common.Address
	MakerAssetData           interface{}
	MakerAssetAmount         *big.Int
	MakerFee                 *big.Int
	MakerFeeAssetData        interface{}
	TakerAddress             common.Address
	TakerAssetData           interface{}
	TakerAssetAmount         *big.Int
	TakerFee                 *big.Int
	TakerFeeAssetData        interface{}
	FeeRecipientAddress      common.Address
	ExpirationTimeSeconds    *big.Int
	Salt                     *big.Int
	Signature                []byte
	FillableTakerAssetAmount *big.Int
}

type OrderEvent struct {
	Timestamp   string
	EndState    string
	SignedOrder *SignedOrder
}

type OrderSubscriptionResponse struct {
	Data struct {
		OrderEvents []*OrderEventPayload `json:"orderEvents"`
	} `json:"data"`
}

func orderPayloadToSignedOrder(decoder *AssetDataDecoder, orderPayload *SignedOrderPayload) *SignedOrder {
	if orderPayload == nil {
		return nil
	}
	makerAssetAmount, _ := new(big.Int).SetString(orderPayload.MakerAssetAmount, 10)
	makerFee, _ := new(big.Int).SetString(orderPayload.MakerFee, 10)
	takerAssetAmount, _ := new(big.Int).SetString(orderPayload.TakerAssetAmount, 10)
	takerFee, _ := new(big.Int).SetString(orderPayload.TakerFee, 10)
	expirationTimeSeconds, _ := new(big.Int).SetString(orderPayload.ExpirationTimeSeconds, 10)
	salt, _ := new(big.Int).SetString(orderPayload.Salt, 10)
	var fillableTakerAssetAmount *big.Int
	if orderPayload.FillableTakerAssetAmount != nil {
		fillableTakerAssetAmount, _ = new(big.Int).SetString(*orderPayload.FillableTakerAssetAmount, 10)
	}
	makerAssetData, _ := decoder.DecodeToInterface(common.FromHex(orderPayload.MakerAssetData))
	makerFeeAssetData, _ := decoder.DecodeToInterface(common.FromHex(orderPayload.MakerFeeAssetData))
	takerAssetData, _ := decoder.DecodeToInterface(common.FromHex(orderPayload.TakerAssetData))
	takerFeeAssetData, _ := decoder.DecodeToInterface(common.FromHex(orderPayload.TakerFeeAssetData))
	return &SignedOrder{
		Hash:                     common.HexToHash(orderPayload.Hash),
		SenderAddress:            orderPayload.SenderAddress,
		MakerAddress:             orderPayload.MakerAddress,
		MakerAssetData:           makerAssetData,
		MakerAssetAmount:         makerAssetAmount,
		MakerFee:                 makerFee,
		MakerFeeAssetData:        makerFeeAssetData,
		TakerAddress:             orderPayload.TakerAddress,
		TakerAssetData:           takerAssetData,
		TakerAssetAmount:         takerAssetAmount,
		TakerFee:                 takerFee,
		TakerFeeAssetData:        takerFeeAssetData,
		FeeRecipientAddress:      orderPayload.FeeRecipientAddress,
		ExpirationTimeSeconds:    expirationTimeSeconds,
		Salt:                     salt,
		Signature:                common.FromHex(orderPayload.Salt),
		FillableTakerAssetAmount: fillableTakerAssetAmount,
	}
}

func (z *Client) SubscribeToOrders(orderChan chan []*OrderEvent) error {
	id := "subscribe"
	c := make(chan []byte)
	z.Listen(id, c)
	q, err := json.Marshal(Query{Query: subscribeQuery})
	if err != nil {
		return err
	}
	log.Infof("Start: %s", string(q))
	if err := z.Start(id, q); err != nil {
		return err
	}
	go func() {
		decoder := NewAssetDataDecoder()
		for msg := range c {
			var messageData OrderSubscriptionResponse
			if err := json.Unmarshal(msg, &messageData); err != nil {
				log.Warnf("%+v", err)
				break
			}
			orderEvents := make([]*OrderEvent, len(messageData.Data.OrderEvents))
			for i, orderEvent := range messageData.Data.OrderEvents {
				orderEvents[i] = &OrderEvent{
					Timestamp:   orderEvent.Timestamp,
					EndState:    orderEvent.EndState,
					SignedOrder: orderPayloadToSignedOrder(decoder, orderEvent.SignedOrder),
				}
			}
			orderChan <- orderEvents
		}
		close(orderChan)
	}()
	return nil
}

const orderQuery = `
query Orders($limit: Int!) {
	orders(limit:$limit) {
		hash
                makerAddress
                takerAddress
                feeRecipientAddress
                senderAddress
                makerAssetAmount
                takerAssetAmount
                makerFee
                takerFee
                expirationTimeSeconds
                salt
                makerAssetData
                takerAssetData
                makerFeeAssetData
                takerFeeAssetData
                hash
                fillableTakerAssetAmount
                signature
	}
}
`

type OrdersResponse struct {
	Data struct {
		Orders []*SignedOrderPayload `json:"orders"`
	} `json:"data"`
}

func (z *Client) Orders(orderChan chan []*SignedOrder, limit int64) error {
	id := "orders"
	c := make(chan []byte)
	z.Listen(id, c)
	q, err := json.Marshal(Query{Query: orderQuery, Variables: map[string]int64{"limit": limit}})
	if err != nil {
		return err
	}
	log.Infof("Start: %s", string(q))
	if err := z.Start(id, q); err != nil {
		return err
	}
	go func() {
		decoder := NewAssetDataDecoder()
		for msg := range c {
			//log.Infof("MSG: %+v", msg)
			var messageData OrdersResponse
			if err := json.Unmarshal(msg, &messageData); err != nil {
				log.Warnf("%+v", err)
				break
			}
			orders := make([]*SignedOrder, len(messageData.Data.Orders))
			for i, order := range messageData.Data.Orders {
				orders[i] = orderPayloadToSignedOrder(decoder, order)
			}
			orderChan <- orders
		}
		close(orderChan)
	}()
	return nil
}
