// Copyright 2018 The QOS Authors

package txplugins

import (
	"database/sql"
	"errors"
	"time"

	qbasetxs "github.com/QOSGroup/qbase/txs"
	"github.com/QOSGroup/qmoon/txplugins/transfer"
	"github.com/gin-gonic/gin"
)

type TxPlugin interface {
	DbInit(driveName string, db *sql.DB) error
	DbClear(driveName string, db *sql.DB) error

	Parse(chainID string, height int64, hash string, t time.Time, itx qbasetxs.ITx) (typeName string, hit bool, err error)
	Type() string

	RegisterGin(r *gin.Engine)
}

var tps = make(map[string]TxPlugin)

func init() {
	register(&transfer.TxTransferPlugin{})
}

func RegisterGin(r *gin.Engine) {
	for _, v := range tps {
		v.RegisterGin(r)
	}
}

func register(tp TxPlugin) error {
	if _, ok := tps[tp.Type()]; ok {
		return errors.New("aleady registered")
	} else {
		tps[tp.Type()] = tp
	}

	return nil
}

func Init(driveName string, db *sql.DB) error {
	for _, tp := range tps {
		if err := tp.DbInit(driveName, db); err != nil {
			return err
		}
	}

	return nil
}

func Parse(chainID string, height int64, hash string, t time.Time, itx qbasetxs.ITx) (name string, err error) {
	var hit bool
	for _, tp := range tps {
		if name, hit, err = tp.Parse(chainID, height, hash, t, itx); hit {
			break
		}
	}

	if !hit {
		return "", errors.New("unsupport tx")
	}

	return
}

/*

	qbasetxs "github.com/QOSGroup/qbase/txplugins"
	"github.com/QOSGroup/qos/txplugins/approve"
	"github.com/QOSGroup/qos/txplugins/qsc"
	"github.com/QOSGroup/qos/txplugins/transfer"
	"github.com/QOSGroup/qstars/x/bank"
	"github.com/QOSGroup/qstars/x/kvstore"

func ParseITx(t qbasetxs.ITx, mt *model.Tx) error {
	cdc := lib.MakeCodec()
	d, err := cdc.MarshalJSON(t)
	if err != nil {
		return err
	}
	mt.JSONTx = utils.NullString(string(d))

	switch t.(type) {
	case *approve.TxCancelApprove:
		mt.TxType = utils.NullString("ApproveCancelTx")
	case *approve.TxCreateApprove:
		mt.TxType = utils.NullString("ApproveCreateTx")
	case *approve.TxDecreaseApprove:
		mt.TxType = utils.NullString("ApproveDecreaseTx")
	case *approve.TxIncreaseApprove:
		mt.TxType = utils.NullString("ApproveIncreaseTx")
	case *approve.TxUseApprove:
		mt.TxType = utils.NullString("ApproveUseTx")
	case *kvstore.KvstoreTx:
		mt.TxType = utils.NullString("KvstoreTx")
	case *qbasetxs.QcpTxResult:
		mt.TxType = utils.NullString("QcpTxResult")
	case *transfer.TxTransfer:
		mt.TxType = utils.NullString("TransferTx")
	case *qsc.TxCreateQSC:
		mt.TxType = utils.NullString("TxCreateQSC")
	case *qsc.TxIssueQSC:
		mt.TxType = utils.NullString("TxIssueQsc")
	case *bank.WrapperSendTx:
		mt.TxType = utils.NullString("WrapperSendTx")
	default:
		mt.TxType = utils.NullString("Unknown")
	}

	return nil
}


*/
