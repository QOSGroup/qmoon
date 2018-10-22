// Copyright 2018 The QOS Authors

package handler

import (
	"net/http"

	"github.com/QOSGroup/qmoon/service/account"
	"github.com/QOSGroup/qmoon/service/auth"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

func AuthGin() func(c *gin.Context) {
	return func(c *gin.Context) {
		a := c.GetHeader("Authorization")
		if a == "" {
			c.AbortWithStatus(http.StatusForbidden)
		}

		sd,err := auth.ParseSign([]byte(a))
		if err !=nil{
			c.AbortWithError(http.StatusForbidden, err)
		}

		ac,err :=account.Get(sd.SecretID)
		if err !=nil{
			c.AbortWithError(http.StatusForbidden, err)
		}

		s, err := sd.Sign([]byte(ac.SecretKey))
		if err !=nil{
			c.AbortWithError(http.StatusForbidden, err)
		}

		if string(s) != a{
			c.AbortWithStatus(http.StatusForbidden)
		}

		c.Set(types.AuthUserKey, ac)
	}
}
