// Copyright 2018 The QOS Authors

package middleware

import (
	"errors"
	"net/http"

	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/service/account"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

var (
	TokenKey = "QToken"
	AuthKey  = "Authorization"
)

func AccountSessionGin() func(c *gin.Context) {
	return func(c *gin.Context) {
		t := c.GetHeader(TokenKey)
		if t == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		acc, err := service.CheckSession(t)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set(types.AuthUserKey, acc)
	}
}

func ApiAuthGin() func(c *gin.Context) {
	return func(c *gin.Context) {
		a := c.GetHeader(AuthKey)
		if a == "" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		app, err := account.AppBySecretKey(a)
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Set(types.AuthAppKey, app)

		acc, err := app.Account()
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Set(types.AuthUserKey, acc)
	}
}

func GetUserFromRequest(c *gin.Context) (*account.Account, error) {
	d, ok := c.Get(types.AuthUserKey)
	if !ok {
		return nil, errors.New("服务异常")
	}

	acc, ok := d.(*account.Account)
	if !ok {
		return nil, errors.New("服务异常")
	}

	return acc, nil
}

func GetAppFromRequest(c *gin.Context) (*account.App, error) {
	d, ok := c.Get(types.AuthAppKey)
	if !ok {
		return nil, errors.New("服务异常")
	}

	app, ok := d.(*account.App)
	if !ok {
		return nil, errors.New("服务异常")
	}

	return app, nil
}
