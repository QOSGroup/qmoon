// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const validatorUrl = "/validators/:address"

func init() {
	hdataHander[validatorUrl] = ValidatorGinRegister
}

// ValidatorGinRegister 注册validator
func ValidatorGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+validatorUrl, middleware.ApiAuthGin(), validatorGin())
}

func validatorGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		address := lib.Bech32AddressToHex(c.Param("address"))
		v, err := node.RetrieveValidator(address)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		var minHeight, maxHeight int64
		maxHeightStr := c.Query("maxHeight")
		maxHeight, _ = strconv.ParseInt(maxHeightStr, 10, 64)

		minHeightStr := c.Query("minHeight")
		minHeight, _ = strconv.ParseInt(minHeightStr, 10, 64)
		bs, err := node.BlockValidatorByAddress(address, minHeight, maxHeight)

		var result types.ResultValidator
		result.Validator = v
		result.Blocks = bs

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", result))
	}
}
