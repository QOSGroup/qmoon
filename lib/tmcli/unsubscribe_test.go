// Copyright 2018 The QOS Authors

package tmcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnsubscribe(t *testing.T) {
	opt, err := NewOption()
	assert.Nil(t, err)

	_, err = NewClient(opt).Unsubscribe.Retrieve(nil, "a")
	assert.NotNil(t, err)
}
