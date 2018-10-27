// Copyright 2018 The QOS Authors

package qstarscli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeVersion(t *testing.T) {
	opt, err := NewOption(SetOptionHost(tmDefaultServer))
	assert.Nil(t, err)

	_, err = NewClient(opt).NodeVersion.Retrieve(nil)
	assert.NotNil(t, err)
}
