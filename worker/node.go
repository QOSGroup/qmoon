// Copyright 2018 The QOS Authors

package worker

import (
	"context"
	"sync"
	"time"

	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/service/syncer"
)

// SyncAllNodePeer 同步peer节点信息
func SyncAllNodePeer() {
	wg := sync.WaitGroup{}

	nodes, err := service.AllNodes()
	if err != nil {
		return
	}

	needSync := make(map[string]*service.Node)
	for _, node := range nodes {
		if node.ChanID != "" && node.BaseURL != "" {
			needSync[node.ChanID] = node
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	for _, v := range needSync {
		wg.Add(1)
		go func(node *service.Node) {
			defer wg.Done()
			syncer.NewSyncer(node).Peer(ctx)
		}(v)
	}

	wg.Wait()
	cancel()
}
