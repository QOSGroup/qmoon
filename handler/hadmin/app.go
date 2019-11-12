// Copyright 2018 The QOS Authors

package hadmin

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

// AppGinRegister 注册app相关API
func AppGinRegister(r *gin.Engine) {
	r.POST("/apps", middleware.AccountSessionGin(), createAppGin())
	r.GET("/apps", middleware.AccountSessionGin(), listAppsGin())
	r.DELETE("/apps/:id", middleware.AccountSessionGin(), deleteAppGin())
}

type createAppQuery struct {
	Name string `json:"name"`
}

func (q createAppQuery) Validator() error {
	if q.Name == "" {
		return errors.New("app名字不能为空")
	}

	return nil
}

func createAppGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		acc, err := middleware.GetUserFromRequest(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		var reqObj createAppQuery
		if err := c.ShouldBind(&reqObj); err != nil {
			c.JSON(http.StatusOK, types.RPCParseError("", err))
			return
		}

		if err := reqObj.Validator(); err != nil {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
			return
		}

		res, err := acc.CreateApp(reqObj.Name)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}

func listAppsGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		acc, err := middleware.GetUserFromRequest(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		res, err := acc.Apps()
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}

func deleteAppGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		acc, err := middleware.GetUserFromRequest(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		p := c.Param("id")
		id, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
			return
		}

		err = acc.DeleteAppByID(id)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", nil))

	}
}
