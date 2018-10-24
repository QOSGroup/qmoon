// Copyright 2018 The QOS Authors

package tmcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	opt, err := NewOption()
	assert.Nil(t, err)

	res, err := NewClient(opt).Status.Retrieve(nil)
	assert.Nil(t, err)
	assert.NotNil(t, res.NodeInfo)
}
