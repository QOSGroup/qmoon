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
	"github.com/QOSGroup/qmoon/types"
	qos_types "github.com/QOSGroup/qos/module/gov/types"
	"github.com/gin-gonic/gin"
)

const (
	proposalsUrl = "/proposals"
	proposalUrl  = "/proposal/:pid"
	votesUrl     = "/votes/:pid"
	depositsUrl  = "/deposits/:pid"
	tallyUrl     = "/tally/:pid"

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

func queryProposal(context *gin.Context) (proposal types.ResultProposal, err error) {
	pId, err := strconv.ParseInt(context.Param("pid"), 10, 64)
	if err != nil {
		err = errors.New("Can't find proposal id")
		return
	}

	k := fmt.Sprintf(proposalCacheKey, pId)
	if v, ok := cache.Get(k); ok {
		if proposal, ok = v.(types.ResultProposal); ok {
			return
		}
	}

	node, err := GetNodeFromUrl(context)
	if err != nil {
		return
	}

	proposal, err = qos.NewQosCli("").QueryProposal(node.BaseURL, pId)
	if err == nil {
		proposal.Deposits, err = queryDeposits(context)
		proposal.Votes, err = queryVotes(context)
	}
	if err == nil{
		cache.Set(k, proposal, time.Minute*5)
	}
	return
}

func queryVotes(context *gin.Context) ([]*types.ResultVote, error) {
	pId, err := strconv.ParseInt(context.Param("pid"), 10, 64)
	if err != nil {
		err = errors.New("Can't find proposal id")
		return nil, err
	}
	//k := fmt.Sprintf(votesCacheKey, pId)
	//if v, ok := cache.Get(k); ok {
	//	if votes, ok = v.(qos_types.Votes); ok {
	//		return
	//	}
	//}

	node, err := GetNodeFromUrl(context)
	if err != nil {
		return nil, err
	}
	votes, err := qos.NewQosCli("").QueryVotes(node.BaseURL, pId)
	result := make([]*types.ResultVote, 0)
	for _, d := range(votes) {
		option := "Empty"
		switch (d.Option){
		case qos_types.OptionEmpty:
			option = "Empty"
			break
		case qos_types.OptionAbstain:
			option = "Abstain"
			break
		case qos_types.OptionNo:
			option = "No"
			break
		case qos_types.OptionNoWithVeto:
			option = "No With Veto"
			break
		case qos_types.OptionYes:
			option = "Yes"
			break
		}
		vote := types.ResultVote{
			Voter: d.Voter.String(),
			Option: option,
		}
		result = append(result, &vote)
	}
	//if err == nil{
	//	cache.Set(k, votes, time.Minute*5)
	//}
	return result, nil
}

func queryDeposits(context *gin.Context) ([]*types.ResultDeposit, error) {
	pId, err := strconv.ParseInt(context.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}
	//
	//k := fmt.Sprintf(depositsCacheKey, pId)
	//if v, ok := cache.Get(k); ok {
	//	if deposits, ok = v.(qos_types.Deposits); ok {
	//		return
	//	}
	//}
	//
	node, err := GetNodeFromUrl(context)
	if err != nil {
		return nil, err
	}
	deposits, err := qos.NewQosCli("").QueryDeposits(node.BaseURL, pId)
	result := make([]*types.ResultDeposit, 0)
	for _, d := range(deposits) {
		deposit := types.ResultDeposit{
			Depositor: d.Depositor.String(),
			Amount: strconv.FormatInt(d.Amount.Int64(), 10),
		}
		result = append(result, &deposit)
	}
	//if err == nil{
	//	cache.Set(k, deposits, time.Minute*5)
	//}
	return result, nil
}

func queryTally(context *gin.Context) (tally qos_types.TallyResult, err error) {
	pId, err := strconv.ParseInt(context.Param("pid"), 10, 64)
	if err != nil {
		err = errors.New("Can't find proposal id")
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
