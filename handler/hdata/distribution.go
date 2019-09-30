package hdata

import (
	"github.com/QOSGroup/qmoon/cache"
	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/lib/qos"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	communityFeePoolUrl = "/community/fee/pool"

	communityFeePoolCacheKey = "community_fee_pool_key"
)

func init() {
	hdataHander[communityFeePoolUrl] = func(r *gin.Engine) {
		r.GET(NodeProxy+communityFeePoolUrl, middleware.ApiAuthGin(), func(context *gin.Context) {
			feePool, err := queryCommunityFeePool(context)
			if err != nil {
				context.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
				return
			}
			context.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", feePool))
		})
	}
}

func queryCommunityFeePool(context *gin.Context) (feePool string, err error) {
	k := communityFeePoolCacheKey
	if v, ok := cache.Get(k); ok {
		if feePool, ok = v.(string); ok {
			return
		}
	}

	node, err := GetNodeFromUrl(context)
	if err != nil {
		return
	}

	feePool, err = qos.NewQosCli("").QueryCommunityFeePool(node.BaseURL)
	if err == nil {
		cache.Set(k, feePool, time.Minute*5)
	}
	return
}
