// Copyright 2018 The QOS Authors

package commands

import (
	"errors"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/migrations"
	"github.com/QOSGroup/qmoon/plugins"
	"github.com/spf13/cobra"
)

// MigrationCmd 数据库初始化命令
var MigrationCmd = &cobra.Command{
	Use:   "migration up/down",
	Short: "migration",
	RunE:  migration,
}

func init() {
	registerFlagsDb(MigrationCmd)
}

func migration(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("需要参数up或down")
	}
	t := args[0]

	err := db.InitDb(config.DB, logger)
	if err != nil {
		return err
	}

	if t == "up" {
		migrations.Up(config.DB.DriverName, db.Db)
		plugins.DbUp(config.DB.DriverName, db.Db)

	} else if t == "down" {
		migrations.Down(config.DB.DriverName, db.Db)
		plugins.DbDown(config.DB.DriverName, db.Db)

	} else {
		return errors.New("需要参数up或down")
	}

	return nil
}
