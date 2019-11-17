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

const blockchainUrl = "/blockchain"

func init() {
	hdataHander[blockchainUrl] = BlockchainGinRegister
}

// BlockchainGinRegister 注册blockchain
func BlockchainGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+blockchainUrl, middleware.ApiAuthGin(), blockchainGin())
}

type blockchainQuery struct {
	MinHeight int64 `pa:"minHeight"`
	MaxHeight int64 `json:"maxHeight"`
}

func (q blockchainQuery) Validator() error {

	return nil
}

type blockchainResp struct {
	LastHeight int64                    `json:"last_height"`
	Blocks     []*types.ResultBlockBase `json:"blocks"`
}

func blockchainGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		var minHeight, maxHeight int64
		maxHeightStr := c.Query("maxHeight")
		maxHeight, err = strconv.ParseInt(maxHeightStr, 10, 64)
		if err != nil {
			maxHeight = 0
		}

		minHeightStr := c.Query("minHeight")
		minHeight, err = strconv.ParseInt(minHeightStr, 10, 64)
		if err != nil {
			minHeight = 1
		}

		//lb, err := node.LatestBlock()
		//if err != nil {
		//	c.JSON(http.StatusOK, types.RPCServerError("", err))
		//	return
		//}

		//status, err := qos.NewQosCli("").QueryStatus(node.BaseURL)

		//if err != nil && status!= nil {
		//		c.JSON(http.StatusOK, types.RPCServerError("", err))
		//		return
		//}
		//cs, err1 := n.ConsensusState()
		//result.Height = status.SyncInfo.LatestBlockHeight
		//result, err = n.BlockByHeight(status.SyncInfo.LatestBlockHeight)

		//if maxHeight == 0 {
		//	//maxHeight
		//	//= status.SyncInfo.LatestBlockHeight
		//	maxHeight = lb.Height
		//}

		// if maxHeight >= 20 {
		// 	minHeight = maxHeight - 19
		// } else {
		// 	minHeight = 1
		// }

		offset, _ := strconv.ParseInt(c.Query("offset"), 10, 64)

		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		if limit == 0 {
			limit = 20
		}

		bs, err := node.Blocks(minHeight, maxHeight, offset, limit)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		res := blockchainResp{
			LastHeight: bs[0].Height,
			Blocks:     bs,
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
