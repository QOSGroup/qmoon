// Copyright 2019 The QOS Authors

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDigit(t *testing.T) {
	assert.Equal(t, true, IsDigit("86400"))
	assert.Equal(t, true, IsDigit("1"))
	assert.Equal(t, true, IsDigit("123"))
	assert.Equal(t, false, IsDigit("a1"))
	assert.Equal(t, false, IsDigit("86400s"))
	assert.Equal(t, false, IsDigit("1a2"))
	assert.Equal(t, false, IsDigit("12a"))
}
