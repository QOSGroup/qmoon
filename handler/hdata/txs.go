// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const txsUrl = "/txs"

func init() {
	hdataHander[txsUrl] = TxsGinRegister
}

// TxsGinRegister 注册txs
func TxsGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+txsUrl, middleware.ApiAuthGin(), txsGin())
}

func txsGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		var minHeight, maxHeight int64
		maxHeightStr := c.Query("maxHeight")
		maxHeight, _ = strconv.ParseInt(maxHeightStr, 10, 64)

		minHeightStr := c.Query("minHeight")
		minHeight, _ = strconv.ParseInt(minHeightStr, 10, 64)

		offset, _ := strconv.ParseInt(c.Query("offset"), 10, 64)

		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		if limit == 0 {
			limit = 20
		}

		ts, err := node.Txs(minHeight, maxHeight, offset, limit)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		res := ts

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
