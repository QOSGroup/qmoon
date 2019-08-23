// Copyright 2018 The QOS Authors

package transfer

import (
	"database/sql"
	"log"
	"time"

	qbasetxs "github.com/QOSGroup/qbase/txs"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
	transfer "github.com/QOSGroup/qos/module/bank/txs"
	transfertypes "github.com/QOSGroup/qos/module/bank/types"
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

func (ttp TxTransferPlugin) Parse(blockHeader types.BlockHeader, itx qbasetxs.ITx) (typeName string, hit bool, err error) {
	tt, ok := itx.(*transfer.TxTransfer)
	if !ok {
		return "", false, nil
	}
	log.Printf("transfer.TxTransfer:%+v", blockHeader.Time)

	//d, err := json.Marshal(tt)
	//if err != nil {
	//	return "", true, err
	//}

	for _, v := range tt.Senders {
		saveTransItem(blockHeader, Sender, v)
	}

	for _, v := range tt.Receivers {
		saveTransItem(blockHeader, Reciever, v)
	}

	return "TxTransfer", true, nil
}

func (ttp TxTransferPlugin) Type() string {
	return "TxTransfer"
}

func (ttp TxTransferPlugin) Doctor() error {
	return nil
}

func (ttp TxTransferPlugin) RegisterGin(r *gin.Engine) {
	AccountTxsGinRegister(r)
}

type Transfer struct {
	ChainID  string
	Height   int64
	Hash     string
	Sender   string
	Reciever string
	Coin     string
	Amount   string
	Time     time.Time
}

func saveTransItem(blockHeader types.BlockHeader, ut TransferType, item transfertypes.TransItem) error {
	if !item.QOS.IsZero() {
		saveTx(blockHeader.ChainID, blockHeader.Height, blockHeader.DataHash, item.Address.String(),
			"QOS", item.QOS.String(), ut, blockHeader.Time)
	}

	for _, v := range item.QSCs {
		saveTx(blockHeader.ChainID, blockHeader.Height, blockHeader.DataHash, item.Address.String(),
			v.GetName(), v.GetAmount().String(), ut, blockHeader.Time)
	}

	return nil
}

func saveTx(chainID string, height int64, hash string, address string, coin string, amount string, ut TransferType, t time.Time) error {
	tt := &models.TxTransfer{}
	tt.Height = height
	tt.Hash = hash
	tt.Address = address
	tt.Coin = coin
	tt.Amount = amount
	tt.Type = int(ut)
	tt.Time = t

	return tt.Insert(chainID)
}
