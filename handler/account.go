// Copyright 2018 The QOS Authors

package handler

import "github.com/gin-gonic/gin"

// AccountGinRegister 注册account相关的api
func AccountGinRegister(r *gin.Engine) {
	r.POST("/accounts", createAccountGin())
	r.GET("/accounts", queryAccountsGin())

	r.GET("/accounts/:id", queryAccountGin())
	r.PUT("/accounts/:id", updateAccountGin())
	r.DELETE("/accounts/:id", deleteAccountGin())

}

func createAccountGin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func queryAccountsGin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func queryAccountGin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func updateAccountGin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func deleteAccountGin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
