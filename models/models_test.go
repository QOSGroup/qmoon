package models

import (
	"testing"

	"github.com/QOSGroup/qmoon/config"
)

func TestMain(m *testing.M) {
	dbTest, err := NewTestEngine(config.TestDBConfig())
	if err != nil {
		panic(err)
	}
	defer dbTest.Close()

	m.Run()
}
