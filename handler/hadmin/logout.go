// Copyright 2018 The QOS Authors

package hadmin

import (
	"net/http"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const logoutUrl = "/admin/logout"

// LogoutGinRegister 注册account
func LogoutGinRegister(r *gin.Engine) {
	r.POST(logoutUrl, middleware.AccountSessionGin(), logoutGin())
}

func logoutGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		acc, err := middleware.GetUserFromRequest(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInternalError("", err))
			return
		}

		err = service.Logout(acc.ID)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", nil))
	}
}
