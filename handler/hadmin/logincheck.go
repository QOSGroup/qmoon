// Copyright 2018 The QOS Authors

package hadmin

import (
	"fmt"
	"net/http"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const loginCheckUrl = "/admin/login/check"

// LoginCheckGinRegister 注册登陆检查
func LoginCheckGinRegister(r *gin.Engine) {
	r.GET(loginCheckUrl, middleware.AccountSessionGin(), loginCheckGin())
}

type loginCheckQuery struct {
}

func (q loginCheckQuery) Validator() error {

	return nil
}

func loginCheckGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqObj loginCheckQuery
		if err := c.ShouldBind(&reqObj); err != nil {
			c.JSON(http.StatusOK, types.RPCParseError("", err))
			return
		}

		fmt.Printf("req:%+v\n", reqObj)
		if err := reqObj.Validator(); err != nil {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", nil))
	}
}
