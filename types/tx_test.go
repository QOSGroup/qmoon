// Copyright 2019 The QOS Authors

package types

import (
	"encoding/json"
	"testing"

	"github.com/tidwall/gjson"
)

func TestCosmosTxParse(t *testing.T) {
	s := `{"chain_id":"","hash":"3128A9423C3532D5E3B8729AF4575832CBBC4260AF4D0E16FF808ED7AAC48496","height":"56912","index":"0","tx_type":"delegate","tx_type_cn":"委托","gas_wanted":"200000","gas_used":"105186","fee":"10000uatom","tx_status":"成功","status":"1","data":[{"type":"delegate","data":{"delegator_address":"cosmos127ax4kl8kfrjtl20gyvlumqqsxfp66hj8fqk4v","validator_address":"cosmosvaloper1clpqr4nrk4khgkxj78fcwwh6dl3uw4epsluffn","value":{"denom":"uatom","amount":"11000000"}}}]}`
	var tx ResultTx
	_ = json.Unmarshal([]byte(s), &tx)
	t.Logf("tx:%+v", string(tx.Data))

	m := make(map[string][]json.RawMessage)
	result := gjson.Parse(string(tx.Data))
	result.ForEach(func(_, value gjson.Result) bool {
		t.Logf("value: %s", value.String())

		if _, ok := m[value.Get("type").String()]; ok {
			m[value.Get("type").String()] = append(m[value.Get("type").String()], json.RawMessage(value.Get("data").String()))
		} else {
			m[value.Get("type").String()] = []json.RawMessage{json.RawMessage(value.Get("data").String())}
		}
		return true
	})

	tx.ITxs = m

	t.Logf("---%+v", tx)
}
