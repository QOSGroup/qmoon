// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const statusQueryUrl = "/status"

func init() {
	hdataHander[statusQueryUrl] = StatusQueryGinRegister
}

// StatusQueryGinRegister 注册statusQuery
func StatusQueryGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+statusQueryUrl, middleware.ApiAuthGin(), statusQueryGin())
}

func statusQueryGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		result, err := node.ChainStatus(false)
		//status, err := qos.NewQosCli("").QueryStatus(node.BaseURL)
		//if err != nil {
		//	c.JSON(http.StatusOK, types.RPCInternalError("Can't get status", err))
		//	return
		//}
		//
		//consensus, err := node.ConsensusState()
		//validators, err := node.Validators()
		//height:=status.SyncInfo.LatestBlockHeight
		//block, err := node.BlockByHeight(height)
		//// txs, err := node.Txs(height, height, 0, 1)
		//current, err := node.ChainStatus(false)
		//
		//result := types.ResultStatus{
		//	ConsensusState: consensus,
		//	TotalValidators: int64(len(validators)),
		//	Height: status.SyncInfo.LatestBlockHeight,
		//	TotalTxs: block.TotalTxs,
		//	BlockTimeAvg: current.BlockTimeAvg,
		//}

		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", result))
	}
}
