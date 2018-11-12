// Copyright 2018 The QOS Authors

package worker

import (
	"log"
	"sync"
	"time"

	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/service/block"
	"github.com/tendermint/tendermint/rpc/client"
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

	needSync := map[string]string{}
	for _, v := range nodes {
		if v.ChanID != "" && v.BaseURL != "" {
			needSync[v.ChanID] = v.BaseURL
		}
	}

	for chanID, nodeUrl := range needSync {
		//log.Printf("--SyncAllNodeBlock start chanID:%s", chanID)
		wg.Add(1)
		go func(chanID, nodeUrl string) {
			defer wg.Done()
			SyncBlock(chanID, nodeUrl, 100, time.Second*5)
		}(chanID, nodeUrl)
	}

	wg.Wait()
}

// SyncBlock 同步块
// maxSync 一次最多同步块数量
// maxTime 程序最多运行时长，如果没有设置，则默认最长60s
func SyncBlock(chanID, remote string, maxSync int64, maxTime time.Duration) error {
	if ok := service.SyncLock(chanID + "-block"); !ok {
		return nil
	}
	defer service.SyncUnlock(chanID + "-block")

	if maxTime == 0 {
		maxTime = time.Minute
	}
	var start int64 = 1

	latest, err := block.Latest(chanID)
	log.Printf("--SyncBlock- latest:%+v, err:%+v", latest, err)
	if err == nil && latest != nil {
		start = latest.Height + 1
	}

	height := start
	syncNum := int64(0)
	tmc := client.NewHTTP(remote, "/websocket")

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

		err = service.CreateBlock(b)
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
