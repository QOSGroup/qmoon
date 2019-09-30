package hdata

import (
	"github.com/QOSGroup/qmoon/lib/qos"
	"github.com/QOSGroup/qmoon/models/errors"
	"net/http"
	"strconv"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/lib"
	qos_types "github.com/QOSGroup/qmoon/lib/qos/gov/types"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const (
	proposalsUrl = "/proposals"
	proposalUrl  = "/proposal"
	votesUrl     = "/votes"
	depositsUrl  = "/deposits"
	tallyUrl     = "/tally"
)

func init() {
	hdataHander[proposalsUrl] = proposalsGinRegister
	hdataHander[proposalUrl] = func(r *gin.Engine) {
		r.GET(NodeProxy+proposalUrl, middleware.ApiAuthGin(), func(context *gin.Context) {
			proposal, err := proposal(context)
			if err != nil {
				context.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
				return
			}
			context.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", proposal))
		})
	}
	hdataHander[votesUrl] = func(r *gin.Engine) {
		r.GET(NodeProxy+votesUrl, middleware.ApiAuthGin(), func(context *gin.Context) {
			votes, err := votes(context)
			if err != nil {
				context.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
				return
			}
			context.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", votes))
		})
	}
	hdataHander[depositsUrl] = func(r *gin.Engine) {
		r.GET(NodeProxy+depositsUrl, middleware.ApiAuthGin(), func(context *gin.Context) {
			deposits, err := deposits(context)
			if err != nil {
				context.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
				return
			}
			context.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", deposits))
		})
	}
	hdataHander[tallyUrl] = func(r *gin.Engine) {
		r.GET(NodeProxy+tallyUrl, middleware.ApiAuthGin(), func(context *gin.Context) {
			tally, err := tally(context)
			if err != nil {
				context.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
				return
			}
			context.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", tally))
		})
	}
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

func proposal(context *gin.Context) (proposal qos_types.Proposal, err error) {
	node, err := GetNodeFromUrl(context)
	pId, err := strconv.ParseInt(context.Query("pId"), 10, 64)
	if err != nil {
		err = errors.New("pid is error")
		return
	}
	proposal, err = qos.NewQosCli("").QueryProposal(node.BaseURL, pId)
	return
}

func votes(context *gin.Context) (votes qos_types.Votes, err error) {
	node, err := GetNodeFromUrl(context)
	if err != nil {
		return
	}
	pId, err := strconv.ParseInt(context.Query("pId"), 10, 64)
	if err != nil {
		err = errors.New("pid is error")
		return
	}
	votes, err = qos.NewQosCli("").QueryVotes(node.BaseURL, pId)
	return
}

func deposits(context *gin.Context) (deposits qos_types.Deposits, err error) {
	node, err := GetNodeFromUrl(context)
	if err != nil {
		return
	}
	pId, err := strconv.ParseInt(context.Query("pId"), 10, 64)
	if err != nil {
		err = errors.New("pid is error")
		return
	}
	deposits, err = qos.NewQosCli("").QueryDeposits(node.BaseURL, pId)
	return
}

func tally(context *gin.Context) (tally qos_types.TallyResult, err error) {
	node, err := GetNodeFromUrl(context)
	if err != nil {
		return
	}
	pId, err := strconv.ParseInt(context.Query("pId"), 10, 64)
	if err != nil {
		err = errors.New("pid is error")
		return
	}
	tally, err = qos.NewQosCli("").QueryTally(node.BaseURL, pId)
	return
}
