// Copyright 2019 The QOS Authors

package service

import (
	"time"

	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
)

type peer struct {
	Remote    string
	Delay     time.Duration
	Status    int // 0: 待验证 1: 有效  2: 无效
	CheckTime time.Time
}

type NetworkSpider struct {
	chainID string
	remote  string

	peers map[string]*peer
}

func (ns *NetworkSpider) Save() error {
	for _, v := range ns.peers {
		if v.Status != 0 {
			_ = models.CreateNetworkOrUpdate(ns.chainID, &models.Network{
				Remote:     v.Remote,
				Delay:      v.Delay.Nanoseconds(),
				Status:     v.Status,
				LastTestAt: v.CheckTime,
			})
		}
	}

	return nil
}

func (ns *NetworkSpider) Check() {
	for _, v := range ns.peers {
		if v.Status == 0 {
			ok, delay, err := utils.CheckHttpHealthy(v.Remote, time.Second*3)
			if err != nil {
				v.Status = models.NetworkStatusInvalid
			} else {
				if ok {
					v.Status = models.NetworkStatusValid
					v.Delay = delay
				} else {

				}
			}
		}
	}

}

func (n Node) ValidNodes() ([]*types.ChainNode, error) {
	ns, err := models.Networks(n.ChanID, &models.NetworkOption{Status: 1})
	if err != nil {
		return nil, err
	}

	var res []*types.ChainNode

	for _, v := range ns {
		res = append(res, &types.ChainNode{
			NodeID: v.ID,
			Delay:  v.Delay,
			Remote: v.Remote,
		})
	}

	return res, nil
}
