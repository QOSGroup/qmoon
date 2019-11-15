// Copyright 2018 The QOS Authors

package middleware

import (
	"errors"
	"net/http"

	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

func GetNodeFromUrl() func(c *gin.Context) {
	return func(c *gin.Context) {
		nodeName := c.Param("nodeName")
		if nodeName == "" {
			c.AbortWithStatusJSON(http.StatusOK, types.RPCMethodNotFoundError(""))
		}
		node, err := service.GetNodeByName(nodeName)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, types.RPCMethodNotFoundError(""))
		}

		c.Set(types.AuthNode, node)
	}
}

func GetNodeFromRequest(c *gin.Context) (*service.Node, error) {
	d, ok := c.Get(types.AuthNode)
	if !ok {
		return nil, errors.New("服务异常")
	}

	node, ok := d.(*service.Node)
	if !ok {
		return nil, errors.New("服务异常")
	}

	return node, nil
}
