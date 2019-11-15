// Copyright 2018 The QOS Authors

package qstarscli

import (
	"testing"
)

func TestMain(m *testing.M) {
	tq := NewTestQstarsServer()
	defer tq.Close()

	m.Run()
}
