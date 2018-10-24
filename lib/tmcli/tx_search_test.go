// Copyright 2018 The QOS Authors

package tmcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTxSearch(t *testing.T) {
	opt, err := NewOption()
	assert.Nil(t, err)

	_, err = NewClient(opt).TxSearch.Retrieve(nil, "", nil)
	assert.NotNil(t, err)
}
