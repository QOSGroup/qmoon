// Copyright 2018 The QOS Authors

package hadmin

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const updatePasswordUrl = "/password"

// UpdatePasswordGinRegister 更新密码
func UpdatePasswordGinRegister(r *gin.Engine) {
	r.PUT(updatePasswordUrl, middleware.AccountSessionGin(), updatePasswordGin())
}

type updatePasswordQuery struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func (q updatePasswordQuery) Validator() error {
	if q.OldPassword == "" {
		return errors.New("旧密码不能为空")
	}

	if q.NewPassword == "" {
		return errors.New("新密码不能为空")
	}

	if q.NewPassword == q.OldPassword {
		return errors.New("新旧密码不能相同")
	}

	return nil
}

func updatePasswordGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqObj updatePasswordQuery
		if err := c.ShouldBind(&reqObj); err != nil {
			c.JSON(http.StatusOK, types.RPCParseError("", err))
			return
		}

		fmt.Printf("req:%+v\n", reqObj)
		if err := reqObj.Validator(); err != nil {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
			return
		}

		acc, err := middleware.GetUserFromRequest(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		err = acc.UpdatePassword(reqObj.NewPassword)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", nil))
	}
}
