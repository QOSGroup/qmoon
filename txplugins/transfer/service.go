// Copyright 2018 The QOS Authors

package transfer

import (
	"fmt"
	"strings"
	"time"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/txplugins/transfer/model"
)

type TxTransfer struct {
	ID      int64        `json:"id"`       // id
	ChainID string       `json:"chain_id"` // chain_id
	Height  int64        `json:"height"`   // height
	Hash    string       `json:"hash"`     // hash
	Address string       `json:"address"`  // address
	Coin    string       `json:"coin"`     // coin
	Amount  string       `json:"amount"`   // amount
	Type    TransferType `json:"type"`     // type
	Time    time.Time    `json:"time"`     // time
}

func converToTxTransfer(mtt model.TxTransfer) TxTransfer {
	return TxTransfer{
		ID:      mtt.ID,
		ChainID: mtt.ChainID.String,
		Height:  mtt.Height.Int64,
		Hash:    mtt.Hash.String,
		Address: mtt.Address.String,
		Coin:    mtt.Coin.String,
		Amount:  mtt.Amount.String,
		Type:    TransferType(mtt.Type.Int64),
		Time:    mtt.Time.Time,
	}
}

func ListByAddress(address string, offset, limint int64) ([]TxTransfer, error) {
	var res []TxTransfer

	var wheres []string
	wheres = append(wheres, fmt.Sprintf(" %s = '%s' ", "address", address))

	mtts, err := model.TxTransferFilter(db.Db, strings.Join(wheres, " and "), " sort by time desc ", offset, limint)
	if err != nil {
		return nil, err
	}

	for _, mtt := range mtts {
		res = append(res, converToTxTransfer(*mtt))
	}

	return res, nil
}
