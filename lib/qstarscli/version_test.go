// Copyright 2018 The QOS Authors

package qstarscli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	opt, err := NewOption(SetOptionHost(tmDefaultServer))
	assert.Nil(t, err)

	_, err = NewClient(opt).Version.Retrieve(nil)
	assert.Nil(t, err)
}
