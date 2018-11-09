// Copyright 2018 The QOS Authors

// Package pkg comments for pkg block
// block 块相关数据封装
package block

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/gin-gonic/gin/json"
	tmtypes "github.com/tendermint/tendermint/rpc/core/types"
)

func convertToBlock(mb *model.Block) *types.ResultBlockBase {
	return &types.ResultBlockBase{
		ID:             mb.ID,
		ChainID:        mb.ChainID.String,
		Height:         mb.Height.Int64,
		NumTxs:         mb.NumTxs.Int64,
		Time:           mb.Time.Time,
		DataHash:       mb.DataHash.String,
		ValidatorsHash: mb.ValidatorsHash.String,
		CreatedAt:      mb.CreatedAt.Time,
	}
}

const sqlOrder = " order by height desc "

// Latest最新的块
func Latest(chainID string) (*types.ResultBlockBase, error) {
	var wheres []string
	wheres = append(wheres, fmt.Sprintf(" chain_id='%s' ", chainID))

	mbs, err := model.BlockFilter(db.Db, strings.Join(wheres, " and ")+sqlOrder, 0, 1)
	if err != nil {
		return nil, err
	}

	if len(mbs) == 0 {
		return nil, errors.New("not found")
	}

	return convertToBlock(mbs[0]), nil
}

// Retrieve 块查询
func Retrieve(chainID string, mheight int64) (*types.ResultBlockBase, error) {
	var wheres []string
	wheres = append(wheres, fmt.Sprintf(" chain_id = '%s' ", chainID))
	wheres = append(wheres, fmt.Sprintf(" height = %d ", mheight))

	mbs, err := model.BlockFilter(db.Db, strings.Join(wheres, " and ")+sqlOrder, 0, -1)
	if err != nil {
		return nil, err
	}

	if len(mbs) != 1 {
		return nil, errors.New("not found")
	}

	return convertToBlock(mbs[0]), err
}

// Search 块查询
func Search(chainID string, minHeight, maxHeight int64) ([]*types.ResultBlockBase, error) {
	var wheres []string
	wheres = append(wheres, fmt.Sprintf(" chain_id = '%s' ", chainID))
	wheres = append(wheres, fmt.Sprintf(" height >= %d ", minHeight))
	wheres = append(wheres, fmt.Sprintf(" height <= %d ", maxHeight))

	mbs, err := model.BlockFilter(db.Db, strings.Join(wheres, " and ")+sqlOrder, 0, -1)
	if err != nil {
		return nil, err
	}

	var res []*types.ResultBlockBase
	for _, v := range mbs {
		res = append(res, convertToBlock(v))
	}

	return res, err
}

// HasTx 有交易的块
func HasTx(chainID string, minHeight, maxHeight int64) ([]*types.ResultBlockBase, error) {
	var wheres []string
	wheres = append(wheres, fmt.Sprintf(" chain_id = '%s' ", chainID))
	wheres = append(wheres, fmt.Sprintf(" height >= %d ", minHeight))
	wheres = append(wheres, fmt.Sprintf(" height <= %d ", maxHeight))
	wheres = append(wheres, fmt.Sprintf(" num_txs != 0 "))

	mbs, err := model.BlockFilter(db.Db, strings.Join(wheres, " and ")+sqlOrder, 0, -1)
	if err != nil {
		return nil, err
	}

	var res []*types.ResultBlockBase
	for _, v := range mbs {
		res = append(res, convertToBlock(v))
	}

	return res, err
}

func Save(b *tmtypes.ResultBlock) error {
	d, err := json.Marshal(b)
	if err != nil {
		return err
	}
	mtb, err := model.TmBlockByChainIDHeight(db.Db, utils.NullString(b.BlockMeta.Header.ChainID),
		utils.NullInt64(b.BlockMeta.Header.Height))
	if err != nil {
		if err == sql.ErrNoRows {
			mtb = &model.TmBlock{
				ChainID:   utils.NullString(b.BlockMeta.Header.ChainID),
				Height:    utils.NullInt64(b.BlockMeta.Header.Height),
				Data:      utils.NullString(string(d)),
				CreatedAt: utils.NullTime(time.Now()),
			}
			err = mtb.Insert(db.Db)
			if err != nil {
				return err
			}
		}
	}

	mb, err := model.BlockByChainIDHeight(db.Db, utils.NullString(b.BlockMeta.Header.ChainID),
		utils.NullInt64(b.BlockMeta.Header.Height))
	if err != nil {
		if err == sql.ErrNoRows {
			mb = &model.Block{
				ChainID:        utils.NullString(b.BlockMeta.Header.ChainID),
				Height:         utils.NullInt64(b.BlockMeta.Header.Height),
				NumTxs:         utils.NullInt64(b.BlockMeta.Header.NumTxs),
				Time:           utils.NullTime(b.BlockMeta.Header.Time),
				DataHash:       utils.NullString(b.BlockMeta.Header.DataHash.String()),
				ValidatorsHash: utils.NullString(b.BlockMeta.Header.ValidatorsHash.String()),
				CreatedAt:      utils.NullTime(time.Now()),
			}
			err = mb.Insert(db.Db)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
