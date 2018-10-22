// Copyright 2018 The QOS Authors

package handler

import (
	"net/http"

	"github.com/QOSGroup/qmoon/version"
	"github.com/gin-gonic/gin"
)

func VersionGinRegister(r *gin.Engine) {
	r.GET("/version",serverInfoGinHandler())
}

type versionRes struct {
	Version string `json:"version"`
}

func serverInfo() *versionRes {
	return &versionRes{
		Version: version.Version,
	}
}

func serverInfoGinHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		si := serverInfo()

		c.JSON(http.StatusOK, WrapperRPCResponse("", si, nil))
	}
}
