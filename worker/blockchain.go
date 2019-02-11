// Copyright 2018 The QOS Authors

package worker

import (
	"context"
	"sync"
	"time"

	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/service/syncer"
)

var syncBlockIsRunning bool

func SyncAllNodeBlock() {
	if syncBlockIsRunning {
		return
	}

	syncBlockIsRunning = true
	defer func() {
		syncBlockIsRunning = false
	}()

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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	for _, v := range needSync {
		wg.Add(1)
		go func(node *service.Node) {
			defer wg.Done()
			syncer.NewSyncer(v).Block(ctx)
		}(v)
	}

	wg.Wait()
	cancel()
}
