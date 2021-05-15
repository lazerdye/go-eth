package zxmesh

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/common"
)

var (
	testERC20AssetData = common.FromHex("0xf47261b0000000000000000000000000a0246c9032bc3a600820415ae600c6388619a14d")
)

func TestAssetDataDecoding(t *testing.T) {
	decoder := NewAssetDataEncoderDecoder()
	assetData, err := decoder.Decode(testERC20AssetData)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(assetData.Params))
	assetDataName, err := decoder.GetName(assetData)
	assert.Nil(t, err)
	assert.Equal(t, "ERC20Token", assetDataName)
}

func TestAssetDataEncoding(t *testing.T) {
	decoder := NewAssetDataEncoderDecoder()
	assetData, err := decoder.Decode(testERC20AssetData)
	assert.Nil(t, err)

	encoded, err := decoder.Encode(assetData)
	assert.Nil(t, err)

	assert.Equal(t, testERC20AssetData, encoded)
}
