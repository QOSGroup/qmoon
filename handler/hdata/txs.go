// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service/tx"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const txsUrl = "/txs"

func init() {
	hdataHander[txsUrl] = TxsGinRegister
}

// TxsGinRegister 注册txs
func TxsGinRegister(r *gin.Engine) {
	r.GET(nodeProxy+txsUrl, middleware.ApiAuthGin(), txsGin())
}

func txsGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		nt, err := getNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		var minHeight, maxHeight int64
		maxHeightStr := c.Query("maxHeight")
		maxHeight, _ = strconv.ParseInt(maxHeightStr, 10, 64)

		minHeightStr := c.Query("minHeight")
		minHeight, _ = strconv.ParseInt(minHeightStr, 10, 64)

		ts, err := tx.List(nt.ChanID, minHeight, maxHeight)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		res := ts

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
