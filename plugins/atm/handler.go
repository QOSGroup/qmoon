// Copyright 2018 The QOS Authors

package atm

import (
	"net/http"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const transferUrl = types.UrlNodeProxy + "/accounts/:address/withdraw"

// AccountWithdrawGinRegister 注册accountWithdraw
func AccountWithdrawGinRegister(r *gin.Engine) {
	r.GET(transferUrl, TransferGin())
}

func TransferGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		chainID := c.Param("nodeName")
		address := c.Param("address")

		ip := c.ClientIP()
		logrus.WithField("model", "atm").WithField("ip", ip).Debug()

		if err := ipCheck(ip, chainID); err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		res, err := Withdraw(address, chainID)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		ipWithdraw(ip, chainID)

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
