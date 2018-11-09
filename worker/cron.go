// Copyright 2018 The QOS Authors

package worker

import (
	"github.com/robfig/cron"
)

//func Start(quit chan os.Signal) {
func Start() {
	c := cron.New()

	c.AddFunc("@every 5s", SyncBlockLoop) // 每5秒
	c.AddFunc("@every 1m", SyncPeersLoop) // 每分钟
	//c.AddFunc("1 1 3 * * *", func() {})  // 每天早上03:01:01

	c.Start()
}
