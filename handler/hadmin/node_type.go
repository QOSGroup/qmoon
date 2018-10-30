// Copyright 2018 The QOS Authors

package hadmin

import (
	"errors"
	"net/http"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/gin-gonic/gin"
)

const nodeTypeUrl = "/nodeTypes"

// NodeTypesGinRegister 注册node type相关的api
func NodeTypesGinRegister(r *gin.Engine) {
	r.POST(nodeTypeUrl, middleware.IPWhitelist(utils.NewLocalIPNet()), createNodeTypeGin())
	r.GET(nodeTypeUrl, listNodeTypesGin())
}

type createNodeTypeQuery struct {
	Name      string                  `json:"name"`
	BaseURL   string                  `json:"baseUrl"`
	SecretKey string                  `json:"secretKey"`
	Routers   []service.NodeTypeRoute `json:"routers"`
}

func (q createNodeTypeQuery) Validator() error {
	if q.Name == "" {
		return errors.New("名字不能为空")
	}

	if q.BaseURL == "" {
		return errors.New("baseUrl不能为空")
	}

	return nil
}

func createNodeTypeGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqObj createNodeTypeQuery
		if err := c.ShouldBind(&reqObj); err != nil {
			c.JSON(http.StatusOK, types.RPCParseError("", err))
			return
		}

		if err := reqObj.Validator(); err != nil {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
			return
		}

		err := service.CreateNodeType(reqObj.Name, reqObj.BaseURL, reqObj.SecretKey, reqObj.Routers)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", nil))
	}
}

func listNodeTypesGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := service.AllNodeTypes()
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
