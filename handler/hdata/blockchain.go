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

		lb, err := node.LatestBlock()
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		if maxHeight == 0 {
			maxHeight = lb.Height
		}

		if maxHeight >= 20 {
			minHeight = maxHeight - 19
		} else {
			minHeight = 1
		}

		bs, err := node.Blocks(minHeight, maxHeight)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		res := blockchainResp{
			LastHeight: lb.Height,
			Blocks:     bs,
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
