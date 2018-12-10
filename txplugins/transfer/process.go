// Copyright 2018 The QOS Authors

package transfer

import (
	"database/sql"
	"log"
	"time"

	qbasetxs "github.com/QOSGroup/qbase/txs"
	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/txplugins/transfer/model"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/QOSGroup/qos/txs/transfer"
	"github.com/gin-gonic/gin"
)

type TxTransferPlugin struct{}

func (ttp TxTransferPlugin) DbInit(driveName string, db *sql.DB) error {

	return DbInit(driveName, db)
}

func (ttp TxTransferPlugin) DbClear(driveName string, db *sql.DB) error {
	return DbClear(driveName, db)
}

type TransferType int

const (
	Sender TransferType = iota
	Reciever
)

func (ut TransferType) String() string {
	switch ut {
	case Sender:
		return "转出"
	case Reciever:
		return "转入"
	default:
		return ""
	}
}

func (ttp TxTransferPlugin) Parse(chainID string, height int64, hash string, t time.Time, itx qbasetxs.ITx) (typeName string, hit bool, err error) {
	tt, ok := itx.(*transfer.TxTransfer)
	if !ok {
		return "", false, nil
	}
	log.Printf("transfer.TxTransfer:%+v", t)

	for _, v := range tt.Senders {
		saveTransItem(chainID, height, hash, Sender, t, v)
	}

	for _, v := range tt.Receivers {
		saveTransItem(chainID, height, hash, Reciever, t, v)
	}

	return "TxTransfer", true, nil
}

func (ttp TxTransferPlugin) Type() string {
	return "TxTransfer"
}

func (ttp TxTransferPlugin) RegisterGin(r *gin.Engine) {
	AccountTxsGinRegister(r)
}

func saveTransItem(chainID string, height int64, hash string, ut TransferType, t time.Time, item transfer.TransItem) error {
	if !item.QOS.IsZero() {
		saveTx(chainID, height, hash, item.Address.String(), "QOS", item.QOS.String(), ut, t)
	}

	for _, v := range item.QSCs {
		saveTx(chainID, height, hash, item.Address.String(), v.GetName(), v.GetAmount().String(), ut, t)
	}

	return nil
}

func saveTx(chainID string, height int64, hash string, address string, coin string, amount string, ut TransferType, t time.Time) error {
	tt := &model.TxTransfer{}
	tt.ChainID = utils.NullString(chainID)
	tt.Height = utils.NullInt64(height)
	tt.Hash = utils.NullString(hash)
	tt.Address = utils.NullString(address)
	tt.Coin = utils.NullString(coin)
	tt.Amount = utils.NullString(amount)
	tt.Type = utils.NullInt64(int64(ut))
	tt.Time = utils.NullTime(t)

	return tt.Insert(db.Db)
}
