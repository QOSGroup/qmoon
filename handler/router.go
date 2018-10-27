// Copyright 2018 The QOS Authors

package handler

import "github.com/gin-gonic/gin"

func NewGinRouter() *gin.Engine {
	r := gin.Default()

	//r.Use(AuthGin())

	VersionGinRegister(r)
	AccountRegisterGinRegister(r)
	AdminAccountGinRegister(r)

	return r
}
