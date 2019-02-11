// Copyright 2018 The QOS Authors

package hdata

import (
	"testing"

	"github.com/QOSGroup/qmoon/config"
	"github.com/QOSGroup/qmoon/handler"
	"github.com/QOSGroup/qmoon/lib/qstarscli"
	"github.com/QOSGroup/qmoon/lib/tmcli"
	"github.com/QOSGroup/qmoon/models"
)

func TestMain(m *testing.M) {
	dbTest, err := models.NewTestEngine(config.TestDBConfig())
	if err != nil {
		panic(err)
	}
	defer dbTest.Close()

	err = handler.CreateTestUser()
	if err != nil {
		panic(err)
	}

	tq := qstarscli.NewTestQstarsServer()
	defer tq.Close()

	tts := tmcli.NewTestTmServer()
	defer tts.Close()

	m.Run()
}
