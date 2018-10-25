// Copyright 2018 The QOS Authors

package handler

import "github.com/gin-gonic/gin"

// AdminNodeTypesGinRegister 注册account相关的api
func AdminNodeTypesGinRegister(r *gin.Engine) {
	r.POST("/admin/node_types", createAdminNodeTypeGin())
	r.GET("/admin/node_types", queryAdminNodeTypesGin())
}

func createAdminNodeTypeGin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func queryAdminNodeTypesGin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
