// Copyright 2018 The QOS Authors

package commands

import (
	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/handler"
	"github.com/QOSGroup/qmoon/handler/hadmin"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// ServerCmd qmoon http server
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "restful api server",
	RunE:  server,
}

func init() {
	registerFlagsHttpServer(ServerCmd)
	registerFlagsDb(ServerCmd)
}

func server(cmd *cobra.Command, args []string) error {
	err := db.InitDb(config.DB, logger)
	if err != nil {
		return err
	}

	r := gin.Default()

	handler.VersionGinRegister(r)
	hadmin.AccountRegisterGinRegister(r)
	hadmin.LoginGinRegister(r)
	hadmin.LogoutGinRegister(r)
	hadmin.AccountGinRegister(r)
	hadmin.AppGinRegister(r)

	if err := r.Run(config.HttpServer.ListenAddress); err != nil {
		return err
	}

	return nil
}
