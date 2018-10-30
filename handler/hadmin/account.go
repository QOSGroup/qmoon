// Copyright 2018 The QOS Authors

package hadmin

import (
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service/account"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

// AccountGinRegister 账户相关API
func AccountGinRegister(r *gin.Engine) {
	r.GET("/admin/accounts", middleware.AccountSessionGin(), listAccountsGin())
	r.GET("/admin/accounts/:id", middleware.AccountSessionGin(), retrieveAccountGin())
	r.PUT("/admin/accounts/:id", middleware.AccountSessionGin(), updateAccountGin())
	r.DELETE("/admin/accounts/:id", middleware.AccountSessionGin(), deleteAccountGin())
}

func listAccountsGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var offset, limit int64 = 0, 20
		offset = c.GetInt64("offset")
		limit = c.GetInt64("limit")

		res, err := account.List(offset, limit)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}

func retrieveAccountGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Param("id")
		id, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
			return
		}

		res, err := account.RetrieveAccountByID(id)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))

	}
}

func updateAccountGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))

	}
}

func deleteAccountGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))

	}
}
