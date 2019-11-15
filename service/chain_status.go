// Copyright 2018 The QOS Authors

package service

import (
	"github.com/QOSGroup/qmoon/lib/qos"
	"time"

	"github.com/QOSGroup/qmoon/lib/cache"
	"github.com/QOSGroup/qmoon/types"
)

const chainStatusCache = "ChainStatusCache"

func (n Node) ChainStatus(cached bool) (*types.ResultStatus, error) {
	result := &types.ResultStatus{}
	if cached {
		d, ok := cache.Get(chainStatusCache)
		if ok {
			if v, okk := d.(*types.ResultStatus); okk {
				return v, nil
			}
		}
	}


	status, err := qos.NewQosCli("").QueryStatus(n.BaseURL)

	cs, err1 := n.ConsensusState()
	result.Height = status.SyncInfo.LatestBlockHeight
	if err1 != nil {
		result.ConsensusState = &types.ResultConsensusState{}
	} else {
		result.ConsensusState = cs
		// latestHeight,_ = strconv.ParseInt(cs.Height, 10, 64)
	}

	blc, err := n.BlockByHeight(result.Height)
	if err == nil {
		result.Block = blc
	}

	vs, err2 := n.Validators()
	if err2 == nil {
		result.TotalValidators = int64(len(vs))
	}

	d, err := n.BlockTimeAvg(100)
	if err == nil {
		result.BlockTimeAvg = d.String()
	}

	// lb, err3 := n.LatestBlock()
	lb, err3 := n.BlockByHeight(result.Height)
	if err3 == nil {
		result.TotalTxs = lb.TotalTxs
		result.Proposer = lb.Proposer
		result.Votes = lb.Votes
	}
	result.ConsensusState.ChainID = n.ChainID
	if err3 == nil && err2 == nil {
		cache.Set(chainStatusCache, result, time.Second*1)
	}

	return result, nil
}
