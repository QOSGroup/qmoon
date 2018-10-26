// Copyright 2018 The QOS Authors

package tmcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBroadcastTxSync(t *testing.T) {
	opt, err := NewOption(SetOptionHost(tmMockServer))
	assert.Nil(t, err)

	res, err := NewClient(opt).BroadcastTxSync.Retrieve(nil, &BroadcastTxSyncOption{Tx: ""})
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
