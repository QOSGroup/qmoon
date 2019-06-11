// Copyright 2018 The QOS Authors

package service

import (
	"encoding/json"
	"strings"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
	"github.com/tidwall/gjson"
)

func convertToTx(mt *models.Tx, address string) *types.ResultTx {
	res := &types.ResultTx{
		ChainID:   mt.ChainId,
		Hash:      mt.Hash,
		Height:    mt.Height,
		Index:     mt.Index,
		TxType:    mt.TxType,
		TxTypeCN:  types.TxCN(mt.TxType, mt.JsonTx, address),
		GasWanted: mt.GasWanted,
		GasUsed:   mt.GasUsed,
		Fee:       mt.Fee,
		Data:      []byte(mt.JsonTx),
		Time:      types.ResultTime(mt.Time),
		TxStatus:  types.TxStatus(mt.TxStatus).String(),
		Status:    mt.TxStatus,
		Log:       mt.Log,
	}

	m := make(map[string][]json.RawMessage)
	if strings.HasPrefix(string(res.Data), "{") {
		m[res.TxType] = []json.RawMessage{res.Data}
	} else {
		result := gjson.Parse(string(res.Data))
		result.ForEach(func(_, value gjson.Result) bool {
			if _, ok := m[value.Get("type").String()]; ok {
				m[value.Get("type").String()] = append(m[value.Get("type").String()], json.RawMessage(value.Get("data").String()))
			} else {
				m[value.Get("type").String()] = []json.RawMessage{json.RawMessage(value.Get("data").String())}
			}
			return true
		})
	}
	res.TxDetail = m

	return res
}

const maxLimit = 20

// TxsByAddress 交易查询
func (n Node) TxsByAddress(address string, tx string, minHeight, maxHeight int64, offset, limit int) ([]*types.ResultTx, error) {
	mbs, err := models.Txs(n.ChanID, &models.TxOption{
		TxType:    tx,
		MinHeight: minHeight, MaxHeight: maxHeight, Address: address, Offset: offset, Limit: limit})
	if err != nil {
		return nil, err
	}

	var res []*types.ResultTx
	for _, v := range mbs {
		res = append(res, convertToTx(v, address))
	}

	return res, err
}

// List 交易查询
func (n Node) Txs(minHeight, maxHeight, offset, limit int64) ([]*types.ResultTx, error) {
	mbs, err := models.Txs(n.ChanID, &models.TxOption{MinHeight: minHeight, MaxHeight: maxHeight, Offset: int(offset), Limit: int(limit)})
	if err != nil {
		return nil, err
	}

	var res []*types.ResultTx
	for _, v := range mbs {
		res = append(res, convertToTx(v, ""))
	}

	return res, err
}

// Search 交易查询
func (n Node) Tx(height, index int64) (*types.ResultTx, error) {
	mt, err := models.TxByHeightIndex(n.ChanID, height, index)
	if err != nil {
		return nil, err
	}

	return convertToTx(mt, ""), err
}

// Search 交易查询
func (n Node) TxByHash(hash string) (*types.ResultTx, error) {
	mt, err := models.TxByHash(n.ChanID, strings.ToUpper(hash))
	if err != nil {
		return nil, err
	}

	return convertToTx(mt, ""), err
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
