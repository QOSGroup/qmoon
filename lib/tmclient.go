// Copyright 2018 The QOS Authors

package lib

import (
	"sync"

	"github.com/tendermint/tendermint/rpc/client"
)

type TmClient struct {
	tmcs map[string]*client.HTTP
	lock *sync.Mutex
}

var tmcs *TmClient

func init() {
	tmcs = &TmClient{
		tmcs: make(map[string]*client.HTTP),
		lock: new(sync.Mutex),
	}
}

func (tc *TmClient) Get(remote string) *client.HTTP {
	tc.lock.Lock()
	defer tc.lock.Unlock()

	c, ok := tc.tmcs[remote]
	if ok {
		return c
	}

	tmc := client.NewHTTP(remote, "/websocket")
	tc.tmcs[remote] = tmc

	return tmc
}

func TendermintClient(remote string) *client.HTTP {
	return tmcs.Get(remote)
}
