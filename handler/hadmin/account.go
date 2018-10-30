// Copyright 2018 The QOS Authors

package hadmin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service/account"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/gin-gonic/gin"
)

const accountsUrl = "/admin/accounts"

// AccountGinRegister 账户相关API
func AccountGinRegister(r *gin.Engine) {
	r.GET(accountsUrl, middleware.IPWhitelist(utils.NewLocalIPNet()), listAccountsGin())
	r.GET(accountsUrl+"/:id", middleware.AccountSessionGin(), retrieveAccountGin())
	r.PUT(accountsUrl+"/:id", middleware.AccountSessionGin(), updateAccountGin())
	r.DELETE(accountsUrl+"/:id", middleware.AccountSessionGin(), deleteAccountGin())
}

func listAccountsGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var offset, limit int64 = 0, 20
		offset = c.GetInt64("offset")
		limit = c.GetInt64("limit")
		if limit == 0 {
			limit = 20
		}

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

type updateAccountQuery struct {
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

func (q updateAccountQuery) Validator() error {
	return nil
}

func updateAccountGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqObj updateAccountQuery
		if err := c.ShouldBind(&reqObj); err != nil {
			c.JSON(http.StatusOK, types.RPCParseError("", err))
			return
		}

		fmt.Printf("req:%+v\n", reqObj)
		if err := reqObj.Validator(); err != nil {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
			return
		}

		p := c.Param("id")
		id, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
			return
		}

		acc, err := account.RetrieveAccountByID(id)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		opt, err := account.NewOption(account.SetAvatar(
			reqObj.Avatar),
			account.SetName(reqObj.Name),
			account.SetDescription(reqObj.Description),
		)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		err = acc.UpdateProfile(opt)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", acc))
	}
}

func deleteAccountGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))

	}
}
