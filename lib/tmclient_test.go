// Copyright 2018 The QOS Authors

package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTendermintClient(t *testing.T) {
	// cli := TendermintClient("http://47.98.253.9:36657")

	// height := int64(1000)
	// validator, err := cli.COSMOSValidator(height)
	// assert.Nil(t, err)
	// t.Log(validator)
	// for _, v := range block.Tx {
	// 	txStatus := cli.RetrieveTxResult(v)
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, types.TxStatusSuccess, txStatus)
	// }
}

func TestConsensusAddressToHex(t *testing.T) {
	addr := "cosmosvalconspub1zcjduepqmqrghhgxs6q8t6h7hlvwygawht78surjg3e47jdjxkkx8pc5n3psylm6gl"
	assert.Equal(t, "981560BCEABD99DA0F20464118628854E3C7555D", Bech32AddressToHex(addr))

	assert.Equal(t, "981560BCEABD99DA0F20464118628854E3C7555D", Bech32AddressToHex("981560BCEABD99DA0F20464118628854E3C7555D"))
}

func TestHexToBech32Address(t *testing.T) {
	assert.Equal(t, "cosmosvalconspub1zcjduepqmqrghhgxs6q8t6h7hlvwygawht78surjg3e47jdjxkkx8pc5n3psylm6gl",
		PubkeyToBech32Address("cosmosvalconspub", "tendermint/PubKeyEd25519", "2AaL3QaGgHXq/r/Y4iOuuvx4cHJEc19JsjWsY4cUnEM="))
}
