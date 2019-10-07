package hdata

import (
	"github.com/QOSGroup/qmoon/cache"
	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/lib/qos"
	qos_mint_types "github.com/QOSGroup/qmoon/lib/qos/mint/types"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	inflationPhrasesUrl = "/inflation/phrases"
	appliedUrl          = "/applied"

	inflationPhrasesCacheKey = "inflation_phrases_key"
	appliedCacheKey          = "applied_key"
)

func init() {
	hdataHander[inflationPhrasesUrl] = func(r *gin.Engine) {
		r.GET(NodeProxy+inflationPhrasesUrl, middleware.ApiAuthGin(), func(context *gin.Context) {
			proposal, err := queryInflationPhrases(context)
			if err != nil {
				context.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
				return
			}
			context.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", proposal))
		})
	}
	hdataHander[appliedUrl] = func(r *gin.Engine) {
		r.GET(NodeProxy+appliedUrl, middleware.ApiAuthGin(), func(context *gin.Context) {
			applied, err := queryApplied(context)
			if err != nil {
				context.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
				return
			}
			context.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", applied))
		})
	}
}

func queryInflationPhrases(context *gin.Context) (inflationPhrases []qos_mint_types.InflationPhrase, err error) {
	k := inflationPhrasesCacheKey
	if v, ok := cache.Get(k); ok {
		if inflationPhrases, ok = v.([]qos_mint_types.InflationPhrase); ok {
			return
		}
	}

	node, err := GetNodeFromUrl(context)
	if err != nil {
		return
	}

	inflationPhrases, err = qos.NewQosCli("").QueryInflationPhrases(node.BaseURL)
	if err == nil {
		cache.Set(k, inflationPhrases, time.Minute*5)
	}
	return
}

func queryApplied(context *gin.Context) (applied string, err error) {
	k := appliedCacheKey
	if v, ok := cache.Get(k); ok {
		if applied, ok = v.(string); ok {
			return
		}
	}

	node, err := GetNodeFromUrl(context)
	if err != nil {
		return
	}

	applied, err = qos.NewQosCli("").QueryApplied(node.BaseURL)
	if err == nil {
		cache.Set(k, applied, time.Minute*5)
	}
	return
}
