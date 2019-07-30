// Copyright 2018 The QOS Authors

package worker

import (
	"testing"

	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/testdata"
	"github.com/QOSGroup/qmoon/types"
)

func TestNetwork(t *testing.T) {
	network(&service.Node{
		BaseURL:  "http://47.105.52.237:36657",
		NodeType: types.NodeTypeQOS.String(),
		ChanID:   testdata.ChainID,
	})
}
