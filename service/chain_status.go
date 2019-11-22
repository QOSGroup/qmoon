// Copyright 2018 The QOS Authors

package service

import (
	"github.com/QOSGroup/qmoon/lib/cache"
	"github.com/QOSGroup/qmoon/lib/qos"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
	"strconv"
	"time"
)

const LatestHeightKey = "latest_height_key"

func (n Node) ChainStatus() (*types.ResultStatus, error) {
	result := &types.ResultStatus{}

	cs, err1 := n.ConsensusState()
	if err1 != nil {
		result.ConsensusState = &types.ResultConsensusState{}
	} else {
		result.ConsensusState = cs
	}

	latestheight, err := n.LatestBlockHeight()
	if err != nil || latestheight == 0{
		return nil, err
	}

	bl, err := n.BlockByHeight(latestheight)
	if err == nil {
		result.Height = bl.Height
		cache.Set(LatestHeightKey, result.Height,  time.Second*7)
		result.Block = bl
		result.TotalTxs = bl.TotalTxs
		result.Proposer = bl.Proposer
		result.Votes = bl.Votes
	}

	vs, err2 := n.Validators(result.Height)
	if err2 == nil {
		result.TotalValidators = int64(len(vs))
	}

	d, err := n.BlockTimeAvg(100)
	if err == nil {
		result.BlockTimeAvg = d.String()
	}

	bonded, err := models.TotalVotingPower(n.ChainID)
	if err == nil {
		result.BondedTokens = bonded
	}
	applied, err := qos.NewQosCli("").QueryApplied(n.BaseURL)
	if err == nil {
		appliedInt, err := strconv.ParseInt(applied, 10, 64)
		if err == nil {
			result.Circulation = appliedInt
		}
	}
	result.ConsensusState.ChainID = n.ChainID
	return result, nil
}
