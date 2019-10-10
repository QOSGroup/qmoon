// Copyright 2018 The QOS Authors

package hdata

import (
	"github.com/QOSGroup/qmoon/cache"
	"github.com/QOSGroup/qmoon/lib/qos"
	"net/http"
	"strconv"
	"time"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	stake_types "github.com/QOSGroup/qmoon/lib/qos/stake/types"
	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const (
	validatorsUrl          = "/validators"
	validatorsBondTokenUrl = "/validator/bond/tokens"

	validatorsCacheKey          = "validators"
	validatorsBondTokenCacheKey = "validators_bond_tokens"
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

		err = updateValidatorsFromAgent(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		vs, err := node.Validators()
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		for i := 0; i < len(vs); i++ {
			vs[i].ConsPubKey = lib.PubkeyToBech32Address(node.Bech32PrefixConsPub(), vs[i].PubKeyType, vs[i].PubKeyValue)
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", vs))
	}
}

func updateValidatorsFromAgent(context *gin.Context) error {
	node, err := GetNodeFromUrl(context)
	if err != nil {
		return err
	}
	vals, err := queryValidators(context)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	for _, val := range vals {
		validator, err := node.ConvertDisplayValidators(val)
		node.CreateValidator(validator)
		//old, err := models.ValidatorByAddress(node.ChanID, val.OperatorAddress)
		//if err != nil {
		//	return err
		//}
		//bondTokens_int64, err := strconv.ParseInt(val.BondTokens, 10, 64)
		//if err != nil {
		//	return types.NewInvalidTypeError(val.BondTokens, "int64")
		//}
		//selfBond_int64, err := strconv.ParseInt(val.SelfBond, 10, 64)
		//if err != nil {
		//	return types.NewInvalidTypeError(val.BondTokens, "int64")
		//}
		//
		//status_int8 := int(0)
		//if val.Status != "active" {
		//	status_int8 = int(1)
		//}
		//inactive_int8 := int64(0)
		//if val.InactiveDesc != "" {
		//	inactive_int8, err = strconv.ParseInt(val.InactiveDesc, 10, 8)
		//	if err != nil {
		//		return types.NewInvalidTypeError(val.InactiveDesc, "int8")
		//	}
		//}
		//
		//old.Status = status_int8
		//old.InactiveCode = int(inactive_int8)
		//old.Name = val.Description.Moniker
		//old.Details = val.Description.Details
		//old.Logo = val.Description.Logo
		//old.Website = val.Description.Website
		//old.Owner = val.Owner
		//old.Commission = val.Commission.Rate
		//old.BondedTokens = bondTokens_int64
		//old.SelfBond = selfBond_int64
		//if err := old.Update(node.ChanID); err != nil {
		//	return err
		//}
	}
	return nil
}

func queryValidators(context *gin.Context) (result []stake_types.ValidatorDisplayInfo, err error) {
	k := validatorsCacheKey
	if v, ok := cache.Get(k); ok {
		if result, ok = v.([]stake_types.ValidatorDisplayInfo); ok {
			return
		}
	}

	node, err := GetNodeFromUrl(context)
	if err != nil {
		return
	}

	result, err = qos.NewQosCli("").QueryValidators(node.BaseURL)
	//result = context.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", vals))
	if err == nil {
		cache.Set(k, result, time.Minute*5)
	}
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
