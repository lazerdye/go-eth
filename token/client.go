package token

import (
	"context"
	"math"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/lazerdye/go-eth/client"
	"github.com/lazerdye/go-eth/gasstation"
	"github.com/lazerdye/go-eth/token/erc20"
	"github.com/lazerdye/go-eth/wallet"
)

const (
	KyberETH         = "eth"
	KyberETHContract = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	KyberETHDecimals = 18

	WETH         = "weth"
	WETHContract = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	WETHDecimals = 18

	ANT         = "ant"
	ANTContract = "0x960b236A07cf122663c4303350609A66A7B288C0"
	ANTDecimals = 18

	BAND         = "band"
	BANDContract = "0xba11d00c5f74255f56a5e366f4f77f5a186d7f55"
	BANDDecimals = 18

	BAT         = "bat"
	BATContract = "0x0D8775F648430679A709E98d2b0Cb6250d2887EF"
	BATDecimals = 18

	BNT         = "bnt"
	BNTContract = "0x1f573d6fb3f13d689ff844b4ce37794d79a7ff1c"
	BNTDecimals = 18

	// Compound dai
	CDAI         = "cdai"
	CDAIContract = "0x5d3a536E4D6DbD6114cc1Ead35777bAB948E3643"
	CDAIDecimals = 8

	// Compound eth
	CETH         = "ceth"
	CETHContract = "0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"
	CETHDecimals = 8

	CUSD         = "cusd"
	CUSDContract = "0x1410d4eC3D276C0eBbf16ccBE88A4383aE734eD0"
	CUSDDecimals = 18

	DAI         = "dai"
	DAIContract = "0x6B175474E89094C44Da98b954EedeAC495271d0F"
	DAIDecimals = 18

	DATA         = "data"
	DATAContract = "0x0Cf0Ee63788A0849fE5297F3407f701E122cC023"
	DATADecimals = 18

	DGD         = "dgd"
	DGDContract = "0xe0b7927c4af23765cb51314a0e0521a9645f0e2a"
	DGDDecimals = 9

	DGX         = "dgx"
	DGXContract = "0x4f3afec4e5a3f2a6a1a411def7d7dfe50ee057bf"
	DGXDecimals = 9

	ENJ         = "enj"
	ENJContract = "0xf629cbd94d3791c9250152bd8dfbdf380e2a3b9c"
	ENJDecimals = 18

	EXMR         = "exmr"
	EXMRContract = "0x331fA6C97c64e47475164b9fC8143b533c5eF529"
	EXMRDecimals = 18

	FOAM         = "foam"
	FOAMContract = "0x4946Fcea7C692606e8908002e55A582af44AC121"
	FOAMDecimals = 18

	FUN         = "fun"
	FUNContract = "0x419D0d8BdD9aF5e606Ae2232ed285Aff190E711b"
	FUNDecimals = 8

	GNO         = "gno"
	GNOContract = "0x6810e776880c02933d47db1b9fc05908e5386b96"
	GNODecimals = 18

	GWIT         = "gwit"
	GWITContract = "0x55D0Bb8D7e7fBf5B863C7923c4645FF83c3D0033"
	GWITDecimals = 18

	HOT         = "hot"
	HOTContract = "0x6c6ee5e31d828de241282b9606c8e98ea48526e2"
	HOTDecimals = 18

	KEY         = "key"
	KEYContract = "0x4cc19356f2d37338b9802aa8e8fc58b0373296e7"
	KEYDecimals = 18

	KNC         = "knc"
	KNCContract = "0xdd974d5c2e2928dea5f71b9825b8b646686bd200"
	KNCDecimals = 18

	LABX         = "labx"
	LABXContract = "0xF0daeC652dD7C9f779e7C0F3ff5610fa3B61f61F"
	LABXDecimals = 18

	LEND         = "lend"
	LENDContract = "0x80fB784B7eD66730e8b1DBd9820aFD29931aab03"
	LENDDecimals = 18

	LINK         = "link"
	LINKContract = "0x514910771AF9Ca656af840dff83E8264EcF986CA"
	LINKDecimals = 18

	LPT         = "lpt"
	LPTContract = "0x58b6a8a3302369daec383334672404ee733ab239"
	LPTDecimals = 18

	LRC         = "lrc"
	LRCContract = "0xBBbbCA6A901c926F240b89EacB641d8Aec7AEafD"
	LRCDecimals = 18

	MANA         = "mana"
	MANAContract = "0x0F5D2fB29fb7d3CFeE444a200298f468908cC942"
	MANADecimals = 18

	MLN         = "mln"
	MLNContract = "0xec67005c4e498ec7f55e092bd1d35cbc47c91892"
	MLNDecimals = 18

	MKR         = "mkr"
	MKRContract = "0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2"
	MKRDecimals = 18

	NMR         = "nmr"
	NMRContract = "0x1776e1f26f98b1a5df9cd347953a26dd3cb46671"
	NMRDecimals = 18

	OHDAI         = "ohdai"
	OHDAIContract = "0x64a03cE1E52B4e579f0A1cf44cF95C0D7898B5A3"
	OHDAIDecimals = 18

	PNK         = "pnk"
	PNKContract = "0x93ed3fbe21207ec2e8f2d3c3de6e058cb73bc04d"
	PNKDecimals = 18

	RCN         = "rcn"
	RCNContract = "0xF970b8E36e23F7fC3FD752EeA86f8Be8D83375A6"
	RCNDecimals = 18

	RDN         = "rdn"
	RDNContract = "0x255Aa6DF07540Cb5d3d297f0D0D4D84cb52bc8e6"
	RDNDecimals = 18

	REP         = "rep"
	REPContract = "0x1985365e9f78359a9B6AD760e32412f4a445E862"
	REPDecimals = 18

	RIGO         = "rigo"
	RIGOContract = "0x4FbB350052Bca5417566f188eB2EBCE5b19BC964"
	RIGODecimals = 18

	SAI         = "sai"
	SAIContract = "0x89d24A6b4CcB1B6fAA2625fE562bDD9a23260359"
	SAIDecimals = 18

	SNT         = "snt"
	SNTContract = "0x744d70fdbe2ba4cf95131626614a1763df805b9e"
	SNTDecimals = 18

	SPANK         = "spank"
	SPANKContract = "0x42d6622deCe394b54999Fbd73D108123806f6a18"
	SPANKDecimals = 18

	STORJ         = "storj"
	STORJContract = "0xb64ef51c888972c908cfacf59b47c1afbc0ab8ac"
	STORJDecimals = 8

	TNT         = "tnt"
	TNTContract = "0x08f5a9235B08173b7569F83645d2c7fB55e8cCD8"
	TNTDecimals = 8

	TUSD         = "tusd"
	TUSDContract = "0x0000000000085d4780B73119b644AE5ecd22b376"
	TUSDDecimals = 18

	UBT         = "ubt"
	UBTContract = "0x8400D94A5cb0fa0D041a3788e395285d61c9ee5e"
	UBTDecimals = 8

	USDC         = "usdc"
	USDCContract = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	USDCDecimals = 6

	ZRX         = "zrx"
	ZRXContract = "0xE41d2489571d322189246DaFA5ebDe1F4699F498"
	ZRXDecimals = 18

	averageGasSpeed  = gasstation.Average
	approveGasLimit  = uint64(200000)
	transferGasLimit = uint64(300000)
)

var (
	DefaultRegistry = NewRegistry()

	maxIntForCap = big.NewInt(1e12)
)

func init() {
	DefaultRegistry.Register(KyberETH, KyberETHContract, KyberETHDecimals)
	DefaultRegistry.Register(WETH, WETHContract, WETHDecimals)
	DefaultRegistry.Register(ANT, ANTContract, ANTDecimals)
	DefaultRegistry.Register(BAND, BANDContract, BANDDecimals)
	DefaultRegistry.Register(BAT, BATContract, BATDecimals)
	DefaultRegistry.Register(BNT, BNTContract, BNTDecimals)
	DefaultRegistry.Register(CDAI, CDAIContract, CDAIDecimals)
	DefaultRegistry.Register(CETH, CETHContract, CETHDecimals)
	DefaultRegistry.Register(CUSD, CUSDContract, CUSDDecimals)
	DefaultRegistry.Register(DAI, DAIContract, DAIDecimals)
	DefaultRegistry.Register(DATA, DATAContract, DATADecimals)
	DefaultRegistry.Register(DGD, DGDContract, DGDDecimals)
	DefaultRegistry.Register(DGX, DGXContract, DGXDecimals)
	DefaultRegistry.Register(ENJ, ENJContract, ENJDecimals)
	DefaultRegistry.Register(EXMR, EXMRContract, EXMRDecimals)
	DefaultRegistry.Register(FOAM, FOAMContract, FOAMDecimals)
	DefaultRegistry.Register(FUN, FUNContract, FUNDecimals)
	DefaultRegistry.Register(GNO, GNOContract, GNODecimals)
	DefaultRegistry.Register(GWIT, GWITContract, GWITDecimals)
	DefaultRegistry.Register(HOT, HOTContract, HOTDecimals)
	DefaultRegistry.Register(KEY, KEYContract, KEYDecimals)
	DefaultRegistry.Register(KNC, KNCContract, KNCDecimals)
	DefaultRegistry.Register(LABX, LABXContract, LABXDecimals)
	DefaultRegistry.Register(LEND, LENDContract, LENDDecimals)
	DefaultRegistry.Register(LINK, LINKContract, LINKDecimals)
	DefaultRegistry.Register(LPT, LPTContract, LPTDecimals)
	DefaultRegistry.Register(LRC, LRCContract, LRCDecimals)
	DefaultRegistry.Register(MKR, MKRContract, MKRDecimals)
	DefaultRegistry.Register(NMR, NMRContract, NMRDecimals)
	DefaultRegistry.Register(MANA, MANAContract, MANADecimals)
	DefaultRegistry.Register(MLN, MLNContract, MLNDecimals)
	DefaultRegistry.Register(OHDAI, OHDAIContract, OHDAIDecimals)
	DefaultRegistry.Register(PNK, PNKContract, PNKDecimals)
	DefaultRegistry.Register(RCN, RCNContract, RCNDecimals)
	DefaultRegistry.Register(RDN, RDNContract, RDNDecimals)
	DefaultRegistry.Register(REP, REPContract, REPDecimals)
	DefaultRegistry.Register(RIGO, RIGOContract, RIGODecimals)
	DefaultRegistry.Register(SAI, SAIContract, SAIDecimals)
	DefaultRegistry.Register(SNT, SNTContract, SNTDecimals)
	DefaultRegistry.Register(SPANK, SPANKContract, SPANKDecimals)
	DefaultRegistry.Register(STORJ, STORJContract, STORJDecimals)
	DefaultRegistry.Register(TNT, TNTContract, TNTDecimals)
	DefaultRegistry.Register(TUSD, TUSDContract, TUSDDecimals)
	DefaultRegistry.Register(UBT, UBTContract, UBTDecimals)
	DefaultRegistry.Register(USDC, USDCContract, USDCDecimals)
	DefaultRegistry.Register(ZRX, ZRXContract, ZRXDecimals)
}

type Token struct {
	name     string
	contract common.Address
	decimals int
}

func (t Token) FromWei(i *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(i), big.NewFloat(math.Pow10(t.decimals)))
}

func (t Token) ToWei(f *big.Float) *big.Int {
	i, _ := new(big.Float).Mul(f, big.NewFloat(math.Pow10(t.decimals))).Int(nil)
	return i
}

func (t Token) Contract() common.Address {
	return t.contract
}

func (t Token) Name() string {
	return t.name
}

func (t Token) Decimals() int {
	return t.decimals
}

type Registry struct {
	data  map[string]Token
	mutex sync.RWMutex
}

func NewRegistry() *Registry {
	return &Registry{
		data: make(map[string]Token),
	}
}

func (r *Registry) Register(name, contract string, decimals int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.data[name]; ok {
		// Name already exists.
		return errors.Errorf("Token with name %s already exists", name)
	}
	r.data[name] = Token{
		name:     name,
		contract: common.HexToAddress(contract),
		decimals: decimals,
	}
	return nil
}

func (r *Registry) ByName(name string) (*Token, bool) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	t, ok := r.data[name]
	return &t, ok
}

func (r *Registry) ByAddress(address common.Address) (*Token, bool) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, t := range r.data {
		if t.contract.Hex() == address.Hex() {
			return &t, true
		}
	}
	return nil, false
}

func (r *Registry) Names() []string {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	ret := make([]string, len(r.data))
	i := 0
	for n, _ := range r.data {
		ret[i] = n
		i++
	}
	return ret
}

type Client struct {
	*client.Client

	instance *erc20.Erc20
	info     *Token
}

func NewClient(client *client.Client, token *Token) (*Client, error) {
	instance, err := erc20.NewErc20(token.contract, client)
	if err != nil {
		return nil, err
	}
	return &Client{client, instance, token}, nil
}

// TODO: This file needs an overhaul.
func ByRawAddress(ctx context.Context, client *client.Client, address common.Address) (*Client, error) {
	opts := bind.CallOpts{Context: ctx}
	instance, err := erc20.NewErc20(address, client)
	if err != nil {
		return nil, err
	}
	decimal, err := instance.Decimals(&opts)
	if err != nil {
		return nil, err
	}
	name, err := instance.Symbol(&opts)
	if err != nil {
		return nil, err
	}
	token := Token{
		name:     name,
		contract: address,
		decimals: int(decimal),
	}
	return &Client{client, instance, &token}, nil
}

func ByAddress(ctx context.Context, client *client.Client, name string, address common.Address) (*Client, error) {
	instance, err := erc20.NewErc20(address, client)
	if err != nil {
		return nil, err
	}
	opts := bind.CallOpts{Context: ctx}
	decimal, err := instance.Decimals(&opts)
	if err != nil {
		return nil, err
	}
	token := Token{
		name:     name,
		contract: address,
		decimals: int(decimal),
	}
	return &Client{client, instance, &token}, nil
}

func ByName(client *client.Client, name string) (*Client, error) {
	token, ok := DefaultRegistry.ByName(name)
	if !ok {
		return nil, errors.Errorf("Unknown token: %s", name)
	}
	return NewClient(client, token)
}

func (c Client) FromWei(i *big.Int) *big.Float {
	return c.info.FromWei(i)
}

func (c Client) ToWei(f *big.Float) *big.Int {
	return c.info.ToWei(f)
}

func (c Client) ToWeiCapped(f *big.Float, address common.Address) (*big.Int, error) {
	balance, err := c.instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}
	amount := c.ToWei(f)
	if amount.Cmp(balance) > 0 {
		diff := new(big.Int).Sub(amount, balance)
		log.Infof("XXX bal: %d amount: %d diff: %d", balance, amount, diff)
		if diff.Cmp(maxIntForCap) > 0 {
			return nil, errors.Errorf("Not enough balance for cap (%d > %d)", amount, balance)
		}
		return balance, nil
	}
	return amount, nil
}

func (c *Client) ContractAddress() common.Address {
	return c.info.contract
}

func (c *Client) Decimals() int {
	return c.info.decimals
}

func (c *Client) Name() string {
	return c.info.Name()
}

func (c *Client) BalanceOf(ctx context.Context, address common.Address) (*big.Float, error) {
	balance, err := c.instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}

	return c.FromWei(balance), nil
}

func (c *Client) Transfer(ctx context.Context, sourceAccount *wallet.Account, destAccount common.Address, amount *big.Float) (*types.Transaction, error) {
	gasPrice, _, err := c.GasPrice(ctx, client.TransferGasSpeed)
	if err != nil {
		return nil, err
	}
	fAmount := new(big.Float).Mul(amount, big.NewFloat(math.Pow10(c.info.decimals)))
	iAmount, _ := fAmount.Int(nil)
	opts, err := sourceAccount.NewTransactor(ctx, nil, gasPrice, transferGasLimit)
	if err != nil {
		return nil, err
	}
	return c.instance.Transfer(opts, destAccount, iAmount)
}

func (c *Client) Approve(ctx context.Context, from *wallet.Account, contract common.Address, value *big.Float) (*types.Transaction, error) {
	gasPrice, _, err := c.GasPrice(ctx, client.TransferGasSpeed)
	if err != nil {
		return nil, err
	}
	opts, err := from.NewTransactor(ctx, nil, gasPrice, approveGasLimit)
	if err != nil {
		return nil, err
	}

	iAmount := c.ToWei(value)
	return c.instance.Approve(opts, contract, iAmount)
}

func (c *Client) Allowance(ctx context.Context, address, contract common.Address) (*big.Float, error) {
	opts := bind.CallOpts{Context: ctx}
	iAmount, err := c.instance.Allowance(&opts, address, contract)
	if err != nil {
		return nil, err
	}
	return c.FromWei(iAmount), nil
}
