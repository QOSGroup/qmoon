// Copyright 2018 The QOS Authors

package hdata

import (
	"github.com/QOSGroup/qmoon/lib/qos"
	"github.com/QOSGroup/qmoon/models"
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"

)

const blockUrl = "/block/:height"
const latestBlockUrl = "/latestblock"

func init() {
	hdataHander[blockUrl] = BlockGinRegister
	hdataHander[latestBlockUrl] = LatestBlockGinRegister
}

// BlockGinRegister 注册block
func BlockGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+blockUrl, middleware.ApiAuthGin(), blockGin())
}

func LatestBlockGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+latestBlockUrl, middleware.ApiAuthGin(), latestBlockGin())
}

func blockGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		height := c.Param("height")
		d, err := strconv.ParseInt(height, 10, 64)
		if err != nil {
			d = 0
		}

		var b *types.ResultBlockBase
		resp := &types.ResultBlock{}

		if d == 0 {
			b, err = node.LatestBlockFromCli()
			// b, err = node.LatestBlock()
			if err != nil {
				c.JSON(http.StatusOK, types.RPCServerError("", err))
				return
			}

		} else {
			b, err = node.RetrieveBlock(d)
			if err != nil {
				b, err = node.BlockByHeight(d)
				if err != nil {
					c.JSON(http.StatusOK, types.RPCServerError("", err))
					return
				}
			}
		}
		offset, _ := strconv.ParseInt(c.Query("offset"), 10, 64)
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)

		ts, _ := node.Txs(b.Height, 0, 0, offset, limit)
		vs, _ := node.BlockValidatorsByHeight(b.Height)
		mv, _ := models.RetrieveMissingValidators(node.ChainID, d)
		ev, _ := models.RetrieveEvidenceByHeight(node.ChainID, d)
		//er := make([]*types.ResultEvidenceValidator, 0)
		if ev != nil && len(ev) > 0 {
			er := make([]*types.ResultEvidenceValidator, 0)
			for _, e := range(ev) {
				ve, err := node.RetrieveValidator(e.ValidatorAddress)
				if err != nil {
					continue
				}
				rev := types.ResultEvidenceValidator{
					Height: e.Height,
					Validator: ve,
					Hash: e.Hash,
					Data: e.Data,
					CreatedAtUnix: e.CreatedAtUnix,
				}
				er = append(er, &rev)
			}
			resp.Evidences = er
		}

		resp.Block = b
		resp.Txs = ts
		resp.Validators = vs
		resp.Missings = mv

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", resp))
	}
}

func latestBlockGin() gin.HandlerFunc {
	return blockGin()
}

func queryBlockIfNotExist(context *gin.Context, height int64) (*types.ResultBlockBase, error) {
	node, err := GetNodeFromUrl(context)
	if err != nil {
		return nil, err
	}

	result, err := qos.NewQosCli("").QueryBlockByHeight(node.BaseURL, height)
	if err != nil {
		return nil, err
	}

	block := types.ResultBlockBase{
		ChainID: result.Block.ChainID,
		Height: result.Block.Height,
		NumTxs: result.Block.NumTxs,
		TotalTxs: result.Block.TotalTxs,
		Data: "",
		Time: types.ResultTime(result.Block.Time),
		DataHash: result.Block.DataHash.String(),
		ValidatorsHash: result.Block.ValidatorsHash.String(),
		CreatedAt: types.ResultTime(result.Block.Time),
	}


	return &block, nil
}
