// Copyright 2018 The QOS Authors

package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendAuthEmail(t *testing.T) {
	//err := sendAuthEmail("xuyz1988@163.com")
	err := send("xuyz1988@163.com")
	assert.Nil(t, err)
}
