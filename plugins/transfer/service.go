// Copyright 2018 The QOS Authors

package transfer

import (
	"time"

	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
)

type TxTransfer struct {
	ID       int64        `json:"id"`       // id
	ChainID  string       `json:"chain_id"` // chain_id
	Height   int64        `json:"height"`   // height
	Hash     string       `json:"hash"`     // hash
	Address  string       `json:"address"`  // address
	Coin     string       `json:"coin"`     // coin
	Amount   string       `json:"amount"`   // amount
	Type     TransferType `json:"type"`     // type
	TxStatus string       `json:"txStatus"`
	Time     time.Time    `json:"time"` // time
}

func converToTxTransfer(mtt models.TxTransfer) *TxTransfer {
	return &TxTransfer{
		ID:       mtt.Id,
		Height:   mtt.Height,
		Hash:     mtt.Hash,
		Address:  mtt.Address,
		Coin:     mtt.Coin,
		Amount:   mtt.Amount,
		Type:     TransferType(mtt.Type),
		TxStatus: types.TxStatus(mtt.TxStatus).String(),
		Time:     mtt.Time,
	}
}

type SearchOpt struct {
	Coin string
}

func ListByAddress(chainId, address string, offset, limint int, opt *SearchOpt) ([]*TxTransfer, error) {
	mtts, err := models.TxTransfers(chainId, &models.TxTransferOption{
		Address: address,
		Coin:    opt.Coin,
		Offset:  offset,
		Limit:   limint,
	})
	if err != nil {
		return nil, err
	}

	var res []*TxTransfer
	for _, mtt := range mtts {
		res = append(res, converToTxTransfer(*mtt))
	}

	return res, nil
}
