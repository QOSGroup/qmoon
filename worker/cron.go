// Copyright 2018 The QOS Authors

package worker

import (
	"github.com/robfig/cron"
)

func Start() {

	go SyncAllNodeConsensusState()
	go SyncAllNodePeer()
	go SyncAllNodeValidator()
	go SyncAllNodeBlock()

	c := cron.New()

	c.AddFunc("@every 5m", SyncAllNodeConsensusState) // 每5分
	//c.AddFunc("@every 3s", SyncAllNodeConsensusState) // 每3s

	c.AddFunc("@every 5m", SyncAllNodePeer) // 每5分
	//c.AddFunc("@every 3s", SyncAllNodePeer) // 每3s

	c.AddFunc("@every 5m", SyncAllNodeValidator) // 每5分
	//c.AddFunc("@every 3s", SyncAllNodeValidator) // 每3s

	c.AddFunc("@every 1m", SyncAllNodeBlock) // 每1分

	c.Start()
}
