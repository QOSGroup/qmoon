// Copyright 2018 The QOS Authors

package db

import (
	"database/sql"
	"errors"

	"github.com/QOSGroup/qmoon/config"
	"github.com/QOSGroup/qmoon/db/model"
	"github.com/tendermint/tendermint/libs/log"

	_ "github.com/lib/pq"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func checkDriverSupport(d string) bool {
	if d == "mysql" {
		return true
	}

	if d == "postgres" {
		return true
	}

	return false
}

func InitDb(cfg *config.DBConfig, logger log.Logger) error {
	if !checkDriverSupport(cfg.DriverName) {
		return errors.New("unsupported db driver")
	}

	d, err := sql.Open(cfg.DriverName, cfg.DataSource())
	if err != nil {
		return err
	}

	if err := d.Ping(); err != nil {
		return err
	}

	Db = d

	model.XOLog = logger.Debug

	return nil
}
