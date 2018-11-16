// Copyright 2018 The QOS Authors

package worker

import (
	"log"
	"sync"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service"
)

var syncConsensusStateIsRunning bool

func SyncAllConsensusState() {
	if syncConsensusStateIsRunning {
		return
	}

	syncConsensusStateIsRunning = true
	defer func() {
		syncConsensusStateIsRunning = false
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
		wg.Add(1)
		go func(chanID, nodeUrl string) {
			defer wg.Done()
			SyncConsensusState(chanID, nodeUrl)
		}(chanID, nodeUrl)
	}

	wg.Wait()
}

// SyncConsensusState 同步共识过程
func SyncConsensusState(chanID, remote string) error {
	//if ok := service.SyncLock(chanID + "-block"); !ok {
	//	return nil
	//}
	//defer service.SyncUnlock(chanID + "-block")

	tmc := lib.TendermintClient(remote)
	cs, err := tmc.ConsensusState()
	if err != nil {
		return err
	}

	if cs == nil {
		return nil
	}

	err = service.CreateConsensusState(chanID, cs)
	if err != nil {
		log.Printf("CreateConsensusState error:[%s]. chaidID:%s", err.Error(), chanID)
	}

	return nil
}
