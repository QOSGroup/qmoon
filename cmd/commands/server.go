// Copyright 2018 The QOS Authors

package commands

import (
	"net/http"

	"github.com/QOSGroup/qmoon/db"
	"github.com/QOSGroup/qmoon/handler"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// ServerCmd qmoon http server
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "server explorer",
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

	r := mux.NewRouter()
	r.Handle("/version", handler.ServerInfoHandler()).Methods(http.MethodGet)

	http.ListenAndServe(config.HttpServer.ListenAddress, r)

	return nil
}
