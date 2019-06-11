// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/gin-gonic/gin"
)

const txSendUrl = "/tx/send"

func init() {
	hdataHander[txSendUrl] = TxSendGinRegister
}

// TxSendGinRegister 注册txSend
func TxSendGinRegister(r *gin.Engine) {
	r.POST(NodeProxy+txSendUrl, middleware.ApiAuthGin(), txSendGin())
}

type txSendBody struct {
	Remote string `json:"remote"`
	Tx     string `json:"tx"`
}

func txSendGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		var d txSendBody
		if err := c.BindJSON(&d); err != nil {
			c.JSON(http.StatusOK, types.RPCParseError("", err))
			return
		}

		tx, err := utils.Base64De(d.Tx)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCParseError("", err))
			return
		}

		resp, err := node.TxSend(d.Remote, tx)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", resp))
	}
}
