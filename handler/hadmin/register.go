// Copyright 2018 The QOS Authors

package hadmin

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/QOSGroup/qmoon/lib"
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
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

func (q registerQuery) Validator() error {
	if q.Mail == "" {
		return errors.New("mail不能为空")
	}

	if q.Password == "" {
		return errors.New("密码不能为空")
	}

	return nil
}

func (q *registerQuery) ServerHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("aaaa"))
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

		res, err := account.CreateAccount(reqObj.Mail, reqObj.Password)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
