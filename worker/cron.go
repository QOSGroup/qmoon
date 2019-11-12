// Copyright 2018 The QOS Authors

package worker

import (
	"fmt"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/service/syncer"
	"os"
	"strings"
	"time"
)

func Start() {
	//先删除自己的锁信息
	models.DeleteKeyBySystemName(syncer.SYSTEM_NAME)

	go startConsensusState()
	go syncAllNodeBlock()
	go startNetwork()
	go startEventListener()
}

func startConsensusState() {
	for {
		startTime := time.Now().Unix()
		syncAllNodeConsensusState()
		endTime := time.Now().Unix()
		//每5分钟
		detTime := 300 - (endTime - startTime)
		if detTime > 0 {
			time.Sleep(time.Second * time.Duration(detTime))
		}
	}
}

func startNetwork() {
	for {
		startTime := time.Now().Unix()
		syncAllNodeNetwork()
		endTime := time.Now().Unix()
		//每60分钟
		detTime := 3600 - (endTime - startTime)
		if detTime > 0 {
			time.Sleep(time.Second * time.Duration(detTime))
		}
	}
}

func startEventListener() {
	nodes, err := service.AllNodes()
	if err != nil {
		fmt.Errorf("No node find")
		os.Exit(1)
	}
	for _, n := range nodes {
		if strings.Index(n.Name, "cosmos") < 0 {
			err = n.SubscribInflation()
			if err != nil {
				os.Exit(1)
			}
		}
	}
}
