// Copyright 2018 The QOS Authors

package hdata

import "github.com/gin-gonic/gin"

// AdminNodeTypesGinRegister 注册account相关的api
func AdminNodeTypesGinRegister(r *gin.Engine) {
	r.GET("/node_types", listNodeTypesGin())
}

func listNodeTypesGin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
