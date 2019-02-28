package main

import (
	bcli "github.com/QOSGroup/qbase/client"
	"github.com/QOSGroup/qmoon/cmd/commands"
	"github.com/QOSGroup/qmoon/lib"
)

func main() {
	rootCmd := commands.RootCmd
	rootCmd.AddCommand(
		commands.InitFilesCmd,
		commands.MigrationCmd,
		commands.NodeCmd,
		commands.TxCmd,
		commands.ServerCmd,
		commands.DoctorCmd,
		commands.VersionCmd)

	rootCmd.AddCommand(
		bcli.KeysCommand(lib.Cdc))

	//cmd := cli.PrepareBaseCmd(rootCmd, "", os.ExpandEnv(filepath.Join("$HOME", cfg.DefaultQMoonDir)))
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
