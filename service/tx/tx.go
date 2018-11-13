// Copyright 2018 The QOS Authors

package tx

import (
	"errors"
	"fmt"
	"strings"
	"time"

	qbasetxs "github.com/QOSGroup/qbase/txs"
	qbasetypes "github.com/QOSGroup/qbase/types"
	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	qostxs "github.com/QOSGroup/qos/txs"
	qosapprove "github.com/QOSGroup/qos/txs/approve"
	"github.com/QOSGroup/qstars/x/bank"
	"github.com/QOSGroup/qstars/x/kvstore"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func convertToTx(mt *model.Tx) *types.ResultTx {
	return &types.ResultTx{
		ChainID:     mt.ChainID.String,
		Height:      mt.Height.Int64,
		Index:       mt.Index.Int64,
		TxType:      mt.TxType.String,
		Maxgas:      mt.Maxgas.Int64,
		QcpFrom:     mt.QcpFrom.String,
		QcpTo:       mt.QcpTo.String,
		QcpSequence: mt.QcpSequence.Int64,
		QcpTxindex:  mt.QcpTxindex.Int64,
		QcpIsresult: mt.QcpIsresult.Bool,
		Data:        []byte(mt.JSONTx.String),
		Time:        mt.Time.Time,
		CreatedAt:   mt.CreatedAt.Time,
	}
}

// Search 交易查询
func List(chainID string, minHeight, maxHeight int64) ([]*types.ResultTx, error) {
	const sqlOrder = " order by height desc "
	var limit string
	var wheres []string
	wheres = append(wheres, fmt.Sprintf(" chain_id = '%s' ", chainID))
	if minHeight != 0 && maxHeight != 0 {
		wheres = append(wheres, fmt.Sprintf(" height >= %d ", minHeight))
		wheres = append(wheres, fmt.Sprintf(" height <= %d ", maxHeight))
	} else {
		limit = " limit 20"
	}

	mbs, err := model.TxFilter(db.Db, strings.Join(wheres, " and ")+sqlOrder+limit, 0, -1)
	if err != nil {
		return nil, err
	}

	var res []*types.ResultTx
	for _, v := range mbs {
		res = append(res, convertToTx(v))
	}

	return res, err
}

// Search 交易查询
func Retrieve(chainID string, height, index int64) (*types.ResultTx, error) {
	mt, err := model.TxByChainIDHeightIndex(db.Db, utils.NullString(chainID), utils.NullInt64(height), utils.NullInt64(index))
	if err != nil {
		return nil, err
	}

	return convertToTx(mt), err
}

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
		mt.OriginTx = utils.NullString(utils.Base64En(v))
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
	cdc := lib.MakeCodec()
	d, err := cdc.MarshalJSON(t)
	if err != nil {
		return err
	}
	mt.JSONTx = utils.NullString(string(d))

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
