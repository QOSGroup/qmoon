// Copyright 2018 The QOS Authors

package lib

import (
	"github.com/tendermint/tendermint/rpc/client"
)

var tmcs map[string]*client.HTTP

func init() {
	tmcs = make(map[string]*client.HTTP)
}

func TendermintClient(remote string) *client.HTTP {

	c, ok := tmcs[remote]
	if ok {
		return c
	}
	tmc := client.NewHTTP(remote, "/websocket")
	tmcs[remote] = tmc

	return tmc
}
