package migrations

import (
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
)

const _MIN_DB_VER = 10

type Migration interface {
	Description() string
	Migrate(*xorm.Engine) error
}

type migration struct {
	description string
	migrate     func(*xorm.Engine) error
}

func NewMigration(desc string, fn func(*xorm.Engine) error) Migration {
	return &migration{desc, fn}
}

func (m *migration) Description() string {
	return m.description
}

func (m *migration) Migrate(x *xorm.Engine) error {
	return m.migrate(x)
}

// The version table. Should have only one row with id==1
type Version struct {
	ID      int64
	Version int64
}

var migrations = []Migration{}

// Migrate database to current version
func Migrate(x *xorm.Engine) error {
	if err := x.Sync2(new(Version)); err != nil {
		return fmt.Errorf("sync: %v", err)
	}

	currentVersion := &Version{ID: 1}
	has, err := x.Get(currentVersion)
	if err != nil {
		return fmt.Errorf("get: %v", err)
	} else if !has {
		// If the version record does not exist we think
		// it is a fresh installation and we can skip all migrations.
		currentVersion.ID = 0
		currentVersion.Version = int64(_MIN_DB_VER + len(migrations))

		if _, err = x.InsertOne(currentVersion); err != nil {
			return fmt.Errorf("insert: %v", err)
		}
	}

	v := currentVersion.Version
	if _MIN_DB_VER > v {
		log.Print(`
Qmoon can not auto-migration from your previously installed version.
`)
		return nil
	}

	if int(v-_MIN_DB_VER) > len(migrations) {
		currentVersion.Version = int64(len(migrations) + _MIN_DB_VER)
		_, err = x.ID(1).Update(currentVersion)
		return err
	}
	for i, m := range migrations[v-_MIN_DB_VER:] {
		if err = m.Migrate(x); err != nil {
			return fmt.Errorf("do migrate: %v", err)
		}
		currentVersion.Version = v + int64(i) + 1
		if _, err = x.ID(1).Update(currentVersion); err != nil {
			return err
		}
	}
	return nil
}
