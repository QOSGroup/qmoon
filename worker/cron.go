// Copyright 2018 The QOS Authors

package worker

import (
	"fmt"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/service/syncer"
	"github.com/tendermint/tendermint/rpc/client"
	"os"
	"strconv"
	"strings"
	"time"
)

func Start() {
	//先删除自己的锁信息
	models.DeleteKeyBySystemName(syncer.SYSTEM_NAME)

	go startEventListener()
	go startConsensusState()
	go syncAllNodeBlock()
	go startNetwork()
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
				url := "tcp" + n.BaseURL[4:len(n.BaseURL)]

				client := client.NewHTTP(url, "/websocket")
				fmt.Println("Starting listening to ", url, "/websocket")
				err := client.Start()
				if err != nil {
					fmt.Println("[Event] Can't start websocket client [%s] - '%s'", url, err)
				}
				//defer client.Stop()
				_, events, err := n.SubscribInflation(client)
				if err != nil {
					fmt.Errorf("Exiting for error:", err)
					client.Stop()
					os.Exit(1)
				}
				go func() {
					for eventData := range events {
						fmt.Println("[Event] Received event from [%s] - '%s'", url, eventData)
						fmt.Println("[Event] event data: ", eventData.Data)
						fmt.Println("[Event] events: ", eventData.Events)
						inf := models.Inflation{}
						for key := range eventData.Events {
							switch key {
							case "mint.height":
								for _, runeValue := range eventData.Events[key][0] {
									inf.Height = int64(runeValue)
								}
								break
							case "mint.tokens":
								if n, err := strconv.ParseInt(eventData.Events[key][0], 10, 64); err == nil {
									inf.Tokens = n
								} else {
									fmt.Println("[Event] Inflation tokens error : %s", eventData.Events[key][0])
								}
								break
							}
						}
						if (inf.Height>0 && inf.Tokens>0) {
							inf.InsertOrUpdate(n.ChainID)
						}
					}
				}()
			}
		}

}
