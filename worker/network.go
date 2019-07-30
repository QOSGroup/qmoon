// Copyright 2018 The QOS Authors

package worker

import (
	"context"
	"sync"

	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/service/syncer"
	"github.com/sirupsen/logrus"
)

var syncNetworkIsRunning bool

func SyncAllNodeNetwork() {
	if syncNetworkIsRunning {
		return
	}

	syncNetworkIsRunning = true
	defer func() {
		syncNetworkIsRunning = false
	}()

	wg := sync.WaitGroup{}

	nodes, err := service.AllNodes()
	if err != nil {
		logrus.Warnf("SyncAllNodeNetwork AllNodes err:%s", err.Error())
		return
	}

	logrus.Warnf("SyncAllNodeNetwork AllNodes nodes:%+v", nodes)

	for _, v := range nodes {
		wg.Add(1)
		go func(node *service.Node) {
			defer wg.Done()
			network(node)
		}(v)
	}

	wg.Wait()
}

func network(node *service.Node) {
	_ = syncer.NewSyncer(node).RpcPeers(context.Background())
	node.NetworkSpider(context.Background())
}
