// Copyright 2018 The QOS Authors

package worker

import (
	"sync"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service"
)

var syncValidatorIsRunning bool

func SyncAllNodeValidator() {
	if syncValidatorIsRunning {
		return
	}

	syncValidatorIsRunning = true
	defer func() {
		syncValidatorIsRunning = false
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
		wg.Add(1)
		go func(node service.Node) {
			defer wg.Done()
			SyncValidator(node)
		}(v)
	}

	wg.Wait()
}

// SyncValidator 同步validator
func SyncValidator(node service.Node) error {
	tmc := lib.TendermintClient(node.BaseURL)

	v, err := service.GetQOSValidator(tmc, 0, node)
	if err != nil {
		return err
	}

	service.CreateValidator(node.ChanID, v)

	return nil
}
