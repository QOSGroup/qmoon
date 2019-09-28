package hdata

import (
	"net/http"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const proposalsUrl = "/proposals"

func init() {
	hdataHander[proposalsUrl] = proposalsGinRegister
}

// proposalsGinRegister 注册proposals
func proposalsGinRegister(r *gin.Engine) {
	r.GET(NodeProxy+proposalsUrl, middleware.ApiAuthGin(), proposalsGin())
}

func proposalsGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		node, err := GetNodeFromUrl(c)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}

		vs, err := node.Proposals()
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
		}

		c.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", vs))
	}
}
