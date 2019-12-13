// Copyright 2018 The QOS Authors

package hdata

import (
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/lib/qos"
	"github.com/QOSGroup/qmoon/models"
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/types"
	stype "github.com/QOSGroup/qmoon/lib/qos/stake/types"
	"github.com/gin-gonic/gin"
)

const accountQueryUrl = "/accounts/:address"
const accountTxsUrl = "/accounts/:address/txs"
const accountDelegationsUrl = "/accounts/:address/delegations"

func init() {
	hdataHander[accountQueryUrl] = AccountQueryGinRegister
	hdataHander[accountTxsUrl] = AccountTxsGinRegister
	hdataHander[accountDelegationsUrl] = AccountDelegationsGinRegister
}

// AccountQueryGinRegister 注册accountQuery
func AccountQueryGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+accountQueryUrl, middleware.ApiAuthGin(), accountQueryGin())
}
// AccountTxsGinRegister 注册accountTxs
func AccountTxsGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+accountTxsUrl, middleware.ApiAuthGin(), accountTxsGin())
}

func AccountDelegationsGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+accountDelegationsUrl, middleware.ApiAuthGin(), accountDelegationsGin())
}

func accountQueryGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		address := c.Param("address")
		result, err := service.QueryAccount(node, address)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		c.JSON(http.StatusOK, types.RPCResponse{JSONRPC: "2.0", ID: "", Result: result})
	}
}

func accountTxsGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		address := c.Param("address")
		maxId, _ := strconv.ParseInt(c.Query("maxHeight"), 10, 64)

		minId, _ := strconv.ParseInt(c.Query("minHeight"), 10, 64)

		offset, limit := 0, 20

		if d, err := strconv.ParseInt(c.Query("offset"), 10, 64); err == nil {
			offset = int(d)
		}

		if d, err := strconv.ParseInt(c.Query("limit"), 10, 64); err == nil {
			limit = int(d)
		}

		ts, err := node.TxsByAddress(address, minId, maxId, offset, limit)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		res := ts

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}

func accountDelegationsGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		dels, err := qos.NewQosCli("").QueryDelegationsWithDelegator(node.BaseURL, c.Param("address"))
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		delsWithVals := stype.DelegationsWithValidator{}
		for _, del := range dels {
			v, err := node.RetrieveValidatorByStakingAddress(del.ValidatorAddr)
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
				v, err = node.RetrieveValidatorByStakingAddress(del.ValidatorAddr)
				if err != nil {
					c.JSON(http.StatusOK, types.RPCServerError("", err))
					return

				}
			}

			validatorHistory, err := models.ValidatorHistoryByAddress(node.ChainID, v.Address, 0,0,1)
			if err == nil && validatorHistory != nil {
				v.Percent = strconv.FormatFloat(float64(validatorHistory[0].VotingPower)/float64(validatorHistory[0].TotalPower)*100, 'f', -2, 64)
			}
			uptimePercent, _ := models.QueryValidatorUptime(node.ChainID, v.Address, 0,1)
			if uptimePercent != nil && len(uptimePercent)>0 {
				v.UptimeFloat, _= strconv.ParseFloat(uptimePercent[0].Y, 64)
				v.Uptime = uptimePercent[0].Y
			}

			delsWithVals = append(delsWithVals, stype.DelegationQueryWithValidatorResult{
				DelegatorAddr: del.DelegatorAddr,
				Validator: v,
				ValidatorConsensusPubKey: del.ValidatorConsensusPubKey,
				Amount: del.Amount,
				IsCompound: del.IsCompound,
			})
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", delsWithVals))
	}
}