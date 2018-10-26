// Copyright 2018 The QOS Authors

package tmcli

import (
	"net/http/httptest"
	"testing"

	"github.com/QOSGroup/qmoon/lib/tmcli/tmmock"
)

var tmMockServer string

func TestMain(m *testing.M) {
	tm := tmmock.NewTendermintMock()
	s := httptest.NewServer(tm)
	defer s.Close()
	tmMockServer = s.URL

	m.Run()
}
