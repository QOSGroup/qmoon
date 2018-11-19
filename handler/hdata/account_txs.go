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

const accountTxsUrl = "/accounts/:address/txs"

func init() {
	hdataHander[accountTxsUrl] = AccountTxsGinRegister
}

// AccountTxsGinRegister 注册accountTxs
func AccountTxsGinRegister(r *gin.Engine) {
	r.GET(nodeProxy+accountTxsUrl, middleware.ApiAuthGin(), accountTxsGin())
}

func accountTxsGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		nt, err := getNodeFromUrl(c)
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

		ts, err := tx.ListByAddress(nt.ChanID, address, minHeight, maxHeight)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		res := ts

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
