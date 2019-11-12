// Copyright 2018 The QOS Authors

package tmcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnconfirmedTxs(t *testing.T) {
	opt, err := NewOption(SetOptionHost(TmDefaultServer))
	assert.Nil(t, err)

	res, err := NewClient(opt).UnconfirmedTxs.Retrieve(nil, &UnconfirmedTxsOption{Limit: 0})
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
