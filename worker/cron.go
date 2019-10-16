// Copyright 2018 The QOS Authors

package worker

import (
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/service/syncer"
	"time"
)

func Start() {
	//先删除自己的锁信息
	models.DeleteKeyBySystemName(syncer.SYSTEM_NAME)

	go startConsensusState()
	go syncAllNodeBlock()
	go startNetwork()
}

func startConsensusState(){
	for{
		startTime:=time.Now().Unix()
		syncAllNodeConsensusState()
		endTime:=time.Now().Unix()
		//每5分钟
		detTime:=300-(endTime-startTime)
		if detTime>0{
			time.Sleep(time.Second*time.Duration(detTime))
		}
	}
}

func startNetwork(){
	for{
		startTime:=time.Now().Unix()
		syncAllNodeNetwork()
		endTime:=time.Now().Unix()
		//每60分钟
		detTime:=3600-(endTime-startTime)
		if detTime>0{
			time.Sleep(time.Second*time.Duration(detTime))
		}
	}
}