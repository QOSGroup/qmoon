// Copyright 2018 The QOS Authors

package service

import (
	"github.com/QOSGroup/qmoon/lib/cache"
	"github.com/QOSGroup/qmoon/lib/qos"
	"github.com/QOSGroup/qmoon/types"
	"time"
)

const LatestHeightKey = "latest_height_key"

func (n Node) ChainStatus() (*types.ResultStatus, error) {
	result := &types.ResultStatus{}
	//if cached {
	//	d, ok := cache.Get(chainStatusCache)
	//	if ok {
	//		if v, okk := d.(*types.ResultStatus); okk {
	//			return v, nil
	//		}
	//	}
	//}

	status, err := qos.NewQosCli("").QueryStatus(n.BaseURL)
	if err != nil {
		return nil, err
	}

	//bl, err := n.LatestBlock();
	if status!=nil {
	//if err == nil {
		bl, err := n.BlockByHeight(status.SyncInfo.LatestBlockHeight)
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

	}

	cs, err1 := n.ConsensusState()
	if err1 != nil {
		result.ConsensusState = &types.ResultConsensusState{}
	} else {
		result.ConsensusState = cs
		// latestHeight,_ = strconv.ParseInt(cs.Height, 10, 64)
	}



	d, err := n.BlockTimeAvg(100)
	if err == nil {
		result.BlockTimeAvg = d.String()
	}

	// lb, err3 := n.LatestBlock()

	result.ConsensusState.ChainID = n.ChainID

	// cache.Set(chainStatusCache, result, time.Second*1)

	return result, nil
}
