// Copyright 2018 The QOS Authors

package tmcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscribe(t *testing.T) {
	opt, err := NewOption(SetOptionHost(TmDefaultServer))
	assert.Nil(t, err)

	_, err = NewClient(opt).Subscribe.Retrieve(nil, "a")
	assert.NotNil(t, err)
}
