// Copyright 2018 The QOS Authors

package middleware

import (
	"errors"
	"net/http"

	"github.com/QOSGroup/qmoon/models"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

func AccountSessionGin() func(c *gin.Context) {
	return func(c *gin.Context) {
		t := c.GetHeader(types.TokenKey)
		if t == "" {
			c.AbortWithStatusJSON(http.StatusOK, types.RPCUnauthorizedError("", errors.New("no QToken")))
			return
		}

		acc, err := service.CheckSession(t)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, types.RPCUnauthorizedError("", errors.New("QToken Invalid")))
			return
		}

		c.Set(types.AuthUserKey, acc)
	}
}

func ApiAuthGin() func(c *gin.Context) {
	return func(c *gin.Context) {
		{
			return
		}
		a := c.GetHeader(types.TokenKey)
		if a == "" {
			c.AbortWithStatusJSON(http.StatusOK, types.RPCForbiddenError("", errors.New("no Authorization")))
			return
		}

		app, err := models.AppBySecretKey(a)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, types.RPCForbiddenError("", errors.New("invalid Authorization")))
			return
		}
		c.Set(types.AuthAppKey, app)

		acc, err := app.Account()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, types.RPCForbiddenError("", errors.New("invalid Authorization")))
			return
		}

		c.Set(types.AuthUserKey, acc)
	}
}

func GetUserFromRequest(c *gin.Context) (*models.Account, error) {
	d, ok := c.Get(types.AuthUserKey)
	if !ok {
		return nil, errors.New("服务异常")
	}

	acc, ok := d.(*models.Account)
	if !ok {
		return nil, errors.New("服务异常")
	}

	return acc, nil
}

func GetAppFromRequest(c *gin.Context) (*models.App, error) {
	d, ok := c.Get(types.AuthAppKey)
	if !ok {
		return nil, errors.New("服务异常")
	}

	app, ok := d.(*models.App)
	if !ok {
		return nil, errors.New("服务异常")
	}

	return app, nil
}
