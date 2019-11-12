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

const accountTxsUrl = "/accounts/:address/txs"

func init() {
	hdataHander[accountTxsUrl] = AccountTxsGinRegister
}

// AccountTxsGinRegister 注册accountTxs
func AccountTxsGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+accountTxsUrl, middleware.ApiAuthGin(), accountTxsGin())
}

func accountTxsGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		address := c.Param("address")
		var minHeight, maxHeight int64
		maxHeightStr := c.Query("maxHeight")
		maxHeight, _ = strconv.ParseInt(maxHeightStr, 10, 64)

		minHeightStr := c.Query("minHeight")
		minHeight, _ = strconv.ParseInt(minHeightStr, 10, 64)

		offset, limit := 0, 20
		tx := c.Query("tx")

		if d, err := strconv.ParseInt(c.Query("offset"), 10, 64); err == nil {
			offset = int(d)
		}

		if d, err := strconv.ParseInt(c.Query("limit"), 10, 64); err == nil {
			limit = int(d)
		}

		ts, err := node.TxsByAddress(address, tx, minHeight, maxHeight, offset, limit)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		res := ts

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
