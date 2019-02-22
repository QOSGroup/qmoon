// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const txHashUrl = NodeProxy + "/tx/:hash"

func init() {
	hdataHander[txHashUrl] = TxHashGinRegister
}

// TxHashGinRegister 注册tx
func TxHashGinRegister(r *gin.Engine) {
	r.GET(txHashUrl, middleware.ApiAuthGin(), txHashGin())
}

func txHashGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		hash := c.Param("hash")
		result, err := node.TxByHash(hash)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", result))
	}
}
