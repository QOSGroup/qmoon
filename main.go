package main

import (
	"github.com/QOSGroup/qmoon/cmd/commands"
)

// 静态文件
//go:generate cp -f docs/swagger.json static/
//go:generate go-bindata -o static/static.go -ignore=static/static.go -pkg=static static/ static/swagger-ui/

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

	//cmd := cli.PrepareBaseCmd(rootCmd, "", os.ExpandEnv(filepath.Join("$HOME", cfg.DefaultQMoonDir)))
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
