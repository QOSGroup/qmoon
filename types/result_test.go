// Copyright 2018 The QOS Authors

package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTxCN(t *testing.T) {
	t.Run("qos TxTransfer out", func(t *testing.T) {
		txType := "qos/txs/TxTransfer"
		tx := `
{ "senders": [ { "addr": "address1pcjs0t9m9vl7vejwttuc2fzfgnutndnrpyw08m", "qos": "1000000000", "qscs": null } ], "receivers": [ { "addr": "address1t368fjgwerattsc4dy70vzs35k7v00fpv6lumk", "qos": "1000000000", "qscs": null } ] }`
		address := "address1pcjs0t9m9vl7vejwttuc2fzfgnutndnrpyw08m"

		assert.Equal(t, "转出", TxCN(txType, tx, address))
	})

	t.Run("qos TxTransfer into", func(t *testing.T) {
		txType := "qos/txs/TxTransfer"
		tx := `
{ "senders": [ { "addr": "address1pcjs0t9m9vl7vejwttuc2fzfgnutndnrpyw08m", "qos": "1000000000", "qscs": null } ], "receivers": [ { "addr": "address1t368fjgwerattsc4dy70vzs35k7v00fpv6lumk", "qos": "1000000000", "qscs": null } ] }`
		address := "address1t368fjgwerattsc4dy70vzs35k7v00fpv6lumk"

		assert.Equal(t, "转入", TxCN(txType, tx, address))
	})

	t.Run("cosmos send out", func(t *testing.T) {
		txType := "send"
		tx := `
[ { "type": "send", "data": { "from_address": "cosmos1pe9pk2seh2sk6q7mfx3qtzhjs2u4xtq38kwufj", "to_address": "cosmos1e5yhp5lkhjuautf4mhhll7l733za8tgpj329d9", "amount": [ { "denom": "stake", "amount": "40000" } ] } } ]
`
		address := "cosmos1pe9pk2seh2sk6q7mfx3qtzhjs2u4xtq38kwufj"

		assert.Equal(t, "转出", TxCN(txType, tx, address))
	})

	t.Run("cosmos send out", func(t *testing.T) {
		txType := "send"
		tx := `
[ { "type": "send", "data": { "from_address": "cosmos1pe9pk2seh2sk6q7mfx3qtzhjs2u4xtq38kwufj", "to_address": "cosmos1e5yhp5lkhjuautf4mhhll7l733za8tgpj329d9", "amount": [ { "denom": "stake", "amount": "40000" } ] } } ]
`
		address := "cosmos1e5yhp5lkhjuautf4mhhll7l733za8tgpj329d9"

		assert.Equal(t, "转入", TxCN(txType, tx, address))
	})
}
