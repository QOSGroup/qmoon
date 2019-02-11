// Copyright 2018 The QOS Authors

package worker

import (
	"github.com/robfig/cron"
)

func Start() {
	c := cron.New()

	//c.AddFunc("@every 5m", SyncAllConsensusState) // 每5分
	c.AddFunc("@every 3s", SyncAllConsensusState) // 每5分

	//c.AddFunc("@every 5m", SyncPeersLoop)         // 每5分
	c.AddFunc("@every 3s", SyncPeersLoop) // 每5分

	//c.AddFunc("@every 5m", SyncAllNodeValidator) // 每5分
	c.AddFunc("@every 3s", SyncAllNodeValidator) // 每3秒
	c.AddFunc("@every 5s", SyncAllNodeBlock)     // 每5秒

	c.Start()
}
