// Copyright 2018 The QOS Authors

package atm

import (
	"net/http"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const transferUrl = "/node/:chainId/accounts/:address/withdraw"

// AccountWithdrawGinRegister 注册accountWithdraw
func AccountWithdrawGinRegister(r *gin.Engine) {
	r.GET(transferUrl, TransferGin())
}

func TransferGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		chainID := c.Param("chainId")
		address := c.Param("address")

		res, err := Withdraw(address, chainID)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
