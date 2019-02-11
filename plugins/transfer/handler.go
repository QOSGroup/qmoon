// Copyright 2018 The QOS Authors

package transfer

import (
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/handler/hdata"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const transferUrl = types.UrlNodeProxy + "/accounts/:address/transfer"

// AccountTxsGinRegister 注册accountTxs
func AccountTxsGinRegister(r *gin.Engine) {
	r.GET(transferUrl, TransferGin())
}

func TransferGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := hdata.GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}
		address := c.Param("address")
		offset, limit := 0, 20
		coin := c.Query("coin")

		if d, err := strconv.ParseInt(c.Query("offset"), 10, 64); err == nil {
			offset = int(d)
		}

		if d, err := strconv.ParseInt(c.Query("limit"), 10, 64); err == nil {
			limit = int(d)
		}

		txs, err := ListByAddress(node.ChanID, address, offset, limit, &SearchOpt{Coin: coin})
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", txs))
	}
}
