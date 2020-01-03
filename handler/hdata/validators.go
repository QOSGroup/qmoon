// Copyright 2018 The QOS Authors

package hdata

import (
	"github.com/QOSGroup/qmoon/cache"
	"github.com/QOSGroup/qmoon/lib/qos"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/service"
	"net/http"
	"strconv"
	"time"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	stake_types "github.com/QOSGroup/qmoon/lib/qos/stake/types"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const (
	validatorsUrl          = "/validators"
	validatorsBondTokenUrl = "/validator/bond/tokens"

	validatorsCacheKey          = "validators"
	validatorsBondTokenCacheKey = "validators_bond_tokens"
	validatorKeyPrifix          = "validator_"
)

func init() {
	hdataHander[validatorsUrl] = ValidatorsGinRegister
	hdataHander[validatorsBondTokenUrl] = func(r *gin.Engine) {
		r.GET(NodeProxy+validatorsBondTokenUrl, middleware.ApiAuthGin(), func(context *gin.Context) {
			bondToken, err := queryBondTokens(context)
			if err != nil {
				context.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
				return
			}
			context.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", bondToken))
		})
	}
}

// ValidatorsGinRegister 注册validators
func ValidatorsGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+validatorsUrl, middleware.ApiAuthGin(), validatorsGin())
}

func validatorsGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}
		vs := types.Validators{}
		vss, ok := cache.Get("qosvalidators")
		if !ok {
			err = updateValidatorsFromAgent(c)
			if err != nil {
				c.JSON(http.StatusOK, types.RPCServerError("", err))
				return
			}
			d, ok := cache.Get(service.LatestHeightKey)
			if ok {
				if v, okk := d.(int64); okk {
					vs, err = node.Validators(v)
				} else {
					h, err := node.LatestBlockHeight()
					if err == nil && h > 0 {
						vs, err = node.Validators(h)
					}
				}
			} else {
				h, err := node.LatestBlockHeight()
				if err == nil && h > 0 {
					vs, err = node.Validators(h)
				}
			}

			if err != nil {
				c.JSON(http.StatusOK, types.RPCServerError("", err))
				return
			}

			if len(vs) != 0 {
				cache.Set("qosvalidators", vs, time.Minute*5)
			}
		} else {
			vs = vss.(types.Validators)
		}

		for i := 0; i < len(vs); i++ {
			vs[i].ConsPubKey = lib.PubkeyToBech32Address(node.Bech32PrefixConsPub(), vs[i].PubKeyType, vs[i].PubKeyValue)
			validatorHistory, err := models.ValidatorHistoryByAddress(node.ChainID, vs[i].Address, 0, 0,1)
			if err == nil && validatorHistory != nil {
				vs[i].Percent = strconv.FormatFloat(float64(validatorHistory[0].VotingPower)/float64(validatorHistory[0].TotalPower)*100, 'f', -2, 64)
			}
			uptimePercent, _ := models.QueryValidatorUptime(node.ChainID, vs[i].Address, 720,1)
			if uptimePercent != nil && len(uptimePercent)>0 {
				vs[i].UptimeFloat, _= strconv.ParseFloat(uptimePercent[0].Y, 64)
				vs[i].Uptime = uptimePercent[0].Y
			}
		}


		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", vs))
	}
}

func updateValidatorsFromAgent(context *gin.Context) error {
	node, err := GetNodeFromUrl(context)
	if err != nil {
		return err
	}
	// vals, err := queryValidators(context)
	vals, err := qos.NewQosCli("").QueryValidators(node.BaseURL)
	if err != nil {
		return err
	}
	for _, val := range vals {
		validator, err := node.ConvertDisplayValidators(val)
		if err != nil {
			return err
		}
		node.CreateValidator(validator)
	}
	return nil
}

func queryValidators(context *gin.Context) (result []stake_types.ValidatorDisplayInfo, err error) {
	//k := validatorsCacheKey
	//if v, ok := cache.Get(k); ok {
	//	if result, ok = v.([]stake_types.ValidatorDisplayInfo); ok {
	//		return
	//	}
	//}

	node, err := GetNodeFromUrl(context)
	if err != nil {
		return
	}

	result, err = qos.NewQosCli("").QueryValidators(node.BaseURL)
	//if err == nil {
	//	cache.Set(k, result, time.Minute*5)
	//}
	return
}

func queryBondTokens(context *gin.Context) (result string, err error) {
	k := validatorsBondTokenCacheKey
	if v, ok := cache.Get(k); ok {
		if result, ok = v.(string); ok {
			return
		}
	}

	node, err := GetNodeFromUrl(context)
	if err != nil {
		return
	}

	result, err = qos.NewQosCli("").QueryTotalValidatorBondTokens(node.BaseURL)
	if err == nil {
		cache.Set(k, result, time.Minute*5)
	}
	return
}
