// Copyright 2019 The QOS Authors

package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCheckHttpHealthy(t *testing.T) {
	u := "http://httpbin.org/anything"
	ok, duration, err := CheckHttpHealthy(u, time.Second*60)
	assert.Nil(t, err)
	t.Logf("ok:%v", ok)
	t.Logf("duration:%v", duration.String())
	t.Logf("duration:%v", duration.Nanoseconds())
}
