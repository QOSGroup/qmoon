// Copyright 2018 The QOS Authors

package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCoin(t *testing.T) {
	c, err := ParseCoin("100qos")
	assert.Nil(t, err)
	assert.Equal(t, c, Coin{Amount: 100, Denom: "qos"})
}

func TestParseCoins(t *testing.T) {
	cs, err := ParseCoins("100qos,20aoe")
	assert.Nil(t, err)
	assert.Equal(t, cs[0], Coin{Amount: 100, Denom: "qos"})
	assert.Equal(t, cs[1], Coin{Amount: 20, Denom: "aoe"})
}
