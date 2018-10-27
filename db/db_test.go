package db

import (
	"database/sql"
	"testing"

	"github.com/QOSGroup/qmoon/db/model"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	dbTest := NewTestDb(m)
	defer dbTest.Close()

	m.Run()
}

func TestVersion(t *testing.T) {
	qs, err := model.QmoonStatusByKey(Db, sql.NullString{String: "qmoon_version", Valid: true})
	assert.Nil(t, err)
	assert.NotEmpty(t, qs.Value.String)
}
