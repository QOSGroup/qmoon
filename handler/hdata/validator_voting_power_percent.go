// Copyright 2018 The QOS Authors

package hdata

import (
	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

const validatorVotingPowerPercentUrl = NodeProxy + "/validators/:address/votingPowerPercent"

func init() {
	hdataHander[validatorVotingPowerPercentUrl] = ValidatorVotingPowerPercentGinRegister
}

// ValidatorVotingPowerGinRegister 注册
func ValidatorVotingPowerPercentGinRegister(r *gin.Engine) {
	r.GET(validatorVotingPowerPercentUrl, middleware.ApiAuthGin(), validatorVotingPowerPercentGin())
}

func validatorVotingPowerPercentGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		address := lib.Bech32AddressToHex(c.Param("address"))

		res, err := models.QueryValidatorVotingPowerPercent(node.ChainID, address)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
