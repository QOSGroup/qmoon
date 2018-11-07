// Copyright 2018 The QOS Authors

package commands

import (
	"fmt"
	"net/http"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/migrations"
	"github.com/QOSGroup/qmoon/handler"
	"github.com/QOSGroup/qmoon/handler/hadmin"
	"github.com/QOSGroup/qmoon/handler/hdata"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// ServerCmd qmoon http server
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "restful api server",
	RunE:  server,
}

var (
	explorer string
)

func init() {
	registerFlagsHttpServer(ServerCmd)
	registerFlagsDb(ServerCmd)

	ServerCmd.PersistentFlags().StringVar(&explorer, "with-explorer", "", "the dir of explorer")
}

func server(cmd *cobra.Command, args []string) error {
	err := db.InitDb(config.DB, logger)
	if err != nil {
		return err
	}

	needUp, err := migrations.NeedMigration(config.DB.DriverName, db.Db)
	if err != nil {
		panic(err)
	}

	if needUp {
		fmt.Println("数据库需要升级")
		fmt.Println("可以执行 qmoon migration up")
		return nil
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, HEAD, OPTIONS, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, QToken") //有使用自定义头 需要这个,Action, Module是例子

		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusOK)
		}
	})

	handler.VersionGinRegister(r)
	hadmin.RegisterGinRegister(r)
	hadmin.LoginGinRegister(r)
	hadmin.LoginCheckGinRegister(r)
	hadmin.LogoutGinRegister(r)
	hadmin.AccountGinRegister(r)
	hadmin.AppGinRegister(r)
	hadmin.UpdatePasswordGinRegister(r)
	hadmin.NodeTypesGinRegister(r)
	hdata.ProxyGinRegister(r)

	if explorer != "" {
		s := gin.Default()
		s.Static("/", explorer)

		go s.Run("0.0.0.0:80")
	}

	if err := r.Run(config.HttpServer.ListenAddress); err != nil {
		return err
	}

	return nil
}
