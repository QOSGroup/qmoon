// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const chainNodesUrl = "/nodes"

func init() {
	hdataHander[chainNodesUrl] = ChainNodesGinRegister
}

// ChainNodesGinRegister 注册chainNodes
func ChainNodesGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+chainNodesUrl, middleware.ApiAuthGin(), chainNodesGin())
}

func chainNodesGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		resp, err := node.ValidNodes()
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", resp))
	}
}
