// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const validatorMissingUrl = NodeProxy + "/validators/:address/missing"

func init() {
	hdataHander[validatorMissingUrl] = ValidatorMissingGinRegister
}

// ValidatorVotingPowerGinRegister 注册
func ValidatorMissingGinRegister(r *gin.Engine) {
	r.GET(validatorMissingUrl, middleware.ApiAuthGin(), validatorMissingGin())
}

func validatorMissingGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		address := lib.Bech32AddressToHex(c.Param("address"))
		res, err := node.Missings(address)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
