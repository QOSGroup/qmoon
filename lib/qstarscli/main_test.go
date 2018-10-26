// Copyright 2018 The QOS Authors

package qstarscli

import (
	"net/http/httptest"
	"testing"

	"github.com/QOSGroup/qmoon/lib/qstarscli/qstarsmock"
)

var qstarsMockServer string

func TestMain(m *testing.M) {
	tm := qstarsmock.NewQstarsMock()
	s := httptest.NewServer(tm)
	defer s.Close()
	qstarsMockServer = s.URL

	m.Run()
}
