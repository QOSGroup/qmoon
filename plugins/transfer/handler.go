// Copyright 2018 The QOS Authors

package transfer

import (
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const transferUrl = "/node/:chainId/accounts/:address/transfer"

// AccountTxsGinRegister 注册accountTxs
func AccountTxsGinRegister(r *gin.Engine) {
	r.GET(transferUrl, TransferGin())
}

func TransferGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//chainID := c.Param("chainId")
		address := c.Param("address")
		offset, limit := int64(0), int64(20)

		if d, err := strconv.ParseInt(c.Param("offset"), 10, 64); err == nil {
			offset = d
		}

		if d, err := strconv.ParseInt(c.Param("limit"), 10, 64); err == nil {
			limit = d
		}

		txs, err := ListByAddress(address, offset, limit)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", txs))
	}
}
