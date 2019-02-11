// Copyright 2018 The QOS Authors

package hdata

import (
	"errors"

	"github.com/QOSGroup/qmoon/service"
	"github.com/gin-gonic/gin"
)

var hdataHander = make(map[string]func(r *gin.Engine))

func GinRegister(r *gin.Engine) {
	for _, v := range hdataHander {
		v(r)
	}
}

func GetNodeFromUrl(c *gin.Context) (*service.Node, error) {
	nodeName := c.Param("nodeName")
	if nodeName == "" {
		return nil, errors.New("invalid nodeName")
	}
	node, err := service.GetNodeByName(nodeName)
	if err != nil {
		return nil, errors.New("nodeName not found")
	}

	return node, nil
}
