// Copyright 2018 The QOS Authors

package commands

import (
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/QOSGroup/qmoon/handler"
	"github.com/QOSGroup/qmoon/handler/hadmin"
	"github.com/QOSGroup/qmoon/handler/hdata"
	"github.com/QOSGroup/qmoon/plugins"
	"github.com/QOSGroup/qmoon/static"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/worker"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"	
)
var slogger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))

// ServerCmd qmoon http server
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "restful api server",
	RunE:  server,
}


func init() {
	registerFlagsHttpServer(ServerCmd)
	registerFlagsDb(ServerCmd)
	ServerCmd.Flags().Int64(types.FlagMaxGas, 20000, "gas limit to set per tx")
	ServerCmd.Flags().String(types.FlagInfluxdbServer, "http://localhost:8086", "influxdb server")
	ServerCmd.Flags().String(types.FlagInfluxdbUser, "", "influxdb user")
	ServerCmd.Flags().String(types.FlagInfluxdbPassword, "", "influxdb password")
}

func initRouter(r *gin.Engine) {
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

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static/swagger-ui/")
	})
	r.NoRoute(
		gin.WrapH(http.FileServer(
			&assetfs.AssetFS{Asset: static.Asset, AssetDir: static.AssetDir, AssetInfo: static.AssetInfo, Prefix: ""})))

	handler.VersionGinRegister(r)
	hadmin.SendCodeGinSendCode(r)
	hadmin.RegisterGinRegister(r)
	hadmin.LoginGinRegister(r)
	hadmin.LoginCheckGinRegister(r)
	hadmin.LogoutGinRegister(r)
	hadmin.AccountGinRegister(r)
	hadmin.AppGinRegister(r)
	hadmin.UpdatePasswordGinRegister(r)
	hadmin.NodesGinRegister(r)

	hdata.GinRegister(r)

	plugins.RegisterGin(r)
}

func server(cmd *cobra.Command, args []string) error {
	wg := sync.WaitGroup{}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	worker.Start()

	r := gin.Default()
	initRouter(r)

	go func() {
		if err := r.Run(config.HttpServer.ListenAddress); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
	common.TrapSignal(slogger,func() {})

	return nil
}
