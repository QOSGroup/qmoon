// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"
	"strconv"
	"time"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service/metric"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/gin-gonic/gin"
)

const validatorUptimeUrl = NodeProxy + "/validators/:address/uptime"

func init() {
	hdataHander[validatorUptimeUrl] = ValidatorUptimeGinRegister
}

// ValidatorVotingPowerGinRegister 注册
func ValidatorUptimeGinRegister(r *gin.Engine) {
	r.GET(validatorUptimeUrl, middleware.ApiAuthGin(), validatorUptimeGin())
}

func validatorUptimeGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		address := lib.Bech32AddressToHex(c.Param("address"))

		now := time.Now()
		start := utils.NDaysAgo(now, 28).Unix()
		end := now.Unix()
		step := "1d"

		if d, err := strconv.ParseInt(c.Query("start"), 10, 64); err == nil {
			start = d
		}
		if d, err := strconv.ParseInt(c.Query("end"), 10, 64); err == nil {
			end = d
		}
		if d := c.Query("step"); d != "" {
			if utils.IsDigit(d) {
				step = d + "s"
			} else {
				step = d
			}
		}

		res, err := metric.QueryValidatorUptime(node.ChainID, address,
			time.Unix(start, 0), time.Unix(end, 0), step)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
