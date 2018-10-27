// Copyright 2018 The QOS Authors

package tmcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	opt, err := NewOption(SetOptionHost(tmDefaultServer))
	assert.Nil(t, err)

	res, err := NewClient(opt).Health.Retrieve(nil)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
