// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const statusQueryUrl = "/status"

func init() {
	hdataHander[statusQueryUrl] = StatusQueryGinRegister
}

// StatusQueryGinRegister 注册statusQuery
func StatusQueryGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+statusQueryUrl, middleware.ApiAuthGin(), statusQueryGin())
}

func statusQueryGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		result, err := node.ChainStatus(false)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", result))
	}
}
