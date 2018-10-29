// Copyright 2018 The QOS Authors

package hadmin

import (
	"testing"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/lib/qstarscli"
	"github.com/QOSGroup/qmoon/lib/tmcli"
)

func TestMain(m *testing.M) {
	dbTest := db.NewTestDb(m)
	defer dbTest.Close()

	tq := qstarscli.NewTestQstarsServer()
	defer tq.Close()

	tts := tmcli.NewTestTmServer()
	defer tts.Close()

	m.Run()
}

func getSession() (string, error) {
	return "", nil
}
