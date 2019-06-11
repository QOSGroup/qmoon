// Copyright 2018 The QOS Authors

package worker

import (
	"github.com/robfig/cron"
)

func Start() {

	go SyncAllNodeConsensusState()
	go SyncAllNodeBlock()
	go SyncAllNodeNetwork()

	c := cron.New()

	_ = c.AddFunc("@every 5m", SyncAllNodeConsensusState) // 每5分

	_ = c.AddFunc("@every 60m", SyncAllNodeNetwork) // 每60分

	c.Start()
}
