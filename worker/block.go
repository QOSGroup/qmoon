// Copyright 2018 The QOS Authors

package worker

import (
	"context"
	"sync"

	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/service/syncer"
)

func syncAllNodeBlock() {
	wg := &sync.WaitGroup{}

	nodes, err := service.AllNodes()
	if err != nil {
		return
	}

	needSync := map[string]*service.Node{}
	for _, v := range nodes {
		if v.ChanID != "" && v.BaseURL != "" {
			needSync[v.ChanID] = v
		}
	}

	//ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	ctx := context.Background()
	for _, v := range needSync {
		wg.Add(1)
		go func(node *service.Node) {
			defer wg.Done()
			syncer.NewSyncer(node).BlockLoop(ctx)
		}(v)
	}

	wg.Wait()
	//cancel()
}
