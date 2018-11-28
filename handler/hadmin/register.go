// Copyright 2018 The QOS Authors

package hadmin

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/service/account"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const registerUrl = "/admin/register"

// RegisterGinRegister 注册account
func RegisterGinRegister(r *gin.Engine) {
	r.POST(registerUrl, registerGin())
}

type registerQuery struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

func (q registerQuery) Validator() error {
	if q.Email == "" {
		return errors.New("email不能为空")
	}

	if q.Password == "" {
		return errors.New("密码不能为空")
	}
	if q.Code == "" {
		return errors.New("验证码不能为空")
	}

	return nil
}

func registerGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqObj registerQuery
		if err := c.ShouldBind(&reqObj); err != nil {
			c.JSON(http.StatusOK, types.RPCParseError("", err))
			return
		}

		fmt.Printf("req:%+v\n", reqObj)
		if err := reqObj.Validator(); err != nil {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
			return
		}

		if ok := service.CheckCode(reqObj.Email, reqObj.Code); !ok {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", errors.New("验证码不正确")))
			return
		}

		res, err := account.CreateAccount(reqObj.Email, reqObj.Password)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
