// Copyright 2018 The QOS Authors

package tmcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlock(t *testing.T) {
	opt, err := NewOption(SetOptionHost(tmMockServer))
	assert.Nil(t, err)

	res, err := NewClient(opt).Block.Retrieve(nil, &BlockOption{Height: 1})
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
