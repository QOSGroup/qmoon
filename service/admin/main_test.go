// Copyright 2018 The QOS Authors

package admin

import (
	"testing"

	"github.com/QOSGroup/qmoon/db"
)

func TestMain(m *testing.M) {
	db.MakeTestMain()(m)
}
