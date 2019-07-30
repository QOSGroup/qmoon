// Copyright 2019 The QOS Authors

package service

import (
	"context"
	"time"

	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/sirupsen/logrus"
)

type NetworkSpider struct {
	chainID string
	remote  string

	checked map[string]*models.Network
	wait    map[string]*models.Network
}

func NewNetworkSpider(chainID string, remote string) *NetworkSpider {
	var ns NetworkSpider

	ns.chainID = chainID
	ns.remote = remote
	ns.checked = make(map[string]*models.Network)
	ns.wait = make(map[string]*models.Network)

	return &ns
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

func (n Node) NetworkSpider(ctx context.Context) {
	spider := NewNetworkSpider(n.ChanID, n.BaseURL)
	spider.wait[n.BaseURL] = &models.Network{
		Remote: n.BaseURL,
	}

	ns, err := models.Networks(n.ChanID, &models.NetworkOption{})
	if err == nil {
		for _, v := range ns {
			spider.wait[v.Remote] = &models.Network{
				Remote: v.Remote,
			}
		}
	}

	spider.Run()
	spider.Check()
	spider.Save()
}

func (ns *NetworkSpider) Run() {

}

func (ns *NetworkSpider) Save() {
	for _, v := range ns.checked {
		err := models.CreateNetworkOrUpdate(ns.chainID, &models.Network{
			Remote:     v.Remote,
			Delay:      v.Delay,
			Status:     v.Status,
			LastTestAt: v.LastTestAt,
		})
		if err != nil {
			logrus.Warnf("Network:%+v save error:%s", v, err.Error())
		}

	}

}

func (ns *NetworkSpider) Check() {
	for _, v := range ns.wait {
		ok, delay, err := utils.CheckHttpHealthy(v.Remote, time.Second*3)
		if err != nil {
			v.Status = models.NetworkStatusInvalid
		} else {
			if ok {
				v.Status = models.NetworkStatusValid
				v.Delay = delay.Nanoseconds()
			} else {
				v.Status = models.NetworkStatusInvalid
			}
		}
		v.LastTestAt = time.Now()

		ns.checked[v.Remote] = v
	}

}
