// Copyright 2018 The QOS Authors

package hadmin

import (
	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/gin-gonic/gin"
)

// AdminNodeTypesGinRegister 注册account相关的api
func AdminNodeTypesGinRegister(r *gin.Engine) {
	r.POST("/account/node_types", middleware.AccountSessionGin(), createAdminNodeTypeGin())
	r.GET("/account/node_types", queryAdminNodeTypesGin())
}

func createAdminNodeTypeGin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func queryAdminNodeTypesGin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
