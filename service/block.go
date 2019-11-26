// Copyright 2018 The QOS Authors

// Package pkg comments for pkg block
// service 块相关数据封装
package service

import (
	"errors"
	tmlib "github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/lib/qos"

	"strconv"
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
func (n Node) LatestBlock() (result *types.ResultBlockBase, err error) {
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
	latestblock.Votes, _ = models.RetrieveVotesByHeight(n.ChainID, mbs[0].Height)
	inf, err := models.InflationByHeight(n.ChainID, mbs[0].Height)
	if err != nil {
		latestblock.Inflation = "Not Available"
	} else {
		latestblock.Inflation = strconv.FormatInt(inf.Tokens, 10)
	}
	return latestblock, nil
}

func (n Node) LatestBlockHeight() (height int64, err error) {
	status, err := qos.NewQosCli("").QueryStatus(n.BaseURL)
	if err != nil || status == nil {
		return 0, err
	}
	height = status.LatestBlockHeight
	return
}

func (n Node) LatestBlockFromCli() (result *types.ResultBlockBase, err error) {
	height, _ := n.LatestBlockHeight()
	if err != nil || height == 0 {
		return nil, err
	}
	b, err := tmlib.TendermintClient(n.BaseURL).RetrieveBlock(&height)
	// b, err := n.BlockByHeight(height)
	if err != nil {
		return nil, err
	}
	n.CreateBlock(b)
	result = &types.ResultBlockBase{
		ChainID:b.Header.ChainID,
		Height:b.Header.Height,
		NumTxs:b.Header.NumTxs,
		TotalTxs:b.Header.TotalTxs,
		Time:types.ResultTime(b.Header.Time),
		DataHash: b.Header.DataHash,
		ValidatorsHash:b.Header.ValidatorsHash,
		CreatedAt: types.ResultTime(b.Header.Time),
	}
	proposer, err := models.ValidatorByAddress(n.ChainID,b.Header.ProposerAddress)
	if err != nil {
		return
	}
	result.Proposer = ConvertToValidator(proposer, height)
	vote, err := models.RetrieveVotesByHeight(n.ChainID,height)
	result.Votes = vote

	result.Inflation = "995474"
	inf, err := models.InflationByHeight(n.ChainID, height)
	if err == nil {
		result.Inflation = strconv.FormatInt(inf.Tokens, 10)
	}
	return
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

	block.Inflation = "995474"
	inf, err1 := models.InflationByHeight(n.ChainID, height)
	if err1 == nil {
		block.Inflation = strconv.FormatInt(inf.Tokens, 10)
	}
	return block, err
}

func (n Node) BlockByHeight(height int64) (*types.ResultBlockBase, error) {
	block, err := qos.NewQosCli("").QueryBlockByHeight(n.BaseURL, height)
	//fmt.Println("Got Block: ", height, " @ ", block)
	if err != nil || block == nil {
		return nil, err
	}
	//blockM := models.Block{Height:block.Height}
	//err = blockM.InsertIfNotExist(n.ChainID)
	numTxs := int64(0)
	totalTxs := int64(0)
	time:=types.ResultTime{}
	dataHash := ""
	validatorsHash := ""
	createAt := types.ResultTime{}
	if block.Block != nil {
		numTxs = block.Block.Header.NumTxs
		totalTxs = block.Block.Header.TotalTxs
		time = types.ResultTime(block.Block.Header.Time)
		dataHash = block.Block.DataHash.String()
		validatorsHash = block.Block.ValidatorsHash.String()
		createAt = types.ResultTime(block.Block.Header.Time)
	}
	resultBlock := types.ResultBlockBase {
		ChainID: block.Block.Header.ChainID,
		Height: height,
		NumTxs: numTxs,
		TotalTxs: totalTxs,
		Time: time,
		DataHash: dataHash,
		ValidatorsHash: validatorsHash,
		CreatedAt: createAt,
	}
	//fmt.Println("finding Proposer by address:", block.Block.Header.ProposerAddress.String())
	proposer, err0 := models.ValidatorByAddress(n.ChainID, block.Block.Header.ProposerAddress.String())
	if err0 == nil {
		resultBlock.Proposer = ConvertToValidator(proposer, height)
	}
	resultBlock.Votes, _ = models.RetrieveVotesByHeight(n.ChainID, height)
	resultBlock.Inflation = "995474"
	inf, err1 := models.InflationByHeight(n.ChainID, height)
	if err1 == nil {
		resultBlock.Inflation = strconv.FormatInt(inf.Tokens, 10)
	}
	return &resultBlock, err
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

		blc.Inflation = "995474"
		inf, err := models.InflationByHeight(n.ChainID, mbs[0].Height)
		if err == nil {
			blc.Inflation = strconv.FormatInt(inf.Tokens, 10)
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
	if err := block.InsertIfNotExist(n.ChainID); err != nil {
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
