// Copyright 2018 The QOS Authors

package worker

import (
	"log"
	"sync"
	"time"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/service/block"
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

	wg := sync.WaitGroup{}

	nodes, err := service.AllNodes()
	if err != nil {
		return
	}

	needSync := map[string]service.Node{}
	for _, v := range nodes {
		if v.ChanID != "" && v.BaseURL != "" {
			needSync[v.ChanID] = *v
		}
	}

	for _, v := range needSync {
		//log.Printf("--SyncAllNodeBlock start chanID:%s", chanID)
		wg.Add(1)
		go func(node service.Node) {
			defer wg.Done()
			SyncBlock(node, 100, time.Second*5)
		}(v)
	}

	wg.Wait()
}

// SyncBlock 同步块
// maxSync 一次最多同步块数量
// maxTime 程序最多运行时长，如果没有设置，则默认最长60s
func SyncBlock(node service.Node, maxSync int64, maxTime time.Duration) error {
	//if ok := service.SyncLock(chanID + "-block"); !ok {
	//	return nil
	//}
	//defer service.SyncUnlock(chanID + "-block")

	if maxTime == 0 {
		maxTime = time.Minute
	}
	var start int64 = 1

	latest, err := block.Latest(node.ChanID)
	//log.Printf("--SyncBlock- latest:%+v, err:%+v", latest, err)
	if err == nil && latest != nil {
		start = latest.Height + 1
	}

	height := start
	syncNum := int64(0)
	tmc := lib.TendermintClient(node.BaseURL)

	ticker := time.NewTicker(maxTime)
LOOP:
	for {
		select {
		case <-ticker.C:
			break LOOP
		default:
		}

		//log.Printf("--------height:%s %d", chanID, height)
		b, err := tmc.Block(&height)
		if err != nil {
			return err
		}

		if b == nil {
			break
		}

		v, err := service.GetQOSValidator(tmc, height, node)
		if err != nil {
			return err
		}

		if v == nil {
			break
		}

		err = service.CreateBlock(b, v)
		if err != nil {
			log.Printf("CreateBlock error:[%s]. chaidID:%s, height:%d", err.Error(), b.Block.ChainID, b.Block.Height)
		}

		height++
		syncNum++
		if syncNum >= maxSync {
			break
		}
	}

	return nil
}
