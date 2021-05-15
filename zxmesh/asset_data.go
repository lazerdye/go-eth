package zxmesh

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	// ERC20AssetDataID is the assetDataId for ERC20 tokens
	ERC20AssetDataID = "f47261b0"

	// ERC721AssetDataID is the assetDataId for ERC721 tokens
	ERC721AssetDataID = "02571792"

	// ERC1155AssetDataID is the assetDataId for ERC721 tokens
	ERC1155AssetDataID = "a7cb5fb7"

	// StaticCallAssetDataID is the assetDataId for staticcalls
	StaticCallAssetDataID = "c339d10a"

	// CheckGasDefaultID is the function selector for the `checkGas` function that does
	// not accept a gasPrice.
	CheckGasPriceDefaultID = "d728f5b7"

	// CheckGasID is the function selector for the `checkGas` function that accepts a gasPrice.
	CheckGasPriceID = "da5b166a"

	// MultiAssetDataID is the assetDataId for multiAsset tokens
	MultiAssetDataID = "94cfcdd7"

	// ERC20BridgeAssetDataID is the assetDataId for ERC20Bridge assets
	ERC20BridgeAssetDataID = "dc1600f3"

	erc20AssetDataAbi                     = "[{\"inputs\":[{\"name\":\"address\",\"type\":\"address\"}],\"name\":\"ERC20Token\",\"type\":\"function\"}]"
	erc721AssetDataAbi                    = "[{\"inputs\":[{\"name\":\"address\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721Token\",\"type\":\"function\"}]"
	erc1155AssetDataAbi                   = "[{\"constant\":false,\"inputs\":[{\"name\":\"address\",\"type\":\"address\"},{\"name\":\"ids\",\"type\":\"uint256[]\"},{\"name\":\"values\",\"type\":\"uint256[]\"},{\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"ERC1155Assets\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	staticCallAssetDataAbi                = "[{\"inputs\":[{\"name\":\"staticCallTargetAddress\",\"type\":\"address\"},{\"name\":\"staticCallData\",\"type\":\"bytes\"},{\"name\":\"expectedReturnHashData\", \"type\":\"bytes32\"}],\"name\":\"StaticCall\",\"type\":\"function\"}]"
	checkGasPriceDefaultStaticCallDataAbi = "[{\"inputs\":[],\"name\":\"checkGasPrice\",\"type\":\"function\"}]"
	checkGasPriceStaticCallDataAbi        = "[{\"inputs\":[{\"name\":\"maxGasPrice\",\"type\":\"uint256\"}],\"name\":\"checkGasPrice\",\"type\":\"function\"}]"
	multiAssetDataAbi                     = "[{\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"name\":\"nestedAssetData\",\"type\":\"bytes[]\"}],\"name\":\"MultiAsset\",\"type\":\"function\"}]"
	erc20BridgeAssetDataAbi               = "[{\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"name\":\"bridgeData\",\"type\":\"bytes\"}],\"name\":\"ERC20Bridge\",\"type\":\"function\"}]"
)

type AssetData struct {
	Id     string
	Params []interface{}
}

type assetDataInfo struct {
	name string
	abi  abi.ABI
}

type AssetDataEncoderDecoder struct {
	idToAssetDataInfo map[string]assetDataInfo
}

func NewAssetDataEncoderDecoder() *AssetDataEncoderDecoder {
	erc20AssetDataABI, err := abi.JSON(strings.NewReader(erc20AssetDataAbi))
	if err != nil {
		log.WithField("erc20AssetDataAbi", erc20AssetDataAbi).Panic("erc20AssetDataAbi should be ABI parsable")
	}
	erc721AssetDataABI, err := abi.JSON(strings.NewReader(erc721AssetDataAbi))
	if err != nil {
		log.WithField("erc721AssetDataAbi", erc721AssetDataAbi).Panic("erc721AssetDataAbi should be ABI parsable")
	}
	erc1155AssetDataABI, err := abi.JSON(strings.NewReader(erc1155AssetDataAbi))
	if err != nil {
		log.WithField("erc1155AssetDataAbi", erc1155AssetDataAbi).Panic("erc1155AssetDataAbi should be ABI parsable")
	}
	staticCallAssetDataABI, err := abi.JSON(strings.NewReader(staticCallAssetDataAbi))
	if err != nil {
		log.WithField("staticCallAssetDataAbi", staticCallAssetDataAbi).Panic("staticCallAssetDataAbi should be ABI parsable")
	}
	checkGasPriceDefaultStaticCallDataABI, err := abi.JSON(strings.NewReader(checkGasPriceDefaultStaticCallDataAbi))
	if err != nil {
		log.WithField("checkGasPriceDefaultStaticCallDataAbi", checkGasPriceDefaultStaticCallDataAbi).Panic("checkGasPriceDefaultStaticCallDataAbi should be ABI parsable")
	}
	checkGasPriceStaticCallDataABI, err := abi.JSON(strings.NewReader(checkGasPriceStaticCallDataAbi))
	if err != nil {
		log.WithField("checkGasPriceStaticCallDataAbi", checkGasPriceStaticCallDataAbi).Panic("checkGasStaticCallDataAbi should be ABI parsable")
	}
	multiAssetDataABI, err := abi.JSON(strings.NewReader(multiAssetDataAbi))
	if err != nil {
		log.WithField("multiAssetDataAbi", multiAssetDataAbi).Panic("multiAssetDataAbi should be ABI parsable")
	}
	erc20BridgeAssetDataABI, err := abi.JSON(strings.NewReader(erc20BridgeAssetDataAbi))
	if err != nil {
		log.WithField("erc20BridgeAssetDataABI", erc20BridgeAssetDataAbi).Panic("erc20BridgeAssetDataABI should be ABI parsable")
	}
	idToAssetDataInfo := map[string]assetDataInfo{
		ERC20AssetDataID: {
			name: "ERC20Token",
			abi:  erc20AssetDataABI,
		},
		ERC721AssetDataID: {
			name: "ERC721Token",
			abi:  erc721AssetDataABI,
		},
		ERC1155AssetDataID: {
			name: "ERC1155Assets",
			abi:  erc1155AssetDataABI,
		},
		StaticCallAssetDataID: {
			name: "StaticCall",
			abi:  staticCallAssetDataABI,
		},
		CheckGasPriceDefaultID: {
			name: "checkGasPrice",
			abi:  checkGasPriceDefaultStaticCallDataABI,
		},
		CheckGasPriceID: {
			name: "checkGasPrice",
			abi:  checkGasPriceStaticCallDataABI,
		},
		MultiAssetDataID: {
			name: "MultiAsset",
			abi:  multiAssetDataABI,
		},
		ERC20BridgeAssetDataID: {
			name: "ERC20Bridge",
			abi:  erc20BridgeAssetDataABI,
		},
	}
	return &AssetDataEncoderDecoder{
		idToAssetDataInfo: idToAssetDataInfo,
	}
}

func (ed *AssetDataEncoderDecoder) GetName(assetData *AssetData) (string, error) {
	info, ok := ed.idToAssetDataInfo[assetData.Id]
	if !ok {
		return "", errors.Errorf("Unrecognized assetData with prefix: %s", assetData.Id)
	}
	return info.name, nil
}

func (ed *AssetDataEncoderDecoder) Decode(assetData []byte) (*AssetData, error) {
	if len(assetData) < 4 {
		return nil, errors.New("assetData must be at least 4 bytes long")
	}
	id := assetData[:4]
	idHex := common.Bytes2Hex(id)
	info, ok := ed.idToAssetDataInfo[idHex]
	if !ok {
		return nil, errors.Errorf("Unrecognized assetData with prefix: %s", idHex)
	}
	// This is necessary to prevent a nil pointer exception for ABIs with no inputs
	if len(info.abi.Methods[info.name].Inputs.NonIndexed()) == 0 {
		return nil, nil
	}
	params, err := info.abi.Methods[info.name].Inputs.UnpackValues(assetData[4:])
	if err != nil {
		return nil, err
	}
	return &AssetData{
		Id:     idHex,
		Params: params,
	}, nil
}

func (ed *AssetDataEncoderDecoder) Encode(assetData *AssetData) ([]byte, error) {
	info, ok := ed.idToAssetDataInfo[assetData.Id]
	if !ok {
		return nil, errors.Errorf("Unrecognized assetData with prefix: %s", assetData.Id)
	}
	data, err := info.abi.Methods[info.name].Inputs.PackValues(assetData.Params)
	if err != nil {
		return nil, err
	}
	prefix := common.Hex2Bytes(assetData.Id)
	return append(prefix, data...), nil
}
