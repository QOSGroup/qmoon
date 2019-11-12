// Copyright 2018 The QOS Authors

package context

import (
	"testing"

	"github.com/QOSGroup/qmoon/lib/log"
	"github.com/magiconair/properties/assert"
	"github.com/sirupsen/logrus"
)

func TestNewContext(t *testing.T) {
	c := NewContext(nil, log.New())
	c = c.WithChainID("chain_id")
	assert.Equal(t, "chain_id", c.ChainID())
	c = c.WithLogger(logrus.New())
	assert.Equal(t, "chain_id", c.ChainID())
}
