// Copyright 2018 The QOS Authors

package qstarscli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatekV(t *testing.T) {
	opt, err := NewOption(SetOptionHost(tmDefaultServer))
	assert.Nil(t, err)

	_, err = NewClient(opt).KV.Create(nil, "key", "value", "", "")
	assert.NotNil(t, err)
}

func TestRetrievekV(t *testing.T) {
	opt, err := NewOption(SetOptionHost(tmDefaultServer))
	assert.Nil(t, err)

	_, err = NewClient(opt).KV.Retrieve(nil, "key")
	assert.NotNil(t, err)
}
