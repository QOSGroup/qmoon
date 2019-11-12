package service

import (
	"fmt"
	"context"
	"time"

	//"github.com/snowdiceX/events"
	"github.com/tendermint/tendermint/rpc/client"
	//"github.com/snowdiceX/events"
)

func (n Node) SubscribInflation() error {
	url := "tcp" + n.BaseURL[4:len(n.BaseURL)]
	//_, events, err := events.SubscribeRemote(url, n.ChainID, "tm.event = 'mint' ")
	//if err != nil {
	//	fmt.Println("Remote [%s] : '%s'", url, err)
	//	return err
	//}
	//go func() {
	//	for eventData := range events {
	//		fmt.Println("Received event from [%s] - '%s'", url, eventData)
	//		fmt.Println("event data: ", eventData.Data)
	//		fmt.Println("events: ", eventData.Events)
	//		// inf := models.Inflation{Height:eventData.Data, eventData.}
	//	}
	//}()
	//return nil


	client := client.NewHTTP(url, "/websocket")
	err := client.Start()
	if err != nil {
	  // handle error
	}
	defer client.Stop()
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()
	query := "tm.event = 'Inflation'"
	events, err := client.Subscribe(ctx, "qmoon", query)
	if err != nil {
		fmt.Println("[Event] Remote [%s] : '%s'", url, err)
	}
	go func() {
		for eventData := range events {
			fmt.Println("[Event] Received event from [%s] - '%s'", url, eventData)
			fmt.Println("[Event] event data: ", eventData.Data)
			fmt.Println("[Event] events: ", eventData.Events)
			// inf := models.Inflation{Height:eventData.Data, eventData.}
		}
	}()
	return nil
}
