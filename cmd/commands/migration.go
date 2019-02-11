// Copyright 2018 The QOS Authors

package commands

import (
	"errors"

	"github.com/QOSGroup/qmoon/models"
	"github.com/spf13/cobra"
)

// MigrationCmd 数据库初始化命令
var MigrationCmd = &cobra.Command{
	Use:   "migration up",
	Short: "migration",
	RunE:  migration,
}

func init() {
	registerFlagsDb(MigrationCmd)
}

func migration(cmd *cobra.Command, args []string) error {
	if len(args) != 1 || args[0] != "up" {
		return errors.New("需要参数up")
	}

	err := models.InitDb(config.DB)
	if err != nil {
		return err
	}

	return nil
}
