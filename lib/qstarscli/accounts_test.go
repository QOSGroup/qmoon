// Copyright 2018 The QOS Authors

package qstarscli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	opt, err := NewOption(SetOptionHost(tmDefaultServer))
	assert.Nil(t, err)

	_, err = NewClient(opt).Accounts.Create(nil)
	assert.Nil(t, err)
}

func TestRetrieveAccount(t *testing.T) {
	opt, err := NewOption(SetOptionHost(tmDefaultServer))
	assert.Nil(t, err)

	_, err = NewClient(opt).Accounts.Retrieve(nil, "address")
	assert.NotNil(t, err)
}
