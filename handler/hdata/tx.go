// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service/tx"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const txUrl = "/tx"

func init() {
	hdataHander[txUrl] = TxGinRegister
}

// TxGinRegister 注册tx
func TxGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+txUrl, middleware.ApiAuthGin(), txGin())
}

func txGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		nt, err := getNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		heightStr := c.Query("height")
		height, err := strconv.ParseInt(heightStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
			return
		}
		indexStr := c.Query("index")
		index, err := strconv.ParseInt(indexStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
			return
		}

		result, err := tx.Retrieve(nt.ChanID, height, index)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", result))
	}
}
