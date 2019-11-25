// Copyright 2018 The QOS Authors

package hdata

import (
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/lib/qos"
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const accountQueryUrl = "/accounts/:address"
const accountTxsUrl = "/accounts/:address/txs"
const accountDelegationsUrl = "/accounts/:address/delegations"

func init() {
	hdataHander[accountQueryUrl] = AccountQueryGinRegister
	hdataHander[accountTxsUrl] = AccountTxsGinRegister
	hdataHander[accountDelegationsUrl] = AccountDelegationsGinRegister
}

// AccountQueryGinRegister 注册accountQuery
func AccountQueryGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+accountQueryUrl, middleware.ApiAuthGin(), accountQueryGin())
}
// AccountTxsGinRegister 注册accountTxs
func AccountTxsGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+accountTxsUrl, middleware.ApiAuthGin(), accountTxsGin())
}

func AccountDelegationsGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+accountDelegationsUrl, middleware.ApiAuthGin(), accountDelegationsGin())
}

func accountQueryGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		address := c.Param("address")
		result, err := service.QueryAccount(node, address)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		c.JSON(http.StatusOK, types.RPCResponse{JSONRPC: "2.0", ID: "", Result: result})
	}
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

		if d, err := strconv.ParseInt(c.Query("offset"), 10, 64); err == nil {
			offset = int(d)
		}

		if d, err := strconv.ParseInt(c.Query("limit"), 10, 64); err == nil {
			limit = int(d)
		}

		ts, err := node.TxsByAddress(address, minHeight, maxHeight, offset, limit)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		res := ts

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}

func accountDelegationsGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		dels, err := qos.NewQosCli("").QueryDelegationsWithDelegator(node.BaseURL, c.Param("address"))
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", dels))
	}
}