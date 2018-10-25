// Copyright 2018 The QOS Authors

package handler

import "github.com/gin-gonic/gin"

// AdminNodeTypesGinRegister 注册account相关的api
func AdminAccountGinRegister(r *gin.Engine) {
	r.POST("/admin/accounts", createAdminAccountGin())
	r.GET("/admin/accounts", queryAdminAccountsGin())

	r.GET("/admin/accounts/:id", queryAdminAccountGin())
	r.PUT("/admin/accounts/:id", updateAdminAccountGin())
	r.DELETE("/admin/accounts/:id", deleteAdminAccountGin())
}

func createAdminAccountGin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func queryAdminAccountsGin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func queryAdminAccountGin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func updateAdminAccountGin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func deleteAdminAccountGin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
