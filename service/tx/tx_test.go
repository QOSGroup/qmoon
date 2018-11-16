// Copyright 2018 The QOS Authors

package tx

import (
	"log"
	"testing"

	qbasetxs "github.com/QOSGroup/qbase/txs"
	qbasetypes "github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/lib/tmcli"
	"github.com/QOSGroup/qmoon/utils"
	qostxs "github.com/QOSGroup/qos/txs"
	"github.com/QOSGroup/qos/txs/approve"
	"github.com/QOSGroup/qstars/x/bank"
	"github.com/QOSGroup/qstars/x/kvstore"
	"github.com/stretchr/testify/assert"
)

const dd = "/C/xVwrQAQpScS2EEgolChSz9n5iYOIL6u+90iOhO8hTmJbWHBIBMBoKCgVxc3RhchIBMRIlChSAx3bUHCBBN7OUgge+LRxq57RtSxIBMBoKCgVxc3RhchIBMRJtCiUWJN5kIBPy2AYUkun4FJUQtbZ9NA0Zn/JPNMhdu71+DfeA6abMEkABJ/tmgFLbRmMMehOcb1IKjinDhNL+Fi13nTx9Yj+lw97gEXJbPP5Gtkyk/olclFerkABIQO6I7MiKd3HIb3QGGO6oBBoIcW9zLXRlc3QiATASC3FzdGFycy10ZXN0Gghxb3MtdGVzdCD0qAQqaQolFiTeZCA7Reelc9iSfiNZfAE+AbXCnVoLHS266D1iVzRYcGeXlBJAo9gX6etadNMV8pc03b3ywHXiYlKA8mtFFCLBRYQbrlPIlA6FpIoyieYOKrXw1fmMVSWHuotqWaxvJCpdbMPQCDCwnyBKO2hlaWd0aDoyNjQxNTIsaGFzaDpGQUU1N0FFREQ1MTdENEE0RTdEQjBEQjk0QTY3RkI5MjA5MDMyRDQ3"

func checkTxStd(t qbasetxs.ITx) {
	switch impl := t.(type) {
	case *approve.ApproveCancelTx:
		log.Printf("qbasetxs.ITx approve.ApproveCancelTx: %+v", impl)
	case *approve.ApproveCreateTx:
		log.Printf("qbasetxs.ITx approve.ApproveCreateTx: %+v", impl)
	case *approve.ApproveDecreaseTx:
		log.Printf("qbasetxs.ITx approve.ApproveDecreaseTx: %+v", impl)
	case *approve.ApproveIncreaseTx:
		log.Printf("qbasetxs.ITx approve.ApproveIncreaseTx: %+v", impl)
	case *approve.ApproveUseTx:
		log.Printf("qbasetxs.ITx approve.ApproveUseTx: %+v", impl)
	case *kvstore.KvstoreTx:
		log.Printf("qbasetxs.ITx kvstore.KvstoreTx: %+v", impl)
	case *qbasetxs.QcpTxResult:
		log.Printf("qbasetxs.ITx qbasetxs.QcpTxResult: %+v", impl)
	case *qostxs.TransferTx:
		log.Printf("qbasetxs.ITx qostxs.TransferTx: %+v", impl)
	case *qostxs.TxCreateQSC:
		log.Printf("qbasetxs.ITx qostxs.TxCreateQSCx: %+v", impl)
	case *qostxs.TxIssueQsc:
		log.Printf("qbasetxs.ITx qostxs.TxIssueQsc: %+v", impl)
	case *bank.WrapperSendTx:
		log.Printf("qbasetxs.ITx bank.WrapperSendTx: %+v", impl.Wrapper.ITx)
	default:
		log.Printf("qbasetxs.ITx not found txstd")
	}
}

/*
txs.TxStd
txs.TxQcp:
*/
func testSend(t *testing.T) {
	cdc := lib.MakeCodec()

	tmc := lib.TendermintClient("http://192.168.1.224:26657")
	var height int64 = 488039
	b, err := tmc.Block(&height)
	assert.Nil(t, err)

	qbasetx, err := qbasetypes.DecoderTx(cdc, []byte(string(b.Block.Txs[0])))
	assert.Nil(t, err)

	//初始化context相关数据

	switch implTx := qbasetx.(type) {
	case *qbasetxs.TxStd:
		t.Logf("---TxStd:%+v- %+v", implTx, implTx.ITx)
		checkTxStd(implTx.ITx)
	case *qbasetxs.TxQcp:
		t.Logf("---TxQcp:%+v, %+v", implTx, implTx.TxStd)
		checkTxStd(implTx.TxStd.ITx)
	default:
		t.Log("not support itx type")
	}
}

func TestSave(t *testing.T) {
	b, err := tmcli.NewClient(nil).Block.Retrieve(nil, nil)
	assert.Nil(t, err)
	err = Save(b)
	assert.Nil(t, err)

	mt, err := model.TxesByChainIDHeight(db.Db, utils.NullString(b.Block.ChainID), utils.NullInt64(b.Block.Height))
	assert.Nil(t, err)
	assert.Equal(t, "TransferTx", mt[0].TxType.String)
}
