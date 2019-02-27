// Copyright 2018 The QOS Authors

package syncer

import (
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
)

func TestQOSVersion(t *testing.T) {
	v1, err := version.NewSemver("v0.0.1")
	assert.Nil(t, err)
	v2, _ := version.NewSemver("0.0.2")
	v3, _ := version.NewSemver("0.0.3")
	v4, _ := version.NewSemver("v0.0.4")
	v5, _ := version.NewSemver("0.0.5")

	assert.True(t, v1.LessThan(qos0_0_3))
	assert.True(t, v1.LessThan(qos0_0_4))

	assert.True(t, v2.LessThan(qos0_0_3))
	assert.True(t, v2.LessThan(qos0_0_4))

	assert.True(t, v3.Equal(qos0_0_3))
	assert.True(t, v3.LessThan(qos0_0_4))

	assert.True(t, v4.GreaterThan(qos0_0_3))
	assert.True(t, v4.Equal(qos0_0_4))

	assert.True(t, v5.GreaterThan(qos0_0_3))
	assert.True(t, v5.GreaterThan(qos0_0_4))
}
