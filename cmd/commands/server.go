// Copyright 2018 The QOS Authors

package commands

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/db/migrations"
	"github.com/QOSGroup/qmoon/handler"
	"github.com/QOSGroup/qmoon/handler/hadmin"
	"github.com/QOSGroup/qmoon/handler/hdata"
	"github.com/QOSGroup/qmoon/plugins"
	"github.com/QOSGroup/qmoon/worker"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/common"
)

// ServerCmd qmoon http server
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "restful api server",
	RunE:  server,
}

var (
	explorer      string
	explorerLaddr string
)

func init() {
	registerFlagsHttpServer(ServerCmd)
	registerFlagsDb(ServerCmd)

	ServerCmd.PersistentFlags().StringVar(&explorer, "with-explorer", "", "the dir of explorer")
	ServerCmd.PersistentFlags().StringVarP(&explorerLaddr, "explorerLaddr", "", "0.0.0.0:9528", "address of explorer listening")
}

func checkDb(db *sql.DB) bool {
	needUp, err := migrations.NeedMigration(config.DB.DriverName, db)
	if err != nil {
		panic(err)
	}
	if needUp {
		fmt.Println("数据库需要升级")
		fmt.Println("可以执行 qmoon migration up")
		return true
	}

	return false
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

func runStaticServer(laddr, dir string) {
	s := gin.Default()
	s.Static("/", dir)

	s.Run(laddr)
}

func server(cmd *cobra.Command, args []string) error {
	wg := sync.WaitGroup{}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	err := db.InitDb(config.DB, logger)
	if err != nil {
		return err
	}

	if err := plugins.DbUp(config.DB.DriverName, db.Db); err != nil {
		return err
	}

	if ok := checkDb(db.Db); ok {
		return nil
	}

	worker.Start()
	if explorer != "" {
		go func() {
			wg.Add(1)
			defer wg.Done()
			runStaticServer(explorerLaddr, explorer)
		}()
	}

	r := gin.Default()
	initRouter(r)

	go func() {
		if err := r.Run(config.HttpServer.ListenAddress); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
	common.TrapSignal(func() {})

	return nil
}
