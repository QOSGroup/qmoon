// Copyright 2018 The QOS Authors

package hdata

import (
	"net/http"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const validatorsUrl = "/validators"

func init() {
	hdataHander[validatorsUrl] = ValidatorsGinRegister
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

		vs, err := node.Validators()
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", vs))
	}
}
