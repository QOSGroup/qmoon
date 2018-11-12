// Copyright 2018 The QOS Authors

package tx

import (
	"errors"
	"time"

	qbasetxs "github.com/QOSGroup/qbase/txs"
	qbasetypes "github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/utils"
	qostxs "github.com/QOSGroup/qos/txs"
	qosapprove "github.com/QOSGroup/qos/txs/approve"
	"github.com/QOSGroup/qstars/x/bank"
	"github.com/QOSGroup/qstars/x/kvstore"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func Save(b *tmctypes.ResultBlock) error {
	if b.Block.NumTxs == 0 {
		return nil
	}
	cdc := lib.MakeCodec()

	now := time.Now()
	for k, v := range b.Block.Data.Txs {
		qbasetx, err := qbasetypes.DecoderTx(cdc, v)
		if err != nil {
			return err
		}

		mt := &model.Tx{}
		mt.ChainID = utils.NullString(b.Block.Header.ChainID)
		mt.Height = utils.NullInt64(b.Block.Header.Height)
		mt.Index = utils.NullInt64(int64(k))
		mt.Data = utils.NullString(utils.Base64En(b.Block.Txs[0]))
		mt.Time = utils.NullTime(b.Block.Time)
		mt.CreatedAt = utils.NullTime(now)

		if err := ParseTx(qbasetx, mt); err != nil {
			return err
		}

		if err := mt.Insert(db.Db); err != nil {
			return err
		}
	}
	return nil
}

func ParseTx(t qbasetypes.Tx, mt *model.Tx) error {
	var std *qbasetxs.TxStd
	switch implTx := t.(type) {
	case *qbasetxs.TxStd:
		std = implTx
	case *qbasetxs.TxQcp:
		mt.QcpFrom = utils.NullString(implTx.From)
		mt.QcpTo = utils.NullString(implTx.To)
		mt.QcpSequence = utils.NullInt64(implTx.Sequence)
		mt.QcpTxindex = utils.NullInt64(implTx.TxIndex)
		mt.QcpIsresult = utils.NullBool(implTx.IsResult)
		std = implTx.TxStd
	default:
		return errors.New("unsupport itx type")
	}

	mt.Maxgas = utils.NullInt64(std.MaxGas.Int64())

	if err := ParseITx(std.ITx, mt); err != nil {
		return err
	}

	return nil
}

func ParseITx(t qbasetxs.ITx, mt *model.Tx) error {
	switch t.(type) {
	case *qosapprove.ApproveCancelTx:
		mt.TxType = utils.NullString("ApproveCancelTx")
	case *qosapprove.ApproveCreateTx:
		mt.TxType = utils.NullString("ApproveCreateTx")
	case *qosapprove.ApproveDecreaseTx:
		mt.TxType = utils.NullString("ApproveDecreaseTx")
	case *qosapprove.ApproveIncreaseTx:
		mt.TxType = utils.NullString("ApproveIncreaseTx")
	case *qosapprove.ApproveUseTx:
		mt.TxType = utils.NullString("ApproveUseTx")
	case *kvstore.KvstoreTx:
		mt.TxType = utils.NullString("KvstoreTx")
	case *qbasetxs.QcpTxResult:
		mt.TxType = utils.NullString("QcpTxResult")
	case *qostxs.TransferTx:
		mt.TxType = utils.NullString("TransferTx")
	case *qostxs.TxCreateQSC:
		mt.TxType = utils.NullString("TxCreateQSC")
	case *qostxs.TxIssueQsc:
		mt.TxType = utils.NullString("TxIssueQsc")
	case *bank.WrapperSendTx:
		mt.TxType = utils.NullString("WrapperSendTx")
	default:
		mt.TxType = utils.NullString("Unknown")
	}

	return nil
}
