// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const accountQueryUrl = "/accounts/:address"

func init() {
	hdataHander[accountQueryUrl] = AccountQueryGinRegister
}

// AccountQueryGinRegister 注册accountQuery
func AccountQueryGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+accountQueryUrl, middleware.ApiAuthGin(), accountQueryGin())
}

func accountQueryGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		nt, err := getNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		address := c.Param("address")
		ctx := lib.NewQstarsClient(nt.BaseURL)
		result, err := ctx.QueryAccount(address)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", result))
	}
}
