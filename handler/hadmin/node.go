// Copyright 2018 The QOS Authors

package hadmin

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"github.com/gin-gonic/gin"
)

const nodeTypeUrl = "/nodes"

// NodesGinRegister 注册node type相关的api
func NodesGinRegister(r *gin.Engine) {
	r.POST(nodeTypeUrl, middleware.IPWhitelist(utils.NewLocalIPNet()), createNodeGin())
	r.GET(nodeTypeUrl, listNodesGin())
}

type createNodeQuery struct {
	Name        string `json:"name"`
	BaseURL     string `json:"baseUrl"`
	NodeType    string `json:"nodeType"`
	NodeVersion string `json:"nodeVersion"`
	SecretKey   string `json:"secretKey"`
}

func (q createNodeQuery) Validator() error {
	if q.Name == "" {
		return errors.New("名字不能为空")
	}

	if q.BaseURL == "" {
		return errors.New("baseUrl不能为空")
	}

	if q.NodeType == "" {
		return errors.New("nodeType不能为空")
	}

	return nil
}

func createNodeGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqObj createNodeQuery
		if err := c.ShouldBind(&reqObj); err != nil {
			c.JSON(http.StatusOK, types.RPCParseError("", err))
			return
		}

		if err := reqObj.Validator(); err != nil {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
			return
		}

		err := service.CreateNode(reqObj.Name, reqObj.BaseURL, reqObj.NodeType, reqObj.NodeVersion, reqObj.SecretKey)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", nil))
	}
}

type listNodesResp struct {
	Nodes []*service.Node `json:"nodes"`
}

func listNodesGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		nts, err := service.AllNodes()
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}
		fmt.Printf("---req:%+v\n", c.Request)
		for _, v := range nts {
			(*v).BaseURL = "http://" + c.Request.Host + "/nodes/" + v.Name
		}

		res := listNodesResp{
			Nodes: nts,
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", res))
	}
}
