// Copyright 2018 The QOS Authors

package hdata

import (
	"encoding/json"
	"github.com/QOSGroup/qmoon/models"
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/lib/qos"
	"github.com/QOSGroup/qmoon/service"
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

		//address := lib.Bech32AddressToHex(c.Param("address"))
		stakingAddress := c.Param("address")
		v, err := node.RetrieveValidatorByStakingAddress(stakingAddress)
		if err != nil {
			//c.JSON(http.StatusOK, types.RPCServerError("", err))
			//return
			vals_display, err := qos.NewQosCli("").QueryValidators(node.BaseURL)
			if err != nil {
				c.JSON(http.StatusOK, types.RPCServerError("", err))
				return
			}
			for _, dist := range vals_display {
				val, err := node.ConvertDisplayValidators(dist)
				if err != nil {
					c.JSON(http.StatusOK, types.RPCServerError("", err))
					return
				}
				node.CreateValidator(val)
			}
			v, err = node.RetrieveValidatorByStakingAddress(stakingAddress)
			if err != nil {
				c.JSON(http.StatusOK, types.RPCServerError("", err))
				return
			}
		}

		var minHeight, maxHeight, offset, limit int64
		maxHeightStr := c.Query("maxHeight")
		maxHeight, _ = strconv.ParseInt(maxHeightStr, 10, 64)

		minHeightStr := c.Query("minHeight")
		minHeight, _ = strconv.ParseInt(minHeightStr, 10, 64)

		offsetStr := c.Query("offset")
		offset, _ = strconv.ParseInt(offsetStr, 10, 64)

		limitStr := c.Query("limit")
		limit, _ = strconv.ParseInt(limitStr, 10, 64)
		bs, err := node.BlockValidatorByAddress(v.Address, minHeight, maxHeight, int(offset), int(limit))

		var result types.ResultValidator


		validatorHistory, err := models.ValidatorHistoryByAddress(node.ChainID, v.Address, 0, 0,1)
		if err == nil && validatorHistory != nil {
			v.Percent = strconv.FormatFloat(float64(validatorHistory[0].VotingPower)/float64(validatorHistory[0].TotalPower)*100, 'f', -2, 64)
		}
		uptimePercent, _ := models.QueryValidatorUptime(node.ChainID, v.Address, 0,1)
		if uptimePercent != nil && len(uptimePercent)>0 {
			v.UptimeFloat, _= strconv.ParseFloat(uptimePercent[0].Y, 64)
			v.Uptime = uptimePercent[0].Y
		}

		//v.ConsPubKey = lib.PubkeyToBech32Address(node.Bech32PrefixConsPub(), v.PubKeyType, v.PubKeyValue)
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

		dels, err := qos.NewQosCli("").QueryDelegationsWithValidator(node.BaseURL, c.Param("address"))
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		var result types.ResultValidator
		v, err := models.ValidatorByStakeAddress(node.ChainID, c.Param("address"))
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		// latest, err := node.LatestBlock()
		latestheight, err := node.LatestBlockHeight()
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		result.Validator = service.ConvertToValidator(v, latestheight)
		validatorHistory, err := models.ValidatorHistoryByAddress(node.ChainID, v.Address, 0, 0, 1)
		if err == nil && validatorHistory != nil {
			result.Validator.Percent = strconv.FormatFloat(float64(validatorHistory[0].VotingPower)/float64(validatorHistory[0].TotalPower)*100, 'f', -2, 64)
		}

		uptimePercent, _ := models.QueryValidatorUptime(node.ChainID, result.Validator.Address, 0,1)
		if uptimePercent != nil && len(uptimePercent)>0 {
			result.Validator.Uptime = uptimePercent[0].Y
			result.Validator.UptimeFloat, _= strconv.ParseFloat(uptimePercent[0].Y, 64)
		}

		for _, d := range dels {
			delegation := types.ResultDelagation{Delegator: d.DelegatorAddr, Amount: d.Amount, Compound: d.IsCompound}
			j, _ := json.Marshal(delegation)
			result.Delegations = append(result.Delegations, j)
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", result))
	}
}
