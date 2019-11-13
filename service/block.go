// Copyright 2018 The QOS Authors

// Package pkg comments for pkg block
// service 块相关数据封装
package service

import (
	"errors"
	"time"

	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
)

func convertToBlock(mb *models.Block) *types.ResultBlockBase {
	return &types.ResultBlockBase{
		ID:             mb.Id,
		ChainID:        mb.ChainId,
		Height:         mb.Height,
		NumTxs:         mb.NumTxs,
		TotalTxs:       mb.TotalTxs,
		Time:           types.ResultTime(mb.Time),
		DataHash:       mb.DataHash,
		ValidatorsHash: mb.ValidatorsHash,
	}
}

// Latest最新的块
func (n Node) LatestBlock() (*types.ResultBlockBase, error) {
	mbs, err := models.Blocks(n.ChainID, &models.BlockOption{Offset: 0, Limit: 1})
	if err != nil {
		return nil, err
	}

	if len(mbs) == 0 {
		return nil, errors.New("not found")
	}
	latestblock := convertToBlock(mbs[0])
	proposer, err := models.ValidatorByAddress(n.ChainID, mbs[0].ProposerAddress)
	if err != nil {
		return nil, err
	}
	latestblock.Proposer = ConvertToValidator(proposer, latestblock.Height)
	latestVote, err := models.RetrieveVotesByHeight(n.ChainID, mbs[0].Height)
	latestblock.Votes = latestVote
	inf, err := models.InflationByHeight(n.ChainID, mbs[0].Height)
	if err != nil {
		latestblock.Inflation = "Not Available"
	} else {
		latestblock.Inflation = string(inf.Tokens)
	}
	return latestblock, nil
}

// Retrieve 块查询
func (n Node) RetrieveBlock(height int64) (*types.ResultBlockBase, error) {
	mbs, err := models.Blocks(n.ChainID, &models.BlockOption{Height: height, Offset: 0, Limit: 1})
	if err != nil {
		return nil, err
	}

	if len(mbs) != 1 {
		return nil, errors.New("not found")
	}

	block := convertToBlock(mbs[0])
	proposer, err := models.ValidatorByAddress(n.ChainID, mbs[0].ProposerAddress)
	if err != nil {
		return nil, err
	}
	block.Proposer = ConvertToValidator(proposer, height)
	vote, err := models.RetrieveVotesByHeight(n.ChainID, mbs[0].Height)
	block.Votes = vote
	inf, err := models.InflationByHeight(n.ChainID, mbs[0].Height)
	if err != nil {
		block.Inflation = "Not Available"
	} else {
		block.Inflation = string(inf.Tokens)
	}
	return block, err
}

// Search 块查询
func (n Node) Blocks(minHeight, maxHeight, offset, limit int64) ([]*types.ResultBlockBase, error) {
	mbs, err := models.Blocks(n.ChainID, &models.BlockOption{MinHeight: minHeight, MaxHeight: maxHeight, Offset: int(offset), Limit: int(limit)})
	if err != nil {
		return nil, err
	}

	var res []*types.ResultBlockBase
	for _, v := range mbs {
		blc := convertToBlock(v)
		proposer, err := models.ValidatorByAddress(n.ChainID, v.ProposerAddress)
		if err != nil {
			return nil, err
		}
		blc.Proposer = ConvertToValidator(proposer, maxHeight)
		vote, err := models.RetrieveVotesByHeight(n.ChainID, mbs[0].Height)
		blc.Votes = vote
		inf, err := models.InflationByHeight(n.ChainID, mbs[0].Height)
		if err != nil {
			blc.Inflation = "Not Available"
		} else {
			blc.Inflation = string(inf.Tokens)
		}
		res = append(res, blc)
	}

	return res, err
}

// Search 最近N块平均打快时间
func (n Node) BlockTimeAvg(blockNum int) (time.Duration, error) {
	mbs, err := models.Blocks(n.ChainID, &models.BlockOption{Offset: 0, Limit: blockNum})
	if err != nil {
		return 0, err
	}

	if len(mbs) <= 1 {
		return 0, nil
	}

	var duration int64
	num := int64(0)
	for k := 1; k < len(mbs); k++ {
		duration += int64(mbs[k-1].Time.Sub(mbs[k].Time))
		num++
	}

	return time.Duration(duration / num), err
}

// HasTx 有交易的块 这是干啥的
func (n Node) HasTxBlocks(minHeight, maxHeight int64) ([]*types.ResultBlockBase, error) {
	mbs, err := models.Blocks(n.ChainID, &models.BlockOption{
		MinHeight: minHeight, MaxHeight: maxHeight, NumTxs: 1})
	if err != nil {
		return nil, err
	}

	var res []*types.ResultBlockBase
	for _, v := range mbs {
		res = append(res, convertToBlock(v))
	}

	return res, err
}

func (n Node) CreateBlock(b *types.Block) error {
	block := &models.Block{}
	block.Height = b.Header.Height
	block.NumTxs = b.Header.NumTxs
	block.TotalTxs = b.Header.TotalTxs
	block.Time = b.Header.Time
	block.DataHash = b.Header.DataHash
	block.ValidatorsHash = b.Header.ValidatorsHash
	block.ProposerAddress = b.Header.ProposerAddress
	if err := block.Insert(n.ChainID); err != nil {
		return err
	}

	return nil
}

func (n Node) CreateEvidence(b *types.Block) error {
	if b.EvidenceList.Evidences == nil || len(b.EvidenceList.Evidences) == 0 {
		return nil
	}

	return models.CreateEvidences(n.ChainID, b.EvidenceList)
}
