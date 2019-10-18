// Copyright 2018 The QOS Authors

package lib

import (
	"fmt"
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
	addr := "qosconspub1zcjduepqjqde5lsuln43wqgx73d4d8t0ecaqd055fanczd3gs37cuep95keqt9yv9r"
	assert.Equal(t, "A1089E4792E1F63A2D67F3067F8205B896D7C355", Bech32AddressToHex(addr))

	//assert.Equal(t, "981560BCEABD99DA0F20464118628854E3C7555D", Bech32AddressToHex("981560BCEABD99DA0F20464118628854E3C7555D"))
}

func TestHexToBech32Address(t *testing.T) {
	fmt.Println(PubkeyToBech32Address("qosconspub", "tendermint/PubKeyEd25519", "qosconspub1zcjduepqmgd4n9eul3r540f7mw8vzwzw9j5dt7q7wgz57lrj39ufr0cxx6zqpu85ct"))
	assert.Equal(t, "A1089E4792E1F63A2D67F3067F8205B896D7C355",
		PubkeyToBech32Address("qosconspub", "tendermint/PubKeyEd25519", "qosconspub1zcjduepqmgd4n9eul3r540f7mw8vzwzw9j5dt7q7wgz57lrj39ufr0cxx6zqpu85ct"))
}
