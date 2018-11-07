// Copyright 2018 The QOS Authors

package hadmin

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const loginUrl = "/admin/login"

// LoginGinRegister 注册account
func LoginGinRegister(r *gin.Engine) {
	r.POST(loginUrl, loginGin())
}

type loginQuery struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q loginQuery) Validator() error {
	if q.Email == "" {
		return errors.New("email不能为空")
	}

	if q.Password == "" {
		return errors.New("密码不能为空")
	}

	return nil
}

func loginGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqObj loginQuery
		if err := c.ShouldBind(&reqObj); err != nil {
			c.JSON(http.StatusOK, types.RPCParseError("", err))
			return
		}

		fmt.Printf("req:%+v\n", reqObj)
		if err := reqObj.Validator(); err != nil {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
			return
		}

		res, err := service.Login(reqObj.Email, reqObj.Password)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
