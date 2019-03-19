// Copyright 2018 The QOS Authors

package service

import (
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
)

func convertToFee(mt *models.Fee) *types.ResultFee {
	cs, _ := types.ParseCoins(mt.Fee)

	return &types.ResultFee{
		Tx:        mt.Tx,
		Fee:       cs,
		GasWanted: mt.GasWanted,
		GasUsed:   mt.GasUsed,
	}
}

// Fee fee查询
func (n Node) Fee(tx string) (*types.ResultFee, error) {
	if tx == "" {
		tx = models.DefaultFeeTx
	}
	mt, err := models.RetrieveFeeByTx(n.ChanID, tx)
	if err != nil {
		return nil, err
	}

	return convertToFee(mt), err
}
