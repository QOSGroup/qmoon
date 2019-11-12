package service

import (
	"fmt"
	"github.com/snowdiceX/events"
)

func (n Node) SubscribInflation() error {
	url := "tcp" + n.BaseURL[4:len(n.BaseURL)-1]
	_, events, err := events.SubscribeRemote(url, n.ChainID, "tm.event = 'mint' ")
	if err != nil {
		fmt.Println("Remote [%s] : '%s'", url, err)
		return err
	}
	go func() {
		for eventData := range events {
			fmt.Println("Received event from [%s] - '%s'", url, eventData)
			fmt.Println("event data: ", eventData.Data)
			fmt.Println("events: ", eventData.Events)
			// inf := models.Inflation{Height:eventData.Data, eventData.}
		}
	}()
	return nil
}
