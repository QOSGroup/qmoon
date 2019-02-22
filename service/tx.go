// Copyright 2018 The QOS Authors

package service

import (
	"strings"

	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
)

func convertToTx(mt *models.Tx) *types.ResultTx {
	return &types.ResultTx{
		ChainID:   mt.ChainId,
		Hash:      mt.Hash,
		Height:    mt.Height,
		Index:     mt.Index,
		TxType:    mt.TxType,
		GasWanted: mt.GasWanted,
		GasUsed:   mt.GasUsed,
		Data:      []byte(mt.JsonTx),
		Time:      types.ResultTime(mt.Time),
		TxStatus:  types.TxStatus(mt.TxStatus).String(),
	}
}

const maxLimit = 20

// TxsByAddress 交易查询
func (n Node) TxsByAddress(address string, minHeight, maxHeight int64) ([]*types.ResultTx, error) {
	mbs, err := models.Txs(n.ChanID, &models.TxOption{
		MinHeight: minHeight, MaxHeight: maxHeight, Address: address, Offset: 0, Limit: maxLimit})
	if err != nil {
		return nil, err
	}

	var res []*types.ResultTx
	for _, v := range mbs {
		res = append(res, convertToTx(v))
	}

	return res, err
}

// List 交易查询
func (n Node) Txs(minHeight, maxHeight int64) ([]*types.ResultTx, error) {
	mbs, err := models.Txs(n.ChanID, &models.TxOption{MinHeight: minHeight, MaxHeight: maxHeight, Offset: 0, Limit: maxLimit})
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
func (n Node) Tx(height, index int64) (*types.ResultTx, error) {
	mt, err := models.TxByHeightIndex(n.ChanID, height, index)
	if err != nil {
		return nil, err
	}

	return convertToTx(mt), err
}

// Search 交易查询
func (n Node) TxByHash(hash string) (*types.ResultTx, error) {
	mt, err := models.TxByHash(n.ChanID, strings.ToUpper(hash))
	if err != nil {
		return nil, err
	}

	return convertToTx(mt), err
}
