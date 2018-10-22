// Copyright 2018 The QOS Authors

package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	secretKey := []byte("afdcd")
	now := time.Now().Unix()
	sd1 := SignData{
		SecretID:    "123",
		expiredTime: int(now) + 60*60,
		currentTime: int(now),
		rand:        "wsxf",
	}

	sign1, err := sd1.Sign(secretKey)
	assert.Nil(t, err)

	sd2, err := ParseSign(sign1)
	assert.Nil(t, err)
	sign2, err := sd2.Sign(secretKey)
	assert.Nil(t, err)

	t.Log(string(sign1))
	assert.Equal(t, sign1, sign2)
}
