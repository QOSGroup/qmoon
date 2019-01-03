// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service/block"
	"github.com/QOSGroup/qmoon/service/tx"
	"github.com/QOSGroup/qmoon/service/validator"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const blockUrl = "/block/:height"

func init() {
	hdataHander[blockUrl] = BlockGinRegister
}

// BlockGinRegister 注册block
func BlockGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+blockUrl, middleware.ApiAuthGin(), blockGin())
}

func blockGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		nt, err := getNodeFromUrl(c)
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
		if d == 0 {
			b, err = block.Latest(nt.ChanID)
			if err != nil {
				c.JSON(http.StatusOK, types.RPCServerError("", err))
				return
			}

		} else {
			b, err = block.Retrieve(nt.ChanID, d)
			if err != nil {
				c.JSON(http.StatusOK, types.RPCServerError("", err))
				return
			}
		}

		ts, _ := tx.List(nt.ChanID, b.Height, b.Height)
		vs, _ := validator.ListBlockValidatorByHeight(nt.ChanID, b.Height)

		resp := &types.ResultBlock{}
		resp.Block = b
		resp.Txs = ts
		resp.Validators = vs

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", resp))
	}
}
