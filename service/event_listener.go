package service

import (
	"fmt"
	"context"
	"time"
	tctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/rpc/client"
)

func (n Node) SubscribInflation(client *client.HTTP) (
	context.CancelFunc, <-chan tctypes.ResultEvent, error ) {
	ctx, cancel := context.WithTimeout(context.Background(), 10000 * time.Second)
	query := "tm.event = 'NewBlock'"
	events, err := client.Subscribe(ctx, n.ChainID, query)
	if err != nil {
		fmt.Println("[Event] subscribe error : '%s'", err)
		return cancel, nil, err
	}
	return cancel, events, nil
}
