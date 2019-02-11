// Copyright 2018 The QOS Authors

package atm

import (
	"net/http"

	"github.com/QOSGroup/qmoon/handler/hdata"
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
		node, err := hdata.GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}
		address := c.Param("address")

		ip := c.ClientIP()
		logrus.WithField("model", "atm").WithField("ip", ip).Debug()

		if err := ipCheck(ip, node.ChanID); err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		res, err := Withdraw(address, node.ChanID)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		ipWithdraw(ip, node.ChanID)

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
