// Copyright 2018 The QOS Authors

package hdata

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/QOSGroup/qmoon/handler/middleware"
	"github.com/QOSGroup/qmoon/service"
	"github.com/QOSGroup/qmoon/types"
	"github.com/gin-gonic/gin"
)

const NodeProxy = "/nodes/:nodeName"

func init() {
	hdataHander[NodeProxy] = ProxyGinRegister
}

type ginHandler struct {
	Method  string
	Handler gin.HandlerFunc
}

const (
	tmAbciInfoUrl           = "/abci_info"
	tmConsensusStateUrl     = "/consensus_state"
	tmdumpConsensusStateUrl = "/dump_consensus_state"
	tmGenesisUrl            = "/genesis"
	tmHealthUrl             = "/health"
	tmNetInfoUrl            = "/net_info"
	tmNumUnconfirmedTxsUrl  = "/num_unconfirmed_txs"
	tmStatusUrl             = "/status"
	tmAbciQueryUrl          = "/abci_query"
	tmHeightUrl             = "/block?height"
	tmBlockResultsUrl       = "/block_results"
	tmBlockchainUrl         = "/blockchain"
	tmBroadcastTxAsyncUrl   = "/broadcast_tx_async"
	tmBroadcastTxCommitUrl  = "/broadcast_tx_commit"
	tmBroadcastTxSyncUrl    = "/broadcast_tx_sync"
	tmCommitUrl             = "/commit"
	tmSubscribeUrl          = "/subscribe"
	tmTxUrl                 = "/tx"
	tmTxSearchUrl           = "/tx_search"
	tmUnconfirmedTxsUrl     = "/unconfirmed_txs"
	tmUnsubscribeUrl        = "/unsubscribe"
	tmUnsubscribeAllUrl     = "/unsubscribe_all"
	tmValidatorsUrl         = "/validators"
)

var tmRouter = map[string]*ginHandler{
	"/abci_info":            nil,
	"/consensus_state":      nil,
	"/dump_consensus_state": nil,
	"/genesis":              nil,
	"/health":               nil,
	"/net_info":             nil,
	"/num_unconfirmed_txs":  nil,
	"/status":               nil,
	"/abci_query":           nil,
	"/block?height":         nil,
	"/block_results":        nil,
	"/blockchain":           nil,
	"/broadcast_tx_async":   nil,
	"/broadcast_tx_commit":  nil,
	"/broadcast_tx_sync":    nil,
	"/commit":               nil,
	"/subscribe":            nil,
	"/tx":                   nil,
	"/tx_search":            nil,
	"/unconfirmed_txs":      nil,
	"/unsubscribe":          nil,
	"/unsubscribe_all":      nil,
	"/validators":           nil,
}

const tendermintPrefix = "/tendermint"

// ProxyGinRegister 代理handler
func ProxyGinRegister(r *gin.Engine) {
	for k, v := range tmRouter {
		u := NodeProxy + tendermintPrefix + k
		if v == nil {
			r.GET(u, middleware.ApiAuthGin(), proxyGin())
		} else {
			if v.Handler == nil {
				r.Handle(v.Method, u, middleware.ApiAuthGin(), proxyGin())
			} else {
				r.Handle(v.Method, u, middleware.ApiAuthGin(), v.Handler)
			}
		}
	}
}

func copyHeaders(dst, src http.Header, keepDestHeaders bool) {
	if !keepDestHeaders {
		for k := range dst {
			dst.Del(k)
		}
	}
	for k, vs := range src {
		for _, v := range vs {
			dst.Add(k, v)
		}
	}
}

func proxyGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		nodeName := c.Param("nodeName")
		if nodeName == "" {
			c.JSON(http.StatusOK, types.RPCMethodNotFoundError(""))
			return
		}
		node, err := service.GetNodeByName(nodeName)
		if err != nil {
			c.JSON(http.StatusOK, types.RPCInvalidParamsError("", errors.New("nodeName not found")))
			return
		}
		us := strings.Split(c.Request.URL.String(), "/")
		log.Printf("us:%+v, l:%d", us, len(us))
		log.Printf("u:%s", node.BaseURL+"/"+strings.Join(us[4:], "/"))
		resp, err := http.Get(node.BaseURL + "/" + strings.Join(us[4:], "/"))
		if err != nil {
			c.JSON(http.StatusOK, types.RPCServerError("", err))
			return
			defer resp.Body.Close()
		}

		c.Writer.WriteHeader(resp.StatusCode)
		io.Copy(c.Writer, resp.Body)
	}
}
