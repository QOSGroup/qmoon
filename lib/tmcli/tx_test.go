// Copyright 2018 The QOS Authors

package tmcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTx(t *testing.T) {
	opt, err := NewOption(SetOptionHost(TmDefaultServer))
	assert.Nil(t, err)

	_, err = NewClient(opt).Tx.Retrieve(nil, []byte("a"), true)
	assert.NotNil(t, err)
}
