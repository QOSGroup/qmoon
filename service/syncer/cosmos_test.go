// Copyright 2019 The QOS Authors

package syncer

import (
	"testing"
)

type QueryValidatorParams struct {
	ValidatorAddr string
}

//func TestCosmosDelegate(t *testing.T) {
//	cli := lib.TendermintClient("http://47.105.52.237:36657")
//	height := int64(187332)
//	b, _ := cli.Block(&height)
//
//	b.Block.Txs[0].Hash()
//}

func TestValidator(t *testing.T) {
	//cli := lib.TendermintClient("http://192.168.1.180:26657")
	//endpoint := "custom/staking/validator"
	//s, err := hex.DecodeString(strings.ToLower("1FBF89B0DC10144877BF8C93D7FBAF00E10FFC24"))
	//assert.Nil(t, err)
	//
	//address, err := bech32.ConvertAndEncode("cosmosvaloper", s)
	//assert.Nil(t, err)
	//t.Logf("address1:%s", address)
	//
	//address = lib.PubkeyToBech32Address("cosmosvalconspub", "tendermint/PubKeyEd25519", "9tK9IT+FPdf2qm+5c2qaxi10sWP+3erWTKgftn2PaQM=")
	//t.Logf("address2:%s", address)
	//
	//p := QueryValidatorParams{ValidatorAddr: address}
	//d, err := lib.Cdc.MarshalJSON(p)
	//t.Logf("d:%s", string(d))
	//assert.Nil(t, err)
	//res, err := cli.ABCIQuery(endpoint, d)
	//assert.Nil(t, err)
	//t.Logf("res:%+v", res)
	//
	//res1, err := cli.ABCIQuery("custom/staking/validators", nil)
	//t.Logf("code:%+v", res1.Response.Code)
	//t.Logf("res1.Response.Value:%+v", string(res1.Response.Value))
	//if res1.Response.Code == 0 {
	//	var sv []StakingValidator
	//	json.Unmarshal(res1.Response.Value, &sv)
	//	//t.Logf("sv:%+v", sv)
	//}
}
