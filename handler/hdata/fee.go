// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const feeUrl = NodeProxy + "/fee"

func init() {
	hdataHander[feeUrl] = FeeGinRegister
}

// FeeGinRegister 注册tx
func FeeGinRegister(r *gin.Engine) {
	r.GET(feeUrl, middleware.ApiAuthGin(), feeGin())
}

func feeGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		tx := c.Query("tx")
		result, err := node.Fee(tx)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", result))
	}
}
