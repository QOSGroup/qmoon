// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const validatorEvidenceUrl = NodeProxy + "/validators/:address/evidence"

func init() {
	hdataHander[validatorEvidenceUrl] = ValidatorEvidenceGinRegister
}

// ValidatorVotingPowerGinRegister 注册
func ValidatorEvidenceGinRegister(r *gin.Engine) {
	r.GET(validatorEvidenceUrl, middleware.ApiAuthGin(), validatorEvidenceGin())
}

func validatorEvidenceGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		address := lib.Bech32AddressToHex(c.Param("address"))
		res, err := node.Evidences(address)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
