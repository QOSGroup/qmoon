package hdata

import (
	"fmt"
	"github.com/QOSGroup/qmoon/cache"
	"github.com/QOSGroup/qmoon/lib/qos"
	"github.com/QOSGroup/qmoon/models/errors"
	"net/http"
	"strconv"
	"time"

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

	proposalCacheKey = "qos_proposal_key_%d"
	votesCacheKey    = "qos_votes_key_%d"
	depositsCacheKey = "qos_deposits_key_%d"
	tallyCacheKey    = "qos_tally_key_%d"
)

func init() {
	hdataHander[proposalsUrl] = proposalsGinRegister
	hdataHander[proposalUrl] = func(r *gin.Engine) {
		r.GET(NodeProxy+proposalUrl, middleware.ApiAuthGin(), func(context *gin.Context) {
			proposal, err := queryProposal(context)
			if err != nil {
				context.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
				return
			}
			context.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", proposal))
		})
	}
	hdataHander[votesUrl] = func(r *gin.Engine) {
		r.GET(NodeProxy+votesUrl, middleware.ApiAuthGin(), func(context *gin.Context) {
			votes, err := queryVotes(context)
			if err != nil {
				context.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
				return
			}
			context.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", votes))
		})
	}
	hdataHander[depositsUrl] = func(r *gin.Engine) {
		r.GET(NodeProxy+depositsUrl, middleware.ApiAuthGin(), func(context *gin.Context) {
			deposits, err := queryDeposits(context)
			if err != nil {
				context.JSON(http.StatusOK, types.RPCInvalidParamsError("", err))
				return
			}
			context.JSON(http.StatusOK, types.NewRPCSuccessResponse(lib.Cdc, "", deposits))
		})
	}
	hdataHander[tallyUrl] = func(r *gin.Engine) {
		r.GET(NodeProxy+tallyUrl, middleware.ApiAuthGin(), func(context *gin.Context) {
			tally, err := queryTally(context)
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

func queryProposal(context *gin.Context) (proposal qos_types.Proposal, err error) {
	pId, err := strconv.ParseInt(context.Query("pId"), 10, 64)
	if err != nil {
		err = errors.New("pid is error")
		return
	}

	k := fmt.Sprintf(proposalCacheKey, pId)
	if v, ok := cache.Get(k); ok {
		if proposal, ok = v.(qos_types.Proposal); ok {
			return
		}
	}

	node, err := GetNodeFromUrl(context)
	if err != nil {
		return
	}

	proposal, err = qos.NewQosCli("").QueryProposal(node.BaseURL, pId)
	if err == nil{
		cache.Set(k, proposal, time.Minute*5)
	}
	return
}

func queryVotes(context *gin.Context) (votes qos_types.Votes, err error) {
	pId, err := strconv.ParseInt(context.Query("pId"), 10, 64)
	if err != nil {
		err = errors.New("pid is error")
		return
	}

	k := fmt.Sprintf(votesCacheKey, pId)
	if v, ok := cache.Get(k); ok {
		if votes, ok = v.(qos_types.Votes); ok {
			return
		}
	}

	node, err := GetNodeFromUrl(context)
	if err != nil {
		return
	}
	votes, err = qos.NewQosCli("").QueryVotes(node.BaseURL, pId)
	if err == nil{
		cache.Set(k, votes, time.Minute*5)
	}
	return
}

func queryDeposits(context *gin.Context) (deposits qos_types.Deposits, err error) {
	pId, err := strconv.ParseInt(context.Query("pId"), 10, 64)
	if err != nil {
		err = errors.New("pid is error")
		return
	}

	k := fmt.Sprintf(depositsCacheKey, pId)
	if v, ok := cache.Get(k); ok {
		if deposits, ok = v.(qos_types.Deposits); ok {
			return
		}
	}

	node, err := GetNodeFromUrl(context)
	if err != nil {
		return
	}
	deposits, err = qos.NewQosCli("").QueryDeposits(node.BaseURL, pId)
	if err == nil{
		cache.Set(k, deposits, time.Minute*5)
	}
	return
}

func queryTally(context *gin.Context) (tally qos_types.TallyResult, err error) {
	pId, err := strconv.ParseInt(context.Query("pId"), 10, 64)
	if err != nil {
		err = errors.New("pid is error")
		return
	}

	k := fmt.Sprintf(tallyCacheKey, pId)
	if v, ok := cache.Get(k); ok {
		if tally, ok = v.(qos_types.TallyResult); ok {
			return
		}
	}

	node, err := GetNodeFromUrl(context)
	if err != nil {
		return
	}
	tally, err = qos.NewQosCli("").QueryTally(node.BaseURL, pId)
	if err == nil{
		cache.Set(k, tally, time.Minute*5)
	}
	return
}
