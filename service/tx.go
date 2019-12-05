// Copyright 2018 The QOS Authors

package service

import (
	"encoding/json"
	"strings"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
)

func convertToTx(mt *models.Tx, address string) *types.ResultTx {
	res := &types.ResultTx{
		ChainID: mt.ChainId,
		Hash:    mt.Hash,
		Height:  mt.Height,
		Index:   mt.Index,
		TxType:  mt.TxType,
		// TxTypeCN:  types.TxCN(mt.TxType, mt.JsonTx, address),
		GasWanted: mt.GasWanted,
		GasUsed:   mt.GasUsed,
		Fee:       mt.Fee,
		Time:      types.ResultTime(mt.Time),
		TxStatus:  types.TxStatus(mt.TxStatus).String(),
		Status:    mt.TxStatus,
		Log:       mt.Log,
		ITxs:      make([]json.RawMessage, 0),
	}

	for _, v := range mt.ITxs {
		j, _ := json.Marshal(v)
		res.ITxs = append(res.ITxs, j)
	}

	return res
}

const maxLimit = 20

// TxsByAddress 交易查询
//func (n Node) TxsByAddress(address string, tx string, minHeight, maxHeight int64, offset, limit int) ([]*types.ResultTx, error) {
//	mbs, err := models.Txs(n.ChainID, &models.TxOption{
//		TxType:    tx,
//		MinHeight: minHeight, MaxHeight: maxHeight, Address: address, Offset: offset, Limit: limit})
//	if err != nil {
//		return nil, err
//	}
//
//	var res []*types.ResultTx
//	for _, v := range mbs {
//		res = append(res, convertToTx(v, address))
//	}
//
//	return res, err
//}

// List 交易查询
func (n Node) Txs(minHeight, maxHeight, offset, limit int64) ([]*types.ResultTx, error) {
	mbs, err := models.Txs(n.ChainID, &models.TxOption{MinHeight: minHeight, MaxHeight: maxHeight, Offset: int(offset), Limit: int(limit)})
	if err != nil {
		return nil, err
	}

	var res []*types.ResultTx
	for _, v := range mbs {
		v.ITxs, err = models.ITxByHash(n.ChainID, v.Hash)
		if err != nil {
			return nil, err
		}
		res = append(res, convertToTx(v, ""))
	}

	return res, err
}

// Search 交易查询
func (n Node) Tx(height, index int64) (*types.ResultTx, error) {
	mt, err := models.TxByHeightIndex(n.ChainID, height, index)
	if err != nil {
		return nil, err
	}

	return convertToTx(mt, ""), err
}

// Search 交易查询
func (n Node) TxByHash(hash string) (*types.ResultTx, error) {
	mt, err := models.TxByHash(n.ChainID, strings.ToUpper(hash))
	if err != nil {
		return nil, err
	}
	return convertToTx(mt, ""), err
}

func (n Node) TxsByAddress(address string, minHeight int64, maxHeight int64, offset int, limit int, txTypes... string)(result *types.ResultTxs, err error) {
	result = &types.ResultTxs{Txs:make([]*types.ResultTx,0)}
	txs, err := models.TxByAddress(n.ChainID, address, minHeight, maxHeight, offset, limit)
	if err != nil {
		return nil, err
	}
	for _, tx := range txs {
		tx.ITxs, err = models.ITxByHash(n.ChainID, tx.Hash)
		if err != nil {
			return nil, err
		}
		result.Txs = append(result.Txs, convertToTx(tx, ""))
	}
	return
}

// TxSend
func (n Node) TxSend(remote string, txData []byte) (*types.TxSendResult, error) {
	cli := lib.TendermintClient(remote)
	res, err := cli.BroadcastTxSync(txData)
	if err != nil {
		return nil, err
	}

	// TODO 检查BroadcastTxSync的返回值
	// xxx

	return &types.TxSendResult{
		Hash: res.Hash.String(),
	}, err
}
