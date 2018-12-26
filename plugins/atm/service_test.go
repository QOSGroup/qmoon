// Copyright 2018 The QOS Authors

package atm

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/QOSGroup/qstars/x/bank"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/go-amino"
	tmltypes "github.com/tendermint/tendermint/rpc/lib/types"
)

func TestSend(t *testing.T) {
	//logrus.SetLevel(logrus.DebugLevel)
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	//addr := "address1pcjs0t9m9vl7vejwttuc2fzfgnutndnrpyw08m"
	//chainid := "capricorn-1000"
	//coin := "qos"
	//amount := getAmount()
	//_, err := send(addr, chainid, fmt.Sprintf("%d%s", amount, coin))
	//assert.Nil(t, err)
}

func TestJson(t *testing.T) {
	body := `{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "hash": "1CFC7C08C9E3CB6F8DF7E8668BEABDEA55022237",
    "error": "",
    "code": "-1",
    "result": "BankStub",
    "heigth": "374095"
  }
}`

	var tmresp tmltypes.RPCResponse
	fmt.Println(string(body))
	err := json.Unmarshal([]byte(body), &tmresp)
	assert.Nil(t, err)

	assert.True(t, tmresp.Error == nil)
	var res bank.SendResult
	var cdc = amino.NewCodec()
	err = cdc.UnmarshalJSON(tmresp.Result, &res)
	assert.Nil(t, err)
	assert.Equal(t, "1CFC7C08C9E3CB6F8DF7E8668BEABDEA55022237", res.Hash)
	assert.Equal(t, "374095", res.Heigth)
}
