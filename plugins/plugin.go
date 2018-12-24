// Copyright 2018 The QOS Authors

package plugins

import (
	"database/sql"
	"errors"

	qbasetxs "github.com/QOSGroup/qbase/txs"
	"github.com/QOSGroup/qmoon/plugins/atm"
	"github.com/QOSGroup/qmoon/plugins/transfer"
	"github.com/gin-gonic/gin"
	tmtypes "github.com/tendermint/tendermint/types"
)

type Pluginer interface {
	DbInit(driveName string, db *sql.DB) error
	DbClear(driveName string, db *sql.DB) error

	Parse(blockHeader tmtypes.Header, itx qbasetxs.ITx) (typeName string, hit bool, err error)
	Type() string

	RegisterGin(r *gin.Engine)
}

var tps = make(map[string]Pluginer)

func init() {
	register(&transfer.TxTransferPlugin{})
	register(&atm.ATMPlugin{})
}

func RegisterGin(r *gin.Engine) {
	for _, v := range tps {
		v.RegisterGin(r)
	}
}

func register(tp Pluginer) error {
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

func Parse(blockHeader tmtypes.Header, itx qbasetxs.ITx) (name string, err error) {
	var hit bool
	for _, tp := range tps {
		if name, hit, err = tp.Parse(blockHeader, itx); hit {
			break
		}
	}

	if !hit {
		return "", errors.New("unsupport tx")
	}

	return
}

/*

	qbasetxs "github.com/QOSGroup/qbase/plugins"
	"github.com/QOSGroup/qos/plugins/approve"
	"github.com/QOSGroup/qos/plugins/qsc"
	"github.com/QOSGroup/qos/plugins/transfer"
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
