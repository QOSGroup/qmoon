package main

import (
	"os"
	"path/filepath"

	"github.com/QOSGroup/qmoon/cmd/commands"
	cfg "github.com/QOSGroup/qmoon/config"
	"github.com/tendermint/tendermint/libs/cli"
)

func main() {
	rootCmd := commands.RootCmd
	rootCmd.AddCommand(
		commands.MigrationCmd,
		commands.ServerCmd,
		commands.VersionCmd)

	cmd := cli.PrepareBaseCmd(rootCmd, "", os.ExpandEnv(filepath.Join("$HOME", cfg.DefaultQMoonDir)))
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
