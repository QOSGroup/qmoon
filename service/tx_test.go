// Copyright 2018 The QOS Authors

package service

import (
	"log"
	"testing"

	qbasetxs "github.com/QOSGroup/qbase/txs"
	qbasetypes "github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qos/module/approve"
	"github.com/QOSGroup/qos/module/qsc"
	"github.com/stretchr/testify/assert"
)

const dd = "/C/xVwrQAQpScS2EEgolChSz9n5iYOIL6u+90iOhO8hTmJbWHBIBMBoKCgVxc3RhchIBMRIlChSAx3bUHCBBN7OUgge+LRxq57RtSxIBMBoKCgVxc3RhchIBMRJtCiUWJN5kIBPy2AYUkun4FJUQtbZ9NA0Zn/JPNMhdu71+DfeA6abMEkABJ/tmgFLbRmMMehOcb1IKjinDhNL+Fi13nTx9Yj+lw97gEXJbPP5Gtkyk/olclFerkABIQO6I7MiKd3HIb3QGGO6oBBoIcW9zLXRlc3QiATASC3FzdGFycy10ZXN0Gghxb3MtdGVzdCD0qAQqaQolFiTeZCA7Reelc9iSfiNZfAE+AbXCnVoLHS266D1iVzRYcGeXlBJAo9gX6etadNMV8pc03b3ywHXiYlKA8mtFFCLBRYQbrlPIlA6FpIoyieYOKrXw1fmMVSWHuotqWaxvJCpdbMPQCDCwnyBKO2hlaWd0aDoyNjQxNTIsaGFzaDpGQUU1N0FFREQ1MTdENEE0RTdEQjBEQjk0QTY3RkI5MjA5MDMyRDQ3"

func checkTxStd(t qbasetxs.ITx) {
	switch impl := t.(type) {
	case *approve.TxCancelApprove:
		log.Printf("qbasetxs.ITx approve.TxCancelApprove: %+v", impl)
	case *approve.TxCreateApprove:
		log.Printf("qbasetxs.ITx approve.TxCreateApprove: %+v", impl)
	case *approve.TxDecreaseApprove:
		log.Printf("qbasetxs.ITx approve.TxDecreaseApprove: %+v", impl)
	case *approve.TxIncreaseApprove:
		log.Printf("qbasetxs.ITx approve.TxIncreaseApprove: %+v", impl)
	case *approve.TxUseApprove:
		log.Printf("qbasetxs.ITx approve.TxUseApprove: %+v", impl)
	case *qbasetxs.QcpTxResult:
		log.Printf("qbasetxs.ITx qbasetxs.QcpTxResult: %+v", impl)
	case *qsc.TxCreateQSC:
		log.Printf("qbasetxs.ITx qsc.TxCreateQSC: %+v", impl)
	case *qsc.TxIssueQSC:
		log.Printf("qbasetxs.ITx qsc.TxIssueQSC: %+v", impl)
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
