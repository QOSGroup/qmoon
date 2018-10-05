// Copyright 2018 The QOS Authors

package commands

import (
	cfg "github.com/QOSGroup/qmoon/config"
	"github.com/spf13/cobra"
)

// InitFilesCmd initialises a fresh Tendermint Core instance.
var InitFilesCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize qmood",
	RunE:  initFiles,
}

func initFiles(cmd *cobra.Command, args []string) error {
	return initConfig(config)
}

func initConfig(config *cfg.Config) error {
	cfg.WriteDefaultConfigFile(config.BaseConfig.RootDir)

	return nil
}
