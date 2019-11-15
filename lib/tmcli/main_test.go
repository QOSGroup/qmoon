// Copyright 2018 The QOS Authors

package tmcli

import (
	"testing"
)

func TestMain(m *testing.M) {
	tts := NewTestTmServer()
	defer tts.Close()

	m.Run()
}
