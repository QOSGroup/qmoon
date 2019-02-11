// Copyright 2018 The QOS Authors

package syncer

import (
	"context"
	"testing"
	"time"

	"github.com/QOSGroup/qmoon/service"
)

func TestSyncBlockV2(t *testing.T) {
	s := NewSyncer(&service.Node{BaseURL: "aaa"})

	context.TODO()
	//ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	go s.Block(ctx)

	time.Sleep(time.Second * 10)
	cancel()
}
