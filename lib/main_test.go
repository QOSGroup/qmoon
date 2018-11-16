// Copyright 2018 The QOS Authors

package lib

import (
	"testing"

	"github.com/QOSGroup/qmoon/lib/tmcli"
)

var tmTestServer string

func TestMain(m *testing.M) {
	tts := tmcli.NewTestTmServer()
	defer tts.Close()
	tmTestServer = tts.URL()

	m.Run()
}
