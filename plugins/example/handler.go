// Copyright 2018 The QOS Authors

package example

import (
	"net/http"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const exampleUrl = types.UrlNodeProxy + "/plugins/example"

// ExampleGinRegister 注册Example
func ExampleGinRegister(r *gin.Engine) {
	r.GET(types.UrlNodeProxy+exampleUrl, ExampleGin())
}

func ExampleGin() gin.HandlerFunc {
	return func(c *gin.Context) {

		res, err := DoNothing()
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
