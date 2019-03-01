// Copyright 2018 The QOS Authors

package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTendermintClient(t *testing.T) {
	//cli := TendermintClient("http://192.168.1.223:26657")
	//
	//height := int64(654547)
	//block, err := cli.RetrieveBlock(&height)
	//assert.Nil(t, err)
	//
	//for _, v := range block.Tx {
	//	txStatus := cli.RetrieveTxResult(v)
	//	assert.Nil(t, err)
	//	assert.Equal(t, types.TxStatusSuccess, txStatus)
	//}
}

func TestConsensusAddressToHex(t *testing.T) {
	addr := "cosmosvalconspub1zcjduepqmqrghhgxs6q8t6h7hlvwygawht78surjg3e47jdjxkkx8pc5n3psylm6gl"
	assert.Equal(t, "981560BCEABD99DA0F20464118628854E3C7555D", Bech32AddressToHex(addr))

	assert.Equal(t, "981560BCEABD99DA0F20464118628854E3C7555D", Bech32AddressToHex("981560BCEABD99DA0F20464118628854E3C7555D"))
}
