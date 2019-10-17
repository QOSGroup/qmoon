// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/lib/qos"
	stake_types "github.com/QOSGroup/qmoon/lib/qos/stake/types"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const validatorUrl = "/validators/:address"
const validatorDelegationUrl = "/validators/:address/delegation"

func init() {
	hdataHander[validatorUrl] = ValidatorGinRegister
	hdataHander[validatorDelegationUrl] = ValidatorDelegationGinRegister
}

// ValidatorGinRegister 注册validator
func ValidatorGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+validatorUrl, middleware.ApiAuthGin(), validatorGin())
}

func ValidatorDelegationGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+validatorDelegationUrl, middleware.ApiAuthGin(), validatorDelegationGin())
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

		v.ConsPubKey = lib.PubkeyToBech32Address(node.Bech32PrefixConsPub(), v.PubKeyType, v.PubKeyValue)
		result.Validator = v
		result.Blocks = bs

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", result))
	}
}

func validatorDelegationGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		address := lib.Bech32AddressToHex(c.Param("address"))
		result, err := qos.NewQosCli("").QueryDelegationsWithValidator(node.BaseURL, address)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		delegations := stake_types.Delegations{}
		delegations.DelagationList = result
		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", delegations))
	}
}
