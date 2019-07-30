// Copyright 2018 The QOS Authors

package syncer

import (
	"context"
	"strings"
	"testing"

	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/types"
	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
)

func TestQOSVersion(t *testing.T) {
	v1, err := version.NewSemver("v0.0.1")
	assert.Nil(t, err)
	v2, _ := version.NewSemver("0.0.2")
	v3, _ := version.NewSemver("0.0.3")
	v4, _ := version.NewSemver("v0.0.4")
	v5, _ := version.NewSemver("0.0.5")

	assert.True(t, v1.LessThan(qos0_0_3))
	assert.True(t, v1.LessThan(qos0_0_4))

	assert.True(t, v2.LessThan(qos0_0_3))
	assert.True(t, v2.LessThan(qos0_0_4))

	assert.True(t, v3.Equal(qos0_0_3))
	assert.True(t, v3.LessThan(qos0_0_4))

	assert.True(t, v4.GreaterThan(qos0_0_3))
	assert.True(t, v4.Equal(qos0_0_4))

	assert.True(t, v5.GreaterThan(qos0_0_3))
	assert.True(t, v5.GreaterThan(qos0_0_4))
}

func TestJson(t *testing.T) {
	arr := `[{"type":"send","data":{"from_address":"cosmos1tsggqpeu8yw82thhy6jp4l889zmruw6eum45tf","to_address":"cosmos1yeygh0y8rfyufdczhzytcl
3pehsnxv9d3wsnlg","amount":[{"denom":"uatom","amount":"175000"}]}}]`

	obj := `{"senders":[{"addr":"address1tqcjlsjgjdcxn3045evgfpr3u8u5h6ut7tu4r0","qos":"1000000","qscs":null}],"receivers":[{"addr":"address1qkf3
jqkxf2rzpkyud6qyn6yk83jhysy2xxlq2d","qos":"1000000","qscs":null}]}`

	t.Logf("arr:%v", strings.HasPrefix(arr, "["))
	t.Logf("obj:%v", strings.HasPrefix(obj, "{"))
}

func TestRpcPeers(t *testing.T) {

	qos := NewSyncer(&service.Node{
		BaseURL:  "http://47.105.52.237:36657",
		NodeType: types.NodeTypeQOS.String(),
	})

	err := qos.RpcPeers(context.Background())
	assert.Nil(t, err)
}
