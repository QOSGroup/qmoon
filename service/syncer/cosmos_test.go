// Copyright 2018 The QOS Authors

package syncer

//
//func TestCOSMOS(t *testing.T) {
//	node := &service.Node{
//		Name:    "",
//		BaseURL: "http://192.168.1.215:18080",
//		ChanID:  "",
//	}
//	cosmos := COSMOS{
//		node:  node,
//		tmcli: lib.TendermintClient(node.BaseURL),
//	}
//
//	cdc := lib.COSMOSCdc
//
//	t.Run("block", func(t *testing.T) {
//		height := int64(1723)
//		b, err := cosmos.tmcli.RetrieveBlock(&height)
//		assert.Nil(t, err)
//
//		assert.NotEqual(t, 0, len(b.Txs))
//
//		var tx auth.StdTx
//		err = cdc.UnmarshalBinaryLengthPrefixed(b.Txs[0], &tx)
//		assert.Nil(t, err)
//		fmt.Printf("%+v", tx.Msgs)
//
//	})
//}
