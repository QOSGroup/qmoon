// Copyright 2018 The QOS Authors

package handler

import (
	"net/http"

	"github.com/QOSGroup/qmoon/service/admin"
	"github.com/gin-gonic/gin"
)

// AccountRegisterGinRegister 注册account
func AccountRegisterGinRegister(r *gin.Engine) {
	r.POST("/register", accountRegisterGin())
}

type accountRegisterQuery struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

func accountRegisterGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqObj accountRegisterQuery
		if c.ShouldBind(&reqObj) != nil {
			c.JSON(http.StatusOK, WrapperRPCResponse("", nil,
				NewRPCError(http.StatusBadRequest, "请求参数错误", "")))
			return
		}

		res, err := admin.CreateAccount(reqObj.Mail, reqObj.Password)
		if err != nil {
			c.JSON(http.StatusOK, WrapperRPCResponse("", nil,
				NewRPCError(http.StatusInternalServerError, "请求参数错误", "")))
			return
		}
		c.JSON(http.StatusOK, WrapperRPCResponse("", res, nil))
	}
}
