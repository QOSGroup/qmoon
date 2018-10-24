// Copyright 2018 The QOS Authors

package tmcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockchain(t *testing.T) {
	opt, err := NewOption()
	assert.Nil(t, err)

	res, err := NewClient(opt).Blockchain.List(nil, &BlockchainOption{MinHeight: 1, MaxHeight: 1})
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, res.BlockMetas)
}
