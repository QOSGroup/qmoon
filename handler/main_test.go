// Copyright 2018 The QOS Authors

// Package pkg comments for pkg handler
// handler ...
package handler

import (
	"testing"

	"github.com/QOSGroup/qmoon/db"
)

func TestMain(m *testing.M) {
	db.MakeTestMain()(m)
}
