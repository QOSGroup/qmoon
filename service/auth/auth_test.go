// Copyright 2018 The QOS Authors

package auth

import (
	"testing"
	"time"

	"github.com/QOSGroup/qmoon/config"
	"github.com/QOSGroup/qmoon/lib/qstarscli"
	"github.com/QOSGroup/qmoon/lib/tmcli"
	"github.com/QOSGroup/qmoon/models"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	dbTest, err := models.NewTestEngine(config.TestDBConfig())
	if err != nil {
		panic(err)
	}
	defer dbTest.Close()

	tq := qstarscli.NewTestQstarsServer()
	defer tq.Close()

	tts := tmcli.NewTestTmServer()
	defer tts.Close()

	m.Run()
}

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

	assert.Equal(t, sign1, sign2)
}
